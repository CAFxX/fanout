package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 string }
func genFn0(p0 uint16, p1 uint32, p2 int32) int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: false, F1: 766, F2: ("foo bar" + "hello")}
        _ = fi6
        _ = v7
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 409
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 255
    v3 := []uint32{4294967295, 1, 2147483647}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{2, 924}
    if (0 == 696) {
        v1 = (v1 | v1)
        var v8 int16 = 0
        delete(v4, 414)
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 uint8 = 42
            var v12 uint32 = 65535
            var v13 int8 = -31
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        switch 128 {
        case 9223372036854775807:
            var v14 float32 = ((0.0 + (0.0 - (((1e10 + -2.25) - (float32(0.1) * 1e10)) * ((((-2.25 * 1.5) - 0.0) / 1.5) * 0.0)))) / 1.0)
            _ = v14
        case 196:
            v15 := GenS1{F0: (0 != 997), F1: 231, F2: (("\n\t" + "a") + "")}
            _ = v15
        }
    case "hello":
        var v16 uint = 296
        _ = v16
    case "foo bar":
        v2 -= (^(v2))
    default:
        v4[2] = 2147483648
    }
    v2 += (v2 / ((1) | 1))
    if (((((1 >= 1) || (8 != 0)) && (42 <= 42)) && (int64(v1) <= 9223372036854775807)) || (-2147483648 == 953)) {
        var v17 int16 = 273
        fmt.Println(v2, v17, v17)
        var v18 uint32 = 42
        _ = v17
        _ = v18
    }
    v2 += v2
    v1 += v1
    var v19 int64 = -1
    switch 113 {
    case 1:
        v20 := GenS1{F0: (32768 < 32768), F1: 0, F2: "\n\t"}
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = 553
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
