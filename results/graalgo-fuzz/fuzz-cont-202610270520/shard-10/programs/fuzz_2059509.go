package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 0.1
    var v3 int32 = 1
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{(fi5 | fi5), 1, 4294967295, fi5}
        _ = fi5
        _ = v6
    }
    if (42 != 75) {
        v4 -= (v4 - 790)
        switch "a" {
        case "\n\t":
            var v7 float64 = (1.5 * 1.5)
            _ = v7
        case "foo bar":
            for fi8 := 0; fi8 < 10; fi8++ {
                switch (fi8 * fi8) {
                case -814:
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 float32 = 1e10
                        _ = fi9
                        _ = v10
                    }
                }
                _ = fi8
            }
        default:
            v1 += 359
        }
    }
    var v11 int = -1
    fmt.Println(v4, v2, v3)
    v2 += (0.0 + 0.0)
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            v13 := []int{v11, fi12, (fi12 ^ -2147483648), (v11 ^ (-(fi12)))}
            _ = v13
        }
        _ = fi12
    }
    v3 = v3
    switch (v11 ^ 128) {
    case -13:
        v14 := GenS1{F0: v4, F1: "foo bar"}
        _ = v14
    default:
        v11 -= (v11 << ((18446744073709551615) & 63))
    }
    v15 := []int{760}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
