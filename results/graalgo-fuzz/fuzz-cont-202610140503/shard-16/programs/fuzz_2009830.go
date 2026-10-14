package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := []string{}
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    v4 := GenS1{F0: 42}
    switch "" {
    case "foo bar":
        v5 := GenS2{F0: 1, F1: 255}
        _ = v5
    case "a":
        var v6 int8 = 127
        _ = v6
    case "":
        var v7 bool = true
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int32 = -2147483648
        _ = len(v2)
        _ = len(v3)
        _ = fi8
        _ = v9
    }
    v1 = 0
    fmt.Println(v1)
    v1 += v1
    switch 128 {
    case 1:
        fmt.Println(v1)
    case -2147483648:
        var v10 string = "hello"
        _ = v10
    case -923:
        v1 = (^(v1))
    }
    v1 -= (^(v1))
    if (!false) {
        var v11 uint64 = 1
        var v12 float32 = 0.0
        fmt.Println(v11)
        _ = v11
        _ = v12
    } else {
        _ = len(v2)
        v1 += v1
        v1 = v1
    }
    _ = len(v2)
    v1 = v1
    switch 32767 {
    case 9223372036854775807:
        var v13 int = 32768
        _ = v13
    case -1:
        v14 := GenS1{F0: 313}
        _ = v14
    case 603:
        if (0.0 < (float64(0.0) + 0.0)) {
            v15 := GenS1{F0: 2147483647}
            var v16 float32 = 0.0
            _ = v15
            _ = v16
        }
    default:
        var v17 int32 = 1
        _ = v17
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
