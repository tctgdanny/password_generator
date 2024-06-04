#!/bin/bash

# This script will download the password generator binary from github and install it on your system
githubLink="https://github.com/thecoretg/password_generator/raw/main/pwgen"
filePath="/usr/local/bin/pwgen"

# Download the binary from github
sudo curl -L -o "$filePath" "$githubLink" && echo "Successfully downloaded binary from github"

# Make the binary executable
sudo chmod +x "$filePath" && echo "Binary is now executable"

# Remove the installer
if [[ -f "/tmp/install.sh" ]]; then
    sudo rm "/tmp/install.sh" && echo "Installer removed"
fi