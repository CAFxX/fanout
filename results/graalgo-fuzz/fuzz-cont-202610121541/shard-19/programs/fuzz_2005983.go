package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    var v1 int64 = -1
    v2 := []uint32{1, 2147483648}
    var v3 float64 = -2.25
    var v4 int64 = v1
    var v5 int8 = -128
    for fi6 := 0; fi6 < 8; fi6++ {
        v1 = v1
        var v7 float32 = 0.0
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: ("" + "foo bar")}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    v4 += v4
    fmt.Println(v5)
    v3 = v3
    v1 = -9223372036854775808
    for fi9 := 0; fi9 < 13; fi9++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = v5
        _ = fi9
    }
    if (2147483647 == 42) {
        v10 := GenS2{F0: 0, F1: (v4 | (v1 + v4))}
        _ = v10
    }
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 20; fi12++ {
        switch ("a" + ("hello" + "hello")) {
        case "foo bar":
            v2 = append(v2, 654)
        case "a":
            _ = len(v2)
        default:
            var v13 uint8 = 127
            _ = v13
        }
        _ = fi12
    }
    var v14 int8 = 72
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
