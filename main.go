// no começo de todo arquivo go vai ter o package, ele vai falar qual o pacote
// que eu estou trabalhando
// se eu estiver no package main e eu tiver uma função chamada main
// (func main() {}) quando eu rodar o meu programa ele vai executar a função
// main primeiro
package main

import (
	"github.com/Rodrigo001-de/iniciando-golang/http"
	"github.com/Rodrigo001-de/iniciando-golang/model"
	uuid "github.com/satori/go.uuid"
)

// type Person struct {
// 	Name string
// 	Age int
// }

// aqui eu estou criando um método
// quando eu indico que essa Person é um ponteiro de Person ele vai valores da
// Person em si porque ele está apontando para o mesmo lugar da memória
// func (p *Person) setName(name string) {
// 	p.Name = name
// 	fmt.Println(p.Name)
// }

func main() {
	// http.HandleFunc("/product", ProductHandle)
	// http.ListenAndServe("9095", nil)
	// aqui eu estou declarando e atribuindo a variável ao mesmo tempo
	// nome := "Rodrigo"
	// com esse & na frente ele vai retornar o local da memória que está guardado
	// o nome
	// fmt.Println(&nome)
	// esse * na frente do tipo é um ponteiro quer dizer que ele é um apontamento
	// para algun lugar na memória
	// var nome2 *string

	// person := Person{
	// 	Name: "Rodrigo",
	// 	Age: 17,
	// }
	// person.setName("Rael")
	// fmt.Println(person.Name)

	product1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carro",
	}

	product2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	products := model.Products{}
	products.Add(product1)
	products.Add(product2)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()
}

// func ProductHandle(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>ola mundo</h1>"))
// }
