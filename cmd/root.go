/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/alexpfx/go-quiz/prova"
	"github.com/alexpfx/go-quiz/screen"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var jsonExample bool
var screenExample bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-quiz",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if jsonExample {
			p := example()
			jsobj := prova.ToJson(&p)
			fmt.Println(jsobj)
			return
		}

		if screenExample {
			p := prova.FromJson(filepath.Join("data/example.json"))
			s := screen.NewScreen(p)
			_, err := s.Run()
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func example() prova.Prova {
	altQ1 := []prova.Alternativa{
		{Texto: "H2O2"},
		{Texto: "H2O"},
		{Texto: "NACL"},
		{Texto: "H2C"},
		{Texto: "O2"},
	}

	altQ2 := []prova.Alternativa{
		{Texto: "Pero Vaz de Caminha"},
		{Texto: "Cristovão Buarque"},
		{Texto: "Cristovão Colombo"},
		{Texto: "Pedro Álvares Cabral"},
		{Texto: "Pedro Álvares de Lara"},
	}
	questoes := []prova.Questao{
		{Enunciado: "Qual a fórmula da água?", Alternativas: altQ1, Correta: altQ1[1]},
		{Enunciado: "Quem descobriu o Brasil?", Alternativas: altQ2, Correta: altQ2[3]},
	}

	return prova.Prova{
		Questoes: questoes,
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&screenExample, "screenExample", "x", false, "Mostra uma tela de prova utilizando o json de exemplo")
	rootCmd.Flags().BoolVarP(&jsonExample, "jsonExample", "X", false, "Mostra na saída um json de exemplo contendo os campos requeridos")
}
