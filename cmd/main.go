package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	cep := "01519-000"

	go func() {
		cepResult := getApicep(cep)
		c1 <- cepResult

	}()

	go func() {
		cepResult := getViaCep(cep)
		c2 <- cepResult
	}()

	select {
	case cepResult := <-c1:
		fmt.Println(cepResult, "Apicep")

	case cepResult := <-c2:
		fmt.Println(cepResult, "Viacep")

	case <-time.After(time.Second * 1):
		fmt.Println("timeout")

	}

}

func getApicep(cep string) string {
	c := http.Client{}

	resp, err := c.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(bodyBytes)

}

func getViaCep(cep string) string {
	c := http.Client{}

	resp, err := c.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(bodyBytes)

}
