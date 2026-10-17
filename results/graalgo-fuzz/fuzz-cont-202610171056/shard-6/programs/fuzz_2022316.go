package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 int8, p1 int64) int8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(p1, p1)
        _ = fi7
    }
    if true {
        var v8 int32 = 2147483647
        _ = v8
    }
    var v9 uint8 = 19
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    var v1 uint = 631
    v2 := make(map[int]int)
    var v3 uint8 = 128
    v4 := []string{"\n\t", "", "hello", ("hello" + "foo bar")}
    v5 := make(map[string]int)
    v6 := []uint32{2147483647}
    var v10 uint32 = 0
    var v11 int8 = 43
    v1 = 4294967295
    switch 32768 {
    case 32767:
        _ = len(v2)
    default:
        v3 = 127
    }
    v12 := GenS1{F0: -637, F1: -2147483648}
    v13 := GenS2{F0: 2147483647, F1: 42}
    for fi14 := 0; fi14 < 7; fi14++ {
        if true {
            fmt.Println(v3)
            v15 := GenS2{F0: 579, F1: 0}
            _ = v15
        }
        _ = fi14
    }
    if true {
        v1 += (v1 % ((v1) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
