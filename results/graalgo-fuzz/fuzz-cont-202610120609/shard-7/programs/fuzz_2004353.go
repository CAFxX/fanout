package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 int8, p1 uint16) string {
    fmt.Println(p0, p1)
    v7 := []int{}
    if true {
        var v8 bool = (65535 != 2)
        _ = len(v7)
        if false {
            var v9 bool = true
            _ = v9
        } else {
            var v10 uint8 = uint8(p1)
            _ = v10
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}
func genFn1() int8 {
    v11 := GenS1{F0: 182, F1: 932}
    v12 := []int{}
    var v13 bool = true
    switch ("a" + ("hello" + "hello")) {
    case "foo bar":
        switch ("a" + "") {
        case "\n\t":
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: 127, F1: 0}
                _ = fi14
                _ = v15
            }
        case "a":
            var v16 bool = (!(815 != 0))
            _ = v16
        case "":
            fmt.Println(v13)
        }
    case "\n\t":
        if v13 {
            fmt.Println(v13)
            _ = len(v12)
        } else {
            v13 = ((2 <= 18446744073709551615) && v13)
        }
    default:
        v13 = (!(0 <= 0))
    }
    v17 := GenS1{F0: 2, F1: 0}
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    return -1
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    v5 := []string{}
    v6 := make(map[string]string)
    var v18 uint = 498
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: 127, F1: 18446744073709551615}
        v4 += (v4 - v4)
        fmt.Println(v18)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = ((true && (!false)) || ((v4 + 1) >= -617))
        if false {
            fmt.Println(v4, v22)
        }
        v18 = 248
        _ = fi21
        _ = v22
    }
    var v23 uint64 = 0
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = -656
    }
    var v24 float64 = 0.0
    v25 := GenS2{F0: 0, F1: true}
    var v26 int32 = -1
    v26 += v26
    v18 += 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
