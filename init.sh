# Download front-end dependencies.
echo "Downloading front-end dependencies..."
yarn || exit $?

# Download Go code dependencies.
echo "Downloading dependencies..."
go get -v -t -d ./... || exit $?