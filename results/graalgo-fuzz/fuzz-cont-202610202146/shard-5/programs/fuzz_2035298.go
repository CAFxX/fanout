package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float64 = (1.5 + float64((((3.14159 - (3.14159 - 1e10)) + 0.0) * (float32(-2.25) + float32((0.0 + (3.14159 + 3.14159)))))))
    var v2 uint8 = 42
    var v3 uint64 = 395
    if false {
        v4 := []int{}
        fmt.Println(v3, v2, v3)
        _ = v4
    }
    if (true || false) {
        v5 := []int{32768, 315, 128, 2147483647}
        v2 = (v2 | v2)
        switch -2147483648 {
        case 2147483647:
            v1 -= v1
        case 290:
            switch "foo bar" {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = 0
                }
            default:
                var v6 uint32 = 0
                _ = v6
            }
        }
        _ = v5
    }
    v7 := GenS1{F0: -9223372036854775808, F1: v3, F2: -37}
    var v8 int64 = -456
    var v9 uint8 = (^(v2))
    var v10 uint16 = 2
    var v11 int32 = 876
    var v12 uint = 65535
    for fi13 := 0; fi13 < 4; fi13++ {
        if (!true) {
            v14 := []int{(fi13 ^ fi13), fi13}
            v8 = (v8 >> (((v12 / (((v12 % (((^(v12))) | 1))) | 1))) & 63))
            _ = v14
        }
        var v15 uint8 = (v2 ^ v2)
        _ = fi13
        _ = v15
    }
    if true {
        var v16 float32 = 0.1
        v3 = (^(v3))
        var v17 int8 = 42
        _ = v16
        _ = v17
    } else {
        v3 = (v3 - uint64(v10))
        v11 += 218
    }
    var v18 uint16 = uint16(v12)
    var v19 int = 32768
    if false {
        var v20 uint8 = 128
        _ = v20
    } else {
        var v21 float32 = (0.1 * 0.1)
        v10 -= 255
        v1 = -2.25
        _ = v21
    }
    v22 := GenS2{F0: 477}
    var v23 int64 = (-(v8))
    v24 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
