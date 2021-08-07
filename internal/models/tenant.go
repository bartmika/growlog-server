package models

import (
	"context"
	"time"
)

type Tenant struct {
	Id                      uint64    `json:"id"`
	Uuid                    string    `json:"uuid"`
	Name                    string    `json:"name"`
	Url                     string    `json:"url"`
	State                   int8      `json:"state"`
	Timezone                string    `json:"timestamp"`
	CreatedTime             time.Time `json:"created_time"`
	ModifiedTime            time.Time `json:"modified_time"`
}

type TenantRepository interface {
	Insert(ctx context.Context, u *Tenant) error
	UpdateById(ctx context.Context, u *Tenant) error
	GetById(ctx context.Context, id uint64) (*Tenant, error)
	CheckIfExistsById(ctx context.Context, id uint64) (bool, error)
	CheckIfExistsByName(ctx context.Context, name string) (bool, error)
	InsertOrUpdateById(ctx context.Context, u *Tenant) error
}
