package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

var (
	NotOwnedByUsError = errors.New("Not owned by us")
)

func GetAuthConfigFilePath() string {
	return viper.ConfigFileUsed()
}

func SaveDataToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0666)
}

func ReadDataFromFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func GetAccessToken() (string, error) {
	data, err := ReadDataFromFile(GetAuthConfigFilePath())
	if err != nil {
		return "", err
	}

	jsonData := map[string]interface{}{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}

	accessToken, ok := jsonData["access_token"].(string)
	if !ok {
		return "", errors.New("Failed to type cast interface into string")
	}
	return string(accessToken), nil
}

func ToMapInterface(data []byte) (map[string]interface{}, error) {
	var mapInterface map[string]interface{}
	err := json.Unmarshal(data, &mapInterface)
	if err != nil {
		return mapInterface, err
	}
	return mapInterface, nil
}

func Owner() string {
	return "InnovoltPasswordManager"
}
