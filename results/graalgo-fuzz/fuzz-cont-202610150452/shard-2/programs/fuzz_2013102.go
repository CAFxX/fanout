package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int; F1 int32; F2 int8 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool; F1 int32; F2 uint8 }
type GenS14 struct { F0 bool; F1 int16 }
type GenS15 struct { F0 string; F1 int8; F2 int32 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 string; F1 uint64 }
type GenS18 struct { F0 bool; F1 bool }
type GenS19 struct { F0 uint64; F1 uint32 }
type GenS20 struct { F0 uint; F1 bool; F2 uint8 }
func genFn0(p0 int8) int64 {
    fmt.Println(p0)
    v4 := []int{247, 32767}
    if (false || (0 >= -673)) {
        p0 -= p0
    } else {
        var v5 int = 32767
        _ = v5
    }
    p0 = (p0 >> ((669) & 63))
    _ = p0
    _ = v4
    return -921
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 76
    v3 := make(map[int]string)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 42, F1: 509, F2: 1000}
        if (("a" + "\n\t") < "") {
            var v8 float64 = 3.14159
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi6, fi9, v8)
                _ = fi9
            }
            _ = v8
        } else {
            v2 = 0
        }
        if (false && false) {
            v10 := GenS2{F0: 128, F1: 42}
            fmt.Println(v2, v2)
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS3{F0: (-746 <= (^(fi11))), F1: 0, F2: "a"}
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS2{F0: 130, F1: 42}
            v15 := GenS4{F0: ("a" + "")}
            _ = fi13
            _ = v14
            _ = v15
        }
        v2 += v2
        _ = fi11
        _ = v12
    }
    v16 := GenS5{F0: 255}
    switch (("hello" + ((("" + (("a" + ("hello" + ("hello" + "hello"))) + ("hello" + "foo bar"))) + "\n\t") + "a")) + ("foo bar" + "foo bar")) {
    case "":
        v17 := GenS5{F0: 255}
        _ = v17
    case "hello":
        var v18 bool = (201 >= -660)
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 4; fi19++ {
            switch (fi19 % ((fi19) | 1)) {
            case -9223372036854775808:
                v2 = -128
            default:
                var v20 int8 = -1
                _ = v20
            }
            v2 = v2
            _ = fi19
        }
    }
    if true {
        v2 += v2
        _ = len(v1)
    }
    var v21 int16 = 229
    v22 := v3[2147483647]
    var v23 uint64 = 1
    switch (("a" + "hello") + "\n\t") {
    case "\n\t":
        switch 2147483647 {
        case -797:
            for fi24 := 0; fi24 < 10; fi24++ {
                for fi25 := 0; fi25 < 6; fi25++ {
                    if (581 <= v21) {
                        if (true && true) {
                            switch 42 {
                            case 0:
                                if false {
                                    v2 = 1
                                    if false {
                                        if false {
                                            fmt.Println(fi24)
                                            var v26 string = ((v22 + "a") + ("a" + "a"))
                                            _ = v26
                                        } else {
                                            v27 := GenS4{F0: v22}
                                            var v28 int8 = 1
                                            var v29 int16 = (v21 & v21)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                    } else {
                                        v30 := GenS5{F0: 0}
                                        v31 := GenS6{F0: (255 > 118)}
                                        v32 := GenS7{F0: (v2 >> ((0) & 63)), F1: v21}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    if true {
                                        _ = len(v3)
                                        var v33 bool = false
                                        v34 := GenS8{F0: 2, F1: 22}
                                        _ = v33
                                        _ = v34
                                    }
                                    v21 = 1
                                    v35 := GenS2{F0: 255, F1: (v23 % ((v23) | 1))}
                                    _ = v35
                                }
                            case -9223372036854775808:
                                delete(v3, fi24)
                            }
                            var v36 bool = false
                            v37 := GenS9{F0: 610, F1: 0, F2: 0}
                            fmt.Println(v23)
                            _ = v36
                            _ = v37
                        } else {
                            v21 = (v21 + v21)
                            v38 := GenS9{F0: 0, F1: 127, F2: 170}
                            _ = v38
                        }
                        v23 = v23
                    }
                    switch v22 {
                    case "foo bar":
                        switch 128 {
                        case -426:
                            var v39 int16 = v21
                            _ = v39
                        case 32767:
                            v3[(fi25 | (fi24 << ((2) & 63)))] = v22
                        }
                    }
                    v40 := GenS10{F0: fi25, F1: 897, F2: (v2 << ((420) & 63))}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi25
                    _ = v40
                }
                var v41 uint64 = v23
                var v42 int16 = v21
                _ = fi24
                _ = v41
                _ = v42
            }
        case 4294967295:
            v43 := GenS11{F0: int64(v23), F1: v2}
            _ = v43
        default:
            fmt.Println(v22, v2)
        }
    case "foo bar":
        v44 := GenS12{F0: 9223372036854775807}
        _ = v44
    default:
        switch (v22 + "a") {
        case "\n\t":
            switch (((((("" + v22) + v22) + (v22 + "")) + (v22 + v22)) + ((v22 + v22) + "foo bar")) + v22) {
            case "":
                v2 = v2
            case "hello":
                var v45 string = (("a" + ("\n\t" + v22)) + (v22 + ""))
                _ = v45
            case "\n\t":
                var v46 int16 = v21
                _ = v46
            }
        }
    }
    if ((!(true && ((v23 <= 1) && false))) && (1 >= 1)) {
        for fi47 := 0; fi47 < 4; fi47++ {
            v48 := GenS13{F0: (42 > 942), F1: 2147483647, F2: 0}
            _ = fi47
            _ = v48
        }
    }
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        v50 := GenS14{F0: (v22 == v22), F1: (-(v21))}
        v21 = v21
        _ = fi49
        _ = v50
    }
    v1 = append(v1, 924)
    var v51 uint16 = 757
    for fi52 := 0; fi52 < 12; fi52++ {
        for fi53 := 0; fi53 < 12; fi53++ {
            if (128 > 2) {
                switch (fi53 - (fi52 ^ -1)) {
                case -2147483648:
                    v54 := GenS11{F0: 42, F1: (v2 & (v2 / ((v2) | 1)))}
                    _ = v54
                case 32768:
                    switch (v22 + v22) {
                    case "":
                        if true {
                            var v55 int16 = v21
                            v1 = append(v1, 2147483648)
                            _ = v55
                        } else {
                            fmt.Println(fi53, v51)
                            var v56 uint32 = 42
                            _ = v56
                        }
                    default:
                        v23 += 18446744073709551615
                    }
                    var v57 uint8 = 42
                    var v58 uint32 = 1
                    if (1 != 0) {
                        delete(v3, fi52)
                        var v59 float32 = ((-2.25 * 1e10) / ((((0.0 * -2.25) / (-2.25 + 0.0)) - (1e10 + 3.14159)) / 1e10))
                        v21 += (v21 % ((377) | 1))
                        switch v22 {
                        case "foo bar":
                            _ = len(v3)
                        case "":
                            v23 = 42
                        default:
                            v60 := GenS12{F0: -784}
                            _ = v60
                        }
                        if false {
                            fmt.Println(v2, fi52, v2)
                            v51 = 42
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = len(v1)
                            }
                            switch fi52 {
                            case 1:
                                if (2 <= (v51 << ((42) & 63))) {
                                    fmt.Println(v23, v22)
                                    var v61 int8 = -104
                                    if true {
                                        _ = len(v1)
                                        if (!(42 <= 0)) {
                                            v62 := GenS15{F0: (v22 + "hello"), F1: (v2 & v2), F2: 1}
                                            switch ("\n\t" + ((("hello" + v22) + "foo bar") + (((("\n\t" + "\n\t") + v22) + "a") + ("hello" + "a")))) {
                                            case "a":
                                                var v63 int64 = -9223372036854775808
                                                _ = v63
                                            case "\n\t":
                                                var v64 float64 = (0.1 * (3.14159 * (0.1 / -2.25)))
                                                _ = v64
                                            default:
                                                var v65 uint = 65535
                                                _ = v65
                                            }
                                            var v66 int8 = (v2 % (((v2 % (((v61 % ((v2) | 1))) | 1))) | 1))
                                            _ = v62
                                            _ = v66
                                        }
                                    } else {
                                        switch (((v22 + v22) + (v22 + "")) + "\n\t") {
                                        case "\n\t":
                                            switch fi52 {
                                            case 592:
                                                switch fi53 {
                                                case 42:
                                                    v22 = (v22 + ((v22 + v22) + "a"))
                                                case 4294967295:
                                                    var v67 int64 = 1
                                                    _ = v67
                                                case -9223372036854775808:
                                                    fmt.Println(v59)
                                                default:
                                                    switch (^(fi53)) {
                                                    case -862:
                                                        v68 := GenS14{F0: true, F1: 0}
                                                        _ = v68
                                                    }
                                                }
                                            case -37:
                                                v69 := GenS10{F0: fi53, F1: 0, F2: -82}
                                                _ = v69
                                            }
                                        default:
                                            fmt.Println(v22)
                                        }
                                    }
                                    _ = v61
                                }
                            case -188:
                                if (false || (1e10 < ((1e10 / (1e10 + 3.14159)) - 1e10))) {
                                    v70 := GenS7{F0: (-(v2)), F1: v21}
                                    if false {
                                        fmt.Println(v57)
                                        if (v22 >= (("a" + "\n\t") + v22)) {
                                            if false {
                                                v23 -= 2
                                                v2 -= (v2 / ((v2) | 1))
                                                var v71 int = (fi53 - (fi53 % (((fi53 ^ fi53)) | 1)))
                                                _ = v71
                                            }
                                            switch fi53 {
                                            case 9223372036854775807:
                                                var v72 int8 = 18
                                                _ = v72
                                            default:
                                                v59 -= 0.0
                                            }
                                            if false {
                                                var v73 bool = true
                                                _ = v73
                                            } else {
                                                if (true && (v57 >= 128)) {
                                                    if true {
                                                        v58 -= v58
                                                        _ = len(v3)
                                                    }
                                                    var v74 string = v22
                                                    fmt.Println(v21)
                                                    _ = v74
                                                }
                                                if (!false) {
                                                    switch fi52 {
                                                    case -126:
                                                        fmt.Println(v21, v23, v22)
                                                    case 71:
                                                        var v75 bool = false
                                                        _ = v75
                                                    }
                                                }
                                            }
                                        }
                                        var v76 int16 = (-(v21))
                                        _ = v76
                                    } else {
                                        var v77 uint8 = uint8(v51)
                                        v77 = 1
                                        v22 = v22
                                        var v78 int32 = -984
                                        if true {
                                            v1 = append(v1, v58)
                                        } else {
                                            v22 = "a"
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v79 := GenS16{F0: (v57 % ((v57) | 1))}
                                            _ = v79
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                    var v80 int16 = -1
                                    _ = v70
                                    _ = v80
                                } else {
                                    switch ("a" + v22) {
                                    case "hello":
                                        if ((v2 * int8(v58)) >= int8(v51)) {
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v23)
                                            case "hello":
                                                v23 = (v23 >> ((255) & 63))
                                            case "foo bar":
                                                if (-2.25 >= 0.0) {
                                                    if len(v1) > 0 {
                                                        v1[0] = v58
                                                    }
                                                    switch (fi53 & fi53) {
                                                    case 42:
                                                        v1 = append(v1, (v58 << ((638) & 63)))
                                                    default:
                                                        v22 = v22
                                                    }
                                                    var v81 int8 = 86
                                                    _ = v81
                                                }
                                            default:
                                                switch (fi52 << ((106) & 63)) {
                                                case 9223372036854775807:
                                                    var v82 uint64 = 2
                                                    _ = v82
                                                case 32767:
                                                    v3[128] = (v22 + v22)
                                                case -2147483648:
                                                    if len(v1) > 0 {
                                                        v1[0] = uint32(v51)
                                                    }
                                                }
                                            }
                                        } else {
                                            if false {
                                                var v83 int64 = -1
                                                if (v57 <= (v57 - (v57 & 224))) {
                                                    fmt.Println(v83, v22, v21)
                                                    if true {
                                                        v84 := GenS15{F0: "hello", F1: v2, F2: 42}
                                                        v22 = "\n\t"
                                                        fmt.Println(v23, v57, v23)
                                                        v85 := v3[fi52]
                                                        v86 := GenS17{F0: ("foo bar" + "\n\t"), F1: 0}
                                                        if len(v1) > 0 {
                                                            v1[0] = 1
                                                        }
                                                        v87 := GenS18{F0: false, F1: (v51 <= 255)}
                                                        var v88 int16 = v21
                                                        v89 := v3[-9223372036854775808]
                                                        _ = v84
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    } else {
                                                        var v90 uint64 = (^(v23))
                                                        var v91 int64 = (^(v83))
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                } else {
                                                    v57 = (v57 << ((255) & 63))
                                                    v92 := GenS1{F0: 4294967295, F1: -618, F2: v21}
                                                    _ = v92
                                                }
                                                _ = v83
                                            } else {
                                                v23 = v23
                                            }
                                            var v93 int16 = int16(fi52)
                                            v22 = "hello"
                                            var v94 int8 = v2
                                            _ = v93
                                            _ = v94
                                        }
                                    case "\n\t":
                                        switch (-(fi52)) {
                                        case 128:
                                            switch (fi52 & fi53) {
                                            case 2147483647:
                                                var v95 int32 = 2147483647
                                                _ = v95
                                            case -504:
                                                v96 := GenS5{F0: (^(v57))}
                                                _ = v96
                                            case 4294967295:
                                                _ = len(v1)
                                            default:
                                                v51 += 223
                                            }
                                        case -233:
                                            if true {
                                                switch ((v22 + v22) + ((v22 + "\n\t") + v22)) {
                                                case "a":
                                                    v97 := GenS15{F0: (v22 + "a"), F1: (v2 / (((v2 / ((v2) | 1))) | 1)), F2: -1}
                                                    _ = v97
                                                case "":
                                                    v98 := v3[fi53]
                                                    _ = v98
                                                }
                                                var v99 bool = true
                                                v100 := GenS19{F0: v23, F1: (v58 - 259)}
                                                _ = v100
                                                _ = v99
                                            }
                                        case 2147483647:
                                            switch v22 {
                                            case "a":
                                                v58 -= (v58 >> ((255) & 63))
                                            default:
                                                fmt.Println(v2)
                                            }
                                        }
                                    default:
                                        switch fi52 {
                                        case 954:
                                            if len(v1) > 0 {
                                                v1[0] = v58
                                            }
                                        case 32767:
                                            fmt.Println(v57)
                                        default:
                                            v101 := GenS14{F0: ((-(fi53)) >= (^(fi52))), F1: -1}
                                            _ = v101
                                        }
                                    }
                                }
                            case 9223372036854775807:
                                switch v22 {
                                case "":
                                    var v102 uint8 = 0
                                    _ = v102
                                }
                            }
                        } else {
                            v22 = (((v22 + v22) + "\n\t") + (v22 + ("" + v22)))
                        }
                        _ = v59
                    }
                    _ = v57
                    _ = v58
                case 0:
                    var v103 uint32 = 727
                    _ = v103
                }
                var v104 uint64 = 490
                fmt.Println(v21, v51)
                switch fi52 {
                case 1:
                    if false {
                        v105 := GenS20{F0: 255, F1: false, F2: 255}
                        var v106 float64 = (0.0 + 0.0)
                        _ = v105
                        _ = v106
                    }
                case 2147483647:
                    var v107 uint32 = 2
                    _ = v107
                }
                _ = v104
            } else {
                v23 += 2
                v22 = ("" + (v22 + v22))
            }
            var v108 float64 = 3.14159
            _ = fi53
            _ = v108
        }
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = fi52
    }
    for fi109 := 0; fi109 < 9; fi109++ {
        v22 = v22
        _ = fi109
    }
    var v110 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v110
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v51
}
