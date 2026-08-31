package main

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}
	novo.proximo = l.inicio
	l.inicio = novo
}

func (l *lista) adicionarFim(valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}
	atual.proximo = novo
}

func (l *lista) adicionarPosicao(valor, posicao int) bool {
	if posicao < 0 {
		return false
	}

	novo := &no{valor: valor}

	if posicao == 0 {
		novo.proximo = l.inicio
		l.inicio = novo
		return true
	}

	anterior := l.inicio
	for i := 0; i < posicao-1; i++ {
		if anterior == nil {
			return false
		}
		anterior = anterior.proximo
	}

	if anterior == nil {
		return false
	}

	novo.proximo = anterior.proximo
	anterior.proximo = novo
	return true
}

func (l *lista) removerInicio() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	valor := l.inicio.valor
	l.inicio = l.inicio.proximo
	return valor, true
}

func (l *lista) removerFim() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	if l.inicio.proximo == nil {
		valor := l.inicio.valor
		l.inicio = nil
		return valor, true
	}

	anterior := l.inicio
	for anterior.proximo.proximo != nil {
		anterior = anterior.proximo
	}

	valor := anterior.proximo.valor
	anterior.proximo = nil
	return valor, true
}

func (l *lista) removerPosicao(posicao int) (int, bool) {
	if posicao < 0 || l.inicio == nil {
		return 0, false
	}

	if posicao == 0 {
		valor := l.inicio.valor
		l.inicio = l.inicio.proximo
		return valor, true
	}

	anterior := l.inicio
	for i := 0; i < posicao-1; i++ {
		if anterior.proximo == nil {
			return 0, false
		}
		anterior = anterior.proximo
	}

	if anterior.proximo == nil {
		return 0, false
	}

	valor := anterior.proximo.valor
	anterior.proximo = anterior.proximo.proximo
	return valor, true
}

func (l *lista) posicao(valorProcurado int) (int, bool) {
	atual := l.inicio
	indice := 0
	for atual != nil {
		if atual.valor == valorProcurado {
			return indice, true
		}
		atual = atual.proximo
		indice++
	}
	return 0, false
}

func (l *lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 {
		return 0, false
	}

	atual := l.inicio
	indice := 0
	for atual != nil {
		if indice == posicaoProcurada {
			return atual.valor, true
		}
		atual = atual.proximo
		indice++
	}
	return 0, false
}

func (l *lista) imprimir() {
	atual := l.inicio
	for atual != nil {
		print(atual.valor, " -> ")
		atual = atual.proximo
	}
	println("nil")
}
