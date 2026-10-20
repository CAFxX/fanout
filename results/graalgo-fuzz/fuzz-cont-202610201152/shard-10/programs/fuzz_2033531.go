package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8) int16 {
    v6 := GenS1{F0: int16(p0), F1: 659, F2: 32767}
    switch ("hello" + "foo bar") {
    case "hello":
        p0 = 131
    case "":
        var v7 int64 = -9223372036854775808
        _ = v7
    case "\n\t":
        var v8 uint8 = 127
        _ = v8
    }
    p0 += p0
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -32768
}

func main() {
    var v1 int32 = 101
    v2 := make(map[uint32]int)
    var v3 int = -582
    v4 := make(map[uint32]uint32)
    var v5 float64 = (0.0 - (1e10 - float64(1.5)))
    var v9 float32 = 0.0
    var v10 int8 = 86
    v11 := GenS2{F0: 555, F1: (-(v1)), F2: false}
    var v12 uint64 = 4294967295
    if false {
        v13 := GenS2{F0: genFn0(127), F1: v1, F2: false}
        _ = v13
    }
    v14 := []int{977, v3, (v3 - v3), (v3 << ((18446744073709551615) & 63))}
    fmt.Println(v5, v3, v10)
    v5 = v5
    var v15 uint64 = v12
    var v16 int8 = v10
    fmt.Println(v3)
    switch (v3 - v3) {
    case -9223372036854775808:
        if (255 <= 2) {
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS3{F0: 9223372036854775807, F1: (v1 - (-(v1)))}
                _ = fi17
                _ = v18
            }
        }
    default:
        v14 = append(v14, (v3 ^ (^(v3))))
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            v10 -= v16
        default:
            _ = len(v2)
        }
        var v19 int = v3
        v15 += (^(v15))
        _ = v19
    }
    if false {
        _ = len(v14)
    }
    fmt.Println(v12, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
