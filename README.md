# Go Confit Client

go client to load configurations from confit

## Usage

```
import "github.com/Skillbill/confit-go"

	// ...

	repoId := "wOgys75iJVWmuL4Ykx1dBHgSsp03"
	secret := "f801cf39-b784-414e-b997-231b9cc51ebe"
	c := confit.Client{RepoId: repoId, Secret: secret}
	data, _ := c.LoadByAlias("production.json")

	cfg := Configuration{}
	json.Unmarshal(data, &cfg)
```
