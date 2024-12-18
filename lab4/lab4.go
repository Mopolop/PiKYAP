package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	err := writeToFile("output.txt", "ПиКЯП 3 семестр Князев")
	if err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
		return
	}

	contents, err := readFile("output.txt")
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}
	fmt.Println("Текст из файла:", contents)

	wordCount := countWords(contents)
	fmt.Println("Количество слов:", wordCount)

	err = copyFile("output.txt", "copy.txt")
	if err != nil {
		fmt.Printf("Ошибка при копировании файла: %v\n", err)
		return
	}
	fmt.Println("Файл успешно скопирован.")
}

func writeToFile(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		return fmt.Errorf("не удалось записать текст в файл: %w", err)
	}
	return writer.Flush()
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать файл: %w", err)
	}
	return string(data), nil
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func copyFile(source, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("не удалось открыть исходный файл: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("не удалось создать целевой файл: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("не удалось скопировать данные: %w", err)
	}
	return nil
}
