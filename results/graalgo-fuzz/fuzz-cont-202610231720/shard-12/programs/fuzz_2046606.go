package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
func genFn0(p0 int8, p1 uint16) int {
    p1 += (p1 + p1)
    p1 -= p1
    v4 := GenS1{F0: (p1 - p1), F1: 0, F2: -128}
    v5 := []int{127, 0, 9223372036854775807, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = ("hello" + "hello")
    var v3 int64 = 0
    v2 = "\n\t"
    var v6 bool = (!false)
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, genFn0(40, 32767)}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    if v6 {
        switch (v2 + "\n\t") {
        case "foo bar":
            v3 = v3
        case "\n\t":
            fmt.Println(v6, v6, v2)
        case "a":
            v2 = "\n\t"
        }
        var v9 uint64 = v1
        _ = v9
    } else {
        v10 := []int{-9223372036854775808, int(v1)}
        var v11 string = (("a" + "foo bar") + ((v2 + (v2 + "a")) + (("a" + v2) + v2)))
        var v12 uint64 = (^(v1))
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int = 0
    v2 = v2
    switch (v13 >> ((2) & 63)) {
    case 1:
        v14 := GenS1{F0: 881, F1: v13, F2: 100}
        _ = v14
    default:
        v15 := GenS2{F0: -1, F1: 65535, F2: 4294967295}
        _ = v15
    }
    var v16 int32 = 13
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            if v6 {
                v19 := GenS2{F0: int16(v16), F1: 0, F2: (fi18 ^ (fi18 | fi18))}
                _ = v19
            } else {
                v20 := GenS2{F0: int16(v16), F1: 0, F2: -2147483648}
                _ = v20
            }
            _ = fi18
        }
        var v21 int = -1
        _ = fi17
        _ = v21
    }
    var v22 int16 = 32767
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v13, v2, v16)
        _ = fi23
    }
    v22 = v22
    fmt.Println(v13, v6)
    if (v6 && v6) {
        var v24 uint = 1
        _ = v24
    }
    v25 := []int{genFn0(-44, 616)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v6
}
