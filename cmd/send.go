/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"mail/cmd/mailUtils"
	"strings"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sender, secretPassword, receiver, subjectMessage, messageBody string
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Helps to send the mail",
	Long:  `Helps to send the mail`,
	// Args:                  cobra.ExactArgs(5),
	DisableFlagsInUseLine: true,
	Example:               `mail send FromMailId Password TOMailId Subject Message`,
	Run: func(cmd *cobra.Command, args []string) {
		var fromMailId, password, toMailId, subject, message string
		fmt.Println(sender + " " + secretPassword + " " + receiver + " " + subjectMessage + " " + messageBody)
		if len(sender) > 0 && len(secretPassword) > 0 && len(receiver) > 0 && len(subjectMessage) > 0 && len(messageBody) > 0 {
			fromMailId = sender
			password = secretPassword
			toMailId = receiver
			subject = subjectMessage
			message = messageBody
		} else if len(args) == 5 {
			fromMailId = args[0]
			password = args[1]
			toMailId = args[2]
			subject = args[3]
			message = args[4]
		} else {
			fmt.Println("Syntax Error")
			return
		}
		toMailIds := strings.Split(toMailId, ",")
		mailUtils.SendMail(fromMailId, password, toMailIds, subject, message)
		fmt.Println("Sent Mail successfully")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.PersistentFlags().StringVarP(&sender, "sender", "s", "", "Mail id of sender")
	sendCmd.PersistentFlags().StringVarP(&secretPassword, "password", "p", "", "Password of sender mail id")
	sendCmd.PersistentFlags().StringVarP(&receiver, "receiver", "r", "", "Mail id of receiver")
	sendCmd.PersistentFlags().StringVarP(&subjectMessage, "subject", "b", "", "Subject of the mail")
	sendCmd.PersistentFlags().StringVarP(&messageBody, "message", "m", "", "Message Body of the mail")
}
