package service

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "errors"

    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
    "github.com/aidantrabs/trinbago-hackathon/backend/internal/email"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

var (
    ErrSubscriptionNotFound = errors.New("subscription not found")
    ErrInvalidToken         = errors.New("invalid token")
)

type SubscriptionService struct {
    queries *db.Queries
    email   *email.Service
}

func NewSubscriptionService(queries *db.Queries, emailSvc *email.Service) *SubscriptionService {
    return &SubscriptionService{
        queries: queries,
        email:   emailSvc,
    }
}

type CreateSubscriptionParams struct {
    Email        string
    DigestWeekly bool
}

func (s *SubscriptionService) Create(ctx context.Context, params CreateSubscriptionParams) (db.Subscription, error) {
    confirmToken, err := generateToken()
    if err != nil {
        return db.Subscription{}, err
    }

    unsubToken, err := generateToken()
    if err != nil {
        return db.Subscription{}, err
    }

    sub, err := s.queries.CreateSubscription(ctx, db.CreateSubscriptionParams{
        Email:             params.Email,
        DigestWeekly:      pgtype.Bool{Bool: params.DigestWeekly, Valid: true},
        FestivalReminders: []byte("[]"),
        ConfirmationToken: pgtype.Text{String: confirmToken, Valid: true},
        UnsubscribeToken:  unsubToken,
    })
    if err != nil {
        return db.Subscription{}, err
    }

    if err := s.email.SendConfirmation(params.Email, confirmToken); err != nil {
        // log error but don't fail the request
    }

    return sub, nil
}

func (s *SubscriptionService) Confirm(ctx context.Context, token string) error {
    sub, err := s.queries.GetSubscriptionByConfirmationToken(ctx, pgtype.Text{String: token, Valid: true})
    if errors.Is(err, pgx.ErrNoRows) {
        return ErrInvalidToken
    }
    if err != nil {
        return err
    }

    if err := s.queries.ConfirmSubscription(ctx, sub.ID); err != nil {
        return err
    }

    if err := s.email.SendWelcome(sub.Email, sub.UnsubscribeToken); err != nil {
        // log error but don't fail the request
    }

    return nil
}

func (s *SubscriptionService) Unsubscribe(ctx context.Context, token string) error {
    sub, err := s.queries.GetSubscriptionByUnsubscribeToken(ctx, token)
    if errors.Is(err, pgx.ErrNoRows) {
        return ErrInvalidToken
    }
    if err != nil {
        return err
    }

    return s.queries.DeleteSubscription(ctx, sub.ID)
}

func (s *SubscriptionService) ListAll(ctx context.Context) ([]db.Subscription, error) {
    return s.queries.ListAllSubscriptions(ctx)
}

func (s *SubscriptionService) Delete(ctx context.Context, id pgtype.UUID) error {
    return s.queries.DeleteSubscription(ctx, id)
}

func generateToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}
