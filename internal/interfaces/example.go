package interfaces

import (
	"context"
	"log/slog"
	"time"
)

/** This is an example on how to use interfaces */

// Interfaces
type (
	AccountNotifier interface {
		// here we added context.Context so that we have the control on when to terminate
		// suppose that the notify using slack takes longer that 10 sec then we can cancel the context making the
		// function not to wait indefinitely
		NotifyAccountCreated(context.Context, Account) error
	}
)

// Structs
type (
	Account struct {
		Username string
		Email    string
	}
	SimpleAccountNotifier struct{}
	BetterAccountNotifier struct {
		timeToDeliver time.Duration
	}
	AccountHandler struct {
		AccountNotifier AccountNotifier
	}
)

func (s *SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("Simple Account Notifier: Account created successfully", "account", account)
	return nil
}

// func (b *BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
// 	slog.Info("Simple Account Notifier: Account created successfully", "account", account)
// 	return nil
// }

func (b *BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	// Simulate long-running external call
	select {
	case <-ctx.Done():
		slog.Error("Better Account Notifier: Notification cancelled", "error", ctx.Err())
		return ctx.Err()
	case <-time.After(b.timeToDeliver): // simulate API call taking 2s
		slog.Info("Better Account Notifier: Sending welcome email...",
			"username", account.Username,
			"email", account.Email,
		)
	}

	slog.Info("Better Account Notifier: Welcome email sent successfully", "account", account)
	return nil
}

func (h *AccountHandler) handleCreateAccount(ctx context.Context, account Account) {
	h.AccountNotifier.NotifyAccountCreated(ctx, account)
}

func ExampleMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	accountHandler := &AccountHandler{
		// Here we can change the notifer behaviour by Changing it to BetterAccountNotifier
		// AccountNotifier: &SimpleAccountNotifier{},
		AccountNotifier: &BetterAccountNotifier{
			// Play with context time and delivery time to see the magic
			timeToDeliver: 10 * time.Second,
		},
	}
	account := Account{
		Username: "Maneesh",
		Email:    "abcd@gmail.com",
	}
	accountHandler.handleCreateAccount(ctx, account)
}
