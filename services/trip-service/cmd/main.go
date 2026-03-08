package main

import (
	"context"
	"log"
	"time"

	"hopon.io/services/trip-service/internal/domain"
	"hopon.io/services/trip-service/internal/infrastructure/repository"
	"hopon.io/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	fare := &domain.RideFareModel{
		UserID: "42",
	}

	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	//To-be-removed
	for {
		time.Sleep(time.Second)
	}
}
