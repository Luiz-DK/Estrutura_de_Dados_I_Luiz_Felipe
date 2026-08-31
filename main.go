package main

import "fmt"

func exibirMenu() {
	fmt.Println()
	fmt.Println("1 - Adicionar no início")
	fmt.Println("2 - Adicionar no fim")
	fmt.Println("3 - Adicionar em posição")
	fmt.Println("4 - Remover do início")
	fmt.Println("5 - Remover do fim")
	fmt.Println("6 - Remover de posição")
	fmt.Println("7 - Buscar por valor")
	fmt.Println("8 - Buscar por posição")
	fmt.Println("9 - Tamanho")
	fmt.Println("10 - Imprimir")
	fmt.Println("0 - Sair")
	fmt.Print("Escolha uma opção: ")
}

func main() {
	l := &lista{}

	for {
		exibirMenu()

		var opcao int
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarInicio(valor)

		case 2:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarFim(valor)

		case 3:
			var valor, posicao int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			if l.adicionarPosicao(valor, posicao) {
				fmt.Println("Inserido com sucesso.")
			} else {
				fmt.Println("Posição inválida.")
			}

		case 4:
			valor, ok := l.removerInicio()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 5:
			valor, ok := l.removerFim()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 6:
			var posicao int
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.removerPosicao(posicao)
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 7:
			var valorProcurado int
			fmt.Print("Valor procurado: ")
			fmt.Scan(&valorProcurado)
			indice, ok := l.posicao(valorProcurado)
			if ok {
				fmt.Println("Encontrado na posição:", indice)
			} else {
				fmt.Println("Valor não encontrado.")
			}

		case 8:
			var posicaoProcurada int
			fmt.Print("Posição procurada: ")
			fmt.Scan(&posicaoProcurada)
			valor, ok := l.valorNaPosicao(posicaoProcurada)
			if ok {
				fmt.Println("Valor:", valor)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 9:
			fmt.Println("Tamanho:", l.tamanho())

		case 10:
			l.imprimir()

		case 0:
			fmt.Println("Encerrando.")
			return

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
