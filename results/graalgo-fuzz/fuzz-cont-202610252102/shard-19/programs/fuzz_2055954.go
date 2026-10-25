package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 bool; F2 int8 }
func genFn0(p0 string, p1 uint) bool {
    v7 := []int{42, 4294967295}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = 5
        v10 := GenS1{F0: (p1 & p1), F1: 1, F2: (fi8 ^ fi8)}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    p1 -= 49
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: p1, F1: 127, F2: fi11}
        p0 = "\n\t"
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (p1 * 1), F1: 1, F2: 127}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return false
}
func genFn1(p0 int16) uint64 {
    p0 -= p0
    if (42 <= 1) {
        p0 -= p0
        v14 := []int{}
        if false {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        _ = v14
    } else {
        if false {
            v15 := []int{-2147483648, 32768}
            _ = v15
        }
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = ((!(!true)) && false)
    var v3 uint64 = 1
    var v4 int64 = 1
    var v5 uint8 = 1
    v6 := []int{-207}
    switch 263 {
    case 9:
        fmt.Println(v3, v4)
    case 4294967295:
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v3)
            if len(v6) > 0 {
                v6[0] = (fi16 + (-(fi16)))
            }
            _ = fi16
        }
    }
    if v2 {
        v17 := GenS1{F0: 316, F1: (v5 + (^(v5))), F2: 2147483647}
        _ = v17
    } else {
        v18 := GenS2{F0: (v4 / ((v4) | 1)), F1: (float64(-2.25) == 1.5), F2: 42}
        _ = len(v6)
        for fi19 := 0; fi19 < 13; fi19++ {
            v2 = true
            v5 -= 42
            _ = fi19
        }
        _ = v18
    }
    v20 := GenS2{F0: (-(v4)), F1: v2, F2: 0}
    var v21 int64 = v4
    var v22 int64 = 1
    v22 += (v22 ^ v4)
    var v23 string = "\n\t"
    var v24 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
