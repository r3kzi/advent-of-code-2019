package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "aocctl",
  Short: "Simplifying Advent of Code 2019",
  Long: "With this command you are able to simplify your daily doing with Advent of Code",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
