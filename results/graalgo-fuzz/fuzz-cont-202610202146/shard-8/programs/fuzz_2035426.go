package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int; F1 bool; F2 bool }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 string; F1 uint; F2 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint64; F1 string }
type GenS18 struct { F0 int32; F1 bool; F2 uint }
type GenS19 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int; F1 uint8; F2 int32 }
type GenS22 struct { F0 uint32 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 int8; F1 string }
type GenS26 struct { F0 int; F1 int32 }
type GenS27 struct { F0 int8 }
func genFn0(p0 int16) uint {
    fmt.Println(p0)
    p0 = (^(p0))
    p0 = p0
    p0 = p0
    v7 := GenS1{F0: false}
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1() string {
    v8 := GenS1{F0: (0.0 == 1.5)}
    var v9 string = ("" + "\n\t")
    switch "a" {
    case "\n\t":
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: (!(v9 < "a"))}
            switch v9 {
            case "foo bar":
                var v12 uint = 65535
                _ = v12
            case "a":
                fmt.Println(fi10)
            case "\n\t":
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := GenS1{F0: false}
                    _ = fi13
                    _ = v14
                }
            }
            v15 := GenS1{F0: false}
            _ = fi10
            _ = v11
            _ = v15
        }
    }
    fmt.Println(v9)
    _ = v8
    _ = v9
    return (v9 + "foo bar")
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 105
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    var v6 string = ""
    v16 := []int{-590}
    v6 = ""
    v17 := GenS1{F0: (((!false) || true) || (2 != (^(v1))))}
    for fi18 := 0; fi18 < 16; fi18++ {
        if (v2 >= (v2 % (((v2 << ((4294967295) & 63))) | 1))) {
            v16 = append(v16, fi18)
        }
        v1 -= v1
        v2 = 227
        _ = fi18
    }
    _ = len(v16)
    _ = len(v16)
    if true {
        switch "hello" {
        case "hello":
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 float64 = ((0.1 * 3.14159) / -2.25)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                var v21 int8 = -125
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        var v22 uint32 = 506
        v6 = (v6 + (v6 + v6))
        _ = v22
    } else {
        fmt.Println(v2)
        for fi23 := 0; fi23 < 11; fi23++ {
            v16 = append(v16, fi23)
            v24 := v5[(-(fi23))]
            v25 := GenS1{F0: ((1.5 <= ((3.14159 * 0.1) + float64(((1.5 * (3.14159 - 1.5)) * -2.25)))) && (2147483647 >= 408))}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v16 = append(v16, 1)
    case "hello":
        for fi26 := 0; fi26 < 2; fi26++ {
            for fi27 := 0; fi27 < 7; fi27++ {
                if true {
                    v2 = v2
                } else {
                    var v28 bool = false
                    if v28 {
                        fmt.Println(v1, v2, v28)
                        if (-15 <= 109) {
                            if ((!((fi27 * (fi26 % ((9223372036854775807) | 1))) < 975)) || v28) {
                                if (-2.25 >= (1e10 / 1.0)) {
                                    var v29 int64 = 620
                                    _ = v29
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v30 := GenS2{F0: -104, F1: uint8(v1)}
                                        _ = v30
                                    case "foo bar":
                                        switch ("foo bar" + "\n\t") {
                                        case "foo bar":
                                            var v31 int = 4294967295
                                            _ = v31
                                        case "hello":
                                            var v32 uint8 = v2
                                            _ = v32
                                        case "a":
                                            fmt.Println(v28, v28)
                                        default:
                                            if ((-1 < 1) || v28) {
                                                var v33 uint64 = v1
                                                v34 := v3["a"]
                                                var v35 float64 = 1.5
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            } else {
                                                v6 = (v6 + ("\n\t" + ("hello" + v6)))
                                                v36 := GenS1{F0: (-32768 > 32767)}
                                                var v37 bool = false
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                    case "hello":
                                        fmt.Println(v6)
                                        v28 = (65535 < 0)
                                        v6 = (v6 + v6)
                                    }
                                    if v28 {
                                        if v28 {
                                            if (v28 || v28) {
                                                fmt.Println(fi27, v6, v1)
                                                v28 = ((-44 <= 0) && false)
                                            } else {
                                                v38 := GenS2{F0: -2147483648, F1: 15}
                                                var v39 uint32 = 2147483647
                                                switch (fi27 + fi26) {
                                                case 4294967295:
                                                    v6 = v6
                                                }
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        switch v6 {
                                        case "":
                                            if (1 <= (^(v2))) {
                                                switch "foo bar" {
                                                case "a":
                                                    v1 = v1
                                                case "\n\t":
                                                    var v40 int8 = 4
                                                    v41 := GenS3{F0: 0, F1: -925, F2: 42}
                                                    _ = v40
                                                    _ = v41
                                                default:
                                                    var v42 int = fi27
                                                    _ = v42
                                                }
                                            } else {
                                                fmt.Println(v6, v28)
                                            }
                                        case "a":
                                            var v43 int8 = -1
                                            _ = v43
                                        }
                                        if true {
                                            if ((!v28) || (("foo bar" + (v6 + "a")) != v6)) {
                                                v44 := GenS1{F0: false}
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                                _ = v44
                                            } else {
                                                v45 := GenS4{F0: 389, F1: 65535, F2: 2147483647}
                                                v46 := GenS2{F0: 1, F1: (^(v2))}
                                                _ = v45
                                                _ = v46
                                            }
                                        } else {
                                            switch (((v6 + "foo bar") + "foo bar") + v6) {
                                            case "\n\t":
                                                fmt.Println(v28)
                                            case "":
                                                var v47 int16 = 999
                                                _ = v47
                                            case "hello":
                                                if (!true) {
                                                    var v48 uint8 = 1
                                                    switch "a" {
                                                    case "":
                                                        if v28 {
                                                            if ((0 != 2147483648) || v28) {
                                                                if v28 {
                                                                    if (22 < v1) {
                                                                        if v28 {
                                                                            _ = len(v3)
                                                                            var v49 int32 = 217
                                                                            if true {
                                                                                v49 -= v49
                                                                            }
                                                                            _ = v49
                                                                        } else {
                                                                            v50 := GenS2{F0: -1, F1: (^(v2))}
                                                                            switch "" {
                                                                            case "":
                                                                                if v28 {
                                                                                    v51 := GenS2{F0: 2147483647, F1: (v48 << ((0) & 63))}
                                                                                    var v52 uint = 787
                                                                                    delete(v4, "a")
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                } else {
                                                                                    var v53 int64 = -9223372036854775808
                                                                                    _ = v53
                                                                                }
                                                                                var v54 int = (fi26 % (((fi27 + fi27)) | 1))
                                                                                _ = v54
                                                                            case "foo bar":
                                                                                fmt.Println(v6, fi27, fi26)
                                                                            }
                                                                            _ = v50
                                                                        }
                                                                        if false {
                                                                            v1 -= v1
                                                                        }
                                                                    } else {
                                                                        var v55 uint = 255
                                                                        var v56 uint8 = 42
                                                                        v57 := GenS5{F0: 1, F1: 731}
                                                                        _ = v55
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                } else {
                                                                    var v58 bool = (571 <= 762)
                                                                    _ = v58
                                                                }
                                                                var v59 bool = v28
                                                                v60 := GenS6{F0: 32767, F1: (0 < 538), F2: (0.0 == (((0.1 / (1e10 + 0.1)) / ((-2.25 - 0.0) / (-2.25 - 0.1))) - 1.5))}
                                                                var v61 uint32 = 560
                                                                if (fi27 >= (fi26 & fi26)) {
                                                                    v62 := GenS2{F0: 0, F1: 1}
                                                                    if false {
                                                                        v63 := GenS4{F0: 2, F1: 486, F2: 2147483647}
                                                                        var v64 int32 = 1
                                                                        v65 := v3[v6]
                                                                        v64 -= 0
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                    } else {
                                                                        v61 = uint32(v1)
                                                                        _ = len(v16)
                                                                    }
                                                                    v48 = 130
                                                                    _ = v62
                                                                } else {
                                                                    switch v6 {
                                                                    case "":
                                                                        var v66 float32 = 1.5
                                                                        _ = v66
                                                                    case "hello":
                                                                        v1 = (v1 | (^(v1)))
                                                                    case "a":
                                                                        if (-2.25 <= 1e10) {
                                                                            v67 := GenS3{F0: 42, F1: 1, F2: -9223372036854775808}
                                                                            v59 = true
                                                                            if true {
                                                                                fmt.Println(v2, v28)
                                                                            }
                                                                            if v28 {
                                                                                v28 = v28
                                                                                switch int(v1) {
                                                                                case 32767:
                                                                                    var v68 int8 = 101
                                                                                    _ = v68
                                                                                default:
                                                                                    v28 = v28
                                                                                }
                                                                                if (!v59) {
                                                                                    v69 := GenS3{F0: 255, F1: 42, F2: -1}
                                                                                    if (!(18446744073709551615 != 1)) {
                                                                                        if (!(32767 <= 992)) {
                                                                                            var v70 float32 = 1.5
                                                                                            fmt.Println(v2, v28)
                                                                                            _ = v70
                                                                                        } else {
                                                                                            v71 := GenS7{F0: -1, F1: 32767}
                                                                                            if v28 {
                                                                                                var v72 bool = (-86 != -1)
                                                                                                var v73 uint32 = v61
                                                                                                if v28 {
                                                                                                    v6 = ((v6 + v6) + ((v6 + ((((((v6 + v6) + v6) + "foo bar") + v6) + v6) + ("a" + "hello"))) + "hello"))
                                                                                                    v72 = (v28 || (v59 && (uint(v48) > genFn0(1000))))
                                                                                                    v74 := v5[int(v48)]
                                                                                                    _ = v74
                                                                                                }
                                                                                                _ = v72
                                                                                                _ = v73
                                                                                            }
                                                                                            _ = v71
                                                                                        }
                                                                                        if (42 == -2147483648) {
                                                                                            v2 = (^(v48))
                                                                                        }
                                                                                    }
                                                                                    _ = v69
                                                                                } else {
                                                                                    _ = len(v16)
                                                                                    if len(v16) > 0 {
                                                                                        v16[0] = 2147483647
                                                                                    }
                                                                                    v75 := GenS6{F0: fi26, F1: (v1 <= (v1 + 15)), F2: (v61 < 93)}
                                                                                    _ = v75
                                                                                }
                                                                                var v76 float32 = (3.14159 - ((float32(float64((1e10 + (0.1 - (0.0 * 0.0))))) - (float32(3.14159) + -2.25)) / -2.25))
                                                                                _ = v76
                                                                            } else {
                                                                                switch ((((v6 + v6) + v6) + v6) + ("\n\t" + v6)) {
                                                                                case "foo bar":
                                                                                    if (0.0 >= float64(float32(0.1))) {
                                                                                        if false {
                                                                                            var v77 uint32 = 9
                                                                                            fmt.Println(v28, v48)
                                                                                            var v78 uint = 42
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        }
                                                                                        v1 = (v1 / ((v1) | 1))
                                                                                    } else {
                                                                                        switch fi26 {
                                                                                        case 1:
                                                                                            switch v6 {
                                                                                            case "a":
                                                                                                if v28 {
                                                                                                    v28 = false
                                                                                                    if v59 {
                                                                                                        if (!((fi27 >> ((uint(v48)) & 63)) >= (fi26 + fi27))) {
                                                                                                            v6 = ("a" + v6)
                                                                                                            v61 = (^(v61))
                                                                                                            v79 := GenS1{F0: (9223372036854775807 < 723)}
                                                                                                            _ = v79
                                                                                                        }
                                                                                                        var v80 uint = uint(v61)
                                                                                                        v80 += v80
                                                                                                        switch (fi26 / ((int(v48)) | 1)) {
                                                                                                        case -637:
                                                                                                            if (v59 && v28) {
                                                                                                                var v81 uint16 = uint16(v48)
                                                                                                                _ = v81
                                                                                                            }
                                                                                                        default:
                                                                                                            var v82 int8 = 22
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        _ = v80
                                                                                                    }
                                                                                                    var v83 int32 = -2147483648
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v84 := GenS5{F0: 65535, F1: 65535}
                                                                                                    v85 := GenS2{F0: 520, F1: 42}
                                                                                                    _ = v84
                                                                                                    _ = v85
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v86 uint8 = v48
                                                                                                v87 := GenS2{F0: 2147483647, F1: v2}
                                                                                                switch ("" + v6) {
                                                                                                case "a":
                                                                                                    v61 -= (v61 | v61)
                                                                                                case "foo bar":
                                                                                                    v88 := GenS1{F0: ((2 < 1) && (v59 || (0 < -1)))}
                                                                                                    _ = v88
                                                                                                case "\n\t":
                                                                                                    var v89 float64 = float64(1.5)
                                                                                                    _ = v89
                                                                                                }
                                                                                                v61 -= 4294967295
                                                                                                var v90 int32 = 736
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                                _ = v90
                                                                                            }
                                                                                        case 0:
                                                                                            v59 = (!true)
                                                                                        case 128:
                                                                                            var v91 int64 = int64(fi27)
                                                                                            _ = v91
                                                                                        }
                                                                                        switch fi27 {
                                                                                        case 2147483647:
                                                                                            var v92 uint64 = (v1 / ((v1) | 1))
                                                                                            _ = v92
                                                                                        case 383:
                                                                                            v2 += (v2 | v48)
                                                                                        }
                                                                                        v93 := GenS8{F0: (^(v48)), F1: (-64 > 42)}
                                                                                        if len(v16) >= 2 {
                                                                                            v16 = v16[1:len(v16)]
                                                                                        }
                                                                                        _ = v93
                                                                                    }
                                                                                }
                                                                                var v94 string = ""
                                                                                var v95 uint8 = 128
                                                                                _ = v94
                                                                                _ = v95
                                                                            }
                                                                            fmt.Println(fi26, v59)
                                                                            _ = v67
                                                                        }
                                                                    default:
                                                                        v96 := GenS9{F0: 42}
                                                                        _ = v96
                                                                    }
                                                                    if len(v16) >= 2 {
                                                                        v16 = v16[1:len(v16)]
                                                                    }
                                                                }
                                                                switch v6 {
                                                                case "\n\t":
                                                                    switch (v6 + v6) {
                                                                    case "a":
                                                                        var v97 int8 = -125
                                                                        _ = v97
                                                                    case "hello":
                                                                        v98 := GenS9{F0: -105}
                                                                        _ = v98
                                                                    }
                                                                case "":
                                                                    var v99 string = v6
                                                                    _ = v99
                                                                default:
                                                                    v48 += v2
                                                                }
                                                                switch ("hello" + (v6 + "a")) {
                                                                case "\n\t":
                                                                    var v100 uint32 = (v61 | (^(v61)))
                                                                    _ = v100
                                                                case "hello":
                                                                    v61 = 42
                                                                }
                                                                _ = v59
                                                                _ = v60
                                                                _ = v61
                                                            } else {
                                                                v101 := GenS5{F0: 257, F1: 417}
                                                                v48 = (v2 ^ v2)
                                                                var v102 int8 = -1
                                                                if (true || (v28 || false)) {
                                                                    var v103 uint32 = 4294967295
                                                                    _ = v103
                                                                } else {
                                                                    v16 = append(v16, fi26)
                                                                }
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                            switch v6 {
                                                            case "":
                                                                switch (("hello" + "a") + (v6 + "foo bar")) {
                                                                case "foo bar":
                                                                    v6 = v6
                                                                case "":
                                                                    fmt.Println(v1, v48, v48)
                                                                case "a":
                                                                    var v104 int16 = -1
                                                                    _ = v104
                                                                }
                                                            }
                                                            v2 += 127
                                                            _ = len(v16)
                                                        } else {
                                                            v2 -= v2
                                                        }
                                                    case "\n\t":
                                                        if v28 {
                                                            fmt.Println(v2, v48, fi27)
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                switch ("foo bar" + v6) {
                                                                case "a":
                                                                    var v105 int64 = 9223372036854775807
                                                                    _ = v105
                                                                case "":
                                                                    fmt.Println(fi27, v6)
                                                                }
                                                            case "foo bar":
                                                                v2 = (v2 / (((v2 * uint8(v1))) | 1))
                                                            case "\n\t":
                                                                var v106 int64 = 42
                                                                _ = v106
                                                            default:
                                                                var v107 uint8 = (v2 % ((v48) | 1))
                                                                _ = v107
                                                            }
                                                        }
                                                    }
                                                    _ = v48
                                                }
                                            }
                                            v108 := GenS10{F0: v6, F1: v2}
                                            v1 -= (v1 + 18446744073709551615)
                                            _ = v108
                                        }
                                        fmt.Println(v1, v6)
                                        v4[v6] = (-(fi27))
                                    }
                                }
                            } else {
                                v109 := v4[v6]
                                if ((100 != 62) && (0.0 > (1.5 - (0.0 / 1.0)))) {
                                    var v110 float32 = (float32(3.14159) - -2.25)
                                    _ = v110
                                }
                                _ = v109
                            }
                        } else {
                            var v111 float64 = -2.25
                            if true {
                                if true {
                                    var v112 bool = v28
                                    var v113 int8 = 0
                                    v112 = ((v111 + (v111 / (-2.25 - v111))) < v111)
                                    v16 = append(v16, (fi27 ^ 42))
                                    v114 := GenS11{F0: fi27, F1: -1}
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                } else {
                                    var v115 uint = 2
                                    v116 := GenS12{F0: ("" + "foo bar"), F1: v115, F2: 42}
                                    var v117 int32 = 0
                                    _ = v115
                                    _ = v116
                                    _ = v117
                                }
                                var v118 uint = 4294967295
                                v119 := GenS13{F0: 2}
                                v2 -= 42
                                var v120 int64 = 1
                                _ = v118
                                _ = v119
                                _ = v120
                            } else {
                                if v28 {
                                    var v121 float32 = float32((v111 * v111))
                                    _ = v121
                                }
                                v6 = v6
                            }
                            v122 := GenS14{F0: 1000, F1: 4294967295}
                            v28 = (!false)
                            _ = v111
                            _ = v122
                        }
                        v28 = (32767 > 42)
                    } else {
                        var v123 float64 = 0.0
                        v124 := GenS15{F0: (!(0 > 1))}
                        v125 := GenS16{F0: 32767}
                        var v126 float64 = 1e10
                        var v127 uint = 2
                        _ = v123
                        _ = v124
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    var v128 uint32 = 2147483648
                    _ = v128
                    _ = v28
                }
                _ = fi27
            }
            var v129 uint = 1
            var v130 uint16 = uint16(v1)
            _ = fi26
            _ = v129
            _ = v130
        }
    case "a":
        var v131 float32 = 1.5
        _ = v131
    }
    var v132 string = v6
    for fi133 := 0; fi133 < 1; fi133++ {
        var v134 int64 = 305
        _ = fi133
        _ = v134
    }
    v135 := v4["\n\t"]
    if true {
        if ((-1 < 0) && (true || false)) {
            if false {
                if ((!false) && (1e10 == 0.0)) {
                    var v136 int16 = -32768
                    _ = v136
                } else {
                    var v137 int16 = 42
                    v138 := GenS14{F0: (v137 - 0), F1: 42}
                    v2 = 127
                    _ = v137
                    _ = v138
                }
                var v139 uint16 = 65535
                _ = v139
            } else {
                switch (v135 << ((2) & 63)) {
                case 32767:
                    switch v135 {
                    case 32767:
                        if (!false) {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            for fi140 := 0; fi140 < 3; fi140++ {
                                _ = len(v3)
                                _ = fi140
                            }
                            switch (v6 + "hello") {
                            case "foo bar":
                                var v141 uint = 189
                                _ = v141
                            case "\n\t":
                                fmt.Println(v6)
                            default:
                                _ = len(v16)
                            }
                        }
                    case 4294967295:
                        v142 := GenS17{F0: 1, F1: (v132 + ("a" + ((v6 + "hello") + v6)))}
                        _ = v142
                    default:
                        for fi143 := 0; fi143 < 17; fi143++ {
                            var v144 int = v135
                            fmt.Println(v1, v1, v2)
                            fmt.Println(v6)
                            _ = fi143
                            _ = v144
                        }
                    }
                }
                for fi145 := 0; fi145 < 11; fi145++ {
                    v146 := v5[790]
                    if (42 == (v1 - v1)) {
                        v147 := GenS18{F0: -2147483648, F1: (true || false), F2: 65535}
                        _ = v147
                    } else {
                        fmt.Println(v135, v1, v132)
                        var v148 int = -707
                        switch (v148 ^ fi145) {
                        case 127:
                            if (-36 < 1) {
                                v135 -= (v148 | (fi145 % ((v148) | 1)))
                                if (v2 > (v2 >> ((18446744073709551615) & 63))) {
                                    v132 = v6
                                }
                            } else {
                                v149 := GenS2{F0: 0, F1: 105}
                                for fi150 := 0; fi150 < 12; fi150++ {
                                    v148 = v148
                                    _ = fi150
                                }
                                _ = v149
                            }
                        case 32767:
                            if true {
                                switch "hello" {
                                case "hello":
                                    var v151 bool = true
                                    _ = v151
                                default:
                                    for fi152 := 0; fi152 < 2; fi152++ {
                                        fmt.Println(v2, v1, v6)
                                        _ = fi152
                                    }
                                }
                                var v153 int8 = -87
                                fmt.Println(v135, v1, v132)
                                _ = v153
                            }
                        case -1:
                            for fi154 := 0; fi154 < 1; fi154++ {
                                if (v6 > "") {
                                    var v155 uint16 = 0
                                    if (true && (-2.25 <= ((-2.25 + -2.25) * 1.5))) {
                                        fmt.Println(v1, fi154)
                                    } else {
                                        if len(v16) > 0 {
                                            v16[0] = fi145
                                        }
                                    }
                                    _ = v155
                                }
                                _ = fi154
                            }
                        default:
                            if (true && false) {
                                var v156 float32 = 0.0
                                var v157 uint32 = v146
                                for fi158 := 0; fi158 < 6; fi158++ {
                                    v132 = "\n\t"
                                    v159 := GenS19{F0: 42, F1: 202, F2: 42}
                                    if (!(-729 != 932)) {
                                        if (true && (v6 == ((v6 + "a") + v6))) {
                                            var v160 uint32 = (^(v157))
                                            v156 += v156
                                            v5[0] = (v160 % ((v157) | 1))
                                            _ = v160
                                        } else {
                                            v161 := GenS20{F0: v146}
                                            if (65535 <= 855) {
                                                v135 = fi158
                                                if true {
                                                    v6 = "a"
                                                    var v162 int8 = 100
                                                    _ = v162
                                                }
                                                v6 = v6
                                            } else {
                                                v163 := GenS21{F0: (v135 - v135), F1: (^(v2)), F2: 42}
                                                switch ((v6 + ((v132 + v132) + "foo bar")) + v6) {
                                                case "\n\t":
                                                    fmt.Println(v6, v132)
                                                }
                                                v148 += fi145
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v164 := GenS14{F0: 0, F1: v157}
                                                    _ = v164
                                                case "":
                                                    v165 := GenS22{F0: 65535}
                                                    _ = v165
                                                default:
                                                    v166 := GenS23{F0: 0}
                                                    _ = v166
                                                }
                                                switch v6 {
                                                case "":
                                                    v6 = (v6 + "foo bar")
                                                case "foo bar":
                                                    var v167 uint8 = (v2 ^ v2)
                                                    _ = v167
                                                }
                                                if ((^(v1)) <= (v1 * v1)) {
                                                    if false {
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                        v148 = (fi145 >> ((2) & 63))
                                                        v168 := GenS4{F0: 213, F1: 42, F2: (^(v157))}
                                                        _ = v168
                                                    }
                                                }
                                                _ = v163
                                            }
                                            v132 = v132
                                            switch fi158 {
                                            case 1:
                                                var v169 uint32 = 563
                                                _ = v169
                                            }
                                            v170 := GenS6{F0: v148, F1: true, F2: false}
                                            if (!true) {
                                                var v171 float64 = (1.5 * (0.1 - 0.0))
                                                _ = v171
                                            }
                                            v6 = v132
                                            _ = v161
                                            _ = v170
                                        }
                                    } else {
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        v172 := GenS24{F0: 0}
                                        var v173 int32 = 42
                                        v6 = ""
                                        v174 := GenS10{F0: ("hello" + ("a" + ((v6 + "") + "a"))), F1: (v2 ^ 1)}
                                        _ = v172
                                        _ = v173
                                        _ = v174
                                    }
                                    v157 = 477
                                    v175 := GenS25{F0: 127, F1: v6}
                                    _ = fi158
                                    _ = v159
                                    _ = v175
                                }
                                _ = v156
                                _ = v157
                            } else {
                                v176 := GenS5{F0: v146, F1: v146}
                                v2 -= 0
                                v1 = 18446744073709551615
                                _ = v176
                            }
                        }
                        _ = v148
                    }
                    _ = fi145
                    _ = v146
                }
                if (-771 == (v135 - (v135 + (v135 * v135)))) {
                    if (true || true) {
                        fmt.Println(v6, v6)
                        if false {
                            var v177 uint64 = 938
                            v178 := v5[int(v177)]
                            switch 127 {
                            case 32767:
                                for fi179 := 0; fi179 < 18; fi179++ {
                                    v180 := GenS18{F0: 1, F1: false, F2: genFn0(-1)}
                                    v3["foo bar"] = false
                                    v181 := GenS26{F0: v135, F1: -2147483648}
                                    _ = fi179
                                    _ = v180
                                    _ = v181
                                }
                            default:
                                fmt.Println(v6)
                            }
                            _ = v177
                            _ = v178
                        } else {
                            switch v135 {
                            case 32768:
                                if (!true) {
                                    for fi182 := 0; fi182 < 4; fi182++ {
                                        if true {
                                            v1 = (v1 | 705)
                                        }
                                        switch (fi182 ^ -579) {
                                        case 32767:
                                            var v183 uint16 = 319
                                            _ = v183
                                        default:
                                            v6 = "hello"
                                        }
                                        var v184 string = "\n\t"
                                        _ = fi182
                                        _ = v184
                                    }
                                    v185 := GenS2{F0: 206, F1: (v2 << ((0) & 63))}
                                    _ = v185
                                } else {
                                    if len(v16) > 0 {
                                        v16[0] = v135
                                    }
                                    if ((^(v135)) != (^(v135))) {
                                        v2 += v2
                                        v6 = "hello"
                                    } else {
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    }
                                    var v186 uint8 = v2
                                    _ = v186
                                }
                            case 904:
                                fmt.Println(v1)
                            case -396:
                                for fi187 := 0; fi187 < 13; fi187++ {
                                    var v188 uint32 = 1
                                    _ = fi187
                                    _ = v188
                                }
                            default:
                                fmt.Println(v135, v6, v2)
                            }
                            if (0 <= 1000) {
                                v189 := GenS26{F0: -58, F1: -1}
                                var v190 int8 = 43
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                _ = v189
                                _ = v190
                            }
                            var v191 uint16 = 65535
                            _ = v191
                        }
                        fmt.Println(v2, v132, v1)
                    } else {
                        v2 = (^(v2))
                    }
                    v192 := GenS21{F0: 32767, F1: v2, F2: 2147483647}
                    var v193 float32 = (float32(1.5) - 0.1)
                    _ = v192
                    _ = v193
                } else {
                    switch (-(v135)) {
                    case -182:
                        switch "hello" {
                        case "":
                            for fi194 := 0; fi194 < 16; fi194++ {
                                v195 := GenS16{F0: 225}
                                fmt.Println(v135, fi194, v2)
                                _ = fi194
                                _ = v195
                            }
                        case "hello":
                            fmt.Println(v2)
                        default:
                            var v196 int32 = 2147483647
                            _ = v196
                        }
                    default:
                        v197 := GenS27{F0: 0}
                        _ = v197
                    }
                }
            }
        }
    }
    v198 := GenS5{F0: 4294967295, F1: 2147483647}
    v199 := GenS12{F0: v132, F1: 42, F2: 42}
    for fi200 := 0; fi200 < 19; fi200++ {
        _ = len(v16)
        v6 = v6
        _ = fi200
    }
    var v201 uint64 = v1
    var v202 uint = genFn0(144)
    var v203 uint = (v202 << (((v202 / ((65535) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v135:", v135)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v202:", v202)
    fmt.Println("v203:", v203)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v132
    _ = v135
    _ = v16
    _ = v17
    _ = v198
    _ = v199
    _ = v2
    _ = v201
    _ = v202
    _ = v203
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
