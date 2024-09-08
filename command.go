package main

import (
    "flag"
)

type CmdFlags struct {
	Add string
	Del int 
	Edit string
	Toggle int 
	List bool
}

func NewCmdFlags() *CmdFlags {

	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title.")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index & specify a new title, id:new_title")
	flag.IntVar(&cf.Del, "Del", -1, "Specify a ToDo Index to be deleted")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Specify a ToDo Index to toggle")
	flag.BoolVar(&cf.List, "List", false, "List all ToDos")

	flag.Parse()

	return &cf
}