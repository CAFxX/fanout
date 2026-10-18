package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int32 = 42
    var v4 uint8 = 2
    v5 := GenS1{F0: false}
    switch "a" {
    case "hello":
        switch "" {
        case "a":
            var v6 float64 = -2.25
            _ = v6
        }
    case "a":
        v4 = v4
    case "":
        v7 := []int{-484, -141, 127}
        _ = v7
    default:
        v8 := GenS2{F0: -73}
        _ = v8
    }
    if (0 > 847) {
        v9 := GenS2{F0: 127}
        if true {
            v10 := GenS1{F0: false}
            switch 128 {
            case -9223372036854775808:
                v4 -= 0
            case 1:
                var v11 int64 = -264
                _ = v11
            }
            _ = v10
        }
        _ = v9
    }
    if (100 > 42) {
        fmt.Println(v3)
    }
    var v12 uint8 = v4
    var v13 uint16 = 1
    v14 := GenS1{F0: true}
    for fi15 := 0; fi15 < 17; fi15++ {
        v3 -= (^(v3))
        _ = fi15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS1{F0: (1.5 < 1e10)}
        var v18 uint64 = 18446744073709551615
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v13)
    switch 32768 {
    case 638:
        var v19 float32 = -2.25
        _ = v19
    case 1:
        v20 := []int{1, -391}
        _ = v20
    case 127:
        fmt.Println(v4, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
