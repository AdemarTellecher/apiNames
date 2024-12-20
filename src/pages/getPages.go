package pages

import (
	"io"
	"net/http"
	"regexp"
)

func GetNomes(pagina string) ([]string, error) {
	respPagina, err := getPagina(pagina)
	if err != nil {
		return nil, err
	}
	//Montando a expressão regular para captudar os dados desejados
	regexPagina := regexp.MustCompile(`<span class="list-wide--name full-w">(.+?)</span>`)

	//Montando a lista de combinações da resposta da REGEXP "Expressão Regular"
	matchesList := regexPagina.FindAllStringSubmatch(string(respPagina), -1)
	nomes := make([]string, 0, 1)
	for _, mChes := range matchesList {
		nomes = append(nomes, mChes[1])
	}
	return nomes, err

}

func getPagina(pagina string) ([]byte, error) {
	//Requisitando a pagina para analise...
	resp, err := http.Get("https://www.dicionariodenomesproprios.com.br/nomes-masculinos/" + pagina)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //Depois de terminar a função, fecha o body da resposta, liberando memoria

	//Montando o corpo da requisição na variavel "respBody"
	respBady, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return (respBady), nil

}
