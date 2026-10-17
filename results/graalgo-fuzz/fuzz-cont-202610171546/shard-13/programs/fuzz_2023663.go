package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint32; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 bool; F2 int16 }
type GenS11 struct { F0 uint8; F1 int64 }
type GenS12 struct { F0 string; F1 uint32; F2 uint }
type GenS13 struct { F0 uint64; F1 int; F2 uint64 }
type GenS14 struct { F0 bool; F1 uint8 }
type GenS15 struct { F0 int8; F1 int8 }
type GenS16 struct { F0 uint; F1 int16; F2 uint8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16; F1 uint64 }
type GenS19 struct { F0 string; F1 int32 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int64; F1 string; F2 int8 }
type GenS22 struct { F0 uint; F1 int16 }
type GenS23 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS24 struct { F0 int8; F1 int }
type GenS25 struct { F0 int8; F1 int64; F2 uint16 }
type GenS26 struct { F0 int; F1 uint16 }
type GenS27 struct { F0 uint32 }
type GenS28 struct { F0 bool }
type GenS29 struct { F0 int; F1 uint8; F2 int64 }
type GenS30 struct { F0 uint64; F1 int8 }
type GenS31 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS32 struct { F0 uint16; F1 uint }
type GenS33 struct { F0 bool; F1 int8 }
type GenS34 struct { F0 uint }
type GenS35 struct { F0 int64 }
type GenS36 struct { F0 int }

func main() {
    var v1 uint32 = 65535
    var v2 int = 32768
    var v3 uint8 = 1
    v4 := []string{"hello", "hello"}
    var v5 int16 = -162
    v6 := make(map[int]bool)
    if false {
        if true {
            if (2 < v3) {
                fmt.Println(v2, v2, v2)
                v1 -= v1
                for fi7 := 0; fi7 < 7; fi7++ {
                    v2 += (-(fi7))
                    v8 := GenS1{F0: 0, F1: 0}
                    v2 = (v2 * 127)
                    _ = fi7
                    _ = v8
                }
            }
            v5 = (v5 << ((246) & 63))
            v9 := GenS1{F0: 42, F1: v5}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 11; fi10++ {
                if true {
                    v5 -= int16(v3)
                    if true {
                        var v11 string = ("\n\t" + "")
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v13 := GenS2{F0: 0, F1: (^(v1)), F2: (v5 >> ((2) & 63))}
                            var v14 uint8 = (^(v3))
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    }
                }
                _ = fi10
            }
        }
        switch int(v1) {
        case 32768:
            if true {
                if true {
                    v3 += v3
                    for fi15 := 0; fi15 < 1; fi15++ {
                        for fi16 := 0; fi16 < 8; fi16++ {
                            switch 82 {
                            case -438:
                                if len(v4) > 0 {
                                    v4[0] = "\n\t"
                                }
                            case 312:
                                switch "foo bar" {
                                case "a":
                                    switch v2 {
                                    case 9223372036854775807:
                                        v17 := GenS1{F0: uint16(v3), F1: v5}
                                        _ = v17
                                    case -733:
                                        v1 = (v1 ^ (v1 >> ((2) & 63)))
                                    case -1:
                                        v18 := GenS3{F0: ""}
                                        _ = v18
                                    }
                                default:
                                    var v19 uint = 628
                                    _ = v19
                                }
                            case 9223372036854775807:
                                if (127 != 128) {
                                    var v20 int8 = -7
                                    var v21 int8 = (v20 * v20)
                                    delete(v6, int(v5))
                                    _ = v20
                                    _ = v21
                                } else {
                                    v1 = (^(v1))
                                }
                            default:
                                var v22 float64 = (-2.25 * ((0.1 * ((0.1 - (1.5 / -2.25)) * 1.5)) / -2.25))
                                _ = v22
                            }
                            if len(v4) > 0 {
                                v4[0] = ((((("foo bar" + "foo bar") + "") + "\n\t") + "\n\t") + ("a" + "a"))
                            }
                            _ = fi16
                        }
                        _ = fi15
                    }
                    v2 -= v2
                } else {
                    if ((false || false) || (-815 != -857)) {
                        switch "hello" {
                        case "\n\t":
                            if false {
                                var v23 string = ("hello" + "")
                                v4 = append(v4, (((("" + "a") + (v23 + v23)) + (v23 + v23)) + v23))
                                fmt.Println(v2)
                                _ = v23
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                }
                v1 -= 821
                v2 -= 1
            }
        }
        var v24 int8 = 100
        _ = v24
    } else {
        if true {
            v2 = (v2 | v2)
        } else {
            var v25 uint8 = v3
            v2 = v2
            var v26 uint8 = (v25 - (^(v25)))
            _ = v25
            _ = v26
        }
        v6[v2] = true
        if false {
            v4 = append(v4, ("hello" + ("foo bar" + ("\n\t" + "\n\t"))))
            if ((v3 | (v3 % ((81) | 1))) > 255) {
                fmt.Println(v2, v2, v1)
            }
        } else {
            var v27 float32 = 0.0
            _ = v27
        }
    }
    var v28 uint8 = 248
    v5 = v5
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v29 uint = 553
    if (((0.1 * 0.1) * 1e10) <= 0.1) {
        v30 := GenS4{F0: v5, F1: 2, F2: -9223372036854775808}
        _ = v30
    } else {
        if true {
            if true {
                var v31 int32 = 1
                _ = v31
            } else {
                if ((!(((false || false) && (0 == 331)) || false)) && true) {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        v33 := GenS2{F0: (-(v5)), F1: v1, F2: (v5 + (v5 << ((65535) & 63)))}
                        for fi34 := 0; fi34 < 10; fi34++ {
                            if ((-1 <= 2147483647) && true) {
                                v3 = 255
                                v5 = -32768
                                fmt.Println(v2, fi32, v5)
                            } else {
                                if (42 != 0) {
                                    switch 9223372036854775807 {
                                    case 127:
                                        v35 := GenS4{F0: int16(v3), F1: (v1 * (v1 | (v1 + (v1 * (v1 >> ((v29) & 63)))))), F2: 1}
                                        _ = v35
                                    case -994:
                                        v36 := GenS2{F0: v5, F1: v1, F2: (v5 + v5)}
                                        v37 := GenS3{F0: ""}
                                        _ = v36
                                        _ = v37
                                    default:
                                        if false {
                                            v5 = (v5 & -32768)
                                            if false {
                                                v38 := GenS5{F0: 100}
                                                switch (v2 - (v2 / ((9223372036854775807) | 1))) {
                                                case -340:
                                                    var v39 uint8 = (^(v28))
                                                    v29 -= (v29 << ((308) & 63))
                                                    _ = v39
                                                }
                                                v40 := GenS2{F0: v5, F1: v1, F2: 42}
                                                _ = v38
                                                _ = v40
                                            } else {
                                                switch "a" {
                                                case "":
                                                    var v41 uint64 = 18446744073709551615
                                                    _ = v41
                                                }
                                            }
                                        } else {
                                            v42 := v6[v2]
                                            fmt.Println(v1, fi32, v42)
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            _ = v42
                                        }
                                        var v43 float32 = 0.0
                                        var v44 uint64 = 1
                                        fmt.Println(v3, fi32)
                                        if (("a" + "\n\t") <= "\n\t") {
                                            _ = len(v6)
                                            _ = len(v6)
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                if (0.0 > ((float32(((0.0 + 0.0) + -2.25)) / (v43 - float32(1.5))) / v43)) {
                                                    v5 -= (-(v5))
                                                    if true {
                                                        var v45 int16 = (v5 | v5)
                                                        var v46 uint8 = 207
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        v47 := GenS2{F0: v5, F1: (v1 & 0), F2: v5}
                                                        var v48 int8 = 0
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    var v49 uint16 = 1
                                                    _ = v49
                                                }
                                            case "a":
                                                v50 := GenS6{F0: 255, F1: "hello"}
                                                v29 = v29
                                                _ = v50
                                            case "":
                                                if true {
                                                    v51 := GenS3{F0: ("" + "")}
                                                    var v52 uint8 = v28
                                                    fmt.Println(v2, v29, v52)
                                                    _ = v51
                                                    _ = v52
                                                }
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    var v53 uint32 = 42
                                                    _ = v53
                                                case "\n\t":
                                                    fmt.Println(fi32, v28)
                                                default:
                                                    v54 := v6[fi34]
                                                    _ = v54
                                                }
                                            }
                                            var v55 uint = v29
                                            fmt.Println(v44)
                                            _ = v55
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                } else {
                                    v56 := GenS1{F0: 2, F1: v5}
                                    _ = v56
                                }
                                v57 := GenS6{F0: 1, F1: ""}
                                _ = v57
                            }
                            _ = fi34
                        }
                        _ = fi32
                        _ = v33
                    }
                } else {
                    switch "a" {
                    case "\n\t":
                        v5 = v5
                    case "":
                        if (-2147483648 == 42) {
                            fmt.Println(v2, v2, v29)
                        }
                    case "foo bar":
                        v5 = v5
                    }
                }
                for fi58 := 0; fi58 < 17; fi58++ {
                    v59 := GenS5{F0: 77}
                    fmt.Println(v2, v29, fi58)
                    for fi60 := 0; fi60 < 15; fi60++ {
                        switch 128 {
                        case 42:
                            if (v5 <= 32767) {
                                v61 := GenS7{F0: (-(v5)), F1: 127}
                                v5 -= (v5 * 1)
                                _ = v61
                            } else {
                                if false {
                                    switch int(v29) {
                                    case 569:
                                        var v62 uint32 = (^(v1))
                                        _ = v62
                                    case 2147483647:
                                        v63 := GenS8{F0: 563}
                                        _ = v63
                                    default:
                                        v3 += 128
                                    }
                                    var v64 string = ("" + ("a" + "\n\t"))
                                    if ((^(fi58)) == 1) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (3.14159 < (0.1 - (1e10 + 3.14159))) {
                                                if false {
                                                    var v65 float64 = (3.14159 - (0.1 + 0.0))
                                                    if true {
                                                        fmt.Println(fi58, fi58, fi60)
                                                        v3 = 0
                                                    }
                                                    _ = v65
                                                } else {
                                                    v29 = v29
                                                }
                                                if (856 != (v29 % (((v29 / ((v29) | 1))) | 1))) {
                                                    var v66 int16 = 0
                                                    if true {
                                                        switch fi60 {
                                                        case 127:
                                                            v5 = 42
                                                        }
                                                    }
                                                    if true {
                                                        if true {
                                                            v28 -= v3
                                                            if (4294967295 < (fi58 & fi58)) {
                                                                _ = len(v4)
                                                                fmt.Println(v29, v28)
                                                                v3 += v3
                                                            } else {
                                                                v67 := GenS9{F0: v5}
                                                                v5 = v5
                                                                var v68 int64 = 42
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        } else {
                                                            switch v64 {
                                                            case "\n\t":
                                                                if false {
                                                                    v69 := GenS1{F0: 32767, F1: (v5 << (((v29 & (v29 >> ((v29) & 63)))) & 63))}
                                                                    if true {
                                                                        switch v64 {
                                                                        case "a":
                                                                            var v70 float32 = (0.0 / -2.25)
                                                                            _ = v70
                                                                        case "":
                                                                            if true {
                                                                                if (false && true) {
                                                                                    var v71 string = (v64 + "hello")
                                                                                    _ = v71
                                                                                } else {
                                                                                    if (-2147483648 != 9223372036854775807) {
                                                                                        fmt.Println(v64)
                                                                                        switch -878 {
                                                                                        case -2147483648:
                                                                                            var v72 float32 = 0.0
                                                                                            _ = v72
                                                                                        case 4294967295:
                                                                                            var v73 int32 = -1
                                                                                            _ = v73
                                                                                        default:
                                                                                            v74 := GenS10{F0: 947, F1: false, F2: 90}
                                                                                            _ = v74
                                                                                        }
                                                                                        fmt.Println(v29, fi60)
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = v64
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v75 int = fi60
                                                                                var v76 string = v64
                                                                                _ = v75
                                                                                _ = v76
                                                                            } else {
                                                                                _ = len(v6)
                                                                                fmt.Println(v5, v29)
                                                                            }
                                                                        case "\n\t":
                                                                            v77 := GenS7{F0: v5, F1: -8}
                                                                            _ = v77
                                                                        default:
                                                                            switch v64 {
                                                                            case "":
                                                                                v3 = (v3 % (((v28 * (v3 & v3))) | 1))
                                                                            case "a":
                                                                                v1 = 42
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v69
                                                                }
                                                            case "hello":
                                                                v78 := GenS1{F0: 277, F1: (^(v5))}
                                                                _ = v78
                                                            }
                                                            if false {
                                                                var v79 int8 = -1
                                                                _ = len(v6)
                                                                _ = v79
                                                            } else {
                                                                if false {
                                                                    var v80 float64 = float64(1e10)
                                                                    v29 = (v29 << ((42) & 63))
                                                                    var v81 uint64 = 1
                                                                    v82 := GenS10{F0: 42, F1: false, F2: (v5 / ((int16(v28)) | 1))}
                                                                    v6[(v2 + fi60)] = true
                                                                    _ = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    var v83 string = ("hello" + "")
                                                                    var v84 bool = (v5 <= v5)
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                switch (v64 + (("foo bar" + v64) + v64)) {
                                                                case "hello":
                                                                    var v85 uint16 = 32767
                                                                    _ = v85
                                                                case "foo bar":
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        switch v64 {
                                                                        case "":
                                                                            switch v64 {
                                                                            case "a":
                                                                                fmt.Println(fi60, v5, fi58)
                                                                            case "":
                                                                                v86 := GenS11{F0: v3, F1: 1}
                                                                                v87 := GenS3{F0: ((((v64 + (v64 + v64)) + "") + ("foo bar" + (v64 + v64))) + v64)}
                                                                                _ = v86
                                                                                _ = v87
                                                                            case "\n\t":
                                                                                fmt.Println(v5)
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        switch v64 {
                                                                        case "":
                                                                            switch fi60 {
                                                                            case 127:
                                                                                v64 = (v64 + "\n\t")
                                                                            }
                                                                            var v88 uint8 = v28
                                                                            _ = v88
                                                                        case "foo bar":
                                                                            switch (fi58 - fi58) {
                                                                            case -9223372036854775808:
                                                                                v3 += v28
                                                                                if false {
                                                                                    if false {
                                                                                        if false {
                                                                                            v2 -= 32768
                                                                                            v89 := GenS12{F0: v64, F1: (v1 - (v1 - v1)), F2: v29}
                                                                                            _ = len(v4)
                                                                                            _ = v89
                                                                                        } else {
                                                                                            if (32767 == 876) {
                                                                                                v3 -= 100
                                                                                                if true {
                                                                                                    v3 = (v3 >> ((2) & 63))
                                                                                                } else {
                                                                                                    v90 := v6[fi58]
                                                                                                    switch -1 {
                                                                                                    case 2147483647:
                                                                                                        if (!(9223372036854775807 <= 0)) {
                                                                                                            fmt.Println(fi60, v64)
                                                                                                            v90 = (uint64(v3) >= 0)
                                                                                                        }
                                                                                                    case 127:
                                                                                                        v2 += (-(fi60))
                                                                                                    case 522:
                                                                                                        v28 += v3
                                                                                                    default:
                                                                                                        if v90 {
                                                                                                            v64 = "foo bar"
                                                                                                        }
                                                                                                    }
                                                                                                    v5 -= v66
                                                                                                    var v91 string = v64
                                                                                                    v1 += (v1 * (v1 >> (((v29 / ((v29) | 1))) & 63)))
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                }
                                                                                                var v92 bool = (32767 == uint16(v1))
                                                                                                var v93 float64 = 3.14159
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (!true) {
                                                                                        var v94 float64 = 0.1
                                                                                        var v95 uint = v29
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                                v66 = (-(v5))
                                                                            case -1:
                                                                                v96 := GenS7{F0: (v66 & (v66 >> ((696) & 63))), F1: -93}
                                                                                _ = v96
                                                                            case 32767:
                                                                                var v97 int64 = 0
                                                                                _ = v97
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(v64, v28, v2)
                                                                        default:
                                                                            fmt.Println(v28, fi60)
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v98 := GenS6{F0: 1, F1: v64}
                                                                    _ = v98
                                                                default:
                                                                    switch (v2 << (((v29 | 255)) & 63)) {
                                                                    case -920:
                                                                        v99 := GenS13{F0: 4294967295, F1: fi58, F2: 2}
                                                                        _ = v99
                                                                    case -2147483648:
                                                                        _ = len(v6)
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v28, v66)
                                                    } else {
                                                        if false {
                                                            v100 := GenS14{F0: (true && false), F1: (v28 * (v3 + (v28 + v28)))}
                                                            _ = v100
                                                        } else {
                                                            v1 -= (v1 / ((29) | 1))
                                                            var v101 uint8 = v28
                                                            var v102 uint64 = 4294967295
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                    }
                                                    var v103 uint = uint(v1)
                                                    fmt.Println(fi60)
                                                    fmt.Println(v29)
                                                    if true {
                                                        if true {
                                                            v29 -= (v29 << ((v103) & 63))
                                                        }
                                                        switch (v2 % ((v2) | 1)) {
                                                        case -522:
                                                            switch "" {
                                                            case "\n\t":
                                                                var v104 int = (-(v2))
                                                                _ = v104
                                                            case "":
                                                                if true {
                                                                    var v105 string = "a"
                                                                    fmt.Println(v3)
                                                                    _ = v105
                                                                } else {
                                                                    switch fi58 {
                                                                    case -1:
                                                                        delete(v6, -9223372036854775808)
                                                                    case -946:
                                                                        if (-768 != -1) {
                                                                            if (!(-1 >= (v66 + v5))) {
                                                                                v106 := GenS7{F0: v5, F1: 0}
                                                                                switch 32768 {
                                                                                case 32767:
                                                                                    fmt.Println(v28, v1)
                                                                                case 42:
                                                                                    if false {
                                                                                        if (((1e10 + 1.5) * 0.1) != 0.1) {
                                                                                            v29 += (v103 ^ v103)
                                                                                        }
                                                                                        v28 = (v28 << ((v29) & 63))
                                                                                    } else {
                                                                                        var v107 uint64 = uint64(v1)
                                                                                        v108 := GenS15{F0: 42, F1: 42}
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                                _ = v106
                                                                            }
                                                                            var v109 int32 = int32(v1)
                                                                            fmt.Println(v64, v29)
                                                                            _ = v109
                                                                        } else {
                                                                            delete(v6, int(v29))
                                                                        }
                                                                    }
                                                                    v110 := GenS14{F0: true, F1: v28}
                                                                    var v111 string = v64
                                                                    if true {
                                                                        v111 = "foo bar"
                                                                    }
                                                                    v1 = (v1 ^ v1)
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v5, v103, v1)
                                                            }
                                                        case 1:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        if (!false) {
                                                            switch 273 {
                                                            case 938:
                                                                var v112 int32 = -1
                                                                _ = v112
                                                            }
                                                            if true {
                                                                switch -961 {
                                                                case 42:
                                                                    if ((-(fi60)) < int(v1)) {
                                                                        var v113 uint16 = 780
                                                                        _ = v113
                                                                    } else {
                                                                        if (1 >= uint64(v1)) {
                                                                            var v114 float32 = 3.14159
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                case 2147483647:
                                                                    var v115 float32 = (1.5 * 1.5)
                                                                    _ = v115
                                                                case 1:
                                                                    v29 += v29
                                                                default:
                                                                    fmt.Println(v28)
                                                                }
                                                                v2 = v2
                                                                fmt.Println(v2)
                                                            } else {
                                                                v29 += (v29 + 1)
                                                                if true {
                                                                    var v116 float64 = (-2.25 * 1.5)
                                                                    _ = v116
                                                                } else {
                                                                    var v117 bool = (950 > 687)
                                                                    _ = v117
                                                                }
                                                            }
                                                            if false {
                                                                if true {
                                                                    switch fi60 {
                                                                    case 0:
                                                                        v118 := GenS2{F0: (^(v5)), F1: 2147483647, F2: v5}
                                                                        var v119 bool = (fi60 <= fi60)
                                                                        v103 = uint(v28)
                                                                        _ = v118
                                                                        _ = v119
                                                                    case 312:
                                                                        v3 = v3
                                                                    case 127:
                                                                        var v120 int64 = 577
                                                                        _ = v120
                                                                    }
                                                                }
                                                                fmt.Println(v29, v103)
                                                                var v121 uint16 = 255
                                                                _ = v121
                                                            }
                                                        }
                                                    }
                                                    var v122 int32 = 953
                                                    v103 -= (^(v103))
                                                    _ = v103
                                                    _ = v122
                                                    _ = v66
                                                } else {
                                                    var v123 uint = (^(v29))
                                                    var v124 float32 = (((float32(3.14159) - 0.0) - 3.14159) * 0.1)
                                                    _ = v123
                                                    _ = v124
                                                }
                                                v29 = v29
                                            } else {
                                                v28 += 128
                                                if false {
                                                    var v125 int8 = 127
                                                    v126 := GenS3{F0: (v64 + (v64 + v64))}
                                                    _ = v125
                                                    _ = v126
                                                }
                                            }
                                            v127 := GenS9{F0: (v5 % ((v5) | 1))}
                                            v128 := GenS3{F0: v64}
                                            _ = v127
                                            _ = v128
                                        case "foo bar":
                                            v3 = 143
                                        case "a":
                                            var v129 uint32 = 4294967295
                                            _ = v129
                                        default:
                                            if false {
                                                var v130 uint = uint(v3)
                                                if false {
                                                    if (127 >= -128) {
                                                        if true {
                                                            var v131 uint64 = uint64(v28)
                                                            _ = v131
                                                        } else {
                                                            var v132 uint64 = 1
                                                            if (!false) {
                                                                if (9223372036854775807 >= 0) {
                                                                    v133 := GenS16{F0: v130, F1: 42, F2: 172}
                                                                    v28 = (v28 | (^(v28)))
                                                                    if false {
                                                                        var v134 float32 = 0.0
                                                                        v135 := v6[fi58]
                                                                        v2 -= (v2 + (-(v2)))
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                    _ = v133
                                                                } else {
                                                                    if (!(!true)) {
                                                                        if (!(uint32(v3) <= v1)) {
                                                                            fmt.Println(v1, v2)
                                                                            v5 = (v5 * 32767)
                                                                            v136 := GenS1{F0: 0, F1: v5}
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                    v3 += (v28 - (v3 | (^(v28))))
                                                                    var v137 string = v64
                                                                    _ = v137
                                                                }
                                                                fmt.Println(fi58, v2)
                                                                var v138 int8 = 32
                                                                _ = v138
                                                            } else {
                                                                v139 := GenS17{F0: 127}
                                                                v140 := GenS18{F0: v5, F1: 415}
                                                                switch (^(fi58)) {
                                                                case 127:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v1 -= 2147483647
                                                                    case "foo bar":
                                                                        v141 := GenS19{F0: v64, F1: 680}
                                                                        _ = v141
                                                                    case "\n\t":
                                                                        v142 := GenS3{F0: "hello"}
                                                                        _ = v142
                                                                    }
                                                                }
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                            var v143 uint32 = (v1 << ((v130) & 63))
                                                            v1 = (v143 & v1)
                                                            _ = v132
                                                            _ = v143
                                                        }
                                                        v144 := GenS15{F0: 46, F1: 42}
                                                        _ = v144
                                                    } else {
                                                        v5 = v5
                                                    }
                                                    var v145 uint64 = 1
                                                    _ = v145
                                                } else {
                                                    var v146 int32 = -2147483648
                                                    if true {
                                                        var v147 string = (v64 + v64)
                                                        _ = v147
                                                    } else {
                                                        v1 -= v1
                                                        v64 = "\n\t"
                                                        switch fi60 {
                                                        case -1:
                                                            if false {
                                                                fmt.Println(v29, v2, v1)
                                                            }
                                                            v64 = (v64 + "hello")
                                                        case 1:
                                                            v148 := GenS20{F0: 42}
                                                            _ = v148
                                                        case 745:
                                                            v149 := GenS21{F0: 49, F1: "hello", F2: -42}
                                                            _ = v149
                                                        default:
                                                            v150 := GenS1{F0: 32767, F1: (-(v5))}
                                                            _ = v150
                                                        }
                                                        var v151 uint16 = 1
                                                        v152 := GenS15{F0: 105, F1: -1}
                                                        _ = v151
                                                        _ = v152
                                                    }
                                                    _ = v146
                                                }
                                                _ = v130
                                            } else {
                                                var v153 int32 = 280
                                                _ = v153
                                            }
                                            v2 = (-(fi58))
                                        }
                                        var v154 int8 = 35
                                        var v155 uint64 = uint64(v1)
                                        _ = v154
                                        _ = v155
                                    } else {
                                        var v156 int16 = 904
                                        fmt.Println(v2)
                                        switch (("foo bar" + "hello") + (v64 + "hello")) {
                                        case "foo bar":
                                            var v157 int64 = -9223372036854775808
                                            v158 := GenS1{F0: 1, F1: v156}
                                            if false {
                                                v4 = append(v4, "a")
                                            } else {
                                                var v159 int = (-(fi58))
                                                v160 := GenS22{F0: v29, F1: 850}
                                                var v161 int64 = v157
                                                _ = v159
                                                _ = v160
                                                _ = v161
                                            }
                                            if (1 != -1) {
                                                if (v157 != -1) {
                                                    switch fi60 {
                                                    case 127:
                                                        v156 -= v5
                                                    case -1:
                                                        var v162 string = (v64 + "\n\t")
                                                        _ = v162
                                                    case 792:
                                                        v157 = int64(v28)
                                                    default:
                                                        var v163 int8 = -109
                                                        _ = v163
                                                    }
                                                    var v164 int64 = v157
                                                    _ = v164
                                                }
                                            } else {
                                                var v165 int32 = -2147483648
                                                _ = v165
                                            }
                                            fmt.Println(fi60)
                                            _ = v157
                                            _ = v158
                                        default:
                                            switch v64 {
                                            case "\n\t":
                                                switch (-(fi60)) {
                                                case 127:
                                                    if (997 < 18446744073709551615) {
                                                        var v166 bool = (v29 <= uint(v1))
                                                        v156 = (v5 / ((int16(v29)) | 1))
                                                        if (837 <= 4294967295) {
                                                            switch fi58 {
                                                            case -873:
                                                                fmt.Println(fi60, fi58)
                                                            case 312:
                                                                v156 -= v5
                                                            default:
                                                                if len(v4) > 0 {
                                                                    v4[0] = v64
                                                                }
                                                            }
                                                            v6[-343] = ((0 > -385) || (-9223372036854775808 <= 172))
                                                            v166 = (9223372036854775807 > 605)
                                                        } else {
                                                            delete(v6, (fi58 << ((1) & 63)))
                                                            switch ((v64 + "hello") + ((v64 + v64) + v64)) {
                                                            case "foo bar":
                                                                fmt.Println(v3, fi58, fi60)
                                                            default:
                                                                var v167 uint = v29
                                                                if ((255 != 32767) && (127 < -63)) {
                                                                    v64 = v64
                                                                    fmt.Println(fi58, v5, fi60)
                                                                    var v168 uint64 = 4294967295
                                                                    var v169 int32 = 1
                                                                    _ = v168
                                                                    _ = v169
                                                                } else {
                                                                    fmt.Println(v167)
                                                                    var v170 int = v2
                                                                    v171 := v6[(-(fi60))]
                                                                    if (4294967295 <= v1) {
                                                                        var v172 int = (fi58 / ((fi58) | 1))
                                                                        _ = v172
                                                                    }
                                                                    v28 = v3
                                                                    _ = v170
                                                                    _ = v171
                                                                }
                                                                switch v64 {
                                                                case "hello":
                                                                    var v173 uint = 588
                                                                    _ = v173
                                                                case "a":
                                                                    v5 = (v156 * 42)
                                                                }
                                                                _ = v167
                                                            }
                                                        }
                                                        _ = v166
                                                    }
                                                case -9223372036854775808:
                                                    if ((("a" + v64) != (v64 + (((v64 + "hello") + "a") + v64))) || false) {
                                                        if (!true) {
                                                            switch 127 {
                                                            case 40:
                                                                v174 := GenS22{F0: v29, F1: -1}
                                                                fmt.Println(v3, v3, fi58)
                                                                var v175 int16 = 1000
                                                                v6[-1] = (v156 > (v175 / ((-1) | 1)))
                                                                _ = v174
                                                                _ = v175
                                                            case 32768:
                                                                var v176 int32 = 2147483647
                                                                _ = v176
                                                            default:
                                                                if (false || false) {
                                                                    var v177 int = int(v29)
                                                                    fmt.Println(v29, v29, fi60)
                                                                    v1 = (v1 & v1)
                                                                    _ = v177
                                                                }
                                                            }
                                                            switch (fi58 >> (((v29 << (((v29 | 42)) & 63))) & 63)) {
                                                            case 230:
                                                                var v178 int = fi58
                                                                _ = v178
                                                            case 287:
                                                                v179 := v6[fi60]
                                                                _ = v179
                                                            }
                                                            var v180 uint8 = (v28 % ((v3) | 1))
                                                            _ = v180
                                                        } else {
                                                            switch (v64 + "a") {
                                                            case "a":
                                                                var v181 int32 = 317
                                                                _ = v181
                                                            case "":
                                                                v3 = (^(v3))
                                                            default:
                                                                var v182 uint64 = 0
                                                                _ = v182
                                                            }
                                                            v4 = append(v4, ((v64 + v64) + ((("foo bar" + "a") + v64) + v64)))
                                                            v28 -= v3
                                                        }
                                                    }
                                                case 42:
                                                    if (!true) {
                                                        if (true && (552 >= 1)) {
                                                            var v183 int16 = v156
                                                            if (((65 <= 42) || true) || false) {
                                                                v6[fi58] = (!true)
                                                                if (v1 <= 65535) {
                                                                    v64 = ("foo bar" + "\n\t")
                                                                    var v184 int64 = -9223372036854775808
                                                                    fmt.Println(v2, v2, v3)
                                                                    v156 -= v156
                                                                    fmt.Println(v5, v2)
                                                                    _ = v184
                                                                }
                                                                if (!false) {
                                                                    v156 += -32768
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if false {
                                                                        if (v1 < uint32(v28)) {
                                                                            v3 = (v28 << ((v29) & 63))
                                                                            var v185 uint32 = 184
                                                                            v186 := GenS2{F0: int16(v29), F1: v1, F2: v183}
                                                                            v28 = 0
                                                                            fmt.Println(v29, v28, v3)
                                                                            switch fi58 {
                                                                            case 4294967295:
                                                                                v4 = append(v4, (v64 + "foo bar"))
                                                                            }
                                                                            _ = v185
                                                                            _ = v186
                                                                        } else {
                                                                            var v187 uint64 = 0
                                                                            if false {
                                                                                var v188 uint16 = 255
                                                                                _ = v188
                                                                            }
                                                                            switch v2 {
                                                                            case 1:
                                                                                v187 -= (v187 << (((v29 >> (((v29 >> (((v29 << ((v29) & 63))) & 63))) & 63))) & 63))
                                                                            case 2147483647:
                                                                                fmt.Println(v5, v156, v64)
                                                                            default:
                                                                                if true {
                                                                                    v183 = v156
                                                                                }
                                                                            }
                                                                            _ = v187
                                                                        }
                                                                        var v189 int32 = -684
                                                                        _ = v189
                                                                    } else {
                                                                        v190 := GenS17{F0: v3}
                                                                        v1 = (v1 + (v1 - (v1 << (((v29 % (((v29 % (((v29 - 0)) | 1))) | 1))) & 63))))
                                                                        var v191 int16 = (v5 * 42)
                                                                        if true {
                                                                            var v192 string = v64
                                                                            fmt.Println(fi60, v1, v183)
                                                                            var v193 uint64 = 1
                                                                            _ = v192
                                                                            _ = v193
                                                                        } else {
                                                                            fmt.Println(v156, fi58)
                                                                            switch "a" {
                                                                            case "":
                                                                                v3 = 1
                                                                            case "a":
                                                                                v4 = append(v4, v64)
                                                                            case "\n\t":
                                                                                var v194 string = "\n\t"
                                                                                fmt.Println(v29, v3)
                                                                                v195 := GenS23{F0: (v28 / ((2) | 1)), F1: 2, F2: (v1 | v1)}
                                                                                _ = v194
                                                                                _ = v195
                                                                            }
                                                                        }
                                                                        if false {
                                                                            v196 := GenS24{F0: 0, F1: (v2 + (fi60 / ((236) | 1)))}
                                                                            if false {
                                                                                var v197 uint8 = (v28 ^ (v3 << ((v29) & 63)))
                                                                                _ = v197
                                                                            }
                                                                            switch "" {
                                                                            case "":
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = "\n\t"
                                                                                }
                                                                                v4 = append(v4, v64)
                                                                            case "\n\t":
                                                                                fmt.Println(v3, v191)
                                                                            default:
                                                                                var v198 int64 = -886
                                                                                _ = v198
                                                                            }
                                                                            _ = v196
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v183, v191)
                                                                            v199 := GenS25{F0: 1, F1: int64(fi60), F2: 2}
                                                                            _ = v199
                                                                        } else {
                                                                            fmt.Println(v29)
                                                                            v200 := v6[-267]
                                                                            _ = v200
                                                                        }
                                                                        _ = v190
                                                                        _ = v191
                                                                    }
                                                                } else {
                                                                    fmt.Println(v1, v156)
                                                                }
                                                            }
                                                            if false {
                                                                var v201 int = (fi58 & 746)
                                                                v202 := GenS11{F0: (v3 - 1), F1: 9223372036854775807}
                                                                v203 := GenS26{F0: fi60, F1: 32767}
                                                                _ = v201
                                                                _ = v202
                                                                _ = v203
                                                            }
                                                            _ = v183
                                                        }
                                                        switch (v2 + (fi58 | fi58)) {
                                                        case 32767:
                                                            var v204 string = v64
                                                            _ = v204
                                                        case 517:
                                                            delete(v6, 616)
                                                        default:
                                                            var v205 bool = (-2.25 == (0.1 + (3.14159 - 3.14159)))
                                                            _ = v205
                                                        }
                                                        var v206 uint64 = 2
                                                        _ = v206
                                                    } else {
                                                        if ((float32(float64((float32(3.14159) - float32((3.14159 / 3.14159))))) <= (-2.25 + 0.1)) || true) {
                                                            v207 := GenS18{F0: (v5 << ((42) & 63)), F1: 958}
                                                            fmt.Println(v3, fi60)
                                                            _ = v207
                                                        } else {
                                                            switch v64 {
                                                            case "hello":
                                                                if ((!false) && (851 != 895)) {
                                                                    fmt.Println(v64, v156, v3)
                                                                    delete(v6, fi58)
                                                                    var v208 uint8 = (v3 & v28)
                                                                    _ = v208
                                                                } else {
                                                                    switch v2 {
                                                                    case 65:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v64
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if (false && (42 < 1)) {
                                                                            fmt.Println(fi58)
                                                                            _ = len(v4)
                                                                        }
                                                                    case -1:
                                                                        v209 := GenS18{F0: 453, F1: 1}
                                                                        _ = v209
                                                                    default:
                                                                        v210 := GenS3{F0: v64}
                                                                        _ = v210
                                                                    }
                                                                    if (907 == 1) {
                                                                        if false {
                                                                            if (42 <= -2147483648) {
                                                                                fmt.Println(v29, v5)
                                                                            }
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                v1 += 42
                                                                            case "hello":
                                                                                var v211 int16 = (-(v156))
                                                                                _ = v211
                                                                            }
                                                                            fmt.Println(v28, v1, v29)
                                                                            if (!false) {
                                                                                var v212 uint16 = uint16(v29)
                                                                                if (4294967295 >= 42) {
                                                                                    if false {
                                                                                        var v213 int = v2
                                                                                        _ = v213
                                                                                    }
                                                                                    v29 -= (v29 << (((v29 - (v29 - uint(v3)))) & 63))
                                                                                } else {
                                                                                    var v214 int32 = 1
                                                                                    switch (-(fi60)) {
                                                                                    case 1:
                                                                                        var v215 bool = (true && true)
                                                                                        _ = v215
                                                                                    }
                                                                                    switch v64 {
                                                                                    case "a":
                                                                                        fmt.Println(fi60)
                                                                                    case "hello":
                                                                                        v156 -= -73
                                                                                        fmt.Println(v28)
                                                                                        if false {
                                                                                            if (false && false) {
                                                                                                if true {
                                                                                                    if true {
                                                                                                        v214 += (v214 | -565)
                                                                                                        v29 += (v29 ^ v29)
                                                                                                    } else {
                                                                                                        var v216 int64 = 1
                                                                                                        var v217 int32 = v214
                                                                                                        if (76 == 100) {
                                                                                                            var v218 float64 = (-2.25 * float64((-2.25 * 0.1)))
                                                                                                            _ = v218
                                                                                                        } else {
                                                                                                            switch (v64 + "") {
                                                                                                            case "hello":
                                                                                                                v2 = (fi58 % (((fi60 & (v2 & (v2 / (((v2 | v2)) | 1))))) | 1))
                                                                                                            case "foo bar":
                                                                                                                _ = len(v4)
                                                                                                                var v219 bool = (v64 <= v64)
                                                                                                                _ = v219
                                                                                                            case "":
                                                                                                                var v220 int8 = 14
                                                                                                                _ = v220
                                                                                                            default:
                                                                                                                _ = len(v6)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v216
                                                                                                        _ = v217
                                                                                                    }
                                                                                                    fmt.Println(v3)
                                                                                                    switch v2 {
                                                                                                    case -203:
                                                                                                        if (true && true) {
                                                                                                            var v221 int8 = 42
                                                                                                            _ = v221
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v5 -= (v5 >> ((v29) & 63))
                                                                                                                if (511 == (v29 ^ (^(v29)))) {
                                                                                                                    fmt.Println(fi60, fi60, v2)
                                                                                                                    var v222 float32 = -2.25
                                                                                                                    fmt.Println(v212)
                                                                                                                    v223 := GenS27{F0: 699}
                                                                                                                    v1 -= v1
                                                                                                                    _ = v222
                                                                                                                    _ = v223
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v28 = (v28 & (v28 >> (((v29 + 4294967295)) & 63)))
                                                                                                                        v3 = v3
                                                                                                                    }
                                                                                                                    fmt.Println(v212, v214, v3)
                                                                                                                    var v224 uint16 = (v212 - (v212 & v212))
                                                                                                                    _ = v224
                                                                                                                }
                                                                                                                v225 := GenS28{F0: (v212 < 32767)}
                                                                                                                _ = v225
                                                                                                            }
                                                                                                        }
                                                                                                    case -531:
                                                                                                        fmt.Println(v214, v212)
                                                                                                    case -9223372036854775808:
                                                                                                        v226 := GenS29{F0: 42, F1: v28, F2: -9223372036854775808}
                                                                                                        _ = v226
                                                                                                    default:
                                                                                                        var v227 uint32 = v1
                                                                                                        _ = v227
                                                                                                    }
                                                                                                } else {
                                                                                                    var v228 uint32 = 2147483647
                                                                                                    v156 = 1000
                                                                                                    var v229 int64 = 9223372036854775807
                                                                                                    _ = v228
                                                                                                    _ = v229
                                                                                                }
                                                                                                if (3.14159 >= 1e10) {
                                                                                                    fmt.Println(v64, v2, v3)
                                                                                                    v230 := GenS15{F0: 1, F1: 72}
                                                                                                    _ = v230
                                                                                                } else {
                                                                                                    v156 = (v5 << ((v29) & 63))
                                                                                                    delete(v6, int(v1))
                                                                                                }
                                                                                                _ = len(v4)
                                                                                            }
                                                                                            var v231 int64 = -619
                                                                                            _ = v231
                                                                                        }
                                                                                    default:
                                                                                        var v232 int64 = 1
                                                                                        _ = v232
                                                                                    }
                                                                                    _ = v214
                                                                                }
                                                                                _ = v212
                                                                            } else {
                                                                                v6[128] = false
                                                                            }
                                                                            switch (v64 + "") {
                                                                            case "hello":
                                                                                if (-32768 < -1) {
                                                                                    v233 := GenS30{F0: 1, F1: 42}
                                                                                    if true {
                                                                                        if (true && false) {
                                                                                            var v234 uint64 = 2
                                                                                            var v235 uint64 = uint64(v3)
                                                                                            switch -98 {
                                                                                            case 70:
                                                                                                switch int(v29) {
                                                                                                case -837:
                                                                                                    v236 := GenS31{F0: 42, F1: -32768, F2: 42}
                                                                                                    _ = v236
                                                                                                case 1:
                                                                                                    switch v64 {
                                                                                                    case "":
                                                                                                        if true {
                                                                                                            v2 += (-(fi60))
                                                                                                            switch v64 {
                                                                                                            case "foo bar":
                                                                                                                var v237 int32 = 42
                                                                                                                _ = v237
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        _ = len(v4)
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v238 := GenS21{F0: -1, F1: v64, F2: -115}
                                                                                                _ = v238
                                                                                            }
                                                                                            _ = v234
                                                                                            _ = v235
                                                                                        }
                                                                                        fmt.Println(v2, v1, v64)
                                                                                    } else {
                                                                                        var v239 int16 = 1000
                                                                                        switch v64 {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v239, v3)
                                                                                        case "":
                                                                                            switch v2 {
                                                                                            case -860:
                                                                                                if (!(!true)) {
                                                                                                    var v240 float64 = (1.5 + 0.0)
                                                                                                    v29 -= v29
                                                                                                    _ = v240
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if ((v28 & v3) != v3) {
                                                                                            v241 := GenS4{F0: 587, F1: 4294967295, F2: 199}
                                                                                            var v242 uint = v29
                                                                                            _ = v241
                                                                                            _ = v242
                                                                                        } else {
                                                                                            var v243 string = (v64 + v64)
                                                                                            _ = v243
                                                                                        }
                                                                                        _ = v239
                                                                                    }
                                                                                    _ = v233
                                                                                }
                                                                            default:
                                                                                v3 += 141
                                                                            }
                                                                        }
                                                                        switch (fi60 / ((fi58) | 1)) {
                                                                        case 127:
                                                                            fmt.Println(v5)
                                                                        case -38:
                                                                            v244 := GenS6{F0: 552, F1: ("" + "hello")}
                                                                            _ = v244
                                                                        }
                                                                        v245 := GenS5{F0: 100}
                                                                        v246 := GenS23{F0: v3, F1: 42, F2: v1}
                                                                        v156 = (v156 % (((v5 | v5)) | 1))
                                                                        _ = v245
                                                                        _ = v246
                                                                    }
                                                                    fmt.Println(v1, v156, v2)
                                                                    v3 = v28
                                                                }
                                                            default:
                                                                var v247 int32 = 0
                                                                _ = v247
                                                            }
                                                            v248 := GenS32{F0: uint16(v29), F1: v29}
                                                            var v249 uint16 = 973
                                                            _ = v248
                                                            _ = v249
                                                        }
                                                        if (!true) {
                                                            if ((^(v1)) == 2147483648) {
                                                                v250 := GenS1{F0: 42, F1: v156}
                                                                var v251 uint32 = (^(v1))
                                                                if ((v29 - (v29 - v29)) > v29) {
                                                                    delete(v6, fi58)
                                                                    fmt.Println(v2, v3, v64)
                                                                    v29 = v29
                                                                    if (-388 != -534) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v252 float64 = 3.14159
                                                                            _ = v252
                                                                        }
                                                                    } else {
                                                                        if (!true) {
                                                                            v253 := GenS4{F0: (-(v156)), F1: (v1 - v1), F2: 9223372036854775807}
                                                                            _ = v253
                                                                        } else {
                                                                            v254 := GenS26{F0: -9223372036854775808, F1: 42}
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (v64 + v64)
                                                                            }
                                                                            switch v2 {
                                                                            case 1:
                                                                                var v255 uint = 18446744073709551615
                                                                                _ = v255
                                                                            case -2147483648:
                                                                                if (42 > 9223372036854775807) {
                                                                                    var v256 uint8 = v28
                                                                                    switch -786 {
                                                                                    case 32767:
                                                                                        v257 := GenS33{F0: (0.0 <= 3.14159), F1: 1}
                                                                                        _ = v257
                                                                                    case 0:
                                                                                        v2 = fi58
                                                                                    default:
                                                                                        switch (fi58 % ((v2) | 1)) {
                                                                                        case -9223372036854775808:
                                                                                            if false {
                                                                                                var v258 uint8 = v3
                                                                                                if (-2.25 < (((float64(float32(0.0)) + 0.0) - 0.1) * (1e10 - 0.1))) {
                                                                                                    v258 += 2
                                                                                                    var v259 uint32 = v251
                                                                                                    v258 -= 194
                                                                                                    v258 -= uint8(v1)
                                                                                                    _ = v259
                                                                                                }
                                                                                                v260 := v6[fi58]
                                                                                                _ = v258
                                                                                                _ = v260
                                                                                            }
                                                                                        default:
                                                                                            v261 := GenS34{F0: (^(v29))}
                                                                                            _ = v261
                                                                                        }
                                                                                    }
                                                                                    _ = v256
                                                                                } else {
                                                                                    fmt.Println(v251, v251)
                                                                                    v251 += (v251 >> ((v29) & 63))
                                                                                    fmt.Println(v251)
                                                                                }
                                                                            default:
                                                                                v156 = -32768
                                                                            }
                                                                            _ = v254
                                                                        }
                                                                    }
                                                                }
                                                                _ = v250
                                                                _ = v251
                                                            } else {
                                                                var v262 uint8 = (v3 & v28)
                                                                v156 += (^(v5))
                                                                fmt.Println(v28, fi60, v28)
                                                                _ = v262
                                                            }
                                                            var v263 uint8 = 127
                                                            _ = v263
                                                        }
                                                        var v264 int32 = -299
                                                        _ = v264
                                                    }
                                                }
                                            case "hello":
                                                v1 = 0
                                            default:
                                                v28 -= v28
                                            }
                                        }
                                        v265 := GenS19{F0: v64, F1: -164}
                                        var v266 uint32 = 896
                                        _ = v156
                                        _ = v265
                                        _ = v266
                                    }
                                    if (true && (((1e10 * float32(1.5)) >= 3.14159) || true)) {
                                        v267 := GenS35{F0: -1}
                                        if (-9223372036854775808 >= 43) {
                                            switch v2 {
                                            case 17:
                                                v6[fi60] = true
                                            case 128:
                                                v2 += (-(v2))
                                            }
                                            var v268 int64 = -180
                                            _ = v268
                                        }
                                        v269 := GenS34{F0: (v29 % (((v29 / (((v29 << (((v29 | 4294967295)) & 63))) | 1))) | 1))}
                                        _ = v267
                                        _ = v269
                                    }
                                    _ = v64
                                }
                                if false {
                                    v270 := GenS25{F0: 0, F1: 1, F2: 32767}
                                    if (!(!(!(!true)))) {
                                        var v271 int8 = 73
                                        v2 = fi58
                                        v3 = (v28 * 93)
                                        v272 := GenS2{F0: v5, F1: v1, F2: (^(v5))}
                                        _ = v271
                                        _ = v272
                                    } else {
                                        var v273 uint8 = v28
                                        fmt.Println(fi60, v273)
                                        switch int(v29) {
                                        case 32767:
                                            v273 = 1
                                        case 127:
                                            v274 := GenS6{F0: uint16(v1), F1: "a"}
                                            _ = v274
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = ("" + "")
                                            }
                                        }
                                        _ = len(v6)
                                        _ = v273
                                    }
                                    v2 -= 4294967295
                                    _ = v270
                                }
                            }
                        default:
                            v4 = append(v4, ("a" + ((("hello" + "\n\t") + (("a" + ("hello" + "foo bar")) + "a")) + "\n\t")))
                        }
                        _ = fi60
                    }
                    _ = fi58
                    _ = v59
                }
            }
        }
        v3 = (v28 << ((v29) & 63))
    }
    if false {
        v275 := GenS36{F0: -490}
        if true {
            v276 := v6[v2]
            var v277 uint32 = (v1 ^ v1)
            var v278 int8 = 127
            _ = v276
            _ = v277
            _ = v278
        } else {
            var v279 string = ("\n\t" + "a")
            v5 += (v5 >> ((v29) & 63))
            if false {
                if len(v4) > 0 {
                    v4[0] = v279
                }
                v1 = 0
                for fi280 := 0; fi280 < 18; fi280++ {
                    delete(v6, 0)
                    v279 = v279
                    var v281 int64 = 9223372036854775807
                    _ = fi280
                    _ = v281
                }
            }
            _ = v279
        }
        _ = v275
    } else {
        for fi282 := 0; fi282 < 6; fi282++ {
            fmt.Println(v1, v5)
            _ = fi282
        }
        fmt.Println(v5, v28)
    }
    var v283 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v283:", v283)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v283
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
