package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int8 = 100
    var v2 int8 = (v1 / ((-1) | 1))
    v3 := []string{}
    v4 := make(map[int]string)
    var v5 uint64 = 1
    if true {
        var v6 uint64 = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v1, v2, fi7)
        v8 := GenS1{F0: 2147483647}
        for fi9 := 0; fi9 < 3; fi9++ {
            _ = len(v3)
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if false {
            var v11 uint8 = 42
            var v12 uint64 = 4294967295
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(v5, v1)
    var v13 int32 = -1
    if true {
        switch 714 {
        case -9223372036854775808:
            var v14 int64 = 1
            _ = v14
        }
        fmt.Println(v13, v2)
    }
    _ = len(v3)
    switch "a" {
    case "a":
        v1 = v2
    case "hello":
        v13 -= 2147483647
    }
    v15 := GenS1{F0: (v13 & v13)}
    var v16 uint16 = 42
    switch -23 {
    case 127:
        var v17 float64 = 0.0
        _ = v17
    case 154:
        v18 := GenS1{F0: v13}
        _ = v18
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
