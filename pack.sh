go build .
zip pack.zip test.png config.toml enrollment
scp pack.zip twt:~/
