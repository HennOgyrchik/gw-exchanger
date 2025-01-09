package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type PSQL struct {
	timeout time.Duration
	url     string
	pool    *pgxpool.Pool
}

func New(url string, timeout time.Duration) PSQL {
	return PSQL{
		timeout: timeout,
		url:     url,
		pool:    nil,
	}
}

func (p *PSQL) Start(ctx context.Context) error {
	const op = "PSQL Start"

	ctxTimeout, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	pool, err := pgxpool.Connect(ctxTimeout, p.url)
	if err != nil {
		return fmt.Errorf("%s (Create pool connect): %w", op, err)
	}

	p.pool = pool

	_, err = p.pool.Exec(ctx, "create table if not exists exchange (id serial not null, currency varchar(5) unique not null, ratio float not null)")
	if err != nil {
		return fmt.Errorf("%s (Create table): %w", op, err)
	}
	return nil
}

func (p *PSQL) Stop() {
	if p.pool != nil {
		p.pool.Close()
	}
}
