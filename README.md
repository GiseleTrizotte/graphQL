
Criar o banco de dados 
sqlite3 data.db

Criar as tabelas:
create table categories ( id string, name string, description string );
create table courses ( id string, name string, description string, category_id string );

Voce pode cnsultar se a table foi criada
select * from categories;

go run cmd/server/server.go

Acessar localhost:8080
criando as mutations para realizar o cadastro, primeiro deve criar a categia do curso e depois o curso, pois o curso precisa passar o id de uma categoria.

mutation createCategory {
  createCategory(input: {name: "Tecnologia", description: "teste"}) {
    id
    name
    description
  }
}

mutation createCouse {
  createCourse(input: {name: "Full Cycle", description: "The best", categoryId: "7474c176-310a-4f09-84fc-b306feef3b84"}) {
    id
    name
  }
}


criar as querys resolvers para buscar as informacoes.


