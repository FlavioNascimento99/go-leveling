/*
	Nivelamento em GoLang
	Flavio Nascimento
	03.2026

	Objetivo: Construir uma aplicação que traga conceitos 
	básicos da linguagem. 

	Aplicação: Criação de uma aplicação para gerenciamento 
	de números em CLI, deverá permitir o usuário adicionar,
	listar, remover e calcular determinadas estatísticas 
	sobre os mesmos.
*/

package main

import (
	"bufio"			// Scanner de I/O
	"fmt"			// Fomatação e saída de dados (CLI), Print(e variáveis amplamente utilizados)
	"os"			// Interações para com Sistema Operacional
	"sort" 			// Ordenação de listas (literalmente utilizado de maneira exclusiva pela função de ordenação)
	"strconv"		// 
	"strings"
)

func main() {

	/*
		Criação de lista de tipagem inteira.

		Declaração e criação de instância (.NewScanner) de variável (tipo Scanner) a ser utilizado como ponteiro dentro das demais funções.
	*/
	numeros := []int{}
	scanner := bufio.NewScanner(os.Stdin) // (os.Std) define a entrada default do OS como parâmetro da função.

	for {
		exibirMenu()
		fmt.Print("Escolha uma opção: ")
		scanner.Scan()
		opcao := strings.TrimSpace(scanner.Text())

		switch opcao {
		case "1":
			numeros = adicionarNumero(numeros, scanner)
		case "2":
			listarNumeros(numeros)
		case "3":
			numeros = removerPorIndice(numeros, scanner)
		case "4":
			calcularEstatisticas(numeros)
		case "5":
			divisaoSegura(scanner)
		case "6":
			numeros = []int{}
			fmt.Println("Lista limpa com sucesso")
		case "7":
			ordenarLista(numeros, scanner)
		case "8":
			exibirPares(numeros)
		case "9":
			exportarParaArquivo(numeros, scanner)
		case "0":
			fmt.Println("Encerrando...")
			return
		default:
			fmt.Println("Opção inválida")
		}
		fmt.Println()
	}
}


/*
	Função para renderização em terminal de menu de interação.
	sem input, ideia resume-se à construção
*/ 
func exibirMenu() {
	fmt.Println("=== Menu ===")
	fmt.Println("1) Adicionar número")
	fmt.Println("2) Listar números")
	fmt.Println("3) Remover por índice")
	fmt.Println("4) Estatísticas")
	fmt.Println("5) Divisão segura")
	fmt.Println("6) Limpar lista")
	fmt.Println("7) Ordenar lista")
	fmt.Println("8) Exibir pares")
	fmt.Println("9) Exportar para arquivo")
	fmt.Println("0) Sair")
}


/*
	Função de adição de número dentro da listagem de valores.
	@param: lista de inteiros
*/
func adicionarNumero(numeros []int, scanner *bufio.Scanner) []int {
	fmt.Print("Digite um número inteiro: ")
	scanner.Scan()

	
	texto := strings.TrimSpace(scanner.Text())

	// Tratamento de excessão: 
	numero, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Valor inválido")
		return numeros
	}

	// Tratamento de excessão: 
	if numero < 0 {
		fmt.Println("Não é permitido adicionar números negativos")
		return numeros
	}

	numeros = append(numeros, numero)
	fmt.Println("Número adicionado com sucesso")
	return numeros
}

/*
	Impressão de listagem de números cadastrados
*/
func listarNumeros(numeros []int) {
	if len(numeros) == 0 {
		fmt.Println("Lista vazia")
		return
	}

	fmt.Println("Números armazenados:")
	for i, num := range numeros {
		fmt.Printf("[%d] %d\n", i, num)
	}
}


/*
	Remover número a partir da posição do mesmo na lista.
*/
func removerPorIndice(numeros []int, scanner *bufio.Scanner) []int {
	if len(numeros) == 0 {
		fmt.Println("Lista vazia")
		return numeros
	}

	listarNumeros(numeros)
	fmt.Print("Digite o índice para remover: ")
	scanner.Scan()
	texto := strings.TrimSpace(scanner.Text())

	indice, err := strconv.Atoi(texto)
	if err != nil || indice < 0 || indice >= len(numeros) {
		fmt.Println("Índice inválido")
		return numeros
	}

	numeros = append(numeros[:indice], numeros[indice+1:]...)
	fmt.Println("Número removido com sucesso")
	return numeros
}


/*
	Impressão (valor mais alto e baixo) e média.
*/
func calcularEstatisticas(numeros []int) {
	if len(numeros) == 0 {
		fmt.Println("Lista vazia")
		return
	}

	minimo, maximo, media := estatisticas(numeros)
	fmt.Printf("Mínimo: %d\n", minimo)
	fmt.Printf("Máximo: %d\n", maximo)
	fmt.Printf("Média: %.2f\n", media)
}


/*

*/
func estatisticas(numeros []int) (int, int, float64) {
	minimo := numeros[0]
	maximo := numeros[0]
	soma := 0

	for _, num := range numeros {
		if num < minimo {
			minimo = num
		}
		if num > maximo {
			maximo = num
		}
		soma += num
	}

	media := float64(soma) / float64(len(numeros))
	return minimo, maximo, media
}


/*
	Divisão de valores utilizando Atoi, para conversão de ASCII para Inteiro
	TrimSpace vai remover espaços em branco enquanto que Scanner.Text espera 
	uma entrada do usuário.
*/
func divisaoSegura(scanner *bufio.Scanner) {
	fmt.Print("Digite o dividendo: ")
	scanner.Scan()
	dividendo, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Valor inválido")
		return
	}

	fmt.Print("Digite o divisor: ")
	scanner.Scan()
	divisor, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Valor inválido")
		return
	}

	resultado, err := dividir(dividendo, divisor)
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		return
	}

	fmt.Printf("Resultado: %d\n", resultado)
}


/*
	Função de divisão feita à mão, impede passagem de valor de divisor como
	0, afinal divisor não pode ser 0.
*/
func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("divisor não pode ser zero")
	}
	return dividendo / divisor, nil
}


/*
	Ordenação de Lista de números cadastrados, selecionável se 
	será crescente ou decrescente, utilizando uma cópia da lista 
	original.

	Para construir uma cópia em ordem crescente, utilizamos sort.Ints()
	Para construir uma cópia em ordem DEcrescente, utilizamos 
*/
func ordenarLista(numeros []int, scanner *bufio.Scanner) {
	if len(numeros) == 0 {
		fmt.Println("Lista vazia")
		return
	}

	copia := make([]int, len(numeros))
	copy(copia, numeros)

	fmt.Println("1) Crescente")
	fmt.Println("2) Decrescente")
	fmt.Print("Escolha: ")
	scanner.Scan()
	escolha := strings.TrimSpace(scanner.Text())

	if escolha == "1" {
		sort.Ints(copia)
	} else if escolha == "2" {
		/*
			IntSlice, vai internamente verificar tamanho da lista, checar valores e ordenar 
			de acordo com a chamada que faz parte, nesse caso o Reverse.
		*/
		sort.Sort(sort.Reverse(sort.IntSlice(copia)))
	} else {
		fmt.Println("Opção inválida")
		return
	}

	/*
		Impressão da ordenação
		Sairá basicamente (valor_indice, valor_numero)
	*/
	fmt.Println("Números ordenados:")
	for valor_indice, valor_numero := range copia {
		fmt.Println(valor_numero)
	}
}


/*
	Recorte de valores numéricos adicionados à lista que podem ser 
	pares.
	Basicamente, valores em que, em caso de divisão por 2, não possui sobra.
*/
func exibirPares(numeros []int) {
	pares := []int{}
	for _, num := range numeros {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}

	if len(pares) == 0 {
		fmt.Println("Nenhum número par encontrado")
		return
	}

	fmt.Println("Números pares:")
	for _, num := range pares {
		fmt.Println(num)
	}
}


/*
	Salvam os dados dentro de um arquivo de texto, em caso em que a lista
	1. Nomeamos o arquivo
	2. Em caso de erro, retorna print informativo
	3. Caso de sucesso, a impressão de valores (Indice, Valor) e arquivo 
	salvo
	Encerra função.
*/
func exportarParaArquivo(numeros []int, scanner *bufio.Scanner) {
	fmt.Print("Digite o nome do arquivo: ")
	scanner.Scan()
	nomeArquivo := strings.TrimSpace(scanner.Text())

	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo: %v\n", err)
		return
	}
	defer arquivo.Close()

	for _, num := range numeros {
		_, err := arquivo.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo: %v\n", err)
			return
		}
	}

	fmt.Println("Lista exportada com sucesso")
}
