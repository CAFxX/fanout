package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    var v1 int8 = 42
    v2 := []uint32{543, 531, 65535}
    v3 := make(map[uint32]bool)
    var v4 int8 = (v1 << ((18446744073709551615) & 63))
    var v5 uint64 = 711
    v6 := GenS1{F0: 42, F1: 65535}
    switch "hello" {
    case "foo bar":
        for fi7 := 0; fi7 < 7; fi7++ {
            switch (fi7 / ((4294967295) | 1)) {
            case -2147483648:
                for fi8 := 0; fi8 < 11; fi8++ {
                    v5 -= (^(v5))
                    v4 = (v1 * -116)
                    _ = fi8
                }
            }
            _ = fi7
        }
    case "hello":
        var v9 bool = (false || (65535 >= 1))
        _ = v9
    default:
        v5 += (v5 << ((2) & 63))
    }
    v10 := GenS1{F0: 100, F1: 2}
    v11 := GenS1{F0: (v4 << ((1) & 63)), F1: 65535}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 bool = true
        var v14 uint = 166
        v13 = (!(!false))
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 = 100
    v5 += 2
    v15 := GenS2{F0: ((-(v1)) >= v4), F1: -757}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
