/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"bases/src/bases"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert numbers from some numeric base to another",
	Long: `Convert some number on any numeric base between 1 to 30 to 
another numeric bases on the same range.

Example:
	-> bases convert 101 -b 2 --to 10
	-> 101₂ = 5₁₀`,
	Run: func(cmd *cobra.Command, args []string) {

		errorMessage := "\nThe parameter and flags -b and --to are required: \n" +
						"\tbases convert [parameter] -b [int] --to [int]\n\n" +
						"Example:\n \tbases convert 101 -b 2 --to 10 \n\n"+
						"NOTE: the -b and --to flags must be grand than 1 \n\n"

		if len(args) <= 0 {
			fmt.Print(errorMessage)
			return
		}

		//Get the flags and parameters
		number := strings.ToUpper(args[0])
		base, _ := cmd.Flags().GetInt("base")
		to, _ := cmd.Flags().GetInt("to")

		if base <= 1 || to <= 1 {
			fmt.Print(errorMessage)
			return
		}

		if base > 30 || to > 30 {
			fmt.Print("bases not supported, use bases only between 2 and 30 ")
			return
		}

		//Do the conversion of the bases
		result, err := bases.ConvertBases(number, base, to)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//Print the result
		fmt.Println(number + bases.SubIndex(base) + " = " + result)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().IntP("base","b",0,"The base of the given number")
	convertCmd.Flags().IntP("to", "t" , 0,"The bases that you want to convert")
}
