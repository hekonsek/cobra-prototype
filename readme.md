# Cobra prototype

This is template for new GoLang cli projects based on [Cobra](https://github.com/spf13/cobra).

## Usage

```
# Configure new project
export GITHUB_USER=myusername
export CLI=mycli

# Download template and customize it
cd ~/go/src
mkdir -p github.com/$GITHUB_USER
git clone git@github.com:hekonsek/cobra-prototype.git github.com/$GITHUB_USER/$CLI
cd github.com/$GITHUB_USER/$CLI
rm -rf .git
mv cobra-prototype.go $CLI.go
sed -i "s/hekonsek/$GITHUB_USER/g" go.mod Makefile $CLI.go .gitignore cmd/*
sed -i "s/cobra-prototype/$CLI/g" go.mod Makefile $CLI.go .gitignore cmd/*

# Enjoy!
make
./mycli

```