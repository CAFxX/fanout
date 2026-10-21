package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32, p1 bool, p2 int) int16 {
    switch (p2 | p2) {
    case 9223372036854775807:
        switch ("\n\t" + "hello") {
        case "\n\t":
            p2 = p2
        case "":
            var v7 uint8 = 2
            _ = v7
        default:
            if (!(42 == 0)) {
                v8 := []int{(-(p2)), p2, (p2 ^ p2), (p2 & (^(p2)))}
                _ = v8
            } else {
                v9 := GenS1{F0: (true && p1), F1: -485, F2: 585}
                _ = v9
            }
        }
    case 595:
        fmt.Println(p2, p0)
    default:
        if p1 {
            p0 = (p0 >> ((4294967295) & 63))
        } else {
            p2 += p2
            p0 -= p0
        }
    }
    var v10 uint = 255
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}
func genFn1(p0 uint64) bool {
    p0 += 4294967295
    var v11 int16 = 32767
    _ = p0
    _ = v11
    return (2 <= 65535)
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]int)
    var v4 float32 = (-2.25 + (3.14159 * 0.0))
    var v5 uint8 = 127
    v6 := make(map[string]bool)
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int32 = 0
        _ = fi12
        _ = v13
    }
    if true {
        if true {
            switch (("foo bar" + "hello") + "hello") {
            case "\n\t":
                var v14 int16 = 386
                _ = v14
            case "hello":
                v1 += (v1 - 1e10)
            default:
                var v15 uint = 1
                _ = v15
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                _ = len(v3)
                v1 += v1
                _ = fi16
            }
        } else {
            var v17 uint32 = 2147483647
            if (65535 == 2) {
                v18 := []int{32767}
                _ = v18
            } else {
                if (100 != -75) {
                    v19 := []int{-9223372036854775808, -102, 42, 1}
                    v4 -= (v4 / v4)
                    fmt.Println(v4)
                    _ = v19
                } else {
                    v1 -= (v1 + float64(-2.25))
                }
                v20 := GenS1{F0: false, F1: 0, F2: 4294967295}
                _ = v20
            }
            _ = v17
        }
        v1 = (((3.14159 / v1) / (float64(((float32((0.0 + v1)) * v4) - (v4 / float32(1e10)))) + -2.25)) - v1)
    }
    var v21 int = -9223372036854775808
    v1 += v1
    v21 -= 0
    fmt.Println(v4)
    _ = len(v2)
    fmt.Println(v21, v5)
    fmt.Println(v4, v5)
    v22 := GenS1{F0: true, F1: 876, F2: 990}
    v23 := GenS1{F0: false, F1: 32767, F2: 340}
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int32 = 0
            switch fi24 {
            case 127:
                if (!((v4 * (v4 + (float32(3.14159) - v4))) == (((v4 - (v4 / (v4 * v4))) * v4) + (-2.25 / 0.1)))) {
                    if true {
                        v27 := []int{}
                        switch 127 {
                        case 9223372036854775807:
                            v4 = ((0.0 / v4) * float32(v1))
                        case -9223372036854775808:
                            fmt.Println(v1, v5)
                        default:
                            var v28 uint32 = 65535
                            _ = v28
                        }
                        v29 := GenS1{F0: false, F1: 42, F2: 405}
                        _ = v27
                        _ = v29
                    } else {
                        var v30 uint32 = 231
                        _ = v30
                    }
                } else {
                    switch "hello" {
                    case "a":
                        switch fi24 {
                        case -43:
                            var v31 string = "\n\t"
                            _ = v31
                        case -224:
                            v32 := GenS2{F0: 0}
                            _ = v32
                        case -14:
                            fmt.Println(fi25, v5, v5)
                        default:
                            v33 := GenS2{F0: 1}
                            _ = v33
                        }
                    }
                }
            }
            v4 = v4
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v34 := []int{}
    v5 += (^(v5))
    var v35 uint64 = 2
    v36 := GenS2{F0: 132}
    var v37 int32 = -89
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
