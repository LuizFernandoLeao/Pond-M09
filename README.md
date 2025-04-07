# Projeto de Aplicação Web

Este projeto é uma aplicação web modularizada que utiliza contêineres para uma implantação segura, por meio do Docker. Apresenta em um front-end construído com React, um back-end desenvolvido em GoLang usando o framework Gin, um banco de dados PostgreSQL e Nginx como servidor web. 

## Estrutura do projeto

```
pondM09
├── docker-compose.yml
├── README.md
├── frontend
│   ├── Dockerfile
│   ├── package.json
│   ├── public
│   │   ├── index.html
│   │   └── manifest.json
│   ├── src
│   │   ├── App.js
│   │   ├── components
│   │   │   ├── UserForm.js
│   │   │   └── UserList.js
│   │   ├── index.js
│   │   └── services
│   │       └── api.js
│   └── webpack.config.js
├── backend
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── handlers
│   │   ├── users.go
│   │   └── products.go
│   ├── models
│   │   ├── user.go
│   │   └── product.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│   │   └── db.go
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
└── db
    ├── Dockerfile
    └── init.sql
```

### Pré-requisitos

- Docker e Docker Compose instalados em sua máquina.

### Instruções de Configuração

1. **Clone o Repositório**
   ```bash
   git clone <repository-url>
   cd pondM09
   ```

2. **Construa e Execute a Aplicação**
   Use o Docker Compose para construir e executar a aplicação:
   ```bash
   docker-compose up --build
   ```

3. **Acesse a Aplicação**
   Uma vez que os contêineres estejam rodando, você pode acessar a aplicação no seu navegador web em `http://localhost:3000`.

### Funcionalidades

- **Gerenciamento de Usuários**: Registrar, autenticar, listar, editar e deletar usuários.
- **Gerenciamento de Produtos**: Registrar, listar, editar e deletar produtos.
- **Armazenamento de Imagens**: Suporta uploads de imagens para perfis de usuário e produtos.

### Considerações de Segurança

- A aplicação é configurada para rodar em uma rede dedicada, expondo apenas as portas necessárias.

## Licença

Este projeto é licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.