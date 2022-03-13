package example

import "fmt"

func Execute() {
	repo := NewStaticRepository()
	h := NewHandler(repo)
	s := h.Do(2)
	fmt.Println(s)
}

// ได้ obj Language
type Language struct {
	ID   uint
	Name string
}

// Repository -> type interface
// uinit -> id ของ => method QueryLang
type Repository interface {
	QueryLang(uint) Language
}

// struct -> type Repository
type Handler struct {
	repo Repository
}

// fn return pointer handler
// interface -> test ไง ?
func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Do(id uint) string {
	lang := h.repo.QueryLang(id)
	return fmt.Sprintf("%s language", lang.Name)
}

// TODO:this interface still bug
func NewStaticRepository() Repository {
	re := Repository{QueryLang(1)}
	return re
}

// Real Repository

var data = map[uint]Language{
	1: {ID: 1, Name: "Go"},
	2: {ID: 2, Name: "Java"},
	3: {ID: 3, Name: "Python"},
	4: {ID: 4, Name: "Rust"},
}

type StaticRepo struct {
	data map[uint]Language
}
