package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0(p0 int32, p1 uint64, p2 int64) uint {
    var v5 int16 = -32768
    var v6 uint32 = 2147483648
    var v7 uint8 = 128
    var v8 bool = (!false)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 549
}
func genFn1() uint8 {
    var v9 uint16 = 0
    fmt.Println(v9)
    v9 = v9
    v9 += v9
    _ = v9
    return 139
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 2
    var v3 int = 565
    var v4 int = -597
    v10 := []int{(v4 << ((2) & 63)), v4, v4, v3}
    v11 := GenS1{F0: 1, F1: ((("foo bar" + "hello") + "a") + "a"), F2: 32767}
    if len(v10) > 0 {
        v10[0] = v4
    }
    v3 = (v3 << ((genFn0(-2147483648, 229, 0)) & 63))
    var v12 int16 = -1
    for fi13 := 0; fi13 < 11; fi13++ {
        v12 -= v12
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 string = "a"
                _ = fi14
                _ = v15
            }
        case -1:
            var v16 string = "a"
            _ = v16
        default:
            var v17 float32 = 1.5
            _ = v17
        }
        _ = fi13
    }
    var v18 uint64 = 4294967295
    v2 -= 1
    v4 += 32768
    var v19 int8 = -56
    fmt.Println(v12, v12)
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int16 = v12
        fmt.Println(v12, v2, v4)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
