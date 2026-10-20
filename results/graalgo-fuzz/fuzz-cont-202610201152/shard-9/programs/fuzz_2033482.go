package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 int16) uint32 {
    v5 := GenS1{F0: (("hello" + "") + "hello"), F1: 655, F2: (p0 - (^(p0)))}
    v6 := GenS1{F0: (("foo bar" + "a") + "a"), F1: 4294967295, F2: 2147483648}
    if p1 {
        var v7 uint16 = 255
        switch "" {
        case "a":
            fmt.Println(p2, p2)
        case "foo bar":
            fmt.Println(v7, p1)
        default:
            switch "hello" {
            case "hello":
                v8 := []int{1, -914}
                _ = v8
            default:
                v9 := []int{-9223372036854775808, 2147483647}
                _ = v9
            }
        }
        var v10 int32 = 0
        _ = v10
        _ = v7
    }
    v11 := GenS2{F0: 65535, F1: 18446744073709551615, F2: 255}
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    _ = v6
    return (p0 + p0)
}

func main() {
    var v1 bool = ((("hello" + "") + "foo bar") < "")
    v2 := make(map[int]bool)
    v3 := []uint32{2, 878}
    v4 := []uint32{4294967295, 65535, 34, 0}
    fmt.Println(v1)
    v1 = v1
    for fi12 := 0; fi12 < 11; fi12++ {
        if (genFn0(1, (!true), 17) <= 354) {
            if (2 >= 18446744073709551615) {
                v1 = (2147483647 <= -1)
            } else {
                if (42 > 1) {
                    v1 = (!(!(42 >= -2147483648)))
                    v1 = v1
                    var v13 uint32 = 4294967295
                    _ = v13
                } else {
                    v1 = ("" != "hello")
                    v1 = v1
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v14 float32 = (1e10 / (((3.14159 * 1.5) * -2.25) * (1.5 / (((1e10 / ((1.5 - (((0.1 * 1.5) + -2.25) - (((0.0 * 3.14159) * 3.14159) * float32(((0.1 / 1.0) + (1.5 * (3.14159 - 3.14159))))))) / 0.1)) + (1.5 / ((1.5 + float32(3.14159)) / (0.1 * 3.14159)))) * -2.25))))
                _ = v14
            }
        }
        var v15 uint32 = 2
        _ = fi12
        _ = v15
    }
    var v16 uint16 = 42
    if (128 <= 916) {
        v17 := GenS2{F0: v16, F1: 0, F2: 255}
        var v18 uint64 = 902
        _ = v17
        _ = v18
    }
    fmt.Println(v1, v16)
    fmt.Println(v1)
    var v19 int = -1
    v20 := GenS1{F0: "a", F1: 18446744073709551615, F2: genFn0(634, ((0 <= -1) || false), 278)}
    _ = len(v3)
    var v21 uint32 = genFn0(2147483647, (!v1), -73)
    var v22 int64 = 1
    v16 = (v16 | v16)
    var v23 int16 = int16(v19)
    v19 = v19
    for fi24 := 0; fi24 < 7; fi24++ {
        v16 += v16
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
