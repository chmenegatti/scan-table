package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scan-table",
	Short: "Um aplicativo CLI para buscar campos em tabelas de banco de dados",
	Long:  `scan-table é um aplicativo CLI para buscar campos em tabelas de banco de dados usando GoLang.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
