package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 uint8; F1 uint32; F2 string }
type GenS9 struct { F0 string; F1 uint64; F2 uint64 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS16 struct { F0 uint8; F1 bool }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint; F1 int16; F2 int32 }
func genFn0(p0 int8, p1 uint8, p2 int) string {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int16 = -290
        v6 := GenS1{F0: 1}
        v7 := GenS1{F0: 42}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (("foo bar" + ("hello" + "")) + ("a" + ((("a" + ("" + "")) + "a") + ("hello" + "foo bar"))))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2, 2147483648, 4294967295}
    var v3 uint32 = 4294967295
    if ("\n\t" < "hello") {
        v1[(v3 ^ (^(v3)))] = (v3 ^ 1)
        _ = len(v1)
    }
    if false {
        v8 := GenS1{F0: -9223372036854775808}
        v9 := GenS2{F0: 32768, F1: true}
        _ = v8
        _ = v9
    } else {
        v3 += 2147483647
        delete(v1, (v3 % (((v3 | (^(v3)))) | 1)))
    }
    var v10 int32 = 42
    v11 := GenS3{F0: 255, F1: 2}
    var v12 uint16 = 1
    v1[1] = (v3 - v3)
    switch ("" + ("foo bar" + "a")) {
    case "a":
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                v10 += 629
                fmt.Println(v12)
            } else {
                v12 += v12
                if false {
                    v3 += (v3 + 42)
                }
            }
            v3 -= v3
            _ = fi13
        }
    case "":
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: -986, F1: (false || (0.0 == 0.0))}
            _ = fi14
            _ = v15
        }
    case "hello":
        switch 42 {
        case 1:
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(v3, v3, v10)
                var v17 bool = (3.14159 == (float64((1e10 / (3.14159 / (float32(1.5) / ((0.1 - 0.0) * 1e10))))) + (0.0 * 0.0)))
                if (false && true) {
                    v18 := v1[(^(v3))]
                    _ = v18
                } else {
                    switch (fi16 % ((128) | 1)) {
                    case 128:
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v2 = append(v2, (v3 - v3))
                            switch -169 {
                            case -1:
                                if v17 {
                                    switch -1 {
                                    case 1:
                                        switch ("hello" + "hello") {
                                        case "\n\t":
                                            var v20 string = "\n\t"
                                            _ = v20
                                        }
                                    case 32767:
                                        fmt.Println(v3, fi19)
                                    case 128:
                                        var v21 float64 = (-2.25 / (-2.25 * 1e10))
                                        _ = v21
                                    }
                                    fmt.Println(v3, v3, fi19)
                                    switch fi19 {
                                    case 4294967295:
                                        v17 = (519 < (v10 << ((818) & 63)))
                                    }
                                }
                            case 127:
                                var v22 int64 = 49
                                _ = v22
                            }
                            var v23 int64 = -1
                            _ = fi19
                            _ = v23
                        }
                    case -9223372036854775808:
                        v17 = v17
                    case 263:
                        var v24 float32 = (1.5 * -2.25)
                        _ = v24
                    default:
                        if v17 {
                            fmt.Println(v10, v17, fi16)
                        } else {
                            switch 0 {
                            case 4294967295:
                                fmt.Println(v17, v17, v3)
                            case 0:
                                fmt.Println(v10, v10)
                            case -9223372036854775808:
                                switch fi16 {
                                case -715:
                                    for fi25 := 0; fi25 < 3; fi25++ {
                                        v17 = (!false)
                                        switch ("a" + ("foo bar" + "a")) {
                                        case "":
                                            if true {
                                                fmt.Println(fi25, fi25)
                                            } else {
                                                var v26 uint = 42
                                                v12 = (v12 << (((^(v26))) & 63))
                                                if (fi16 == int(v3)) {
                                                    var v27 int16 = 1000
                                                    if ((65535 <= (^(v26))) || (v27 != (v27 & v27))) {
                                                        fmt.Println(v26, v27, v26)
                                                        var v28 int32 = 42
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                } else {
                                                    var v29 uint32 = (v3 ^ 65535)
                                                    v30 := GenS4{F0: 32767, F1: 127}
                                                    _ = v29
                                                    _ = v30
                                                }
                                                _ = v26
                                            }
                                        case "hello":
                                            if (!(!(65535 == 2))) {
                                                if ((-2.25 + (float32(1.5) + 0.1)) > (-2.25 / float32(1.5))) {
                                                    var v31 uint8 = 42
                                                    _ = v31
                                                } else {
                                                    if v17 {
                                                        v32 := GenS5{F0: ("a" < ("" + "hello")), F1: (fi16 ^ (fi16 % ((fi25) | 1)))}
                                                        var v33 int = fi25
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        var v34 bool = (!v17)
                                                        var v35 int32 = v10
                                                        _ = len(v1)
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    v36 := GenS3{F0: 128, F1: (v12 - v12)}
                                                    _ = v36
                                                }
                                                if v17 {
                                                    var v37 float32 = 1.5
                                                    switch "a" {
                                                    case "hello":
                                                        if false {
                                                            if true {
                                                                v17 = (!v17)
                                                                v38 := GenS6{F0: (v12 >> ((42) & 63))}
                                                                var v39 int16 = -32768
                                                                var v40 float64 = ((-2.25 / 0.1) - 1e10)
                                                                var v41 int64 = 0
                                                                _ = v38
                                                                _ = v39
                                                                _ = v40
                                                                _ = v41
                                                            } else {
                                                                delete(v1, v3)
                                                                if true {
                                                                    if ((1 == 0) || (!v17)) {
                                                                        if (32767 <= v12) {
                                                                            var v42 uint32 = 383
                                                                            _ = v42
                                                                        } else {
                                                                            v43 := GenS7{F0: 127, F1: 228}
                                                                            _ = v43
                                                                        }
                                                                    }
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if false {
                                                                            v10 = v10
                                                                            v17 = (false && false)
                                                                            v10 += 42
                                                                            var v44 float32 = (v37 * 1e10)
                                                                            if (v17 && v17) {
                                                                                if (69 == 6) {
                                                                                    if (v10 != 81) {
                                                                                        if (!v17) {
                                                                                            var v45 uint64 = 346
                                                                                            var v46 uint32 = v3
                                                                                            _ = v45
                                                                                            _ = v46
                                                                                        } else {
                                                                                            if (0 < 100) {
                                                                                                _ = len(v2)
                                                                                            } else {
                                                                                                fmt.Println(fi16, v10, v3)
                                                                                                var v47 int64 = 1
                                                                                                v12 = (^(v12))
                                                                                                _ = v47
                                                                                            }
                                                                                            var v48 int32 = (v10 % (((v10 << ((255) & 63))) | 1))
                                                                                            _ = v48
                                                                                        }
                                                                                        v49 := GenS7{F0: 1, F1: 2}
                                                                                        fmt.Println(fi16)
                                                                                        _ = v49
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v37 = v37
                                                                                    v12 = (v12 - v12)
                                                                                    delete(v1, (^(v3)))
                                                                                    v50 := GenS4{F0: 42, F1: 1}
                                                                                    var v51 int32 = (v10 - (v10 * (v10 + (-(v10)))))
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                            _ = v44
                                                                        } else {
                                                                            var v52 int = fi16
                                                                            switch (fi16 & 128) {
                                                                            case 9223372036854775807:
                                                                                v37 = 0.0
                                                                            case 415:
                                                                                v53 := GenS3{F0: 0, F1: v12}
                                                                                _ = v53
                                                                            case 127:
                                                                                _ = len(v2)
                                                                            default:
                                                                                v10 -= (v10 + (v10 & v10))
                                                                            }
                                                                            if v17 {
                                                                                fmt.Println(v10, v12, v37)
                                                                                if v17 {
                                                                                    if ((2147483648 > (v3 << ((1) & 63))) || (-1 <= (v10 % ((int32(fi16)) | 1)))) {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            var v54 int8 = int8(v52)
                                                                                            _ = v54
                                                                                        case "\n\t":
                                                                                            if ((!v17) || v17) {
                                                                                                var v55 string = ""
                                                                                                var v56 int = fi25
                                                                                                v10 = v10
                                                                                                _ = v55
                                                                                                _ = v56
                                                                                            }
                                                                                        case "foo bar":
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                v10 = v10
                                                                                            case "a":
                                                                                                var v57 uint = 421
                                                                                                _ = v57
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    var v58 int = 42
                                                                                                    _ = v58
                                                                                                } else {
                                                                                                    var v59 uint = 721
                                                                                                    v60 := GenS8{F0: uint8(v3), F1: (^(v3)), F2: "hello"}
                                                                                                    _ = v59
                                                                                                    _ = v60
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v10 = (v10 + v10)
                                                                                        }
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    v3 = 2147483648
                                                                                } else {
                                                                                    v61 := GenS1{F0: 1}
                                                                                    _ = len(v2)
                                                                                    var v62 uint = 18446744073709551615
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                }
                                                                            } else {
                                                                                v12 = (v12 % ((v12) | 1))
                                                                                if (1 < 315) {
                                                                                    var v63 uint32 = 2
                                                                                    _ = len(v1)
                                                                                    if (!(false && true)) {
                                                                                        v64 := GenS3{F0: 2, F1: v12}
                                                                                        _ = v64
                                                                                    } else {
                                                                                        fmt.Println(v12)
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v3 | v63)
                                                                                        }
                                                                                    }
                                                                                    _ = v63
                                                                                } else {
                                                                                    var v65 int16 = 42
                                                                                    if (!(!(int32(v12) < -532))) {
                                                                                        v66 := GenS9{F0: "", F1: 490, F2: 527}
                                                                                        switch fi25 {
                                                                                        case 0:
                                                                                            if v17 {
                                                                                                delete(v1, (v3 - 2147483648))
                                                                                            }
                                                                                        default:
                                                                                            v65 -= (-(v65))
                                                                                        }
                                                                                        fmt.Println(fi16, v52, v37)
                                                                                        _ = v66
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v3)
                                                                                            if (v65 < 1) {
                                                                                                v65 -= v65
                                                                                                fmt.Println(v52)
                                                                                                v1[(v3 + v3)] = v3
                                                                                                switch (fi16 << ((18446744073709551615) & 63)) {
                                                                                                case -38:
                                                                                                    var v67 uint64 = 67
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = uint32(v12)
                                                                                                    }
                                                                                                    var v68 bool = v17
                                                                                                    _ = v67
                                                                                                    _ = v68
                                                                                                case 0:
                                                                                                    var v69 uint8 = 209
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        v1[4294967295] = (v3 % ((v3) | 1))
                                                                                                    case "\n\t":
                                                                                                        if (95 > 100) {
                                                                                                            if true {
                                                                                                                _ = len(v1)
                                                                                                                var v70 bool = v17
                                                                                                                _ = v70
                                                                                                            }
                                                                                                            var v71 float32 = v37
                                                                                                            var v72 int32 = -845
                                                                                                            _ = v71
                                                                                                            _ = v72
                                                                                                        }
                                                                                                    case "a":
                                                                                                        fmt.Println(v12, v52)
                                                                                                    }
                                                                                                    _ = v69
                                                                                                default:
                                                                                                    v52 -= (fi25 / (((fi16 / ((fi25) | 1))) | 1))
                                                                                                }
                                                                                            }
                                                                                            if (128 != 70) {
                                                                                                if v17 {
                                                                                                    var v73 int32 = (-(v10))
                                                                                                    if (v52 == fi16) {
                                                                                                        v74 := v1[(^(v3))]
                                                                                                        fmt.Println(v12, v73)
                                                                                                        var v75 uint32 = (^(v3))
                                                                                                        _ = v74
                                                                                                        _ = v75
                                                                                                    }
                                                                                                    v76 := v1[v3]
                                                                                                    _ = v73
                                                                                                    _ = v76
                                                                                                } else {
                                                                                                    v77 := GenS4{F0: 1, F1: -1}
                                                                                                    var v78 uint16 = 2
                                                                                                    _ = v77
                                                                                                    _ = v78
                                                                                                }
                                                                                                fmt.Println(fi16, fi16)
                                                                                            } else {
                                                                                                var v79 int16 = (v65 | (v65 | (v65 % ((v65) | 1))))
                                                                                                _ = v79
                                                                                            }
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                            var v80 string = ""
                                                                                            _ = v80
                                                                                        }
                                                                                        var v81 uint32 = v3
                                                                                        _ = v81
                                                                                    }
                                                                                    v82 := GenS6{F0: v12}
                                                                                    _ = v65
                                                                                    _ = v82
                                                                                }
                                                                                fmt.Println(v3)
                                                                            }
                                                                            _ = v52
                                                                        }
                                                                    case "\n\t":
                                                                        if ((0 == 128) && (9223372036854775807 < -316)) {
                                                                            if (!(v37 >= v37)) {
                                                                                if true {
                                                                                    v83 := GenS10{F0: 10, F1: (v10 + -1)}
                                                                                    _ = len(v2)
                                                                                    switch ("foo bar" + "hello") {
                                                                                    case "":
                                                                                        fmt.Println(fi25)
                                                                                    }
                                                                                    _ = v83
                                                                                } else {
                                                                                    v17 = v17
                                                                                    if v17 {
                                                                                        v2 = append(v2, (v3 << ((18446744073709551615) & 63)))
                                                                                        var v84 uint8 = 157
                                                                                        if (!v17) {
                                                                                            var v85 int32 = v10
                                                                                            _ = v85
                                                                                        } else {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            if v17 {
                                                                                                var v86 uint32 = v3
                                                                                                _ = v86
                                                                                            } else {
                                                                                                v87 := GenS4{F0: 964, F1: -83}
                                                                                                _ = v87
                                                                                            }
                                                                                        }
                                                                                        _ = v84
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v88 string = (("hello" + ("hello" + "foo bar")) + "a")
                                                                            v88 = "\n\t"
                                                                            fmt.Println(fi25, fi25, v12)
                                                                            var v89 uint32 = v3
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    case "":
                                                                        v90 := GenS4{F0: -861, F1: -1}
                                                                        _ = v90
                                                                    }
                                                                    fmt.Println(fi25)
                                                                } else {
                                                                    fmt.Println(v3, fi25)
                                                                    v12 += v12
                                                                    var v91 int32 = 487
                                                                    _ = v91
                                                                }
                                                                switch (("foo bar" + ("\n\t" + ("" + "a"))) + "hello") {
                                                                case "\n\t":
                                                                    v92 := GenS11{F0: 32767, F1: 0}
                                                                    _ = v92
                                                                case "hello":
                                                                    var v93 int32 = v10
                                                                    _ = v93
                                                                case "foo bar":
                                                                    v37 = ((v37 / v37) * (v37 + (0.0 - (3.14159 + (v37 / ((float32(((0.1 * 0.1) - (0.1 - (1e10 / -2.25)))) / v37) + v37))))))
                                                                }
                                                            }
                                                            var v94 string = ("hello" + ("hello" + ""))
                                                            delete(v1, v3)
                                                            _ = v94
                                                        } else {
                                                            if v17 {
                                                                if ((v12 >> ((65535) & 63)) != 2) {
                                                                    v3 = 65535
                                                                    v95 := GenS11{F0: 0, F1: (v3 % ((v3) | 1))}
                                                                    _ = v95
                                                                }
                                                                var v96 uint16 = v12
                                                                v37 = v37
                                                                _ = v96
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v97 int64 = 9223372036854775807
                                                        _ = v97
                                                    }
                                                    _ = len(v2)
                                                    _ = v37
                                                }
                                            }
                                        case "\n\t":
                                            if (54 >= 39) {
                                                fmt.Println(v3, v3)
                                                var v98 int16 = -32768
                                                v12 = 2
                                                _ = v98
                                            } else {
                                                fmt.Println(fi16, v12, fi25)
                                                v10 = 2147483647
                                                switch fi25 {
                                                case 42:
                                                    if (18446744073709551615 != 931) {
                                                        switch 1 {
                                                        case -159:
                                                            _ = len(v1)
                                                        case 1:
                                                            if v17 {
                                                                v10 -= 1
                                                            } else {
                                                                fmt.Println(v3, v3, fi16)
                                                                var v99 int16 = -43
                                                                v10 -= (v10 - v10)
                                                                _ = v99
                                                            }
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            var v100 uint = 65535
                                                            _ = v100
                                                        case "\n\t":
                                                            v2 = append(v2, uint32(v12))
                                                        case "":
                                                            if true {
                                                                var v101 int16 = -32768
                                                                _ = v101
                                                            } else {
                                                                if (v17 || ((1.5 / (1.5 / 1.5)) > 0.0)) {
                                                                    var v102 float32 = 0.0
                                                                    var v103 float32 = (((((v102 * (0.0 * 0.0)) + ((3.14159 * ((v102 * v102) + -2.25)) * 0.0)) + -2.25) / v102) - v102)
                                                                    var v104 bool = (-128 >= 0)
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                fmt.Println(fi16, v17)
                                                                if false {
                                                                    switch (fi25 | fi16) {
                                                                    case 876:
                                                                        switch fi16 {
                                                                        case -315:
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                if (false || v17) {
                                                                                    v105 := GenS3{F0: 0, F1: 0}
                                                                                    _ = len(v1)
                                                                                    var v106 uint8 = 252
                                                                                    _ = v105
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                        case -552:
                                                                            if v17 {
                                                                                delete(v1, 570)
                                                                                v107 := GenS12{F0: v3}
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                    case 32767:
                                                                        switch fi25 {
                                                                        case 0:
                                                                            fmt.Println(v10, fi16)
                                                                            v10 -= 2147483647
                                                                            v17 = (99 != 0)
                                                                            fmt.Println(fi16, v12, v12)
                                                                        case 125:
                                                                            var v108 uint32 = (v3 % ((2147483647) | 1))
                                                                            _ = v108
                                                                        case 9223372036854775807:
                                                                            v109 := GenS10{F0: 42, F1: (v10 * (v10 ^ (v10 % (((v10 ^ (v10 * -1))) | 1))))}
                                                                            _ = v109
                                                                        default:
                                                                            v110 := GenS7{F0: -107, F1: 162}
                                                                            _ = v110
                                                                        }
                                                                    case 2147483647:
                                                                        var v111 int8 = 127
                                                                        _ = v111
                                                                    default:
                                                                        v3 = 760
                                                                    }
                                                                } else {
                                                                    if (!false) {
                                                                        var v112 int = (-(fi25))
                                                                        v112 -= fi25
                                                                        if (false && v17) {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v10 = (v10 + (v10 ^ v10))
                                                                            case "":
                                                                                _ = len(v1)
                                                                            }
                                                                            switch "a" {
                                                                            case "hello":
                                                                                var v113 int = fi16
                                                                                _ = v113
                                                                            case "":
                                                                                if (!(9223372036854775807 > 1)) {
                                                                                    var v114 uint16 = 750
                                                                                    if v17 {
                                                                                        var v115 int8 = int8(v10)
                                                                                        v112 = v112
                                                                                        _ = v115
                                                                                    }
                                                                                    v116 := GenS13{F0: uint64(v114)}
                                                                                    _ = v114
                                                                                    _ = v116
                                                                                } else {
                                                                                    var v117 float32 = (3.14159 + 0.0)
                                                                                    _ = v117
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v112
                                                                    } else {
                                                                        v17 = (("hello" != ("a" + "")) || v17)
                                                                        v1[2] = (v3 ^ v3)
                                                                    }
                                                                    v118 := GenS14{F0: 9223372036854775807, F1: 863}
                                                                    v119 := GenS14{F0: -9223372036854775808, F1: 4294967295}
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                        v120 := GenS15{F0: 1, F1: 4294967295, F2: 0}
                                                        _ = v120
                                                    }
                                                case 128:
                                                    v17 = false
                                                case -678:
                                                    var v121 int16 = -424
                                                    _ = v121
                                                }
                                            }
                                        default:
                                            delete(v1, uint32(v12))
                                        }
                                        var v122 string = ("" + "a")
                                        _ = fi25
                                        _ = v122
                                    }
                                case 127:
                                    var v123 int8 = 0
                                    _ = v123
                                }
                            default:
                                v124 := GenS3{F0: 128, F1: v12}
                                _ = v124
                            }
                            var v125 uint8 = 42
                            _ = v125
                        }
                    }
                }
                _ = fi16
                _ = v17
            }
        }
    default:
        var v126 uint64 = 42
        _ = v126
    }
    for fi127 := 0; fi127 < 1; fi127++ {
        switch fi127 {
        case 204:
            var v128 uint32 = (v3 >> ((255) & 63))
            _ = v128
        }
        for fi129 := 0; fi129 < 18; fi129++ {
            v10 -= (v10 % (((v10 * (v10 * v10))) | 1))
            v130 := v1[65535]
            _ = fi129
            _ = v130
        }
        _ = fi127
    }
    if true {
        v131 := GenS16{F0: 133, F1: true}
        if false {
            fmt.Println(v12)
            if true {
                if (-104 >= 100) {
                    var v132 int64 = -1
                    v133 := GenS2{F0: 1, F1: (!(4294967295 < uint(v12)))}
                    _ = v132
                    _ = v133
                }
            } else {
                switch 280 {
                case 128:
                    if false {
                        v134 := GenS11{F0: 32767, F1: (v3 >> ((1) & 63))}
                        v10 += (v10 - v10)
                        _ = v134
                    } else {
                        var v135 uint32 = (v3 << ((65535) & 63))
                        fmt.Println(v135, v12)
                        _ = v135
                    }
                case 463:
                    if len(v2) > 0 {
                        v2[0] = (^(v3))
                    }
                }
                var v136 int = -9223372036854775808
                _ = v136
            }
        }
        switch 32767 {
        case -476:
            v3 = (v3 % ((0) | 1))
        default:
            var v137 uint16 = 0
            _ = v137
        }
        _ = v131
    } else {
        var v138 int = -721
        _ = v138
    }
    v12 = (v12 << ((1) & 63))
    v3 = (^(v3))
    v3 = v3
    for fi139 := 0; fi139 < 3; fi139++ {
        for fi140 := 0; fi140 < 15; fi140++ {
            _ = len(v2)
            if true {
                var v141 int8 = -110
                _ = len(v2)
                if (false && ((1e10 + (1.5 + (float64(1e10) * (float64(0.1) - (1e10 - 1e10))))) == float64(((0.1 - (3.14159 - 1e10)) * -2.25)))) {
                    switch "\n\t" {
                    case "a":
                        var v142 uint8 = 167
                        _ = v142
                    case "foo bar":
                        v143 := GenS11{F0: -1, F1: v3}
                        fmt.Println(v141, v12, v141)
                        var v144 int = 32768
                        _ = v143
                        _ = v144
                    default:
                        var v145 int64 = 42
                        _ = v145
                    }
                    var v146 int = (fi139 / ((fi139) | 1))
                    v141 = 100
                    _ = v146
                }
                _ = v141
            } else {
                if (5 != 127) {
                    _ = len(v1)
                    var v147 string = ""
                    _ = v147
                }
                var v148 uint = 65535
                if false {
                    var v149 int16 = 42
                    if (!false) {
                        if (false || (42 < 42)) {
                            _ = len(v2)
                        }
                    } else {
                        fmt.Println(fi139, v3, v3)
                        var v150 uint64 = 4294967295
                        _ = v150
                    }
                    _ = v149
                }
                _ = len(v2)
                var v151 uint32 = 766
                _ = v148
                _ = v151
            }
            _ = fi140
        }
        v152 := GenS6{F0: 42}
        _ = fi139
        _ = v152
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    delete(v1, 4294967295)
    var v153 int32 = (v10 / ((1) | 1))
    if false {
        switch "\n\t" {
        case "foo bar":
            if false {
                v153 += 967
                if (150 <= 65535) {
                    fmt.Println(v3)
                }
            } else {
                var v154 uint8 = 0
                switch 209 {
                case 4294967295:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 549:
                    if (false || true) {
                        if (v3 < v3) {
                            v12 -= (v12 ^ v12)
                            switch -1 {
                            case 670:
                                if false {
                                    if true {
                                        var v155 int64 = -211
                                        v156 := GenS15{F0: 2, F1: 4294967295, F2: 2}
                                        switch 32768 {
                                        case 42:
                                            v155 -= (v155 | v155)
                                        }
                                        _ = v155
                                        _ = v156
                                    } else {
                                        var v157 float64 = 3.14159
                                        if (v154 < v154) {
                                            var v158 float32 = (float32(1.5) / (1e10 + (float32(v157) * ((-2.25 + ((-2.25 - 1e10) / 3.14159)) + (3.14159 + (0.0 / (1.5 * 0.1)))))))
                                            switch -89 {
                                            case 444:
                                                switch ("hello" + ("a" + "foo bar")) {
                                                case "a":
                                                    for fi159 := 0; fi159 < 8; fi159++ {
                                                        var v160 int32 = v153
                                                        if (((float32(-2.25) <= ((((v158 * 3.14159) - v158) + v158) + v158)) && true) || true) {
                                                            if true {
                                                                var v161 bool = (127 == 52)
                                                                _ = v161
                                                            }
                                                        }
                                                        _ = fi159
                                                        _ = v160
                                                    }
                                                case "foo bar":
                                                    v12 += (v12 >> ((2) & 63))
                                                case "hello":
                                                    if true {
                                                        v1[v3] = (v3 - v3)
                                                    }
                                                default:
                                                    fmt.Println(v10, v154, v153)
                                                }
                                            case 889:
                                                switch ("\n\t" + ("hello" + "")) {
                                                case "hello":
                                                    if ("a" > "") {
                                                        v162 := GenS17{F0: v12}
                                                        for fi163 := 0; fi163 < 13; fi163++ {
                                                            v164 := GenS18{F0: "\n\t"}
                                                            var v165 int = fi163
                                                            v10 = v10
                                                            _ = fi163
                                                            _ = v164
                                                            _ = v165
                                                        }
                                                        _ = v162
                                                    } else {
                                                        v154 = v154
                                                    }
                                                case "a":
                                                    if (-268 < 0) {
                                                        if (true || (true && false)) {
                                                            v154 = (v154 + (v154 << ((42) & 63)))
                                                            for fi166 := 0; fi166 < 1; fi166++ {
                                                                v12 += 615
                                                                delete(v1, 4294967295)
                                                                fmt.Println(fi166, fi166)
                                                                _ = fi166
                                                            }
                                                        }
                                                        for fi167 := 0; fi167 < 13; fi167++ {
                                                            delete(v1, uint32(v154))
                                                            v154 -= (v154 % ((78) | 1))
                                                            _ = fi167
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v153)
                                                }
                                            default:
                                                v2 = append(v2, (v3 + v3))
                                            }
                                            for fi168 := 0; fi168 < 14; fi168++ {
                                                var v169 uint8 = (v154 / (((v154 + (v154 / ((v154) | 1)))) | 1))
                                                for fi170 := 0; fi170 < 20; fi170++ {
                                                    var v171 float32 = (((v158 - v158) / v158) * (float32(v157) / ((v158 - (-2.25 * 3.14159)) / (1.5 * v158))))
                                                    v172 := GenS19{F0: 42, F1: 1000, F2: -1}
                                                    if false {
                                                        _ = len(v1)
                                                        v173 := GenS12{F0: (v3 | 2)}
                                                        var v174 string = ""
                                                        _ = v173
                                                        _ = v174
                                                    } else {
                                                        var v175 uint32 = (v3 + (v3 / ((v3) | 1)))
                                                        if true {
                                                            v3 = 479
                                                            var v176 int64 = int64(fi168)
                                                            if (v158 != v171) {
                                                                v3 += 454
                                                            }
                                                            _ = v176
                                                        } else {
                                                            var v177 uint32 = 892
                                                            _ = v177
                                                        }
                                                        _ = v175
                                                    }
                                                    switch fi168 {
                                                    case 0:
                                                        v153 -= v153
                                                    default:
                                                        if true {
                                                            fmt.Println(v158, v171)
                                                        }
                                                    }
                                                    _ = fi170
                                                    _ = v171
                                                    _ = v172
                                                }
                                                _ = fi168
                                                _ = v169
                                            }
                                            _ = v158
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            var v178 float64 = 1.5
                                            _ = v178
                                        }
                                        _ = v157
                                    }
                                }
                            }
                        }
                        fmt.Println(v3, v3)
                    } else {
                        v153 += int32(v154)
                        if false {
                            var v179 bool = ((!(-1 < 315)) && (255 > uint(v3)))
                            v12 -= v12
                            if false {
                                for fi180 := 0; fi180 < 12; fi180++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi180
                                }
                                delete(v1, v3)
                            } else {
                                var v181 uint8 = (^(v154))
                                v10 -= (^(v153))
                                _ = v181
                            }
                            _ = v179
                        }
                        fmt.Println(v12, v10)
                    }
                case 32767:
                    v3 = 2147483648
                default:
                    var v182 bool = true
                    _ = v182
                }
                v1[(v3 + v3)] = v3
                _ = v154
            }
        case "":
            v12 = (v12 - v12)
        default:
            fmt.Println(v10, v3)
        }
        var v183 int32 = 42
        _ = v183
    } else {
        for fi184 := 0; fi184 < 20; fi184++ {
            v10 = v10
            _ = fi184
        }
        if false {
            v185 := v1[(v3 * v3)]
            _ = v185
        } else {
            v3 = (v3 >> ((534) & 63))
            if ((v10 % (((v10 & (v10 / (((v153 & (v10 ^ v153))) | 1)))) | 1)) <= v10) {
                _ = len(v1)
                for fi186 := 0; fi186 < 20; fi186++ {
                    fmt.Println(v12, v153)
                    delete(v1, 1)
                    var v187 int16 = -32768
                    _ = fi186
                    _ = v187
                }
            }
            v188 := GenS18{F0: ""}
            _ = v188
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v153
    _ = v2
    _ = v3
}
