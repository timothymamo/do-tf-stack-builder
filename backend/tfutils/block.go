package tfutils

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func InitResource(root *hclwrite.Body, r, n string, a int) *hclwrite.Body {
	block := root.AppendNewBlock("resource", []string{r, n})
	blockBody := block.Body()
	if a > 1 {
		blockBody.SetAttributeValue("count", cty.NumberIntVal(int64(a)))
		blockBody.AppendNewline()
	}

	return blockBody
}

func CreateVarIfNotNil(root *hclwrite.Body, v, d string, typ interface{}) {

	switch reflect.Indirect(reflect.ValueOf(typ)).Kind() {
	case reflect.String, reflect.Int, reflect.Bool, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		if !reflect.ValueOf(typ).IsNil() {
			ctyType, _ := gocty.ImpliedType(typ)
			if ctyType != cty.NilType {
				varBlock := root.AppendNewBlock("variable", []string{v})
				varBody := varBlock.Body()
				varBody.SetAttributeValue("description", cty.StringVal(d))
				varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
				root.AppendNewline()
			}
		}
	case reflect.Struct:
		values := reflect.ValueOf(typ)
		typName := reflect.TypeOf(typ).Elem().Name()
		for j := 0; j < values.Elem().NumField(); j++ {
			field := reflect.TypeOf(typ).Elem().Field(j).Tag
			CreateVarIfNotNil(root, typName+"_"+string(field.Get("json")), string(field.Get("description")), values.Elem().Field(j).Interface())
		}
	}
}

func typeExprTokens(ty cty.Type) hclwrite.Tokens {
	switch ty {
	case cty.String:
		return hclwrite.TokensForIdentifier("string")
	case cty.Bool:
		return hclwrite.TokensForIdentifier("bool")
	case cty.Number:
		return hclwrite.TokensForIdentifier("number")
	case cty.DynamicPseudoType:
		return hclwrite.TokensForIdentifier("any")
	}

	if ty.IsCollectionType() {
		etyTokens := typeExprTokens(ty.ElementType())
		switch {
		case ty.IsListType():
			return hclwrite.TokensForFunctionCall("list", etyTokens)
		case ty.IsSetType():
			return hclwrite.TokensForFunctionCall("set", etyTokens)
		case ty.IsMapType():
			return hclwrite.TokensForFunctionCall("map", etyTokens)
		default:
			// Should never happen because the above is exhaustive
			panic("unsupported collection type")
		}
	}

	if ty.IsObjectType() {
		atys := ty.AttributeTypes()
		names := make([]string, 0, len(atys))
		for name := range atys {
			names = append(names, name)
		}
		sort.Strings(names)

		items := make([]hclwrite.ObjectAttrTokens, len(names))
		for i, name := range names {
			items[i] = hclwrite.ObjectAttrTokens{
				Name:  hclwrite.TokensForIdentifier(name),
				Value: typeExprTokens(atys[name]),
			}
		}

		return hclwrite.TokensForObject(items)
	}

	if ty.IsTupleType() {
		etys := ty.TupleElementTypes()
		items := make([]hclwrite.Tokens, len(etys))
		for i, ety := range etys {
			items[i] = typeExprTokens(ety)
		}
		return hclwrite.TokensForTuple(items)
	}

	panic(fmt.Errorf("unsupported type %#v", ty))
}
