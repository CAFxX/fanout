package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }

func main() {
    var v1 uint64 = 2
    var v2 string = "a"
    v3 := make(map[string]bool)
    var v4 uint = 2
    var v5 uint64 = (^(v1))
    v5 = (v5 % ((v1) | 1))
    v6 := GenS1{F0: 32767}
    switch v2 {
    case "foo bar":
        v7 := GenS2{F0: v2, F1: 877}
        _ = v7
    case "":
        var v8 float64 = 1e10
        _ = v8
    case "\n\t":
        fmt.Println(v2, v1, v5)
    default:
        if false {
            var v9 int = 795
            _ = v9
        } else {
            fmt.Println(v5, v5)
            v2 = v2
        }
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = 0
        switch v2 {
        case "":
            var v12 int32 = 1
            _ = v12
        case "foo bar":
            v13 := GenS3{F0: -1, F1: 0, F2: 0}
            _ = v13
        case "hello":
            var v14 float32 = float32(1e10)
            _ = v14
        }
        fmt.Println(v4, v5)
        _ = fi10
        _ = v11
    }
    v5 += (^(v5))
    var v15 int = 0
    var v16 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
