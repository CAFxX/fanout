package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int8; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint }
type GenS8 struct { F0 int8; F1 uint16; F2 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint8; F2 uint8 }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string; F1 int32; F2 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 string; F1 uint32; F2 int }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int; F1 int8; F2 uint16 }
type GenS22 struct { F0 uint; F1 uint8; F2 int }
type GenS23 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint16) uint64 {
    switch 1 {
    case 18:
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint16 = p0
            _ = fi4
            _ = v5
        }
    case 812:
        var v6 uint = 4294967295
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 508, F2: 100}
    _ = p0
    _ = v7
    return 110
}
func genFn1() uint64 {
    var v8 uint8 = 215
    v8 = 34
    v9 := GenS2{F0: -1, F1: 2147483647}
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]int)
    var v3 uint64 = 763
    for fi10 := 0; fi10 < 2; fi10++ {
        if (1000 > -275) {
            v11 := GenS1{F0: 42, F1: 255, F2: v1}
            v2[fi10] = (fi10 % (((fi10 ^ (fi10 << ((42) & 63)))) | 1))
            _ = v11
        } else {
            if false {
                v1 += -56
                var v12 float32 = 1.5
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 float32 = (float32((3.14159 / (1.5 + 1e10))) - 3.14159)
                    _ = fi13
                    _ = v14
                }
            }
            switch ("\n\t" + ((("a" + ((("" + "") + "\n\t") + "foo bar")) + "a") + "\n\t")) {
            case "a":
                fmt.Println(v3, v3, v1)
            case "hello":
                var v15 float32 = 0.0
                _ = v15
            case "":
                if (!(-1 == 1)) {
                    if false {
                        var v16 float64 = 0.1
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v18 := []int{}
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    v19 := GenS3{F0: 33, F1: 734}
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(fi10, fi20, fi20)
                        v21 := []int{}
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            default:
                if false {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v1 -= (v1 | (v1 * (v1 % ((v1) | 1))))
                        if false {
                            v23 := GenS1{F0: -100, F1: 2, F2: v1}
                            v3 = v3
                            _ = v23
                        } else {
                            v24 := GenS4{F0: fi22, F1: 0}
                            _ = v24
                        }
                        _ = fi22
                    }
                    v1 = (v1 << ((42) & 63))
                } else {
                    switch "" {
                    case "hello":
                        if false {
                            if false {
                                var v25 float32 = 0.0
                                if (!false) {
                                    v26 := []int{}
                                    _ = v26
                                } else {
                                    var v27 uint = 450
                                    _ = v27
                                }
                                var v28 uint = 262
                                _ = v25
                                _ = v28
                            } else {
                                if (!true) {
                                    v29 := []int{fi10}
                                    v30 := GenS5{F0: fi10}
                                    _ = v29
                                    _ = v30
                                }
                                var v31 float32 = 0.1
                                fmt.Println(v3, fi10, fi10)
                                _ = v31
                            }
                            v32 := []int{(fi10 % ((fi10) | 1)), int(v3), (fi10 >> ((255) & 63)), -1}
                            v2[-208] = (fi10 | (-(fi10)))
                            _ = v32
                        } else {
                            v33 := v2[(fi10 ^ fi10)]
                            switch "foo bar" {
                            case "foo bar":
                                delete(v2, -744)
                            case "\n\t":
                                switch 500 {
                                case -9223372036854775808:
                                    v34 := []int{(v33 >> ((1) & 63)), 32767, (^(fi10))}
                                    _ = v34
                                default:
                                    for fi35 := 0; fi35 < 16; fi35++ {
                                        if (!(127 >= 128)) {
                                            v2[(-(fi10))] = fi35
                                            var v36 int64 = -51
                                            if false {
                                                v37 := GenS3{F0: 255, F1: 603}
                                                if true {
                                                    v38 := []int{fi10, 495, fi35, (v33 << ((18446744073709551615) & 63))}
                                                    v39 := GenS6{F0: "", F1: (-(v1)), F2: 42}
                                                    if ((0 <= 42) && ((v36 | (^(v36))) > (v36 ^ v36))) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            _ = len(v38)
                                                        default:
                                                            var v40 float32 = (float32(((0.0 - 1e10) + (1e10 + 0.1))) + float32(-2.25))
                                                            _ = v40
                                                        }
                                                        v33 -= (v33 >> ((0) & 63))
                                                    }
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    var v41 int8 = int8(fi10)
                                                    _ = len(v2)
                                                    var v42 int32 = 42
                                                    fmt.Println(v3)
                                                    _ = v41
                                                    _ = v42
                                                }
                                                var v43 uint16 = 634
                                                _ = v37
                                                _ = v43
                                            }
                                            _ = v36
                                        }
                                        var v44 uint32 = 65535
                                        switch fi35 {
                                        case 96:
                                            var v45 float32 = 1.5
                                            _ = v45
                                        case -1:
                                            fmt.Println(fi35, fi35)
                                        default:
                                            v46 := GenS7{F0: -79, F1: (^(v44)), F2: 4294967295}
                                            _ = v46
                                        }
                                        v1 -= (v1 & v1)
                                        v44 = (v44 + (v44 | v44))
                                        if (!false) {
                                            if ((!false) || false) {
                                                var v47 string = "\n\t"
                                                _ = v47
                                            }
                                            _ = len(v2)
                                        }
                                        _ = len(v2)
                                        switch v33 {
                                        case 4294967295:
                                            v44 += (v44 & 2147483647)
                                        case 789:
                                            var v48 uint8 = 255
                                            v49 := []int{}
                                            _ = v48
                                            _ = v49
                                        default:
                                            v50 := []int{(v33 % ((fi10) | 1))}
                                            _ = v50
                                        }
                                        if false {
                                            fmt.Println(v3, v33)
                                            fmt.Println(fi10)
                                        }
                                        v51 := GenS3{F0: 2, F1: -975}
                                        var v52 string = "a"
                                        _ = fi35
                                        _ = v44
                                        _ = v51
                                        _ = v52
                                    }
                                }
                            case "hello":
                                fmt.Println(fi10, v3, v3)
                            }
                            _ = v33
                        }
                    case "":
                        switch fi10 {
                        case 697:
                            v1 += (v1 - v1)
                        case -78:
                            for fi53 := 0; fi53 < 4; fi53++ {
                                v3 += v3
                                if (!false) {
                                    var v54 int16 = -32768
                                    var v55 string = ""
                                    switch fi10 {
                                    case -766:
                                        v56 := GenS1{F0: v1, F1: 2, F2: 127}
                                        _ = v56
                                    default:
                                        v1 = -1
                                    }
                                    _ = v54
                                    _ = v55
                                } else {
                                    v57 := GenS4{F0: (-(fi10)), F1: 255}
                                    _ = v57
                                }
                                if (!false) {
                                    var v58 uint8 = 1
                                    fmt.Println(fi10, v58, fi10)
                                    _ = v58
                                }
                                _ = fi53
                            }
                        case 1:
                            for fi59 := 0; fi59 < 16; fi59++ {
                                v1 += int8(fi10)
                                switch "\n\t" {
                                case "":
                                    _ = len(v2)
                                    v2[fi10] = (fi59 * (fi59 & 0))
                                    fmt.Println(fi59)
                                default:
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        switch (("a" + "foo bar") + ("" + "")) {
                                        case "foo bar":
                                            v60 := GenS5{F0: fi10}
                                            _ = v60
                                        }
                                    default:
                                        var v61 bool = true
                                        _ = v61
                                    }
                                }
                                if true {
                                    v62 := GenS1{F0: v1, F1: 255, F2: 42}
                                    if (!true) {
                                        v3 = (v3 | genFn1())
                                        v2[fi59] = fi10
                                        switch ("hello" + ("a" + "a")) {
                                        case "":
                                            var v63 int32 = 42
                                            _ = v63
                                        case "a":
                                            fmt.Println(v3, v1, v3)
                                        case "foo bar":
                                            if false {
                                                fmt.Println(fi10)
                                                v3 -= (v3 * v3)
                                            } else {
                                                v1 = v1
                                                v64 := []int{(fi10 + (fi59 >> ((0) & 63)))}
                                                _ = v64
                                            }
                                        }
                                    } else {
                                        v3 += v3
                                        fmt.Println(fi59, v1)
                                        fmt.Println(fi59)
                                    }
                                    var v65 int8 = (^(v1))
                                    _ = v62
                                    _ = v65
                                } else {
                                    switch fi59 {
                                    case 9223372036854775807:
                                        if ((v1 - 6) <= (^(v1))) {
                                            v66 := GenS3{F0: 1, F1: 1}
                                            switch ("" + "a") {
                                            case "foo bar":
                                                fmt.Println(v1, v1)
                                            case "a":
                                                v3 = 937
                                                if true {
                                                    switch fi59 {
                                                    case -743:
                                                        v67 := []int{}
                                                        _ = v67
                                                    case 2147483647:
                                                        switch ("" + (("a" + ((("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + ("\n\t" + ("a" + (("foo bar" + ("" + ("" + ("a" + ("hello" + ("\n\t" + ("hello" + (("\n\t" + "hello") + "\n\t")))))))) + ("hello" + ("" + "hello")))))) + "a")) + ("hello" + ("a" + "")))) {
                                                        case "hello":
                                                            if false {
                                                                v68 := []int{}
                                                                fmt.Println(fi59, fi59, v3)
                                                                switch ("\n\t" + "a") {
                                                                case "a":
                                                                    v2[fi59] = 1
                                                                case "foo bar":
                                                                    v69 := GenS8{F0: (-(v1)), F1: 255, F2: "foo bar"}
                                                                    _ = v69
                                                                default:
                                                                    switch ("\n\t" + "a") {
                                                                    case "":
                                                                        var v70 float32 = 1.5
                                                                        var v71 int16 = 0
                                                                        switch ((("foo bar" + "\n\t") + "a") + ("a" + "\n\t")) {
                                                                        case "foo bar":
                                                                            v70 += v70
                                                                        case "\n\t":
                                                                            v72 := GenS9{F0: fi10}
                                                                            _ = v72
                                                                        default:
                                                                            switch int(v3) {
                                                                            case 118:
                                                                                var v73 bool = true
                                                                                _ = v73
                                                                            case -2147483648:
                                                                                if false {
                                                                                    if true {
                                                                                        if ((65535 == 65535) || (false || (!(128 < 66)))) {
                                                                                            if ((((^(v3)) > v3) || (2 <= (v3 % (((v3 & v3)) | 1)))) && false) {
                                                                                                delete(v2, 32768)
                                                                                            }
                                                                                            var v74 uint16 = 32767
                                                                                            v75 := GenS2{F0: 0, F1: 2147483648}
                                                                                            _ = v74
                                                                                            _ = v75
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v70 = (float32((1e10 / 3.14159)) / (float32(float64(v70)) * (1e10 - (v70 * 0.0))))
                                                                                    var v76 int64 = -9223372036854775808
                                                                                    v70 += (v70 + v70)
                                                                                    _ = v76
                                                                                }
                                                                            case 32767:
                                                                                delete(v2, (^(fi10)))
                                                                            }
                                                                        }
                                                                        _ = v70
                                                                        _ = v71
                                                                    case "hello":
                                                                        v3 = genFn1()
                                                                    default:
                                                                        switch (fi10 << ((255) & 63)) {
                                                                        case 1:
                                                                            v1 = (v1 ^ v1)
                                                                        default:
                                                                            if true {
                                                                                var v77 float64 = 3.14159
                                                                                _ = v77
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v68
                                                            }
                                                        }
                                                    default:
                                                        v1 = v1
                                                    }
                                                } else {
                                                    var v78 uint64 = v3
                                                    v79 := GenS10{F0: fi10, F1: 255, F2: 0}
                                                    _ = v78
                                                    _ = v79
                                                }
                                                v80 := []int{0, fi10, 32768, fi10}
                                                var v81 float32 = (1e10 + 1.5)
                                                var v82 float32 = (v81 + 0.1)
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            case "hello":
                                                v3 -= (v3 & v3)
                                            default:
                                                if false {
                                                    v83 := GenS11{F0: 255, F1: 0, F2: 937}
                                                    v84 := []int{(fi59 << ((605) & 63))}
                                                    v84 = append(v84, 32767)
                                                    v3 += v3
                                                    v85 := GenS3{F0: 0, F1: 32767}
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    v3 -= (v3 >> ((42) & 63))
                                                    if (!(v3 == (v3 / ((v3) | 1)))) {
                                                        v1 = v1
                                                        v1 = int8(v3)
                                                    }
                                                }
                                            }
                                            _ = v66
                                        }
                                    }
                                }
                                _ = fi59
                            }
                        default:
                            var v86 float32 = (0.1 - 1e10)
                            _ = v86
                        }
                    case "\n\t":
                        v1 = (v1 | v1)
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            var v87 string = "\n\t"
                            _ = v87
                        default:
                            if true {
                                v88 := []int{(fi10 / ((fi10) | 1))}
                                v89 := GenS4{F0: fi10, F1: 128}
                                _ = v88
                                _ = v89
                            } else {
                                var v90 int32 = 0
                                _ = v90
                            }
                        }
                    }
                    if true {
                        v1 -= v1
                        fmt.Println(v3)
                    }
                    for fi91 := 0; fi91 < 12; fi91++ {
                        var v92 string = "hello"
                        v93 := []int{(fi91 % ((947) | 1)), 466, fi91}
                        _ = fi91
                        _ = v92
                        _ = v93
                    }
                }
            }
            var v94 uint = 42
            _ = v94
        }
        var v95 int64 = 10
        _ = fi10
        _ = v95
    }
    fmt.Println(v3, v3)
    delete(v2, -9223372036854775808)
    var v96 uint32 = 65535
    v97 := GenS12{F0: v96}
    if (42 < 32767) {
        v98 := GenS12{F0: v96}
        _ = v98
    }
    for fi99 := 0; fi99 < 4; fi99++ {
        for fi100 := 0; fi100 < 14; fi100++ {
            fmt.Println(v96, v96, fi100)
            _ = fi100
        }
        v3 -= v3
        for fi101 := 0; fi101 < 11; fi101++ {
            if (42 != uint16(v96)) {
                v1 += -13
            }
            _ = fi101
        }
        _ = fi99
    }
    v3 = v3
    var v102 int16 = 42
    var v103 float32 = 1.5
    var v104 uint32 = v96
    v105 := GenS13{F0: 0}
    var v106 int16 = int16(v3)
    v103 -= (v103 - (3.14159 + ((1e10 * (v103 / v103)) - ((v103 / float32(1.5)) - ((v103 + (1e10 / float32(-2.25))) * v103)))))
    for fi107 := 0; fi107 < 2; fi107++ {
        if (float32(float64(3.14159)) != float32((1e10 - (1.5 - (((0.0 / (((3.14159 + (3.14159 + -2.25)) + 1.5) - 0.1)) + float64(v103)) / 1e10))))) {
            if true {
                for fi108 := 0; fi108 < 10; fi108++ {
                    fmt.Println(v104, v106, fi107)
                    if ((42 != 1) && (1 < 1)) {
                        v109 := []int{(fi108 % ((fi107) | 1)), (fi107 % (((fi107 << ((65535) & 63))) | 1)), (fi107 | fi108)}
                        if false {
                            switch "" {
                            case "":
                                fmt.Println(v104)
                            case "\n\t":
                                if (0 <= 128) {
                                    v96 += (^(v104))
                                }
                                if (v103 <= 1.5) {
                                    delete(v2, -1)
                                    v102 = (-(v106))
                                }
                            case "a":
                                v110 := GenS14{F0: ("a" + "foo bar"), F1: -1, F2: 42}
                                _ = v110
                            default:
                                v111 := GenS9{F0: fi107}
                                _ = v111
                            }
                            _ = len(v109)
                        } else {
                            if (((1 == 630) && (true && true)) && (!true)) {
                                var v112 uint16 = 198
                                v109 = append(v109, -9223372036854775808)
                                var v113 bool = false
                                switch "" {
                                case "hello":
                                    if (v96 <= v104) {
                                        if true {
                                            if (!v113) {
                                                v114 := GenS15{F0: 255}
                                                v3 = (^(v3))
                                                v1 -= v1
                                                _ = v114
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    v2[542] = fi107
                                                default:
                                                    v115 := GenS1{F0: (-(v1)), F1: v112, F2: (v1 | (v1 * -100))}
                                                    _ = v115
                                                }
                                                v116 := GenS5{F0: fi107}
                                                _ = v116
                                            }
                                            v96 += uint32(v3)
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                if v113 {
                                                    if v113 {
                                                        switch (fi107 & fi107) {
                                                        case -156:
                                                            if ((fi107 != (fi108 | (fi107 << ((12) & 63)))) || ((v112 == (v112 & (v112 - uint16(v3)))) || (float32((3.14159 * 1.5)) > v103))) {
                                                                v117 := GenS7{F0: v1, F1: 667, F2: 970}
                                                                _ = v117
                                                            } else {
                                                                _ = len(v109)
                                                                v113 = (v112 >= 1)
                                                            }
                                                            var v118 float32 = 3.14159
                                                            if false {
                                                                var v119 int = (fi108 / (((fi108 - (fi108 ^ fi108))) | 1))
                                                                _ = v119
                                                            } else {
                                                                if v113 {
                                                                    var v120 int32 = 2147483647
                                                                    _ = v120
                                                                }
                                                            }
                                                            _ = v118
                                                        case 415:
                                                            fmt.Println(v96)
                                                        }
                                                    } else {
                                                        var v121 uint16 = (v112 - v112)
                                                        var v122 uint = 0
                                                        var v123 float32 = v103
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                } else {
                                                    fmt.Println(v96)
                                                }
                                            }
                                            var v124 int8 = int8(v3)
                                            v125 := GenS3{F0: 1, F1: v102}
                                            v124 += v124
                                            _ = v124
                                            _ = v125
                                        }
                                        if v113 {
                                            v103 = v103
                                        }
                                        if false {
                                            var v126 int16 = v106
                                            v112 = (v112 | (v112 & v112))
                                            if ((("\n\t" + ("\n\t" + "hello")) + "") >= "a") {
                                                var v127 uint32 = 2
                                                v128 := GenS16{F0: "foo bar", F1: v96, F2: (fi107 / ((-9223372036854775808) | 1))}
                                                var v129 string = (("hello" + "hello") + "hello")
                                                _ = v127
                                                _ = v128
                                                _ = v129
                                            } else {
                                                if true {
                                                    if (1 == 42) {
                                                        var v130 int32 = -1
                                                        v130 += v130
                                                        fmt.Println(v103, v126)
                                                        _ = v130
                                                    }
                                                    if (((v113 || v113) && (-140 < 42)) || v113) {
                                                        _ = len(v109)
                                                        fmt.Println(fi108)
                                                        _ = len(v109)
                                                    }
                                                    v126 -= (v126 | v126)
                                                } else {
                                                    v96 += v96
                                                }
                                                var v131 int8 = v1
                                                v132 := GenS17{F0: (365 > (^(v112))), F1: -1}
                                                _ = v131
                                                _ = v132
                                            }
                                            _ = v126
                                        }
                                    }
                                case "\n\t":
                                    v133 := v2[fi107]
                                    v134 := GenS16{F0: ("\n\t" + "hello"), F1: (v104 >> ((950) & 63)), F2: (^(fi108))}
                                    _ = v133
                                    _ = v134
                                default:
                                    if (v112 == v112) {
                                        switch ("hello" + "") {
                                        case "foo bar":
                                            var v135 uint8 = 2
                                            _ = v135
                                        case "":
                                            if v113 {
                                                v109 = append(v109, 32768)
                                            } else {
                                                var v136 uint16 = v112
                                                if v113 {
                                                    var v137 uint64 = (v3 << ((4294967295) & 63))
                                                    if (!(!true)) {
                                                        if (127 == -1) {
                                                            v3 = (v3 ^ v137)
                                                        }
                                                        v102 = (^(v106))
                                                    } else {
                                                        switch fi107 {
                                                        case 709:
                                                            v109 = append(v109, (fi107 * fi107))
                                                            var v138 uint = 0
                                                            var v139 int8 = (-(v1))
                                                            _ = v138
                                                            _ = v139
                                                        default:
                                                            v140 := GenS18{F0: 841, F1: v112, F2: (v136 >> ((255) & 63))}
                                                            _ = v140
                                                        }
                                                    }
                                                    _ = v137
                                                }
                                                _ = v136
                                            }
                                        case "a":
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v113, fi108)
                                            case "":
                                                v141 := GenS19{F0: 0}
                                                _ = v141
                                            case "a":
                                                var v142 uint64 = v3
                                                _ = v142
                                            default:
                                                fmt.Println(v96, v102)
                                            }
                                        }
                                    }
                                    var v143 uint16 = v112
                                    if false {
                                        var v144 string = ("" + ("\n\t" + ("a" + "\n\t")))
                                        v145 := GenS14{F0: "\n\t", F1: 2147483647, F2: v143}
                                        if (42 <= (v3 - (v3 % ((v3) | 1)))) {
                                            delete(v2, fi107)
                                            var v146 uint64 = (v3 / ((18446744073709551615) | 1))
                                            _ = v146
                                        }
                                        _ = v144
                                        _ = v145
                                    } else {
                                        v103 = float32(float64(v103))
                                        var v147 uint8 = 0
                                        _ = v147
                                    }
                                    _ = v143
                                }
                                _ = v112
                                _ = v113
                            }
                            _ = len(v2)
                        }
                        var v148 string = "foo bar"
                        v1 += (v1 - v1)
                        _ = v109
                        _ = v148
                    } else {
                        switch "hello" {
                        case "hello":
                            var v149 string = ("" + "a")
                            v150 := GenS5{F0: (-(fi108))}
                            _ = v149
                            _ = v150
                        case "a":
                            v106 += int16(v96)
                        }
                        v151 := GenS12{F0: v96}
                        v152 := GenS16{F0: "hello", F1: (v96 - 65535), F2: fi107}
                        _ = v151
                        _ = v152
                    }
                    var v153 uint8 = 1
                    _ = fi108
                    _ = v153
                }
                switch "hello" {
                case "":
                    switch "" {
                    case "a":
                        switch (fi107 << ((4294967295) & 63)) {
                        case 32768:
                            v154 := GenS20{F0: "hello"}
                            _ = v154
                        case -1:
                            for fi155 := 0; fi155 < 16; fi155++ {
                                v156 := GenS21{F0: (fi155 / ((fi107) | 1)), F1: -59, F2: 87}
                                v3 += (v3 >> ((715) & 63))
                                _ = fi155
                                _ = v156
                            }
                        }
                    default:
                        if false {
                            if (1e10 <= float64(v103)) {
                                v157 := GenS20{F0: ("hello" + (("a" + "a") + "a"))}
                                var v158 bool = true
                                _ = v157
                                _ = v158
                            } else {
                                for fi159 := 0; fi159 < 10; fi159++ {
                                    v160 := []int{(fi107 | (-(fi159)))}
                                    _ = fi159
                                    _ = v160
                                }
                                for fi161 := 0; fi161 < 16; fi161++ {
                                    v162 := []int{(fi107 / ((2147483647) | 1)), (fi161 - (-(fi161))), (fi107 << ((uint(v3)) & 63)), -356}
                                    switch "a" {
                                    case "hello":
                                        v104 = (v104 & 4294967295)
                                    }
                                    _ = fi161
                                    _ = v162
                                }
                            }
                            v163 := GenS10{F0: 32768, F1: 0, F2: 255}
                            _ = v163
                        } else {
                            var v164 int64 = -1
                            if true {
                                if (!((^(v102)) == v102)) {
                                    var v165 float64 = float64(float32((3.14159 + 0.1)))
                                    fmt.Println(v104, v3, v1)
                                    v1 += v1
                                    _ = v165
                                }
                            } else {
                                for fi166 := 0; fi166 < 15; fi166++ {
                                    fmt.Println(v1)
                                    _ = fi166
                                }
                            }
                            _ = v164
                        }
                    }
                case "a":
                    switch 127 {
                    case 20:
                        var v167 uint64 = genFn1()
                        _ = v167
                    case -2147483648:
                        v3 -= 2
                    }
                case "foo bar":
                    v168 := GenS22{F0: 996, F1: 49, F2: (^(fi107))}
                    _ = v168
                default:
                    v106 = (v102 << ((2) & 63))
                }
            }
            for fi169 := 0; fi169 < 2; fi169++ {
                var v170 int = (fi107 >> ((255) & 63))
                _ = fi169
                _ = v170
            }
        }
        if true {
            var v171 uint = 0
            delete(v2, fi107)
            _ = v171
        } else {
            v1 -= v1
            switch "hello" {
            case "a":
                for fi172 := 0; fi172 < 16; fi172++ {
                    v103 = (float32(-2.25) / (v103 + (v103 - 1.5)))
                    v1 = (v1 ^ (v1 * v1))
                    if true {
                        switch ("hello" + "\n\t") {
                        case "hello":
                            fmt.Println(fi107, fi107, fi107)
                        case "a":
                            var v173 float32 = v103
                            _ = v173
                        default:
                            v174 := v2[127]
                            _ = v174
                        }
                        fmt.Println(v96, v106, fi107)
                        v175 := GenS15{F0: 170}
                        v176 := GenS23{F0: 1, F1: 32767}
                        var v177 uint64 = v3
                        var v178 uint8 = 2
                        _ = v175
                        _ = v176
                        _ = v177
                        _ = v178
                    }
                    _ = fi172
                }
            default:
                v106 = 756
            }
            fmt.Println(v102)
        }
        _ = fi107
    }
    v106 = v102
    switch 32768 {
    case -772:
        v102 = (v106 & (v106 | int16(v96)))
    case 889:
        for fi179 := 0; fi179 < 4; fi179++ {
            for fi180 := 0; fi180 < 13; fi180++ {
                var v181 float64 = 3.14159
                _ = fi180
                _ = v181
            }
            _ = fi179
        }
    case 9223372036854775807:
        if (!true) {
            v182 := v2[349]
            _ = v182
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v96
    _ = v97
}
