package main

import (
	"errors"
	"fmt"
	"github.com/caiquelira/huffman"
	"os"
	"strings"
)

func main() {
	// ler argumentos passados pela linha de comando
	args := os.Args[1:]

	// checar se argumentos estao de acordos com o formato especificado
	if len(args) < 1 {
		e := errors.New("Invalid Arguments. Run with argument \"help.\"")
		panic(e)
	} else if len(args) == 1 {
		if args[0] == "help" {
			fmt.Println("----------------------------------\n\nrun: ./main <compress | decompress> <file_to_compress.txt | file_to_decompress.ch>\n\n----------------------------------")
			return
		} else {
			e := errors.New("Invalid Arguments. Run with argument \"help.\"")
			panic(e)
		}
	} else if len(args) > 2 {
		e := errors.New("Too many arguments! Run with argument \"help.\"")
		panic(e)
	} else if args[0] != "compress" && args[0] != "decompress" {
		e := errors.New("First argument must be 'compress' or 'decompress'! Run with argument \"help.\"")
		panic(e)
	} else if args[0] == "compress" {
		if !strings.HasSuffix(args[1], ".txt") {
			e := errors.New("The file to compress must be .txt extension.")
			panic(e)
		}
		//tentar abrir arquivo pedido
		fileName := args[1]
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		outputName := strings.TrimSuffix(fileName, ".txt") + ".ch"
		fmt.Printf("Compressing file: %s", fileName)
		huffman.Compress(file, outputName)
	} else if args[0] == "decompress" {
		if !strings.HasSuffix(args[1], ".ch") {
			e := errors.New("The file to decompress must be .ch extension.")
			panic(e)
		}
		//tentar abrir arquivo pedido
		fileName := args[1]
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		outputName := strings.TrimSuffix(fileName, ".ch") + ".txt"
		fmt.Printf("Decompressing file: %s ...\n", fileName)
		huffman.Decompress(file, outputName)
	}

}
