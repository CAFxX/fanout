package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int16 = 42
    var v2 uint16 = 54
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 int = 127
    var v6 bool = (!(v5 > v5))
    v6 = false
    fmt.Println(v6, v2)
    v7 := GenS1{F0: "a"}
    for fi8 := 0; fi8 < 16; fi8++ {
        v1 -= (-(v1))
        v1 = (v1 / ((v1) | 1))
        for fi9 := 0; fi9 < 11; fi9++ {
            _ = len(v3)
            _ = fi9
        }
        _ = fi8
    }
    v1 += v1
    v6 = (v6 || (((1.5 / (float32(0.1) * 0.1)) + 3.14159) > float32(3.14159)))
    var v10 float32 = (0.0 * 1e10)
    v1 = 32767
    v4[v5] = 951
    fmt.Println(v1)
    if (726 > 4294967295) {
        v5 -= -9223372036854775808
        v11 := GenS1{F0: "foo bar"}
        var v12 uint16 = v2
        _ = v11
        _ = v12
    }
    switch (v5 - -897) {
    case 1:
        switch ("hello" + "foo bar") {
        case "a":
            _ = len(v4)
        case "hello":
            var v13 uint16 = (v2 + (^(v2)))
            _ = v13
        default:
            switch v5 {
            case 4294967295:
                v10 += v10
            case 128:
                v2 -= v2
            default:
                v1 += v1
            }
        }
    case 100:
        var v14 int16 = (v1 | v1)
        _ = v14
    default:
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := []int{(v5 * v5), -589}
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS2{F0: v6, F1: 226, F2: 0}
        v19 := GenS3{F0: -275, F1: 98}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
