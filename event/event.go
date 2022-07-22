package event

type Event interface {
	Apply(currState *CurrentState)
	Display()
}
