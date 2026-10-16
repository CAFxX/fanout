package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int; F1 uint; F2 int8 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint; F1 int }
type GenS16 struct { F0 bool; F1 int32 }
type GenS17 struct { F0 int32; F1 int16 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 bool; F1 int64; F2 uint32 }
type GenS21 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS22 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 bool; F1 int16; F2 int64 }
type GenS25 struct { F0 uint32; F1 uint8; F2 uint }
type GenS26 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS27 struct { F0 int64; F1 bool; F2 uint16 }
type GenS28 struct { F0 bool }
type GenS29 struct { F0 int16 }
type GenS30 struct { F0 uint32; F1 int32 }
type GenS31 struct { F0 int64; F1 string; F2 int }
type GenS32 struct { F0 string }
type GenS33 struct { F0 int16 }
func genFn0(p0 int, p1 int32) uint8 {
    var v6 int16 = -32768
    var v7 int32 = (p1 << ((42) & 63))
    var v8 bool = (!(255 != 255))
    var v9 int32 = -1
    v9 = v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 42
    v5 := make(map[uint32]string)
    v10 := []int{1}
    var v11 int8 = 0
    if false {
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                v2[(fi13 - fi13)] = -1
                v10 = append(v10, fi12)
                v4 = (v4 * v4)
                _ = fi13
            }
            switch ((("a" + ("\n\t" + "foo bar")) + "hello") + "") {
            case "":
                fmt.Println(v4, fi12, fi12)
            case "hello":
                if true {
                    v4 = (^(v4))
                    if false {
                        fmt.Println(v1, fi12, fi12)
                    } else {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            var v15 uint64 = 741
                            if true {
                                fmt.Println(fi14, fi14)
                            }
                            v1 = (^(v1))
                            v1 = v1
                            _ = fi14
                            _ = v15
                        }
                        var v16 float32 = (-2.25 - 1.5)
                        _ = v16
                    }
                    v17 := GenS1{F0: (v1 % (((v1 >> ((382) & 63))) | 1)), F1: 2}
                    _ = v17
                }
            case "\n\t":
                for fi18 := 0; fi18 < 5; fi18++ {
                    switch ("foo bar" + "") {
                    case "hello":
                        _ = len(v2)
                    case "foo bar":
                        v11 += v11
                    }
                    switch (("a" + (("foo bar" + "foo bar") + (((("\n\t" + "foo bar") + "\n\t") + "") + "a"))) + "hello") {
                    case "foo bar":
                        v4 += 42
                    case "hello":
                        switch (fi12 ^ (fi12 << ((18446744073709551615) & 63))) {
                        case 2147483647:
                            v19 := v2[2147483647]
                            _ = v19
                        case -371:
                            var v20 float64 = 0.1
                            _ = v20
                        case 32767:
                            if (false || (false || (801 >= 18446744073709551615))) {
                                var v21 float32 = (((1e10 / (1e10 / 1e10)) / float32((float64(float32(3.14159)) - (-2.25 * (((1e10 - float64(1e10)) * 1e10) - (((1.5 * float64(-2.25)) - (3.14159 - float64(0.0))) * (float64(3.14159) / 1.0))))))) * -2.25)
                                switch 127 {
                                case 128:
                                    switch fi12 {
                                    case 32768:
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                case -452:
                                    v21 = v21
                                    var v22 uint = 1
                                    v4 += v4
                                    _ = v22
                                case -970:
                                    v23 := GenS1{F0: (^(v1)), F1: 852}
                                    _ = v23
                                default:
                                    var v24 uint8 = 42
                                    _ = v24
                                }
                                if (!((true && (true || (true || true))) && false)) {
                                    v21 -= 1e10
                                } else {
                                    var v25 int8 = (v11 / (((v11 * (v11 << ((1) & 63)))) | 1))
                                    v26 := GenS2{F0: (v25 * -128), F1: v4, F2: 747}
                                    var v27 string = ("hello" + ("\n\t" + (("foo bar" + "\n\t") + "a")))
                                    if true {
                                        v11 += (v11 % ((-128) | 1))
                                        switch fi18 {
                                        case 128:
                                            v28 := GenS3{F0: genFn0((fi12 % ((fi18) | 1)), 42), F1: 128}
                                            _ = v28
                                        }
                                    }
                                    v25 -= 1
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v21
                            }
                        default:
                            if (((127 == 42) || true) || ((^(v11)) != 1)) {
                                fmt.Println(v4, v4, v4)
                                switch (^(fi18)) {
                                case -313:
                                    var v29 string = "hello"
                                    _ = v29
                                case 32768:
                                    v11 += (-(v11))
                                case 9223372036854775807:
                                    var v30 float32 = float32((0.0 + 3.14159))
                                    _ = v30
                                default:
                                    if len(v10) > 0 {
                                        v10[0] = fi18
                                    }
                                }
                            }
                        }
                    default:
                        if ((1e10 * 1e10) > 1e10) {
                            var v31 uint = 42
                            if true {
                                var v32 uint64 = (v4 % ((v4) | 1))
                                if (145 <= 127) {
                                    if false {
                                        var v33 float64 = 3.14159
                                        switch (-(fi12)) {
                                        case 818:
                                            var v34 uint16 = (v1 % ((1) | 1))
                                            _ = v34
                                        case 1:
                                            var v35 uint32 = 2147483647
                                            _ = v35
                                        case 127:
                                            if true {
                                                fmt.Println(fi18)
                                                var v36 float32 = -2.25
                                                _ = v36
                                            }
                                        default:
                                            switch "" {
                                            case "":
                                                v37 := GenS1{F0: 42, F1: 1}
                                                v33 = (v33 / (v33 - float64((0.0 + 1.5))))
                                                _ = v37
                                            case "hello":
                                                if len(v10) > 0 {
                                                    v10[0] = (fi12 - -375)
                                                }
                                            case "foo bar":
                                                var v38 int16 = int16(v11)
                                                v1 = (v1 << ((uint(v1)) & 63))
                                                v11 -= -17
                                                _ = v38
                                            }
                                            var v39 uint16 = v1
                                            if (-1 <= 1) {
                                                v40 := GenS4{F0: fi12, F1: 42}
                                                v39 = 2
                                                if true {
                                                    var v41 uint64 = 56
                                                    fmt.Println(v39)
                                                    v42 := GenS4{F0: fi18, F1: genFn0(fi12, 538)}
                                                    v43 := GenS5{F0: 765, F1: -1}
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v40
                                            }
                                            _ = v39
                                        }
                                        if (-9223372036854775808 < 2147483647) {
                                            v44 := GenS6{F0: 360, F1: v31, F2: 32767}
                                            _ = v44
                                        }
                                        _ = v33
                                    }
                                } else {
                                    switch ("hello" + ("a" + "hello")) {
                                    case "\n\t":
                                        fmt.Println(v1, v4)
                                    case "":
                                        v11 = (^(v11))
                                    case "a":
                                        var v45 int8 = -53
                                        _ = v45
                                    }
                                }
                                v11 += v11
                                _ = v32
                            }
                            var v46 uint16 = v1
                            v1 = 568
                            var v47 int8 = (-(v11))
                            _ = v31
                            _ = v46
                            _ = v47
                        } else {
                            v1 = v1
                        }
                    }
                    var v48 float64 = float64(0.1)
                    _ = fi18
                    _ = v48
                }
            default:
                if (255 > genFn0(fi12, 2147483647)) {
                    if (1 <= genFn0((fi12 / ((fi12) | 1)), 42)) {
                        v4 += (v4 << ((42) & 63))
                        if ((!false) && true) {
                            v49 := GenS3{F0: 127, F1: 101}
                            if false {
                                v11 -= v11
                                if (true && false) {
                                    for fi50 := 0; fi50 < 15; fi50++ {
                                        var v51 uint = uint(v4)
                                        _ = fi50
                                        _ = v51
                                    }
                                    var v52 int32 = -2
                                    var v53 int8 = (v11 ^ (-(v11)))
                                    _ = v52
                                    _ = v53
                                }
                                switch "" {
                                case "":
                                    fmt.Println(v4, v4)
                                case "\n\t":
                                    switch (("\n\t" + "hello") + "foo bar") {
                                    case "a":
                                        v54 := GenS7{F0: 1}
                                        _ = v54
                                    default:
                                        if false {
                                            v1 = (v1 - (v1 + (v1 / ((183) | 1))))
                                        } else {
                                            v11 = (-(v11))
                                            if (205 < (^(v4))) {
                                                for fi55 := 0; fi55 < 8; fi55++ {
                                                    var v56 uint8 = genFn0(fi55, 0)
                                                    _ = fi55
                                                    _ = v56
                                                }
                                                if true {
                                                    if (!(false || (0.0 == float64(float32(((1.5 * (0.0 - -2.25)) / float64(float32(1.5)))))))) {
                                                        v1 = 335
                                                        v1 = v1
                                                    }
                                                } else {
                                                    v57 := GenS8{F0: 1}
                                                    for fi58 := 0; fi58 < 17; fi58++ {
                                                        switch fi58 {
                                                        case -670:
                                                            var v59 uint64 = 4294967295
                                                            _ = v59
                                                        default:
                                                            _ = len(v10)
                                                        }
                                                        _ = fi58
                                                    }
                                                    var v60 int32 = 2147483647
                                                    _ = v57
                                                    _ = v60
                                                }
                                            } else {
                                                v61 := v3[65535]
                                                _ = v61
                                            }
                                        }
                                    }
                                }
                            } else {
                                for fi62 := 0; fi62 < 3; fi62++ {
                                    fmt.Println(fi12, v1, v11)
                                    fmt.Println(v1)
                                    fmt.Println(v1, fi62)
                                    _ = fi62
                                }
                            }
                            v1 = (v1 >> ((65535) & 63))
                            _ = v49
                        } else {
                            if (float64(0.0) == ((1e10 * 1e10) - 1.5)) {
                                v11 = (v11 << ((1) & 63))
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    for fi63 := 0; fi63 < 1; fi63++ {
                                        var v64 uint = 165
                                        _ = fi63
                                        _ = v64
                                    }
                                }
                                var v65 int32 = -1
                                _ = v65
                            } else {
                                if false {
                                    var v66 int = (fi12 ^ fi12)
                                    for fi67 := 0; fi67 < 2; fi67++ {
                                        if ((32767 == 3) && (!((!(!((false || false) || true))) || true))) {
                                            v11 -= (v11 * (^(v11)))
                                            fmt.Println(v4, v4, fi12)
                                        } else {
                                            if false {
                                                switch 1 {
                                                case -2147483648:
                                                    _ = len(v10)
                                                case 984:
                                                    v68 := GenS9{F0: 952}
                                                    _ = v68
                                                case -389:
                                                    var v69 uint32 = 982
                                                    _ = v69
                                                default:
                                                    v1 = (v1 % (((v1 * (v1 - (^(v1))))) | 1))
                                                }
                                                var v70 uint = 710
                                                v71 := GenS10{F0: true, F1: (v70 - (v70 | 18446744073709551615))}
                                                _ = v70
                                                _ = v71
                                            }
                                            if (("hello" + ("hello" + ("\n\t" + "foo bar"))) >= ("foo bar" + ("foo bar" + ("" + ("a" + "foo bar"))))) {
                                                if false {
                                                    var v72 int8 = v11
                                                    var v73 int16 = -1
                                                    v2[fi67] = -1
                                                    _ = v72
                                                    _ = v73
                                                }
                                            }
                                        }
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                v74 := GenS2{F0: (v11 % ((42) | 1)), F1: uint64(v1), F2: 42}
                                                if (float64((0.1 * (float32(1.5) + 1e10))) < ((1e10 - 1e10) * (0.0 - float64(float32(3.14159))))) {
                                                    if (false && true) {
                                                        v75 := GenS11{F0: 355, F1: 18446744073709551615, F2: (v11 + (v11 ^ (v11 ^ (v11 ^ v11))))}
                                                        _ = v75
                                                    }
                                                    v76 := GenS12{F0: v11, F1: 42}
                                                    var v77 string = ("a" + "a")
                                                    v11 = (-(v11))
                                                    fmt.Println(v4, v4)
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    if (("foo bar" + ("hello" + ("foo bar" + "foo bar"))) > (("hello" + "") + ("" + ""))) {
                                                        v66 -= int(v11)
                                                        switch fi12 {
                                                        case -1:
                                                            v11 = (v11 & (v11 * -1))
                                                        case -9223372036854775808:
                                                            var v78 float32 = 0.0
                                                            _ = v78
                                                        case 2147483647:
                                                            var v79 uint32 = 42
                                                            _ = v79
                                                        default:
                                                            v80 := GenS8{F0: -1}
                                                            _ = v80
                                                        }
                                                    }
                                                    var v81 int8 = (v11 % ((v11) | 1))
                                                    _ = v81
                                                }
                                                _ = v74
                                            case "a":
                                                if len(v10) > 0 {
                                                    v10[0] = fi12
                                                }
                                            case "foo bar":
                                                var v82 int = (fi67 - fi12)
                                                _ = v82
                                            }
                                            fmt.Println(v1)
                                            if len(v10) > 0 {
                                                v10[0] = (v66 + v66)
                                            }
                                        } else {
                                            fmt.Println(v11)
                                            if false {
                                                if (false && true) {
                                                    fmt.Println(v4, v66, v66)
                                                    if ((v4 / (((v4 ^ 1)) | 1)) > (v4 * (v4 % ((v4) | 1)))) {
                                                        _ = len(v10)
                                                    }
                                                    if (-999 != 0) {
                                                        v66 = (-(fi67))
                                                        v1 = v1
                                                    }
                                                }
                                                if true {
                                                    switch "a" {
                                                    case "a":
                                                        v11 -= (v11 ^ -2)
                                                    case "":
                                                        if (1000 <= -411) {
                                                            var v83 int8 = (v11 + (v11 & (v11 % (((-(v11))) | 1))))
                                                            if false {
                                                                if false {
                                                                    var v84 float64 = 0.0
                                                                    v4 -= v4
                                                                    v4 += (v4 >> ((447) & 63))
                                                                    _ = v84
                                                                } else {
                                                                    switch "a" {
                                                                    case "a":
                                                                        fmt.Println(v83, v66)
                                                                    case "hello":
                                                                        v85 := GenS13{F0: "hello"}
                                                                        _ = v85
                                                                    case "\n\t":
                                                                        var v86 bool = true
                                                                        _ = v86
                                                                    default:
                                                                        v87 := GenS1{F0: 65535, F1: 42}
                                                                        _ = v87
                                                                    }
                                                                    var v88 int8 = (v83 ^ 1)
                                                                    _ = v88
                                                                }
                                                                if (128 <= 42) {
                                                                    var v89 float32 = (((1e10 - float32((0.1 * 1.5))) / 1e10) * 0.0)
                                                                    _ = v89
                                                                } else {
                                                                    v1 += v1
                                                                }
                                                                if (!(v11 > v83)) {
                                                                    var v90 uint = 42
                                                                    _ = v90
                                                                }
                                                            } else {
                                                                v91 := GenS3{F0: 255, F1: 127}
                                                                if (fi67 < (fi12 >> ((42) & 63))) {
                                                                    delete(v5, 2147483648)
                                                                    _ = len(v5)
                                                                }
                                                                v92 := GenS10{F0: true, F1: 42}
                                                                v66 = 836
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                            v11 -= -128
                                                            var v93 int = -2147483648
                                                            v3[1] = 115
                                                            _ = v83
                                                            _ = v93
                                                        }
                                                    }
                                                }
                                                var v94 bool = ((v11 << ((42) & 63)) <= int8(v4))
                                                _ = v94
                                            } else {
                                                var v95 int16 = -32768
                                                _ = v95
                                            }
                                            if true {
                                                v11 = v11
                                                var v96 uint = 2
                                                _ = v96
                                            }
                                        }
                                        v97 := GenS2{F0: v11, F1: v4, F2: 0}
                                        _ = fi67
                                        _ = v97
                                    }
                                    v98 := GenS14{F0: 4294967295}
                                    _ = v66
                                    _ = v98
                                } else {
                                    for fi99 := 0; fi99 < 11; fi99++ {
                                        v100 := GenS3{F0: 255, F1: 255}
                                        v101 := GenS15{F0: 0, F1: 657}
                                        _ = fi99
                                        _ = v100
                                        _ = v101
                                    }
                                    if (!false) {
                                        var v102 int = (-(fi12))
                                        var v103 uint64 = v4
                                        _ = v102
                                        _ = v103
                                    }
                                    v104 := GenS3{F0: 255, F1: 241}
                                    _ = v104
                                }
                                v2[fi12] = 1
                                var v105 int64 = 1
                                _ = v105
                            }
                            for fi106 := 0; fi106 < 18; fi106++ {
                                if (2147483648 < 2) {
                                    v4 += uint64(v1)
                                } else {
                                    if true {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        if (3.14159 != (1.5 / (1e10 / 1.0))) {
                                            switch fi106 {
                                            case 978:
                                                v1 = v1
                                            case 32767:
                                                if (true || ((false && ((1 >= 1) || true)) && (1 != 647))) {
                                                    v1 = 65535
                                                    v4 = 702
                                                }
                                            case -9223372036854775808:
                                                if (128 < 2) {
                                                    v107 := GenS16{F0: (0.0 != float64(float32(1e10))), F1: 559}
                                                    fmt.Println(v11)
                                                    switch fi12 {
                                                    case -9223372036854775808:
                                                        v10 = append(v10, 1)
                                                    case 105:
                                                        var v108 uint16 = v1
                                                        _ = v108
                                                    }
                                                    _ = v107
                                                }
                                            }
                                            v4 = (^(v4))
                                            v4 -= (v4 ^ (v4 & v4))
                                        } else {
                                            _ = len(v3)
                                            switch (fi12 / ((fi106) | 1)) {
                                            case 4294967295:
                                                switch (fi12 + (fi106 & fi106)) {
                                                case 2147483647:
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                case -1:
                                                    delete(v5, 4294967295)
                                                case 911:
                                                    switch int(v11) {
                                                    case 4294967295:
                                                        if false {
                                                            if false {
                                                                if (42 > v1) {
                                                                    var v109 uint32 = 458
                                                                    var v110 uint = 2
                                                                    _ = v109
                                                                    _ = v110
                                                                } else {
                                                                    var v111 float32 = 1.5
                                                                    switch (fi12 + (fi12 % ((9223372036854775807) | 1))) {
                                                                    case -9223372036854775808:
                                                                        v112 := GenS16{F0: false, F1: -401}
                                                                        _ = v112
                                                                    case 634:
                                                                        _ = len(v10)
                                                                    case 32767:
                                                                        v113 := GenS17{F0: -591, F1: 0}
                                                                        _ = v113
                                                                    default:
                                                                        var v114 float64 = 1.5
                                                                        _ = v114
                                                                    }
                                                                    switch (fi106 | (fi12 * fi12)) {
                                                                    case -461:
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v4 -= 4294967295
                                                                        case "\n\t":
                                                                            fmt.Println(v111, v1, fi12)
                                                                        case "a":
                                                                            switch fi106 {
                                                                            case 9223372036854775807:
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                            case -2147483648:
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                            }
                                                                        default:
                                                                            v1 = 42
                                                                            _ = len(v5)
                                                                            fmt.Println(v1, fi12, v11)
                                                                            if false {
                                                                                switch fi12 {
                                                                                case 1:
                                                                                    v11 -= (v11 << ((1) & 63))
                                                                                case 128:
                                                                                    v115 := GenS18{F0: v1}
                                                                                    _ = v115
                                                                                }
                                                                                v116 := GenS19{F0: v11}
                                                                                _ = v116
                                                                            } else {
                                                                                switch 1 {
                                                                                case -2147483648:
                                                                                    _ = len(v2)
                                                                                }
                                                                            }
                                                                            fmt.Println(fi12, v1, v111)
                                                                        }
                                                                    case 0:
                                                                        v117 := GenS20{F0: false, F1: 363, F2: 860}
                                                                        fmt.Println(fi12, v11, fi106)
                                                                        var v118 int64 = 429
                                                                        var v119 int = fi106
                                                                        _ = v117
                                                                        _ = v118
                                                                        _ = v119
                                                                    default:
                                                                        var v120 uint32 = uint32(v1)
                                                                        _ = v120
                                                                    }
                                                                    _ = v111
                                                                }
                                                                var v121 int64 = 42
                                                                _ = v121
                                                            } else {
                                                                var v122 uint16 = (v1 + (v1 * (^(v1))))
                                                                v123 := GenS12{F0: 127, F1: 255}
                                                                if (4294967295 <= 255) {
                                                                    if (false || (("a" <= "\n\t") && (1.5 == (((0.1 * 3.14159) + 3.14159) + -2.25)))) {
                                                                        v4 -= v4
                                                                    }
                                                                    v11 = int8(fi106)
                                                                    v124 := GenS21{F0: (v4 * (v4 / ((uint64(v1)) | 1))), F1: 4294967295, F2: 32767}
                                                                    _ = v124
                                                                }
                                                                var v125 uint64 = (v4 >> ((4294967295) & 63))
                                                                _ = v122
                                                                _ = v123
                                                                _ = v125
                                                            }
                                                            v126 := GenS15{F0: 255, F1: fi106}
                                                            _ = v126
                                                        } else {
                                                            switch "hello" {
                                                            case "a":
                                                                fmt.Println(fi12, fi106)
                                                            }
                                                            delete(v3, 746)
                                                        }
                                                    case -729:
                                                        fmt.Println(fi12)
                                                    }
                                                default:
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "":
                                                            fmt.Println(v4)
                                                        case "a":
                                                            v4 -= v4
                                                        }
                                                    case "a":
                                                        if (0 < -2147483648) {
                                                            var v127 float64 = 1.5
                                                            if false {
                                                                fmt.Println(fi106, v4)
                                                            } else {
                                                                v4 -= v4
                                                                v127 -= ((1e10 * (0.0 - float64(3.14159))) * (((v127 * v127) + ((float64(0.1) - 0.0) * (0.1 - (v127 / v127)))) - (v127 - ((-2.25 + 0.1) / ((v127 + float64((float32(v127) * float32(v127)))) - v127)))))
                                                            }
                                                            switch fi12 {
                                                            case 127:
                                                                fmt.Println(v1, v4, v4)
                                                            case 0:
                                                                v128 := GenS22{F0: -1, F1: v4, F2: v1}
                                                                v127 += v127
                                                                switch fi106 {
                                                                case -687:
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v129 uint16 = (^(v1))
                                                                        _ = v129
                                                                    default:
                                                                        v2[fi12] = 32768
                                                                    }
                                                                case 2147483647:
                                                                    var v130 uint8 = 1
                                                                    _ = v130
                                                                case 897:
                                                                    delete(v3, 229)
                                                                }
                                                                v4 -= v4
                                                                _ = v128
                                                            }
                                                            _ = v127
                                                        } else {
                                                            var v131 uint8 = 255
                                                            if (true || false) {
                                                                var v132 uint8 = 42
                                                                _ = v132
                                                            }
                                                            _ = v131
                                                        }
                                                    }
                                                }
                                            case 127:
                                                var v133 int32 = 0
                                                _ = v133
                                            case -1:
                                                if ((v4 % ((v4) | 1)) != 4294967295) {
                                                    v134 := v5[705]
                                                    v4 = v4
                                                    _ = v134
                                                }
                                            default:
                                                switch ("hello" + "hello") {
                                                case "hello":
                                                    fmt.Println(fi12, v1, v4)
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v135 int32 = -2147483648
                                                    _ = v135
                                                case "foo bar":
                                                    var v136 uint16 = v1
                                                    _ = v136
                                                case "a":
                                                    v137 := GenS8{F0: 1}
                                                    _ = v137
                                                default:
                                                    fmt.Println(v4)
                                                }
                                            }
                                        }
                                        v138 := GenS23{F0: 2147483647}
                                        if (false || (false && ((1.5 > 1.5) && true))) {
                                            v139 := GenS15{F0: 42, F1: (fi12 >> ((uint(v1)) & 63))}
                                            v1 += (v1 & v1)
                                            if false {
                                                if true {
                                                    switch 32768 {
                                                    case -474:
                                                        v140 := GenS24{F0: true, F1: 32767, F2: 9223372036854775807}
                                                        _ = v140
                                                    case 737:
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                    }
                                                    v141 := GenS13{F0: "a"}
                                                    var v142 uint = 4294967295
                                                    v142 += v142
                                                    v4 -= (v4 / ((v4) | 1))
                                                    if (-32768 <= -1) {
                                                        if len(v10) > 0 {
                                                            v10[0] = fi106
                                                        }
                                                    } else {
                                                        if false {
                                                            v5[2147483648] = ""
                                                            v1 = 32767
                                                            if true {
                                                                _ = len(v3)
                                                                var v143 uint = uint(v1)
                                                                v144 := GenS15{F0: 18446744073709551615, F1: (fi106 / (((^(fi106))) | 1))}
                                                                _ = v143
                                                                _ = v144
                                                            } else {
                                                                switch (-(fi106)) {
                                                                case 1:
                                                                    var v145 uint32 = uint32(v1)
                                                                    _ = v145
                                                                case 42:
                                                                    fmt.Println(v11)
                                                                default:
                                                                    v142 = (v142 << ((4294967295) & 63))
                                                                    v1 = v1
                                                                    var v146 uint64 = 1
                                                                    if (v1 > 2) {
                                                                        _ = len(v10)
                                                                        _ = len(v10)
                                                                        v147 := GenS4{F0: (fi12 << (((v142 | 870)) & 63)), F1: 2}
                                                                        _ = v147
                                                                    }
                                                                    v1 -= 65535
                                                                    _ = v146
                                                                }
                                                                v148 := GenS25{F0: 2, F1: 0, F2: 2}
                                                                if ((^(fi12)) > fi106) {
                                                                    v149 := GenS26{F0: 18446744073709551615, F1: uint16(v4), F2: 42}
                                                                    delete(v5, 65535)
                                                                    v1 -= v1
                                                                    _ = v149
                                                                } else {
                                                                    v11 -= -112
                                                                }
                                                                v1 = v1
                                                                _ = v148
                                                            }
                                                        } else {
                                                            v150 := GenS18{F0: uint16(v4)}
                                                            var v151 uint16 = (^(v1))
                                                            if false {
                                                                v11 -= (v11 | 127)
                                                                v152 := GenS27{F0: 1, F1: false, F2: 65535}
                                                                _ = v152
                                                            }
                                                            v151 -= v1
                                                            _ = v150
                                                            _ = v151
                                                        }
                                                        v153 := GenS28{F0: false}
                                                        v154 := GenS18{F0: 927}
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    _ = v141
                                                    _ = v142
                                                } else {
                                                    if true {
                                                        fmt.Println(fi12, v4, fi12)
                                                        v155 := GenS29{F0: 1000}
                                                        v1 += v1
                                                        _ = v155
                                                    }
                                                    fmt.Println(fi106, v1, fi12)
                                                    var v156 bool = false
                                                    _ = v156
                                                }
                                                v11 = v11
                                                v1 -= (v1 + 2)
                                            }
                                            _ = v139
                                        } else {
                                            var v157 bool = (("a" + "foo bar") < ("" + ""))
                                            _ = v157
                                        }
                                        _ = v138
                                    }
                                    fmt.Println(fi12, v4, v4)
                                }
                                _ = fi106
                            }
                            var v158 float64 = (1.5 * 1e10)
                            _ = v158
                        }
                        if (-894 != (^(fi12))) {
                            var v159 uint16 = v1
                            v160 := GenS30{F0: 2, F1: -1}
                            if true {
                                if (1.5 == 0.0) {
                                    if false {
                                        for fi161 := 0; fi161 < 3; fi161++ {
                                            if true {
                                                if (int8(fi12) < v11) {
                                                    v4 -= (^(v4))
                                                    switch "hello" {
                                                    case "a":
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            var v162 bool = false
                                                            _ = v162
                                                        case "hello":
                                                            var v163 uint64 = v4
                                                            _ = v163
                                                        case "a":
                                                            v4 = (v4 / ((v4) | 1))
                                                        default:
                                                            v159 = (v159 / (((v1 + v1)) | 1))
                                                        }
                                                    default:
                                                        var v164 uint32 = 0
                                                        _ = v164
                                                    }
                                                    v165 := GenS24{F0: (false || true), F1: 1, F2: -1}
                                                    _ = v165
                                                }
                                            } else {
                                                v166 := GenS2{F0: 0, F1: v4, F2: 0}
                                                _ = v166
                                            }
                                            _ = fi161
                                        }
                                        v11 += -1
                                    } else {
                                        v167 := GenS29{F0: 1}
                                        v11 -= v11
                                        var v168 uint8 = 127
                                        _ = v167
                                        _ = v168
                                    }
                                    v11 -= v11
                                } else {
                                    var v169 uint8 = 127
                                    if true {
                                        _ = len(v10)
                                        for fi170 := 0; fi170 < 6; fi170++ {
                                            var v171 bool = true
                                            var v172 int = -9223372036854775808
                                            _ = fi170
                                            _ = v171
                                            _ = v172
                                        }
                                        v11 -= (-(v11))
                                    }
                                    for fi173 := 0; fi173 < 18; fi173++ {
                                        var v174 int32 = -1
                                        _ = fi173
                                        _ = v174
                                    }
                                    _ = v169
                                }
                                switch fi12 {
                                case 368:
                                    switch 42 {
                                    case 128:
                                        switch (fi12 % ((fi12) | 1)) {
                                        case 127:
                                            for fi175 := 0; fi175 < 6; fi175++ {
                                                v1 -= 205
                                                v4 = (v4 + v4)
                                                fmt.Println(fi12, fi12, fi12)
                                                _ = fi175
                                            }
                                        case 42:
                                            v176 := GenS31{F0: -1, F1: "\n\t", F2: fi12}
                                            _ = v176
                                        case 218:
                                            var v177 float64 = -2.25
                                            _ = v177
                                        }
                                    case -385:
                                        fmt.Println(v11, v1)
                                    case -2147483648:
                                        v4 = (v4 - 18446744073709551615)
                                    default:
                                        v1 = v159
                                    }
                                default:
                                    for fi178 := 0; fi178 < 2; fi178++ {
                                        var v179 bool = (!(genFn0((-(fi178)), 0) >= 2))
                                        v159 = v159
                                        var v180 float32 = (1.5 - 3.14159)
                                        _ = fi178
                                        _ = v179
                                        _ = v180
                                    }
                                }
                                for fi181 := 0; fi181 < 2; fi181++ {
                                    switch fi181 {
                                    case -84:
                                        if (42 <= 42) {
                                            var v182 uint = 2
                                            var v183 float64 = 0.0
                                            _ = v182
                                            _ = v183
                                        } else {
                                            if (!((0.1 == float64(3.14159)) && false)) {
                                                var v184 int64 = 0
                                                _ = v184
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v4 += 18446744073709551615
                                            default:
                                                fmt.Println(v11)
                                            }
                                        }
                                    case 42:
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v2[fi181] = (^(fi181))
                                            case "":
                                                v4 -= (v4 % ((v4) | 1))
                                            case "a":
                                                var v185 int32 = 0
                                                _ = v185
                                            }
                                        } else {
                                            var v186 string = ("\n\t" + "a")
                                            v11 = (v11 | (-(v11)))
                                            _ = v186
                                        }
                                    default:
                                        var v187 uint32 = 2
                                        _ = v187
                                    }
                                    v188 := GenS31{F0: 42, F1: "foo bar", F2: fi181}
                                    _ = fi181
                                    _ = v188
                                }
                            }
                            _ = v159
                            _ = v160
                        }
                    } else {
                        v189 := GenS32{F0: "foo bar"}
                        _ = v189
                    }
                    v190 := GenS3{F0: 1, F1: genFn0(fi12, -2147483648)}
                    _ = v190
                }
            }
            _ = fi12
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v1)
        }
    }
    fmt.Println(v11)
    v191 := GenS33{F0: 42}
    v1 = (v1 / (((v1 - (v1 + 1))) | 1))
    var v192 uint = 113
    if (!(2147483648 <= 1)) {
        var v193 float32 = 0.1
        _ = v193
    } else {
        v1 -= v1
        fmt.Println(v4, v4)
        for fi194 := 0; fi194 < 11; fi194++ {
            if (0 >= -9223372036854775808) {
                fmt.Println(v4, v11)
            }
            if true {
                for fi195 := 0; fi195 < 14; fi195++ {
                    var v196 float32 = 1e10
                    v196 += -2.25
                    v196 = v196
                    _ = fi195
                    _ = v196
                }
                v192 = (^(v192))
                v11 += (v11 >> ((v192) & 63))
            } else {
                var v197 bool = false
                var v198 string = ""
                _ = v197
                _ = v198
            }
            _ = fi194
        }
    }
    _ = len(v5)
    v199 := GenS19{F0: v11}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v192:", v192)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v191
    _ = v192
    _ = v199
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
