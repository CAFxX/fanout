package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint8; F2 int }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 int64, p1 uint8) bool {
    if (1 > 42) {
        v6 := []int{481}
        _ = v6
    } else {
        v7 := []int{-682, -914}
        p1 = p1
        _ = v7
    }
    p1 += (p1 * 2)
    if (549 <= 0) {
        v8 := GenS1{F0: true, F1: (p1 ^ p1), F2: 2147483648}
        v9 := GenS2{F0: 1, F1: -2147483648, F2: (!(0.0 > float64(1.5)))}
        v10 := []int{128, 9223372036854775807, 1}
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return (p1 == (p1 << ((180) & 63)))
}
func genFn1() uint8 {
    var v11 float32 = float32((((((0.1 - float64(1e10)) * -2.25) / (0.0 + 1.5)) - (3.14159 * 0.1)) - (((0.1 / 1.0) + 0.0) / (float64(((0.1 - (-2.25 - 0.1)) - 3.14159)) * -2.25))))
    var v12 float64 = ((0.0 / 1e10) / (3.14159 - 0.1))
    var v13 int8 = 97
    for fi14 := 0; fi14 < 16; fi14++ {
        switch fi14 {
        case 687:
            if (!true) {
                v13 += (v13 ^ (-(v13)))
                fmt.Println(v13)
            } else {
                var v15 float32 = v11
                var v16 uint64 = 42
                _ = v15
                _ = v16
            }
        }
        v17 := []int{(fi14 + fi14)}
        _ = fi14
        _ = v17
    }
    _ = v11
    _ = v12
    _ = v13
    return 227
}

func main() {
    var v1 int8 = 100
    var v2 int64 = 9223372036854775807
    var v3 int = 660
    var v4 int32 = 0
    var v5 int32 = (v4 & 42)
    v4 = 0
    v1 = (v1 ^ v1)
    switch ("a" + "a") {
    case "hello":
        v5 = v4
    case "\n\t":
        for fi18 := 0; fi18 < 15; fi18++ {
            fmt.Println(fi18, v5, fi18)
            _ = fi18
        }
    }
    v19 := GenS3{F0: 83, F1: 72, F2: 128}
    for fi20 := 0; fi20 < 11; fi20++ {
        fmt.Println(v2, v5, v1)
        var v21 uint32 = 753
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: 571, F1: 127, F2: v3}
    var v23 float64 = (1e10 + -2.25)
    var v24 uint8 = 2
    v1 -= 0
    v25 := []int{(v3 % ((int(v24)) | 1)), v3, (v3 ^ v3), 9223372036854775807}
    fmt.Println(v24, v24, v4)
    var v26 int16 = -352
    v27 := GenS4{F0: 18446744073709551615, F1: (^(v1))}
    for fi28 := 0; fi28 < 20; fi28++ {
        v1 += v1
        var v29 uint8 = v24
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
