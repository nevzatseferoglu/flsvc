package cmd

import "encoding/json"

func indentJSONWithByteArray(jsonData []byte, path RemoteHostPathType) (string, error) {
	var result []byte
	switch path {
	case ipAddr:
		var data map[string]interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return "", err
		}
		result, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", err
		}
	case contactInfo:
		var data []interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return "", err
		}
		result, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", err
		}
	}
	return string(result), nil
}
