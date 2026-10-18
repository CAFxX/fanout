package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint; F1 int16; F2 int8 }
type GenS8 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int64; F1 int16; F2 uint }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 bool; F1 bool; F2 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 int; F1 bool; F2 bool }
type GenS15 struct { F0 uint64; F1 string }
type GenS16 struct { F0 uint32; F1 uint16 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int8; F1 uint32 }
type GenS19 struct { F0 uint; F1 uint16 }
func genFn0() uint8 {
    var v7 string = "foo bar"
    var v8 int64 = 9223372036854775807
    fmt.Println(v7, v8)
    var v9 uint64 = 42
    _ = v7
    _ = v8
    _ = v9
    return 140
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 int64 = -353
    v6 := []uint32{(^(v2))}
    v10 := GenS1{F0: -32768, F1: 127}
    switch "a" {
    case "foo bar":
        switch 777 {
        case 127:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v2)
            case "a":
                _ = len(v4)
            case "hello":
                v2 = (^(v2))
            }
        case 2147483647:
            v2 = 984
        case 860:
            v11 := GenS2{F0: 166}
            _ = v11
        }
    default:
        if false {
            v2 = (^(v2))
            var v12 float64 = 0.1
            var v13 uint32 = (v2 & 341)
            _ = v12
            _ = v13
        } else {
            if (!true) {
                var v14 uint = 2
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v14
            } else {
                var v15 uint32 = v2
                for fi16 := 0; fi16 < 20; fi16++ {
                    fmt.Println(fi16, v2, v2)
                    _ = fi16
                }
                var v17 uint8 = 42
                _ = v15
                _ = v17
            }
        }
    }
    switch -1 {
    case -1:
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 string = "hello"
            if (true && ((^(fi18)) < 32768)) {
                v2 += v2
                var v20 uint = 4294967295
                for fi21 := 0; fi21 < 13; fi21++ {
                    switch fi21 {
                    case 623:
                        v6 = append(v6, uint32(v20))
                    case -583:
                        switch v19 {
                        case "\n\t":
                            var v22 float32 = (-2.25 * float32(1e10))
                            _ = v22
                        case "":
                            v5 = v5
                        case "a":
                            if (true || true) {
                                v2 -= 65535
                                v19 = v19
                                fmt.Println(fi21)
                            }
                        default:
                            if false {
                                var v23 uint8 = 0
                                if (!(!false)) {
                                    if (65535 <= 42) {
                                        fmt.Println(v20)
                                    }
                                    v19 = "foo bar"
                                } else {
                                    v3[(^(v2))] = v2
                                    if (v2 >= v2) {
                                        v24 := v1[v19]
                                        _ = v24
                                    }
                                }
                                switch "foo bar" {
                                case "hello":
                                    delete(v4, (v2 & (v2 >> ((18446744073709551615) & 63))))
                                case "":
                                    var v25 uint8 = (v23 * v23)
                                    switch -1 {
                                    case 127:
                                        var v26 int16 = 1
                                        _ = v26
                                    case 630:
                                        delete(v3, 65535)
                                        if true {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        v20 = v20
                                    case 713:
                                        v2 = (^(v2))
                                    }
                                    _ = v25
                                case "foo bar":
                                    switch (fi21 * (fi21 >> ((18446744073709551615) & 63))) {
                                    case 4294967295:
                                        v27 := GenS1{F0: -882, F1: 127}
                                        _ = v27
                                    }
                                }
                                _ = v23
                            } else {
                                v5 += (v5 * (-(v5)))
                                var v28 float32 = -2.25
                                _ = v28
                            }
                        }
                    }
                    var v29 uint = (v20 - 2)
                    v30 := GenS1{F0: -182, F1: 10}
                    if true {
                        v20 -= 511
                    } else {
                        var v31 int8 = 42
                        var v32 int64 = v5
                        v32 += (^(v32))
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS2{F0: 0}
                    if true {
                        fmt.Println(v19)
                        v2 = 65535
                        var v34 int64 = v5
                        _ = v34
                    } else {
                        v35 := v3[v2]
                        switch v19 {
                        case "foo bar":
                            var v36 string = v19
                            _ = v36
                        case "hello":
                            if (false && (0.1 == (3.14159 * float32(((0.1 - float64((float32(float64(((0.0 - -2.25) / 0.1))) * 0.0))) + (float64((1e10 * (1e10 - -2.25))) * 1.5)))))) {
                                switch fi21 {
                                case -868:
                                    v19 = (v19 + ((("hello" + "foo bar") + "hello") + "a"))
                                default:
                                    if (!true) {
                                        v37 := GenS1{F0: -17, F1: 106}
                                        fmt.Println(v2, v5, fi21)
                                        var v38 bool = false
                                        _ = v37
                                        _ = v38
                                    }
                                    switch (-(fi18)) {
                                    case 32768:
                                        v5 -= v5
                                    case -2147483648:
                                        fmt.Println(fi21)
                                    }
                                }
                                v39 := GenS1{F0: 1000, F1: 42}
                                _ = v39
                            }
                        case "a":
                            fmt.Println(fi18, v19, fi18)
                        }
                        _ = v35
                    }
                    v5 = -818
                    _ = fi21
                    _ = v29
                    _ = v30
                    _ = v33
                }
                _ = v20
            } else {
                var v40 int16 = 1000
                _ = v40
            }
            _ = fi18
            _ = v19
        }
    case 42:
        if len(v6) > 0 {
            v6[0] = 65535
        }
    case 0:
        _ = len(v6)
    default:
        var v41 int = -2147483648
        _ = v41
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        v2 = (^(v2))
        _ = fi42
    }
    v2 = (v2 % ((v2) | 1))
    switch ("hello" + (("a" + "") + "hello")) {
    case "a":
        v2 = 780
    case "":
        if (100 <= 42) {
            for fi43 := 0; fi43 < 11; fi43++ {
                var v44 int = (fi43 + fi43)
                var v45 uint32 = 2147483647
                _ = fi43
                _ = v44
                _ = v45
            }
            if false {
                var v46 int64 = v5
                if true {
                    v46 = v46
                    var v47 int8 = 127
                    _ = v47
                } else {
                    switch (("foo bar" + (("foo bar" + "") + "foo bar")) + "\n\t") {
                    case "\n\t":
                        for fi48 := 0; fi48 < 10; fi48++ {
                            if len(v6) > 0 {
                                v6[0] = v2
                            }
                            v1[""] = (("a" + ("a" + (("a" + "\n\t") + ((("\n\t" + "hello") + "foo bar") + "")))) + "")
                            _ = fi48
                        }
                    default:
                        for fi49 := 0; fi49 < 20; fi49++ {
                            var v50 uint8 = 46
                            _ = fi49
                            _ = v50
                        }
                    }
                    v51 := GenS1{F0: 32767, F1: 0}
                    if ((-(v5)) < (v46 + -588)) {
                        delete(v1, "a")
                        fmt.Println(v2)
                        if true {
                            var v52 float32 = 0.0
                            _ = v52
                        }
                    } else {
                        v53 := GenS2{F0: 0}
                        _ = v53
                    }
                    _ = v51
                }
                var v54 int32 = 616
                _ = v46
                _ = v54
            }
        }
    case "hello":
        if (456 > 32767) {
            fmt.Println(v2, v2)
            for fi55 := 0; fi55 < 16; fi55++ {
                if false {
                    v56 := GenS1{F0: 32767, F1: 42}
                    _ = v56
                }
                var v57 int = fi55
                if (-565 == 402) {
                    var v58 int = fi55
                    v59 := GenS1{F0: 102, F1: 127}
                    var v60 uint32 = v2
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    if (true || ((v5 + (v5 / (((v5 / (((v5 + (v5 + v5))) | 1))) | 1))) < v5)) {
                        v61 := GenS3{F0: 9223372036854775807, F1: 1, F2: 356}
                        _ = v61
                    } else {
                        if true {
                            for fi62 := 0; fi62 < 11; fi62++ {
                                if false {
                                    if len(v6) > 0 {
                                        v6[0] = v2
                                    }
                                }
                                _ = fi62
                            }
                            switch "\n\t" {
                            case "hello":
                                var v63 uint16 = 696
                                _ = v63
                            case "foo bar":
                                v64 := GenS4{F0: v57, F1: 2, F2: v2}
                                _ = v64
                            default:
                                switch 42 {
                                case -857:
                                    if (v2 == v2) {
                                        var v65 string = "hello"
                                        var v66 float64 = 3.14159
                                        _ = v65
                                        _ = v66
                                    }
                                case 2147483647:
                                    var v67 int64 = (-(v5))
                                    _ = v67
                                case 127:
                                    switch (^(fi55)) {
                                    case 4294967295:
                                        _ = len(v6)
                                    case 2147483647:
                                        var v68 uint = 42
                                        _ = v68
                                    default:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                            }
                        } else {
                            v69 := GenS5{F0: (375 > 18446744073709551615), F1: (v2 << ((0) & 63)), F2: genFn0()}
                            v2 = v2
                            switch v57 {
                            case -462:
                                v5 -= v5
                            case 32767:
                                var v70 float64 = 3.14159
                                _ = v70
                            }
                            _ = v69
                        }
                        var v71 uint64 = 783
                        fmt.Println(fi55, v5, v2)
                        _ = v71
                    }
                    fmt.Println(v57)
                }
                _ = fi55
                _ = v57
            }
            if true {
                switch 32768 {
                case 32768:
                    v72 := GenS5{F0: true, F1: 0, F2: 78}
                    _ = v72
                default:
                    v2 = (v2 >> ((0) & 63))
                }
                for fi73 := 0; fi73 < 2; fi73++ {
                    v2 = v2
                    fmt.Println(fi73, v5)
                    for fi74 := 0; fi74 < 13; fi74++ {
                        switch (fi74 + 2147483647) {
                        case -934:
                            v5 -= (v5 / (((-(v5))) | 1))
                        case 1:
                            switch ("\n\t" + "") {
                            case "":
                                if (true || false) {
                                    if true {
                                        v5 -= v5
                                        var v75 float64 = 0.0
                                        v6 = append(v6, v2)
                                        switch (fi74 / ((fi74) | 1)) {
                                        case 291:
                                            v76 := GenS6{F0: 52, F1: (v2 * (v2 - (v2 % (((v2 << ((4294967295) & 63))) | 1))))}
                                            _ = v76
                                        case 32768:
                                            var v77 int64 = 1
                                            _ = v77
                                        default:
                                            var v78 int8 = 15
                                            delete(v1, "\n\t")
                                            var v79 int = fi74
                                            _ = v78
                                            _ = v79
                                        }
                                        _ = v75
                                    }
                                    var v80 uint8 = 0
                                    _ = v80
                                }
                            case "\n\t":
                                delete(v4, (v2 & v2))
                            case "hello":
                                var v81 int32 = -1
                                _ = v81
                            default:
                                v5 = (^(v5))
                            }
                            var v82 uint8 = 184
                            _ = v82
                        default:
                            if false {
                                v2 = v2
                                _ = len(v1)
                                switch 1 {
                                case 127:
                                    fmt.Println(v5, fi74, v5)
                                case 2147483647:
                                    switch "\n\t" {
                                    case "":
                                        if (3.14159 > 0.0) {
                                            if (33 == 65535) {
                                                if (!(0 >= 0)) {
                                                    v2 = v2
                                                    if (1 < 838) {
                                                        var v83 int16 = 560
                                                        switch (fi74 - fi73) {
                                                        case 32767:
                                                            if (-2147483648 < 42) {
                                                                var v84 uint8 = 128
                                                                fmt.Println(v83, fi73, v83)
                                                                _ = v84
                                                            }
                                                        }
                                                        if (!false) {
                                                            switch ("foo bar" + ("\n\t" + ("" + ("a" + (("foo bar" + "") + "hello"))))) {
                                                            case "":
                                                                v85 := GenS7{F0: 561, F1: v83, F2: 35}
                                                                _ = v85
                                                            case "\n\t":
                                                                fmt.Println(fi73, v2, fi73)
                                                            case "hello":
                                                                v86 := GenS4{F0: (fi74 + (fi73 >> ((0) & 63))), F1: 65535, F2: v2}
                                                                _ = v86
                                                            default:
                                                                if false {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = v2
                                                                    }
                                                                }
                                                                v87 := GenS5{F0: (620 < (v2 << ((1) & 63))), F1: 2147483648, F2: 254}
                                                                _ = v87
                                                            }
                                                            if false {
                                                                fmt.Println(fi73)
                                                            } else {
                                                                v5 += (-(v5))
                                                            }
                                                        }
                                                        _ = v83
                                                    } else {
                                                        var v88 string = ""
                                                        _ = v88
                                                    }
                                                    var v89 string = "foo bar"
                                                    _ = v89
                                                } else {
                                                    v5 = 1
                                                    v90 := GenS8{F0: ((30 > 81) && (1 > 37)), F1: 42, F2: 1}
                                                    _ = v90
                                                }
                                                if true {
                                                    delete(v4, (v2 >> ((0) & 63)))
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v91 int32 = 2147483647
                                                        _ = v91
                                                    case "foo bar":
                                                        v2 = (v2 % ((v2) | 1))
                                                    case "hello":
                                                        v92 := GenS5{F0: true, F1: (v2 + v2), F2: 0}
                                                        _ = v92
                                                    default:
                                                        v5 -= (v5 << ((370) & 63))
                                                    }
                                                    switch (fi73 | fi73) {
                                                    case 0:
                                                        var v93 int16 = 4
                                                        _ = v93
                                                    default:
                                                        var v94 bool = (-128 > 100)
                                                        v94 = ((v5 * v5) >= 0)
                                                        var v95 bool = v94
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                            } else {
                                                v5 -= (-(v5))
                                            }
                                            var v96 bool = false
                                            var v97 int32 = 385
                                            _ = v96
                                            _ = v97
                                        }
                                    case "hello":
                                        v2 += v2
                                    default:
                                        _ = len(v6)
                                    }
                                case 0:
                                    v98 := GenS2{F0: 42}
                                    switch ("hello" + ((("\n\t" + "hello") + "") + "foo bar")) {
                                    case "":
                                        v5 = 9223372036854775807
                                    case "hello":
                                        var v99 int8 = 1
                                        _ = v99
                                    case "a":
                                        v5 = v5
                                    default:
                                        if (((v2 | (v2 + v2)) < (v2 ^ (v2 - v2))) && true) {
                                            v100 := GenS9{F0: (v5 * v5), F1: 32767, F2: 0}
                                            if len(v6) > 0 {
                                                v6[0] = (v2 + (v2 + 65535))
                                            }
                                            v2 -= v2
                                            if false {
                                                delete(v3, 0)
                                                switch ("foo bar" + "a") {
                                                case "":
                                                    var v101 int8 = 127
                                                    _ = v101
                                                case "a":
                                                    if false {
                                                        _ = len(v6)
                                                    } else {
                                                        if true {
                                                            if false {
                                                                v102 := v1["hello"]
                                                                _ = v102
                                                            } else {
                                                                v5 = v5
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v103 := GenS10{F0: genFn0(), F1: true}
                                                                    _ = v103
                                                                case "foo bar":
                                                                    var v104 float32 = 1.5
                                                                    _ = v104
                                                                }
                                                            }
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                default:
                                                    if ((fi73 + fi74) >= fi73) {
                                                        if false {
                                                            var v105 uint16 = 32767
                                                            _ = v105
                                                        }
                                                        var v106 uint8 = 128
                                                        fmt.Println(v2)
                                                        v5 = (v5 & 1)
                                                        switch "a" {
                                                        case "\n\t":
                                                            v106 -= 2
                                                        case "a":
                                                            if (-97 != 0) {
                                                                v2 = v2
                                                                var v107 float32 = 3.14159
                                                                v108 := v1[""]
                                                                if (4294967295 <= (v2 % (((v2 % ((2147483647) | 1))) | 1))) {
                                                                    fmt.Println(v2, v2)
                                                                }
                                                                switch fi74 {
                                                                case 2147483647:
                                                                    if ((v108 <= "a") || (275 > 0)) {
                                                                        if ((v5 / (((v5 ^ 0)) | 1)) > 9223372036854775807) {
                                                                            var v109 uint64 = 18446744073709551615
                                                                            _ = v109
                                                                        }
                                                                    } else {
                                                                        if (false || (803 <= 65535)) {
                                                                            v108 = "a"
                                                                        }
                                                                    }
                                                                case 1:
                                                                    var v110 int32 = 1
                                                                    _ = v110
                                                                default:
                                                                    if (32767 >= 1) {
                                                                        if (293 <= 913) {
                                                                            if false {
                                                                                var v111 uint64 = 0
                                                                                _ = v111
                                                                            }
                                                                            if len(v6) > 0 {
                                                                                v6[0] = v2
                                                                            }
                                                                            switch v108 {
                                                                            case "\n\t":
                                                                                if (!(1.5 >= float64(-2.25))) {
                                                                                    var v112 float32 = v107
                                                                                    _ = v112
                                                                                }
                                                                                v107 = ((v107 / float32(-2.25)) * (0.1 - float32(float64(3.14159))))
                                                                            case "a":
                                                                                fmt.Println(fi73, fi73)
                                                                            default:
                                                                                if true {
                                                                                    v113 := GenS9{F0: v5, F1: int16(fi74), F2: uint(v106)}
                                                                                    fmt.Println(v107, v5)
                                                                                    var v114 int = (fi73 * fi74)
                                                                                    v114 = fi74
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                } else {
                                                                                    var v115 uint16 = 32767
                                                                                    fmt.Println(v106, fi74)
                                                                                    _ = v115
                                                                                }
                                                                            }
                                                                            if (((!((^(v2)) != v2)) && ((1.5 / ((3.14159 / 1.5) * float64(v107))) == float64(v107))) || false) {
                                                                                v2 = v2
                                                                                fmt.Println(fi74)
                                                                                v106 = (v106 / ((v106) | 1))
                                                                            }
                                                                        }
                                                                        v116 := GenS11{F0: (142 != v106), F1: ((365 < -126) && (true && false)), F2: (fi74 | -135)}
                                                                        _ = v116
                                                                    } else {
                                                                        var v117 string = v108
                                                                        if true {
                                                                            v118 := GenS12{F0: ((v117 + (v117 + v108)) + v117)}
                                                                            var v119 uint32 = v2
                                                                            _ = v118
                                                                            _ = v119
                                                                        } else {
                                                                            if false {
                                                                                if ((!true) && (true && (true || true))) {
                                                                                    var v120 int32 = -14
                                                                                    v121 := GenS11{F0: true, F1: true, F2: 0}
                                                                                    switch fi73 {
                                                                                    case 652:
                                                                                        v122 := GenS13{F0: 18446744073709551615, F1: (v2 * (v2 & 2))}
                                                                                        _ = v122
                                                                                    case -1:
                                                                                        var v123 int8 = -1
                                                                                        _ = v123
                                                                                    case 32767:
                                                                                        v3[v2] = 897
                                                                                    }
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                }
                                                                            } else {
                                                                                v124 := GenS14{F0: fi73, F1: false, F2: (!false)}
                                                                                _ = v124
                                                                            }
                                                                            v125 := GenS3{F0: fi74, F1: v106, F2: 714}
                                                                            v106 = (v106 * 40)
                                                                            _ = v125
                                                                        }
                                                                        _ = v117
                                                                    }
                                                                }
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                        case "foo bar":
                                                            v106 -= genFn0()
                                                            switch (^(fi73)) {
                                                            case 32768:
                                                                var v126 uint = 65535
                                                                _ = v126
                                                            case 128:
                                                                v127 := GenS7{F0: 310, F1: 0, F2: -62}
                                                                _ = v127
                                                            case 127:
                                                                v128 := GenS11{F0: true, F1: true, F2: (fi73 % ((fi73) | 1))}
                                                                _ = v128
                                                            default:
                                                                fmt.Println(v2, fi74, v2)
                                                            }
                                                        }
                                                        v5 += v5
                                                        var v129 int16 = int16(v5)
                                                        var v130 int = (fi74 >> ((255) & 63))
                                                        _ = v106
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v131 := GenS14{F0: fi74, F1: false, F2: true}
                                                    _ = v131
                                                }
                                            }
                                            v132 := GenS15{F0: 632, F1: ("foo bar" + "\n\t")}
                                            var v133 uint = 18446744073709551615
                                            switch (((("" + "") + "a") + "") + ("a" + ("a" + "\n\t"))) {
                                            case "foo bar":
                                                if false {
                                                    v2 = (^(v2))
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v134 uint16 = 32767
                                                        _ = v134
                                                    }
                                                }
                                            case "\n\t":
                                                if true {
                                                    v135 := GenS16{F0: 297, F1: 1}
                                                    v136 := GenS17{F0: (v5 % (((v5 * (v5 << ((v133) & 63)))) | 1))}
                                                    if (("foo bar" < "hello") && ((false || false) || (127 != 42))) {
                                                        fmt.Println(v2, fi74, fi73)
                                                    }
                                                    _ = len(v6)
                                                    var v137 bool = (127 <= uint8(v133))
                                                    _ = v135
                                                    _ = v136
                                                    _ = v137
                                                }
                                            }
                                            _ = v100
                                            _ = v132
                                            _ = v133
                                        }
                                    }
                                    var v138 int64 = int64(v2)
                                    _ = v138
                                    _ = v98
                                }
                            }
                        }
                        var v139 uint32 = (^(v2))
                        var v140 float64 = 1.5
                        v141 := GenS16{F0: v139, F1: 32767}
                        _ = fi74
                        _ = v139
                        _ = v140
                        _ = v141
                    }
                    _ = fi73
                }
                var v142 float32 = (1e10 / (((1e10 + 0.1) + (0.1 + 0.0)) / 1.0))
                _ = v142
            } else {
                if (v5 > v5) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
        }
    default:
        if (1.5 > 0.0) {
            v143 := GenS14{F0: 32768, F1: false, F2: (true || (42 != 42))}
            var v144 uint64 = 1
            if true {
                v145 := GenS18{F0: 42, F1: (v2 - v2)}
                _ = v145
            }
            _ = v143
            _ = v144
        }
    }
    if false {
        if (-2.25 < 3.14159) {
            fmt.Println(v5, v2)
        }
        if false {
            v6 = append(v6, (v2 & v2))
            v146 := GenS3{F0: 4294967295, F1: 255, F2: 0}
            fmt.Println(v5)
            _ = v146
        }
        var v147 uint8 = 99
        _ = v147
    } else {
        var v148 uint32 = (v2 << ((65535) & 63))
        if (!(((0.1 / 1.0) + 0.0) == 1.5)) {
            var v149 int8 = -49
            var v150 int64 = 516
            for fi151 := 0; fi151 < 4; fi151++ {
                v2 -= (v2 - 0)
                var v152 int64 = v150
                if true {
                    fmt.Println(v149, v2, fi151)
                } else {
                    var v153 int8 = (v149 & v149)
                    if false {
                        for fi154 := 0; fi154 < 7; fi154++ {
                            var v155 int = 2147483647
                            v156 := GenS15{F0: 328, F1: ""}
                            switch 412 {
                            case 32768:
                                if (42 <= 0) {
                                    v5 -= (v5 / ((0) | 1))
                                } else {
                                    if (false || ((false || true) && (false && true))) {
                                        fmt.Println(v149, fi154)
                                        v149 += (v149 - v153)
                                    }
                                }
                            default:
                                v153 = (^(v149))
                            }
                            _ = fi154
                            _ = v155
                            _ = v156
                        }
                    }
                    _ = v153
                }
                _ = fi151
                _ = v152
            }
            _ = v149
            _ = v150
        }
        for fi157 := 0; fi157 < 17; fi157++ {
            if len(v6) > 0 {
                v6[0] = (^(v2))
            }
            v158 := GenS12{F0: (("hello" + "hello") + "hello")}
            v5 = v5
            _ = fi157
            _ = v158
        }
        _ = v148
    }
    v5 = (v5 - v5)
    for fi159 := 0; fi159 < 20; fi159++ {
        v6 = append(v6, (v2 * 735))
        _ = fi159
    }
    v2 += 4294967295
    if (0 <= 518) {
        v160 := GenS19{F0: 0, F1: 0}
        var v161 int64 = 1
        var v162 float32 = 0.0
        _ = v160
        _ = v161
        _ = v162
    }
    switch 32768 {
    case 128:
        var v163 int8 = 51
        _ = v163
    }
    var v164 int32 = -1
    var v165 int = -11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v164
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
