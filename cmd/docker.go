/*
Copyright © 2023 Nevzat Seferoglu nevzatseferoglu@gmail.com
*/
package cmd

import (
	"fmt"
	"net"
	"io"

	"github.com/spf13/cobra"
)

var (
	dockerIPAddrData net.IP
	dockerDepStates bool
	dockerIns bool
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Install and get the states of the docker dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case dockerDepStates:
			buf, err := remoteHostGetWithGivenPath(docker, map[RemoteHostPathType]string{
				docker: dockerIPAddrData.String(),
			})
			if err != nil {
				return fmt.Errorf("remoteHostGetWithGivenPath is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, dockerStates)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		case dockerIns:
			url := fmt.Sprintf("%s/%s/%s/%s", rootURL.String(), docker.String(), dockerInstall.String(), dockerIPAddrData.String())
			resp, err := client.Post(url, "", nil)
			if err != nil {
				return fmt.Errorf("Request is unsuccessful, err: %w\n", err)
			}
			defer resp.Body.Close()
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("ReadAll is unsuccessful, err: %w\n", err)
			}
			str, err := indentJSONWithByteArray(buf, dockerInstall)
			if err != nil {
				return fmt.Errorf("indentJSON is unsuccessful, err: %w\n", err)
			}
			cmd.Println(str)
		}
		return nil
	},
}

func init() {
	dockerCmd.Flags().IPVar(&dockerIPAddrData, "ip-addr", nil, "IP address of the remote host")
	_ = dockerCmd.MarkFlagRequired("ip-addr")
	dockerCmd.Flags().BoolVar(&dockerDepStates, "states", false, "Get the states of the docker dependencies of the remote host")
	dockerCmd.Flags().BoolVar(&dockerIns, "install", false, "Install the docker to remote host (Ubuntu 20.04 (focal))")
}
