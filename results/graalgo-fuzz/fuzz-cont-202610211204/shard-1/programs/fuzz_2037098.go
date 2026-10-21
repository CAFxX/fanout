package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    v3 -= (-(v3))
    v4 := GenS1{F0: 1000, F1: 255, F2: (v3 >> ((454) & 63))}
    v5 := GenS2{F0: (v3 & v3), F1: 4294967295, F2: 264}
    switch ("a" + "a") {
    case "\n\t":
        if true {
            fmt.Println(v3)
            fmt.Println(v3)
            v6 := GenS1{F0: 398, F1: 1, F2: (v3 | -9223372036854775808)}
            _ = v6
        } else {
            var v7 uint16 = 835
            var v8 uint32 = 4294967295
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: 803, F1: 128, F2: (v3 - (^(v3)))}
                v11 := GenS3{F0: v3, F1: (0.1 == 3.14159)}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v7
            _ = v8
        }
    case "":
        v12 := []int{}
        _ = v12
    }
    var v13 int64 = v3
    switch "foo bar" {
    case "hello":
        v3 += 1
    case "\n\t":
        var v14 string = "hello"
        _ = v14
    default:
        if ((42 < 0) && (false || (2 == 2))) {
            var v15 int32 = -1
            var v16 uint16 = 1
            _ = v15
            _ = v16
        }
    }
    var v17 float32 = -2.25
    if (!false) {
        switch ("foo bar" + "hello") {
        case "a":
            for fi18 := 0; fi18 < 13; fi18++ {
                for fi19 := 0; fi19 < 16; fi19++ {
                    v20 := GenS2{F0: v3, F1: 4294967295, F2: 42}
                    _ = fi19
                    _ = v20
                }
                v3 = 794
                _ = fi18
            }
        case "\n\t":
            var v21 float64 = (3.14159 * 1.5)
            _ = v21
        case "foo bar":
            _ = len(v1)
        default:
            v22 := GenS1{F0: 0, F1: 58, F2: v13}
            _ = v22
        }
        _ = len(v1)
    }
    var v23 int64 = (v13 + (^(v13)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
