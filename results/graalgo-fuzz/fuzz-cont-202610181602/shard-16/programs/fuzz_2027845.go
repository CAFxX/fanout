package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{226}
    var v2 int64 = 1
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 42
    v2 -= v2
    switch "a" {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi5 := 0; fi5 < 8; fi5++ {
            v4 += v4
            v3 += v3
            _ = fi5
        }
    }
    switch 964 {
    case 4294967295:
        var v6 uint32 = 2147483647
        _ = v6
    }
    var v7 float64 = 0.0
    var v8 uint16 = v4
    if false {
        if len(v1) > 0 {
            v1[0] = 191
        }
    } else {
        v9 := GenS1{F0: 128}
        var v10 uint64 = (v3 * 1)
        var v11 uint8 = 42
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
