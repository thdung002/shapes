package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"shape/controllers"
	"shape/entities"
)

func main() {
	secretkey := "go-rbac-secret-key"

	localconf, _ := entities.LoadConfig("./config/conf.json")
	userdb := "./auth/users.json"

	auth := entities.NewAuth(userdb, &entities.Options{Secret: secretkey, Nonce: 5})

	svc := controllers.NewController(auth, localconf, secretkey)
	svc.Start()
	defer svc.Stop()

	// load provides initial load an reload
	svc.Load()

	if localconf.API.TLSEnabled {
		go svc.ListenAndServeTLS()

	} else {
		go svc.ListenAndServe()
	}

	log.Println("\nServer started.")

	sigHandler := func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt)
		sig := <-signChan
		log.Println("Cleanup processes started by ", sig, " signal")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		svc.Shutdown(ctx)
		os.Exit(1)
	}

	sigHandler()

}
