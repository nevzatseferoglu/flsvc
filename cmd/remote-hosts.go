/*
Copyright © 2023 Nevzat Seferoglu nevzatseferoglu@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var remoteHostData RemoteHostCommand

var remoteHostCmd = &cobra.Command{
	Use:   "remote-hosts",
	Short: "Remote host operations",
	RunE: func(cmd *cobra.Command, args []string) error {
		mvalue := make(map[RemoteHostPathType]string, 0)
		switch {
		// print properties of the host machine which has given IP address
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
		// print all remote hosts which are belonging to given contact information
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
		// print all remote hosts which are belonging to given contact information
		case len(remoteHostData.FLIdentifier) != 0:
			mvalue[flIdentifier] = remoteHostData.FLIdentifier
			buf, err := remoteHostGetWithGivenPath(flIdentifier, mvalue)
			if err != nil {
				return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, flIdentifier)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		// print all recorded remote hosts
		default:
			buf, err := remoteHostGetWithGivenPath(remoteHosts, nil)
			if err != nil {
				return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, remoteHosts)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		}
		return nil
	},
}

func init() {
	remoteHostCmd.Flags().IPVar(&remoteHostData.IpAddr, "ip-addr", nil, "IP address of the remote host")
	remoteHostCmd.Flags().StringVar(&remoteHostData.ContackInfo, "contact-info", "", "Contact info of the remote hosts")
	remoteHostCmd.Flags().StringVar(&remoteHostData.FLIdentifier, "fl-identifier", "", "Flower federeated learning cluster identifier")
}
