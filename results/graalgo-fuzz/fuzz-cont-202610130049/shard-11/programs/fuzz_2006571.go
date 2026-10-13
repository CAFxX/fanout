package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 uint64 }

func main() {
    var v1 float32 = (-2.25 + ((3.14159 / float32(3.14159)) * (1.5 + 1.5)))
    v2 := make(map[uint32]string)
    v3 := []uint32{65535, 2147483647, 65535}
    var v4 int64 = -871
    if (false && true) {
        fmt.Println(v1)
    } else {
        v5 := GenS1{F0: 862, F1: 0, F2: 42}
        v4 = 481
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v4 -= (v4 - (v4 - 42))
                var v8 int8 = 127
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    }
    v9 := GenS1{F0: 0, F1: 920, F2: 1}
    var v10 int32 = 42
    fmt.Println(v1, v4, v4)
    fmt.Println(v1, v1, v10)
    v11 := GenS2{F0: 890, F1: (-(v4)), F2: 4294967295}
    switch 32767 {
    case -774:
        if true {
            fmt.Println(v4, v10, v1)
        }
    case 326:
        var v12 int64 = v4
        _ = v12
    }
    v13 := v2[2147483648]
    v2[241] = v13
    var v14 int = 32767
    var v15 string = v13
    v16 := v2[4294967295]
    var v17 bool = (1 >= 95)
    v18 := GenS2{F0: (v14 << ((4294967295) & 63)), F1: (v4 & (v4 >> ((489) & 63))), F2: 369}
    v14 = (v14 * (v14 * (v14 | v14)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
