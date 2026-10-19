package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
func genFn0(p0 int16) uint32 {
    p0 -= p0
    v5 := []int{-9223372036854775808, -1, -2147483648, 2147483647}
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 uint64, p1 uint64, p2 int16) int {
    p2 -= (p2 - p2)
    p2 = -1
    v6 := GenS1{F0: 1}
    p1 -= (p0 % (((p1 + p0)) | 1))
    var v7 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []int{0, -2147483648, 9223372036854775807}
    var v2 uint8 = 60
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]bool)
    switch 32768 {
    case -2147483648:
        v8 := GenS1{F0: 778}
        _ = v8
    case 1:
        switch "foo bar" {
        case "":
            v9 := GenS1{F0: v3}
            _ = v9
        case "foo bar":
            v3 -= (v3 * v3)
        default:
            fmt.Println(v3)
        }
    case 423:
        fmt.Println(v3)
    }
    var v10 bool = (42 != 236)
    var v11 int16 = 0
    v12 := v4[9223372036854775807]
    v12 = (v2 <= v2)
    if v10 {
        delete(v4, 128)
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        switch (fi13 & (fi13 * fi13)) {
        case -1:
            for fi14 := 0; fi14 < 11; fi14++ {
                if (false || false) {
                    fmt.Println(v11, v3)
                    var v15 uint64 = 42
                    var v16 float32 = ((float32(-2.25) + ((0.0 / 1e10) + (0.0 - (0.0 - 3.14159)))) / float32((-2.25 + float64(3.14159))))
                    _ = v15
                    _ = v16
                }
                if (458 == 0) {
                    if (true && false) {
                        v2 -= (v2 ^ v2)
                    } else {
                        v4[(fi13 + fi14)] = ((("" + (("a" + "foo bar") + "\n\t")) < "a") && false)
                    }
                    v2 = 9
                }
                v17 := GenS2{F0: 2, F1: (v2 >> ((1) & 63))}
                _ = fi14
                _ = v17
            }
        case -506:
            switch "a" {
            case "":
                var v18 float32 = (float32(float64((3.14159 / (float32((1e10 / float64(1.5))) + float32((1.5 * ((1e10 * ((0.1 * 1.5) * ((0.0 - float64(3.14159)) + 0.0))) / 1e10))))))) / -2.25)
                _ = v18
            case "\n\t":
                v19 := GenS1{F0: v3}
                _ = v19
            case "hello":
                var v20 int16 = -32768
                _ = v20
            }
        case 4294967295:
            if (718 < 255) {
                fmt.Println(v3, fi13, v10)
            } else {
                var v21 uint8 = 255
                var v22 int16 = (v11 ^ v11)
                _ = v21
                _ = v22
            }
        }
        fmt.Println(fi13, v11)
        _ = fi13
    }
    var v23 float64 = 0.0
    for fi24 := 0; fi24 < 18; fi24++ {
        if ((1e10 * ((3.14159 + 1e10) * (1.5 / (((3.14159 / (float32(v23) * (1.5 + (((1e10 + 0.0) / 1e10) / 0.1)))) / -2.25) / 1.0)))) == (0.1 + 1.5)) {
            if (42 != 18446744073709551615) {
                fmt.Println(v10)
                fmt.Println(v10)
            } else {
                v12 = (18446744073709551615 >= 643)
                v2 = (v2 * 128)
            }
        } else {
            var v25 bool = (uint64(v2) <= 0)
            _ = v25
        }
        _ = fi24
    }
    switch ("hello" + "foo bar") {
    case "a":
        v26 := GenS1{F0: -9223372036854775808}
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v3, v2, fi27)
        _ = fi27
    }
    v28 := GenS3{F0: (v12 && (("" + "foo bar") <= "")), F1: 32768, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
}
