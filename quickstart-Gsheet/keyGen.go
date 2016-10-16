package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
)

// SyncStruct Used to sync go rutines
type SyncStruct struct {
	PrintChan chan string
	DoneChan  chan bool
	LogChan   chan string
	FilesChan map[string]chan string
	Wgs       map[string]sync.WaitGroup
}

func newSyscStruct() *SyncStruct {
	var syncStruct SyncStruct
	syncStruct.DoneChan = make(chan bool, 5)
	syncStruct.PrintChan = make(chan string, 5)
	syncStruct.LogChan = make(chan string, 10)
	syncStruct.FilesChan = make(map[string]chan string) // NOTE each channel needs to be created when filepath is determined
	syncStruct.Wgs = make(map[string]sync.WaitGroup)
	return &syncStruct
}

func makeDir(path string) error {
	return os.MkdirAll(path, 0700)
}

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		return err
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err = pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pubKeyPath, ssh.MarshalAuthorizedKey(pub), 0655)
}

func connectToHost(user, host, pass string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

func checkForFile(path string) error {

	_, err := os.Open(path)

	return err
}
