package gac

type (
	Environment interface {
		GetUri() string
		GetName() string
	}

	PrototypeEnvironment struct {
	}
)
