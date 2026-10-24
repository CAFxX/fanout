package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32; F2 int64 }
type GenS3 struct { F0 int32; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8; F1 int; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 int16; F2 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 int8; F1 int32; F2 uint8 }
func genFn0(p0 int64, p1 uint, p2 int) uint32 {
    v4 := []int{}
    switch 4294967295 {
    case 0:
        for fi5 := 0; fi5 < 3; fi5++ {
            if (true && false) {
                switch "a" {
                case "a":
                    var v6 uint = (p1 / (((p1 + 4294967295)) | 1))
                    _ = v6
                }
                if false {
                    var v7 uint = (p1 - 255)
                    _ = v7
                } else {
                    v8 := GenS1{F0: 100}
                    for fi9 := 0; fi9 < 1; fi9++ {
                        var v10 float64 = 3.14159
                        v11 := GenS1{F0: -1}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    if false {
                        var v12 int64 = p0
                        _ = v12
                    }
                    _ = v8
                }
                for fi13 := 0; fi13 < 14; fi13++ {
                    _ = len(v4)
                    _ = fi13
                }
            }
            if true {
                p0 += p0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            if len(v4) > 0 {
                v4[0] = fi5
            }
            _ = fi5
        }
    default:
        if true {
            if (42 != 1) {
                for fi14 := 0; fi14 < 9; fi14++ {
                    fmt.Println(p1, fi14)
                    _ = fi14
                }
            }
        } else {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS2{F0: fi15, F1: 531, F2: p0}
                if true {
                    p1 += (p1 << (((p1 - 65535)) & 63))
                    v17 := GenS3{F0: 1, F1: 0, F2: int32(p0)}
                    _ = v17
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS3{F0: 0, F1: (p0 >> (((p1 / ((255) | 1))) & 63)), F2: -1}
                    _ = len(v4)
                    v20 := GenS1{F0: -128}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = fi15
                _ = v16
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int32, p1 bool) bool {
    if p1 {
        v21 := []int{191}
        v22 := GenS2{F0: 384, F1: 2147483648, F2: 42}
        switch -1 {
        case -1:
            v23 := GenS4{F0: 59, F1: 1, F2: (!(4294967295 <= 42))}
            _ = v23
        case -275:
            p0 = (p0 / (((^(p0))) | 1))
        case 563:
            p1 = (255 != 42)
        default:
            var v24 int32 = p0
            _ = v24
        }
        _ = v21
        _ = v22
    } else {
        v25 := []int{446}
        _ = v25
    }
    p1 = p1
    fmt.Println(p0)
    p0 -= (p0 * -2147483648)
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int16 = 240
    var v2 int = -777
    var v3 uint8 = 1
    fmt.Println(v3)
    var v26 int32 = 1
    v26 += (v26 / ((v26) | 1))
    switch "a" {
    case "hello":
        for fi27 := 0; fi27 < 12; fi27++ {
            v3 = 80
            v26 -= (v26 | 2147483647)
            _ = fi27
        }
    default:
        switch 2147483647 {
        case -2147483648:
            var v28 int16 = 42
            _ = v28
        case 42:
            if false {
                v26 -= (v26 | 1)
            }
        default:
            switch v2 {
            case -9223372036854775808:
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        for fi29 := 0; fi29 < 10; fi29++ {
                            if (-2.25 <= 0.1) {
                                switch "hello" {
                                case "foo bar":
                                    if true {
                                        for fi30 := 0; fi30 < 3; fi30++ {
                                            if false {
                                                switch v2 {
                                                case 0:
                                                    switch (fi30 >> ((639) & 63)) {
                                                    case 4294967295:
                                                        fmt.Println(v3)
                                                    case 9223372036854775807:
                                                        v2 = (-(fi30))
                                                    case -247:
                                                        var v31 string = (("a" + "") + ("" + "a"))
                                                        _ = v31
                                                    }
                                                default:
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        v26 += v26
                                                    }
                                                }
                                                var v32 int16 = 1000
                                                var v33 int8 = -1
                                                _ = v32
                                                _ = v33
                                            }
                                            v26 = v26
                                            _ = fi30
                                        }
                                        var v34 uint8 = v3
                                        _ = v34
                                    }
                                case "\n\t":
                                    fmt.Println(v3, fi29)
                                case "a":
                                    var v35 bool = (!false)
                                    _ = v35
                                }
                                v36 := []int{-1, (v2 % (((^(v2))) | 1)), (v2 | v2), v2}
                                _ = v36
                            } else {
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    var v37 uint64 = 42
                                    _ = v37
                                case "\n\t":
                                    for fi38 := 0; fi38 < 15; fi38++ {
                                        fmt.Println(fi29, v26)
                                        _ = fi38
                                    }
                                }
                                v3 -= (^(v3))
                                var v39 uint16 = 255
                                _ = v39
                            }
                            if true {
                                v40 := []int{(v2 - fi29)}
                                _ = v40
                            }
                            _ = fi29
                        }
                    default:
                        for fi41 := 0; fi41 < 10; fi41++ {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                var v43 float32 = (((3.14159 + (float32(float64(0.1)) + (1.5 + 1e10))) * (-2.25 + 1e10)) + 3.14159)
                                switch (fi41 + fi41) {
                                case 1:
                                    switch (fi41 % ((9223372036854775807) | 1)) {
                                    case 32768:
                                        if true {
                                            fmt.Println(v1)
                                        }
                                    }
                                default:
                                    v3 = (v3 << ((834) & 63))
                                }
                                _ = fi42
                                _ = v43
                            }
                            _ = fi41
                        }
                    }
                }
            case -989:
                for fi44 := 0; fi44 < 20; fi44++ {
                    switch (("hello" + ("foo bar" + "\n\t")) + "") {
                    case "foo bar":
                        switch (-(v2)) {
                        case 438:
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v46 := GenS5{F0: ((("" + ("a" + (("hello" + ("hello" + "hello")) + ""))) + "a") + "\n\t")}
                                _ = fi45
                                _ = v46
                            }
                        case -2147483648:
                            if (302 >= 2) {
                                v47 := GenS6{F0: "", F1: 0, F2: 2147483647}
                                fmt.Println(fi44, v2)
                                _ = v47
                            } else {
                                var v48 int8 = -121
                                _ = v48
                            }
                        case 32767:
                            if (!true) {
                                fmt.Println(v2, v26, v1)
                                if true {
                                    if true {
                                        v3 = v3
                                    } else {
                                        v49 := GenS3{F0: v26, F1: 9223372036854775807, F2: (v26 ^ v26)}
                                        v50 := []int{4294967295, (^(fi44))}
                                        _ = v49
                                        _ = v50
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v51 := []int{fi44, (v2 ^ (fi44 - -2147483648)), (fi44 ^ v2), (v2 & fi44)}
                                        _ = v51
                                    case "\n\t":
                                        v3 -= (v3 - 2)
                                    default:
                                        var v52 uint8 = v3
                                        _ = v52
                                    }
                                    for fi53 := 0; fi53 < 16; fi53++ {
                                        if true {
                                            v1 = v1
                                        }
                                        _ = fi53
                                    }
                                }
                                var v54 uint32 = genFn0(42, 2, v2)
                                _ = v54
                            }
                        default:
                            if true {
                                for fi55 := 0; fi55 < 5; fi55++ {
                                    if true {
                                        if false {
                                            fmt.Println(fi55, v1)
                                        } else {
                                            if (!(334 >= 0)) {
                                                fmt.Println(fi55)
                                                v26 -= 0
                                            } else {
                                                v56 := []int{4294967295, -1}
                                                _ = v56
                                            }
                                            v3 = 255
                                        }
                                        if (65535 != 2) {
                                            v1 -= (-(v1))
                                        }
                                        v57 := []int{(fi55 & (fi44 / (((^(fi44))) | 1))), v2, (-(v2)), v2}
                                        var v58 int32 = (v26 - (v26 ^ (v26 & -867)))
                                        v59 := GenS3{F0: -1, F1: 1, F2: -9}
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi55
                                }
                                fmt.Println(v26)
                            }
                        }
                    case "a":
                        v60 := []int{(^(fi44)), (v2 + v2)}
                        _ = v60
                    case "\n\t":
                        var v61 bool = true
                        _ = v61
                    default:
                        for fi62 := 0; fi62 < 19; fi62++ {
                            v26 = (v26 ^ v26)
                            if true {
                                fmt.Println(v2, fi44)
                                var v63 float64 = float64(1e10)
                                fmt.Println(fi44, fi44, v2)
                                _ = v63
                            }
                            v64 := []int{-2147483648, fi62, fi44}
                            _ = fi62
                            _ = v64
                        }
                    }
                    _ = fi44
                }
            }
        }
    }
    v26 -= int32(v1)
    fmt.Println(v2, v26, v3)
    v65 := []int{1, 1}
    var v66 uint16 = 1
    fmt.Println(v2, v3)
    for fi67 := 0; fi67 < 10; fi67++ {
        v68 := GenS7{F0: 546}
        v69 := GenS8{F0: v3, F1: "\n\t"}
        _ = fi67
        _ = v68
        _ = v69
    }
    v70 := GenS9{F0: int8(v3), F1: 329, F2: 2}
    _ = len(v65)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v65
    _ = v66
    _ = v70
}
