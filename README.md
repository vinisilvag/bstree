# bstree

Esse repositório armazena uma implementação de Árvore Binária de Busca na linguagem de programação Go.

Ela foi desenvolvida como sistema alvo para a criação de pipelines de CI/CD para as aulas práticas 5 e 6 da disciplina de Engenharia de Software 2.

## Como executar

O sistema não possui um ponto inicial onde possa ser executado propriamente, ele foi pensado como uma biblioteca que exporta as estruturas e funções utilizadas. No entanto, os testes de unidade desenvolvidos para os *structs* `Node` e `BinarySearchTree` podem ser verificados ao executar

```
go test -v ./...
```

na pasta raiz do projeto.
