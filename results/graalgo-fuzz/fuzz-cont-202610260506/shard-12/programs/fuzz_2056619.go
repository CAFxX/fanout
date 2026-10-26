package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2
    v4 := []uint32{v3}
    if true {
        fmt.Println(v3)
    }
    for fi5 := 0; fi5 < 5; fi5++ {
        v3 = (v3 + v3)
        _ = fi5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "" {
        case "hello":
            var v7 int = 971
            _ = v7
        case "foo bar":
            v3 = v3
        default:
            v3 = 2147483647
        }
        _ = fi6
    }
    var v8 uint = 702
    v9 := v1["foo bar"]
    v1["\n\t"] = (^(v9))
    v9 -= 2147483647
    v3 -= 0
    if len(v4) > 0 {
        v4[0] = 932
    }
    var v10 int32 = 42
    v11 := GenS1{F0: -2147483648, F1: 2147483647}
    delete(v1, "\n\t")
    v3 += (^(v3))
    v10 = v10
    v12 := GenS1{F0: -966, F1: 1}
    v10 += v10
    var v13 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
