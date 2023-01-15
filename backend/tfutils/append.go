package tfutils

import (
	"log"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func AppendAttrIfNotNil(body *hclwrite.Body, attr, size string, typ interface{}) {

	switch reflect.Indirect(reflect.ValueOf(typ)).Kind() {
	case reflect.String, reflect.Int, reflect.Bool, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		if !reflect.ValueOf(typ).IsNil() {
			switch typ.(type) {
			case *string:
				body.SetAttributeValue(attr, cty.StringVal(*typ.(*string)))
			case *int:
				body.SetAttributeValue(attr, cty.NumberIntVal(int64(*typ.(*int))))
			case *bool:
				body.SetAttributeValue(attr, cty.BoolVal(*typ.(*bool)))
			case *[]string:
				val, err := gocty.ToCtyValue(*typ.(*[]string), cty.List(cty.String))
				if err != nil {
					log.Println(err)
				}
				body.SetAttributeValue(attr, val)
			case *map[string]string:
				val, err := gocty.ToCtyValue(*typ.(*map[string]string), cty.Map(cty.String))
				if err != nil {
					log.Println(err)
				}
				body.SetAttributeValue(attr, val)
			}
		}
	case reflect.Struct:
		if size == "small" {
			values := reflect.ValueOf(typ)
			typName := reflect.TypeOf(typ).Elem().Name()
			varBlock := body.AppendNewBlock(typName, []string{})
			varBody := varBlock.Body()
			for j := 0; j < values.Elem().NumField(); j++ {
				field := reflect.TypeOf(typ).Elem().Field(j).Tag
				AppendAttrIfNotNil(varBody, string(field.Get("json")), size, values.Elem().Field(j).Interface())
			}
		} else if size == "medium" {
			values := reflect.ValueOf(typ)
			typName := reflect.TypeOf(typ).Elem().Name()
			for j := 0; j < values.Elem().NumField(); j++ {
				field := reflect.TypeOf(typ).Elem().Field(j).Tag
				AppendAttrIfNotNil(body, typName+"_"+string(field.Get("json")), size, values.Elem().Field(j).Interface())
			}
		}
	}
}

func AppendAttrTravIfNotNil(body *hclwrite.Body, attr, vstr string, typ interface{}) {
	switch reflect.Indirect(reflect.ValueOf(typ)).Kind() {
	case reflect.String, reflect.Int, reflect.Bool, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		if !reflect.ValueOf(typ).IsNil() {
			body.SetAttributeTraversal(attr, hcl.Traversal{hcl.TraverseRoot{Name: "var"}, hcl.TraverseAttr{Name: vstr}})
		}
	case reflect.Struct:
		values := reflect.ValueOf(typ)
		typName := reflect.TypeOf(typ).Elem().Name()
		varBlock := body.AppendNewBlock(typName, []string{})
		varBody := varBlock.Body()
		for j := 0; j < values.Elem().NumField(); j++ {
			field := reflect.TypeOf(typ).Elem().Field(j).Tag
			AppendAttrTravIfNotNil(varBody, string(field.Get("json")), typName+"_"+string(field.Get("json")), values.Elem().Field(j).Interface())
		}
	}
}
