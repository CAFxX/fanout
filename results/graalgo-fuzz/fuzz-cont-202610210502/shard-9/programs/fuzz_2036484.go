package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 128
    var v3 float64 = ((3.14159 - ((0.1 * float64(0.1)) * 1.5)) - 1e10)
    var v4 int32 = 552
    var v5 uint8 = 42
    var v6 string = "hello"
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v4, v3)
        var v8 int64 = -9223372036854775808
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 4294967295, F1: (v2 ^ (v2 >> ((251) & 63)))}
    switch 127 {
    case 128:
        v6 = (v6 + ("hello" + ("" + (v6 + ""))))
    default:
        var v10 bool = (-9223372036854775808 < 42)
        _ = v10
    }
    if false {
        var v11 uint32 = 2147483647
        v2 -= int16(v5)
        v11 += (v11 * v11)
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 uint32 = 65535
        _ = fi12
        _ = v13
    }
    if true {
        v2 += v2
        v14 := GenS1{F0: 700, F1: v2}
        var v15 uint64 = 828
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 742, F1: v2}
    switch 1 {
    case 2147483647:
        fmt.Println(v6, v3, v6)
    case -9223372036854775808:
        fmt.Println(v3, v4)
    case 0:
        var v17 int16 = -32768
        _ = v17
    default:
        var v18 int = 9223372036854775807
        _ = v18
    }
    if ((v5 / ((v5) | 1)) <= 2) {
        v5 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
