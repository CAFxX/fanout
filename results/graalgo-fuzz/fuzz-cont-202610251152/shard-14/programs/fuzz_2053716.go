package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int32 = -1
    var v2 int8 = -1
    v3 := make(map[int]int)
    v4 := []string{("" + "\n\t"), (("\n\t" + "foo bar") + "")}
    v5 := []uint32{}
    var v6 float64 = 0.0
    var v7 uint = 104
    switch 9223372036854775807 {
    case 2147483647:
        v8 := GenS1{F0: 2, F1: v1}
        _ = v8
    default:
        if ((((0.0 - 0.0) * 1e10) <= 3.14159) || (333 >= 42)) {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 float32 = (-2.25 - (3.14159 - 0.1))
                v11 := GenS1{F0: 1, F1: (^(v1))}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint64 = 4294967295
            _ = v12
        } else {
            v2 -= int8(v7)
            if ((uint64(v7) < 929) && true) {
                var v13 float32 = (1.5 / float32(((((float64(1.5) / 1.5) / (v6 - -2.25)) + 0.1) / 3.14159)))
                _ = v13
            }
            v6 += 0.1
        }
    }
    v2 = (v2 >> (((v7 | v7)) & 63))
    fmt.Println(v7)
    var v14 bool = (false || true)
    if ((v2 ^ v2) >= v2) {
        v7 += 18446744073709551615
        switch "hello" {
        case "":
            v15 := GenS1{F0: 185, F1: (v1 - (v1 - -295))}
            _ = v15
        }
        v7 += 255
    } else {
        var v16 uint64 = 1
        v1 = (v1 ^ (v1 << ((v7) & 63)))
        v7 = (v7 | uint(v16))
        _ = v16
    }
    v7 += v7
    if true {
        v17 := GenS1{F0: 32767, F1: v1}
        _ = v17
    } else {
        var v18 uint16 = 134
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int8 = (v2 >> ((v7) & 63))
            _ = fi19
            _ = v20
        }
        v14 = (int8(v1) > (-(v2)))
        _ = v18
    }
    v3[127] = 332
    var v21 uint32 = 2
    if (0 == 3) {
        var v22 int16 = 0
        fmt.Println(v21, v22)
        var v23 uint8 = 128
        _ = v22
        _ = v23
    } else {
        v1 = v1
        var v24 float32 = 0.0
        _ = v24
    }
    var v25 int32 = (v1 & (v1 * (v1 << ((v7) & 63))))
    v2 = int8(v21)
    if false {
        if (!(!(266 == 1))) {
            var v26 int16 = int16(v7)
            v7 += (^(v7))
            _ = v26
        }
        v2 = (v2 - -12)
    }
    v6 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
