package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 uint64, p2 int32) uint16 {
    p2 -= p2
    fmt.Println(p0)
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.0
    var v3 float32 = 0.1
    v4 := make(map[int]string)
    var v5 int8 = 42
    v6 := make(map[uint32]uint32)
    v4[-2147483648] = ("hello" + "\n\t")
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint = 4294967295
        v10 := GenS1{F0: 42}
        v5 += v5
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint16 = 32767
        v5 = -128
        _ = fi11
        _ = v12
    }
    v3 = (3.14159 * ((1.5 + v3) * (3.14159 - 3.14159)))
    v13 := []int{}
    fmt.Println(v3)
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v3 += (0.0 + v3)
            v16 := v4[fi15]
            v17 := GenS1{F0: 32767}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 42}
        v19 := GenS2{F0: 647, F1: 1, F2: -1}
        _ = fi14
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        switch 32768 {
        case 1:
            v2 = (((0.0 / v2) - ((1.5 - v3) - v2)) / v3)
        default:
            var v21 string = "hello"
            _ = v21
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS1{F0: int16(fi22)}
            _ = fi22
            _ = v23
        }
        _ = fi20
    }
    var v24 int8 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
