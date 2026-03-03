# Gerenciador de Números em Go

Um aplicativo de linha de comando (CLI) desenvolvido em Go que permite gerenciar uma coleção de números inteiros com várias operações como adição, listagem, remoção e cálculo de estatísticas.

## 📋 Sobre o Projeto

Este é um projeto educacional de nivelamento em GoLang que explora os conceitos fundamentais da linguagem, incluindo:

- Declaração de variáveis
- Estruturas condicionais
- Laços de repetição
- Funções
- Slices
- Funções com múltiplos retornos
- Tratamento de erros

## ✨ Funcionalidades Principais

| Opção | Descrição |
|-------|-----------|
| 1 | Adicionar número - Adiciona um número inteiro não-negativo ao slice |
| 2 | Listar números - Exibe todos os números armazenados com seus índices |
| 3 | Remover por índice - Remove um elemento específico da lista |
| 4 | Estatísticas - Calcula mínimo, máximo e média dos números |
| 5 | Divisão segura - Realiza divisão com tratamento de erro para divisor zero |
| 6 | Limpar lista - Esvazia completamente o slice de números |
| 0 | Sair - Encerra a aplicação |

## 🎁 Funcionalidades Bônus

| Opção | Descrição |
|-------|-----------|
| 7 | Ordenar lista - Ordena números em ordem crescente ou decrescente |
| 8 | Exibir pares - Mostra apenas os números pares da lista |
| 9 | Exportar arquivo - Salva a lista de números em um arquivo de texto |

## 🚀 Como Executar

### Pré-requisitos
- Go 1.16 ou superior instalado

### Executar a Aplicação

```bash
go run main.go
```

### Compilar Executável

```bash
go build -o gerenciador
./gerenciador
```

## 📖 Exemplo de Uso

```
=== Menu ===
1) Adicionar número
2) Listar números
3) Remover por índice
4) Estatísticas
5) Divisão segura
6) Limpar lista
7) Ordenar lista
8) Exibir pares
9) Exportar para arquivo
0) Sair
Escolha uma opção: 1
Digite um número inteiro: 42
Número adicionado com sucesso

Escolha uma opção: 1
Digite um número inteiro: 15
Número adicionado com sucesso

Escolha uma opção: 2
Números armazenados:
[0] 42
[1] 15

Escolha uma opção: 4
Mínimo: 15
Máximo: 42
Média: 28.50

Escolha uma opção: 0
Encerrando...
```

## 🔑 Conceitos Go Utilizados

### Múltiplos Retornos
A função `estatisticas()` demonstra como retornar múltiplos valores em Go:

```go
func estatisticas(numeros []int) (int, int, float64) {
    // retorna minimo, maximo e media
}
```

### Tratamento de Erro Padrão
A função `dividir()` segue o padrão Go de retornar erro:

```go
func dividir(dividendo, divisor int) (int, error) {
    if divisor == 0 {
        return 0, fmt.Errorf("divisor não pode ser zero")
    }
    return dividendo / divisor, nil
}
```

### Slices Dinâmicos
O programa usa slices para armazenar números com operações de:
- Adição: `append()`
- Remoção: slicing
- Cópia: `copy()` e `make()`

## 📂 Estrutura do Projeto

```
go-leveling/
├── main.go      # Arquivo principal com todas as funções
└── README.md    # Este arquivo
```

## 🛠️ Funções Implementadas

- `main()` - Loop principal com menu
- `exibirMenu()` - Exibe opções do menu
- `adicionarNumero()` - Adiciona número com validação
- `listarNumeros()` - Lista todos os números
- `removerPorIndice()` - Remove número por índice
- `calcularEstatisticas()` - Exibe estatísticas
- `estatisticas()` - Calcula min, max e média
- `divisaoSegura()` - Realiza divisão com erro handling
- `dividir()` - Função auxiliar de divisão
- `ordenarLista()` - Ordena crescente ou decrescente
- `exibirPares()` - Filtra e exibe números pares
- `exportarParaArquivo()` - Salva lista em arquivo

## ⚙️ Requisitos Técnicos Atendidos

✅ Implementado em um único arquivo `main.go`  
✅ Compilável e executável com `go run main.go`  
✅ Uso de slice para armazenar números  
✅ Funções modulares bem definidas  
✅ Tratamento de erro seguindo padrão Go  
✅ Múltiplos retornos em funções  
✅ Validação de entrada do usuário  
✅ Menu interativo com loop contínuo  

## 📝 Anotações

- O programa não permite adicionar números negativos
- Índices são exibidos a partir de 0
- A divisão segura validada contra divisor zero
- Exportação cria novo arquivo ou sobrescreve existente
- Ordenação não modifica a lista original

## 👤 Autor

Flavio Nascimento - Março de 2026

## 📄 Licença

Projeto educacional
