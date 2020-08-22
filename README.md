# golang-TDD
to run test you need to run:
```
go test -timeout 30s -run ^TestCall$
```

```
go test -timeout 30s -run ^TestContent$
```
## install:
You need to install
```
go get github.com/stretchr/testify
go get github.com/jarcoal/httpmock
```

## Test 1:

the first test is to check if the call function is there and if we have no problem to call it the response is replace by a mock so is a fack response.
## Test 2:

the second test is to check if the content is html and contain
and 
