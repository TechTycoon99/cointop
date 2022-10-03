package cointop

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ReadAPIKeyFromStdin reads the user inputed API from the stdin prompt
func (ct *Cointop) ReadAPIKeyFromStdin(name string) (string, error) {
	log.Debug("ReadAPIKeyFromStdin()")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s API Key: ", name)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
