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

const remoteHostRouter = "remote_hosts"

var ip_addr net.IP

var remoteHostCmd = &cobra.Command{
	Use:   "remote-hosts",
	Short: "Remote host operations",
	RunE: func(cmd *cobra.Command, args []string) error {
		// query according to ip address
		if ip_addr != nil {
			newURL := rootURL
			newURL.Path = fmt.Sprintf("%s/%s", remoteHostRouter, ip_addr)
			resp, err := client.Get(newURL.String())
			if err != nil {
				return fmt.Errorf("Request is unsuccessful, err: %w\n", err)
			}
			defer resp.Body.Close()
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("ReadAll is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		}
		return nil
	},
}

func init() {
	remoteHostCmd.Flags().IPVar(&ip_addr, "ip-addr", net.ParseIP("127.0.0.1"), "IP address of the remote host")
	_ = remoteHostCmd.MarkFlagRequired("ip-addr")
}
