package tfutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func CreateFile(dirname, filename string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(dirname), 0770); err != nil {
		fmt.Println(err)
	}

	tfFile, err := os.Create(dirname + filename + ".tf")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tfFile, nil

}

func ProviderFile(layer, filename string) {

	tfProvider := hclwrite.NewEmptyFile()
	tfProviderFile, _ := CreateFile("do-terraform/tf-modules/layer-"+layer+"/"+filename+"/", "provider")
	rootBody := tfProvider.Body()
	tfBlock := rootBody.AppendNewBlock("terraform", nil)
	tfBlockBody := tfBlock.Body()

	requiredProviderBlock := tfBlockBody.AppendNewBlock("required_providers", nil)
	requiredProviderBlockBody := requiredProviderBlock.Body()
	requiredProviderBlockBody.SetAttributeValue("digitalocean", cty.ObjectVal(map[string]cty.Value{
		"source":  cty.StringVal("digitalocean/digitalocean"),
		"version": cty.StringVal("~> 2.0"),
	}))

	rootBody.AppendNewline()

	providerBlock := rootBody.AppendNewBlock("provider", []string{"digitalocean"})
	providerBlockBody := providerBlock.Body()
	providerBlockBody.AppendNewline()
	// providerBlockBody.SetAttributeValue("token", cty.StringVal("var.do_token"))

	rootBody.AppendNewline()

	// varBlock := rootBody.AppendNewBlock("variable", []string{"do_token"})
	// varBlockBody := varBlock.Body()
	// varBlockBody.AppendNewline()

	formattedContent := hclwrite.Format(tfProvider.Bytes())
	tfProviderFile.Write(formattedContent)

}

func InitModLayerFile(layer, filename string) (*hclwrite.Body, *hclwrite.File, *os.File) {

	tfBlock := hclwrite.NewEmptyFile()
	tfBlockFile, _ := CreateFile("do-terraform/tf-modules/layer-"+layer+"/", filename)
	rootBody := tfBlock.Body()

	return rootBody, tfBlock, tfBlockFile
}

func EndModule(tfMod *hclwrite.File, tfFile *os.File, size, layer, filename string) {

	formattedContent := hclwrite.Format(tfMod.Bytes())
	tfFile.Write(formattedContent)
	if size == "medium" {
		ProviderFile(layer, filename)
	}
}

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

func CreateVar(root *hclwrite.Body, v, d string, ty interface{}) {

	ctyType, _ := gocty.ImpliedType(ty)

	varBlock := root.AppendNewBlock("variable", []string{v})
	varBody := varBlock.Body()
	varBody.SetAttributeValue("description", cty.StringVal(d))
	varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
	root.AppendNewline()
}

func CreateVarStrIfNotNil(root *hclwrite.Body, v, d string, st *string) {
	if st != nil {
		ctyType, _ := gocty.ImpliedType(st)

		varBlock := root.AppendNewBlock("variable", []string{v})
		varBody := varBlock.Body()
		varBody.SetAttributeValue("description", cty.StringVal(d))
		varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
		root.AppendNewline()
	}
}

func CreateVarBoolIfNotNil(root *hclwrite.Body, v, d string, bl *bool) {
	if bl != nil {
		ctyType, _ := gocty.ImpliedType(bl)

		varBlock := root.AppendNewBlock("variable", []string{v})
		varBody := varBlock.Body()
		varBody.SetAttributeValue("description", cty.StringVal(d))
		varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
		root.AppendNewline()
	}
}

func CreateVarIntIfNotNil(root *hclwrite.Body, v, d string, i *int) {
	if i != nil {
		ctyType, _ := gocty.ImpliedType(i)

		varBlock := root.AppendNewBlock("variable", []string{v})
		varBody := varBlock.Body()
		varBody.SetAttributeValue("description", cty.StringVal(d))
		varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
		root.AppendNewline()
	}
}

func CreateVarSliceIfNotNil(root *hclwrite.Body, v, d string, sl *[]string) {
	if sl != nil {
		ctyType, _ := gocty.ImpliedType(sl)

		varBlock := root.AppendNewBlock("variable", []string{v})
		varBody := varBlock.Body()
		varBody.SetAttributeValue("description", cty.StringVal(d))
		varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
		root.AppendNewline()
	}
}

func CreateVarMapIfNotNil(root *hclwrite.Body, v, d string, mp *map[string]string) {
	if mp != nil {
		ctyType, _ := gocty.ImpliedType(mp)

		varBlock := root.AppendNewBlock("variable", []string{v})
		varBody := varBlock.Body()
		varBody.SetAttributeValue("description", cty.StringVal(d))
		varBody.SetAttributeRaw("type", typeExprTokens(ctyType))
		root.AppendNewline()
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

func InitEnvFile(module, name string, rootEnvBody *hclwrite.Body) *hclwrite.Body {

	moduleBlock := rootEnvBody.AppendNewBlock("module", []string{module + "_" + strings.Replace(name, "-", "_", -1)})
	moduleBlockBody := moduleBlock.Body()
	moduleBlockBody.SetAttributeValue("source", cty.StringVal("../../tf-modules/layer-"+module+"/"+name+"/"))
	moduleBlockBody.AppendNewline()

	rootEnvBody.AppendNewline()

	outputBlock := rootEnvBody.AppendNewBlock("output", []string{"module_" + module + "_" + strings.Replace(name, "-", "_", -1)})
	outputBlockBody := outputBlock.Body()
	outputVal := hclwrite.Tokens{
		{Type: hclsyntax.TokenIdent, Bytes: []byte(`module.` + module + "_" + strings.Replace(name, "-", "_", -1))},
	}
	outputBlockBody.SetAttributeRaw("value", outputVal)

	rootEnvBody.AppendNewline()

	return moduleBlockBody
}

func EndEnvFile(env string, tfEnvMain *hclwrite.File) {

	formattedContent := hclwrite.Format(tfEnvMain.Bytes())

	tfEnvMainFile, _ := CreateFile("do-terraform/tf-envs/"+env+"/", "main")
	tfEnvMainFile.Write(formattedContent)

}

func BackendFile(env, state string) {

	tfEnvBackend := hclwrite.NewEmptyFile()
	rootBody := tfEnvBackend.Body()
	tfBlock := rootBody.AppendNewBlock("terraform", nil)
	tfBlockBody := tfBlock.Body()
	if state == "cloud" {
		cloudBlock := tfBlockBody.AppendNewBlock("cloud", nil)
		cloudBlockBody := cloudBlock.Body()
		cloudBlockBody.SetAttributeValue("organization", cty.StringVal("my-org"))
		cloudBlockBody.AppendNewline()
		workspaceBlock := cloudBlockBody.AppendNewBlock("workspaces", nil)
		workspaceBlockBody := workspaceBlock.Body()
		workspaceBlockBody.SetAttributeValue("tags", cty.ListVal([]cty.Value{cty.StringVal(env), cty.StringVal("test")}))
	} else if state == "local" {
		localBlock := tfBlockBody.AppendNewBlock("backend", []string{"local"})
		localBlock.Body()
	}
	rootBody.AppendNewline()
	formattedContent := hclwrite.Format(tfEnvBackend.Bytes())

	tfEnvBackendFile, _ := CreateFile("do-terraform/tf-envs/"+env+"/", "backend")
	tfEnvBackendFile.Write(formattedContent)

}
