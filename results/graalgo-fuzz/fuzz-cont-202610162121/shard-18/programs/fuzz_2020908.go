package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 uint64 = 42
    v2[633] = 2
    for fi5 := 0; fi5 < 5; fi5++ {
        fmt.Println(v1, fi5, v1)
        _ = fi5
    }
    var v6 uint32 = 42
    v7 := []int{}
    var v8 uint16 = 332
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS1{F0: (("\n\t" + "a") + ""), F1: -2147483648, F2: 32767}
        _ = fi9
        _ = v10
    }
    switch 4294967295 {
    case 921:
        v4 = v4
    case 127:
        v8 = 32767
    }
    v6 -= v6
    if (((0.0 * 0.1) > 1.5) || (v6 <= v6)) {
        v1 = v1
        fmt.Println(v4, v1, v4)
        var v11 uint32 = v6
        _ = v11
    } else {
        if false {
            v12 := GenS1{F0: "\n\t", F1: (v1 >> ((758) & 63)), F2: -1}
            v4 = v4
            _ = v12
        } else {
            var v13 int16 = 0
            _ = len(v7)
            _ = v13
        }
    }
    var v14 int8 = 0
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint64 = (v4 << ((18446744073709551615) & 63))
        v8 = uint16(v16)
        _ = fi15
        _ = v16
    }
    if false {
        fmt.Println(v1, v4)
    } else {
        var v17 uint8 = 42
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
