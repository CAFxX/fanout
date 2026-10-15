package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int8, p1 uint) uint16 {
    p1 -= (^(p1))
    var v7 int = 994
    p0 -= (p0 ^ p0)
    p1 = p1
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    return 348
}

func main() {
    var v1 int32 = -170
    v2 := make(map[int]bool)
    var v3 bool = false
    v4 := make(map[string]bool)
    var v5 int = 42
    v6 := make(map[string]bool)
    fmt.Println(v1, v1, v5)
    v8 := GenS1{F0: 1000, F1: true}
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v3, v1, v1)
        v10 := []int{}
        var v11 int = (^(v5))
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v3, v3)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint32 = 210
            _ = fi13
            _ = v14
        }
        v3 = true
        _ = fi12
    }
    if ((!(!v3)) || v3) {
        v15 := []int{(v5 ^ (v5 + v5)), (v5 << ((2) & 63)), v5, (v5 & v5)}
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if (2 >= 40) {
            v17 := []int{fi16, (fi16 - v5)}
            _ = v17
        }
        v18 := GenS2{F0: (v1 / (((v1 + (v1 >> ((42) & 63)))) | 1)), F1: 42, F2: 566}
        var v19 uint8 = 255
        _ = fi16
        _ = v18
        _ = v19
    }
    v3 = true
    v20 := []int{4294967295, (v5 % (((v5 + (v5 ^ v5))) | 1))}
    for fi21 := 0; fi21 < 9; fi21++ {
        v6[("" + "foo bar")] = true
        var v22 int64 = -1
        switch "\n\t" {
        case "hello":
            fmt.Println(v5)
        default:
            var v23 int = v5
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS1{F0: -1, F1: v3}
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: v5, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
