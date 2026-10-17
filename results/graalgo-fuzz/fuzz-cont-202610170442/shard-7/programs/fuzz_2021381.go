package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint, p1 int32, p2 int16) int64 {
    v4 := []int{32767, 32767, -1, 386}
    var v5 int8 = 1
    var v6 float64 = (0.0 / 1e10)
    switch 32768 {
    case 127:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        fmt.Println(v6)
    }
    if (p2 != p2) {
        v7 := GenS1{F0: p0, F1: v5}
        p1 -= (^(p1))
        if (2147483647 < p1) {
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
        }
        _ = v7
    } else {
        switch "\n\t" {
        case "\n\t":
            p0 = 0
        }
        var v8 uint8 = 69
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 int64 = int64(v1)
    var v3 uint16 = 255
    fmt.Println(v2)
    var v9 float64 = (0.0 + (3.14159 / (3.14159 / 1.5)))
    fmt.Println(v2, v2)
    v10 := GenS1{F0: 42, F1: v1}
    if (32767 == 576) {
        fmt.Println(v9, v2, v1)
        var v11 int = 127
        _ = v11
    } else {
        var v12 int32 = -214
        v12 += v12
        _ = v12
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 6; fi14++ {
        switch fi14 {
        case 42:
            switch "hello" {
            case "\n\t":
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := []int{}
                    var v17 uint64 = 42
                    var v18 uint = 4294967295
                    var v19 int8 = -45
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            case "hello":
                v20 := GenS2{F0: -359, F1: 0, F2: 42}
                _ = v20
            }
        case -212:
            v21 := []int{(fi14 / ((127) | 1))}
            _ = v21
        }
        v22 := GenS2{F0: fi14, F1: 18446744073709551615, F2: int32(v1)}
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(v3, fi23, v1)
            v24 := []int{}
            fmt.Println(fi23, v3)
            switch v13 {
            case "\n\t":
                v25 := GenS1{F0: 625, F1: (v1 | v1)}
                _ = v25
            }
            v9 += -2.25
            _ = fi23
            _ = v24
        }
        _ = fi14
        _ = v22
    }
    v13 = "foo bar"
    var v26 int16 = -485
    if false {
        switch 127 {
        case -301:
            if (false && (-387 == 127)) {
                v27 := []int{32767, -270}
                v28 := GenS3{F0: 1}
                _ = v27
                _ = v28
            } else {
                switch -9223372036854775808 {
                case 128:
                    if (0 > 164) {
                        switch v13 {
                        case "\n\t":
                            if (-1 == -238) {
                                v29 := []int{127, 877, -92}
                                v30 := GenS1{F0: 612, F1: v1}
                                _ = v29
                                _ = v30
                            }
                        default:
                            var v31 uint = 470
                            _ = v31
                        }
                        v2 = 1
                    }
                case 940:
                    v32 := GenS1{F0: 706, F1: (-(v1))}
                    _ = v32
                case -1:
                    v33 := GenS4{F0: v26, F1: 2147483648}
                    _ = v33
                }
                v34 := []int{}
                _ = v34
            }
        case 2147483647:
            var v35 uint = 0
            _ = v35
        case 438:
            for fi36 := 0; fi36 < 6; fi36++ {
                for fi37 := 0; fi37 < 11; fi37++ {
                    v3 = (^(v3))
                    _ = fi37
                }
                v26 = (v26 << ((65535) & 63))
                _ = fi36
            }
        default:
            for fi38 := 0; fi38 < 4; fi38++ {
                v13 = (("" + "\n\t") + "")
                v39 := GenS5{F0: (true && false)}
                _ = fi38
                _ = v39
            }
        }
    }
    var v40 float32 = (1e10 / 1.0)
    v41 := []int{123, -2147483648, -9223372036854775808}
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(fi42, v3)
        if (!true) {
            var v43 float32 = (0.0 / float32((1.5 * v9)))
            _ = v43
        }
        v3 = 0
        _ = fi42
    }
    switch 32768 {
    case 128:
        v44 := GenS1{F0: 42, F1: v1}
        _ = v44
    default:
        var v45 uint64 = 1
        _ = v45
    }
    v1 = 22
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v40
    _ = v41
    _ = v9
}
