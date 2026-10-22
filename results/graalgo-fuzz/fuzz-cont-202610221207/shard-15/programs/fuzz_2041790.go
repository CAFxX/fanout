package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]int)
    var v3 int32 = 0
    var v4 float64 = (0.0 - -2.25)
    var v5 int = v1
    var v6 int = (v5 - (-(v5)))
    if true {
        v4 = -2.25
        _ = len(v2)
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                switch "hello" {
                case "a":
                    var v9 uint32 = 2147483648
                    _ = v9
                case "\n\t":
                    var v10 bool = ((v6 | (v1 % ((fi8) | 1))) != -9223372036854775808)
                    _ = v10
                case "hello":
                    var v11 uint8 = 127
                    _ = v11
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if (!true) {
                            v4 -= (((((-2.25 * v4) - 1e10) - (((v4 - float64(((float32(float64(float32(v4))) / ((0.1 * 1e10) / -2.25)) / (1.5 * ((0.0 + 1e10) / ((0.0 / (1.5 / 1.0)) - ((-2.25 - 0.0) / (-2.25 * (3.14159 + -2.25))))))))) * v4) - 0.0)) * v4) - (((v4 - float64(0.0)) + (v4 / v4)) * (-2.25 / float64((3.14159 - 0.1)))))
                        } else {
                            switch "\n\t" {
                            case "a":
                                if (693 <= 506) {
                                    var v12 int = (v5 >> ((18446744073709551615) & 63))
                                    _ = v12
                                } else {
                                    fmt.Println(v5, fi8)
                                    var v13 int32 = v3
                                    fmt.Println(v13, fi7, v13)
                                    var v14 int64 = 1
                                    _ = v13
                                    _ = v14
                                }
                            default:
                                var v15 uint64 = 0
                                _ = v15
                            }
                            v16 := GenS1{F0: -32768, F1: 127, F2: 41}
                            _ = v16
                        }
                    default:
                        if (1000 > int16(v3)) {
                            v6 = v5
                        } else {
                            v17 := GenS2{F0: false}
                            v18 := GenS1{F0: -32768, F1: 0, F2: 1}
                            var v19 float64 = -2.25
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    }
                    v6 = v6
                    v6 += v6
                }
                v20 := []int{v1, (v6 % ((v1) | 1))}
                v1 = v6
                _ = fi8
                _ = v20
            }
            v21 := GenS1{F0: 42, F1: 1, F2: 121}
            fmt.Println(v3)
            _ = fi7
            _ = v21
        }
    } else {
        fmt.Println(v3, v4)
        if (743 == 0) {
            switch "a" {
            case "foo bar":
                fmt.Println(v6, v6)
            case "":
                if (1 <= 128) {
                    v5 = (-(v6))
                    v22 := GenS1{F0: -32768, F1: -99, F2: 100}
                    _ = v22
                }
            default:
                var v23 int8 = 42
                _ = v23
            }
            var v24 int = v5
            _ = v24
        } else {
            v5 += (^(v1))
            switch (v1 - -9223372036854775808) {
            case 127:
                if false {
                    v25 := GenS1{F0: -400, F1: -128, F2: 1}
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v3 += v3
                        v5 += v1
                        _ = fi26
                    }
                    v5 -= v5
                    _ = v25
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        for fi27 := 0; fi27 < 4; fi27++ {
                            v1 = (v1 / ((v6) | 1))
                            if (648 >= (v1 * (v6 << ((42) & 63)))) {
                                v28 := GenS3{F0: 83, F1: -28, F2: 1}
                                v5 = -556
                                var v29 uint64 = 1
                                _ = v28
                                _ = v29
                            }
                            _ = fi27
                        }
                    case "":
                        v30 := []int{706, 0}
                        _ = v30
                    }
                }
            default:
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v4, fi31, v3)
                    v32 := GenS3{F0: -128, F1: 27, F2: -32768}
                    _ = fi31
                    _ = v32
                }
            }
        }
    }
    v33 := []int{-2147483648, (v1 * v1)}
    v4 += v4
    fmt.Println(v3)
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := GenS2{F0: true}
        v4 += 1e10
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 int32 = (v3 >> ((565) & 63))
            var v38 uint32 = 2
            fmt.Println(fi34, v37, v4)
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v1 += -545
        v4 = v4
        _ = fi39
    }
    v40 := GenS4{F0: "\n\t"}
    v3 = (v3 - 0)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
