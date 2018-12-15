# Build RESTful API in go using gorrila framework
### Install: 
clone source code from github:
```
git clone https://github.com/roynguyen/backend-go
```
set path to source code in your desktop:
```
export GOPATH= "$HOME/your_path"
```
### run: 
```
go run main.go 
```
### server start on:
http://localhost:8000
# Demo get balance and nonce of an Account in Ethereum network using go-web3 
### get balance 
```
http://localhost:8000/eth/0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d/balance
```
address : 0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d \\
response :
```
{
    "address": "0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d",
    "balance": 4527301006945768810
}
```
### get Nonce 
```
http://localhost:8000/eth/0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d/nonce
```
address : 0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d \\
response :
```
{
    "address": "0x5dCaA1D8D8132e5bF9CF12DECCFC0CeCF26a780d",
    "nonce": 161578
}
```