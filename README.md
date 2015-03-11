GoBlockr
==========
Library in Go to work with Blockr.io ( http://btc.blockr.io/documentation/api ) API.

## Installation


```bash
$ go get github.com/conejoninja/goblockr
```

## Documentation
See [Go Doc](http://godoc.org/github.com/conejoninja/goblockr) or [Go Walker](http://gowalker.org/github.com/conejoninja/goblockr) for usage and details.

## Example of use

```go
package main


import (
	"github.com/conejoninja/goblockr"
	"fmt"
)

func main() {

	var api *goblockr.API

	api = goblockr.NewAPI("btc")

	res, err := api.CoinInfo()
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.ExchangeRate()
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.BlockInfo("0000000000000000210b10d620600dc1cc2380bb58eb2408f9767eb792ed31fa")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.BlockTxs("0000000000000000210b10d620600dc1cc2380bb58eb2408f9767eb792ed31fa")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.BlockRaw("0000000000000000210b10d620600dc1cc2380bb58eb2408f9767eb792ed31fa")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.TxInfo("60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471", "string")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.TxRaw("60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.UnconfirmedTxInfo("60c1f1a3160042152114e2bba45600a5045711c3a8a458016248acec59653471")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.AddressInfo("198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi", 15, "string")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.AddressBalance("198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi", 15)
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.AddressTxs("198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi")
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.AddressUnspent("198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi", true, false)
	fmt.Println(res)
	fmt.Println(err)

	res, err = api.AddressUnconfirmed("198aMn6ZYAczwrE5NvNTUMyJ5qkfy4g3Hi")
	fmt.Println(res)
	fmt.Println(err)

}
```

## Noted
I wouldn't use for anything serious or important.

## Contributing to GoBlockr:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.


## License

This is distributed under the Apache License v2.0

Copyright 2015 Daniel Esteban  -  conejo@conejo.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

