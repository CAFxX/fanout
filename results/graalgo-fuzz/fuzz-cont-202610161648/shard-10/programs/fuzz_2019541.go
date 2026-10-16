package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint32 = 95
    var v2 int8 = 1
    var v3 uint64 = 4294967295
    var v4 int32 = 2147483647
    v2 += v2
    for fi5 := 0; fi5 < 9; fi5++ {
        switch ("hello" + "hello") {
        case "hello":
            fmt.Println(v4, v2)
        default:
            var v6 uint = 18446744073709551615
            _ = v6
        }
        v7 := GenS1{F0: 1000, F1: (v3 | (v3 << ((255) & 63)))}
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v1 += (v1 / ((2147483648) | 1))
        v2 = (v2 >> ((905) & 63))
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    v4 = v4
    v10 := GenS1{F0: 42, F1: (v3 | (^(v3)))}
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := []int{}
        switch ("a" + "") {
        case "foo bar":
            var v13 float32 = (0.1 * (0.1 / (float32(1.5) / (((((0.0 - (1.5 / ((1.5 / float32(1.5)) * -2.25))) / (((((0.1 / (3.14159 - 1.5)) / float32(1e10)) + 3.14159) * -2.25) / -2.25)) / 1.0) * (-2.25 + 0.0)) / 1e10))))
            _ = v13
        }
        v12 = append(v12, (fi11 << ((18446744073709551615) & 63)))
        _ = fi11
        _ = v12
    }
    v2 = v2
    v2 = 71
    v3 = 18446744073709551615
    v14 := []int{}
    _ = len(v14)
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15, v4, v4)
        v3 -= (v3 + uint64(v1))
        _ = fi15
    }
    var v16 int32 = -929
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
