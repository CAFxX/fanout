package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int8; F1 int; F2 int8 }
type GenS8 struct { F0 uint16; F1 int8; F2 int }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 bool }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 int64; F1 int16 }
type GenS16 struct { F0 int64; F1 string; F2 int64 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int64; F1 uint; F2 int }
type GenS22 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint) int32 {
    var v6 float32 = float32(1e10)
    var v7 uint = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1() int64 {
    var v8 uint64 = 0
    v8 += 476
    _ = v8
    return -1
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 540
    var v3 uint = 4294967295
    v4 := []string{"a"}
    var v5 uint8 = 1
    v2 -= v2
    switch 127 {
    case 32768:
        v2 += v2
    case -743:
        var v9 uint8 = (^(v5))
        _ = v9
    case 9223372036854775807:
        switch 239 {
        case 2147483647:
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint16 = 255
                _ = fi10
                _ = v11
            }
        }
    }
    v12 := GenS1{F0: -67, F1: ("a" + ((("hello" + "hello") + "\n\t") + ("" + "hello"))), F2: 32767}
    v13 := GenS1{F0: 1, F1: ("a" + "hello"), F2: -208}
    var v14 string = ""
    v5 = (v5 | v5)
    var v15 int = 9223372036854775807
    fmt.Println(v2)
    fmt.Println(v14, v15)
    v16 := GenS1{F0: 0, F1: ((v14 + v14) + (v14 + (v14 + "a"))), F2: 1}
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v14, v5, fi17)
        if (!false) {
            switch v14 {
            case "hello":
                v2 = (v2 | v2)
            case "foo bar":
                v15 -= 430
            default:
                if (uint16(v2) <= 802) {
                    if (!false) {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            var v19 uint32 = v2
                            var v20 uint8 = v5
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            if false {
                                v15 -= 1
                                var v22 string = (v14 + v14)
                                _ = v22
                            }
                            v23 := GenS1{F0: -1, F1: v14, F2: 32767}
                            v24 := GenS1{F0: 74, F1: "", F2: 256}
                            _ = fi21
                            _ = v23
                            _ = v24
                        }
                        v1 = ((v1 - v1) + v1)
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if true {
                                var v26 float64 = float64(v1)
                                v26 = v26
                                var v27 float64 = -2.25
                                switch (fi17 << ((65535) & 63)) {
                                case 317:
                                    if (false && false) {
                                        var v28 int = (^(fi17))
                                        v26 += float64((((float32(1.5) * v1) * -2.25) / float32((float64((v1 + v1)) * v26))))
                                        switch (^(fi25)) {
                                        case 2147483647:
                                            if len(v4) > 0 {
                                                v4[0] = v14
                                            }
                                        case 1:
                                            if ((v15 ^ fi25) >= (-(v15))) {
                                                if false {
                                                    v3 += v3
                                                    switch (v14 + v14) {
                                                    case "":
                                                        var v29 uint8 = v5
                                                        _ = v29
                                                    default:
                                                        var v30 int16 = 38
                                                        _ = v30
                                                    }
                                                    switch v14 {
                                                    case "hello":
                                                        v15 += fi17
                                                    case "\n\t":
                                                        v3 += v3
                                                    default:
                                                        v14 = v14
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = (((v14 + v14) + v14) + (v14 + v14))
                                                    }
                                                    v31 := GenS1{F0: 35, F1: "hello", F2: -32768}
                                                    _ = v31
                                                } else {
                                                    v32 := GenS2{F0: 28, F1: v5, F2: (v3 & (v3 & (v3 & 0)))}
                                                    switch ("hello" + ("a" + v14)) {
                                                    case "hello":
                                                        v2 = (^(v2))
                                                    case "foo bar":
                                                        var v33 string = v14
                                                        _ = v33
                                                    case "a":
                                                        v1 = v1
                                                        switch (v14 + ("a" + ("" + v14))) {
                                                        case "":
                                                            if (4294967295 >= 367) {
                                                                if (374 <= 746) {
                                                                    switch (v15 * 128) {
                                                                    case 0:
                                                                        if false {
                                                                            v15 = (v28 + (fi25 + (fi25 % ((v28) | 1))))
                                                                            if (true || (!true)) {
                                                                                if true {
                                                                                    if true {
                                                                                        if (false && true) {
                                                                                            v34 := GenS3{F0: 973, F1: v3, F2: -76}
                                                                                            _ = v34
                                                                                        }
                                                                                        v35 := GenS2{F0: 0, F1: (v5 % (((v5 - (v5 - (v5 << ((v3) & 63))))) | 1)), F2: v3}
                                                                                        _ = v35
                                                                                    } else {
                                                                                        if (true && false) {
                                                                                            var v36 uint = uint(v2)
                                                                                            fmt.Println(v36, v36, v15)
                                                                                            if (-2147483648 < 1) {
                                                                                                fmt.Println(v14)
                                                                                            }
                                                                                            if (2147483647 != v2) {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                v37 := GenS1{F0: -128, F1: ("\n\t" + v14), F2: 1}
                                                                                                _ = v37
                                                                                            }
                                                                                            switch v14 {
                                                                                            case "foo bar":
                                                                                                var v38 uint32 = (v2 / (((^(v2))) | 1))
                                                                                                _ = v38
                                                                                            }
                                                                                            v28 = (v15 % ((128) | 1))
                                                                                            if false {
                                                                                                v26 = v26
                                                                                                v5 = (v5 ^ 255)
                                                                                                if (v14 >= v14) {
                                                                                                    var v39 uint = 0
                                                                                                    var v40 uint8 = v5
                                                                                                    if (false || false) {
                                                                                                        fmt.Println(v2, v14)
                                                                                                    }
                                                                                                    _ = v39
                                                                                                    _ = v40
                                                                                                } else {
                                                                                                    var v41 float32 = -2.25
                                                                                                    fmt.Println(v36, v36, v41)
                                                                                                    var v42 string = v14
                                                                                                    _ = v41
                                                                                                    _ = v42
                                                                                                }
                                                                                            } else {
                                                                                                v5 = v5
                                                                                            }
                                                                                            v3 = 65535
                                                                                            switch v15 {
                                                                                            case -769:
                                                                                                v15 += (fi25 % (((-(v28))) | 1))
                                                                                            case -2147483648:
                                                                                                if (315 < 255) {
                                                                                                    if ((v36 * (v3 % ((v3) | 1))) < (v3 * v36)) {
                                                                                                        v15 += 2147483647
                                                                                                    }
                                                                                                    _ = len(v4)
                                                                                                    var v43 int = v15
                                                                                                    v26 -= 0.0
                                                                                                    _ = v43
                                                                                                }
                                                                                            default:
                                                                                                switch 4294967295 {
                                                                                                case 4294967295:
                                                                                                    v44 := GenS4{F0: -61}
                                                                                                    _ = v44
                                                                                                case -319:
                                                                                                    if false {
                                                                                                        v28 = (v15 * v15)
                                                                                                    }
                                                                                                    var v45 bool = (-68 <= 0)
                                                                                                    _ = v45
                                                                                                }
                                                                                            }
                                                                                            v3 = v36
                                                                                            fmt.Println(v5)
                                                                                            _ = v36
                                                                                        } else {
                                                                                            if true {
                                                                                                v3 = (v3 | (v3 >> ((86) & 63)))
                                                                                                if true {
                                                                                                    v5 = (v5 << ((v3) & 63))
                                                                                                    switch ((v14 + "\n\t") + "a") {
                                                                                                    case "a":
                                                                                                        var v46 uint16 = 2
                                                                                                        _ = v46
                                                                                                    }
                                                                                                    var v47 uint64 = 0
                                                                                                    _ = v47
                                                                                                }
                                                                                                fmt.Println(v3)
                                                                                            }
                                                                                            v1 += ((v1 + float32(3.14159)) - ((3.14159 / float32(((0.0 - v26) * v27))) * (3.14159 - (v1 + v1))))
                                                                                        }
                                                                                        v26 = v26
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    var v48 int = (fi25 & fi17)
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    _ = v48
                                                                                }
                                                                            }
                                                                            v28 -= (fi17 + -2147483648)
                                                                        }
                                                                        v28 = v28
                                                                    case -325:
                                                                        v4 = append(v4, "")
                                                                    case 42:
                                                                        if true {
                                                                            if (true && (v14 <= v14)) {
                                                                                v49 := GenS5{F0: (v2 % (((^(v2))) | 1)), F1: (^(v3)), F2: v3}
                                                                                if false {
                                                                                    v50 := GenS6{F0: 100, F1: 128}
                                                                                    var v51 float32 = v1
                                                                                    fmt.Println(v5)
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                } else {
                                                                                    var v52 float64 = v27
                                                                                    v1 += v1
                                                                                    v53 := GenS7{F0: 100, F1: -1, F2: -95}
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                                _ = v49
                                                                            } else {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                fmt.Println(v26, v3)
                                                                            }
                                                                            v54 := GenS8{F0: 32767, F1: -121, F2: fi25}
                                                                            _ = v54
                                                                        }
                                                                    default:
                                                                        fmt.Println(v26, v1)
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi17, v15)
                                                                    var v55 uint64 = 18446744073709551615
                                                                    _ = v55
                                                                }
                                                                var v56 uint8 = v5
                                                                if true {
                                                                    v14 = v14
                                                                    v15 = 2147483647
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                _ = v56
                                                            }
                                                        case "hello":
                                                            fmt.Println(v15, v3, v15)
                                                            fmt.Println(v2, fi25)
                                                            var v57 int32 = int32(v2)
                                                            _ = v57
                                                        case "a":
                                                            var v58 int = fi17
                                                            _ = v58
                                                        }
                                                        var v59 string = ("hello" + "foo bar")
                                                        _ = v59
                                                    }
                                                    if true {
                                                        v60 := GenS9{F0: 42, F1: 2}
                                                        _ = len(v4)
                                                        v15 = (v15 + (fi25 << ((uint(v5)) & 63)))
                                                        if false {
                                                            fmt.Println(v27, v28, v28)
                                                        } else {
                                                            v61 := GenS1{F0: 1, F1: v14, F2: -457}
                                                            _ = v61
                                                        }
                                                        _ = v60
                                                    }
                                                    _ = v32
                                                }
                                                v62 := GenS3{F0: 42, F1: v3, F2: 100}
                                                v5 = v5
                                                _ = v62
                                            }
                                        default:
                                            if true {
                                                var v63 bool = false
                                                _ = v63
                                            } else {
                                                v64 := GenS10{F0: (v5 | 201), F1: 42}
                                                v14 = "a"
                                                var v65 int64 = genFn1()
                                                _ = v64
                                                _ = v65
                                            }
                                        }
                                        _ = v28
                                    }
                                case 778:
                                    switch (fi25 / (((v15 * 0)) | 1)) {
                                    case 32767:
                                        if len(v4) > 0 {
                                            v4[0] = v14
                                        }
                                    case 582:
                                        if false {
                                            _ = len(v4)
                                            var v66 uint = 987
                                            _ = v66
                                        } else {
                                            if ((^(v2)) > (v2 << ((v3) & 63))) {
                                                var v67 float32 = v1
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v68 float64 = 3.14159
                                                _ = v67
                                                _ = v68
                                            }
                                            if true {
                                                if (0 != 32767) {
                                                    if true {
                                                        v3 += v3
                                                        v27 = float64(float32((v27 - (((v26 / (float64(v1) / v26)) / 1.0) - 0.0))))
                                                    } else {
                                                        switch (^(fi17)) {
                                                        case 127:
                                                            v69 := GenS11{F0: 226, F1: genFn0(v2, 4294967295)}
                                                            _ = v69
                                                        case -930:
                                                            var v70 uint8 = (v5 << ((266) & 63))
                                                            _ = v70
                                                        }
                                                        v14 = "a"
                                                    }
                                                    switch ((v14 + "hello") + "") {
                                                    case "hello":
                                                        if (false && (!(-32768 < -1))) {
                                                            var v71 uint = v3
                                                            fmt.Println(fi17)
                                                            if (152 >= uint8(v2)) {
                                                                switch v14 {
                                                                case "foo bar":
                                                                    v72 := GenS12{F0: v5}
                                                                    _ = v72
                                                                case "hello":
                                                                    var v73 int = (v15 % ((42) | 1))
                                                                    var v74 bool = true
                                                                    _ = v73
                                                                    _ = v74
                                                                case "a":
                                                                    var v75 int8 = 72
                                                                    _ = v75
                                                                }
                                                                v76 := GenS13{F0: -579, F1: (!false)}
                                                                v77 := GenS8{F0: 0, F1: 0, F2: (v15 ^ v15)}
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                            _ = v71
                                                        }
                                                    default:
                                                        var v78 uint8 = (v5 ^ v5)
                                                        _ = v78
                                                    }
                                                }
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = v14
                                                }
                                                var v79 uint16 = 42
                                                fmt.Println(v2)
                                                if false {
                                                    v80 := GenS14{F0: (fi17 + fi17), F1: (v2 / (((v2 - (^(v2)))) | 1))}
                                                    v81 := GenS15{F0: genFn1(), F1: 32767}
                                                    if false {
                                                        if true {
                                                            v26 = 0.0
                                                        } else {
                                                            var v82 int16 = 42
                                                            fmt.Println(v15)
                                                            _ = v82
                                                        }
                                                        switch v14 {
                                                        case "\n\t":
                                                            v1 -= (float32(v27) - (float32(float64((v1 + (1.5 * -2.25)))) * -2.25))
                                                        case "foo bar":
                                                            v2 += 826
                                                        case "hello":
                                                            v1 = v1
                                                        default:
                                                            v83 := GenS10{F0: v5, F1: 1}
                                                            _ = v83
                                                        }
                                                        var v84 uint = v3
                                                        _ = v84
                                                    } else {
                                                        if false {
                                                            fmt.Println(v5, v26)
                                                            fmt.Println(fi25, fi17)
                                                            fmt.Println(v26)
                                                            v85 := GenS16{F0: 9223372036854775807, F1: (v14 + ((("foo bar" + "a") + v14) + ("a" + v14))), F2: -989}
                                                            _ = v85
                                                        } else {
                                                            v27 -= float64(3.14159)
                                                            var v86 uint64 = 2
                                                            v87 := GenS9{F0: 189, F1: (v79 / (((v79 & 42)) | 1))}
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        var v88 uint32 = v2
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v88
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    v1 -= (3.14159 - ((v1 * v1) - 1.5))
                                                }
                                                v5 += v5
                                                if (!true) {
                                                    var v89 uint64 = 0
                                                    _ = v89
                                                }
                                                _ = v79
                                            }
                                            v1 = float32((((v26 / 3.14159) - float64(float32(((-2.25 - v27) + v27)))) - (v27 - (v26 / (0.1 / v26)))))
                                        }
                                    }
                                }
                                var v90 int8 = 127
                                var v91 uint16 = 42
                                _ = v26
                                _ = v27
                                _ = v90
                                _ = v91
                            } else {
                                if (false || true) {
                                    var v92 float32 = (v1 - v1)
                                    switch fi25 {
                                    case 42:
                                        v2 = (v2 + v2)
                                    case -1:
                                        v14 = "hello"
                                    case 128:
                                        v93 := GenS12{F0: 0}
                                        _ = v93
                                    default:
                                        if true {
                                            switch "foo bar" {
                                            case "foo bar":
                                                _ = len(v4)
                                            case "a":
                                                v3 = (^(v3))
                                            case "":
                                                if (v2 < (v2 | v2)) {
                                                    if ((528 <= 609) && (!true)) {
                                                        if (((v1 - -2.25) - 3.14159) < v1) {
                                                            var v94 int16 = 1
                                                            v94 = v94
                                                            _ = v94
                                                        }
                                                        var v95 int8 = -1
                                                        _ = v95
                                                    } else {
                                                        v3 = 42
                                                        var v96 int8 = -12
                                                        _ = v96
                                                    }
                                                } else {
                                                    var v97 int16 = 266
                                                    if (int16(v2) <= (v97 * v97)) {
                                                        if (v97 >= int16(v3)) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            if (-32768 >= (^(v97))) {
                                                                switch ("hello" + v14) {
                                                                case "\n\t":
                                                                    v3 -= uint(v5)
                                                                case "hello":
                                                                    fmt.Println(v5, v1)
                                                                default:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = "\n\t"
                                                                    }
                                                                }
                                                                v98 := GenS1{F0: 49, F1: ((v14 + v14) + v14), F2: v97}
                                                                _ = v98
                                                            } else {
                                                                switch v14 {
                                                                case "\n\t":
                                                                    v4 = append(v4, ((v14 + v14) + v14))
                                                                case "foo bar":
                                                                    fmt.Println(v5, fi25, v97)
                                                                case "":
                                                                    var v99 int32 = 1
                                                                    _ = v99
                                                                }
                                                                var v100 string = (v14 + ((v14 + (v14 + v14)) + "hello"))
                                                                _ = v100
                                                            }
                                                            var v101 int64 = 0
                                                            _ = v101
                                                        } else {
                                                            v102 := GenS8{F0: uint16(v2), F1: -1, F2: fi17}
                                                            var v103 int32 = 0
                                                            if ((v5 + (v5 % ((v5) | 1))) < 127) {
                                                                if ((0.0 + (1e10 / v92)) >= float32(1e10)) {
                                                                    var v104 int8 = -90
                                                                    _ = v104
                                                                }
                                                            }
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        var v105 string = ((((((v14 + "hello") + "hello") + "a") + "a") + v14) + (v14 + "hello"))
                                                        _ = v105
                                                    } else {
                                                        if ((false && false) && false) {
                                                            v106 := GenS2{F0: 42, F1: v5, F2: v3}
                                                            v107 := GenS17{F0: 454}
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    }
                                                    _ = v97
                                                }
                                            default:
                                                switch v14 {
                                                case "foo bar":
                                                    fmt.Println(v2, fi25)
                                                }
                                            }
                                            var v108 bool = true
                                            fmt.Println(v108, v3)
                                            _ = v108
                                        }
                                    }
                                    if false {
                                        var v109 uint32 = v2
                                        v4 = append(v4, "\n\t")
                                        _ = v109
                                    }
                                    fmt.Println(v2, v2)
                                    _ = v92
                                } else {
                                    v15 += (v15 ^ -9223372036854775808)
                                }
                                v14 = "hello"
                                var v110 uint32 = (^(v2))
                                v5 = v5
                                _ = v110
                            }
                            if true {
                                v14 = (v14 + v14)
                                switch v15 {
                                case 268:
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(v5)
                                    case "foo bar":
                                        if true {
                                            v111 := GenS12{F0: (v5 - 127)}
                                            _ = len(v4)
                                            _ = v111
                                        }
                                    }
                                case -1:
                                    var v112 float64 = (0.1 / float64((v1 + 3.14159)))
                                    _ = v112
                                case 105:
                                    v2 = (v2 * (v2 % (((v2 + (v2 % ((v2) | 1)))) | 1)))
                                }
                            }
                            v1 = v1
                            _ = fi25
                        }
                    }
                }
            }
            v113 := GenS4{F0: 94}
            _ = v113
        }
        v114 := GenS18{F0: -411}
        _ = fi17
        _ = v114
    }
    v1 = v1
    if (v1 >= v1) {
        if ((1 <= 0) && true) {
            if true {
                v115 := GenS19{F0: (v5 + v5)}
                v116 := GenS2{F0: uint64(v2), F1: (v5 >> ((255) & 63)), F2: v3}
                _ = v115
                _ = v116
            } else {
                switch v15 {
                case -1:
                    for fi117 := 0; fi117 < 7; fi117++ {
                        var v118 uint64 = 472
                        if true {
                            if false {
                                v15 = v15
                                fmt.Println(v5, v14, v15)
                                v119 := GenS9{F0: 32767, F1: 65535}
                                _ = v119
                            } else {
                                for fi120 := 0; fi120 < 14; fi120++ {
                                    fmt.Println(fi117, v5)
                                    v118 = uint64(v3)
                                    v121 := GenS20{F0: v2}
                                    _ = fi120
                                    _ = v121
                                }
                                for fi122 := 0; fi122 < 1; fi122++ {
                                    switch "hello" {
                                    case "foo bar":
                                        if false {
                                            v2 += uint32(v118)
                                        }
                                        if true {
                                            v1 = v1
                                        }
                                        switch -249 {
                                        case 128:
                                            var v123 int32 = -1
                                            _ = v123
                                        case 987:
                                            v124 := GenS21{F0: 521, F1: 18446744073709551615, F2: fi122}
                                            _ = v124
                                        case 1:
                                            v125 := GenS17{F0: (v118 << ((v3) & 63))}
                                            _ = v125
                                        default:
                                            if (32767 < 1) {
                                                v15 -= 42
                                            }
                                        }
                                    case "":
                                        var v126 int = (v15 * (v15 - (v15 ^ v15)))
                                        _ = v126
                                    }
                                    var v127 bool = ((!(((!(float32(-2.25) >= v1)) || true) && (false || true))) && (-2.25 > v1))
                                    _ = fi122
                                    _ = v127
                                }
                                v128 := GenS17{F0: 707}
                                _ = v128
                            }
                        } else {
                            v129 := GenS22{F0: 1}
                            switch "\n\t" {
                            case "":
                                var v130 uint16 = 983
                                _ = v130
                            }
                            if true {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                } else {
                                    var v131 int64 = 9223372036854775807
                                    var v132 uint16 = 255
                                    v4 = append(v4, v14)
                                    _ = v131
                                    _ = v132
                                }
                                fmt.Println(v3, v5, v1)
                            }
                            _ = v129
                        }
                        _ = fi117
                        _ = v118
                    }
                case 418:
                    if len(v4) > 0 {
                        v4[0] = v14
                    }
                default:
                    var v133 float32 = 1e10
                    _ = v133
                }
            }
            var v134 int16 = -32768
            v15 -= (v15 / (((v15 / ((v15) | 1))) | 1))
            _ = v134
        }
        if len(v4) > 0 {
            v4[0] = v14
        }
    } else {
        v1 = (-2.25 * v1)
    }
    var v135 float32 = ((v1 * (v1 - (v1 * v1))) * v1)
    v14 = ("hello" + (v14 + "foo bar"))
    var v136 uint32 = uint32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v135
    _ = v136
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
