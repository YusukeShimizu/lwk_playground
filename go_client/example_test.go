package lwkgo_test

import (
	"context"
	"fmt"
	"log"

	"github.com/YusukeShimizu/lwkgo/client"
)

func Example_address() {
	c := client.NewLwk("http://localhost:32112")
	res, err := c.Address(context.Background(), &client.AddressRequest{
		WalletName: "w1",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("address: %s\n", res.Address)
	// Output: address: el1qqgz7ugka4x68w22cqndpvsq9r7nu98qaa73qker74peqlrv7lfg3nsmne60r4rd4tupa68sdlddrdczclfe6yvmnp0k2pn73t
}

func Example_send() {
	ctx := context.Background()
	c := client.NewLwk("http://localhost:32112")
	res, err := c.Address(ctx, &client.AddressRequest{
		WalletName: "w2",
	})
	if err != nil {
		log.Fatal(err)
	}
	//send from w1 to w2
	sendres, err := c.Send(ctx, &client.SendRequest{
		WalletName: "w1",
		Addressees: []*client.UnvalidatedAddressee{
			{
				Address: res.Address,
				Satoshi: 1000,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	signed, err := c.Sign(ctx, &client.SignRequest{
		SignerName: "s1",
		Pset:       sendres.Pset,
	})
	if err != nil {
		log.Fatal(err)
	}
	broadcastres, err := c.Broadcast(ctx, &client.BroadcastRequest{
		WalletName: "w1",
		Pset:       signed.Pset,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bloadcast", broadcastres.Txid)
	// Output: dfdc5849b05c85a40fc6a3c89c3a30413c645c476f1fe634924ff61ec831867f
}
