package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func() {
		for {
			x, ok := <-(*c).Input
			if !ok {
				close((*c).Output)
				break
			}
			(*c).Output <- x * x
		}
	}()
}
