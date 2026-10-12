package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 382
    var v2 uint = 18446744073709551615
    var v3 int = 42
    var v4 uint64 = 0
    for fi5 := 0; fi5 < 13; fi5++ {
        v3 = (v3 % ((2147483647) | 1))
        _ = fi5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{}
        fmt.Println(fi6, v2)
        _ = fi6
        _ = v7
    }
    var v8 int32 = 2147483647
    var v9 uint = 65535
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v3, fi10)
        _ = fi10
    }
    var v11 int32 = v8
    switch "a" {
    case "a":
        fmt.Println(v8)
    case "foo bar":
        var v12 float64 = -2.25
        _ = v12
    case "":
        var v13 int32 = (v8 % ((v8) | 1))
        _ = v13
    default:
        var v14 uint32 = 0
        _ = v14
    }
    var v15 float64 = 0.0
    v8 -= v8
    v16 := GenS1{F0: v1}
    v17 := []int{127}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
