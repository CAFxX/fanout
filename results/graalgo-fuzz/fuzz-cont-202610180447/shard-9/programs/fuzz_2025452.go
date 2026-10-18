package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
func genFn0(p0 uint) uint16 {
    var v5 bool = false
    switch "foo bar" {
    case "foo bar":
        v6 := []int{128, 1, 32768, 9223372036854775807}
        _ = v6
    case "":
        switch "hello" {
        case "foo bar":
            fmt.Println(v5, p0)
        case "":
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 1; fi8++ {
                    if v5 {
                        switch 952 {
                        case 32767:
                            fmt.Println(fi7)
                        }
                        var v9 int8 = -128
                        _ = v9
                    }
                    if true {
                        v10 := []int{(fi8 * -1)}
                        var v11 int16 = 1
                        p0 = (p0 >> ((p0) & 63))
                        var v12 int16 = v11
                        v11 = v11
                        fmt.Println(v12, v12, fi8)
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        p0 = p0
                        p0 -= 4294967295
                        fmt.Println(fi8)
                    }
                    _ = fi8
                }
                _ = fi7
            }
        case "\n\t":
            if (42 <= 42) {
                switch -131 {
                case 128:
                    switch "foo bar" {
                    case "\n\t":
                        var v13 string = ("" + "foo bar")
                        _ = v13
                    }
                default:
                    var v14 uint8 = 156
                    _ = v14
                }
                for fi15 := 0; fi15 < 5; fi15++ {
                    if (!(1 <= 32767)) {
                        fmt.Println(fi15)
                        v16 := GenS1{F0: 42, F1: 12, F2: -304}
                        _ = v16
                    }
                    if v5 {
                        var v17 int = 483
                        var v18 int64 = 0
                        v17 -= (fi15 >> (((p0 | p0)) & 63))
                        _ = v17
                        _ = v18
                    }
                    _ = fi15
                }
            }
        default:
            var v19 uint16 = 1
            _ = v19
        }
    case "hello":
        v20 := []int{128, 127, 4294967295, -658}
        _ = v20
    default:
        v21 := []int{}
        _ = v21
    }
    switch -9223372036854775808 {
    case -636:
        var v22 int64 = 9223372036854775807
        _ = v22
    case 417:
        fmt.Println(p0, v5)
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            v5 = (!v5)
            _ = fi24
        }
        p0 -= 42
        _ = fi23
    }
    _ = p0
    _ = v5
    return 32767
}

func main() {
    var v1 uint64 = 701
    v2 := make(map[uint32]uint32)
    var v3 int64 = 42
    v4 := make(map[int]uint32)
    var v25 bool = (false || (0 > 42))
    _ = len(v4)
    fmt.Println(v25)
    v26 := []int{-9223372036854775808, 128}
    for fi27 := 0; fi27 < 7; fi27++ {
        if (!(-1 < -567)) {
            _ = len(v2)
        } else {
            fmt.Println(fi27)
            v28 := GenS1{F0: 2147483648, F1: 0, F2: 0}
            v29 := GenS1{F0: 2147483648, F1: 32, F2: -525}
            _ = v28
            _ = v29
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            v2[42] = uint32(v1)
            _ = fi31
        }
        v32 := v4[(fi30 | fi30)]
        v25 = ((v1 << ((287) & 63)) > (v1 & (v1 ^ v1)))
        _ = fi30
        _ = v32
    }
    fmt.Println(v1, v3)
    switch 18 {
    case -9223372036854775808:
        switch ("hello" + "a") {
        case "foo bar":
            var v33 int8 = 100
            _ = v33
        case "":
            switch "hello" {
            case "foo bar":
                v34 := GenS1{F0: 280, F1: 2, F2: 2147483647}
                _ = v34
            case "":
                v1 = v1
            default:
                v35 := GenS1{F0: 42, F1: 697, F2: -1}
                _ = v35
            }
        case "a":
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        default:
            v3 = 316
        }
    case 32768:
        for fi36 := 0; fi36 < 16; fi36++ {
            if (!true) {
                v37 := GenS1{F0: 42, F1: 1, F2: -851}
                fmt.Println(v3)
                delete(v2, 1)
                _ = v37
            }
            _ = fi36
        }
    case -295:
        v25 = (966 >= 2)
    default:
        var v38 float64 = 0.0
        _ = v38
    }
    var v39 int32 = 42
    var v40 float64 = -2.25
    var v41 int64 = (v3 % (((v3 * -1)) | 1))
    v3 -= (-(v41))
    delete(v2, 1)
    if v25 {
        delete(v2, 65535)
    }
    _ = len(v2)
    v41 -= (v3 / ((v3) | 1))
    for fi42 := 0; fi42 < 15; fi42++ {
        v43 := v2[1]
        _ = fi42
        _ = v43
    }
    v44 := GenS1{F0: 2147483647, F1: 762, F2: v39}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
}
