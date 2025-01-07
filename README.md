Documentação para execução do projeto

 ## Esse foi meu primeiro projeto na linguagem GoLang, linguagem está que estou ainda em processo de aprendizagem, mas queria comentar quanto foi desafiador realizar a execução desse projeto, e também o quanto foi gratificante conseguir finalizar ele e entregar dentro do prazo estipulado. Abaixo irei listar as funcionalidades da API e orientar a forma de execução dela. E a escolha do PostgreSQL foi mais pela questão da facilidade, já que é uma ferramente que tenho maior dominio que o MongoDB.

 ## Funcionalidades da API

- **CreatBook**: Função reponsável por: Adicionar um novo livro ao banco de dados, onde o mesmo terá um ID associado a ele, seguindo por: Título, autor, categoria e sinopse.
- **GetBook**: Função reponsável por: Consultar informações de um livro pelo seu ID.
- **GetAllBooks**: Função reponsável por: Mostar uma lista de todos os livros no bando de dados.
- **UpdateBook**: Função reponsável por: Atualizar informações de um livro em especifico.
- **DeleteBook**: Função reponsável por: Remover um livro do banco de dados.

 ## Como Executar a API
 **1. Configuração do Banco de Dados**:

 - Crie um banco de dados no PostgreSQL, dê o nome (meu_banco) para esse banco de dados, o username deve ser: postgres e a senha: 210420

 - Crie a tabela books no PostgreSQL usando o script a seguir:
 - CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    synopsis TEXT );
   

 **2. Instalar Dependências:** 

 - Execute o comando abaixo para instalar as dependências:
         go mod tidy 


 **3. Rodar a Aplicação:** 

 - Execute o comando abaixo para rodar o servidor:
        go run main.go 

  - A API estará disponível em http://localhost:8080/.
  - Recomendo que use o Postman para testar as rotas.
- POST /books: Cria um novo livro.
- GET /books: Lista todos os livros:
- GET /books/id: Mostra os detalhes de um livro específico:
- PUT /books/id: Atualiza as informações de um livro:
- DELETE /books/id: Remove um livro:
    
