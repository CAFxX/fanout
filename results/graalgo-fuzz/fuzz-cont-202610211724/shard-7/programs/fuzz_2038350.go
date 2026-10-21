package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
func genFn0() int {
    switch "\n\t" {
    case "\n\t":
        var v7 int64 = 538
        _ = v7
    case "hello":
        var v8 float32 = 0.0
        _ = v8
    case "a":
        var v9 float64 = 0.0
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: "foo bar", F2: 2147483647}
    switch -2147483648 {
    case -2147483648:
        v11 := []int{32767, -705}
        _ = v11
    case -894:
        var v12 int64 = 42
        _ = v12
    default:
        v13 := GenS1{F0: 1, F1: "foo bar", F2: 571}
        _ = v13
    }
    var v14 int8 = -1
    _ = v10
    _ = v14
    return -1
}
func genFn1(p0 uint) bool {
    var v15 uint8 = 2
    var v16 int = -886
    var v17 int16 = 1
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ""
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v6 uint64 = 4294967295
    var v18 int = genFn0()
    var v19 float64 = 0.0
    v18 -= (v18 >> ((93) & 63))
    _ = len(v3)
    var v20 float64 = (v19 * 0.0)
    v20 += 0.0
    v21 := GenS2{F0: 128, F1: 65}
    v22 := GenS3{F0: -66, F1: 42, F2: 1}
    if false {
        v23 := GenS3{F0: 95, F1: 2, F2: -128}
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS4{F0: 2147483647}
            _ = fi24
            _ = v25
        }
        v26 := GenS5{F0: -225}
        _ = v23
        _ = v26
    }
    v27 := []int{(^(v18)), (-(v18)), int(v6), (v18 ^ genFn0())}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
