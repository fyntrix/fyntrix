package main

import (
	"github.com/fyntrix/fyntrix/version"
	"github.com/spf13/cobra"
)

var configPath string

func main() {
	rootCmd := &cobra.Command{
		Use:   "fyntrix",
		Short: "Fyntrix Service",
		Long: `
Fyntrix is an open-source service for real-time image and video transformation.
It lets you modify media on-the-fly using simple URL parameters or API requests.
Ideal for use in websites, apps, CDNs, or any system needing fast, dynamic media delivery.
`,
		Version:           version.Version.String(),
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	// Hide the "help" sub-command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().StringVar(&configPath, "config", "",
		"Path to TOML config file (required)")
	_ = rootCmd.MarkPersistentFlagRequired("config")

	// register commands
	startCMD(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
