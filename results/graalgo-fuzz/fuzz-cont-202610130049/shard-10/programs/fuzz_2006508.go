package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool; F1 int8; F2 int8 }
func genFn0(p0 int64, p1 bool) int8 {
    p1 = p1
    p0 += p0
    if p1 {
        var v5 uint8 = 68
        var v6 float64 = (((((1e10 * 3.14159) - 1.5) - 3.14159) / (0.0 - ((1e10 - 0.0) + (1e10 / -2.25)))) * 0.1)
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi7)
            var v8 int16 = 1000
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint16 = 32767
                if true {
                    var v12 int32 = 42
                    _ = v12
                }
                v13 := GenS1{F0: -259, F1: "", F2: (-(fi9))}
                if (p0 > 9223372036854775807) {
                    switch "foo bar" {
                    case "":
                        var v14 float64 = 0.1
                        _ = v14
                    case "hello":
                        var v15 uint16 = v11
                        _ = v15
                    }
                }
                if ((0 != 2) && ((1e10 + (0.0 - 1.5)) != ((0.1 + 1.5) / 3.14159))) {
                    v11 += (v11 & (v11 + v11))
                    v16 := GenS1{F0: 2147483647, F1: "\n\t", F2: (fi9 >> ((714) & 63))}
                    fmt.Println(v11)
                    v17 := []int{fi9, (fi9 | fi9)}
                    fmt.Println(p1)
                    fmt.Println(fi10, p0)
                    var v18 int = fi9
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    fmt.Println(p1)
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(p1, p0)
                    if (p1 && false) {
                        p0 -= p0
                        fmt.Println(p0, fi9, fi10)
                    } else {
                        var v19 uint = 93
                        _ = v19
                    }
                }
                var v20 int8 = -114
                var v21 bool = ((843 > 42) || (!(!true)))
                _ = fi10
                _ = v11
                _ = v13
                _ = v20
                _ = v21
            }
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1(p0 int) bool {
    p0 = (p0 - (p0 | p0))
    var v22 int8 = 100
    v22 = (v22 >> ((65535) & 63))
    var v23 int8 = (v22 * 127)
    _ = p0
    _ = v22
    _ = v23
    return (false && false)
}

func main() {
    var v1 uint = 1
    v2 := []string{"a", ("" + ""), ("foo bar" + "a")}
    v3 := make(map[string]int)
    v4 := []uint32{479, 4294967295, 65535, 65535}
    if true {
        _ = len(v2)
        switch (((("hello" + "a") + ("foo bar" + "hello")) + "hello") + "\n\t") {
        case "":
            fmt.Println(v1)
        case "\n\t":
            v1 += v1
        }
    } else {
        fmt.Println(v1)
        if ((!(189 < 127)) || false) {
            for fi24 := 0; fi24 < 9; fi24++ {
                v2 = append(v2, ("" + ("" + (((("hello" + "hello") + "foo bar") + "a") + ("" + ("\n\t" + "\n\t"))))))
                fmt.Println(v1, fi24)
                v1 = 216
                _ = fi24
            }
            for fi25 := 0; fi25 < 20; fi25++ {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int8 = 100
                    _ = fi26
                    _ = v27
                }
                if true {
                    fmt.Println(fi25, v1)
                    var v28 uint64 = 4294967295
                    for fi29 := 0; fi29 < 17; fi29++ {
                        fmt.Println(fi25, fi25, v1)
                        var v30 int = fi29
                        switch "a" {
                        case "a":
                            fmt.Println(v30)
                        case "":
                            v28 += v28
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
                var v31 uint8 = 42
                _ = fi25
                _ = v31
            }
            switch -9223372036854775808 {
            case 32767:
                fmt.Println(v1)
            }
        } else {
            v1 = (^(v1))
            v1 = (v1 ^ v1)
        }
        delete(v3, "foo bar")
    }
    if (true || (0 != -1)) {
        if (1 != 2147483647) {
            v32 := GenS2{F0: (332 >= -9223372036854775808), F1: 2147483648, F2: 32768}
            if false {
                if false {
                    v1 += 4294967295
                    v33 := GenS3{F0: 1}
                    _ = v33
                } else {
                    v1 += v1
                    var v34 int64 = 0
                    _ = v34
                }
                fmt.Println(v1)
            } else {
                if len(v2) > 0 {
                    v2[0] = ("hello" + "a")
                }
            }
            _ = v32
        } else {
            v1 += (v1 ^ v1)
        }
        switch "foo bar" {
        case "":
            var v35 int16 = 1000
            _ = v35
        case "hello":
            v1 -= v1
        case "a":
            v36 := GenS4{F0: 0, F1: 255}
            _ = v36
        default:
            for fi37 := 0; fi37 < 13; fi37++ {
                if (false || ((1e10 - 3.14159) < ((0.0 + 3.14159) + 1e10))) {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        if (262 != 32767) {
                            if ((float64(-2.25) / ((1.5 - 1e10) + 0.1)) <= float64(-2.25)) {
                                var v39 uint8 = 42
                                if ((("a" + ((("" + "a") + ((("a" + "\n\t") + ((("a" + "\n\t") + ("hello" + ("\n\t" + ("hello" + "\n\t")))) + "hello")) + ((("foo bar" + "") + "") + "foo bar"))) + "")) + (("foo bar" + "") + "a")) == ("foo bar" + "a")) {
                                    v1 -= 65535
                                    v3[(("a" + "a") + "a")] = fi38
                                    v40 := v3[("foo bar" + "hello")]
                                    _ = v40
                                }
                                _ = v39
                            }
                        } else {
                            fmt.Println(v1, fi37, fi37)
                            if true {
                                var v41 int64 = -9223372036854775808
                                v41 = v41
                                fmt.Println(v1)
                                var v42 int64 = (v41 ^ v41)
                                if (!true) {
                                    if false {
                                        var v43 int = fi38
                                        if (4294967295 > 341) {
                                            var v44 float64 = 0.0
                                            v1 = (v1 - (v1 / ((v1) | 1)))
                                            if (!true) {
                                                v45 := GenS3{F0: 21}
                                                _ = v45
                                            } else {
                                                var v46 float64 = v44
                                                var v47 int8 = -128
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = v44
                                        } else {
                                            delete(v3, "foo bar")
                                            var v48 string = "foo bar"
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = len(v4)
                                            var v49 float32 = 1e10
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = v43
                                    }
                                }
                                _ = v41
                                _ = v42
                            } else {
                                fmt.Println(v1, fi38, fi37)
                                v1 = (v1 % ((472) | 1))
                                v1 += 184
                            }
                        }
                        switch ("\n\t" + ("" + "\n\t")) {
                        case "hello":
                            v3["hello"] = (fi38 + fi38)
                        case "foo bar":
                            v2 = append(v2, "a")
                            if (!false) {
                                if false {
                                    var v50 int = (fi37 / ((fi37) | 1))
                                    v3[((("\n\t" + "") + "a") + "")] = (fi38 & (fi37 ^ (fi38 % ((32768) | 1))))
                                    fmt.Println(fi38, fi37)
                                    _ = v50
                                }
                                var v51 int32 = -2147483648
                                v52 := GenS3{F0: -20}
                                _ = v51
                                _ = v52
                            } else {
                                v1 += v1
                                v1 = v1
                            }
                        }
                        if false {
                            v53 := GenS3{F0: 100}
                            _ = v53
                        }
                        _ = fi38
                    }
                    v54 := GenS3{F0: genFn0(-27, (0.0 != 1e10))}
                    if (202 != 693) {
                        fmt.Println(fi37, v1)
                        v55 := v3["hello"]
                        if (false && (118 <= 127)) {
                            if true {
                                if true {
                                    var v56 float64 = 0.0
                                    switch v55 {
                                    case 4294967295:
                                        v57 := GenS2{F0: true, F1: 2147483648, F2: fi37}
                                        _ = v57
                                    case 9223372036854775807:
                                        var v58 int16 = 32767
                                        _ = v58
                                    case 42:
                                        var v59 int64 = 114
                                        _ = v59
                                    }
                                    _ = v56
                                }
                                if true {
                                    var v60 int16 = -1
                                    for fi61 := 0; fi61 < 8; fi61++ {
                                        v60 = (v60 % (((v60 | (v60 * (v60 | (v60 | (v60 - v60)))))) | 1))
                                        v62 := GenS4{F0: 394, F1: 1}
                                        switch "a" {
                                        case "\n\t":
                                            switch 32768 {
                                            case -659:
                                                var v63 int64 = int64(v1)
                                                _ = v63
                                            case 128:
                                                var v64 int64 = 1
                                                _ = v64
                                            case -113:
                                                _ = len(v2)
                                            }
                                        case "hello":
                                            if (!((1.5 + (-2.25 * 1.5)) != (1.5 + -2.25))) {
                                                var v65 float32 = 1.5
                                                v66 := GenS4{F0: 687, F1: 1}
                                                v67 := GenS5{F0: v55, F1: 2147483647}
                                                _ = v65
                                                _ = v66
                                                _ = v67
                                            } else {
                                                var v68 int64 = -9223372036854775808
                                                switch "" {
                                                case "hello":
                                                    switch ("\n\t" + "foo bar") {
                                                    case "hello":
                                                        var v69 uint = (v1 | v1)
                                                        _ = v69
                                                    case "foo bar":
                                                        if (!((!true) || ((true || (721 == 42)) || true))) {
                                                            if false {
                                                                v70 := GenS6{F0: 65535}
                                                                v1 = (v1 << ((v1) & 63))
                                                                var v71 int8 = 100
                                                                switch (("" + "") + ("hello" + ("a" + "foo bar"))) {
                                                                case "a":
                                                                    _ = len(v3)
                                                                default:
                                                                    if (-9223372036854775808 <= 876) {
                                                                        if true {
                                                                            v60 -= 628
                                                                        }
                                                                    } else {
                                                                        v4 = append(v4, 2147483648)
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v72 := GenS3{F0: v71}
                                                                            _ = v72
                                                                        case "\n\t":
                                                                            var v73 uint = (v1 * (^(v1)))
                                                                            switch "" {
                                                                            case "":
                                                                                v74 := GenS7{F0: v73, F1: 100}
                                                                                _ = v74
                                                                            case "hello":
                                                                                if (0.0 >= ((0.0 * (0.0 - float32(3.14159))) - 0.1)) {
                                                                                    v60 = v60
                                                                                    v55 = 0
                                                                                } else {
                                                                                    _ = len(v3)
                                                                                    if (((v68 << (((^(v73))) & 63)) < v68) && (1.5 > float64(-2.25))) {
                                                                                        var v75 bool = true
                                                                                        v68 -= (v68 | v68)
                                                                                        _ = v75
                                                                                    }
                                                                                }
                                                                                var v76 uint = 18446744073709551615
                                                                                var v77 float64 = float64(0.1)
                                                                                var v78 int = v55
                                                                                _ = v76
                                                                                _ = v77
                                                                                _ = v78
                                                                            }
                                                                            _ = v73
                                                                        case "foo bar":
                                                                            if ((v71 & v71) != v71) {
                                                                                v1 += (^(v1))
                                                                                if (float32((0.1 / 0.1)) > 0.0) {
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v60 += v60
                                                                                    case "a":
                                                                                        v55 -= fi37
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    v3[("\n\t" + (("a" + "hello") + ("\n\t" + "\n\t")))] = (fi37 << ((v1) & 63))
                                                                                }
                                                                                v55 -= fi61
                                                                                var v79 bool = false
                                                                                _ = v79
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v55, v71, v55)
                                                                                    fmt.Println(v55)
                                                                                } else {
                                                                                    var v80 int32 = 1
                                                                                    var v81 uint8 = 128
                                                                                    if (3.14159 < (0.1 + (1.5 / ((3.14159 - float32((3.14159 + 0.1))) - 1e10)))) {
                                                                                        switch fi37 {
                                                                                        case 359:
                                                                                            v71 = 127
                                                                                        case 9223372036854775807:
                                                                                            switch fi61 {
                                                                                            case 2147483647:
                                                                                                v71 += genFn0(v68, (!true))
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = "hello"
                                                                                                }
                                                                                            default:
                                                                                                var v82 uint16 = 2
                                                                                                _ = v82
                                                                                            }
                                                                                        default:
                                                                                            if (1 == v81) {
                                                                                                v60 = 981
                                                                                                var v83 float32 = float32(((0.0 / 1.5) - -2.25))
                                                                                                var v84 uint = (v1 >> ((18446744073709551615) & 63))
                                                                                                var v85 int16 = v60
                                                                                                v3["hello"] = (-(v55))
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                                _ = v85
                                                                                            }
                                                                                        }
                                                                                        if (953 == 909) {
                                                                                            v68 -= v68
                                                                                            if (-42 >= v80) {
                                                                                                _ = len(v3)
                                                                                                v68 = -9223372036854775808
                                                                                            } else {
                                                                                                v71 += (v71 - (v71 & int8(v1)))
                                                                                                fmt.Println(v71, v55)
                                                                                            }
                                                                                            fmt.Println(v1, v1, v71)
                                                                                            v60 -= (v60 << ((v1) & 63))
                                                                                        } else {
                                                                                            switch ("" + "\n\t") {
                                                                                            case "a":
                                                                                                v1 -= (v1 + v1)
                                                                                            case "hello":
                                                                                                fmt.Println(v80)
                                                                                            case "":
                                                                                                fmt.Println(v71, v1)
                                                                                            }
                                                                                        }
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = 4294967295
                                                                                        }
                                                                                    } else {
                                                                                        if ((v60 / (((v60 * v60)) | 1)) >= 42) {
                                                                                            var v86 bool = (((v1 >> ((v1) & 63)) < v1) && (int(v60) == fi37))
                                                                                            _ = v86
                                                                                        } else {
                                                                                            if false {
                                                                                                var v87 int8 = v71
                                                                                                _ = v87
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v88 string = "hello"
                                                                                                    var v89 int16 = v60
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                }
                                                                                                v90 := GenS1{F0: (v80 * 0), F1: "\n\t", F2: fi37}
                                                                                                _ = v90
                                                                                            }
                                                                                            v1 += v1
                                                                                            var v91 bool = true
                                                                                            var v92 int16 = (v60 << (((v1 % ((v1) | 1))) & 63))
                                                                                            v60 += v60
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                        }
                                                                                        v93 := GenS8{F0: (v81 / ((v81) | 1)), F1: (v81 / (((v81 >> ((v1) & 63))) | 1))}
                                                                                        _ = v93
                                                                                    }
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                }
                                                                                v94 := GenS9{F0: 357, F1: ("" + ("hello" + (("a" + "\n\t") + "hello")))}
                                                                                if true {
                                                                                    switch ("\n\t" + ("" + "")) {
                                                                                    case "a":
                                                                                        var v95 float32 = ((0.1 - 0.0) + float32(0.0))
                                                                                        _ = v95
                                                                                    case "foo bar":
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            v71 = (v71 & (v71 ^ (v71 << ((v1) & 63))))
                                                                                        case "\n\t":
                                                                                            if (0 == 2) {
                                                                                                switch ((("\n\t" + "hello") + "\n\t") + "\n\t") {
                                                                                                case "":
                                                                                                    var v96 int = 358
                                                                                                    _ = v96
                                                                                                case "a":
                                                                                                    fmt.Println(fi61)
                                                                                                }
                                                                                                v55 = -705
                                                                                            } else {
                                                                                                if (0.1 > float64((1.5 / -2.25))) {
                                                                                                    v71 = -85
                                                                                                }
                                                                                                v68 = (v68 ^ int64(v71))
                                                                                                fmt.Println(v60, fi61)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v97 int16 = 0
                                                                                        _ = v97
                                                                                    }
                                                                                } else {
                                                                                    if (!((false || (("a" + "") == "foo bar")) && true)) {
                                                                                        var v98 uint64 = 838
                                                                                        _ = v98
                                                                                    } else {
                                                                                        var v99 bool = true
                                                                                        fmt.Println(v55, v99)
                                                                                        _ = v99
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v100 := GenS1{F0: -2147483648, F1: "\n\t", F2: 1}
                                                                                var v101 uint64 = 42
                                                                                _ = v100
                                                                                _ = v101
                                                                                _ = v94
                                                                            }
                                                                        default:
                                                                            var v102 string = ""
                                                                            _ = v102
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    v71 = 1
                                                                    switch v55 {
                                                                    case 83:
                                                                        v1 = (v1 & (v1 + 0))
                                                                        var v103 uint = (v1 % (((v1 | (v1 | 2))) | 1))
                                                                        _ = v103
                                                                    case 42:
                                                                        switch fi37 {
                                                                        case 14:
                                                                            if (true || true) {
                                                                                v3["hello"] = (fi61 + (fi61 >> ((v1) & 63)))
                                                                                fmt.Println(v68, v60)
                                                                            }
                                                                        default:
                                                                            fmt.Println(v55)
                                                                        }
                                                                    default:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 0
                                                                        }
                                                                        v3[("foo bar" + "foo bar")] = -2147483648
                                                                    }
                                                                }
                                                                var v104 float64 = (0.1 / float64((1.5 * float32(1e10))))
                                                                _ = v104
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            v68 = v68
                                                            var v105 int64 = (v68 / ((-9223372036854775808) | 1))
                                                            var v106 int = fi61
                                                            var v107 uint8 = 1
                                                            _ = v105
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    default:
                                                        switch -959 {
                                                        case 656:
                                                            if false {
                                                                switch ("hello" + (("foo bar" + "\n\t") + "foo bar")) {
                                                                case "foo bar":
                                                                    _ = len(v2)
                                                                case "hello":
                                                                    var v108 int = (fi61 >> (((v1 - 4294967295)) & 63))
                                                                    _ = v108
                                                                default:
                                                                    v60 = v60
                                                                }
                                                                v1 = (v1 & v1)
                                                                var v109 uint32 = 511
                                                                _ = v109
                                                            }
                                                        case 424:
                                                            v110 := GenS8{F0: 83, F1: 0}
                                                            _ = v110
                                                        default:
                                                            delete(v3, "\n\t")
                                                            v55 = fi37
                                                        }
                                                        switch 128 {
                                                        case 128:
                                                            v111 := v3["hello"]
                                                            _ = v111
                                                        }
                                                        v112 := GenS3{F0: 100}
                                                        _ = v112
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v1, fi37, fi37)
                                                }
                                                v60 += (v60 + (v60 & -1))
                                                _ = v68
                                            }
                                        }
                                        _ = fi61
                                        _ = v62
                                    }
                                    switch v55 {
                                    case 32768:
                                        _ = len(v2)
                                    case 42:
                                        if (false || true) {
                                            var v113 int32 = -819
                                            switch v55 {
                                            case 127:
                                                v113 -= (v113 << ((v1) & 63))
                                            default:
                                                if true {
                                                    if ("" != "hello") {
                                                        fmt.Println(v60)
                                                    } else {
                                                        v113 = 2147483647
                                                        for fi114 := 0; fi114 < 12; fi114++ {
                                                            var v115 float32 = 0.0
                                                            v116 := GenS10{F0: 42}
                                                            var v117 uint32 = 2
                                                            v60 += -32768
                                                            if true {
                                                                v118 := v3[("a" + ("a" + "\n\t"))]
                                                                if len(v2) > 0 {
                                                                    v2[0] = "\n\t"
                                                                }
                                                                v55 -= (fi37 - (fi114 % (((v118 & 32768)) | 1)))
                                                                _ = v118
                                                            }
                                                            _ = fi114
                                                            _ = v115
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        fmt.Println(fi37, v1, v1)
                                                    }
                                                    if ((0.1 != (0.1 - 1.5)) && true) {
                                                        v3[""] = v55
                                                    }
                                                    if true {
                                                        if false {
                                                            v55 -= fi37
                                                            var v119 uint32 = 356
                                                            _ = v119
                                                        } else {
                                                            delete(v3, ("a" + ("a" + "")))
                                                            v55 -= (fi37 + fi37)
                                                        }
                                                    }
                                                } else {
                                                    v60 += v60
                                                    var v120 int32 = (^(v113))
                                                    _ = v120
                                                }
                                            }
                                            _ = v113
                                        }
                                    }
                                    _ = v60
                                }
                            } else {
                                v55 = v55
                                switch (v55 - (v55 >> ((v1) & 63))) {
                                case 112:
                                    delete(v3, "foo bar")
                                case 32767:
                                    v121 := GenS2{F0: (!true), F1: 2147483648, F2: (fi37 % ((fi37) | 1))}
                                    _ = v121
                                }
                            }
                        } else {
                            v1 += (^(v1))
                        }
                        _ = v55
                    }
                    _ = v54
                } else {
                    v122 := GenS8{F0: 255, F1: 127}
                    for fi123 := 0; fi123 < 12; fi123++ {
                        var v124 int32 = 257
                        _ = fi123
                        _ = v124
                    }
                    v125 := GenS5{F0: (fi37 / ((721) | 1)), F1: 505}
                    _ = v122
                    _ = v125
                }
                switch "" {
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    delete(v3, "foo bar")
                }
                _ = fi37
            }
        }
        var v126 int8 = -48
        _ = v126
    } else {
        v1 += 0
    }
    if true {
        fmt.Println(v1)
    }
    var v127 int16 = -32768
    if true {
        fmt.Println(v127, v1)
        switch -1 {
        case -9223372036854775808:
            var v128 uint8 = 255
            _ = v128
        default:
            if true {
                if false {
                    var v129 int8 = -128
                    _ = len(v2)
                    v127 += 32767
                    _ = v129
                } else {
                    v130 := GenS11{F0: 520}
                    _ = v130
                }
                var v131 string = ("hello" + ("\n\t" + ("" + ("hello" + "\n\t"))))
                var v132 float32 = (float32((3.14159 - 1.5)) / (1e10 + ((float32(3.14159) * 0.1) + 0.1)))
                _ = v131
                _ = v132
            }
        }
        switch 4294967295 {
        case 127:
            fmt.Println(v127, v127)
        case -1:
            v127 = (v127 + v127)
        default:
            if false {
                var v133 uint8 = 127
                _ = v133
            }
        }
    } else {
        v134 := GenS1{F0: 185, F1: "", F2: 128}
        _ = v134
    }
    v135 := GenS12{F0: true, F1: genFn0(42, true), F2: 42}
    if false {
        v1 = 42
        v136 := GenS7{F0: 255, F1: 104}
        _ = v136
    }
    v127 = 0
    var v137 uint = 4294967295
    fmt.Println(v127, v127, v1)
    v138 := GenS10{F0: v127}
    switch 32768 {
    case 330:
        if (9223372036854775807 < 42) {
            v2 = append(v2, "hello")
        }
    case 993:
        v137 = v137
    default:
        v1 += 4294967295
    }
    switch -940 {
    case 42:
        fmt.Println(v137, v1, v127)
    case 472:
        v139 := GenS11{F0: 519}
        _ = v139
    case -1:
        if (127 == 0) {
            delete(v3, "\n\t")
        } else {
            v137 = v1
            v137 = v1
        }
    default:
        var v140 uint64 = 1
        _ = v140
    }
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v137:", v137)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v127
    _ = v135
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v4
}
