package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = (0.0 / (1.5 + 0.1))
    var v2 uint = 42
    var v3 int = -2147483648
    var v4 float32 = (v1 * float32(1e10))
    v5 := []string{"\n\t", (("foo bar" + ("a" + "")) + "")}
    fmt.Println(v4, v3)
    v6 := GenS1{F0: 32767}
    v7 := GenS1{F0: -660}
    var v8 int8 = -128
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v9 := GenS1{F0: 32767}
    v10 := GenS2{F0: 42}
    var v11 string = ""
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    var v12 bool = true
    var v13 bool = (!false)
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int32 = -707
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 0}
    var v17 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
