package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var logger = logrus.New()

// Messager for general messaging/logging
type Messager interface {
	GetType() string
	GetMessage() string
}

// MsgChan short cut for make(chan Messager)
type MsgChan chan Messager

type message struct {
	kind string
	text string
}

func (msg message) GetType() string {
	return msg.kind
}

func (msg message) GetMessage() string {
	return msg.text
}

func connect(ip string) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	logFile, err := os.Open("logs.json")
	if err != nil {
		logger.Errorf("Error opening log file%s", err)
		return err
	}

	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = logFile
	sshConfig := &ssh.ClientConfig{
		User: "your_user_name",
		Auth: []ssh.AuthMethod{
			ssh.Password("your_password"),
		},
	}
	connStr := fmt.Sprintf("%s:%d", ip, 22)
	connection, err := ssh.Dial("tcp", connStr, sshConfig)
	if err != nil {

		return err
	}
	session, err := connection.NewSession()
	if err != nil {
		logger.Errorf("Failed to create session: %s", err)
		return err
	}
	/*
	   Before we will be able to run the command on the remote machine, we should create a pseudo terminal on the remote machine.
	   A pseudoterminal (or “pty”) is a pair of virtual character devices that provide a bidirectional communication channel.

	   We should create an xterm terminal that has 80 columns and 40 rows.
	*/
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err = session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		logger.Errorf("request for pseudo terminal failed: %s", err)
		return err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		logger.Errorf("Unable to setup stdin for session: %v", err)
		return err
	}
	go io.Copy(stdin, os.Stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		logger.Errorf("Unable to setup stdout for session: %v", err)
		return err
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		logger.Errorf("Unable to setup stderr for session: %v", err)
		return err
	}
	go io.Copy(os.Stderr, stderr)
	<-c
	return err
}
