package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"webservice/Http/server/customhandler"
)

var rootCmd = &cobra.Command{
	Use: "server",
	Short:"http server",
	Run: func(cmd *cobra.Command, args []string) {
		customhandler.StartServer()
	},
}

func Execute(){

	rootCmd.Flags().String("name","yuzhipeng","the name of author")
	rootCmd.Flags().String("port","3000","the name of author")

	_ = viper.BindPFlag("name", rootCmd.Flags().Lookup("name"))
	_ = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))

	if err := rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
