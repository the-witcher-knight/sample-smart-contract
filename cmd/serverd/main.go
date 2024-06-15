package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-chi/chi/v5"

	"github.com/the-witcher-knight/store-contract/internal/config"
	"github.com/the-witcher-knight/store-contract/internal/contracts"
	"github.com/the-witcher-knight/store-contract/internal/controller"
	"github.com/the-witcher-knight/store-contract/internal/handler"
	"github.com/the-witcher-knight/store-contract/internal/system"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := config.ReadConfigFromEnv()

	s, err := system.New(cfg)
	if err != nil {
		return err
	}

	chainID, err := s.EthClient().ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := authByPrivateKey(cfg, *chainID)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(cfg.ContractAddr)
	store, err := contracts.NewStorage(addr, s.EthClient())
	if err != nil {
		return err
	}

	ctrl := controller.New(*store, auth, *s.EthClient())
	hdl := handler.New(ctrl)

	addRoutes(s.Mux(), hdl)

	if err := startServer(ctx, s); err != nil {
		return err
	}

	return nil
}

func addRoutes(mux *chi.Mux, hdl handler.Handler) {
	mux.Route("/api", func(r chi.Router) {
		r.Mount("/storage", storeRoute(hdl))
	})
}

func storeRoute(hdl handler.Handler) http.Handler {
	return chi.NewRouter().Route("/storage", func(v1 chi.Router) {
		v1.Get("/", hdl.Retrieve)
		v1.Post("/{val}", hdl.Store)
	})
}

func startServer(ctx context.Context, s system.System) error {
	port := "8080"

	webServer := &http.Server{
		Addr:    ":" + port,
		Handler: s.Mux(),
	}

	fmt.Printf("Starting server on port %s\n", port)

	return webServer.ListenAndServe()
}

func authByPrivateKey(cfg config.Config, chainID big.Int) (bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(cfg.SecretKey)
	if err != nil {
		return bind.TransactOpts{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, &chainID)
	if err != nil {
		log.Fatalf("Failed to create keyed transactor: %v", err)
	}

	return *auth, nil
}
