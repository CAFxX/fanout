package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 uint8, p1 uint, p2 int) uint16 {
    if (1 > 65535) {
        var v7 uint64 = 4294967295
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            p2 = (fi8 | (fi8 + 1))
            p0 -= 127
            _ = fi8
        }
    }
    fmt.Println(p1, p2)
    if (true || false) {
        var v9 uint16 = 42
        v10 := GenS1{F0: 0, F1: (-(p2)), F2: 255}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 952
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    var v4 int64 = 132
    var v5 float32 = 0.0
    var v6 int8 = -1
    var v11 int8 = v6
    v12 := []int{9223372036854775807, -2147483648, 158, 826}
    switch "\n\t" {
    case "hello":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    case "":
        switch -184 {
        case 128:
            fmt.Println(v4, v6)
        }
    case "a":
        fmt.Println(v4)
    default:
        v13 := GenS1{F0: 5, F1: 32767, F2: 908}
        _ = v13
    }
    var v14 float32 = ((v5 + ((1.5 / v5) / float32((1.5 / 1.0)))) - v5)
    if true {
        v5 = ((-2.25 / float32(((3.14159 / float64(3.14159)) / (3.14159 * (-2.25 - 1e10))))) + (v14 * float32(0.1)))
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v14, v4, fi15)
            _ = fi16
        }
        v4 -= (v4 % (((-(v4))) | 1))
        _ = fi15
    }
    var v17 int64 = v4
    for fi18 := 0; fi18 < 20; fi18++ {
        v14 = v14
        v11 = -1
        _ = fi18
    }
    delete(v2, 0)
    v11 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
