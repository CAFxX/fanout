package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
func genFn0() uint {
    var v6 uint = 42
    var v7 int32 = 42
    v7 = (v7 >> ((65535) & 63))
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 int) int8 {
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 13; fi9++ {
        v8 += v8
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int32 = 2147483647
            _ = fi10
            _ = v11
        }
        v12 := []int{(p1 & (^(fi9))), (p1 ^ p1), fi9, p1}
        _ = fi9
        _ = v12
    }
    var v13 uint64 = v8
    switch "hello" {
    case "foo bar":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int64 = int64(p0)
            var v16 int8 = -128
            _ = fi14
            _ = v15
            _ = v16
        }
    case "\n\t":
        var v17 uint16 = 42
        _ = v17
    }
    v8 -= v8
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 uint = 255
    v4 := []string{}
    v5 := make(map[uint32]uint32)
    fmt.Println(v3)
    var v18 int8 = 42
    fmt.Println(v3)
    v18 = (v18 % (((^(v18))) | 1))
    var v19 int8 = v18
    v20 := GenS1{F0: 2147483647, F1: (!true), F2: (v19 & 1)}
    switch -9223372036854775808 {
    case -339:
        var v21 int8 = 100
        _ = v21
    }
    var v22 int = -1
    v23 := GenS1{F0: -9223372036854775808, F1: true, F2: v19}
    v3 += v3
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int = (-(fi24))
        var v26 uint64 = 863
        var v27 float64 = 3.14159
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if true {
        var v28 uint8 = 0
        _ = v28
    }
    v29 := GenS2{F0: (43 <= genFn0()), F1: 0, F2: -1}
    v3 = v3
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(v3)
        _ = fi30
    }
    var v31 uint64 = 18446744073709551615
    v32 := GenS2{F0: true, F1: 255, F2: v22}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
