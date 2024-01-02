# Projeto Cache LRU em Go


## Descrição

Este projeto tem como objetivo desenvolver uma estrutura de dados que simula um Cache LRU (Least Recently Used), que é um tipo de dado que armazena até X itens, removendo o item com acesso mais antigo quando a capacidade máxima é atingida.

## Estado do Projeto

**Andamento:** Concluído

## Problema

Nesse desafio, a tarefa é implementar uma estrutura/classe que represente o Cache LRU. A estrutura deve possuir os seguintes métodos/funções:

1. **Função Construtora:** Recebe a capacidade do cache como um número inteiro.

2. **Método Get:** Recebe uma chave (string) e retorna um valor de qualquer tipo. Retorna -1 se não houver um item correspondente à chave. Este método memoriza a ordem de acesso de cada chave.

3. **Método Set:** Recebe a chave e o valor. Salva um valor com base na chave se o número de itens salvos estiver abaixo da capacidade. Substitui o item com o acesso mais antigo se a capacidade de itens for atingida.

## Requisitos

- Escreva uma estrutura/classe que representa o Cache LRU.
- A estrutura deve possuir os três métodos/funções mencionados acima.
- O método Get retorna -1 se não existir item para a chave recebida.
- O método Set salva um valor com base na chave se o número de itens salvos for abaixo da capacidade.
- O método Set substitui o item com acesso mais antigo se a capacidade de itens foi atingida.

## O que vou explorar neste projeto

- Algoritmos
- Estrutura de Dados

