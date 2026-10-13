package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 string) bool {
    switch (p1 + "a") {
    case "a":
        for fi7 := 0; fi7 < 15; fi7++ {
            if (((18446744073709551615 == 0) || (42 == 4294967295)) && (p1 >= "")) {
                p0 -= (fi7 + -1)
            }
            switch 32767 {
            case 9223372036854775807:
                v8 := []int{p0, (p0 / (((fi7 / ((p0) | 1))) | 1)), 1}
                _ = v8
            case -9223372036854775808:
                fmt.Println(p0, p1, p1)
            case -1:
                if (2 > 2) {
                    var v9 float32 = -2.25
                    _ = v9
                }
            }
            _ = fi7
        }
    case "hello":
        p0 += 127
    case "foo bar":
        p1 = (("hello" + "") + p1)
    }
    fmt.Println(p1, p1)
    p0 = (p0 | 0)
    p1 = (p1 + p1)
    v10 := []int{p0, (p0 % ((p0) | 1)), 32767}
    _ = p0
    _ = p1
    _ = v10
    return (true && false)
}
func genFn1(p0 uint16, p1 uint64, p2 uint8) int64 {
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint8 = uint8(p0)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 127}
    switch "\n\t" {
    case "a":
        v14 := GenS2{F0: 32767, F1: -2147483648}
        _ = v14
    case "foo bar":
        v15 := []int{982}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 9223372036854775807
}

func main() {
    v1 := []string{"hello", "hello", ("" + "")}
    v2 := make(map[uint32]int)
    var v3 float32 = (0.1 * 0.1)
    v4 := make(map[uint32]int)
    var v5 float32 = 1.5
    v6 := make(map[string]uint32)
    if false {
        v16 := GenS3{F0: -32768, F1: -2147483648}
        _ = v16
    } else {
        if (36 >= 0) {
            switch 42 {
            case 9223372036854775807:
                fmt.Println(v3, v3)
            case -9223372036854775808:
                v5 -= (0.1 + 3.14159)
            default:
                v5 += (v5 / 1.0)
            }
        }
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if (-2147483648 >= -2147483648) {
            var v18 int8 = int8(fi17)
            if true {
                var v19 bool = false
                if false {
                    if true {
                        fmt.Println(fi17, v19)
                        v20 := GenS2{F0: -32768, F1: -1}
                        _ = v20
                    }
                } else {
                    var v21 int64 = 9223372036854775807
                    _ = v21
                }
                v3 = (((v5 * float32(3.14159)) * 1.5) + v5)
                _ = v19
            }
            delete(v4, 2147483648)
            _ = v18
        } else {
            fmt.Println(fi17)
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS4{F0: 2147483647}
            _ = fi22
            _ = v23
        }
        _ = fi17
    }
    switch (("a" + "hello") + "a") {
    case "\n\t":
        var v24 uint64 = 2
        _ = v24
    default:
        if ("foo bar" != "foo bar") {
            var v25 bool = false
            var v26 bool = false
            _ = v25
            _ = v26
        }
    }
    v3 -= (-2.25 - float32((3.14159 * float64(v3))))
    delete(v2, 2147483648)
    fmt.Println(v3)
    if ((((0.0 / v3) * v5) * -2.25) < (v3 / (((((1.5 + 3.14159) + ((v3 / 1.0) * (v5 - (float32((3.14159 - float64(float32((1.5 / 1.0))))) - v5)))) + 1e10) - (v5 - v5)) - -2.25))) {
        v5 = 1e10
    }
    v27 := v2[0]
    if (!false) {
        v3 -= float32(-2.25)
        var v28 int64 = 845
        _ = v28
    }
    v27 = 1
    var v29 uint32 = 956
    v27 += (-(v27))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
