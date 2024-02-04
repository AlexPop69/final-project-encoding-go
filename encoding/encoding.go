package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v2"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	fileData, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Println("Ошибка чтения файла", j.FileInput)
	}

	if err := json.Unmarshal(fileData, &j.DockerCompose); err != nil {
		fmt.Println("Ошибка перекодировки из JSON в YAML")
		return err
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в yaml: %s", err.Error())
		return err
	}

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Println("Ошибка создания YAML файла:", err)
		return err
	}
	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		fmt.Println("Ошибка записи YAML файла:", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	fileData, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Println("Ошибка чтения файла", y.FileInput)
	}

	if err := yaml.Unmarshal(fileData, &y.DockerCompose); err != nil {
		fmt.Println("Ошибка перекодировки из YAML в JSON")
		return err
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s", err.Error())
		return err
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Println("Ошибка создания JSON файла:", err)
		return err
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи JSON файла:", err)
	}

	return nil
}
