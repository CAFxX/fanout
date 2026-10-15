package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int64, p1 int64, p2 uint64) int32 {
    var v4 int8 = 1
    v4 = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int16, p1 int8, p2 uint16) int64 {
    fmt.Println(p0, p2)
    var v5 float32 = (1e10 / ((3.14159 + (0.0 * float32(1e10))) / 1e10))
    p1 -= p1
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    v3 := []int{}
    switch 1 {
    case 32768:
        v2 = v2
    case 549:
        var v6 float64 = (0.0 / 3.14159)
        _ = v6
    default:
        v2 = v2
    }
    var v7 uint8 = 255
    v2 = v2
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: "a", F1: 35}
        var v10 string = "foo bar"
        v7 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    if (-9223372036854775808 > 42) {
        if false {
            for fi11 := 0; fi11 < 11; fi11++ {
                v3 = append(v3, fi11)
                v2 = ((("\n\t" + "a") + "") < (("hello" + ("a" + "")) + "a"))
                var v12 bool = (0 != 979)
                _ = fi11
                _ = v12
            }
            v13 := v1[9223372036854775807]
            v14 := GenS1{F0: v13, F1: -128}
            _ = v13
            _ = v14
        }
        var v15 uint32 = 2147483647
        v16 := GenS2{F0: -1, F1: 42}
        _ = v15
        _ = v16
    }
    var v17 int8 = -120
    var v18 int = 9223372036854775807
    v7 = v7
    v19 := GenS3{F0: "", F1: 18446744073709551615}
    v20 := GenS1{F0: "", F1: -38}
    delete(v1, (v18 << ((4294967295) & 63)))
    fmt.Println(v7)
    switch v18 {
    case -775:
        fmt.Println(v18, v17)
    default:
        var v21 uint16 = uint16(v7)
        _ = v21
    }
    if (v17 > v17) {
        for fi22 := 0; fi22 < 5; fi22++ {
            if (0 == 18446744073709551615) {
                if v2 {
                    v23 := GenS2{F0: v17, F1: 1}
                    _ = v23
                }
            } else {
                v17 = v17
            }
            var v24 float32 = float32(3.14159)
            _ = fi22
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
