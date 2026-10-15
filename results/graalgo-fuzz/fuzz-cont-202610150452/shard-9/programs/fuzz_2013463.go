package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 int, p1 int) uint32 {
    if (!(541 < 65535)) {
        var v5 bool = (1 > 1)
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint16 = 496
            _ = fi6
            _ = v7
        }
        var v8 int16 = 0
        _ = v5
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 0, F1: 0, F2: "foo bar"}
        var v11 int32 = 2147483647
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS2{F0: 128, F1: "hello", F2: -1}
            v14 := []int{}
            p0 += (fi12 << ((4294967295) & 63))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 int8 = 100
    var v15 float64 = 3.14159
    fmt.Println(v4, v1, v4)
    v2[0] = (!((int32(v1) != 70) && false))
    if true {
        v4 = (v4 + v4)
        v15 += v15
    } else {
        fmt.Println(v4, v4)
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v15 = v15
        v2[2] = (fi16 <= 0)
        var v17 string = "a"
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: v1, F1: "", F2: 1000}
    var v19 string = "hello"
    v4 += v4
    var v20 float64 = ((-2.25 / v15) / 0.1)
    var v21 int = (v1 ^ (v1 - v1))
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
