package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 uint32; F2 int32 }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 int8; F1 string; F2 int32 }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 uint32 }

func main() {
    var v1 int8 = -74
    var v2 int8 = v1
    var v3 uint = 42
    v4 := []int{-9223372036854775808, 32767}
    v5 := make(map[string]bool)
    v6 := make(map[string]string)
    var v7 int = 42
    fmt.Println(v2, v2, v2)
    v2 += (v2 * v2)
    var v8 bool = ((!(-153 == 1)) && false)
    var v9 int32 = 1
    v10 := GenS1{F0: v1, F1: 65535}
    if ((v8 || v8) && true) {
        var v11 int64 = 0
        if (1.5 == ((1e10 + -2.25) * float64(0.1))) {
            fmt.Println(v11, v9, v1)
            var v12 float32 = ((0.1 / (1.5 / float32((0.1 / 1.5)))) + 1.5)
            _ = len(v4)
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (fi13 | fi13) {
            case 127:
                for fi14 := 0; fi14 < 2; fi14++ {
                    if ((float32(0.1) / (((1.5 / (-2.25 + -2.25)) * 1.5) * 1e10)) != (0.0 / 1.5)) {
                        fmt.Println(fi14, fi13, v9)
                        if v8 {
                            v15 := GenS1{F0: int8(v9), F1: 600}
                            _ = v15
                        }
                        v3 = v3
                        if (((true && (!(v8 && true))) && (!(v2 > (v1 % ((-2) | 1))))) && v8) {
                            var v16 uint8 = 213
                            _ = v16
                        } else {
                            v8 = v8
                            var v17 float64 = float64((((3.14159 - 3.14159) + (-2.25 * (0.0 + ((0.1 * 0.1) * 0.1)))) / 0.1))
                            var v18 int64 = v11
                            v8 = (2 <= 484)
                            var v19 uint64 = 4294967295
                            v17 = v17
                            var v20 uint32 = 359
                            v17 = 1.5
                            v21 := GenS1{F0: v2, F1: 65535}
                            v22 := v6[(("" + "hello") + "foo bar")]
                            var v23 float64 = 1.5
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    } else {
                        if true {
                            if len(v4) > 0 {
                                v4[0] = (v7 | (fi14 % (((fi14 * fi14)) | 1)))
                            }
                            switch "" {
                            case "foo bar":
                                var v24 uint32 = 52
                                _ = v24
                            case "hello":
                                var v25 float32 = -2.25
                                _ = v25
                            default:
                                v8 = (-2.25 <= ((float32((float64(1.5) + float64((0.0 * 1.5)))) * -2.25) - -2.25))
                            }
                        }
                        var v26 int8 = v2
                        _ = len(v4)
                        _ = len(v4)
                        _ = v26
                    }
                    v27 := GenS2{F0: 128}
                    delete(v5, "hello")
                    _ = fi14
                    _ = v27
                }
            }
            fmt.Println(v2, v8, v3)
            _ = fi13
        }
        _ = v11
    } else {
        if true {
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 uint = v3
                var v30 float32 = -2.25
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            v4 = append(v4, 4294967295)
            v1 = (^(v1))
        }
    }
    v7 -= 32767
    v8 = v8
    switch v7 {
    case 4294967295:
        switch "\n\t" {
        case "\n\t":
            if (1 > v3) {
                var v31 uint8 = 141
                for fi32 := 0; fi32 < 7; fi32++ {
                    for fi33 := 0; fi33 < 9; fi33++ {
                        var v34 int16 = -118
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS3{F0: 18446744073709551615}
                    var v36 bool = true
                    _ = fi32
                    _ = v35
                    _ = v36
                }
                _ = v31
            }
        case "a":
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 int8 = v2
                for fi39 := 0; fi39 < 11; fi39++ {
                    switch "a" {
                    case "a":
                        v8 = (v8 || false)
                    case "foo bar":
                        v38 = (v1 | (v38 & v2))
                    }
                    if v8 {
                        if v8 {
                            v8 = v8
                            var v40 uint32 = 663
                            v1 += (-(v2))
                            v5[("\n\t" + ("a" + ("foo bar" + (("\n\t" + ("foo bar" + ("a" + ("a" + ("hello" + ("\n\t" + ("foo bar" + ("a" + "a")))))))) + (("foo bar" + ("" + ("foo bar" + "\n\t"))) + "foo bar")))))] = v8
                            _ = v40
                        }
                    } else {
                        var v41 int64 = -9223372036854775808
                        if v8 {
                            var v42 uint32 = 4294967295
                            _ = v42
                        }
                        switch (("\n\t" + "") + "foo bar") {
                        case "foo bar":
                            var v43 int16 = 188
                            var v44 uint64 = 4294967295
                            _ = v43
                            _ = v44
                        case "":
                            v41 += (v41 | (v41 >> (((^(v3))) & 63)))
                        case "\n\t":
                            if ((!v8) && v8) {
                                v45 := GenS4{F0: "hello"}
                                v9 -= 42
                                _ = v45
                            } else {
                                v7 += (fi39 + (fi39 << ((1) & 63)))
                                var v46 uint32 = 23
                                switch "hello" {
                                case "foo bar":
                                    if v8 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if false {
                                            v1 -= (^(v38))
                                            var v47 uint16 = 14
                                            var v48 int64 = v41
                                            v6["\n\t"] = "hello"
                                            _ = v47
                                            _ = v48
                                        }
                                        delete(v5, (("" + "\n\t") + ""))
                                        switch ("a" + "hello") {
                                        case "a":
                                            var v49 string = "foo bar"
                                            _ = v49
                                        case "":
                                            var v50 int32 = int32(v3)
                                            if false {
                                                if v8 {
                                                    switch ((((("\n\t" + "foo bar") + "foo bar") + "\n\t") + "foo bar") + "\n\t") {
                                                    case "hello":
                                                        switch ("a" + "") {
                                                        case "\n\t":
                                                            _ = len(v6)
                                                        case "hello":
                                                            var v51 uint8 = 255
                                                            _ = v51
                                                        default:
                                                            v8 = (v8 || v8)
                                                            if v8 {
                                                                var v52 uint32 = v46
                                                                var v53 int16 = 32767
                                                                _ = v52
                                                                _ = v53
                                                            } else {
                                                                var v54 float64 = ((3.14159 - float64((1.5 + 1e10))) / float64(-2.25))
                                                                var v55 uint8 = 1
                                                                switch fi37 {
                                                                case 806:
                                                                    v56 := GenS3{F0: 2}
                                                                    _ = v56
                                                                }
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            _ = len(v6)
                                                        }
                                                    case "a":
                                                        if (!v8) {
                                                            if true {
                                                                if (!v8) {
                                                                    v57 := GenS4{F0: ""}
                                                                    v3 = 0
                                                                    v58 := GenS5{F0: 9223372036854775807}
                                                                    var v59 int64 = v41
                                                                    v41 = int64(v46)
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                } else {
                                                                    var v60 float64 = 3.14159
                                                                    v61 := GenS3{F0: 0}
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                            } else {
                                                                v62 := GenS6{F0: 1, F1: (-(v50))}
                                                                if (!(v8 || (42 < 177))) {
                                                                    v8 = v8
                                                                    if v8 {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        fmt.Println(v2)
                                                                    }
                                                                    var v63 uint32 = (v46 / ((v46) | 1))
                                                                    _ = v63
                                                                } else {
                                                                    v2 = (v38 % ((v1) | 1))
                                                                    v64 := GenS7{F0: 25}
                                                                    _ = v64
                                                                }
                                                                if false {
                                                                    if (v50 < int32(fi37)) {
                                                                        if (1 != v3) {
                                                                            v4 = append(v4, fi39)
                                                                            switch -618 {
                                                                            case 633:
                                                                                v65 := GenS8{F0: 255, F1: (v46 + 42), F2: 32767}
                                                                                _ = v65
                                                                            case -2147483648:
                                                                                fmt.Println(v38, v7, fi39)
                                                                            default:
                                                                                var v66 int64 = (v41 & -965)
                                                                                _ = v66
                                                                            }
                                                                        } else {
                                                                            switch fi37 {
                                                                            case 127:
                                                                                v7 += v7
                                                                            case 42:
                                                                                if ((true || (float64((0.1 - 3.14159)) >= (3.14159 * 0.1))) && v8) {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v41, v8)
                                                                                    }
                                                                                    if ((0.0 * float64(0.0)) < -2.25) {
                                                                                        switch (v7 & (fi37 % ((int(v41)) | 1))) {
                                                                                        case 32768:
                                                                                            fmt.Println(v2, fi37)
                                                                                        case -2147483648:
                                                                                            v38 -= 1
                                                                                        default:
                                                                                            v9 = v9
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            v4 = append(v4, (^(fi39)))
                                                                                            v2 = v1
                                                                                            var v67 float64 = ((float64((1.5 * float32((0.0 + 0.0)))) * 1.5) / (-2.25 + ((0.0 + 0.1) * (0.0 - ((float64(float32(0.0)) + 3.14159) * float64(1.5))))))
                                                                                            fmt.Println(v46, v2, v7)
                                                                                            _ = v67
                                                                                        } else {
                                                                                            v41 = (v41 * v41)
                                                                                            var v68 int = (fi39 & (v7 | 42))
                                                                                            switch (fi37 | (fi39 / ((v7) | 1))) {
                                                                                            case 2147483647:
                                                                                                if (!false) {
                                                                                                    if (!(1e10 <= 3.14159)) {
                                                                                                        v69 := GenS9{F0: 2}
                                                                                                        _ = v69
                                                                                                    } else {
                                                                                                        fmt.Println(v38, fi39)
                                                                                                        v1 -= (v1 ^ (v38 * v2))
                                                                                                        v1 = -128
                                                                                                    }
                                                                                                    var v70 string = ""
                                                                                                    var v71 int = v68
                                                                                                    _ = v70
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    _ = len(v4)
                                                                                                    fmt.Println(v46, v9, v50)
                                                                                                }
                                                                                            }
                                                                                            v1 += (v38 + v38)
                                                                                            _ = v68
                                                                                        }
                                                                                        v7 = v7
                                                                                    }
                                                                                    fmt.Println(fi39, v3)
                                                                                    var v72 int16 = 1
                                                                                    _ = v72
                                                                                } else {
                                                                                    v8 = v8
                                                                                }
                                                                            case 128:
                                                                                var v73 int = fi39
                                                                                _ = v73
                                                                            default:
                                                                                v50 += (v50 ^ 1)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            fmt.Println(fi39, v8, v3)
                                                                            var v74 bool = v8
                                                                            _ = v74
                                                                        } else {
                                                                            var v75 uint = 605
                                                                            _ = v75
                                                                        }
                                                                        var v76 int32 = v50
                                                                        _ = v76
                                                                    } else {
                                                                        if v8 {
                                                                            var v77 float64 = -2.25
                                                                            v7 = fi37
                                                                            fmt.Println(v9, v8, v38)
                                                                            if false {
                                                                                var v78 int = fi39
                                                                                var v79 float32 = 1.5
                                                                                switch "hello" {
                                                                                case "":
                                                                                    v80 := GenS3{F0: 513}
                                                                                    _ = v80
                                                                                default:
                                                                                    if ((!v8) && v8) {
                                                                                        fmt.Println(v38, v79, v41)
                                                                                        var v81 int = v7
                                                                                        fmt.Println(v9, v78)
                                                                                        _ = v81
                                                                                    }
                                                                                }
                                                                                _ = v78
                                                                                _ = v79
                                                                            } else {
                                                                                var v82 uint32 = v46
                                                                                _ = v82
                                                                            }
                                                                            v2 -= v1
                                                                            v9 += (-(v9))
                                                                            fmt.Println(v3, v3)
                                                                            v7 += (fi37 | (fi39 % (((-(fi39))) | 1)))
                                                                            var v83 int64 = v41
                                                                            var v84 int8 = (v2 << ((1) & 63))
                                                                            _ = v77
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            v1 += (v2 << ((v3) & 63))
                                                                        }
                                                                    }
                                                                    var v85 float64 = (((3.14159 + 0.0) / 1e10) + -2.25)
                                                                    _ = v85
                                                                } else {
                                                                    v3 += (^(v3))
                                                                    v8 = v8
                                                                }
                                                                if (!((v46 ^ v46) < 1)) {
                                                                    v86 := GenS4{F0: "foo bar"}
                                                                    var v87 bool = v8
                                                                    var v88 string = ("hello" + (("\n\t" + ("hello" + "foo bar")) + "hello"))
                                                                    _ = v86
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                                _ = v62
                                                            }
                                                            switch (("hello" + "hello") + ("hello" + "")) {
                                                            case "\n\t":
                                                                v8 = (v8 && false)
                                                            default:
                                                                if ("\n\t" >= "\n\t") {
                                                                    v7 = int(v41)
                                                                }
                                                            }
                                                            v89 := GenS10{F0: "a", F1: v46, F2: (v9 * (v9 + (v9 ^ (v50 & v9))))}
                                                            _ = v89
                                                        }
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    fmt.Println(v46, v9)
                                                    if (true || v8) {
                                                        v90 := GenS11{F0: fi37, F1: ""}
                                                        _ = v90
                                                    }
                                                }
                                            } else {
                                                v1 += v1
                                                var v91 string = "hello"
                                                var v92 int16 = int16(fi39)
                                                v46 += v46
                                                switch (v91 + "") {
                                                case "hello":
                                                    var v93 int64 = int64(v92)
                                                    _ = v93
                                                case "\n\t":
                                                    var v94 bool = false
                                                    _ = v94
                                                case "":
                                                    var v95 string = v91
                                                    _ = v95
                                                }
                                                _ = v91
                                                _ = v92
                                            }
                                            _ = v50
                                        case "foo bar":
                                            var v96 uint8 = 255
                                            _ = v96
                                        default:
                                            v97 := GenS12{F0: 26, F1: "a", F2: (v9 + v9)}
                                            _ = v97
                                        }
                                        fmt.Println(v46)
                                        if true {
                                            switch (fi39 - (fi37 | -1)) {
                                            case 9223372036854775807:
                                                if v8 {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v41 += 42
                                                } else {
                                                    fmt.Println(v1, v9, v1)
                                                }
                                            }
                                            if false {
                                                _ = len(v5)
                                            }
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            switch 127 {
                                            case 32767:
                                                var v98 int16 = 1000
                                                _ = v98
                                            default:
                                                v99 := GenS10{F0: "", F1: (v46 / (((^(v46))) | 1)), F2: 0}
                                                _ = v99
                                            }
                                        }
                                        if (626 > 354) {
                                            var v100 uint16 = 0
                                            _ = v100
                                        } else {
                                            var v101 int16 = -555
                                            _ = v101
                                        }
                                    }
                                    var v102 float64 = (3.14159 / 1.0)
                                    if v8 {
                                        if (("foo bar" + "foo bar") != ("a" + "")) {
                                            if false {
                                                _ = len(v5)
                                            } else {
                                                var v103 uint32 = v46
                                                fmt.Println(v7)
                                                v46 = 65535
                                                _ = v103
                                            }
                                            v8 = ((((v2 != int8(v7)) || v8) && (("hello" + ("foo bar" + ("foo bar" + "a"))) == "")) || v8)
                                            var v104 uint = 255
                                            _ = v104
                                        }
                                        _ = len(v4)
                                        _ = len(v5)
                                    }
                                    _ = v102
                                case "hello":
                                    if v8 {
                                        v9 += (^(v9))
                                        v4 = append(v4, v7)
                                        var v105 float64 = (float64((0.0 + float32(1e10))) * 1e10)
                                        var v106 int8 = 1
                                        _ = v105
                                        _ = v106
                                    }
                                    v107 := GenS7{F0: v9}
                                    _ = v107
                                }
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v8, v46, v1)
                                case "hello":
                                    var v108 uint = v3
                                    _ = v108
                                case "":
                                    v3 = (^(v3))
                                }
                                _ = v46
                            }
                        }
                        _ = v41
                    }
                    _ = fi39
                }
                var v109 uint32 = 268
                _ = fi37
                _ = v109
                _ = v38
            }
        case "hello":
            fmt.Println(v7)
        }
    case 42:
        if ((v8 && false) && v8) {
            fmt.Println(v9)
        } else {
            var v110 uint32 = 484
            v111 := GenS2{F0: 255}
            _ = v110
            _ = v111
        }
    }
    v112 := GenS12{F0: (v1 - (v2 - v2)), F1: "hello", F2: 575}
    switch "a" {
    case "hello":
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "foo bar":
        v9 = (v9 ^ (v9 / (((v9 * -875)) | 1)))
    case "a":
        var v113 int16 = 42
        _ = v113
    default:
        if true {
            var v114 uint64 = 42
            _ = v114
        } else {
            switch ("hello" + ("a" + ("\n\t" + "\n\t"))) {
            case "hello":
                for fi115 := 0; fi115 < 12; fi115++ {
                    v3 += (v3 - v3)
                    for fi116 := 0; fi116 < 18; fi116++ {
                        v117 := GenS13{F0: (^(v3)), F1: 65535}
                        _ = fi116
                        _ = v117
                    }
                    _ = fi115
                }
            case "\n\t":
                var v118 int = (-(v7))
                _ = v118
            }
            if true {
                for fi119 := 0; fi119 < 17; fi119++ {
                    v7 += (-(fi119))
                    _ = fi119
                }
                var v120 uint = v3
                _ = v120
            } else {
                v121 := GenS8{F0: 202, F1: 172, F2: 1}
                var v122 float64 = (0.0 - float64(0.0))
                _ = v121
                _ = v122
            }
            var v123 uint8 = 83
            _ = v123
        }
    }
    var v124 uint32 = 42
    for fi125 := 0; fi125 < 17; fi125++ {
        switch v7 {
        case 42:
            for fi126 := 0; fi126 < 11; fi126++ {
                switch (("hello" + "") + "foo bar") {
                case "hello":
                    fmt.Println(fi125)
                case "\n\t":
                    v2 -= v1
                case "":
                    v127 := v6["hello"]
                    _ = v127
                }
                var v128 uint8 = 0
                v3 = (v3 + 255)
                _ = fi126
                _ = v128
            }
        case -2147483648:
            var v129 int16 = -1
            _ = v129
        }
        v124 -= v124
        _ = fi125
    }
    var v130 int32 = 442
    if ((!(("hello" + ("foo bar" + (((("hello" + "a") + "foo bar") + ("" + "")) + "\n\t"))) <= "")) || v8) {
        v131 := GenS14{F0: (v124 * (v124 >> (((v3 % ((65535) | 1))) & 63)))}
        _ = v131
    } else {
        v124 += 4294967295
        var v132 uint = v3
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v112
    _ = v124
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
