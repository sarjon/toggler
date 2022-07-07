package toggler

type Storage interface {
	Get(name string) (*Toggle, bool)
}
