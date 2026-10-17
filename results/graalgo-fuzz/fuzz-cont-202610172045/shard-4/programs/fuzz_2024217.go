package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint32; F1 int64; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64; F1 bool; F2 int32 }
type GenS11 struct { F0 uint64; F1 string }
func genFn0() uint8 {
    if false {
        var v5 int16 = 1
        _ = v5
    }
    if true {
        v6 := []int{506, 1, 106}
        _ = v6
    } else {
        v7 := []int{-9223372036854775808, -463}
        switch -354 {
        case 32768:
            var v8 int8 = 127
            _ = v8
        case 127:
            v9 := GenS1{F0: "hello", F1: 0}
            _ = v9
        default:
            var v10 int64 = 0
            _ = v10
        }
        _ = v7
    }
    var v11 uint32 = 589
    v12 := []int{32767, 839}
    _ = v11
    _ = v12
    return 42
}
func genFn1(p0 uint32, p1 int64, p2 int64) uint8 {
    p2 += p1
    v13 := GenS2{F0: 2}
    if false {
        v14 := GenS2{F0: p0}
        _ = v14
    }
    v15 := []int{32768, 1, -881}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v15
    return 93
}

func main() {
    var v1 uint = 319
    var v2 uint64 = 2
    var v3 int32 = 2147483647
    var v4 float64 = 1e10
    for fi16 := 0; fi16 < 6; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            switch "\n\t" {
            case "a":
                v2 = 4294967295
            case "hello":
                var v18 bool = true
                _ = v18
            case "foo bar":
                v19 := []int{(fi16 / ((fi16) | 1)), (fi17 - (-(fi16))), fi17, fi17}
                _ = v19
            }
            _ = fi17
        }
        var v20 uint = v1
        v21 := GenS3{F0: true, F1: (^(v1))}
        _ = fi16
        _ = v20
        _ = v21
    }
    v4 += (1e10 * v4)
    if true {
        v22 := []int{}
        _ = v22
    }
    v2 = v2
    v23 := []int{-908}
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS3{F0: true, F1: v1}
        switch 64 {
        case 0:
            if len(v23) > 0 {
                v23[0] = -931
            }
        }
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        if ((v3 << ((v1) & 63)) < 657) {
            switch "hello" {
            case "":
                if true {
                    var v27 bool = true
                    fmt.Println(v4)
                    fmt.Println(v3, v1, v1)
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        if (!false) {
                            v29 := GenS3{F0: ((float32(v4) * (0.0 + float32(float64(-2.25)))) >= (3.14159 / 3.14159)), F1: v1}
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            v4 += ((-2.25 - v4) - v4)
                            _ = v29
                        } else {
                            if false {
                                switch "\n\t" {
                                case "":
                                    switch "\n\t" {
                                    case "a":
                                        v23 = append(v23, fi26)
                                    case "hello":
                                        v30 := GenS1{F0: "foo bar", F1: fi28}
                                        _ = v30
                                    }
                                case "\n\t":
                                    fmt.Println(fi26)
                                case "hello":
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                default:
                                    v31 := GenS4{F0: ("a" + ("foo bar" + "a")), F1: (-(fi28))}
                                    _ = v31
                                }
                            } else {
                                var v32 float64 = v4
                                v1 -= (v1 | (v1 - v1))
                                _ = v32
                            }
                            var v33 bool = false
                            if v33 {
                                v4 = 1e10
                                v1 = v1
                                var v34 uint8 = 0
                                _ = v34
                            }
                            _ = v33
                        }
                        fmt.Println(fi26, v2, v3)
                        _ = fi28
                    }
                    if false {
                        v35 := GenS2{F0: 2}
                        switch (("a" + ("foo bar" + (((("a" + "") + "hello") + "\n\t") + ""))) + "\n\t") {
                        case "a":
                            for fi36 := 0; fi36 < 16; fi36++ {
                                _ = len(v23)
                                v3 += (v3 / ((v3) | 1))
                                if false {
                                    var v37 int64 = 0
                                    if (4294967295 == uint32(v1)) {
                                        if (!true) {
                                            v2 = v2
                                            fmt.Println(v37, v3)
                                            if len(v23) > 0 {
                                                v23[0] = 1
                                            }
                                        }
                                    }
                                    _ = v37
                                }
                                _ = fi36
                            }
                        case "hello":
                            switch fi26 {
                            case 1:
                                v2 = (v2 & v2)
                            case 329:
                                v4 += (v4 + ((v4 - 0.0) / (-2.25 / v4)))
                            case 192:
                                var v38 float32 = (0.0 / -2.25)
                                _ = v38
                            }
                        case "\n\t":
                            v4 = 0.0
                        }
                        _ = v35
                    } else {
                        v39 := GenS2{F0: 435}
                        v40 := GenS1{F0: (("foo bar" + ("a" + "\n\t")) + ("hello" + "hello")), F1: -958}
                        var v41 uint64 = v2
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
            case "hello":
                for fi42 := 0; fi42 < 12; fi42++ {
                    if false {
                        v43 := GenS5{F0: 0, F1: 17, F2: (((("a" + "\n\t") + ("" + "a")) + ("foo bar" + "\n\t")) + (("hello" + "a") + "foo bar"))}
                        fmt.Println(v4, fi26, v1)
                        _ = v43
                    } else {
                        var v44 string = ("" + (("hello" + "a") + ("hello" + "")))
                        switch fi42 {
                        case 4294967295:
                            v1 -= (v1 ^ 2)
                        case -315:
                            switch ((((v44 + (v44 + "a")) + "\n\t") + "foo bar") + ((v44 + (("\n\t" + "") + v44)) + "hello")) {
                            case "hello":
                                var v45 uint32 = 2
                                _ = v45
                            case "\n\t":
                                v46 := GenS3{F0: (fi42 <= (fi26 + fi42)), F1: (v1 ^ v1)}
                                _ = v46
                            case "":
                                v44 = ("hello" + (v44 + ""))
                            }
                        case -929:
                            var v47 int = (fi26 % ((798) | 1))
                            _ = v47
                        }
                        if false {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            switch "a" {
                            case "\n\t":
                                var v48 int64 = 0
                                if (v48 == v48) {
                                    if len(v23) > 0 {
                                        v23[0] = (^(fi42))
                                    }
                                } else {
                                    switch (((v44 + ("\n\t" + (v44 + "a"))) + "") + "foo bar") {
                                    case "\n\t":
                                        fmt.Println(v44, v4, v1)
                                    default:
                                        if (fi42 <= (fi42 + (-(fi42)))) {
                                            var v49 uint32 = 42
                                            if len(v23) > 0 {
                                                v23[0] = -55
                                            }
                                            _ = v49
                                        }
                                    }
                                    v48 = 176
                                }
                                _ = v48
                            case "hello":
                                var v50 int16 = -1
                                _ = v50
                            }
                            fmt.Println(fi42)
                        } else {
                            fmt.Println(v4, fi26, v4)
                            v51 := GenS3{F0: true, F1: v1}
                            switch v44 {
                            case "foo bar":
                                v3 = 1
                            case "":
                                if false {
                                    var v52 string = ((v44 + v44) + v44)
                                    switch 32768 {
                                    case 32768:
                                        v44 = "hello"
                                    case 475:
                                        var v53 uint32 = 2147483648
                                        v2 -= v2
                                        _ = v53
                                    case 4294967295:
                                        var v54 int = (fi42 >> (((v1 | (v1 % ((42) | 1)))) & 63))
                                        _ = v54
                                    default:
                                        var v55 int32 = v3
                                        _ = v55
                                    }
                                    v2 = (v2 | v2)
                                    switch fi42 {
                                    case 1:
                                        v23 = append(v23, -1)
                                    case -1:
                                        var v56 float32 = float32(v4)
                                        _ = v56
                                    }
                                    v2 = v2
                                    var v57 string = (v44 + v44)
                                    if true {
                                        v2 = v2
                                        var v58 uint8 = 2
                                        v59 := GenS5{F0: 540, F1: -481, F2: ("hello" + v44)}
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v60 := GenS6{F0: v1}
                                        if (!(((false && true) && (!true)) || (true || (v1 < (v1 >> ((212) & 63)))))) {
                                            v44 = v52
                                            v23 = append(v23, 32767)
                                        }
                                        _ = v60
                                    }
                                    _ = v52
                                    _ = v57
                                }
                            default:
                                if true {
                                    switch (fi26 % ((127) | 1)) {
                                    case 1:
                                        if false {
                                            _ = len(v23)
                                            var v61 int64 = -349
                                            _ = v61
                                        } else {
                                            var v62 float64 = (v4 + v4)
                                            v62 += 0.0
                                            if (1 < -1) {
                                                fmt.Println(fi42, v4)
                                                var v63 int = -678
                                                _ = v63
                                            } else {
                                                if (((false && false) || (((false && ((genFn1(uint32(v2), 1, -1) < 0) || true)) || false) && false)) || true) {
                                                    var v64 uint8 = 42
                                                    if true {
                                                        fmt.Println(v1)
                                                    } else {
                                                        v23 = append(v23, fi42)
                                                    }
                                                    v65 := GenS7{F0: 0, F1: (v3 << (((v1 / (((v1 >> ((57) & 63))) | 1))) & 63))}
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    var v66 string = v44
                                                    v23 = append(v23, fi42)
                                                    if (!(-32768 != 1000)) {
                                                        v67 := GenS8{F0: 2}
                                                        v1 -= 967
                                                        _ = v67
                                                    }
                                                    var v68 uint32 = 2
                                                    var v69 uint16 = 595
                                                    fmt.Println(fi26)
                                                    switch fi42 {
                                                    case -2147483648:
                                                        switch (fi42 | fi42) {
                                                        case 0:
                                                            v69 = (v69 | 257)
                                                        case 127:
                                                            v66 = ("" + v66)
                                                        case -974:
                                                            if false {
                                                                var v70 int16 = 0
                                                                v1 = 2
                                                                _ = v70
                                                            } else {
                                                                switch (fi42 - fi26) {
                                                                case 0:
                                                                    if true {
                                                                        var v71 float32 = 0.0
                                                                        _ = v71
                                                                    }
                                                                case -1:
                                                                    v66 = ("a" + v66)
                                                                    fmt.Println(v66, fi26, v44)
                                                                case -416:
                                                                    var v72 bool = (1 != 0)
                                                                    _ = v72
                                                                }
                                                                var v73 uint = (v1 / ((501) | 1))
                                                                fmt.Println(v66, v68, v68)
                                                                v74 := GenS3{F0: true, F1: v73}
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                    case -1:
                                                        var v75 int16 = 32767
                                                        _ = v75
                                                    case -9223372036854775808:
                                                        _ = len(v23)
                                                    }
                                                    _ = v66
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                            _ = v62
                                        }
                                    case -879:
                                        if len(v23) > 0 {
                                            v23[0] = fi26
                                        }
                                    }
                                    fmt.Println(v1)
                                } else {
                                    v76 := GenS4{F0: ("" + ""), F1: 891}
                                    v1 = v1
                                    v44 = (v44 + (v44 + "foo bar"))
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v2 = v2
                                    _ = v76
                                }
                            }
                            var v77 bool = true
                            v78 := GenS9{F0: -2147483648}
                            _ = v51
                            _ = v77
                            _ = v78
                        }
                        _ = v44
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if (!(true && (65535 != 65535))) {
                            v79 := GenS4{F0: ((("" + ("" + "a")) + "hello") + "foo bar"), F1: 837}
                            _ = v79
                        } else {
                            v80 := GenS10{F0: 2, F1: false, F2: 0}
                            _ = v80
                        }
                    case "\n\t":
                        v23 = append(v23, 0)
                    }
                    _ = fi42
                }
            }
        }
        if true {
            v1 -= v1
        } else {
            if (false || (-1 < 608)) {
                fmt.Println(v3, fi26)
                v81 := GenS6{F0: 2}
                _ = v81
            } else {
                var v82 int8 = 102
                if true {
                    if (v2 == (v2 / ((v2) | 1))) {
                        v2 -= (v2 / ((v2) | 1))
                        v4 += v4
                    }
                }
                v83 := GenS11{F0: (v2 + v2), F1: "foo bar"}
                _ = v82
                _ = v83
            }
        }
        _ = fi26
    }
    switch -945 {
    case 2147483647:
        fmt.Println(v2)
    default:
        var v84 string = "hello"
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
