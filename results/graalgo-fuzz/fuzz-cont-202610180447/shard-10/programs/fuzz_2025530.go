package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -1
    v3 := make(map[string]bool)
    var v4 int8 = 1
    var v5 uint32 = 42
    v6 := []uint32{}
    v7 := GenS1{F0: 243}
    v8 := GenS1{F0: 42}
    fmt.Println(v4)
    v4 += (v4 % (((v4 >> ((0) & 63))) | 1))
    v2 -= v2
    for fi9 := 0; fi9 < 11; fi9++ {
        switch "hello" {
        case "hello":
            fmt.Println(v2, v5, fi9)
        case "a":
            var v10 bool = false
            _ = v10
        case "":
            v1[(v5 << ((18446744073709551615) & 63))] = -2147483648
        default:
            v5 = (v5 * v5)
        }
        _ = fi9
    }
    v5 += v5
    switch "hello" {
    case "foo bar":
        v5 = v5
    case "hello":
        if (!("a" < "\n\t")) {
            switch (v2 - (v2 ^ v2)) {
            case -145:
                var v11 uint = 1
                _ = v11
            case -744:
                for fi12 := 0; fi12 < 2; fi12++ {
                    fmt.Println(v4, v4, fi12)
                    v2 = (fi12 << ((2) & 63))
                    _ = fi12
                }
            case -9223372036854775808:
                var v13 int32 = -612
                _ = v13
            }
            v14 := v3["a"]
            _ = v14
        }
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v5, v2, v4)
            v16 := v3["foo bar"]
            _ = fi15
            _ = v16
        }
    default:
        _ = len(v1)
    }
    if false {
        v17 := GenS1{F0: 2}
        v18 := GenS1{F0: 0}
        fmt.Println(v5)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v4 -= v4
        if false {
            if false {
                var v20 uint = 255
                _ = v20
            }
            v21 := GenS1{F0: 810}
            switch ("\n\t" + "hello") {
            case "\n\t":
                v22 := GenS1{F0: 2}
                _ = v22
            case "a":
                v4 += v4
            case "hello":
                if ((false && (v4 <= (v4 - 60))) && (1 <= 1)) {
                    v23 := GenS2{F0: ((!true) || true), F1: 42, F2: 128}
                    if (-32768 == 0) {
                        if false {
                            if false {
                                var v24 float32 = ((0.1 - 3.14159) - (1e10 / -2.25))
                                if len(v6) > 0 {
                                    v6[0] = v5
                                }
                                v25 := v1[v5]
                                _ = v24
                                _ = v25
                            }
                        }
                        if ((false && (1.5 <= (1e10 / 0.1))) && (2147483647 > 1)) {
                            var v26 uint = 255
                            switch "a" {
                            case "hello":
                                v27 := GenS1{F0: 953}
                                _ = v27
                            case "a":
                                _ = len(v1)
                            case "\n\t":
                                if false {
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        v29 := v1[(^(v5))]
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch "" {
                                            case "":
                                                v5 = 2
                                                v30 := v3[("\n\t" + "\n\t")]
                                                v4 -= v4
                                                _ = v30
                                            }
                                        case "":
                                            v4 -= (v4 | v4)
                                        case "a":
                                            var v31 bool = (0.1 != -2.25)
                                            _ = v31
                                        default:
                                            var v32 string = ""
                                            v5 += v5
                                            _ = v32
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                    v2 += (fi19 - v2)
                                    if false {
                                        if true {
                                            v5 -= v5
                                            for fi33 := 0; fi33 < 15; fi33++ {
                                                switch (fi19 / ((4294967295) | 1)) {
                                                case 342:
                                                    v34 := GenS2{F0: (!(fi33 >= (fi19 & v2))), F1: 0, F2: 127}
                                                    _ = v34
                                                case -55:
                                                    switch int(v5) {
                                                    case 127:
                                                        var v35 uint16 = 32767
                                                        _ = v35
                                                    case -964:
                                                        v2 -= fi33
                                                    default:
                                                        switch fi33 {
                                                        case -2147483648:
                                                            switch fi19 {
                                                            case 4294967295:
                                                                var v36 float64 = float64(1e10)
                                                                _ = v36
                                                            case 127:
                                                                var v37 uint = v26
                                                                _ = v37
                                                            case 1:
                                                                var v38 float32 = 1e10
                                                                _ = v38
                                                            }
                                                        }
                                                    }
                                                }
                                                var v39 int8 = -1
                                                _ = fi33
                                                _ = v39
                                            }
                                        }
                                        var v40 float64 = 1e10
                                        var v41 bool = (277 < 1)
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            default:
                                v26 = (^(v26))
                            }
                            if (473 != v5) {
                                v4 -= 0
                                if true {
                                    switch (("a" + ("hello" + ("a" + "a"))) + "hello") {
                                    case "":
                                        if true {
                                            var v42 float64 = 3.14159
                                            var v43 int = 42
                                            v44 := GenS1{F0: 18446744073709551615}
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        } else {
                                            var v45 float32 = ((3.14159 + 0.1) - 0.1)
                                            for fi46 := 0; fi46 < 18; fi46++ {
                                                switch (fi46 << ((v26) & 63)) {
                                                case 873:
                                                    if ((42 <= 1) && false) {
                                                        v47 := GenS1{F0: 847}
                                                        _ = v47
                                                    } else {
                                                        v6 = append(v6, v5)
                                                        v26 = v26
                                                        v45 = v45
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(fi19, v2, fi46)
                                                }
                                                _ = fi46
                                            }
                                            switch -379 {
                                            case 486:
                                                v48 := v3[("foo bar" + "\n\t")]
                                                _ = v48
                                            case 1:
                                                v45 += 1e10
                                            default:
                                                if len(v6) > 0 {
                                                    v6[0] = v5
                                                }
                                            }
                                            _ = v45
                                        }
                                    }
                                    for fi49 := 0; fi49 < 12; fi49++ {
                                        var v50 float64 = 0.0
                                        switch (fi49 - (v2 & (v2 + fi49))) {
                                        case -689:
                                            v50 -= v50
                                        case 4294967295:
                                            fmt.Println(v50)
                                        case -870:
                                            switch "foo bar" {
                                            case "hello":
                                                switch ("hello" + (("foo bar" + ("hello" + "foo bar")) + "\n\t")) {
                                                case "hello":
                                                    var v51 uint8 = 1
                                                    _ = v51
                                                default:
                                                    switch "" {
                                                    case "a":
                                                        if false {
                                                            if (v26 < (v26 / (((^(v26))) | 1))) {
                                                                v1[371] = (v2 * (fi49 + (fi19 | v2)))
                                                            }
                                                            v50 = v50
                                                            switch -1 {
                                                            case 4294967295:
                                                                v52 := GenS1{F0: 4294967295}
                                                                _ = v52
                                                            default:
                                                                var v53 int = 2147483647
                                                                _ = v53
                                                            }
                                                        }
                                                    }
                                                    v4 = v4
                                                    v54 := GenS3{F0: 1}
                                                    _ = v54
                                                }
                                            case "\n\t":
                                                v55 := GenS4{F0: 1, F1: -922}
                                                _ = v55
                                            case "":
                                                var v56 uint32 = v5
                                                _ = v56
                                            }
                                        }
                                        v4 = -32
                                        _ = fi49
                                        _ = v50
                                    }
                                } else {
                                    v26 += v26
                                    switch "a" {
                                    case "foo bar":
                                        for fi57 := 0; fi57 < 17; fi57++ {
                                            if len(v6) > 0 {
                                                v6[0] = v5
                                            }
                                            _ = fi57
                                        }
                                    case "hello":
                                        fmt.Println(fi19)
                                    }
                                    v2 -= v2
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v58 int16 = -32768
                                    _ = v58
                                case "\n\t":
                                    v59 := GenS5{F0: v4, F1: 65535, F2: ("" + "")}
                                    _ = v59
                                }
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    v2 = 32768
                                case "\n\t":
                                    for fi60 := 0; fi60 < 17; fi60++ {
                                        var v61 int8 = (v4 | 100)
                                        v62 := GenS6{F0: 1000}
                                        v63 := GenS7{F0: 18446744073709551615, F1: 26}
                                        if true {
                                            fmt.Println(v61, v5)
                                            var v64 int64 = 42
                                            v4 = v4
                                            v65 := GenS1{F0: 0}
                                            fmt.Println(v5, v61, v64)
                                            _ = v64
                                            _ = v65
                                        } else {
                                            switch (-(fi19)) {
                                            case 411:
                                                fmt.Println(v26, v5)
                                            case -665:
                                                fmt.Println(v4, fi19)
                                            default:
                                                var v66 bool = (true && false)
                                                _ = v66
                                            }
                                            if (false && true) {
                                                if false {
                                                    var v67 float32 = 0.1
                                                    var v68 bool = false
                                                    fmt.Println(fi19)
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    var v69 uint = v26
                                                    v61 = 42
                                                    fmt.Println(fi60, fi19)
                                                    var v70 uint16 = 2
                                                    _ = v69
                                                    _ = v70
                                                }
                                                var v71 uint8 = 255
                                                v2 = (^(v2))
                                                _ = v71
                                            } else {
                                                var v72 uint64 = 1
                                                _ = v72
                                            }
                                            fmt.Println(fi19, v2)
                                            var v73 uint32 = v5
                                            _ = v73
                                        }
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                default:
                                    if (!true) {
                                        v74 := GenS8{F0: false}
                                        for fi75 := 0; fi75 < 6; fi75++ {
                                            if true {
                                                var v76 string = "\n\t"
                                                _ = v76
                                            } else {
                                                v4 = v4
                                                v4 = (v4 + (v4 - (v4 ^ v4)))
                                                v4 = v4
                                                var v77 int16 = -1
                                                v2 -= 231
                                                v5 += (v5 * v5)
                                                var v78 int = (fi19 >> (((v26 / ((v26) | 1))) & 63))
                                                _ = v77
                                                _ = v78
                                            }
                                            v26 = (v26 >> ((v26) & 63))
                                            _ = fi75
                                        }
                                        v2 -= (-(v2))
                                        _ = v74
                                    }
                                }
                                if false {
                                    var v79 uint8 = 1
                                    v80 := GenS4{F0: (^(v5)), F1: -1}
                                    v6 = append(v6, 1)
                                    _ = v79
                                    _ = v80
                                } else {
                                    v81 := GenS2{F0: false, F1: -939, F2: 105}
                                    _ = v81
                                }
                            }
                            _ = v26
                        }
                        switch (fi19 + (v2 % (((-(fi19))) | 1))) {
                        case 244:
                            var v82 int16 = 1
                            _ = v82
                        case -697:
                            if (-558 != -9223372036854775808) {
                                fmt.Println(v4)
                                v2 -= (^(v2))
                            } else {
                                for fi83 := 0; fi83 < 2; fi83++ {
                                    if ((v4 / ((v4) | 1)) != 127) {
                                        _ = len(v6)
                                    } else {
                                        if true {
                                            _ = len(v6)
                                        } else {
                                            if true {
                                                if (42 <= 481) {
                                                    var v84 uint64 = 768
                                                    var v85 int32 = -676
                                                    if true {
                                                        v86 := v1[(v5 - v5)]
                                                        v86 = -9223372036854775808
                                                        v87 := GenS9{F0: v85, F1: 65535, F2: 65535}
                                                        _ = v86
                                                        _ = v87
                                                    } else {
                                                        v85 = (v85 << ((249) & 63))
                                                        var v88 uint = 380
                                                        _ = v88
                                                    }
                                                    _ = v84
                                                    _ = v85
                                                }
                                                var v89 int64 = 1
                                                _ = v89
                                            }
                                        }
                                    }
                                    delete(v1, v5)
                                    _ = fi83
                                }
                                var v90 uint8 = 128
                                var v91 int = fi19
                                _ = v90
                                _ = v91
                            }
                        default:
                            for fi92 := 0; fi92 < 2; fi92++ {
                                fmt.Println(fi92)
                                var v93 int16 = -1
                                v94 := GenS6{F0: (^(v93))}
                                _ = fi92
                                _ = v93
                                _ = v94
                            }
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v95 := GenS10{F0: 42}
                        _ = v95
                    case "foo bar":
                        for fi96 := 0; fi96 < 14; fi96++ {
                            fmt.Println(fi19, v2)
                            var v97 int16 = 42
                            _ = fi96
                            _ = v97
                        }
                    case "a":
                        v98 := GenS3{F0: -2147483648}
                        _ = v98
                    default:
                        var v99 bool = false
                        _ = v99
                    }
                    _ = v23
                } else {
                    switch ("foo bar" + "") {
                    case "foo bar":
                        var v100 uint = 1
                        _ = v100
                    case "":
                        if (459 != 42) {
                            if false {
                                fmt.Println(v5, v5, fi19)
                                fmt.Println(fi19)
                                var v101 uint64 = 4294967295
                                _ = v101
                            }
                            delete(v1, 4294967295)
                            if false {
                                fmt.Println(v5, v2, v4)
                            } else {
                                if false {
                                    var v102 uint64 = 863
                                    for fi103 := 0; fi103 < 7; fi103++ {
                                        if (219 == 191) {
                                            if (!(1e10 <= (1.5 - 1e10))) {
                                                var v104 float64 = 0.0
                                                var v105 bool = false
                                                switch fi19 {
                                                case -2147483648:
                                                    var v106 uint64 = 0
                                                    _ = v106
                                                case -513:
                                                    var v107 bool = (32767 >= 381)
                                                    _ = v107
                                                }
                                                v108 := GenS7{F0: (v102 + 478), F1: v4}
                                                var v109 bool = (805 != 1)
                                                v102 = (v102 * v102)
                                                _ = v104
                                                _ = v105
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                        _ = fi103
                                    }
                                    for fi110 := 0; fi110 < 16; fi110++ {
                                        var v111 uint = 65535
                                        if ((("a" + ("a" + (("foo bar" + ("" + "\n\t")) + ("" + ("" + ("" + "hello")))))) + "\n\t") > "\n\t") {
                                            var v112 int32 = 2147483647
                                            v102 -= v102
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                switch (("" + "a") + ("foo bar" + (("a" + "hello") + "a"))) {
                                                case "\n\t":
                                                    if (0 == 1) {
                                                        v113 := GenS11{F0: 127, F1: (v5 * (v5 - (v5 | 0)))}
                                                        var v114 int = fi110
                                                        _ = v113
                                                        _ = v114
                                                    } else {
                                                        switch int(v102) {
                                                        case -9223372036854775808:
                                                            switch -1 {
                                                            case 9223372036854775807:
                                                                v2 += 9223372036854775807
                                                            default:
                                                                var v115 uint16 = 32767
                                                                _ = v115
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    v102 += (v102 >> ((v111) & 63))
                                                }
                                            case "hello":
                                                v116 := GenS3{F0: v112}
                                                _ = v116
                                            case "foo bar":
                                                fmt.Println(fi19, v102)
                                            default:
                                                fmt.Println(v4, fi19, v102)
                                            }
                                            _ = v112
                                        }
                                        _ = fi110
                                        _ = v111
                                    }
                                    _ = v102
                                }
                                v117 := GenS12{F0: uint8(v5)}
                                _ = v117
                            }
                        } else {
                            switch v2 {
                            case 128:
                                for fi118 := 0; fi118 < 13; fi118++ {
                                    var v119 uint = 93
                                    switch fi19 {
                                    case -9223372036854775808:
                                        v4 = 127
                                    case 128:
                                        if true {
                                            switch "hello" {
                                            case "foo bar":
                                                v119 = 0
                                            default:
                                                v120 := GenS4{F0: v5, F1: 9223372036854775807}
                                                _ = v120
                                            }
                                        } else {
                                            fmt.Println(v4, fi19, v2)
                                            v119 += 42
                                        }
                                    }
                                    _ = fi118
                                    _ = v119
                                }
                            case 127:
                                v121 := GenS13{F0: (true && (42 > 1))}
                                _ = v121
                            }
                            v4 = v4
                            fmt.Println(v2)
                        }
                    }
                }
            default:
                if false {
                    for fi122 := 0; fi122 < 19; fi122++ {
                        v5 = (v5 * v5)
                        _ = fi122
                    }
                    for fi123 := 0; fi123 < 3; fi123++ {
                        var v124 uint = 1
                        _ = fi123
                        _ = v124
                    }
                } else {
                    if (v4 <= -1) {
                        if true {
                            fmt.Println(fi19, v4)
                        }
                        v2 = (v2 | v2)
                    } else {
                        for fi125 := 0; fi125 < 8; fi125++ {
                            if true {
                                v5 += v5
                            }
                            _ = fi125
                        }
                    }
                    v5 += (^(v5))
                }
            }
            _ = v21
        }
        for fi126 := 0; fi126 < 17; fi126++ {
            if (!(18446744073709551615 != uint64(v5))) {
                delete(v1, (v5 / ((2) | 1)))
                fmt.Println(fi126, v4, v5)
            } else {
                v3["a"] = false
            }
            if true {
                switch "a" {
                case "":
                    v2 -= 32767
                case "\n\t":
                    fmt.Println(v2, fi19, fi126)
                default:
                    v127 := GenS5{F0: (^(v4)), F1: v5, F2: "hello"}
                    switch 128 {
                    case 30:
                        v128 := GenS14{F0: v2}
                        _ = v128
                    case -1:
                        if false {
                            var v129 int = (v2 ^ 4294967295)
                            _ = v129
                        }
                    case 9223372036854775807:
                        v2 = (v2 & (v2 + (fi19 ^ (fi19 << ((42) & 63)))))
                    default:
                        if (!true) {
                            v130 := GenS13{F0: (42 != 0)}
                            if (1 <= 42) {
                                v4 = (-(v4))
                                var v131 float64 = 0.1
                                _ = v131
                            } else {
                                v4 = -128
                                v2 = (^(fi126))
                                var v132 uint8 = 2
                                fmt.Println(fi126, v132)
                                _ = v132
                            }
                            _ = v130
                        } else {
                            v2 += (^(v2))
                            if true {
                                var v133 uint = 2
                                fmt.Println(v5)
                                var v134 float64 = float64(((((((((3.14159 * (-2.25 - float32(-2.25))) * 1e10) * float32(float64(0.1))) - (1.5 / ((3.14159 / 0.1) - ((0.1 / 1.5) + ((-2.25 - (float32(-2.25) - 1.5)) + (float32(-2.25) + 0.0)))))) / (1.5 + 3.14159)) - -2.25) / 1.0) / 1.0))
                                v135 := GenS15{F0: 676, F1: (65535 >= uint32(v133))}
                                var v136 uint = v133
                                _ = v133
                                _ = v134
                                _ = v135
                                _ = v136
                            } else {
                                v137 := GenS1{F0: 4294967295}
                                _ = v137
                            }
                            v6 = append(v6, v5)
                        }
                        v1[(v5 << ((708) & 63))] = (fi19 % (((fi126 >> ((525) & 63))) | 1))
                    }
                    v2 = -621
                    _ = v127
                }
                switch (^(fi19)) {
                case 756:
                    v4 = v4
                }
                var v138 int32 = 92
                _ = v138
            } else {
                var v139 uint64 = 18446744073709551615
                v4 = int8(v5)
                v2 += (-(fi126))
                _ = v139
            }
            switch (fi126 & 128) {
            case 4294967295:
                var v140 bool = (0 != -1)
                _ = v140
            default:
                var v141 uint32 = 2
                var v142 int = -633
                _ = v141
                _ = v142
            }
            _ = fi126
        }
        _ = fi19
    }
    v4 = 83
    var v143 uint64 = 4294967295
    v144 := GenS16{F0: (-(v2)), F1: 2147483647}
    _ = len(v6)
    v145 := GenS1{F0: v143}
    switch "hello" {
    case "a":
        var v146 uint32 = 597
        _ = v146
    case "\n\t":
        fmt.Println(v2, v4)
    default:
        switch "hello" {
        case "a":
            if false {
                var v147 int = (v2 | (v2 ^ v2))
                for fi148 := 0; fi148 < 17; fi148++ {
                    switch "\n\t" {
                    case "\n\t":
                        v5 -= (v5 + 2147483648)
                    default:
                        v147 += v2
                    }
                    _ = fi148
                }
                if (!false) {
                    _ = len(v6)
                    var v149 uint = 255
                    _ = v149
                } else {
                    var v150 float32 = 0.0
                    _ = v150
                }
                _ = v147
            } else {
                v5 -= (^(v5))
            }
        default:
            v4 += 1
        }
    }
    var v151 uint32 = 860
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v143:", v143)
    fmt.Println("v151:", v151)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v143
    _ = v144
    _ = v145
    _ = v151
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
