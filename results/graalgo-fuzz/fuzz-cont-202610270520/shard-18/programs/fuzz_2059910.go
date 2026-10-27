package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() uint64 {
    var v7 float64 = 0.0
    if ((!(!(1 > 178))) || true) {
        v7 += 0.1
        v8 := []int{0, 9223372036854775807, 127, 2147483647}
        for fi9 := 0; fi9 < 12; fi9++ {
            if ((-32768 <= -107) || true) {
                fmt.Println(v7, fi9)
                fmt.Println(v7, v7)
            } else {
                fmt.Println(fi9, fi9)
                var v10 int16 = 42
                _ = v10
            }
            _ = fi9
        }
        _ = v8
    }
    fmt.Println(v7)
    _ = v7
    return 810
}
func genFn1(p0 uint, p1 uint64) int8 {
    p0 = (p0 + (p0 * p0))
    for fi11 := 0; fi11 < 19; fi11++ {
        p0 = p0
        _ = fi11
    }
    switch 443 {
    case 788:
        p0 += p0
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]string)
    v3 := []uint32{65535, 2147483647, 625, 1}
    v4 := make(map[int]bool)
    var v5 int32 = 3
    v6 := []int{744, 9223372036854775807}
    v12 := GenS1{F0: true, F1: v1}
    fmt.Println(v1)
    v5 += 0
    var v13 int = 127
    _ = len(v6)
    v5 -= 42
    switch (("" + "\n\t") + ("\n\t" + (("hello" + ("foo bar" + "hello")) + "\n\t"))) {
    case "foo bar":
        delete(v4, v13)
    case "a":
        if (!true) {
            if false {
                if (!true) {
                    v1 -= v1
                    v13 -= (v13 >> ((18446744073709551615) & 63))
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    switch "" {
                    case "hello":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v4[(v13 ^ v13)] = true
                            var v16 int16 = -43
                            v17 := v4[(fi15 % ((fi15) | 1))]
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        _ = len(v6)
                    default:
                        v18 := GenS1{F0: false, F1: 255}
                        _ = v18
                    }
                    v5 = (v5 / (((-(v5))) | 1))
                    var v19 int16 = -468
                    _ = fi14
                    _ = v19
                }
            }
            if (1 == -55) {
                var v20 uint16 = 32767
                _ = v20
            }
        } else {
            v13 -= 127
            var v21 uint32 = 1
            v22 := GenS2{F0: (v21 / (((v21 >> ((0) & 63))) | 1)), F1: 255}
            _ = v21
            _ = v22
        }
    case "":
        switch int(v1) {
        case 9223372036854775807:
            fmt.Println(v1)
        case 4294967295:
            v1 -= v1
        }
    default:
        _ = len(v2)
    }
    _ = len(v6)
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint = 83
        _ = fi23
        _ = v24
    }
    fmt.Println(v5, v13)
    var v25 uint8 = v1
    var v26 string = ("foo bar" + ((((("\n\t" + "a") + "") + "") + "foo bar") + ("" + "")))
    var v27 uint8 = 42
    fmt.Println(v27)
    if false {
        v13 += v13
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
