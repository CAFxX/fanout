package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 int) int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{(-(fi7)), -9223372036854775808}
            p0 -= (fi7 ^ -9223372036854775808)
            _ = fi8
            _ = v9
        }
        var v10 float64 = (1.5 / 1.5)
        v11 := GenS1{F0: "foo bar"}
        _ = fi7
        _ = v10
        _ = v11
    }
    p0 = (p0 << ((1) & 63))
    p0 -= p0
    _ = p0
    return -249
}
func genFn1(p0 string, p1 string, p2 uint8) int32 {
    p0 = ("\n\t" + ((p0 + p1) + "hello"))
    var v12 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2147483647
}

func main() {
    v1 := []string{}
    var v2 int32 = -2147483648
    var v3 uint32 = 2147483647
    var v4 int16 = 460
    var v5 uint16 = 1
    var v6 uint8 = 29
    fmt.Println(v6, v3)
    if (643 < 490) {
        var v13 int16 = v4
        _ = v13
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v14 uint8 = (^(v6))
        var v15 int = -551
        _ = v14
        _ = v15
    }
    if ((-(v4)) <= v4) {
        v1 = append(v1, "\n\t")
    } else {
        v6 = 63
    }
    fmt.Println(v5, v2, v6)
    v3 += (^(v3))
    for fi16 := 0; fi16 < 20; fi16++ {
        v2 = (v2 - v2)
        var v17 uint64 = 1
        _ = fi16
        _ = v17
    }
    var v18 float64 = 3.14159
    v5 = uint16(v3)
    v5 -= v5
    v19 := GenS2{F0: -688, F1: v6}
    for fi20 := 0; fi20 < 20; fi20++ {
        v1 = append(v1, "")
        _ = fi20
    }
    v1 = append(v1, "hello")
    v1 = append(v1, "foo bar")
    v1 = append(v1, "a")
    fmt.Println(v18, v6)
    v1 = append(v1, "hello")
    v1 = append(v1, ("hello" + "\n\t"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
