package view

import (
	"encoding/json"
	"fmt"

	"bufio"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"financial-cli/internal/config/utils"
	"financial-cli/internal/domain/category"
)

func BaseOutput(i interface{}) {
	jsonData, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func ClearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PromptCategory(categories []category.Category, txDescription string, amount float64, date string) (uint, error) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Transação: %s | Valor: %.2f | Data: %s\n", txDescription, amount, utils.ParseOFXDate(date))
	fmt.Println("Escolha uma categoria para essa transação:")

	for _, c := range categories {
		fmt.Printf(" [%d] %s\n", c.ID, c.Name)
	}

	fmt.Print("\nDigite o ID da categoria: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("erro ao ler entrada: %w", err)
	}

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("ID inválido: %w", err)
	}
	ClearScreen()
	return uint(id), nil
}

func PromptNewCategory() (*category.Category, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Criar nova categoria")

	// Nome
	fmt.Print("Nome da categoria: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("erro ao ler nome: %w", err)
	}
	name = strings.TrimSpace(name)

	// Descrição
	fmt.Print("Descrição: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("erro ao ler descrição: %w", err)
	}
	description = strings.TrimSpace(description)

	// Valor esperado
	fmt.Print("Valor esperado (ex: 500.00): ")
	expStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("erro ao ler valor esperado: %w", err)
	}
	expStr = strings.TrimSpace(expStr)

	expected, err := strconv.ParseFloat(expStr, 64)
	if err != nil {
		return nil, fmt.Errorf("valor inválido para 'Expected': %w", err)
	}

	ClearScreen()
	return &category.Category{
		Name:        name,
		Description: description,
		Expected:    expected,
	}, nil
}
