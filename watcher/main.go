package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/valkey-io/valkey-go"
)

func main() {
	// Créez un nouveau client Valkey
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"localhost:6379"}})
	if err != nil {
		log.Fatalf("Erreur lors de la création du client : %v", err)
	}
	defer client.Close()

	// Contexte avec annulation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Abonnez-vous au canal d'expiration de la base de données 0
	go func() {
		err := client.Receive(ctx, client.B().Subscribe().Channel("__keyevent@0__:expired").Build(), func(msg valkey.PubSubMessage) {
			fmt.Printf("Clé expirée : %s\n", msg.Message)
		})
		if err != nil {
			log.Printf("Erreur lors de la réception des messages : %v", err)
		}
	}()

	// Exemple : définir une clé avec une expiration de 10 secondes
	err = client.Do(ctx, client.B().Set().Key("ma_cle").Value("ma_valeur").ExSeconds(10).Build()).Error()
	if err != nil {
		log.Fatalf("Erreur lors de la définition de la clé : %v", err)
	}

	// Attendre suffisamment longtemps pour que la clé expire
	time.Sleep(15 * time.Second)
}
