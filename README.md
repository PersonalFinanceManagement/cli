## Cli application
the finance manager cli application.



Development steup
Edit the bash rc as follows
```
export GOROOT="/usr/local/go"
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
export PATH="$PATH:$GOROOT/bin"
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bash_profile
```

Required go packages to be installed.
```
go install github.com/amacneil/dbmate@latest
go install github.com/aarondl/sqlboiler/v4@latest
go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
```
