package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint32, p1 uint) uint16 {
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = 569
        p1 = (p1 & p1)
        _ = fi8
        _ = v9
    }
    if true {
        if (!(0 <= -2147483648)) {
            if true {
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 bool = false
                    _ = fi10
                    _ = v11
                }
            }
            if true {
                v12 := []int{(v7 >> (((p1 ^ 65535)) & 63)), v7, 127}
                var v13 uint = (p1 + (p1 * p1))
                v14 := GenS1{F0: ("foo bar" + "foo bar"), F1: (v13 * 1), F2: p0}
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: 42, F1: 18446744073709551615}
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
    }
    if false {
        v18 := []int{}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := []int{-1}
    var v2 int32 = -2147483648
    var v3 uint16 = 2
    var v4 float64 = 0.1
    var v5 string = ""
    v6 := []uint32{2147483647, 65535, 758, 65535}
    var v19 uint8 = 5
    for fi20 := 0; fi20 < 2; fi20++ {
        v5 = ((v5 + v5) + (v5 + "foo bar"))
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    var v22 uint = 18446744073709551615
    v4 = (-2.25 / ((v4 + (-2.25 / v4)) - v4))
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int8 = 1
            v2 = 708
            var v26 uint = (v22 & (v22 + uint(v3)))
            _ = fi24
            _ = v25
            _ = v26
        }
        v19 += (v19 & (v19 / ((v19) | 1)))
        _ = fi23
    }
    v22 -= v22
    fmt.Println(v4, v4)
    var v27 int16 = 1000
    v28 := GenS2{F0: v19, F1: 293}
    var v29 uint8 = v19
    fmt.Println(v5)
    switch 0 {
    case 9223372036854775807:
        v4 = 0.0
    case 679:
        var v30 bool = false
        _ = v30
    }
    var v31 int16 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
