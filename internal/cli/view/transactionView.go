package view

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptNote() (string, error) {
	fmt.Print("Adicione uma nota ou observação para essa transação: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("erro ao ler entrada: %w", err)
	}

	note := strings.TrimSpace(input)
	return note, nil
}
