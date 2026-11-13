package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float64 = 0.1
    v2 := []uint32{1}
    var v3 string = ("foo bar" + "a")
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 float64 = v1
    if (!(!(3.14159 != 0.1))) {
        fmt.Println(v1, v3, v6)
        var v7 float32 = float32(0.0)
        v8 := GenS1{F0: 18446744073709551615, F1: 2147483648, F2: 42}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 493, F2: 0}
    v10 := GenS1{F0: 566, F1: 42, F2: 127}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = ((v3 + v3) + (v3 + ""))
    v2 = append(v2, 0)
    delete(v5, 2)
    if ((false && ((0 == 255) || (!false))) || false) {
        switch -135 {
        case 42:
            var v11 uint32 = 2147483647
            _ = v11
        case 127:
            delete(v4, -9223372036854775808)
        }
        if true {
            var v12 uint16 = 849
            _ = v12
        }
        v3 = v3
    } else {
        if false {
            var v13 float32 = 3.14159
            if false {
                if ((0.1 + v13) != v13) {
                    v14 := GenS1{F0: 42, F1: 2147483648, F2: 127}
                    v1 = (v6 / v6)
                    v15 := GenS2{F0: -1}
                    _ = v14
                    _ = v15
                } else {
                    v13 = (v13 + (v13 - (v13 * v13)))
                    fmt.Println(v13, v1)
                }
            } else {
                v13 -= v13
            }
            _ = v13
        } else {
            if len(v2) > 0 {
                v2[0] = 556
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v16 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
