package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 141
    var v3 int16 = 0
    var v4 uint8 = v2
    var v5 int = 42
    var v6 uint32 = 42
    v7 := GenS1{F0: 126}
    v2 -= v4
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 127
        var v10 uint32 = (v6 - (v6 % ((v6) | 1)))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = 42
    v5 -= v5
    v12 := []int{}
    v6 = 2147483647
    switch ("hello" + (("foo bar" + "foo bar") + "a")) {
    case "foo bar":
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint8 = v4
            fmt.Println(v4, v14)
            var v15 int8 = 32
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        var v16 int8 = -71
        _ = v16
    }
    v17 := GenS2{F0: ("hello" + "a"), F1: (v11 - v11)}
    v3 = 32767
    v6 += v6
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS2{F0: "foo bar", F1: uint(v2)}
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
