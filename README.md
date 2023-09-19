# go-core


t.Fail()
t.FailNow()
t.Error()
t.Errorf()
t.Fatal()
t.Fatalf()


go test [inside package]
go test ./...
go test -run FuncName [ex: go test -run Sampletest ]
go test -v
go test -short

go test -run TestShort
go test -v -short -run TestShort


go test -run TestTableSimple
go test -v -run TestTableSimple
go test -v -run TestTableSlow
go test -v -run TestTableParallel