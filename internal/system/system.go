package system

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/v5"

	"github.com/the-witcher-knight/store-contract/internal/config"
)

type System interface {
	EthClient() *ethclient.Client
	Mux() *chi.Mux
}

type system struct {
	ethClient *ethclient.Client
	mux       *chi.Mux
}

func New(cfg config.Config) (System, error) {
	s := system{}

	if err := s.initEthClient(cfg); err != nil {
		return nil, err
	}

	s.initMux()

	return &s, nil
}

func (s *system) initEthClient(cfg config.Config) error {
	conn, err := ethclient.Dial(cfg.Network)
	if err != nil {
		return err
	}

	s.ethClient = conn
	return nil
}

func (s *system) EthClient() *ethclient.Client {
	return s.ethClient
}

func (s *system) initMux() {
	s.mux = chi.NewMux()
}

func (s *system) Mux() *chi.Mux {
	return s.mux
}
