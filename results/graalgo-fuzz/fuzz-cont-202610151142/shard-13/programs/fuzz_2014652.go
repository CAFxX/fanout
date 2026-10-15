package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0() uint16 {
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6, 32767, fi6}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "a":
        fmt.Println(v5)
    case "":
        v5 -= (v5 & v5)
    case "foo bar":
        var v8 float32 = 1e10
        _ = v8
    default:
        v9 := []int{-198}
        _ = v9
    }
    var v10 uint8 = 152
    _ = v10
    _ = v5
    return 685
}
func genFn1(p0 uint, p1 string) uint32 {
    fmt.Println(p0, p0)
    var v11 int8 = 0
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}

func main() {
    v1 := []string{"a", "", "\n\t", "a"}
    v2 := []uint32{4294967295, 2147483647}
    var v3 int8 = 92
    v4 := make(map[int]int)
    var v12 int16 = 32767
    v12 += -1
    v13 := v4[696]
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint32 = 2
    var v15 uint = 4294967295
    switch -745 {
    case 128:
        if (9223372036854775807 != 64) {
            var v16 uint64 = uint64(v15)
            var v17 float32 = -2.25
            _ = v16
            _ = v17
        } else {
            v18 := GenS1{F0: false, F1: "\n\t"}
            var v19 uint64 = 1
            _ = v18
            _ = v19
        }
    }
    v4[42] = v13
    v3 += 1
    var v20 int16 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
