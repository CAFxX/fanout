package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 uint8 = 1
    var v4 float64 = float64((-2.25 * (3.14159 + 0.1)))
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]bool)
    v3 -= (v3 * 42)
    var v7 int32 = 2147483647
    if true {
        var v8 uint8 = v3
        v4 = float64(float32((v4 - 1e10)))
        v9 := v5[42]
        _ = v8
        _ = v9
    } else {
        v4 -= ((v4 + (float64(1e10) / v4)) + v4)
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v7 = v7
        fmt.Println(v7, fi10, v3)
        var v11 int64 = -289
        _ = fi10
        _ = v11
    }
    v12 := []int{32767, 128, 1}
    var v13 uint64 = 749
    for fi14 := 0; fi14 < 17; fi14++ {
        v4 = (1e10 * v4)
        if ((v7 ^ (v7 & 0)) < (v7 & (v7 & v7))) {
            var v15 uint32 = 2147483647
            if true {
                if ((("" >= ("" + "hello")) || ("" == "")) && false) {
                    v7 = (v7 - (v7 & (v7 - -2147483648)))
                    if false {
                        delete(v6, 42)
                    }
                }
                if (v3 != (v3 - v3)) {
                    switch -2147483648 {
                    case 127:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    default:
                        if ((v7 / (((-(v7))) | 1)) <= -1) {
                            fmt.Println(v4, v15, v4)
                            v3 = (v3 | (^(v3)))
                            v15 += v15
                        }
                    }
                    var v16 uint16 = 53
                    v17 := GenS1{F0: 18446744073709551615}
                    _ = v16
                    _ = v17
                }
                v18 := GenS1{F0: 255}
                _ = v18
            } else {
                var v19 int = fi14
                var v20 uint8 = uint8(v13)
                v13 = 0
                _ = v19
                _ = v20
            }
            v7 += (v7 ^ v7)
            _ = v15
        } else {
            fmt.Println(v7)
            v4 -= float64(((1e10 * -2.25) / float32(v4)))
            fmt.Println(fi14)
        }
        _ = len(v12)
        _ = fi14
    }
    v21 := GenS1{F0: 4294967295}
    v4 -= 0.1
    var v22 bool = true
    fmt.Println(v7, v3, v4)
    fmt.Println(v3, v13, v7)
    var v23 string = "\n\t"
    if (5 < 42) {
        var v24 float32 = 0.1
        _ = v24
    } else {
        v3 -= 2
        v22 = v22
        switch 374 {
        case 127:
            var v25 uint8 = 255
            _ = v25
        case 9223372036854775807:
            v22 = (65535 != 2)
        }
    }
    v4 -= v4
    fmt.Println(v22, v3, v13)
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint = 255
        v4 = -2.25
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
