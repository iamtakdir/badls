package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "badls",
	Short: "badls is going to kill ls ",
	Long:  "A cli tool ",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {

			mode := f.Mode()
			name := f.Name()
			fmt.Printf(" %s         %s  \n  ", mode ,name)
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error:", err)
	}
}
