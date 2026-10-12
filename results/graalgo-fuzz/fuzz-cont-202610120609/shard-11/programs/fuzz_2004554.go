package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 bool; F2 int }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int; F1 uint8; F2 int }
type GenS13 struct { F0 bool; F1 uint8; F2 bool }
type GenS14 struct { F0 bool; F1 int16; F2 uint16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint32, p1 int) uint32 {
    p1 = (-(p1))
    p1 = p1
    v7 := GenS1{F0: "hello"}
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 string = ("\n\t" + ((("a" + ("foo bar" + "")) + "hello") + (((("\n\t" + "hello") + (("hello" + "\n\t") + "")) + "\n\t") + "foo bar")))
        _ = fi8
        _ = v9
    }
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    var v6 int16 = int16(v1)
    var v11 int64 = 701
    switch ("foo bar" + "") {
    case "a":
        var v12 int8 = 0
        _ = v12
    case "\n\t":
        v1 = v1
    case "":
        v6 -= 42
    default:
        v11 = (v11 >> ((48) & 63))
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if false {
            var v14 uint64 = 42
            _ = v14
        }
        v15 := GenS2{F0: 42, F1: -9223372036854775808}
        if false {
            switch "foo bar" {
            case "":
                switch fi13 {
                case 127:
                    if false {
                        if (0 <= 2147483647) {
                            v1 = 0
                            for fi16 := 0; fi16 < 14; fi16++ {
                                v6 = -729
                                _ = fi16
                            }
                            fmt.Println(v6)
                        } else {
                            _ = len(v3)
                            delete(v3, 42)
                            v17 := []int{fi13, (fi13 & fi13), fi13, 127}
                            _ = v17
                        }
                        if (42 != 281) {
                            delete(v3, 42)
                            v18 := []int{-1, (fi13 & fi13)}
                            v19 := GenS3{F0: 2, F1: v6, F2: v11}
                            _ = v18
                            _ = v19
                        }
                    }
                }
            }
            v20 := GenS4{F0: fi13}
            fmt.Println(v11, fi13)
            _ = v20
        }
        _ = fi13
        _ = v15
    }
    var v21 float64 = (1e10 / (((0.1 / (1e10 + 0.0)) + (1e10 + 1.5)) * (1.5 / 1.0)))
    v22 := GenS3{F0: genFn0(4294967295, 2147483647), F1: (v6 % ((v6) | 1)), F2: 9223372036854775807}
    v23 := GenS2{F0: 0, F1: 948}
    fmt.Println(v1, v1)
    v24 := GenS4{F0: 32767}
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS5{F0: v1}
        if false {
            v5[(("" + "a") + "\n\t")] = (0 > 2)
        } else {
            for fi27 := 0; fi27 < 18; fi27++ {
                if true {
                    if ((!true) && ((v6 >> ((2) & 63)) <= (v6 % ((v6) | 1)))) {
                        v21 += ((0.0 * float64((v2 * v2))) + (((v21 + 0.1) + float64(((1e10 - (v2 / v2)) - v2))) / v21))
                        if true {
                            var v28 int64 = v11
                            _ = v28
                        }
                        var v29 uint32 = 2147483647
                        _ = v29
                    }
                }
                v30 := GenS4{F0: 0}
                var v31 uint64 = 4294967295
                var v32 int64 = (v11 + v11)
                v33 := []int{fi25, (fi25 % (((fi27 ^ (fi25 - fi27))) | 1))}
                _ = fi27
                _ = v30
                _ = v31
                _ = v32
                _ = v33
            }
            if false {
                for fi34 := 0; fi34 < 5; fi34++ {
                    fmt.Println(v11, v6, v21)
                    var v35 bool = false
                    _ = fi34
                    _ = v35
                }
                v1 += (v1 % (((v1 ^ (^(v1)))) | 1))
                if (1.5 != v2) {
                    if (0 > -761) {
                        v36 := []int{fi25}
                        _ = v36
                    } else {
                        fmt.Println(fi25)
                        var v37 uint32 = 2147483647
                        switch "" {
                        case "hello":
                            switch int(v11) {
                            case -1:
                                for fi38 := 0; fi38 < 20; fi38++ {
                                    switch (^(fi38)) {
                                    case 128:
                                        if ((^(v37)) != (v37 + v37)) {
                                            var v39 string = "foo bar"
                                            v2 = (v2 * ((v2 / v2) * v2))
                                            v40 := GenS2{F0: (v37 - (v37 + v37)), F1: fi38}
                                            switch ("hello" + (("a" + v39) + (v39 + (v39 + v39)))) {
                                            case "hello":
                                                switch -188 {
                                                case 9223372036854775807:
                                                    v21 = ((v21 * v21) + 0.1)
                                                case -9223372036854775808:
                                                    if false {
                                                        v41 := GenS2{F0: (^(v37)), F1: (fi38 - fi38)}
                                                        if (!true) {
                                                            var v42 int8 = 1
                                                            v43 := []int{}
                                                            if false {
                                                                v6 = (v6 * v6)
                                                                v1 = 828
                                                            } else {
                                                                var v44 string = "foo bar"
                                                                v37 -= genFn0(943, -9223372036854775808)
                                                                switch ((v44 + v39) + v39) {
                                                                case "foo bar":
                                                                    fmt.Println(fi38, v39, v42)
                                                                }
                                                                fmt.Println(v1, v42)
                                                                _ = v44
                                                            }
                                                            _ = len(v43)
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        v45 := GenS6{F0: ((v39 > (v39 + v39)) || (4294967295 == 255)), F1: false, F2: fi38}
                                                        _ = v41
                                                        _ = v45
                                                    }
                                                default:
                                                    var v46 float64 = v21
                                                    _ = v46
                                                }
                                            case "\n\t":
                                                fmt.Println(v1)
                                            default:
                                                var v47 string = ("foo bar" + ("foo bar" + "foo bar"))
                                                _ = v47
                                            }
                                            fmt.Println(fi25, v6, v1)
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v48 int32 = -1
                                            fmt.Println(fi25)
                                            _ = v48
                                        }
                                    case 604:
                                        var v49 int = (^(fi38))
                                        _ = v49
                                    case 32768:
                                        var v50 float32 = v2
                                        _ = v50
                                    default:
                                        v37 -= v37
                                    }
                                    _ = fi38
                                }
                            case -9223372036854775808:
                                v11 += int64(fi25)
                            case 42:
                                v1 += v1
                            default:
                                if true {
                                    v51 := GenS3{F0: (v37 - v37), F1: v6, F2: (v11 & -127)}
                                    _ = v51
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        if (!false) {
                                            fmt.Println(fi25, v6, v2)
                                            switch "foo bar" {
                                            case "hello":
                                                v1 = v1
                                            }
                                            if false {
                                                v52 := GenS7{F0: fi25, F1: v1}
                                                fmt.Println(v1)
                                                _ = v52
                                            }
                                        }
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v2 = ((v2 / (0.1 + v2)) - 0.0)
                                    default:
                                        switch (fi25 * (fi25 | (^(fi25)))) {
                                        case 9223372036854775807:
                                            var v53 int16 = (v6 / (((v6 << ((4294967295) & 63))) | 1))
                                            _ = v53
                                        default:
                                            if (!false) {
                                                fmt.Println(v37, v37, v2)
                                                var v54 string = "\n\t"
                                                switch ((v54 + v54) + "hello") {
                                                case "\n\t":
                                                    fmt.Println(v54, v54, v6)
                                                case "a":
                                                    v55 := []int{}
                                                    _ = v55
                                                case "hello":
                                                    fmt.Println(v37)
                                                }
                                                _ = v54
                                            }
                                        }
                                    }
                                }
                            }
                        case "a":
                            if false {
                                if (255 < 286) {
                                    switch "foo bar" {
                                    case "hello":
                                        var v56 uint = 318
                                        _ = v56
                                    default:
                                        var v57 int64 = (^(v11))
                                        _ = v57
                                    }
                                    fmt.Println(v21)
                                    var v58 uint8 = 42
                                    _ = v58
                                }
                            } else {
                                v21 -= v21
                                if (false && true) {
                                    v37 -= v37
                                    var v59 float64 = v21
                                    _ = v59
                                }
                                if ((4294967295 <= 4294967295) || (2147483647 == 1)) {
                                    v60 := []int{(fi25 - (fi25 + fi25)), (fi25 ^ (fi25 ^ 9223372036854775807)), fi25, fi25}
                                    if false {
                                        var v61 uint32 = (v37 * v37)
                                        _ = v61
                                    }
                                    var v62 uint32 = (v37 - 4294967295)
                                    _ = v60
                                    _ = v62
                                }
                            }
                        case "":
                            v21 -= v21
                        default:
                            switch fi25 {
                            case 986:
                                var v63 int8 = -1
                                _ = v63
                            case 42:
                                v64 := []int{}
                                _ = v64
                            case -2147483648:
                                if (!("\n\t" <= "\n\t")) {
                                    v65 := []int{(^(fi25)), (fi25 >> ((2) & 63)), 4294967295}
                                    _ = v65
                                } else {
                                    v66 := GenS2{F0: genFn0(2, fi25), F1: -216}
                                    if ((v11 * v11) > v11) {
                                        var v67 int = fi25
                                        var v68 uint64 = uint64(v1)
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = v66
                                }
                            default:
                                v5["a"] = true
                            }
                        }
                        _ = v37
                    }
                } else {
                    v6 = (v6 / (((v6 - v6)) | 1))
                    var v69 uint64 = 42
                    _ = v69
                }
            }
        }
        v70 := GenS6{F0: (42 == 42), F1: (-17 <= 77), F2: fi25}
        _ = fi25
        _ = v26
        _ = v70
    }
    var v71 uint32 = 973
    for fi72 := 0; fi72 < 3; fi72++ {
        v73 := GenS3{F0: (^(v71)), F1: (-(v6)), F2: 0}
        if (230 >= 2) {
            if (127 >= int8(v6)) {
                for fi74 := 0; fi74 < 19; fi74++ {
                    v75 := GenS8{F0: 42, F1: -807}
                    v6 = (v6 % ((1) | 1))
                    v6 += -461
                    _ = fi74
                    _ = v75
                }
                fmt.Println(fi72, v11, v11)
                var v76 int32 = 1
                _ = v76
            } else {
                for fi77 := 0; fi77 < 12; fi77++ {
                    v21 += v21
                    var v78 uint = 2
                    _ = fi77
                    _ = v78
                }
            }
            v79 := []int{}
            v80 := GenS5{F0: 255}
            _ = v79
            _ = v80
        } else {
            if false {
                switch int(v6) {
                case -2147483648:
                    switch "hello" {
                    case "\n\t":
                        v1 = (v1 + (v1 + v1))
                    case "hello":
                        v81 := GenS6{F0: true, F1: false, F2: (fi72 << ((18446744073709551615) & 63))}
                        _ = v81
                    default:
                        var v82 int64 = (^(v11))
                        _ = v82
                    }
                default:
                    v83 := []int{}
                    _ = v83
                }
            }
            fmt.Println(v71, fi72)
            switch (^(fi72)) {
            case 974:
                switch (fi72 | fi72) {
                case 32768:
                    v84 := GenS9{F0: fi72}
                    _ = v84
                case 2147483647:
                    var v85 uint16 = (v1 - v1)
                    _ = v85
                default:
                    switch int(v1) {
                    case 4294967295:
                        if (2147483647 == -1) {
                            if (!("hello" <= "")) {
                                fmt.Println(v21, v11, v71)
                                switch "a" {
                                case "a":
                                    if (genFn0((^(v71)), (fi72 * (fi72 & 42))) <= 749) {
                                        fmt.Println(v1, v1)
                                        v86 := GenS10{F0: genFn0((v71 % (((^(v71))) | 1)), -9223372036854775808), F1: v6}
                                        _ = v86
                                    }
                                case "foo bar":
                                    if false {
                                        for fi87 := 0; fi87 < 2; fi87++ {
                                            if (true || true) {
                                                v11 -= -1
                                                var v88 string = (("foo bar" + "") + ("\n\t" + "\n\t"))
                                                _ = v88
                                            }
                                            if true {
                                                var v89 int8 = 42
                                                var v90 float64 = v21
                                                _ = v89
                                                _ = v90
                                            }
                                            switch ("hello" + "") {
                                            case "a":
                                                fmt.Println(v71, v2)
                                            case "":
                                                fmt.Println(v21)
                                            case "hello":
                                                fmt.Println(v2)
                                            }
                                            _ = fi87
                                        }
                                        var v91 uint8 = 1
                                        _ = v91
                                    } else {
                                        delete(v4, -891)
                                        var v92 uint8 = 128
                                        v92 = (v92 << ((65535) & 63))
                                        _ = v92
                                    }
                                default:
                                    if false {
                                        v1 = (v1 << ((42) & 63))
                                        switch (fi72 >> ((836) & 63)) {
                                        case 32768:
                                            for fi93 := 0; fi93 < 7; fi93++ {
                                                if (false && false) {
                                                    if false {
                                                        v94 := []int{-9223372036854775808, fi93}
                                                        var v95 int32 = 1
                                                        v96 := GenS1{F0: "a"}
                                                        _ = v94
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    switch 127 {
                                                    case 4294967295:
                                                        v1 += v1
                                                    }
                                                } else {
                                                    var v97 string = "hello"
                                                    v6 = 278
                                                    _ = v97
                                                }
                                                _ = fi93
                                            }
                                        case 127:
                                            fmt.Println(v6, v21)
                                        default:
                                            for fi98 := 0; fi98 < 2; fi98++ {
                                                _ = len(v3)
                                                v1 = (v1 ^ 0)
                                                _ = fi98
                                            }
                                        }
                                        for fi99 := 0; fi99 < 15; fi99++ {
                                            if (false || false) {
                                                var v100 uint32 = v71
                                                _ = v100
                                            } else {
                                                if true {
                                                    if true {
                                                        v6 = v6
                                                        var v101 uint8 = 255
                                                        v21 = v21
                                                        _ = v101
                                                    }
                                                } else {
                                                    v102 := GenS11{F0: 42}
                                                    fmt.Println(fi72, v21)
                                                    v1 = v1
                                                    _ = v102
                                                }
                                                v103 := GenS12{F0: (fi72 | fi72), F1: 1, F2: 9223372036854775807}
                                                _ = v103
                                            }
                                            v71 -= (v71 >> ((42) & 63))
                                            v21 -= v21
                                            if (!false) {
                                                if false {
                                                    switch (("a" + ("foo bar" + "\n\t")) + ("a" + "foo bar")) {
                                                    case "hello":
                                                        v5["a"] = (!false)
                                                        if (!((true || (true && (-2.25 < v2))) || (true || true))) {
                                                            switch (fi99 << ((389) & 63)) {
                                                            case -405:
                                                                _ = len(v3)
                                                            }
                                                            if (1 >= 18446744073709551615) {
                                                                fmt.Println(v6, v2, v6)
                                                                v2 += 0.0
                                                                var v104 int = -2147483648
                                                                _ = v104
                                                            } else {
                                                                v105 := GenS11{F0: 762}
                                                                _ = v105
                                                            }
                                                        }
                                                        var v106 uint8 = uint8(v71)
                                                        _ = v106
                                                    default:
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v107 int64 = (-(v11))
                                                            _ = v107
                                                        case "a":
                                                            if (!((995 <= -1) || ((((1e10 + (v2 + (3.14159 * 3.14159))) - (v2 - v2)) / (v2 - 0.0)) != v2))) {
                                                                v108 := GenS13{F0: (102 > -61), F1: 127, F2: ((true || ((v1 ^ (^(v1))) < v1)) || (1 > 0))}
                                                                _ = v108
                                                            } else {
                                                                if (!true) {
                                                                    fmt.Println(v71, v1, v6)
                                                                    v11 = v11
                                                                    v109 := GenS14{F0: true, F1: 42, F2: v1}
                                                                    v4[fi99] = "foo bar"
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v110 float32 = 3.14159
                                                                        _ = v110
                                                                    }
                                                                    _ = v109
                                                                }
                                                                v11 += 924
                                                            }
                                                            if true {
                                                                if (((true && (!true)) || false) || (950 == 32767)) {
                                                                    if (0 < uint16(v71)) {
                                                                        var v111 int8 = 100
                                                                        v3[v71] = true
                                                                        var v112 uint = 4294967295
                                                                        _ = v111
                                                                        _ = v112
                                                                    }
                                                                    v2 = (v2 - (float32(v21) - float32((v21 * v21))))
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v6, v21)
                                                                        v113 := []int{}
                                                                        _ = v113
                                                                    } else {
                                                                        if (!true) {
                                                                            var v114 float64 = v21
                                                                            var v115 int8 = -1
                                                                            _ = v114
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                    var v116 uint64 = 44
                                                                    _ = v116
                                                                }
                                                                if (v6 <= v6) {
                                                                    if false {
                                                                        v21 -= v21
                                                                    }
                                                                } else {
                                                                    v117 := GenS4{F0: fi99}
                                                                    v6 += v6
                                                                    _ = v117
                                                                }
                                                            } else {
                                                                v11 = v11
                                                            }
                                                        default:
                                                            v118 := []int{}
                                                            v119 := GenS8{F0: 1, F1: (v6 ^ v6)}
                                                            switch (fi72 + fi99) {
                                                            case 9223372036854775807:
                                                                var v120 string = "hello"
                                                                _ = v120
                                                            case -9223372036854775808:
                                                                v6 = 42
                                                            case 4294967295:
                                                                v2 += (((float32(v21) + float32(((float64((v2 + v2)) * v21) - v21))) / (v2 + ((-2.25 + float32(0.1)) + (v2 - (-2.25 / v2))))) + v2)
                                                            }
                                                            if ((v21 + 0.1) <= float64((0.0 * v2))) {
                                                                if true {
                                                                    if len(v118) > 0 {
                                                                        v118[0] = (fi99 & 32767)
                                                                    }
                                                                    v121 := GenS15{F0: ("" + ("hello" + "hello"))}
                                                                    _ = v121
                                                                }
                                                                var v122 float64 = v21
                                                                _ = v122
                                                            } else {
                                                                var v123 int64 = -756
                                                                v2 += 1.5
                                                                _ = v123
                                                            }
                                                            _ = v118
                                                            _ = v119
                                                        }
                                                    }
                                                    v21 += 1e10
                                                    v124 := []int{}
                                                    _ = v124
                                                } else {
                                                    v125 := []int{42}
                                                    switch -2147483648 {
                                                    case 2147483647:
                                                        v11 += v11
                                                    }
                                                    switch (fi72 - (fi99 << ((255) & 63))) {
                                                    case -903:
                                                        if len(v125) >= 2 {
                                                            v125 = v125[1:len(v125)]
                                                        }
                                                    case 32768:
                                                        _ = len(v4)
                                                    default:
                                                        if false {
                                                            switch (fi72 % ((fi99) | 1)) {
                                                            case 32767:
                                                                var v126 uint = 2
                                                                _ = v126
                                                            case 128:
                                                                v11 += v11
                                                            case 4294967295:
                                                                v127 := GenS7{F0: (fi99 ^ 9223372036854775807), F1: v1}
                                                                _ = v127
                                                            }
                                                            v128 := GenS6{F0: false, F1: false, F2: (fi99 & (^(fi99)))}
                                                            _ = v128
                                                        }
                                                    }
                                                    if (-1 >= int(v1)) {
                                                        if false {
                                                            v129 := GenS15{F0: "foo bar"}
                                                            v21 = v21
                                                            var v130 uint64 = 18446744073709551615
                                                            _ = v129
                                                            _ = v130
                                                        } else {
                                                            v6 = (v6 << ((42) & 63))
                                                            v21 += v21
                                                            if (138 != 628) {
                                                                v131 := GenS16{F0: 127}
                                                                _ = v131
                                                            }
                                                        }
                                                        var v132 int16 = 42
                                                        _ = v132
                                                    }
                                                    _ = v125
                                                }
                                            } else {
                                                fmt.Println(fi72, fi72)
                                                fmt.Println(v1)
                                                var v133 float32 = v2
                                                switch int(v11) {
                                                case 4294967295:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v134 uint32 = v71
                                                        _ = v134
                                                    default:
                                                        v135 := GenS17{F0: 42, F1: 166}
                                                        _ = v135
                                                    }
                                                case 32767:
                                                    v21 -= 0.0
                                                case 32768:
                                                    fmt.Println(v1)
                                                }
                                                _ = v133
                                            }
                                            _ = fi99
                                        }
                                    } else {
                                        var v136 uint16 = v1
                                        var v137 int16 = (v6 - int16(v136))
                                        _ = v136
                                        _ = v137
                                    }
                                }
                                var v138 int32 = 1
                                _ = v138
                            }
                        } else {
                            for fi139 := 0; fi139 < 12; fi139++ {
                                v6 += (v6 | (v6 << ((42) & 63)))
                                _ = fi139
                            }
                            v140 := v4[-2147483648]
                            var v141 int = fi72
                            _ = v140
                            _ = v141
                        }
                    case 32768:
                        var v142 float64 = v21
                        _ = v142
                    case -9223372036854775808:
                        v11 = (v11 | v11)
                    }
                }
            case -2147483648:
                switch "foo bar" {
                case "foo bar":
                    _ = len(v5)
                case "\n\t":
                    switch ("foo bar" + ("foo bar" + ("" + ""))) {
                    case "\n\t":
                        var v143 int32 = 514
                        _ = v143
                    default:
                        v144 := []int{32768, fi72, (fi72 % ((fi72) | 1)), 4294967295}
                        _ = v144
                    }
                case "a":
                    var v145 uint32 = 0
                    _ = v145
                }
            default:
                var v146 int16 = v6
                _ = v146
            }
        }
        _ = fi72
        _ = v73
    }
    for fi147 := 0; fi147 < 4; fi147++ {
        v148 := []int{-783}
        _ = fi147
        _ = v148
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
