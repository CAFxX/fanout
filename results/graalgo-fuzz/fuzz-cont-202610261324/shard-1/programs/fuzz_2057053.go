package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    v1 := []uint32{65535, 2147483648, 42, 0}
    var v2 int = 42
    var v3 float64 = (1e10 - (((-2.25 / -2.25) - -2.25) * ((((((1e10 + 0.0) + (1e10 + ((0.0 * ((1.5 - 3.14159) * ((-2.25 - 1e10) + 1e10))) / 0.1))) - ((3.14159 / 3.14159) + 3.14159)) * float64(-2.25)) / -2.25) / ((((0.1 + (((1e10 / (1.5 + -2.25)) + 0.0) - 1e10)) + float64(0.0)) / 1.0) - ((0.0 - 0.1) * 0.0)))))
    if true {
        if false {
            v4 := GenS1{F0: 879}
            if true {
                v2 -= -914
                switch "hello" {
                case "\n\t":
                    v5 := GenS1{F0: 0}
                    _ = v5
                default:
                    var v6 float64 = (3.14159 - v3)
                    _ = v6
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v4
        } else {
            fmt.Println(v3)
        }
    }
    v7 := GenS2{F0: "", F1: 1, F2: 4294967295}
    v2 -= (v2 << ((255) & 63))
    fmt.Println(v2, v3)
    var v8 int32 = 42
    for fi9 := 0; fi9 < 3; fi9++ {
        v2 += 139
        _ = fi9
    }
    var v10 uint = 0
    switch ((("\n\t" + "hello") + "\n\t") + "foo bar") {
    case "hello":
        var v11 uint8 = 255
        _ = v11
    case "":
        if (true || (!true)) {
            var v12 uint16 = 255
            _ = v12
        }
    case "a":
        var v13 float32 = 1e10
        _ = v13
    default:
        if (!true) {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
    }
    if true {
        v2 = (v2 % ((4294967295) | 1))
        var v14 float64 = v3
        v15 := GenS2{F0: "foo bar", F1: 61, F2: 0}
        _ = v14
        _ = v15
    }
    if (float32(float64(float32((float64(0.1) / v3)))) == 1e10) {
        v16 := GenS3{F0: ((0.0 + ((1e10 / 1e10) + -2.25)) == (1.5 / -2.25))}
        var v17 float64 = 1.5
        for fi18 := 0; fi18 < 17; fi18++ {
            v2 += -2147483648
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 int8 = -128
                var v21 uint = v10
                fmt.Println(fi19)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    } else {
        fmt.Println(v3, v2, v2)
    }
    v22 := GenS4{F0: 255}
    v23 := GenS4{F0: (v10 - 255)}
    if (false && (255 <= 32767)) {
        v2 += v2
    } else {
        fmt.Println(v2)
        _ = len(v1)
    }
    v2 = (-(v2))
    v24 := GenS5{F0: 2, F1: 12}
    v10 = 172
    if ((^(v2)) != (v2 + (-(v2)))) {
        var v25 uint16 = 367
        v26 := GenS1{F0: -579}
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
