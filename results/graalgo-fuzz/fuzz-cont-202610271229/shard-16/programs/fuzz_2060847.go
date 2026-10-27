package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint16, p1 int32, p2 int64) int8 {
    p2 += -910
    v5 := GenS1{F0: ("a" + "hello"), F1: p0}
    var v6 int32 = (p1 | p1)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{fi7}
        switch "\n\t" {
        case "\n\t":
            v6 = -79
        case "a":
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(p0, p1, p0)
                _ = len(v8)
                p1 += 667
                _ = fi9
            }
        case "":
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    var v11 int32 = (p1 | -31)
                    p0 -= p0
                    fmt.Println(p1, p1)
                    _ = v11
                }
                switch "a" {
                case "hello":
                    var v12 uint8 = 255
                    _ = v12
                }
                _ = fi10
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int, p1 uint32) uint16 {
    var v13 uint = 438
    var v14 uint8 = 0
    switch (p0 << (((v13 - v13)) & 63)) {
    case 32768:
        v15 := []int{p0, (p0 / ((p0) | 1))}
        _ = v15
    default:
        switch (p0 ^ p0) {
        case -9223372036854775808:
            var v16 int = 2147483647
            _ = v16
        case 42:
            var v17 uint = v13
            _ = v17
        case -224:
            fmt.Println(v13, v14, p1)
        default:
            fmt.Println(v13, v14)
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 338
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = -850
    var v3 bool = true
    var v4 uint = 0
    v18 := []int{-9223372036854775808}
    var v19 uint8 = 2
    v4 = 273
    switch -662 {
    case 155:
        var v20 string = ""
        _ = v20
    }
    if v3 {
        v19 = 124
    }
    fmt.Println(v1, v19, v4)
    switch 32767 {
    case 0:
        switch 128 {
        case 213:
            v21 := GenS1{F0: "foo bar", F1: genFn1(2147483647, 2147483648)}
            _ = v21
        case 790:
            v18 = append(v18, 478)
        case 128:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        default:
            var v22 uint = (v4 | (v4 & (^(v4))))
            _ = v22
        }
    }
    v3 = (!v3)
    v3 = v3
    fmt.Println(v1, v2, v4)
    if (false || v3) {
        var v23 float32 = float32(3.14159)
        _ = v23
    }
    v24 := GenS2{F0: (^(v19))}
    v25 := GenS3{F0: v1, F1: -128}
    var v26 float64 = (0.1 * ((0.0 - (-2.25 / (0.1 - (0.0 * float64((((((1e10 - -2.25) - 0.0) - 1e10) + (1.5 - 0.0)) + (float32((3.14159 + 0.0)) / -2.25))))))) / (((1.5 - 1.5) * (1.5 * (float64(((-2.25 - (1.5 + 0.1)) * (-2.25 / 3.14159))) + 0.0))) - (3.14159 / 1e10))))
    v19 += v19
    var v27 uint = 614
    v18 = append(v18, 143)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
