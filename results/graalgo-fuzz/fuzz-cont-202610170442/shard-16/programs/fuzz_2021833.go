package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16) int64 {
    switch 588 {
    case 2147483647:
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int64 = 132
            v8 := []int{fi6, (^(fi6)), fi6, (fi6 >> ((255) & 63))}
            switch -965 {
            case -1:
                var v9 int16 = 0
                _ = v9
            case 4294967295:
                var v10 uint16 = 255
                _ = v10
            case 1:
                p0 = p0
            default:
                if true {
                    var v11 string = "a"
                    var v12 uint64 = 1
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint8 = 42
                    v14 := GenS1{F0: 0, F1: 1}
                    _ = v13
                    _ = v14
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case -9223372036854775808:
        if false {
            var v15 int = 32767
            switch v15 {
            case 32767:
                p0 = p0
            default:
                fmt.Println(p0, p0)
            }
            _ = v15
        } else {
            p0 = (p0 & p0)
            if true {
                v16 := []int{127, -1}
                var v17 uint16 = p0
                v17 = 242
                _ = v16
                _ = v17
            } else {
                var v18 float32 = 1.5
                var v19 uint8 = 42
                _ = v18
                _ = v19
            }
            v20 := []int{0, -2147483648, 4294967295}
            _ = v20
        }
    }
    if ((^(p0)) <= 255) {
        p0 += p0
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := []int{128, 195, 624, 1}
    var v2 float64 = 1e10
    v3 := []int{-799, 617, 127}
    var v4 float32 = (float32((((v2 * 1.5) / (3.14159 - ((float64(1.5) + 0.1) / v2))) - float64(-2.25))) + (float32(v2) + float32((3.14159 * -2.25))))
    v5 := make(map[string]int)
    switch -588 {
    case 0:
        delete(v5, (("" + "hello") + "\n\t"))
    case -9223372036854775808:
        var v21 int8 = -88
        _ = v21
    }
    fmt.Println(v4, v2)
    v4 -= (-2.25 + (v4 - 0.1))
    v4 -= 1e10
    v4 = (((((v4 + v4) / v4) * 0.0) + (((((v4 - -2.25) / float32(v2)) - v4) / (0.1 + -2.25)) / 0.1)) * v4)
    switch -310 {
    case 706:
        v22 := GenS1{F0: 74, F1: 0}
        _ = v22
    case -238:
        var v23 uint16 = 0
        _ = v23
    default:
        if ((1e10 / (v4 * (v4 + float32((v2 - 1e10))))) >= (v4 * v4)) {
            v4 = v4
            fmt.Println(v2, v2)
            v2 = 0.1
        } else {
            v4 += v4
        }
    }
    var v24 uint = 410
    v25 := GenS2{F0: 1}
    delete(v5, "foo bar")
    fmt.Println(v24, v4, v2)
    if len(v1) > 0 {
        v1[0] = 897
    }
    var v26 int32 = 998
    if len(v3) > 0 {
        v3[0] = 127
    }
    var v27 int64 = genFn0(0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
