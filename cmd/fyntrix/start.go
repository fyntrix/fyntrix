package main

import (
	"fmt"

	"github.com/fyntrix/fyntrix/config"
	"github.com/fyntrix/fyntrix/image/vips"
	"github.com/fyntrix/fyntrix/pkg/logger"
	"github.com/spf13/cobra"
)

func startCMD(root *cobra.Command) *cobra.Command {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start fyntrix server",
	}

	root.AddCommand(startCmd)

	startCmd.RunE = func(_ *cobra.Command, _ []string) error {
		cfg, err := config.Load(configPath)
		if err != nil {
			return err
		}

		logger.InitGlobalLogger(cfg.Logger)

		vip := vips.New(cfg.Image.Vips)

		if err := vip.Init(); err != nil {
			return err
		}

		defer func() {
			_ = vip.Close()
		}()

		fmt.Printf("vips initialized: %v\n", vip.Inited())
		fmt.Println(vip.Version())

		return nil
	}

	return startCmd
}
