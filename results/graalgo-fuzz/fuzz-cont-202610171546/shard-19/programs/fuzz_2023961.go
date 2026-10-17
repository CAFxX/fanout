package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]int)
    var v3 int = 720
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: 65535, F1: 22}
    fmt.Println(v3, v3)
    if (128 >= 36) {
        for fi6 := 0; fi6 < 15; fi6++ {
            v3 = -856
            var v7 int = (fi6 << ((1) & 63))
            _ = fi6
            _ = v7
        }
        var v8 int32 = 0
        if false {
            switch "hello" {
            case "\n\t":
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 bool = true
                    _ = fi9
                    _ = v10
                }
            case "a":
                if (-128 <= -6) {
                    v8 += 42
                }
            default:
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(fi11, v3)
                    _ = fi11
                }
            }
            var v12 int64 = -750
            _ = v12
        }
        _ = v8
    }
    var v13 int = (v3 << ((255) & 63))
    v1 -= v1
    fmt.Println(v1, v1)
    for fi14 := 0; fi14 < 1; fi14++ {
        switch "\n\t" {
        case "":
            for fi15 := 0; fi15 < 20; fi15++ {
                v13 = v3
                v13 = 2147483647
                _ = fi15
            }
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                fmt.Println(v13, v1)
                v17 := GenS1{F0: 2, F1: -128}
                v13 = -9223372036854775808
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
    }
    fmt.Println(v13)
    v13 -= 0
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v13, v1)
        v4[2147483647] = ("" + ("foo bar" + "\n\t"))
        var v19 bool = true
        _ = fi18
        _ = v19
    }
    v1 = float32((1.5 + ((1e10 - 3.14159) / float64(((v1 + (3.14159 + ((1.5 - v1) * (v1 * v1)))) / (v1 + float32(float64(0.1))))))))
    for fi20 := 0; fi20 < 3; fi20++ {
        if (false && true) {
            v21 := []int{}
            _ = v21
        }
        switch ("hello" + "\n\t") {
        case "a":
            var v22 int16 = -330
            _ = v22
        case "":
            var v23 int32 = -2147483648
            _ = v23
        }
        switch "" {
        case "hello":
            v3 -= (v3 & (fi20 - v3))
        default:
            var v24 int64 = 42
            _ = v24
        }
        _ = fi20
    }
    v25 := GenS1{F0: 2, F1: 127}
    v1 = (0.1 + (((float32(0.0) + 3.14159) / v1) / (1e10 * (v1 / v1))))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
