package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Oomat-Dzhumagulov/6SprintFinal/internal/service"
)

func MainHandle(w http.ResponseWriter, r *http.Request) {
	html := "../../index.html"
	http.ServeFile(w, r, html)
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Неправельный метод", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении формы", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	conver := service.Convert(string(data))

	extFile := filepath.Ext(header.Filename)

	fileName := fmt.Sprintf("JamesBond_%s%s", time.Now().Format("2006-01-02_15-04-05"), extFile)

	JamesBond, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer JamesBond.Close()

	_, err = JamesBond.Write([]byte(conver))
	if err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(conver))
}
