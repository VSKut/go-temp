package main

func main() {
	c := NewClient(NewConfig())

	c.Start()
}
