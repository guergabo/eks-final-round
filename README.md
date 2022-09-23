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

### Hexagonal Architecture Overview  
![alt text](https://miro.medium.com/max/1400/1*ERYx0IB1pN-5ZX98cKAoUw.png)
