package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 127
    var v3 uint64 = 4294967295
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    v7 := GenS1{F0: 9223372036854775807, F1: 2, F2: 0}
    v8 := GenS1{F0: -1, F1: 127, F2: 0}
    v2 -= (v2 | 4294967295)
    if ("\n\t" < "") {
        var v9 int = v2
        var v10 int16 = 859
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: -1}
    var v12 uint16 = 461
    if true {
        _ = len(v1)
        v3 += v3
    } else {
        if ((674 < -2147483648) && true) {
            v3 -= (v3 & (v3 | 1))
            if (!true) {
                v12 = (v12 + 2)
                v1[817] = "foo bar"
            }
            v13 := GenS2{F0: (v6 % (((v6 / ((v6) | 1))) | 1))}
            _ = v13
        } else {
            v3 = 33
        }
    }
    var v14 uint64 = (v3 % ((0) | 1))
    v15 := GenS1{F0: 9223372036854775807, F1: 255, F2: -128}
    v16 := GenS1{F0: 602, F1: 1, F2: -128}
    var v17 uint16 = v12
    var v18 uint8 = 196
    var v19 uint16 = (v17 ^ 424)
    fmt.Println(v3)
    switch (v2 << ((4294967295) & 63)) {
    case 32767:
        var v20 int32 = 218
        _ = v20
    case 411:
        v3 -= 177
    default:
        delete(v4, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
