package mocks

import (
	"github.com/andrenormanlang/common"
	"fmt"
)

type DatabaseMock struct {
	GetPostHandler    func(int) (common.Post, error)
	GetPostsHandler   func(int, int) ([]common.Post, error)
	AddPostHandler    func(string, string, string) (int, error)
	DeletePostHandler func(int) (int, error)
	AddPageHandler 		func(string, string, string) (int, error)
}

func (db DatabaseMock) GetPosts(limit int, offset int) ([]common.Post, error) {
	return db.GetPostsHandler(limit, offset)
}

func (db DatabaseMock) GetPost(post_id int) (common.Post, error) {
	return db.GetPostHandler(post_id)
}

func (db DatabaseMock) AddPost(title string, excerpt string, content string) (int, error) {
	return db.AddPostHandler(title, excerpt, content)
}

func (db DatabaseMock) ChangePost(id int, title string, excerpt string, content string) error {
	return nil
}

func (db DatabaseMock) DeletePost(id int) (int, error) {
	return db.DeletePostHandler(id)
}

func (db DatabaseMock) AddPage(title string, content string, link string) (int, error) {
	// Call the mock handler instead of the method itself to avoid recursion
	if db.AddPageHandler != nil {
		return db.AddPageHandler(title, content, link)
	}
	return 0, fmt.Errorf("AddPage not implemented")
}