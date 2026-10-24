package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 uint64) uint64 {
    fmt.Println(p2, p1, p1)
    switch "" {
    case "a":
        fmt.Println(p0, p2)
    case "":
        fmt.Println(p0, p2)
    case "hello":
        var v6 uint32 = 211
        _ = v6
    default:
        var v7 uint8 = (p1 & p1)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 bool, p1 uint32) string {
    switch 886 {
    case -319:
        var v8 uint16 = 42
        _ = v8
    }
    var v9 uint8 = 2
    _ = p0
    _ = p1
    _ = v9
    return ("foo bar" + ("a" + "\n\t"))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    var v4 int16 = -304
    var v5 uint64 = v3
    v4 -= v4
    var v10 uint = 18446744073709551615
    if true {
        v11 := GenS1{F0: -278, F1: uint16(v10)}
        _ = v11
    } else {
        var v12 float32 = 0.1
        _ = v12
    }
    var v13 uint = (v10 >> (((^(v10))) & 63))
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 1
        switch "foo bar" {
        case "hello":
            v16 := []int{-1, (-(fi14)), 57, -663}
            _ = v16
        case "foo bar":
            var v17 float64 = 1.5
            _ = v17
        case "\n\t":
            var v18 uint64 = v5
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    if (float32((0.1 - ((3.14159 / 1.0) + (0.0 - -2.25)))) >= 3.14159) {
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(fi20)
                _ = fi20
            }
            v21 := GenS2{F0: 946, F1: 82}
            _ = fi19
            _ = v21
        }
    }
    if (true || (!false)) {
        var v22 uint32 = 2147483647
        _ = v22
    } else {
        v23 := GenS3{F0: 845}
        var v24 uint = v10
        _ = v23
        _ = v24
    }
    if ((false || (127 == 42)) || (false || (("\n\t" + "a") != "a"))) {
        var v25 int32 = -1
        _ = v25
    }
    var v26 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
