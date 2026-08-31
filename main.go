package main

func main() {
	l := &lista{}

	l.adicionarFim(20)
	l.adicionarFim(30)
	l.adicionarInicio(10)
	l.adicionarFim(40)
	l.adicionarInicio(5)

	l.imprimir()

	l.adicionarPosicao(15, 2)
	l.imprimir()

	ok := l.adicionarPosicao(0, 0)
	println("inserir na posição 0:", ok)
	l.imprimir()

	okFim := l.adicionarPosicao(50, 7)
	println("inserir no fim (posição 7):", okFim)
	l.imprimir()

	okInvalida := l.adicionarPosicao(99, 100)
	println("posição inválida (100):", okInvalida)

	okNegativa := l.adicionarPosicao(99, -1)
	println("posição negativa (-1):", okNegativa)

	valorIni, okIni := l.removerInicio()
	println("removerInicio ->", valorIni, okIni)
	l.imprimir()

	valorFim, okFim2 := l.removerFim()
	println("removerFim ->", valorFim, okFim2)
	l.imprimir()

	unica := &lista{}
	unica.adicionarInicio(7)
	v, ok := unica.removerFim()
	println("removerFim (lista com 1 elemento) ->", v, ok)
	unica.imprimir()

	vazia := &lista{}
	vIni, okVazia1 := vazia.removerInicio()
	println("removerInicio (lista vazia) ->", vIni, okVazia1)
	vFim, okVazia2 := vazia.removerFim()
	println("removerFim (lista vazia) ->", vFim, okVazia2)
}
