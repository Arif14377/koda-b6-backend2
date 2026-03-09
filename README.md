# Latihan implementasi fitur dan backend Project Coffee Shop

## Ringkasan implementasi
1. Dependency Injection (CRUD user dan product)
2. Koneksi database (bukan tugas, tapi kebutuhan)
3. Redis

## Koneksi Database (PostgreSQL)
1. 

## Redis
1. Pull image redis container
```bash
docker pull redis:8.6.1-trixie
```
2. Run Container Redis (with port forward & restart always)
```bash
docker run -p 6379:6379 -d --restart always redis:8.6.1-trixie
```
3. Install extension Redis for Vscode
4. Connect database melalui extension Redis for VsCode (tidak perlu setting apapun)
5. Install (package) go redis untuk komunikasi go dan redis
```bash
go get -u github.com/redis/go-redis/v9
```
6. Install godotenv untuk loads env variables dari file .env
```bash
go get -u github.com/joho/godotenv
```
7. Konfigurasi HOST dan PORT redis di file .env
```
REDIS_HOST="localhost"
REDIS_PORT="6379"
```
8. Buat NewClient Redis di main.go untuk quick start / di /di/container.go
```go
rdb := redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
    DB:       0, //default DB
    Password: "",
})
defer rdb.Close()
```

9. 