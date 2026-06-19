package todo

type Service struct {
	todos []Todo
}

func (s *Service) Add(t Todo) (string, error) {
	
}
func (s *Service) Delete(title string) (string, error) {}
func (s *Service) Update(title string, t Todo) (string, error) {}
func (s *Service) List() ([]Todo, error) {}
func (s *Service) Search(title string) (string, error) {}