package tfutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func CreateFile(dirname, filename string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(dirname), 0755); err != nil {
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

	// rootBody.AppendNewline()

	// varBlock := rootBody.AppendNewBlock("variable", []string{"do_token"})
	// varBlockBody := varBlock.Body()
	// varBlockBody.AppendNewline()

	rootBody.AppendNewline()

	formattedContent := hclwrite.Format(tfProvider.Bytes())
	tfProviderFile.Write(formattedContent)

}

func InitModuleFile(layer, filename string) (*hclwrite.Body, *hclwrite.File, *os.File) {

	tfBlock := hclwrite.NewEmptyFile()
	tfBlockFile, _ := CreateFile("do-terraform/tf-modules/layer-"+layer+"/", filename)
	rootBody := tfBlock.Body()

	return rootBody, tfBlock, tfBlockFile
}

func EndModuleFile(tfMod *hclwrite.File, tfFile *os.File, size, layer, filename string) {

	formattedContent := hclwrite.Format(tfMod.Bytes())
	tfFile.Write(formattedContent)
	if size == "medium" {
		ProviderFile(layer, filename)
	}
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
