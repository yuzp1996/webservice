package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "server",
	Short:"http server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute(){




	if err := rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}