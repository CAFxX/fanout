package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]int)
    var v3 int64 = 886
    var v4 float32 = float32(((-2.25 / 1.0) + 0.1))
    v5 := []uint32{65535}
    var v6 int8 = -30
    v4 = (v4 * (v4 / (float32(3.14159) / v4)))
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 646
        var v9 float32 = v4
        v10 := GenS1{F0: ((-(fi7)) < -972), F1: v3, F2: 42}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := v2[2147483648]
        _ = fi11
        _ = v12
    }
    v1 = true
    fmt.Println(v6)
    _ = len(v5)
    switch "" {
    case "a":
        v13 := GenS1{F0: (2147483647 > 2147483648), F1: v3, F2: 57}
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 9; fi14++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v1)
            case "hello":
                var v15 uint = 65535
                _ = v15
            case "a":
                v16 := GenS1{F0: ((!((3.14159 - float64(v4)) == float64((v4 - (v4 / (float32(3.14159) * v4)))))) || true), F1: v3, F2: 93}
                _ = v16
            }
            v3 -= v3
            _ = len(v5)
            _ = fi14
        }
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if v1 {
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v4, fi17)
            case "":
                v4 = 1.5
            case "hello":
                var v18 uint = 1
                _ = v18
            default:
                v19 := GenS1{F0: false, F1: v3, F2: 34}
                _ = v19
            }
            v6 = (^(v6))
            var v20 uint8 = 255
            _ = v20
        }
        v21 := GenS1{F0: v1, F1: -33, F2: 128}
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        _ = len(v5)
        v4 -= v4
        _ = fi22
    }
    switch "" {
    case "hello":
        var v23 uint16 = 324
        _ = v23
    }
    fmt.Println(v1)
    _ = len(v5)
    for fi24 := 0; fi24 < 7; fi24++ {
        v4 += 1e10
        _ = fi24
    }
    v1 = v1
    v1 = true
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
