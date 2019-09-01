# Download Go code dependencies.
echo "Downloading dependencies..."
go get -v -t -d ./...
echo

# Generate embedded templates.
echo "Generating embedded templates..."
rm -rf server/generated
pushd server/http
go generate
popd