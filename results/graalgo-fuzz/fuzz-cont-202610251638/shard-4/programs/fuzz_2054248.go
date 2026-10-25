package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    var v3 string = "foo bar"
    var v4 uint8 = 253
    var v5 int8 = 100
    fmt.Println(v5)
    switch 4294967295 {
    case 128:
        fmt.Println(v4, v4)
    case 0:
        v6 := []int{}
        _ = v6
    case -9223372036854775808:
        v1 -= (^(v1))
    }
    fmt.Println(v4, v5, v4)
    v1 += (v1 << ((2) & 63))
    for fi7 := 0; fi7 < 7; fi7++ {
        switch ((v3 + "") + (v3 + ((v3 + v3) + ("\n\t" + "hello")))) {
        case "\n\t":
            v8 := GenS1{F0: 0, F1: 4294967295}
            _ = v8
        case "hello":
            v9 := GenS2{F0: 0}
            _ = v9
        }
        _ = fi7
    }
    if true {
        var v10 string = "\n\t"
        v3 = (("\n\t" + v3) + (v10 + v10))
        var v11 int16 = 1
        _ = v10
        _ = v11
    }
    var v12 int64 = (^(v1))
    v13 := []int{0, 896, -9223372036854775808, -2147483648}
    fmt.Println(v4, v4)
    if (-2.25 > 0.0) {
        var v14 int64 = (-(v1))
        v4 += (v4 + v4)
        _ = v14
    }
    if (!(!true)) {
        var v15 uint32 = 2147483647
        _ = len(v2)
        fmt.Println(v5, v12)
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
