### Desafio

Considere uma senha sendo válida quando a mesma possuir as seguintes definições:

- Nove ou mais caracteres
- Ao menos 1 dígito
- Ao menos 1 letra minúscula
- Ao menos 1 letra maiúscula
- Ao menos 1 caractere especial
 - Considere como especial os seguintes caracteres: !@#$%^&*()-+
- Não possuir caracteres repetidos dentro do conjunto

---

- Input: uma senha (string)
- Output: um boolean indicando se a senha é válida

### Como executar

1. Na raiz da aplicação onde contém o arquivo `makefile`
1. Execute o comando `make build` para construir e em seguida `make run` para iniciar o container Docker contendo a aplicação. Como alternativa execute o comando `make up` para construir e iniciar a aplicação
1. O serviço é exposto na porta `8080` e amarrado na porta 8080 da máquina local.
1. Para executar os testes, com o ambiente Go instalado, execute o comando `make tests` - atualmente 66.2% da base coberta por testes (100% dos handlers e validators testados).
1. Para testar, fazer uma requisção utilizado `curl` -> `curl -v -X POST localhost:8080/v1/validate -d '{"password": "AbTp9!fok"}' | jq`
1. O comando `make stop` para a execução a aplicação.

### Abordagem

Tecnologia utilizadas:
- go
- gin - go web framework
- zap - uber's go logging package
- docker

A abordagem principal é a criação de um package em go (validation/validators/...) contendo os critérios de validação de senha. Os critérios podem ser utilizados de forma única ou em conjunto, utilizando um padrão semelhante como os middlewares são utilizados em go, ou o padrão Chain of Responsibility. Dentro do package cada arquivo implementa uma validação, onde, se necessário, a modificação ou extensão com mais validadores, a extensão acontece em único ponto.

O endpoint é versionado e declarado como v1/validate, e validação é feita encaminhando no corpo do payload um JSON contendo o campo `password` do tipo string; a aplicação responde com um JSON contendo o campo `is_valid` do tipo boolean, sendo `true` caso a senha atenda os critérios ou `false` caso contrário.

Para a critérios de validação de tipo de caractere, foi utilizado mapas (validation/validators/validators.go) em golang, que se comporta como uma hashtable, o qual contém os caracteres válidos, esses mapas são inicializados junto a aplicação. Caso não indique no código quais validadores utilizar, a aplicação inicia utilizando o DefaultValidator que emprega os critérios definidos pelo desafio.
