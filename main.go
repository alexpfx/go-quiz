/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/alexpfx/go-quiz/cmd"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	cmd.Execute()
}
