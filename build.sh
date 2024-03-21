solc --pretty-json --overwrite --abi -o abi ./Zorotocol.sol
mv abi/Zorotocol.abi Zorotocol.json
rm -rf abi
solc --optimize --overwrite --bin -o bin ./Zorotocol.sol
mv bin/Zorotocol.bin Zorotocol.bin
rm -rf bin
abigen --bin=Zorotocol.bin --abi=Zorotocol.json --pkg=contract --out=Zorotocol.go