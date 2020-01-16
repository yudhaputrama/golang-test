package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	sumCmd := &cobra.Command{
		Use:   "sum",
		Short: "Sum all argument",
		Args:  cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			total := SumString(args)
			fmt.Printf("Sum of arguments is %d\n", total)
		},
	}

	rootCmd := &cobra.Command{
		Use:   "go",
		Short: "Simple app doing calculation",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.AddCommand(sumCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
