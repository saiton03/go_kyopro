package assemble

import (
	"github.com/spf13/cobra"
	"log"
)

var RootCmd = &cobra.Command{
	Use:   "go-kyopro",
	Short: "assemble kyopro library sources to one file",
	Run:   Root,
}

func init() {
	RootCmd.PersistentFlags().BoolP("clipBoard", "c", false, "if true, the output is copied to clip board")
}

func Root(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("%s: no go file specified", cmd.Use)
	}
	err := Assemble(cmd, args[0])
	if err != nil {
		log.Fatal(err)
	}

}
