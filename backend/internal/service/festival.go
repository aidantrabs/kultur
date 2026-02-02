package service

import (
    "context"
    "errors"

    "github.com/aidantrabs/kultur/backend/internal/db"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
)

var (
    ErrFestivalNotFound     = errors.New("festival not found")
    ErrFestivalDateNotFound = errors.New("festival date not found")
)

type FestivalService struct {
    queries *db.Queries
}

func NewFestivalService(queries *db.Queries) *FestivalService {
    return &FestivalService{queries: queries}
}

type ListFestivalsParams struct {
    Region   string
    Heritage string
}

func (s *FestivalService) List(ctx context.Context, params ListFestivalsParams) ([]db.Festival, error) {
    switch {
    case params.Region != "":
        return s.queries.ListFestivalsByRegion(ctx, params.Region)
    case params.Heritage != "":
        return s.queries.ListFestivalsByHeritage(ctx, params.Heritage)
    default:
        return s.queries.ListFestivals(ctx)
    }
}

func (s *FestivalService) ListUpcoming(ctx context.Context) ([]db.ListUpcomingFestivalDatesRow, error) {
    return s.queries.ListUpcomingFestivalDates(ctx)
}

func (s *FestivalService) ListByYear(ctx context.Context, year int32) ([]db.ListFestivalDatesByYearRow, error) {
    return s.queries.ListFestivalDatesByYear(ctx, year)
}

func (s *FestivalService) GetBySlug(ctx context.Context, slug string) (db.Festival, error) {
    festival, err := s.queries.GetFestivalBySlug(ctx, slug)
    if errors.Is(err, pgx.ErrNoRows) {
        return db.Festival{}, ErrFestivalNotFound
    }

    return festival, err
}

func (s *FestivalService) GetByID(ctx context.Context, id pgtype.UUID) (db.Festival, error) {
    festival, err := s.queries.GetFestivalByID(ctx, id)
    if errors.Is(err, pgx.ErrNoRows) {
        return db.Festival{}, ErrFestivalNotFound
    }

    return festival, err
}

func (s *FestivalService) Create(ctx context.Context, params db.CreateFestivalParams) (db.Festival, error) {
    return s.queries.CreateFestival(ctx, params)
}

func (s *FestivalService) Update(ctx context.Context, params db.UpdateFestivalParams) (db.Festival, error) {
    festival, err := s.queries.UpdateFestival(ctx, params)
    if errors.Is(err, pgx.ErrNoRows) {
        return db.Festival{}, ErrFestivalNotFound
    }

    return festival, err
}

func (s *FestivalService) Delete(ctx context.Context, id pgtype.UUID) error {
    return s.queries.DeleteFestival(ctx, id)
}

func (s *FestivalService) GetDates(ctx context.Context, festivalID pgtype.UUID) ([]db.FestivalDate, error) {
    return s.queries.GetFestivalDatesByFestivalID(ctx, festivalID)
}

func (s *FestivalService) GetDateByYear(ctx context.Context, festivalID pgtype.UUID, year int32) (db.FestivalDate, error) {
    date, err := s.queries.GetFestivalDateByYear(ctx, db.GetFestivalDateByYearParams{
        FestivalID: festivalID,
        Year:       year,
    })
    if errors.Is(err, pgx.ErrNoRows) {
        return db.FestivalDate{}, ErrFestivalDateNotFound
    }

    return date, err
}

func (s *FestivalService) CreateDate(ctx context.Context, params db.CreateFestivalDateParams) (db.FestivalDate, error) {
    return s.queries.CreateFestivalDate(ctx, params)
}

func (s *FestivalService) UpdateDate(ctx context.Context, params db.UpdateFestivalDateParams) (db.FestivalDate, error) {
    date, err := s.queries.UpdateFestivalDate(ctx, params)
    if errors.Is(err, pgx.ErrNoRows) {
        return db.FestivalDate{}, ErrFestivalDateNotFound
    }

    return date, err
}

func (s *FestivalService) DeleteDate(ctx context.Context, id pgtype.UUID) error {
    return s.queries.DeleteFestivalDate(ctx, id)
}
