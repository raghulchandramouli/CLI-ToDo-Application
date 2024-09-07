package main

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	todos.add("Buy milk")
	todos.add("Buy bread")

	todos.toggle(0)
	todos.print()

	Storage.Save(todos)

}