# Download front-end dependencies.
echo "Downloading front-end dependencies..."
yarn

# Download Go code dependencies.
echo "Downloading dependencies..."
go get -v -t -d ./...
echo
