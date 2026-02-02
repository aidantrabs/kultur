package service

import (
    "context"
    "errors"

    "github.com/aidantrabs/kultur/backend/internal/db"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

var ErrMemoryNotFound = errors.New("memory not found")

type MemoryService struct {
    queries         *db.Queries
    festivalService *FestivalService
}

func NewMemoryService(queries *db.Queries, festivalService *FestivalService) *MemoryService {
    return &MemoryService{
        queries:         queries,
        festivalService: festivalService,
    }
}

func (s *MemoryService) ListByFestivalSlug(ctx context.Context, slug string) ([]db.Memory, error) {
    festival, err := s.festivalService.GetBySlug(ctx, slug)
    if err != nil {
        return nil, err
    }

    return s.queries.ListMemoriesByFestival(ctx, festival.ID)
}

func (s *MemoryService) ListAll(ctx context.Context) ([]db.Memory, error) {
    return s.queries.ListAllMemories(ctx)
}

func (s *MemoryService) GetByID(ctx context.Context, id pgtype.UUID) (db.Memory, error) {
    memory, err := s.queries.GetMemoryByID(ctx, id)
    if errors.Is(err, pgx.ErrNoRows) {
        return db.Memory{}, ErrMemoryNotFound
    }

    return memory, err
}

type CreateMemoryParams struct {
    FestivalID   pgtype.UUID
    AuthorName   string
    AuthorEmail  string
    Content      string
    YearOfMemory string
}

func (s *MemoryService) Create(ctx context.Context, params CreateMemoryParams) (db.Memory, error) {
    return s.queries.CreateMemory(ctx, db.CreateMemoryParams{
        FestivalID:   params.FestivalID,
        AuthorName:   pgtype.Text{String: params.AuthorName, Valid: params.AuthorName != ""},
        AuthorEmail:  pgtype.Text{String: params.AuthorEmail, Valid: params.AuthorEmail != ""},
        Content:      params.Content,
        YearOfMemory: pgtype.Text{String: params.YearOfMemory, Valid: params.YearOfMemory != ""},
    })
}

func (s *MemoryService) UpdateStatus(ctx context.Context, id pgtype.UUID, status string) error {
    _, err := s.GetByID(ctx, id)
    if err != nil {
        return err
    }

    return s.queries.UpdateMemoryStatus(ctx, db.UpdateMemoryStatusParams{
        ID:     id,
        Status: pgtype.Text{String: status, Valid: true},
    })
}

func (s *MemoryService) Delete(ctx context.Context, id pgtype.UUID) error {
    return s.queries.DeleteMemory(ctx, id)
}
