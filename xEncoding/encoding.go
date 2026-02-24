package xEncoding

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"os"
)

// FileToTarget 读取文件并解码到 target
func FileToTarget[T any](target *T, targetType string, filename string) error {
	// 读取文件到比特流
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// 比特流解码
	switch targetType {
	case "json":
		return json.Unmarshal(bytes, target)
	case "xml":
		return xml.Unmarshal(bytes, target)
	default:
		return errors.New("target type not supported: " + targetType)
	}
}

// URLToTarget 读取网址并解码到 target
func URLToTarget[T any](target *T, targetType string, url string) error {
	// Get
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	// body 解码
	switch targetType {
	case "json":
		return json.NewDecoder(response.Body).Decode(target)
	case "xml":
		return xml.NewDecoder(response.Body).Decode(target)
	default:
		return errors.New("target type not supported: " + targetType)
	}
}

// TargetToBytes target 编码到比特流
func TargetToBytes[T any](target *T, targetType string) ([]byte, error) {
	switch targetType {
	case "json":
		return json.Marshal(target)
	case "xml":
		return xml.Marshal(target)
	default:
		return nil, errors.New("target type not supported: " + targetType)
	}
}

// TargetToFile target 编码到文件
func TargetToFile[T any](target *T, targetType string, filename string) error {
	bytes, err := TargetToBytes(target, targetType)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, bytes, 0644)
}
