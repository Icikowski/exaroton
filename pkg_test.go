package exaroton_test

import (
	"context"
	"fmt"

	"pkg.icikowski.pl/exaroton"
)

func Example() {
	client, err := exaroton.NewClient(
		"your_token_here",
		// Can also specify options here:
		// exaroton.WithHTTPClient(http.DefaultClient),
		// exaroton.WithBaseURL("https://api.exaroton.com/v1"),
	)
	if err != nil {
		panic(err)
	}

	// Get account information
	account, _, err := client.GetAccount(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(account.Name)

	// Fetch all servers
	servers, _, err := client.GetServers(context.Background())
	if err != nil {
		panic(err)
	}

	for _, server := range servers {
		fmt.Println(server.Name)
	}

	// Get a specific server
	server := client.Server(servers[0].ID)
	serverInfo, _, err := server.GetServer(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(serverInfo.Motd)

	// Start the server
	_, err = server.Start(context.Background())
	if err != nil {
		panic(err)
	}

	// Get and set server RAM
	ram, _, err := server.GetRAM(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server %s has %dGB of RAM\n", serverInfo.Address, ram)

	if ram < 8 {
		_, _, err := server.SetRAM(context.Background(), 8)
		if err != nil {
			panic(err)
		}
		fmt.Println("Server RAM increased to 8GB")
	}

	// Get server logs
	logs, _, err := server.GetLogs(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(logs)

	// Stop the server
	_, err = server.Stop(context.Background())
	if err != nil {
		panic(err)
	}
}
