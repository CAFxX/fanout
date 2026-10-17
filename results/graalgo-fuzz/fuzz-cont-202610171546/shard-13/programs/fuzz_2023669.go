package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 uint; F2 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 int; F2 uint64 }
type GenS11 struct { F0 string; F1 int32; F2 uint64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint16; F1 int8; F2 int8 }
type GenS18 struct { F0 int }
func genFn0(p0 int16) bool {
    var v5 uint = 2
    switch 873 {
    case 32768:
        var v6 float32 = 3.14159
        _ = v6
    default:
        v5 += v5
    }
    p0 += -1
    _ = p0
    _ = v5
    return ((-2.25 - 0.0) > (3.14159 + 1e10))
}
func genFn1() uint8 {
    if (false && false) {
        var v7 int8 = 1
        fmt.Println(v7)
        v7 = (v7 ^ (v7 ^ v7))
        _ = v7
    } else {
        v8 := GenS1{F0: (-2147483648 != 42)}
        var v9 uint16 = 0
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int = (fi10 + (fi10 & 9223372036854775807))
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    var v12 bool = false
    v12 = v12
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint16 = 2
        var v15 int16 = 1000
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v12
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 2
    var v3 uint16 = 24
    v4 := []uint32{}
    v16 := GenS1{F0: false}
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            if true {
                if (!false) {
                    var v19 int = (^(fi18))
                    if ((!true) && true) {
                        v20 := GenS1{F0: false}
                        _ = v20
                    }
                    var v21 float64 = (1.5 / float64(((((float32((float64(0.1) - (float64(0.1) + 0.0))) - ((-2.25 - float32((float64(float32(1.5)) + (0.0 + (0.1 + 0.0))))) / (((0.0 + (float32(0.1) / 1.0)) - -2.25) / (1.5 + 1e10)))) / 1.5) - (-2.25 * 3.14159)) + 3.14159)))
                    _ = v19
                    _ = v21
                }
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                if false {
                    var v22 uint = 1
                    var v23 uint32 = 2147483648
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v22)
                    case "\n\t":
                        if ((^(v2)) < (^(v2))) {
                            fmt.Println(v22, fi17, v22)
                        } else {
                            if (-2.25 > 0.1) {
                                switch "" {
                                case "a":
                                    switch 0 {
                                    case 584:
                                        v22 += uint(v2)
                                    case 42:
                                        if (!((^(v23)) < (v23 / (((v23 * v23)) | 1)))) {
                                            v24 := GenS1{F0: ((("\n\t" + "") < "a") || false)}
                                            _ = v24
                                        }
                                    default:
                                        v25 := GenS2{F0: ("" + "a"), F1: 395, F2: (v22 ^ (v22 * 4294967295))}
                                        _ = v25
                                    }
                                case "\n\t":
                                    fmt.Println(v23, fi18, v23)
                                case "hello":
                                    v3 += 59
                                    switch "hello" {
                                    case "foo bar":
                                        delete(v1, (fi17 - (fi17 * (fi17 / (((fi17 << ((112) & 63))) | 1)))))
                                        fmt.Println(fi18)
                                    case "hello":
                                        _ = len(v4)
                                    }
                                }
                                v26 := GenS3{F0: 128, F1: 0}
                                _ = v26
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v27 float32 = 0.0
                            _ = v27
                        }
                    case "a":
                        if false {
                            v28 := GenS4{F0: -32768, F1: v23, F2: 1}
                            v2 = uint8(v22)
                            var v29 int32 = 343
                            _ = v28
                            _ = v29
                        } else {
                            fmt.Println(v23, fi18, fi17)
                            var v30 uint16 = v3
                            switch "a" {
                            case "":
                                var v31 uint = v22
                                _ = v31
                            case "a":
                                switch (-(fi17)) {
                                case 9223372036854775807:
                                    v2 -= (v2 % ((v2) | 1))
                                }
                            case "\n\t":
                                switch "hello" {
                                case "a":
                                    v32 := GenS5{F0: (v30 ^ v30), F1: 168}
                                    _ = v32
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            default:
                                if (0 != 0) {
                                    if false {
                                        var v33 uint = (^(v22))
                                        v2 = v2
                                        _ = v33
                                    } else {
                                        if false {
                                            if true {
                                                switch (^(fi17)) {
                                                case -883:
                                                    v34 := GenS6{F0: 972}
                                                    fmt.Println(v2)
                                                    var v35 float64 = 1.5
                                                    _ = v34
                                                    _ = v35
                                                case -1:
                                                    v36 := GenS1{F0: false}
                                                    _ = v36
                                                default:
                                                    var v37 uint16 = (v30 >> ((18446744073709551615) & 63))
                                                    _ = v37
                                                }
                                                switch ("" + ("\n\t" + "a")) {
                                                case "":
                                                    v38 := GenS7{F0: uint(v23), F1: 1, F2: false}
                                                    var v39 string = "a"
                                                    _ = v38
                                                    _ = v39
                                                case "hello":
                                                    v40 := GenS5{F0: (^(v30)), F1: v30}
                                                    _ = v40
                                                case "foo bar":
                                                    if (((972 == 0) || false) || ((v2 | uint8(v30)) == (v2 ^ 235))) {
                                                        var v41 int64 = 0
                                                        v1[(fi18 % (((fi18 | fi17)) | 1))] = (v23 - v23)
                                                        if true {
                                                            if (!((v23 << (((v22 >> ((v22) & 63))) & 63)) != (v23 - uint32(v2)))) {
                                                                var v42 uint32 = (^(v23))
                                                                v43 := v1[(fi18 + 2147483647)]
                                                                v41 += v41
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                if ((("foo bar" + "foo bar") == "foo bar") && (int8(v3) <= 1)) {
                                                                    if true {
                                                                        _ = len(v1)
                                                                        if (3.14159 >= (0.0 * (((0.0 * (((3.14159 / ((3.14159 - (float32((0.1 + (0.0 / -2.25))) * 3.14159)) * 1.5)) / (1.5 * (0.1 / -2.25))) / -2.25)) + -2.25) - 1.5))) {
                                                                            var v44 uint8 = v2
                                                                            var v45 int = (fi17 - (fi17 % ((fi17) | 1)))
                                                                            v46 := GenS3{F0: uint8(v3), F1: 42}
                                                                            v23 += (v23 + (v23 ^ 817))
                                                                            var v47 uint = v22
                                                                            fmt.Println(v23)
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                if ((v47 * 255) != v22) {
                                                                                    v30 += (v30 & v30)
                                                                                } else {
                                                                                    v1[fi18] = (v23 - 1)
                                                                                    v48 := v1[(-(v45))]
                                                                                    _ = v48
                                                                                }
                                                                            }
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            v49 := GenS2{F0: ("foo bar" + "a"), F1: 2, F2: (v22 + v22)}
                                                                            switch ("" + ("" + "")) {
                                                                            case "a":
                                                                                var v50 uint = uint(v3)
                                                                                _ = v50
                                                                            default:
                                                                                if ((!(false || (!true))) || false) {
                                                                                    v51 := GenS8{F0: -332}
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    default:
                                                                                        switch ("hello" + ("\n\t" + "\n\t")) {
                                                                                        case "\n\t":
                                                                                            v52 := GenS6{F0: -1}
                                                                                            _ = v52
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            switch ("hello" + "foo bar") {
                                                                                            case "hello":
                                                                                                if false {
                                                                                                    var v53 bool = true
                                                                                                    switch (fi18 & 62) {
                                                                                                    case -883:
                                                                                                        v53 = false
                                                                                                    case -9223372036854775808:
                                                                                                        delete(v1, (fi18 >> (((^(v22))) & 63)))
                                                                                                    }
                                                                                                    fmt.Println(v53, fi17, fi17)
                                                                                                    _ = v53
                                                                                                } else {
                                                                                                    var v54 float64 = -2.25
                                                                                                    fmt.Println(v23, v22, v2)
                                                                                                    _ = v54
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v22 -= v22
                                                                                            case "foo bar":
                                                                                                var v55 uint = v22
                                                                                                _ = v55
                                                                                            }
                                                                                        case "":
                                                                                            switch (fi18 / ((fi18) | 1)) {
                                                                                            case 9223372036854775807:
                                                                                                switch (fi17 / ((int(v22)) | 1)) {
                                                                                                case -2147483648:
                                                                                                    if (0.1 <= 3.14159) {
                                                                                                        if ((-1 == -128) || (1e10 > -2.25)) {
                                                                                                            v30 = (v3 << ((42) & 63))
                                                                                                        }
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = 65535
                                                                                                        }
                                                                                                        v3 += (^(v30))
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    v30 += v30
                                                                                                case 32768:
                                                                                                    var v56 bool = (1000 == 42)
                                                                                                    _ = v56
                                                                                                }
                                                                                            case 32767:
                                                                                                v30 += (v3 + (v3 ^ uint16(v2)))
                                                                                            default:
                                                                                                fmt.Println(v3, v41, v41)
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v57 := GenS9{F0: (2147483647 >= -439)}
                                                                                            var v58 uint = (^(v22))
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                        }
                                                                                        v23 += 4294967295
                                                                                        v23 += (v23 / ((v23) | 1))
                                                                                    }
                                                                                    v3 -= (v3 & v30)
                                                                                    _ = v51
                                                                                } else {
                                                                                    v2 = (v2 | (v2 << ((v22) & 63)))
                                                                                    fmt.Println(v30)
                                                                                    v59 := v1[(fi18 - (fi18 ^ fi18))]
                                                                                    _ = v59
                                                                                }
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                        v41 -= int64(v3)
                                                                    } else {
                                                                        fmt.Println(v3, fi18, v23)
                                                                        var v60 bool = false
                                                                        v61 := GenS10{F0: 1, F1: (fi17 % (((^(fi17))) | 1)), F2: 4294967295}
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    var v62 bool = false
                                                                    _ = v62
                                                                }
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v63 int8 = 64
                                                                    _ = v63
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                var v64 string = (("hello" + "foo bar") + "foo bar")
                                                                _ = v64
                                                            }
                                                        } else {
                                                            v65 := v1[fi17]
                                                            v66 := GenS11{F0: ((("\n\t" + "") + "\n\t") + "a"), F1: 912, F2: 42}
                                                            switch 32768 {
                                                            case 127:
                                                                var v67 float32 = (1e10 + (float32(float64(1e10)) / -2.25))
                                                                _ = v67
                                                            case 618:
                                                                _ = len(v4)
                                                            case -1:
                                                                v4 = append(v4, (v65 >> ((255) & 63)))
                                                            }
                                                            fmt.Println(fi18, v30, v30)
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        _ = v41
                                                    }
                                                default:
                                                    if ((((true || false) || true) && true) && (!(!true))) {
                                                        v3 = v30
                                                        if ((-128 > -75) && (("foo bar" + "hello") <= "")) {
                                                            v3 = (v3 | v3)
                                                            fmt.Println(fi18, v2)
                                                        } else {
                                                            v30 = uint16(v2)
                                                            var v68 int32 = int32(v22)
                                                            v4 = append(v4, v23)
                                                            _ = v68
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                }
                                                var v69 int16 = 1
                                                var v70 float64 = 0.1
                                                _ = v69
                                                _ = v70
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    if true {
                                                        var v71 string = "a"
                                                        _ = v71
                                                    } else {
                                                        var v72 int32 = -2147483648
                                                        _ = v72
                                                    }
                                                case "":
                                                    v4 = append(v4, v23)
                                                default:
                                                    fmt.Println(v23, v3)
                                                    v73 := GenS4{F0: 1, F1: v23, F2: v3}
                                                    _ = v73
                                                }
                                                var v74 int32 = 0
                                                v75 := GenS10{F0: 42, F1: (fi18 + fi18), F2: 0}
                                                _ = v74
                                                _ = v75
                                            }
                                            switch (fi17 / (((fi17 >> ((107) & 63))) | 1)) {
                                            case 127:
                                                var v76 uint32 = v23
                                                _ = v76
                                            case -800:
                                                v30 = v30
                                            }
                                            var v77 float64 = (0.0 * 3.14159)
                                            _ = v77
                                        } else {
                                            switch (fi17 ^ fi18) {
                                            case 9223372036854775807:
                                                v78 := GenS12{F0: (^(v23))}
                                                _ = v78
                                            }
                                            delete(v1, fi17)
                                            _ = len(v4)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        fmt.Println(v2, fi17, v22)
                                        switch "foo bar" {
                                        case "":
                                            switch (("" + "\n\t") + ("" + ("" + (("" + "hello") + ("foo bar" + ""))))) {
                                            case "":
                                                if len(v4) > 0 {
                                                    v4[0] = (v23 * (^(v23)))
                                                }
                                            case "\n\t":
                                                var v79 string = "foo bar"
                                                _ = v79
                                            default:
                                                v23 += 65535
                                            }
                                        case "a":
                                            v80 := GenS13{F0: "foo bar"}
                                            _ = v80
                                        default:
                                            if true {
                                                var v81 string = "hello"
                                                if (42 > 942) {
                                                    var v82 uint = v22
                                                    fmt.Println(v22)
                                                    _ = v82
                                                }
                                                _ = len(v1)
                                                v83 := GenS6{F0: -1}
                                                fmt.Println(v3, v22, v30)
                                                if (-1 <= 81) {
                                                    var v84 int16 = 905
                                                    v4 = append(v4, v23)
                                                    v85 := GenS5{F0: (v3 * 0), F1: v3}
                                                    if ((-2.25 + ((((float64(3.14159) * (float64(float32((((1.5 / -2.25) * 0.0) + 0.1))) + (0.0 - 3.14159))) / 0.1) - 0.0) - float64(0.1))) < (0.0 + (1.5 / 1.0))) {
                                                        var v86 uint16 = v3
                                                        if true {
                                                            if false {
                                                                v23 += v23
                                                                v87 := GenS14{F0: false}
                                                                _ = v87
                                                            }
                                                            fmt.Println(fi17, v2, v3)
                                                        } else {
                                                            var v88 uint16 = 865
                                                            if true {
                                                                var v89 uint64 = 42
                                                                if true {
                                                                    v81 = "foo bar"
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = 1
                                                                }
                                                                v90 := GenS15{F0: 1, F1: (v84 & (v84 & 1))}
                                                                v3 -= (v88 & v88)
                                                                if (!(((!(((v22 + v22) >= 731) && true)) || false) && ((^(v88)) != uint16(v22)))) {
                                                                    v84 = v84
                                                                    var v91 int = -476
                                                                    var v92 int32 = -813
                                                                    v88 += v30
                                                                    v81 = v81
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v89 -= (v89 * (v89 | (v89 & uint64(v22))))
                                                                    var v93 uint16 = 1
                                                                    _ = v93
                                                                }
                                                                var v94 int16 = v84
                                                                var v95 int8 = 127
                                                                var v96 int = (fi17 * fi18)
                                                                fmt.Println(v88, v86)
                                                                var v97 bool = false
                                                                _ = v89
                                                                _ = v90
                                                                _ = v94
                                                                _ = v95
                                                                _ = v96
                                                                _ = v97
                                                            } else {
                                                                v2 -= v2
                                                                v98 := v1[fi17]
                                                                fmt.Println(v84, v23, v84)
                                                                _ = v98
                                                            }
                                                            var v99 bool = ((v23 << ((65535) & 63)) <= v23)
                                                            _ = v88
                                                            _ = v99
                                                        }
                                                        v100 := GenS16{F0: (!(!(!true)))}
                                                        _ = v100
                                                        _ = v86
                                                    } else {
                                                        var v101 int16 = v84
                                                        v102 := GenS12{F0: (v23 ^ 204)}
                                                        fmt.Println(v23)
                                                        var v103 int32 = 2147483647
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    v104 := GenS2{F0: "hello", F1: 2, F2: 4294967295}
                                                    if (v30 < v3) {
                                                        var v105 int = fi17
                                                        _ = v105
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = 2147483648
                                                        }
                                                        if true {
                                                            var v106 int64 = 1
                                                            switch 591 {
                                                            case 42:
                                                                v30 += 612
                                                            default:
                                                                v22 -= (v22 >> ((v22) & 63))
                                                                var v107 float64 = (0.0 * 3.14159)
                                                                _ = v107
                                                            }
                                                            _ = v106
                                                        } else {
                                                            switch (fi18 << (((^(v22))) & 63)) {
                                                            case -259:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            case 42:
                                                                v108 := GenS17{F0: 519, F1: -1, F2: 90}
                                                                _ = v108
                                                            }
                                                            v81 = ""
                                                            _ = len(v4)
                                                        }
                                                    }
                                                    var v109 uint32 = v23
                                                    if ((v84 - 404) >= v84) {
                                                        v110 := GenS14{F0: false}
                                                        switch (fi18 - fi18) {
                                                        case 42:
                                                            v111 := GenS18{F0: -9223372036854775808}
                                                            _ = v111
                                                        }
                                                        _ = v110
                                                    }
                                                    var v112 bool = ((v2 & genFn1()) == (v2 ^ (v2 * v2)))
                                                    if true {
                                                        delete(v1, int(v3))
                                                        fmt.Println(v2, v22, v2)
                                                        _ = len(v1)
                                                    } else {
                                                        v3 = v3
                                                        switch (^(fi18)) {
                                                        case 0:
                                                            fmt.Println(v3, v22, v22)
                                                        case 32767:
                                                            v22 = v22
                                                        }
                                                    }
                                                    v113 := GenS12{F0: (v23 / ((707) | 1))}
                                                    _ = len(v4)
                                                    _ = v104
                                                    _ = v109
                                                    _ = v112
                                                    _ = v113
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = v81
                                                _ = v83
                                            } else {
                                                v30 -= (v30 % ((v3) | 1))
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v30
                        }
                    default:
                        v114 := GenS16{F0: (((!(!(1e10 != 0.1))) || false) || false)}
                        _ = v114
                    }
                    _ = v22
                    _ = v23
                }
            }
            v115 := v1[fi17]
            switch (("foo bar" + ("a" + "foo bar")) + "") {
            case "hello":
                if (((0.1 + 0.0) * ((3.14159 - 0.0) + ((1e10 * 1e10) / 1.5))) >= 1e10) {
                    var v116 int32 = 42
                    v115 = (v115 >> ((4294967295) & 63))
                    v117 := GenS7{F0: 2, F1: 2, F2: (((0.1 + (0.1 * 1e10)) + float64(0.1)) >= (0.1 * 0.1))}
                    _ = v116
                    _ = v117
                } else {
                    switch ("foo bar" + (("\n\t" + "a") + "")) {
                    case "":
                        v115 -= 420
                        if ((fi17 & 42) != fi18) {
                            if (v115 >= (^(v115))) {
                                fmt.Println(v115, v2, fi18)
                                fmt.Println(fi18, v115, v115)
                            }
                        }
                        v118 := v1[(fi17 << ((4294967295) & 63))]
                        _ = v118
                    default:
                        v2 -= (v2 ^ genFn1())
                    }
                    var v119 int32 = 2147483647
                    _ = v119
                }
            case "a":
                v115 -= v115
            case "":
                v120 := v1[fi17]
                _ = v120
            }
            _ = fi18
            _ = v115
        }
        _ = fi17
    }
    var v121 int16 = 0
    for fi122 := 0; fi122 < 10; fi122++ {
        v4 = append(v4, 2)
        _ = fi122
    }
    v4 = append(v4, 1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi123 := 0; fi123 < 20; fi123++ {
        var v124 string = "foo bar"
        fmt.Println(v124, v124)
        _ = fi123
        _ = v124
    }
    v125 := GenS8{F0: (-(v121))}
    v126 := GenS1{F0: false}
    var v127 uint32 = 299
    v127 = (v127 >> ((uint(v127)) & 63))
    if ((v121 >> ((4294967295) & 63)) < 714) {
        fmt.Println(v3, v127, v2)
    }
    v128 := v1[32768]
    v128 = 0
    var v129 bool = ((v128 / (((v127 + (v128 % ((1) | 1)))) | 1)) < v128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v121:", v121)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v121
    _ = v125
    _ = v126
    _ = v127
    _ = v128
    _ = v129
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
