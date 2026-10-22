package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint }
type GenS6 struct { F0 int16; F1 uint32; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 int16; F2 uint32 }
func genFn0(p0 uint, p1 uint) bool {
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int64 = -9223372036854775808
        _ = fi8
        _ = v9
    }
    var v10 uint = (p0 * (p0 / (((p1 | p1)) | 1)))
    if true {
        v7 -= (^(v7))
    }
    if true {
        if (0 < 1) {
            v11 := []int{32768, 9223372036854775807}
            _ = v11
        }
        fmt.Println(v10)
        switch 32767 {
        case 32767:
            switch 4294967295 {
            case -2147483648:
                for fi12 := 0; fi12 < 15; fi12++ {
                    if (42 > 0) {
                        p1 = 628
                        if true {
                            v13 := GenS1{F0: (fi12 + (fi12 | -248)), F1: 2147483648}
                            v14 := GenS1{F0: (fi12 % ((fi12) | 1)), F1: 2147483648}
                            _ = v13
                            _ = v14
                        } else {
                            var v15 float32 = (1.5 + 3.14159)
                            if true {
                                var v16 float32 = v15
                                _ = v16
                            }
                            _ = v15
                        }
                    } else {
                        v10 = 255
                        fmt.Println(fi12, fi12)
                        v17 := []int{fi12}
                        _ = v17
                    }
                    _ = fi12
                }
            case 4294967295:
                fmt.Println(v10)
            }
        case 704:
            if ((^(v7)) > v7) {
                for fi18 := 0; fi18 < 19; fi18++ {
                    if false {
                        fmt.Println(p0, fi18)
                        switch "\n\t" {
                        case "foo bar":
                            v19 := GenS1{F0: (fi18 & 9223372036854775807), F1: 221}
                            _ = v19
                        default:
                            v20 := []int{fi18, (fi18 << ((0) & 63))}
                            _ = v20
                        }
                        p0 += (^(p0))
                    } else {
                        switch "\n\t" {
                        case "hello":
                            var v21 string = "a"
                            _ = v21
                        case "\n\t":
                            for fi22 := 0; fi22 < 10; fi22++ {
                                var v23 string = ("foo bar" + (("\n\t" + "") + "\n\t"))
                                _ = fi22
                                _ = v23
                            }
                        default:
                            v24 := []int{-125, fi18}
                            _ = v24
                        }
                    }
                    _ = fi18
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return false
}
func genFn1() uint64 {
    var v25 uint8 = 0
    v25 = v25
    if false {
        if (127 < 127) {
            v26 := GenS2{F0: -406, F1: 2}
            _ = v26
        }
        fmt.Println(v25)
    }
    v25 = (v25 / ((v25) | 1))
    v27 := []int{32767}
    _ = v25
    _ = v27
    return 180
}

func main() {
    var v1 int64 = -1
    var v2 int64 = 565
    var v3 int16 = 32767
    var v4 int32 = -2147483648
    var v5 int32 = 1
    var v6 int = -1
    v3 -= v3
    v28 := GenS3{F0: "foo bar", F1: ""}
    v29 := GenS3{F0: ("\n\t" + ("a" + "foo bar")), F1: "\n\t"}
    if true {
        v3 += 25
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := []int{(fi30 & (fi30 >> ((65535) & 63)))}
            _ = fi30
            _ = v31
        }
        v1 = v2
        for fi32 := 0; fi32 < 12; fi32++ {
            v3 -= (v3 | -32768)
            v2 = (v2 * v1)
            _ = fi32
        }
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        for fi34 := 0; fi34 < 7; fi34++ {
            v3 = (v3 * (v3 * v3))
            _ = fi34
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            if (v5 < v5) {
                v36 := []int{v6, 4294967295}
                _ = v36
            }
            var v37 string = ((("\n\t" + "foo bar") + "hello") + ("" + (("a" + ("\n\t" + (((("foo bar" + ("foo bar" + ((("" + "\n\t") + (("" + ("a" + (("hello" + "\n\t") + ("" + (("foo bar" + "hello") + ""))))) + "\n\t")) + "a"))) + ("\n\t" + ("a" + "a"))) + ("foo bar" + "a")) + ("\n\t" + "hello")))) + "")))
            _ = fi35
            _ = v37
        }
        _ = fi33
    }
    if (-495 != -1) {
        var v38 float64 = 0.1
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        fmt.Println(v4, v3, v1)
        v40 := []int{}
        for fi41 := 0; fi41 < 8; fi41++ {
            v42 := GenS3{F0: "a", F1: ("" + "foo bar")}
            var v43 float64 = -2.25
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    if false {
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 int16 = (v3 / ((v3) | 1))
            var v46 bool = false
            for fi47 := 0; fi47 < 5; fi47++ {
                switch fi44 {
                case 32767:
                    if v46 {
                        var v48 int16 = v45
                        v49 := GenS4{F0: 68}
                        _ = v48
                        _ = v49
                    } else {
                        var v50 bool = (false || (true || v46))
                        if ((v50 || true) && (652 == 759)) {
                            if ((v2 - (v2 << ((65535) & 63))) == 1) {
                                var v51 string = "a"
                                if (255 >= 240) {
                                    switch (-(fi47)) {
                                    case 128:
                                        switch (v51 + (v51 + v51)) {
                                        case "\n\t":
                                            if v46 {
                                                if (((v5 | v5) != v5) || (234 <= 603)) {
                                                    v52 := GenS5{F0: v45, F1: 100, F2: 18446744073709551615}
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v53 bool = false
                                                        _ = v53
                                                    case "foo bar":
                                                        var v54 int = (fi44 * (^(fi44)))
                                                        _ = v54
                                                    case "hello":
                                                        v5 = (v5 / ((int32(v45)) | 1))
                                                    }
                                                    _ = v52
                                                } else {
                                                    switch (v6 - (v6 ^ v6)) {
                                                    case -2147483648:
                                                        switch v51 {
                                                        case "a":
                                                            var v55 float64 = (3.14159 - 1.5)
                                                            _ = v55
                                                        case "":
                                                            switch (((v51 + v51) + v51) + "\n\t") {
                                                            case "a":
                                                                v51 = v51
                                                            case "":
                                                                var v56 string = "a"
                                                                _ = v56
                                                            default:
                                                                v57 := GenS1{F0: (fi47 / ((int(v5)) | 1)), F1: 2}
                                                                _ = v57
                                                            }
                                                        case "foo bar":
                                                            var v58 uint = 255
                                                            _ = v58
                                                        default:
                                                            v59 := GenS6{F0: 1, F1: 0, F2: -86}
                                                            _ = v59
                                                        }
                                                    case 127:
                                                        if (fi44 >= int(v2)) {
                                                            v51 = ("a" + v51)
                                                            switch -294 {
                                                            case -53:
                                                                v45 = (v45 | v3)
                                                            }
                                                        } else {
                                                            fmt.Println(v4, v4, v46)
                                                            var v60 float32 = 1.5
                                                            v61 := []int{}
                                                            if v46 {
                                                                switch (v51 + "foo bar") {
                                                                case "":
                                                                    v45 += (^(v3))
                                                                case "a":
                                                                    fmt.Println(v50, v60)
                                                                case "foo bar":
                                                                    if (!v46) {
                                                                        v45 = v3
                                                                    } else {
                                                                        v46 = v50
                                                                        if v50 {
                                                                            v61 = append(v61, -422)
                                                                        }
                                                                    }
                                                                }
                                                                v45 -= (-(v3))
                                                            } else {
                                                                if len(v61) >= 2 {
                                                                    v61 = v61[1:len(v61)]
                                                                }
                                                                if len(v61) >= 2 {
                                                                    v61 = v61[1:len(v61)]
                                                                }
                                                            }
                                                            fmt.Println(fi44, v3)
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    default:
                                                        v1 = 1
                                                        var v62 uint8 = 127
                                                        _ = v62
                                                    }
                                                }
                                                var v63 int = fi47
                                                v6 += v63
                                                _ = v63
                                            } else {
                                                var v64 int32 = v4
                                                v65 := GenS1{F0: fi47, F1: 989}
                                                v6 = 1
                                                _ = v64
                                                _ = v65
                                            }
                                        case "":
                                            fmt.Println(v45, v45)
                                        default:
                                            v66 := GenS5{F0: int16(v1), F1: -92, F2: 4294967295}
                                            _ = v66
                                        }
                                    case 0:
                                        switch ("\n\t" + "hello") {
                                        case "\n\t":
                                            v3 = (^(v45))
                                        case "":
                                            v67 := GenS3{F0: v51, F1: ((v51 + v51) + "a")}
                                            _ = v67
                                        case "foo bar":
                                            if (127 <= -128) {
                                                var v68 uint16 = 781
                                                v69 := []int{fi47, 127, fi44}
                                                if true {
                                                    var v70 int32 = (-(v4))
                                                    var v71 uint = 771
                                                    _ = v70
                                                    _ = v71
                                                }
                                                _ = v68
                                                _ = v69
                                            } else {
                                                var v72 uint8 = 0
                                                v73 := GenS7{F0: (v3 - (v45 * (v45 | 32767)))}
                                                var v74 int16 = (v45 / ((954) | 1))
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                            var v75 bool = ((!true) && v46)
                                            v1 = 0
                                            _ = v75
                                        default:
                                            v76 := []int{fi47, (fi47 * int(v1)), fi44}
                                            _ = v76
                                        }
                                    case -1:
                                        fmt.Println(v45)
                                    }
                                    if true {
                                        var v77 uint32 = 1
                                        v50 = (!true)
                                        switch int(v1) {
                                        case -873:
                                            if false {
                                                v6 += 32767
                                            }
                                        case 1:
                                            if (!v50) {
                                                v51 = ""
                                            } else {
                                                var v78 string = v51
                                                if (!((-(v1)) == 42)) {
                                                    v79 := GenS1{F0: fi44, F1: 42}
                                                    _ = v79
                                                }
                                                var v80 float64 = (-2.25 / (0.1 / 1.0))
                                                _ = v78
                                                _ = v80
                                            }
                                        default:
                                            v81 := GenS8{F0: 32767}
                                            _ = v81
                                        }
                                        _ = v77
                                    }
                                    var v82 uint16 = 335
                                    fmt.Println(v46, v5, v45)
                                    v45 += v3
                                    var v83 uint8 = 1
                                    _ = v82
                                    _ = v83
                                } else {
                                    var v84 int64 = (v2 * (v1 * (^(v1))))
                                    fmt.Println(fi44)
                                    _ = v84
                                }
                                var v85 int8 = 0
                                var v86 int8 = (-(v85))
                                _ = v51
                                _ = v85
                                _ = v86
                            }
                            var v87 int16 = (v45 % ((v3) | 1))
                            var v88 uint32 = 0
                            _ = v87
                            _ = v88
                        } else {
                            v89 := []int{(fi47 >> ((763) & 63)), fi47, (fi44 * fi44)}
                            v90 := GenS4{F0: -44}
                            _ = v89
                            _ = v90
                        }
                        v91 := []int{32767, (fi47 % ((-24) | 1)), (fi47 / ((v6) | 1)), fi47}
                        _ = v50
                        _ = v91
                    }
                case 32768:
                    var v92 int32 = -1
                    _ = v92
                case 128:
                    v46 = (458 < fi44)
                default:
                    var v93 int = fi47
                    _ = v93
                }
                _ = fi47
            }
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    for fi94 := 0; fi94 < 5; fi94++ {
        fmt.Println(v3, v2)
        v95 := []int{9223372036854775807}
        _ = fi94
        _ = v95
    }
    switch ("hello" + "\n\t") {
    case "a":
        if ((616 > 42) && true) {
            if true {
                switch "" {
                case "hello":
                    if (((2 > 2) || (v2 != 42)) && true) {
                        var v96 int64 = v1
                        v2 -= (v2 >> ((0) & 63))
                        _ = v96
                    } else {
                        var v97 int8 = -60
                        v97 -= v97
                        var v98 bool = true
                        _ = v97
                        _ = v98
                    }
                default:
                    if true {
                        v99 := GenS9{F0: 2147483647, F1: 42}
                        v3 = 1000
                        v100 := []int{-1}
                        _ = v100
                        _ = v99
                    } else {
                        var v101 int8 = 0
                        v102 := GenS7{F0: v3}
                        for fi103 := 0; fi103 < 9; fi103++ {
                            v104 := GenS5{F0: v3, F1: 127, F2: 2}
                            if (197 > 224) {
                                v105 := []int{fi103, -61, v6}
                                _ = v105
                            } else {
                                v106 := []int{fi103, -9223372036854775808, 32767}
                                v107 := GenS10{F0: v1}
                                var v108 int32 = (v4 ^ v5)
                                _ = v106
                                _ = v107
                                _ = v108
                            }
                            for fi109 := 0; fi109 < 8; fi109++ {
                                if false {
                                    switch -2147483648 {
                                    case -1:
                                        v110 := GenS7{F0: v3}
                                        _ = v110
                                    case 32768:
                                        v111 := GenS11{F0: 0, F1: v3, F2: 845}
                                        _ = v111
                                    }
                                } else {
                                    var v112 int32 = (v4 << ((321) & 63))
                                    _ = v112
                                }
                                v113 := GenS1{F0: v6, F1: 668}
                                var v114 int32 = -31
                                _ = fi109
                                _ = v113
                                _ = v114
                            }
                            _ = fi103
                            _ = v104
                        }
                        _ = v101
                        _ = v102
                    }
                }
            } else {
                var v115 uint = 65535
                v5 = v5
                _ = v115
            }
            var v116 uint16 = 207
            _ = v116
        }
    }
    for fi117 := 0; fi117 < 7; fi117++ {
        fmt.Println(v1)
        _ = fi117
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
