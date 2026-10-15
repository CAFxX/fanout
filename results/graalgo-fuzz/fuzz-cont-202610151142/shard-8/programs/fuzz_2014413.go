package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 string; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint8; F1 int32; F2 int64 }
type GenS10 struct { F0 uint32; F1 int8; F2 bool }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint8; F1 int32 }
type GenS15 struct { F0 uint; F1 int32; F2 bool }
type GenS16 struct { F0 uint8; F1 int8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int64; F1 int16; F2 uint16 }
type GenS19 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint8; F1 string }
type GenS23 struct { F0 uint32; F1 int; F2 int8 }
type GenS24 struct { F0 uint64; F1 int16 }
type GenS25 struct { F0 int64; F1 int8; F2 uint64 }
type GenS26 struct { F0 int32; F1 uint32 }
type GenS27 struct { F0 int8; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (882 != 0)
    var v3 uint = 18446744073709551615
    var v4 uint16 = 671
    var v5 float64 = 0.0
    if v2 {
        var v6 float32 = 0.0
        var v7 bool = false
        var v8 uint8 = 77
        _ = v6
        _ = v7
        _ = v8
    } else {
        switch "a" {
        case "\n\t":
            v9 := GenS1{F0: 0, F1: v2, F2: 42}
            _ = v9
        case "hello":
            switch 1 {
            case 2147483647:
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 int16 = -32768
                    v12 := GenS1{F0: 2147483647, F1: v2, F2: 1}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case 42:
                v3 += (v3 >> (((^(v3))) & 63))
            default:
                if true {
                    v13 := []int{}
                    v14 := GenS1{F0: -854, F1: false, F2: 42}
                    _ = v13
                    _ = v14
                } else {
                    v2 = (!v2)
                }
            }
        }
        v15 := v1[uint32(v3)]
        _ = v15
    }
    if (v2 && v2) {
        v2 = (((v5 >= v5) && v2) || v2)
    } else {
        var v16 int8 = 1
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS1{F0: 42, F1: (v2 && ((fi17 << ((v3) & 63)) <= fi17)), F2: -9223372036854775808}
            v19 := []int{fi17, 9223372036854775807}
            var v20 uint = 0
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int = 2147483647
        _ = v16
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 uint64 = 82
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 3; fi25++ {
            if true {
                v4 += uint16(v3)
                v26 := GenS1{F0: -175, F1: (("foo bar" + "foo bar") > "\n\t"), F2: 0}
                _ = v26
            } else {
                var v27 uint8 = 0
                _ = v27
            }
            v28 := []int{}
            switch 946 {
            case -543:
                if v2 {
                    switch ("" + "") {
                    case "a":
                        if true {
                            var v29 uint = (v3 ^ v3)
                            _ = v29
                        }
                    case "":
                        if ((1.5 - v5) <= (1.5 / -2.25)) {
                            var v30 float32 = float32(v5)
                            v2 = (((v4 & (v4 % ((200) | 1))) > v4) || v2)
                            v30 = (v30 / float32((v5 + float64((v30 + float32(v5))))))
                            _ = v30
                        } else {
                            if false {
                                var v31 bool = false
                                var v32 int32 = -256
                                if (1.5 != (1.5 + (((1.5 - 3.14159) * 1.5) * -2.25))) {
                                    fmt.Println(v31)
                                    switch (fi24 >> (((v3 - (v3 % ((v3) | 1)))) & 63)) {
                                    case -9223372036854775808:
                                        v33 := GenS1{F0: (v32 * 42), F1: (2 == 4294967295), F2: -9223372036854775808}
                                        _ = v33
                                    case 32768:
                                        _ = len(v1)
                                    default:
                                        v34 := GenS2{F0: 4294967295, F1: -128, F2: ("a" + "a")}
                                        _ = v34
                                    }
                                    if v2 {
                                        if ((!(v31 || true)) && true) {
                                            switch "foo bar" {
                                            case "a":
                                                var v35 uint8 = 2
                                                _ = v35
                                            case "":
                                                fmt.Println(v2, v4)
                                            }
                                            if true {
                                                var v36 int8 = -128
                                                delete(v1, 638)
                                                _ = v36
                                            } else {
                                                if ((-128 < 126) && (!v2)) {
                                                    v31 = v2
                                                    v31 = v2
                                                    var v37 bool = true
                                                    _ = v37
                                                }
                                                v31 = v31
                                                v38 := v1[243]
                                                _ = v38
                                            }
                                        } else {
                                            var v39 int32 = 180
                                            fmt.Println(fi25, fi25)
                                            _ = v39
                                        }
                                    }
                                }
                                _ = v31
                                _ = v32
                            } else {
                                var v40 int32 = 1
                                if (-510 > -875) {
                                    v3 += (v3 >> (((v3 | (^(v3)))) & 63))
                                    if v2 {
                                        delete(v1, 1)
                                        if false {
                                            v41 := GenS1{F0: (^(v40)), F1: (0.1 != (v5 * (v5 / 3.14159))), F2: 1}
                                            v40 = (^(v40))
                                            _ = v41
                                        } else {
                                            v42 := GenS1{F0: v40, F1: (!(96 > 752)), F2: 440}
                                            fmt.Println(v3)
                                            _ = v42
                                        }
                                    }
                                    var v43 int8 = 1
                                    if v2 {
                                        v44 := GenS3{F0: 88, F1: 9223372036854775807, F2: v40}
                                        _ = v44
                                    } else {
                                        v45 := GenS2{F0: 4294967295, F1: v43, F2: ""}
                                        _ = v45
                                    }
                                    _ = v43
                                } else {
                                    v3 -= v3
                                    v46 := GenS2{F0: 4294967295, F1: 42, F2: "hello"}
                                    v5 += ((v5 * (v5 - (float64((((float32(v5) / 1e10) + 0.1) - 3.14159)) / v5))) / v5)
                                    _ = v46
                                }
                                v3 += v3
                                _ = v40
                            }
                            if ((v3 / ((726) | 1)) != v3) {
                                v5 = (v5 / -2.25)
                            }
                            if (!(254 == 1)) {
                                v3 = 42
                            } else {
                                fmt.Println(v2)
                            }
                        }
                    }
                    switch "foo bar" {
                    case "a":
                        var v47 uint32 = 65535
                        _ = v47
                    case "hello":
                        v5 = v5
                    case "foo bar":
                        v5 += v5
                    default:
                        switch (fi24 >> ((uint(v4)) & 63)) {
                        case -1:
                            switch fi25 {
                            case 127:
                                fmt.Println(fi25, fi24)
                            case 128:
                                v48 := GenS4{F0: 2147483647}
                                _ = v48
                            }
                        case 422:
                            v49 := GenS1{F0: 2147483647, F1: v2, F2: -9223372036854775808}
                            _ = v49
                        case 1:
                            if v2 {
                                v4 += 2
                                v4 += v4
                                var v50 float32 = (((1.5 - 0.0) + 1e10) - 0.0)
                                fmt.Println(v4)
                                _ = v50
                            }
                        default:
                            var v51 uint16 = (v4 | (v4 / ((uint16(v3)) | 1)))
                            _ = v51
                        }
                    }
                    var v52 bool = (580 < 4294967295)
                    _ = v52
                } else {
                    switch "\n\t" {
                    case "":
                        v3 += (v3 & v3)
                    case "a":
                        switch fi25 {
                        case 32767:
                            var v53 bool = v2
                            _ = v53
                        case -823:
                            fmt.Println(v3, v2, v2)
                        default:
                            if (!((-32768 == 422) && v2)) {
                                _ = len(v1)
                            }
                        }
                    case "\n\t":
                        switch int(v3) {
                        case 2147483647:
                            if v2 {
                                _ = len(v1)
                                v54 := GenS5{F0: 539, F1: 4294967295}
                                switch ("" + ("foo bar" + "a")) {
                                case "hello":
                                    _ = len(v28)
                                default:
                                    v5 -= (v5 - v5)
                                }
                                _ = v54
                            }
                        }
                    default:
                        v55 := GenS1{F0: 0, F1: v2, F2: -9223372036854775808}
                        _ = v55
                    }
                    if true {
                        switch (fi25 << ((4294967295) & 63)) {
                        case 128:
                            var v56 int64 = -1
                            _ = v56
                        case -155:
                            v3 = (v3 - v3)
                        }
                    } else {
                        v57 := GenS1{F0: -459, F1: false, F2: 42}
                        fmt.Println(v3)
                        var v58 uint64 = 4294967295
                        fmt.Println(v2, v2, v4)
                        fmt.Println(v2)
                        var v59 int16 = 32767
                        switch 0 {
                        case 1:
                            if len(v28) > 0 {
                                v28[0] = 32768
                            }
                        }
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                }
            case 0:
                if (float32(v5) != 0.0) {
                    v2 = (!v2)
                    _ = len(v1)
                } else {
                    var v60 int16 = 1
                    v61 := v1[2]
                    if (4294967295 >= 260) {
                        v61 = (v61 * (v61 << (((v3 - uint(v61))) & 63)))
                        v5 = ((v5 - 0.1) + ((3.14159 - v5) + v5))
                        v62 := GenS2{F0: 545, F1: 127, F2: ""}
                        _ = v62
                    }
                    _ = v60
                    _ = v61
                }
            case 9223372036854775807:
                if true {
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v63 float64 = v5
                        _ = v63
                    }
                    v64 := GenS6{F0: 181, F1: "foo bar", F2: (0 <= -1)}
                    _ = v64
                } else {
                    if v2 {
                        if v2 {
                            if ((v3 != (v3 << ((18446744073709551615) & 63))) || false) {
                                var v65 float64 = (v5 - (v5 + ((v5 / v5) * v5)))
                                v65 += ((0.1 / v65) - v5)
                                _ = v65
                            } else {
                                switch (fi24 - fi24) {
                                case 127:
                                    var v66 uint16 = (v4 | v4)
                                    switch (("" + "foo bar") + ("foo bar" + "\n\t")) {
                                    case "":
                                        switch "\n\t" {
                                        case "foo bar":
                                            if len(v28) > 0 {
                                                v28[0] = (-(fi24))
                                            }
                                        case "\n\t":
                                            var v67 int16 = 42
                                            _ = v67
                                        case "hello":
                                            var v68 int8 = -1
                                            _ = v68
                                        default:
                                            v69 := v1[4294967295]
                                            switch fi24 {
                                            case -1:
                                                v5 += 0.0
                                            default:
                                                v2 = v2
                                            }
                                            _ = v69
                                        }
                                    default:
                                        var v70 float64 = v5
                                        v70 += 3.14159
                                        var v71 uint16 = (v4 | 1)
                                        _ = v70
                                        _ = v71
                                    }
                                    var v72 int16 = -1
                                    if (816 != int32(v4)) {
                                        if (v66 == v4) {
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                if v2 {
                                                    var v73 uint16 = v4
                                                    var v74 int64 = 701
                                                    var v75 uint32 = 568
                                                    var v76 uint = v3
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    if v2 {
                                                        if (!(3.14159 != float32((-2.25 / 3.14159)))) {
                                                            if ((1.5 * float32(v5)) <= 0.0) {
                                                                var v77 float32 = ((-2.25 * 3.14159) + (1e10 - float32(v5)))
                                                                v66 = v66
                                                                v78 := GenS6{F0: 2, F1: "", F2: ((v72 * (v72 | v72)) < (v72 - v72))}
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            v4 = v4
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                var v79 int8 = -64
                                                _ = v79
                                            default:
                                                if (!v2) {
                                                    v5 += 3.14159
                                                    if v2 {
                                                        var v80 uint32 = 1
                                                        switch (fi24 | 0) {
                                                        case -9223372036854775808:
                                                            switch fi25 {
                                                            case 4294967295:
                                                                fmt.Println(v5)
                                                            case 1:
                                                                fmt.Println(v3, v72)
                                                            }
                                                        case 0:
                                                            v81 := GenS6{F0: 8, F1: "a", F2: (0 <= -2147483648)}
                                                            _ = v81
                                                        case 9223372036854775807:
                                                            var v82 uint16 = 303
                                                            _ = v82
                                                        default:
                                                            if true {
                                                                var v83 float64 = (v5 * ((v5 * ((v5 * (v5 + float64(1.5))) - v5)) / v5))
                                                                if (2 < (^(v4))) {
                                                                    var v84 int = (-(fi24))
                                                                    if (42 <= 900) {
                                                                        var v85 bool = ((v2 || v2) && true)
                                                                        v85 = false
                                                                        fmt.Println(v72, fi24)
                                                                        _ = v85
                                                                    } else {
                                                                        var v86 bool = v2
                                                                        v5 = v83
                                                                        _ = v86
                                                                    }
                                                                    _ = v84
                                                                } else {
                                                                    if true {
                                                                        var v87 int8 = 1
                                                                        v88 := v1[v80]
                                                                        fmt.Println(v3)
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                    var v89 int32 = 0
                                                                    _ = v89
                                                                }
                                                                var v90 uint64 = 42
                                                                var v91 float32 = 3.14159
                                                                var v92 uint16 = v66
                                                                if v2 {
                                                                    v93 := v1[2]
                                                                    _ = v93
                                                                }
                                                                _ = v83
                                                                _ = v90
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                v5 -= v5
                                                                if (-1 != 693) {
                                                                    fmt.Println(fi24)
                                                                    var v94 int8 = -128
                                                                    v5 = float64((3.14159 - 1e10))
                                                                    if v2 {
                                                                        fmt.Println(v2)
                                                                        v94 += (v94 - v94)
                                                                    } else {
                                                                        v72 = (v72 - v72)
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v95 uint16 = 2
                                                                            _ = v95
                                                                        case "a":
                                                                            var v96 int8 = (v94 << (((^(v3))) & 63))
                                                                            v97 := v1[v80]
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "a":
                                                                        var v98 string = ("\n\t" + "\n\t")
                                                                        _ = v98
                                                                    case "":
                                                                        if (v2 && v2) {
                                                                            v99 := GenS7{F0: v3}
                                                                            fmt.Println(v2, v5)
                                                                            var v100 uint = v3
                                                                            if len(v28) >= 2 {
                                                                                v28 = v28[1:len(v28)]
                                                                            }
                                                                            _ = v100
                                                                            _ = v99
                                                                        } else {
                                                                            v2 = (1 != 114)
                                                                            if len(v28) >= 2 {
                                                                                v28 = v28[1:len(v28)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        v101 := GenS7{F0: (^(v3))}
                                                                        _ = v101
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                        }
                                                        _ = v80
                                                    } else {
                                                        v3 = v3
                                                    }
                                                } else {
                                                    switch (fi25 << ((117) & 63)) {
                                                    case 128:
                                                        if len(v28) >= 2 {
                                                            v28 = v28[1:len(v28)]
                                                        }
                                                    case 32768:
                                                        v102 := GenS8{F0: v72, F1: v3}
                                                        _ = v102
                                                    default:
                                                        if v2 {
                                                            v2 = ((-2.25 < 3.14159) || true)
                                                            var v103 string = "\n\t"
                                                            _ = v103
                                                        } else {
                                                            if ("" <= ("a" + "foo bar")) {
                                                                var v104 int8 = 42
                                                                v105 := GenS1{F0: 1, F1: v2, F2: -146}
                                                                _ = v104
                                                                _ = v105
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    if v2 {
                                                                        v1[2147483648] = uint32(v3)
                                                                        v106 := GenS9{F0: 42, F1: 0, F2: int64(v66)}
                                                                        var v107 int64 = int64(fi24)
                                                                        _ = v106
                                                                        _ = v107
                                                                    } else {
                                                                        if len(v28) >= 2 {
                                                                            v28 = v28[1:len(v28)]
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    if (853 >= 172) {
                                                                        switch 861 {
                                                                        case 9223372036854775807:
                                                                            if (v2 || v2) {
                                                                                v5 += v5
                                                                                v108 := GenS10{F0: 195, F1: 127, F2: (-47 <= 100)}
                                                                                fmt.Println(v72, v5)
                                                                                v109 := GenS11{F0: 4294967295, F1: 478}
                                                                                var v110 uint16 = v66
                                                                                var v111 string = "a"
                                                                                switch ((v111 + ((v111 + v111) + "\n\t")) + ("" + v111)) {
                                                                                case "hello":
                                                                                    v111 = v111
                                                                                }
                                                                                if len(v28) >= 2 {
                                                                                    v28 = v28[1:len(v28)]
                                                                                }
                                                                                if v2 {
                                                                                    v110 = v110
                                                                                    fmt.Println(v66)
                                                                                    fmt.Println(v5, v2)
                                                                                    var v112 int64 = -9
                                                                                    _ = v112
                                                                                }
                                                                                var v113 int16 = (v72 % (((v72 >> ((v3) & 63))) | 1))
                                                                                _ = v108
                                                                                _ = v109
                                                                                _ = v110
                                                                                _ = v111
                                                                                _ = v113
                                                                            } else {
                                                                                if v2 {
                                                                                    if ("\n\t" != "foo bar") {
                                                                                        if ((18446744073709551615 > 519) || v2) {
                                                                                            var v114 int32 = 0
                                                                                            v3 = 1
                                                                                            _ = v114
                                                                                        } else {
                                                                                            v3 = v3
                                                                                            fmt.Println(v72)
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v3, v2)
                                                                                }
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v115 := GenS12{F0: v3, F1: 4294967295}
                                                                                    _ = v115
                                                                                case "":
                                                                                    _ = len(v1)
                                                                                case "a":
                                                                                    var v116 string = "a"
                                                                                    _ = v116
                                                                                }
                                                                            }
                                                                        case 127:
                                                                            var v117 int16 = 1000
                                                                            _ = v117
                                                                        case 1:
                                                                            if (v2 && true) {
                                                                                var v118 uint = 392
                                                                                v119 := GenS13{F0: -1}
                                                                                _ = v118
                                                                                _ = v119
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    fmt.Println(v5, v4)
                                                                                case "hello":
                                                                                    switch (fi25 & fi24) {
                                                                                    case 641:
                                                                                        fmt.Println(v2)
                                                                                    default:
                                                                                        v120 := GenS14{F0: 42, F1: -526}
                                                                                        _ = v120
                                                                                    }
                                                                                case "a":
                                                                                    switch fi24 {
                                                                                    case -2147483648:
                                                                                        fmt.Println(v2, v4, v3)
                                                                                    case 0:
                                                                                        fmt.Println(v5)
                                                                                    case 1:
                                                                                        if v2 {
                                                                                            if len(v28) >= 2 {
                                                                                                v28 = v28[1:len(v28)]
                                                                                            }
                                                                                            if (!true) {
                                                                                                var v121 int64 = -9223372036854775808
                                                                                                v3 = (v3 * v3)
                                                                                                switch (^(fi25)) {
                                                                                                case -9223372036854775808:
                                                                                                    fmt.Println(fi24)
                                                                                                }
                                                                                                var v122 uint32 = 2
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            } else {
                                                                                                var v123 int8 = -128
                                                                                                v3 = v3
                                                                                                var v124 uint32 = 2
                                                                                                fmt.Println(fi25, v4)
                                                                                                _ = v123
                                                                                                _ = v124
                                                                                            }
                                                                                            v125 := GenS15{F0: (^(v3)), F1: -604, F2: v2}
                                                                                            _ = v125
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v28 = append(v28, (fi25 / (((^(fi25))) | 1)))
                                                                        }
                                                                    } else {
                                                                        var v126 uint8 = 0
                                                                        v1[0] = 971
                                                                        _ = v126
                                                                    }
                                                                case "foo bar":
                                                                    v72 -= (v72 + (v72 % (((v72 + 612)) | 1)))
                                                                }
                                                            }
                                                            var v127 string = ""
                                                            _ = v127
                                                        }
                                                    }
                                                    fmt.Println(v66, v72)
                                                }
                                            }
                                            v2 = true
                                            if true {
                                                var v128 int32 = int32(v66)
                                                if ((v3 >> ((v3) & 63)) != v3) {
                                                    if len(v28) >= 2 {
                                                        v28 = v28[1:len(v28)]
                                                    }
                                                    v72 = (v72 + 1000)
                                                } else {
                                                    if (!true) {
                                                        var v129 uint8 = 127
                                                        switch (fi24 % ((fi25) | 1)) {
                                                        case 42:
                                                            v129 += (v129 - v129)
                                                        case -2147483648:
                                                            if v2 {
                                                                v130 := GenS6{F0: 127, F1: (((("" + "") + ("a" + "hello")) + "") + "a"), F2: false}
                                                                v129 = 127
                                                                var v131 int32 = v128
                                                                _ = v130
                                                                _ = v131
                                                            }
                                                        default:
                                                            v4 -= (v4 ^ (^(v4)))
                                                        }
                                                        v3 = uint(v129)
                                                        if v2 {
                                                            v1[102] = 2
                                                            var v132 uint64 = 1
                                                            fmt.Println(fi25)
                                                            _ = v132
                                                        } else {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v133 float32 = 1e10
                                                                _ = v133
                                                            case "a":
                                                                v129 = (v129 >> ((v3) & 63))
                                                            }
                                                            var v134 uint64 = 18446744073709551615
                                                            _ = v134
                                                        }
                                                        switch 32768 {
                                                        case -9223372036854775808:
                                                            v66 += (v4 >> ((v3) & 63))
                                                        case -2147483648:
                                                            var v135 bool = false
                                                            _ = v135
                                                        case -604:
                                                            v5 = (((0.1 / float64((((((-2.25 * (float32(v5) + 3.14159)) / 0.1) * float32((3.14159 / ((((v5 / 1.5) + v5) * v5) - 0.0)))) - 0.0) + 1e10))) - (v5 / float64(float32(float64(3.14159))))) * (v5 + -2.25))
                                                            if len(v28) > 0 {
                                                                v28[0] = 946
                                                            }
                                                        }
                                                        _ = v129
                                                    } else {
                                                        switch "a" {
                                                        case "hello":
                                                            var v136 float32 = 3.14159
                                                            _ = v136
                                                        }
                                                        v28 = append(v28, fi25)
                                                        v2 = (((v66 + v4) >= 255) && (!false))
                                                    }
                                                    v137 := GenS11{F0: 18446744073709551615, F1: 2}
                                                    _ = v137
                                                }
                                                _ = v128
                                            } else {
                                                if len(v28) > 0 {
                                                    v28[0] = (fi24 & (^(fi24)))
                                                }
                                                fmt.Println(v3, v66, v5)
                                                var v138 bool = ((!v2) && ((1e10 <= v5) && v2))
                                                if (9223372036854775807 > -1) {
                                                    var v139 uint32 = 2
                                                    if v2 {
                                                        if false {
                                                            v138 = v2
                                                            v140 := GenS14{F0: 0, F1: -937}
                                                            switch ("hello" + "") {
                                                            case "\n\t":
                                                                v66 = (v4 | uint16(v3))
                                                            case "hello":
                                                                v66 = v66
                                                            case "":
                                                                v5 += v5
                                                                v141 := GenS16{F0: 223, F1: 56}
                                                                if v2 {
                                                                    if v138 {
                                                                        v3 = 0
                                                                        var v142 uint8 = uint8(v3)
                                                                        v143 := GenS17{F0: (v142 >> (((v3 % ((uint(v4)) | 1))) & 63))}
                                                                        v4 -= uint16(v142)
                                                                        _ = v142
                                                                        _ = v143
                                                                    } else {
                                                                        if v2 {
                                                                            fmt.Println(v3)
                                                                            var v144 int = (^(fi24))
                                                                            var v145 int32 = 42
                                                                            _ = v144
                                                                            _ = v145
                                                                        }
                                                                        v146 := GenS18{F0: 566, F1: -1, F2: v4}
                                                                        _ = v146
                                                                    }
                                                                    v147 := GenS17{F0: 2}
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v2 = (!v2)
                                                                    }
                                                                    _ = v147
                                                                } else {
                                                                    if (!(-2147483648 > int32(v72))) {
                                                                        var v148 bool = v138
                                                                        _ = v148
                                                                    } else {
                                                                        v4 = 65535
                                                                        if v2 {
                                                                            var v149 uint16 = (v66 >> ((v3) & 63))
                                                                            switch int(v72) {
                                                                            case 2147483647:
                                                                                v72 += (-(v72))
                                                                            case 32768:
                                                                                var v150 uint8 = 108
                                                                                _ = v150
                                                                            }
                                                                            _ = v149
                                                                        } else {
                                                                            if (128 < 0) {
                                                                                if false {
                                                                                    if ("\n\t" < ((((("a" + "\n\t") + (((("" + ("" + "foo bar")) + "") + "\n\t") + "\n\t")) + "foo bar") + "\n\t") + "foo bar")) {
                                                                                        var v151 float32 = -2.25
                                                                                        _ = v151
                                                                                    }
                                                                                    v139 = (^(v139))
                                                                                    var v152 bool = (!true)
                                                                                    _ = v152
                                                                                } else {
                                                                                    v139 = 2147483648
                                                                                    v4 -= (v66 | 255)
                                                                                    v72 -= v72
                                                                                }
                                                                                v66 = (v4 / ((v66) | 1))
                                                                                fmt.Println(v66, v139, v139)
                                                                                v2 = ((v72 ^ (v72 + -32768)) <= (v72 % ((-523) | 1)))
                                                                            }
                                                                            var v153 uint8 = 128
                                                                            v154 := GenS19{F0: 0, F1: v153, F2: -2147483648}
                                                                            v155 := GenS20{F0: -46}
                                                                            _ = v153
                                                                            _ = v154
                                                                            _ = v155
                                                                        }
                                                                    }
                                                                }
                                                                var v156 uint16 = (v4 % ((42) | 1))
                                                                _ = v141
                                                                _ = v156
                                                            default:
                                                                var v157 int64 = 42
                                                                _ = v157
                                                            }
                                                            _ = v140
                                                        } else {
                                                            v3 = v3
                                                            v158 := GenS19{F0: (^(v4)), F1: 128, F2: 1}
                                                            if ((-(v72)) == (v72 ^ (v72 << (((v3 + (v3 - (v3 / ((592) | 1))))) & 63)))) {
                                                                v159 := GenS21{F0: v138}
                                                                var v160 int16 = v72
                                                                _ = v159
                                                                _ = v160
                                                            } else {
                                                                var v161 uint8 = 233
                                                                _ = len(v1)
                                                                switch fi25 {
                                                                case 1:
                                                                    v162 := GenS19{F0: v4, F1: (^(v161)), F2: 42}
                                                                    _ = v162
                                                                default:
                                                                    v72 = -984
                                                                }
                                                                _ = v161
                                                            }
                                                            if (v138 && (true && v138)) {
                                                                var v163 int32 = 42
                                                                _ = v163
                                                            } else {
                                                                switch 1 {
                                                                case 445:
                                                                    if len(v28) >= 2 {
                                                                        v28 = v28[1:len(v28)]
                                                                    }
                                                                    var v164 string = ("\n\t" + "a")
                                                                    _ = v164
                                                                default:
                                                                    v139 = (v139 % ((4294967295) | 1))
                                                                }
                                                                v1[(^(v139))] = (v139 & v139)
                                                            }
                                                            fmt.Println(v4)
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(v139, v138, v138)
                                                            }
                                                            _ = v158
                                                        }
                                                    } else {
                                                        fmt.Println(v138)
                                                    }
                                                    if v2 {
                                                        fmt.Println(fi24, v139, v139)
                                                        v2 = false
                                                        v165 := GenS22{F0: 137, F1: "hello"}
                                                        _ = v165
                                                    }
                                                    _ = v139
                                                }
                                                v2 = ((!(!true)) || v138)
                                                _ = v138
                                            }
                                        } else {
                                            v3 = 65535
                                            fmt.Println(fi25, fi25)
                                        }
                                        var v166 uint = (v3 | 593)
                                        _ = v166
                                    }
                                    _ = len(v28)
                                    _ = v66
                                    _ = v72
                                case -326:
                                    switch "foo bar" {
                                    case "":
                                        var v167 uint32 = 17
                                        _ = v167
                                    }
                                default:
                                    v1[1] = 42
                                }
                                var v168 int16 = 69
                                var v169 float64 = (v5 / v5)
                                _ = v168
                                _ = v169
                            }
                            if len(v28) > 0 {
                                v28[0] = fi25
                            }
                        }
                        v4 += (v4 % (((v4 + 255)) | 1))
                    } else {
                        var v170 int = fi24
                        _ = v170
                    }
                    if (407 >= 830) {
                        v4 = (v4 ^ (v4 + 392))
                    }
                    v4 = 0
                    fmt.Println(fi25, v4)
                    if (!((((0.0 / 0.1) / 1.0) <= float32(v5)) || v2)) {
                        if len(v28) > 0 {
                            v28[0] = 514
                        }
                        v171 := GenS18{F0: 9223372036854775807, F1: 1000, F2: v4}
                        var v172 string = ("foo bar" + "a")
                        _ = v171
                        _ = v172
                    } else {
                        _ = len(v1)
                        v28 = append(v28, fi25)
                        v2 = v2
                        var v173 int8 = 0
                        _ = v173
                    }
                    var v174 uint8 = 1
                    _ = v174
                }
            }
            _ = fi25
            _ = v28
        }
        _ = fi24
    }
    for fi175 := 0; fi175 < 8; fi175++ {
        for fi176 := 0; fi176 < 5; fi176++ {
            if ("" == ("" + "")) {
                switch "\n\t" {
                case "a":
                    v177 := []int{(fi176 >> ((v3) & 63)), (fi175 | -467), fi175}
                    _ = v177
                case "hello":
                    if (fi175 <= (fi176 * (fi175 ^ fi175))) {
                        fmt.Println(v3)
                    }
                case "\n\t":
                    var v178 float32 = (3.14159 + float32(v5))
                    var v179 uint8 = 127
                    var v180 uint = v3
                    _ = v178
                    _ = v179
                    _ = v180
                }
                var v181 uint16 = 1
                v1[42] = 1
                _ = v181
            } else {
                if true {
                    v2 = false
                    switch fi176 {
                    case -1:
                        v1[0] = uint32(v3)
                    default:
                        if (32767 > -32768) {
                            switch "foo bar" {
                            case "foo bar":
                                var v182 string = ("a" + "")
                                _ = v182
                            case "":
                                var v183 bool = (2147483648 == 2147483647)
                                v3 -= (^(v3))
                                _ = v183
                            case "a":
                                var v184 float64 = -2.25
                                _ = v184
                            }
                            v185 := GenS23{F0: 4294967295, F1: 0, F2: 1}
                            _ = v185
                        }
                    }
                    if (2147483647 == 65535) {
                        fmt.Println(v3, v5, fi176)
                        var v186 float32 = 0.0
                        v187 := GenS24{F0: 18446744073709551615, F1: 0}
                        v188 := GenS25{F0: -9223372036854775808, F1: 78, F2: 4294967295}
                        _ = v186
                        _ = v187
                        _ = v188
                    } else {
                        fmt.Println(fi176, fi176)
                        v4 += (v4 / ((uint16(v3)) | 1))
                        v189 := []int{(fi176 + 677)}
                        _ = v189
                    }
                }
                v190 := GenS26{F0: -56, F1: 793}
                _ = v190
            }
            var v191 uint = (v3 & v3)
            _ = fi176
            _ = v191
        }
        fmt.Println(v4, v3, v4)
        _ = fi175
    }
    for fi192 := 0; fi192 < 3; fi192++ {
        v4 -= v4
        v3 = v3
        _ = fi192
    }
    v193 := GenS9{F0: 42, F1: 828, F2: -9223372036854775808}
    if (65535 == uint16(v3)) {
        v194 := []int{}
        v195 := GenS22{F0: 205, F1: "foo bar"}
        v5 += v5
        _ = v194
        _ = v195
    } else {
        if (0.0 > v5) {
            var v196 int32 = -2147483648
            _ = v196
        }
        fmt.Println(v3, v5)
    }
    var v197 uint8 = 42
    v198 := GenS27{F0: -128, F1: 0, F2: v4}
    switch -959 {
    case 86:
        var v199 bool = (v2 || v2)
        _ = v199
    case -1:
        v200 := []int{-302, -1}
        _ = v200
    case 858:
        for fi201 := 0; fi201 < 5; fi201++ {
            v3 = v3
            v2 = v2
            v202 := []int{-2147483648, fi201}
            _ = fi201
            _ = v202
        }
    }
    var v203 uint64 = 0
    var v204 int8 = int8(v203)
    switch "hello" {
    case "a":
        v205 := GenS14{F0: v197, F1: 560}
        _ = v205
    case "\n\t":
        fmt.Println(v203, v4, v5)
    default:
        v197 = (^(v197))
    }
    v206 := GenS13{F0: v204}
    for fi207 := 0; fi207 < 16; fi207++ {
        var v208 float32 = (((0.1 / float32(v5)) * ((0.1 / 1.0) + 0.1)) * 1e10)
        var v209 int64 = int64(fi207)
        var v210 uint32 = uint32(v4)
        _ = fi207
        _ = v208
        _ = v209
        _ = v210
    }
    for fi211 := 0; fi211 < 12; fi211++ {
        v4 = (v4 / (((v4 + 586)) | 1))
        _ = fi211
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v197:", v197)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v193
    _ = v197
    _ = v198
    _ = v2
    _ = v203
    _ = v204
    _ = v206
    _ = v3
    _ = v4
    _ = v5
}
