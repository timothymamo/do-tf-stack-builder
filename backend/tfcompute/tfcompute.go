package tfcompute

import (
	"do-tf-stack-builder/tfutils"
	"net/http"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Compute struct {
	Droplets []Droplets `json:"droplet,omitempty"`
	K8s      []K8s      `json:"k8s,omitempty"`
}

func CreateComputeFiles(w http.ResponseWriter, compute Compute, p tfutils.Project, rootEnvBody *hclwrite.Body, tfEnvMain *hclwrite.File) {

	layer := strings.ToLower(reflect.TypeOf(compute).Name())

	if compute.Droplets != nil {
		for i := range compute.Droplets {
			modBody, tfMod, tfModFile := tfutils.InitModLayerFile(layer+"/"+strings.Replace(*compute.Droplets[i].Name, "_", "-", -1), strings.Replace(*compute.Droplets[i].Name, "_", "-", -1))
			tfutils.InitModule(w, modBody, compute.Droplets[i], "digitalocean_droplet", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *compute.Droplets[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, compute.Droplets[i], moduleBlockBody)
			}
			tfutils.EndModule(tfMod, tfModFile, p.Size, layer, *compute.Droplets[i].Name)
		}
	}

	if compute.K8s != nil {
		for i := range compute.K8s {
			k8sBody, tfK8s, tfK8sFile := tfutils.InitModLayerFile(layer+"/"+strings.Replace(*compute.K8s[i].Name, "_", "-", -1), strings.Replace(*compute.K8s[i].Name, "_", "-", -1))
			tfutils.InitModule(w, k8sBody, compute.K8s[i], "digitalocean_cluster", p.Size, layer)
			if p.Size == "medium" {
				moduleBlockBody := tfutils.InitEnvFile(p.Modules[1], *compute.K8s[i].Name, rootEnvBody)

				tfutils.TFVarModule(p.Size, compute.K8s[i], moduleBlockBody)
			}
			tfutils.EndModule(tfK8s, tfK8sFile, p.Size, layer, *compute.K8s[i].Name)
		}
	}
}
