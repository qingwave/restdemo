package service

import "github.com/qingwave/weave/pkg/model"

type UserService interface {
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Get(string) (*model.User, error)
	CreateOAuthUser(user *model.User) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
	Delete(string) error
	Validate(*model.User) error
	Auth(*model.AuthUser) (*model.User, error)
	Default(*model.User)
	GetGroups(string) ([]model.Group, error)
}

type GroupService interface {
	List() ([]model.Group, error)
	Create(*model.User, *model.Group) (*model.Group, error)
	Get(string) (*model.Group, error)
	Update(string, *model.Group) (*model.Group, error)
	Delete(string) error
	GetUsers(gid string) ([]model.UserRole, error)
	AddUser(user *model.UserRole, gid string) error
	DelUser(user *model.UserRole, gid string) error
}

type PostService interface {
	List() ([]model.Post, error)
	Create(*model.User, *model.Post) (*model.Post, error)
	Get(user *model.User, id string) (*model.Post, error)
	Update(id string, post *model.Post) (*model.Post, error)
	Delete(id string) error
	GetTags(id string) ([]model.Tag, error)
	GetCategories(id string) ([]model.Category, error)
	AddLike(user *model.User, pid string) error
	DelLike(user *model.User, pid string) error
	AddComment(user *model.User, pid string, comment *model.Comment) (*model.Comment, error)
	DelComment(id string) error
}
