package logger

func save(logStr string) {
	if _, err := fileLocal.Write([]byte(logStr)); err != nil {
		errorInterno("Erro ao salvar no arquivo:", err)
	}
}
