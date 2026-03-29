package service

type Notifier interface {
	Notify(message string) error
}

type UserService struct {
	notifier Notifier
}

func (s *UserService) WelcomeUser(name string) error {
	return s.notifier.Notify("Welcome " + name)
}
