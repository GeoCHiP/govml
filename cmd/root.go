package cmd

import (
    "os"

    "github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
    Use:   "govml",
    Short: "Tool for easy and transparent management of qemu virtual machines",
    Long: `vml is a tool for easy and transparent management of
qemu virtual machines. Virtual machines are presented as
directories with vml.toml files in it. vml is able to initialize
images with cloud-init. Virtual machines with ALT, Centos, Debian,
Fedora, openSUSE and Ubuntu could be created with just one command`,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
}
