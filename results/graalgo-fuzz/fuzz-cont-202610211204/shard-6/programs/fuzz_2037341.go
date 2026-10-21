package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32; F1 bool }
func genFn0() string {
    var v6 uint32 = 2147483647
    var v7 int16 = 1
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 string, p1 uint) int16 {
    v8 := []int{128}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{1, 0, 128}
    v3 := []string{}
    v4 := make(map[uint32]int)
    var v5 float64 = 0.1
    var v9 uint32 = 2147483647
    v9 = v9
    v5 -= (v5 + v5)
    for fi10 := 0; fi10 < 9; fi10++ {
        switch fi10 {
        case -778:
            var v11 int64 = -497
            _ = v11
        default:
            v9 += v9
        }
        _ = fi10
    }
    var v12 int64 = -1
    var v13 int8 = 119
    if (v5 == v5) {
        var v14 float64 = 1.5
        v9 = (v9 + (v9 % ((v9) | 1)))
        v9 -= 42
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 42}
    if ((((float64(1e10) + v5) / (v5 * v5)) <= v5) && true) {
        var v16 float64 = 1.5
        var v17 string = "\n\t"
        if (!(42 < 42)) {
            var v18 uint8 = 0
            _ = v18
        }
        _ = v16
        _ = v17
    }
    if false {
        if true {
            _ = len(v3)
            v12 -= (v12 | (v12 - v12))
        }
    } else {
        v19 := GenS2{F0: 2147483647, F1: (v5 == (v5 * float64(1.5)))}
        switch -1 {
        case -462:
            var v20 uint = 745
            _ = v20
        case 287:
            for fi21 := 0; fi21 < 7; fi21++ {
                _ = len(v4)
                _ = fi21
            }
        case -426:
            v13 -= 0
        }
        _ = v19
    }
    _ = len(v3)
    var v22 int64 = v12
    fmt.Println(v9, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
