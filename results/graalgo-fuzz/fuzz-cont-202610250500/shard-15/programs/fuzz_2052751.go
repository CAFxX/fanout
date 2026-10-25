package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 uint64, p1 int32, p2 uint32) int32 {
    if (!true) {
        var v7 uint32 = 813
        fmt.Println(p2, p0)
        _ = v7
    } else {
        p1 -= (p1 >> ((0) & 63))
        var v8 uint8 = 1
        _ = v8
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ p1)
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 0
    var v3 uint = 1
    v4 := make(map[string]bool)
    var v5 float64 = (((((1e10 * float64(1e10)) + float64(3.14159)) + float64(-2.25)) - 0.0) / 1.5)
    var v6 int = 536
    v9 := GenS1{F0: 42, F1: 756}
    v5 += 0.1
    fmt.Println(v5)
    var v10 string = ("foo bar" + "hello")
    switch "" {
    case "hello":
        for fi11 := 0; fi11 < 13; fi11++ {
            v1["hello"] = "\n\t"
            var v12 int32 = 0
            fmt.Println(v10, v6)
            _ = fi11
            _ = v12
        }
    case "\n\t":
        v13 := GenS2{F0: 255, F1: 1, F2: (65535 == 2147483647)}
        _ = v13
    case "":
        switch "a" {
        case "":
            switch v6 {
            case -9223372036854775808:
                v14 := GenS1{F0: -347, F1: 0}
                _ = v14
            default:
                var v15 int64 = -1
                _ = v15
            }
        case "\n\t":
            v16 := GenS1{F0: 9223372036854775807, F1: 2}
            _ = v16
        case "hello":
            v17 := v4[v10]
            _ = v17
        default:
            var v18 bool = true
            _ = v18
        }
    default:
        var v19 uint = v3
        _ = v19
    }
    v20 := GenS2{F0: 255, F1: 1, F2: true}
    fmt.Println(v5, v3, v3)
    v2 = (v2 ^ (v2 << ((v3) & 63)))
    v21 := GenS3{F0: ((!(0 != 506)) || true)}
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 string = v10
        if (3.14159 > (3.14159 * 0.1)) {
            v3 -= (v3 % ((0) | 1))
            var v24 int8 = 0
            _ = v24
        } else {
            switch (v6 ^ (v6 ^ 62)) {
            case 9223372036854775807:
                v10 = (v23 + "\n\t")
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
