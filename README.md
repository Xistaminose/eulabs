# EulabsApi

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)

## About <a name = "about"></a>

Projeto desenvolvido para o teste da empresa eulabs. O projeto consiste em uma API que realiza o CRUD de produtos utilizando o framework Echo.

## Getting Started <a name = "getting_started"></a>
Para rodar o projeto basta ter o docker em sua maquina, e rodar o comando `make run` na raiz do projeto.

Você pode alterar as configurações de banco de dados usando as variáveis de ambiente `DB_TYPE`, onde você pode escolher entre sqlite (default) e postgres. Caso escolha postgres, você deve preencher as variáveis `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT` e `DB_NAME`.