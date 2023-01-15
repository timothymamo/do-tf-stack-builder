package tfbase

import (
	"do-tf-stack-builder/tfutils"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Base struct {
	Vpc []Vpc `json:"vpc"`
}

func CreateBaseFiles(w http.ResponseWriter, base Base, p tfutils.Project, rootEnvBody *hclwrite.Body, tfEnvMain *hclwrite.File) {

	layer := strings.ToLower(reflect.TypeOf(base).Name())

	if base.Vpc != nil {
		for i := range base.Vpc {
			modBody, tfMod, tfModFile := tfutils.InitModuleFile(layer+"/"+strings.Replace(*base.Vpc[i].Name, "_", "-", -1), strings.Replace(*base.Vpc[i].Name, "_", "-", -1))
			tfutils.InitModule(w, modBody, base.Vpc[i], "digitalocean_vpc", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *base.Vpc[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, base.Vpc[i], moduleBlockBody)
			}
			tfutils.EndModuleFile(tfMod, tfModFile, p.Size, layer, *base.Vpc[i].Name)
		}
	}
}
