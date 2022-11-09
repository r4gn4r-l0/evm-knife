/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/r4gn4r-l0/evm-knife/cmd/get"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evm-knife",
	Short: "The swiss army knife for EVM-Hacking",
	Long: `
	________  __     __  __       __        __    __            __   ______           
	|        \|  \   |  \|  \     /  \      |  \  /  \          |  \ /      \          
	| $$$$$$$$| $$   | $$| $$\   /  $$      | $$ /  $$ _______   \$$|  $$$$$$\ ______  
	| $$__    | $$   | $$| $$$\ /  $$$      | $$/  $$ |       \ |  \| $$_  \$$/      \ 
	| $$  \    \$$\ /  $$| $$$$\  $$$$      | $$  $$  | $$$$$$$\| $$| $$ \   |  $$$$$$\
	| $$$$$     \$$\  $$ | $$\$$ $$ $$      | $$$$$\  | $$  | $$| $$| $$$$   | $$    $$
	| $$_____    \$$ $$  | $$ \$$$| $$      | $$ \$$\ | $$  | $$| $$| $$     | $$$$$$$$
	| $$     \    \$$$   | $$  \$ | $$      | $$  \$$\| $$  | $$| $$| $$      \$$     \
	 \$$$$$$$$     \$     \$$      \$$       \$$   \$$ \$$   \$$ \$$ \$$       \$$$$$$$
																					   	
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(get.GetCmd)
}
