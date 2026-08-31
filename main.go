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

	testePos := &lista{}
	testePos.adicionarFim(10)
	testePos.adicionarFim(20)
	testePos.adicionarFim(30)
	testePos.adicionarFim(40)
	testePos.imprimir()

	v0, ok0 := testePos.removerPosicao(0)
	println("removerPosicao(0) ->", v0, ok0)
	testePos.imprimir()

	v1, ok1 := testePos.removerPosicao(1)
	println("removerPosicao(1) ->", v1, ok1)
	testePos.imprimir()

	vUlt, okUlt := testePos.removerPosicao(1)
	println("removerPosicao(fim) ->", vUlt, okUlt)
	testePos.imprimir()

	vInv, okInv := testePos.removerPosicao(10)
	println("removerPosicao(inválida) ->", vInv, okInv)

	vNeg, okNeg := testePos.removerPosicao(-1)
	println("removerPosicao(negativa) ->", vNeg, okNeg)

	vaziaPos := &lista{}
	vVazia, okVaziaPos := vaziaPos.removerPosicao(0)
	println("removerPosicao (lista vazia) ->", vVazia, okVaziaPos)

	testeBusca := &lista{}
	testeBusca.adicionarFim(10)
	testeBusca.adicionarFim(20)
	testeBusca.adicionarFim(30)
	testeBusca.imprimir()

	iInicio, okInicio := testeBusca.posicao(10)
	println("posicao(10) ->", iInicio, okInicio)

	iMeio, okMeio := testeBusca.posicao(20)
	println("posicao(20) ->", iMeio, okMeio)

	iFim, okFimBusca := testeBusca.posicao(30)
	println("posicao(30) ->", iFim, okFimBusca)

	iInexistente, okInexistente := testeBusca.posicao(99)
	println("posicao(99) ->", iInexistente, okInexistente)

	buscaVazia := &lista{}
	iVazia, okBuscaVazia := buscaVazia.posicao(5)
	println("posicao (lista vazia) ->", iVazia, okBuscaVazia)
}
