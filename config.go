package logger

import (
	"os"
)

// SetStatusExit Determina o estado de erro
func SetStatusExit(status int) {
	statusExit = status
}

// SetDebug Determina se os loggers de DEBUG aparecerá
func SetDebug(ative bool) {
	debugAtive = ative
}

// SetClock Determina se as informações de data serão mostradas
func SetClock(ative bool) {
	clockAtive = ative
}

// SetFileLocal Determina o local para gravação do log
func SetFileLocal(local string) {
	if local != "" {
		opFile, err := os.OpenFile(local, os.O_WRONLY|os.O_SYNC|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			errorInterno("Erro ao abrir o arquivo:", err)
		}
		fileLocal = opFile
		fileSave = true
	}
}

// CloseFile Fecha o arquivo
func CloseFile() {
	if fileSave {
		fileLocal.Close()
	}
}
