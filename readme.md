# Cobra prototype

```
export GITHUB_USER=myusername
export CLI=mycli

cd ~/go/src
mkdir -p github.com/$GITHUB_USER
git clone git@github.com:hekonsek/cobra-prototype.git github.com/$GITHUB_USER/$CLI
cd github.com/$GITHUB_USER/$CLI
rm -rf .git
mv cobra-prototype $CLI
sed -i "s/hekonsek/$GITHUB_USER/g" * cmd/*
sed -i "s/cobra-prototype/$CLI/g" * cmd/*
```