package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 int16; F1 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64; F1 uint64; F2 uint }
type GenS15 struct { F0 uint; F1 uint8 }
type GenS16 struct { F0 bool; F1 int; F2 bool }
type GenS17 struct { F0 uint; F1 uint8 }
type GenS18 struct { F0 int16; F1 int64 }
type GenS19 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 float32 = 3.14159
    v3 = (v3 * float32(-2.25))
    v3 += (v3 + 3.14159)
    if true {
        v3 += float32(0.0)
        if (20 <= 0) {
            v4 := GenS1{F0: 18446744073709551615}
            switch 128 {
            case 0:
                var v5 int8 = -1
                _ = v5
            case 4294967295:
                var v6 uint32 = 65535
                _ = v6
            case 42:
                var v7 int64 = -9223372036854775808
                _ = v7
            default:
                var v8 uint64 = 487
                _ = v8
            }
            fmt.Println(v3)
            _ = v4
        } else {
            v3 -= 1.5
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS2{F0: -912, F1: 2147483648}
                fmt.Println(v3, v3)
                _ = fi9
                _ = v10
            }
            v3 += v3
        }
    }
    v11 := GenS3{F0: -313, F1: (786 != -116), F2: 1}
    if (32767 == 1) {
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 uint16 = 2
            for fi14 := 0; fi14 < 4; fi14++ {
                v3 = float32((float64(1e10) + ((0.0 * 1.5) - (3.14159 + 1e10))))
                switch "foo bar" {
                case "a":
                    v2[((("hello" + "a") + ("a" + ("" + ((((("\n\t" + ("hello" + ("hello" + ("\n\t" + "\n\t")))) + ("\n\t" + "\n\t")) + "") + (("a" + "\n\t") + "foo bar")) + "foo bar")))) + "foo bar")] = (true || (0.0 <= 0.1))
                }
                switch (("hello" + ("a" + ("foo bar" + "foo bar"))) + "hello") {
                case "a":
                    var v15 uint32 = 2
                    _ = v15
                case "hello":
                    switch fi12 {
                    case -985:
                        v13 = 32767
                    case 128:
                        switch "a" {
                        case "a":
                            v16 := []int{fi12, (fi14 + fi14), (^(fi14)), -2147483648}
                            _ = v16
                        case "":
                            v3 = -2.25
                        case "hello":
                            fmt.Println(fi14)
                        default:
                            switch ("" + ("\n\t" + "foo bar")) {
                            case "hello":
                                var v17 int8 = -121
                                _ = v17
                            }
                        }
                    case 32767:
                        if false {
                            var v18 int16 = -1
                            _ = v18
                        } else {
                            var v19 int8 = 100
                            if (!(2147483648 < 698)) {
                                v20 := []int{fi12, (fi12 << ((18446744073709551615) & 63)), 42, 0}
                                _ = v20
                            } else {
                                fmt.Println(fi14, v13, v13)
                                if (0 <= 255) {
                                    v19 -= (v19 ^ 100)
                                    if (0.1 > (-2.25 - 1.5)) {
                                        v19 -= (v19 >> ((2) & 63))
                                        v19 = v19
                                        if (-9223372036854775808 < 9223372036854775807) {
                                            var v21 int8 = (v19 - v19)
                                            v3 -= float32(1.5)
                                            var v22 float32 = 3.14159
                                            _ = v21
                                            _ = v22
                                        }
                                        v19 = 42
                                    }
                                } else {
                                    v23 := []int{4294967295, -1, fi14, 9223372036854775807}
                                    v13 += v13
                                    _ = v23
                                }
                                if (v19 <= (v19 << ((783) & 63))) {
                                    fmt.Println(fi14, fi12)
                                    v24 := []int{fi12, 32767, fi12}
                                    var v25 string = ("" + "hello")
                                    var v26 string = v25
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                } else {
                                    switch (((((("foo bar" + "foo bar") + ("a" + "a")) + ("" + "foo bar")) + ("a" + ("foo bar" + ("foo bar" + "foo bar")))) + (("\n\t" + (("foo bar" + ("a" + ("foo bar" + ("" + (("" + "a") + "\n\t"))))) + (("\n\t" + ("a" + "hello")) + ""))) + "hello")) + ("" + (("foo bar" + "hello") + "hello"))) {
                                    case "a":
                                        v27 := GenS2{F0: -971, F1: 2}
                                        _ = v27
                                    case "\n\t":
                                        var v28 int32 = 42
                                        _ = v28
                                    case "hello":
                                        v19 -= (v19 >> ((42) & 63))
                                    default:
                                        v29 := GenS3{F0: 0, F1: false, F2: 4294967295}
                                        _ = v29
                                    }
                                }
                                if (((v19 < v19) || (!(true && true))) && false) {
                                    fmt.Println(fi14, v19, fi14)
                                }
                            }
                            _ = v19
                        }
                    default:
                        switch (fi12 & (^(fi12))) {
                        case 4294967295:
                            v30 := GenS1{F0: 18446744073709551615}
                            v31 := GenS4{F0: 0}
                            _ = v30
                            _ = v31
                        }
                        if true {
                            v32 := GenS5{F0: "", F1: 0, F2: 42}
                            var v33 uint32 = 42
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                var v34 float32 = v3
                                _ = v34
                            case "":
                                switch (("a" + "\n\t") + "foo bar") {
                                case "hello":
                                    v33 -= (v33 + 2147483647)
                                case "foo bar":
                                    v1[128] = "a"
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        if (18446744073709551615 != 65535) {
                                            var v35 int32 = 151
                                            _ = v35
                                        } else {
                                            v3 -= v3
                                        }
                                    }
                                }
                            default:
                                v36 := []int{fi12, fi12, -61}
                                _ = v36
                            }
                            _ = v32
                            _ = v33
                        } else {
                            v37 := []int{(fi14 * fi14), (fi14 >> ((678) & 63)), (fi12 & fi12)}
                            delete(v2, ("\n\t" + ("" + "foo bar")))
                            v38 := GenS6{F0: (v13 / (((^(v13))) | 1)), F1: 100}
                            _ = v37
                            _ = v38
                        }
                        v39 := GenS7{F0: 1}
                        _ = v39
                    }
                case "":
                    var v40 uint16 = v13
                    _ = v40
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        for fi41 := 0; fi41 < 19; fi41++ {
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 uint64 = 2
                v44 := GenS3{F0: 255, F1: true, F2: v43}
                switch ("hello" + "") {
                case "hello":
                    v43 = v43
                case "":
                    if ((3.14159 * 0.0) != 3.14159) {
                        v45 := GenS3{F0: 1000, F1: (0 > 255), F2: (v43 * (v43 ^ 262))}
                        v46 := []int{fi42}
                        v47 := GenS4{F0: 1}
                        _ = v45
                        _ = v46
                        _ = v47
                    } else {
                        _ = len(v2)
                    }
                default:
                    v48 := []int{1, fi41, (-(fi41))}
                    _ = v48
                }
                _ = fi42
                _ = v43
                _ = v44
            }
            switch (fi41 + fi41) {
            case -1:
                switch (fi41 << ((65535) & 63)) {
                case 128:
                    var v49 float32 = (1e10 / v3)
                    _ = v49
                case 4294967295:
                    var v50 uint = 18446744073709551615
                    _ = v50
                }
            default:
                v51 := GenS8{F0: -1}
                _ = v51
            }
            _ = fi41
        }
        var v52 uint16 = 0
        _ = v52
    } else {
        if false {
            if (!(18446744073709551615 > 1)) {
                if false {
                    v3 -= 1e10
                } else {
                    for fi53 := 0; fi53 < 6; fi53++ {
                        if (0 > 255) {
                            v54 := GenS9{F0: -783, F1: 1}
                            var v55 int = (fi53 >> ((176) & 63))
                            if false {
                                v55 = v55
                                if false {
                                    if false {
                                        var v56 uint = 0
                                        if (!(42 < 2147483647)) {
                                            var v57 bool = false
                                            v58 := v2[""]
                                            v3 += (v3 / v3)
                                            _ = v57
                                            _ = v58
                                        } else {
                                            fmt.Println(v55)
                                            v56 = v56
                                        }
                                        v3 += v3
                                        _ = v56
                                    } else {
                                        if true {
                                            fmt.Println(v55, v3, v3)
                                            for fi59 := 0; fi59 < 10; fi59++ {
                                                v3 += (float32(float64(float32(0.1))) - (-2.25 / 1.5))
                                                _ = fi59
                                            }
                                        } else {
                                            v60 := []int{(fi53 ^ 42), (fi53 | 128)}
                                            for fi61 := 0; fi61 < 8; fi61++ {
                                                if len(v60) > 0 {
                                                    v60[0] = (fi61 + (fi53 & (-(v55))))
                                                }
                                                v62 := GenS3{F0: 32767, F1: true, F2: 407}
                                                var v63 uint64 = 1
                                                _ = fi61
                                                _ = v62
                                                _ = v63
                                            }
                                            var v64 float32 = v3
                                            _ = v60
                                            _ = v64
                                        }
                                        v3 += 1.5
                                        v55 = (fi53 & fi53)
                                    }
                                    v65 := []int{v55, v55, (fi53 * (v55 ^ fi53)), -9223372036854775808}
                                    var v66 int = (fi53 ^ fi53)
                                    _ = v65
                                    _ = v66
                                } else {
                                    v1[fi53] = "foo bar"
                                }
                            }
                            _ = v54
                            _ = v55
                        } else {
                            v67 := GenS9{F0: 2147483647, F1: -1}
                            _ = v67
                        }
                        _ = fi53
                    }
                    if false {
                        var v68 uint64 = 385
                        _ = v68
                    } else {
                        for fi69 := 0; fi69 < 6; fi69++ {
                            var v70 int32 = 42
                            var v71 uint32 = 725
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                        switch ("\n\t" + "") {
                        case "hello":
                            if true {
                                if ((!(0.1 >= 1e10)) || false) {
                                    for fi72 := 0; fi72 < 16; fi72++ {
                                        if (0 < -693) {
                                            switch fi72 {
                                            case -9223372036854775808:
                                                if false {
                                                    for fi73 := 0; fi73 < 4; fi73++ {
                                                        v3 -= v3
                                                        v3 = ((0.1 * (v3 + v3)) - ((v3 / ((1.5 + ((v3 * ((v3 + float32((1e10 * -2.25))) - 3.14159)) * (1.5 * v3))) + float32(3.14159))) - 1e10))
                                                        var v74 bool = true
                                                        _ = fi73
                                                        _ = v74
                                                    }
                                                    for fi75 := 0; fi75 < 13; fi75++ {
                                                        v3 -= v3
                                                        _ = fi75
                                                    }
                                                    if (!false) {
                                                        for fi76 := 0; fi76 < 3; fi76++ {
                                                            switch 32767 {
                                                            case 42:
                                                                v77 := []int{(fi76 % (((fi72 & fi72)) | 1))}
                                                                _ = v77
                                                            }
                                                            var v78 bool = (-448 < -459)
                                                            _ = fi76
                                                            _ = v78
                                                        }
                                                        _ = len(v1)
                                                    } else {
                                                        var v79 int16 = 32767
                                                        _ = v79
                                                    }
                                                }
                                            case -2147483648:
                                                fmt.Println(fi72, fi72)
                                            default:
                                                switch 2147483647 {
                                                case -2147483648:
                                                    v80 := GenS10{F0: 2}
                                                    _ = v80
                                                case 128:
                                                    var v81 uint64 = 59
                                                    _ = v81
                                                }
                                            }
                                            if (true || true) {
                                                for fi82 := 0; fi82 < 8; fi82++ {
                                                    v83 := []int{-1}
                                                    v3 = float32(-2.25)
                                                    _ = fi82
                                                    _ = v83
                                                }
                                                for fi84 := 0; fi84 < 16; fi84++ {
                                                    v85 := GenS11{F0: "\n\t", F1: 885}
                                                    var v86 string = "foo bar"
                                                    v2[v86] = (!false)
                                                    _ = fi84
                                                    _ = v85
                                                    _ = v86
                                                }
                                                v87 := GenS2{F0: 704, F1: 717}
                                                _ = v87
                                            } else {
                                                fmt.Println(fi72, v3)
                                                v88 := GenS9{F0: 42, F1: 0}
                                                _ = v88
                                            }
                                        }
                                        var v89 int8 = 127
                                        _ = fi72
                                        _ = v89
                                    }
                                } else {
                                    for fi90 := 0; fi90 < 19; fi90++ {
                                        var v91 int = 730
                                        switch "" {
                                        case "foo bar":
                                            if false {
                                                var v92 uint = 0
                                                v3 = 1.5
                                                var v93 uint8 = 243
                                                _ = v92
                                                _ = v93
                                            }
                                        case "":
                                            var v94 int16 = 42
                                            _ = v94
                                        case "a":
                                            v91 += v91
                                        }
                                        _ = fi90
                                        _ = v91
                                    }
                                }
                                fmt.Println(v3)
                                v95 := []int{127, 4294967295, -9223372036854775808}
                                _ = v95
                            }
                        case "a":
                            v3 = v3
                        default:
                            for fi96 := 0; fi96 < 3; fi96++ {
                                for fi97 := 0; fi97 < 18; fi97++ {
                                    switch fi97 {
                                    case 1:
                                        v3 -= v3
                                    case 32767:
                                        if false {
                                            v98 := []int{}
                                            if true {
                                                if false {
                                                    v3 -= v3
                                                } else {
                                                    switch fi97 {
                                                    case -101:
                                                        fmt.Println(fi96)
                                                    case -248:
                                                        v99 := GenS12{F0: -32768, F1: 4294967295}
                                                        _ = v99
                                                    }
                                                }
                                                v100 := GenS5{F0: ((("a" + ((("\n\t" + ("hello" + ((("" + ("\n\t" + ("a" + "a"))) + "hello") + "a"))) + "") + "")) + ("hello" + "")) + ("foo bar" + "foo bar")), F1: 2147483647, F2: 1}
                                                _ = len(v2)
                                                _ = v100
                                            } else {
                                                v101 := GenS3{F0: 1, F1: true, F2: 1}
                                                _ = v101
                                            }
                                            v102 := GenS13{F0: 523}
                                            _ = v102
                                            _ = v98
                                        } else {
                                            if true {
                                                var v103 int16 = -1
                                                var v104 int8 = 42
                                                v105 := GenS11{F0: "foo bar", F1: 4294967295}
                                                var v106 bool = true
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                            } else {
                                                if (!(!false)) {
                                                    switch fi97 {
                                                    case 796:
                                                        v3 += (v3 - ((1e10 * v3) * v3))
                                                    default:
                                                        var v107 float64 = float64((v3 * ((0.0 * v3) / (1.5 / 3.14159))))
                                                        _ = v107
                                                    }
                                                    var v108 int16 = 612
                                                    _ = v108
                                                } else {
                                                    fmt.Println(fi97, v3)
                                                    if false {
                                                        var v109 float64 = ((float64(-2.25) - -2.25) + float64(v3))
                                                        var v110 uint32 = 1
                                                        _ = v109
                                                        _ = v110
                                                    } else {
                                                        v2["foo bar"] = ((0 >= 1) && (!false))
                                                        var v111 uint = 65535
                                                        v111 = (v111 | v111)
                                                        _ = v111
                                                    }
                                                    switch 1 {
                                                    case 32767:
                                                        v3 = v3
                                                    }
                                                    v3 -= (-2.25 - v3)
                                                }
                                            }
                                            if false {
                                                v3 = 0.1
                                            }
                                            v3 -= v3
                                        }
                                    }
                                    v112 := GenS2{F0: 219, F1: 732}
                                    _ = fi97
                                    _ = v112
                                }
                                var v113 uint32 = 0
                                v114 := []int{(fi96 % ((fi96) | 1))}
                                _ = fi96
                                _ = v113
                                _ = v114
                            }
                        }
                        for fi115 := 0; fi115 < 8; fi115++ {
                            v3 = (((float32(0.0) - v3) / float32(3.14159)) + v3)
                            var v116 uint8 = 2
                            _ = fi115
                            _ = v116
                        }
                    }
                    var v117 int8 = 42
                    _ = v117
                }
            } else {
                fmt.Println(v3)
            }
            v118 := GenS14{F0: 4294967295, F1: 1, F2: 255}
            _ = v118
        } else {
            for fi119 := 0; fi119 < 2; fi119++ {
                v120 := []int{523, (fi119 ^ fi119), 9223372036854775807}
                for fi121 := 0; fi121 < 10; fi121++ {
                    v122 := v2["\n\t"]
                    v122 = (42 != 2147483647)
                    v3 = v3
                    _ = fi121
                    _ = v122
                }
                switch (("\n\t" + (((("\n\t" + "hello") + "") + ("a" + ("foo bar" + ""))) + ("" + ("a" + "foo bar")))) + "") {
                case "a":
                    v123 := GenS7{F0: 255}
                    _ = v123
                }
                _ = fi119
                _ = v120
            }
            v3 = v3
        }
    }
    var v124 uint16 = 32767
    for fi125 := 0; fi125 < 15; fi125++ {
        if true {
            _ = len(v2)
            switch ("" + ("\n\t" + "hello")) {
            case "\n\t":
                switch (fi125 + (fi125 / ((fi125) | 1))) {
                case -9223372036854775808:
                    v124 = v124
                case -238:
                    if (1 != 1) {
                        var v126 uint = 1
                        var v127 bool = (2147483647 >= 2147483647)
                        _ = v126
                        _ = v127
                    } else {
                        for fi128 := 0; fi128 < 4; fi128++ {
                            switch 1 {
                            case 128:
                                switch "a" {
                                case "hello":
                                    v129 := GenS13{F0: 2147483647}
                                    _ = v129
                                case "foo bar":
                                    v130 := GenS15{F0: 1, F1: 2}
                                    _ = v130
                                default:
                                    var v131 uint32 = 2
                                    _ = v131
                                }
                            case 42:
                                switch "" {
                                case "foo bar":
                                    delete(v1, fi128)
                                case "\n\t":
                                    if false {
                                        switch ("" + ("a" + "foo bar")) {
                                        case "hello":
                                            switch "hello" {
                                            case "hello":
                                                v3 = v3
                                            default:
                                                v3 = 3.14159
                                            }
                                        case "":
                                            if ((1.5 - (0.0 * (-2.25 - (-2.25 / 3.14159)))) < 1e10) {
                                                v132 := []int{fi125}
                                                _ = v132
                                            }
                                        case "a":
                                            v133 := GenS16{F0: (-2147483648 <= 981), F1: 32768, F2: (false || true)}
                                            _ = v133
                                        default:
                                            v134 := []int{}
                                            _ = v134
                                        }
                                    }
                                default:
                                    var v135 uint8 = 255
                                    _ = v135
                                }
                            case 2147483647:
                                v124 = v124
                            default:
                                if (0 <= 1) {
                                    if ((((3.14159 * (3.14159 * (1e10 - 1e10))) / 1.0) * -2.25) == 0.0) {
                                        v124 = 0
                                        if (0 < 733) {
                                            fmt.Println(fi128, fi128, v3)
                                        } else {
                                            var v136 bool = ((!true) && false)
                                            _ = v136
                                        }
                                    }
                                }
                            }
                            var v137 int64 = -1
                            _ = fi128
                            _ = v137
                        }
                        var v138 uint8 = 209
                        _ = v138
                    }
                }
            }
        }
        v3 = v3
        _ = fi125
    }
    v139 := []int{}
    v124 = 2
    switch 0 {
    case 4294967295:
        if len(v139) > 0 {
            v139[0] = 42
        }
    case 2147483647:
        var v140 int8 = 119
        _ = v140
    }
    switch 990 {
    case -9223372036854775808:
        v3 = -2.25
    case -2147483648:
        fmt.Println(v3, v3)
    case -771:
        var v141 bool = ((255 >= 636) || (!false))
        _ = v141
    }
    switch 42 {
    case 754:
        v124 = 2
    case 128:
        _ = len(v1)
    case 32768:
        v124 += 685
    default:
        v1[-1] = "a"
    }
    for fi142 := 0; fi142 < 17; fi142++ {
        for fi143 := 0; fi143 < 6; fi143++ {
            var v144 int = (fi143 >> ((42) & 63))
            v145 := GenS17{F0: 319, F1: 2}
            v3 -= (v3 - (v3 * v3))
            var v146 string = (("hello" + "\n\t") + "")
            if (false || (42 == 42)) {
                if len(v139) >= 2 {
                    v139 = v139[1:len(v139)]
                }
            }
            v124 = 34
            _ = fi143
            _ = v144
            _ = v145
            _ = v146
        }
        if (617 != 65535) {
            var v147 int16 = 0
            _ = v147
        } else {
            switch "hello" {
            case "a":
                v148 := GenS16{F0: true, F1: fi142, F2: true}
                _ = v148
            default:
                var v149 string = ""
                _ = v149
            }
        }
        _ = fi142
    }
    for fi150 := 0; fi150 < 3; fi150++ {
        for fi151 := 0; fi151 < 17; fi151++ {
            if (!(182 == 2147483648)) {
                switch ((("\n\t" + "a") + ((("\n\t" + ("" + "hello")) + "") + "a")) + (("foo bar" + ("foo bar" + "hello")) + "\n\t")) {
                case "":
                    var v152 float64 = 1e10
                    _ = v152
                case "a":
                    v124 += v124
                case "\n\t":
                    v124 -= 32767
                default:
                    var v153 uint16 = v124
                    fmt.Println(fi150, v3, v124)
                    _ = v153
                }
                v154 := GenS18{F0: 749, F1: 9223372036854775807}
                _ = v154
            }
            v155 := GenS19{F0: 32767, F1: 32768}
            _ = fi151
            _ = v155
        }
        for fi156 := 0; fi156 < 3; fi156++ {
            var v157 int32 = 2147483647
            v157 += 1
            v124 = v124
            _ = fi156
            _ = v157
        }
        _ = fi150
    }
    for fi158 := 0; fi158 < 8; fi158++ {
        for fi159 := 0; fi159 < 14; fi159++ {
            var v160 int32 = 2147483647
            _ = fi159
            _ = v160
        }
        if (!((0.0 / 0.1) != 1e10)) {
            v3 -= v3
            v161 := GenS4{F0: 0}
            for fi162 := 0; fi162 < 4; fi162++ {
                v124 = v124
                fmt.Println(v3)
                var v163 float64 = 3.14159
                _ = fi162
                _ = v163
            }
            _ = v161
        }
        _ = fi158
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v124:", v124)
    fmt.Println("slicelen v139:", len(v139), v139)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v124
    _ = v139
    _ = v2
    _ = v3
}
