package singleton

type singleton struct {
	counter int64
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int64 {
	s.counter++
	return s.counter
}
