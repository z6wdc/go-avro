package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "go-avro",
    Short: "AVRO encoding/decoding demo tool",
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.AddCommand(encodeCmd)
    rootCmd.AddCommand(decodeCmd)
}
