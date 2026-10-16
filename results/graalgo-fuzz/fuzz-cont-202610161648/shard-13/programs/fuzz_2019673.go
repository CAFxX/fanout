package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int8) int {
    var v7 uint8 = 0
    fmt.Println(v7, v7)
    v7 += v7
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return 32767
}
func genFn1(p0 uint32, p1 uint8) int {
    p1 = 1
    var v8 uint = 42
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: v8, F1: (p1 ^ p1)}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        p0 -= (p0 ^ (p0 >> (((v8 & v8)) & 63)))
        _ = fi11
    }
    var v12 uint64 = 801
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    var v2 uint32 = 891
    var v3 int8 = -1
    var v4 bool = (((((0.0 / 1e10) * (1e10 - 1.5)) / 1e10) / float32((0.0 + (1.5 / 1.0)))) != 3.14159)
    v5 := make(map[string]bool)
    var v6 int = 32768
    fmt.Println(v6, v3)
    _ = len(v5)
    var v13 uint8 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v5, ("foo bar" + ""))
    v6 += (v6 * (v6 % ((v6) | 1)))
    switch "foo bar" {
    case "\n\t":
        var v14 float64 = ((1e10 + 0.1) + 0.1)
        _ = v14
    case "a":
        var v15 int16 = int16(v3)
        _ = v15
    }
    switch "hello" {
    case "foo bar":
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
    case "":
        fmt.Println(v6)
    }
    var v17 uint = 42
    v18 := GenS2{F0: 598, F1: (("hello" + "hello") + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
