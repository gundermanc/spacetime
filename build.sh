# Build...
echo "Building..."
go build -v ./cmd/spacetime || exit $?

# Embed assets...
echo "Embedding assets..."
pushd server
go run github.com/GeertJohan/go.rice/rice embed-go || exit $?
popd