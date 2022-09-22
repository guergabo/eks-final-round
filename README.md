# eks-final-round  

### Installation 
```bash
go get github.com/guergabo/eks-final-round
```

### Usage
```bash 
make build
cd bin/
./airgabe BOOK A0 1 --help --show
./airgabe CANCEL A0 1 --help --show
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

### Architecture Overview  



