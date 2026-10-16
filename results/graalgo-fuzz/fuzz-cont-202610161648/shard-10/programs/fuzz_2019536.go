package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int8, p1 uint32) int64 {
    p0 -= (p0 & 27)
    p0 -= (p0 >> ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    return 842
}
func genFn1(p0 bool, p1 uint8, p2 uint32) uint64 {
    switch "" {
    case "hello":
        v5 := GenS1{F0: 18446744073709551615}
        _ = v5
    case "\n\t":
        if (!(!(!(-32768 == 1000)))) {
            fmt.Println(p0, p2)
        } else {
            p1 = uint8(p2)
        }
    case "a":
        v6 := []int{18, 538}
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    if p0 {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{(^(fi8))}
            p1 = (p1 & 127)
            var v10 int16 = 600
            _ = fi8
            _ = v10
            _ = v9
        }
        p1 = (p1 + (p1 << ((18446744073709551615) & 63)))
    }
    p1 += 42
    v11 := GenS2{F0: p0, F1: -873}
    var v12 float32 = ((3.14159 * 3.14159) - 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []string{"hello", "\n\t", ""}
    v2 := make(map[string]int)
    v3 := []int{-671}
    v4 := make(map[string]uint32)
    if (42 == 42) {
        if (18 == 23) {
            var v13 float64 = 0.1
            v13 = (v13 - (0.0 + ((v13 / ((1e10 - -2.25) * (((-2.25 + (float64(0.0) - v13)) - v13) / (float64((((-2.25 - 1e10) / -2.25) + 0.1)) - v13)))) - 0.1)))
            _ = v13
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint16 = 42
                var v16 bool = (((1e10 * 3.14159) * 3.14159) < (3.14159 - 3.14159))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v1)
        for fi17 := 0; fi17 < 9; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint = 401
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 string = ""
                switch ("\n\t" + "foo bar") {
                case "hello":
                    v22 := v4["\n\t"]
                    _ = v22
                case "foo bar":
                    v21 = "a"
                }
                var v23 float64 = (1.5 - (1e10 - (1.5 + -2.25)))
                _ = fi20
                _ = v21
                _ = v23
            }
            _ = fi17
        }
    }
    var v24 float32 = (1.5 / 1.0)
    v24 -= v24
    var v25 int = -353
    var v26 uint64 = 1
    var v27 int = (v25 ^ 2147483647)
    fmt.Println(v25)
    v24 += (v24 * (((v24 + (0.1 - (0.1 / float32(3.14159)))) / (float32((float64(1.5) * 3.14159)) + v24)) * ((v24 + (float32(1e10) * v24)) / v24)))
    v26 = (v26 >> ((uint(v26)) & 63))
    v24 = v24
    var v28 uint64 = (v26 & v26)
    v4["foo bar"] = 2
    var v29 float32 = ((v24 / -2.25) * v24)
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
