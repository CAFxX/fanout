package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 int; F1 uint; F2 bool }

func main() {
    var v1 int = -9223372036854775808
    var v2 bool = (!true)
    var v3 uint16 = 2
    var v4 int8 = -1
    var v5 string = "hello"
    var v6 int32 = 2147483647
    var v7 int32 = (v6 / ((v6) | 1))
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{int(v4), v1}
        if ((^(v8)) > 42) {
            var v11 int = -1
            _ = v11
        }
        v12 := GenS1{F0: 0, F1: 2, F2: (uint8(v3) == 226)}
        _ = fi9
        _ = v10
        _ = v12
    }
    v13 := []int{}
    v14 := GenS2{F0: (v1 * 2147483647), F1: 0, F2: ((^(v8)) < 2147483647)}
    var v15 string = ("" + "")
    for fi16 := 0; fi16 < 18; fi16++ {
        v7 += -2147483648
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 uint8 = 128
            v13 = append(v13, 9223372036854775807)
            v1 = -1
            _ = fi17
            _ = v18
        }
        var v19 int16 = 1
        _ = fi16
        _ = v19
    }
    switch v5 {
    case "\n\t":
        var v20 uint64 = 42
        _ = v20
    }
    var v21 float32 = -2.25
    var v22 string = ("foo bar" + ((v15 + (v15 + "a")) + "\n\t"))
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    if len(v13) > 0 {
        v13[0] = (-(v1))
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        switch v15 {
        case "a":
            v1 -= -9223372036854775808
        default:
            v3 = 2
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
