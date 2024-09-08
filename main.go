package main

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	Storage.Save(todos)

}