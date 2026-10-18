package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint; F2 bool }

func main() {
    v1 := []string{"", ""}
    v2 := make(map[int]int)
    var v3 uint = 4294967295
    var v4 bool = (!((0 > 9223372036854775807) || (-1 == -1)))
    if (1 <= 951) {
        if (0.1 >= 0.0) {
            v4 = v4
            v5 := GenS1{F0: 1, F1: (v3 ^ (^(v3))), F2: 18446744073709551615}
            v6 := GenS2{F0: v4}
            _ = v5
            _ = v6
        }
        v3 -= (v3 / ((v3) | 1))
        fmt.Println(v3, v3)
    }
    if v4 {
        var v7 int16 = -32768
        v8 := GenS1{F0: 2, F1: 65535, F2: 0}
        var v9 int64 = 42
        _ = v7
        _ = v8
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 int = 247
    v4 = (v4 || true)
    v11 := GenS2{F0: true}
    if (0 == 1) {
        for fi12 := 0; fi12 < 4; fi12++ {
            v2[(v10 >> ((779) & 63))] = (fi12 & (fi12 / ((32768) | 1)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (v10 >> ((v3) & 63)) {
            case 1:
                v13 := GenS2{F0: v4}
                _ = v13
            case 243:
                var v14 int8 = -128
                _ = v14
            case -9223372036854775808:
                _ = len(v1)
            }
            _ = fi12
        }
        v10 += v10
    } else {
        var v15 float32 = 1.5
        _ = len(v1)
        _ = v15
    }
    fmt.Println(v3, v3)
    v3 = (v3 & 4294967295)
    for fi16 := 0; fi16 < 2; fi16++ {
        switch "foo bar" {
        case "foo bar":
            var v17 float32 = float32((((0.1 + 3.14159) - (0.0 + (3.14159 - float64(((3.14159 + 1e10) * ((-2.25 + 3.14159) * ((1e10 / (3.14159 + 3.14159)) / -2.25))))))) - 1.5))
            _ = v17
        default:
            switch int(v3) {
            case 42:
                v2[v10] = fi16
            case -9223372036854775808:
                var v18 int8 = 100
                _ = v18
            case -556:
                var v19 float32 = (-2.25 / 0.1)
                _ = v19
            default:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
        }
        fmt.Println(fi16, v10, v3)
        _ = fi16
    }
    switch "a" {
    case "foo bar":
        if v4 {
            var v20 float32 = (float32((((1e10 + 0.1) * ((float64(float32((0.1 * 3.14159))) + (3.14159 * ((1.5 * 3.14159) * -2.25))) + 1.5)) * ((-2.25 + (3.14159 * -2.25)) - 0.1))) + 3.14159)
            v3 = v3
            for fi21 := 0; fi21 < 5; fi21++ {
                v20 -= v20
                v22 := GenS3{F0: (v4 || true), F1: (v3 - v3), F2: v4}
                var v23 uint32 = 2147483648
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        } else {
            switch (v10 << (((v3 ^ v3)) & 63)) {
            case -2147483648:
                if false {
                    var v24 int32 = 42
                    fmt.Println(v3, v10)
                    switch "foo bar" {
                    case "a":
                        v24 += (v24 * 42)
                    case "foo bar":
                        fmt.Println(v3, v3, v24)
                    case "":
                        v25 := v2[(^(v10))]
                        _ = v25
                    }
                    _ = v24
                }
            case 92:
                v26 := v2[-1]
                _ = v26
            }
            var v27 uint64 = 2
            _ = v27
        }
    }
    switch (v10 + -9223372036854775808) {
    case 1:
        v3 = v3
    case 410:
        fmt.Println(v10)
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        _ = len(v2)
        _ = fi28
    }
    var v29 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
