package feed

//Feeder feeds a feed to multiple subscribers
type Feeder struct {
	subscribers []chan interface{}
	close       chan chan interface{}
	input       chan interface{}
}

//Feed  interface
type Feed interface {
	Feed() chan<- interface{}
	Close() <-chan interface{}
	Subscribe() <-chan interface{}
}

func (f *Feeder) open() {
	f.input = make(chan interface{})
	f.close = make(chan chan interface{})
	go func() {
		defer func() {
			close(f.close)
			close(f.input)
		}()
		for {
			select {
			case i := <-f.input:
				for _, l := range f.subscribers {
					l <- i
				}
			case c := <-f.close:
				for _, l := range f.subscribers {
					close(l)
				}
				c <- 435
				return
			}
		}
	}()
}

//Feed the subsrcibers
func (f *Feeder) Feed() chan<- interface{} {
	if f.input == nil {
		f.open()
	}
	return f.input
}

//Close the feeder
func (f *Feeder) Close() <-chan interface{} {
	c := make(chan interface{})
	f.close <- c
	return c
}

//Subscribe to the feed
func (f *Feeder) Subscribe() <-chan interface{} {
	c := make(chan interface{})
	f.subscribers = append(f.subscribers, c)
	return c
}
