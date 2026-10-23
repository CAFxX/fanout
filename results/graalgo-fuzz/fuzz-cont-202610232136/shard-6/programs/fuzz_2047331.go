package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint) int {
    switch ("foo bar" + "") {
    case "hello":
        p0 = (p0 / ((p0) | 1))
    default:
        switch 0 {
        case -699:
            fmt.Println(p0)
        }
    }
    var v5 uint = (p0 | 42)
    for fi6 := 0; fi6 < 11; fi6++ {
        if false {
            p0 -= v5
            fmt.Println(p0, fi6)
        } else {
            v5 = 2
            v7 := GenS1{F0: 431, F1: 932}
            var v8 uint16 = 42
            _ = v7
            _ = v8
        }
        fmt.Println(p0, fi6, p0)
        p0 -= (p0 ^ 255)
        _ = fi6
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{4294967295, 2, 0, 42}
    var v3 int16 = 42
    v4 := make(map[string]string)
    v3 = (-(v3))
    v1 = (v1 >> ((171) & 63))
    var v9 uint32 = 0
    v3 += (v3 & v3)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = (v1 ^ v1)
        v3 = (-(v3))
        _ = fi10
        _ = v11
    }
    _ = len(v2)
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v13 int = 525
        _ = fi12
        _ = v13
    }
    v1 = (^(v1))
    var v14 uint32 = v9
    var v15 uint = uint(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
