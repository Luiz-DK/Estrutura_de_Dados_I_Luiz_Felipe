package main

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

// adicionarInicio insere um novo nó no início da lista — O(1)
func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}
	novo.proximo = l.inicio
	l.inicio = novo
}

// adicionarFim insere um novo nó no final da lista — O(n)
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

// imprimir exibe a lista no formato "10 -> 20 -> 30 -> nil"
func (l *lista) imprimir() {
	atual := l.inicio
	for atual != nil {
		print(atual.valor, " -> ")
		atual = atual.proximo
	}
	println("nil")
}