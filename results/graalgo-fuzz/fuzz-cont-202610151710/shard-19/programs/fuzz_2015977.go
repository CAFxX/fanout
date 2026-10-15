package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8, p1 uint32, p2 int16) uint16 {
    if false {
        p2 += (p2 % ((p2) | 1))
        for fi4 := 0; fi4 < 20; fi4++ {
            var v5 int16 = (p2 | 0)
            fmt.Println(v5, fi4)
            _ = fi4
            _ = v5
        }
        v6 := []int{1, -2147483648, 4294967295}
        _ = v6
    }
    v7 := []int{}
    p2 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 686
}
func genFn1(p0 int8, p1 uint32) int {
    v8 := GenS1{F0: 18446744073709551615}
    var v9 string = "foo bar"
    v10 := []int{0, -9223372036854775808, 2147483647}
    for fi11 := 0; fi11 < 6; fi11++ {
        if ((p0 - (p0 % ((p0) | 1))) == p0) {
            if (false || (-1 == 42)) {
                fmt.Println(fi11)
                var v12 float32 = (float32((1.5 * -2.25)) + float32(1.5))
                var v13 string = "\n\t"
                _ = v12
                _ = v13
            } else {
                if true {
                    p0 = p0
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        var v14 int64 = -9223372036854775808
                        _ = v14
                    case "a":
                        fmt.Println(p0, v9)
                    default:
                        if true {
                            v15 := GenS1{F0: 2}
                            switch v9 {
                            case "\n\t":
                                v16 := GenS2{F0: (p0 ^ p0), F1: 1}
                                _ = v16
                            case "":
                                fmt.Println(fi11)
                            case "foo bar":
                                v10 = append(v10, (fi11 & (fi11 + fi11)))
                            default:
                                if (1 == 228) {
                                    v17 := GenS2{F0: (p0 / ((p0) | 1)), F1: 989}
                                    _ = v17
                                } else {
                                    fmt.Println(p0)
                                }
                            }
                            _ = v15
                        } else {
                            fmt.Println(p0)
                        }
                    }
                }
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint64 = 2
    var v2 float32 = (((3.14159 - 0.1) * -2.25) - 0.0)
    var v3 bool = false
    var v18 uint = 392
    v19 := []int{1, -2147483648}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v20 uint64 = 42
    v2 -= (v2 / (-2.25 * float32(float64(v2))))
    var v21 uint16 = uint16(v20)
    v22 := GenS3{F0: 771, F1: 65535}
    var v23 float32 = v2
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1, v23, v20)
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if true {
            v19 = append(v19, fi24)
            v25 := GenS4{F0: v20, F1: 591}
            switch genFn1(100, 0) {
            case -856:
                v3 = true
            }
            _ = v25
        } else {
            v23 -= -2.25
            var v26 uint32 = 1
            fmt.Println(v23, v26)
            _ = v26
        }
        _ = fi24
    }
    switch "a" {
    case "hello":
        v3 = (genFn1(59, 1) > 1)
    }
    v3 = v3
    if (false && (42 != 576)) {
        fmt.Println(v18, v3)
        v2 = -2.25
        for fi27 := 0; fi27 < 3; fi27++ {
            if false {
                var v28 uint64 = v20
                v29 := GenS4{F0: (v1 / (((v28 / ((1) | 1))) | 1)), F1: -32768}
                v19 = append(v19, fi27)
                _ = v28
                _ = v29
            } else {
                v3 = (((0.1 < (float64(v2) - -2.25)) && v3) || ((v21 << (((v18 & v18)) & 63)) == 2))
            }
            _ = fi27
        }
    }
    v30 := GenS5{F0: 1, F1: 128}
    var v31 uint = (v18 >> ((v18) & 63))
    fmt.Println(v31)
    var v32 bool = (32767 >= 42)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
