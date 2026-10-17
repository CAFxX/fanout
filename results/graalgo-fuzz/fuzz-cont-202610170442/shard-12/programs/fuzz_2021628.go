package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 int64) bool {
    v7 := []int{-225, -9223372036854775808, -172}
    v8 := GenS1{F0: (1 >= 1)}
    if len(v7) > 0 {
        v7[0] = -735
    }
    if (-2147483648 == -1) {
        if len(v7) > 0 {
            v7[0] = 0
        }
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            v7 = append(v7, 482)
            if false {
                v7 = append(v7, fi9)
                for fi10 := 0; fi10 < 11; fi10++ {
                    p0 = (p0 - p0)
                    _ = fi10
                }
                p0 -= (^(p0))
            } else {
                v11 := GenS2{F0: 189}
                var v12 uint64 = 832
                _ = v11
                _ = v12
            }
            var v13 int = fi9
            _ = fi9
            _ = v13
        }
        var v14 float64 = 3.14159
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return ((("a" + ("hello" + ("foo bar" + "hello"))) + "a") <= (("" + "") + "hello"))
}

func main() {
    var v1 float64 = 0.1
    var v2 int = 9223372036854775807
    var v3 float32 = float32((0.1 / v1))
    var v4 bool = (0 != 1)
    var v5 uint8 = 255
    v6 := make(map[uint32]bool)
    var v15 float64 = (v1 + 3.14159)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 bool = v4
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v2, v3)
    var v18 uint = 4294967295
    switch -80 {
    case 2147483647:
        v3 += -2.25
    case -1:
        v2 += v2
    case 888:
        v19 := GenS1{F0: false}
        _ = v19
    default:
        var v20 bool = (!v4)
        _ = v20
    }
    v21 := GenS1{F0: (v4 || (v3 != (0.1 / float32(v1))))}
    var v22 int16 = -1
    _ = len(v6)
    if true {
        v23 := GenS1{F0: false}
        _ = v23
    } else {
        var v24 int = -234
        v1 = float64(v3)
        var v25 int16 = (v22 - (v22 / ((v22) | 1)))
        _ = v24
        _ = v25
    }
    var v26 float64 = 3.14159
    v27 := []int{(v2 >> ((4294967295) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
