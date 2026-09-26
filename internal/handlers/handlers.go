package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetHTML(w http.ResponseWriter, r *http.Request) {
	dir := "/home/vladislav/dev/vladislav_go1fl-sprint6-final-tpl/"
	htmlPath := filepath.Join(dir, "index.html")

	data, err := os.ReadFile(htmlPath)
	if err != nil {
		fmt.Println(err, err)
		http.Error(w, "Внутренняя ошибка сервера, GetHTML", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Внутренняя ошибка сервера, ParseMultipartForm", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	convertedData, err := service.MorseTextConv(string(data))
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка конвертации: %v", err), http.StatusInternalServerError)
		return
	}

	localFileName := time.Now().UTC().Format("20060102_150405")
	localFileExt := strings.ToLower(filepath.Ext(header.Filename))
	if localFileExt == "" {
		localFileExt = ".txt"
	}
	localFile, err := os.Create(localFileName + localFileExt)
	if err != nil {
		http.Error(w, "Не удалось создать локальный файл", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	_, err = io.WriteString(localFile, convertedData)
	if err != nil {
		http.Error(w, "Ошибка записи в локальный файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedData))
}
