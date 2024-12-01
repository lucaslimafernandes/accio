package utilities

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func SaveOutputAsFile(filepath, text *string) {

	newFile, err := os.Create(*filepath)
	if err != nil {
		log.Fatalf("failed to create file: %v\n", err)
	}
	defer newFile.Close()

	_, err = newFile.WriteString(*text)
	if err != nil {
		log.Fatalf("failed to write in file: %v\n", err)
	}

}

func ExtractVars(text string, vars map[string]string) string {

	start := strings.Index(text, "{{") + 2
	end := strings.Index(text, "}}")

	if start < end {
		extracted := text[start:end]

		return strings.TrimSpace(extracted)
	}

	return ""
}

func ExtractVarsReturn(text string, vars map[string]string) string {

	var result string
	start := 0

	for {
		// Encontrar as posições de {{ e }}
		s := strings.Index(text[start:], "{{")
		if s == -1 {
			// Se não encontrar mais, acrescenta o restante da string
			result += text[start:]
			break
		}

		// Adiciona o texto anterior ao resultado
		result += text[start : start+s]

		// Atualiza o índice de início para a próxima iteração
		start += s + 2
		e := strings.Index(text[start:], "}}")
		if e == -1 {
			// Caso não encontre o fechamento, interrompe
			break
		}

		// Extrai o nome da variável entre as chaves
		variable := strings.TrimSpace(text[start : start+e])

		// Substitui a variável pelo valor no mapa
		val, ok := vars[variable]
		if ok {
			result += val
		} else {
			// Se não encontrar a chave, mantém o texto original
			result += "{{" + variable + "}}"
		}

		// Atualiza o índice de início para o próximo bloco de texto
		start += e + 2
	}

	// Exibe o resultado final
	fmt.Println(result)

	return result
}
