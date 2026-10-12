package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 32768
    v3 := []uint32{1, 1}
    var v4 uint8 = uint8(v1)
    var v5 uint32 = 4294967295
    v6 := []uint32{}
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v7 int8 = -20
    if false {
        v5 -= (v5 + (v5 << ((65535) & 63)))
        for fi8 := 0; fi8 < 19; fi8++ {
            v4 = (v4 & uint8(v5))
            _ = fi8
        }
    } else {
        if (!true) {
            fmt.Println(v5, v7, v1)
            v7 = (v7 % ((v7) | 1))
            v4 -= uint8(v5)
        } else {
            var v9 string = "hello"
            _ = v9
        }
        v10 := GenS1{F0: -2147483648, F1: (v4 * v4), F2: ((("hello" + "\n\t") + "\n\t") + "a")}
        _ = v10
    }
    var v11 int8 = v7
    fmt.Println(v2)
    var v12 uint16 = 65535
    v3 = append(v3, 4294967295)
    fmt.Println(v2, v7)
    v12 = 70
    fmt.Println(v11, v2)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS2{F0: (v1 + v1), F1: 0}
        if len(v6) > 0 {
            v6[0] = v5
        }
        _ = fi13
        _ = v14
    }
    var v15 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
