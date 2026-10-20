package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int; F2 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 bool; F1 string; F2 int64 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]uint32)
    var v3 uint = 2
    var v4 int8 = 100
    v5 := v2["hello"]
    v4 -= 97
    var v6 float32 = v1
    if false {
        var v7 float64 = (0.0 - (0.0 + 0.0))
        v8 := GenS1{F0: (v5 & (v5 ^ 1))}
        var v9 int32 = 42
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{-2147483648, (fi10 >> ((255) & 63)), fi10}
            _ = fi10
            _ = v11
        }
        v2[(((("a" + "\n\t") + (("a" + ("hello" + "hello")) + ("foo bar" + ""))) + "a") + "\n\t")] = (v5 >> ((v3) & 63))
        for fi12 := 0; fi12 < 18; fi12++ {
            v2["\n\t"] = v5
            _ = fi12
        }
    }
    v13 := GenS1{F0: v5}
    v14 := GenS1{F0: v5}
    if (!(v4 <= -1)) {
        var v15 int8 = v4
        _ = v15
    }
    if false {
        if false {
            if true {
                var v16 bool = true
                switch "a" {
                case "a":
                    if (((-2.25 + v6) - 0.0) > v6) {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v18 := []int{fi17, fi17, (fi17 * (fi17 | -9223372036854775808))}
                            var v19 int = fi17
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                case "foo bar":
                    v6 -= 0.0
                case "":
                    if (!(v16 && v16)) {
                        if (!v16) {
                            for fi20 := 0; fi20 < 11; fi20++ {
                                var v21 int = fi20
                                if ("" == "hello") {
                                    v22 := []int{}
                                    switch "hello" {
                                    case "\n\t":
                                        for fi23 := 0; fi23 < 16; fi23++ {
                                            v22 = append(v22, (fi23 / ((fi23) | 1)))
                                            var v24 int = (v21 ^ fi20)
                                            if true {
                                                v3 = v3
                                                v22 = append(v22, (fi23 ^ fi20))
                                            }
                                            _ = fi23
                                            _ = v24
                                        }
                                    case "hello":
                                        fmt.Println(fi20)
                                    case "a":
                                        switch "" {
                                        case "":
                                            var v25 int = 32768
                                            _ = v25
                                        case "\n\t":
                                            var v26 float64 = (0.0 - ((float64(float32(0.1)) / 1.5) - -2.25))
                                            _ = v26
                                        }
                                    default:
                                        v16 = v16
                                    }
                                    _ = v22
                                } else {
                                    v27 := GenS1{F0: 737}
                                    _ = v27
                                }
                                v28 := GenS2{F0: "\n\t", F1: (42 == (v5 % (((v5 * v5)) | 1)))}
                                _ = fi20
                                _ = v21
                                _ = v28
                            }
                            v6 -= (v1 + ((v6 / v1) - (((1e10 * 0.0) * float32(((-2.25 - ((0.0 - (3.14159 + ((-2.25 * (0.1 + -2.25)) + float64(0.1)))) - 0.1)) - 0.1))) - ((0.0 - ((v1 * v1) + (v6 / v1))) - (v6 - (-2.25 + v6))))))
                            v3 = v3
                        } else {
                            fmt.Println(v5)
                        }
                    }
                default:
                    switch 0 {
                    case 32768:
                        var v29 uint16 = 42
                        _ = v29
                    }
                }
                if true {
                    var v30 bool = (!(true || (42 > 140)))
                    v1 += (v1 + 0.1)
                    var v31 uint8 = 255
                    _ = v30
                    _ = v31
                }
                _ = v16
            }
            var v32 int32 = 42
            if ((v32 >> (((v3 | (v3 % ((v3) | 1)))) & 63)) == v32) {
                fmt.Println(v3, v5)
                v6 += (v6 / (float32(3.14159) * v6))
                if (v32 == (v32 - v32)) {
                    var v33 int64 = 660
                    var v34 float64 = ((1.5 / 1.0) - (1.5 + 1.5))
                    _ = v33
                    _ = v34
                } else {
                    switch ((("hello" + "\n\t") + ("" + ("" + (("\n\t" + ("hello" + "hello")) + (("a" + "foo bar") + "hello"))))) + "") {
                    case "":
                        v3 += (^(v3))
                    }
                    var v35 uint = v3
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v5 = 380
                        _ = fi36
                    }
                    _ = v35
                }
            }
            _ = v32
        } else {
            fmt.Println(v1, v4, v1)
            if ((!(9223372036854775807 >= 42)) || true) {
                v6 = -2.25
                if ((3.14159 - ((1e10 + -2.25) + float64(v1))) < 0.1) {
                    fmt.Println(v6, v5)
                    v37 := GenS2{F0: "foo bar", F1: false}
                    v4 = 127
                    _ = v37
                } else {
                    if ("" >= "foo bar") {
                        if true {
                            if (true || true) {
                                v3 = (v3 ^ (v3 & (v3 - (v3 + 65535))))
                                fmt.Println(v4)
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    for fi38 := 0; fi38 < 7; fi38++ {
                                        if (((((((1.5 - 0.1) * 1e10) * 3.14159) + 3.14159) / float64(v6)) + float64(float32((0.0 * (1e10 * 3.14159))))) != 1e10) {
                                            delete(v2, (("hello" + (("foo bar" + "a") + "\n\t")) + ""))
                                            v39 := GenS3{F0: false, F1: "a"}
                                            for fi40 := 0; fi40 < 2; fi40++ {
                                                var v41 uint8 = 127
                                                v41 -= (^(v41))
                                                switch "" {
                                                case "\n\t":
                                                    if false {
                                                        if true {
                                                            if ((1.5 >= 0.0) || (0 >= 1)) {
                                                                v42 := []int{}
                                                                var v43 int16 = 1000
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                var v44 bool = (!false)
                                                                v44 = false
                                                                v4 -= v4
                                                                _ = v44
                                                            }
                                                            v3 -= v3
                                                            v5 -= 968
                                                        }
                                                        if false {
                                                            v45 := GenS1{F0: v5}
                                                            v46 := []int{fi40, (fi38 - -2147483648), (fi38 & 127)}
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        switch "" {
                                                        case "hello":
                                                            _ = len(v2)
                                                        }
                                                        v47 := v2[("a" + ("foo bar" + "\n\t"))]
                                                        v48 := GenS3{F0: false, F1: ""}
                                                        var v49 float64 = 1.5
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        v50 := GenS2{F0: "a", F1: true}
                                                        switch (^(fi38)) {
                                                        case -1:
                                                            fmt.Println(v5, v6, v4)
                                                        case 4294967295:
                                                            v51 := []int{fi38}
                                                            _ = v51
                                                        case 32767:
                                                            var v52 int32 = 1
                                                            _ = v52
                                                        default:
                                                            v5 += (^(v5))
                                                        }
                                                        var v53 uint64 = 843
                                                        _ = v50
                                                        _ = v53
                                                    }
                                                }
                                                _ = fi40
                                                _ = v41
                                            }
                                            _ = v39
                                        } else {
                                            v54 := GenS1{F0: uint32(v3)}
                                            if ("\n\t" <= "") {
                                                var v55 int8 = (v4 >> ((1) & 63))
                                                _ = v55
                                            } else {
                                                v5 += v5
                                                fmt.Println(v1, v4, v6)
                                                var v56 int64 = 42
                                                _ = v56
                                            }
                                            v1 = 1e10
                                            _ = v54
                                        }
                                        switch "a" {
                                        case "":
                                            switch "" {
                                            case "a":
                                                v57 := v2[("hello" + "hello")]
                                                _ = v57
                                            case "hello":
                                                for fi58 := 0; fi58 < 10; fi58++ {
                                                    delete(v2, "a")
                                                    _ = fi58
                                                }
                                            case "\n\t":
                                                v4 = 0
                                            default:
                                                v59 := GenS4{F0: (v3 * (v3 * (v3 - (v3 - (v3 | (v3 | (v3 + (v3 % (((^(v3))) | 1)))))))))}
                                                _ = v59
                                            }
                                        case "hello":
                                            switch 42 {
                                            case 1:
                                                if (!true) {
                                                    var v60 uint8 = 222
                                                    var v61 uint16 = 0
                                                    _ = v60
                                                    _ = v61
                                                }
                                            case -9223372036854775808:
                                                switch 0 {
                                                case -1:
                                                    var v62 float32 = 1.5
                                                    _ = v62
                                                case -957:
                                                    fmt.Println(fi38, v3)
                                                default:
                                                    var v63 float64 = (-2.25 + 1e10)
                                                    _ = v63
                                                }
                                            }
                                        case "a":
                                            for fi64 := 0; fi64 < 7; fi64++ {
                                                v1 = (1e10 / v1)
                                                v3 = (v3 + 0)
                                                _ = fi64
                                            }
                                        }
                                        _ = fi38
                                    }
                                default:
                                    switch (("hello" + "\n\t") + ("\n\t" + ("foo bar" + "hello"))) {
                                    case "":
                                        var v65 string = ("" + "a")
                                        _ = v65
                                    case "\n\t":
                                        switch ("\n\t" + ("" + ((("a" + ("foo bar" + "foo bar")) + "hello") + "foo bar"))) {
                                        case "a":
                                            v6 = -2.25
                                        case "\n\t":
                                            if true {
                                                var v66 uint64 = 0
                                                _ = v66
                                            } else {
                                                fmt.Println(v6)
                                                v5 = (v5 * 415)
                                                fmt.Println(v3, v3)
                                            }
                                        case "":
                                            var v67 int64 = -1
                                            _ = v67
                                        default:
                                            v68 := []int{42}
                                            _ = v68
                                        }
                                    case "foo bar":
                                        v6 = 0.1
                                    default:
                                        fmt.Println(v1, v5)
                                    }
                                }
                            }
                        }
                        for fi69 := 0; fi69 < 11; fi69++ {
                            if (!(1.5 < 1e10)) {
                                fmt.Println(fi69, v3)
                                fmt.Println(v1, v1, v4)
                            }
                            var v70 int64 = 930
                            if (!(true && false)) {
                                fmt.Println(v70)
                                v71 := GenS5{F0: ("foo bar" == ((("hello" + "\n\t") + "hello") + (("a" + (("" + "foo bar") + "")) + "\n\t")))}
                                _ = v71
                            } else {
                                v3 = v3
                                var v72 uint16 = 65535
                                fmt.Println(v72, v72)
                                _ = v72
                            }
                            _ = fi69
                            _ = v70
                        }
                    }
                    switch 9223372036854775807 {
                    case 32767:
                        var v73 float64 = 1e10
                        _ = v73
                    case -898:
                        for fi74 := 0; fi74 < 19; fi74++ {
                            v5 -= v5
                            fmt.Println(v3, v3)
                            v75 := []int{}
                            _ = fi74
                            _ = v75
                        }
                    }
                    v76 := v2["foo bar"]
                    _ = v76
                }
            }
        }
        v5 += 2147483648
    } else {
        var v77 int32 = 2147483647
        var v78 uint = (^(v3))
        v79 := GenS1{F0: (v5 - v5)}
        _ = v77
        _ = v78
        _ = v79
    }
    if (false && false) {
        v3 = (v3 / (((v3 >> ((v3) & 63))) | 1))
        v3 = (v3 % (((^(v3))) | 1))
    }
    v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
    if true {
        if true {
            var v80 int8 = v4
            _ = v80
        }
    } else {
        v81 := GenS2{F0: "a", F1: (2147483647 > 128)}
        for fi82 := 0; fi82 < 15; fi82++ {
            var v83 uint32 = (v5 & v5)
            fmt.Println(v3)
            _ = fi82
            _ = v83
        }
        v6 += ((v6 + -2.25) - 3.14159)
        _ = v81
    }
    if (true || (true || (2 >= 894))) {
        var v84 uint16 = 936
        if (42 <= -32768) {
            v85 := GenS2{F0: ("a" + "\n\t"), F1: true}
            v4 -= (v4 >> ((255) & 63))
            if (((3.14159 + float32(1e10)) == 1e10) && true) {
                switch 42 {
                case 32768:
                    if (127 > 255) {
                        fmt.Println(v6)
                    }
                case 128:
                    v86 := []int{2147483647}
                    _ = v86
                case 215:
                    delete(v2, "hello")
                }
                if false {
                    v87 := GenS1{F0: (v5 - (v5 >> (((v3 << (((^(v3))) & 63))) & 63)))}
                    delete(v2, ("\n\t" + "a"))
                    _ = v87
                } else {
                    var v88 uint64 = 4294967295
                    _ = v88
                }
                switch (("foo bar" + ("a" + "")) + ("\n\t" + "\n\t")) {
                case "":
                    if (((v3 >> ((4294967295) & 63)) == 65) || false) {
                        var v89 uint16 = uint16(v5)
                        for fi90 := 0; fi90 < 6; fi90++ {
                            if true {
                                var v91 bool = (false && (0 == -301))
                                var v92 string = ("foo bar" + "a")
                                fmt.Println(v5)
                                _ = v91
                                _ = v92
                            }
                            var v93 int32 = -2147483648
                            var v94 uint64 = 2
                            _ = fi90
                            _ = v93
                            _ = v94
                        }
                        _ = v89
                    }
                case "\n\t":
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        v84 = uint16(v3)
                    case "\n\t":
                        v95 := GenS2{F0: ((("\n\t" + "a") + "foo bar") + ((("hello" + "") + ((("foo bar" + ("hello" + (("a" + "") + "hello"))) + "foo bar") + "\n\t")) + "foo bar")), F1: true}
                        _ = v95
                    default:
                        for fi96 := 0; fi96 < 3; fi96++ {
                            var v97 bool = false
                            var v98 uint = (v3 << ((18446744073709551615) & 63))
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                    }
                default:
                    fmt.Println(v84, v5, v6)
                }
            }
            _ = v85
        } else {
            v1 += v1
            if ((!(!((!true) || false))) && (!true)) {
                if false {
                    for fi99 := 0; fi99 < 11; fi99++ {
                        v6 = v1
                        v100 := GenS4{F0: 902}
                        _ = fi99
                        _ = v100
                    }
                    var v101 int8 = (v4 ^ v4)
                    fmt.Println(v6, v6)
                    _ = v101
                } else {
                    var v102 uint32 = 1
                    var v103 int = 0
                    v104 := GenS2{F0: ("" + "\n\t"), F1: (-183 > -1)}
                    _ = v102
                    _ = v103
                    _ = v104
                }
                v4 += (v4 >> ((0) & 63))
                fmt.Println(v84, v1, v5)
            } else {
                fmt.Println(v6, v5)
                v105 := v2["foo bar"]
                v2["a"] = (v105 / ((2147483648) | 1))
                _ = v105
            }
        }
        _ = v84
    }
    if false {
        if true {
            var v106 uint64 = uint64(v3)
            _ = v106
        }
    }
    if (18446744073709551615 != 29) {
        var v107 uint8 = 51
        _ = v107
    }
    if (0.1 != float64(0.0)) {
        for fi108 := 0; fi108 < 10; fi108++ {
            for fi109 := 0; fi109 < 10; fi109++ {
                _ = len(v2)
                switch fi108 {
                case -1:
                    if false {
                        switch fi109 {
                        case -763:
                            var v110 uint64 = 1
                            _ = v110
                        default:
                            var v111 int32 = -2147483648
                            _ = v111
                        }
                        fmt.Println(fi109, v1)
                        fmt.Println(v3, fi109)
                    } else {
                        if (true || false) {
                            var v112 int8 = v4
                            switch 994 {
                            case 4294967295:
                                v113 := []int{913, (fi109 * fi108), 127}
                                _ = v113
                            case -104:
                                if true {
                                    var v114 float64 = (0.1 - (float64((v1 + float32(1.5))) + 3.14159))
                                    _ = v114
                                } else {
                                    switch (-(fi108)) {
                                    case 2147483647:
                                        switch "" {
                                        case "foo bar":
                                            var v115 int32 = -466
                                            _ = v115
                                        case "hello":
                                            var v116 int = (fi109 | 32768)
                                            _ = v116
                                        case "\n\t":
                                            fmt.Println(v6, fi108, v4)
                                        default:
                                            v117 := GenS6{F0: v4}
                                            _ = v117
                                        }
                                    case 32767:
                                        fmt.Println(v1, v1)
                                    default:
                                        var v118 int = (-(fi108))
                                        _ = v118
                                    }
                                    if (2 == 607) {
                                        fmt.Println(v5, v4)
                                    } else {
                                        v2["hello"] = 129
                                        if false {
                                            v112 = 1
                                            v119 := []int{}
                                            v4 = v112
                                            _ = v119
                                        }
                                    }
                                    fmt.Println(v4)
                                }
                            case -1:
                                fmt.Println(v112)
                            default:
                                v120 := []int{42, (fi109 >> ((2) & 63)), (^(fi108))}
                                _ = v120
                            }
                            var v121 string = ("a" + ((("foo bar" + ("a" + "foo bar")) + ("a" + "a")) + "hello"))
                            _ = v112
                            _ = v121
                        }
                        var v122 uint16 = 926
                        v123 := GenS7{F0: 0}
                        switch int(v122) {
                        case 429:
                            if (v3 < 1) {
                                v124 := []int{(fi109 * -9223372036854775808), -1}
                                switch fi108 {
                                case 0:
                                    v125 := GenS1{F0: v5}
                                    _ = v125
                                }
                                switch "\n\t" {
                                case "hello":
                                    v1 += (v1 / float32(((float64(v1) / 0.1) / (-2.25 - 1e10))))
                                case "foo bar":
                                    var v126 int32 = 0
                                    _ = v126
                                default:
                                    fmt.Println(fi108)
                                }
                                _ = v124
                            } else {
                                if true {
                                    v6 = 1.5
                                }
                                v127 := GenS3{F0: (-643 > 1000), F1: (("\n\t" + ((("hello" + ((("foo bar" + ("" + "hello")) + "hello") + "hello")) + "a") + "")) + ("foo bar" + ""))}
                                _ = v127
                            }
                        default:
                            delete(v2, "")
                        }
                        _ = v122
                        _ = v123
                    }
                case -234:
                    var v128 uint32 = (v5 >> (((v3 | (^(v3)))) & 63))
                    _ = v128
                case 406:
                    var v129 uint8 = 179
                    switch fi108 {
                    case 32767:
                        var v130 uint64 = uint64(v3)
                        _ = v130
                    default:
                        v4 -= (v4 | (^(v4)))
                    }
                    v131 := GenS5{F0: true}
                    _ = v129
                    _ = v131
                }
                fmt.Println(v4, v5, v6)
                if (!(255 >= 65535)) {
                    if false {
                        v132 := v2["a"]
                        var v133 int32 = 1
                        _ = v132
                        _ = v133
                    } else {
                        v134 := GenS2{F0: "", F1: (1 > 2147483647)}
                        v1 -= 0.1
                        switch (^(fi109)) {
                        case 0:
                            var v135 string = "a"
                            _ = v135
                        case 42:
                            if ((!false) && (-240 <= 1)) {
                                var v136 bool = false
                                v1 = v6
                                delete(v2, ("foo bar" + (("a" + "") + "a")))
                                _ = v136
                            }
                        default:
                            if (-2.25 == 0.0) {
                                v3 = v3
                                if true {
                                    var v137 int = (fi108 / (((fi109 + 127)) | 1))
                                    var v138 int32 = -284
                                    v3 += (v3 / ((809) | 1))
                                    if false {
                                        delete(v2, ("a" + "hello"))
                                    } else {
                                        if (1000 > 32767) {
                                            switch v137 {
                                            case 849:
                                                switch fi109 {
                                                case 380:
                                                    if true {
                                                        v139 := v2["hello"]
                                                        v2[""] = (v5 ^ 1)
                                                        v140 := []int{-75, fi108}
                                                        _ = v139
                                                        _ = v140
                                                    }
                                                case -901:
                                                    fmt.Println(v6)
                                                default:
                                                    if true {
                                                        v5 += v5
                                                        v141 := GenS8{F0: (false || (!((v138 ^ (v138 & v138)) <= v138))), F1: fi108, F2: false}
                                                        var v142 float32 = 3.14159
                                                        _ = v141
                                                        _ = v142
                                                    } else {
                                                        v4 += (v4 * 42)
                                                    }
                                                }
                                                if true {
                                                    v138 = (v138 ^ v138)
                                                }
                                                var v143 float64 = (3.14159 + ((float64(v6) * ((1e10 - 0.1) + ((float64((v6 + 3.14159)) * ((((float64(((float32(1e10) * float32((0.1 * 0.1))) + (float32(((float64((((((v6 - (v1 + (v6 - v6))) / (float32((-2.25 - ((0.1 + 1e10) * 3.14159))) + 0.1)) / 1.5) * float32((0.0 * -2.25))) / (float32((3.14159 - (float64(((0.1 / v1) - v6)) * 1.5))) / 1.5))) / -2.25) + (float64(v6) / -2.25))) * 1e10))) * (3.14159 * -2.25)) - 1e10) * -2.25) - 0.1)) - 1e10))) * 1e10))
                                                var v144 uint64 = 0
                                                v145 := GenS8{F0: (!true), F1: (fi109 ^ v137), F2: true}
                                                v138 = 42
                                                var v146 int8 = int8(v138)
                                                v4 += (v4 - v4)
                                                v147 := GenS4{F0: 4294967295}
                                                v148 := []int{fi108, (fi108 - fi109), fi109}
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                                _ = v146
                                                _ = v147
                                                _ = v148
                                            case -498:
                                                v149 := []int{(fi108 + (fi108 / ((fi108) | 1))), (v137 / (((fi109 - (fi108 ^ (fi108 ^ int(v138))))) | 1)), v137}
                                                _ = v149
                                            case 1:
                                                var v150 uint16 = 117
                                                if (1 >= 0) {
                                                    if true {
                                                        v150 += (^(v150))
                                                        if (false && (3.14159 > (((1.5 + (1e10 - (3.14159 + float64((v6 * float32(0.0)))))) * ((((float64(((v6 * v6) * 3.14159)) - 1e10) * -2.25) / (((float64((0.1 * (float32(0.1) / -2.25))) / 0.1) + ((1.5 + 1e10) - 1.5)) + float64((v6 / v1)))) / float64(v1))) + 3.14159))) {
                                                            var v151 uint16 = (v150 >> ((65535) & 63))
                                                            switch (v137 % ((v137) | 1)) {
                                                            case -1:
                                                                v3 -= 468
                                                            case 32767:
                                                                if (0 <= 1) {
                                                                    v138 = 0
                                                                    var v152 uint8 = 255
                                                                    v153 := GenS9{F0: 42}
                                                                    v3 = (v3 << (((v3 - v3)) & 63))
                                                                    _ = v152
                                                                    _ = v153
                                                                } else {
                                                                    v154 := GenS10{F0: v5, F1: v5}
                                                                    _ = v154
                                                                }
                                                            }
                                                            if true {
                                                                if false {
                                                                    v1 -= v1
                                                                    if true {
                                                                        v151 += 65535
                                                                        v155 := []int{1, -916, fi108, v137}
                                                                        var v156 int = fi109
                                                                        _ = v155
                                                                        _ = v156
                                                                    }
                                                                    v137 = 128
                                                                    var v157 int8 = 28
                                                                    fmt.Println(v1, v3)
                                                                    v151 = 0
                                                                    _ = v157
                                                                } else {
                                                                    v2[""] = (v5 | 4294967295)
                                                                    if (1 < (v4 * v4)) {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            if (1e10 > (3.14159 + 1.5)) {
                                                                                var v158 float32 = (1.5 / (1e10 / v6))
                                                                                v150 = v151
                                                                                _ = len(v2)
                                                                                _ = v158
                                                                            }
                                                                            switch "a" {
                                                                            case "":
                                                                                fmt.Println(v150, v6)
                                                                            case "a":
                                                                                if (v138 == v138) {
                                                                                    if false {
                                                                                        v6 -= (v1 * (1.5 * v6))
                                                                                    }
                                                                                    fmt.Println(v5, v6)
                                                                                    delete(v2, "")
                                                                                } else {
                                                                                    v151 = 654
                                                                                    switch (fi109 & (v137 - v137)) {
                                                                                    case 32767:
                                                                                        fmt.Println(fi108)
                                                                                    case 127:
                                                                                        v150 -= 2
                                                                                        v1 += 0.1
                                                                                        v159 := []int{fi109, int(v3), (-(fi109)), int(v151)}
                                                                                        _ = v159
                                                                                    case -566:
                                                                                        var v160 int16 = 198
                                                                                        _ = v160
                                                                                    default:
                                                                                        fmt.Println(v5, v1, v137)
                                                                                    }
                                                                                    if ((!false) || true) {
                                                                                        var v161 string = ""
                                                                                        _ = v161
                                                                                    } else {
                                                                                        v162 := GenS2{F0: "foo bar", F1: false}
                                                                                        v150 -= v150
                                                                                        _ = v162
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(fi108, v1, v138)
                                                                            }
                                                                        case "foo bar":
                                                                            var v163 int64 = -9223372036854775808
                                                                            _ = v163
                                                                        case "":
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    var v164 int64 = 1
                                                                                    _ = v164
                                                                                case "hello":
                                                                                    v150 -= v151
                                                                                case "a":
                                                                                    if true {
                                                                                        delete(v2, "hello")
                                                                                    }
                                                                                    var v165 int8 = 1
                                                                                    _ = v165
                                                                                default:
                                                                                    fmt.Println(fi109)
                                                                                    fmt.Println(v1, v1, v4)
                                                                                    v166 := []int{v137}
                                                                                    _ = v166
                                                                                }
                                                                            case "foo bar":
                                                                                v151 -= 391
                                                                            case "":
                                                                                var v167 uint64 = 2
                                                                                _ = v167
                                                                            default:
                                                                                v168 := GenS10{F0: (v5 % ((v5) | 1)), F1: 582}
                                                                                _ = v168
                                                                            }
                                                                        default:
                                                                            _ = len(v2)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v169 string = "hello"
                                                            var v170 uint8 = 2
                                                            v171 := GenS8{F0: (!(!true)), F1: (^(fi108)), F2: false}
                                                            v3 -= 4294967295
                                                            v170 -= v170
                                                            switch (fi108 % (((fi109 + 42)) | 1)) {
                                                            case 2147483647:
                                                                var v172 bool = (890 != 0)
                                                                _ = v172
                                                            }
                                                            _ = v151
                                                            _ = v169
                                                            _ = v170
                                                            _ = v171
                                                        } else {
                                                            if true {
                                                                v3 -= (v3 % (((v3 << ((v3) & 63))) | 1))
                                                                var v173 string = "foo bar"
                                                                _ = v173
                                                            } else {
                                                                v174 := []int{fi109, fi108, (fi108 >> ((v3) & 63)), int(v5)}
                                                                var v175 string = "hello"
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                        }
                                                    } else {
                                                        var v176 int8 = (^(v4))
                                                        v3 = uint(v150)
                                                        _ = v176
                                                    }
                                                    var v177 bool = true
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v150 = v150
                                                    default:
                                                        var v178 int16 = 1
                                                        _ = v178
                                                    }
                                                    _ = v177
                                                }
                                                if false {
                                                    _ = len(v2)
                                                    var v179 int = fi108
                                                    var v180 uint64 = 42
                                                    _ = v179
                                                    _ = v180
                                                } else {
                                                    switch v137 {
                                                    case 184:
                                                        var v181 int32 = v138
                                                        _ = v181
                                                    }
                                                    v182 := []int{0, (fi109 / ((fi108) | 1)), (v137 ^ 42)}
                                                    _ = v182
                                                }
                                                fmt.Println(v1, v3)
                                                var v183 uint = v3
                                                _ = v150
                                                _ = v183
                                            default:
                                                v5 -= v5
                                            }
                                        } else {
                                            v1 = 0.1
                                        }
                                        v138 -= -2147483648
                                    }
                                    v2[("a" + "hello")] = (^(v5))
                                    _ = v137
                                    _ = v138
                                }
                                fmt.Println(v5, fi108)
                            } else {
                                if ((true || (1 >= -9223372036854775808)) && true) {
                                    var v184 float32 = 1e10
                                    _ = v184
                                }
                                fmt.Println(v5, v1, v4)
                            }
                        }
                        _ = v134
                    }
                    if (((((("\n\t" + "") + ((("hello" + "\n\t") + "foo bar") + (("\n\t" + ("a" + ("hello" + ""))) + "\n\t"))) + "hello") + "") + "hello") > "hello") {
                        var v185 uint8 = 67
                        _ = v185
                    } else {
                        var v186 uint32 = 218
                        v3 -= 516
                        var v187 bool = (9223372036854775807 > 0)
                        delete(v2, "")
                        if (!false) {
                            v188 := GenS11{F0: false, F1: "a", F2: 0}
                            _ = v188
                        } else {
                            if (v4 <= (v4 ^ (v4 / ((-1) | 1)))) {
                                var v189 float64 = ((1.5 - 0.0) + (float64(v6) - 1e10))
                                var v190 uint32 = 65535
                                var v191 bool = (42 == 42)
                                _ = v189
                                _ = v190
                                _ = v191
                            }
                            delete(v2, ("foo bar" + (("" + ("foo bar" + "")) + "")))
                            fmt.Println(v187, v1, v186)
                        }
                        _ = v186
                        _ = v187
                    }
                    v192 := []int{121, fi108, (fi108 * (fi108 ^ fi108)), fi108}
                    v4 += -46
                    var v193 uint8 = 255
                    if (true || ((true && true) && (3.14159 < -2.25))) {
                        var v194 int16 = 32767
                        if (42 >= 4294967295) {
                            var v195 int = fi108
                            if len(v192) > 0 {
                                v192[0] = (fi109 << (((^(v3))) & 63))
                            }
                            _ = v195
                        } else {
                            _ = len(v192)
                        }
                        _ = v194
                    }
                    v6 = 3.14159
                    _ = v192
                    _ = v193
                } else {
                    if false {
                        var v196 uint8 = 127
                        _ = v196
                    }
                }
                _ = fi109
            }
            v197 := GenS6{F0: v4}
            switch "" {
            case "foo bar":
                if (128 >= fi108) {
                    fmt.Println(v4, v1)
                    fmt.Println(fi108, v3, v6)
                    fmt.Println(fi108, v5, v6)
                }
            default:
                var v198 bool = true
                _ = v198
            }
            _ = fi108
            _ = v197
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
