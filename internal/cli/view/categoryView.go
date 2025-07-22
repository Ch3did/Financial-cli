package view

import (
	"bufio"
	"financial-cli/internal/config/utils"
	"financial-cli/internal/domain/category"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func PromptCategory(categories []category.Category, txDescription string, amount float64, date string) (uint, error) {
	RunIfNotDebug(ClearScreen)
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
		err := fmt.Errorf("erro ao ler entrada: %w", err)
		time.Sleep(3 * time.Second)
		return 0, err
	}

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		err := fmt.Errorf("ID inválido: %w", err)
		time.Sleep(3 * time.Second)
		return 0, err
	}
	return uint(id), nil
}

func PromptNewCategory() (*category.Category, error) {
	RunIfNotDebug(ClearScreen)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Criar nova categoria")

	// Nome
	fmt.Print("Nome da categoria: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("erro ao ler nome: %w", err)
		time.Sleep(3 * time.Second)
		return nil, err
	}
	name = strings.TrimSpace(name)

	// Descrição
	fmt.Print("Descrição: ")
	description, err := reader.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("erro ao ler descrição: %w", err)
		time.Sleep(3 * time.Second)
		return nil, err
	}
	description = strings.TrimSpace(description)

	// Valor esperado
	fmt.Print("Valor esperado (ex: 500.00): ")
	expStr, err := reader.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("erro ao ler valor esperado: %w", err)
		time.Sleep(3 * time.Second)
		return nil, err
	}
	expStr = strings.TrimSpace(expStr)

	expected, err := strconv.ParseFloat(expStr, 64)
	if err != nil {
		err := fmt.Errorf("valor inválido para 'Expected': %w", err)
		time.Sleep(3 * time.Second)
		return nil, err
	}

	return &category.Category{
		Name:        name,
		Description: description,
		Expected:    expected,
	}, nil
}
