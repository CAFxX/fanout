package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0(p0 uint32, p1 int8) uint32 {
    v4 := []int{-117, 276, -840}
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 5; fi6++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() int16 {
    v7 := []int{4294967295, -9223372036854775808, 0}
    if true {
        var v8 uint8 = 213
        _ = v8
    } else {
        var v9 uint16 = 42
        var v10 uint64 = 42
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 255}
    for fi12 := 0; fi12 < 17; fi12++ {
        v7 = append(v7, (fi12 ^ (fi12 >> ((228) & 63))))
        _ = fi12
    }
    _ = v11
    _ = v7
    return 1
}

func main() {
    v1 := []int{488}
    v2 := make(map[uint32]string)
    var v3 int8 = -45
    v13 := GenS2{F0: 0, F1: 2, F2: -73}
    var v14 bool = (false || true)
    var v15 float64 = 1e10
    switch "foo bar" {
    case "":
        switch "\n\t" {
        case "a":
            var v16 uint64 = 18446744073709551615
            _ = v16
        default:
            var v17 uint16 = 65535
            _ = v17
        }
    default:
        v3 -= v3
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 171, F1: 272, F2: (v3 + (v3 << ((2) & 63)))}
        _ = fi18
        _ = v19
    }
    v15 += (0.0 + ((v15 - float64((float32((float64(0.1) - 0.1)) * 1e10))) + 3.14159))
    v20 := GenS2{F0: 2, F1: 65535, F2: v3}
    var v21 int16 = -32768
    fmt.Println(v15)
    v22 := GenS3{F0: 0, F1: 2147483647, F2: 42}
    v15 = v15
    v15 = float64(-2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
