package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int16, p2 uint) uint64 {
    p0 -= p0
    p2 = (p2 - (p2 & 4294967295))
    if false {
        if (false && true) {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 float32 = 0.0
                if false {
                    var v9 int16 = p1
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        }
        var v10 int16 = p1
        _ = v10
    } else {
        var v11 uint64 = 42
        var v12 uint16 = 32767
        _ = v11
        _ = v12
    }
    fmt.Println(p0)
    if (!(!false)) {
        var v13 bool = (3.14159 > (-2.25 / (1e10 + 0.1)))
        p2 += (p2 & p2)
        _ = v13
    } else {
        p2 -= p2
        var v14 int32 = 1
        var v15 uint = p2
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 892
}
func genFn1(p0 uint, p1 int64) bool {
    var v16 uint = p0
    v17 := []int{459}
    switch -1 {
    case 42:
        var v18 int = -599
        _ = v18
    case 4294967295:
        v19 := GenS1{F0: ("a" + ""), F1: 1}
        _ = v19
    default:
        for fi20 := 0; fi20 < 1; fi20++ {
            p1 += p1
            _ = fi20
        }
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return (83 <= 420)
}

func main() {
    var v1 int32 = -874
    v2 := make(map[int]string)
    var v3 int32 = (v1 << ((575) & 63))
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 4294967295
    v6 := make(map[uint32]int)
    fmt.Println(v3, v1, v3)
    var v21 string = ((("foo bar" + "hello") + "") + "a")
    v22 := []int{-2147483648, 359, -1, -340}
    if len(v22) > 0 {
        v22[0] = 0
    }
    var v23 uint64 = v5
    v3 = v1
    v24 := GenS2{F0: 42}
    v5 += 469
    v22 = append(v22, -888)
    fmt.Println(v21, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
