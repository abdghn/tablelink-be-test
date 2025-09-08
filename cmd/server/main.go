package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"tablelink-be-test/internal/application/usecase"
	"tablelink-be-test/internal/infrastructure/cache"
	"tablelink-be-test/internal/infrastructure/database"
	grpcHandler "tablelink-be-test/internal/infrastructure/grpc"
	"tablelink-be-test/internal/infrastructure/repository"
	authpb "tablelink-be-test/proto/auth"
	userspb "tablelink-be-test/proto/users"
)

func main() {
	db, err := database.NewPostgresDB("localhost", "5432", "promac", "postgres", "tablelink_db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	if err := db.CreateTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	redisCache := cache.NewRedisCache("localhost:6379", "", 0)

	userRepo := repository.NewUserRepositoryImpl(db)
	roleRepo := repository.NewRoleRepositoryImpl(db)
	roleRightRepo := repository.NewRoleRightRepositoryImpl(db)

	authUsecase := usecase.NewAuthUsecaseImpl(userRepo, roleRepo, roleRightRepo, redisCache)
	userUsecase := usecase.NewUserUsecaseImpl(userRepo, roleRepo)

	authHandler := grpcHandler.NewAuthHandler(authUsecase)
	userHandler := grpcHandler.NewUserHandler(userUsecase, authUsecase)

	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, authHandler)
	userspb.RegisterUserServiceServer(s, userHandler)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
