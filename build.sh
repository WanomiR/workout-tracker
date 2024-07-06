cd backend
swag init -g ./cmd/api/main.go
cd ..
docker-compose up --force-recreate --build