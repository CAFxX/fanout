package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
func genFn0(p0 uint, p1 uint64, p2 uint32) uint {
    p2 += p2
    p0 = (^(p0))
    v4 := GenS1{F0: p1, F1: (false && false), F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 & p0)
}

func main() {
    v1 := []int{0, -2147483648, 663, 127}
    v2 := []int{128, -977}
    var v3 bool = (true && true)
    v3 = false
    _ = len(v1)
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint16 = 1
        var v7 int8 = -128
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = (-850 != 1)
    var v8 int16 = -1
    for fi9 := 0; fi9 < 11; fi9++ {
        if v3 {
            _ = len(v2)
            var v10 float64 = (0.0 / 1.0)
            v11 := GenS2{F0: 1000, F1: true, F2: fi9}
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    v8 = -71
    var v12 uint16 = 954
    if v3 {
        var v13 int32 = 0
        v14 := GenS2{F0: (v8 ^ 0), F1: true, F2: 127}
        _ = v13
        _ = v14
    } else {
        v12 += v12
        if (!false) {
            var v15 uint32 = 0
            if v3 {
                switch -1 {
                case 2147483647:
                    v12 -= 42
                case -838:
                    _ = len(v1)
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    if (-105 >= -1) {
                        var v17 float64 = -2.25
                        _ = v17
                    }
                    _ = fi16
                }
                fmt.Println(v15)
            }
            _ = v15
        }
    }
    _ = len(v2)
    v18 := GenS1{F0: 42, F1: v3, F2: 0}
    if v3 {
        v12 = v12
    }
    v12 = (v12 >> ((0) & 63))
    if v3 {
        if (3.14159 == ((3.14159 * ((float32((3.14159 * float64(1e10))) * 1.5) - 0.1)) - 3.14159)) {
            if (!(false && (true || false))) {
                v19 := GenS1{F0: 2, F1: (!true), F2: 1}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 20; fi20++ {
                    v21 := GenS2{F0: (v8 * -32768), F1: ((-54 <= -664) || (-2.25 > 1e10)), F2: fi20}
                    _ = fi20
                    _ = v21
                }
            }
        }
        v3 = false
        v22 := GenS1{F0: 2, F1: (42 > 566), F2: 42}
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
