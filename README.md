```
docker volume create mysql-volume
docker run --name mysql-crud -p 3306:3306 -v mysql-volume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest
docker exec -it mysql-crud mysql -u root -p
```

```
go get github.com/go-sql-driver/mysql
```