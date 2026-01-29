# Desafio CEP API - Go

API para consulta de CEP em paralelo, integrando dados de **BrasilAPI** e **ViaCEP** de forma concorrente usando Go.

## 🚀 Descrição

Projeto desenvolvido em Go que implementa uma API RESTful para buscar informações de endereços através do código de endereçamento postal (CEP). A aplicação realiza requisições paralelas a múltiplas fontes de dados, demonstrando o uso de multithreading/concorrência em Go.

## 📋 Pré-requisitos

- Go 1.19 ou superior
- Git

## 🏃 Como Executar o Servidor


1. Instale as dependências
   ```bash
   go mod tidy
   ```

2. Vá até o caminho /cmd/server 
 ```
 cd cmd/server
 ```

3. Configure as variáveis de ambiente (crie um arquivo `.env`) no caminho /cmd/server
   ```env
   API_HOST=https://brasilapi.com.br/api/address/v1/public
   OTHER_API_HOST=https://viacep.com.br/ws
   ```

4. Execute o servidor
   ```bash
   go run main.go
   ```

O servidor estará disponível em `http://localhost:8000`

## 🧪 Como Rodar os Testes

Para executar todos os testes do projeto:

```bash
go test ./...
```

Para rodar testes de um pacote específico:

```bash
go test ./internal/infra/service/...
```

Para rodar testes com cobertura de código:

```bash
go test -cover ./...
```

## 📚 Como Abrir o Swagger

Com o servidor executando, acesse a documentação da API no seu navegador:

```
http://localhost:8000/docs/index.html
```

Lá você encontrará:
- ✅ Todos os endpoints disponíveis
- ✅ Modelos de requisição e resposta
- ✅ Exemplos de uso
- ✅ Possibilidade de testar os endpoints diretamente

## 🔌 Como Usar Extensão HTTP REST

### Usando a extensão REST Client

1. **Instale a extensão** no VS Code:
   - Procure por "REST Client" (publicada por Huachao Mao)
   - Ou execute: `ext install humao.rest-client`

2. **Use o arquivo** `test/cep.http` incluído no projeto:
   - Abra o arquivo `test/cep.http`
   - Clique em "Send Request" (ou use `Ctrl+Alt+R`)
   - Veja a resposta no painel de output

3. **Exemplo de requisição**:
   ```http
   GET http://localhost:8000/cep?cep=01001000 HTTP/1.1
   ```

## 📝 Endpoints Disponíveis

### Buscar CEP
```
GET /cep?cep=01001000
```

Retorna informações do endereço em formato JSON com dados de múltiplas fontes.


## 📞 Contato

Desenvolvido por Luana Andrade - luanaands@gmail.com

---

**Aproveite! 🚀**
