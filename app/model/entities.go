package model

import (
	"time"
)

// type byteString string

// the following structs vill be store appengine
type Article struct {
	Title      string
	Content    string
	CreateTime time.Time
	ID         string
}

func (self *Article) GetCreateTimeStr() string {
	timeStr := self.CreateTime.String()[0:16]
	return timeStr
}

type ArticleEntity struct {
	Title      string
	Content    []byte
	CreateTime time.Time
}

type Admin struct {
	GoogleAcount string
	ID           string
}

type AdminEntity struct {
	GoogleAcount string
}

type BlogCube struct {
	BlogName string
	NoRight  bool
	ID       string
}

type BlogCubeEntity struct {
	BlogName string
	NoRight  bool
}

type AFloat struct {
	Title      string
	Content    string
	CreateTime time.Time
	ID         string
}

type AFloatEntity struct {
	Title      string
	Content    []byte
	CreateTime time.Time
}
