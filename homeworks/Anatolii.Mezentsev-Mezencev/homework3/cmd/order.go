package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var order string

var orderCmd = &cobra.Command{
	Use: "my_flags",
}

func Execute() {
	if err := orderCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	orderCmd.PersistentFlags().StringVar(&order, "order", "1 2 1 3 4", "order status.")
}
