package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 uint8) bool {
    fmt.Println(p0)
    var v5 uint32 = 2147483647
    _ = p0
    _ = v5
    return true
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]uint32)
    var v3 uint32 = 456
    v4 := make(map[uint32]uint32)
    switch ("\n\t" + "hello") {
    case "\n\t":
        v6 := v2["a"]
        _ = v6
    case "hello":
        var v7 float64 = (((((0.1 + (1.5 + (-2.25 * (((((0.0 * (((1.5 * 0.0) - (3.14159 / 1e10)) / 0.1)) / float64((((1.5 - ((0.0 / -2.25) * 3.14159)) - float32(3.14159)) - float32(0.1)))) / float64((float32(float64(1e10)) + ((float32(0.1) - (0.1 - 0.0)) + (1e10 + 1.5))))) + -2.25) + (3.14159 * 0.0))))) / (((0.1 / (1e10 + ((1.5 / 3.14159) / 1.0))) + (-2.25 - 1e10)) + ((1.5 + (0.1 / 3.14159)) - float64((float32((0.0 - 0.0)) - 1e10))))) * -2.25) + float64(3.14159)) - 1.5)
        _ = v7
    case "a":
        switch 128 {
        case 32767:
            v1 = (v1 | 42)
        case 0:
            v8 := v2[("foo bar" + "a")]
            _ = v8
        }
    default:
        var v9 bool = ((v1 == v1) && ((1e10 * 0.0) > 1e10))
        _ = v9
    }
    v1 = (v1 * 155)
    v10 := []int{42, -9223372036854775808, -477, 9223372036854775807}
    v3 = (v3 << ((v1) & 63))
    for fi11 := 0; fi11 < 15; fi11++ {
        v1 += 65535
        switch "hello" {
        case "":
            v12 := GenS1{F0: -1, F1: (v1 % ((255) | 1)), F2: v1}
            _ = v12
        case "foo bar":
            var v13 uint16 = 652
            _ = v13
        default:
            v1 += (^(v1))
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            v3 += v3
            if len(v10) > 0 {
                v10[0] = -9223372036854775808
            }
            _ = fi14
        }
        _ = fi11
    }
    fmt.Println(v1, v3)
    switch (("" + "") + "") {
    case "hello":
        v15 := GenS2{F0: v3}
        _ = v15
    default:
        var v16 int32 = 42
        _ = v16
    }
    var v17 bool = true
    v18 := GenS3{F0: ("hello" + ""), F1: 2}
    switch 42 {
    case 649:
        var v19 int16 = -335
        _ = v19
    case 128:
        fmt.Println(v3)
    case -9223372036854775808:
        if false {
            if v17 {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v17 = (1 > -2147483648)
                    var v21 int32 = 2147483647
                    for fi22 := 0; fi22 < 15; fi22++ {
                        fmt.Println(v21)
                        var v23 int = 128
                        fmt.Println(v23, v1, v17)
                        v24 := GenS4{F0: -9223372036854775808, F1: 127}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
            }
            var v25 uint16 = 650
            v3 = (^(v3))
            _ = v25
        }
    }
    var v26 int8 = 119
    var v27 bool = (42 > 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
