package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler возвращает HTML из файла index.html
func RootHandler(w http.ResponseWriter, r *http.Request) {

	// Указываем путь к index.html явно
	path := filepath.Join("..", "index.html") // Поднимаемся на уровень выше

	// Чтение файла index.html
	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read index.html", http.StatusInternalServerError)
		log.Printf("Error reading index.html: %v\n", err)
		return
	}

	// Установка заголовка Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Отправка содержимого файла в ответ
	w.Write(data)

}

// UploadHandler обрабатывает загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Конвертируем данные
	result, err := service.TextAndMorseCodeConversion(string(data))
	if err != nil {
		http.Error(w, "the string contains invalid characters", http.StatusMethodNotAllowed)
		return
	}

	// Генерируем имя файла для сохранения
	fileName := fmt.Sprintf("%s%s", time.Now().UTC().Format("20060102150405"), filepath.Ext("output.txt"))

	// Записываем результат в локальный файл
	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Unable to write file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации строки
	fmt.Fprintf(w, "Converted data saved to %s", fileName)
}
