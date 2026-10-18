package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 uint16 = 42
    var v5 bool = false
    v4 += v4
    v6 := []int{}
    if (float64(float32(1.5)) == (0.0 - float64(0.0))) {
        switch -981 {
        case 1:
            v1[2147483648] = "foo bar"
        case 127:
            v7 := GenS1{F0: 0, F1: -369}
            _ = v7
        case -2147483648:
            var v8 uint64 = 981
            _ = v8
        default:
            var v9 int32 = 47
            _ = v9
        }
        v4 -= (v4 << ((523) & 63))
    } else {
        v4 += (v4 ^ (v4 % ((v4) | 1)))
        v10 := GenS2{F0: -14, F1: (v5 && ((!true) || ((!(190 < uint8(v4))) && v5))), F2: -1}
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v4, v4)
        _ = fi11
    }
    var v12 string = "foo bar"
    v12 = "a"
    if (-1 == 1) {
        v13 := GenS2{F0: 156, F1: v5, F2: 32767}
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint16 = 65535
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    fmt.Println(v4)
    v16 := GenS3{F0: 520}
    v17 := GenS2{F0: 32767, F1: true, F2: 1000}
    var v18 uint16 = v4
    switch 127 {
    case 128:
        var v19 int64 = 9223372036854775807
        _ = v19
    case 32767:
        var v20 int32 = 205
        _ = v20
    default:
        var v21 int = -1
        _ = v21
    }
    fmt.Println(v18, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
