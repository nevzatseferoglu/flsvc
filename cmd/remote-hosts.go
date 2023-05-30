/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/cobra"
)

type RemoteHostPathType int

const (
	root RemoteHostPathType = iota
	ipAddr
	contactInfo
)

func (p RemoteHostPathType) String() string {
	return [...]string{"remote-hosts", "", "contact-info"}[p]
}

type RemoteHostCommand struct {
	IpAddr      net.IP
	ContackInfo string
}

var remoteHostData RemoteHostCommand

var remoteHostCmd = &cobra.Command{
	Use:   "remote-hosts",
	Short: "Remote host operations",
	RunE: func(cmd *cobra.Command, args []string) error {
		mvalue := make(map[RemoteHostPathType]string, 0)
		switch {
		case remoteHostData.IpAddr != nil:
			mvalue[ipAddr] = remoteHostData.IpAddr.String()
			buf, err := remoteHostGetWithGivenPath(ipAddr, mvalue)
			if err != nil {
				return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, ipAddr)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		case len(remoteHostData.ContackInfo) != 0:
			mvalue[contactInfo] = remoteHostData.ContackInfo
			buf, err := remoteHostGetWithGivenPath(contactInfo, mvalue)
			if err != nil {
				return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, contactInfo)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		}
		return nil
	},
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

func init() {
	remoteHostCmd.Flags().IPVar(&remoteHostData.IpAddr, "ip-addr", nil, "IP address of the remote host")
	remoteHostCmd.Flags().StringVar(&remoteHostData.ContackInfo, "contact-info", "", "Contact info of the remote hosts")
}
