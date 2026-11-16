package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 int }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 string, p1 uint8) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            var v6 int16 = 42
            v7 := GenS1{F0: 42, F1: -2147483648, F2: 370}
            var v8 float64 = (1e10 * 0.0)
            _ = v6
            _ = v7
            _ = v8
        } else {
            fmt.Println(p1)
        }
        v9 := GenS1{F0: 42, F1: int32(fi5), F2: 2}
        _ = fi5
        _ = v9
    }
    switch -9223372036854775808 {
    case -845:
        var v10 uint32 = 0
        _ = v10
    case -844:
        v11 := GenS2{F0: p1, F1: 20, F2: -1}
        _ = v11
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: 1, F1: int32(p1), F2: 32767}
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS3{F0: 425, F1: 42, F2: (fi14 / ((fi14) | 1))}
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 string = p0
                v18 := GenS1{F0: 0, F1: 2147483647, F2: 65535}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi12
            _ = v13
        }
    }
    var v19 float32 = 3.14159
    v20 := GenS3{F0: 1, F1: 444, F2: 1}
    v21 := GenS3{F0: -32768, F1: 42, F2: 187}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 74
}
func genFn1(p0 int16, p1 uint32) uint8 {
    for fi22 := 0; fi22 < 2; fi22++ {
        switch ("a" + "hello") {
        case "foo bar":
            if true {
                p0 = (p0 >> ((42) & 63))
            } else {
                v23 := []int{(^(fi22))}
                p0 = (p0 & (p0 / ((42) | 1)))
                _ = v23
            }
        case "hello":
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS4{F0: 1, F1: fi22, F2: 851}
                var v26 int64 = -1
                v27 := []int{(fi22 * fi22), (-(fi24)), fi22}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        default:
            switch ("" + "a") {
            case "foo bar":
                fmt.Println(fi22)
            case "hello":
                fmt.Println(p1, p0, p0)
            default:
                v28 := []int{}
                _ = v28
            }
        }
        _ = fi22
    }
    switch 4294967295 {
    case 587:
        var v29 float64 = (float64(1.5) * float64((-2.25 - ((1e10 * 0.1) / 3.14159))))
        _ = v29
    case 42:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return 31
}

func main() {
    var v1 uint = 1
    var v2 int16 = 0
    var v3 int64 = 1
    var v4 float32 = 1e10
    var v30 uint64 = 4294967295
    switch genFn0("a", 128) {
    case -505:
        if (128 > 295) {
            if (!((^(v30)) != 961)) {
                v31 := GenS4{F0: 471, F1: 32767, F2: uint(v30)}
                _ = v31
            }
            v30 -= v30
        }
    }
    if false {
        v32 := []int{}
        var v33 bool = (0.0 == (1.5 + 1.5))
        _ = v32
        _ = v33
    }
    var v34 int64 = (v3 % ((0) | 1))
    var v35 float32 = v4
    for fi36 := 0; fi36 < 15; fi36++ {
        v37 := []int{(fi36 * fi36), 32768, (fi36 - 128)}
        var v38 int64 = v34
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint64 = v30
    for fi40 := 0; fi40 < 2; fi40++ {
        if (!((^(v1)) < (v1 << (((v1 - 4294967295)) & 63)))) {
            if ((v1 << ((v1) & 63)) == 4294967295) {
                v41 := []int{fi40, fi40, fi40}
                var v42 int8 = -128
                for fi43 := 0; fi43 < 4; fi43++ {
                    switch (fi40 % (((fi40 | fi40)) | 1)) {
                    case 32768:
                        switch (("hello" + "a") + ("hello" + "a")) {
                        case "foo bar":
                            v39 += (v39 + v30)
                        }
                    case -2147483648:
                        v44 := GenS5{F0: (v1 - (v1 >> ((v1) & 63)))}
                        _ = v44
                    case 128:
                        if (float64(v35) == 1e10) {
                            switch genFn0(("" + "hello"), 42) {
                            case 32768:
                                if len(v41) >= 2 {
                                    v41 = v41[1:len(v41)]
                                }
                            case 32767:
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        var v45 bool = (v39 < v39)
                                        v30 += (v30 ^ v30)
                                        var v46 uint = v1
                                        _ = v45
                                        _ = v46
                                    }
                                case "\n\t":
                                    if len(v41) > 0 {
                                        v41[0] = (^(fi43))
                                    }
                                case "foo bar":
                                    var v47 float32 = ((1.5 / 1e10) + v35)
                                    _ = v47
                                }
                            case -2147483648:
                                if false {
                                    if true {
                                        if len(v41) > 0 {
                                            v41[0] = (fi40 | (fi40 % (((fi40 >> ((v1) & 63))) | 1)))
                                        }
                                    } else {
                                        switch (fi40 | genFn0("a", 15)) {
                                        case 981:
                                            fmt.Println(v34, v2)
                                        case -269:
                                            var v48 float32 = v35
                                            _ = v48
                                        default:
                                            v41 = append(v41, fi43)
                                        }
                                        var v49 bool = false
                                        _ = v49
                                    }
                                    v30 = v39
                                    v50 := GenS6{F0: "foo bar", F1: 811}
                                    _ = v50
                                } else {
                                    var v51 uint = v1
                                    switch "a" {
                                    case "a":
                                        switch "a" {
                                        case "a":
                                            v41 = append(v41, (fi43 - (-(fi43))))
                                        case "\n\t":
                                            switch (("hello" + "\n\t") + "hello") {
                                            case "foo bar":
                                                var v52 int16 = (v2 * 32767)
                                                _ = v52
                                            case "hello":
                                                if false {
                                                    v4 = v4
                                                    v42 += v42
                                                    var v53 bool = (!(!false))
                                                    _ = v53
                                                } else {
                                                    var v54 bool = (!(true || true))
                                                    fmt.Println(fi40, v39)
                                                    _ = v54
                                                }
                                            case "\n\t":
                                                switch (fi40 << (((^(v1))) & 63)) {
                                                case 29:
                                                    fmt.Println(v2)
                                                case -9223372036854775808:
                                                    var v55 int16 = (v2 & (v2 >> (((^(v51))) & 63)))
                                                    _ = v55
                                                case 128:
                                                    var v56 uint = 961
                                                    _ = v56
                                                default:
                                                    switch (((((("foo bar" + ("foo bar" + "foo bar")) + (("hello" + ("\n\t" + ("a" + "\n\t"))) + ((("\n\t" + "hello") + ("foo bar" + ("" + ("hello" + "foo bar")))) + ("hello" + "")))) + "foo bar") + "hello") + ("hello" + "foo bar")) + (("foo bar" + "a") + "hello")) {
                                                    case "a":
                                                        var v57 int = (fi40 << ((4294967295) & 63))
                                                        _ = v57
                                                    case "\n\t":
                                                        var v58 int8 = (v42 % ((-84) | 1))
                                                        var v59 float32 = v35
                                                        _ = v58
                                                        _ = v59
                                                    default:
                                                        var v60 uint8 = 1
                                                        _ = v60
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        var v61 int32 = -2147483648
                                                        _ = v61
                                                    case "foo bar":
                                                        if len(v41) > 0 {
                                                            v41[0] = 128
                                                        }
                                                    case "hello":
                                                        v62 := GenS7{F0: (v2 << (((v1 - (^(v51)))) & 63)), F1: uint32(v51), F2: 129}
                                                        _ = v62
                                                    default:
                                                        v34 -= (v34 >> (((^(v1))) & 63))
                                                    }
                                                    var v63 uint64 = v30
                                                    _ = v63
                                                }
                                            }
                                        default:
                                            var v64 int16 = v2
                                            _ = v64
                                        }
                                    case "foo bar":
                                        var v65 bool = true
                                        _ = v65
                                    case "hello":
                                        fmt.Println(v1, v1)
                                    default:
                                        v51 -= (v1 & (v1 >> (((v1 & (^(v1)))) & 63)))
                                    }
                                    var v66 uint = 255
                                    _ = v51
                                    _ = v66
                                }
                            default:
                                switch ("a" + "foo bar") {
                                case "hello":
                                    v2 -= 32767
                                case "a":
                                    var v67 uint8 = 0
                                    _ = v67
                                case "\n\t":
                                    v39 -= (v30 & (v30 << (((^(v1))) & 63)))
                                }
                            }
                            var v68 int16 = 32767
                            v69 := GenS7{F0: (v68 * int16(fi40)), F1: 65535, F2: genFn1(v2, 2)}
                            _ = v68
                            _ = v69
                        }
                    default:
                        fmt.Println(fi40)
                    }
                    var v70 int32 = 2147483647
                    _ = fi43
                    _ = v70
                }
                _ = v41
                _ = v42
            }
            for fi71 := 0; fi71 < 18; fi71++ {
                if true {
                    switch -2147483648 {
                    case 9223372036854775807:
                        v72 := GenS4{F0: 1, F1: fi71, F2: (v1 >> ((1) & 63))}
                        var v73 uint8 = 128
                        if false {
                            fmt.Println(v34)
                        }
                        _ = v72
                        _ = v73
                    case 32767:
                        if true {
                            v1 = (v1 << (((^(v1))) & 63))
                            switch 32767 {
                            case 499:
                                if (255 < 255) {
                                    v74 := []int{fi40}
                                    _ = v74
                                } else {
                                    if (!(!(-32768 >= 0))) {
                                        var v75 int64 = (v34 - v34)
                                        _ = v75
                                    } else {
                                        v76 := GenS5{F0: v1}
                                        switch "foo bar" {
                                        case "hello":
                                            if (false && false) {
                                                v3 += -9223372036854775808
                                                if (!false) {
                                                    v77 := []int{}
                                                    _ = v77
                                                } else {
                                                    v78 := GenS8{F0: 0, F1: 2147483647}
                                                    v1 = 1
                                                    _ = v78
                                                }
                                                var v79 int16 = v2
                                                _ = v79
                                            } else {
                                                if true {
                                                    fmt.Println(fi40, v1, v2)
                                                    fmt.Println(v1, v2)
                                                    if (!true) {
                                                        var v80 uint32 = 713
                                                        var v81 uint = v1
                                                        var v82 int = (fi71 >> ((v1) & 63))
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                    } else {
                                                        fmt.Println(v30, v4, v2)
                                                        var v83 uint16 = 1
                                                        _ = v83
                                                    }
                                                }
                                            }
                                        default:
                                            v84 := []int{fi71, -678, (fi71 >> ((v1) & 63))}
                                            _ = v84
                                        }
                                        switch fi40 {
                                        case 1:
                                            var v85 uint32 = 42
                                            _ = v85
                                        case -1:
                                            v86 := []int{}
                                            _ = v86
                                        default:
                                            v30 = 702
                                        }
                                        _ = v76
                                    }
                                }
                            }
                        }
                    case -2147483648:
                        v30 -= 974
                    default:
                        switch int(v30) {
                        case -1:
                            v87 := GenS9{F0: 65535}
                            _ = v87
                        }
                    }
                }
                var v88 uint = (v1 * 4294967295)
                var v89 uint32 = 0
                _ = fi71
                _ = v88
                _ = v89
            }
            v90 := []int{fi40, fi40, fi40, -708}
            _ = v90
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
