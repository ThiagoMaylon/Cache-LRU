package main

import "fmt"

type CacheLru struct {
	Capacity int
	Item     map[string]interface{}
	Order    []string
}

func NewCacheLru(capacity int) *CacheLru {
	return &CacheLru{
		Capacity: capacity,
		Item:     make(map[string]interface{}),
		Order:    make([]string, 0),
	}
}

func (c *CacheLru) get(key string) interface{} {
	// Recebe uma chave (string) e retorna um valor de qualquer tipo. Retorna -1 se não houver um item correspondente à chave. Este método memoriza a ordem de acesso de cada chave.
	_, exist := c.Item[key]
	if !exist {
		return -1
	} else {
		value := c.Item[key]
		for i, k := range c.Order {
			if k == key {
				c.Order = append(c.Order[:i], c.Order[i+1:]...)
				delete(c.Item, key)
			}
		}
		c.Item[key] = value
		c.Order = append(c.Order, key)
		return 0
	}
}

func (c *CacheLru) set(key string, value interface{}) {
	// Recebe a chave e o valor. Salva um valor com base na chave se o número de itens salvos estiver abaixo da capacidade. Substitui o item com o acesso mais antigo se a capacidade de itens for atingida.
	_, exist := c.Item[key]
	if exist {
		for i, k := range c.Order {
			if k == key {
				c.Order = append(c.Order[:i], c.Order[i+1:]...)
			}
		}
		c.Item[key] = value
		c.Order = append(c.Order, key)

	} else {
		c.Item[key] = value
		c.Order = append(c.Order, key)

		if len(c.Item) > c.Capacity {
			deleteKey := c.Order[0]
			c.Order = c.Order[1:]
			delete(c.Item, deleteKey)
			c.Item[key] = value
		}

	}

}

func main() {
	// Cria um novo cache com capacidade 4
	var cache = NewCacheLru(4)

	// Testando o método set adicionando alguns itens
	cache.set("Abacaxi", 3)
	cache.set("Abacate", 5)
	cache.set("Banana", 8)
	cache.set("Maçã", 2)

	// Testando o método get para a chave "Abacaxi"
	cache.get("Abacaxi")

	// Exibe os itens presentes no cache após as operações
	count := 1
	for _, key := range cache.Order {
		value := cache.Item[key]
		fmt.Printf("%v) %v : %v\n", count, key, value)

		count++
	}
}
