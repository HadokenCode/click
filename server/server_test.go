//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/server/mock_contract_test.go github.com/kamilsk/click/server Service
package server_test
