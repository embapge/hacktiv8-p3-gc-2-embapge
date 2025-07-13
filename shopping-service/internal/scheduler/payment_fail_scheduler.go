package scheduler

import (
	"context"
	"log"
	"time"

	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/repository"

	"github.com/go-co-op/gocron"
)

func StartPaymentFailScheduler(txRepo repository.TransactionRepository) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(func() {
		ctx := context.Background()
		cutoff := time.Now().Add(-3 * time.Hour)

		transactions, err := txRepo.FindAll(ctx)
		if err != nil {
			log.Println("[Scheduler] Error fetching transactions:", err)
			return
		}
		for _, tx := range transactions {
			if (tx.Status == "pending" || tx.Status == "success") && tx.CreatedAt.Before(cutoff) {
				tx.Status = "failed"
				tx.UpdatedAt = time.Now()
				_, err := txRepo.Update(ctx, tx.ID, &tx)
				if err != nil {
					log.Println("[Scheduler] Failed to update transaction status:", err)
				}
			}
		}
	})
	s.StartAsync()
}
