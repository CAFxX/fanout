package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int16; F1 uint; F2 int }

func main() {
    var v1 bool = false
    var v2 int32 = 42
    var v3 string = ("" + ("hello" + "foo bar"))
    v4 := []int{32767}
    var v5 float64 = 0.1
    v6 := make(map[string]bool)
    v3 = "\n\t"
    v5 += v5
    v2 = -1
    v7 := GenS1{F0: 0, F1: "\n\t"}
    for fi8 := 0; fi8 < 15; fi8++ {
        v2 -= (v2 / (((-(v2))) | 1))
        v5 = v5
        switch (-(fi8)) {
        case -1:
            var v9 bool = (!(!v1))
            _ = v9
        case 32767:
            var v10 int16 = 0
            _ = v10
        default:
            var v11 bool = (259 >= 2147483647)
            _ = v11
        }
        _ = fi8
    }
    switch 2147483647 {
    case -406:
        fmt.Println(v1, v5)
    case 128:
        fmt.Println(v5, v1)
    default:
        v3 = (("hello" + v3) + "foo bar")
    }
    switch 32768 {
    case 9223372036854775807:
        _ = len(v4)
    case 128:
        var v12 bool = true
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        switch (v3 + v3) {
        case "\n\t":
            var v14 uint8 = 42
            _ = v14
        }
        _ = fi13
    }
    v5 += v5
    if true {
        v15 := GenS1{F0: -394, F1: v3}
        fmt.Println(v2)
        if ((4294967295 >= 1) || true) {
            v16 := GenS1{F0: -1, F1: (v3 + v3)}
            _ = v16
        } else {
            v5 -= ((v5 * 1e10) / v5)
        }
        _ = v15
    } else {
        fmt.Println(v2, v5, v5)
        v4 = append(v4, -9223372036854775808)
    }
    fmt.Println(v5, v3, v2)
    for fi17 := 0; fi17 < 9; fi17++ {
        v3 = v3
        var v18 uint64 = 18446744073709551615
        v4 = append(v4, (^(fi17)))
        _ = fi17
        _ = v18
    }
    switch v3 {
    case "a":
        switch 32767 {
        case 994:
            v19 := v6["a"]
            _ = v19
        }
    default:
        v20 := GenS2{F0: 42, F1: 255, F2: -253}
        _ = v20
    }
    v2 = (v2 ^ v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
