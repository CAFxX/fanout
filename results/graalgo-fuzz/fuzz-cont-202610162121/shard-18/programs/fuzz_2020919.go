package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64; F1 string }
func genFn0(p0 string, p1 uint32) int32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        p1 += 42
        p1 += (p1 >> ((4294967295) & 63))
        p1 += 2147483647
        _ = fi4
    }
    p1 = 4294967295
    p0 = p0
    for fi5 := 0; fi5 < 16; fi5++ {
        if ((true || true) && (42 < 1)) {
            var v6 float64 = (0.1 + 1e10)
            if (-323 != -32768) {
                for fi7 := 0; fi7 < 20; fi7++ {
                    v8 := GenS1{F0: 18446744073709551615}
                    v6 = v6
                    p0 = p0
                    _ = fi7
                    _ = v8
                }
                if (true && (v6 <= (v6 * v6))) {
                    if false {
                        v9 := GenS1{F0: 4294967295}
                        if (v6 < (v6 / v6)) {
                            for fi10 := 0; fi10 < 16; fi10++ {
                                var v11 uint16 = 42
                                v12 := GenS1{F0: 0}
                                v6 -= float64(0.1)
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        } else {
                            v13 := []int{}
                            _ = len(v13)
                            _ = v13
                        }
                        _ = v9
                    }
                    v14 := GenS1{F0: 942}
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v16 := GenS2{F0: 2, F1: 0, F2: p0}
                        var v17 uint8 = 128
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
            }
            _ = v6
        } else {
            fmt.Println(fi5, p1, p0)
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            switch ((p0 + p0) + (p0 + p0)) {
            case "hello":
                fmt.Println(p1, p1, p0)
            case "":
                v19 := GenS3{F0: 0}
                _ = v19
            }
            var v20 uint = 1
            _ = fi18
            _ = v20
        }
        _ = fi5
    }
    switch (("hello" + (("\n\t" + "\n\t") + p0)) + (p0 + p0)) {
    case "a":
        fmt.Println(p1)
    case "hello":
        p0 = p0
    case "":
        switch 32768 {
        case 776:
            if (!(p1 == (p1 >> ((42) & 63)))) {
                p1 -= 772
                v21 := []int{-472}
                _ = v21
            }
        case 259:
            p1 = p1
        case -9223372036854775808:
            var v22 string = ""
            _ = v22
        }
    default:
        v23 := []int{512, 4294967295, 32768, 259}
        _ = v23
    }
    _ = p0
    _ = p1
    return 680
}

func main() {
    var v1 uint64 = 116
    var v2 int64 = -1
    v3 := make(map[int]bool)
    v24 := []int{127, 9223372036854775807}
    switch 127 {
    case -2147483648:
        fmt.Println(v2)
    }
    var v25 uint8 = 1
    var v26 uint64 = 0
    fmt.Println(v2, v26)
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS4{F0: 848}
        switch fi27 {
        case -163:
            v29 := v3[fi27]
            _ = v29
        case -360:
            v25 = v25
        default:
            v26 += v1
        }
        v30 := GenS2{F0: (v1 % ((42) | 1)), F1: (v1 % ((4294967295) | 1)), F2: "\n\t"}
        _ = fi27
        _ = v28
        _ = v30
    }
    if ((3.14159 + 1.5) != 1.5) {
        v2 = v2
        v31 := GenS5{F0: false, F1: 1, F2: 534}
        _ = v31
    }
    var v32 uint32 = 65535
    v33 := GenS1{F0: v26}
    v34 := GenS4{F0: 1000}
    v32 -= (v32 ^ 131)
    for fi35 := 0; fi35 < 14; fi35++ {
        v32 = (v32 ^ v32)
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS6{F0: 2, F1: v25}
            if true {
                switch ("foo bar" + (("foo bar" + "hello") + ("a" + "hello"))) {
                case "foo bar":
                    var v38 uint64 = (v26 + v26)
                    if (true || (((true && (42 == 9223372036854775807)) || (!true)) && false)) {
                        if (1 >= 127) {
                            if (("" < "\n\t") && (v32 != v32)) {
                                v26 = v1
                            } else {
                                var v39 int = -1
                                v26 += 992
                                var v40 uint64 = v1
                                _ = v39
                                _ = v40
                            }
                            v2 = -572
                            if len(v24) > 0 {
                                v24[0] = (fi35 << ((uint(v1)) & 63))
                            }
                            v41 := GenS1{F0: (v38 * 18446744073709551615)}
                            _ = v41
                        } else {
                            var v42 uint32 = (^(v32))
                            var v43 bool = (v38 == (v1 >> ((42) & 63)))
                            if ((1 <= 934) || v43) {
                                v1 = 0
                            }
                            var v44 int64 = (v2 * v2)
                            var v45 string = "hello"
                            var v46 uint16 = uint16(v32)
                            if v43 {
                                var v47 float64 = 3.14159
                                var v48 int8 = 1
                                switch ("hello" + v45) {
                                case "a":
                                    var v49 float64 = v47
                                    _ = v49
                                case "hello":
                                    fmt.Println(v43, v46, v46)
                                default:
                                    var v50 int = fi35
                                    _ = v50
                                }
                                _ = v47
                                _ = v48
                            }
                            _ = v42
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        var v51 uint8 = (v25 << ((42) & 63))
                        _ = v51
                    } else {
                        if false {
                            var v52 int = fi36
                            switch "foo bar" {
                            case "foo bar":
                                switch "a" {
                                case "a":
                                    if false {
                                        var v53 uint = 65535
                                        v54 := v3[(fi36 % (((fi35 << ((4294967295) & 63))) | 1))]
                                        var v55 uint8 = (v25 - (v25 >> ((0) & 63)))
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    switch -2147483648 {
                                    case 32768:
                                        v56 := v3[(v52 + (v52 - (fi36 ^ 4294967295)))]
                                        _ = v56
                                    case -2147483648:
                                        var v57 int16 = 42
                                        _ = v57
                                    default:
                                        switch (v52 << ((255) & 63)) {
                                        case 128:
                                            if (9223372036854775807 < int(v25)) {
                                                v2 -= -1
                                            } else {
                                                v58 := GenS7{F0: (v38 << ((42) & 63))}
                                                if false {
                                                    v52 -= 2147483647
                                                    v59 := v3[(fi35 / ((v52) | 1))]
                                                    fmt.Println(v26, v59)
                                                    _ = v59
                                                }
                                                if ((1e10 == 1.5) && true) {
                                                    var v60 uint = 2
                                                    v1 += v1
                                                    _ = v60
                                                } else {
                                                    if true {
                                                        var v61 int16 = 466
                                                        fmt.Println(v25, fi36)
                                                        _ = v61
                                                    }
                                                }
                                                _ = v58
                                            }
                                        default:
                                            _ = len(v24)
                                        }
                                        if len(v24) >= 2 {
                                            v24 = v24[1:len(v24)]
                                        }
                                    }
                                    var v62 string = (("a" + "") + "")
                                    _ = v62
                                default:
                                    v26 -= (^(v1))
                                }
                            default:
                                v63 := GenS5{F0: (32767 != 255), F1: 42, F2: 2147483648}
                                _ = v63
                            }
                            switch "hello" {
                            case "a":
                                v52 -= (fi35 % ((int(v25)) | 1))
                            case "hello":
                                if (!true) {
                                    if true {
                                        v52 -= (fi36 * int(v32))
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            if (v32 > v32) {
                                                var v64 float64 = 3.14159
                                                var v65 uint32 = (v32 & (^(v32)))
                                                switch "foo bar" {
                                                case "hello":
                                                    v66 := GenS1{F0: 42}
                                                    _ = v66
                                                }
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v2 = (^(v2))
                                                v67 := GenS5{F0: ((v32 / ((v32) | 1)) != v32), F1: 32767, F2: (^(v32))}
                                                var v68 bool = (!((v1 >> ((42) & 63)) > (v26 >> ((1) & 63))))
                                                _ = v67
                                                _ = v68
                                            }
                                        default:
                                            switch "foo bar" {
                                            case "a":
                                                var v69 uint = 1
                                                _ = v69
                                            case "foo bar":
                                                switch "" {
                                                case "hello":
                                                    var v70 uint32 = (v32 << ((876) & 63))
                                                    _ = v70
                                                case "a":
                                                    v71 := v3[(fi36 * -2147483648)]
                                                    _ = v71
                                                case "":
                                                    _ = len(v3)
                                                }
                                            default:
                                                fmt.Println(v1)
                                            }
                                        }
                                        v25 = 255
                                        v24 = append(v24, (v52 + 62))
                                        switch ("\n\t" + ("a" + ("hello" + ""))) {
                                        case "foo bar":
                                            v3[(fi35 / (((^(fi36))) | 1))] = (1e10 < (3.14159 - 1.5))
                                        default:
                                            var v72 uint8 = (v25 ^ v25)
                                            _ = v72
                                        }
                                        v73 := GenS5{F0: true, F1: 1, F2: 654}
                                        if (!((-2.25 + 3.14159) == 3.14159)) {
                                            _ = len(v3)
                                            delete(v3, -1)
                                        }
                                        _ = v73
                                    } else {
                                        var v74 int64 = v2
                                        var v75 bool = false
                                        if (-59 >= -19) {
                                            fmt.Println(v52)
                                            if ((626 <= v32) && ((!(!v75)) && false)) {
                                                v74 -= v2
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v76 uint64 = v26
                                                v75 = (v75 || false)
                                                v77 := GenS8{F0: ((((-2.25 * 1e10) + 1e10) * 1e10) == 0.1), F1: 0}
                                                _ = v76
                                                _ = v77
                                            case "hello":
                                                v52 += (-(v52))
                                            }
                                            var v78 uint = 2
                                            v79 := GenS4{F0: 1}
                                            v32 -= (v32 - 0)
                                            _ = v78
                                            _ = v79
                                        }
                                        _ = v74
                                        _ = v75
                                    }
                                    v3[200] = true
                                } else {
                                    if false {
                                        switch "hello" {
                                        case "hello":
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        }
                                        var v80 float32 = (float32((1.5 / 1.5)) - -2.25)
                                        _ = v80
                                    }
                                    v32 = (v32 << ((42) & 63))
                                    switch -1 {
                                    case 9223372036854775807:
                                        if len(v24) >= 2 {
                                            v24 = v24[1:len(v24)]
                                        }
                                    case 2147483647:
                                        switch (fi36 >> ((65535) & 63)) {
                                        case 95:
                                            v1 += (^(v1))
                                        case 0:
                                            v81 := GenS9{F0: v2}
                                            _ = v81
                                        }
                                    case 42:
                                        v25 -= (v25 | v25)
                                    default:
                                        var v82 int64 = (v2 >> ((271) & 63))
                                        _ = v82
                                    }
                                    var v83 uint16 = 42
                                    _ = v83
                                }
                            case "\n\t":
                                _ = len(v3)
                            }
                            _ = v52
                        }
                        var v84 int = fi35
                        var v85 uint64 = (v38 & v1)
                        v38 = v85
                        _ = v84
                        _ = v85
                    }
                    var v86 uint64 = (v38 ^ (v38 & (v26 >> ((18446744073709551615) & 63))))
                    if true {
                        fmt.Println(v1, fi36, fi35)
                        if (fi36 < -1) {
                            if true {
                                fmt.Println(v32)
                                v87 := GenS10{F0: (v2 % (((v2 % (((v2 & (v2 % (((v2 / ((int64(fi36)) | 1))) | 1)))) | 1))) | 1)), F1: "\n\t"}
                                v32 = v32
                                v38 = uint64(v32)
                                _ = v87
                            } else {
                                v26 = v1
                            }
                        } else {
                            v88 := GenS10{F0: (v2 | int64(fi35)), F1: "a"}
                            _ = v88
                        }
                        switch fi35 {
                        case 783:
                            var v89 int16 = 0
                            _ = v89
                        case -445:
                            v38 = v1
                        default:
                            if (true || (!(505 <= 0))) {
                                if true {
                                    switch 128 {
                                    case -822:
                                        var v90 int = (fi36 / ((fi36) | 1))
                                        _ = v90
                                    default:
                                        if false {
                                            var v91 bool = (!false)
                                            _ = v91
                                        }
                                    }
                                    v92 := GenS2{F0: 993, F1: (v1 >> ((744) & 63)), F2: "\n\t"}
                                    _ = v92
                                } else {
                                    switch (("foo bar" + "\n\t") + ("a" + "hello")) {
                                    case "a":
                                        if false {
                                            v25 -= v25
                                            _ = len(v3)
                                            _ = len(v24)
                                        }
                                    case "\n\t":
                                        if (1 > 255) {
                                            var v93 float64 = -2.25
                                            switch "a" {
                                            case "a":
                                                v93 = v93
                                            }
                                            _ = v93
                                        }
                                    default:
                                        fmt.Println(v1, v2)
                                    }
                                    fmt.Println(v26)
                                }
                                var v94 bool = (2 > 0)
                                var v95 int32 = -2147483648
                                _ = v94
                                _ = v95
                            }
                        }
                    } else {
                        switch ((("\n\t" + "") + "") + "\n\t") {
                        case "hello":
                            v26 = v1
                        default:
                            if (!(true || (v2 > 42))) {
                                fmt.Println(v32, fi36, v2)
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                            } else {
                                var v96 int32 = -2147483648
                                var v97 bool = ((!(v25 >= (v25 & 127))) || (0 >= 42))
                                v98 := GenS6{F0: 631, F1: v25}
                                var v99 float32 = -2.25
                                v26 -= 1
                                var v100 int16 = -563
                                _ = v100
                                _ = v96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                        }
                        var v101 int = 2147483647
                        var v102 int64 = v2
                        _ = v101
                        _ = v102
                    }
                    v103 := GenS2{F0: v26, F1: v38, F2: ("a" + "")}
                    _ = v103
                    _ = v38
                    _ = v86
                case "a":
                    if (true && true) {
                        v3[fi35] = false
                        var v104 bool = false
                        var v105 int8 = -40
                        _ = v104
                        _ = v105
                    }
                case "":
                    var v106 uint16 = 0
                    _ = v106
                }
            }
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
