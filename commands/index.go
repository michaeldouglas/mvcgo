package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func CreateController(controllerName string) error {
	funcName := fmt.Sprintf("%sController", controllerName)

	controllerContent := []byte(`package controllers
	import (
			"github.com/gofiber/fiber/v2"
	)

	func ` + controllerName + `(c *fiber.Ctx) error {
			return c.SendString("Hello, World! FuncName: ` + controllerName + `")
	}`)

	// Obter o caminho atual
	dir, errGetWd := os.Getwd()
	if errGetWd != nil {
		return errGetWd
	}

	// Concatenar com um novo caminho
	controllersPath := path.Join(dir, "controllers")

	// Criar o diretório caso não exista
	err := os.MkdirAll(controllersPath, 0755)
	if err != nil {
		return err
	}

	// Salva o arquivo caso não exista
	controllerFilePath := filepath.Join(controllersPath, funcName+".go")
	if _, err := os.Stat(controllerFilePath); os.IsNotExist(err) {
		err = ioutil.WriteFile(controllerFilePath, controllerContent, 0644)
		if err != nil {
			return err
		}
	} else {
		erroFileExist := fmt.Sprintf("A Controller: %s já existe!", controllerName)
		panic(erroFileExist)
	}

	return nil
}

func Run(args []string) {
	if args[0] == "--controller" {
		if CreateController(args[1]) == nil {
			fmt.Println("Controller criada com sucesso!")
		}
	}
}
