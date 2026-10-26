package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint16; F1 string; F2 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS13 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0() int {
    v6 := []int{}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint = 142
        _ = fi7
        _ = v8
    }
    var v9 int64 = 1
    if true {
        switch 127 {
        case 4294967295:
            fmt.Println(v9)
        default:
            var v10 uint32 = 1
            _ = v10
        }
    }
    _ = v6
    _ = v9
    return -571
}

func main() {
    var v1 int64 = -667
    v2 := []uint32{549, 485}
    var v3 int16 = -944
    v4 := make(map[uint32]bool)
    var v5 int = int(v3)
    fmt.Println(v1, v3)
    if false {
        v1 -= v1
        if true {
            fmt.Println(v1, v3)
        }
        switch (("" + "a") + "hello") {
        case "a":
            var v11 float32 = 1.5
            _ = v11
        default:
            switch ("" + "hello") {
            case "":
                if false {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        delete(v4, 327)
                        var v13 int64 = (v1 * -682)
                        switch (fi12 * (fi12 % (((-(fi12))) | 1))) {
                        case -9223372036854775808:
                            v5 += v5
                        case -260:
                            for fi14 := 0; fi14 < 4; fi14++ {
                                if true {
                                    delete(v4, 979)
                                    v15 := GenS1{F0: 18446744073709551615}
                                    _ = v15
                                } else {
                                    v1 += v1
                                    fmt.Println(fi14, v13)
                                    if (33 != -30) {
                                        var v16 int16 = v3
                                        v16 -= (v3 << ((255) & 63))
                                        _ = v16
                                    }
                                }
                                _ = fi14
                            }
                        default:
                            var v17 uint64 = 864
                            _ = v17
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v18 uint64 = 0
                    _ = v18
                }
            case "a":
                _ = len(v4)
            }
        }
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint8 = 245
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(fi19, fi21)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        var v22 uint64 = 2
        _ = v22
    }
    var v23 uint32 = 2147483647
    v24 := GenS1{F0: 2}
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            switch 0 {
            case 9223372036854775807:
                v3 = (v3 / ((v3) | 1))
            case 42:
                var v27 int16 = (v3 ^ (v3 << ((374) & 63)))
                _ = v27
            default:
                delete(v4, (v23 / (((^(v23))) | 1)))
                switch "" {
                case "\n\t":
                    var v28 float32 = 0.1
                    _ = v28
                }
                _ = len(v2)
            }
            if true {
                var v29 uint64 = 64
                _ = v29
            }
            v5 += (^(fi26))
            var v30 uint64 = 386
            v2 = append(v2, (v23 << ((65535) & 63)))
            fmt.Println(v30)
            _ = fi26
            _ = v30
        }
        _ = fi25
    }
    v31 := GenS1{F0: 960}
    if ((42 >= 127) || (v23 <= v23)) {
        v32 := GenS1{F0: 338}
        var v33 int64 = (-(v1))
        _ = v32
        _ = v33
    } else {
        v5 = v5
    }
    switch 127 {
    case -215:
        v34 := GenS1{F0: 2}
        _ = v34
    default:
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS1{F0: 2}
            _ = fi35
            _ = v36
        }
    }
    v37 := GenS2{F0: v3, F1: v3}
    var v38 int32 = 42
    for fi39 := 0; fi39 < 5; fi39++ {
        switch "foo bar" {
        case "\n\t":
            for fi40 := 0; fi40 < 16; fi40++ {
                v5 += (v5 ^ (^(fi39)))
                switch "" {
                case "a":
                    switch genFn0() {
                    case -26:
                        fmt.Println(v38, fi39)
                        var v41 string = ("" + "\n\t")
                        v5 = 128
                        v42 := GenS3{F0: -83}
                        _ = v41
                        _ = v42
                    case 127:
                        v43 := GenS4{F0: 2147483647, F1: (v3 * (v3 >> ((65535) & 63)))}
                        _ = v43
                    default:
                        var v44 uint8 = 255
                        v45 := v4[v23]
                        _ = v44
                        _ = v45
                    }
                }
                v3 = int16(v38)
                _ = fi40
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (!(false && false)) {
            v46 := GenS3{F0: 1}
            _ = v46
        } else {
            var v47 string = "a"
            var v48 int8 = 1
            _ = v47
            _ = v48
        }
        _ = fi39
    }
    switch (v5 | (v5 ^ 2147483647)) {
    case -2147483648:
        v49 := GenS2{F0: 1000, F1: (v3 >> ((0) & 63))}
        _ = v49
    case 32767:
        v23 = 2
    case 0:
        for fi50 := 0; fi50 < 11; fi50++ {
            v51 := GenS4{F0: 1, F1: 42}
            _ = fi50
            _ = v51
        }
    }
    if (4294967295 > 65535) {
        if true {
            var v52 uint = 434
            var v53 uint16 = 42
            _ = v52
            _ = v53
        }
        _ = len(v4)
    } else {
        for fi54 := 0; fi54 < 17; fi54++ {
            if false {
                if (!(!(((fi54 - (fi54 * int(v3))) < -9223372036854775808) || false))) {
                    var v55 int64 = v1
                    _ = v55
                } else {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v3, v23)
                    case "":
                        var v56 uint16 = 932
                        _ = v56
                    case "\n\t":
                        fmt.Println(v3, v23, v23)
                    default:
                        v38 -= (v38 % ((v38) | 1))
                    }
                    switch "hello" {
                    case "":
                        var v57 uint8 = 1
                        _ = v57
                    case "a":
                        for fi58 := 0; fi58 < 8; fi58++ {
                            var v59 int32 = v38
                            fmt.Println(fi58)
                            _ = fi58
                            _ = v59
                        }
                    default:
                        v60 := GenS5{F0: (v23 % (((v23 << ((1) & 63))) | 1)), F1: 422}
                        _ = v60
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                switch (-(fi54)) {
                case 2147483647:
                    if (0 >= (v23 << ((635) & 63))) {
                        for fi61 := 0; fi61 < 20; fi61++ {
                            if false {
                                var v62 int8 = 42
                                _ = v62
                            }
                            _ = fi61
                        }
                        if false {
                            v1 -= (v1 - -1)
                        }
                        switch (^(fi54)) {
                        case -966:
                            switch ("hello" + (("hello" + "\n\t") + "")) {
                            case "":
                                fmt.Println(v23, fi54, v23)
                            default:
                                fmt.Println(v1, v1)
                            }
                        }
                    }
                default:
                    var v63 float64 = (((1e10 / (((0.1 / (1e10 / 3.14159)) / (((0.1 + -2.25) + 0.0) / 0.1)) - 1.5)) - ((1.5 * 0.1) / -2.25)) - 0.1)
                    _ = v63
                }
                if ((42 != (v38 << ((2) & 63))) || ((42 == 448) || ((false || (42 <= 421)) || true))) {
                    for fi64 := 0; fi64 < 7; fi64++ {
                        if true {
                            if false {
                                switch fi54 {
                                case 262:
                                    v65 := GenS4{F0: 42, F1: (-(v3))}
                                    v38 += (v38 * (v38 % (((^(v38))) | 1)))
                                    _ = v65
                                case 0:
                                    var v66 float64 = (0.1 / 3.14159)
                                    _ = v66
                                case -9223372036854775808:
                                    if ((v3 % ((v3) | 1)) == v3) {
                                        var v67 int16 = v3
                                        v3 += (v3 % (((^(v3))) | 1))
                                        _ = v67
                                    }
                                    switch "" {
                                    case "a":
                                        v68 := GenS5{F0: (v23 | v23), F1: 887}
                                        _ = v68
                                    case "foo bar":
                                        var v69 int32 = (^(v38))
                                        _ = v69
                                    }
                                    switch "" {
                                    case "a":
                                        v3 = (v3 & (^(v3)))
                                        fmt.Println(v38, v5)
                                    case "":
                                        switch v5 {
                                        case 4294967295:
                                            if (true || (v3 >= int16(v5))) {
                                                if true {
                                                    v1 += (v1 % (((v1 >> ((255) & 63))) | 1))
                                                    var v70 int16 = (v3 & 874)
                                                    v71 := v4[v23]
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    v5 += v5
                                                }
                                                if (0.0 < ((float32(3.14159) + 0.1) / 0.1)) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                var v72 int8 = 1
                                                if true {
                                                    if false {
                                                        v3 = v3
                                                        if true {
                                                            switch genFn0() {
                                                            case 142:
                                                                v73 := GenS6{F0: -719, F1: v3}
                                                                _ = v73
                                                            default:
                                                                fmt.Println(fi64, v72, v5)
                                                            }
                                                            switch "foo bar" {
                                                            case "":
                                                                var v74 uint16 = 255
                                                                _ = v74
                                                            case "a":
                                                                v38 = v38
                                                            default:
                                                                if (127 < 2) {
                                                                    var v75 int = genFn0()
                                                                    v5 -= v75
                                                                    if false {
                                                                        if ((((-(v72)) >= v72) || true) && true) {
                                                                            v76 := GenS7{F0: 820, F1: 42}
                                                                            _ = v76
                                                                        } else {
                                                                            v38 -= -76
                                                                            _ = len(v4)
                                                                            v77 := GenS2{F0: (v3 | v3), F1: v3}
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v78 int = -2147483648
                                                                        _ = v78
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v23
                                                                        }
                                                                        v79 := GenS6{F0: v3, F1: v3}
                                                                        _ = v79
                                                                    }
                                                                    _ = v75
                                                                } else {
                                                                    fmt.Println(fi64, v38, v38)
                                                                    switch 128 {
                                                                    case 953:
                                                                        switch (^(v5)) {
                                                                        case -2147483648:
                                                                            fmt.Println(fi54, fi54, fi64)
                                                                        default:
                                                                            switch fi64 {
                                                                            case 128:
                                                                                if true {
                                                                                    var v80 bool = (0.0 == (3.14159 - 0.0))
                                                                                    v38 -= 200
                                                                                    _ = v80
                                                                                } else {
                                                                                    v72 = 127
                                                                                }
                                                                            default:
                                                                                if (!((!true) && false)) {
                                                                                    v3 += v3
                                                                                    v3 = -255
                                                                                    switch ((("a" + ("foo bar" + "")) + "") + "hello") {
                                                                                    case "hello":
                                                                                        var v81 bool = false
                                                                                        _ = v81
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        fmt.Println(v3, v3, fi54)
                                                                                        v5 -= v5
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    if false {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        switch (-(fi54)) {
                                                                                        case 127:
                                                                                            v82 := GenS8{F0: (fi54 + (^(v5)))}
                                                                                            _ = v82
                                                                                        case 128:
                                                                                            switch -242 {
                                                                                            case 40:
                                                                                                v38 += v38
                                                                                            case -1:
                                                                                                var v83 int8 = (v72 | 1)
                                                                                                _ = v83
                                                                                            case -9223372036854775808:
                                                                                                if (("hello" + "\n\t") > ("a" + "a")) {
                                                                                                    if true {
                                                                                                        var v84 uint16 = 376
                                                                                                        var v85 bool = (v23 != v23)
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    if (v1 > (v1 - (v1 << ((548) & 63)))) {
                                                                                                        if true {
                                                                                                            v3 += (^(v3))
                                                                                                            switch genFn0() {
                                                                                                            case 4294967295:
                                                                                                                v4[(v23 + v23)] = true
                                                                                                            }
                                                                                                            if false {
                                                                                                                fmt.Println(v5, fi54)
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 65535
                                                                                                                }
                                                                                                                if true {
                                                                                                                    var v86 int64 = (v1 & v1)
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                                v72 += (v72 >> ((2) & 63))
                                                                                                            }
                                                                                                        } else {
                                                                                                            v3 = (v3 ^ (^(v3)))
                                                                                                        }
                                                                                                        v4[(v23 >> ((1) & 63))] = false
                                                                                                    }
                                                                                                    fmt.Println(fi64, v72, fi54)
                                                                                                } else {
                                                                                                    var v87 uint16 = 2
                                                                                                    v23 = v23
                                                                                                    if (true || (((1.5 > 1.5) || false) && false)) {
                                                                                                        var v88 int = genFn0()
                                                                                                        _ = v88
                                                                                                    }
                                                                                                    _ = v87
                                                                                                }
                                                                                            }
                                                                                            v89 := GenS7{F0: 255, F1: 32767}
                                                                                            _ = v89
                                                                                        default:
                                                                                            fmt.Println(fi54, v5)
                                                                                        }
                                                                                        v38 += v38
                                                                                    } else {
                                                                                        var v90 uint8 = 42
                                                                                        _ = v90
                                                                                    }
                                                                                    v38 = (-(v38))
                                                                                    var v91 uint64 = 4294967295
                                                                                    if false {
                                                                                        switch fi54 {
                                                                                        case 496:
                                                                                            if ((("a" >= "") || false) && true) {
                                                                                                if (false || true) {
                                                                                                    var v92 float64 = (float64((0.0 + (0.0 + 0.0))) - (3.14159 * (-2.25 * -2.25)))
                                                                                                    v4[v23] = true
                                                                                                    var v93 uint = 42
                                                                                                    switch ("hello" + (("hello" + "hello") + "")) {
                                                                                                    case "foo bar":
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                    var v94 uint8 = 36
                                                                                                    v23 -= 0
                                                                                                    v5 -= (-(v5))
                                                                                                    _ = v92
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                }
                                                                                                if false {
                                                                                                    v23 = (v23 * 208)
                                                                                                    var v95 uint8 = 17
                                                                                                    _ = v95
                                                                                                } else {
                                                                                                    v1 += v1
                                                                                                }
                                                                                            } else {
                                                                                                v23 -= v23
                                                                                            }
                                                                                        case 32767:
                                                                                            v2 = append(v2, (v23 & v23))
                                                                                        case 42:
                                                                                            if ((false && false) && true) {
                                                                                                v1 -= -9223372036854775808
                                                                                                var v96 float32 = 1e10
                                                                                                var v97 int32 = -2147483648
                                                                                                if (!(1 > 2)) {
                                                                                                    v3 = -1
                                                                                                    v38 = v38
                                                                                                }
                                                                                                var v98 int = (v5 * 127)
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                                _ = v98
                                                                                            } else {
                                                                                                fmt.Println(v91)
                                                                                                v91 -= v91
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v72, v91)
                                                                                        }
                                                                                        v1 += v1
                                                                                        if true {
                                                                                            delete(v4, (^(v23)))
                                                                                        }
                                                                                    } else {
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v3 += (^(v3))
                                                                                        case "a":
                                                                                            v99 := GenS3{F0: v72}
                                                                                            _ = v99
                                                                                        case "\n\t":
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        v23 += (v23 * (v23 % (((v23 << ((871) & 63))) | 1)))
                                                                                        var v100 float32 = float32(3.14159)
                                                                                        v101 := GenS9{F0: v3, F1: true}
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    }
                                                                                    _ = v91
                                                                                }
                                                                            }
                                                                        }
                                                                    case 128:
                                                                        v38 += v38
                                                                    case 655:
                                                                        var v102 int = (fi64 * (fi64 % (((v5 << ((514) & 63))) | 1)))
                                                                        _ = v102
                                                                    }
                                                                    v1 = v1
                                                                }
                                                            }
                                                            fmt.Println(fi64, v1, fi54)
                                                        }
                                                        fmt.Println(v23, v23)
                                                    }
                                                    switch genFn0() {
                                                    case 2147483647:
                                                        switch genFn0() {
                                                        case 4294967295:
                                                            v1 += (^(v1))
                                                        case 0:
                                                            v38 -= (v38 | (v38 % (((v38 ^ int32(v1))) | 1)))
                                                        default:
                                                            var v103 uint32 = v23
                                                            _ = v103
                                                        }
                                                    case 877:
                                                        var v104 int32 = (v38 | v38)
                                                        _ = v104
                                                    }
                                                    v72 += (v72 * (v72 & (^(v72))))
                                                    var v105 uint16 = 2
                                                    _ = v105
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = (v23 / (((v23 * v23)) | 1))
                                                    }
                                                    fmt.Println(v23, fi54, v38)
                                                    switch "hello" {
                                                    case "hello":
                                                        v106 := GenS10{F0: 467, F1: (("a" + (("\n\t" + "\n\t") + "\n\t")) + ""), F2: 42}
                                                        _ = v106
                                                    case "a":
                                                        v107 := GenS3{F0: (v72 ^ (v72 ^ v72))}
                                                        _ = v107
                                                    case "\n\t":
                                                        v2 = append(v2, (v23 % ((v23) | 1)))
                                                        if len(v2) > 0 {
                                                            v2[0] = v23
                                                        }
                                                        if false {
                                                            v38 = (-(v38))
                                                            var v108 int = (v5 / ((fi64) | 1))
                                                            var v109 int16 = v3
                                                            fmt.Println(fi54)
                                                            _ = v108
                                                            _ = v109
                                                        } else {
                                                            v38 += -870
                                                            v110 := GenS11{F0: 4294967295}
                                                            _ = v110
                                                        }
                                                        var v111 string = "foo bar"
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483647
                                                        }
                                                        v1 = -1
                                                        _ = v111
                                                    default:
                                                        v112 := v4[(^(v23))]
                                                        _ = v112
                                                    }
                                                }
                                                _ = v72
                                            } else {
                                                v1 -= (v1 % (((^(v1))) | 1))
                                            }
                                            var v113 int32 = v38
                                            _ = v113
                                        case 42:
                                            v38 = v38
                                        case -1:
                                            if len(v2) > 0 {
                                                v2[0] = v23
                                            }
                                        }
                                    default:
                                        var v114 int64 = 1
                                        _ = v114
                                    }
                                default:
                                    v115 := GenS12{F0: 1, F1: 2, F2: 127}
                                    var v116 bool = false
                                    v116 = (580 < (v23 >> ((476) & 63)))
                                    _ = v115
                                    _ = v116
                                }
                                switch v5 {
                                case 2147483647:
                                    if true {
                                        if false {
                                            v117 := GenS3{F0: 100}
                                            v118 := GenS8{F0: fi54}
                                            _ = v117
                                            _ = v118
                                        } else {
                                            switch (fi64 + (-(fi54))) {
                                            case 2147483647:
                                                fmt.Println(v1, fi64)
                                            case 127:
                                                switch fi64 {
                                                case 1:
                                                    switch -233 {
                                                    case 32767:
                                                        if true {
                                                            v23 = (v23 * v23)
                                                        }
                                                    }
                                                case 32767:
                                                    if false {
                                                        v1 = v1
                                                        var v119 string = "foo bar"
                                                        _ = v119
                                                    }
                                                case -9223372036854775808:
                                                    v1 += v1
                                                default:
                                                    delete(v4, 127)
                                                }
                                            case 441:
                                                var v120 uint16 = 32767
                                                if true {
                                                    var v121 int = v5
                                                    _ = v121
                                                }
                                                var v122 float64 = (3.14159 / 1e10)
                                                _ = v120
                                                _ = v122
                                            }
                                            _ = len(v4)
                                        }
                                        if (v38 >= (v38 ^ v38)) {
                                            switch 0 {
                                            case -9223372036854775808:
                                                fmt.Println(v1, v3)
                                            case -1:
                                                v5 = 1
                                            case 62:
                                                var v123 uint8 = 212
                                                _ = v123
                                            }
                                            fmt.Println(fi54, v3, fi64)
                                            var v124 float64 = 1e10
                                            _ = v124
                                        }
                                        fmt.Println(v5, v3)
                                    } else {
                                        if false {
                                            v125 := v4[(v23 % ((531) | 1))]
                                            _ = v125
                                        }
                                        var v126 int32 = v38
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v126
                                    }
                                }
                            } else {
                                var v127 int16 = v3
                                switch ("hello" + ("hello" + "foo bar")) {
                                case "\n\t":
                                    _ = len(v4)
                                case "foo bar":
                                    v4[v23] = (68 >= 173)
                                }
                                fmt.Println(v3, v127, v3)
                                _ = v127
                            }
                            var v128 uint64 = 18446744073709551615
                            switch 83 {
                            case 0:
                                fmt.Println(v38)
                            case 2147483647:
                                switch ("a" + ((((("foo bar" + "foo bar") + "a") + "foo bar") + "a") + "\n\t")) {
                                case "\n\t":
                                    v129 := GenS3{F0: -61}
                                    v130 := GenS6{F0: v3, F1: 1}
                                    _ = v129
                                    _ = v130
                                case "a":
                                    v1 += v1
                                case "hello":
                                    fmt.Println(v1)
                                    if len(v2) > 0 {
                                        v2[0] = v23
                                    }
                                }
                            case 128:
                                switch 9223372036854775807 {
                                case 32768:
                                    var v131 int8 = 124
                                    _ = v131
                                case 127:
                                    if (892 <= 42) {
                                        var v132 bool = true
                                        v133 := GenS13{F0: (v132 || (0.1 >= ((float32((-2.25 - 3.14159)) / 1.0) + ((((1e10 / ((float32((1e10 * float64(0.0))) / (1.5 * (((((0.1 / float32(-2.25)) * 3.14159) + 1e10) - 1.5) * -2.25))) + 1.5)) * 1e10) * 1e10) * 3.14159)))), F1: 1, F2: v23}
                                        fmt.Println(v128)
                                        _ = v132
                                        _ = v133
                                    } else {
                                        var v134 bool = ((v1 + v1) < -726)
                                        _ = v134
                                    }
                                }
                            default:
                                v3 = v3
                            }
                            _ = v128
                        }
                        fmt.Println(fi54, v23, fi64)
                        var v135 float64 = (3.14159 / 1.5)
                        _ = fi64
                        _ = v135
                    }
                    delete(v4, (v23 / ((65535) | 1)))
                    fmt.Println(v1, fi54)
                }
            }
            fmt.Println(v5, fi54)
            _ = fi54
        }
        for fi136 := 0; fi136 < 16; fi136++ {
            v137 := GenS3{F0: -128}
            _ = fi136
            _ = v137
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
