#!/bin/bash
ADDRESS=$1

# Start
bitcoinzd \
  -mineraddress=$ADDRESS \
  -nuparams=5ba81b19:10  \
  -nuparams=76b809bb:20  \
  -nuparams=2bb40e60:30  \
  -nuparams=f5b9230b:40  \
  -nuparams=e9ff75a6:50
sleep 10

echo "BITCOINZ_ADDRESS=$ADDRESS"

# Import the address
bitcoinz-cli importaddress $ADDRESS

# Generate enough block to pass the maturation tim=
bitcoinz-cli generate 101

# Simulate mining
while :
do
    bitcoinz-cli generate 1
    sleep 10
done
