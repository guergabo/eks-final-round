# eks-final-round  

### Installation 
```bash
go get github.com/guergabo/eks-final-round
```

### Usage
```bash 
make build
cd bin/
./airgabe --help
./airgabe BOOK A0 1  
./airgabe CANCEL A0 1 
```
![alt text](https://giphy.com/embed/UYKVF23jjI4FR5Gemy)


### Run Tests
#### Test Sample Expected Output
```bash 
make build
cd bin/
../scripts/test.sh
```
#### Unit Test Output
```bash
go test ./internal/... -cover
```

### Architecture Overview  
![alt text](https://miro.medium.com/max/1400/1*ERYx0IB1pN-5ZX98cKAoUw.png)
