package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
func genFn0(p0 uint16) uint64 {
    v5 := []int{243, 32767, 4294967295, -781}
    p0 -= (p0 + (p0 & p0))
    _ = p0
    _ = v5
    return 770
}

func main() {
    var v1 int64 = -1
    var v2 int = 9223372036854775807
    var v3 uint8 = 2
    var v4 string = "hello"
    var v6 int32 = -461
    if (0 < 255) {
        var v7 uint = 981
        var v8 uint32 = 2147483647
        var v9 int8 = 0
        _ = v7
        _ = v8
        _ = v9
    }
    v2 += 4294967295
    v10 := GenS1{F0: -272, F1: -32768}
    v4 = v4
    v11 := []int{(v2 ^ 32768), (v2 * v2), 262, (v2 | (v2 | (^(v2))))}
    var v12 uint32 = 13
    v2 += (v2 & (v2 + (v2 ^ 9223372036854775807)))
    var v13 int32 = (v6 ^ (v6 + -1))
    v11 = append(v11, v2)
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: -32768, F1: 42}
        _ = fi14
        _ = v15
    }
    v3 = v3
    if true {
        v16 := GenS1{F0: int16(v2), F1: -1}
        _ = v16
    } else {
        var v17 int32 = (v13 >> ((102) & 63))
        _ = v17
    }
    var v18 bool = false
    v11 = append(v11, (v2 ^ v2))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
