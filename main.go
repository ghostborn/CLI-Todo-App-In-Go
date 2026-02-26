package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	// todos.add("Buy Milk")
	// todos.add("Buy Bread")
	// todos.toggle(0)
	// todos.print()
	storage.Save(todos)
}
