package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 int16) bool {
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            p0 -= p0
            _ = fi6
        }
        var v7 uint64 = 18446744073709551615
        _ = v7
    } else {
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{128, -1, fi8, (fi8 & (-(fi8)))}
        p0 -= (-(p0))
        var v10 float32 = 1e10
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int32 = 42
        v13 := GenS1{F0: 42, F1: ("a" + "\n\t")}
        var v14 int32 = (^(v12))
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    return (p0 > 32767)
}

func main() {
    var v1 float64 = ((1.5 - 1.5) - (-2.25 - 1e10))
    v2 := []int{42}
    v3 := []string{"\n\t", (("" + ("\n\t" + "foo bar")) + ""), "a"}
    var v4 uint64 = 0
    var v5 uint8 = 0
    var v15 uint = 2
    var v16 uint = v15
    v15 += 4294967295
    var v17 float32 = (0.1 / ((float32(v1) / (((-2.25 - (1.5 - (((1e10 + (1e10 - 0.1)) * -2.25) - (((0.0 * 0.1) - (1.5 / float32(v1))) / 0.1)))) - 3.14159) * float32(((v1 - (((v1 * 0.0) / -2.25) + float64(((1e10 - 0.0) - 0.0)))) - v1)))) - 1.5))
    fmt.Println(v4, v5)
    var v18 float32 = float32(v1)
    var v19 int16 = 32767
    fmt.Println(v18, v4)
    fmt.Println(v4, v1, v19)
    var v20 uint16 = 32767
    for fi21 := 0; fi21 < 14; fi21++ {
        v19 = v19
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
