package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 string; F1 int; F2 uint16 }
type GenS12 struct { F0 int32; F1 int; F2 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int }
func genFn0(p0 int16, p1 int16, p2 bool) int32 {
    switch ("a" + "hello") {
    case "":
        p1 += (p0 / ((p1) | 1))
    case "foo bar":
        v7 := GenS1{F0: "a", F1: 0}
        _ = v7
    }
    switch (("hello" + "a") + ("hello" + "hello")) {
    case "":
        switch 991 {
        case 127:
            v8 := []int{128, -366, 42, 127}
            _ = v8
        case 42:
            fmt.Println(p1, p0)
        }
    case "\n\t":
        if (42 < -1) {
            var v9 int8 = 42
            var v10 float32 = 0.1
            fmt.Println(v9, p0)
            _ = v10
            _ = v9
        } else {
            if ((p2 && p2) || p2) {
                if (255 > 0) {
                    v11 := []int{32767, 32767}
                    p0 = p1
                    for fi12 := 0; fi12 < 12; fi12++ {
                        switch "" {
                        case "hello":
                            var v13 uint16 = 42
                            _ = v13
                        case "":
                            p0 -= (p1 % ((591) | 1))
                        case "a":
                            var v14 int = (fi12 << ((0) & 63))
                            _ = v14
                        }
                        v15 := GenS2{F0: 42, F1: -235, F2: 247}
                        p0 += p1
                        _ = fi12
                        _ = v15
                    }
                    _ = v11
                } else {
                    v16 := []int{}
                    v16 = append(v16, 128)
                    _ = v16
                }
            } else {
                fmt.Println(p1, p2, p0)
            }
        }
    }
    v17 := []int{32768, 32768, -871, 1}
    for fi18 := 0; fi18 < 2; fi18++ {
        p0 = (p1 | -423)
        var v19 int16 = p0
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    v1 := []string{(("hello" + "foo bar") + ("foo bar" + "foo bar")), "a", "a", (("\n\t" + "\n\t") + ("a" + ""))}
    var v2 uint8 = 128
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    var v5 int64 = -617
    var v6 int = -9223372036854775808
    var v20 int64 = v5
    v2 = v2
    for fi21 := 0; fi21 < 16; fi21++ {
        if (32 != int8(v5)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 = 46
            var v22 int8 = 1
            _ = v22
        } else {
            if (false || (127 >= -128)) {
                if ((v2 ^ v2) <= 128) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v20 -= (v20 + v5)
                        v6 = -421
                        if (false && ((0.1 - 0.1) >= -2.25)) {
                            switch -2147483648 {
                            case 42:
                                switch ("foo bar" + ("" + "hello")) {
                                case "hello":
                                    v24 := GenS3{F0: 954, F1: 42}
                                    _ = v24
                                case "a":
                                    var v25 uint32 = 853
                                    _ = v25
                                default:
                                    v5 += -9223372036854775808
                                }
                            case 127:
                                var v26 float64 = (((((-2.25 * 1.5) - -2.25) * 1.5) - 0.1) - -2.25)
                                _ = v26
                            }
                            fmt.Println(fi21)
                            v27 := GenS2{F0: 0, F1: (fi23 << ((4294967295) & 63)), F2: 65535}
                            _ = v27
                        } else {
                            if false {
                                v3[((("hello" + "a") + "\n\t") + "foo bar")] = 549
                                fmt.Println(fi21, v6)
                            }
                        }
                        _ = fi23
                    }
                    var v28 int8 = int8(v4)
                    v2 -= 0
                    _ = v28
                } else {
                    v29 := GenS4{F0: ((v5 + (v20 + 42)) <= (v20 * 0))}
                    switch (fi21 * 4294967295) {
                    case -2147483648:
                        if false {
                            var v30 uint8 = v2
                            if (65535 <= 2) {
                                var v31 uint = 695
                                delete(v3, "")
                                _ = v31
                            }
                            v30 += v30
                            _ = v30
                        } else {
                            var v32 float32 = ((-2.25 / 3.14159) * 0.0)
                            v33 := GenS5{F0: 42}
                            _ = v32
                            _ = v33
                        }
                    case 127:
                        for fi34 := 0; fi34 < 14; fi34++ {
                            switch (v6 / (((fi34 % (((-(v6))) | 1))) | 1)) {
                            case 475:
                                v5 -= (v20 & (v20 & (v20 + (v5 / ((v5) | 1)))))
                            case 127:
                                v3["\n\t"] = 42
                            case 0:
                                fmt.Println(fi34, fi21, v2)
                            }
                            v4 += genFn0(506, 1000, (!true))
                            fmt.Println(v4, v6)
                            _ = fi34
                        }
                    case -528:
                        var v35 int = (fi21 + (fi21 * (fi21 ^ v6)))
                        _ = v35
                    }
                    v36 := GenS4{F0: ("hello" != "a")}
                    _ = v29
                    _ = v36
                }
            } else {
                var v37 int = v6
                _ = v37
            }
        }
        v20 += v5
        _ = fi21
    }
    v38 := GenS2{F0: 42, F1: (v6 * 9223372036854775807), F2: 1}
    _ = len(v3)
    v39 := GenS6{F0: v6, F1: 2147483647}
    for fi40 := 0; fi40 < 20; fi40++ {
        var v41 uint64 = 42
        v5 = v5
        _ = fi40
        _ = v41
    }
    _ = len(v3)
    for fi42 := 0; fi42 < 11; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            if (0.1 > 3.14159) {
                v4 += -1
                if (2147483648 >= 2147483647) {
                    var v44 uint = 42
                    _ = v44
                } else {
                    if true {
                        var v45 string = "hello"
                        _ = v45
                    }
                    var v46 int32 = (v4 | v4)
                    var v47 uint8 = 127
                    _ = v46
                    _ = v47
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v48 uint64 = 42
                _ = v48
            } else {
                v20 += int64(v2)
            }
            var v49 uint64 = 0
            if ((false && true) || true) {
                fmt.Println(fi42)
            }
            v2 = (v2 * (v2 << ((4294967295) & 63)))
            _ = fi43
            _ = v49
        }
        v3[("\n\t" + ("" + "hello"))] = 256
        _ = fi42
    }
    switch v6 {
    case 1:
        v50 := v3["a"]
        _ = v50
    case 9223372036854775807:
        if (false || (!((false && false) || (323 < 0)))) {
            for fi51 := 0; fi51 < 13; fi51++ {
                for fi52 := 0; fi52 < 8; fi52++ {
                    v53 := GenS4{F0: false}
                    switch -9223372036854775808 {
                    case -796:
                        _ = len(v3)
                    case 127:
                        var v54 uint = 56
                        _ = v54
                    case 128:
                        switch (fi51 + fi51) {
                        case 127:
                            if true {
                                var v55 int64 = v20
                                if (v20 >= 42) {
                                    var v56 int16 = -32768
                                    if true {
                                        v4 = (v4 * 0)
                                    }
                                    _ = v56
                                } else {
                                    _ = len(v1)
                                }
                                v2 -= v2
                                v57 := GenS7{F0: 32767}
                                _ = v55
                                _ = v57
                            }
                            var v58 int64 = v5
                            fmt.Println(v2)
                            _ = v58
                        default:
                            v1 = append(v1, ("a" + ""))
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                v1 = append(v1, "hello")
                switch v6 {
                case 32767:
                    for fi59 := 0; fi59 < 9; fi59++ {
                        switch ("hello" + ("" + "a")) {
                        case "foo bar":
                            _ = len(v1)
                        case "\n\t":
                            if (1e10 < (1.5 + 3.14159)) {
                                v3["\n\t"] = 2147483648
                                v3["a"] = 42
                                v4 -= v4
                            } else {
                                v5 = (v20 % (((v5 ^ (-(v20)))) | 1))
                                if ((v6 / (((fi51 / (((fi51 * fi51)) | 1))) | 1)) == -2147483648) {
                                    v5 += 9223372036854775807
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (((((("a" + "\n\t") + "") + "") + ("hello" + ("a" + ("" + "hello")))) + "a") + "a")
                                    }
                                }
                            }
                        }
                        v2 = (v2 * (v2 << ((255) & 63)))
                        if true {
                            delete(v3, "foo bar")
                        }
                        _ = fi59
                    }
                case -2147483648:
                    var v60 bool = true
                    _ = v60
                }
                _ = fi51
            }
            for fi61 := 0; fi61 < 12; fi61++ {
                fmt.Println(v6, v20)
                if true {
                    var v62 int16 = -363
                    v1 = append(v1, "foo bar")
                    v2 += (v2 / ((132) | 1))
                    _ = v62
                } else {
                    var v63 uint32 = 0
                    for fi64 := 0; fi64 < 9; fi64++ {
                        if false {
                            var v65 int8 = 1
                            v63 = 835
                            _ = v65
                        } else {
                            v20 -= 9223372036854775807
                            var v66 uint = 4294967295
                            if ((((("\n\t" + ("foo bar" + ("\n\t" + ""))) + ("foo bar" + "hello")) + "\n\t") + "a") != "hello") {
                                v1 = append(v1, "hello")
                            } else {
                                v6 = fi61
                                v67 := v3["a"]
                                if true {
                                    var v68 float64 = (0.1 * float64(-2.25))
                                    _ = v68
                                }
                                v69 := GenS8{F0: (v2 >> ((0) & 63))}
                                _ = v67
                                _ = v69
                            }
                            var v70 uint32 = v63
                            if (174 > (v2 << ((18446744073709551615) & 63))) {
                                v71 := GenS9{F0: (v4 >> ((955) & 63)), F1: v2}
                                switch ("a" + "\n\t") {
                                case "hello":
                                    _ = len(v1)
                                case "foo bar":
                                    var v72 int = (v6 << ((v66) & 63))
                                    _ = v72
                                case "a":
                                    var v73 int16 = int16(v66)
                                    _ = v73
                                }
                                v66 = v66
                                if false {
                                    _ = len(v1)
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        if true {
                                            v70 = (v70 | 0)
                                        } else {
                                            v74 := v3["foo bar"]
                                            var v75 int = -329
                                            v70 += 65535
                                            _ = v74
                                            _ = v75
                                        }
                                    case "\n\t":
                                        if true {
                                            v70 = (^(v63))
                                        }
                                    case "":
                                        v76 := GenS2{F0: v70, F1: 128, F2: v66}
                                        _ = v76
                                    }
                                    v77 := GenS2{F0: v63, F1: (v6 >> (((^(v66))) & 63)), F2: 566}
                                    _ = v77
                                }
                                v78 := GenS3{F0: 0, F1: v2}
                                v20 -= (v5 + (v20 >> ((v66) & 63)))
                                if (!false) {
                                    v1 = append(v1, "foo bar")
                                    v79 := GenS10{F0: "a"}
                                    switch "\n\t" {
                                    case "hello":
                                        var v80 uint64 = 0
                                        _ = v80
                                    case "foo bar":
                                        var v81 uint8 = v2
                                        _ = v81
                                    case "\n\t":
                                        var v82 int32 = (v4 | (v4 + (v4 % ((-927) | 1))))
                                        _ = v82
                                    }
                                    switch (v6 - (fi64 & (v6 * (-(fi61))))) {
                                    case 451:
                                        var v83 uint32 = 42
                                        _ = v83
                                    case 297:
                                        var v84 int16 = -32768
                                        _ = v84
                                    case -393:
                                        if true {
                                            switch v6 {
                                            case 4294967295:
                                                v1 = append(v1, "foo bar")
                                            case 128:
                                                var v85 string = ("hello" + "\n\t")
                                                var v86 int32 = v4
                                                v63 -= v63
                                                _ = v85
                                                _ = v86
                                            case -9223372036854775808:
                                                if ((v20 >= (v5 - v5)) || (!((v66 % ((v66) | 1)) < 18446744073709551615))) {
                                                    fmt.Println(fi61)
                                                    switch "" {
                                                    case "a":
                                                        v6 = (fi61 << ((255) & 63))
                                                    }
                                                    if ((v2 * (v2 + 255)) == v2) {
                                                        var v87 uint8 = 42
                                                        v5 -= v5
                                                        var v88 uint64 = 1
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                } else {
                                                    v1 = append(v1, "foo bar")
                                                    v63 = v63
                                                }
                                            }
                                            var v89 int16 = 1000
                                            _ = v89
                                        } else {
                                            v5 = v20
                                        }
                                    }
                                    _ = v79
                                }
                                _ = v71
                                _ = v78
                            } else {
                                if false {
                                    v63 = (v63 - (v70 | 2))
                                    switch int(v66) {
                                    case 32767:
                                        v70 = (v63 & (v63 % ((v63) | 1)))
                                    default:
                                        v70 = uint32(v66)
                                    }
                                    var v90 int16 = -246
                                    _ = v90
                                } else {
                                    var v91 int = (fi64 + (fi61 & v6))
                                    switch (fi61 % (((fi61 - v6)) | 1)) {
                                    case 2147483647:
                                        if ("\n\t" < "foo bar") {
                                            var v92 string = ("\n\t" + "a")
                                            _ = v92
                                        } else {
                                            v93 := GenS11{F0: "foo bar", F1: (v91 * 1), F2: 336}
                                            _ = v93
                                        }
                                        var v94 int64 = (v5 - -799)
                                        if (!true) {
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            var v95 uint8 = v2
                                            _ = v95
                                        } else {
                                            v66 += v66
                                            var v96 bool = (!(!false))
                                            _ = v96
                                        }
                                        var v97 int = (fi64 & v6)
                                        _ = v94
                                        _ = v97
                                    case 42:
                                        var v98 uint8 = (^(v2))
                                        _ = v98
                                    }
                                    _ = v91
                                }
                                if (!(1 >= v5)) {
                                    fmt.Println(v2)
                                    fmt.Println(v63, v63)
                                    v99 := GenS12{F0: v4, F1: v6, F2: v66}
                                    _ = v99
                                }
                            }
                            _ = v66
                            _ = v70
                        }
                        _ = fi64
                    }
                    _ = v63
                }
                _ = fi61
            }
            if false {
                fmt.Println(v5, v20, v6)
                v6 = v6
                v20 = 0
            }
        }
    case 660:
        fmt.Println(v2, v5)
    default:
        v2 += v2
    }
    if (1 >= 18446744073709551615) {
        var v100 float64 = 3.14159
        var v101 uint16 = 32767
        v20 += 821
        _ = v100
        _ = v101
    } else {
        var v102 bool = false
        for fi103 := 0; fi103 < 8; fi103++ {
            var v104 uint = 18446744073709551615
            _ = fi103
            _ = v104
        }
        if v102 {
            if v102 {
                v105 := GenS13{F0: v6}
                _ = v105
            }
            _ = len(v1)
        } else {
            v20 += (^(v5))
        }
        _ = v102
    }
    if (42 < 0) {
        var v106 uint64 = 154
        var v107 float64 = 0.1
        v106 += v106
        _ = v106
        _ = v107
    } else {
        for fi108 := 0; fi108 < 11; fi108++ {
            for fi109 := 0; fi109 < 8; fi109++ {
                v20 -= 25
                v110 := GenS14{F0: fi109}
                _ = fi109
                _ = v110
            }
            _ = fi108
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
