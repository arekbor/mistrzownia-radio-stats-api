package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arekbor/mistrzownia-radio-stats-api/types"
	"github.com/arekbor/mistrzownia-radio-stats-api/utils"
	_ "github.com/lib/pq"
)

const (
	statsDBTableName = "steamStats"
)

type storeConfig struct {
	conn string
}

type Store struct {
	Db *sql.DB
}

func New(host string, port string, user string, pwd string, dbname string, timeout time.Duration) *storeConfig {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbname)
	return &storeConfig{
		conn: conn,
	}
}

func (s *storeConfig) Init() (*Store, error) {
	db, err := sql.Open("postgres", s.conn)
	if err != nil {
		return nil, err
	}
	timeout, err := utils.GetMaxDbTimeout()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		Db: db,
	}, nil
}

func (s *Store) GetPaginatedStats(page int, limit int) ([]types.Stats, error) {
	timeout, err := utils.GetMaxDbTimeout()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	q := fmt.Sprintf(`SELECT * FROM "%s" 
	ORDER BY "DateTime" DESC
	LIMIT $1 OFFSET $2
	`, statsDBTableName)

	offset := limit * (page - 1)

	rows, err := s.Db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}

	slice := []types.Stats{}

	for rows.Next() {
		stats := &types.Stats{}
		err := rows.Scan(
			&stats.Id,
			&stats.SteamId,
			&stats.Username,
			&stats.AvatarURL,
			&stats.ProfileURL,
			&stats.YoutubeURL,
			&stats.YoutubeName,
			&stats.Datetime,
		)
		if err != nil {
			return nil, err
		}
		slice = append(slice, *stats)
	}

	return slice, nil
}

func (s *Store) GetCountOfStats() (int, error) {
	timeout, err := utils.GetMaxDbTimeout()
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	q := fmt.Sprintf(`select count(*) from "%s"`, statsDBTableName)

	rows, err := s.Db.QueryContext(ctx, q)
	if err != nil {
		return 0, err
	}
	var count int

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
