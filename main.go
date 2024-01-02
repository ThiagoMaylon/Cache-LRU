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
	var value interface{}
	for k, v := range c.Item {
		if k == key {
			value = v
		}
	}
	if value != 0 {
		return value
	} else {
		return -1
	}
}

func (c *CacheLru) set(key string, value interface{}) {
	// Recebe a chave e o valor. Salva um valor com base na chave se o número de itens salvos estiver abaixo da capacidade. Substitui o item com o acesso mais antigo se a capacidade de itens for atingida.
	_, exist := c.Item[key]
	if exist {
		c.Item[key] = value

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
	var cache = NewCacheLru(5)
	cache.set("Abacaxi", 3)
	cache.set("Abacate", 5)
	cache.set("Banana", 5)
	cache.set("Maça", 2)
	cache.set("Melancia", 2)
	cache.set("Uva", 2)

	count := 1
	for _, key := range cache.Order{
		value := cache.Item[key]
		fmt.Printf("%v) %v : %v\n",count, key,value)
		
		count++
	}
}
