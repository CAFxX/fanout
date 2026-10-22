package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 int32 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint8; F1 int; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 uint; F2 int }
type GenS9 struct { F0 int16; F1 uint64; F2 bool }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 int; F1 uint64; F2 bool }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int64; F1 uint32; F2 int }
func genFn0() bool {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 438, F1: 1}
        var v6 float64 = 0.0
        if (-2.25 >= 1.5) {
            var v7 float32 = 0.1
            _ = v7
        } else {
            v6 = v6
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v8 float64 = float64(0.1)
    if true {
        v8 -= float64(1e10)
        if (0 > 1) {
            v8 = ((v8 - 1.5) / v8)
            for fi9 := 0; fi9 < 16; fi9++ {
                fmt.Println(v8, fi9)
                var v10 int8 = 42
                v8 = v8
                _ = fi9
                _ = v10
            }
        } else {
            if (true || (42 <= 462)) {
                v11 := []int{32767}
                v8 -= (3.14159 + ((v8 * v8) / 1e10))
                var v12 uint64 = 42
                _ = v11
                _ = v12
            }
        }
        if (425 == 4294967295) {
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int16 = 32767
                _ = fi13
                _ = v14
            }
        }
    } else {
        if (!(1000 != 1000)) {
            fmt.Println(v8)
        } else {
            v15 := GenS2{F0: 350}
            v16 := GenS1{F0: 9223372036854775807, F1: 42}
            if (42 <= 72) {
                v17 := GenS3{F0: 32767}
                _ = v17
            }
            _ = v15
            _ = v16
        }
        var v18 string = "hello"
        for fi19 := 0; fi19 < 19; fi19++ {
            v18 = v18
            _ = fi19
        }
        _ = v18
    }
    _ = v8
    return true
}
func genFn1(p0 int8, p1 int64) int32 {
    for fi20 := 0; fi20 < 1; fi20++ {
        p1 = p1
        _ = fi20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := []int{fi21, 128, 32768}
        if false {
            v23 := GenS4{F0: 68, F1: "a", F2: -1}
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    if true {
        if true {
            v24 := []int{32768, 1, 0}
            _ = v24
        } else {
            p1 = 9223372036854775807
            var v25 int32 = 722
            for fi26 := 0; fi26 < 5; fi26++ {
                if false {
                    fmt.Println(p0, p1, p0)
                    p1 += -9223372036854775808
                    p0 -= 1
                } else {
                    for fi27 := 0; fi27 < 8; fi27++ {
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                v25 = v25
                            case "a":
                                v25 = v25
                            case "":
                                var v28 bool = (!true)
                                _ = v28
                            }
                            if ((p0 / ((p0) | 1)) < p0) {
                                v29 := GenS5{F0: 2147483647, F1: fi27}
                                v30 := []int{(fi26 | 9223372036854775807), 2147483647, fi26, (fi27 >> ((4294967295) & 63))}
                                if false {
                                    switch 32768 {
                                    case 4294967295:
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v25)
                                            p1 = 0
                                        }
                                    }
                                    v25 += (v25 & (-(v25)))
                                    p1 += -733
                                } else {
                                    if true {
                                        v25 += (v25 >> ((1) & 63))
                                        if true {
                                            var v31 float64 = -2.25
                                            v25 += -964
                                            _ = v31
                                        } else {
                                            p0 = (p0 + 83)
                                            v30 = append(v30, (fi26 << ((1) & 63)))
                                        }
                                        if len(v30) >= 2 {
                                            v30 = v30[1:len(v30)]
                                        }
                                        fmt.Println(fi27, fi26)
                                    }
                                    v32 := GenS2{F0: 65535}
                                    _ = v32
                                }
                                v30 = append(v30, (fi26 << ((2) & 63)))
                                fmt.Println(fi26, v25)
                                var v33 float32 = 1.5
                                p1 = (p1 | int64(fi27))
                                _ = v29
                                _ = v30
                                _ = v33
                            }
                        } else {
                            var v34 float64 = -2.25
                            _ = v34
                        }
                        v25 += int32(p1)
                        if true {
                            v35 := GenS3{F0: 1}
                            v36 := GenS5{F0: 701, F1: (fi26 + -1)}
                            _ = v35
                            _ = v36
                        }
                        _ = fi27
                    }
                }
                _ = fi26
            }
            _ = v25
        }
        v37 := []int{32768, -2147483648, -1}
        _ = v37
    }
    _ = p0
    _ = p1
    return 463
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    v3 := make(map[int]uint32)
    v2 = v2
    var v38 uint64 = 4294967295
    for fi39 := 0; fi39 < 20; fi39++ {
        switch ("\n\t" + "") {
        case "a":
            v38 += 552
        case "":
            var v40 float32 = (1.5 * 1.5)
            _ = v40
        case "hello":
            switch "" {
            case "hello":
                for fi41 := 0; fi41 < 5; fi41++ {
                    v38 -= (v38 / (((v38 ^ (v38 / (((v38 >> ((0) & 63))) | 1)))) | 1))
                    v42 := GenS2{F0: 2}
                    _ = fi41
                    _ = v42
                }
            case "foo bar":
                v43 := GenS4{F0: 127, F1: ("" + "\n\t"), F2: 42}
                _ = v43
            default:
                var v44 int16 = 32767
                _ = v44
            }
        }
        v3[(fi39 - (fi39 << ((1) & 63)))] = 42
        _ = fi39
    }
    var v45 int64 = -199
    v45 = (v45 >> ((65535) & 63))
    v2 = 0.0
    if (!true) {
        if (!(!(42 <= 2))) {
            for fi46 := 0; fi46 < 8; fi46++ {
                v47 := GenS6{F0: 63, F1: (fi46 * -394), F2: ("foo bar" + "a")}
                v2 -= (v2 - 0.0)
                switch fi46 {
                case 32767:
                    var v48 int = fi46
                    _ = v48
                case -308:
                    v49 := GenS2{F0: 255}
                    _ = v49
                default:
                    for fi50 := 0; fi50 < 3; fi50++ {
                        v51 := []int{1, (-(fi50)), fi50, 42}
                        v2 -= (v2 - ((-2.25 + v2) * float64(1.5)))
                        v38 = v38
                        var v52 float32 = float32(v2)
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
                _ = fi46
                _ = v47
            }
            v38 = 2
            if true {
                fmt.Println(v38, v45, v2)
                if ((1e10 + (((float64(0.0) + v2) * 0.1) / 0.1)) <= (float64((0.0 / (float32((float64(1.5) * v2)) - 0.1))) / -2.25)) {
                    var v53 uint64 = 4294967295
                    if (4294967295 >= 847) {
                        for fi54 := 0; fi54 < 17; fi54++ {
                            if ((fi54 | fi54) <= fi54) {
                                v45 -= (v45 << ((65535) & 63))
                                v38 += v38
                            }
                            if (!(687 == 474)) {
                                v55 := GenS3{F0: 32767}
                                _ = v55
                            } else {
                                var v56 bool = false
                                _ = v56
                            }
                            v38 += v38
                            _ = fi54
                        }
                    } else {
                        v57 := []int{0, -2147483648, 4294967295}
                        switch "" {
                        case "foo bar":
                            var v58 int16 = 0
                            _ = v58
                        case "a":
                            v38 = 0
                        }
                        _ = v57
                    }
                    _ = v53
                }
            } else {
                _ = len(v1)
                v59 := GenS4{F0: 1, F1: ("a" + (("\n\t" + ("hello" + "")) + "")), F2: 42}
                _ = v59
            }
        } else {
            for fi60 := 0; fi60 < 8; fi60++ {
                if false {
                    fmt.Println(v38, v2)
                }
                for fi61 := 0; fi61 < 13; fi61++ {
                    v45 -= v45
                    v62 := []int{(fi61 >> ((0) & 63))}
                    v2 = 0.1
                    _ = fi61
                    _ = v62
                }
                _ = fi60
            }
        }
        v38 -= (^(v38))
        v63 := []int{}
        _ = v63
    } else {
        _ = len(v1)
        if true {
            delete(v3, 715)
            for fi64 := 0; fi64 < 4; fi64++ {
                v2 -= v2
                v65 := GenS7{F0: fi64}
                v2 -= 1.5
                _ = fi64
                _ = v65
            }
        } else {
            var v66 string = "a"
            _ = v66
        }
        fmt.Println(v2, v2)
    }
    v45 = v45
    v67 := v3[32767]
    if true {
        v68 := GenS3{F0: 2}
        v69 := v3[1]
        switch ("hello" + "hello") {
        case "hello":
            var v70 string = ("hello" + ("" + "a"))
            _ = v70
        case "":
            if (939 != 0) {
                fmt.Println(v69)
                _ = len(v1)
            }
        }
        _ = v68
        _ = v69
    } else {
        if (!(3.14159 >= (((0.1 / 1.5) + float32(1e10)) - 3.14159))) {
            if false {
                var v71 uint = 401
                switch "\n\t" {
                case "":
                    switch 827 {
                    case 9223372036854775807:
                        switch "" {
                        case "\n\t":
                            fmt.Println(v45, v71, v45)
                        case "a":
                            if (!true) {
                                fmt.Println(v38, v38)
                                var v72 int32 = 2147483647
                                _ = v72
                            }
                        default:
                            if ((v71 >> ((v71) & 63)) != v71) {
                                var v73 float32 = 1.5
                                fmt.Println(v67)
                                _ = v73
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    for fi74 := 0; fi74 < 8; fi74++ {
                                        if true {
                                            if true {
                                                if true {
                                                    delete(v3, (fi74 & 921))
                                                    if false {
                                                        for fi75 := 0; fi75 < 11; fi75++ {
                                                            var v76 bool = (genFn1(1, (-(v45))) <= -1)
                                                            _ = fi75
                                                            _ = v76
                                                        }
                                                        if false {
                                                            fmt.Println(v71, v45)
                                                        }
                                                        fmt.Println(fi74, fi74, v2)
                                                    }
                                                    if (976 <= -384) {
                                                        if (!(v71 >= (v71 % ((v71) | 1)))) {
                                                            v67 = 404
                                                        }
                                                        switch 128 {
                                                        case 32768:
                                                            switch (fi74 | (fi74 % ((-9223372036854775808) | 1))) {
                                                            case 0:
                                                                for fi77 := 0; fi77 < 4; fi77++ {
                                                                    v38 -= (v38 / (((v38 + 42)) | 1))
                                                                    switch (^(fi74)) {
                                                                    case 128:
                                                                        if true {
                                                                            var v78 float64 = v2
                                                                            var v79 int = fi74
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "a":
                                                                                var v80 float64 = (v78 + (v2 / v2))
                                                                                _ = v80
                                                                            case "\n\t":
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    var v81 int8 = 0
                                                                                    _ = v81
                                                                                case "":
                                                                                    v2 -= (v2 / (1e10 - v78))
                                                                                default:
                                                                                    var v82 int8 = 100
                                                                                    _ = v82
                                                                                }
                                                                                v67 -= (v67 << (((v71 << (((v71 << (((v71 * uint(v38))) & 63))) & 63))) & 63))
                                                                                var v83 float64 = 0.1
                                                                                _ = v83
                                                                            case "":
                                                                                var v84 uint64 = (v38 | (v38 + (v38 & v38)))
                                                                                _ = v84
                                                                            }
                                                                            v85 := GenS3{F0: 65535}
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v85
                                                                        } else {
                                                                            fmt.Println(fi77, v67, fi77)
                                                                            var v86 int64 = v45
                                                                            var v87 float32 = float32(float64(3.14159))
                                                                            v88 := GenS8{F0: 42, F1: (^(v71)), F2: (fi74 + fi74)}
                                                                            v89 := GenS9{F0: 198, F1: 42, F2: true}
                                                                            _ = v86
                                                                            _ = v87
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    case 844:
                                                                        v90 := GenS9{F0: -609, F1: 18446744073709551615, F2: (127 == 128)}
                                                                        _ = v90
                                                                    }
                                                                    if (((((v45 * (v45 - (^(v45)))) != v45) && true) && (fi74 == 4294967295)) || (true && false)) {
                                                                        v1[(fi77 | (fi74 % ((fi74) | 1)))] = false
                                                                        v91 := []int{fi74}
                                                                        fmt.Println(fi74, v67, fi74)
                                                                        _ = v91
                                                                    }
                                                                    _ = fi77
                                                                }
                                                            case -2147483648:
                                                                v67 = (v67 ^ (v67 - (v67 ^ v67)))
                                                            case -1:
                                                                v92 := []int{fi74}
                                                                _ = v92
                                                            }
                                                        case 4294967295:
                                                            v93 := GenS8{F0: genFn1(-1, v45), F1: v71, F2: -9223372036854775808}
                                                            _ = v93
                                                        default:
                                                            var v94 uint8 = 127
                                                            _ = v94
                                                        }
                                                        v95 := GenS10{F0: (v38 << ((4294967295) & 63)), F1: 0}
                                                        _ = v95
                                                    } else {
                                                        v71 = 498
                                                        var v96 int8 = 55
                                                        _ = v96
                                                    }
                                                }
                                                fmt.Println(fi74, v45)
                                                v97 := []int{}
                                                _ = v97
                                            } else {
                                                fmt.Println(v38, v2)
                                                var v98 uint64 = (v38 * v38)
                                                _ = v98
                                            }
                                            fmt.Println(v45, v45)
                                        } else {
                                            var v99 float64 = ((0.0 / (0.1 * v2)) - float64(1.5))
                                            _ = v99
                                        }
                                        var v100 float64 = (v2 * v2)
                                        for fi101 := 0; fi101 < 14; fi101++ {
                                            v100 -= (v100 + ((v100 / v100) / v100))
                                            v102 := []int{}
                                            v103 := GenS1{F0: (v45 >> (((v71 >> (((v71 % ((42) | 1))) & 63))) & 63)), F1: -1}
                                            _ = fi101
                                            _ = v102
                                            _ = v103
                                        }
                                        _ = fi74
                                        _ = v100
                                    }
                                default:
                                    if (((v71 / (((v71 / ((v71) | 1))) | 1)) < (v71 - 4294967295)) && (-302 < 1)) {
                                        fmt.Println(v67, v45, v38)
                                    }
                                }
                            }
                        }
                    default:
                        for fi104 := 0; fi104 < 2; fi104++ {
                            v38 = 18446744073709551615
                            var v105 int = 45
                            v71 -= 42
                            _ = fi104
                            _ = v105
                        }
                    }
                case "hello":
                    fmt.Println(v2, v2)
                default:
                    v106 := GenS10{F0: (v38 >> (((v71 ^ v71)) & 63)), F1: 2}
                    _ = v106
                }
                _ = v71
            }
            for fi107 := 0; fi107 < 2; fi107++ {
                v45 += 42
                for fi108 := 0; fi108 < 3; fi108++ {
                    v109 := GenS6{F0: 2, F1: 1, F2: ("hello" + "")}
                    if true {
                        if (2 <= 245) {
                            switch (fi108 * (fi108 / ((fi107) | 1))) {
                            case 1:
                                v67 += v67
                            }
                            v110 := GenS11{F0: (fi108 - (fi108 - fi108)), F1: (v38 & (v38 - v38)), F2: false}
                            _ = v110
                        } else {
                            v2 = 0.0
                            if true {
                                if (-1 >= int8(v67)) {
                                    var v111 float64 = (1e10 / 1e10)
                                    v2 += 3.14159
                                    var v112 int8 = int8(v45)
                                    v113 := v1[fi108]
                                    v113 = (135 == 42)
                                    switch ("" + "hello") {
                                    case "":
                                        var v114 bool = (!(!v113))
                                        _ = v114
                                    default:
                                        fmt.Println(v45, fi108)
                                    }
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                } else {
                                    if true {
                                        v2 += float64(((float32(v2) + 3.14159) / 0.1))
                                        v115 := GenS12{F0: v67, F1: -2147483648}
                                        _ = v115
                                    } else {
                                        var v116 float32 = ((-2.25 * (0.0 * -2.25)) * (1e10 / float32(1.5)))
                                        var v117 int16 = 1
                                        _ = v116
                                        _ = v117
                                    }
                                    var v118 uint16 = 655
                                    _ = v118
                                }
                            }
                        }
                        fmt.Println(v45)
                    }
                    _ = fi108
                    _ = v109
                }
                var v119 uint16 = 76
                _ = fi107
                _ = v119
            }
        }
        var v120 int16 = 1
        for fi121 := 0; fi121 < 3; fi121++ {
            fmt.Println(v45)
            _ = fi121
        }
        _ = v120
    }
    v122 := GenS13{F0: -128}
    v123 := []int{571, 244}
    v124 := GenS11{F0: 32767, F1: v38, F2: (!(1.5 <= 3.14159))}
    var v125 bool = false
    var v126 uint8 = 219
    v127 := GenS14{F0: v45, F1: (^(v67)), F2: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v123:", len(v123), v123)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v45:", v45)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v38
    _ = v45
    _ = v67
}
