# Go Confit Client

go client to load configurations from confit

## Install

```
$ go get github.com/Skillbill/confit-go
```

## Usage

```
import (
	"github.com/Skillbill/confit-go"
	"encoding/json"
)

// ...

	repoId := "wOgys75iJVWmuL4Ykx1dBHgSsp03"
	secret := "f801cf39-b784-414e-b997-231b9cc51ebe"
	c := confit.Client{RepoId: repoId, Secret: secret}
	data, _ := c.LoadByPath("/prod/config.json")
	cfg := Configuration{}
	json.Unmarshal(data, &cfg)
```
