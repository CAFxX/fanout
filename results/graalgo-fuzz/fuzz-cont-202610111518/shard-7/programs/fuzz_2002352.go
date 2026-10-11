package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 bool; F1 int8; F2 int64 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8; F1 int64 }
type GenS13 struct { F0 int; F1 int64 }
type GenS14 struct { F0 int32; F1 string; F2 int8 }
type GenS15 struct { F0 int64; F1 int64 }
type GenS16 struct { F0 int8; F1 int32; F2 int16 }
type GenS17 struct { F0 string; F1 uint16; F2 string }
type GenS18 struct { F0 int; F1 int32 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int8; F1 int; F2 int }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        for fi7 := 0; fi7 < 8; fi7++ {
            if (!true) {
                var v8 int32 = -544
                _ = v8
            } else {
                v9 := []int{(fi6 | (fi7 << ((856) & 63))), fi6, fi6}
                var v10 uint64 = 42
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        v11 := GenS1{F0: 2, F1: 132, F2: "a"}
        _ = fi6
        _ = v11
    }
    v12 := []int{847, 749, -854}
    var v13 int = -9223372036854775808
    var v14 int16 = -622
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 uint32 = 2147483647
    var v5 uint8 = 7
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint64 = 1
        switch "\n\t" {
        case "\n\t":
            v17 := GenS1{F0: 2, F1: v5, F2: "hello"}
            _ = v17
        }
        if false {
            if (1 >= 1) {
                v18 := GenS2{F0: true}
                for fi19 := 0; fi19 < 5; fi19++ {
                    _ = len(v2)
                    v1 -= v1
                    _ = fi19
                }
                _ = v18
            }
            switch "a" {
            case "hello":
                for fi20 := 0; fi20 < 8; fi20++ {
                    v4 = (v4 * 2147483647)
                    switch "a" {
                    case "foo bar":
                        v5 -= v5
                    }
                    var v21 string = (("\n\t" + "foo bar") + "hello")
                    fmt.Println(fi20, v1)
                    _ = fi20
                    _ = v21
                }
            case "a":
                var v22 int16 = 0
                _ = v22
            default:
                v5 = (v5 % ((v5) | 1))
            }
            fmt.Println(v4)
        }
        _ = fi15
        _ = v16
    }
    v23 := GenS3{F0: true, F1: 1, F2: -128}
    var v24 uint = 1
    v2[v4] = false
    var v25 float32 = (-2.25 / 1.0)
    if (!true) {
        if len(v3) > 0 {
            v3[0] = -807
        }
        var v26 float64 = 0.1
        if false {
            var v27 uint16 = 2
            _ = v27
        } else {
            v4 -= (v4 - v4)
            switch "\n\t" {
            case "\n\t":
                for fi28 := 0; fi28 < 2; fi28++ {
                    v2[65535] = (-1 < (^(v1)))
                    if false {
                        for fi29 := 0; fi29 < 20; fi29++ {
                            var v30 int64 = 443
                            fmt.Println(fi29, v25, v26)
                            var v31 int32 = int32(fi29)
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 14; fi32++ {
                            v33 := GenS2{F0: (!(((v25 - (v25 / -2.25)) - float32(-2.25)) <= v25))}
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            switch 9223372036854775807 {
                            case 0:
                                if false {
                                    v35 := GenS4{F0: 42}
                                    v24 -= uint(v4)
                                    _ = v35
                                }
                            case 1:
                                fmt.Println(v5)
                            default:
                                switch fi28 {
                                case 465:
                                    if (!false) {
                                        if true {
                                            v3 = append(v3, (v1 - (fi34 + (^(fi34)))))
                                            v25 = v25
                                        } else {
                                            if (!(!(-326 < 42))) {
                                                if ((127 <= 1) && true) {
                                                    v36 := GenS5{F0: (v4 * (v4 * 0)), F1: 42}
                                                    _ = v36
                                                } else {
                                                    var v37 uint8 = (v5 / (((v5 % ((42) | 1))) | 1))
                                                    var v38 float32 = (1.5 + v25)
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v5 = (v5 | v5)
                                                if ((v26 / (v26 - v26)) > -2.25) {
                                                    v39 := GenS4{F0: -1}
                                                    v40 := GenS5{F0: v4, F1: -128}
                                                    _ = v39
                                                    _ = v40
                                                }
                                                fmt.Println(v25, v25, v24)
                                            }
                                            var v41 uint32 = v4
                                            fmt.Println(v24, v4)
                                            _ = v41
                                        }
                                        v42 := v2[(v4 ^ 777)]
                                        switch fi28 {
                                        case 503:
                                            v4 += 2
                                        case 0:
                                            v2[v4] = (((v42 && (v42 || (v25 != 1.5))) && v42) || ((fi34 >= 0) && v42))
                                        case 187:
                                            v25 -= 0.0
                                        }
                                        var v43 int32 = 2147483647
                                        if ((v24 >> (((v24 * v24)) & 63)) >= 2) {
                                            switch "a" {
                                            case "":
                                                v42 = v42
                                            case "\n\t":
                                                var v44 uint64 = genFn0()
                                                switch fi34 {
                                                case 38:
                                                    if ((91 >= 60) && (false || true)) {
                                                        fmt.Println(v26, v43, v1)
                                                        v45 := GenS1{F0: v24, F1: (v5 >> (((v24 + v24)) & 63)), F2: "foo bar"}
                                                        v24 = (v24 | v24)
                                                        _ = v45
                                                    }
                                                default:
                                                    fmt.Println(v26, v4, v25)
                                                }
                                                _ = v44
                                            }
                                            if true {
                                                v46 := GenS2{F0: (!false)}
                                                delete(v2, v4)
                                                _ = v46
                                            } else {
                                                v43 = int32(fi28)
                                                var v47 bool = (fi34 <= (fi34 * (v1 * (fi28 + -531))))
                                                _ = v47
                                            }
                                            fmt.Println(v26, v5, v5)
                                            if (!(v1 < (v1 >> (((v24 + 915)) & 63)))) {
                                                v48 := GenS6{F0: v42, F1: 42, F2: -9223372036854775808}
                                                _ = v48
                                            } else {
                                                switch (("" + "foo bar") + "foo bar") {
                                                case "":
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi28 >> ((2) & 63))
                                                    }
                                                case "foo bar":
                                                    if v42 {
                                                        if (v43 > v43) {
                                                            var v49 string = "hello"
                                                            if (((v4 - uint32(v5)) > (v4 + v4)) || v42) {
                                                                switch v49 {
                                                                case "foo bar":
                                                                    fmt.Println(v5, v25)
                                                                case "\n\t":
                                                                    fmt.Println(v5)
                                                                default:
                                                                    v50 := GenS6{F0: (100 != 88), F1: 83, F2: int64(v43)}
                                                                    _ = v50
                                                                }
                                                                if v42 {
                                                                    v51 := GenS1{F0: 42, F1: 127, F2: v49}
                                                                    v1 = (fi34 >> ((v24) & 63))
                                                                    v24 += v24
                                                                    _ = v51
                                                                } else {
                                                                    if ((1 <= 42) && (float32(v26) <= 1.5)) {
                                                                        if (v42 && (((1e10 / v26) * float64(v25)) < (float64((v25 * (1.5 + ((((float32(1.5) - (v25 * v25)) / (v25 - float32(0.0))) / (((0.1 * 3.14159) + (float32(v26) + float32(1e10))) / (v25 * -2.25))) - (v25 * v25))))) * (3.14159 * -2.25)))) {
                                                                            v52 := GenS4{F0: 42}
                                                                            if false {
                                                                                switch v49 {
                                                                                case "foo bar":
                                                                                    v26 -= v26
                                                                                case "a":
                                                                                    if true {
                                                                                        v49 = (v49 + ((((v49 + "hello") + v49) + ((("a" + (("\n\t" + (v49 + (v49 + v49))) + v49)) + v49) + "hello")) + ("\n\t" + "a")))
                                                                                        switch (fi28 | -128) {
                                                                                        case -388:
                                                                                            var v53 int8 = -1
                                                                                            _ = v53
                                                                                        case -1:
                                                                                            fmt.Println(v25, v1)
                                                                                        case -2147483648:
                                                                                            if (!(v42 || (!v42))) {
                                                                                                if true {
                                                                                                    if false {
                                                                                                        fmt.Println(v5, v25, v1)
                                                                                                    } else {
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            v42 = (!(!((65535 != (v4 * (v4 ^ v4))) && (!v42))))
                                                                                                        case "":
                                                                                                            v26 -= 3.14159
                                                                                                        case "foo bar":
                                                                                                            v54 := GenS2{F0: ((42 <= 66) && (v49 != ""))}
                                                                                                            _ = v54
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    if (!v42) {
                                                                                                        fmt.Println(v24, v5)
                                                                                                        var v55 uint16 = uint16(v4)
                                                                                                        v56 := v2[(v4 + (v4 % ((2147483647) | 1)))]
                                                                                                        _ = v55
                                                                                                        _ = v56
                                                                                                    }
                                                                                                    v57 := GenS7{F0: (v4 & 865), F1: v49}
                                                                                                    delete(v2, 2)
                                                                                                    v58 := GenS2{F0: v42}
                                                                                                    _ = v57
                                                                                                    _ = v58
                                                                                                }
                                                                                                v59 := GenS6{F0: v42, F1: -7, F2: 42}
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                if true {
                                                                                                    if ((v24 % (((v24 % ((v24) | 1))) | 1)) < 65535) {
                                                                                                        v2[v4] = v42
                                                                                                        var v60 bool = (309 == -9223372036854775808)
                                                                                                        fmt.Println(v1, v49, v43)
                                                                                                        v4 = (v4 & (v4 * v4))
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    switch (fi34 ^ fi28) {
                                                                                                    case 32767:
                                                                                                        v61 := GenS5{F0: 404, F1: -128}
                                                                                                        _ = v61
                                                                                                    case 9223372036854775807:
                                                                                                        v62 := GenS8{F0: 0}
                                                                                                        _ = v62
                                                                                                    case 554:
                                                                                                        var v63 uint64 = 1
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    switch (v1 | -2147483648) {
                                                                                                    case 910:
                                                                                                        switch v49 {
                                                                                                        case "":
                                                                                                            fmt.Println(fi28, fi34, v42)
                                                                                                        case "hello":
                                                                                                            if v42 {
                                                                                                                v64 := GenS7{F0: v4, F1: "foo bar"}
                                                                                                                v5 += 1
                                                                                                                _ = v64
                                                                                                            } else {
                                                                                                                fmt.Println(v26, fi34)
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v65 := GenS1{F0: v24, F1: 198, F2: v49}
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    case 524:
                                                                                                        var v66 bool = (!v42)
                                                                                                        _ = v66
                                                                                                    case -1:
                                                                                                        if false {
                                                                                                            fmt.Println(fi28, v49)
                                                                                                            v2[(^(v4))] = (v25 > ((float32(v26) / v25) - float32(v26)))
                                                                                                        } else {
                                                                                                            v3 = append(v3, 9223372036854775807)
                                                                                                            delete(v2, (^(v4)))
                                                                                                        }
                                                                                                    default:
                                                                                                        fmt.Println(v43, v25)
                                                                                                    }
                                                                                                } else {
                                                                                                    switch v49 {
                                                                                                    case "\n\t":
                                                                                                        var v67 int64 = 0
                                                                                                        _ = v67
                                                                                                    case "foo bar":
                                                                                                        v43 += (v43 * (v43 % ((v43) | 1)))
                                                                                                    }
                                                                                                    v49 = v49
                                                                                                    fmt.Println(v5, fi28)
                                                                                                }
                                                                                                if (!(914 <= 0)) {
                                                                                                    if false {
                                                                                                        fmt.Println(v43, v24, v43)
                                                                                                        var v68 float64 = float64((v25 + ((v25 * (v25 / (v25 * 1e10))) + v25)))
                                                                                                        v26 = 3.14159
                                                                                                        var v69 int = 128
                                                                                                        _ = v68
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    var v70 uint32 = v4
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    var v71 float32 = (1.5 - float32(v26))
                                                                                                    var v72 string = v49
                                                                                                    _ = v71
                                                                                                    _ = v72
                                                                                                }
                                                                                                _ = v59
                                                                                            }
                                                                                        }
                                                                                        if v42 {
                                                                                            var v73 uint8 = (v5 - 42)
                                                                                            _ = v73
                                                                                        }
                                                                                        v74 := GenS7{F0: v4, F1: v49}
                                                                                        v49 = v49
                                                                                        switch v49 {
                                                                                        case "\n\t":
                                                                                            switch v49 {
                                                                                            case "":
                                                                                                switch (fi28 - v1) {
                                                                                                case 0:
                                                                                                    v75 := GenS6{F0: (0 == 42), F1: -80, F2: 509}
                                                                                                    _ = v75
                                                                                                case 2147483647:
                                                                                                    v5 += v5
                                                                                                case -9223372036854775808:
                                                                                                    var v76 bool = v42
                                                                                                    _ = v76
                                                                                                default:
                                                                                                    if false {
                                                                                                        v77 := GenS9{F0: v5}
                                                                                                        if v42 {
                                                                                                            var v78 int64 = -9223372036854775808
                                                                                                            _ = v78
                                                                                                        }
                                                                                                        v79 := GenS10{F0: v24}
                                                                                                        _ = v77
                                                                                                        _ = v79
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                        switch fi34 {
                                                                                                        case 32768:
                                                                                                            fmt.Println(v43)
                                                                                                        case 1:
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    v26 += 1.5
                                                                                                                    if (!(v49 <= ("hello" + v49))) {
                                                                                                                        v25 += (0.0 + (float32(0.1) + (float32(-2.25) / v25)))
                                                                                                                    }
                                                                                                                    switch v49 {
                                                                                                                    case "hello":
                                                                                                                        switch "hello" {
                                                                                                                        case "\n\t":
                                                                                                                            var v80 uint16 = 65535
                                                                                                                            _ = v80
                                                                                                                        case "foo bar":
                                                                                                                            switch v49 {
                                                                                                                            case "\n\t":
                                                                                                                                v81 := GenS3{F0: false, F1: 9223372036854775807, F2: -7}
                                                                                                                                _ = v81
                                                                                                                            case "hello":
                                                                                                                                v82 := GenS10{F0: (v24 / (((v24 << ((v24) & 63))) | 1))}
                                                                                                                                _ = v82
                                                                                                                            case "":
                                                                                                                                var v83 int8 = -14
                                                                                                                                _ = v83
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            switch fi28 {
                                                                                                                            case 127:
                                                                                                                                var v84 int16 = 1
                                                                                                                                _ = v84
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        var v85 int16 = 1
                                                                                                                        _ = v85
                                                                                                                    case "":
                                                                                                                        switch (((v49 + v49) + v49) + v49) {
                                                                                                                        case "a":
                                                                                                                            var v86 uint = v24
                                                                                                                            _ = v86
                                                                                                                        default:
                                                                                                                            switch v49 {
                                                                                                                            case "hello":
                                                                                                                                var v87 int16 = 32767
                                                                                                                                _ = v87
                                                                                                                            case "a":
                                                                                                                                if v42 {
                                                                                                                                    v5 -= (^(v5))
                                                                                                                                    v1 = int(v5)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v88 := v2[652]
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                v89 := GenS9{F0: (v5 * v5)}
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            case "hello":
                                                                                                delete(v2, 442)
                                                                                            default:
                                                                                                v90 := GenS7{F0: (v4 / (((v4 - 4294967295)) | 1)), F1: v49}
                                                                                                _ = v90
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v3 = append(v3, v1)
                                                                                            v91 := GenS3{F0: true, F1: int64(v5), F2: 46}
                                                                                            _ = v91
                                                                                        case "hello":
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (-(fi34))
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi34, v26, v4)
                                                                                        delete(v2, (v4 - (v4 | (v4 & v4))))
                                                                                        _ = v74
                                                                                    }
                                                                                }
                                                                                fmt.Println(v26, fi28, v42)
                                                                                fmt.Println(fi28, fi28, v4)
                                                                            } else {
                                                                                switch (fi28 << ((42) & 63)) {
                                                                                case -63:
                                                                                    fmt.Println(v5)
                                                                                case 0:
                                                                                    if (82 == 127) {
                                                                                        if ((v49 <= v49) || ((-667 == 0) && (0 <= 525))) {
                                                                                            fmt.Println(fi34, v42, fi28)
                                                                                            v92 := GenS11{F0: 128}
                                                                                            fmt.Println(v24, v25, v5)
                                                                                            _ = v92
                                                                                        } else {
                                                                                            v93 := GenS3{F0: false, F1: 310, F2: 0}
                                                                                            if (v42 && true) {
                                                                                                fmt.Println(v43)
                                                                                                v24 += v24
                                                                                            } else {
                                                                                                switch v49 {
                                                                                                case "foo bar":
                                                                                                    var v94 uint8 = uint8(v4)
                                                                                                    _ = v94
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        var v95 bool = (65535 > (v24 - (^(v24))))
                                                                                                        fmt.Println(v5, fi28, v4)
                                                                                                        v96 := GenS12{F0: 100, F1: -1}
                                                                                                        _ = v95
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    v97 := GenS13{F0: (fi34 - (fi34 + 2147483647)), F1: -670}
                                                                                                    _ = v97
                                                                                                case "":
                                                                                                    if v42 {
                                                                                                        if v42 {
                                                                                                            v4 = 65535
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v93
                                                                                        }
                                                                                        var v98 int16 = 32767
                                                                                        delete(v2, (v4 % (((^(v4))) | 1)))
                                                                                        _ = v98
                                                                                    } else {
                                                                                        v99 := GenS14{F0: v43, F1: v49, F2: 95}
                                                                                        v26 += (v26 / v26)
                                                                                        _ = v99
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    v100 := GenS12{F0: 65, F1: 1}
                                                                                    _ = v100
                                                                                default:
                                                                                    switch v49 {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v24, v43)
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (v1 * (fi34 << ((2) & 63)))
                                                                                        }
                                                                                    }
                                                                                    if (!(!v42)) {
                                                                                        v5 = v5
                                                                                    } else {
                                                                                        var v101 uint16 = 1
                                                                                        v102 := GenS15{F0: 9223372036854775807, F1: 42}
                                                                                        var v103 string = (v49 + ((v49 + v49) + ((v49 + ((v49 + (v49 + ((v49 + "a") + "hello"))) + "\n\t")) + v49)))
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    }
                                                                                    if true {
                                                                                        switch fi34 {
                                                                                        case 32768:
                                                                                            v24 = 0
                                                                                        default:
                                                                                            var v104 float32 = (v25 / (float32(v26) + (v25 + (v25 + float32(v26)))))
                                                                                            _ = v104
                                                                                        }
                                                                                        fmt.Println(fi34, fi28, v25)
                                                                                        if v42 {
                                                                                            fmt.Println(v49)
                                                                                            switch "\n\t" {
                                                                                            case "a":
                                                                                                v105 := GenS16{F0: 1, F1: v43, F2: -1}
                                                                                                _ = v105
                                                                                            case "":
                                                                                                var v106 float32 = 1.5
                                                                                                _ = v106
                                                                                            }
                                                                                            if ("hello" > ("hello" + v49)) {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                _ = len(v3)
                                                                                                fmt.Println(fi34)
                                                                                                var v107 float32 = (v25 / v25)
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            if v42 {
                                                                                                var v108 float32 = (v25 * 0.0)
                                                                                                if v42 {
                                                                                                    fmt.Println(v49, v5, v42)
                                                                                                    v4 = v4
                                                                                                } else {
                                                                                                    delete(v2, 65535)
                                                                                                }
                                                                                                _ = v108
                                                                                            }
                                                                                            v109 := GenS12{F0: -128, F1: -1}
                                                                                            if false {
                                                                                                v1 = (fi34 << (((^(v24))) & 63))
                                                                                                v4 += 2147483648
                                                                                            } else {
                                                                                                fmt.Println(v42, v25, v5)
                                                                                                if false {
                                                                                                    var v110 float64 = v26
                                                                                                    fmt.Println(v43)
                                                                                                    _ = v110
                                                                                                }
                                                                                            }
                                                                                            v111 := GenS11{F0: (v1 + (v1 << (((v24 / ((408) | 1))) & 63)))}
                                                                                            _ = v109
                                                                                            _ = v111
                                                                                        } else {
                                                                                            v25 += v25
                                                                                            if v42 {
                                                                                                var v112 int64 = 1
                                                                                                _ = v112
                                                                                            }
                                                                                            v113 := GenS17{F0: (("hello" + v49) + (v49 + v49)), F1: 894, F2: v49}
                                                                                            switch ("" + v49) {
                                                                                            case "\n\t":
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = -1
                                                                                                }
                                                                                            case "":
                                                                                                var v114 uint64 = genFn0()
                                                                                                _ = v114
                                                                                            }
                                                                                            switch fi28 {
                                                                                            case -480:
                                                                                                v115 := GenS5{F0: v4, F1: -1}
                                                                                                var v116 uint32 = (v4 - (v4 ^ v4))
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            case -9223372036854775808:
                                                                                                v117 := GenS16{F0: 42, F1: (v43 % ((v43) | 1)), F2: 32767}
                                                                                                _ = v117
                                                                                            }
                                                                                            _ = v113
                                                                                        }
                                                                                    }
                                                                                    var v118 uint32 = v4
                                                                                    v119 := GenS9{F0: v5}
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                }
                                                                                v26 += ((v26 - (v26 + (v26 + (v26 / v26)))) / (1e10 - 1.5))
                                                                            }
                                                                            var v120 int32 = (v43 - v43)
                                                                            _ = v120
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                }
                                                                v121 := GenS18{F0: (fi34 & -824), F1: 2147483647}
                                                                _ = v121
                                                            }
                                                            _ = v49
                                                        }
                                                    }
                                                case "\n\t":
                                                    if (v5 == (v5 | (v5 * v5))) {
                                                        var v122 int = 2147483647
                                                        fmt.Println(v1)
                                                        _ = v122
                                                    }
                                                }
                                                v123 := GenS12{F0: 42, F1: 8}
                                                var v124 uint16 = 700
                                                _ = v123
                                                _ = v124
                                            }
                                        } else {
                                            v4 += 0
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            }
                            _ = fi34
                        }
                    }
                    if ("hello" != ("\n\t" + (("a" + (("hello" + "\n\t") + "\n\t")) + "\n\t"))) {
                        v125 := GenS19{F0: 100}
                        for fi126 := 0; fi126 < 11; fi126++ {
                            var v127 string = "\n\t"
                            v128 := GenS10{F0: uint(v4)}
                            fmt.Println(fi28)
                            v26 += v26
                            _ = fi126
                            _ = v127
                            _ = v128
                        }
                        _ = v125
                    }
                    _ = fi28
                }
            }
        }
        _ = v26
    } else {
        if false {
            var v129 int = v1
            v130 := GenS20{F0: 42, F1: (v129 | -479), F2: -2147483648}
            v131 := GenS9{F0: (v5 << (((v24 + (v24 ^ 65535))) & 63))}
            _ = v129
            _ = v130
            _ = v131
        }
    }
    var v132 int8 = 42
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v133 int8 = v132
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v132
    _ = v133
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
