package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 uint; F2 int16 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint32; F1 int8; F2 int64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 string; F1 string }
type GenS16 struct { F0 string; F1 string; F2 int32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int8; F1 bool; F2 uint8 }
type GenS19 struct { F0 uint8 }
func genFn0() int8 {
    switch "a" {
    case "foo bar":
        var v4 uint8 = 44
        _ = v4
    case "a":
        var v5 string = ""
        _ = v5
    default:
        var v6 string = (("foo bar" + ((("a" + ("a" + "\n\t")) + "a") + "a")) + "hello")
        _ = v6
    }
    if true {
        v7 := []int{127, -943, -1, 514}
        _ = v7
    } else {
        v8 := []int{4294967295}
        _ = v8
    }
    return 62
}
func genFn1() int64 {
    var v9 int8 = -35
    if false {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint32 = 42
            v12 := []int{(fi10 - (fi10 / (((fi10 << ((255) & 63))) | 1))), fi10}
            var v13 uint64 = 2
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    } else {
        if false {
            v9 += (v9 / ((-98) | 1))
        } else {
            v9 = -128
            v9 += (v9 >> ((1) & 63))
            v14 := []int{32767, -1, -520, -1}
            _ = v14
        }
        v15 := GenS1{F0: -595, F1: 80}
        v16 := []int{32767}
        _ = v15
        _ = v16
    }
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{1, 1, 65535}
    v2 := make(map[string]string)
    var v3 int = 9223372036854775807
    if true {
        var v17 float32 = 0.0
        var v18 float32 = 0.0
        switch v3 {
        case 1:
            switch "foo bar" {
            case "hello":
                if false {
                    v3 += (-(v3))
                } else {
                    _ = len(v1)
                    v3 += v3
                    v18 = (float32(0.1) - float32((((float64(3.14159) / 0.1) / ((3.14159 * 0.0) + (1.5 / (((3.14159 * (3.14159 / 3.14159)) / ((3.14159 / ((0.1 - -2.25) / (0.1 * -2.25))) - ((0.0 / 1.0) / 3.14159))) / 1.0)))) / -2.25)))
                }
            default:
                if (42 < 255) {
                    var v19 string = "foo bar"
                    if (0 >= 42) {
                        _ = len(v2)
                        v20 := GenS2{F0: v19, F1: -1}
                        _ = v20
                    } else {
                        var v21 int = v3
                        if len(v1) > 0 {
                            v1[0] = 877
                        }
                        _ = v21
                    }
                    _ = v19
                }
            }
        case -1:
            v22 := v2["a"]
            _ = v22
        default:
            v17 += 0.0
        }
        _ = v17
        _ = v18
    } else {
        v23 := GenS3{F0: 2, F1: (v3 - (-(v3)))}
        if ((1.5 - 1.5) > -2.25) {
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v24 uint32 = 2
                    if (true && (1 >= 126)) {
                        v25 := GenS3{F0: 369, F1: 9223372036854775807}
                        switch "\n\t" {
                        case "":
                            var v26 uint = 783
                            _ = v26
                        case "\n\t":
                            v27 := GenS4{F0: (v24 << ((817) & 63)), F1: 42}
                            _ = v27
                        default:
                            v28 := GenS3{F0: 1, F1: (v3 % ((v3) | 1))}
                            _ = v28
                        }
                        _ = v25
                    } else {
                        v24 += (v24 ^ (^(v24)))
                    }
                    _ = v24
                } else {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 string = ("foo bar" + "\n\t")
                        v3 = fi29
                        _ = fi29
                        _ = v30
                    }
                }
            case "a":
                v31 := GenS5{F0: 1}
                _ = v31
            }
            if (-9223372036854775808 == genFn1()) {
                v3 -= (v3 % ((v3) | 1))
                v32 := GenS6{F0: 0, F1: 42}
                _ = v32
            } else {
                if ((!true) || (-635 >= 42)) {
                    if (1e10 <= (((1.5 + (-2.25 / 1.0)) + (1.5 - float32((0.0 - 0.1)))) - 0.0)) {
                        var v33 uint64 = 42
                        _ = v33
                    } else {
                        var v34 int8 = -1
                        if true {
                            _ = len(v2)
                            v3 = int(v34)
                            v3 += (-(v3))
                        }
                        _ = v34
                    }
                }
                var v35 uint16 = 255
                _ = v35
            }
            v36 := GenS7{F0: 255}
            _ = v36
        } else {
            fmt.Println(v3)
            switch ("a" + (("a" + (("foo bar" + "foo bar") + ("" + ("a" + "\n\t")))) + "\n\t")) {
            case "hello":
                if (255 < 1) {
                    if true {
                        fmt.Println(v3)
                        var v37 uint = 693
                        var v38 float32 = 3.14159
                        _ = v37
                        _ = v38
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v39 := GenS1{F0: 32767, F1: 42}
                        var v40 uint8 = 101
                        _ = v39
                        _ = v40
                    }
                }
            default:
                switch v3 {
                case 0:
                    if ((32767 != 65535) && (!(!true))) {
                        var v41 bool = (!((70 != 868) && (((1e10 / float32(1e10)) * 3.14159) < 3.14159)))
                        var v42 int64 = 42
                        fmt.Println(v41, v41)
                        _ = v41
                        _ = v42
                    } else {
                        var v43 int16 = -673
                        fmt.Println(v43, v43)
                        if true {
                            var v44 uint = 2
                            _ = v44
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (1 < 116) {
                                if true {
                                    delete(v2, "")
                                }
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    v3 = (-(v3))
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v43)
                                    case "a":
                                        v46 := GenS8{F0: 4294967295}
                                        _ = v46
                                    }
                                    for fi47 := 0; fi47 < 4; fi47++ {
                                        v48 := v2[""]
                                        _ = fi47
                                        _ = v48
                                    }
                                    _ = fi45
                                }
                            } else {
                                var v49 int32 = 135
                                var v50 int64 = 9223372036854775807
                                _ = v49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 19; fi51++ {
                                switch (fi51 ^ fi51) {
                                case 128:
                                    if true {
                                        v43 += v43
                                        for fi52 := 0; fi52 < 10; fi52++ {
                                            if false {
                                                if (255 >= 2) {
                                                    if (((((3.14159 * 1e10) + 3.14159) + (((0.1 / 1.0) - -2.25) + 0.1)) - 0.1) < 0.0) {
                                                        v3 = fi51
                                                        v53 := GenS9{F0: 42, F1: 4294967295, F2: (^(v43))}
                                                        _ = v53
                                                    }
                                                    var v54 int16 = (-(v43))
                                                    var v55 uint = 85
                                                    v56 := GenS3{F0: 42, F1: (^(fi51))}
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v43)
                                                    if (127 == 20) {
                                                        var v57 uint8 = 250
                                                        v58 := v2["hello"]
                                                        if false {
                                                            v3 = -9223372036854775808
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        switch "a" {
                                                        case "hello":
                                                            v3 += (fi52 >> ((65535) & 63))
                                                        case "\n\t":
                                                            if (true && (0 != -1)) {
                                                                v59 := GenS10{F0: genFn0(), F1: 0}
                                                                _ = v59
                                                            } else {
                                                                var v60 bool = true
                                                                var v61 int16 = 32767
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        default:
                                                            v62 := GenS10{F0: 100, F1: 65535}
                                                            _ = v62
                                                        }
                                                    }
                                                    switch (^(fi52)) {
                                                    case -2147483648:
                                                        switch (fi51 | v3) {
                                                        case 383:
                                                            var v63 bool = true
                                                            _ = v63
                                                        case 32768:
                                                            var v64 uint8 = 173
                                                            _ = v64
                                                        case 9223372036854775807:
                                                            switch fi52 {
                                                            case -290:
                                                                v43 -= -748
                                                            case 25:
                                                                var v65 float32 = (0.0 + ((((3.14159 / 3.14159) - 1.5) * 3.14159) - ((0.1 / 1.5) + (-2.25 + 1e10))))
                                                                _ = v65
                                                            case 0:
                                                                if (3.14159 > 0.1) {
                                                                    if (100 < 127) {
                                                                        v3 = 0
                                                                        v66 := GenS4{F0: 42, F1: 42}
                                                                        switch -490 {
                                                                        case 4294967295:
                                                                            var v67 uint8 = 1
                                                                            _ = v67
                                                                        case 2147483647:
                                                                            fmt.Println(v3, fi51, fi52)
                                                                        default:
                                                                            var v68 int8 = -9
                                                                            _ = v68
                                                                        }
                                                                        _ = v66
                                                                    }
                                                                }
                                                            default:
                                                                v43 = (v43 & (v43 | v43))
                                                            }
                                                        default:
                                                            if (18446744073709551615 < 4294967295) {
                                                                if false {
                                                                    v69 := GenS11{F0: 4294967295, F1: 282}
                                                                    var v70 uint64 = 382
                                                                    var v71 uint8 = 124
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                if false {
                                                                    if true {
                                                                        v43 -= (^(v43))
                                                                        v72 := GenS12{F0: 2147483647}
                                                                        fmt.Println(v43)
                                                                        _ = v72
                                                                    }
                                                                }
                                                                v73 := GenS6{F0: 255, F1: 1}
                                                                _ = v73
                                                            } else {
                                                                if false {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v43 -= v43
                                                                    case "hello":
                                                                        var v74 int64 = 744
                                                                        if (false && (-2.25 < 1e10)) {
                                                                            if (!(!(((v74 << ((0) & 63)) >= -1) || (!true)))) {
                                                                                if (65535 < 42) {
                                                                                    v1 = append(v1, 4294967295)
                                                                                    v74 = v74
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v75 uint = 628
                                                                            _ = v75
                                                                        }
                                                                        _ = len(v2)
                                                                        if (601 < 264) {
                                                                            v74 = (v74 | genFn1())
                                                                            v76 := GenS4{F0: 21, F1: 2147483647}
                                                                            _ = v76
                                                                        }
                                                                        _ = v74
                                                                    }
                                                                    if (!(127 != 1)) {
                                                                        v3 = 665
                                                                        var v77 float64 = 1.5
                                                                        if true {
                                                                            if true {
                                                                                v78 := GenS8{F0: 295}
                                                                                _ = v78
                                                                            } else {
                                                                                var v79 float64 = -2.25
                                                                                var v80 uint8 = 127
                                                                                _ = v79
                                                                                _ = v80
                                                                            }
                                                                            if true {
                                                                                v81 := GenS13{F0: 476, F1: 100, F2: -807}
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                        v3 += fi51
                                                                        var v82 int = (fi52 >> ((255) & 63))
                                                                        _ = v77
                                                                        _ = v82
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2
                                                                        }
                                                                        var v83 int32 = 0
                                                                        v43 = 42
                                                                        _ = v83
                                                                    }
                                                                }
                                                                v84 := GenS12{F0: 2147483647}
                                                                delete(v2, ("" + (("foo bar" + "hello") + "\n\t")))
                                                                _ = v84
                                                            }
                                                        }
                                                    case 128:
                                                        v1 = append(v1, 822)
                                                    case 9223372036854775807:
                                                        var v85 uint = 65535
                                                        _ = v85
                                                    default:
                                                        if true {
                                                            if false {
                                                                var v86 uint32 = 896
                                                                fmt.Println(fi51, fi51)
                                                                _ = v86
                                                            } else {
                                                                v3 += (fi51 << ((2) & 63))
                                                                v3 -= (v3 - (fi52 % ((fi51) | 1)))
                                                            }
                                                            if (("foo bar" + "\n\t") >= "\n\t") {
                                                                v87 := v2[("" + "hello")]
                                                                var v88 int8 = -117
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                        }
                                                    }
                                                }
                                                v3 = (fi51 + (fi51 / ((32768) | 1)))
                                            } else {
                                                var v89 uint16 = 27
                                                _ = v89
                                            }
                                            _ = fi52
                                        }
                                    }
                                }
                                for fi90 := 0; fi90 < 9; fi90++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch (fi90 | 9223372036854775807) {
                                    case 252:
                                        if (255 == 1) {
                                            fmt.Println(fi90, fi51, v43)
                                            if true {
                                                delete(v2, "")
                                            } else {
                                                if false {
                                                    var v91 uint16 = 851
                                                    _ = v91
                                                } else {
                                                    fmt.Println(fi51)
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 363
                                                }
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v92 bool = false
                                                _ = v92
                                            case "\n\t":
                                                if true {
                                                    fmt.Println(v3, v43)
                                                    v43 += -577
                                                    v3 += v3
                                                } else {
                                                    _ = len(v1)
                                                    var v93 bool = (!(42 == 32767))
                                                    var v94 uint16 = 0
                                                    v93 = v93
                                                    _ = v93
                                                    _ = v94
                                                }
                                            default:
                                                fmt.Println(fi90)
                                            }
                                        } else {
                                            if true {
                                                var v95 string = ("" + (("" + "hello") + ""))
                                                var v96 float64 = (1.5 + (1e10 + 0.1))
                                                var v97 int8 = 0
                                                var v98 uint8 = 108
                                                var v99 int16 = -32768
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                                _ = v98
                                                _ = v99
                                            }
                                            v43 = -809
                                            v43 += 1000
                                        }
                                    case -1:
                                        if (!false) {
                                            delete(v2, (("foo bar" + ("" + "a")) + (("" + "foo bar") + ("\n\t" + ""))))
                                            v100 := GenS12{F0: 888}
                                            _ = v100
                                        }
                                    case 0:
                                        switch (fi90 ^ (^(fi51))) {
                                        case -801:
                                            v101 := GenS14{F0: -275}
                                            _ = v101
                                        case 4294967295:
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = 245
                                                }
                                                v43 += v43
                                                var v102 int = v3
                                                var v103 int = (fi90 | fi51)
                                                _ = v102
                                                _ = v103
                                            }
                                        case 2147483647:
                                            switch "hello" {
                                            case "a":
                                                if false {
                                                    v104 := GenS2{F0: ("\n\t" + "foo bar"), F1: (-(v43))}
                                                    if ((!(-1 < 2147483647)) || (((0.0 / ((((1.5 / (1e10 / (float32(0.0) - -2.25))) - -2.25) * 1.5) + (float32((0.1 / 1.0)) - 0.0))) * 1.5) <= 0.0)) {
                                                        if true {
                                                            v3 += v3
                                                        }
                                                        fmt.Println(fi90, v3, fi51)
                                                    } else {
                                                        v43 -= (^(v43))
                                                        var v105 uint16 = 860
                                                        if (100 > genFn0()) {
                                                            if (1 >= 2) {
                                                                var v106 int = (v3 % ((fi90) | 1))
                                                                switch fi51 {
                                                                case 32768:
                                                                    var v107 uint = 758
                                                                    _ = v107
                                                                case -129:
                                                                    _ = len(v1)
                                                                default:
                                                                    fmt.Println(fi51)
                                                                }
                                                                _ = v106
                                                            } else {
                                                                delete(v2, "")
                                                                fmt.Println(v105)
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v108 int64 = -76
                                                                    _ = v108
                                                                case "a":
                                                                    v109 := GenS2{F0: "", F1: (v43 ^ v43)}
                                                                    _ = v109
                                                                default:
                                                                    var v110 string = ""
                                                                    _ = v110
                                                                }
                                                                var v111 string = ""
                                                                _ = len(v2)
                                                                _ = v111
                                                            }
                                                        } else {
                                                            v112 := GenS8{F0: 268}
                                                            var v113 uint16 = (^(v105))
                                                            v3 = fi90
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch 12 {
                                                                case 32768:
                                                                    var v114 float32 = ((1.5 / 3.14159) * ((float32((-2.25 * float64(0.1))) * 3.14159) + (1e10 + 0.1)))
                                                                    _ = v114
                                                                }
                                                            default:
                                                                v115 := GenS15{F0: "hello", F1: "foo bar"}
                                                                _ = v115
                                                            }
                                                            v43 += v43
                                                            _ = v112
                                                            _ = v113
                                                        }
                                                        var v116 uint32 = uint32(v105)
                                                        _ = v105
                                                        _ = v116
                                                    }
                                                    v43 -= v43
                                                    v3 += (v3 * v3)
                                                    v117 := GenS7{F0: 147}
                                                    _ = v104
                                                    _ = v117
                                                } else {
                                                    var v118 uint16 = 42
                                                    v119 := GenS16{F0: ((("foo bar" + "\n\t") + "\n\t") + "hello"), F1: ("hello" + ((("a" + "foo bar") + (("a" + "hello") + "")) + "\n\t")), F2: 176}
                                                    _ = v118
                                                    _ = v119
                                                }
                                            case "\n\t":
                                                if (!false) {
                                                    v3 -= fi90
                                                } else {
                                                    fmt.Println(v3, fi90)
                                                    fmt.Println(fi90, v43)
                                                    if len(v1) > 0 {
                                                        v1[0] = 2147483648
                                                    }
                                                }
                                            case "hello":
                                                var v120 int16 = -218
                                                _ = v120
                                            default:
                                                if false {
                                                    switch fi51 {
                                                    case -305:
                                                        v121 := GenS17{F0: 1}
                                                        _ = v121
                                                    case 443:
                                                        delete(v2, ("\n\t" + "\n\t"))
                                                    default:
                                                        var v122 uint = 2
                                                        _ = v122
                                                    }
                                                    var v123 uint = 65535
                                                    _ = v123
                                                }
                                            }
                                        }
                                    default:
                                        var v124 string = ("\n\t" + "\n\t")
                                        _ = v124
                                    }
                                    v3 -= 1
                                    _ = fi90
                                }
                                _ = fi51
                            }
                        }
                        _ = v43
                    }
                case 4294967295:
                    v125 := GenS14{F0: 42}
                    _ = v125
                case 846:
                    if (854 <= 0) {
                        v3 = (-(v3))
                        for fi126 := 0; fi126 < 16; fi126++ {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            v127 := GenS11{F0: 964, F1: 1}
                            _ = fi126
                            _ = v127
                        }
                    } else {
                        switch (("hello" + ("\n\t" + "foo bar")) + "\n\t") {
                        case "hello":
                            v1 = append(v1, 0)
                        }
                    }
                default:
                    v3 = (v3 + 2147483647)
                }
            }
            v3 = 0
        }
        fmt.Println(v3)
        _ = v23
    }
    var v128 uint64 = 22
    var v129 uint64 = 42
    switch "\n\t" {
    case "foo bar":
        v130 := GenS9{F0: 622, F1: 510, F2: 650}
        _ = v130
    case "a":
        fmt.Println(v3, v128, v3)
    }
    if ((true && (!(2 >= 430))) || false) {
        var v131 string = ""
        _ = v131
    }
    v132 := GenS15{F0: "a", F1: "hello"}
    var v133 int64 = genFn1()
    if true {
        for fi134 := 0; fi134 < 18; fi134++ {
            fmt.Println(v128, v133)
            _ = fi134
        }
        var v135 uint8 = 194
        v3 -= v3
        _ = v135
    } else {
        for fi136 := 0; fi136 < 15; fi136++ {
            for fi137 := 0; fi137 < 15; fi137++ {
                v138 := GenS2{F0: "foo bar", F1: 0}
                _ = fi137
                _ = v138
            }
            v129 -= 64
            v139 := GenS11{F0: v128, F1: 2147483647}
            _ = fi136
            _ = v139
        }
        var v140 int64 = v133
        for fi141 := 0; fi141 < 7; fi141++ {
            for fi142 := 0; fi142 < 4; fi142++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = 687
                    }
                    if true {
                        v128 += 428
                        if (834 <= 42) {
                            var v143 uint64 = v129
                            _ = v143
                        } else {
                            v144 := GenS2{F0: "\n\t", F1: 32767}
                            if false {
                                delete(v2, (("foo bar" + "foo bar") + "hello"))
                                if (13 > 65535) {
                                    var v145 uint64 = (v128 << ((255) & 63))
                                    v146 := v2[("a" + "\n\t")]
                                    _ = v145
                                    _ = v146
                                }
                            } else {
                                fmt.Println(v140, fi142, v3)
                                var v147 int = v3
                                _ = v147
                            }
                            _ = v144
                        }
                        v128 = v129
                    }
                    fmt.Println(fi142, v133, v133)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v2["hello"] = "hello"
                _ = fi142
            }
            var v148 uint64 = (^(v128))
            v149 := GenS18{F0: 127, F1: ((2147483648 <= 42) || false), F2: 255}
            _ = fi141
            _ = v148
            _ = v149
        }
        _ = v140
    }
    for fi150 := 0; fi150 < 14; fi150++ {
        for fi151 := 0; fi151 < 3; fi151++ {
            v3 = fi151
            _ = fi151
        }
        switch ("foo bar" + "a") {
        case "\n\t":
            v128 -= v128
        case "":
            v152 := GenS5{F0: -1}
            _ = v152
        case "foo bar":
            v153 := GenS19{F0: 0}
            _ = v153
        default:
            var v154 int16 = -987
            _ = v154
        }
        _ = fi150
    }
    if (42 == 2147483647) {
        for fi155 := 0; fi155 < 9; fi155++ {
            v129 = v128
            _ = fi155
        }
        v129 = (v129 >> ((42) & 63))
        switch (v3 & v3) {
        case 297:
            var v156 float64 = 0.1
            _ = v156
        case 9223372036854775807:
            var v157 int8 = -128
            _ = v157
        default:
            if (!true) {
                var v158 int32 = 1
                _ = v158
            }
        }
    } else {
        for fi159 := 0; fi159 < 19; fi159++ {
            if false {
                var v160 uint16 = 65535
                v129 -= (v129 ^ uint64(v160))
                _ = v160
            }
            _ = fi159
        }
        switch (("\n\t" + "hello") + "\n\t") {
        case "a":
            for fi161 := 0; fi161 < 1; fi161++ {
                v162 := GenS3{F0: 448, F1: -732}
                _ = fi161
                _ = v162
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v133 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v133:", v133)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v128
    _ = v129
    _ = v132
    _ = v133
    _ = v2
    _ = v3
}
