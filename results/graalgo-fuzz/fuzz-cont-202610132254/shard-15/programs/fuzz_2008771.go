package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
func genFn0() uint32 {
    v7 := []int{-251, 124}
    var v8 int32 = 2147483647
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int8) uint64 {
    var v9 float32 = 1e10
    switch 182 {
    case 1:
        v10 := []int{}
        _ = v10
    case -9223372036854775808:
        var v11 uint32 = genFn0()
        _ = v11
    default:
        v12 := []int{-574, 0, 4294967295}
        _ = v12
    }
    v13 := []int{128, 1, 127, 32768}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float32 = v9
        var v16 float64 = 0.0
        var v17 uint32 = 2147483648
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v13
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 2
    v6 := make(map[uint32]string)
    var v18 uint = 4294967295
    for fi19 := 0; fi19 < 7; fi19++ {
        v18 = v18
        _ = fi19
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v18, v5, v18)
    case "foo bar":
        v20 := GenS1{F0: 0, F1: 32767}
        _ = v20
    }
    v5 -= v5
    _ = len(v4)
    v21 := []int{127, 1}
    v22 := GenS1{F0: 255, F1: -32768}
    var v23 int8 = 1
    if true {
        var v24 int64 = -1
        var v25 uint16 = 934
        _ = v24
        _ = v25
    } else {
        var v26 int8 = v23
        switch ("a" + "a") {
        case "a":
            v6[238] = ((("hello" + "") + "a") + ("a" + ("a" + "a")))
        }
        var v27 float64 = 1e10
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 uint32 = genFn0()
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 0, F1: -32768}
    fmt.Println(v3, v3)
    v31 := GenS1{F0: 255, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
