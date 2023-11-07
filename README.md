
<h1 align="center"><img src="/assets/img/graphql-reasons.png" width="600px"></h1>

<h3 align="center">Iniciando com GraphQl 🧠</h3>

<p align="center">
  <a href="#about">Sobre</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#install">Instalação</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#challenge">Desafios</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#technologies">Tecnologias</a>
</p>

## :speech_balloon: Sobre <a name="about"></a>

> Essa aplicação back-end com GraphQL desenvolvido durante o curso da Full Cycle, ele consiste em persistir os dados vindos do front-end no banco de dados.

## :warning: Instalação <a name="install"></a>

Realizar a criação da base de dados com o sqlite

```bash
$ sqlite3 data.db
```
Criar as tabelas:

```bash
$ create table categories ( id string, name string, description string );
$ create table courses ( id string, name string, description string, category_id string );
```

Subir o serviço 

```bash
$ go run cmd/server/server.go
```

Acessar pelo navegador localhost:8080
Crie as mutations para realizar o cadastro, criar primeiramente a categoria do curso e depois o curso, pois o curso precisa do id da categoria

```bash
mutation createCategory {
  createCategory(input: {name: "Tecnologia", description: "GoLang"}) {
    id
    name
    description
  }
}
```

```bash
mutation createCouse {
  createCourse(input: {name: "Arquitetura", description: "The best", categoryId: "7474c176-310a-4f09-84fc-b306feef3b84"}) {
    id
    name
  }
}
```

Criar as querys resolvers para buscar informações.

```bash
query queryCategories {
  categories {
    id
    name
    description
  }
}
```

```bash
query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
    }
  }
}
```

```bash
query queryCourses {
  courses {
    id
    name
  }
}
```

```bash
query queryCoursesWithCategory {
  courses {
    id
    name
    description
    category {
      id 
      name
      description
    }
  }
}
```

## :triangular_flag_on_post: Desafio <a name="challenge"></a>

> Foi o meu primeiro contato com GraphQL.

## :heavy_check_mark: Tecnologias <a name="technologies"></a>

- [GO](https://go.dev/)
- [GraphQL](https://graphql.org/)
---

by [Gisele Trizotte](https://www.github.com/giseleTrizotte) ❤️



