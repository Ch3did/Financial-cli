package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func PromptNote() (string, error) {
	fmt.Print("Adicione uma nota ou observação para essa transação: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("erro ao ler entrada: %w", err)
		time.Sleep(3 * time.Second)
		return "", err
	}

	note := strings.TrimSpace(input)
	return note, nil
}
