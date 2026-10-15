package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }

func main() {
    v1 := []int{-426, 32767, 2147483647}
    var v2 int = 128
    var v3 string = "\n\t"
    var v4 uint32 = 162
    v2 -= v2
    v1 = append(v1, (^(v2)))
    for fi5 := 0; fi5 < 11; fi5++ {
        if len(v1) > 0 {
            v1[0] = fi5
        }
        var v6 uint = 512
        v3 = v3
        _ = fi5
        _ = v6
    }
    v3 = (v3 + v3)
    v7 := GenS1{F0: 1, F1: 42}
    v3 = (v3 + ((v3 + v3) + ("hello" + (v3 + v3))))
    if false {
        var v8 int32 = 2147483647
        _ = v8
    } else {
        var v9 int8 = 127
        _ = v9
    }
    v10 := GenS1{F0: 2, F1: v4}
    var v11 bool = false
    v4 += v4
    v4 = (v4 * (v4 - 1))
    v12 := GenS2{F0: 2, F1: v11, F2: 615}
    var v13 int = (v2 >> ((473) & 63))
    v14 := GenS1{F0: 42, F1: (v4 << ((65535) & 63))}
    var v15 uint8 = 2
    if ((!((float32(1.5) > -2.25) && false)) || (73 <= 18446744073709551615)) {
        var v16 uint64 = 2
        v17 := GenS1{F0: (v16 & v16), F1: v4}
        _ = v16
        _ = v17
    }
    var v18 float64 = (0.0 * (1e10 * (0.1 - 3.14159)))
    v15 += (v15 / ((v15) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
