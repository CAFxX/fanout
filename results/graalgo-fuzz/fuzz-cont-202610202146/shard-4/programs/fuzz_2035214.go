package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }

func main() {
    var v1 int16 = 42
    var v2 float32 = 3.14159
    var v3 int = -478
    var v4 int16 = (v1 ^ (v1 & (v1 & -170)))
    v5 := make(map[string]int)
    v6 := []int{(v3 & (v3 * (v3 - 434))), 127, 32768}
    var v7 float32 = v2
    if false {
        switch "hello" {
        case "foo bar":
            switch "" {
            case "foo bar":
                var v8 uint8 = 0
                _ = v8
            }
        case "":
            fmt.Println(v2, v1, v2)
        default:
            fmt.Println(v4, v2)
        }
        fmt.Println(v3)
    }
    var v9 int = v3
    var v10 int32 = 42
    fmt.Println(v10, v7)
    v1 += 0
    v11 := GenS1{F0: 1}
    _ = len(v5)
    fmt.Println(v9, v4, v9)
    if false {
        v12 := GenS1{F0: 4294967295}
        var v13 int = (v3 >> ((255) & 63))
        switch v13 {
        case -2147483648:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint64 = 18446744073709551615
                v16 := GenS1{F0: 0}
                _ = fi14
                _ = v15
                _ = v16
            }
        case 1:
            var v17 int32 = v10
            _ = v17
        }
        _ = v12
        _ = v13
    }
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    v1 = 295
    v9 = (v9 ^ v3)
    v18 := GenS2{F0: 2, F1: 104, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
