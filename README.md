# Full Cycle 3.0

## Event Driven Architecture

Este repositório foi criado para cumprir o desafio da trilha **Event Driven Architecture** do curso **Full Cycle 3.0**

## Executando

Para executar este projeto, acesse a pasta raiz do projeto através do terminal e então execute:

```sh
$ docker compose up -d
```

Na inicialização, serão criadas duas instâncias do MySQL, cada um com um banco de dados:
* `walletcore`: Persiste os dados de clientes e contas bancárias.
* `transactions`: Persiste os dados de transações bancárias.

O banco de dados `walletcore` será previamente preenchidos com alguns dados de clientes e contas bancárias.

## Consultando clientes

Para consultar os clientes disponíveis, execute a requisição denominada **listCustomers** no arquivo `api.http`, localizado na pasta `api` relativa à raiz do projeto. A resposta dessa requisição será um documento JSON com a lista de todos os clientes disponíveis.

```json
{
  "customers": [
    {
      "id": "a9ef943d-39e7-11ee-b8f7-0242ac120004",
      "name": "Josimar Zimermann",
      "email": "josimarz@yahoo.com.br",
      "createdAt": "2023-08-13T14:42:44Z",
      "updatedAt": "2023-08-13T14:42:44Z"
    },
    {
      "id": "a9f1a411-39e7-11ee-b8f7-0242ac120004",
      "name": "Gustavo Kuerten",
      "email": "guga@tennis.com",
      "createdAt": "2023-08-13T14:42:44Z",
      "updatedAt": "2023-08-13T14:42:44Z"
    },
    {
      "id": "a9f30a18-39e7-11ee-b8f7-0242ac120004",
      "name": "Ana Ivanovic",
      "email": "ivanovic@wta.com",
      "createdAt": "2023-08-13T14:42:44Z",
      "updatedAt": "2023-08-13T14:42:44Z"
    },
    {
      "id": "a9f493fd-39e7-11ee-b8f7-0242ac120004",
      "name": "Maria Sharapova",
      "email": "sharapova@wta.com",
      "createdAt": "2023-08-13T14:42:44Z",
      "updatedAt": "2023-08-13T14:42:44Z"
    }
  ]
}
```

Copie o `id` do cliente desejado para consultar as contas desse cliente.

## Consultando contas do cliente

Ainda no arquivo `api.http`, utilize a requisição denominada `listCustomerAccounts` para consultar as contas vinculadas com um determinado cliente. Na URL da requisição, substitua o `id` de exemplo pelo `id` do cliente desejado. A resposta da requisição será uma lista com as contas vinculadas ao cliente.

```json
{
  "accounts": [
    {
      "id": "a9f80299-39e7-11ee-b8f7-0242ac120004",
      "balance": 500,
      "createdAt": "2023-08-13T14:42:44Z",
      "updatedAt": "2023-08-13T14:42:44Z"
    }
  ]
}
```

## Realizando transações

Para efetuar uma transação, no arquivo `api.http` procure pela requisição denominada `createTransaction`. No corpo da requisição atribua para o campo `from` o `id` da conta de origem, isto é, a conta da qual o valor será debitado. Para o campo `to`, atribua o `id` da conta de destino, isto é, a conta na qual o valor será creditado. No campo `amount`, informe o valor da operação. A resposta da requisição exibirá a condição atualizada das contas envolvidas na operação.

```json
{
  "from": {
    "id": "a9f3fd64-39e7-11ee-b8f7-0242ac120004",
    "balance": 4000,
    "createdAt": "2023-08-13T14:42:44Z",
    "updatedAt": "2023-08-13T15:08:59.208585952Z"
  },
  "to": {
    "id": "a9f80299-39e7-11ee-b8f7-0242ac120004",
    "balance": 1500,
    "createdAt": "2023-08-13T14:42:44Z",
    "updatedAt": "2023-08-13T15:08:59.208585874Z"
  },
  "amount": 1000
}
```

Sempre que uma transação é criada, uma nova mensagem é enviada para o Apache Kafka. As mensagens são consumidas pelo serviço WalletCore, responsável pela atualização do balanço das contas envolvidas na transação.

## Consultando o balanço das contas

Para consultar o balanço atualizado das contas envolvidas na transação, utilize a requisição denominada `showAccountBalance`, disponível no arquivo `api.http`. A resposta da requisição será um documento JSON exibindo a condição atual da conta.