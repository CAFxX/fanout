package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 uint32; F2 int }
type GenS9 struct { F0 uint; F1 int32; F2 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 string; F2 uint64 }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64; F1 int64 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 int8; F1 int64; F2 bool }
type GenS17 struct { F0 int64; F1 string }
type GenS18 struct { F0 uint; F1 string }
type GenS19 struct { F0 string; F1 int64 }
type GenS20 struct { F0 uint8; F1 uint; F2 int8 }
type GenS21 struct { F0 int16; F1 int; F2 uint32 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 uint16; F1 uint }
type GenS24 struct { F0 uint16 }
type GenS25 struct { F0 int64; F1 int64 }
type GenS26 struct { F0 int64; F1 uint }
type GenS27 struct { F0 int64; F1 int8 }
type GenS28 struct { F0 bool; F1 int16 }
type GenS29 struct { F0 uint8 }
func genFn0(p0 string, p1 uint) int32 {
    if (4294967295 <= p1) {
        var v5 int8 = -1
        var v6 uint16 = 462
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 0}
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}
func genFn1(p0 uint, p1 uint) int32 {
    var v8 uint = 2
    switch 32767 {
    case 32767:
        var v9 bool = false
        _ = v9
    case 127:
        var v10 int64 = 42
        _ = v10
    default:
        fmt.Println(v8, p1, p0)
    }
    v11 := []int{}
    p1 += (p0 >> (((p0 | (p0 & p1))) & 63))
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := []string{}
    var v12 uint64 = 1
    fmt.Println(v12)
    switch "foo bar" {
    case "hello":
        v12 -= 2
    case "foo bar":
        var v13 uint32 = 65535
        _ = v13
    case "":
        var v14 int16 = 1
        _ = v14
    }
    v1 = v1
    switch 32768 {
    case 2147483647:
        v1 = v1
    case 276:
        v15 := GenS2{F0: 1, F1: 451}
        _ = v15
    case -635:
        if false {
            for fi16 := 0; fi16 < 18; fi16++ {
                v1 = (v1 << ((255) & 63))
                _ = fi16
            }
        }
    }
    if (!true) {
        fmt.Println(v1, v12)
        if (0.0 > 3.14159) {
            switch 32767 {
            case 0:
                v1 += 42
            case 128:
                v17 := GenS2{F0: 51, F1: 9223372036854775807}
                _ = v17
            case -9223372036854775808:
                if true {
                    if (((1e10 - -2.25) / 1.0) > 1.5) {
                        switch 42 {
                        case 42:
                            v12 += (v12 << ((0) & 63))
                        case 4294967295:
                            fmt.Println(v12, v12)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v1 -= v1
                    v4 = append(v4, "hello")
                }
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v20 int16 = 0
                    if true {
                        v12 += v12
                        if (-9223372036854775808 == -1) {
                            v21 := GenS3{F0: 18446744073709551615}
                            v22 := GenS4{F0: false, F1: 65535}
                            v20 -= (v20 / (((v20 >> ((18446744073709551615) & 63))) | 1))
                            _ = v21
                            _ = v22
                        } else {
                            v1 += (v1 << ((255) & 63))
                            var v23 int64 = (v1 >> ((255) & 63))
                            _ = v23
                        }
                        fmt.Println(fi18, v1, v20)
                    }
                    v24 := GenS5{F0: "foo bar", F1: (v1 ^ (v1 % ((v1) | 1)))}
                    switch fi18 {
                    case -1:
                        if false {
                            v25 := GenS4{F0: false, F1: 4294967295}
                            _ = len(v2)
                            _ = v25
                        }
                        if ((fi19 & 2147483647) <= fi19) {
                            var v26 int8 = -128
                            v1 -= (v1 - -9223372036854775808)
                            _ = v26
                        }
                        var v27 int32 = -1
                        _ = v27
                    case 32768:
                        v28 := GenS5{F0: "foo bar", F1: (v1 << ((1) & 63))}
                        _ = v28
                    case 42:
                        v1 += (v1 * (v1 % (((^(v1))) | 1)))
                    default:
                        if true {
                            var v29 string = (("a" + "\n\t") + "foo bar")
                            switch (fi18 << ((18446744073709551615) & 63)) {
                            case -887:
                                var v30 float32 = 3.14159
                                var v31 string = (("a" + v29) + (v29 + (v29 + "\n\t")))
                                v4 = append(v4, (((v29 + v29) + v29) + v29))
                                _ = v30
                                _ = v31
                            case 128:
                                if (-1 <= -635) {
                                    var v32 int = 127
                                    fmt.Println(v12)
                                    _ = v32
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = v29
                                    }
                                }
                            case -25:
                                switch (fi18 / (((fi19 % (((fi18 % ((fi19) | 1))) | 1))) | 1)) {
                                case 128:
                                    fmt.Println(fi18, v20)
                                case -841:
                                    fmt.Println(fi18)
                                default:
                                    var v33 float64 = ((1.5 - (-2.25 / (((((((1.5 - (0.1 * -2.25)) * 0.0) / 1.0) * (-2.25 * 0.1)) + 3.14159) / (1e10 / float64((1.5 * (((0.0 - 1e10) / 3.14159) / 1e10))))) + float64(0.1)))) / (0.1 - float64(float32(1e10))))
                                    _ = v33
                                }
                            }
                            _ = v29
                        } else {
                            v4 = append(v4, ("a" + "hello"))
                        }
                        switch fi18 {
                        case 2147483647:
                            if (!true) {
                                fmt.Println(v20, v12, fi19)
                            }
                        case -1:
                            if true {
                                var v34 uint8 = 127
                                fmt.Println(v1, fi18)
                                if true {
                                    delete(v2, fi19)
                                    fmt.Println(v34)
                                    fmt.Println(fi19)
                                    v1 += v1
                                    v1 = -218
                                }
                                var v35 uint = 2
                                fmt.Println(fi18, v34, v20)
                                v12 = (^(v12))
                                _ = v34
                                _ = v35
                            } else {
                                switch fi18 {
                                case -1:
                                    v3[-253] = (!true)
                                }
                                v1 = 0
                                if false {
                                    if true {
                                        switch (fi18 - fi18) {
                                        case 0:
                                            fmt.Println(fi19)
                                            var v36 uint64 = (v12 & (v12 | v12))
                                            _ = v36
                                        }
                                        if ((-1 < -57) || (!(false && true))) {
                                            v37 := GenS1{F0: 334, F1: 32767}
                                            var v38 int32 = -1
                                            v20 = (v20 ^ v20)
                                            _ = v37
                                            _ = v38
                                        } else {
                                            fmt.Println(fi18, fi19)
                                            if (1e10 != 3.14159) {
                                                switch 1 {
                                                case 128:
                                                    fmt.Println(v12, v1)
                                                case 2147483647:
                                                    v12 = 692
                                                case 0:
                                                    if (v12 == (v12 + (v12 | v12))) {
                                                        if false {
                                                            var v39 uint16 = 42
                                                            v40 := GenS5{F0: "\n\t", F1: v1}
                                                            var v41 int32 = 527
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        if ((uint(v12) == 1) && true) {
                                                            v20 = v20
                                                        } else {
                                                            var v42 float64 = float64(0.0)
                                                            v42 = (v42 * (float64(-2.25) * v42))
                                                            var v43 string = "\n\t"
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        v20 += 42
                                                    } else {
                                                        var v44 uint = 42
                                                        switch 4294967295 {
                                                        case 32767:
                                                            v1 -= v1
                                                        case -529:
                                                            fmt.Println(v20, v1)
                                                        case 0:
                                                            var v45 float64 = (-2.25 + (-2.25 * 0.1))
                                                            _ = v45
                                                        default:
                                                            v46 := GenS6{F0: -2147483648}
                                                            _ = v46
                                                        }
                                                        v1 += v1
                                                        v3[4294967295] = false
                                                        var v47 uint32 = 42
                                                        _ = v44
                                                        _ = v47
                                                    }
                                                    v48 := GenS5{F0: "", F1: int64(fi18)}
                                                    _ = v48
                                                default:
                                                    v49 := GenS7{F0: true}
                                                    _ = v49
                                                }
                                                if true {
                                                    if true {
                                                        v20 += v20
                                                        var v50 int = fi19
                                                        switch (-(fi19)) {
                                                        case 127:
                                                            var v51 uint8 = 128
                                                            _ = v51
                                                        case 42:
                                                            var v52 int64 = (v1 - -9223372036854775808)
                                                            _ = v52
                                                        default:
                                                            v4 = append(v4, "a")
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                            v53 := v2[v50]
                                                            _ = v53
                                                        }
                                                        var v54 bool = ((v12 & (^(v12))) <= 215)
                                                        _ = v50
                                                        _ = v54
                                                    }
                                                    switch (("foo bar" + "\n\t") + "foo bar") {
                                                    case "\n\t":
                                                        var v55 float32 = 0.1
                                                        _ = v55
                                                    case "a":
                                                        switch fi18 {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi19, fi19, fi18)
                                                        case 32768:
                                                            switch 0 {
                                                            case 42:
                                                                v56 := GenS8{F0: 42, F1: 65535, F2: (fi19 << ((255) & 63))}
                                                                _ = v56
                                                            case 914:
                                                                switch int(v1) {
                                                                case 900:
                                                                    v57 := GenS9{F0: 42, F1: int32(v20), F2: 915}
                                                                    _ = v57
                                                                case 128:
                                                                    v12 = v12
                                                                case 127:
                                                                    if false {
                                                                        if false {
                                                                            if false {
                                                                                var v58 uint = 0
                                                                                if (float64(float32((0.1 * (1e10 - 1e10)))) > -2.25) {
                                                                                    _ = len(v4)
                                                                                    var v59 uint = (v58 << ((v58) & 63))
                                                                                    _ = v59
                                                                                } else {
                                                                                    var v60 float32 = ((1.5 / (-2.25 + (float32((float64(0.1) + 1e10)) - 3.14159))) * 0.0)
                                                                                    var v61 uint16 = 5
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                                switch fi19 {
                                                                                case -1:
                                                                                    _ = len(v3)
                                                                                case 2147483647:
                                                                                    v62 := GenS7{F0: true}
                                                                                    _ = v62
                                                                                case 42:
                                                                                    switch (fi19 ^ (^(fi18))) {
                                                                                    case -2147483648:
                                                                                        if (v12 < v12) {
                                                                                            v63 := GenS10{F0: v20}
                                                                                            _ = v63
                                                                                        }
                                                                                    case 267:
                                                                                        v64 := GenS6{F0: 2147483647}
                                                                                        _ = v64
                                                                                    default:
                                                                                        v65 := GenS4{F0: true, F1: 243}
                                                                                        _ = v65
                                                                                    }
                                                                                }
                                                                                _ = v58
                                                                            } else {
                                                                                v12 -= 4294967295
                                                                                switch -9223372036854775808 {
                                                                                case -2147483648:
                                                                                    switch 4294967295 {
                                                                                    case -9223372036854775808:
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v20)
                                                                                        }
                                                                                    case -1:
                                                                                        if ((v12 >> ((65535) & 63)) == v12) {
                                                                                            fmt.Println(v12, v1, v12)
                                                                                        }
                                                                                    case 32767:
                                                                                        v12 += v12
                                                                                    }
                                                                                case 1:
                                                                                    if false {
                                                                                        if (true || (false && ((3.14159 * float32((3.14159 + ((-2.25 / 1e10) + -2.25)))) <= 1.5))) {
                                                                                            v66 := GenS11{F0: false, F1: "hello", F2: 2}
                                                                                            fmt.Println(v1, v20)
                                                                                            if true {
                                                                                                if (!((int64(v20) < -57) && ((true || (v20 != int16(v1))) || true))) {
                                                                                                    fmt.Println(fi18, fi18, fi19)
                                                                                                    v1 += v1
                                                                                                    if (18446744073709551615 != 290) {
                                                                                                        fmt.Println(v1)
                                                                                                    } else {
                                                                                                        v67 := GenS8{F0: 519, F1: 992, F2: fi19}
                                                                                                        switch fi19 {
                                                                                                        case 2147483647:
                                                                                                            var v68 int64 = v1
                                                                                                            _ = v68
                                                                                                        default:
                                                                                                            if (("\n\t" + "foo bar") >= "hello") {
                                                                                                                if (v1 > v1) {
                                                                                                                    if (true && (127 >= 0)) {
                                                                                                                        v1 -= v1
                                                                                                                        v69 := v2[(fi19 * fi19)]
                                                                                                                        var v70 int8 = 1
                                                                                                                        _ = v69
                                                                                                                        _ = v70
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            if (18446744073709551615 < 42) {
                                                                                                                                v71 := GenS12{F0: (^(fi19)), F1: 4294967295}
                                                                                                                                _ = v71
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v1 += v1
                                                                                                                                v12 = (^(v12))
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v20 += (v20 << ((42) & 63))
                                                                                                                            var v72 int64 = v1
                                                                                                                            var v73 string = "a"
                                                                                                                            var v74 int32 = 0
                                                                                                                            v74 += (-(v74))
                                                                                                                            _ = v72
                                                                                                                            _ = v73
                                                                                                                            _ = v74
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v20, v20)
                                                                                                                v20 = -56
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v67
                                                                                                    }
                                                                                                } else {
                                                                                                    v20 += (^(v20))
                                                                                                }
                                                                                            } else {
                                                                                                if true {
                                                                                                    fmt.Println(v20, v12, v20)
                                                                                                    if (false && false) {
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = "foo bar"
                                                                                                        }
                                                                                                    }
                                                                                                    _ = len(v4)
                                                                                                } else {
                                                                                                    v75 := GenS13{F0: v1}
                                                                                                    var v76 float64 = 3.14159
                                                                                                    _ = v75
                                                                                                    _ = v76
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    var v77 uint32 = 0
                                                                                                    _ = v77
                                                                                                case "foo bar":
                                                                                                    if (false && true) {
                                                                                                        fmt.Println(fi19, v1)
                                                                                                    }
                                                                                                    var v78 int32 = 1
                                                                                                    v78 = genFn0("hello", 65535)
                                                                                                    if false {
                                                                                                        if ("" != "a") {
                                                                                                            var v79 uint32 = 42
                                                                                                            if false {
                                                                                                                var v80 int8 = 127
                                                                                                                var v81 int8 = (v80 >> ((434) & 63))
                                                                                                                _ = v80
                                                                                                                _ = v81
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    v20 = v20
                                                                                                                    fmt.Println(v78, v78, fi18)
                                                                                                                }
                                                                                                                var v82 uint32 = v79
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = ((("a" + (("\n\t" + ("foo bar" + "\n\t")) + ("" + "foo bar"))) + "") + "foo bar")
                                                                                                            }
                                                                                                            _ = v79
                                                                                                        } else {
                                                                                                            var v83 float32 = (-2.25 * 0.1)
                                                                                                            fmt.Println(v1, v20)
                                                                                                            if (!((4294967295 > 65535) && ("hello" != ("" + "hello")))) {
                                                                                                                if ((fi18 - (fi19 | fi19)) >= (fi18 | (fi19 % (((fi18 + fi19)) | 1)))) {
                                                                                                                    v1 = v1
                                                                                                                    fmt.Println(v20, v1, v78)
                                                                                                                } else {
                                                                                                                    var v84 float32 = v83
                                                                                                                    _ = v84
                                                                                                                }
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = "hello"
                                                                                                                }
                                                                                                            } else {
                                                                                                                v1 = 0
                                                                                                                v1 = v1
                                                                                                            }
                                                                                                            _ = v83
                                                                                                        }
                                                                                                        fmt.Println(v78)
                                                                                                    } else {
                                                                                                        v12 -= v12
                                                                                                        var v85 int = fi19
                                                                                                        fmt.Println(v1, v78, v85)
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    _ = v78
                                                                                                case "a":
                                                                                                    var v86 int8 = 1
                                                                                                    _ = v86
                                                                                                }
                                                                                                var v87 uint32 = 2147483648
                                                                                                _ = v87
                                                                                            }
                                                                                            _ = v66
                                                                                        } else {
                                                                                            delete(v2, 42)
                                                                                            var v88 uint8 = 255
                                                                                            v88 = 2
                                                                                            _ = v88
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v89 uint64 = (v12 ^ 18446744073709551615)
                                                                                _ = v89
                                                                            }
                                                                            var v90 int64 = v1
                                                                            _ = v90
                                                                        }
                                                                        v1 += 535
                                                                    }
                                                                }
                                                            case 4294967295:
                                                                if (1 != 255) {
                                                                    v91 := GenS14{F0: 0, F1: 0}
                                                                    v12 = (^(v12))
                                                                    var v92 int64 = 1
                                                                    switch (("foo bar" + ("\n\t" + "hello")) + "hello") {
                                                                    case "hello":
                                                                        v93 := GenS7{F0: true}
                                                                        _ = v93
                                                                    }
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    v94 := GenS11{F0: true, F1: "\n\t", F2: 0}
                                                                    v95 := GenS15{F0: false, F1: 65535}
                                                                    var v96 float32 = (3.14159 * float32(0.1))
                                                                    _ = len(v4)
                                                                    if len(v4) > 0 {
                                                                        v4[0] = ("" + ((("" + "a") + (("foo bar" + ("hello" + (("a" + ("hello" + "foo bar")) + "\n\t"))) + ("" + "foo bar"))) + "a"))
                                                                    }
                                                                    _ = v94
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            var v97 int64 = (v1 ^ 0)
                                                            _ = v97
                                                        }
                                                    default:
                                                        if true {
                                                            var v98 uint = 1
                                                            var v99 int16 = 1000
                                                            var v100 uint16 = 32767
                                                            _ = v100
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    }
                                                    switch (("" + "foo bar") + "hello") {
                                                    case "\n\t":
                                                        var v101 float64 = float64((0.0 - (1e10 + -2.25)))
                                                        _ = v101
                                                    default:
                                                        var v102 int64 = (v1 & 9223372036854775807)
                                                        _ = v102
                                                    }
                                                } else {
                                                    fmt.Println(v20, v20)
                                                    if true {
                                                        if (false || true) {
                                                            var v103 uint32 = 367
                                                            _ = v103
                                                        }
                                                        v1 = (v1 | v1)
                                                    } else {
                                                        if false {
                                                            if true {
                                                                v12 = v12
                                                                fmt.Println(fi19, fi18, v20)
                                                            }
                                                        } else {
                                                            fmt.Println(v20)
                                                        }
                                                    }
                                                    v104 := v2[-1]
                                                    _ = v104
                                                }
                                            }
                                        }
                                        var v105 uint32 = 42
                                        _ = v105
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        v106 := GenS16{F0: -60, F1: -9223372036854775808, F2: true}
                                        _ = v106
                                    case "hello":
                                        if false {
                                            var v107 string = ("a" + ("foo bar" + ""))
                                            _ = v107
                                        } else {
                                            if (42 == 441) {
                                                v12 = 180
                                                var v108 uint32 = 305
                                                var v109 bool = false
                                                _ = v108
                                                _ = v109
                                            }
                                            v110 := GenS17{F0: 391, F1: ("" + ("hello" + "hello"))}
                                            var v111 int16 = (v20 % ((v20) | 1))
                                            v20 = v111
                                            switch fi18 {
                                            case 32767:
                                                v1 = -85
                                                v20 = (v111 & -260)
                                                v112 := GenS9{F0: 18446744073709551615, F1: 46, F2: 2}
                                                _ = v112
                                            default:
                                                v113 := GenS9{F0: 0, F1: genFn1(255, 0), F2: 0}
                                                _ = v113
                                            }
                                            _ = v110
                                            _ = v111
                                        }
                                    }
                                } else {
                                    if false {
                                        if (0.1 >= 0.0) {
                                            fmt.Println(v1)
                                            v114 := GenS18{F0: 699, F1: ("a" + ("hello" + "a"))}
                                            _ = v114
                                        }
                                        var v115 bool = (1 >= 4294967295)
                                        if v115 {
                                            if v115 {
                                                fmt.Println(v12, v1, v115)
                                            }
                                            fmt.Println(fi18)
                                        }
                                        var v116 int64 = (v1 + v1)
                                        _ = v115
                                        _ = v116
                                    }
                                }
                            }
                        case 42:
                            v12 = v12
                        default:
                            v117 := GenS19{F0: ("" + "a"), F1: -1}
                            _ = v117
                        }
                    }
                    v12 -= (v12 / ((1) | 1))
                    if (!(v20 < (v20 % (((v20 & 42)) | 1)))) {
                        v20 += 44
                    } else {
                        v20 -= v20
                        switch "\n\t" {
                        case "\n\t":
                            switch (fi18 + fi19) {
                            case 4294967295:
                                if (1 <= uint16(v12)) {
                                    v20 = 0
                                }
                            }
                        default:
                            if true {
                                v1 = (v1 - (-(v1)))
                                if ((!true) || (((32767 >= (fi18 ^ (^(fi19)))) || true) || (v1 > v1))) {
                                    v118 := GenS18{F0: 165, F1: "foo bar"}
                                    if false {
                                        var v119 int32 = -1
                                        _ = v119
                                    }
                                    var v120 string = ("a" + "hello")
                                    _ = v118
                                    _ = v120
                                } else {
                                    var v121 uint16 = 824
                                    _ = v121
                                }
                                switch fi18 {
                                case 128:
                                    _ = len(v4)
                                default:
                                    v122 := GenS20{F0: 128, F1: 42, F2: 42}
                                    _ = v122
                                }
                                v123 := v2[(fi19 / ((950) | 1))]
                                _ = v123
                            } else {
                                fmt.Println(fi18)
                            }
                        }
                        if ((1 <= 161) || (("\n\t" + "hello") >= ("foo bar" + ("a" + "a")))) {
                            if (-1 > 335) {
                                fmt.Println(v12, fi19, fi18)
                                v124 := GenS18{F0: 120, F1: "a"}
                                _ = v124
                            }
                            v20 = (v20 << ((217) & 63))
                            var v125 bool = (!true)
                            _ = v125
                        }
                    }
                    _ = len(v4)
                    switch "foo bar" {
                    case "\n\t":
                        var v126 int8 = 97
                        _ = v126
                    }
                    var v127 float32 = (-2.25 - 1e10)
                    _ = fi19
                    _ = v127
                    _ = v20
                    _ = v24
                }
                switch ((("hello" + ("hello" + ("foo bar" + "\n\t"))) + ("hello" + ("\n\t" + "a"))) + "\n\t") {
                case "\n\t":
                    v128 := GenS4{F0: false, F1: 42}
                    _ = v128
                case "a":
                    var v129 uint = 134
                    _ = v129
                default:
                    var v130 float64 = 0.1
                    _ = v130
                }
                fmt.Println(v12, v12, fi18)
                _ = fi18
            }
            v131 := GenS21{F0: -731, F1: 9223372036854775807, F2: 2}
            _ = v131
        }
        var v132 float64 = 0.0
        _ = v132
    } else {
        v12 = v12
        for fi133 := 0; fi133 < 17; fi133++ {
            v134 := GenS22{F0: "hello"}
            fmt.Println(v1, v12, v12)
            _ = fi133
            _ = v134
        }
    }
    _ = len(v4)
    for fi135 := 0; fi135 < 12; fi135++ {
        for fi136 := 0; fi136 < 5; fi136++ {
            switch fi136 {
            case 9223372036854775807:
                if (int16(v1) <= -32768) {
                    var v137 int = fi135
                    _ = v137
                }
            default:
                if true {
                    switch (fi135 | (fi135 / (((fi136 ^ 0)) | 1))) {
                    case 128:
                        var v138 uint16 = 32767
                        _ = v138
                    case -2147483648:
                        if (true || false) {
                            fmt.Println(v1)
                            if false {
                                if true {
                                    if false {
                                        if true {
                                            v1 = 9223372036854775807
                                        } else {
                                            var v139 int = (fi136 << ((42) & 63))
                                            var v140 int16 = -32768
                                            switch (v139 & fi136) {
                                            case -495:
                                                var v141 string = "hello"
                                                _ = v141
                                            case -2147483648:
                                                v142 := GenS23{F0: 856, F1: 1}
                                                _ = v142
                                            case 0:
                                                if (!(!(!(2 != 2)))) {
                                                    v143 := GenS20{F0: 1, F1: 18446744073709551615, F2: 16}
                                                    _ = v143
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v144 := GenS11{F0: (!(-1 >= int8(fi135))), F1: "", F2: v12}
                                                    var v145 float64 = 1.5
                                                    _ = v144
                                                    _ = v145
                                                }
                                                var v146 uint16 = 2
                                                if false {
                                                    if ((-1 != v140) || true) {
                                                        fmt.Println(v1)
                                                        _ = len(v4)
                                                    }
                                                } else {
                                                    var v147 int16 = 1
                                                    var v148 int8 = -1
                                                    if (false && false) {
                                                        if true {
                                                            v148 += (v148 ^ v148)
                                                            v149 := GenS6{F0: genFn1(0, 18446744073709551615)}
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                switch -2147483648 {
                                                                case 127:
                                                                    var v150 uint32 = 1
                                                                    _ = v150
                                                                default:
                                                                    fmt.Println(v147, fi136)
                                                                }
                                                            case "a":
                                                                var v151 float32 = 3.14159
                                                                _ = v151
                                                            }
                                                            v152 := v3[(fi135 / ((32767) | 1))]
                                                            _ = v149
                                                            _ = v152
                                                        }
                                                    }
                                                    _ = v147
                                                    _ = v148
                                                }
                                                _ = v146
                                            }
                                            _ = v139
                                            _ = v140
                                        }
                                        var v153 float64 = 1.5
                                        _ = v153
                                    }
                                    switch fi136 {
                                    case 9223372036854775807:
                                        v154 := GenS9{F0: 2, F1: 807, F2: 0}
                                        _ = v154
                                    }
                                    if false {
                                        if (!true) {
                                            v4 = append(v4, "\n\t")
                                            switch -2147483648 {
                                            case 9223372036854775807:
                                                if ((0.0 * 3.14159) <= 0.1) {
                                                    v155 := GenS24{F0: 42}
                                                    v1 = (v1 << ((18446744073709551615) & 63))
                                                    _ = v155
                                                }
                                            case 42:
                                                v12 -= 2
                                            case 813:
                                                switch (("a" + "\n\t") + (("hello" + ("" + "\n\t")) + ("hello" + (("foo bar" + "hello") + "\n\t")))) {
                                                case "hello":
                                                    fmt.Println(v1, fi136, v1)
                                                case "a":
                                                    var v156 uint32 = uint32(v12)
                                                    v12 -= 18446744073709551615
                                                    _ = v156
                                                case "":
                                                    _ = len(v4)
                                                }
                                            }
                                            delete(v3, 1)
                                            v157 := GenS2{F0: 0, F1: 42}
                                            if false {
                                                if (true || (!true)) {
                                                    v1 = 9223372036854775807
                                                }
                                            } else {
                                                v12 -= (v12 / ((18446744073709551615) | 1))
                                                var v158 int8 = 1
                                                _ = v158
                                            }
                                            _ = v157
                                        }
                                        var v159 int64 = (-(v1))
                                        switch "a" {
                                        case "foo bar":
                                            v159 = v159
                                        case "hello":
                                            v4 = append(v4, ("hello" + "a"))
                                        case "\n\t":
                                            v160 := v3[(fi136 & 835)]
                                            _ = v160
                                        }
                                        _ = v159
                                    }
                                } else {
                                    switch fi136 {
                                    case -9223372036854775808:
                                        fmt.Println(v1)
                                        var v161 uint = 0
                                        if (!(float32(((1e10 - 3.14159) + ((0.1 / ((3.14159 * 1.5) + 1.5)) / -2.25))) != 1e10)) {
                                            v161 += 4294967295
                                            if true {
                                                v162 := GenS25{F0: 0, F1: (-(v1))}
                                                switch fi136 {
                                                case 1:
                                                    if (float64(0.1) < 1.5) {
                                                        switch (fi136 << (((v161 & 861)) & 63)) {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi135)
                                                        default:
                                                            v163 := GenS26{F0: (v1 % ((v1) | 1)), F1: v161}
                                                            _ = v163
                                                        }
                                                    } else {
                                                        if false {
                                                            var v164 int8 = 127
                                                            var v165 uint32 = 4294967295
                                                            _ = v164
                                                            _ = v165
                                                        }
                                                    }
                                                case 32768:
                                                    v166 := v2[-902]
                                                    var v167 float32 = ((float32(-2.25) * (float32((1e10 + 1.5)) / 1e10)) + (0.0 * ((1.5 - 1e10) * 1e10)))
                                                    var v168 string = "a"
                                                    _ = v166
                                                    _ = v167
                                                    _ = v168
                                                case -109:
                                                    v1 = v1
                                                }
                                                var v169 float32 = -2.25
                                                _ = v162
                                                _ = v169
                                            }
                                        } else {
                                            var v170 float64 = float64(-2.25)
                                            _ = v170
                                        }
                                        _ = v161
                                    default:
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(v1)
                                        case "a":
                                            v12 = (^(v12))
                                            if ((!(fi136 < (fi136 ^ fi136))) && (!false)) {
                                                var v171 uint16 = 255
                                                _ = v171
                                            }
                                        case "\n\t":
                                            var v172 uint8 = 33
                                            _ = v172
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v1, fi136, v12)
                                v12 = v12
                                switch (("" + ("a" + "a")) + ("\n\t" + "a")) {
                                case "hello":
                                    v12 = 467
                                default:
                                    if ((1 >= -1) && (0.1 == 1.5)) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        switch -1 {
                                        case 1:
                                            v12 += v12
                                        default:
                                            var v173 uint = 969
                                            _ = v173
                                        }
                                    } else {
                                        v174 := GenS27{F0: (v1 >> ((112) & 63)), F1: 72}
                                        _ = v174
                                    }
                                }
                            }
                            v1 -= -881
                        }
                    case 32767:
                        var v175 uint64 = v12
                        _ = v175
                    default:
                        var v176 int8 = -85
                        _ = v176
                    }
                    v12 = (v12 / ((744) | 1))
                    switch (("a" + "") + "a") {
                    case "":
                        v1 += -770
                    default:
                        if false {
                            v1 = 422
                            if false {
                                if true {
                                    var v177 uint16 = 65535
                                    delete(v2, fi136)
                                    _ = v177
                                } else {
                                    switch "a" {
                                    case "":
                                        v178 := GenS11{F0: (19 < 2147483647), F1: ((("foo bar" + "a") + "hello") + "\n\t"), F2: v12}
                                        _ = v178
                                    default:
                                        var v179 uint32 = 199
                                        _ = v179
                                    }
                                    v4 = append(v4, ("hello" + "foo bar"))
                                    _ = len(v4)
                                    switch (fi136 % ((fi135) | 1)) {
                                    case 0:
                                        var v180 float64 = (1.5 - 1e10)
                                        v181 := GenS28{F0: true, F1: 1000}
                                        _ = v180
                                        _ = v181
                                    default:
                                        v1 -= -9223372036854775808
                                    }
                                    v12 -= 1
                                }
                                fmt.Println(fi136, fi136)
                            } else {
                                fmt.Println(fi136)
                            }
                            var v182 uint16 = 0
                            _ = v182
                        } else {
                            var v183 int32 = 212
                            _ = len(v4)
                            _ = v183
                        }
                    }
                    v184 := GenS29{F0: 1}
                    var v185 int = (fi136 ^ fi135)
                    if (((fi135 ^ (fi135 & 608)) < (v185 ^ 32768)) || false) {
                        v1 += (v1 << ((42) & 63))
                        var v186 uint64 = (^(v12))
                        _ = v186
                    }
                    _ = v184
                    _ = v185
                }
            }
            _ = fi136
        }
        delete(v2, int(v1))
        switch "foo bar" {
        case "a":
            switch "hello" {
            case "":
                if (0 < 1) {
                    var v187 string = (((("" + "") + ("" + "foo bar")) + "") + (((("" + "\n\t") + ((((("a" + "hello") + "hello") + "foo bar") + ("\n\t" + ("\n\t" + ("\n\t" + "")))) + ("a" + (((("" + "foo bar") + "hello") + "hello") + "\n\t")))) + "a") + "a"))
                    if false {
                        v12 += 2
                        var v188 uint16 = 32767
                        switch ((v187 + v187) + ("\n\t" + v187)) {
                        case "foo bar":
                            switch fi135 {
                            case -1:
                                for fi189 := 0; fi189 < 3; fi189++ {
                                    v188 = 32767
                                    _ = fi189
                                }
                            case 9223372036854775807:
                                if true {
                                    var v190 uint32 = 0
                                    fmt.Println(v188)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v190
                                } else {
                                    for fi191 := 0; fi191 < 16; fi191++ {
                                        v192 := GenS21{F0: -1, F1: fi135, F2: 724}
                                        switch (v187 + "foo bar") {
                                        case "foo bar":
                                            if true {
                                                switch (fi191 << ((65535) & 63)) {
                                                case 39:
                                                    fmt.Println(fi135)
                                                }
                                            }
                                        case "a":
                                            v12 = 0
                                            var v193 float64 = 0.0
                                            _ = v193
                                        case "":
                                            fmt.Println(v187, fi135, v1)
                                        }
                                        v12 -= 4294967295
                                        _ = fi191
                                        _ = v192
                                    }
                                    v4 = append(v4, v187)
                                }
                            }
                        case "":
                            var v194 uint8 = 140
                            _ = v194
                        case "\n\t":
                            v2[fi135] = true
                        }
                        _ = v188
                    }
                    _ = v187
                } else {
                    for fi195 := 0; fi195 < 5; fi195++ {
                        var v196 int32 = 2147483647
                        _ = fi195
                        _ = v196
                    }
                }
            case "hello":
                var v197 bool = true
                _ = v197
            }
        case "foo bar":
            switch ("" + "a") {
            case "a":
                var v198 int = (fi135 % ((fi135) | 1))
                _ = v198
            default:
                for fi199 := 0; fi199 < 2; fi199++ {
                    fmt.Println(fi135, fi199)
                    fmt.Println(v1)
                    v1 = (v1 * (v1 / ((v1) | 1)))
                    _ = fi199
                }
            }
        case "hello":
            for fi200 := 0; fi200 < 9; fi200++ {
                v12 = v12
                _ = fi200
            }
        default:
            v12 = (^(v12))
        }
        _ = fi135
    }
    if false {
        v12 += (v12 + (v12 ^ v12))
    } else {
        v201 := GenS11{F0: true, F1: "", F2: 42}
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v1)
        }
        _ = v201
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v202 := GenS9{F0: 42, F1: 611, F2: 681}
    if false {
        if false {
            for fi203 := 0; fi203 < 10; fi203++ {
                fmt.Println(v1)
                _ = fi203
            }
        }
        v12 = 18446744073709551615
    }
    _ = len(v2)
    fmt.Println(v1, v12)
    if (false || false) {
        switch 0 {
        case -414:
            v1 = -492
        default:
            v12 = v12
        }
        v3[2147483647] = true
    }
    fmt.Println(v12, v12)
    var v204 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v204:", v204)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v202
    _ = v204
    _ = v3
    _ = v4
}
