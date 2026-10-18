package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 int32, p1 uint8) int {
    fmt.Println(p0)
    fmt.Println(p0, p1)
    var v6 int64 = -1
    _ = p0
    _ = p1
    _ = v6
    return -179
}
func genFn1() uint32 {
    v7 := []int{42, 128, 4294967295}
    var v8 int16 = 708
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{1, 0, 227}
    var v3 uint8 = 255
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    if (false || ((^(v3)) != (v3 / ((v3) | 1)))) {
        v3 = v3
        if false {
            for fi9 := 0; fi9 < 11; fi9++ {
                for fi10 := 0; fi10 < 7; fi10++ {
                    v3 += v3
                    _ = fi10
                }
                v11 := GenS1{F0: true, F1: 208, F2: (uint(v3) <= 0)}
                _ = fi9
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(fi12, v3)
            _ = fi12
        }
    }
    var v13 int64 = 42
    var v14 int16 = -298
    v14 = v14
    var v15 uint64 = 0
    v13 = -494
    if (!true) {
        if true {
            var v16 int16 = -169
            fmt.Println(v13, v3)
            _ = v16
        } else {
            v17 := GenS1{F0: true, F1: 2147483647, F2: false}
            _ = v17
        }
        var v18 float64 = float64((1.5 + 0.1))
        var v19 uint8 = v3
        _ = v18
        _ = v19
    }
    var v20 uint8 = (v3 / ((0) | 1))
    for fi21 := 0; fi21 < 19; fi21++ {
        if len(v2) > 0 {
            v2[0] = genFn1()
        }
        _ = fi21
    }
    if false {
        var v22 uint = 18446744073709551615
        _ = v22
    }
    v23 := GenS1{F0: (false && (0.0 < 0.0)), F1: -2147483648, F2: false}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
