package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 32767
    var v3 int = 2147483647
    var v4 int8 = 1
    var v5 uint = 65535
    var v6 string = "hello"
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v6, v1)
        _ = fi7
    }
    fmt.Println(v2, v2, v3)
    v5 = v5
    v8 := []int{(v3 - (v3 ^ (v3 + 2147483647)))}
    var v9 int64 = -9223372036854775808
    if len(v8) > 0 {
        v8[0] = (^(v3))
    }
    switch (-(v3)) {
    case 2147483647:
        v10 := GenS1{F0: "hello"}
        _ = v10
    case 86:
        v4 = (v4 ^ v4)
    default:
        var v11 int = v3
        _ = v11
    }
    var v12 int16 = 1
    v13 := GenS1{F0: v6}
    v9 = v9
    var v14 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
