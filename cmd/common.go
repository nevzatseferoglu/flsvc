/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type RemoteHostPathType int

const (
	root RemoteHostPathType = iota
	ipAddr
	contactInfo
	flIdentifier
	ping
)

func (p RemoteHostPathType) String() string {
	return [...]string{"remote-hosts", "", "contact-info", "fl-identifier", "ping"}[p]
}

type RemoteHostCommand struct {
	IpAddr       net.IP
	ContackInfo  string
	FLIdentifier string
}

func remoteHostGetWithGivenPath(path RemoteHostPathType, m map[RemoteHostPathType]string) ([]byte, error) {
	url := ""
	switch path {
	case root:
		url = fmt.Sprintf("%s/%s", rootURL.String(), path.String())
	case ipAddr:
		url = fmt.Sprintf("%s/%s/%s", rootURL.String(), root.String(), m[ipAddr])
	case contactInfo:
		url = fmt.Sprintf("%s/%s/%s/%s", rootURL.String(), root.String(), path.String(), m[contactInfo])
	case flIdentifier:
		url = fmt.Sprintf("%s/%s/%s/%s", rootURL.String(), root.String(), path.String(), m[flIdentifier])
	case ping:
		url = fmt.Sprintf("%s/%s/%s", rootURL.String(), ping.String(), m[ping])
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Request is unsuccessful, err: %w\n", err)
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll is unsuccessful, err: %w\n", err)
	}
	return buf, nil
}

func indentJSONWithByteArray(jsonData []byte, path RemoteHostPathType) (string, error) {
	var result []byte
	switch path {
	case ipAddr, ping:
		var data map[string]interface{}
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			return "", err
		}
		result, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", err
		}
	case contactInfo, flIdentifier, root:
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
