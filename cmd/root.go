package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"todo-cli/app"
)

const TodoRoot = "todo"

var (
	dateStr string

	rootCmd = &cobra.Command{
		Use:   "todo-cli",
		Short: "todo kanri on cli",
		Run: func(cmd *cobra.Command, args []string) {
			dDst, err := app.BaseNameToTime(dateStr)
			if err != nil {
				log.Fatal(err)
			}
			dSrc := dDst.AddDate(0, 0, -1)
			pDst := filepath.Join(TodoRoot, fmt.Sprintf("%s.md", dDst.Format(app.TimeFormat)))
			pSrc := filepath.Join(TodoRoot, fmt.Sprintf("%s.md", dSrc.Format(app.TimeFormat)))
			_, err = os.Stat(pSrc)
			if os.IsNotExist(err) {
				log.Fatal(err)
			}
			_, err = os.Stat(pDst)
			if !os.IsNotExist(err) {
				log.Fatal(fmt.Errorf("%s is alredy exist", pDst))
			}

			fSrc, err := os.Open(pSrc)
			if err != nil {
				log.Fatal(err)
			}
			defer fSrc.Close()

			fDst, err := os.Create(pDst)
			if err != nil {
				log.Fatal(err)
			}
			defer fDst.Close()

			app.CarryOver(fSrc, fDst)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dateStr, "date", "D", app.TimeToBasename(time.Now()), "Date to show todo")
}
