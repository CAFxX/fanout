package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
func genFn0(p0 uint64, p1 uint32, p2 int8) int64 {
    var v5 bool = false
    p1 += (p1 + p1)
    v5 = (!v5)
    p0 += (p0 | (p0 * 814))
    if v5 {
        var v6 int32 = 2147483647
        if (0 < 818) {
            v7 := GenS1{F0: -9223372036854775808, F1: v6}
            if (p1 < p1) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    var v9 int = fi8
                    for fi10 := 0; fi10 < 5; fi10++ {
                        v11 := GenS2{F0: p2, F1: -559, F2: -9223372036854775808}
                        _ = fi10
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
                var v12 float64 = 1.5
                _ = v12
            }
            v6 -= 1
            _ = v7
        }
        if v5 {
            v13 := []int{128, 9223372036854775807}
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v5)
                v13 = append(v13, -574)
                _ = fi14
            }
            _ = v13
        } else {
            switch ("foo bar" + "") {
            case "foo bar":
                var v15 float64 = 1e10
                _ = v15
            default:
                v16 := GenS2{F0: (p2 >> ((817) & 63)), F1: 131, F2: 9223372036854775807}
                _ = v16
            }
            v17 := []int{}
            if (v5 || ((3.14159 <= -2.25) || v5)) {
                var v18 uint16 = 1
                if v5 {
                    if (("" != "a") && ((-2.25 / 0.1) <= 3.14159)) {
                        fmt.Println(p2, v6, v6)
                        fmt.Println(v6, v6)
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                }
                var v19 float32 = 0.1
                _ = v18
                _ = v19
            } else {
                if (!v5) {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
            }
            _ = v17
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 uint64, p1 uint32) int64 {
    var v20 bool = true
    fmt.Println(p0, v20, p0)
    fmt.Println(v20, p1, p0)
    if ((1e10 + (0.0 + (1e10 - (1.5 * 0.0)))) == 1.5) {
        switch "hello" {
        case "\n\t":
            var v21 string = (("hello" + ("" + "hello")) + ("a" + "hello"))
            _ = v21
        default:
            for fi22 := 0; fi22 < 10; fi22++ {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 int32 = 42
                    var v25 float32 = 0.1
                    fmt.Println(fi22, v20)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        }
        v26 := GenS1{F0: 185, F1: 0}
        _ = v26
    } else {
        v27 := []int{9223372036854775807, 1, 1, 781}
        v20 = (!true)
        _ = v27
    }
    v28 := []int{}
    _ = p0
    _ = p1
    _ = v20
    _ = v28
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483647
    var v3 int = 843
    v4 := make(map[int]uint32)
    var v29 int = v3
    fmt.Println(v29)
    v2 = 417
    delete(v4, v3)
    switch "a" {
    case "\n\t":
        delete(v1, v29)
    case "hello":
        var v30 bool = false
        _ = v30
    case "foo bar":
        _ = len(v1)
    default:
        v31 := []int{v29, v3, v3}
        _ = v31
    }
    if (0.0 <= (-2.25 - ((-2.25 * 0.1) / (((0.0 / 1.5) - (1.5 / 1.0)) + -2.25)))) {
        v29 -= v29
    } else {
        fmt.Println(v29, v3)
        v32 := GenS2{F0: 7, F1: 9223372036854775807, F2: 15}
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := GenS1{F0: (fi33 >> ((65535) & 63)), F1: 423}
        _ = fi33
        _ = v34
    }
    var v35 int = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
