package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 uint8, p2 uint16) int8 {
    if false {
        if true {
            for fi6 := 0; fi6 < 20; fi6++ {
                v7 := []int{}
                var v8 int = 1
                switch "foo bar" {
                case "hello":
                    var v9 int8 = 0
                    _ = v9
                case "foo bar":
                    fmt.Println(fi6, p0)
                case "":
                    if true {
                        if len(v7) > 0 {
                            v7[0] = (-(v8))
                        }
                    }
                default:
                    v10 := GenS1{F0: 2147483648, F1: -738, F2: -9223372036854775808}
                    _ = v10
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            v11 := []int{-804, 536, 128}
            _ = v11
        } else {
            switch 42 {
            case 2147483647:
                var v12 int64 = 184
                _ = v12
            default:
                p0 = p0
            }
            var v13 uint64 = 444
            _ = v13
        }
        v14 := GenS2{F0: p0}
        _ = v14
    } else {
        fmt.Println(p0)
        p1 += 255
    }
    p1 = (p1 & 119)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := []int{4294967295, 2147483647}
    v5 := []int{128}
    switch 32768 {
    case 32768:
        var v15 int32 = 420
        _ = v15
    }
    var v16 bool = (!true)
    switch "foo bar" {
    case "":
        v16 = (!(((!(!((v16 || v16) && false))) || ((1e10 + 3.14159) >= -2.25)) || v16))
    }
    var v17 int32 = -11
    var v18 float32 = (0.0 - 1.5)
    v18 += ((((v18 - v18) * v18) * v18) * (v18 * (-2.25 + (((((v18 / v18) - v18) / v18) + v18) * (((v18 * (v18 - (0.0 / -2.25))) / (1e10 * (v18 * -2.25))) + (3.14159 / v18))))))
    v18 = v18
    fmt.Println(v17, v1, v16)
    if (v18 >= 1.5) {
        _ = len(v5)
        if ((v1 * 9223372036854775807) == (v1 * (v1 << ((65535) & 63)))) {
            v1 -= -771
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 bool = true
        if (729 >= 0) {
            var v21 uint8 = 255
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if v16 {
        var v22 int64 = (v1 * 42)
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
