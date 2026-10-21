package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 int8; F1 uint64; F2 int16 }
type GenS11 struct { F0 int32; F1 int16; F2 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 bool; F1 int }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int64; F1 uint8 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS20 struct { F0 uint16 }

func main() {
    v1 := []int{2147483647, 32767, -2147483648, 127}
    v2 := []uint32{980, 762}
    var v3 float64 = 0.1
    v4 := []string{"", ("a" + "a")}
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 11; fi7++ {
        v3 = v3
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(v3, fi8, fi7)
            if false {
                if (!true) {
                    var v9 uint8 = 128
                    var v10 uint = 71
                    _ = v10
                    _ = v9
                }
            }
            _ = fi8
        }
        var v11 uint64 = 129
        _ = fi7
        _ = v11
    }
    var v12 uint16 = 1
    fmt.Println(v3)
    if true {
        v3 -= v3
        v12 += 65535
        v13 := v5[1]
        _ = v13
    } else {
        var v14 uint8 = 127
        fmt.Println(v14, v12)
        _ = v14
    }
    switch "\n\t" {
    case "":
        v3 -= (1e10 - v3)
    case "foo bar":
        if true {
            var v15 bool = false
            v15 = (!((v3 >= (v3 - (float64((3.14159 / (-2.25 + 3.14159))) / (v3 / (((1e10 * v3) * v3) * float64((((0.1 / (3.14159 / 1.5)) + float32(1.5)) * -2.25))))))) || v15))
            var v16 uint = 2
            _ = v15
            _ = v16
        }
    case "hello":
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                v3 += v3
                var v19 int32 = -2147483648
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    var v20 float32 = (((-2.25 - 0.1) + -2.25) + 0.1)
    v21 := v5[uint32(v12)]
    v12 += (v12 + (v12 ^ 65535))
    switch "foo bar" {
    case "a":
        switch "foo bar" {
        case "a":
            v21 += (v21 >> ((18446744073709551615) & 63))
        case "foo bar":
            var v22 uint8 = 0
            _ = v22
        case "\n\t":
            v23 := GenS1{F0: false}
            _ = v23
        default:
            var v24 bool = false
            _ = v24
        }
    case "":
        v21 += (v21 << ((976) & 63))
    default:
        for fi25 := 0; fi25 < 4; fi25++ {
            if true {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v21 += (v21 << ((18446744073709551615) & 63))
                    v3 = ((float64(v20) - v3) - 0.1)
                    var v27 int32 = 0
                    fmt.Println(fi26)
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v20, v20)
                for fi28 := 0; fi28 < 13; fi28++ {
                    switch "a" {
                    case "a":
                        var v29 int16 = 1000
                        _ = v29
                    case "":
                        switch "" {
                        case "foo bar":
                            _ = len(v4)
                        default:
                            var v30 uint32 = 0
                            var v31 bool = true
                            _ = v30
                            _ = v31
                        }
                    case "foo bar":
                        v32 := GenS2{F0: uint32(v12), F1: uint64(v12), F2: (v21 ^ fi28)}
                        var v33 bool = (v21 < v21)
                        _ = v32
                        _ = v33
                    default:
                        v34 := GenS3{F0: 2147483648, F1: (v20 > (((v20 - 3.14159) - v20) - (float32(v3) - float32((1e10 / (v3 + v3))))))}
                        _ = v34
                    }
                    v6[65535] = fi25
                    v35 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: v21}
                    _ = fi28
                    _ = v35
                }
            } else {
                var v36 int8 = 84
                fmt.Println(v3, v36)
                _ = v36
            }
            fmt.Println(v21, fi25)
            fmt.Println(v21)
            _ = fi25
        }
    }
    v3 = (v3 / float64((((v20 + v20) + v20) + v20)))
    fmt.Println(v21, v12)
    v37 := GenS2{F0: 773, F1: 0, F2: int(v12)}
    var v38 int16 = 32767
    for fi39 := 0; fi39 < 16; fi39++ {
        if (true || false) {
            if (((v38 + (v38 * v38)) < 231) || true) {
                fmt.Println(v3, v12)
                switch ((("a" + "hello") + "hello") + "hello") {
                case "":
                    switch fi39 {
                    case 9223372036854775807:
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            var v40 int = 32767
                            _ = v40
                        case "foo bar":
                            for fi41 := 0; fi41 < 4; fi41++ {
                                if false {
                                    switch ("hello" + "\n\t") {
                                    case "foo bar":
                                        var v42 int = 128
                                        _ = v42
                                    case "hello":
                                        switch (-(fi41)) {
                                        case 1:
                                            v21 -= (fi41 & (fi41 | (fi39 / ((fi39) | 1))))
                                        default:
                                            if false {
                                                v38 = (v38 & (v38 % ((v38) | 1)))
                                                var v43 uint32 = 2147483648
                                                _ = v43
                                            } else {
                                                var v44 int32 = -308
                                                _ = v44
                                            }
                                        }
                                    case "":
                                        switch "" {
                                        case "foo bar":
                                            var v45 int32 = 0
                                            _ = v45
                                        default:
                                            if true {
                                                if (!true) {
                                                    var v46 uint32 = 2147483647
                                                    if len(v2) > 0 {
                                                        v2[0] = v46
                                                    }
                                                    v47 := GenS2{F0: (v46 / (((v46 ^ v46)) | 1)), F1: uint64(v46), F2: v21}
                                                    v38 = v38
                                                    v48 := GenS4{F0: (v21 ^ int(v46)), F1: -128}
                                                    if (true && false) {
                                                        _ = len(v2)
                                                        var v49 int = -648
                                                        v3 += v3
                                                        _ = v49
                                                    }
                                                    if true {
                                                        if true {
                                                            if (v3 == v3) {
                                                                v20 = 1.5
                                                            } else {
                                                                delete(v6, v46)
                                                            }
                                                            var v50 int32 = 42
                                                            var v51 int32 = 0
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        v52 := GenS1{F0: false}
                                                        _ = v52
                                                    } else {
                                                        v20 = 1.5
                                                        if false {
                                                            var v53 int64 = -1
                                                            if (false || true) {
                                                                v12 = v12
                                                                v54 := GenS3{F0: (v46 << ((255) & 63)), F1: true}
                                                                fmt.Println(v3)
                                                                _ = v54
                                                            }
                                                            _ = v53
                                                        }
                                                        var v55 string = "hello"
                                                        _ = v55
                                                    }
                                                    fmt.Println(v38)
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if true {
                                                        switch -2147483648 {
                                                        case -520:
                                                            var v56 int32 = 0
                                                            _ = v56
                                                        case 32767:
                                                            var v57 uint8 = 112
                                                            _ = v57
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            if false {
                                                                v21 = fi41
                                                            } else {
                                                                v58 := GenS2{F0: 0, F1: 18446744073709551615, F2: fi41}
                                                                fmt.Println(v20, v12, fi39)
                                                                v20 += (v20 / v20)
                                                                v3 += (float64(v20) - 3.14159)
                                                                if false {
                                                                    switch (fi41 >> ((218) & 63)) {
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v12 -= v12
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (fi39 / (((fi39 >> ((934) & 63))) | 1))
                                                                            }
                                                                        } else {
                                                                            if ((!(-412 <= -2147483648)) || false) {
                                                                                _ = len(v4)
                                                                                fmt.Println(fi41, v21)
                                                                            }
                                                                            var v59 float32 = 0.1
                                                                            v59 += -2.25
                                                                            _ = v59
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v60 bool = true
                                                                            if v60 {
                                                                                v61 := GenS5{F0: 555}
                                                                                var v62 int64 = 9223372036854775807
                                                                                _ = v61
                                                                                _ = v62
                                                                            } else {
                                                                                v63 := GenS3{F0: 2147483647, F1: true}
                                                                                v12 += (v12 >> ((uint(v12)) & 63))
                                                                                _ = v63
                                                                            }
                                                                            v21 = (-(fi41))
                                                                            if (int16(fi41) < 32767) {
                                                                                v20 = v20
                                                                                var v64 int16 = (v38 | (v38 % ((42) | 1)))
                                                                                v20 = v20
                                                                                _ = v64
                                                                            }
                                                                            _ = v60
                                                                        }
                                                                        v3 = 0.1
                                                                    }
                                                                    fmt.Println(fi39)
                                                                    var v65 uint32 = 0
                                                                    _ = v65
                                                                }
                                                                _ = v58
                                                            }
                                                        case "":
                                                            v20 = v20
                                                        }
                                                        var v66 bool = ((!(true || false)) || true)
                                                        _ = v66
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = (-(fi39))
                                                    }
                                                    var v67 uint16 = v12
                                                    _ = v67
                                                }
                                                _ = len(v5)
                                                v12 += v12
                                            }
                                        }
                                    }
                                    v68 := GenS3{F0: 42, F1: false}
                                    v69 := GenS6{F0: (!false)}
                                    _ = v68
                                    _ = v69
                                } else {
                                    var v70 uint32 = 1
                                    v71 := GenS2{F0: v70, F1: 18446744073709551615, F2: 350}
                                    if true {
                                        v12 = v12
                                    } else {
                                        var v72 bool = (!true)
                                        v73 := GenS7{F0: (v21 & (fi39 % (((v21 << ((0) & 63))) | 1))), F1: int32(v21)}
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                var v74 int = fi39
                                _ = fi41
                                _ = v74
                            }
                        }
                    case -1:
                        var v75 string = ((("" + "a") + "") + "foo bar")
                        _ = v75
                    case 32767:
                        if ((fi39 + 9223372036854775807) > v21) {
                            if true {
                                fmt.Println(fi39)
                                if (65535 == (^(v12))) {
                                    if ("hello" < "\n\t") {
                                        var v76 uint16 = v12
                                        switch "hello" {
                                        case "foo bar":
                                            for fi77 := 0; fi77 < 8; fi77++ {
                                                v38 += 0
                                                if true {
                                                    var v78 uint8 = 1
                                                    fmt.Println(v12)
                                                    v12 = 597
                                                    switch (fi39 ^ (v21 << ((2) & 63))) {
                                                    case 32768:
                                                        switch "hello" {
                                                        case "\n\t":
                                                            switch ("" + "") {
                                                            case "":
                                                                fmt.Println(v12, v12, v38)
                                                            }
                                                        case "a":
                                                            v12 = v76
                                                            fmt.Println(v78, fi39)
                                                        default:
                                                            delete(v5, 1)
                                                            switch (fi39 / ((fi39) | 1)) {
                                                            case 632:
                                                                v78 += v78
                                                            case -320:
                                                                v21 = 127
                                                            case 127:
                                                                switch v21 {
                                                                case 1:
                                                                    v20 = float32(float64(v20))
                                                                case 609:
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        fmt.Println(v38)
                                                                    case "a":
                                                                        _ = len(v6)
                                                                    default:
                                                                        if false {
                                                                            v78 -= 71
                                                                            var v79 uint32 = 1
                                                                            var v80 int16 = (v38 * (^(v38)))
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                case 4294967295:
                                                                    if false {
                                                                        v1 = append(v1, (fi39 - 0))
                                                                        var v81 float64 = (v3 / float64(float32((0.0 / ((v3 - v3) / (float64(v20) * (v3 / (0.0 / (v3 * ((0.0 - v3) - ((((((((0.1 * v3) + ((v3 + v3) / (((float64(v20) + 0.0) / v3) + 1.5))) * -2.25) - float64(0.1)) * v3) * v3) * (v3 + 1.5)) * v3)))))))))))
                                                                        v82 := GenS8{F0: -128}
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                default:
                                                                    if false {
                                                                        var v83 uint8 = (v78 * v78)
                                                                        v84 := GenS9{F0: 454, F1: 276}
                                                                        switch ("\n\t" + (("" + ("" + "")) + "foo bar")) {
                                                                        case "":
                                                                            var v85 uint64 = 459
                                                                            _ = v85
                                                                        case "foo bar":
                                                                            v21 -= v21
                                                                        default:
                                                                            switch v21 {
                                                                            case 128:
                                                                                v86 := v5[2147483648]
                                                                                _ = v86
                                                                            case 4294967295:
                                                                                v76 = (v12 / (((v76 - (v76 + (v12 >> ((18446744073709551615) & 63))))) | 1))
                                                                            case -9223372036854775808:
                                                                                if true {
                                                                                    v20 -= v20
                                                                                    fmt.Println(v20)
                                                                                } else {
                                                                                    if false {
                                                                                        v6[2147483647] = v21
                                                                                        fmt.Println(v3)
                                                                                    }
                                                                                    var v87 int8 = 127
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v83
                                                                        _ = v84
                                                                    } else {
                                                                        fmt.Println(v38, v38, v20)
                                                                        v20 = -2.25
                                                                    }
                                                                }
                                                            default:
                                                                var v88 string = "hello"
                                                                _ = v88
                                                            }
                                                            var v89 int32 = 2147483647
                                                            v90 := GenS5{F0: (v76 | v76)}
                                                            var v91 int8 = -128
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                    }
                                                    switch ("\n\t" + ("a" + "")) {
                                                    case "\n\t":
                                                        switch 128 {
                                                        case -2147483648:
                                                            switch fi77 {
                                                            case 32768:
                                                                v20 -= v20
                                                                v20 = 0.0
                                                            case -229:
                                                                fmt.Println(v78, v78, fi77)
                                                            case 367:
                                                                v21 = v21
                                                            default:
                                                                if false {
                                                                    if false {
                                                                        if true {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            switch ("\n\t" + "") {
                                                                            case "hello":
                                                                                v20 -= v20
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(fi77, v12)
                                                                                var v92 uint = 2
                                                                                _ = v92
                                                                            } else {
                                                                                if (!(false && (!(((1e10 * float32(v3)) / float32(((((1e10 + ((1e10 - (v3 + 0.0)) - v3)) * 1.5) * v3) * (v3 - (1.5 + v3))))) > v20)))) {
                                                                                    v12 += (v12 - (v76 + (v76 << ((65535) & 63))))
                                                                                }
                                                                                fmt.Println(v21)
                                                                            }
                                                                            v93 := v6[2147483647]
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    v1 = append(v1, fi77)
                                                                    if true {
                                                                        v12 += 554
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                if (-9223372036854775808 != 1) {
                                                                    v94 := GenS10{F0: -1, F1: 4294967295, F2: 32767}
                                                                    _ = len(v2)
                                                                    v78 = (v78 + (v78 * (v78 << ((964) & 63))))
                                                                    _ = v94
                                                                } else {
                                                                    _ = len(v5)
                                                                    if (1 <= 2) {
                                                                        var v95 uint16 = (v76 * (v76 >> ((313) & 63)))
                                                                        fmt.Println(v78, v3)
                                                                        _ = v95
                                                                    } else {
                                                                        var v96 bool = ((v3 * 0.1) != 1.5)
                                                                        _ = v96
                                                                    }
                                                                }
                                                                if false {
                                                                    v3 += (v3 - (v3 * 0.0))
                                                                    if (!false) {
                                                                        var v97 float64 = v3
                                                                        _ = v97
                                                                    }
                                                                    v3 += (1e10 / v3)
                                                                } else {
                                                                    _ = len(v4)
                                                                    switch (v21 - fi39) {
                                                                    case 2147483647:
                                                                        var v98 string = "a"
                                                                        _ = v98
                                                                    default:
                                                                        var v99 int8 = 0
                                                                        _ = v99
                                                                    }
                                                                }
                                                            } else {
                                                                if true {
                                                                    v5[117] = (fi39 - fi39)
                                                                    v3 = 0.0
                                                                    fmt.Println(fi39, v38)
                                                                    if true {
                                                                        v5[2147483647] = fi39
                                                                    } else {
                                                                        v100 := GenS11{F0: 42, F1: v38, F2: (v12 ^ v76)}
                                                                        v101 := GenS10{F0: 96, F1: 18446744073709551615, F2: v38}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    v102 := GenS12{F0: 18446744073709551615}
                                                                    _ = v102
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if ((v78 << ((uint(v76)) & 63)) < (^(v78))) {
                                                        fmt.Println(v12, v3)
                                                        v20 = ((0.1 + ((float32(0.1) * v20) - (3.14159 / v20))) + v20)
                                                        if true {
                                                            var v103 int16 = (v38 / ((v38) | 1))
                                                            _ = len(v6)
                                                            _ = v103
                                                        } else {
                                                            v12 -= 90
                                                            var v104 uint64 = 42
                                                            fmt.Println(v104, v3)
                                                            if true {
                                                                fmt.Println(v21, fi77, v104)
                                                                var v105 int = fi77
                                                                var v106 int32 = -2147483648
                                                                _ = v105
                                                                _ = v106
                                                            } else {
                                                                v107 := v6[uint32(v78)]
                                                                switch "foo bar" {
                                                                case "":
                                                                    v108 := GenS5{F0: 588}
                                                                    _ = v108
                                                                case "a":
                                                                    var v109 uint32 = 2
                                                                    _ = v109
                                                                case "\n\t":
                                                                    var v110 uint64 = 212
                                                                    _ = v110
                                                                default:
                                                                    v111 := v5[4294967295]
                                                                    _ = v111
                                                                }
                                                                _ = v107
                                                            }
                                                            _ = v104
                                                        }
                                                    }
                                                    v76 += (v12 >> ((2) & 63))
                                                    v112 := GenS6{F0: false}
                                                    if true {
                                                        var v113 uint64 = 4294967295
                                                        var v114 uint = 0
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                    v115 := GenS9{F0: 42, F1: 689}
                                                    _ = v112
                                                    _ = v115
                                                    _ = v78
                                                }
                                                _ = fi77
                                            }
                                        case "\n\t":
                                            var v116 uint16 = v12
                                            _ = v116
                                        }
                                        _ = v76
                                    } else {
                                        var v117 float64 = float64((-2.25 * 1e10))
                                        v12 = (^(v12))
                                        _ = v117
                                    }
                                    v38 -= v38
                                    var v118 uint32 = 126
                                    _ = v118
                                } else {
                                    var v119 int = fi39
                                    _ = v119
                                }
                            }
                            fmt.Println(v3)
                        }
                    default:
                        switch fi39 {
                        case 842:
                            switch (fi39 / (((fi39 * -9223372036854775808)) | 1)) {
                            case 2147483647:
                                var v120 uint64 = 18446744073709551615
                                _ = v120
                            case -1:
                                for fi121 := 0; fi121 < 9; fi121++ {
                                    delete(v6, 2147483647)
                                    _ = fi121
                                }
                            case 1:
                                for fi122 := 0; fi122 < 9; fi122++ {
                                    v123 := GenS3{F0: 4294967295, F1: (!false)}
                                    _ = fi122
                                    _ = v123
                                }
                            default:
                                switch (v21 % (((fi39 / ((v21) | 1))) | 1)) {
                                case -1:
                                    var v124 uint32 = 2147483648
                                    _ = v124
                                }
                            }
                        }
                    }
                case "hello":
                    switch (("hello" + "hello") + "\n\t") {
                    case "":
                        fmt.Println(v21)
                    case "hello":
                        fmt.Println(fi39, v3, v38)
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        for fi125 := 0; fi125 < 15; fi125++ {
                            v126 := GenS2{F0: uint32(v12), F1: 0, F2: v21}
                            _ = fi125
                            _ = v126
                        }
                    }
                default:
                    for fi127 := 0; fi127 < 4; fi127++ {
                        switch fi127 {
                        case 2147483647:
                            v128 := GenS13{F0: 0, F1: v38}
                            _ = v128
                        case -837:
                            v38 = (-(v38))
                        default:
                            switch fi127 {
                            case 127:
                                var v129 uint = 18446744073709551615
                                _ = v129
                            case 4294967295:
                                if (-9223372036854775808 < 0) {
                                    v21 = fi39
                                    v130 := GenS2{F0: 603, F1: 0, F2: 42}
                                    v12 = v12
                                    _ = v130
                                } else {
                                    v131 := GenS14{F0: v12, F1: -9223372036854775808}
                                    switch -9223372036854775808 {
                                    case -2147483648:
                                        _ = len(v6)
                                    case 2147483647:
                                        v132 := GenS10{F0: -54, F1: 1, F2: (v38 ^ v38)}
                                        _ = v132
                                    default:
                                        v3 -= ((((1e10 + (float64(1e10) + (v3 / v3))) + v3) * -2.25) / 1.0)
                                        switch -9223372036854775808 {
                                        case -393:
                                            if false {
                                                v12 += v12
                                                switch "\n\t" {
                                                case "a":
                                                    var v133 uint64 = 42
                                                    _ = v133
                                                case "":
                                                    switch 9223372036854775807 {
                                                    case -20:
                                                        if (float64(float32(v3)) > v3) {
                                                            var v134 int32 = 972
                                                            if (2 != 42) {
                                                                v3 = v3
                                                                v38 -= (v38 + 1)
                                                                if true {
                                                                    switch v21 {
                                                                    case 453:
                                                                        var v135 uint32 = 4294967295
                                                                        _ = v135
                                                                    case 42:
                                                                        v20 += (((((0.0 * (v20 + v20)) - (float32((1.5 / (1e10 - v3))) / v20)) * ((0.0 * (((v20 / 3.14159) + float32(3.14159)) + -2.25)) - float32(1.5))) * (v20 * v20)) + (float32((1.5 / v3)) + v20))
                                                                    default:
                                                                        v136 := GenS5{F0: (v12 * (^(v12)))}
                                                                        _ = v136
                                                                    }
                                                                    var v137 float64 = 3.14159
                                                                    v1 = append(v1, -2147483648)
                                                                    _ = v137
                                                                } else {
                                                                    var v138 int16 = v38
                                                                    _ = v138
                                                                }
                                                            } else {
                                                                var v139 uint8 = 255
                                                                if true {
                                                                    v134 = 2147483647
                                                                } else {
                                                                    fmt.Println(v139, fi39, v38)
                                                                    var v140 int = 32767
                                                                    if true {
                                                                        var v141 uint = 789
                                                                        _ = v141
                                                                    } else {
                                                                        var v142 int = -275
                                                                        v143 := GenS11{F0: (v134 >> ((250) & 63)), F1: 0, F2: v12}
                                                                        var v144 int8 = 100
                                                                        var v145 uint16 = v12
                                                                        if (int(v139) < -2147483648) {
                                                                            switch (fi127 + fi127) {
                                                                            case 1:
                                                                                v146 := v6[42]
                                                                                _ = v146
                                                                            case -758:
                                                                                v145 = v145
                                                                                var v147 uint8 = v139
                                                                                var v148 int8 = int8(v134)
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                            v149 := GenS9{F0: 65535, F1: 2}
                                                                            switch ("\n\t" + "a") {
                                                                            case "a":
                                                                                var v150 uint64 = 4294967295
                                                                                _ = v150
                                                                            case "foo bar":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = -9223372036854775808
                                                                                }
                                                                            case "":
                                                                                if ("hello" > "a") {
                                                                                    v142 += fi127
                                                                                }
                                                                            }
                                                                            _ = v149
                                                                        }
                                                                        _ = v142
                                                                        _ = v143
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                    _ = v140
                                                                }
                                                                _ = v139
                                                            }
                                                            var v151 uint8 = 42
                                                            _ = v134
                                                            _ = v151
                                                        } else {
                                                            v20 -= (v20 - (v20 / v20))
                                                            v38 = 32767
                                                        }
                                                    }
                                                case "foo bar":
                                                    v152 := GenS15{F0: (false && (4294967295 != 0)), F1: 4294967295}
                                                    _ = v152
                                                }
                                                v153 := GenS2{F0: 825, F1: 691, F2: fi39}
                                                if (4294967295 >= 667) {
                                                    if (true && false) {
                                                        v154 := GenS4{F0: (fi39 - (fi127 * (v21 << ((0) & 63)))), F1: 100}
                                                        v3 = (v3 * v3)
                                                        var v155 string = "foo bar"
                                                        v156 := GenS16{F0: 233}
                                                        _ = v154
                                                        _ = v155
                                                        _ = v156
                                                    } else {
                                                        var v157 int64 = 9223372036854775807
                                                        if false {
                                                            if true {
                                                                _ = len(v6)
                                                                v12 = v12
                                                                var v158 uint32 = 2147483647
                                                                v12 += v12
                                                                v157 -= v157
                                                                v159 := v5[v158]
                                                                var v160 bool = false
                                                                fmt.Println(v3, v12)
                                                                v6[(v158 << ((65535) & 63))] = fi39
                                                                v161 := GenS17{F0: v157, F1: 255}
                                                                fmt.Println(v160, v158, v21)
                                                                _ = v158
                                                                _ = v159
                                                                _ = v160
                                                                _ = v161
                                                            }
                                                            if false {
                                                                var v162 bool = (18446744073709551615 < 18446744073709551615)
                                                                v163 := GenS18{F0: 42}
                                                                var v164 int8 = -128
                                                                v3 = 0.0
                                                                _ = v162
                                                                _ = v163
                                                                _ = v164
                                                            } else {
                                                                v165 := GenS19{F0: 255, F1: 482, F2: 128}
                                                                var v166 bool = false
                                                                _ = v165
                                                                _ = v166
                                                            }
                                                            var v167 string = ""
                                                            _ = v167
                                                        }
                                                        _ = v157
                                                    }
                                                }
                                                if false {
                                                    v12 += v12
                                                } else {
                                                    delete(v5, 2)
                                                    v168 := GenS11{F0: 42, F1: (v38 - 1000), F2: (v12 * (v12 << ((2) & 63)))}
                                                    _ = v168
                                                }
                                                _ = v153
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    if (1 != 65535) {
                                                        if false {
                                                            var v169 string = "hello"
                                                            v20 = 3.14159
                                                            _ = v169
                                                        }
                                                        switch (-(fi127)) {
                                                        case -388:
                                                            var v170 uint8 = uint8(v12)
                                                            _ = v170
                                                        case 32768:
                                                            v21 += fi127
                                                        case 127:
                                                            var v171 uint32 = 2147483648
                                                            _ = v171
                                                        default:
                                                            v172 := GenS4{F0: fi39, F1: 127}
                                                            _ = v172
                                                        }
                                                    }
                                                case "hello":
                                                    v173 := GenS7{F0: (v21 ^ (fi39 * fi39)), F1: -464}
                                                    _ = v173
                                                case "":
                                                    var v174 int32 = -1
                                                    _ = v174
                                                }
                                                v175 := GenS20{F0: v12}
                                                if false {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch (^(v21)) {
                                                        case 207:
                                                            var v176 string = ("a" + ("a" + "hello"))
                                                            _ = v176
                                                        default:
                                                            fmt.Println(fi127, v38, fi127)
                                                        }
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    var v177 int64 = 9223372036854775807
                                                    v177 -= (v177 >> ((255) & 63))
                                                    _ = v177
                                                } else {
                                                    delete(v5, 509)
                                                    if false {
                                                        fmt.Println(v21)
                                                        var v178 uint32 = 0
                                                        var v179 float32 = (v20 + v20)
                                                        _ = v178
                                                        _ = v179
                                                    }
                                                }
                                                _ = v175
                                            }
                                        case 4294967295:
                                            v12 -= (v12 / ((v12) | 1))
                                        case 269:
                                            fmt.Println(fi39, v12, v3)
                                        default:
                                            var v180 int32 = 1
                                            _ = v180
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ("hello" + "foo bar")
                                    }
                                    _ = v131
                                }
                            case -9223372036854775808:
                                switch (("foo bar" + ("" + "")) + "a") {
                                case "hello":
                                    switch "foo bar" {
                                    case "hello":
                                        if ((v38 - (v38 / ((int16(v21)) | 1))) >= -459) {
                                            v3 = 1e10
                                            var v181 string = ("" + ("a" + ("a" + ("a" + "foo bar"))))
                                            _ = v181
                                        }
                                    case "\n\t":
                                        if false {
                                            v21 = (fi127 & (v21 >> ((1) & 63)))
                                            v182 := GenS9{F0: 42, F1: 4294967295}
                                            _ = v182
                                        }
                                    case "":
                                        delete(v6, 1)
                                    default:
                                        switch v21 {
                                        case 42:
                                            var v183 string = ("\n\t" + ("\n\t" + ((("hello" + "hello") + "foo bar") + "hello")))
                                            _ = v183
                                        case -9223372036854775808:
                                            var v184 int = 1
                                            _ = v184
                                        default:
                                            fmt.Println(v20, v38, fi127)
                                        }
                                    }
                                case "":
                                    var v185 int = (fi39 * -1)
                                    _ = v185
                                default:
                                    v21 -= 2147483647
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = v21
                                        }
                                    case "":
                                        v12 -= 787
                                    default:
                                        v20 -= 1e10
                                    }
                                    v186 := GenS17{F0: 0, F1: 255}
                                    _ = v186
                                }
                            }
                        }
                        _ = fi127
                    }
                }
            } else {
                var v187 int32 = 0
                v6[65535] = (fi39 * int(v12))
                _ = v187
            }
            v188 := GenS12{F0: 0}
            var v189 float32 = -2.25
            _ = v188
            _ = v189
        }
        var v190 uint64 = 42
        _ = fi39
        _ = v190
    }
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
