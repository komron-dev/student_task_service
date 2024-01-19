package main

import (
	"net"

	"student/student_task_service/config"
	pb "student/student_task_service/genproto/task_service"
	"student/student_task_service/pkg/db"
	"student/student_task_service/pkg/logger"
	"student/student_task_service/services"
	"student/student_task_service/services/grpcClient"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)
//
func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "task-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)
	grpcClient, err := grpcClient.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}

	coonDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	taskService := services.NewTaskService(coonDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))

	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterTaskServiceServer(s, taskService)
	log.Info("main: server running",
		logger.String("port:", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}
