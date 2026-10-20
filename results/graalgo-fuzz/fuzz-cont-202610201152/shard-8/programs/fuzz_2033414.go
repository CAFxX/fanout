package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint32 }
func genFn0(p0 string) int {
    v4 := GenS1{F0: 2}
    if (!true) {
        v5 := GenS1{F0: 168}
        switch 127 {
        case 42:
            switch "a" {
            case "":
                v6 := []int{-1, 2147483647, -286}
                _ = v6
            case "hello":
                fmt.Println(p0)
            case "foo bar":
                v7 := []int{}
                _ = v7
            }
        case -466:
            var v8 bool = false
            _ = v8
        }
        _ = v5
    }
    v9 := GenS1{F0: 65535}
    p0 = ""
    _ = p0
    _ = v4
    _ = v9
    return 2147483647
}
func genFn1(p0 bool) int8 {
    v10 := []int{32768, -128, 127}
    p0 = (!(!(123 >= 32767)))
    p0 = (!(42 <= 585))
    fmt.Println(p0)
    v11 := GenS2{F0: 65535, F1: 2}
    _ = p0
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{}
    var v3 int64 = 0
    if (1 == 926) {
        v3 -= (v3 * v3)
        v1 -= ((v1 - 0.0) / v1)
    } else {
        var v12 string = "\n\t"
        _ = v12
    }
    var v13 int32 = -2147483648
    var v14 float32 = 1.5
    v1 -= (v14 - (1e10 * 3.14159))
    if true {
        fmt.Println(v3, v14, v3)
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            fmt.Println(v14, v13, fi15)
            _ = fi15
        }
    }
    var v16 uint32 = 0
    for fi17 := 0; fi17 < 11; fi17++ {
        switch "\n\t" {
        case "\n\t":
            switch (fi17 % ((fi17) | 1)) {
            case 9223372036854775807:
                v18 := GenS2{F0: 255, F1: 2}
                _ = v18
            case 2147483647:
                switch "foo bar" {
                case "hello":
                    v19 := GenS2{F0: uint16(v16), F1: 717}
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 uint32 = v16
                        v14 = 0.0
                        v2 = append(v2, genFn0("\n\t"))
                        _ = fi20
                        _ = v21
                    }
                case "foo bar":
                    v14 = v14
                }
            case 128:
                var v22 string = "a"
                _ = v22
            default:
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 int16 = -32768
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi23
                    _ = v24
                }
            }
        case "foo bar":
            switch ("hello" + "") {
            case "hello":
                v2 = append(v2, (fi17 ^ genFn0(("\n\t" + ("" + "")))))
            case "\n\t":
                v3 = v3
            default:
                v14 += (v14 + 1.5)
            }
        }
        _ = fi17
    }
    fmt.Println(v13, v13, v13)
    switch "\n\t" {
    case "foo bar":
        v25 := GenS2{F0: 657, F1: 4294967295}
        _ = v25
    case "":
        v26 := GenS2{F0: 551, F1: 42}
        _ = v26
    }
    v27 := GenS3{F0: -9223372036854775808, F1: ("hello" + ""), F2: 32767}
    for fi28 := 0; fi28 < 12; fi28++ {
        v1 += (3.14159 + -2.25)
        for fi29 := 0; fi29 < 17; fi29++ {
            switch "hello" {
            case "hello":
                v30 := GenS3{F0: -1, F1: "\n\t", F2: 0}
                _ = v30
            default:
                fmt.Println(fi29, v14)
            }
            switch (fi28 / (((fi29 - (fi29 >> ((2) & 63)))) | 1)) {
            case -1:
                var v31 uint32 = 116
                _ = v31
            case 0:
                fmt.Println(v16)
            default:
                v13 = v13
            }
            _ = fi29
        }
        v14 = v1
        _ = fi28
    }
    v32 := GenS1{F0: 42}
    if true {
        for fi33 := 0; fi33 < 13; fi33++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v34 int32 = 42
            var v35 int64 = (v3 / ((-705) | 1))
            _ = fi33
            _ = v34
            _ = v35
        }
        var v36 int64 = (v3 - v3)
        _ = v36
    }
    if (((^(v13)) <= (v13 ^ 989)) || true) {
        v37 := GenS2{F0: 277, F1: 18446744073709551615}
        for fi38 := 0; fi38 < 15; fi38++ {
            v39 := GenS1{F0: 0}
            if ((("hello" + "hello") + (("\n\t" + "a") + "\n\t")) <= "a") {
                var v40 uint8 = 42
                if (v16 >= (v16 - (v16 << ((42) & 63)))) {
                    var v41 float64 = 3.14159
                    _ = v41
                } else {
                    switch "a" {
                    case "hello":
                        for fi42 := 0; fi42 < 18; fi42++ {
                            var v43 float32 = (1e10 * v1)
                            var v44 string = ("\n\t" + "\n\t")
                            v45 := GenS4{F0: "foo bar", F1: 32767, F2: 135}
                            v40 = 255
                            if (827 == 255) {
                                if (true && false) {
                                    _ = len(v2)
                                }
                            }
                            var v46 int64 = 42
                            if ((false || true) || true) {
                                fmt.Println(v16)
                                v44 = (v44 + v44)
                                switch 127 {
                                case -648:
                                    v1 += float32(1e10)
                                case 4294967295:
                                    v47 := GenS5{F0: true, F1: 42}
                                    fmt.Println(v14)
                                    _ = v47
                                case 0:
                                    if true {
                                        v48 := GenS2{F0: 32767, F1: 2}
                                        v43 += 0.1
                                        switch (fi38 + genFn0("hello")) {
                                        case 10:
                                            v40 = (v40 & (v40 + (v40 ^ 0)))
                                        case 32768:
                                            fmt.Println(v43, v13)
                                        case 0:
                                            v2 = append(v2, (fi38 - (^(fi38))))
                                        }
                                        _ = v48
                                    } else {
                                        if (1 > 32767) {
                                            switch fi38 {
                                            case -9223372036854775808:
                                                var v49 uint16 = 65535
                                                _ = v49
                                            case 0:
                                                switch fi38 {
                                                case 4294967295:
                                                    v1 = v14
                                                case 127:
                                                    if (v13 == -2147483648) {
                                                        switch fi38 {
                                                        case 0:
                                                            if (-2.25 != float64(float32((0.1 / 0.1)))) {
                                                                v1 = v14
                                                                switch v44 {
                                                                case "foo bar":
                                                                    fmt.Println(v3, v40, fi42)
                                                                case "a":
                                                                    if false {
                                                                        if (((v46 <= v3) || ((!(v14 < v43)) && false)) && (v44 == "hello")) {
                                                                            switch fi38 {
                                                                            case 1:
                                                                                v50 := GenS6{F0: 32767}
                                                                                _ = v50
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch fi38 {
                                                                        case 298:
                                                                            fmt.Println(v13, v43, fi38)
                                                                        default:
                                                                            if ((!(493 <= 18446744073709551615)) || (false || ((v16 - (^(v16))) >= 2147483647))) {
                                                                                if ((!(255 == 18446744073709551615)) && false) {
                                                                                    v51 := GenS7{F0: v46, F1: v16, F2: v40}
                                                                                    _ = v51
                                                                                } else {
                                                                                    var v52 uint16 = 1
                                                                                    _ = v52
                                                                                }
                                                                                v13 = (v13 | 42)
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v53 float64 = 1e10
                                                                    _ = v53
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            if (42 >= 0) {
                                                                var v54 int16 = -32768
                                                                fmt.Println(v54, v43)
                                                                _ = v54
                                                            }
                                                        case 32768:
                                                            _ = len(v2)
                                                        default:
                                                            fmt.Println(fi42, v14, v1)
                                                        }
                                                        if false {
                                                            var v55 int16 = -32768
                                                            v3 += (v46 >> ((65535) & 63))
                                                            switch (((v44 + ((v44 + (("" + (("\n\t" + "foo bar") + v44)) + ("a" + ("" + ("\n\t" + ""))))) + "\n\t")) + v44) + "\n\t") {
                                                            case "foo bar":
                                                                fmt.Println(v1, v44, v43)
                                                            }
                                                            var v56 uint16 = 65535
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            if false {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v43, v44, v44)
                                                    }
                                                }
                                            }
                                            var v57 float64 = 3.14159
                                            _ = v57
                                        } else {
                                            switch fi38 {
                                            case 0:
                                                fmt.Println(fi38, v43, v16)
                                            case -2147483648:
                                                fmt.Println(v46)
                                            }
                                        }
                                        v58 := GenS5{F0: (!true), F1: 39}
                                        v1 -= (1.5 + (v14 * v14))
                                        _ = v58
                                    }
                                default:
                                    var v59 float32 = v1
                                    _ = v59
                                }
                            }
                            _ = len(v2)
                            if true {
                                var v60 float32 = (3.14159 * 3.14159)
                                v16 = 42
                                var v61 uint32 = v16
                                _ = v60
                                _ = v61
                            } else {
                                fmt.Println(v1, v46)
                                v13 = (-(v13))
                                var v62 uint32 = v16
                                _ = v62
                            }
                            _ = fi42
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                    }
                    v2 = append(v2, fi38)
                    v13 = v13
                }
                for fi63 := 0; fi63 < 5; fi63++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi38, v14)
                    _ = fi63
                }
                _ = v40
            } else {
                v14 = v1
                for fi64 := 0; fi64 < 4; fi64++ {
                    v1 -= (v14 - v1)
                    var v65 int = genFn0("\n\t")
                    _ = len(v2)
                    _ = fi64
                    _ = v65
                }
                for fi66 := 0; fi66 < 15; fi66++ {
                    var v67 int32 = (v13 + v13)
                    var v68 int16 = 1000
                    v69 := GenS5{F0: false, F1: 0}
                    _ = fi66
                    _ = v67
                    _ = v68
                    _ = v69
                }
            }
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    if true {
        fmt.Println(v1, v13, v13)
        var v70 bool = (3.14159 > 0.0)
        var v71 int8 = genFn1(((v3 != (-(v3))) || (!(0 < 97))))
        _ = v70
        _ = v71
    } else {
        v72 := GenS5{F0: false, F1: -1}
        switch -986 {
        case 653:
            if (100 == 43) {
                v1 = v14
                fmt.Println(v13, v16, v1)
                v2 = append(v2, 128)
            } else {
                v73 := GenS1{F0: 1}
                fmt.Println(v14, v1, v13)
                var v74 uint8 = 248
                _ = v73
                _ = v74
            }
        case -1:
            switch "\n\t" {
            case "a":
                v75 := GenS8{F0: 42, F1: v16}
                _ = v75
            case "":
                v76 := GenS8{F0: 0, F1: v16}
                _ = v76
            }
        case 400:
            var v77 bool = (0.1 > 1e10)
            _ = v77
        }
        for fi78 := 0; fi78 < 9; fi78++ {
            var v79 int64 = -907
            _ = fi78
            _ = v79
        }
        _ = v72
    }
    fmt.Println(v3, v16)
    v80 := GenS4{F0: "\n\t", F1: 255, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v80
}
