package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 int8, p2 uint) bool {
    p1 -= (p1 & (p1 / ((p1) | 1)))
    p0 -= 458
    _ = p0
    _ = p1
    _ = p2
    return (!false)
}
func genFn1(p0 int8, p1 int) int32 {
    switch (("\n\t" + ("a" + "a")) + "\n\t") {
    case "a":
        p1 -= (p1 ^ (-(p1)))
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int8 = p0
            var v9 int16 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        p1 += 4294967295
        if true {
            v10 := GenS1{F0: (p0 % (((^(p0))) | 1)), F1: (905 > 1)}
            var v11 bool = true
            var v12 int32 = 42
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    v13 := []int{p1}
    if (!(127 != (p0 >> ((2) & 63)))) {
        v13 = append(v13, p1)
        p1 += p1
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = (1e10 * -2.25)
    var v3 float32 = (((v2 - v2) + float32(-2.25)) - (v2 + v2))
    var v4 int32 = 2147483647
    var v5 int64 = 9223372036854775807
    fmt.Println(v1, v5, v1)
    var v14 uint16 = 236
    if (true || true) {
        v3 += (v2 / 0.1)
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: -109, F1: (v14 > (v14 ^ 65535))}
            _ = fi15
            _ = v16
        }
        v17 := []int{32768, 1, int(v5), int(v4)}
        _ = v17
    }
    v18 := []int{32767, 32767, -6, 42}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int16 = 32767
        switch "a" {
        case "\n\t":
            v18 = append(v18, -518)
        case "hello":
            var v21 int32 = (v4 & v4)
            _ = v21
        case "":
            v3 += v3
        default:
            fmt.Println(v1, v3)
        }
        v20 -= 315
        _ = fi19
        _ = v20
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v4 = (v4 - v4)
        v18 = append(v18, (fi22 / ((-9223372036854775808) | 1)))
        switch "foo bar" {
        case "a":
            var v23 uint8 = 255
            _ = v23
        case "hello":
            var v24 int32 = 42
            _ = v24
        case "\n\t":
            var v25 uint8 = 223
            _ = v25
        default:
            if (4294967295 > 65535) {
                fmt.Println(v2, v5, v5)
                var v26 int = (-(fi22))
                _ = v26
            }
        }
        _ = fi22
    }
    v3 = v2
    var v27 uint8 = 255
    v28 := GenS2{F0: ((("hello" + "hello") + "foo bar") + "foo bar")}
    v29 := GenS1{F0: 42, F1: false}
    var v30 float64 = 1.5
    fmt.Println(v2, v14, v2)
    var v31 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
