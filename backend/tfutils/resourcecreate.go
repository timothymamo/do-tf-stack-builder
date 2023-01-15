package tfutils

import (
	"do-tf-stack-builder/utils"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func InitModule(w http.ResponseWriter, rootBody *hclwrite.Body, typ interface{}, r, s, l string) {

	if err := utils.ValidateStruct(typ); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		values := reflect.ValueOf(typ)
		blockBody := InitResource(rootBody, r, strings.Replace(*values.FieldByName("Name").Interface().(*string), "-", "_", -1), int(values.FieldByName("Amount").Int()))
		if s == "small" {
			for i := 0; i < values.NumField(); i++ {
				field := reflect.TypeOf(typ).Field(i).Tag
				if string(field.Get("json")) != "amount" || string(field.Get("json")) != "output" {
					if string(field.Get("json")) != "name" {
						AppendAttrIfNotNil(blockBody, string(field.Get("json")), s, values.Field(i).Interface())
					} else if string(field.Get("json")) == "name" {
						if int(values.FieldByName("Amount").Int()) > 1 {
							userTokens := hclwrite.Tokens{
								{
									Type:  hclsyntax.TokenStringLit,
									Bytes: []byte(`"` + *values.Field(i).Interface().(*string) + `-${count.index}"`),
								},
							}
							blockBody.SetAttributeRaw("name", userTokens)
						} else {
							blockBody.SetAttributeValue(string(field.Get("json")), cty.StringVal(*values.Field(i).Interface().(*string)))
						}
					}
				}
			}
			rootBody.AppendNewline()

		} else if s == "medium" {
			values := reflect.ValueOf(typ)
			for i := 0; i < values.NumField(); i++ {
				field := reflect.TypeOf(typ).Field(i).Tag
				if string(field.Get("json")) != "amount" || string(field.Get("json")) != "output" {
					if string(field.Get("json")) != "name" {
						AppendAttrTravIfNotNil(blockBody, string(field.Get("json")), string(field.Get("json")), values.Field(i).Interface())
					} else if string(field.Get("json")) == "name" {
						if int(values.FieldByName("Amount").Int()) > 1 {
							blockBody.SetAttributeTraversal("name", hcl.Traversal{hcl.TraverseRoot{Name: "\"${var"}, hcl.TraverseAttr{Name: string(field.Get("json")) + "}-${count.index}\""}})
						} else {
							blockBody.SetAttributeTraversal("name", hcl.Traversal{hcl.TraverseRoot{Name: "var"}, hcl.TraverseAttr{Name: string(field.Get("json"))}})
						}
					}
				}
			}
			rootBody.AppendNewline()

			name := strings.Replace(*values.FieldByName("Name").Interface().(*string), "-", "_", -1)
			varBody, tfVar, tfVarFile := InitModuleFile(l+"/"+strings.Replace(name, "_", "-", -1), "variables-"+strings.Replace(name, "_", "-", -1))
			for i := 0; i < values.NumField(); i++ {
				field := reflect.TypeOf(typ).Field(i).Tag
				if string(field.Get("json")) != "amount" || string(field.Get("json")) != "output" {
					CreateVarIfNotNil(varBody, string(field.Get("json")), string(field.Get("description")), values.Field(i).Interface())
				}
			}
			formattedContent := hclwrite.Format(tfVar.Bytes())
			tfVarFile.Write(formattedContent)
		}
		rootBody.AppendNewline()
	}
}

func TFVarModule(s string, typ interface{}, moduleBlockBody *hclwrite.Body) {
	values := reflect.ValueOf(typ)
	for i := 1; i < values.NumField(); i++ {
		field := reflect.TypeOf(typ).Field(i).Tag
		AppendAttrIfNotNil(moduleBlockBody, string(field.Get("json")), s, values.Field(i).Interface())
	}
}
