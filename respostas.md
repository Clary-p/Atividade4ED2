# Atividade Avaliativa 4 - Respostas

Nome:Clarissa Pietrangelo e Sabrina Dorigoni Pelentir

GRR:20255551 e 20255562

Se a atividade foi feita em dupla, escreva aqui os dois nomes e os dois GRR. As
duas pessoas enviam o mesmo `.zip` na UFPR Virtual.

## Pergunta

1. Explique por que sua implementação respeita a complexidade computacional exigida 
   pelo enunciado: preprocessamento do grafo em tempo linear e apresentação da 
   distância entre qualquer par de vértices em tempo constante.

No preprocessamento, a implementação respeita a complexidade computacional porque cresce proporcionalmente ao número de vérticas (O(V)), já que ele precisa primeiro criar o vetor de distâncias, que precisa reservar espao de acordo com o número de vértices, crescendo conforme eles crescem (O(V)), no seguinte trecho: 
n := g.V()
distancias := make([]int, n)
E depois, o algoritmo precisa encontrar uma ponta ( que possui grau 1) como podemos ver no seguinte trecho:
    
    ponta := 0
     for v := 0; v < n; v++ {
    if g.Grau(v) == 1 {
        ponta = v
        break
    }
      };
 processo que, no pior dos casos, precisa percorrer todos os vértices até acha-la, comprovando que o algoritmo é O(V). Além disso, agora na parte de percorrer os grafos e encontrar as distâncias o algoritmo não chega em O(V²) apesar de ter um for dentro do outro,por conta de uma característica do grafo da atividade que é ter no máxinmo 2 vizinhos, então nessa operação (considerando que as arestas são E = V-1) realizamos (V-1)x2, uma no primeiro laço e outra no segundo, o que resulta em : O(2V), proporcional a duas vezes o número de vértices que é igual a O(V), já que não consideramos as constantes nesse processo de avaliação de complexidade. A soma ocorre de forma ocnstante no seguinte trecho e não influencia: distancias[vizinho] = distancias[atual] + a.Peso.. Total: O(V)+O(V)+O(V).

 Já a distância, é constante (O(1)) porque no pré-processamento já calculamos a distância de cada vértice até a ponta, então ela apenas faz o acesso ao vetor usando seu índice e faz a subtração: d := p.distancias[1] - p.distancias[7]
 e se o sinal for negativo, inverte: if d < 0 { d = -d}. Nenhum desses processos depende de uma variável como vértice ou arestas, é constante - O(1).

## Uso de IA

Diga se usou ferramentas de IA generativa nesta atividade e em que partes. Não
influencia a nota.

> Pedi para explicar certos trechos de main.go e grafo.go que havia tido dificuldade de entender
> Pedi para checar o código ao final para ver se estava correto, corrigi um erro no for que passava nos testes
