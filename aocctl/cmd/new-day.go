package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var newDayCmd = &cobra.Command{
	Use:   "new-day",
	Short: "Creates a new day",

	Run: func(cmd *cobra.Command, args []string) {
		number, _:= cmd.Flags().GetString("number")

		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range files {
			if f.IsDir() && f.Name() == "day-"+number {
				fmt.Println("Directory already exists")
				os.Exit(1)
			}
		}

		if err := os.Mkdir("day-"+number, os.ModePerm); err != nil {
			fmt.Println(err)
		}

		addReadme, _ := cmd.Flags().GetBool("readme")
		if addReadme {
			out, err := os.Create("day-"+number+"/README.md")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(newDayCmd)
	newDayCmd.Flags().StringP("number", "n", "", "Set the number of the day you want to create.")
	newDayCmd.Flags().BoolP("readme", "r", false, "Set if you want a README.md file." )

	err := cobra.MarkFlagRequired(newDayCmd.Flags(), "number")
	if err != nil {
		fmt.Println(err)
	}
}
