package eddsa

import (
	"crypto/ed25519"
	"encoding/pem"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Generate(cmd *cobra.Command, args []string) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		logrus.Errorf("[ERR] %s", err.Error())
		return
	}

	out, err := cmd.Flags().GetString("out")
	if err != nil {
		logrus.Panicf("Error when retrieving charset flag, %s", err.Error())
		return
	}

	pubpem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: []byte(pub),
	})

	privpem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: []byte(priv),
	})

	outClean := path.Clean(out)
	dir := path.Dir(outClean)

	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		logrus.Errorf("[Err] %s", err.Error())
		return
	}

	f, err := os.OpenFile(outClean, os.O_CREATE, os.FileMode(os.O_RDWR))
	if err != nil {
		logrus.Errorf("[Err] %s", err.Error())
		return
	}
	f.WriteString(string(privpem))
	f.Close()

	f, err = os.OpenFile(outClean+".pub", os.O_CREATE, os.FileMode(os.O_RDWR))
	if err != nil {
		logrus.Errorf("[Err] %s", err.Error())
		return
	}
	f.WriteString(string(pubpem))
	f.Close()

	logrus.Infof("Private Key generated at: %s", outClean)
	logrus.Infof("Public Key generated at: %s", outClean+".pub")
	logrus.Infof("Base64 Public Key: %s", string(pub))
}
