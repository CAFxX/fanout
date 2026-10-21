package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int64; F1 int64; F2 int64 }
type GenS7 struct { F0 int; F1 uint64; F2 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint; F1 uint64; F2 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint32; F1 uint16 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 string) uint64 {
    var v6 int64 = 0
    var v7 uint64 = 2
    _ = p0
    _ = v6
    _ = v7
    return (v7 << ((2) & 63))
}

func main() {
    var v1 uint = 1
    v2 := []string{"hello", ("hello" + ("\n\t" + "")), "hello"}
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 uint16 = 42
    if (false || (!((1e10 / float32(1e10)) > float32(3.14159)))) {
        _ = len(v2)
        switch 4294967295 {
        case -570:
            v5 += 2
        case -9223372036854775808:
            if (!false) {
                switch -233 {
                case -723:
                    if false {
                        v1 = v1
                        v5 = (v5 ^ v5)
                    }
                case 128:
                    var v8 float32 = float32(3.14159)
                    _ = v8
                case 2147483647:
                    var v9 uint = (v1 | (^(v1)))
                    _ = v9
                default:
                    v1 += 4294967295
                }
                var v10 float64 = -2.25
                _ = v10
            } else {
                v5 = v5
                v11 := GenS1{F0: (v1 | 2), F1: "hello", F2: 127}
                var v12 string = "\n\t"
                _ = v11
                _ = v12
            }
        case -2147483648:
            if (float64(1.5) < 3.14159) {
                v13 := GenS2{F0: 2}
                v14 := GenS2{F0: 1}
                v5 = v5
                _ = v13
                _ = v14
            } else {
                var v15 float64 = (1.5 * (3.14159 / 1.0))
                v3[628] = 127
                _ = v15
            }
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5)
        default:
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int64 = 42
                if false {
                    var v18 uint16 = 255
                    if false {
                        if (!((v5 - v5) == v5)) {
                            var v19 int16 = 32767
                            v19 -= v19
                            _ = v19
                        } else {
                            var v20 int = fi16
                            _ = v20
                        }
                        v21 := GenS3{F0: false}
                        var v22 int = (fi16 + fi16)
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v24 := GenS4{F0: (^(v1))}
                        fmt.Println(v1)
                        v25 := GenS5{F0: 42, F1: "foo bar"}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v18
                }
                v1 = v1
                _ = fi16
                _ = v17
            }
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        switch ("a" + "a") {
        case "\n\t":
            v27 := GenS1{F0: (v1 | 855), F1: ((("\n\t" + "hello") + "a") + ""), F2: -75}
            _ = v27
        case "a":
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS6{F0: -252, F1: 129, F2: 42}
                switch 0 {
                case -1:
                    if true {
                        v2 = append(v2, (("\n\t" + "a") + "a"))
                    } else {
                        if (2147483647 > 42) {
                            var v30 uint32 = 65535
                            var v31 uint16 = v5
                            v31 += 64
                            _ = v30
                            _ = v31
                        }
                    }
                }
                v4[(^(fi26))] = 559
                _ = fi28
                _ = v29
            }
        default:
            v32 := GenS4{F0: 255}
            _ = v32
        }
        var v33 uint8 = 42
        var v34 string = "a"
        _ = fi26
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: 960}
    var v36 uint32 = 2147483647
    if true {
        for fi37 := 0; fi37 < 4; fi37++ {
            if (!false) {
                var v38 string = "foo bar"
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 string = v38
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch (v40 + "a") {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "":
                        if true {
                            fmt.Println(v1, v1)
                            var v41 int64 = 9223372036854775807
                            v38 = ((("" + "foo bar") + "") + v40)
                            var v42 uint8 = 255
                            var v43 int16 = 136
                            _ = v41
                            _ = v42
                            _ = v43
                        } else {
                            if ((v5 + 255) != (v5 - v5)) {
                                fmt.Println(fi37, v1)
                            }
                        }
                    default:
                        fmt.Println(fi39, v40, v40)
                    }
                    _ = fi39
                    _ = v40
                }
                switch "" {
                case "\n\t":
                    v5 += (v5 ^ v5)
                case "a":
                    if (28 < 516) {
                        if false {
                            switch (-(fi37)) {
                            case -832:
                                var v44 bool = false
                                _ = v44
                            case 128:
                                if false {
                                    v2 = append(v2, "hello")
                                    for fi45 := 0; fi45 < 10; fi45++ {
                                        _ = len(v3)
                                        var v46 int32 = 42
                                        _ = fi45
                                        _ = v46
                                    }
                                } else {
                                    if ((v1 / (((v1 >> ((v1) & 63))) | 1)) != (^(v1))) {
                                        for fi47 := 0; fi47 < 1; fi47++ {
                                            var v48 int64 = 44
                                            switch (fi37 << ((0) & 63)) {
                                            case -1:
                                                switch (fi47 << ((255) & 63)) {
                                                case 0:
                                                    v1 = (v1 + 880)
                                                }
                                            default:
                                                if (0.1 == (0.1 * (float64(((1.5 / 3.14159) * (float32(0.0) + 1e10))) + -2.25))) {
                                                    var v49 int = -1
                                                    _ = v49
                                                }
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                        if true {
                                            if (v36 <= v36) {
                                                var v50 uint16 = 255
                                                v5 = 110
                                                var v51 float32 = ((0.1 + 0.1) / 0.1)
                                                _ = v50
                                                _ = v51
                                            } else {
                                                if ((v5 + (v5 / (((v5 / (((v5 % ((32767) | 1))) | 1))) | 1))) == uint16(v36)) {
                                                    var v52 int8 = 127
                                                    switch 9223372036854775807 {
                                                    case -9223372036854775808:
                                                        for fi53 := 0; fi53 < 4; fi53++ {
                                                            fmt.Println(v36, fi53, v5)
                                                            if false {
                                                                if (false && ((int(v1) < 1) && (75 > 128))) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v38
                                                                    }
                                                                }
                                                            }
                                                            _ = fi53
                                                        }
                                                    case 9223372036854775807:
                                                        fmt.Println(v52, v1, v38)
                                                    }
                                                    _ = v52
                                                } else {
                                                    v54 := GenS7{F0: fi37, F1: genFn0("a"), F2: 42}
                                                    _ = v54
                                                }
                                                v36 -= v36
                                                v38 = (v38 + v38)
                                            }
                                            if (!(0 == -1)) {
                                                if (true && (((true && (!(v36 < v36))) && false) || true)) {
                                                    v38 = (v38 + v38)
                                                    v1 += 42
                                                    _ = len(v2)
                                                }
                                                fmt.Println(v1, v1)
                                            } else {
                                                v36 = (v36 & 2)
                                                if (989 <= 345) {
                                                    switch (fi37 / (((fi37 % ((4294967295) | 1))) | 1)) {
                                                    case 1:
                                                        var v55 uint8 = 42
                                                        _ = v55
                                                    case 183:
                                                        for fi56 := 0; fi56 < 7; fi56++ {
                                                            var v57 float32 = 0.0
                                                            v5 += (^(v5))
                                                            var v58 int16 = 42
                                                            _ = fi56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case 145:
                                                        switch ((v38 + "\n\t") + v38) {
                                                        case "\n\t":
                                                            if (!false) {
                                                                if true {
                                                                    v38 = ((v38 + "\n\t") + "foo bar")
                                                                    _ = len(v2)
                                                                }
                                                                v59 := GenS3{F0: false}
                                                                _ = v59
                                                            } else {
                                                                if (1 != v5) {
                                                                    for fi60 := 0; fi60 < 20; fi60++ {
                                                                        switch (fi60 << ((42) & 63)) {
                                                                        case 42:
                                                                            fmt.Println(v38, v38, fi60)
                                                                        case 541:
                                                                            fmt.Println(v1, fi60)
                                                                            v38 = "\n\t"
                                                                            var v61 int8 = 42
                                                                            _ = v61
                                                                        case 898:
                                                                            _ = len(v2)
                                                                        }
                                                                        _ = fi60
                                                                    }
                                                                    for fi62 := 0; fi62 < 1; fi62++ {
                                                                        var v63 uint16 = (v5 & v5)
                                                                        _ = fi62
                                                                        _ = v63
                                                                    }
                                                                } else {
                                                                    var v64 uint8 = 2
                                                                    for fi65 := 0; fi65 < 7; fi65++ {
                                                                        v5 = 57
                                                                        _ = len(v2)
                                                                        _ = fi65
                                                                    }
                                                                    _ = v64
                                                                }
                                                                v66 := GenS6{F0: 0, F1: 9223372036854775807, F2: 1}
                                                                _ = v66
                                                            }
                                                        case "a":
                                                            switch -1 {
                                                            case 914:
                                                                v1 += v1
                                                            case -2147483648:
                                                                var v67 int16 = 1
                                                                _ = v67
                                                            }
                                                        case "foo bar":
                                                            v38 = v38
                                                        }
                                                    }
                                                    var v68 uint16 = 348
                                                    if len(v2) > 0 {
                                                        v2[0] = ((("foo bar" + "") + v38) + ((("foo bar" + ("hello" + v38)) + "hello") + "foo bar"))
                                                    }
                                                    _ = v68
                                                } else {
                                                    for fi69 := 0; fi69 < 15; fi69++ {
                                                        if len(v2) > 0 {
                                                            v2[0] = v38
                                                        }
                                                        fmt.Println(v5, v38)
                                                        _ = fi69
                                                    }
                                                }
                                            }
                                            fmt.Println(v1)
                                        } else {
                                            var v70 int8 = 0
                                            _ = v70
                                        }
                                    }
                                    if true {
                                        if (v1 < v1) {
                                            var v71 int8 = 100
                                            v72 := GenS8{F0: 1}
                                            _ = v71
                                            _ = v72
                                        } else {
                                            var v73 int16 = 32767
                                            for fi74 := 0; fi74 < 19; fi74++ {
                                                var v75 uint64 = 42
                                                _ = fi74
                                                _ = v75
                                            }
                                            var v76 uint64 = 0
                                            _ = v73
                                            _ = v76
                                        }
                                        if (!false) {
                                            var v77 int8 = 42
                                            _ = v77
                                        } else {
                                            switch fi37 {
                                            case 1:
                                                fmt.Println(v1, v5, v5)
                                            case 128:
                                                var v78 int64 = 9223372036854775807
                                                _ = v78
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v79 := v3[fi37]
                                        _ = v79
                                    }
                                    if (!true) {
                                        switch ((v38 + v38) + v38) {
                                        case "":
                                            fmt.Println(fi37, v1, v36)
                                        case "a":
                                            if (true || (1 >= uint8(v1))) {
                                                if true {
                                                    var v80 uint8 = 171
                                                    v38 = (v38 + v38)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v80
                                                }
                                            }
                                        case "\n\t":
                                            v2 = append(v2, "a")
                                        }
                                        fmt.Println(v36, v1, v1)
                                        v1 = 0
                                    } else {
                                        switch v38 {
                                        case "":
                                            for fi81 := 0; fi81 < 18; fi81++ {
                                                _ = len(v2)
                                                v38 = v38
                                                _ = fi81
                                            }
                                        case "\n\t":
                                            if (v38 == v38) {
                                                switch (fi37 - (-(fi37))) {
                                                case -270:
                                                    for fi82 := 0; fi82 < 8; fi82++ {
                                                        fmt.Println(v36, v36)
                                                        var v83 int16 = 32767
                                                        var v84 int32 = -1
                                                        _ = fi82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                case 542:
                                                    if true {
                                                        switch fi37 {
                                                        case -1:
                                                            v85 := GenS8{F0: 9223372036854775807}
                                                            _ = v85
                                                        case 127:
                                                            var v86 string = "foo bar"
                                                            _ = v86
                                                        }
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = v38
                                                }
                                            }
                                        default:
                                            switch fi37 {
                                            case -2147483648:
                                                var v87 uint16 = v5
                                                _ = v87
                                            default:
                                                var v88 int8 = -58
                                                _ = v88
                                            }
                                        }
                                        v38 = ((v38 + v38) + (("" + v38) + ("" + "")))
                                        var v89 uint64 = 42
                                        _ = v89
                                    }
                                }
                            default:
                                for fi90 := 0; fi90 < 3; fi90++ {
                                    switch fi37 {
                                    case 127:
                                        delete(v4, (fi37 - fi37))
                                        if false {
                                            var v91 bool = false
                                            switch (^(fi37)) {
                                            case -1:
                                                v1 = (v1 + (^(v1)))
                                            default:
                                                delete(v4, fi37)
                                            }
                                            _ = v91
                                        }
                                        v38 = (((v38 + v38) + v38) + v38)
                                        if true {
                                            var v92 uint64 = 1
                                            _ = v92
                                        } else {
                                            if ((!(0.1 < -2.25)) || true) {
                                                if (281 >= 584) {
                                                    delete(v3, (fi37 % (((^(fi37))) | 1)))
                                                } else {
                                                    switch "a" {
                                                    case "a":
                                                        v38 = v38
                                                    }
                                                    switch v38 {
                                                    case "a":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        var v93 float32 = ((float32(0.0) * 1e10) / 0.1)
                                                        _ = v93
                                                    }
                                                    var v94 uint = v1
                                                    var v95 uint32 = uint32(v1)
                                                    v94 += 0
                                                    _ = v94
                                                    _ = v95
                                                }
                                                fmt.Println(fi90, v5)
                                                v36 = (^(v36))
                                            }
                                            if ((float64(float32(3.14159)) - 3.14159) != 1e10) {
                                                v36 += 1
                                                var v96 uint16 = 2
                                                _ = len(v2)
                                                _ = v96
                                            } else {
                                                var v97 int = fi90
                                                if (!((68 <= 28) && (18446744073709551615 <= v1))) {
                                                    fmt.Println(fi37, fi37)
                                                    switch ("foo bar" + v38) {
                                                    case "\n\t":
                                                        switch v38 {
                                                        case "foo bar":
                                                            v98 := v3[127]
                                                            _ = v98
                                                        case "":
                                                            var v99 float64 = ((1.5 / ((1.5 * 1e10) / 0.1)) + (3.14159 * 0.0))
                                                            _ = v99
                                                        default:
                                                            var v100 uint = (v1 << (((^(v1))) & 63))
                                                            _ = v100
                                                        }
                                                    case "":
                                                        var v101 int8 = 0
                                                        _ = v101
                                                    default:
                                                        var v102 bool = (255 > 127)
                                                        _ = v102
                                                    }
                                                    _ = len(v2)
                                                } else {
                                                    v38 = ("\n\t" + v38)
                                                }
                                                fmt.Println(fi90, fi90)
                                                var v103 uint16 = (v5 % ((564) | 1))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v103
                                                _ = v97
                                            }
                                            switch v38 {
                                            case "\n\t":
                                                switch "" {
                                                case "":
                                                    if true {
                                                        v1 = v1
                                                    } else {
                                                        v3[fi90] = 127
                                                        v36 += (v36 | (^(v36)))
                                                        var v104 int32 = 2147483647
                                                        _ = v104
                                                    }
                                                }
                                            case "foo bar":
                                                switch 2147483647 {
                                                case 1:
                                                    var v105 int = 17
                                                    _ = v105
                                                case 2147483647:
                                                    if (!(true || true)) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v36 += (v36 & 814)
                                                        switch ((v38 + ("" + (((v38 + (v38 + ((v38 + (("foo bar" + (v38 + v38)) + (("a" + v38) + v38))) + ((v38 + v38) + "a")))) + ("" + v38)) + v38))) + ((v38 + ("a" + ((v38 + v38) + v38))) + ((v38 + v38) + v38))) {
                                                        case "\n\t":
                                                            v5 = v5
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v106 int64 = 750
                                                                _ = v106
                                                            case "\n\t":
                                                                var v107 uint32 = v36
                                                                _ = v107
                                                            case "a":
                                                                fmt.Println(fi90)
                                                            default:
                                                                v108 := GenS1{F0: (v1 >> ((v1) & 63)), F1: v38, F2: -102}
                                                                _ = v108
                                                            }
                                                        }
                                                    } else {
                                                        var v109 int16 = -807
                                                        v38 = v38
                                                        v110 := GenS2{F0: 1}
                                                        _ = len(v2)
                                                        var v111 float32 = 1e10
                                                        v109 = int16(v1)
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                }
                                            case "":
                                                switch (fi37 * (-(fi37))) {
                                                case 42:
                                                    v112 := GenS5{F0: 42, F1: ("a" + (v38 + v38))}
                                                    _ = v112
                                                default:
                                                    var v113 int64 = 9223372036854775807
                                                    _ = v113
                                                }
                                            }
                                        }
                                    case -153:
                                        v36 = v36
                                    case 42:
                                        var v114 uint64 = 1
                                        _ = v114
                                    }
                                    _ = fi90
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                v115 := v3[1]
                                _ = v115
                            case "foo bar":
                                for fi116 := 0; fi116 < 14; fi116++ {
                                    switch v38 {
                                    case "a":
                                        v117 := GenS9{F0: 1}
                                        _ = v117
                                    default:
                                        switch fi116 {
                                        case 128:
                                            fmt.Println(v5, v5)
                                        }
                                    }
                                    v118 := GenS1{F0: (v1 & 42), F1: (v38 + ((v38 + "") + v38)), F2: 127}
                                    _ = fi116
                                    _ = v118
                                }
                            default:
                                v38 = v38
                            }
                            fmt.Println(v1, v5, v5)
                        } else {
                            var v119 float64 = 0.0
                            fmt.Println(v1)
                            _ = v119
                        }
                    }
                case "foo bar":
                    v1 = (v1 - (v1 << (((v1 / ((v1) | 1))) & 63)))
                default:
                    fmt.Println(fi37, fi37, fi37)
                }
                _ = v38
            }
            var v120 int32 = -2147483648
            _ = fi37
            _ = v120
        }
        for fi121 := 0; fi121 < 6; fi121++ {
            var v122 float32 = 0.0
            _ = fi121
            _ = v122
        }
    } else {
        fmt.Println(v1, v36, v36)
        fmt.Println(v36, v36)
        fmt.Println(v1, v36, v1)
    }
    var v123 float64 = (0.0 + 1.5)
    v1 -= v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((v5 & (v5 * v5)) > (v5 << ((uint(v5)) & 63))) {
        v36 = 1
    } else {
        var v124 uint16 = (v5 ^ v5)
        v4[414] = (v36 >> (((v1 * 18446744073709551615)) & 63))
        _ = v124
    }
    if ((0.1 >= 1.5) || (!true)) {
        if true {
            v125 := GenS10{F0: v5, F1: 100, F2: uint8(v1)}
            var v126 uint32 = 1
            for fi127 := 0; fi127 < 9; fi127++ {
                fmt.Println(v123, v36)
                var v128 uint32 = (v126 + (v36 * v126))
                _ = fi127
                _ = v128
            }
            _ = v125
            _ = v126
        } else {
            for fi129 := 0; fi129 < 20; fi129++ {
                if false {
                    var v130 int16 = 0
                    var v131 int = fi129
                    var v132 uint = (^(v1))
                    _ = v130
                    _ = v131
                    _ = v132
                } else {
                    for fi133 := 0; fi133 < 17; fi133++ {
                        switch (fi129 - fi129) {
                        case -354:
                            v1 = v1
                        case 1:
                            v1 += (v1 ^ v1)
                        case 127:
                            _ = len(v2)
                        default:
                            v134 := GenS3{F0: false}
                            _ = v134
                        }
                        var v135 int16 = -1
                        _ = fi133
                        _ = v135
                    }
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
                if false {
                    for fi136 := 0; fi136 < 12; fi136++ {
                        fmt.Println(v36)
                        var v137 int32 = 456
                        _ = fi136
                        _ = v137
                    }
                }
                _ = fi129
            }
        }
        switch -2147483648 {
        case 128:
            if true {
                for fi138 := 0; fi138 < 8; fi138++ {
                    fmt.Println(v36)
                    fmt.Println(v36)
                    _ = fi138
                }
            } else {
                delete(v3, 32768)
            }
        case 1:
            v139 := GenS8{F0: 9223372036854775807}
            _ = v139
        default:
            _ = len(v4)
        }
    }
    v5 = (v5 * 0)
    if false {
        switch -2147483648 {
        case 128:
            if (3.14159 >= (0.1 - float32((3.14159 / v123)))) {
                v140 := GenS11{F0: v1, F1: 18446744073709551615, F2: "foo bar"}
                v1 += uint(v5)
                var v141 uint16 = v5
                _ = v140
                _ = v141
            } else {
                v142 := GenS2{F0: 1}
                _ = v142
            }
        case -9223372036854775808:
            if (((^(v5)) == v5) && false) {
                if (false && (false && (42 > -91))) {
                    fmt.Println(v36, v1)
                    var v143 int16 = 32767
                    _ = len(v3)
                    _ = v143
                }
                if true {
                    if (1 == 2) {
                        for fi144 := 0; fi144 < 16; fi144++ {
                            v123 -= ((0.1 * (float64((float32(v123) + float32(((-2.25 + 0.0) * v123)))) + (float64(-2.25) - v123))) + v123)
                            fmt.Println(v1, v123)
                            var v145 uint32 = v36
                            _ = fi144
                            _ = v145
                        }
                        var v146 int64 = 0
                        _ = v146
                    } else {
                        v123 = (v123 * 1e10)
                        fmt.Println(v1)
                    }
                }
                v147 := GenS2{F0: 2}
                _ = v147
            } else {
                switch 128 {
                case 32768:
                    v148 := GenS12{F0: 1}
                    _ = v148
                case 33:
                    fmt.Println(v1, v5)
                case 42:
                    v149 := GenS13{F0: (v36 & v36), F1: (v5 | (v5 % (((v5 & (v5 - (^(v5))))) | 1)))}
                    _ = v149
                default:
                    v123 = v123
                }
            }
        case -958:
            v150 := GenS14{F0: 42}
            _ = v150
        }
        switch 2147483647 {
        case 2147483647:
            var v151 string = "a"
            _ = v151
        case -9223372036854775808:
            switch "a" {
            case "":
                v3[4294967295] = -1
            default:
                v152 := GenS15{F0: v1, F1: "hello", F2: v1}
                _ = v152
            }
        case -2147483648:
            v36 = (v36 + v36)
        }
    }
    var v153 int8 = 100
    var v154 uint64 = 4294967295
    var v155 uint32 = v36
    var v156 uint64 = (v154 * (v154 - (v154 * (v154 - v154))))
    _ = len(v2)
    for fi157 := 0; fi157 < 9; fi157++ {
        var v158 bool = false
        var v159 int8 = (-(v153))
        _ = fi157
        _ = v158
        _ = v159
    }
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v153:", v153)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v123
    _ = v153
    _ = v154
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
