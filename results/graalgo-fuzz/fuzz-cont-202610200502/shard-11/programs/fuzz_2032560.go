package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 bool, p1 uint8) uint16 {
    var v7 bool = (false || (!(false || p0)))
    p0 = false
    _ = p0
    _ = p1
    _ = v7
    return 124
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint = 18446744073709551615
    var v5 float32 = 0.0
    v6 := make(map[int]uint32)
    var v8 string = "foo bar"
    fmt.Println(v8, v4)
    v9 := []int{}
    fmt.Println(v8)
    v10 := GenS1{F0: 65, F1: (false || false)}
    for fi11 := 0; fi11 < 12; fi11++ {
        v8 = v8
        _ = fi11
    }
    var v12 uint32 = 2147483647
    switch 0 {
    case 0:
        var v13 bool = false
        _ = v13
    default:
        switch 2147483647 {
        case 739:
            switch "foo bar" {
            case "\n\t":
                var v14 int16 = 0
                _ = v14
            case "hello":
                v4 -= (v4 << (((v4 - (v4 * v4))) & 63))
            }
        default:
            v5 += v5
        }
    }
    v8 = v8
    for fi15 := 0; fi15 < 15; fi15++ {
        switch v8 {
        case "foo bar":
            v1 = (-(v1))
        case "":
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int = (fi16 / ((1) | 1))
                var v18 bool = (!(32767 != 42))
                var v19 uint16 = 461
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        case "hello":
            v5 -= v5
        }
        var v20 uint8 = uint8(v12)
        var v21 uint16 = 42
        _ = fi15
        _ = v20
        _ = v21
    }
    v12 -= (v12 ^ v12)
    v8 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
