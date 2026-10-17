package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }

func main() {
    var v1 uint = 42
    v2 := []string{""}
    var v3 string = "hello"
    v4 := make(map[int]uint32)
    var v5 uint32 = 65535
    v6 := GenS1{F0: (true && true), F1: 255, F2: -444}
    if (false && false) {
        v1 = v1
    }
    var v7 uint32 = v5
    v8 := GenS1{F0: ((-677 <= -2147483648) && (true && false)), F1: 255, F2: 0}
    v3 = (v3 + "foo bar")
    switch "foo bar" {
    case "foo bar":
        var v9 float32 = 3.14159
        _ = v9
    case "a":
        v5 = 42
    }
    v7 -= (v7 + 2)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int64 = -783
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v3)
    for fi14 := 0; fi14 < 1; fi14++ {
        if (0.0 <= (-2.25 * -2.25)) {
            switch fi14 {
            case -2147483648:
                v5 -= v5
            case 2147483647:
                v15 := GenS2{F0: 64, F1: 9223372036854775807, F2: 0}
                _ = v15
            default:
                v4[(fi14 << (((v1 / ((v1) | 1))) & 63))] = 2147483647
            }
        }
        v4[(^(fi14))] = 497
        _ = fi14
    }
    if (((32767 <= -949) && false) && (uint16(v1) != 538)) {
        switch 4294967295 {
        case 32767:
            var v16 int32 = 110
            _ = v16
        }
        _ = len(v4)
        v17 := v4[-393]
        _ = v17
    }
    var v18 uint32 = (v7 & v5)
    fmt.Println(v18)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
