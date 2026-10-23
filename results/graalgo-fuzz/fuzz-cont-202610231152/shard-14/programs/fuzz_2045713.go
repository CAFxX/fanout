package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 uint16; F1 string; F2 uint16 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 uint8; F1 string; F2 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 uint16; F2 int8 }
type GenS16 struct { F0 string; F1 bool; F2 int8 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int64; F1 int64 }
type GenS20 struct { F0 uint64; F1 uint8 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 int16; F1 uint; F2 uint32 }
type GenS23 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16, p2 string) int8 {
    if true {
        var v7 int16 = 29
        p2 = (p2 + ((((p2 + p2) + p2) + p2) + (p2 + ((p2 + p2) + p2))))
        switch 42 {
        case 32768:
            v8 := []int{-2147483648, -9223372036854775808, 127, 608}
            _ = v8
        case 0:
            for fi9 := 0; fi9 < 19; fi9++ {
                p2 = p2
                _ = fi9
            }
        }
        _ = v7
    } else {
        v10 := []int{292, -1, 4294967295}
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
        _ = v10
    }
    var v11 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 127
}
func genFn1() uint64 {
    var v12 uint32 = 2
    v13 := GenS1{F0: (^(v12)), F1: -984}
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int8 = -1
    var v4 float64 = (0.0 + -2.25)
    v5 := []string{("foo bar" + "hello"), "\n\t", "hello", "hello"}
    v6 := make(map[uint32]string)
    v4 -= (v4 - 3.14159)
    v4 += v4
    v4 += 0.1
    var v14 uint8 = 2
    if (!((^(v3)) > v3)) {
        delete(v1, 65535)
        v15 := GenS2{F0: "a"}
        var v16 string = "hello"
        _ = v15
        _ = v16
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v14, v4)
    case "a":
        var v17 bool = true
        _ = v17
    }
    switch "a" {
    case "a":
        for fi18 := 0; fi18 < 12; fi18++ {
            if len(v5) > 0 {
                v5[0] = ""
            }
            var v19 int32 = -2147483648
            _ = fi18
            _ = v19
        }
    }
    v3 = -128
    v3 += 45
    if false {
        switch 128 {
        case 991:
            v20 := GenS3{F0: 121}
            _ = v20
        case 794:
            for fi21 := 0; fi21 < 2; fi21++ {
                v4 -= ((((v4 / (1.5 * (((v4 + ((float64(3.14159) * v4) - 0.0)) / (float64(float32(v4)) / (((v4 - 0.1) / 1e10) * v4))) - float64(-2.25)))) / float64(-2.25)) - v4) - float64((0.1 - (float32(3.14159) * 1e10))))
                var v22 uint32 = 2147483648
                if false {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        if false {
                            fmt.Println(fi23, fi21)
                        } else {
                            if (!(!false)) {
                                if true {
                                    v5 = append(v5, "")
                                    var v24 float64 = (-2.25 - 1.5)
                                    v4 -= v24
                                    if false {
                                        var v25 int64 = 9223372036854775807
                                        if (72 <= (v14 & v14)) {
                                            if (uint32(v14) != (v22 >> ((1) & 63))) {
                                                var v26 float32 = float32(v4)
                                                var v27 float64 = float64(-2.25)
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 uint32 = (v22 << ((42) & 63))
                                                var v29 uint16 = 1
                                                _ = len(v1)
                                                _ = v28
                                                _ = v29
                                            }
                                            v24 = (v24 / 1.5)
                                            v24 += 1.5
                                        }
                                        v14 -= 2
                                        if true {
                                            v4 = (v4 * v4)
                                            v1[v22] = ""
                                            v30 := GenS4{F0: 42, F1: (v3 % ((v3) | 1))}
                                            v25 = -9223372036854775808
                                            v31 := GenS5{F0: 42}
                                            v22 += v22
                                            fmt.Println(v22, v14)
                                            _ = v30
                                            _ = v31
                                        } else {
                                            v32 := GenS3{F0: 127}
                                            switch "\n\t" {
                                            case "a":
                                                if false {
                                                    v33 := GenS1{F0: v22, F1: 42}
                                                    var v34 float64 = v4
                                                    v14 -= v14
                                                    var v35 bool = (false || true)
                                                    switch "hello" {
                                                    case "hello":
                                                        var v36 uint64 = 2
                                                        _ = v36
                                                    default:
                                                        v4 -= (float64((3.14159 - float32(-2.25))) / (v34 + 0.0))
                                                    }
                                                    v14 = (^(v14))
                                                    v37 := v6[(^(v22))]
                                                    switch "\n\t" {
                                                    case "":
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case "\n\t":
                                                        var v38 int = (fi21 << ((2) & 63))
                                                        _ = v38
                                                    case "foo bar":
                                                        v39 := GenS6{F0: (v22 << ((19) & 63)), F1: (!(true && (0.0 == 3.14159)))}
                                                        _ = v39
                                                    }
                                                    v4 -= (v34 + v34)
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                    _ = v37
                                                } else {
                                                    var v40 float64 = v24
                                                    if true {
                                                        var v41 float64 = v40
                                                        v41 -= 1e10
                                                        switch (("foo bar" + "\n\t") + "foo bar") {
                                                        case "a":
                                                            _ = len(v2)
                                                        case "\n\t":
                                                            v40 = v24
                                                        case "hello":
                                                            _ = len(v2)
                                                        }
                                                        var v42 uint8 = 214
                                                        switch 0 {
                                                        case -66:
                                                            var v43 int32 = -2147483648
                                                            _ = v43
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        fmt.Println(v42, fi23, v4)
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        if false {
                                                            if false {
                                                                v44 := GenS3{F0: fi23}
                                                                var v45 bool = (!(255 >= 384))
                                                                v46 := GenS5{F0: (-(v3))}
                                                                _ = v44
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                _ = len(v1)
                                                                v4 = ((((v40 - -2.25) - v40) / (v24 + (float64(1.5) - float64(-2.25)))) / (v24 * v4))
                                                                var v47 uint16 = 0
                                                                _ = v47
                                                            }
                                                        }
                                                        if false {
                                                            fmt.Println(fi23, v25, v24)
                                                            v48 := GenS6{F0: v22, F1: false}
                                                            if true {
                                                                v22 = (v22 ^ v22)
                                                                if false {
                                                                    v40 -= float64(-2.25)
                                                                    var v49 uint8 = (v14 ^ v14)
                                                                    _ = v49
                                                                } else {
                                                                    v50 := GenS1{F0: v22, F1: 390}
                                                                    v25 += 923
                                                                    _ = v50
                                                                }
                                                                fmt.Println(fi21)
                                                                var v51 string = "foo bar"
                                                                _ = v51
                                                            }
                                                            _ = v48
                                                        } else {
                                                            v4 = 3.14159
                                                            v24 -= v4
                                                            fmt.Println(v25, v14, v40)
                                                        }
                                                    }
                                                    v52 := GenS2{F0: ("\n\t" + ("" + "foo bar"))}
                                                    var v53 int64 = v25
                                                    v54 := GenS7{F0: 42, F1: 42}
                                                    _ = v40
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            default:
                                                fmt.Println(v14, v3)
                                            }
                                            if true {
                                                if ((^(v25)) != (v25 % ((int64(v22)) | 1))) {
                                                    v55 := GenS8{F0: true}
                                                    var v56 float32 = -2.25
                                                    _ = v55
                                                    _ = v56
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v24 = (0.1 * (v4 - (v4 - v4)))
                                            } else {
                                                if true {
                                                    _ = len(v5)
                                                }
                                                if (!(v14 == 0)) {
                                                    var v57 uint64 = 916
                                                    if false {
                                                        if (fi23 >= (fi21 / (((fi21 / ((-276) | 1))) | 1))) {
                                                            if len(v5) > 0 {
                                                                v5[0] = ("hello" + "foo bar")
                                                            }
                                                            if true {
                                                                var v58 int32 = -350
                                                                _ = v58
                                                            }
                                                            if (1.5 > -2.25) {
                                                                v3 = genFn0(255, 0, "foo bar")
                                                                v2["hello"] = (^(fi23))
                                                                if true {
                                                                    var v59 uint16 = uint16(v57)
                                                                    if (!false) {
                                                                        v25 -= (v25 - (v25 & v25))
                                                                        if (-1 < 32767) {
                                                                            _ = len(v5)
                                                                            var v60 bool = (v22 < 455)
                                                                            var v61 int16 = 0
                                                                            _ = v60
                                                                            _ = v61
                                                                        } else {
                                                                            var v62 int16 = 1
                                                                            if true {
                                                                                var v63 int = (fi23 & (fi21 | (fi21 << ((42) & 63))))
                                                                                v64 := v1[v22]
                                                                                _ = v63
                                                                                _ = v64
                                                                            } else {
                                                                                if true {
                                                                                    if ((v4 * v4) > (float64((3.14159 / 1.0)) / v24)) {
                                                                                        fmt.Println(fi23)
                                                                                    }
                                                                                    v25 = (^(v25))
                                                                                    if (!true) {
                                                                                        var v65 bool = false
                                                                                        var v66 int16 = v62
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = "a"
                                                                                        }
                                                                                        switch int(v25) {
                                                                                        case -2147483648:
                                                                                            fmt.Println(fi21, v57, v24)
                                                                                            var v67 int8 = v3
                                                                                            _ = v67
                                                                                        case 4294967295:
                                                                                            fmt.Println(fi23)
                                                                                        case 32767:
                                                                                            v66 -= 1
                                                                                        }
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                }
                                                                                var v68 int = (fi21 / ((9223372036854775807) | 1))
                                                                                _ = v68
                                                                            }
                                                                            v69 := GenS9{F0: 4294967295, F1: 42}
                                                                            _ = v62
                                                                            _ = v69
                                                                        }
                                                                        _ = len(v6)
                                                                        delete(v6, 532)
                                                                    }
                                                                    v70 := GenS10{F0: (v59 >> ((4294967295) & 63)), F1: "a", F2: 294}
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    if (3.14159 != 1e10) {
                                                                        var v71 int16 = 1000
                                                                        _ = v71
                                                                    } else {
                                                                        v22 -= v22
                                                                        var v72 int8 = (v3 + (v3 | 96))
                                                                        v3 = v3
                                                                        if (-1 == 26) {
                                                                            v73 := GenS11{F0: true, F1: (fi21 % (((fi23 ^ 127)) | 1))}
                                                                            _ = v73
                                                                        } else {
                                                                            var v74 uint8 = v14
                                                                            fmt.Println(v22, v57)
                                                                            _ = v74
                                                                        }
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "":
                                                                            if false {
                                                                                var v75 int8 = (v72 >> ((42) & 63))
                                                                                _ = len(v5)
                                                                                var v76 uint32 = (v22 + v22)
                                                                                _ = v75
                                                                                _ = v76
                                                                            }
                                                                        case "\n\t":
                                                                            switch 0 {
                                                                            case 1:
                                                                                v14 += 0
                                                                            default:
                                                                                v6[2] = ("\n\t" + ((("\n\t" + "\n\t") + "a") + ("" + ("foo bar" + ("foo bar" + (((("a" + "hello") + "a") + "hello") + ("a" + "foo bar")))))))
                                                                            }
                                                                        default:
                                                                            if len(v5) > 0 {
                                                                                v5[0] = ""
                                                                            }
                                                                        }
                                                                        _ = v72
                                                                    }
                                                                    _ = v59
                                                                    _ = v70
                                                                } else {
                                                                    var v77 string = (("foo bar" + "a") + "hello")
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v78 uint16 = 65535
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            }
                                                        }
                                                        if (v22 == 2) {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v79 := GenS2{F0: "hello"}
                                                                _ = v79
                                                            default:
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v80 int16 = -1
                                                                    _ = v80
                                                                case "foo bar":
                                                                    if ((-(fi23)) == fi21) {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = "foo bar"
                                                                        }
                                                                        if ((^(v25)) >= (v25 + (v25 | 42))) {
                                                                            var v81 float32 = 0.1
                                                                            v82 := v1[v22]
                                                                            switch ((v82 + v82) + (((("a" + v82) + "\n\t") + ((v82 + "a") + (v82 + v82))) + (v82 + "a"))) {
                                                                            case "foo bar":
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case "\n\t":
                                                                                v83 := GenS12{F0: (-32768 > 140), F1: (!true)}
                                                                                _ = v83
                                                                            case "a":
                                                                                var v84 float64 = v4
                                                                                _ = v84
                                                                            }
                                                                            _ = len(v5)
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            var v85 uint64 = v57
                                                                            var v86 uint = 774
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                        v4 -= 0.0
                                                                        if ((!false) || false) {
                                                                            switch fi23 {
                                                                            case 32767:
                                                                                v24 += ((0.0 + (1e10 * (v24 - v4))) * 0.0)
                                                                            case 883:
                                                                                fmt.Println(v24, v3, v24)
                                                                            case 9223372036854775807:
                                                                                if (2 <= 127) {
                                                                                    var v87 float32 = float32(((v24 - v24) + v24))
                                                                                    v57 += v57
                                                                                    v2["hello"] = (fi23 ^ (fi23 + 32767))
                                                                                    _ = v87
                                                                                }
                                                                            default:
                                                                                v5 = append(v5, "foo bar")
                                                                            }
                                                                            v88 := GenS13{F0: (v14 + (v14 ^ v14)), F1: "foo bar", F2: -729}
                                                                            v89 := GenS14{F0: 1000}
                                                                            v22 = (v22 >> ((810) & 63))
                                                                            _ = v88
                                                                            _ = v89
                                                                        } else {
                                                                            v24 += 0.1
                                                                            if ((v57 / (((^(v57))) | 1)) >= (v57 ^ 42)) {
                                                                                var v90 uint64 = (^(v57))
                                                                                v4 += 3.14159
                                                                                _ = v90
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                    }
                                                                    var v91 int32 = 0
                                                                    _ = v91
                                                                default:
                                                                    var v92 int16 = int16(fi21)
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        var v93 uint8 = v14
                                                        v93 -= (v14 >> ((4294967295) & 63))
                                                        _ = v93
                                                    }
                                                    _ = v57
                                                } else {
                                                    v94 := GenS15{F0: 1000, F1: 0, F2: (v3 & (v3 << ((1) & 63)))}
                                                    fmt.Println(fi21)
                                                    _ = v94
                                                }
                                                _ = len(v6)
                                            }
                                            _ = v32
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            } else {
                                var v95 float32 = (((float32((v4 + 1e10)) * 1e10) + (0.1 - (float32((1e10 * 1.5)) / -2.25))) * (-2.25 / 1.5))
                                switch "foo bar" {
                                case "hello":
                                    if false {
                                        if ((18446744073709551615 <= 2) && false) {
                                            v3 += (-(v3))
                                        }
                                    } else {
                                        v96 := GenS5{F0: v3}
                                        var v97 int16 = -32768
                                        switch (fi23 / ((fi23) | 1)) {
                                        case 127:
                                            if false {
                                                fmt.Println(fi23, fi23)
                                                fmt.Println(v14, v14)
                                            }
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = "a"
                                            }
                                        }
                                        _ = v96
                                        _ = v97
                                    }
                                case "\n\t":
                                    if ((846 > -32768) && false) {
                                        var v98 uint64 = 0
                                        var v99 uint = 724
                                        switch ("foo bar" + ("foo bar" + "\n\t")) {
                                        case "\n\t":
                                            v95 = (-2.25 - 0.1)
                                        case "hello":
                                            switch "a" {
                                            case "a":
                                                v100 := GenS7{F0: (v98 - (v98 - v98)), F1: 32767}
                                                _ = v100
                                            case "\n\t":
                                                v14 = v14
                                            case "foo bar":
                                                var v101 int16 = 1000
                                                _ = v101
                                            default:
                                                if true {
                                                    var v102 float64 = 0.1
                                                    _ = v102
                                                }
                                            }
                                        default:
                                            v3 -= (v3 >> ((v99) & 63))
                                        }
                                        _ = v98
                                        _ = v99
                                    }
                                case "foo bar":
                                    if false {
                                        v3 = (v3 << ((1) & 63))
                                        switch "" {
                                        case "":
                                            if len(v5) > 0 {
                                                v5[0] = "foo bar"
                                            }
                                        case "hello":
                                            switch ("a" + "foo bar") {
                                            case "foo bar":
                                                fmt.Println(fi23)
                                            case "\n\t":
                                                v22 += (v22 & (v22 + 0))
                                            case "a":
                                                v103 := GenS14{F0: 42}
                                                _ = v103
                                            }
                                        default:
                                            var v104 uint64 = 18
                                            _ = v104
                                        }
                                    }
                                    v105 := v1[v22]
                                    fmt.Println(v105, v14, v22)
                                    _ = v105
                                }
                                _ = v95
                            }
                            v5 = append(v5, "hello")
                            fmt.Println(v22, v22, v14)
                            delete(v1, v22)
                            if (false || (18446744073709551615 > 2)) {
                                var v106 float64 = 0.0
                                v106 -= v106
                                var v107 uint16 = 255
                                var v108 int32 = 42
                                _ = v106
                                _ = v107
                                _ = v108
                            } else {
                                v22 = (^(v22))
                                if len(v5) > 0 {
                                    v5[0] = "hello"
                                }
                                var v109 uint8 = (v14 * uint8(v22))
                                _ = v109
                            }
                            switch (fi21 * 506) {
                            case 127:
                                v110 := GenS16{F0: "foo bar", F1: true, F2: v3}
                                _ = v110
                            case 680:
                                v111 := GenS17{F0: (v22 | (v22 * v22))}
                                _ = v111
                            default:
                                v6[(v22 * v22)] = ((("foo bar" + "a") + ("" + ("a" + ("a" + ("" + "a"))))) + "\n\t")
                            }
                        }
                        _ = fi23
                    }
                    var v112 bool = (fi21 != (fi21 & fi21))
                    _ = v112
                }
                _ = fi21
                _ = v22
            }
        case 1:
            if false {
                if true {
                    var v113 float64 = (v4 + 3.14159)
                    _ = v113
                } else {
                    v114 := GenS5{F0: genFn0(32767, -1, (("\n\t" + "") + "hello"))}
                    if (false || true) {
                        _ = len(v1)
                        var v115 uint = 2
                        if (!true) {
                            v5 = append(v5, "hello")
                            v115 -= 294
                            switch "foo bar" {
                            case "hello":
                                var v116 uint = 255
                                _ = v116
                            case "a":
                                v117 := v1[2147483647]
                                _ = v117
                            case "foo bar":
                                if (!(!false)) {
                                    v4 += (v4 / 3.14159)
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = v3
                                    case "a":
                                        var v118 uint32 = 2147483648
                                        _ = v118
                                    case "":
                                        v14 += 127
                                    default:
                                        for fi119 := 0; fi119 < 8; fi119++ {
                                            if true {
                                                for fi120 := 0; fi120 < 19; fi120++ {
                                                    fmt.Println(v14, v4, v3)
                                                    v115 += (v115 / ((v115) | 1))
                                                    fmt.Println(v14, v115)
                                                    _ = fi120
                                                }
                                                for fi121 := 0; fi121 < 1; fi121++ {
                                                    var v122 int16 = 1000
                                                    fmt.Println(v122)
                                                    _ = len(v6)
                                                    if (float32(v4) >= 3.14159) {
                                                        var v123 uint = (v115 & (v115 * v115))
                                                        v5 = append(v5, "\n\t")
                                                        v115 -= v123
                                                        _ = v123
                                                    } else {
                                                        v124 := GenS14{F0: v122}
                                                        _ = v124
                                                    }
                                                    v125 := GenS5{F0: int8(v14)}
                                                    v126 := v2["hello"]
                                                    _ = fi121
                                                    _ = v122
                                                    _ = v125
                                                    _ = v126
                                                }
                                                for fi127 := 0; fi127 < 9; fi127++ {
                                                    var v128 int32 = -1
                                                    v4 += v4
                                                    _ = fi127
                                                    _ = v128
                                                }
                                            }
                                            if false {
                                                var v129 string = ("foo bar" + "foo bar")
                                                _ = v129
                                            }
                                            for fi130 := 0; fi130 < 15; fi130++ {
                                                if false {
                                                    v131 := GenS18{F0: 4294967295}
                                                    v14 = 0
                                                    v132 := GenS1{F0: 2, F1: 2147483647}
                                                    _ = v131
                                                    _ = v132
                                                }
                                                if true {
                                                    fmt.Println(v3, v3)
                                                    var v133 string = "foo bar"
                                                    _ = v133
                                                } else {
                                                    var v134 int8 = v3
                                                    var v135 bool = false
                                                    _ = v134
                                                    _ = v135
                                                }
                                                var v136 float32 = (-2.25 - float32(0.0))
                                                _ = fi130
                                                _ = v136
                                            }
                                            _ = fi119
                                        }
                                    }
                                    for fi137 := 0; fi137 < 8; fi137++ {
                                        var v138 float32 = (0.0 / -2.25)
                                        var v139 uint32 = 1
                                        _ = fi137
                                        _ = v138
                                        _ = v139
                                    }
                                } else {
                                    if (-427 <= 42) {
                                        v140 := GenS19{F0: -857, F1: 9223372036854775807}
                                        var v141 int = -428
                                        v142 := GenS20{F0: 0, F1: v14}
                                        _ = v140
                                        _ = v141
                                        _ = v142
                                    }
                                }
                            default:
                                if (!(!(2 == (v14 * v14)))) {
                                    switch 1 {
                                    case -1:
                                        var v143 int8 = v3
                                        _ = v143
                                    case -306:
                                        for fi144 := 0; fi144 < 6; fi144++ {
                                            if (fi144 == fi144) {
                                                v4 = 1.5
                                                v1[2147483647] = "\n\t"
                                                var v145 int8 = v3
                                                _ = v145
                                            } else {
                                                var v146 int16 = -32768
                                                if ((^(v3)) < (v3 % ((23) | 1))) {
                                                    var v147 int = (fi144 / ((fi144) | 1))
                                                    var v148 uint64 = genFn1()
                                                    _ = v147
                                                    _ = v148
                                                } else {
                                                    v4 = v4
                                                    switch 4294967295 {
                                                    case 32768:
                                                        var v149 uint16 = 2
                                                        _ = v149
                                                    case 42:
                                                        for fi150 := 0; fi150 < 5; fi150++ {
                                                            v4 -= v4
                                                            _ = fi150
                                                        }
                                                    default:
                                                        switch (^(fi144)) {
                                                        case 1:
                                                            v3 -= 127
                                                        case 127:
                                                            v14 += (v14 << ((253) & 63))
                                                        case 9223372036854775807:
                                                            fmt.Println(v146, fi144, v115)
                                                        }
                                                    }
                                                }
                                                switch 166 {
                                                case -403:
                                                    for fi151 := 0; fi151 < 14; fi151++ {
                                                        switch int(v146) {
                                                        case 128:
                                                            v152 := GenS11{F0: false, F1: -2147483648}
                                                            _ = v152
                                                        case -9223372036854775808:
                                                            v4 += v4
                                                        default:
                                                            switch (fi151 * fi144) {
                                                            case -236:
                                                                v153 := GenS1{F0: 4294967295, F1: -591}
                                                                _ = v153
                                                            case -810:
                                                                if (!(0 != uint32(v115))) {
                                                                    var v154 uint64 = 1
                                                                    v155 := GenS21{F0: 1}
                                                                    _ = len(v5)
                                                                    _ = v154
                                                                    _ = v155
                                                                } else {
                                                                    var v156 int16 = 32767
                                                                    _ = v156
                                                                }
                                                                if false {
                                                                    v157 := GenS18{F0: 804}
                                                                    v158 := GenS22{F0: v146, F1: (v115 | (v115 / ((0) | 1))), F2: 42}
                                                                    if (((v146 >> ((v115) & 63)) != (v146 ^ -1)) || false) {
                                                                        fmt.Println(fi151, v4)
                                                                    }
                                                                    v6[1] = ("hello" + "a")
                                                                    _ = v157
                                                                    _ = v158
                                                                }
                                                            case 32767:
                                                                var v159 uint64 = 2
                                                                _ = v159
                                                            }
                                                        }
                                                        _ = fi151
                                                    }
                                                }
                                                _ = v146
                                            }
                                            fmt.Println(v3)
                                            _ = fi144
                                        }
                                    default:
                                        fmt.Println(v3, v14, v115)
                                    }
                                    _ = len(v5)
                                    fmt.Println(v14, v14, v14)
                                }
                            }
                        }
                        _ = v115
                    }
                    v160 := GenS7{F0: 232, F1: 255}
                    _ = v114
                    _ = v160
                }
                v14 += 1
                _ = len(v5)
            }
        }
        v4 += (float64((0.1 - -2.25)) + v4)
    } else {
        v161 := GenS23{F0: 0}
        _ = v161
    }
    switch "hello" {
    case "hello":
        fmt.Println(v14, v4, v4)
    case "":
        var v162 uint64 = 42
        _ = v162
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = ("a" + "hello")
        }
    }
    var v163 int32 = 132
    v14 = 42
    for fi164 := 0; fi164 < 6; fi164++ {
        fmt.Println(v3)
        for fi165 := 0; fi165 < 5; fi165++ {
            if true {
                var v166 bool = false
                _ = v166
            } else {
                var v167 float64 = (3.14159 / v4)
                v14 -= (v14 / (((^(v14))) | 1))
                _ = v167
            }
            v6[768] = ((("hello" + "hello") + "\n\t") + ("a" + (("foo bar" + ("a" + (((("" + "hello") + "a") + "hello") + "a"))) + (("hello" + "hello") + ("foo bar" + ("foo bar" + "\n\t"))))))
            v168 := GenS8{F0: (false && (1 == 1))}
            _ = fi165
            _ = v168
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi164
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v163:", v163)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
