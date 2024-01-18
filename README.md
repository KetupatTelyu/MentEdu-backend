# MentEdu

## Tim
Nama Tim: Ketupat Team

Hacker: Rahman Hakim

### Api Collection
![img.png](assets/img.png)

## How to run
### 1. Install dependencies
```bash
go mod download
go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate
```

### 2. Setup Database

### 3. Put your database url in .env file

### 4. Run migrations
```bash
make migrate_up DATABASE_URL=<database_url>
```

### 5. Setup Necessary Directories
```bash
make setup
```

### 6. Generate Keys
```bash
make generate-keys
```

### 7. Setup Google Cloud Storage

### 8. Put your google cloud storage credentials in .env file

### 9. Seed Database
```bash
make seed DATABASE_URL=<database_url>
```

### 10. Build
```bash
make compile-server
```

### 11. Run
```bash
./bin/server/main
```