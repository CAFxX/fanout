package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int16 = 32767
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 2
    v1 += -903
    v1 -= v1
    v2 = v2
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int32 = -2147483648
        _ = fi5
        _ = v6
    }
    v4 += v4
    switch "hello" {
    case "":
        var v7 uint32 = 2147483648
        _ = v7
    case "a":
        if (!(42 != 1)) {
            v2 -= (^(v2))
        } else {
            if (!false) {
                var v8 uint = 1
                _ = v8
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := []int{(fi10 ^ (fi10 >> ((4294967295) & 63)))}
                    var v12 float64 = (1e10 * ((0.0 / (float64(3.14159) - 1.5)) + (0.1 * (float64(0.0) * 1.5))))
                    v2 -= v2
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := v3[0]
                _ = fi9
                _ = v13
            }
            var v14 float64 = 1.5
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    v16 := GenS1{F0: (0 < 116)}
    switch ("" + "foo bar") {
    case "a":
        v17 := GenS2{F0: 1}
        _ = v17
    case "":
        var v18 int = -661
        _ = v18
    case "hello":
        fmt.Println(v2, v1, v1)
    default:
        v4 -= 255
    }
    v2 = (v2 & v2)
    v2 += v2
    v19 := []int{int(v1)}
    v20 := GenS3{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
