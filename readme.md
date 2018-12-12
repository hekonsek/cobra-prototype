# Cobra prototype

```
export GITHUB_USER=xxx
export CLI=mycli

git clone git@github.com:hekonsek/cobra-prototype.git
mkdir -p $GITHUB_USER
mv cobra-prototype $GITHUB_USER/$CLI
cd $GITHUB_USER/$CLI
sed -i "s/hekonsek/$GITHUB_USER/g" go.mod
sed -i "s/cobra-prototype/$CLI/g" go.mod
```