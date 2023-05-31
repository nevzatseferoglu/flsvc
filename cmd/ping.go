/*
Copyright © 2023 Nevzat Seferoglu nevzatseferoglu@gmail.com
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var pingIPAddrData net.IP

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping the remote host which has the given IPAddress",
	RunE: func(cmd *cobra.Command, args []string) error {
		buf, err := remoteHostGetWithGivenPath(ping, map[RemoteHostPathType]string{
			ping: pingIPAddrData.String(),
		})
		if err != nil {
			return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
		}
		str, err := indentJSONWithByteArray(buf, ping)
		if err != nil {
			return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
		}
		cmd.Println(str)
		return nil
	},
}

func init() {
	pingCmd.Flags().IPVar(&pingIPAddrData, "ip-addr", nil, "IP address of the remote host")
	_ = pingCmd.MarkFlagRequired("ip-addr")
}
