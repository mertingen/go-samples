# go-samples
I'd like to share random apps in the spare times. Thus, I'm going to try learning some concepts of Go and as much as I can I try to clarify each line.

## CRUD APP
It basically handles CRUD operations in Go. I find the following topics important. Thus, I can refer to this app to get a reference for them.

- Dockerfiles, Containers and Images
- Migration
- Gorilla Mux Routing
- Gracefully Shutdown
- MySQL Join Relations
- Handlers, Services, Models

### Warm-up
First thing is creating a network for Docker with the following command.

```
docker create network samples-crud
```


There are 2 containers such as crud-app and mysql-database.

-- Run this command to start mysql db container.

```
docker run -d \
     --network samples-crud --network-alias mysql \
     -v todo-mysql-data:/var/lib/mysql \
     -e MYSQL_ROOT_PASSWORD=secret \
     -e MYSQL_DATABASE=crud \
     mysql:5.7
```

-- Build migrator image in /migrator directory by Dockerfile.
```
docker build -t migrator .
```

-- Run this command to start migration container to UP.
```
docker run --rm \
	--network samples-crud \
	migrator \
	--path=migrations/ \
	-database "mysql://root:secret@tcp(mysql)/crud" up
```

-- Run this command to start migration container to DOWN.
```
docker run --rm \
	--network samples-crud \
	migrator \
	--path=migrations/ \
	-database "mysql://root:secret@tcp(mysql)/crud" down --all
```

-- Build app image in root directory by Dockerfile.
```
docker build -t samples-crud .
```

-- Run this command to start app container.
```
 docker run -d -it \
     -p8000:8000 \
     --network samples-crud \
     -e MYSQL_HOST=mysql \
     -e MYSQL_USER=root \
     -e MYSQL_PASSWORD=secret \
     -e MYSQL_DB=crud \
     -e MYSQL_PORT=3306 \
     -e APP_PORT=8000 \
   samples-crud
 ```