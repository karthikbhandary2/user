#! /usr/bin/env bash

username=$1
name=$2
sudo adduser -q --disabled-password --gecos "$name" $username --force-badname
sudo mkdir -p /home/$username/.ssh
public_key=$(aws ec2 describe-key-pairs --key-names $username --include-public-key | jq -r '.KeyPairs[].PublicKey')
echo "$public_key" | sudo tee -a /home/$username/.ssh/authorized_keys > /dev/null
sudo chown -R $username:$username /home/$username/.ssh
sudo chmod 700 /home/"$username"/.ssh
sudo chmod 600 /home/"$username"/.ssh/authorized_keys
