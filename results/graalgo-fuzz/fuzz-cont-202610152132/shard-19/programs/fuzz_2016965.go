package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string) bool {
    var v4 bool = ((0 < 18446744073709551615) || ((true || true) && (2 == 1)))
    fmt.Println(p0, p0)
    var v5 int8 = 22
    _ = p0
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int, p1 uint32) int32 {
    var v6 float64 = (1.5 / 1e10)
    var v7 uint = 1
    p1 = (p1 % ((2147483647) | 1))
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-9223372036854775808, 1, -2147483648, -987}
    var v3 int8 = 23
    v9 := GenS1{F0: 1}
    var v10 int = 4294967295
    v3 -= v3
    var v11 int64 = int64(v3)
    if false {
        var v12 uint8 = 255
        if false {
            v13 := GenS1{F0: 1}
            var v14 int = v10
            _ = v13
            _ = v14
        }
        _ = v12
    }
    var v15 float32 = (-2.25 + 0.1)
    v11 += -9223372036854775808
    v15 = v15
    if ((!true) && ((0 <= 255) && (-32768 >= -311))) {
        var v16 bool = false
        _ = v16
    } else {
        if len(v2) > 0 {
            v2[0] = 32768
        }
    }
    var v17 int32 = -817
    for fi18 := 0; fi18 < 16; fi18++ {
        v17 = v17
        switch "a" {
        case "":
            v3 -= 0
        case "a":
            delete(v1, ("hello" + "hello"))
        case "\n\t":
            v11 -= v11
        default:
            var v19 uint32 = 2
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            v15 += v15
            v3 -= -55
            _ = fi20
        }
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
