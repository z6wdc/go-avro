package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/infra/avro"
	"github.com/z6wdc/go-avro/internal/usecase"
)

var (
	encodeInput  string
	encodeOutput string
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode a Notification into Avro binary",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(encodeInput)
		if err != nil {
			return fmt.Errorf("failed to read input JSON: %w", err)
		}

		var n entity.Notification
		if err := json.Unmarshal(data, &n); err != nil {
			return fmt.Errorf("invalid JSON format: %w", err)
		}

		codec, err := avro.NewNotificationCodec("internal/infra/avro/schema_v1.avsc")
		if err != nil {
			return err
		}

		uc := usecase.NewEncodeNotificationUseCase(codec)
		encoded, err := uc.Execute(&n)
		if err != nil {
			return err
		}

		if err := os.WriteFile(encodeOutput, encoded, 0644); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}

		fmt.Printf("Encoded to %s (%d bytes)\n", encodeOutput, len(encoded))
		return nil
	},
}

func init() {
	encodeCmd.Flags().StringVar(&encodeInput, "input", "input.json", "Input JSON file")
	encodeCmd.Flags().StringVar(&encodeOutput, "output", "output.avro", "Output binary file")
}
