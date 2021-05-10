/*
Copyright © 2021 DeltaX 13975001197@163.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/1000Delta/wifi-locate/pkg/gateway"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run gateway service",
	Long: `
Run gateway service, support wifi-locate request.`,
	Run: func(cmd *cobra.Command, args []string) {
	
		mode := viper.GetString("gin.mode")
		switch mode {
		case gin.DebugMode, gin.ReleaseMode:
			gin.SetMode(mode)
		default:
			cobra.CheckErr("config gin.mode is not a valid value: [debug, release]")
		}

		// init server
		e := gin.Default()

		gateway.RegisterRouter(e)

		port := viper.GetInt("gin.port")

		e.Run(fmt.Sprintf(":%d", port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
