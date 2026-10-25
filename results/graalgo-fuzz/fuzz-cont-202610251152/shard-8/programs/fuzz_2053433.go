package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 bool) uint32 {
    switch "\n\t" {
    case "hello":
        switch "a" {
        case "hello":
            var v6 uint32 = 4294967295
            _ = v6
        case "\n\t":
            var v7 int32 = -553
            _ = v7
        }
    default:
        v8 := GenS1{F0: 32767}
        _ = v8
    }
    if (p0 && (!false)) {
        if p0 {
            v9 := GenS2{F0: 18446744073709551615}
            _ = v9
        }
    }
    v10 := GenS2{F0: 18446744073709551615}
    v11 := GenS1{F0: 32768}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := []int{-2147483648, -222}
        v14 := GenS3{F0: -2147483648}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 int = 42
    v2 := []int{2147483647, (^(v1)), 51, 127}
    var v3 int8 = 0
    var v4 float32 = 3.14159
    var v5 uint32 = 65535
    v3 = v3
    for fi15 := 0; fi15 < 17; fi15++ {
        switch fi15 {
        case 0:
            if ((v1 + fi15) >= v1) {
                fmt.Println(v4, v3, fi15)
                var v16 bool = (v3 < int8(fi15))
                _ = v16
            }
        case 127:
            v4 = v4
        }
        v17 := GenS4{F0: 1, F1: 0}
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint8 = 127
            _ = fi18
            _ = v19
        }
        _ = fi15
        _ = v17
    }
    v5 += v5
    fmt.Println(v3, v3, v5)
    if true {
        var v20 uint8 = 63
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 float64 = ((((1.5 * float64(float32(3.14159))) - 0.0) + -2.25) - float64((((v4 + v4) * 3.14159) + ((-2.25 - 1e10) + v4))))
            var v23 int64 = 9223372036854775807
            _ = fi21
            _ = v22
            _ = v23
        }
        if len(v2) > 0 {
            v2[0] = (v1 * (v1 ^ (v1 & -2147483648)))
        }
        _ = v20
    } else {
        fmt.Println(v5)
        if (65535 < 1) {
            _ = len(v2)
            fmt.Println(v4, v5, v3)
        }
    }
    var v24 uint = 395
    v4 = (v4 / 0.1)
    fmt.Println(v5, v1)
    v4 -= (((v4 + (1e10 + ((-2.25 - v4) * (v4 + -2.25)))) + v4) + (v4 * (v4 * v4)))
    v24 -= (^(v24))
    v25 := GenS4{F0: 42, F1: 65535}
    for fi26 := 0; fi26 < 5; fi26++ {
        v3 = v3
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint16 = 255
            var v29 uint = 4294967295
            v2 = append(v2, 0)
            v24 += v24
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    v30 := GenS5{F0: (v3 | 0)}
    if false {
        v31 := GenS6{F0: 2}
        var v32 float32 = (((v4 / v4) / (v4 * -2.25)) + (3.14159 + (v4 / ((v4 + (float32(float64(0.1)) + float32((0.1 * (float64((v4 / v4)) + (-2.25 / (((0.0 / -2.25) - 1e10) + 0.0))))))) + 3.14159))))
        _ = v31
        _ = v32
    } else {
        switch ("\n\t" + (("foo bar" + "hello") + "")) {
        case "":
            for fi33 := 0; fi33 < 10; fi33++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (605 >= 0) {
                    v34 := GenS2{F0: uint(v5)}
                    _ = v34
                }
                fmt.Println(v24, v24, v24)
                _ = fi33
            }
        case "foo bar":
            v4 = ((v4 / 1.0) - 3.14159)
        default:
            if (!(((v24 * v24) > 509) || true)) {
                v1 -= v1
                var v35 int = v1
                _ = v35
            }
        }
        var v36 int8 = 1
        v37 := GenS5{F0: 0}
        _ = v36
        _ = v37
    }
    v5 = v5
    v1 = v1
    v38 := GenS3{F0: -1}
    v24 += v24
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
}
