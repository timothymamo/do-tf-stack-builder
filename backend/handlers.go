package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"do-tf-stack-builder/tfbase"
	"do-tf-stack-builder/tfcompute"
	"do-tf-stack-builder/tfinfra"
	"do-tf-stack-builder/tfutils"
	"do-tf-stack-builder/utils"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

const JSON = "application/json"

type TfModules struct {
	Project *tfutils.Project   `json:"project"`
	Base    *tfbase.Base       `json:"base,omitempty"`
	Compute *tfcompute.Compute `json:"compute,omitempty"`
	Infra   *tfinfra.Infra     `json:"infra,omitempty"`
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func (app *App) CreateTFFiles(w http.ResponseWriter, r *http.Request) {

	var m TfModules
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	tfEnvMain := hclwrite.NewEmptyFile()
	rootEnvBody := tfEnvMain.Body()

	if m.Base != nil {
		tfbase.CreateBaseFiles(w, *m.Base, *m.Project, rootEnvBody, tfEnvMain)
	}
	if m.Compute != nil {
		tfcompute.CreateComputeFiles(w, *m.Compute, *m.Project, rootEnvBody, tfEnvMain)
	}
	if m.Infra != nil {
		tfinfra.CreateInfraFiles(w, *m.Infra, *m.Project, rootEnvBody, tfEnvMain)
	}
	for _, env := range m.Project.Envs {
		if m.Project.Size == "medium" {
			tfutils.EndEnvFile(env, tfEnvMain)
		}

		tfutils.BackendFile(env, m.Project.State)
	}

	if m.Project.Size == "small" {

		tfutils.ProviderFile("base", "small")

		listFiles("do-terraform")

		// tfoutput, err := ioutil.ReadFile("do-terraform.tf")
		// if err != nil {
		// 	fmt.Printf("Could not read the file due to this %s error \n", err)
		// }
		// utils.RespondWithJSON(w, http.StatusOK, string(tfoutput))
		sendTFFile(w, r)

		if err := os.RemoveAll("do-terraform.tf"); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else if m.Project.Size == "medium" {
		zipWriter()
		sendZipFile(w, r)
		if err := os.RemoveAll("do-terraform.zip"); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err := os.RemoveAll("do-terraform"); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

}

func listFiles(path string) {

	f, err := os.OpenFile("do-terraform.tf", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			subDirectoryPath := filepath.Join(path, file.Name())
			listFiles(subDirectoryPath)
		} else {
			b, _ := ioutil.ReadFile(path + "/" + file.Name())
			if _, err = f.Write(b); err != nil {
				panic(err)
			}
		}
	}
}

func sendTFFile(w http.ResponseWriter, r *http.Request) {

	downloadBytes, err := ioutil.ReadFile("do-terraform.tf")
	if err != nil {
		fmt.Println(err)
	}
	fileSize := len(string(downloadBytes))

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=do-terraform.tf")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, "do-terraform.tf", time.Now(), bytes.NewReader(downloadBytes))
}

func sendZipFile(w http.ResponseWriter, r *http.Request) {

	downloadBytes, err := ioutil.ReadFile("do-terraform.zip")
	if err != nil {
		fmt.Println(err)
	}
	fileSize := len(string(downloadBytes))

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=do-terraform.zip")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, "do-terraform.zip", time.Now(), bytes.NewReader(downloadBytes))
}

func zipWriter() {
	outFile, err := os.Create("do-terraform.zip")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk("do-terraform", walker)
	if err != nil {
		panic(err)
	}
}
