package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint, p2 uint16) bool {
    p2 += (p2 * (p2 | p2))
    if (-1 != -123) {
        v6 := []int{127}
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint = (p1 * (p1 ^ (p1 | (p1 + 975))))
            if (true && true) {
                switch "foo bar" {
                case "\n\t":
                    switch (fi7 % (((fi7 / (((fi7 ^ (fi7 % ((fi7) | 1)))) | 1))) | 1)) {
                    case 9223372036854775807:
                        switch fi7 {
                        case 2147483647:
                            if (!(!(false && false))) {
                                fmt.Println(p2, p0)
                            } else {
                                var v9 uint16 = (^(p2))
                                v8 += 1
                                if len(v6) > 0 {
                                    v6[0] = fi7
                                }
                                _ = v9
                            }
                        case -980:
                            v10 := GenS1{F0: 0, F1: 2, F2: 709}
                            _ = v10
                        case -9223372036854775808:
                            v8 += (v8 | 214)
                        default:
                            fmt.Println(p0, p0)
                        }
                    case -1:
                        p2 = (p2 >> ((998) & 63))
                    }
                default:
                    if (0.0 != 1e10) {
                        p2 = (p2 * p2)
                    }
                }
                v6 = append(v6, (^(fi7)))
            }
            _ = fi7
            _ = v8
        }
        var v11 bool = (!true)
        _ = v11
        _ = v6
    }
    p0 += uint8(p2)
    _ = p0
    _ = p1
    _ = p2
    return (!true)
}
func genFn1(p0 uint64) uint8 {
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    v13 := GenS2{F0: 255, F1: (("\n\t" + ("a" + "a")) + "a"), F2: 92}
    var v14 float64 = -2.25
    _ = p0
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int = 42
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    var v4 uint64 = 773
    v5 := make(map[string]bool)
    v15 := GenS3{F0: 9}
    v3 -= v3
    v16 := GenS2{F0: 247, F1: "hello", F2: 2}
    fmt.Println(v1, v3, v4)
    fmt.Println(v1, v1)
    switch "\n\t" {
    case "hello":
        v17 := []int{(-(v1)), v1, v1, (v1 | (v1 << ((42) & 63)))}
        _ = v17
    default:
        switch (^(v1)) {
        case 42:
            v1 -= (v1 % ((v1) | 1))
        default:
            var v18 int32 = 0
            _ = v18
        }
    }
    _ = len(v5)
    if false {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch (fi19 / ((500) | 1)) {
            case -830:
                if false {
                    v20 := GenS2{F0: 127, F1: "foo bar", F2: 0}
                    _ = v20
                } else {
                    v21 := GenS4{F0: "a"}
                    v1 = v1
                    _ = v21
                }
            case 42:
                v22 := GenS5{F0: 396}
                _ = v22
            case 9223372036854775807:
                var v23 int8 = -1
                _ = v23
            }
            _ = fi19
        }
    } else {
        v24 := GenS2{F0: 0, F1: "a", F2: genFn1(v4)}
        _ = v24
    }
    var v25 uint64 = v4
    v1 += (v1 ^ 602)
    v26 := []int{v1}
    if len(v26) > 0 {
        v26[0] = (v1 % ((0) | 1))
    }
    var v27 int = (v1 - (v1 >> ((42) & 63)))
    var v28 int32 = 2147483647
    v29 := GenS6{F0: (v3 >> ((846) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
