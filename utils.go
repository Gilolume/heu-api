package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

//Renvoie le timestamp actuelle
func getTimeStamp() string {
	t := time.Now()
	return t.Format("20060102150405")
}

//Decode le body et met les champs dans la structure passer en parametre
func decodeBody(req *http.Request, s interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(s)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
}

//Copie un fichier
func copyFile(src string, dest string) bool {
	cp := exec.Command("cp", src, dest)
	err := cp.Run()
	if err != nil {
		log.Fatal(err)
	}
	return true
}

//Return true si fichier ou dossier existe
func ExistsFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//Creation d'un fichier
//Prend en parametre :
//Un dossier pour verifier son existe et le créer si ce n'est pas le cas
//Et le nom du fichier à créer
func CreateFile(folderPath string, filename string) {
	exists, err := ExistsFile(folderPath)
	if err != nil {
		log.Println("SaveFile:", err)
	}
	if exists == false {
		os.MkdirAll(folderPath, os.ModePerm)
	}
	err = ioutil.WriteFile(folderPath+filename, nil, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//Supprimer un fichier
func DeleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}
