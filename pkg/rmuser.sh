#! /usr/bin/env bash

username=$1
sudo userdel $username
sudo rm -r /home/$username