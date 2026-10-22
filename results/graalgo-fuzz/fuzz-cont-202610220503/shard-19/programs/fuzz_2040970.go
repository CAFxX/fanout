package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
func genFn0(p0 int, p1 int16, p2 int8) uint64 {
    v5 := GenS1{F0: 528}
    p2 = (p2 & (^(p2)))
    p1 -= 32767
    var v6 int = 843
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 65535
    var v3 uint32 = 65535
    var v4 int16 = -32768
    v7 := GenS1{F0: uint32(v2)}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = (v1 | 262)
        v10 := GenS2{F0: uint16(v3), F1: 128, F2: 220}
        v1 -= v1
        _ = fi8
        _ = v10
        _ = v9
    }
    if (true || true) {
        v1 += genFn0(128, v4, 0)
    } else {
        if ((!((814 < 42) || true)) || false) {
            v1 += (v1 % ((v1) | 1))
        } else {
            v11 := []int{-272, -2147483648, -801, 128}
            for fi12 := 0; fi12 < 7; fi12++ {
                if true {
                    var v13 uint = 723
                    v4 = (v4 * v4)
                    _ = v13
                }
                var v14 int = fi12
                _ = fi12
                _ = v14
            }
            var v15 string = "hello"
            _ = v11
            _ = v15
        }
        fmt.Println(v2, v1, v3)
        var v16 uint32 = (v3 ^ v3)
        _ = v16
    }
    v1 = genFn0(-456, v4, 102)
    v17 := []int{9223372036854775807, 32768}
    v17 = append(v17, 32768)
    var v18 int64 = 761
    var v19 int64 = v18
    v2 = v2
    for fi20 := 0; fi20 < 6; fi20++ {
        v18 = 9223372036854775807
        _ = fi20
    }
    fmt.Println(v3)
    _ = len(v17)
    if true {
        for fi21 := 0; fi21 < 13; fi21++ {
            if false {
                if len(v17) > 0 {
                    v17[0] = 2147483647
                }
                v19 += (v18 >> ((946) & 63))
                v22 := GenS3{F0: uint(v3)}
                _ = v22
            } else {
                fmt.Println(v3, v18, v1)
            }
            _ = fi21
        }
        var v23 int16 = (v4 | 842)
        _ = v23
    }
    _ = len(v17)
    switch 2147483647 {
    case 128:
        _ = len(v17)
    case -1:
        var v24 uint32 = (v3 << ((0) & 63))
        _ = v24
    case 1:
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                v4 += -918
                for fi26 := 0; fi26 < 14; fi26++ {
                    _ = len(v17)
                    var v27 uint32 = v3
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    if (uint(v1) != 547) {
                        v1 = genFn0(fi26, (v4 % ((v4) | 1)), 100)
                        v3 = 2147483648
                        v28 := GenS1{F0: (v3 + (v3 + 0))}
                        _ = v28
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            v4 += (v4 & (^(v4)))
                        case "foo bar":
                            v19 -= (v18 >> ((2) & 63))
                        default:
                            if (65535 == 42) {
                                if false {
                                    if false {
                                        if (fi26 > (fi25 & -244)) {
                                            fmt.Println(v27)
                                        }
                                        fmt.Println(v18, v2, v2)
                                    }
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    }
                                    fmt.Println(v19, v18, v19)
                                    fmt.Println(v19, v4, v27)
                                    v29 := GenS4{F0: (v1 - uint64(v2)), F1: 85, F2: 9223372036854775807}
                                    v30 := GenS5{F0: (^(fi25))}
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v3 = v3
                            }
                        }
                    }
                    v31 := GenS3{F0: 0}
                    _ = fi26
                    _ = v27
                    _ = v31
                }
                _ = fi25
            }
            var v32 int64 = (v18 / (((v18 % (((-(v19))) | 1))) | 1))
            if true {
                v3 += (^(v3))
                fmt.Println(v18, v18, v4)
                v18 -= 1
            } else {
                var v33 float64 = 0.0
                _ = v33
            }
            _ = v32
        }
    default:
        switch "" {
        case "a":
            v34 := GenS4{F0: (^(v1)), F1: 128, F2: v18}
            _ = v34
        case "\n\t":
            v35 := GenS6{F0: -67}
            _ = v35
        }
    }
    var v36 int64 = -9223372036854775808
    switch -9223372036854775808 {
    case 2147483647:
        v17 = append(v17, 32768)
    case 127:
        if (true && true) {
            var v37 float64 = 1e10
            _ = v37
        } else {
            v18 = v36
            fmt.Println(v2)
        }
    }
    v2 = (v2 ^ (v2 / (((v2 << ((4294967295) & 63))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
