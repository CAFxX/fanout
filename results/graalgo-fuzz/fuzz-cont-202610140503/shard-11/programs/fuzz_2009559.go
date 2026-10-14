package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 string) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 913
            fmt.Println(p0, fi7)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 42}
    v10 := []int{961}
    _ = p0
    _ = v10
    _ = v9
    return 2
}
func genFn1(p0 int8) uint64 {
    fmt.Println(p0)
    v11 := GenS2{F0: p0, F1: "foo bar"}
    var v12 int8 = genFn0(("\n\t" + ("\n\t" + "a")))
    var v13 string = ("\n\t" + "a")
    if false {
        v14 := []int{42, 32767, -1, -1}
        var v15 int32 = 0
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 int = 650
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 42
    var v5 int8 = 100
    v4 = 2147483647
    if true {
        v16 := v3[v4]
        v17 := GenS2{F0: v5, F1: "a"}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int32 = 2147483647
            var v20 string = "a"
            for fi21 := 0; fi21 < 8; fi21++ {
                v4 = v4
                v22 := GenS2{F0: 9, F1: "a"}
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        v23 := []int{(v1 - v1)}
        fmt.Println(v1)
        _ = v23
    }
    fmt.Println(v1)
    fmt.Println(v1)
    var v24 uint = 217
    switch ("" + "\n\t") {
    case "foo bar":
        var v25 uint8 = 42
        _ = v25
    }
    v5 = v5
    v26 := []int{v1, (v1 / (((v1 - int(v5))) | 1))}
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint16 = 1
        _ = fi27
        _ = v28
    }
    switch "\n\t" {
    case "":
        v29 := v2[(v1 | (-(v1)))]
        _ = v29
    case "a":
        v4 = v4
    default:
        v24 = 18446744073709551615
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
