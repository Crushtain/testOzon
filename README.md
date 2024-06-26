# OZON

## Запуск 
### Обычный запуск
```
go run .\cmd\app\server.go
```
### С использованием postgres вместо inMemory
```
go run .\cmd\app\server.go --s postgres
```

### Примеры запросов, мутаций и подписок:
#### Создать пост:
```graphql
mutation {
  createPublication(title: "...", content: "...", commentsDisabled: false) {
    id
    title
    content
    commentsDisabled
  }
}
```
#### Создать комментарий:
```graphql
mutation {
  createComment(author: "...", publicationId: "...", parentId: null, content: "...") {
    id
    publicationId
    parentId
    content
  }
}
```
#### Получить пост по ID:
```graphql
query {
  publication(id: "...") {
    id
    title
    content
    comments {
      id
      author
      content
    }
    commentsDisabled
  }
}
```

#### Получить все посты
```graphql
query {
    publications {
        id
        title
        content
        commentsDisabled
    }
}

```

#### Получить по ID поста его комментарии
```graphql
query {
    comments(publicationId:"...",
        limit:10, offset:0)
    {
        id
        author
        publicationId
        content
    }
}

```
#### Подписка на комментарии к посту
```graphql
subscription {
    commentNotification(publicationId: "...") {
        id
        author
        publicationId
        content
        parentId
    }
}
```