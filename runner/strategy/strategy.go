package strategy

type Strategy interface {
	Init(dir, uri string)
	ShouldUse() (bool, error)
	Setup() (env map[string]string, err error)
}

var strategies = []Strategy{}

func InitAll(dir, uri string) {
	for _, strategy := range strategies {
		strategy.Init(dir, uri)
	}
}

func SetupAll() (map[string]string, error) {
	env := make(map[string]string)

	for _, strategy := range strategies {
		shouldUse, err := strategy.ShouldUse()
		if err != nil {
			return nil, err
		}

		if shouldUse {
			additionalEnv, err := strategy.Setup()
			if err != nil {
				return nil, err
			}

			for key, val := range additionalEnv {
				env[key] = env[key] + val
			}
		}
	}

	return env, nil
}
