package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 uint16, p1 int8, p2 bool) uint32 {
    p0 = (^(p0))
    switch 127 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint16 = p0
            var v7 int64 = 74
            _ = fi5
            _ = v6
            _ = v7
        }
    case 4294967295:
        if p2 {
            var v8 int64 = -1
            _ = v8
        } else {
            p1 -= (p1 << ((4294967295) & 63))
            switch "\n\t" {
            case "\n\t":
                switch ("" + ("a" + "foo bar")) {
                case "foo bar":
                    p1 = 1
                case "\n\t":
                    var v9 int = -1
                    _ = v9
                case "a":
                    if false {
                        var v10 uint = 65535
                        p0 = (p0 << (((^(v10))) & 63))
                        var v11 uint8 = 121
                        _ = v10
                        _ = v11
                    }
                default:
                    fmt.Println(p0)
                }
            case "":
                p0 = (p0 % ((0) | 1))
            case "foo bar":
                var v12 int8 = (p1 & p1)
                _ = v12
            }
        }
    default:
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := []int{773, (fi13 / ((fi13) | 1)), -9223372036854775808, fi13}
            v15 := GenS1{F0: 1, F1: false, F2: 917}
            v14 = append(v14, (-(fi13)))
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 4294967295
}
func genFn1(p0 uint32) bool {
    var v17 uint32 = genFn0(42, -128, (!(!false)))
    var v18 int32 = -755
    switch 32767 {
    case 380:
        var v19 int = 9223372036854775807
        _ = v19
    case 127:
        var v20 float64 = 3.14159
        _ = v20
    case 42:
        v21 := GenS1{F0: (p0 / ((v17) | 1)), F1: true, F2: v17}
        _ = v21
    default:
        fmt.Println(v18, v17, v17)
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int32 = (v18 / (((v18 & (^(v18)))) | 1))
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = v17
    _ = v18
    return (false || (9223372036854775807 >= 42))
}

func main() {
    var v1 int = 32767
    var v2 int16 = 1000
    var v3 int = (v1 << ((255) & 63))
    var v4 uint16 = 1
    var v24 int32 = -2147483648
    fmt.Println(v24, v2, v1)
    if false {
        var v25 int16 = (v2 << ((883) & 63))
        v26 := GenS1{F0: 4294967295, F1: (!(1 != 9223372036854775807)), F2: 0}
        _ = v25
        _ = v26
    }
    var v27 float64 = ((3.14159 - -2.25) * -2.25)
    fmt.Println(v1, v4, v1)
    v2 = -32768
    v3 += (^(v3))
    var v28 int64 = 42
    fmt.Println(v24, v4)
    v29 := []int{(v3 * -9223372036854775808), (v1 + v1), (v3 >> ((0) & 63))}
    v30 := GenS2{F0: v24, F1: 2147483647, F2: v24}
    var v31 float64 = (0.1 / ((v27 + 0.1) / 1.0))
    v32 := GenS3{F0: (v24 <= (v24 << ((uint(v4)) & 63))), F1: 1, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
