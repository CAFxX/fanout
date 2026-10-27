package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 uint16; F1 uint8; F2 uint }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 int8; F1 uint; F2 int32 }
type GenS11 struct { F0 int32; F1 bool }
type GenS12 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint16; F1 int32 }
type GenS17 struct { F0 int; F1 int }
type GenS18 struct { F0 int; F1 uint16 }
type GenS19 struct { F0 uint8 }
func genFn0() uint16 {
    switch ("\n\t" + "a") {
    case "foo bar":
        switch -2147483648 {
        case 32768:
            v5 := GenS1{F0: 42, F1: 2147483647, F2: 9223372036854775807}
            _ = v5
        case 256:
            var v6 string = "a"
            _ = v6
        case -1:
            var v7 uint32 = 741
            _ = v7
        }
    }
    var v8 float64 = -2.25
    v9 := GenS2{F0: 42}
    if (!(55 < -106)) {
        var v10 uint64 = 353
        _ = v10
    } else {
        switch 1 {
        case -456:
            v11 := []int{-1, 42, -9223372036854775808, -1}
            _ = v11
        }
        v12 := GenS3{F0: 65535}
        _ = v12
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    var v3 int16 = 196
    v4 := []int{-391, 127, -1}
    if false {
        var v13 uint16 = 993
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(fi14, fi14, v2)
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    v3 = v3
    if ((109 <= 9223372036854775807) && ((((-996 != -38) || false) || (float64(0.0) == -2.25)) || (false && ((v2 + (v2 + v2)) != "a")))) {
        delete(v1, 42)
        v16 := GenS3{F0: genFn0()}
        _ = v16
    } else {
        switch "" {
        case "a":
            v3 = v3
        case "hello":
            if false {
                v1[42] = 4294967295
                for fi17 := 0; fi17 < 10; fi17++ {
                    v3 += (v3 ^ (v3 + v3))
                    fmt.Println(v3, v2, fi17)
                    v3 -= v3
                    _ = fi17
                }
                v4 = append(v4, -2147483648)
            } else {
                var v18 string = ((v2 + v2) + "hello")
                _ = v18
            }
        case "\n\t":
            for fi19 := 0; fi19 < 14; fi19++ {
                if false {
                    switch fi19 {
                    case 2147483647:
                        switch (-(fi19)) {
                        case 32768:
                            if true {
                                v3 -= v3
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    v2 = (v2 + v2)
                                    fmt.Println(fi19, fi20)
                                    v21 := GenS4{F0: (v2 + v2), F1: 1, F2: "\n\t"}
                                    _ = fi20
                                    _ = v21
                                }
                                _ = len(v4)
                            } else {
                                switch (fi19 >> ((61) & 63)) {
                                case 32767:
                                    switch v2 {
                                    case "a":
                                        fmt.Println(v2, v2)
                                    case "foo bar":
                                        var v22 int16 = (v3 / ((v3) | 1))
                                        _ = v22
                                    case "":
                                        for fi23 := 0; fi23 < 7; fi23++ {
                                            v2 = "hello"
                                            var v24 string = v2
                                            v25 := GenS5{F0: v3}
                                            if (false || (false && ((v2 + (v2 + v24)) >= "hello"))) {
                                                var v26 float32 = 0.0
                                                _ = v26
                                            } else {
                                                if false {
                                                    var v27 uint = 1
                                                    _ = v27
                                                }
                                            }
                                            fmt.Println(v3, v3)
                                            var v28 bool = true
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                            _ = v28
                                        }
                                    default:
                                        for fi29 := 0; fi29 < 10; fi29++ {
                                            v30 := GenS2{F0: 163}
                                            _ = fi29
                                            _ = v30
                                        }
                                    }
                                default:
                                    for fi31 := 0; fi31 < 11; fi31++ {
                                        if false {
                                            var v32 uint64 = 617
                                            _ = v32
                                        }
                                        v33 := GenS4{F0: "hello", F1: -271, F2: "a"}
                                        v2 = "foo bar"
                                        if (42 <= 86) {
                                            v3 += (v3 & (v3 >> ((0) & 63)))
                                        }
                                        v2 = "hello"
                                        if (v2 < (v2 + (((((("\n\t" + v2) + v2) + ((v2 + (v2 + v2)) + v2)) + v2) + "\n\t") + "hello"))) {
                                            v3 = (v3 >> ((52) & 63))
                                            var v34 int64 = -739
                                            var v35 float32 = (-2.25 - (1.5 / 0.1))
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = fi31
                                        _ = v33
                                    }
                                }
                            }
                        case 1:
                            v3 += 0
                        default:
                            switch (^(fi19)) {
                            case -2147483648:
                                v2 = v2
                            case 428:
                                v3 = v3
                            }
                        }
                    case -1:
                        v3 += (v3 - (v3 % ((32767) | 1)))
                    default:
                        for fi36 := 0; fi36 < 5; fi36++ {
                            var v37 int16 = (v3 - 1000)
                            switch (fi36 / (((^(fi36))) | 1)) {
                            case -1:
                                var v38 int64 = 42
                                _ = v38
                            case -952:
                                switch (fi19 & 0) {
                                case 32768:
                                    if (false || false) {
                                        if true {
                                            switch 146 {
                                            case 4294967295:
                                                v2 = "a"
                                            case 2147483647:
                                                switch (^(fi19)) {
                                                case 32767:
                                                    var v39 float64 = ((1e10 + -2.25) / 1e10)
                                                    _ = v39
                                                case 42:
                                                    fmt.Println(fi36)
                                                    if (1e10 > 0.0) {
                                                        v4 = append(v4, fi36)
                                                        var v40 uint32 = 0
                                                        var v41 string = ((v2 + v2) + ("" + ("foo bar" + v2)))
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    var v42 string = (("foo bar" + "\n\t") + (v2 + v2))
                                                    _ = v42
                                                case 327:
                                                    v37 = (v3 << ((976) & 63))
                                                default:
                                                    var v43 int16 = 0
                                                    _ = v43
                                                }
                                            }
                                            v37 -= v37
                                        }
                                        v3 -= -605
                                    }
                                case 9223372036854775807:
                                    if ((42 != 686) && (v2 < "hello")) {
                                        switch v2 {
                                        case "foo bar":
                                            switch (v2 + v2) {
                                            case "":
                                                switch v2 {
                                                case "hello":
                                                    v44 := GenS6{F0: "foo bar", F1: v37}
                                                    _ = v44
                                                case "foo bar":
                                                    v37 = v37
                                                case "a":
                                                    var v45 bool = ((float32(-2.25) > 1.5) || false)
                                                    _ = v45
                                                }
                                            case "foo bar":
                                                v1[(^(fi36))] = (fi19 + (fi36 >> ((209) & 63)))
                                            default:
                                                _ = len(v4)
                                            }
                                            fmt.Println(fi19)
                                            var v46 int32 = 1
                                            v37 += v3
                                            v2 = "hello"
                                            if (1e10 < (0.0 * 0.0)) {
                                                v2 = v2
                                                if (!true) {
                                                    fmt.Println(v46)
                                                }
                                                fmt.Println(fi36, v2)
                                            } else {
                                                var v47 float64 = 1.5
                                                var v48 int16 = v37
                                                switch v2 {
                                                case "a":
                                                    fmt.Println(v37, fi19, v47)
                                                    v37 = v3
                                                case "hello":
                                                    v48 = v48
                                                case "":
                                                    switch 4294967295 {
                                                    case -838:
                                                        v49 := GenS3{F0: 2}
                                                        v50 := GenS7{F0: 780, F1: 401}
                                                        switch (^(fi36)) {
                                                        case 32768:
                                                            v51 := GenS8{F0: 65535, F1: 172, F2: 953}
                                                            _ = v51
                                                        case 9223372036854775807:
                                                            v52 := GenS8{F0: 2, F1: 42, F2: 4294967295}
                                                            _ = v52
                                                        case 951:
                                                            var v53 int = fi19
                                                            _ = v53
                                                        default:
                                                            var v54 float32 = 0.0
                                                            _ = v54
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    default:
                                                        if (((v2 + ("foo bar" + "")) + v2) > ("a" + v2)) {
                                                            if (!(v46 < (v46 - (v46 * 2147483647)))) {
                                                                fmt.Println(v48)
                                                            } else {
                                                                fmt.Println(v47, fi19)
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        if true {
                                                            v55 := GenS9{F0: 0, F1: (true && (((1 < 0) && true) || (v47 != v47)))}
                                                            _ = v55
                                                        }
                                                        var v56 uint8 = 127
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(v48, v46)
                                                        fmt.Println(fi19, v37, fi19)
                                                        if (((494 != 4294967295) && false) || false) {
                                                            switch v2 {
                                                            case "\n\t":
                                                                var v57 float32 = ((-2.25 / 3.14159) + 3.14159)
                                                                _ = v57
                                                            case "hello":
                                                                fmt.Println(fi36, v47, fi19)
                                                            default:
                                                                v4 = append(v4, fi19)
                                                            }
                                                            var v58 int16 = (v3 | v3)
                                                            _ = v58
                                                        }
                                                    }
                                                }
                                                var v59 int64 = -1
                                                v60 := GenS8{F0: 255, F1: 0, F2: 690}
                                                _ = v47
                                                _ = v48
                                                _ = v59
                                                _ = v60
                                            }
                                            v61 := GenS10{F0: 42, F1: 954, F2: v46}
                                            _ = v46
                                            _ = v61
                                        default:
                                            if true {
                                                fmt.Println(fi36)
                                                v62 := GenS11{F0: -1, F1: ((1 <= 65535) || (!((1e10 > 1.5) || (2147483648 >= 0))))}
                                                v63 := GenS12{F0: 308, F1: genFn0(), F2: 1}
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                        v2 = (v2 + "\n\t")
                                        v3 += v3
                                    } else {
                                        if (-402 > -1) {
                                            v64 := GenS13{F0: 1}
                                            var v65 float32 = 0.1
                                            _ = v64
                                            _ = v65
                                        }
                                        fmt.Println(fi19)
                                    }
                                }
                            case 4294967295:
                                if (1 != 244) {
                                    switch fi36 {
                                    case 1:
                                        delete(v1, (-(fi36)))
                                    }
                                    if (138 == 42) {
                                        v3 = -94
                                        v66 := GenS14{F0: 32767, F1: -346}
                                        v3 -= -1
                                        v67 := GenS7{F0: (v37 << ((208) & 63)), F1: 1}
                                        _ = v66
                                        _ = v67
                                    } else {
                                        _ = len(v4)
                                        switch "\n\t" {
                                        case "\n\t":
                                            if false {
                                                if (v2 <= (v2 + ((v2 + ("" + v2)) + (v2 + "")))) {
                                                    v68 := v1[(fi19 % ((fi36) | 1))]
                                                    v2 = (((("a" + "a") + "") + ("\n\t" + "\n\t")) + "a")
                                                    _ = len(v4)
                                                    _ = v68
                                                } else {
                                                    var v69 float64 = -2.25
                                                    if (2147483647 < 442) {
                                                        v70 := GenS8{F0: 489, F1: 0, F2: 176}
                                                        v71 := GenS13{F0: 42}
                                                        var v72 float64 = (1.5 - v69)
                                                        v73 := GenS3{F0: 546}
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    var v74 uint16 = 766
                                                    _ = v69
                                                    _ = v74
                                                }
                                                var v75 int64 = -9223372036854775808
                                                _ = v75
                                            } else {
                                                if false {
                                                    v37 -= 840
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        switch fi36 {
                                                        case -1:
                                                            var v76 float32 = 1.5
                                                            _ = v76
                                                        case -2147483648:
                                                            v77 := GenS15{F0: 739}
                                                            _ = v77
                                                        default:
                                                            v3 += -32768
                                                        }
                                                    case "\n\t":
                                                        switch v2 {
                                                        case "foo bar":
                                                            switch fi36 {
                                                            case -9223372036854775808:
                                                                v2 = ("" + v2)
                                                            case 0:
                                                                v78 := GenS7{F0: v3, F1: 4294967295}
                                                                _ = v78
                                                            default:
                                                                v79 := v1[(fi19 * 128)]
                                                                _ = v79
                                                            }
                                                        case "":
                                                            v80 := v1[fi36]
                                                            _ = v80
                                                        }
                                                    case "a":
                                                        v81 := GenS12{F0: 4294967295, F1: 42, F2: 0}
                                                        _ = v81
                                                    default:
                                                        if true {
                                                            v2 = v2
                                                            v82 := GenS4{F0: (("\n\t" + "a") + ("a" + (v2 + v2))), F1: -506, F2: "\n\t"}
                                                            _ = v82
                                                        }
                                                    }
                                                    if true {
                                                        v3 = 1000
                                                        fmt.Println(v3)
                                                    } else {
                                                        switch (("\n\t" + "\n\t") + "") {
                                                        case "foo bar":
                                                            v1[fi36] = fi19
                                                        case "\n\t":
                                                            v2 = v2
                                                        }
                                                        fmt.Println(v2, fi19, v2)
                                                        v83 := GenS15{F0: 65535}
                                                        var v84 int32 = 0
                                                        if (v37 >= v3) {
                                                            var v85 int64 = 42
                                                            if ((!(!(v2 > (v2 + "hello")))) && true) {
                                                                var v86 float32 = 0.1
                                                                _ = v86
                                                            } else {
                                                                v1[(fi19 - (fi19 - fi36))] = 32768
                                                            }
                                                            v87 := GenS16{F0: genFn0(), F1: v84}
                                                            _ = v85
                                                            _ = v87
                                                        } else {
                                                            var v88 float64 = (((0.1 * (0.1 / 3.14159)) / 0.1) / float64((float32((1e10 * 1e10)) + 1e10)))
                                                            v2 = (v2 + v2)
                                                            _ = v88
                                                        }
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    switch v2 {
                                                    case "foo bar":
                                                        var v89 uint = 255
                                                        _ = v89
                                                    }
                                                    var v90 int8 = 0
                                                    _ = v90
                                                }
                                            }
                                        case "a":
                                            var v91 uint64 = 18446744073709551615
                                            _ = v91
                                        }
                                        var v92 uint64 = 913
                                        v3 = v37
                                        var v93 uint8 = 137
                                        if (!(v92 != uint64(v93))) {
                                            switch v2 {
                                            case "foo bar":
                                                if true {
                                                    switch 32767 {
                                                    case 1:
                                                        fmt.Println(fi36, v93)
                                                    case 128:
                                                        fmt.Println(v93, v93)
                                                    case 9223372036854775807:
                                                        fmt.Println(fi36, v37, v37)
                                                    default:
                                                        fmt.Println(v93, v92)
                                                    }
                                                    v94 := GenS1{F0: 255, F1: -216, F2: -840}
                                                    _ = v94
                                                }
                                            case "\n\t":
                                                var v95 int8 = 100
                                                _ = v95
                                            }
                                            fmt.Println(v93, v92, v93)
                                            var v96 int = fi19
                                            _ = v96
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v92
                                        _ = v93
                                    }
                                } else {
                                    switch fi19 {
                                    case -1:
                                        switch v2 {
                                        case "foo bar":
                                            var v97 uint = 0
                                            _ = v97
                                        case "a":
                                            v3 += v37
                                            v1[32767] = 0
                                        default:
                                            switch (v2 + (((v2 + "hello") + (v2 + ("hello" + v2))) + v2)) {
                                            case "\n\t":
                                                v2 = v2
                                            case "":
                                                var v98 uint = 0
                                                v3 = 1000
                                                if true {
                                                    v4 = append(v4, fi36)
                                                    var v99 float64 = (1e10 / (-2.25 + float64(1e10)))
                                                    v37 = (v3 >> (((v98 | (v98 / ((4294967295) | 1)))) & 63))
                                                    _ = v99
                                                }
                                                _ = v98
                                            }
                                        }
                                    }
                                }
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                    if true {
                        v2 = v2
                    }
                }
                v2 = v2
                _ = fi19
            }
        }
        _ = len(v1)
        v1[32767] = 357
    }
    if (true || true) {
        fmt.Println(v3, v2)
        v100 := v1[-2147483648]
        _ = v100
    }
    v101 := GenS17{F0: 329, F1: 9223372036854775807}
    var v102 uint8 = 0
    v103 := GenS18{F0: 837, F1: 2}
    v2 = v2
    v2 = "\n\t"
    v104 := GenS19{F0: (v102 >> ((277) & 63))}
    var v105 uint = 587
    _ = len(v1)
    switch 47 {
    case 981:
        v106 := v1[1]
        _ = v106
    case 42:
        var v107 uint8 = v102
        _ = v107
    case 660:
        var v108 int16 = 42
        _ = v108
    default:
        fmt.Println(v3, v3, v2)
    }
    var v109 uint = 255
    _ = len(v4)
    switch 4294967295 {
    case 32767:
        var v110 int8 = -50
        _ = v110
    case 127:
        v102 = v102
    }
    v1[32767] = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v105:", v105)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v109
    _ = v2
    _ = v3
    _ = v4
}
