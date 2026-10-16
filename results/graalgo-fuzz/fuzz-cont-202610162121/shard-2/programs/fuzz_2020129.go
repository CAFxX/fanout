package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int64 {
    switch 4294967295 {
    case 127:
        var v4 int64 = 42
        _ = v4
    default:
        var v5 int32 = 768
        _ = v5
    }
    var v6 uint16 = 0
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 42
    var v3 uint = 708
    var v7 bool = (!false)
    v3 = v3
    if (false && v7) {
        v8 := GenS1{F0: 830}
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    fmt.Println(v9, v9)
    v3 -= v3
    var v10 bool = false
    v10 = (!v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
