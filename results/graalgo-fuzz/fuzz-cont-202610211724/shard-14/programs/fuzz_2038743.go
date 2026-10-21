package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 bool; F1 int32; F2 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 int64; F2 uint8 }
type GenS10 struct { F0 int32; F1 uint16; F2 bool }
type GenS11 struct { F0 uint64; F1 int64; F2 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 string; F1 string }
type GenS14 struct { F0 int8; F1 string; F2 int8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint; F1 int64 }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 int64; F1 bool; F2 uint32 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint16; F1 int; F2 uint64 }
func genFn0(p0 int16, p1 int16, p2 uint64) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 = (p0 * 0)
        _ = fi5
    }
    var v6 string = "a"
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 1, F1: -9223372036854775808, F2: false}
    v9 := GenS2{F0: 239}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint64) string {
    v10 := []int{1}
    var v11 int64 = int64(p0)
    var v12 int16 = 90
    var v13 bool = (!((!true) || (true || true)))
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return ("hello" + "")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 float32 = float32(3.14159)
    v4 := []string{"hello", (("hello" + "hello") + "\n\t"), (("\n\t" + "hello") + "a")}
    if (4294967295 <= 42) {
        if false {
            var v14 int16 = -32768
            _ = v14
        }
    } else {
        var v15 int = 127
        v15 += v15
        fmt.Println(v15, v3)
        _ = v15
    }
    var v16 int = 128
    switch v16 {
    case -223:
        v3 += v3
    default:
        v3 -= 3.14159
    }
    v16 = v16
    fmt.Println(v3, v3)
    if (!true) {
        if true {
            for fi17 := 0; fi17 < 17; fi17++ {
                v3 = float32(((0.1 + 0.1) * 0.0))
                _ = fi17
            }
            var v18 bool = true
            fmt.Println(v16, v18, v3)
            _ = v18
        }
        switch (^(v16)) {
        case 42:
            if (2 <= 1) {
                delete(v2, "a")
                var v19 uint32 = 42
                _ = v19
            }
        case 32768:
            var v20 uint8 = 128
            _ = v20
        }
        switch (("" + "hello") + "hello") {
        case "\n\t":
            switch ("" + "") {
            case "a":
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 int32 = 0
                    switch -2147483648 {
                    case 32767:
                        var v23 uint32 = 245
                        _ = v23
                    case -1:
                        v24 := GenS3{F0: -1, F1: -1, F2: -197}
                        _ = v24
                    case 319:
                        v16 = -9223372036854775808
                    }
                    for fi25 := 0; fi25 < 9; fi25++ {
                        if (2 <= 17) {
                            var v26 int16 = 1
                            if false {
                                if (v22 < (v22 % (((v22 / (((v22 % ((-1) | 1))) | 1))) | 1))) {
                                    v27 := v1["hello"]
                                    switch (("" + "\n\t") + "foo bar") {
                                    case "a":
                                        var v28 int16 = (v26 + v26)
                                        _ = v28
                                    }
                                    var v29 uint64 = 1
                                    _ = v27
                                    _ = v29
                                } else {
                                    if true {
                                        v30 := v2["\n\t"]
                                        switch "\n\t" {
                                        case "hello":
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                v16 = (-(v16))
                                            } else {
                                                v3 = v3
                                            }
                                        default:
                                            var v31 float32 = (0.0 * 0.0)
                                            _ = v31
                                        }
                                        _ = len(v4)
                                        _ = v30
                                    } else {
                                        v32 := GenS4{F0: -621}
                                        var v33 uint = 255
                                        _ = v32
                                        _ = v33
                                    }
                                    if (!((float64(((-2.25 + v3) + v3)) + 1.5) == (0.0 / 1.5))) {
                                        _ = len(v4)
                                        v16 = -608
                                        if true {
                                            v34 := GenS4{F0: 1}
                                            v35 := GenS5{F0: 42, F1: -662}
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if false {
                                                v36 := GenS2{F0: 45}
                                                _ = v36
                                            }
                                            switch 47 {
                                            case 313:
                                                if (v26 >= (v26 + (v26 + (v26 * 1000)))) {
                                                    if (!(42 >= 32767)) {
                                                        switch 9223372036854775807 {
                                                        case 42:
                                                            _ = len(v4)
                                                            if (0 >= 428) {
                                                                var v37 uint16 = 2
                                                                switch fi21 {
                                                                case 32768:
                                                                    var v38 int16 = v26
                                                                    _ = v38
                                                                case -2147483648:
                                                                    delete(v2, "\n\t")
                                                                }
                                                                if ((v22 - v22) <= (-(v22))) {
                                                                    if ((0 > 0) || (47 == 1)) {
                                                                        fmt.Println(v3, v3, v26)
                                                                        v37 = v37
                                                                    }
                                                                } else {
                                                                    v22 -= v22
                                                                    var v39 uint16 = v37
                                                                    v39 -= 0
                                                                    _ = v39
                                                                }
                                                                fmt.Println(v37)
                                                                switch (("\n\t" + "hello") + "hello") {
                                                                case "":
                                                                    var v40 int8 = 1
                                                                    _ = v40
                                                                case "foo bar":
                                                                    v37 = (v37 | v37)
                                                                case "\n\t":
                                                                    v26 -= (-(v26))
                                                                }
                                                                v41 := v2[(("foo bar" + "a") + "hello")]
                                                                _ = v37
                                                                _ = v41
                                                            } else {
                                                                v16 += 0
                                                            }
                                                        case -9223372036854775808:
                                                            var v42 int8 = 0
                                                            _ = v42
                                                        case 32768:
                                                            var v43 bool = ((879 == 1) || true)
                                                            _ = v43
                                                        default:
                                                            v26 += v26
                                                        }
                                                    } else {
                                                        v3 -= v3
                                                    }
                                                } else {
                                                    var v44 float32 = v3
                                                    if true {
                                                        v44 = v44
                                                        switch fi21 {
                                                        case 42:
                                                            var v45 uint8 = 127
                                                            _ = v45
                                                        case 32767:
                                                            var v46 int16 = v26
                                                            _ = v46
                                                        case -141:
                                                            if (v16 != (-(fi21))) {
                                                                if (v26 != 32767) {
                                                                    switch -2147483648 {
                                                                    case 0:
                                                                        v47 := GenS4{F0: -1}
                                                                        _ = v47
                                                                    }
                                                                    var v48 int64 = 0
                                                                    _ = v48
                                                                }
                                                                var v49 uint32 = 38
                                                                _ = v49
                                                            }
                                                        }
                                                        fmt.Println(v44, v26, fi25)
                                                        if (!(v44 != 1e10)) {
                                                            v16 -= (-(fi25))
                                                            v4 = append(v4, "hello")
                                                        } else {
                                                            v50 := GenS6{F0: false, F1: v22, F2: 986}
                                                            fmt.Println(v3, v3)
                                                            if (375 < 42) {
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v51 int16 = v26
                                                                    _ = v51
                                                                case "a":
                                                                    if true {
                                                                        v52 := GenS7{F0: -128}
                                                                        switch fi21 {
                                                                        case 128:
                                                                            if (("\n\t" > "foo bar") || ((9223372036854775807 < (fi25 & v16)) || true)) {
                                                                                var v53 uint = 820
                                                                                v54 := GenS8{F0: ((-244 == 9223372036854775807) || false)}
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                v16 = (fi21 << ((334) & 63))
                                                                                v44 -= (v3 / 1e10)
                                                                                v55 := v2["foo bar"]
                                                                                v56 := GenS1{F0: 0, F1: -1, F2: (((4294967295 > 42) || true) && true)}
                                                                                _ = v55
                                                                                _ = v56
                                                                            }
                                                                        default:
                                                                            var v57 bool = ((!(4294967295 == 0)) && true)
                                                                            _ = v57
                                                                        }
                                                                        if (2 >= 32767) {
                                                                            v58 := GenS4{F0: 81}
                                                                            _ = v58
                                                                        } else {
                                                                            if (!(100 == -4)) {
                                                                                v59 := GenS9{F0: 592, F1: 936, F2: 1}
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    var v60 uint = 0
                                                                                    _ = v60
                                                                                case "hello":
                                                                                    v61 := GenS8{F0: false}
                                                                                    _ = v61
                                                                                case "foo bar":
                                                                                    v16 += fi21
                                                                                }
                                                                                _ = v59
                                                                            }
                                                                            var v62 int8 = 127
                                                                            v3 += ((v3 - (v3 + 0.1)) - -2.25)
                                                                            _ = v62
                                                                        }
                                                                        _ = v52
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v63 bool = (255 == 4294967295)
                                                                            var v64 float64 = -2.25
                                                                            if (42 <= -1) {
                                                                                fmt.Println(v63, fi25)
                                                                                if v63 {
                                                                                    var v65 float32 = v44
                                                                                    if (79 == 85) {
                                                                                        _ = len(v1)
                                                                                        fmt.Println(v22, v3)
                                                                                    }
                                                                                    var v66 uint = 2
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                }
                                                                                var v67 uint8 = 10
                                                                                _ = v67
                                                                            } else {
                                                                                var v68 uint32 = 2
                                                                                _ = v68
                                                                            }
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                case "":
                                                                    switch ("" + "foo bar") {
                                                                    case "a":
                                                                        var v69 int32 = (v22 & (^(v22)))
                                                                        _ = v69
                                                                    case "":
                                                                        var v70 int32 = v22
                                                                        _ = v70
                                                                    }
                                                                }
                                                                v71 := GenS4{F0: 9223372036854775807}
                                                                switch fi21 {
                                                                case 42:
                                                                    var v72 uint = 0
                                                                    _ = v72
                                                                case -715:
                                                                    var v73 int64 = -978
                                                                    _ = v73
                                                                default:
                                                                    v74 := GenS1{F0: 32767, F1: -394, F2: (42 < 127)}
                                                                    _ = v74
                                                                }
                                                                _ = v71
                                                            }
                                                            switch (fi25 << ((18446744073709551615) & 63)) {
                                                            case 32767:
                                                                switch "" {
                                                                case "a":
                                                                    v16 = fi25
                                                                case "foo bar":
                                                                    switch (("\n\t" + "hello") + "hello") {
                                                                    case "":
                                                                        if false {
                                                                            var v75 string = ""
                                                                            v76 := GenS9{F0: 255, F1: 1, F2: 1}
                                                                            switch v16 {
                                                                            case 9223372036854775807:
                                                                                v22 = v22
                                                                            case 32768:
                                                                                v16 = (fi25 / ((fi21) | 1))
                                                                            case 4294967295:
                                                                                v3 += v44
                                                                            default:
                                                                                switch (fi21 & (fi25 ^ fi21)) {
                                                                                case 4294967295:
                                                                                    var v77 uint16 = 599
                                                                                    _ = v77
                                                                                }
                                                                            }
                                                                            _ = v75
                                                                            _ = v76
                                                                        } else {
                                                                            if (!true) {
                                                                                _ = len(v4)
                                                                                v22 = 2147483647
                                                                                switch ("foo bar" + "hello") {
                                                                                case "a":
                                                                                    v4 = append(v4, "")
                                                                                case "hello":
                                                                                    v44 -= (3.14159 / (v3 - v3))
                                                                                    var v78 uint8 = 127
                                                                                    _ = v78
                                                                                default:
                                                                                    v79 := GenS7{F0: -87}
                                                                                    _ = v79
                                                                                }
                                                                            } else {
                                                                                v3 = (1.5 * (1.5 - 0.1))
                                                                                v80 := GenS5{F0: 363, F1: 444}
                                                                                _ = v80
                                                                            }
                                                                            var v81 uint32 = 0
                                                                            _ = v81
                                                                        }
                                                                    case "hello":
                                                                        fmt.Println(v26)
                                                                    default:
                                                                        if (42 < 42) {
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v82 int64 = -1
                                                                    _ = v82
                                                                }
                                                            case 127:
                                                                var v83 int32 = v22
                                                                _ = v83
                                                            default:
                                                                switch 881 {
                                                                case -1:
                                                                    var v84 int32 = v22
                                                                    _ = v84
                                                                }
                                                            }
                                                            fmt.Println(v16, fi21, fi25)
                                                            v3 += v44
                                                            _ = v50
                                                        }
                                                        v22 += (v22 * (v22 & v22))
                                                    }
                                                    _ = v44
                                                }
                                            case 652:
                                                var v85 int32 = v22
                                                _ = v85
                                            }
                                            v86 := GenS10{F0: v22, F1: 32767, F2: (false || (!(!true)))}
                                            _ = v86
                                        }
                                    } else {
                                        v87 := GenS11{F0: 557, F1: -138, F2: v16}
                                        if (0 < 2147483648) {
                                            fmt.Println(v22, fi25)
                                        } else {
                                            if false {
                                                switch "\n\t" {
                                                case "":
                                                    if true {
                                                        v16 -= 127
                                                    } else {
                                                        switch 1 {
                                                        case 127:
                                                            v88 := GenS9{F0: 815, F1: -1, F2: 254}
                                                            _ = v88
                                                        case 128:
                                                            if true {
                                                                var v89 int16 = (v26 | (v26 << ((2) & 63)))
                                                                delete(v2, (("" + ("a" + "")) + "a"))
                                                                v16 -= fi21
                                                                v89 = -32768
                                                                _ = v89
                                                            } else {
                                                                if (!((false || ((v22 % (((v22 << ((255) & 63))) | 1)) <= -2147483648)) || false)) {
                                                                    if (18446744073709551615 <= 1) {
                                                                        _ = len(v4)
                                                                        if false {
                                                                            var v90 uint32 = 46
                                                                            switch "" {
                                                                            case "a":
                                                                                v3 += (v3 - (v3 + 1e10))
                                                                            case "":
                                                                                var v91 uint16 = 2
                                                                                _ = v91
                                                                            default:
                                                                                if false {
                                                                                    v1[("a" + ("" + "\n\t"))] = (^(v90))
                                                                                } else {
                                                                                    var v92 float64 = -2.25
                                                                                    switch 127 {
                                                                                    case 9223372036854775807:
                                                                                        if ((v92 >= 1e10) && true) {
                                                                                            var v93 string = ""
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v93
                                                                                        } else {
                                                                                            if false {
                                                                                                v26 -= v26
                                                                                                if false {
                                                                                                    v94 := v2["a"]
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        var v95 bool = false
                                                                                                        _ = v95
                                                                                                    default:
                                                                                                        var v96 float64 = (v92 + (float64(0.1) + float64(1e10)))
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    if (!true) {
                                                                                                        var v97 uint16 = 32767
                                                                                                        var v98 float64 = v92
                                                                                                        var v99 int = fi21
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                        _ = v99
                                                                                                    }
                                                                                                    fmt.Println(v22, fi21)
                                                                                                    v1["a"] = (v90 * 2147483648)
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    v100 := GenS12{F0: fi25}
                                                                                                    _ = v100
                                                                                                }
                                                                                            }
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = ""
                                                                                            }
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    var v101 uint32 = 2
                                                                                                    _ = v101
                                                                                                case "hello":
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        switch "\n\t" {
                                                                                                        case "\n\t":
                                                                                                            v26 += (v26 + 0)
                                                                                                        case "foo bar":
                                                                                                            switch fi25 {
                                                                                                            case 4294967295:
                                                                                                                switch ("" + "\n\t") {
                                                                                                                case "foo bar":
                                                                                                                    v102 := GenS13{F0: ((("" + "hello") + "hello") + "a"), F1: ""}
                                                                                                                    _ = v102
                                                                                                                case "a":
                                                                                                                    _ = len(v2)
                                                                                                                case "\n\t":
                                                                                                                    v16 += fi25
                                                                                                                default:
                                                                                                                    var v103 uint8 = 128
                                                                                                                    _ = v103
                                                                                                                }
                                                                                                            case -7:
                                                                                                                var v104 int32 = (v22 + (v22 % ((v22) | 1)))
                                                                                                                _ = v104
                                                                                                            case 9223372036854775807:
                                                                                                                fmt.Println(fi25, v92)
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v105 int8 = 60
                                                                                                            _ = v105
                                                                                                        default:
                                                                                                            v26 += v26
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v106 := GenS5{F0: 1, F1: -682}
                                                                                                    _ = v106
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case -1:
                                                                                        if true {
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                var v107 float32 = v3
                                                                                                _ = v107
                                                                                            case "\n\t":
                                                                                                if (!false) {
                                                                                                    if false {
                                                                                                        if ((true || true) || (v92 >= v92)) {
                                                                                                            _ = len(v1)
                                                                                                            fmt.Println(fi21, v22, v26)
                                                                                                        } else {
                                                                                                            fmt.Println(v26, fi21, v16)
                                                                                                            v108 := GenS5{F0: -1, F1: 1}
                                                                                                            switch v16 {
                                                                                                            case 32768:
                                                                                                                v16 += 4294967295
                                                                                                            case 127:
                                                                                                                v109 := GenS14{F0: -54, F1: ("foo bar" + (("hello" + "\n\t") + (("foo bar" + "foo bar") + ("\n\t" + ("a" + "\n\t"))))), F2: 71}
                                                                                                                _ = v109
                                                                                                            case -1:
                                                                                                                v92 -= (3.14159 - v92)
                                                                                                            default:
                                                                                                                var v110 int32 = (v22 << ((18446744073709551615) & 63))
                                                                                                                _ = v110
                                                                                                            }
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        v111 := GenS9{F0: 431, F1: -1, F2: 1}
                                                                                                        v4 = append(v4, ("a" + "a"))
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    var v112 int = 32767
                                                                                                    v16 = 32767
                                                                                                    var v113 uint = 42
                                                                                                    _ = v112
                                                                                                    _ = v113
                                                                                                }
                                                                                            default:
                                                                                                v114 := GenS11{F0: 18446744073709551615, F1: 1, F2: fi21}
                                                                                                _ = v114
                                                                                            }
                                                                                            v3 = ((v3 + v3) - ((3.14159 / float32(float64(1.5))) - v3))
                                                                                        }
                                                                                    default:
                                                                                        var v115 uint32 = v90
                                                                                        _ = v115
                                                                                    }
                                                                                    v3 -= 1e10
                                                                                    _ = v92
                                                                                }
                                                                            }
                                                                            _ = v90
                                                                        }
                                                                    } else {
                                                                        var v116 int64 = 42
                                                                        switch (v16 % ((-750) | 1)) {
                                                                        case 42:
                                                                            v116 = (v116 >> ((0) & 63))
                                                                        case -878:
                                                                            fmt.Println(v26, v3, fi25)
                                                                            v117 := GenS15{F0: 127}
                                                                            _ = v117
                                                                        default:
                                                                            fmt.Println(fi25)
                                                                        }
                                                                        _ = v116
                                                                    }
                                                                    v3 -= ((((3.14159 * float32(1e10)) * 1e10) / 3.14159) / (1.5 / 0.1))
                                                                    if (2147483647 >= (v22 % (((v22 << ((255) & 63))) | 1))) {
                                                                        v118 := GenS15{F0: -18}
                                                                        if (true || ((1.5 + float64((float32(0.1) - ((v3 / 1e10) + v3)))) > (float64(1.5) + (0.1 + (0.0 - 3.14159))))) {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v119 uint64 = 2
                                                                                _ = v119
                                                                            case "":
                                                                                var v120 float32 = 0.1
                                                                                _ = v120
                                                                            case "\n\t":
                                                                                var v121 uint64 = 18446744073709551615
                                                                                _ = v121
                                                                            }
                                                                            var v122 float64 = 3.14159
                                                                            _ = v122
                                                                        }
                                                                        _ = v118
                                                                    } else {
                                                                        if (!(!(!(!(false && true))))) {
                                                                            var v123 int8 = -1
                                                                            _ = v123
                                                                        }
                                                                        if false {
                                                                            var v124 int32 = (^(v22))
                                                                            if (("foo bar" + ((("a" + "a") + "\n\t") + ("foo bar" + "a"))) != "") {
                                                                                var v125 int64 = -1
                                                                                v1[(("foo bar" + "") + ("hello" + (("a" + "a") + "")))] = 65535
                                                                                _ = v125
                                                                            } else {
                                                                                fmt.Println(v3, v22)
                                                                                var v126 int8 = -17
                                                                                if false {
                                                                                    var v127 int64 = 139
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    v4 = append(v4, "a")
                                                                                    _ = v127
                                                                                } else {
                                                                                    fmt.Println(v3)
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    v128 := GenS4{F0: -766}
                                                                                    var v129 uint64 = 396
                                                                                    fmt.Println(v22, v124)
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                }
                                                                                _ = v126
                                                                            }
                                                                            _ = v124
                                                                        }
                                                                        var v130 float64 = -2.25
                                                                        _ = v130
                                                                    }
                                                                    v131 := GenS16{F0: 750}
                                                                    _ = v131
                                                                }
                                                            }
                                                        default:
                                                            var v132 bool = (!(3.14159 == (-2.25 - (v3 - 3.14159))))
                                                            _ = v132
                                                        }
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        if true {
                                                            v133 := GenS17{F0: 65535, F1: 1}
                                                            _ = len(v4)
                                                            if (!true) {
                                                                var v134 int32 = v22
                                                                _ = v134
                                                            } else {
                                                                var v135 float64 = float64(1e10)
                                                                if (654 <= -9223372036854775808) {
                                                                    v3 = v3
                                                                    var v136 int64 = 9223372036854775807
                                                                    var v137 uint32 = 2147483648
                                                                    _ = v136
                                                                    _ = v137
                                                                } else {
                                                                    if (((v3 + v3) / (0.0 - v3)) <= (v3 / ((v3 / 1e10) + v3))) {
                                                                        if (!false) {
                                                                            if false {
                                                                                fmt.Println(v26)
                                                                                var v138 uint32 = 42
                                                                                if (42 > 1) {
                                                                                    var v139 string = ("a" + "")
                                                                                    _ = v139
                                                                                }
                                                                                v140 := GenS16{F0: 18446744073709551615}
                                                                                _ = v138
                                                                                _ = v140
                                                                            } else {
                                                                                _ = len(v1)
                                                                            }
                                                                            fmt.Println(v135, v16)
                                                                        }
                                                                        var v141 float64 = v135
                                                                        var v142 uint64 = 462
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    if (!true) {
                                                                        fmt.Println(fi25)
                                                                    } else {
                                                                        v26 -= 42
                                                                        fmt.Println(fi25)
                                                                        if (!(!(!true))) {
                                                                            var v143 string = ("hello" + ("\n\t" + ("\n\t" + (("" + "hello") + "hello"))))
                                                                            v144 := v1["a"]
                                                                            _ = v143
                                                                            _ = v144
                                                                        }
                                                                        if (v135 <= v135) {
                                                                            if true {
                                                                                v16 += (v16 >> ((2) & 63))
                                                                            }
                                                                        } else {
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                }
                                                                var v145 float64 = (0.0 * float64((-2.25 * (v3 - (v3 / 1.0)))))
                                                                _ = v135
                                                                _ = v145
                                                            }
                                                            _ = v133
                                                        }
                                                    } else {
                                                        v3 = (((v3 * (((v3 / v3) - v3) + v3)) + v3) / 0.1)
                                                        var v146 int16 = v26
                                                        v147 := GenS11{F0: 1, F1: 0, F2: 2147483647}
                                                        _ = v146
                                                        _ = v147
                                                    }
                                                case "hello":
                                                    if (true || ((("\n\t" + ("" + "\n\t")) != "foo bar") || false)) {
                                                        var v148 float64 = (1.5 - float64(-2.25))
                                                        _ = v148
                                                    }
                                                    v26 -= (-(v26))
                                                    var v149 float32 = 3.14159
                                                    _ = v149
                                                default:
                                                    var v150 uint32 = 2147483647
                                                    _ = v150
                                                }
                                                v151 := GenS18{F0: (((false && ((fi25 / ((fi25) | 1)) < -1)) && ((v26 | (v26 ^ v26)) != (v26 - (v26 - (v26 + (v26 >> ((255) & 63))))))) && (211 >= -1)), F1: 1}
                                                _ = v151
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        v1[""] = 4294967295
                                                        switch ("foo bar" + "a") {
                                                        case "hello":
                                                            if (!true) {
                                                                var v152 uint16 = 2
                                                                if (!(false || ((!true) || true))) {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v153 uint64 = 1
                                                                        _ = v153
                                                                    default:
                                                                        fmt.Println(v152)
                                                                    }
                                                                    v154 := GenS19{F0: -9223372036854775808, F1: (false && false), F2: 559}
                                                                    v155 := GenS7{F0: 121}
                                                                    _ = v154
                                                                    _ = v155
                                                                } else {
                                                                    fmt.Println(v22, fi25)
                                                                    v156 := GenS14{F0: 24, F1: "\n\t", F2: 42}
                                                                    _ = v156
                                                                }
                                                                _ = v152
                                                            }
                                                        case "foo bar":
                                                            v26 += v26
                                                        case "a":
                                                            fmt.Println(v3)
                                                        default:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if (0 == 42) {
                                                                    var v157 string = "a"
                                                                    v157 = v157
                                                                    _ = v157
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v158 := GenS20{F0: 1}
                                                                        _ = v158
                                                                    case "a":
                                                                        if false {
                                                                            v3 -= (v3 + v3)
                                                                            v26 = 32767
                                                                            var v159 uint8 = 1
                                                                            _ = v159
                                                                        } else {
                                                                            v4 = append(v4, "")
                                                                        }
                                                                    case "hello":
                                                                        v16 = 0
                                                                    }
                                                                    v4 = append(v4, "")
                                                                    v160 := GenS19{F0: -1, F1: ((v26 % (((v26 % (((v26 + (v26 | (v26 / ((325) | 1))))) | 1))) | 1)) >= v26), F2: 4294967295}
                                                                    _ = v160
                                                                }
                                                            case "foo bar":
                                                                v161 := GenS7{F0: -1}
                                                                _ = v161
                                                            default:
                                                                fmt.Println(v22, v22, fi25)
                                                            }
                                                        }
                                                        switch fi25 {
                                                        case -9223372036854775808:
                                                            var v162 bool = (false && true)
                                                            _ = v162
                                                        }
                                                    case "\n\t":
                                                        var v163 uint32 = 2
                                                        _ = v163
                                                    case "hello":
                                                        v16 = (fi21 << ((111) & 63))
                                                    default:
                                                        v3 += ((float32(float64(0.0)) + (v3 * v3)) + 1.5)
                                                    }
                                                default:
                                                    v164 := GenS17{F0: 18446744073709551615, F1: -1}
                                                    var v165 float64 = (0.1 + 3.14159)
                                                    var v166 uint32 = 51
                                                    _ = v164
                                                    _ = v165
                                                    _ = v166
                                                }
                                                var v167 int8 = -13
                                                var v168 string = ("foo bar" + "foo bar")
                                                _ = v167
                                                _ = v168
                                            }
                                            var v169 int32 = v22
                                            switch fi21 {
                                            case 614:
                                                var v170 int64 = -9223372036854775808
                                                _ = v170
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                            v22 = (v169 - (^(v169)))
                                            _ = v169
                                        }
                                        _ = v87
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = (("hello" + "a") + ("hello" + "\n\t"))
                                }
                            }
                            fmt.Println(fi21)
                            _ = v26
                        } else {
                            v171 := GenS21{F0: 1, F1: (fi25 * fi25), F2: 787}
                            _ = v171
                        }
                        var v172 int64 = 9223372036854775807
                        var v173 uint = 2
                        var v174 uint64 = 1
                        v4 = append(v4, "\n\t")
                        _ = fi25
                        _ = v172
                        _ = v173
                        _ = v174
                    }
                    _ = fi21
                    _ = v22
                }
            default:
                v3 = float32((1.5 * 0.0))
            }
        case "foo bar":
            v16 = 4294967295
        }
    }
    v3 -= (v3 * 1e10)
    fmt.Println(v16, v16)
    var v175 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v175:", v175)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v175
    _ = v2
    _ = v3
    _ = v4
}
