package sshlib

import (
	"bytes"
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CopyFile(host string, user string, key string, src string, dest string) {
	fmt.Printf("%s %s\n", host, key)
	clientConfig, err := auth.PrivateKey(user, key, ssh.InsecureIgnoreHostKey())
	client := scp.NewClient(fmt.Sprintf("%s:22", host), &clientConfig)
	if err != nil {
		panic(err)
	}
	err = client.Connect()
	if err != nil {
		fmt.Println("Couldn't establisch a connection to the remote server ", err)
		panic(err)
	}
	defer client.Close()

	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := strings.Split(src, "/")
	baseName := s[len(s)-1]

	fmt.Printf("%s", baseName)

	client.CopyFile(f, fmt.Sprintf("/tmp/%s", baseName), "0655")

}

func ExecuteCommand(command string, hostname string, user string, keyPath string) string {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, "22"), config)
	if err != nil {
		panic(err)
	}
	session, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)
	return stdoutBuf.String()
}
