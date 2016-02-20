package guber

type Event struct {
	Message string
	Count   int
	Source  struct {
		Host string
	}
}
