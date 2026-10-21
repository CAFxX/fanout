package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 string) int8 {
    p0 = "a"
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int32 = 0
        _ = fi4
        _ = v5
    }
    v6 := []int{-9223372036854775808, 948, 0}
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 string = (("\n\t" + "hello") + ("\n\t" + "hello"))
    v2 := make(map[string]int)
    var v3 int = 128
    if (!true) {
        v3 = v3
        delete(v2, "foo bar")
    } else {
        var v7 uint32 = 1
        v7 = v7
        _ = v7
    }
    v3 -= v3
    if (113 < genFn0(v1)) {
        v8 := GenS1{F0: false, F1: 522, F2: ((false || (168 >= 65535)) && (!false))}
        fmt.Println(v3, v1)
        v9 := []int{(v3 << ((1) & 63)), v3, (-(v3)), 42}
        _ = v8
        _ = v9
    }
    v3 = 786
    v10 := []int{(v3 / (((v3 << ((2) & 63))) | 1)), 32767}
    var v11 bool = false
    if v11 {
        if v11 {
            switch v3 {
            case 42:
                v12 := GenS1{F0: true, F1: -2147483648, F2: (-9223372036854775808 > v3)}
                _ = v12
            case 128:
                fmt.Println(v1, v3, v1)
            default:
                v13 := GenS1{F0: false, F1: 42, F2: v11}
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int8 = 0
                _ = fi14
                _ = v15
            }
            var v16 int = 0
            if v11 {
                fmt.Println(v3, v11)
            }
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 float32 = float32((float64(0.1) - (((3.14159 * (3.14159 - float64(3.14159))) + (3.14159 - 3.14159)) + -2.25)))
            v11 = (673 >= 42)
            if ((v18 > (v18 + v18)) || (32767 == -32768)) {
                v1 = (v1 + ("foo bar" + "a"))
            }
            _ = fi17
            _ = v18
        }
    }
    v3 -= -1
    v19 := GenS2{F0: -2147483648}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v20 := GenS1{F0: false, F1: -2147483648, F2: v11}
    v21 := GenS3{F0: v3, F1: 127, F2: 286}
    v22 := GenS3{F0: v3, F1: -128, F2: 474}
    v11 = (((1 <= 4294967295) || (2 < 44)) || false)
    v3 -= (v3 * (-(v3)))
    var v23 int8 = -1
    switch (v3 / (((v3 + v3)) | 1)) {
    case 883:
        if (9223372036854775807 == 9223372036854775807) {
            for fi24 := 0; fi24 < 18; fi24++ {
                v23 = (v23 * v23)
                _ = fi24
            }
            v2[v1] = 4294967295
            v25 := GenS4{F0: (1.5 < (((0.0 * (float32(0.1) / (0.0 + float32((float64(3.14159) + (3.14159 + (3.14159 + 0.1))))))) * -2.25) - 0.0)), F1: 18446744073709551615}
            _ = v25
        }
    case -730:
        fmt.Println(v11)
    case 9223372036854775807:
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 int16 = 1
            v1 = (v1 + v1)
            _ = fi26
            _ = v27
        }
    default:
        if v11 {
            if (-1 < -32768) {
                var v28 float32 = (0.1 * -2.25)
                var v29 string = v1
                fmt.Println(v29, v29, v3)
                _ = v28
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
