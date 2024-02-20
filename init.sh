#!/usr/bin/env sh

./elements-cli createwallet lwkelements
./elements-cli -rpcwallet=lwkelements -generate 3
./elements-cli -rpcwallet=lwkelements rescanblockchain
./elements-cli -rpcwallet=lwkelements getbalance

./lwk-cli signer load-software --mnemonic "current enforce ignore west hammer neutral obscure shiver say welcome you license" --signer s1
DESCRIPTOR=$(./lwk-cli signer singlesig-desc --signer s1 --descriptor-blinding-key slip77 --kind wpkh | jq -r .descriptor)
./lwk-cli wallet load --wallet w1 -d "$DESCRIPTOR"
ADDRESS=$(./lwk-cli wallet address --wallet w1 | jq -r .address)

./elements-cli -rpcwallet=lwkelements sendtoaddress "$ADDRESS" 1
./elements-cli -rpcwallet=lwkelements -generate 3
./elements-cli -rpcwallet=lwkelements rescanblockchain

./lwk-cli signer load-software --mnemonic "bargain west position crumble code cruel nasty wise zero rapid parent pumpkin" --signer s2
DESCRIPTOR=$(./lwk-cli signer singlesig-desc --signer s2 --descriptor-blinding-key slip77 --kind wpkh | jq -r .descriptor)
./lwk-cli wallet load --wallet w2 -d "$DESCRIPTOR"
./lwk-cli wallet address --wallet w2
