<h2>GoSpinner <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="https://opensource.org/licenses/MIT"/></h2>

<div>
    <p>Package Go is simple to handle loading spinners at terminals with various variants.</p>
</div>

https://github.com/rakarmp/GoSpinner/assets/83684256/44acbd80-3e4d-4314-b0f8-2265631e5b81

### Installation

```shell
go get github.com/rakarmp/GoSpinner
```

### Use

```go
package main

import (
	"fmt"
	"time"

	gospinner "github.com/rakarmp/GoSpinner/GoSpinner"
)

func main() {
	spinner := gospinner.NewLoadingSpinner()

	// Default spinner
	spinner.Start("default", 5*time.Second)
	fmt.Println("Done!")

	// Arrow spinner
	spinner.Start("arrow", 3*time.Second)
	fmt.Println("Done!")

	// Circle spinner
	spinner.Start("circle", 7*time.Second)
	fmt.Println("Done!")

	// Star spinner
	spinner.Start("star", 4*time.Second)
	fmt.Println("Done!")
}
```
