package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG115 = []CodeSample{
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint32 = math.MaxUint32
    b := int32(a)
    fmt.Println(b)
}
	`}, 1, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint16 = math.MaxUint16
    b := int32(a)
    fmt.Println(b)
}
	`}, 0, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint32 = math.MaxUint32
    b := uint16(a)
    fmt.Println(b)
}
	`}, 1, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a int32 = math.MaxInt32
    b := int16(a)
    fmt.Println(b)
}
	`}, 1, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a int16 = math.MaxInt16
    b := int32(a)
    fmt.Println(b)
}
	`}, 0, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a int32 = math.MaxInt32
    b := uint32(a)
    fmt.Println(b)
}
	`}, 0, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint = math.MaxUint
    b := int16(a)
    fmt.Println(b)
}
	`}, 1, gosec.NewConfig()},
	{[]string{`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint = math.MaxUint
    b := int64(a)
    fmt.Println(b)
}
	`}, 1, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint = math.MaxUint
	// #nosec G115
	b := int64(a)
	fmt.Println(b)
}
		`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
	"math"
)

func main() {
    var a uint = math.MaxUint
	// #nosec G115
    b := int64(a)
    fmt.Println(b)
}
	`, `
package main

func ExampleFunction() {
}
`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
	"math"
)

type Uint uint

func main() {
    var a uint8 = math.MaxUint8
    b := Uint(a)
    fmt.Println(b)
}
	`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
)

func main() {
    var a byte = '\xff'
    b := int64(a)
    fmt.Println(b)
}
	`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
)

func main() {
    var a int8 = -1
    b := int64(a)
    fmt.Println(b)
}
	`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
	"math"
)

type CustomType int

func main() {
    var a uint = math.MaxUint
    b := CustomType(a)
    fmt.Println(b)
}
	`,
	}, 1, gosec.NewConfig()},
	{[]string{
		`
package main

import (
	"fmt"
)

func main() {
    a := []int{1,2,3}
    b := uint32(len(a))
    fmt.Println(b)
}
	`,
	}, 1, gosec.NewConfig()},
	{[]string{
		`
package main

import (
        "fmt"
)

func main() {
        a := "A\xFF"
        b := int64(a[0])
        fmt.Printf("%d\n", b)
}
	`,
	}, 0, gosec.NewConfig()},
	{[]string{
		`
package main

import (
        "fmt"
)

func main() {
        var a uint8 = 13
        b := int(a)
        fmt.Printf("%d\n", b)
}
	`,
	}, 0, gosec.NewConfig()},
}
