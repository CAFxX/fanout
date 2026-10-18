package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 82
    var v3 uint = 42
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 float32 = 1e10
        v3 = 65535
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "a":
        v2 = v2
    case "\n\t":
        if (v2 > 2) {
            v6 := v1[2147483647]
            _ = v6
        } else {
            v2 += 85
        }
    case "hello":
        var v7 float32 = 0.0
        _ = v7
    }
    if true {
        delete(v1, 451)
        _ = len(v1)
        fmt.Println(v3)
    }
    var v8 float64 = 0.1
    v9 := []int{457, 128}
    v2 = v2
    if true {
        v3 = 192
    } else {
        var v10 uint = (v3 - (v3 | v3))
        v11 := GenS1{F0: 0}
        switch 42 {
        case 9223372036854775807:
            if len(v9) > 0 {
                v9[0] = -1
            }
        case -484:
            fmt.Println(v2)
        case 716:
            var v12 uint64 = (^(v2))
            _ = v12
        default:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 float32 = (float32(0.0) - (0.0 * 3.14159))
                _ = fi13
                _ = v14
            }
        }
        _ = v10
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = int(v3)
    }
    fmt.Println(v3, v2)
    v15 := GenS1{F0: 42}
    if (-2.25 >= 1e10) {
        var v16 int64 = -1
        v17 := GenS2{F0: -1}
        v9 = append(v9, 2147483647)
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v9 = append(v9, 32767)
        var v19 uint64 = v2
        _ = fi18
        _ = v19
    }
    if true {
        fmt.Println(v8)
        var v20 float64 = v8
        var v21 int64 = -54
        _ = v20
        _ = v21
    } else {
        var v22 uint16 = 32767
        _ = v22
    }
    fmt.Println(v2, v8)
    if (-1 >= 9223372036854775807) {
        v3 = 1
    }
    var v23 int64 = -850
    v24 := GenS1{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
