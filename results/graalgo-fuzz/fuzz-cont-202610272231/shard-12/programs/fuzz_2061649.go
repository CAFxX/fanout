package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0() uint {
    if true {
        var v4 uint32 = 381
        _ = v4
    } else {
        var v5 int = 82
        var v6 int8 = -128
        _ = v5
        _ = v6
    }
    var v7 bool = (false || (-32768 <= -1))
    fmt.Println(v7)
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{(fi9 << ((65535) & 63)), -592, (fi10 << ((18446744073709551615) & 63)), fi9}
            v12 := GenS1{F0: -1}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    v3 := make(map[uint32]string)
    fmt.Println(v2, v2)
    switch "\n\t" {
    case "\n\t":
        v1 -= v2
    case "a":
        v1 -= v2
    case "foo bar":
        v13 := GenS2{F0: 849, F1: 676}
        _ = v13
    }
    fmt.Println(v1, v1)
    if true {
        var v14 uint64 = 0
        _ = v14
    } else {
        var v15 int16 = -1
        var v16 uint8 = 128
        v17 := v3[65]
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := []int{-1, 459, -48, 0}
    v2 = v1
    v19 := GenS1{F0: 994}
    v3[65535] = "foo bar"
    fmt.Println(v2, v1)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
