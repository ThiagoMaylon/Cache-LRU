package main


type CacheLru struct {
	Capacity int
	Item     map[string]int
	Cache    []interface{}
}

func NewCacheLru(capacity int) *CacheLru {
	return &CacheLru{
        Capacity: capacity,
        Item:     make(map[string]int),
    }
}

func (c *CacheLru) get(chave string) interface{} {
	// Recebe uma chave (string) e retorna um valor de qualquer tipo. Retorna -1 se não houver um item correspondente à chave. Este método memoriza a ordem de acesso de cada chave.
	
}

func (c *CacheLru) set(key string, value int) {
	// Recebe a chave e o valor. Salva um valor com base na chave se o número de itens salvos estiver abaixo da capacidade. Substitui o item com o acesso mais antigo se a capacidade de itens for atingida.
	
}
func (c *CacheLru) updateOrder(){
	// Implemente a lógica para atualizar a ordem de acesso
}

func main() {
	
}
