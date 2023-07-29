package strings

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	alpha   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	numeric = []rune("0123456789")
	symbol  = []rune("!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/`~")
	urlsafe = []rune(string(alpha) + string(numeric) + "-._~()'!*:@,")
)

func randInt(max int) (*big.Int, error) {
	curridx, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		logrus.Panicf("Error generating index, %s", err.Error())
		return nil, err
	}

	return curridx, nil
}

func Generate(cmd *cobra.Command, args []string) {
	charsetArg, err := cmd.Flags().GetString("charset")
	if err != nil {
		logrus.Panicf("Error when retrieving charset flag, %s", err.Error())
		return
	}
	charsetArg = strings.TrimSpace(charsetArg)

	length, err := strconv.Atoi(args[0])
	if err != nil {
		logrus.Fatal("Cannot get or parse parameter <length>")
	}

	var charset []rune
	switch charsetArg {
	case "none":
		charset = []rune(string(alpha) + string(numeric) + string(symbol))
	case "urlsafe":
		charset = urlsafe
	case "alpha":
		charset = alpha
	case "numeric":
		charset = numeric
	case "symbol":
		charset = symbol
	case "alphanumeric":
		charset = []rune(string(alpha) + string(numeric))
	case "symnumeric":
		charset = []rune(string(numeric) + string(symbol))
	case "alphasymbolic":
		charset = []rune(string(alpha) + string(symbol))
	default:
		logrus.Fatalf("unsupported charset type %s", charsetArg)
	}

	result := []rune{}
	for i := 0; i < length; i++ {
		idx, _ := randInt(len(charset))
		result = append(result, charset[idx.Int64()])
	}
	fmt.Println(string(result))
}
