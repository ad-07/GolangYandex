package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func run() error {
	// Проверяем количество аргументов
	if len(os.Args) != 4 {
		return errors.New("usage: main <rows> <cols> <percent>")
	}

	// Выводим аргументы для отладки
	fmt.Println("Received arguments:", os.Args)

	// Считываем параметры из командной строки
	rows, err := strconv.Atoi(os.Args[1])
	if err != nil || rows <= 0 {
		return errors.New("rows must be a positive integer")
	}

	cols, err := strconv.Atoi(os.Args[2])
	if err != nil || cols <= 0 {
		return errors.New("cols must be a positive integer")
	}

	percent, err := strconv.Atoi(os.Args[3])
	if err != nil || percent < 0 || percent > 100 {
		return errors.New("percent must be an integer between 0 and 100")
	}

	// Формируем строку для записи в файл
	output := fmt.Sprintf("%dx%d %d%%", rows, cols, percent)

	// Записываем в файл config.txt
	file, err := os.Create("config.txt")
	if err != nil {
		return fmt.Errorf("failed to create config.txt: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(output)
	if err != nil {
		return fmt.Errorf("failed to write to config.txt: %w", err)
	}

	fmt.Println("Configuration saved to config.txt")
	return nil
}
