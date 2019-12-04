package main

func main() {
	s := NewServer(NewConfig())

	s.Start()
}
