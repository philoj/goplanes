A 2D game experiment using WebSockets and WebAssembly. 

# Goose commands

Check status
```shell
goose postgres "user=root dbname=goplanes sslmode=disable host=localhost" -dir=server/migrations status
```

Add migration:
```shell
goose -dir=server/migrations create <name> sql
```


# Preview

https://user-images.githubusercontent.com/13179046/198898565-b7c96ac9-24f2-4bb5-bf1c-45f057edb7ef.mov