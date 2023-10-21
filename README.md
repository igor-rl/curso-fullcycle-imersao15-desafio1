Iniciar o projeto
```yml
go run main.go
```

acessar o serviço
```yml
package github.com.igorlage.curso.fullcycle.desafio1;
```


```yml
docker compose up && docker exec -it desafio1-app-1 bash && go run main.go
```

em uma nova aba
```yml
docker exec -it desafio1-app-1  bash && evans -r repl
```



### Informações do desafio

Dado o seguinte protofile gRPC:
https://gist.github.com/argentinaluiz/8a2e1a7a32da3d1b107d88ec9f1007b2


Crie um servidor gRPC em Golang que realize as duas operações:


> Criar produtos <br>
> Consultar produtos

A aplicação deve usar o SQLite como banco de dados e o GORM como ORM. Use o modo AutoMigrate para gerar as tabelas automaticamente.


Você deve disponibilizar um o script main.go que a rodar go run main.go deve levantar o servidor e deixar a porta 50051 acessível via localhost.

### Gerar o binário do .proto
```yml
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```