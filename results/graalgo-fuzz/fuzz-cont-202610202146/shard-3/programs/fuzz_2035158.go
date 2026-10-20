package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 132
    var v2 int8 = -123
    v3 := make(map[uint32]int)
    var v4 uint16 = 255
    v1 -= (v1 / ((v1) | 1))
    var v5 uint = 0
    var v6 uint32 = uint32(v4)
    var v7 uint16 = v4
    var v8 int = 32767
    v9 := GenS1{F0: (-(v8)), F1: 376}
    v7 = v7
    v7 += 2
    var v10 bool = true
    fmt.Println(v6, v1)
    v11 := GenS1{F0: (v8 / ((4294967295) | 1)), F1: (v6 * (v6 % ((175) | 1)))}
    var v12 uint16 = (v7 + v4)
    var v13 uint64 = 0
    v2 += int8(v6)
    switch 2147483647 {
    case -1:
        var v14 int16 = -506
        _ = v14
    case 1:
        if (((9223372036854775807 == -9223372036854775808) && (v10 && true)) && true) {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 int8 = (-(v2))
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS2{F0: 839}
                    v4 += 32767
                    v5 -= (^(v5))
                    _ = fi17
                    _ = v18
                }
                v8 -= (v8 + (v8 << ((v5) & 63)))
                _ = fi15
                _ = v16
            }
        }
    default:
        if ((v1 - v1) == v1) {
            v13 = (v13 - v13)
            switch (v8 | (v8 * v8)) {
            case 127:
                var v19 float64 = -2.25
                _ = v19
            case 9223372036854775807:
                var v20 string = "hello"
                _ = v20
            default:
                v21 := GenS1{F0: (v8 * v8), F1: (^(v6))}
                _ = v21
            }
            v22 := GenS2{F0: 898}
            _ = v22
        }
    }
    v13 -= (v13 - v13)
    delete(v3, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
