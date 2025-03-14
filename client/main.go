package main

import (
	"context"
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

	// Exemple : définir une clé avec une expiration de 10 secondes
	err = client.Do(ctx, client.B().Set().Key("ma_cle").Value("ma_valeur").ExSeconds(10).Build()).Error()
	if err != nil {
		log.Fatalf("Erreur lors de la définition de la clé : %v", err)
	}

	// Attendre suffisamment longtemps pour que la clé expire
	time.Sleep(15 * time.Second)
}
