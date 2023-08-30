package cmd

import (
    "log"
    "os/exec"

    "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a virtual machine",
    Long: `Create a virtual machine using qemu.`,
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}
