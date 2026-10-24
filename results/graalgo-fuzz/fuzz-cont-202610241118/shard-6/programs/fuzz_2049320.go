package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint8; F1 int8; F2 int16 }
type GenS11 struct { F0 uint64; F1 bool; F2 int64 }
type GenS12 struct { F0 string }
func genFn0() int16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 18446744073709551615}
        var v6 uint16 = 32767
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(fi7, fi4)
            var v8 int = fi4
            _ = fi7
            _ = v8
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: fi9, F1: (fi9 > (-(fi9)))}
        switch 863 {
        case 32768:
            if false {
                var v11 uint64 = 823
                var v12 int32 = -265
                v13 := []int{fi9}
                _ = v11
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 float32 = -2.25
                    var v16 string = "foo bar"
                    v17 := []int{(fi14 + fi14), (^(fi14))}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
        case 127:
            var v18 float32 = (3.14159 * (3.14159 + 3.14159))
            _ = v18
        case -9223372036854775808:
            var v19 uint = 255
            _ = v19
        default:
            v20 := GenS3{F0: "a"}
            _ = v20
        }
        var v21 int8 = 100
        _ = fi9
        _ = v10
        _ = v21
    }
    var v22 int = 538
    _ = v22
    return 0
}
func genFn1(p0 int8) int16 {
    if false {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 float64 = (0.0 / (1e10 + (0.1 - 3.14159)))
            p0 -= p0
            fmt.Println(p0)
            _ = fi23
            _ = v24
        }
        switch -1 {
        case 42:
            var v25 int16 = 1000
            _ = v25
        default:
            var v26 uint16 = 32767
            _ = v26
        }
        switch "foo bar" {
        case "":
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 uint8 = 2
                v29 := GenS3{F0: ("a" + ("hello" + "a"))}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    var v30 float32 = (float32(0.0) + 0.1)
    if (!(-148 != -1)) {
        if true {
            if (742 > 412) {
                var v31 string = (("foo bar" + "hello") + "a")
                if false {
                    fmt.Println(v31, v30)
                    v32 := GenS3{F0: v31}
                    _ = v32
                } else {
                    switch 32767 {
                    case 32768:
                        if false {
                            switch 42 {
                            case 2147483647:
                                v33 := []int{4294967295, 128}
                                _ = v33
                            case 0:
                                for fi34 := 0; fi34 < 1; fi34++ {
                                    v31 = ("foo bar" + "")
                                    _ = fi34
                                }
                            case 128:
                                for fi35 := 0; fi35 < 16; fi35++ {
                                    v36 := GenS4{F0: 2, F1: p0, F2: 1}
                                    if true {
                                        switch (fi35 >> ((1) & 63)) {
                                        case 32768:
                                            var v37 bool = (v31 != v31)
                                            _ = v37
                                        case -9223372036854775808:
                                            v30 -= ((((v30 / float32(float64(((v30 - 1e10) * v30)))) + ((0.1 / 3.14159) + 1.5)) / 1e10) * v30)
                                        case -697:
                                            for fi38 := 0; fi38 < 9; fi38++ {
                                                var v39 bool = (255 > 255)
                                                if v39 {
                                                    switch (("hello" + v31) + v31) {
                                                    case "":
                                                        v40 := GenS1{F0: 1}
                                                        _ = v40
                                                    }
                                                }
                                                v41 := []int{128, fi38}
                                                _ = fi38
                                                _ = v39
                                                _ = v41
                                            }
                                        }
                                        if (127 != 1) {
                                            v42 := GenS5{F0: genFn0(), F1: (fi35 * 2147483647)}
                                            _ = v42
                                        } else {
                                            var v43 float32 = v30
                                            v44 := GenS3{F0: (((("\n\t" + (v31 + v31)) + (("\n\t" + (v31 + v31)) + "hello")) + v31) + "hello")}
                                            _ = v43
                                            _ = v44
                                        }
                                        p0 += -1
                                    }
                                    v45 := []int{fi35}
                                    _ = fi35
                                    _ = v36
                                    _ = v45
                                }
                            }
                        } else {
                            for fi46 := 0; fi46 < 1; fi46++ {
                                if (true && ((0 < 42) || false)) {
                                    switch 1 {
                                    case -319:
                                        v47 := GenS3{F0: v31}
                                        _ = v47
                                    }
                                    if false {
                                        if false {
                                            if (!true) {
                                                v48 := GenS4{F0: 312, F1: p0, F2: 32767}
                                                v49 := []int{fi46, fi46, -923}
                                                _ = len(v49)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                v30 = 0.0
                                            }
                                            var v50 int32 = 1
                                            v31 = ("foo bar" + v31)
                                            _ = v50
                                        }
                                        if (127 <= p0) {
                                            if false {
                                                v51 := []int{(fi46 << ((4294967295) & 63))}
                                                _ = v51
                                            } else {
                                                for fi52 := 0; fi52 < 1; fi52++ {
                                                    var v53 uint32 = 65535
                                                    v54 := GenS6{F0: (!(!true))}
                                                    v31 = v31
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                for fi55 := 0; fi55 < 3; fi55++ {
                                                    var v56 float64 = 0.0
                                                    var v57 int64 = 0
                                                    v30 = v30
                                                    _ = fi55
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        }
                                    } else {
                                        for fi58 := 0; fi58 < 2; fi58++ {
                                            fmt.Println(fi58, v30, p0)
                                            p0 += (p0 << ((1) & 63))
                                            _ = fi58
                                        }
                                        p0 -= p0
                                    }
                                    v59 := []int{fi46, fi46, fi46}
                                    _ = v59
                                } else {
                                    v60 := []int{fi46, (fi46 >> ((0) & 63))}
                                    _ = v60
                                }
                                _ = fi46
                            }
                        }
                    case 32767:
                        for fi61 := 0; fi61 < 6; fi61++ {
                            v31 = (v31 + v31)
                            _ = fi61
                        }
                    case 127:
                        var v62 int64 = 42
                        _ = v62
                    default:
                        v63 := []int{1}
                        _ = v63
                    }
                    for fi64 := 0; fi64 < 8; fi64++ {
                        var v65 string = v31
                        _ = fi64
                        _ = v65
                    }
                }
                if true {
                    fmt.Println(v31)
                } else {
                    v66 := []int{}
                    v30 = (v30 / v30)
                    if (764 != 400) {
                        for fi67 := 0; fi67 < 3; fi67++ {
                            p0 = p0
                            if true {
                                if len(v66) > 0 {
                                    v66[0] = (fi67 - 127)
                                }
                                if len(v66) >= 2 {
                                    v66 = v66[1:len(v66)]
                                }
                                fmt.Println(v30, v30)
                            }
                            _ = fi67
                        }
                    }
                    _ = v66
                }
                _ = v31
            } else {
                v68 := []int{247, 42}
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 int = fi69
                    _ = fi69
                    _ = v70
                }
                _ = v68
            }
            var v71 uint64 = 0
            var v72 uint32 = 0
            _ = v71
            _ = v72
        } else {
            v73 := []int{0}
            var v74 uint = 65535
            v75 := GenS7{F0: 109, F1: 1}
            _ = v73
            _ = v74
            _ = v75
        }
    } else {
        for fi76 := 0; fi76 < 3; fi76++ {
            v30 += 0.0
            if true {
                v77 := []int{586, (fi76 | (fi76 | 0)), (fi76 ^ (^(fi76))), (-(fi76))}
                _ = v77
            } else {
                v78 := []int{-2147483648}
                _ = v78
            }
            _ = fi76
        }
        v79 := GenS5{F0: 1, F1: 32767}
        v30 += v30
        _ = v79
    }
    var v80 float32 = v30
    _ = p0
    _ = v30
    _ = v80
    return 1000
}

func main() {
    v1 := []uint32{251, 2147483648, 42, 2147483648}
    v2 := make(map[string]uint32)
    var v3 int32 = 2147483647
    fmt.Println(v3)
    switch ("foo bar" + "\n\t") {
    case "a":
        if false {
            var v81 int16 = 42
            _ = v81
        }
    case "hello":
        fmt.Println(v3)
    default:
        v2[("hello" + ("" + ("a" + "foo bar")))] = 2
    }
    for fi82 := 0; fi82 < 11; fi82++ {
        if (!false) {
            var v83 int8 = 0
            v3 = -1
            _ = v83
        }
        v3 += 2147483647
        _ = fi82
    }
    v3 = 42
    for fi84 := 0; fi84 < 7; fi84++ {
        if (v3 <= v3) {
            for fi85 := 0; fi85 < 3; fi85++ {
                switch 42 {
                case 32767:
                    if (!true) {
                        if (2147483647 <= 1) {
                            v86 := v2[("" + ("foo bar" + (((("hello" + "hello") + "hello") + "a") + ("a" + "a"))))]
                            var v87 int64 = 1
                            _ = v86
                            _ = v87
                        }
                        if false {
                            v3 -= 0
                        }
                        v3 = (v3 * v3)
                        var v88 int16 = 0
                        _ = v88
                    } else {
                        v3 += 2147483647
                        v89 := GenS7{F0: fi85, F1: -1}
                        var v90 int = 32768
                        _ = v89
                        _ = v90
                    }
                case 127:
                    fmt.Println(v3)
                }
                v3 = -2147483648
                var v91 uint8 = 2
                var v92 int64 = 971
                if (!((0 < 548) && (true || ((-(fi85)) <= fi85)))) {
                    v2["hello"] = 1
                    v93 := GenS8{F0: 705, F1: v92}
                    _ = v93
                }
                _ = fi85
                _ = v91
                _ = v92
            }
            for fi94 := 0; fi94 < 8; fi94++ {
                switch ((("a" + "\n\t") + "a") + ("a" + ("a" + "hello"))) {
                case "foo bar":
                    v95 := GenS7{F0: (fi84 << ((2) & 63)), F1: (v3 * (v3 - -2147483648))}
                    if false {
                        v96 := GenS2{F0: fi94, F1: true}
                        _ = v96
                    }
                    v3 = -1
                    _ = v95
                case "\n\t":
                    if (42 <= 962) {
                        fmt.Println(v3)
                    }
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi94, fi84)
                case "a":
                    if (((-1 > (v3 / (((v3 * (v3 % (((v3 - v3)) | 1)))) | 1))) || true) && ((false || (fi94 != fi94)) && (-1 <= genFn1(1)))) {
                        v97 := v2[""]
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                        var v98 int64 = -1
                        _ = v97
                        _ = v98
                    }
                default:
                    if (1 < 18446744073709551615) {
                        var v99 uint16 = 65535
                        if true {
                            var v100 uint8 = 1
                            v3 -= (v3 + -128)
                            _ = v100
                        } else {
                            v3 += v3
                            var v101 int8 = 127
                            _ = v101
                        }
                        _ = v99
                    } else {
                        switch (fi84 & fi84) {
                        case -1:
                            if true {
                                v3 = v3
                                switch (fi94 % ((-2147483648) | 1)) {
                                case 128:
                                    if (42 >= 32767) {
                                        delete(v2, "foo bar")
                                        var v102 string = "foo bar"
                                        _ = v102
                                    }
                                case 445:
                                    fmt.Println(fi94)
                                default:
                                    if (!true) {
                                        switch fi94 {
                                        case 32768:
                                            v103 := GenS9{F0: 1, F1: true}
                                            _ = v103
                                        case -2147483648:
                                            switch 32768 {
                                            case -2147483648:
                                                switch "foo bar" {
                                                case "a":
                                                    v104 := GenS9{F0: 0, F1: true}
                                                    _ = v104
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        _ = len(v1)
                                                    case "foo bar":
                                                        v3 -= (v3 | v3)
                                                    default:
                                                        v105 := GenS10{F0: 0, F1: -1, F2: 42}
                                                        _ = v105
                                                    }
                                                case "hello":
                                                    var v106 uint32 = 2
                                                    _ = v106
                                                default:
                                                    fmt.Println(fi84, v3)
                                                }
                                            default:
                                                v3 = v3
                                            }
                                        case 127:
                                            _ = len(v1)
                                        default:
                                            v3 = (v3 + 258)
                                        }
                                    }
                                }
                                v107 := GenS11{F0: 42, F1: false, F2: 0}
                                _ = v107
                            }
                        case 127:
                            var v108 float64 = -2.25
                            _ = v108
                        case 2147483647:
                            _ = len(v1)
                        }
                    }
                }
                v3 += -1
                switch int(v3) {
                case 32768:
                    v109 := GenS12{F0: (("a" + "hello") + "foo bar")}
                    _ = v109
                default:
                    fmt.Println(fi84, fi84)
                }
                v3 += v3
                var v110 uint32 = 2147483647
                _ = fi94
                _ = v110
            }
            var v111 uint16 = 42
            _ = v111
        }
        _ = fi84
    }
    fmt.Println(v3)
    v3 = (v3 << ((42) & 63))
    fmt.Println(v3)
    var v112 uint16 = 673
    switch "hello" {
    case "":
        for fi113 := 0; fi113 < 12; fi113++ {
            if len(v1) > 0 {
                v1[0] = 312
            }
            var v114 bool = false
            var v115 uint32 = 79
            _ = fi113
            _ = v114
            _ = v115
        }
    }
    for fi116 := 0; fi116 < 5; fi116++ {
        v3 = 1
        _ = fi116
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v112
    _ = v2
    _ = v3
}
