package main //Principal pacote da aplicação

import "fmt"
import "os"

//Pacote para fazer chamadas http em go
import "net/http"

// para compilar go build <nome_do_arquivo>.go
func main() {
	// Go n tem while. Se eu quiser fazer um loop infinito, eu posso usar o for {} (for sem complemento)
	for {
		exibeIntroducao()
		exibeMenu()
		comando := leComando()
		// O IF é  a mesma ideia do Java mas sem os ()
		// if comando == 1 {
		// 	fmt.Println("Monitorando")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")

		// } else if comando == 0 {
		// 	fmt.Println("Saindo do Programa")
		// } else {
		// 	fmt.Println("Não conheço o comando", comando)
		// }

		//Opção ao if (Não precisa de break a cada case)
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço o comando", comando)
			// Apenas boa pratica, informar para o Sistema Operacional se o programa terminou com erro (-1 com erro, 0 ok)
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	//Go infere o tipo, não preciso informar o tipo. É fortemente tipado
	var nome string = "Montival"
	var versao float32 = 1.1

	// A forma mais comum no Go é declarar a variavel ja atribuindo um valor. Nesse caso o var é dispoensável
	// nome := "Montival"

	// versao := 1.1

	fmt.Println("Ola Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comando int
	//O & significa que eu quero colocar o valor no endereço dessa variável
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variavel 'comando' é", &comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.google.com.br"
	// Funções em Go podem retornar mais de uma resposta. Para desprezar basta informar _ no lugar da variavel q a recebe
	response, err := http.Get(site)
	if err != nil {
		fmt.Print("Deu erro")
	}
	statusCode := response.StatusCode
	if statusCode == 200 {
		fmt.Println("Site: ", site, "Foi carregado com sucesso")

	} else {

		fmt.Println("Site: ", site, " está com problema. Status Code:", statusCode)
	}
}
