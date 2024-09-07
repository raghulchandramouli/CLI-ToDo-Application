package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title 			string
	Completed 		bool
	CreatedAt 		time.Time
	CompletedAt 	*time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)

}

func (todos *Todos) validateIndex(index int) error {
	
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}
func (todos *Todos) delete(index int) error {
    if err := todos.validateIndex(index); err != nil {
        return err
    }

    *todos = append((*todos)[:index], (*todos)[index+1:]...)

    return nil
}
