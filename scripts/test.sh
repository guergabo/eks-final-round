# clean up first
rm -rf current-state.json
rm -rf airgabe.log

# execute sample expected output from assignment
./airgabe --help
./airgabe BOOK A0 1
./airgabe CANCEL A0 1
./airgabe BOOK A0 1
./airgabe BOOK A0 1
./airgabe BOOK A1 1
./airgabe BOOK A2 4
./airgabe BOOK A5 1
./airgabe BOOK A6 3
./airgabe BOOK A8 1
./airgabe BOOK U1 1

# clean up again
rm -rf current-state.json
rm -rf airgabe.log
