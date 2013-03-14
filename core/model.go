package core

type Model struct {
	Theses []Thesis
}

type Thesis struct {
  Text string
  Arguments []Argument
}

type Argument struct {
  Text string
  Votes int32
  CounterArguments []Argument
}
