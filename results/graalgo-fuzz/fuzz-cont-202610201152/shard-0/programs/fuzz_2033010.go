package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint16; F2 int }
type GenS8 struct { F0 int8; F1 uint16; F2 int64 }
type GenS9 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 int16 }
func genFn0(p0 int64) int64 {
    v6 := GenS1{F0: 2147483647, F1: 9223372036854775807}
    var v7 int16 = -32768
    p0 -= p0
    v8 := GenS2{F0: 0, F1: (p0 ^ (p0 / ((p0) | 1)))}
    for fi9 := 0; fi9 < 15; fi9++ {
        v7 = (v7 - 1000)
        _ = fi9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 int64 = -808
    var v5 int8 = -128
    if (!(42 < -32768)) {
        switch -1 {
        case 2147483647:
            var v10 int = 648
            _ = v10
        }
    } else {
        fmt.Println(v1)
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                switch "\n\t" {
                case "a":
                    var v13 int32 = 832
                    _ = v13
                }
                var v14 int8 = -128
                fmt.Println(v5, v5, v1)
                var v15 uint32 = 42
                v5 = v5
                _ = fi12
                _ = v14
                _ = v15
            }
            _ = fi11
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            v5 -= (v5 * (^(v5)))
        }
        var v17 int = (fi16 % ((-9223372036854775808) | 1))
        v18 := GenS2{F0: 715, F1: v4}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (v3 && true) {
        v19 := GenS2{F0: 0, F1: (v4 - (v4 - 1))}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi20)
            case "hello":
                var v21 int = (fi20 - 4294967295)
                _ = v21
            }
            switch "foo bar" {
            case "a":
                delete(v2, 2147483647)
            default:
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 uint = v1
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi20
        }
        v2[2] = 4294967295
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 float32 = 1.5
            _ = fi24
            _ = v25
        }
    }
    var v26 uint32 = 2147483647
    for fi27 := 0; fi27 < 12; fi27++ {
        v1 -= v1
        v28 := GenS3{F0: v26}
        _ = fi27
        _ = v28
    }
    if v3 {
        v29 := GenS2{F0: 18446744073709551615, F1: v4}
        v30 := GenS4{F0: false, F1: v26}
        v4 += -1
        _ = v29
        _ = v30
    } else {
        v31 := []int{791, int(v1)}
        _ = v31
    }
    switch ("\n\t" + ("foo bar" + ("a" + ""))) {
    case "foo bar":
        var v32 string = ("a" + "\n\t")
        _ = v32
    case "":
        switch ("hello" + "hello") {
        case "\n\t":
            if false {
                var v33 int = 32767
                _ = v33
            }
        case "hello":
            if (v3 || v3) {
                v34 := GenS1{F0: -9223372036854775808, F1: 288}
                var v35 float32 = float32(-2.25)
                _ = v34
                _ = v35
            }
        case "":
            var v36 uint16 = 32767
            _ = v36
        default:
            switch ("hello" + "\n\t") {
            case "a":
                var v37 float32 = -2.25
                _ = v37
            case "foo bar":
                if true {
                    v26 = (^(v26))
                    if v3 {
                        for fi38 := 0; fi38 < 13; fi38++ {
                            if ((v5 * (v5 + -128)) < v5) {
                                var v39 uint = (v1 * v1)
                                var v40 int32 = 42
                                v2[v26] = v26
                                _ = v39
                                _ = v40
                            }
                            if v3 {
                                v41 := []int{0, (fi38 & fi38)}
                                fmt.Println(v4, v5)
                                _ = v41
                            } else {
                                var v42 string = ""
                                v43 := GenS1{F0: fi38, F1: fi38}
                                _ = v42
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 19; fi44++ {
                                _ = len(v2)
                                fmt.Println(v26, v4, v3)
                                if v3 {
                                    if true {
                                        v3 = (v3 || (168 != 127))
                                        var v45 uint16 = 65535
                                        _ = v45
                                    } else {
                                        var v46 int8 = (v5 >> ((18446744073709551615) & 63))
                                        switch (fi38 + fi38) {
                                        case 127:
                                            v46 -= 127
                                        case 1:
                                            switch "a" {
                                            case "a":
                                                v47 := GenS4{F0: v3, F1: uint32(v1)}
                                                _ = v47
                                            case "foo bar":
                                                switch "" {
                                                case "a":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v48 := GenS5{F0: 508}
                                                        _ = v48
                                                    case "a":
                                                        v49 := v2[(v26 & (v26 - v26))]
                                                        _ = v49
                                                    }
                                                case "hello":
                                                    delete(v2, 2147483648)
                                                }
                                            default:
                                                var v50 int = 32768
                                                _ = v50
                                            }
                                            switch int(v26) {
                                            case 9223372036854775807:
                                                v51 := []int{128, fi38}
                                                _ = v51
                                            case 683:
                                                v52 := []int{}
                                                _ = v52
                                            }
                                        }
                                        var v53 int8 = v46
                                        v54 := GenS2{F0: 18446744073709551615, F1: v4}
                                        delete(v2, v26)
                                        _ = v46
                                        _ = v53
                                        _ = v54
                                    }
                                    v55 := v2[v26]
                                    v56 := GenS6{F0: (v55 * (v55 | v26))}
                                    v57 := GenS7{F0: 0, F1: 2, F2: fi38}
                                    var v58 float64 = float64(1e10)
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                } else {
                                    if (v3 || false) {
                                        var v59 uint = v1
                                        fmt.Println(v5, fi44)
                                        var v60 uint32 = (v26 / (((v26 + v26)) | 1))
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                _ = fi44
                            }
                            _ = fi38
                        }
                        var v61 string = "\n\t"
                        _ = v61
                    }
                } else {
                    var v62 bool = (v3 || (627 >= 2))
                    var v63 bool = v3
                    if (!false) {
                        switch -447 {
                        case 32767:
                            fmt.Println(v26, v3)
                        case 128:
                            var v64 float64 = (0.0 / float64(((((float32(3.14159) + 1.5) / 0.1) + -2.25) * ((-2.25 + 0.1) / -2.25))))
                            _ = v64
                        case -1:
                            for fi65 := 0; fi65 < 20; fi65++ {
                                v66 := []int{(fi65 % (((fi65 % (((^(fi65))) | 1))) | 1)), 42, 9223372036854775807, (fi65 & fi65)}
                                if (!v62) {
                                    var v67 int16 = 788
                                    v66 = append(v66, (fi65 / ((fi65) | 1)))
                                    _ = v67
                                }
                                _ = fi65
                                _ = v66
                            }
                        default:
                            var v68 uint16 = 322
                            _ = v68
                        }
                        if (v63 || v63) {
                            v4 = (^(v4))
                            var v69 int32 = 2147483647
                            var v70 int16 = 944
                            _ = v69
                            _ = v70
                        }
                    }
                    _ = v62
                    _ = v63
                }
            }
        }
    case "\n\t":
        v71 := GenS4{F0: (-2.25 > (-2.25 + 1e10)), F1: v26}
        _ = v71
    }
    for fi72 := 0; fi72 < 6; fi72++ {
        v3 = (-110 != 2147483647)
        v26 -= (v26 & (^(v26)))
        _ = fi72
    }
    v4 -= (v4 * genFn0(v4))
    v26 = (v26 | (v26 | (v26 << (((v1 - 756)) & 63))))
    for fi73 := 0; fi73 < 12; fi73++ {
        v5 -= v5
        switch fi73 {
        case -2147483648:
            var v74 uint = v1
            _ = v74
        case 603:
            if v3 {
                v75 := GenS2{F0: 0, F1: genFn0((v4 * int64(v26)))}
                _ = v75
            } else {
                v1 = v1
                if (v5 <= v5) {
                    switch -467 {
                    case 128:
                        var v76 uint = 0
                        _ = v76
                    case 32768:
                        if true {
                            delete(v2, uint32(v1))
                            v77 := GenS8{F0: (^(v5)), F1: 32767, F2: v4}
                            switch ("foo bar" + ("a" + "\n\t")) {
                            case "a":
                                if true {
                                    for fi78 := 0; fi78 < 7; fi78++ {
                                        v79 := v2[(v26 - v26)]
                                        var v80 float32 = 3.14159
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                }
                            default:
                                v2[(v26 << ((4294967295) & 63))] = (^(v26))
                            }
                            _ = v77
                        } else {
                            v3 = v3
                            v81 := GenS6{F0: 99}
                            _ = v81
                        }
                    default:
                        for fi82 := 0; fi82 < 20; fi82++ {
                            v83 := GenS9{F0: v5, F1: 4294967295, F2: 18446744073709551615}
                            if (int64(v26) < v4) {
                                v84 := GenS8{F0: v5, F1: uint16(v1), F2: (v4 + 42)}
                                var v85 int16 = 0
                                v1 -= 187
                                _ = v84
                                _ = v85
                            } else {
                                switch ("a" + "") {
                                case "":
                                    v3 = (!false)
                                case "a":
                                    v86 := GenS3{F0: (^(v26))}
                                    _ = v86
                                case "hello":
                                    v3 = (v5 <= 127)
                                default:
                                    v2[v26] = (v26 << ((364) & 63))
                                }
                                v5 -= v5
                                v87 := GenS10{F0: (v5 | (v5 / ((v5) | 1)))}
                                fmt.Println(v26, fi82, fi82)
                                _ = v87
                            }
                            v5 = 100
                            var v88 uint64 = 4294967295
                            fmt.Println(v5)
                            if v3 {
                                v5 = (v5 - 74)
                                var v89 uint32 = 42
                                _ = v89
                            }
                            if v3 {
                                if ((true || v3) || (v88 <= (^(v88)))) {
                                    var v90 string = (((("\n\t" + ("" + ("\n\t" + (("foo bar" + ("\n\t" + "a")) + "\n\t")))) + ("hello" + "a")) + "foo bar") + "hello")
                                    _ = v90
                                } else {
                                    var v91 uint16 = 255
                                    _ = v91
                                }
                                if (false && (!(1e10 == 3.14159))) {
                                    v88 -= (v88 >> ((v1) & 63))
                                } else {
                                    var v92 int16 = 0
                                    var v93 string = ((((("" + "a") + ("a" + "a")) + "\n\t") + ((("\n\t" + "hello") + "hello") + ((("" + "\n\t") + "foo bar") + "a"))) + ("" + (("\n\t" + "\n\t") + ("\n\t" + ("hello" + "hello")))))
                                    delete(v2, (^(v26)))
                                    _ = v92
                                    _ = v93
                                }
                                v5 = (^(v5))
                            }
                            _ = fi82
                            _ = v83
                            _ = v88
                        }
                    }
                    v94 := []int{fi73, (fi73 - fi73), fi73}
                    _ = v94
                } else {
                    fmt.Println(v4)
                    for fi95 := 0; fi95 < 19; fi95++ {
                        v5 -= v5
                        switch ((("foo bar" + "") + "") + "") {
                        case "a":
                            v4 -= 727
                        }
                        v96 := GenS7{F0: -1, F1: 65535, F2: fi73}
                        _ = fi95
                        _ = v96
                    }
                    v4 -= genFn0(genFn0(1))
                }
            }
        case 128:
            switch ("\n\t" + "") {
            case "hello":
                v97 := GenS2{F0: 4294967295, F1: v4}
                _ = v97
            case "foo bar":
                var v98 uint64 = 174
                _ = v98
            case "\n\t":
                if ("" == "foo bar") {
                    v99 := GenS11{F0: true, F1: -250}
                    _ = v99
                }
            }
        default:
            v1 = (v1 % (((v1 % (((v1 & (v1 << ((v1) & 63)))) | 1))) | 1))
        }
        if v3 {
            var v100 uint64 = 167
            v101 := []int{749, (fi73 / ((128) | 1))}
            _ = len(v101)
            _ = v100
            _ = v101
        }
        _ = fi73
    }
    for fi102 := 0; fi102 < 18; fi102++ {
        v1 -= (^(v1))
        _ = fi102
    }
    var v103 int64 = (v4 - (v4 | (v4 >> ((18446744073709551615) & 63))))
    for fi104 := 0; fi104 < 19; fi104++ {
        v105 := []int{(fi104 - 9223372036854775807), 9223372036854775807, 798}
        fmt.Println(v5, v5)
        fmt.Println(v26, v4, v5)
        _ = fi104
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
