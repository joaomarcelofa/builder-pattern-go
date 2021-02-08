# builder-pattern-go

O intuito desse código é criar MacBook's utilizando as configurações descritas pelo site da Apple:

* [MacBook Air](https://www.apple.com/shop/buy-mac/macbook-air)
* [MacBook Pro 13"](https://www.apple.com/shop/buy-mac/macbook-pro/13-inch)
* [MacBook Pro 16"](https://www.apple.com/shop/buy-mac/macbook-pro/16-inch)

**Nota**: (Essas configurações foram acessadas no dia 04/02/2021)

A construção dos MacBook's é feita a partir da implementação do design pattern **Builder**.

---

## Rodando o projeto

Com a variável `GO111MODULE=on` (go modules ligado):

```shell
$ go run main.go
```

Com a variável `GO111MODULE=off` (go modules desligado):

```shell
$ GO111MODULE=on go run main.go
```