package tfinfra

import (
	"do-tf-stack-builder/tfutils"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Infra struct {
	Database     []Database     `json:"database_cluster,omitempty"`
	LoadBalancer []LoadBalancer `json:"load_balancer,omitempty"`
	Spaces       []Spaces       `json:"spaces,omitempty"`
}

func CreateInfraFiles(w http.ResponseWriter, infra Infra, p tfutils.Project, rootEnvBody *hclwrite.Body, tfEnvMain *hclwrite.File) {

	layer := strings.ToLower(reflect.TypeOf(infra).Name())

	if infra.Database != nil {
		for i := range infra.Database {
			modBody, tfMod, tfModFile := tfutils.InitModLayerFile(layer+"/"+strings.Replace(*infra.Database[i].Name, "_", "-", -1), strings.Replace(*infra.Database[i].Name, "_", "-", -1))
			tfutils.InitModule(w, modBody, infra.Database[i], "digitalocean_database_cluster", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *infra.Database[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, infra.Database[i], moduleBlockBody)
			}
			tfutils.EndModule(tfMod, tfModFile, p.Size, layer, *infra.Database[i].Name)
		}
	}

	if infra.LoadBalancer != nil {
		for i := range infra.LoadBalancer {
			modBody, tfMod, tfModFile := tfutils.InitModLayerFile(layer+"/"+strings.Replace(*infra.LoadBalancer[i].Name, "_", "-", -1), strings.Replace(*infra.LoadBalancer[i].Name, "_", "-", -1))
			tfutils.InitModule(w, modBody, infra.LoadBalancer[i], "digitalocean_loadbalancer", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *infra.LoadBalancer[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, infra.LoadBalancer[i], moduleBlockBody)
			}
			tfutils.EndModule(tfMod, tfModFile, p.Size, layer, *infra.LoadBalancer[i].Name)
		}
	}

	if infra.Spaces != nil {
		for i := range infra.Spaces {
			modBody, tfMod, tfModFile := tfutils.InitModLayerFile(layer+"/"+strings.Replace(*infra.Spaces[i].Name, "_", "-", -1), strings.Replace(*infra.Spaces[i].Name, "_", "-", -1))
			tfutils.InitModule(w, modBody, infra.Spaces[i], "digitalocean_spaces_bucket", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *infra.Spaces[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, infra.Spaces[i], moduleBlockBody)
			}
			tfutils.EndModule(tfMod, tfModFile, p.Size, layer, *infra.Spaces[i].Name)
		}
	}
}
