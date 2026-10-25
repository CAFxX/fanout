package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 int64 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int; F1 string }

func main() {
    var v1 uint8 = 42
    var v2 string = ("a" + "a")
    var v3 int16 = 32767
    v4 := make(map[string]int)
    var v5 int = 128
    var v6 uint8 = v1
    v7 := GenS1{F0: 0, F1: -2147483648}
    v8 := GenS2{F0: 2147483647, F1: 4294967295}
    if true {
        var v9 bool = (!(1 >= 42))
        v4[""] = (v5 - (v5 / (((v5 - v5)) | 1)))
        _ = v9
    }
    var v10 int = v5
    v10 = (v5 % ((v5) | 1))
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            v10 += (v10 & 32767)
            _ = fi11
        }
    } else {
        v2 = ((v2 + v2) + (v2 + (v2 + v2)))
        if (!(((v5 > int(v1)) && (false || false)) || true)) {
            switch (v5 % ((2147483647) | 1)) {
            case -1:
                v12 := GenS3{F0: 9223372036854775807}
                _ = v12
            case -671:
                var v13 float32 = (0.0 / (0.0 + 3.14159))
                _ = v13
            case -9223372036854775808:
                v5 -= (v10 ^ (^(v10)))
            default:
                v5 -= v10
            }
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v3 = 42
        v15 := GenS4{F0: 2, F1: -128, F2: 42}
        _ = fi14
        _ = v15
    }
    v4[v2] = v5
    if false {
        var v16 int16 = (v3 & -32768)
        var v17 uint16 = 1
        _ = v16
        _ = v17
    }
    var v18 uint16 = uint16(v6)
    v19 := []int{v5, (v10 ^ (v10 + int(v1))), v10}
    if ((0.0 + 3.14159) != -2.25) {
        if len(v19) > 0 {
            v19[0] = -574
        }
        var v20 float32 = (float32((((-2.25 / (1.5 + (0.1 + ((0.0 + (3.14159 * float64(3.14159))) * 1e10)))) - (0.0 - 0.0)) / (1e10 / 1e10))) - 0.1)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v1, v2, v18)
            fmt.Println(fi21, fi21)
            _ = fi21
        }
        v1 -= (v6 >> ((18446744073709551615) & 63))
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v6, v10)
            delete(v4, ("" + (("foo bar" + v2) + "hello")))
            _ = fi22
        }
    }
    v4[("a" + v2)] = (-(v5))
    var v23 int = (-(v10))
    fmt.Println(v23)
    v24 := GenS5{F0: 128, F1: ((^(v18)) <= (v18 >> ((4294967295) & 63)))}
    v25 := GenS6{F0: v5, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
