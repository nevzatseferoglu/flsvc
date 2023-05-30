package cmd

import "encoding/json"

func indentJSONWithByteArray(jsonData []byte) (string, error) {
	var data map[string]interface{}
	// Unmarshal the JSON data into the map
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}
	indentedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(indentedJSON), nil
}
