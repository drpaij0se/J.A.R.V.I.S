echo "Downloading"
git clone https://github.com/drpaij0se/J.A.R.V.I.S.
cd J.A.R.V.I.S.
echo "Installing"
go get .
echo "compiling"
go build mani.go
echo "Done!"