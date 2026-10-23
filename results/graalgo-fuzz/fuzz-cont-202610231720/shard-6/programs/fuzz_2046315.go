package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
func genFn0(p0 int16, p1 string, p2 int) uint8 {
    if false {
        p0 += p0
        p1 = p1
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: -128, F1: 255, F2: "\n\t"}
            p1 = ((((("foo bar" + "\n\t") + p1) + p1) + "\n\t") + (p1 + "foo bar"))
            _ = fi7
            _ = v8
        }
    }
    var v9 uint16 = 65535
    switch (p2 << ((2) & 63)) {
    case 127:
        p2 += p2
    case 128:
        v10 := []int{p2, (p2 - 128), p2, (p2 ^ (p2 * (p2 / (((p2 * (p2 ^ 0))) | 1))))}
        _ = v10
    case 2147483647:
        fmt.Println(p0, p1)
    }
    p0 = (p0 + p0)
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 127
}
func genFn1(p0 uint16, p1 uint32) int {
    switch "foo bar" {
    case "":
        var v11 uint = 1
        _ = v11
    case "\n\t":
        v12 := GenS1{F0: -128, F1: 128, F2: ("hello" + ("a" + ""))}
        _ = v12
    case "hello":
        v13 := GenS1{F0: -1, F1: 128, F2: "a"}
        _ = v13
    }
    var v14 int = 2147483647
    switch (-(v14)) {
    case 456:
        v15 := []int{v14, v14}
        _ = v15
    case -1:
        v14 = v14
    default:
        var v16 uint = 18446744073709551615
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        fmt.Println(fi17, p1)
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v14
    return (v14 | 2147483647)
}

func main() {
    v1 := []string{"hello", (("\n\t" + "a") + "a"), ("foo bar" + "hello")}
    var v2 int32 = -2147483648
    var v3 string = "hello"
    v4 := []string{v3, "a", "a"}
    var v5 uint8 = 199
    var v6 uint64 = 4294967295
    var v18 uint16 = 473
    switch -860 {
    case 128:
        fmt.Println(v18, v2)
    case 1:
        var v19 uint32 = 2147483647
        _ = v19
    case 2147483647:
        var v20 float64 = (0.0 / (1e10 / 3.14159))
        _ = v20
    }
    var v21 int = genFn1((^(v18)), 4294967295)
    if true {
        v22 := GenS1{F0: 59, F1: 52, F2: v3}
        var v23 int32 = (-(v2))
        _ = v22
        _ = v23
    }
    if (!false) {
        var v24 int64 = 123
        _ = v24
    } else {
        v18 -= v18
        v2 += 0
    }
    if (-1 < 591) {
        fmt.Println(v6)
    } else {
        var v25 uint16 = (v18 ^ v18)
        _ = v25
    }
    v26 := GenS2{F0: (!false)}
    v27 := GenS2{F0: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
