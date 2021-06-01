# Sobre

Essa implementação tem como objetivo fixar conhecimento adquirido durante o módulo de Arquitetura Hexagonal da Formação Full Cycle by Code.Education.
Dessa forma, será demonstrado a utilização de dois adaptadores para consumo de um serviço de cadastro de produtos, sendo o primeiro via CLI e o segundo via Web Server.

---

## Dicas

Geração das classe mock
```
mockgen -destination=application/mocks/application.go -source=application/product.go
```

Geração CLI com Cobra
```
cobra init --pkg-name github.com/roaires/fullcycle-arquitetura-hexagonal-golang
```

Execução dos testes
```
go test ./...
```

---

## Exemplo de utilização via CLI

```
root@3a31c132639a:/go/src# go run main.go cli -a=create -n="Produto via CLI" -p1.99
Prouto ac4c4ecf-990c-4cbb-bdb5-38da62416a86 - Produto via CLI criado com sucesso. 
Preço: 1.990000
Status: Disable

root@3a31c132639a:/go/src# go run main.go cli -a=get -i=ac4c4ecf-990c-4cbb-bdb5-38da62416a86
Informações do produto:
ID: ac4c4ecf-990c-4cbb-bdb5-38da62416a86
Name: Produto via CLI
Price: 1.990000
Status: Disable
```

---

## Exemplo utilizando Web Server

Inicalizar Web Server
```
go run main.go http
```

Consultar produto cadastrado via CLI
```
root@3a31c132639a:/go/src# curl http://localhost:9000/products/ac4c4ecf-990c-4cbb-bdb5-38da62416a86
{"ID":"ac4c4ecf-990c-4cbb-bdb5-38da62416a86","Name":"Produto via CLI","Price":1.99,"Status":"Disable"}
```

Criação de um novo produto e consulta do produto cadastrado via Web Server
```
root@3a31c132639a:/go/src# curl --header "Content-Type: application/json" \
>   --request POST \
>   --data '{"name":"Produto via POST","price":2.99}' \
>   http://localhost:9000/products
{"ID":"6ea015d6-8748-41b8-ab93-599057b6e99c","Name":"Produto via POST","Price":2.99,"Status":"Disable"}


root@3a31c132639a:/go/src# curl http://localhost:9000/products/6ea015d6-8748-41b8-ab93-599057b6e99c
{"ID":"6ea015d6-8748-41b8-ab93-599057b6e99c","Name":"Produto via POST","Price":2.99,"Status":"Disable"}
```
---
Melhorias futuras
- Refactor para remover códigos duplicados

---