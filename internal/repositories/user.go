package repositories

import (
	"context"
	// "database/sql"
	"time"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bartmika/growlog-server/internal/models"
)

type UserRepo struct {
	dbpool             *pgxpool.Pool
}

func NewUserRepo(dbpool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		dbpool: dbpool,
	}
}

func (r *UserRepo) Insert(ctx context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    INSERT INTO users (
        uuid, tenant_id, email, first_name, last_name, password_algorithm, password_hash, state,
		role_id, timezone, created_time, modified_time, joined_time, salt, was_email_activated,
		pr_access_code, pr_expiry_time, old_id
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
    )`

	err := r.dbpool.QueryRow(ctx, query).Scan(&m)
	if err != nil {
		log.Println("UserRepo|Insert|err", err)
		return err
	}
	return nil
}

func (r *UserRepo) UpdateById(ctx context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE
        users
    SET
        tenant_id = $1, email = $2, first_name = $3, last_name = $4, password_algorithm = $5, password_hash = $6, state = $7,
		role_id = $8, timezone = $9, created_time = $10, modified_time = $11, joined_time = $12, salt = $13, was_email_activated = $14,
		pr_access_code = $15, pr_expiry_time = $16
    WHERE
        id = $17`

	err := r.dbpool.QueryRow(ctx, query).Scan(&m)
	if err != nil {
		log.Println("UserRepo|UpdateById|err", err)
		return err
	}
	return nil
}
//
func (r *UserRepo) UpdateByEmail(ctx context.Context, m *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE
        users
    SET
        tenant_id = $1, email = $2, first_name = $3, last_name = $4, password_algorithm = $5, password_hash = $6, state = $7,
		role_id = $8, timezone = $9, created_time = $10, modified_time = $11, joined_time = $12, salt = $13, was_email_activated = $14,
		pr_access_code = $15, pr_expiry_time = $16
    WHERE
        email = $2`

	err := r.dbpool.QueryRow(ctx, query).Scan(&m)
	if err != nil {
		log.Println("UserRepo|UpdateByEmail|err", err)
		return err
	}
	return nil
}

func (r *UserRepo) GetById(ctx context.Context, id uint64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := `
    SELECT
        id, uuid, tenant_id, email, first_name, last_name, password_algorithm, password_hash, state,
		role_id, timezone, created_time, modified_time, joined_time, salt, was_email_activated, pr_access_code, pr_expiry_time
    FROM
        users
    WHERE
        id = $1`

	err := r.dbpool.QueryRow(ctx, query).Scan(&m)
	if err != nil {
		log.Println("UserRepo|GetById|err", err)
		return nil, err
	}
	return m, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := `
    SELECT
        id, uuid, tenant_id, email, first_name, last_name, password_algorithm, password_hash, state,
		role_id, timezone, created_time, modified_time, joined_time, salt, was_email_activated, pr_access_code, pr_expiry_time
    FROM
        users
    WHERE
        email = $1`

	err := r.dbpool.QueryRow(ctx, query).Scan(&m)
	if err != nil {
		log.Println("UserRepo|GetByEmail|err", err)
		return nil, err
	}
	return m, nil
}

func (r *UserRepo) CheckIfExistsById(ctx context.Context, id uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool

	query := `
    SELECT
        1
    FROM
        users
    WHERE
        id = $1`

	err := r.dbpool.QueryRow(ctx, query).Scan(&exists)
	if err != nil {
		log.Println("UserRepo|CheckIfExistsById|err", err)
		return false, err
	}
	return exists, nil
}

func (r *UserRepo) CheckIfExistsByEmail(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool

	query := `
    SELECT
        1
    FROM
        users
    WHERE
        email = $1`

	err := r.dbpool.QueryRow(ctx, query).Scan(&exists)
	if err != nil {
		log.Println("UserRepo|CheckIfExistsByEmail|err", err)
		return false, err
	}
	return exists, nil
}

func (r *UserRepo) InsertOrUpdateById(ctx context.Context, m *models.User) error {
	if m.Id == 0 {
		return r.Insert(ctx, m)
	}

	doesExist, err := r.CheckIfExistsById(ctx, m.Id)
	if err != nil {
		return err
	}

	if doesExist == false {
		return r.Insert(ctx, m)
	}
	return r.UpdateById(ctx, m)
}

func (r *UserRepo) InsertOrUpdateByEmail(ctx context.Context, m *models.User) error {
	if m.Id == 0 {
		return r.Insert(ctx, m)
	}

	doesExist, err := r.CheckIfExistsByEmail(ctx, m.Email)
	if err != nil {
		return err
	}

	if doesExist == false {
		return r.Insert(ctx, m)
	}
	return r.UpdateByEmail(ctx, m)
}
