package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int16; F1 int; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 2, 2147483647}
    var v3 uint = 817
    v4 := []string{"\n\t", "foo bar", "\n\t"}
    var v5 int16 = 1000
    v6 := GenS1{F0: 100, F1: 100}
    if true {
        v7 := GenS1{F0: 30, F1: 0}
        v8 := GenS2{F0: 886, F1: 9223372036854775807}
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                if false {
                    v11 := v1[2147483647]
                    v3 += v3
                    var v12 int8 = -1
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS3{F0: (v5 / ((v5) | 1)), F1: int(v3), F2: 42}
                    v3 -= 18446744073709551615
                    v5 = (^(v5))
                    _ = v13
                }
                _ = fi10
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v5)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi15
        }
        _ = len(v1)
        fmt.Println(fi14)
        _ = fi14
    }
    v3 = 22
    v16 := GenS4{F0: 1, F1: (v3 << ((v3) & 63))}
    if (int32(v3) >= 95) {
        var v17 uint8 = 255
        _ = v17
    } else {
        var v18 uint = 65535
        v18 = (^(v3))
        _ = len(v4)
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int32 = 0
        v3 = v3
        v5 = 42
        _ = fi19
        _ = v20
    }
    _ = len(v1)
    fmt.Println(v5, v3)
    var v21 int32 = -362
    v4 = append(v4, ("foo bar" + "foo bar"))
    switch "hello" {
    case "hello":
        v22 := GenS4{F0: 4294967295, F1: (v3 | v3)}
        _ = v22
    default:
        _ = len(v2)
    }
    var v23 uint64 = 82
    var v24 float32 = -2.25
    v25 := GenS3{F0: (-(v5)), F1: -195, F2: 127}
    var v26 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
