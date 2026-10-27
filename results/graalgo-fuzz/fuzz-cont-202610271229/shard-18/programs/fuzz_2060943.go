package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    v7 := []int{2147483647, 811, 4294967295}
    fmt.Println(p0)
    p0 = 107
    var v8 float64 = float64((1e10 / float32(-2.25)))
    _ = p0
    _ = v7
    _ = v8
    return false
}
func genFn1() int {
    v9 := GenS1{F0: 128, F1: 1, F2: 9223372036854775807}
    v10 := GenS1{F0: 128, F1: 1, F2: -528}
    v11 := GenS2{F0: 91, F1: 595, F2: 663}
    var v12 int = 2147483647
    fmt.Println(v12)
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return v12
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{0, 42, 0}
    var v3 float32 = 1.5
    v4 := []uint32{1, 42, 2147483647, 131}
    var v5 string = "a"
    var v6 int8 = -83
    v6 += (-(v6))
    var v13 bool = (false || true)
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = -1
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        switch v5 {
        case "":
            v15 -= v15
        case "\n\t":
            _ = len(v4)
        default:
            v3 += (0.0 * 0.1)
        }
        _ = fi14
        _ = v15
    }
    v6 = v6
    var v16 int = -2147483648
    var v17 int16 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v13 = ((226 != 65535) || v13)
    var v18 uint8 = (v1 | v1)
    v13 = v13
    if ((2147483647 < 2) && ((3.14159 * v3) <= float32(0.1))) {
        v13 = (!(!(42 > -315)))
        switch "\n\t" {
        case "\n\t":
            if v13 {
                v19 := GenS2{F0: (v6 ^ (v6 >> ((0) & 63))), F1: 0, F2: 255}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v5 = "a"
                    fmt.Println(v6)
                    v13 = false
                    _ = fi20
                }
                v1 = (v18 - (v1 ^ v1))
            }
        case "hello":
            v5 = (v5 + (v5 + v5))
        case "":
            v16 = (v16 % (((v16 % (((^(v16))) | 1))) | 1))
        default:
            if false {
                v4 = append(v4, 258)
            }
        }
        var v21 int = 9223372036854775807
        _ = v21
    } else {
        _ = len(v2)
    }
    v3 += (float32(1.5) + v3)
    v4 = append(v4, 0)
    var v22 int = (v16 << ((568) & 63))
    v4 = append(v4, 0)
    var v23 int = 349
    var v24 int8 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
