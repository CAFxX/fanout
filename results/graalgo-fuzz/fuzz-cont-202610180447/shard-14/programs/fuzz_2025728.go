package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 uint32, p1 string, p2 int64) uint16 {
    p1 = ((p1 + "") + p1)
    var v4 int16 = -32768
    v4 = -721
    if true {
        v5 := GenS1{F0: v4}
        v6 := GenS2{F0: 32768, F1: ((^(p0)) >= p0), F2: 2}
        _ = v5
        _ = v6
    } else {
        var v7 uint16 = 1
        var v8 int64 = p2
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint16(p0)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    if false {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{}
            if false {
                v10 = append(v10, 32768)
                if len(v10) > 0 {
                    v10[0] = fi9
                }
                var v11 bool = (!true)
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 uint16 = 837
        v12 = v12
        _ = v12
    } else {
        var v13 float64 = 0.0
        _ = v13
    }
    switch v2 {
    case "":
        fmt.Println(v2)
    case "foo bar":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int32 = 776
            _ = fi14
            _ = v15
        }
    case "hello":
        var v16 int64 = 382
        _ = v16
    default:
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    fmt.Println(v2)
    v2 = (v2 + (v2 + "hello"))
    v18 := GenS1{F0: 696}
    v19 := GenS1{F0: -289}
    fmt.Println(v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        if false {
            var v21 float32 = ((0.0 * 0.1) - 0.0)
            _ = v21
        } else {
            v22 := GenS1{F0: -161}
            _ = v22
        }
        v2 = (v2 + "a")
        _ = fi20
    }
    v2 = ((v2 + v2) + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
