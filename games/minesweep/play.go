package minesweep

type Input struct {
	X, Y int
}

func (i *Input) Type() string {
	return "standard"
}

type Response int

// for feedback I can return some enum esque type

type MineSweep struct {
	*Board
	Moves      int
	Bombs      int
	Won        bool
	Player, ID string
}

func (m *MineSweep) EndGame() {

}

func (m *MineSweep) Play(input <-chan *Input) <-chan Response {
	out := make(chan Response)
	go func() {
		defer close(out)
		for in := range input {
			m.Moves++
			if m.Valid(in) {
				result, bomb := m.Uncover(in)
				if bomb {
					m.EndGame()
					break
				}
				out <- Response(result)
			}
		}
		close(out)
	}()
	return out
}
