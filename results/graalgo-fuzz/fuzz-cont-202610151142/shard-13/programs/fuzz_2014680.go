package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 97
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: uint32(v2), F1: 42, F2: 0}
    fmt.Println(v2)
    v2 += (^(v2))
    var v6 int = -160
    v6 = v6
    if true {
        switch v6 {
        case 128:
            v3[2147483647] = 128
        default:
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := []int{-934}
                _ = fi7
                _ = v8
            }
        }
    } else {
        v9 := []int{(v6 | (v6 ^ (v6 ^ (-(v6))))), (v6 ^ v6)}
        _ = v9
    }
    if true {
        var v10 int = 4294967295
        v11 := []int{}
        v12 := GenS1{F0: 65535, F1: 0, F2: 344}
        _ = v10
        _ = v11
        _ = v12
    }
    v2 -= (v2 | 127)
    fmt.Println(v6, v2)
    switch (v6 - v6) {
    case 0:
        v13 := GenS1{F0: 152, F1: 2, F2: 2147483647}
        _ = v13
    case 800:
        if (false || false) {
            var v14 int32 = 42
            v15 := []int{}
            _ = v14
            _ = v15
        } else {
            fmt.Println(v6)
            if (-1 < 42) {
                var v16 int8 = -1
                v2 = 127
                _ = v16
            }
            v3[2147483648] = (v6 + (-(v6)))
        }
    case 9223372036854775807:
        var v17 float64 = (0.0 + float64(1.5))
        _ = v17
    default:
        var v18 string = "hello"
        _ = v18
    }
    v6 -= (v6 >> ((1) & 63))
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 uint32 = 2147483647
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        fmt.Println(v2, v2, v6)
        v22 := GenS1{F0: 4294967295, F1: 895, F2: 993}
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := []int{fi23, -182, (fi23 ^ int(v2))}
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    if true {
        v6 = v6
        v2 += 255
    }
    fmt.Println(v2, v2)
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v6, v2)
        fmt.Println(fi25, v6, v2)
        var v26 int64 = -1
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 42, F1: (^(v6)), F2: (v6 * v6)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
