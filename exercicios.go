package main

// Preprocessamento guarda o que Preprocessa calcula e Distancia consulta.
type Preprocessamento struct {
	distancias []int
}

// Preprocessa recebe um grafo linha, com pelo menos dois vértices e pesos
// positivos, e prepara as consultas de distância. O custo deve ser O(V).
func Preprocessa(g *GrafoP) *Preprocessamento {

	n := g.V() // g.v() devolve o número de vértices do grafo
	distancias := make([]int, n)

	ponta := 0
	// procura a ponta do grafo, ou seja, vértice tem grau 1, ou seja, só tem uma aresta conectada
	for v := 0; v < n; v++ {
		if g.Grau(v) == 1 {
			ponta = v
			break
		}
	}

	anterior := -1 // por enquanto, ninguém
	atual := ponta
	// percorre o grafo, calculando as distâncias
	for i := 0; i < n-1; i++ {
		for _, a := range g.Adj(atual) {
			vizinho := a.Outro(atual)
			if vizinho != anterior {
				distancias[vizinho] = distancias[atual] + a.Peso
				anterior = atual
				atual = vizinho
				break
			}
		}
	}
	return &Preprocessamento{distancias: distancias}
}

// Distancia devolve o comprimento do caminho mais curto entre os vértices v e
// w. O custo deve ser O(1).
func (p *Preprocessamento) Distancia(v, w int) int {
	d := p.distancias[v] - p.distancias[w]
	// calcula o módulo
	if d < 0 {
		d = -d
	}
	return d
}
