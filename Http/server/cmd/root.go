package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog"
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


type ExtraConfig struct {
	BaseDomain           string
}

func Run(){
	flags := rootCmd.Flags()

	flags.String("name","yuzhipeng","the name of author")
	flags.String("port","3000","the name of author")


	flags.String("basedomain", "alauda.io",
		"Used to specify the default system namespace (formerly alauda-system).")


	_ = viper.BindPFlag("name", rootCmd.Flags().Lookup("name"))
	_ = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag("basedomain", rootCmd.Flags().Lookup("basedomain"))


	basedomain := viper.GetString("basedomain")
	klog.Errorf("basedomain is %v", basedomain)


	if err := rootCmd.Execute(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
