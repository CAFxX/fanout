package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0() int8 {
    var v5 int = 136
    switch "" {
    case "":
        var v6 uint = 18446744073709551615
        _ = v6
    case "\n\t":
        v5 += v5
    }
    v7 := []int{v5, (v5 >> ((4294967295) & 63))}
    v5 += -640
    if (!(v5 <= 512)) {
        var v8 int64 = 42
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v8
    }
    _ = v5
    _ = v7
    return 0
}
func genFn1(p0 uint16) uint64 {
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!false) {
            fmt.Println(p0, p0, fi10)
            v9 = (v9 / (((v9 * v9) / float64(0.1)) * ((v9 * 1.5) + v9)))
            p0 += (p0 >> ((80) & 63))
        }
        _ = fi10
    }
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (float64((3.14159 - (float32(0.1) - 0.1))) / (3.14159 / -2.25))
    v3 := []uint32{2}
    var v4 int16 = 0
    var v11 uint8 = 1
    v2 = float64(1e10)
    var v12 uint64 = 31
    v12 = (v12 | 0)
    fmt.Println(v2, v4)
    v3 = append(v3, 2147483647)
    v13 := GenS1{F0: 1, F1: (v11 << ((1) & 63)), F2: 0}
    v2 = v2
    fmt.Println(v12)
    fmt.Println(v2)
    v11 = v11
    v14 := GenS1{F0: 1, F1: v11, F2: 100}
    switch "\n\t" {
    case "a":
        switch "\n\t" {
        case "hello":
            if false {
                var v15 uint64 = (^(v12))
                var v16 float32 = 1e10
                _ = v15
                _ = v16
            }
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 float64 = float64((0.1 / 1.0))
            v4 -= 42
            if false {
                var v19 int32 = 2147483647
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: 315, F1: (v11 * v11), F2: -128}
        switch fi20 {
        case 32767:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 4294967295:
            v22 := GenS1{F0: 509, F1: v11, F2: -128}
            _ = v22
        case 9223372036854775807:
            v23 := GenS2{F0: 0, F1: 42}
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
