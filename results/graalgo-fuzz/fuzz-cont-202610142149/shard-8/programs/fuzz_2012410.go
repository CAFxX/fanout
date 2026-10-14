package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS6 struct { F0 string; F1 uint16; F2 string }
type GenS7 struct { F0 bool; F1 string; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 uint64 }
func genFn0() string {
    switch "foo bar" {
    case "foo bar":
        var v5 int32 = 42
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int8 = 66
        var v8 uint = 323
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if ((fi9 * -9223372036854775808) >= (fi9 ^ fi9)) {
            var v10 int16 = 0
            v10 = v10
            _ = v10
        } else {
            v11 := GenS1{F0: 18446744073709551615}
            var v12 int8 = 42
            var v13 float32 = ((1.5 * 0.0) - 0.1)
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: (("hello" + "\n\t") + "hello")}
            var v16 uint16 = 732
            var v17 float64 = -2.25
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 uint = 331
        _ = fi9
        _ = v18
    }
    switch 1 {
    case 0:
        switch ("\n\t" + "a") {
        case "\n\t":
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS2{F0: "foo bar"}
                _ = fi19
                _ = v20
            }
        case "hello":
            for fi21 := 0; fi21 < 8; fi21++ {
                if (178 >= 2) {
                    fmt.Println(fi21)
                } else {
                    var v22 float64 = 1.5
                    _ = v22
                }
                v23 := GenS3{F0: 1000, F1: 1}
                var v24 uint16 = 1
                _ = fi21
                _ = v23
                _ = v24
            }
        case "foo bar":
            for fi25 := 0; fi25 < 18; fi25++ {
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 int16 = 0
                    v27 = (v27 | v27)
                    v27 = 1
                    _ = fi26
                    _ = v27
                }
                fmt.Println(fi25)
                var v28 uint = 18446744073709551615
                _ = fi25
                _ = v28
            }
        default:
            if false {
                var v29 uint8 = 0
                for fi30 := 0; fi30 < 8; fi30++ {
                    v29 = (v29 >> ((42) & 63))
                    v29 += v29
                    _ = fi30
                }
                _ = v29
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 bool = true
                    switch 127 {
                    case 512:
                        var v33 float64 = -2.25
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
            }
        }
    default:
        var v34 float64 = (1e10 / float64(3.14159))
        _ = v34
    }
    v35 := GenS3{F0: 42, F1: 1}
    _ = v35
    return ("a" + ("hello" + "a"))
}
func genFn1(p0 uint32, p1 int8, p2 int8) int {
    p0 = (^(p0))
    var v36 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    return int(p2)
}

func main() {
    var v1 int8 = -104
    var v2 uint16 = 32767
    var v3 int16 = -847
    var v4 int = 586
    switch 1 {
    case -88:
        if ((!false) && false) {
            if false {
                v37 := []int{}
                v3 += v3
                if true {
                    var v38 int64 = 1
                    switch 642 {
                    case -245:
                        var v39 uint = 18446744073709551615
                        _ = v39
                    case 2147483647:
                        for fi40 := 0; fi40 < 2; fi40++ {
                            if ((v1 >> ((255) & 63)) == (v1 / ((1) | 1))) {
                                v38 = (v38 / (((v38 + (v38 * v38))) | 1))
                                for fi41 := 0; fi41 < 15; fi41++ {
                                    v2 += (v2 | v2)
                                    switch 924 {
                                    case 1:
                                        if (!((0 == 18446744073709551615) || (346 != -32768))) {
                                            v4 += -2147483648
                                            if len(v37) >= 2 {
                                                v37 = v37[1:len(v37)]
                                            }
                                            v3 = (v3 >> ((42) & 63))
                                            switch 2147483647 {
                                            case 2147483647:
                                                var v42 uint32 = 2
                                                _ = v42
                                            case -2147483648:
                                                v43 := GenS4{F0: fi40}
                                                _ = v43
                                            case 4294967295:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v44 int64 = v38
                                                    _ = v44
                                                }
                                            }
                                            if (65535 < 2147483648) {
                                                if ((-(v3)) < -1) {
                                                    v45 := GenS5{F0: (v1 + 100), F1: 4294967295, F2: 42}
                                                    _ = v45
                                                } else {
                                                    if len(v37) > 0 {
                                                        v37[0] = genFn1(684, v1, v1)
                                                    }
                                                    if (-2147483648 <= 1) {
                                                        fmt.Println(fi40, v2)
                                                        v1 -= -128
                                                    }
                                                    var v46 uint = 663
                                                    _ = v46
                                                }
                                                var v47 int16 = v3
                                                _ = v47
                                            } else {
                                                v3 += (v3 << ((18446744073709551615) & 63))
                                            }
                                        }
                                    default:
                                        if len(v37) > 0 {
                                            v37[0] = fi40
                                        }
                                    }
                                    _ = fi41
                                }
                            } else {
                                v3 = 1
                                var v48 int32 = 0
                                _ = v48
                            }
                            v49 := GenS2{F0: ""}
                            _ = fi40
                            _ = v49
                        }
                    case 128:
                        v50 := GenS2{F0: ("a" + ((("" + "hello") + "foo bar") + "a"))}
                        _ = v50
                    }
                    v37 = append(v37, v4)
                    _ = v38
                }
                _ = v37
            }
            if true {
                v2 = (v2 >> ((489) & 63))
            } else {
                v51 := GenS5{F0: (v1 | (v1 * (v1 * v1))), F1: 28, F2: 491}
                _ = v51
            }
        }
    case 42:
        v1 += (v1 & v1)
    case -2147483648:
        v1 += (-(v1))
    default:
        v3 += -893
    }
    v2 = v2
    var v52 string = ("a" + "a")
    if (9223372036854775807 <= 826) {
        if (float64(float32(1e10)) > ((-2.25 - 0.1) + (-2.25 + (0.1 / (1e10 - float64(0.0)))))) {
            switch v52 {
            case "foo bar":
                fmt.Println(v3, v1, v52)
            case "":
                for fi53 := 0; fi53 < 1; fi53++ {
                    v54 := GenS2{F0: v52}
                    _ = fi53
                    _ = v54
                }
            default:
                fmt.Println(v1, v52, v52)
            }
            switch v52 {
            case "":
                v55 := []int{9223372036854775807, genFn1(2, (v1 - v1), 0)}
                _ = v55
            }
            for fi56 := 0; fi56 < 3; fi56++ {
                for fi57 := 0; fi57 < 11; fi57++ {
                    var v58 string = v52
                    _ = fi57
                    _ = v58
                }
                if (v1 > v1) {
                    v52 = ("foo bar" + v52)
                    for fi59 := 0; fi59 < 2; fi59++ {
                        var v60 uint16 = 534
                        fmt.Println(fi56, v52, v3)
                        if (0 <= -217) {
                            v60 -= (v2 + v60)
                            if false {
                                v60 -= (v60 | v60)
                                v61 := GenS5{F0: v1, F1: 2, F2: 2147483647}
                                v62 := GenS6{F0: v52, F1: v2, F2: (v52 + "hello")}
                                _ = v61
                                _ = v62
                            }
                        } else {
                            var v63 int64 = int64(v60)
                            v64 := GenS2{F0: (v52 + "\n\t")}
                            v65 := GenS1{F0: 179}
                            v4 += fi59
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = fi59
                        _ = v60
                    }
                    var v66 float32 = -2.25
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 7; fi67++ {
                        v52 = "a"
                        _ = fi67
                    }
                    switch -9223372036854775808 {
                    case 1:
                        var v68 uint = 42
                        _ = v68
                    case -1:
                        switch ((v52 + "foo bar") + v52) {
                        case "":
                            v69 := GenS7{F0: (true && ((-361 != v3) && (true && true))), F1: v52, F2: 1}
                            _ = v69
                        }
                    }
                }
                v70 := GenS3{F0: (v3 % ((v3) | 1)), F1: 0}
                _ = fi56
                _ = v70
            }
        } else {
            fmt.Println(v4, v2, v3)
            if false {
                for fi71 := 0; fi71 < 8; fi71++ {
                    v72 := []int{}
                    v1 += v1
                    switch (^(fi71)) {
                    case -345:
                        if len(v72) > 0 {
                            v72[0] = 849
                        }
                    case -3:
                        v4 = (^(v4))
                    default:
                        for fi73 := 0; fi73 < 18; fi73++ {
                            if (true && false) {
                                v52 = v52
                                var v74 uint8 = 1
                                _ = v74
                            }
                            if (true && false) {
                                var v75 uint8 = 128
                                var v76 int64 = 9223372036854775807
                                var v77 int8 = v1
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                            var v78 string = "foo bar"
                            _ = fi73
                            _ = v78
                        }
                    }
                    _ = fi71
                    _ = v72
                }
            } else {
                var v79 int64 = 0
                v80 := GenS8{F0: (v79 / ((v79) | 1)), F1: 1}
                var v81 string = v52
                _ = v79
                _ = v80
                _ = v81
            }
            v1 = v1
        }
    }
    fmt.Println(v3)
    var v82 uint16 = 32767
    if true {
        if (((v52 + v52) != ((v52 + v52) + ("hello" + ("" + v52)))) || true) {
            if (-9223372036854775808 > 1) {
                for fi83 := 0; fi83 < 19; fi83++ {
                    fmt.Println(v82, fi83, v3)
                    v84 := GenS1{F0: 314}
                    v52 = v52
                    _ = fi83
                    _ = v84
                }
            }
            v85 := GenS7{F0: false, F1: ((v52 + (("a" + v52) + "")) + ((v52 + v52) + "\n\t")), F2: 4294967295}
            _ = v85
        }
        v86 := []int{v4, (v4 % ((v4) | 1)), v4}
        _ = v86
    } else {
        switch (v52 + v52) {
        case "foo bar":
            var v87 uint = 18446744073709551615
            _ = v87
        }
        v2 -= v2
        if (true && false) {
            for fi88 := 0; fi88 < 20; fi88++ {
                v89 := GenS5{F0: (v1 & (v1 >> ((0) & 63))), F1: 2, F2: 564}
                var v90 string = ("a" + (v52 + v52))
                v91 := GenS3{F0: v3, F1: 2147483647}
                _ = fi88
                _ = v89
                _ = v90
                _ = v91
            }
        }
    }
    v92 := []int{2147483647, v4}
    for fi93 := 0; fi93 < 6; fi93++ {
        _ = len(v92)
        _ = fi93
    }
    switch v52 {
    case "a":
        v94 := GenS4{F0: (v4 / ((9223372036854775807) | 1))}
        _ = v94
    case "":
        fmt.Println(v1, v4, v3)
    case "foo bar":
        if len(v92) > 0 {
            v92[0] = (v4 | -45)
        }
    }
    var v95 int = v4
    v1 += (v1 - (v1 * (v1 % ((1) | 1))))
    switch v52 {
    case "a":
        if (0 <= 0) {
            if (367 < 4294967295) {
                fmt.Println(v1)
                if true {
                    for fi96 := 0; fi96 < 11; fi96++ {
                        v4 = (-(fi96))
                        _ = fi96
                    }
                } else {
                    v97 := GenS9{F0: 255, F1: 0}
                    var v98 int8 = (v1 & 0)
                    if false {
                        fmt.Println(v4, v82)
                    } else {
                        var v99 int32 = -141
                        switch (v95 << ((65535) & 63)) {
                        case -1:
                            var v100 int = v4
                            _ = v100
                        }
                        if false {
                            var v101 uint32 = 2147483648
                            for fi102 := 0; fi102 < 11; fi102++ {
                                v95 += v95
                                var v103 uint8 = 0
                                _ = fi102
                                _ = v103
                            }
                            _ = v101
                        }
                        _ = v99
                    }
                    _ = v97
                    _ = v98
                }
                if false {
                    v104 := GenS10{F0: (v3 + (v3 + (v3 / ((v3) | 1)))), F1: -1}
                    var v105 uint32 = 0
                    _ = v104
                    _ = v105
                } else {
                    for fi106 := 0; fi106 < 9; fi106++ {
                        for fi107 := 0; fi107 < 6; fi107++ {
                            var v108 int = -312
                            _ = fi107
                            _ = v108
                        }
                        _ = fi106
                    }
                    v4 -= v4
                }
            } else {
                switch "a" {
                case "foo bar":
                    switch "" {
                    case "foo bar":
                        if len(v92) >= 2 {
                            v92 = v92[1:len(v92)]
                        }
                    case "":
                        fmt.Println(v3)
                    case "a":
                        if len(v92) >= 2 {
                            v92 = v92[1:len(v92)]
                        }
                    }
                }
                for fi109 := 0; fi109 < 17; fi109++ {
                    fmt.Println(v4, v2, v95)
                    var v110 uint64 = 919
                    v111 := GenS7{F0: true, F1: ("\n\t" + v52), F2: (^(v110))}
                    _ = fi109
                    _ = v110
                    _ = v111
                }
                v95 -= v4
            }
        } else {
            var v112 float64 = (-2.25 * 1e10)
            _ = v112
        }
    case "foo bar":
        v92 = append(v92, (v95 + v95))
    }
    if len(v92) > 0 {
        v92[0] = v4
    }
    v113 := GenS11{F0: 2}
    fmt.Println(v4, v3)
    v114 := GenS8{F0: 1, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v82:", v82)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v82
    _ = v92
    _ = v95
}
