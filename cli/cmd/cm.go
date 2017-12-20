// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/painterdrown/go-agenda/cli/entity/AgendaLog"
	"github.com/painterdrown/go-agenda/cli/entity/Meeting"
	"github.com/painterdrown/go-agenda/cli/entity/User"
	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentUser := User.UserState()
		if currentUser == "" {
			AgendaLog.OperateLog("[error]", "cm error => "+"not login")
			return
		}
		if cmd.Flags().NFlag() < 4 {
			fmt.Println("you must provide --title --participator --stime --etime")
			AgendaLog.OperateLog("[error]", "cm error => "+"you don't provide all flags --title --participator --stime --etime")
			return
		}
		participator, _ := cmd.Flags().GetString("participator")
		title, _ := cmd.Flags().GetString("title")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		meeting := Meeting.Meetings{currentUser, participator, title, stime, etime}
		err := Meeting.AddOneMeeting(meeting)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("cm a meeting successfully")
			AgendaLog.OperateLog("[info]", "cm a meeting successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "", "Meeting's title")
	cmCmd.Flags().StringP("participator", "p", "", "Meeting's participator, Use as like -p = weimumu,weimuuiui,weimujd")
	cmCmd.Flags().StringP("stime", "s", "", "Meeting's start time, Use like as -s = 2006-01-02 09:09:09")
	cmCmd.Flags().StringP("etime", "e", "", "Meeting's end time, Use like as -e = 2006-01-02 09:09:09")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
