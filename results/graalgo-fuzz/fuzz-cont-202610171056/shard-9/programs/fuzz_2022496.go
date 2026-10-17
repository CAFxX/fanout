package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
func genFn0(p0 int) uint32 {
    var v6 uint = 2
    fmt.Println(v6)
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 uint32, p2 uint8) int16 {
    fmt.Println(p2, p0)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint32 = 2147483647
        _ = fi7
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 0, F1: -13, F2: 214}
        _ = v9
    }
    v10 := GenS1{F0: -404, F1: p0, F2: -666}
    var v11 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := []int{32768, -2147483648, 2147483647}
    var v2 float32 = 1.5
    var v3 int = 127
    v4 := []uint32{65535, 159, 591}
    var v5 uint8 = 0
    fmt.Println(v3, v5, v2)
    v12 := GenS1{F0: -1, F1: -128, F2: 2147483647}
    if true {
        fmt.Println(v3)
        v5 = (v5 - (v5 / ((219) | 1)))
        fmt.Println(v2, v3, v3)
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int64 = -1
        v15 := GenS1{F0: 2147483647, F1: 127, F2: 42}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v5, v5)
    if (42 != 18446744073709551615) {
        for fi16 := 0; fi16 < 20; fi16++ {
            v3 = fi16
            v4 = append(v4, genFn0((fi16 >> ((174) & 63))))
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(v2, v2)
                var v18 int32 = 898
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        if ((21 >= 18446744073709551615) || false) {
            if true {
                fmt.Println(v2, v5, v5)
            } else {
                v19 := GenS1{F0: 42, F1: 1, F2: 2147483647}
                _ = v19
            }
            v5 -= (^(v5))
            var v20 uint16 = 42
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int16 = -126
                fmt.Println(v3, v5)
                var v23 int32 = 0
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    var v24 float64 = float64(v2)
    var v25 float32 = (v2 - (3.14159 * 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
