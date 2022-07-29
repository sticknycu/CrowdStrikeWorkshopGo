package gateway

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

const fileName = "/tmp/dat"
const myFileName = "/tmp/my_dat"

type FileStorage struct {
}

type MyFileStorage struct {
}

func NewFileRepo() *FileStorage {
	return &FileStorage{}
}

func NewMyFileRepo() *MyFileStorage {
	return &MyFileStorage{}
}

func (f *MyFileStorage) GetMyContent(id string) (string, error) {
	return "", nil
}

func (f *FileStorage) GetContent(id string) (string, error) {
	return "", nil
}

func (f *MyFileStorage) WriteMyContent(id string, content string) error {
	// citesc din fisier
	data, err := os.ReadFile(myFileName)

	// in cazul in care nu merge sa citesc din fisier, eroarea nu e nil pointer (null lul)
	if err != nil {
		log.WithError(err).Errorf("Failed to read my file")
		return err
	}

	// cream frumi un map aici cu datele
	var myJsonData map[string]string
	// facem unmarshel, deci citim ce avem si il bagam frumi in map-ul nostru
	err = json.Unmarshal(data, &myJsonData)
	// hopa eroare
	if err != nil {
		log.WithError(err).Errorf("Failed to unmarshal my file")
		return err
	}

	// si luam contentul de il avem frumi aici si il stocam in json-ul nostru cu ce mai avem
	myJsonData[id] = content

	// si facem marshal
	myFileContent, err := json.Marshal(myJsonData)
	// hopa iar eroare
	if err != nil {
		log.WithError(err).Errorf("Failed to marshal my file")
		return err
	}
	// scriem in fisier
	err = os.WriteFile(myFileName, myFileContent, 777)
	// hopa iar eroare
	if err != nil {
		log.WithError(err).Errorf("Failed to write to my file")
		return err
	}

	// gata toate cazurile bos
	return nil
}

func (f *FileStorage) WriteContent(id string, content string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.WithError(err).Errorf("Failed to read file")
		return err
	}
	var jsonData map[string]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.WithError(err).Errorf("Failed to unmashal file")
		return err
	}
	jsonData[id] = content
	fileContent, err := json.Marshal(jsonData)
	if err != nil {
		log.WithError(err).Errorf("Failed to marshal file")
		return err
	}
	err = os.WriteFile(fileName, fileContent, 777)
	if err != nil {
		log.WithError(err).Errorf("Failed to write file")
		return err
	}
	return nil
}
