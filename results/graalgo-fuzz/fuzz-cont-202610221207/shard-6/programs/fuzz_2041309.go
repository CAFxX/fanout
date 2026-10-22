package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint64, p1 bool, p2 int32) uint64 {
    var v4 bool = true
    var v5 int32 = p2
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{9223372036854775807, (fi7 % ((fi7) | 1)), 2147483647, fi7}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (^(p0))
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    v1 -= (-(v1))
    if true {
        fmt.Println(v1)
        v1 -= (v1 / ((0) | 1))
        v9 := GenS1{F0: 725, F1: 100}
        _ = v9
    } else {
        var v10 uint = 255
        var v11 int64 = 9223372036854775807
        _ = v10
        _ = v11
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v12 uint = 103
        _ = v12
    case "":
        v1 = v1
    case "hello":
        v13 := GenS2{F0: 204, F1: 1}
        _ = v13
    default:
        v2[519] = v1
    }
    if (-264 > 42) {
        var v14 float32 = (1e10 - 0.1)
        var v15 int16 = 42
        _ = v14
        _ = v15
    } else {
        fmt.Println(v1)
        delete(v2, 65535)
    }
    v16 := GenS1{F0: 405, F1: 42}
    v1 = v1
    fmt.Println(v1)
    for fi17 := 0; fi17 < 16; fi17++ {
        switch (fi17 | (fi17 | fi17)) {
        case 2147483647:
            var v18 uint = 412
            _ = v18
        case 0:
            v19 := []int{249, (fi17 + (fi17 % (((fi17 + (-(v1)))) | 1))), v1, (^(v1))}
            _ = v19
        case 766:
            v20 := GenS2{F0: 82, F1: 0}
            _ = v20
        }
        var v21 uint64 = 18446744073709551615
        fmt.Println(v1, v1)
        _ = fi17
        _ = v21
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
