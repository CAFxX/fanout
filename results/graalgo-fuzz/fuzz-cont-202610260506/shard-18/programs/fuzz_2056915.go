package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 uint, p1 int16) int {
    p0 = 19
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 358, F1: 74, F2: 775}
            p0 += (p0 / ((18446744073709551615) | 1))
            var v10 int32 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    p0 = 634
    var v11 int = 42
    for fi12 := 0; fi12 < 3; fi12++ {
        p0 = p0
        v13 := []int{(-(v11)), (^(v11))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    return (-(v11))
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = 1
    v3 := []uint32{42, 0, 4294967295, 42}
    v4 := []string{""}
    var v5 uint32 = 2
    var v6 uint16 = 0
    for fi14 := 0; fi14 < 11; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    v5 -= (v5 / (((v5 & (v5 / ((2147483648) | 1)))) | 1))
    v6 = (v6 * (^(v6)))
    var v15 uint = 2
    var v16 uint32 = (^(v5))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v16 = v16
    if len(v4) > 0 {
        v4[0] = ("foo bar" + "foo bar")
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        _ = len(v4)
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v19 uint = v15
    var v20 float64 = 0.0
    v20 -= -2.25
    v21 := GenS1{F0: 97, F1: 641, F2: uint64(v6)}
    for fi22 := 0; fi22 < 16; fi22++ {
        if len(v3) > 0 {
            v3[0] = v16
        }
        _ = fi22
    }
    v23 := GenS2{F0: 751, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
