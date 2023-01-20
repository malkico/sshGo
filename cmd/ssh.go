package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"log"
)

var RootCommandSSh = &cobra.Command{

	Use:   "sshCopy <user> <host> <psw>",
	Short: "Allows you to connect to a remote host and copy",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		hostKeyCallback, err := knownhosts.New("/home/hicham/.ssh/known_hosts")
		if err != nil {
			log.Fatal(err)
		}

		user := args[0]
		host := args[1]
		psw := args[2]

		fmt.Println("Connecting ...")

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(psw),
			},
			HostKeyCallback: hostKeyCallback,
		}
		// connect ot ssh server

		conn, err := ssh.Dial("tcp", host+":22", config)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		fmt.Println("Connecting successfull")

	},
}
