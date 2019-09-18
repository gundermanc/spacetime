# Download front-end dependencies.
echo "Downloading front-end dependencies..."
yarn || exit $?

# Download Go code dependencies.
echo "Downloading dependencies..."
go get -v -t -d ./... || exit $?

# TODO: are these redundant?
go get github.com/GeertJohan/go.rice || exit $?
go get github.com/GeertJohan/go.rice/rice || exit $?