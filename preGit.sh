#!/bin/bash
# Quick formatting tool

find . -iname '*.go' -exec gofmt -w {} \;

# Tests
echo "Running tests..... "
go test -coverprofile=c0.out github.com/mchirico/go-facebook/grabutil
go vet github.com/mchirico/septa/utils github.com/mchirico/go-facebook/grabutil

echo -e 'Want to test build?\n\n   go build github.com/mchirico/septa/routefirebase \n'
echo -e ' Rebase:   git rebase develop -i'
echo -e '          git push origin feature --force\n'
echo -e ' Force develop to match head:'
echo -e '          git checkout develop'
echo -e '          git fetch origin'
echo -e '          git reset --hard origin/master\n\n'
