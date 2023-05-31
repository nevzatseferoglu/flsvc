/*
Copyright © 2023 Nevzat Seferoglu nevzatseferoglu@gmail.com
*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/nevzatseferoglu/fl-service-cli/config"
	"github.com/spf13/cobra"
)

const defaultTimeout = 5 * time.Second

var (
	rootURL url.URL
	cfgFile string
	client  *http.Client
	conf    config.Conf
)

const long = `This API caters to data scientists, simplifying remote host communication with service endpoints.It allows users to efficiently manage
flower federated learning clusters.

API doc: https://fl-service-api-doc.netlify.app/
`

var rootCmd = &cobra.Command{
	Use:   "flsvc",
	Short: "fl-service project command line client",
	Long:  long,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if len(cfgFile) == 0 {
			cfgFile, err = config.DefaultConfigurationFile()
			if err != nil {
				return fmt.Errorf("Cannot get default configuration file, %w\n", err)
			}
		}
		_, err = conf.ValidateConfigurationFile(cfgFile)
		if err != nil {
			return fmt.Errorf("Cannot validate configuration file, %w\n", err)
		}
		if conf.Timeout != time.Duration(0) {
			conf.Timeout = defaultTimeout
		}
		client = newHTTPClient(conf.Timeout)
		// set root url (http://localhost:8080)
		rootURL = url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%d", conf.Hostname, conf.Port),
		}
		return nil
	},
}

func newHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: time.Duration(timeout)}
}

func Execute() {
	rootCmd.AddCommand(
		remoteHostCmd,
		pingCmd,
		dockerCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flsvc.yaml)")
}
