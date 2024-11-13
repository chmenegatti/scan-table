package cmd

import (
	"fmt"
	"log"
	"scan-table/database"
	"scan-table/searcher"
	"strings"

	"github.com/spf13/cobra"
)

var fields string

// fieldsCmd representa o comando fields
var fieldsCmd = &cobra.Command{
	Use:   "fields",
	Short: "Busca por campos específicos nas tabelas do banco de dados",
	Long:  `Busca por campos específicos nas tabelas do banco de dados e retorna as tabelas que contenham pelo menos um desses campos.`,
	Run: func(cmd *cobra.Command, args []string) {
		runFields()
	},
}

func init() {
	rootCmd.AddCommand(fieldsCmd)

	fieldsCmd.Flags().StringVarP(&fields, "fields", "n", "", "Campos a serem buscados, separados por vírgula")
	fieldsCmd.MarkFlagRequired("fields")
}

func runFields() {
	db := database.GetDBInstance()
	searcher := searcher.TableSearcherFactory{}.CreateSearcher(db)

	tables, err := searcher.GetTables()
	if err != nil {
		log.Fatalf("Error getting tables: %v", err)
	}

	fieldsList := strings.Split(fields, ",")
	resultTables, err := searcher.SearchTablesWithFields(tables, fieldsList)
	if err != nil {
		log.Fatalf("Error searching tables: %v", err)
	}

	fmt.Println("Tabelas que contêm pelo menos um dos campos especificados:")
	for _, table := range resultTables {
		fmt.Println(table)
	}
}
