# description
This lwk Playground provides a docker stack that comprises elementsd, electrs and lwk.  
Operation can be checked in the regtest environment.

## usage

```sh
docker compose up -d
```

Create two wallets and add balances.

```sh
./init.sh
./lwk-cli wallet balance --wallet w1
./lwk-cli wallet balance --wallet w2
```