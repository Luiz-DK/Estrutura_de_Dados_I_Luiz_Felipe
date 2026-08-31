package main

func main() {
	l := &lista{}

	l.adicionarFim(20)
	l.adicionarFim(30)
	l.adicionarInicio(10)
	l.adicionarFim(40)
	l.adicionarInicio(5)

	l.imprimir() // esperado: 5 -> 10 -> 20 -> 30 -> 40 -> nil
}