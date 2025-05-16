package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/z6wdc/go-avro/internal/infra/avro"
	"github.com/z6wdc/go-avro/internal/usecase"
)

var (
	decodeInput  string
	decodeOutput string
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode Avro binary into a Notification",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(decodeInput)
		if err != nil {
			return fmt.Errorf("failed to read input binary: %w", err)
		}

		codec, err := avro.NewNotificationCodec("internal/infra/avro/schema_v1.avsc")
		if err != nil {
			return err
		}

		uc := usecase.NewDecodeNotificationUseCase(codec)
		notification, err := uc.Execute(data)
		if err != nil {
			return err
		}

		result, err := json.MarshalIndent(notification, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to encode JSON: %w", err)
		}

		if err := os.WriteFile(decodeOutput, result, 0644); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}

		fmt.Printf("Decoded to %s\n", decodeOutput)
		return nil
	},
}

func init() {
	decodeCmd.Flags().StringVar(&decodeInput, "input", "output.avro", "Input Avro binary file")
	decodeCmd.Flags().StringVar(&decodeOutput, "output", "output.json", "Output JSON file")
}
