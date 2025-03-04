package serra

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "unknown"

var count int64
var limit float64
var interactive bool
var name string
var oracle string
var rarity string
var set string
var since string
var sortby string
var cardType string
var unique bool

var rootCmd = &cobra.Command{
	Version:               Version,
	Long:                  `serra - Magic: The Gathering Collection Tracker`,
	Use:                   "serra",
	DisableFlagsInUseLine: true,
	SilenceErrors:         true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
