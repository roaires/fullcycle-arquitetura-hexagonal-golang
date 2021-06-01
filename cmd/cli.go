/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {
		// Executa a ação solicitada ao service e imprime o resultado
		res, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / disable")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Id do Produto")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Nome do Procuto")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Preço do produto")
}
