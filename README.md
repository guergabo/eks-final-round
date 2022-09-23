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
##### Test Sample Expected Output
```bash 
make build
cd bin/
../scripts/test.sh
```
##### Unit Test Output
```bash
go test ./internal/... -cover
```
### Project File Structure  
```bash
.
├── Makefile
├── README.md
├── bin
│   ├── README.md
│   ├── airgabe
│   ├── airgabe.log
│   └── init-state.json
├── cmd
│   ├── README.md
│   └── airgabe
│       └── main.go
├── go.mod
├── internal
│   ├── README.md
│   ├── core
│   │   ├── domain
│   │   │   ├── actionConfig.go
│   │   │   ├── airplane.go
│   │   │   ├── booking.go
│   │   │   ├── cancellation.go
│   │   │   ├── domain_test.go
│   │   │   ├── row.go
│   │   │   └── seat.go
│   │   ├── dto
│   │   │   ├── dto_test.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── ports
│   │   │   ├── repositories.go
│   │   │   └── services.go
│   │   └── services
│   │       └── airgabesrv
│   │           ├── service.go
│   │           └── service_test.go
│   ├── handlers
│   │   └── airgabehdl
│   │       ├── handler.go
│   │       └── handler_test.go
│   └── repositories
│       └── airgaberepo
│           ├── airgaberepo_test.go
│           ├── init-state.json
│           └── localFile.go
├── pkg
│   ├── README.md
│   ├── logger
│   │   └── logger.go
│   └── utils
│       ├── utils.go
│       └── utils_test.go
└── scripts
    ├── README.md
    └── test.sh
```
### Architecture Overview  
![alt text](https://miro.medium.com/max/1400/1*ERYx0IB1pN-5ZX98cKAoUw.png)
