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

func NewCmdFLags() *CmdFlags {

	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title.")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo by index & specify a new title, id:new_title")
}