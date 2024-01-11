/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package th

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	teleportCluster string
	user            string
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "This will set the Teleport Cluster into your local machine",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
		// Logic

	},
}

func init() {

	setCmd.Flags().StringVar(&teleportCluster, "clusterDetails", "", "Teleport Cluster you want to connect")
	setCmd.Flags().StringVar(&user, "user", "", "Teleport Cluster you want to connect")

	if err := setCmd.MarkFlagRequired("clusterDetails"); err != nil {
		fmt.Println(err)
	}
	if err := setCmd.MarkFlagRequired("user"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	TshCmd.AddCommand(setCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
