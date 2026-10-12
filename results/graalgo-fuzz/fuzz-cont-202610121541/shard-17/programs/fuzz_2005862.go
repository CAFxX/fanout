package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 bool; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int16; F1 uint8; F2 uint }
type GenS7 struct { F0 int }
func genFn0(p0 int16, p1 bool, p2 uint32) int32 {
    if true {
        v7 := []int{}
        _ = v7
    }
    p1 = (false && (4294967295 < 42))
    _ = p0
    _ = p1
    _ = p2
    return 469
}

func main() {
    var v1 uint32 = 221
    var v2 bool = true
    var v3 uint = 0
    var v4 uint16 = 32767
    var v5 uint8 = 255
    v6 := make(map[int]uint32)
    var v8 uint16 = v4
    v8 = (v8 & v4)
    v4 = 42
    v9 := GenS1{F0: (v3 ^ uint(v4)), F1: (v5 >> ((uint(v8)) & 63))}
    v10 := v6[1]
    switch 0 {
    case 311:
        v5 = (v5 & uint8(v4))
    default:
        var v11 bool = true
        _ = v11
    }
    v5 -= v5
    switch "a" {
    case "foo bar":
        var v12 uint = v3
        _ = v12
    case "\n\t":
        v13 := []int{-2147483648, 9223372036854775807, 42, 701}
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int = (fi14 - fi14)
            switch "foo bar" {
            case "foo bar":
                v3 = v3
            default:
                v16 := GenS1{F0: (^(v3)), F1: v5}
                _ = v16
            }
            var v17 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
            _ = v17
        }
    default:
        if (!(871 <= -1)) {
            if v2 {
                switch -662 {
                case -1:
                    v18 := v6[0]
                    _ = v18
                }
                v19 := GenS2{F0: (v3 >> ((v3) & 63)), F1: v2}
                v6[127] = (v1 | (v1 >> (((^(v3))) & 63)))
                _ = v19
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                switch "foo bar" {
                case "a":
                    fmt.Println(v3, v1)
                case "\n\t":
                    if (-1 > 1) {
                        fmt.Println(v3, v2)
                        v21 := GenS3{F0: (fi20 ^ fi20), F1: (32767 <= -32768), F2: 9223372036854775807}
                        _ = v21
                    } else {
                        fmt.Println(v2, v3, v8)
                        var v22 uint64 = 0
                        _ = v22
                    }
                case "hello":
                    var v23 uint = (v3 ^ v3)
                    _ = v23
                }
                _ = fi20
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                switch (^(fi24)) {
                case 9223372036854775807:
                    if (12 < 78) {
                        if (18446744073709551615 >= 18446744073709551615) {
                            var v25 float64 = ((1e10 / float64((3.14159 + ((3.14159 * -2.25) * (((3.14159 - 0.1) / 1e10) + 1e10))))) / 1e10)
                            var v26 float64 = (v25 / 3.14159)
                            _ = v25
                            _ = v26
                        }
                    }
                case 1000:
                    fmt.Println(v4)
                case -179:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        v28 := GenS4{F0: -32768}
                        var v29 int = (fi27 & (^(fi27)))
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    delete(v6, 0)
                }
                _ = fi24
            }
        }
    }
    var v30 uint = v3
    v31 := GenS5{F0: 1, F1: ""}
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 float64 = -2.25
        v30 -= 4294967295
        _ = fi32
        _ = v33
    }
    var v34 uint64 = uint64(v3)
    v35 := GenS4{F0: 1000}
    switch -494 {
    case 32767:
        v3 -= v3
    case 647:
        v3 -= (v30 >> (((v3 & 217)) & 63))
    case 2147483647:
        if v2 {
            v1 -= (v1 << (((v30 << ((v3) & 63))) & 63))
        }
    default:
        var v36 int32 = genFn0(42, v2, (v1 ^ (v1 + 302)))
        _ = v36
    }
    v37 := GenS6{F0: 205, F1: uint8(v3), F2: (v3 * 912)}
    v38 := GenS7{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
