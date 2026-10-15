package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int8; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint8; F1 uint8 }
func genFn0() uint64 {
    var v5 uint32 = 648
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                v5 += (v5 << ((733) & 63))
                var v8 string = (("hello" + "\n\t") + (("foo bar" + "\n\t") + ""))
                if true {
                    v8 = "hello"
                    fmt.Println(fi7, v8, fi6)
                    v8 = ((v8 + v8) + "hello")
                } else {
                    switch fi7 {
                    case 9223372036854775807:
                        switch v8 {
                        case "a":
                            switch fi6 {
                            case 495:
                                var v9 int = (fi7 | fi7)
                                _ = v9
                            case -376:
                                v10 := []int{fi6, fi7}
                                _ = v10
                            }
                        case "":
                            v11 := GenS1{F0: 1, F1: 520, F2: (fi6 ^ fi7)}
                            _ = v11
                        default:
                            v5 += (v5 | 65535)
                        }
                    case 705:
                        if false {
                            if (927 > 1) {
                                v8 = v8
                                var v12 string = "foo bar"
                                _ = v12
                            } else {
                                v13 := []int{(fi7 % ((42) | 1)), fi7, 32768}
                                v8 = v8
                                _ = v13
                            }
                            var v14 uint64 = 18446744073709551615
                            fmt.Println(v8)
                            _ = v14
                        } else {
                            v15 := []int{fi6, fi7, (fi7 - fi7), (fi7 / ((32767) | 1))}
                            v5 += 2147483648
                            var v16 uint = 42
                            _ = v15
                            _ = v16
                        }
                    }
                }
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    _ = v5
    return 4294967295
}
func genFn1(p0 int64, p1 uint32) uint {
    p0 = p0
    for fi17 := 0; fi17 < 3; fi17++ {
        switch ("" + ("\n\t" + "hello")) {
        case "":
            switch -413 {
            case 0:
                var v18 bool = true
                _ = v18
            case 4294967295:
                var v19 uint8 = 209
                _ = v19
            case 42:
                var v20 uint8 = 128
                _ = v20
            }
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                p1 = p1
            default:
                for fi21 := 0; fi21 < 18; fi21++ {
                    if (true || false) {
                        if false {
                            v22 := GenS2{F0: (("foo bar" + "foo bar") + ("hello" + "a")), F1: (p0 / (((^(p0))) | 1)), F2: p0}
                            _ = v22
                        } else {
                            p0 -= -712
                            v23 := GenS2{F0: "\n\t", F1: (^(p0)), F2: -135}
                            _ = v23
                        }
                        v24 := GenS1{F0: 42, F1: 2, F2: 2147483647}
                        _ = v24
                    } else {
                        if (255 < 65535) {
                            p1 = (^(p1))
                            var v25 uint32 = 65535
                            _ = v25
                        } else {
                            p0 -= 9223372036854775807
                            v26 := GenS2{F0: "\n\t", F1: 9223372036854775807, F2: (p0 << ((42) & 63))}
                            switch (("foo bar" + "foo bar") + "") {
                            case "hello":
                                fmt.Println(fi21, fi17, fi21)
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    switch "\n\t" {
                                    case "hello":
                                        p1 += 2147483648
                                    }
                                case "foo bar":
                                    var v27 int64 = p0
                                    _ = v27
                                }
                            case "\n\t":
                                switch (("a" + "foo bar") + "") {
                                case "hello":
                                    v28 := GenS2{F0: ("\n\t" + ("a" + ("a" + ""))), F1: int64(p1), F2: (p0 + p0)}
                                    _ = v28
                                case "\n\t":
                                    v29 := GenS1{F0: 136, F1: 255, F2: (fi17 % ((fi17) | 1))}
                                    _ = v29
                                }
                            case "foo bar":
                                v30 := []int{int(p0), fi17}
                                _ = v30
                            default:
                                if (1e10 >= float64(1.5)) {
                                    var v31 uint16 = 42
                                    _ = v31
                                }
                            }
                            _ = v26
                        }
                        p1 = p1
                        var v32 string = "a"
                        _ = v32
                    }
                    v33 := []int{(^(fi17)), (-(fi21)), fi17}
                    _ = fi21
                    _ = v33
                }
            }
        default:
            v34 := GenS2{F0: ("" + (("a" + (((("\n\t" + ("a" + "foo bar")) + "hello") + ("a" + (("\n\t" + ("hello" + "a")) + "a"))) + "")) + "hello")), F1: -1, F2: 9223372036854775807}
            _ = v34
        }
        v35 := GenS1{F0: 42, F1: 255, F2: fi17}
        var v36 int64 = p0
        _ = fi17
        _ = v35
        _ = v36
    }
    fmt.Println(p1)
    switch (("\n\t" + "\n\t") + "hello") {
    case "a":
        for fi37 := 0; fi37 < 1; fi37++ {
            switch (fi37 << ((4294967295) & 63)) {
            case 32767:
                if (false || ((p0 >> ((0) & 63)) == p0)) {
                    if (42 != 860) {
                        var v38 uint = 4294967295
                        if (2 != (v38 ^ (v38 + v38))) {
                            p0 -= p0
                            fmt.Println(v38, v38, p1)
                            if (0 != 1) {
                                v39 := []int{-9223372036854775808, (fi37 / (((^(fi37))) | 1)), fi37, (fi37 * (-(fi37)))}
                                p1 = (p1 | 2147483647)
                                _ = v39
                            }
                        } else {
                            v40 := []int{-668}
                            _ = v40
                        }
                        p1 = p1
                        _ = v38
                    }
                    if true {
                        for fi41 := 0; fi41 < 8; fi41++ {
                            if false {
                                p1 -= p1
                                var v42 uint = 65535
                                var v43 int32 = 2147483647
                                _ = v42
                                _ = v43
                            }
                            v44 := []int{}
                            _ = fi41
                            _ = v44
                        }
                        if ((!false) && true) {
                            if (p1 > (p1 ^ (p1 % (((p1 >> ((698) & 63))) | 1)))) {
                                if ((!(4294967295 == 0)) && ((p0 >= p0) && false)) {
                                    for fi45 := 0; fi45 < 18; fi45++ {
                                        p1 = p1
                                        if false {
                                            if (42 <= 539) {
                                                v46 := GenS2{F0: ("" + (((((("\n\t" + "foo bar") + "a") + "foo bar") + "") + "hello") + "\n\t")), F1: p0, F2: p0}
                                                switch (fi45 & 42) {
                                                case -9223372036854775808:
                                                    if true {
                                                        if (!(!(393 < 4294967295))) {
                                                            switch fi45 {
                                                            case 32767:
                                                                v47 := []int{fi45, fi37, fi45}
                                                                _ = v47
                                                            case -2147483648:
                                                                fmt.Println(p0)
                                                            default:
                                                                v48 := []int{(fi45 % ((fi45) | 1)), (-(fi45)), (fi45 % ((fi45) | 1)), 20}
                                                                _ = v48
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v46
                                            }
                                            p1 = p1
                                        }
                                        _ = fi45
                                    }
                                    if (2 != 2) {
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(p1, p0, fi37)
                                        case "a":
                                            p1 = (p1 + p1)
                                        default:
                                            if (1 <= -1) {
                                                p0 -= (p0 ^ p0)
                                            } else {
                                                p0 = (p0 << ((683) & 63))
                                                var v49 int16 = 962
                                                v50 := GenS1{F0: 1, F1: 341, F2: fi37}
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                        var v51 uint64 = genFn0()
                                        _ = v51
                                    }
                                }
                                if (42 <= -735) {
                                    for fi52 := 0; fi52 < 20; fi52++ {
                                        p1 = (p1 >> ((18446744073709551615) & 63))
                                        v53 := GenS3{F0: 369}
                                        fmt.Println(p0, fi37)
                                        var v54 uint = 578
                                        var v55 int64 = (p0 << ((979) & 63))
                                        if true {
                                            fmt.Println(p1)
                                        }
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    var v56 float32 = (1.5 - ((1.5 + 1.5) * 3.14159))
                                    _ = v56
                                }
                            } else {
                                switch 4294967295 {
                                case 631:
                                    if (p0 == p0) {
                                        if (!(false && (!true))) {
                                            for fi57 := 0; fi57 < 12; fi57++ {
                                                v58 := GenS3{F0: (fi57 ^ fi57)}
                                                fmt.Println(fi57, fi37, fi57)
                                                var v59 int32 = 42
                                                _ = fi57
                                                _ = v58
                                                _ = v59
                                            }
                                            if true {
                                                for fi60 := 0; fi60 < 12; fi60++ {
                                                    fmt.Println(p0, fi60, p1)
                                                    _ = fi60
                                                }
                                                v61 := GenS2{F0: "", F1: (p0 - -9223372036854775808), F2: p0}
                                                p1 += (p1 + (p1 / (((p1 * p1)) | 1)))
                                                _ = v61
                                            } else {
                                                fmt.Println(fi37, p1, fi37)
                                                p0 += 1
                                            }
                                            if false {
                                                var v62 uint32 = p1
                                                _ = v62
                                            }
                                        }
                                    } else {
                                        var v63 float32 = (1e10 - float32((float64(0.1) + (0.0 * 1.5))))
                                        _ = v63
                                    }
                                }
                                var v64 uint = 78
                                _ = v64
                            }
                            fmt.Println(p0, p1, p1)
                        }
                    }
                }
            default:
                p1 = p1
            }
            _ = fi37
        }
    case "foo bar":
        for fi65 := 0; fi65 < 3; fi65++ {
            switch "a" {
            case "a":
                v66 := []int{fi65, (fi65 + 32767), (fi65 | -2147483648), (-(fi65))}
                _ = v66
            }
            var v67 bool = false
            _ = fi65
            _ = v67
        }
    case "hello":
        v68 := GenS3{F0: 127}
        _ = v68
    default:
        var v69 float32 = -2.25
        _ = v69
    }
    v70 := GenS4{F0: 42, F1: -32768}
    _ = p0
    _ = p1
    _ = v70
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint32 = 65535
    var v3 float64 = ((0.0 * 1e10) - 1e10)
    var v4 int64 = 9223372036854775807
    if ((!true) && v1) {
        for fi71 := 0; fi71 < 20; fi71++ {
            v72 := GenS2{F0: "hello", F1: (v4 - (v4 ^ (^(v4)))), F2: v4}
            fmt.Println(v4, v1)
            _ = fi71
            _ = v72
        }
        if (0.1 > 1.5) {
            for fi73 := 0; fi73 < 11; fi73++ {
                for fi74 := 0; fi74 < 10; fi74++ {
                    if v1 {
                        if false {
                            if (!true) {
                                v75 := []int{int(v2), -600, (fi73 + fi74)}
                                _ = v75
                            } else {
                                var v76 uint = 707
                                _ = v76
                            }
                            v77 := GenS5{F0: uint(v2), F1: "foo bar"}
                            if v1 {
                                var v78 int = 42
                                var v79 int64 = v4
                                switch v78 {
                                case 4294967295:
                                    fmt.Println(v4, v79, v78)
                                }
                                switch ("foo bar" + ("" + "foo bar")) {
                                case "a":
                                    fmt.Println(v4, v4, v3)
                                case "\n\t":
                                    switch 4294967295 {
                                    case -203:
                                        if true {
                                            v80 := []int{(fi73 + 549), v78, 42, fi73}
                                            _ = v80
                                        } else {
                                            v78 = -9223372036854775808
                                            v81 := []int{}
                                            _ = len(v81)
                                            _ = v81
                                        }
                                    default:
                                        var v82 float32 = (((1.5 * (float32(1e10) - 0.0)) + 1.5) / 0.1)
                                        _ = v82
                                    }
                                case "":
                                    v83 := []int{fi73}
                                    _ = v83
                                }
                                _ = v78
                                _ = v79
                            }
                            _ = v77
                        }
                        if v1 {
                            v84 := GenS5{F0: 673, F1: ("hello" + "\n\t")}
                            _ = v84
                        }
                        var v85 int8 = 42
                        _ = v85
                    } else {
                        v2 = v2
                    }
                    if v1 {
                        v1 = v1
                    } else {
                        var v86 int8 = -58
                        _ = v86
                    }
                    fmt.Println(v2, fi73)
                    v1 = (v1 && false)
                    var v87 float64 = v3
                    _ = fi74
                    _ = v87
                }
                _ = fi73
            }
            v88 := []int{-9223372036854775808, 9223372036854775807, 128}
            for fi89 := 0; fi89 < 16; fi89++ {
                var v90 int64 = 992
                _ = fi89
                _ = v90
            }
            _ = v88
        }
        v91 := GenS4{F0: -128, F1: -32768}
        _ = v91
    }
    v3 += v3
    v92 := []int{0, -9223372036854775808, 2147483647}
    var v93 bool = (4294967295 < 255)
    if false {
        v94 := GenS6{F0: 0, F1: 973, F2: 42}
        _ = v94
    }
    if (v1 && (32767 == 32767)) {
        if v1 {
            if (v4 <= (-(v4))) {
                var v95 uint8 = 248
                v92 = append(v92, 870)
                _ = v95
            }
            v96 := GenS7{F0: 2, F1: 1}
            _ = v96
        }
        v1 = v1
        switch 9223372036854775807 {
        case 313:
            v97 := GenS1{F0: 2, F1: 0, F2: 2147483647}
            _ = v97
        default:
            fmt.Println(v1, v3)
        }
    }
    var v98 uint16 = 42
    var v99 string = "foo bar"
    if v1 {
        fmt.Println(v1, v99)
        var v100 float64 = v3
        fmt.Println(v1)
        _ = v100
    }
    if v1 {
        v101 := GenS5{F0: 65535, F1: ("" + ("foo bar" + v99))}
        if len(v92) >= 2 {
            v92 = v92[1:len(v92)]
        }
        _ = v101
    } else {
        if v93 {
            if v1 {
                v99 = ("a" + ((v99 + v99) + "a"))
            }
            var v102 uint = 110
            _ = v102
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v93:", v93)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v92
    _ = v93
    _ = v98
    _ = v99
}
