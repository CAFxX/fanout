package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint8; F1 int; F2 int64 }
type GenS8 struct { F0 int32; F1 uint; F2 bool }
type GenS9 struct { F0 uint32; F1 int64; F2 int64 }
type GenS10 struct { F0 uint8 }
func genFn0(p0 uint, p1 string) uint64 {
    v7 := GenS1{F0: p0, F1: 0, F2: -27}
    v8 := []int{-1, 0, 4294967295, -629}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 int32) int64 {
    p0 = (^(p0))
    var v9 uint64 = 101
    v10 := GenS2{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return int64(v9)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 100
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := []string{"", ("" + (("a" + "hello") + ("a" + "hello"))), ("hello" + "")}
    var v6 uint32 = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        _ = len(v5)
        _ = fi11
    }
    var v12 int8 = v2
    delete(v3, (^(v6)))
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v6, v2)
        _ = fi13
    }
    if ((!false) && false) {
        v6 -= (v6 + (v6 | (v6 ^ (^(v6)))))
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                if true {
                    if (!false) {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            var v16 bool = false
                            var v17 float32 = float32(1e10)
                            fmt.Println(v12)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                }
                v18 := GenS2{F0: 0}
                var v19 int64 = -1
                _ = v18
                _ = v19
            } else {
                v2 -= 42
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 float32 = (-2.25 / 3.14159)
                    var v22 bool = true
                    var v23 uint16 = 255
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                v12 = (v12 & -20)
                _ = len(v4)
                _ = fi24
            }
            if (!(42 == (fi14 % ((fi14) | 1)))) {
                if (4294967295 == 1) {
                    v2 = (v12 * v2)
                } else {
                    if (1 > 108) {
                        if ((127 == 0) && (!false)) {
                            var v25 string = "a"
                            var v26 uint64 = 18446744073709551615
                            _ = v25
                            _ = v26
                        }
                        switch "hello" {
                        case "a":
                            v2 -= v2
                        case "foo bar":
                            var v27 float32 = 1e10
                            _ = v27
                        case "\n\t":
                            v6 -= v6
                        default:
                            v2 = v12
                        }
                    } else {
                        if (-1 >= 0) {
                            v28 := GenS2{F0: -9223372036854775808}
                            if true {
                                switch -1 {
                                case -150:
                                    var v29 uint8 = 0
                                    _ = v29
                                default:
                                    v2 -= 127
                                }
                                var v30 uint32 = (^(v6))
                                _ = v30
                            } else {
                                for fi31 := 0; fi31 < 20; fi31++ {
                                    v6 -= (v6 / ((v6) | 1))
                                    _ = fi31
                                }
                                var v32 bool = false
                                _ = v32
                            }
                            _ = v28
                        } else {
                            var v33 float64 = (float64((-2.25 * ((float32(1.5) / 3.14159) / (1e10 + float32(0.0))))) + (float64(0.0) * 3.14159))
                            switch fi14 {
                            case 32768:
                                var v34 uint64 = 4294967295
                                _ = v34
                            default:
                                var v35 bool = (false && (fi14 <= (fi14 - (-(fi14)))))
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v6 = v6
                                switch ("a" + "hello") {
                                case "":
                                    v2 += (v2 & (^(v2)))
                                    v2 = (v12 / ((2) | 1))
                                case "hello":
                                    switch (("" + ("hello" + "foo bar")) + ("hello" + "a")) {
                                    case "\n\t":
                                        fmt.Println(v12, fi36, v33)
                                    case "foo bar":
                                        if false {
                                            if ("hello" == "a") {
                                                v33 -= -2.25
                                                if (true && true) {
                                                    switch fi36 {
                                                    case 2147483647:
                                                        switch (fi14 << ((2) & 63)) {
                                                        case -777:
                                                            v12 += v2
                                                        case -9223372036854775808:
                                                            var v37 bool = (1 >= genFn0(uint(v6), ""))
                                                            _ = v37
                                                        case -1:
                                                            v33 -= (float64(0.1) / v33)
                                                        }
                                                    case 128:
                                                        v38 := v3[(v6 / ((v6) | 1))]
                                                        _ = v38
                                                    default:
                                                        switch (fi36 | fi36) {
                                                        case 9223372036854775807:
                                                            v39 := GenS3{F0: 0, F1: 2147483647, F2: (false && (!(!(v33 != -2.25))))}
                                                            var v40 string = "foo bar"
                                                            v41 := GenS3{F0: 311, F1: -2147483648, F2: false}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        case 42:
                                                            fmt.Println(v33)
                                                        }
                                                    }
                                                    v12 += (v2 + (v2 / ((v2) | 1)))
                                                    var v42 uint32 = (v6 | v6)
                                                    if (1 > 1) {
                                                        v42 -= 2147483647
                                                        if true {
                                                            v43 := GenS4{F0: 128}
                                                            v44 := GenS5{F0: "a", F1: 165}
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        if (42 > 42) {
                                                            var v45 int64 = 1
                                                            v46 := GenS6{F0: "foo bar", F1: ((v33 == (((v33 + 0.0) + v33) * v33)) && false)}
                                                            v47 := GenS6{F0: "", F1: (!false)}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            var v48 int16 = -1
                                                            v2 += (v2 & v12)
                                                            _ = v48
                                                        }
                                                    }
                                                    var v49 float64 = (v33 / v33)
                                                    _ = len(v5)
                                                    v2 += (v2 * (v12 + v12))
                                                    switch (fi36 >> ((65535) & 63)) {
                                                    case 9223372036854775807:
                                                        fmt.Println(fi36, v2)
                                                    case 42:
                                                        var v50 uint = 45
                                                        _ = v50
                                                    case -9223372036854775808:
                                                        v51 := GenS1{F0: 18446744073709551615, F1: genFn0(1, "a"), F2: (v12 | (v2 % (((v12 | -128)) | 1)))}
                                                        _ = v51
                                                    default:
                                                        if false {
                                                            var v52 uint8 = 60
                                                            v12 = v12
                                                            v53 := GenS4{F0: v52}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        fmt.Println(v6, v12)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v12 -= v12
                                                        }
                                                    }
                                                    v12 += (-(v2))
                                                    _ = v42
                                                    _ = v49
                                                }
                                            } else {
                                                v54 := GenS5{F0: "\n\t", F1: 90}
                                                _ = v54
                                            }
                                        } else {
                                            switch "" {
                                            case "":
                                                v55 := GenS7{F0: 1, F1: 42, F2: 721}
                                                _ = v55
                                            }
                                            fmt.Println(fi14)
                                            if (true || true) {
                                                var v56 uint32 = 2
                                                switch "a" {
                                                case "a":
                                                    v56 = (v56 & (v6 * v6))
                                                case "\n\t":
                                                    if (-2.25 != 1e10) {
                                                        fmt.Println(v2)
                                                        switch "foo bar" {
                                                        case "":
                                                            v2 = (v2 + (v2 - (v2 | (v12 & v12))))
                                                        case "a":
                                                            v2 = 42
                                                            if true {
                                                                var v57 float32 = ((((0.0 + 1.5) / 1.0) / 1.0) + 3.14159)
                                                                switch (((("" + "foo bar") + ("a" + "foo bar")) + (("\n\t" + ("" + "\n\t")) + "hello")) + ("hello" + "a")) {
                                                                case "":
                                                                    v57 -= 1e10
                                                                case "hello":
                                                                    fmt.Println(v6)
                                                                case "a":
                                                                    v58 := GenS6{F0: (((((("hello" + "") + "hello") + "\n\t") + "") + "") + (("a" + ("" + "\n\t")) + "foo bar")), F1: true}
                                                                    _ = v58
                                                                default:
                                                                    v59 := GenS3{F0: 4294967295, F1: 179, F2: ("foo bar" < "a")}
                                                                    _ = v59
                                                                }
                                                                _ = v57
                                                            } else {
                                                                fmt.Println(v33, v6)
                                                                var v60 int8 = (^(v2))
                                                                _ = v60
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = ("\n\t" + "a")
                                                            }
                                                            fmt.Println(v6)
                                                            switch fi14 {
                                                            case 2147483647:
                                                                v12 = v2
                                                            default:
                                                                var v61 float32 = 0.1
                                                                _ = v61
                                                            }
                                                            if false {
                                                                if true {
                                                                    var v62 int8 = (v2 / (((v12 << ((553) & 63))) | 1))
                                                                    fmt.Println(fi14)
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if true {
                                                                            fmt.Println(fi36, v12)
                                                                            switch 936 {
                                                                            case -331:
                                                                                switch (^(fi14)) {
                                                                                case 32768:
                                                                                    fmt.Println(v6)
                                                                                case -826:
                                                                                    var v63 int32 = -1
                                                                                    _ = v63
                                                                                default:
                                                                                    _ = len(v1)
                                                                                }
                                                                            case 32768:
                                                                                v64 := GenS3{F0: 426, F1: 2147483647, F2: true}
                                                                                _ = v64
                                                                            case -2147483648:
                                                                                var v65 string = ("\n\t" + "a")
                                                                                _ = v65
                                                                            }
                                                                            var v66 int32 = -2147483648
                                                                            _ = v66
                                                                        } else {
                                                                            var v67 int = fi36
                                                                            _ = v67
                                                                        }
                                                                    case "":
                                                                        var v68 float64 = v33
                                                                        _ = v68
                                                                    case "foo bar":
                                                                        var v69 uint32 = v56
                                                                        _ = v69
                                                                    default:
                                                                        v70 := v3[(v6 * v56)]
                                                                        _ = v70
                                                                    }
                                                                    if ((false && (127 < 255)) || false) {
                                                                        if false {
                                                                            v56 += v56
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v62 = v2
                                                                            v33 = 0.1
                                                                        } else {
                                                                            var v71 int32 = 0
                                                                            v56 += v6
                                                                            v72 := GenS8{F0: (v71 + (v71 - v71)), F1: 967, F2: true}
                                                                            fmt.Println(v33, v33)
                                                                            v73 := GenS1{F0: 2, F1: 2, F2: v12}
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                    _ = v62
                                                                }
                                                            }
                                                        }
                                                        v56 += v56
                                                    } else {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v74 uint8 = 0
                                                        v75 := GenS9{F0: v56, F1: -947, F2: 824}
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                                _ = v56
                                            }
                                        }
                                    case "a":
                                        if false {
                                            var v76 uint = 255
                                            _ = v76
                                        } else {
                                            var v77 uint64 = 1
                                            _ = v77
                                        }
                                        _ = len(v1)
                                    }
                                case "foo bar":
                                    v78 := GenS8{F0: -455, F1: 1, F2: (false || ("a" <= "foo bar"))}
                                    _ = v78
                                }
                                _ = fi36
                            }
                            _ = v33
                        }
                        v2 -= (v12 + v2)
                        fmt.Println(v2)
                    }
                    v2 = v12
                    var v79 int32 = 0
                    _ = v79
                }
                var v80 int64 = genFn1(2, 0)
                if (((0.1 + (0.0 * 0.0)) < 1.5) && (false || false)) {
                    v6 = v6
                    if false {
                        if (("" + "") > "foo bar") {
                            v81 := GenS8{F0: 1, F1: 18446744073709551615, F2: true}
                            var v82 int8 = v2
                            _ = v81
                            _ = v82
                        } else {
                            var v83 bool = true
                            _ = v83
                        }
                        if true {
                            v1[(v6 & 65535)] = ""
                        } else {
                            var v84 int16 = -640
                            fmt.Println(v12)
                            v12 = (v12 / ((v2) | 1))
                            _ = v84
                        }
                    } else {
                        fmt.Println(v12, v12)
                        var v85 int32 = 393
                        _ = v85
                    }
                } else {
                    var v86 bool = true
                    if (((0 == 100) && ((^(fi14)) >= (fi14 << ((2) & 63)))) || ((!((v86 || (v86 || true)) && v86)) && ((fi14 << ((2) & 63)) >= fi14))) {
                        switch "foo bar" {
                        case "\n\t":
                            for fi87 := 0; fi87 < 9; fi87++ {
                                var v88 uint16 = 885
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi87
                                _ = v88
                            }
                        case "hello":
                            var v89 uint16 = 1
                            _ = v89
                        default:
                            fmt.Println(v12)
                        }
                        for fi90 := 0; fi90 < 5; fi90++ {
                            var v91 uint8 = 255
                            v6 += (v6 + 29)
                            _ = fi90
                            _ = v91
                        }
                        v2 = (v2 - (v2 + int8(v80)))
                    } else {
                        fmt.Println(v12)
                    }
                    v86 = v86
                    _ = v86
                }
                _ = v80
            }
            _ = fi14
        }
    } else {
        if true {
            var v92 float32 = 1.5
            _ = v92
        } else {
            if false {
                var v93 uint16 = 65535
                var v94 bool = false
                _ = v93
                _ = v94
            } else {
                v95 := GenS6{F0: "foo bar", F1: false}
                _ = v95
            }
        }
        switch (("hello" + ("" + "foo bar")) + ("hello" + ("hello" + "\n\t"))) {
        case "hello":
            v6 = v6
        case "":
            var v96 int32 = 749
            _ = v96
        default:
            v2 = v12
        }
    }
    v2 -= v2
    v97 := v3[(^(v6))]
    v98 := GenS10{F0: 255}
    v5 = append(v5, "a")
    var v99 int16 = 42
    switch "foo bar" {
    case "hello":
        _ = len(v3)
    case "\n\t":
        v99 += int16(v97)
    default:
        v5 = append(v5, "hello")
    }
    v6 = v6
    v12 = 0
    if true {
        var v100 int8 = (v2 * (v2 ^ v12))
        _ = v100
    }
    v101 := GenS6{F0: "foo bar", F1: (v97 == v97)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v101
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
    _ = v98
    _ = v99
}
