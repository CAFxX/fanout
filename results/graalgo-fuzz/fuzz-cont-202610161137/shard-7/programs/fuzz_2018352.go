package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 bool; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 int; F2 uint }
type GenS11 struct { F0 uint16; F1 uint64; F2 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 int8 }
type GenS15 struct { F0 int16; F1 string }

func main() {
    var v1 uint64 = 0
    var v2 uint8 = 255
    v3 := make(map[string]string)
    var v4 uint64 = v1
    v5 := []string{"", "a"}
    v6 := GenS1{F0: (^(v1))}
    fmt.Println(v4, v4, v1)
    v1 -= (^(v1))
    v7 := GenS1{F0: v1}
    switch 2147483647 {
    case -394:
        switch "hello" {
        case "a":
            v8 := v3[""]
            _ = v8
        case "\n\t":
            if (("hello" + "a") == ("foo bar" + (("\n\t" + (((("hello" + "foo bar") + "foo bar") + "a") + (((("foo bar" + ("\n\t" + "foo bar")) + "") + "a") + ("hello" + "a")))) + ("hello" + ("foo bar" + "a"))))) {
                v3["hello"] = ""
                delete(v3, "a")
            } else {
                if ((v2 >> ((65535) & 63)) < (v2 >> ((283) & 63))) {
                    v1 = v4
                    fmt.Println(v2, v2, v2)
                    for fi9 := 0; fi9 < 17; fi9++ {
                        for fi10 := 0; fi10 < 19; fi10++ {
                            if true {
                                var v11 uint64 = v1
                                v1 -= (^(v4))
                                var v12 int16 = 1
                                _ = v11
                                _ = v12
                            }
                            if (!(!(774 == 65535))) {
                                if (4294967295 != 2147483647) {
                                    if (9223372036854775807 == -9223372036854775808) {
                                        fmt.Println(v4, v1)
                                        v1 = v1
                                        if (-1 < 21) {
                                            if false {
                                                var v13 uint = 2
                                                _ = v13
                                            } else {
                                                if (18446744073709551615 == 833) {
                                                    v1 += 18446744073709551615
                                                    fmt.Println(v4, v2)
                                                    switch fi10 {
                                                    case 4294967295:
                                                        if ((-9223372036854775808 <= int64(v2)) || true) {
                                                            fmt.Println(v4, fi9)
                                                        } else {
                                                            switch fi9 {
                                                            case 4294967295:
                                                                v14 := GenS2{F0: true}
                                                                _ = v14
                                                            case 0:
                                                                if (0 < 1) {
                                                                    switch ("a" + "\n\t") {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v15 uint8 = (^(v2))
                                                                            _ = v15
                                                                        } else {
                                                                            v2 += v2
                                                                            v1 = v4
                                                                            var v16 float32 = 1.5
                                                                            _ = v16
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(v1)
                                                                    }
                                                                }
                                                            case 923:
                                                                v17 := GenS3{F0: 32767}
                                                                _ = v17
                                                            }
                                                            fmt.Println(fi9, v4)
                                                            var v18 int64 = 767
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            if true {
                                                                v5 = append(v5, "hello")
                                                                v2 += v2
                                                            } else {
                                                                v3["\n\t"] = "a"
                                                                v18 = -9223372036854775808
                                                                var v19 float32 = -2.25
                                                                var v20 uint8 = (v2 / ((2) | 1))
                                                                _ = v19
                                                                _ = v20
                                                            }
                                                            if true {
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v21 int8 = 1
                                                                    _ = v21
                                                                case "hello":
                                                                    var v22 uint16 = 2
                                                                    _ = v22
                                                                case "a":
                                                                    v18 -= (v18 << ((42) & 63))
                                                                default:
                                                                    var v23 uint64 = (v1 * (v4 - v4))
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v24 bool = (("a" <= "") && ((-2.25 / 1e10) <= (float64((float32(1.5) * float32(((3.14159 * float64(-2.25)) / (((1.5 * ((float64(1e10) / ((0.0 - -2.25) * (1e10 / (1e10 + (-2.25 * (1.5 + 3.14159)))))) - 0.1)) + (((3.14159 - (0.0 / ((0.1 * 3.14159) / -2.25))) * 0.1) * -2.25)) * (-2.25 * 0.1)))))) - 0.1)))
                                                                        _ = v24
                                                                    case "foo bar":
                                                                        v23 -= v23
                                                                    }
                                                                    _ = v23
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    if (v18 >= v18) {
                                                                        v25 := GenS3{F0: 126}
                                                                        var v26 int64 = -9223372036854775808
                                                                        _ = v25
                                                                        _ = v26
                                                                    } else {
                                                                        var v27 uint8 = (v2 / (((^(v2))) | 1))
                                                                        fmt.Println(v27)
                                                                        _ = v27
                                                                    }
                                                                }
                                                                var v28 uint = 18446744073709551615
                                                                fmt.Println(v18)
                                                                _ = v28
                                                            } else {
                                                                _ = len(v3)
                                                                if false {
                                                                    var v29 int = fi10
                                                                    v1 = v1
                                                                    v30 := GenS2{F0: ((!true) && false)}
                                                                    _ = v29
                                                                    _ = v30
                                                                }
                                                            }
                                                            v4 = 973
                                                            v31 := GenS4{F0: (v2 / ((v2) | 1))}
                                                            _ = v18
                                                            _ = v31
                                                        }
                                                    case 32768:
                                                        fmt.Println(v2, fi9, fi9)
                                                    case 395:
                                                        var v32 uint64 = (v1 % (((v1 / ((v4) | 1))) | 1))
                                                        _ = v32
                                                    default:
                                                        var v33 int64 = -9223372036854775808
                                                        _ = v33
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        _ = len(v3)
                                                    case "hello":
                                                        fmt.Println(v1, fi9)
                                                    default:
                                                        v2 = 47
                                                    }
                                                    var v34 string = "hello"
                                                    _ = v34
                                                } else {
                                                    v1 = (v1 + (v4 >> ((0) & 63)))
                                                    switch fi9 {
                                                    case 32768:
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v35 uint8 = v2
                                                            _ = v35
                                                        case "foo bar":
                                                            var v36 int8 = -128
                                                            _ = v36
                                                        default:
                                                            v37 := GenS5{F0: 0, F1: 255}
                                                            _ = v37
                                                        }
                                                    case 9223372036854775807:
                                                        v38 := v3[""]
                                                        _ = v38
                                                    case -693:
                                                        _ = len(v5)
                                                    }
                                                }
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            if false {
                                                                switch 127 {
                                                                case 9223372036854775807:
                                                                    v1 = v1
                                                                case -845:
                                                                    var v39 int64 = 649
                                                                    _ = v39
                                                                case -2147483648:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v1 += v4
                                                                    default:
                                                                        v40 := GenS6{F0: "foo bar", F1: v2}
                                                                        _ = v40
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v4 -= 2
                                                        var v41 uint16 = 2
                                                        var v42 uint64 = 25
                                                        fmt.Println(v4)
                                                        v41 = 42
                                                        var v43 uint32 = uint32(v2)
                                                        v44 := GenS4{F0: uint8(v43)}
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    case "foo bar":
                                                        v45 := GenS5{F0: 866, F1: 128}
                                                        _ = v45
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v46 := v3[""]
                                                            _ = v46
                                                        default:
                                                            var v47 float64 = 1e10
                                                            _ = v47
                                                        }
                                                    }
                                                    v48 := GenS3{F0: 32767}
                                                    fmt.Println(v1, v4)
                                                    var v49 bool = (42 != 255)
                                                    var v50 int8 = -128
                                                    switch ("a" + "foo bar") {
                                                    case "hello":
                                                        v51 := v3[("hello" + "hello")]
                                                        var v52 uint16 = 283
                                                        if v49 {
                                                            if (-2147483648 > -222) {
                                                                var v53 int8 = v50
                                                                _ = v53
                                                            } else {
                                                                v5 = append(v5, v51)
                                                                if len(v5) > 0 {
                                                                    v5[0] = (((v51 + v51) + ((v51 + v51) + (v51 + v51))) + v51)
                                                                }
                                                                var v54 uint32 = 3
                                                                _ = v54
                                                            }
                                                            v55 := GenS3{F0: v52}
                                                            _ = v55
                                                        } else {
                                                            var v56 bool = v49
                                                            if false {
                                                                v4 = v4
                                                                var v57 int64 = 42
                                                                _ = v57
                                                            } else {
                                                                v58 := GenS7{F0: v56, F1: (!false), F2: v2}
                                                                _ = v58
                                                            }
                                                            var v59 uint = 0
                                                            v60 := GenS2{F0: (-1 == 42)}
                                                            _ = v56
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        var v61 int8 = (v50 / ((v50) | 1))
                                                        var v62 int64 = -594
                                                        switch (-(fi9)) {
                                                        case 0:
                                                            var v63 int64 = (-(v62))
                                                            _ = v63
                                                        default:
                                                            fmt.Println(v1, v1)
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                        _ = v61
                                                        _ = v62
                                                    case "foo bar":
                                                        v64 := GenS6{F0: "\n\t", F1: 0}
                                                        _ = v64
                                                    case "":
                                                        fmt.Println(fi10)
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v65 float32 = 0.1
                                                        _ = v65
                                                    case "\n\t":
                                                        v3[""] = "a"
                                                    case "hello":
                                                        v5 = append(v5, "\n\t")
                                                    default:
                                                        if (true || (-2147483648 <= 42)) {
                                                            _ = len(v5)
                                                            var v66 uint32 = 42
                                                            _ = v66
                                                        }
                                                    }
                                                    if false {
                                                        fmt.Println(v1)
                                                        _ = len(v5)
                                                        var v67 int64 = 779
                                                        _ = v67
                                                    }
                                                }
                                                if false {
                                                    switch ("" + ("hello" + ((("a" + "") + "foo bar") + ""))) {
                                                    case "foo bar":
                                                        var v68 uint = 18446744073709551615
                                                        _ = v68
                                                    case "\n\t":
                                                        if false {
                                                            if (-743 >= 0) {
                                                                var v69 int16 = 229
                                                                v70 := GenS2{F0: false}
                                                                _ = v69
                                                                _ = v70
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = ("a" + "\n\t")
                                                                }
                                                                var v71 int16 = 42
                                                                _ = v71
                                                            }
                                                        } else {
                                                            v2 = v2
                                                            if (1.5 == 1e10) {
                                                                var v72 uint64 = v1
                                                                if true {
                                                                    v73 := v3[""]
                                                                    _ = v73
                                                                }
                                                                v4 = v72
                                                                _ = v72
                                                            } else {
                                                                if true {
                                                                    v5 = append(v5, ("hello" + "hello"))
                                                                    if (-1 != -2147483648) {
                                                                        v74 := GenS7{F0: true, F1: (!(v1 <= 18446744073709551615)), F2: v2}
                                                                        if false {
                                                                            v75 := GenS8{F0: 566, F1: ("a" + ("a" + ("foo bar" + "a")))}
                                                                            fmt.Println(v4)
                                                                            _ = v75
                                                                        } else {
                                                                            delete(v3, "\n\t")
                                                                            var v76 int8 = -31
                                                                            var v77 string = "hello"
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        v1 = v1
                                                                        _ = v74
                                                                    }
                                                                    if true {
                                                                        v4 -= 4294967295
                                                                        var v78 bool = true
                                                                        _ = v78
                                                                    }
                                                                    var v79 string = "hello"
                                                                    switch (fi9 << ((255) & 63)) {
                                                                    case -639:
                                                                        v80 := GenS9{F0: v4}
                                                                        _ = v80
                                                                    case 75:
                                                                        v79 = v79
                                                                    case 346:
                                                                        switch 0 {
                                                                        case 32767:
                                                                            _ = len(v3)
                                                                        }
                                                                    default:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                    if (-9223372036854775808 == (fi10 / ((0) | 1))) {
                                                                        delete(v3, (v79 + ((v79 + "\n\t") + v79)))
                                                                        v1 = (^(v4))
                                                                        if (true && (!true)) {
                                                                            var v81 int32 = -2147483648
                                                                            v82 := GenS10{F0: 42, F1: fi9, F2: uint(v2)}
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            v4 = (v1 | (^(v4)))
                                                                        }
                                                                    } else {
                                                                        v3[v79] = (((((v79 + "hello") + ("a" + ((v79 + ("a" + "a")) + v79))) + "\n\t") + v79) + v79)
                                                                    }
                                                                    _ = v79
                                                                } else {
                                                                    v4 += (v4 % (((v1 % ((v4) | 1))) | 1))
                                                                    fmt.Println(v2, v4)
                                                                    _ = len(v3)
                                                                }
                                                                var v83 string = ""
                                                                _ = v83
                                                            }
                                                            switch fi10 {
                                                            case 32768:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            default:
                                                                if (v1 >= v4) {
                                                                    v84 := GenS4{F0: (v2 ^ (^(v2)))}
                                                                    _ = v84
                                                                } else {
                                                                    if true {
                                                                        switch (fi10 - -134) {
                                                                        case 794:
                                                                            switch fi9 {
                                                                            case 724:
                                                                                switch fi10 {
                                                                                case 2147483647:
                                                                                    var v85 int = (fi9 >> ((4294967295) & 63))
                                                                                    _ = v85
                                                                                }
                                                                            case 32768:
                                                                                var v86 float32 = 1.5
                                                                                _ = v86
                                                                            case 128:
                                                                                if (!((912 == 2147483647) || false)) {
                                                                                    v1 -= 751
                                                                                    var v87 bool = (0 != 127)
                                                                                    v1 += (^(v1))
                                                                                    v2 = v2
                                                                                    if true {
                                                                                        var v88 uint8 = v2
                                                                                        v3["foo bar"] = ("" + ("a" + "a"))
                                                                                        v4 += 0
                                                                                        v5 = append(v5, "")
                                                                                        if (!(9223372036854775807 == -9223372036854775808)) {
                                                                                            var v89 uint64 = 837
                                                                                            _ = v89
                                                                                        } else {
                                                                                            var v90 uint32 = 2147483648
                                                                                            if v87 {
                                                                                                v91 := GenS9{F0: 969}
                                                                                                _ = v91
                                                                                            }
                                                                                            var v92 float32 = ((1e10 * -2.25) - -2.25)
                                                                                            _ = len(v3)
                                                                                            v87 = (!((v4 % ((42) | 1)) < 42))
                                                                                            v88 = (v2 * 114)
                                                                                            switch (fi10 % ((0) | 1)) {
                                                                                            case 4294967295:
                                                                                                if v87 {
                                                                                                    v93 := GenS4{F0: v88}
                                                                                                    v94 := GenS3{F0: 532}
                                                                                                    v95 := v3[(("" + "foo bar") + (("\n\t" + "a") + ("a" + ("" + "hello"))))]
                                                                                                    if (0.1 != ((v92 / v92) - v92)) {
                                                                                                        fmt.Println(v2, v88, v1)
                                                                                                        var v96 int8 = 127
                                                                                                        var v97 int8 = (v96 / ((v96) | 1))
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                    } else {
                                                                                                        var v98 int32 = -723
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    var v99 uint8 = 42
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                    _ = v95
                                                                                                    _ = v99
                                                                                                }
                                                                                            case 2147483647:
                                                                                                fmt.Println(v1, fi10, v92)
                                                                                            default:
                                                                                                var v100 int64 = int64(v2)
                                                                                                _ = v100
                                                                                            }
                                                                                            v90 = 2147483648
                                                                                            _ = v90
                                                                                            _ = v92
                                                                                        }
                                                                                        v87 = v87
                                                                                        _ = v88
                                                                                    }
                                                                                    _ = v87
                                                                                } else {
                                                                                    fmt.Println(fi10, fi9)
                                                                                }
                                                                            }
                                                                        case -2147483648:
                                                                            v1 -= 177
                                                                        case 295:
                                                                            if false {
                                                                                if false {
                                                                                    v101 := GenS8{F0: 255, F1: "a"}
                                                                                    fmt.Println(v1, fi9, v2)
                                                                                    _ = v101
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v4 = v1
                                                                                    case "\n\t":
                                                                                        var v102 float32 = ((((0.0 - 0.0) / 1.5) / 3.14159) * (1.5 / (3.14159 + 0.0)))
                                                                                        _ = v102
                                                                                    case "hello":
                                                                                        if (!true) {
                                                                                            if true {
                                                                                                v1 -= (^(v4))
                                                                                                v4 = v4
                                                                                                var v103 string = "foo bar"
                                                                                                _ = v103
                                                                                            } else {
                                                                                                v2 += 58
                                                                                                var v104 uint32 = 1
                                                                                                fmt.Println(v2)
                                                                                                _ = v104
                                                                                            }
                                                                                            var v105 int32 = 1
                                                                                            var v106 int16 = 995
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    default:
                                                                                        if (-2.25 == -2.25) {
                                                                                            fmt.Println(fi10, v2, fi10)
                                                                                            v1 += (v1 * (^(v4)))
                                                                                        } else {
                                                                                            v1 -= 1
                                                                                            var v107 uint16 = 360
                                                                                            if (false || (uint64(v2) >= 577)) {
                                                                                                if (!(0.1 > -2.25)) {
                                                                                                    _ = len(v3)
                                                                                                    if true {
                                                                                                        var v108 uint = 4294967295
                                                                                                        if false {
                                                                                                            v107 -= 870
                                                                                                            if (2147483647 > 4294967295) {
                                                                                                                v109 := GenS6{F0: "foo bar", F1: v2}
                                                                                                                v108 = (v108 ^ (v108 & 0))
                                                                                                                v4 += 2
                                                                                                                fmt.Println(v2)
                                                                                                                _ = v109
                                                                                                            } else {
                                                                                                                v110 := GenS3{F0: 32767}
                                                                                                                _ = v110
                                                                                                            }
                                                                                                            v111 := GenS11{F0: v107, F1: v4, F2: (((0.0 > -2.25) || false) || true)}
                                                                                                            var v112 string = "a"
                                                                                                            v112 = ("" + v112)
                                                                                                            _ = v111
                                                                                                            _ = v112
                                                                                                        }
                                                                                                        var v113 bool = true
                                                                                                        _ = v108
                                                                                                        _ = v113
                                                                                                    }
                                                                                                } else {
                                                                                                    v2 = v2
                                                                                                    v3[""] = "hello"
                                                                                                    if true {
                                                                                                        var v114 int8 = 127
                                                                                                        v1 += (^(v1))
                                                                                                        fmt.Println(v4)
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        v2 -= 241
                                                                                                    }
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v115 int16 = -464
                                                                                                var v116 float32 = 3.14159
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            }
                                                                                            _ = v107
                                                                                        }
                                                                                    }
                                                                                    v117 := GenS2{F0: (4294967295 < 915)}
                                                                                    var v118 float64 = (3.14159 + 0.0)
                                                                                    v118 -= v118
                                                                                    var v119 int64 = -9223372036854775808
                                                                                    if false {
                                                                                        var v120 int16 = 243
                                                                                        switch (fi10 % ((fi9) | 1)) {
                                                                                        case 4294967295:
                                                                                            _ = len(v5)
                                                                                        case 42:
                                                                                            v119 += (v119 - (^(v119)))
                                                                                        default:
                                                                                            var v121 int16 = v120
                                                                                            _ = v121
                                                                                        }
                                                                                        if (!false) {
                                                                                            var v122 uint16 = 0
                                                                                            _ = v122
                                                                                        }
                                                                                        _ = v120
                                                                                    } else {
                                                                                        var v123 float32 = 1e10
                                                                                        _ = v123
                                                                                    }
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    _ = v117
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                }
                                                                                fmt.Println(v1)
                                                                                var v124 uint8 = (^(v2))
                                                                                v2 = v124
                                                                                v1 = (v1 - (v4 * v1))
                                                                                _ = v124
                                                                            } else {
                                                                                v125 := GenS7{F0: true, F1: ("a" < "hello"), F2: (^(v2))}
                                                                                v126 := GenS9{F0: (v1 ^ v1)}
                                                                                _ = v125
                                                                                _ = v126
                                                                            }
                                                                        default:
                                                                            var v127 uint8 = (v2 | v2)
                                                                            _ = v127
                                                                        }
                                                                    } else {
                                                                        if ((((("a" + ("foo bar" + ("a" + ("foo bar" + (("a" + "\n\t") + "foo bar"))))) + "a") != "a") || true) || (((fi9 >= (fi10 - 0)) || true) || true)) {
                                                                            var v128 float64 = 3.14159
                                                                            _ = v128
                                                                        }
                                                                        switch (("" + ("\n\t" + "hello")) + "\n\t") {
                                                                        case "foo bar":
                                                                            var v129 int32 = 2147483647
                                                                            _ = v129
                                                                        }
                                                                        if (v2 >= (v2 % (((v2 ^ 108)) | 1))) {
                                                                            var v130 float64 = 1.5
                                                                            switch (fi9 + (-(fi9))) {
                                                                            case 2147483647:
                                                                                var v131 int32 = 96
                                                                                _ = v131
                                                                            case -2147483648:
                                                                                var v132 uint8 = (v2 + uint8(v1))
                                                                                _ = v132
                                                                            default:
                                                                                var v133 bool = (false && true)
                                                                                _ = v133
                                                                            }
                                                                            _ = v130
                                                                        } else {
                                                                            v5 = append(v5, "")
                                                                            var v134 float32 = 1.5
                                                                            if len(v5) > 0 {
                                                                                v5[0] = "hello"
                                                                            }
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                    fmt.Println(v4)
                                                                }
                                                            }
                                                            v2 = (v2 * v2)
                                                            fmt.Println(fi9)
                                                        }
                                                    }
                                                    v1 = 126
                                                    fmt.Println(v4, v4)
                                                }
                                            }
                                            _ = len(v5)
                                            var v135 int8 = 100
                                            v136 := GenS1{F0: 0}
                                            _ = v135
                                            _ = v136
                                        }
                                    } else {
                                        switch fi10 {
                                        case 42:
                                            v137 := GenS2{F0: (false && (false && true))}
                                            _ = v137
                                        case -430:
                                            if (false || true) {
                                                var v138 float32 = float32(float64(1e10))
                                                _ = v138
                                            }
                                        default:
                                            fmt.Println(v4, fi10)
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            if (1.5 <= 3.14159) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v139 := v3["hello"]
                                                    _ = v139
                                                case "\n\t":
                                                    switch "a" {
                                                    case "foo bar":
                                                        v4 = 794
                                                    case "a":
                                                        _ = len(v3)
                                                    }
                                                }
                                                fmt.Println(v4, v1, fi9)
                                                var v140 float64 = 1.5
                                                _ = v140
                                            } else {
                                                var v141 float32 = 0.1
                                                if ((0.1 <= (float32(0.1) * v141)) || (!(!false))) {
                                                    if false {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    } else {
                                                        v141 -= v141
                                                    }
                                                    var v142 float32 = v141
                                                    var v143 uint16 = 2
                                                    _ = v142
                                                    _ = v143
                                                } else {
                                                    v2 -= v2
                                                }
                                                _ = v141
                                            }
                                        }
                                    }
                                    if (true || true) {
                                        v144 := GenS12{F0: (!true)}
                                        if ((!false) && false) {
                                            var v145 uint16 = 65535
                                            _ = v145
                                        }
                                        var v146 int32 = 2147483647
                                        _ = v144
                                        _ = v146
                                    } else {
                                        v1 += v1
                                        fmt.Println(v1, v4, fi9)
                                    }
                                } else {
                                    var v147 string = "\n\t"
                                    _ = v147
                                }
                            }
                            _ = fi10
                        }
                        var v148 uint8 = (v2 ^ 128)
                        _ = fi9
                        _ = v148
                    }
                } else {
                    v2 -= v2
                }
            }
        case "foo bar":
            v149 := GenS6{F0: "\n\t", F1: (^(v2))}
            _ = v149
        }
    case -1:
        v2 += v2
    case 127:
        if true {
            fmt.Println(v1, v4, v2)
            switch 127 {
            case 928:
                for fi150 := 0; fi150 < 14; fi150++ {
                    _ = len(v5)
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            v2 = 253
                        }
                    case "foo bar":
                        v1 -= 0
                    case "":
                        v151 := GenS13{F0: 824}
                        _ = v151
                    default:
                        if false {
                            v152 := GenS14{F0: 2, F1: 42}
                            v4 = v1
                            fmt.Println(v2, v1)
                            _ = v152
                        } else {
                            v2 = (v2 & v2)
                            for fi153 := 0; fi153 < 17; fi153++ {
                                v2 = v2
                                _ = fi153
                            }
                        }
                    }
                    _ = fi150
                }
            default:
                v1 = v4
            }
        }
    }
    var v154 uint8 = v2
    fmt.Println(v1, v4)
    var v155 uint16 = 32767
    v4 += v1
    if (-2.25 == 0.1) {
        for fi156 := 0; fi156 < 8; fi156++ {
            v2 += 127
            _ = fi156
        }
        for fi157 := 0; fi157 < 13; fi157++ {
            v2 -= (v2 >> ((0) & 63))
            for fi158 := 0; fi158 < 19; fi158++ {
                v4 = 0
                _ = fi158
            }
            switch (("" + ("foo bar" + ("\n\t" + ("foo bar" + "hello")))) + ("hello" + "\n\t")) {
            case "\n\t":
                var v159 uint = 4294967295
                _ = v159
            case "hello":
                for fi160 := 0; fi160 < 8; fi160++ {
                    switch "a" {
                    case "":
                        var v161 bool = (674 < 1)
                        if true {
                            if len(v5) > 0 {
                                v5[0] = ""
                            }
                            v155 += 1
                        }
                        switch ((((("hello" + "") + "hello") + ("" + "a")) + (("" + "a") + "a")) + (("" + "\n\t") + "hello")) {
                        case "a":
                            var v162 uint16 = v155
                            _ = v162
                        case "hello":
                            v1 += v1
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            delete(v3, "")
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v154 -= (v2 * uint8(v1))
                            v1 -= 4294967295
                        case "\n\t":
                            v155 = 255
                        default:
                            var v163 uint32 = 2147483647
                            if v161 {
                                var v164 int8 = int8(v154)
                                _ = v164
                            }
                            v163 = v163
                            _ = v163
                        }
                        v165 := GenS15{F0: -32768, F1: "\n\t"}
                        _ = v161
                        _ = v165
                    }
                    var v166 bool = true
                    var v167 int8 = 0
                    _ = fi160
                    _ = v166
                    _ = v167
                }
            default:
                if (true && (false || false)) {
                    v154 += (v154 << ((uint(v155)) & 63))
                }
            }
            _ = fi157
        }
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v154
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
