package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (p *Publication) Prepare() error{
	if err := p.validation(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *Publication) validation() error{     
	if p.Title == ""{
		return	errors.New("title is Required")
	}

	if len(p.Title) < 5 { 
        return errors.New("title must be at least 5 characters")
    }

	if p.Content == ""{
		return	errors.New("content is Required")
	}

	if len(p.Content) < 15 {
		return  errors.New("content must be at least 15 characters")
	}

	return	nil
}

func (p *Publication) format(){
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}