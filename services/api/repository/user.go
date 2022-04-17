package repository

import (
	"encoding/json"
	"sync"

	"github.com/rafaelbreno/nuveo-test/entity"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/rafaelbreno/nuveo-test/services/api/storage"
	"gorm.io/gorm"
)

type (
	// UserRepoI handles all DB actions
	// related to User entity
	UserRepoI interface {
		Create(u entity.User) (entity.User, error)
		Update(id string, u entity.User) (entity.User, error)
		Read(id string) (entity.User, error)
		ReadAll() ([]entity.User, error)
		Delete(id string) error
	}

	// UserRepo handles all DB actions
	// related to User entity
	UserRepo struct {
		st    *storage.Storage
		in    *internal.Internal
		queue *queue.Queue
		mu    sync.Mutex
	}
)

// NewUserRepo returns an instance of UserRepo
// given Storage, Internal and Queue.
func NewUserRepo(st *storage.Storage, in *internal.Internal, queue *queue.Queue) UserRepoI {
	return UserRepoI(&UserRepo{
		st:    st,
		in:    in,
		queue: queue,
		mu:    sync.Mutex{},
	})
}

// Create receives an entity and inserts into DB.
func (ur *UserRepo) Create(u entity.User) (entity.User, error) {
	if err := ur.DB().Create(u).Error; err != nil {
		ur.in.L.Error(err.Error())
		return entity.User{}, err
	}

	go ur.Publish(u, "create")

	return u, nil
}

// Update receives an entity and id, and update it.
func (ur *UserRepo) Update(id string, u entity.User) (entity.User, error) {
	user := new(entity.User)
	if err := ur.DB().Where("uuid = ?", id).First(user).Error; err != nil {
		ur.in.L.Error(err.Error())
		return entity.User{}, err
	}
	user.UpdateFields(u)

	if err := ur.DB().Save(user).Error; err != nil {
		ur.in.L.Error(err.Error())
		return entity.User{}, err
	}

	go ur.Publish(u, "update")

	return *user, nil
}

// Read receives an ID and returns a user.
func (ur *UserRepo) Read(id string) (entity.User, error) {
	user := new(entity.User)
	if err := ur.DB().Where("uuid = ?", id).First(user).Error; err != nil {
		ur.in.L.Error(err.Error())
		return *user, err
	}
	return *user, nil
}

// ReadAll returns all users in DB.
func (ur *UserRepo) ReadAll() ([]entity.User, error) {
	users := new([]entity.User)
	if err := ur.DB().Find(users).Error; err != nil {
		ur.in.L.Error(err.Error())
		return *users, err
	}
	return *users, nil
}

// Delete deletes a row from DB with given ID.
func (ur *UserRepo) Delete(id string) error {
	err := ur.DB().Where("uuid = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		ur.in.L.Error(err.Error())
	}

	go ur.Publish(entity.User{
		UUID: id,
	}, "update")

	return nil
}

// DB shortcut for *gorm.DB value.
func (ur *UserRepo) DB() *gorm.DB {
	return ur.st.SQL.Client
}

// Publish is a shortcut to insert
// data into Queue.
func (ur *UserRepo) Publish(u entity.User, action string) {
	ur.mu.Lock()
	defer ur.mu.Unlock()
	b, err := json.Marshal(u)
	if err != nil {
		ur.in.L.Error(err.Error())
		return
	}

	switch action {
	case "create":
		fallthrough
	case "update":
		if err := ur.queue.PublishCreate(b); err != nil {
			ur.in.L.Error(err.Error())
		}
		return
	case "delete":
		if err := ur.queue.PublishDelete(b); err != nil {
			ur.in.L.Error(err.Error())
		}
		return
	}
}
