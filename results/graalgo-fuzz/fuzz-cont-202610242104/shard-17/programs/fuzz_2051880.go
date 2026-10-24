package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint64, p1 uint16, p2 int16) uint {
    switch "\n\t" {
    case "foo bar":
        if true {
            p1 -= p1
            fmt.Println(p0, p2, p2)
            v6 := GenS1{F0: false, F1: 127}
            _ = v6
        }
    case "\n\t":
        switch -793 {
        case -614:
            p2 += p2
        }
    }
    p2 = (p2 / ((42) | 1))
    p0 -= p0
    p0 = 2
    _ = p0
    _ = p1
    _ = p2
    return 630
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (!(42 < -9223372036854775808))
    var v3 uint32 = 485
    var v4 string = "a"
    var v5 int32 = 2147483647
    var v7 int8 = -128
    for fi8 := 0; fi8 < 8; fi8++ {
        if (0.0 < ((((float32(3.14159) * ((3.14159 * 1.5) - 0.0)) / float32(1.5)) / 3.14159) / 0.1)) {
            v4 = v4
            v5 -= int32(v7)
            v9 := GenS2{F0: 9223372036854775807, F1: v4}
            _ = v9
        }
        if v2 {
            fmt.Println(v5, v7)
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 bool = (false || (v2 && (v3 > (v3 + (v3 * 42)))))
        var v12 int8 = v7
        var v13 uint = genFn0(0, 514, 0)
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v3 -= 2
    var v14 int64 = 0
    fmt.Println(v5, v5)
    if v2 {
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    var v16 int32 = (v5 + (-(v5)))
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int64 = (v14 + -999)
        _ = fi17
        _ = v18
    }
    v1[-2147483648] = -794
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
