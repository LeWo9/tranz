package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ввод доходов
	fmt.Print("Введите ваш месячный доход: ")
	incomeInput, _ := reader.ReadString('\n')
	income, _ := strconv.Atoi(strings.TrimSpace(incomeInput))

	// Категории и расходы
	expenses := make(map[string]int)
	categories := []string{"еда", "транспорт", "жилье"}
	for _, category := range categories {
		fmt.Printf("Введите расходы на %s: ", category)
		expenseInput, _ := reader.ReadString('\n')
		expense, _ := strconv.Atoi(strings.TrimSpace(expenseInput))
		expenses[category] = expense
	}

	// Расчет баланса
	totalExpenses := 0
	for _, expense := range expenses {
		totalExpenses += expense
	}
	netIncome := income - totalExpenses
	fmt.Printf("Чистый доход: %d\n", netIncome)

	// Анализ бюджета
	for category, expense := range expenses {
		percentage := (expense * 100) / income
		if percentage > 30 {
			fmt.Printf("Внимание: расходы на %s составляют %d%% от вашего дохода. Рекомендуется снизить траты в этой категории.\n", category, percentage)
		}
	}
}
