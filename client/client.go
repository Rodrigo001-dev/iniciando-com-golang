package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// o requestId vai ser um channel que é como se fosse um canal para as threads
	// se comunicarem
	requestId := make(chan int)

	concurrency := 10
	// vai rodar a função worker 2 vezes
	for i := 1; i <= concurrency; i++ {
		// quando eu coloco o go na frente do worker significa que vai gerar uma
		// outra thread, o worker vai ficar rodando em backgrond em paralelo com o
		// fluxo principal na aplicação
		// para que essas threads consigam se falar vai ser utilizando um channel
		go worker(requestId, i)
	}

	for i := 0; i < 20; i++ {
		// é dessa forma que é passado um valor para um channel
		requestId <- i
	}
}

// função que vai fazer a requisição
func worker(requestId chan int, worker int) {
	// o worker vai pegar os dados que estão sendo mandados pelo requestId e
	// utilizar no for
	for r := range requestId {
		res, err := http.Get("http://localhost:8585/product")
		if err != nil {
			log.Fatal("Erro")
		}
		// o defer executa o Close depois que tudo for executado nessa função
		defer res.Body.Close()
		content, _ := ioutil.ReadAll(res.Body)
		// vai imprimir o resultado
		fmt.Printf("Worker %d. RequestId: %d. Content: %s", worker, r, string(content))
		r := rand.Intn(5)
		time.Sleep(time.Duration(r) * time.Second)
	}

}
