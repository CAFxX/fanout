package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int; F1 int8; F2 int }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 int16; F2 uint8 }
type GenS15 struct { F0 int32; F1 int64; F2 uint32 }
type GenS16 struct { F0 int64; F1 bool; F2 uint32 }
type GenS17 struct { F0 int32; F1 string }
type GenS18 struct { F0 uint8; F1 string }
type GenS19 struct { F0 int32; F1 int; F2 int16 }
type GenS20 struct { F0 int64; F1 int }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 uint64; F1 int64; F2 int32 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 int8; F1 int }
func genFn0(p0 int, p1 int8) int32 {
    var v4 int8 = (^(p1))
    if ((84 == 255) && true) {
        v5 := []int{(p0 >> ((42) & 63)), p0, p0, p0}
        _ = v5
    }
    v6 := []int{p0, p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 2147483647
}
func genFn1(p0 uint64) int32 {
    v7 := []int{}
    p0 -= 734
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    _ = p0
    _ = v7
    return 897
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v8 float32 = (0.1 + 1e10)
    _ = len(v1)
    delete(v2, -1)
    var v9 string = (((("foo bar" + ("" + ("a" + ""))) + "foo bar") + "hello") + "a")
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := []int{-1}
        var v12 uint32 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    v8 = 0.1
    v9 = (v9 + v9)
    if true {
        for fi13 := 0; fi13 < 4; fi13++ {
            if (-986 > 0) {
                v14 := GenS1{F0: 42, F1: -9223372036854775808, F2: 247}
                if (946 > 2) {
                    v15 := GenS1{F0: 58, F1: 488, F2: 127}
                    v9 = ("hello" + ("\n\t" + (v9 + ((("" + "a") + v9) + "foo bar"))))
                    _ = v15
                } else {
                    v16 := GenS2{F0: 1, F1: 42, F2: 42}
                    fmt.Println(v8, v9, v9)
                    var v17 uint32 = 591
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
            _ = fi13
        }
        fmt.Println(v8)
        fmt.Println(v9)
    }
    if (((float64(((3.14159 * 0.1) * (3.14159 + 3.14159))) / float64(v8)) - (3.14159 + -2.25)) <= ((1.5 / (float64(1e10) + 1e10)) + (1.5 * float64(-2.25)))) {
        if ((!(2 > 1)) && true) {
            v18 := GenS3{F0: 247, F1: 0, F2: ((false || (!(65535 >= 2))) || (0 < 599))}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := GenS3{F0: 1, F1: 18, F2: false}
                var v21 uint64 = 1
                switch ((v9 + ("hello" + ((("" + "hello") + "a") + ("hello" + "")))) + v9) {
                case "foo bar":
                    v1[v9] = (!(!(4294967295 == 469)))
                default:
                    v21 = (^(v21))
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            if false {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v9 = v9
                    switch v9 {
                    case "hello":
                        switch (fi22 % ((fi22) | 1)) {
                        case -601:
                            var v23 bool = (1 > -22)
                            _ = v23
                        case 175:
                            var v24 uint32 = 2147483648
                            _ = v24
                        default:
                            if (42 <= 2) {
                                var v25 bool = (!true)
                                _ = v25
                            } else {
                                if true {
                                    var v26 int64 = 59
                                    v8 = ((v8 / -2.25) * float32(0.1))
                                    _ = v26
                                }
                            }
                        }
                    }
                    _ = fi22
                }
            } else {
                v27 := GenS4{F0: 0, F1: 65}
                v28 := GenS5{F0: 0}
                if (42 != 32767) {
                    _ = len(v1)
                    if true {
                        var v29 int8 = 127
                        if (false || (0 == 65535)) {
                            v30 := []int{}
                            v29 += (v29 / (((^(v29))) | 1))
                            _ = v30
                        } else {
                            var v31 int16 = -237
                            v29 -= -1
                            _ = v31
                        }
                        _ = v29
                    }
                    v32 := GenS4{F0: 1, F1: 1}
                    _ = v32
                } else {
                    if ((0 <= 9223372036854775807) || true) {
                        v9 = v9
                    } else {
                        v33 := GenS6{F0: -291, F1: 0}
                        v9 = v9
                        _ = v33
                    }
                    switch 42 {
                    case 32768:
                        var v34 int8 = 42
                        _ = v34
                    case -2147483648:
                        _ = len(v2)
                    default:
                        if (!(4294967295 < 425)) {
                            if true {
                                if (-1 > 0) {
                                    if true {
                                        var v35 uint16 = 0
                                        v9 = v9
                                        fmt.Println(v8)
                                        _ = v35
                                    }
                                } else {
                                    v36 := GenS7{F0: 50}
                                    v9 = v9
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 3; fi37++ {
                                    var v38 string = (v9 + ((v9 + (v9 + (v9 + ""))) + v9))
                                    v9 = "a"
                                    for fi39 := 0; fi39 < 13; fi39++ {
                                        if true {
                                            if (-108 == 127) {
                                                if (32767 >= 973) {
                                                    v40 := GenS8{F0: -1, F1: 2}
                                                    fmt.Println(fi37)
                                                    _ = v40
                                                } else {
                                                    var v41 uint = 18446744073709551615
                                                    switch v38 {
                                                    case "":
                                                        var v42 string = v9
                                                        _ = v42
                                                    default:
                                                        var v43 uint64 = 1
                                                        _ = v43
                                                    }
                                                    _ = v41
                                                }
                                                v9 = v38
                                            } else {
                                                var v44 bool = true
                                                v9 = v38
                                                _ = v44
                                            }
                                            if (0.0 == 0.0) {
                                                switch v38 {
                                                case "":
                                                    v8 = v8
                                                    var v45 int32 = 2147483647
                                                    _ = v45
                                                case "a":
                                                    if false {
                                                        if true {
                                                            v3[2] = 0
                                                        } else {
                                                            fmt.Println(fi39, v8)
                                                            fmt.Println(v38)
                                                            v8 += (float32(1.5) - (v8 / 1.0))
                                                            v8 = 0.0
                                                            var v46 uint16 = 1
                                                            _ = v46
                                                        }
                                                        _ = len(v1)
                                                        v2[fi39] = "hello"
                                                    } else {
                                                        var v47 uint64 = 2
                                                        v47 += 4294967295
                                                        var v48 uint = 18446744073709551615
                                                        if true {
                                                            if true {
                                                                if (109 <= 0) {
                                                                    v48 -= 515
                                                                    if true {
                                                                        var v49 uint16 = 255
                                                                        var v50 uint8 = 128
                                                                        fmt.Println(v47, v49, fi39)
                                                                        delete(v3, 1)
                                                                        _ = v49
                                                                        _ = v50
                                                                    } else {
                                                                        if (!true) {
                                                                            v51 := GenS7{F0: 1}
                                                                            v9 = "\n\t"
                                                                            var v52 float64 = 0.0
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            fmt.Println(v8, fi39, fi37)
                                                                        }
                                                                        v48 -= (v48 / ((42) | 1))
                                                                        var v53 float64 = (0.0 + 0.1)
                                                                        _ = v53
                                                                    }
                                                                    switch (fi39 / ((fi39) | 1)) {
                                                                    case 32768:
                                                                        fmt.Println(v9, v38, v9)
                                                                    case -1:
                                                                        v54 := []int{583}
                                                                        _ = v54
                                                                    }
                                                                }
                                                            } else {
                                                                v55 := GenS2{F0: 167, F1: v47, F2: genFn0(0, -48)}
                                                                if false {
                                                                    v56 := v2[(fi39 << ((uint(v47)) & 63))]
                                                                    fmt.Println(v48, v8, v47)
                                                                    _ = v56
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v38 = v38
                                                                    default:
                                                                        if false {
                                                                            v47 += v47
                                                                            _ = len(v3)
                                                                        } else {
                                                                            fmt.Println(fi37, fi39, fi39)
                                                                        }
                                                                    }
                                                                    v57 := GenS6{F0: (fi37 / (((fi37 % (((fi37 + fi37)) | 1))) | 1)), F1: -152}
                                                                    if false {
                                                                        v48 += v48
                                                                        v58 := GenS4{F0: 2, F1: 127}
                                                                        _ = len(v1)
                                                                        if (!(65535 != 2147483647)) {
                                                                            var v59 int32 = 0
                                                                            v8 -= 3.14159
                                                                            _ = v59
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v8, fi39, fi39)
                                                                            } else {
                                                                                v48 = v48
                                                                                v48 -= 0
                                                                                if (!((!((!(0.1 >= ((float64(v8) - 0.1) + 1.5))) || (2147483647 != 1))) && true)) {
                                                                                    var v60 float32 = v8
                                                                                    _ = v60
                                                                                } else {
                                                                                    switch (fi39 & (fi37 % ((fi39) | 1))) {
                                                                                    case 199:
                                                                                        switch -511 {
                                                                                        case 0:
                                                                                            if (0 >= 2) {
                                                                                                fmt.Println(v9, v8)
                                                                                                var v61 uint64 = (v47 / ((v47) | 1))
                                                                                                if false {
                                                                                                    var v62 float64 = (-2.25 + (0.0 * -2.25))
                                                                                                    v47 += 1
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    v47 += (v61 + 463)
                                                                                                    v9 = v38
                                                                                                }
                                                                                                v1[v9] = false
                                                                                                var v63 bool = true
                                                                                                _ = v61
                                                                                                _ = v63
                                                                                            } else {
                                                                                                v38 = ((v38 + "") + v38)
                                                                                                fmt.Println(v8)
                                                                                            }
                                                                                        }
                                                                                    case -2147483648:
                                                                                        v9 = (("a" + (v9 + "")) + v38)
                                                                                    case 2147483647:
                                                                                        if (true && false) {
                                                                                            v9 = (v9 + "hello")
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi39, v8, fi37)
                                                                                    v2[(-(fi39))] = ("a" + ("" + (v9 + "hello")))
                                                                                }
                                                                            }
                                                                            var v64 int = fi37
                                                                            _ = v64
                                                                        }
                                                                        _ = v58
                                                                    } else {
                                                                        if false {
                                                                            v38 = v9
                                                                        }
                                                                        if false {
                                                                            v65 := v1[v38]
                                                                            var v66 float64 = 0.1
                                                                            switch v9 {
                                                                            case "":
                                                                                v67 := v3[1]
                                                                                _ = v67
                                                                            case "a":
                                                                                v68 := []int{}
                                                                                _ = v68
                                                                            default:
                                                                                var v69 float64 = v66
                                                                                _ = v69
                                                                            }
                                                                            switch v9 {
                                                                            case "":
                                                                                v38 = ((v38 + ("hello" + v9)) + "")
                                                                                switch fi39 {
                                                                                case 2147483647:
                                                                                    var v70 float64 = v66
                                                                                    _ = v70
                                                                                case -2147483648:
                                                                                    fmt.Println(v48, v66)
                                                                                case 42:
                                                                                    var v71 uint16 = 65535
                                                                                    _ = v71
                                                                                default:
                                                                                    v65 = (471 == -9223372036854775808)
                                                                                }
                                                                            case "a":
                                                                                switch v9 {
                                                                                case "":
                                                                                    v9 = "foo bar"
                                                                                case "hello":
                                                                                    v47 -= v47
                                                                                }
                                                                                if (v65 && ((((v66 / float64((-2.25 * (((((((0.1 / v8) / v8) - v8) + (-2.25 - v8)) / ((v8 / v8) / float32(v66))) + 1.5) * (((v8 + 0.0) * 1e10) * v8))))) + 1e10) + 1.5) >= (v66 - v66))) {
                                                                                    if false {
                                                                                        var v72 int = 9223372036854775807
                                                                                        _ = v72
                                                                                    }
                                                                                    v73 := GenS6{F0: fi37, F1: 1000}
                                                                                    _ = v73
                                                                                } else {
                                                                                    v38 = v38
                                                                                }
                                                                                v66 = v66
                                                                            default:
                                                                                v48 += 65535
                                                                            }
                                                                            _ = v65
                                                                            _ = v66
                                                                        } else {
                                                                            switch ((("" + v9) + v9) + "a") {
                                                                            case "a":
                                                                                v1["\n\t"] = false
                                                                            default:
                                                                                v38 = v38
                                                                            }
                                                                        }
                                                                        v48 = uint(v47)
                                                                        delete(v1, ((v38 + "foo bar") + (v9 + "foo bar")))
                                                                    }
                                                                    _ = v57
                                                                }
                                                                _ = v55
                                                            }
                                                            v74 := GenS9{F0: 484}
                                                            _ = v74
                                                        } else {
                                                            v38 = (v38 + "a")
                                                            var v75 int64 = 1
                                                            v76 := GenS10{F0: v47}
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        switch (v38 + v9) {
                                                        case "a":
                                                            if true {
                                                                v8 = v8
                                                                v77 := GenS11{F0: fi37, F1: 68, F2: fi39}
                                                                _ = v77
                                                            }
                                                            if (1 == 1000) {
                                                                switch (fi39 % (((fi39 | (fi37 | 42))) | 1)) {
                                                                case 32768:
                                                                    var v78 bool = ((^(v47)) != v47)
                                                                    _ = v78
                                                                case 127:
                                                                    v79 := GenS6{F0: fi37, F1: 32767}
                                                                    _ = v79
                                                                case -141:
                                                                    v80 := GenS12{F0: 1, F1: (("\n\t" + ((v38 + v38) + ("a" + "\n\t"))) + v38)}
                                                                    _ = v80
                                                                default:
                                                                    v1[v38] = true
                                                                }
                                                                v2[(fi37 >> ((v48) & 63))] = "a"
                                                            }
                                                            var v81 int16 = 0
                                                            _ = v81
                                                        case "foo bar":
                                                            if (true && (291 != -1)) {
                                                                v82 := v1[("" + "\n\t")]
                                                                _ = v82
                                                            }
                                                            var v83 float64 = float64(v8)
                                                            v38 = "\n\t"
                                                            _ = v83
                                                        case "":
                                                            v8 += v8
                                                        default:
                                                            v47 = 296
                                                        }
                                                        v9 = ("hello" + v9)
                                                        var v84 string = v9
                                                        v85 := []int{}
                                                        if (2 < 662) {
                                                            fmt.Println(fi39)
                                                        }
                                                        if len(v85) >= 2 {
                                                            v85 = v85[1:len(v85)]
                                                        }
                                                        _ = v47
                                                        _ = v48
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                }
                                                if false {
                                                    v2[(fi37 | fi39)] = ((v38 + ((v9 + "") + v38)) + v9)
                                                    v2[846] = (v9 + (v9 + (((v38 + "\n\t") + (("\n\t" + v38) + "hello")) + ("hello" + v38))))
                                                    switch (fi37 - (fi39 / ((76) | 1))) {
                                                    case 128:
                                                        v8 = v8
                                                    }
                                                }
                                            } else {
                                                fmt.Println(fi39, v8, v38)
                                                delete(v3, 929)
                                            }
                                        } else {
                                            v8 += v8
                                            if true {
                                                v86 := []int{}
                                                switch fi37 {
                                                case 128:
                                                    fmt.Println(fi37, v8)
                                                    var v87 bool = (2147483648 < 545)
                                                    var v88 float64 = 1.5
                                                    _ = v87
                                                    _ = v88
                                                case 9223372036854775807:
                                                    if (false && (0 <= -1)) {
                                                        if (false && false) {
                                                            fmt.Println(v8, v9)
                                                            v8 -= 0.1
                                                            v89 := GenS13{F0: 42}
                                                            _ = v89
                                                        }
                                                        v90 := v2[(fi39 & (fi37 << ((2) & 63)))]
                                                        _ = v90
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                case 0:
                                                    var v91 int = (fi39 >> ((554) & 63))
                                                    _ = v91
                                                }
                                                if false {
                                                    var v92 uint16 = 255
                                                    var v93 uint = 315
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    v94 := GenS14{F0: 2147483647, F1: -837, F2: 0}
                                                    var v95 float32 = (v8 / 1.0)
                                                    _ = v94
                                                    _ = v95
                                                }
                                                _ = v86
                                            }
                                        }
                                        switch (fi37 / ((32767) | 1)) {
                                        case -772:
                                            v96 := GenS15{F0: -1, F1: -9223372036854775808, F2: 457}
                                            _ = v96
                                        case 1:
                                            _ = len(v2)
                                        case 127:
                                            _ = len(v3)
                                        }
                                        _ = fi39
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                for fi97 := 0; fi97 < 5; fi97++ {
                                    v98 := GenS16{F0: -964, F1: true, F2: 2147483648}
                                    _ = fi97
                                    _ = v98
                                }
                            } else {
                                if (1 > 1000) {
                                    var v99 float32 = v8
                                    _ = v99
                                } else {
                                    if (v8 > 0.1) {
                                        var v100 uint16 = 32767
                                        if (false || false) {
                                            var v101 uint32 = 2147483648
                                            _ = v101
                                        }
                                        if false {
                                            v102 := GenS17{F0: genFn1(0), F1: (v9 + "hello")}
                                            fmt.Println(v8)
                                            _ = v102
                                        }
                                        _ = v100
                                    }
                                }
                            }
                            if false {
                                var v103 bool = true
                                _ = v103
                            } else {
                                if (505 >= 2) {
                                    if true {
                                        if false {
                                            for fi104 := 0; fi104 < 4; fi104++ {
                                                if false {
                                                    switch fi104 {
                                                    case -9223372036854775808:
                                                        if ((798 <= 42) && false) {
                                                            _ = len(v1)
                                                        } else {
                                                            var v105 bool = ((13 >= -905) || ((v8 - -2.25) <= float32((0.1 * (0.1 - (0.1 / 0.1))))))
                                                            var v106 int8 = 127
                                                            for fi107 := 0; fi107 < 17; fi107++ {
                                                                v106 = 1
                                                                if (142 < -2147483648) {
                                                                    if v105 {
                                                                        fmt.Println(v9)
                                                                        switch fi104 {
                                                                        case 2147483647:
                                                                            if (v9 < "") {
                                                                                _ = len(v1)
                                                                                delete(v3, 1)
                                                                            }
                                                                        default:
                                                                            v1["a"] = v105
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch "" {
                                                                    case "a":
                                                                        v108 := []int{fi104, (fi107 - (fi104 ^ 32768)), 42}
                                                                        _ = v108
                                                                    case "hello":
                                                                        var v109 int32 = 42
                                                                        _ = v109
                                                                    case "foo bar":
                                                                        var v110 int64 = 9223372036854775807
                                                                        _ = v110
                                                                    }
                                                                }
                                                                _ = fi107
                                                            }
                                                            _ = v105
                                                            _ = v106
                                                        }
                                                    case 9223372036854775807:
                                                        v8 -= 3.14159
                                                    case -1:
                                                        var v111 float64 = 3.14159
                                                        _ = v111
                                                    }
                                                    if false {
                                                        var v112 int64 = 9223372036854775807
                                                        var v113 uint16 = 2
                                                        _ = v112
                                                        _ = v113
                                                    } else {
                                                        switch ((v9 + "a") + "hello") {
                                                        case "a":
                                                            v114 := GenS18{F0: 127, F1: ""}
                                                            _ = v114
                                                        case "foo bar":
                                                            v9 = ("\n\t" + ((("" + v9) + (v9 + ("\n\t" + "\n\t"))) + v9))
                                                        case "hello":
                                                            fmt.Println(v9, v8)
                                                        default:
                                                            var v115 uint16 = 65535
                                                            _ = v115
                                                        }
                                                    }
                                                }
                                                var v116 int = fi104
                                                for fi117 := 0; fi117 < 7; fi117++ {
                                                    v9 = "\n\t"
                                                    var v118 int8 = 18
                                                    _ = fi117
                                                    _ = v118
                                                }
                                                _ = fi104
                                                _ = v116
                                            }
                                        } else {
                                            fmt.Println(v8, v9)
                                            if (2 < 4294967295) {
                                                v8 += v8
                                                v8 += v8
                                                for fi119 := 0; fi119 < 19; fi119++ {
                                                    if (0.0 == v8) {
                                                        v120 := v3[2]
                                                        _ = v120
                                                    } else {
                                                        for fi121 := 0; fi121 < 11; fi121++ {
                                                            if false {
                                                                fmt.Println(fi121, fi121, v8)
                                                                v122 := []int{fi119}
                                                                var v123 float64 = float64(0.0)
                                                                _ = v122
                                                                _ = v123
                                                            } else {
                                                                v9 = ""
                                                                var v124 float32 = (v8 - float32((3.14159 * 1.5)))
                                                                var v125 int16 = -32768
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                            v8 -= v8
                                                            _ = fi121
                                                        }
                                                    }
                                                    _ = fi119
                                                }
                                            } else {
                                                fmt.Println(v8)
                                            }
                                            v8 += (v8 * v8)
                                        }
                                        var v126 int32 = 2147483647
                                        v127 := GenS3{F0: 0, F1: -1, F2: (2 >= 290)}
                                        _ = v126
                                        _ = v127
                                    }
                                }
                                var v128 string = (((("hello" + ((((v9 + "foo bar") + "") + "\n\t") + ((v9 + ((v9 + (v9 + (v9 + v9))) + v9)) + "hello"))) + "foo bar") + ((v9 + ("hello" + ((v9 + "hello") + ("a" + "\n\t")))) + "hello")) + v9)
                                _ = v128
                            }
                        }
                    }
                    for fi129 := 0; fi129 < 4; fi129++ {
                        if (((1.5 / (float64(3.14159) * (1e10 - 0.0))) <= (float64(-2.25) - 0.1)) && (fi129 != (fi129 & (-(fi129))))) {
                            for fi130 := 0; fi130 < 15; fi130++ {
                                v8 -= v8
                                var v131 uint16 = 255
                                var v132 string = ""
                                _ = fi130
                                _ = v131
                                _ = v132
                            }
                            for fi133 := 0; fi133 < 12; fi133++ {
                                var v134 uint8 = 128
                                if true {
                                    _ = len(v1)
                                    fmt.Println(v8)
                                    var v135 uint = 42
                                    _ = v135
                                }
                                if true {
                                    v9 = (v9 + v9)
                                    var v136 float32 = 3.14159
                                    v137 := GenS19{F0: 2147483647, F1: fi129, F2: int16(v134)}
                                    _ = v136
                                    _ = v137
                                }
                                _ = fi133
                                _ = v134
                            }
                            fmt.Println(v8, v9)
                        } else {
                            var v138 float64 = 1.5
                            switch "foo bar" {
                            case "":
                                var v139 float64 = ((0.0 * v138) * 0.1)
                                _ = v139
                            }
                            var v140 int = (fi129 - fi129)
                            _ = v138
                            _ = v140
                        }
                        var v141 uint = 255
                        _ = fi129
                        _ = v141
                    }
                }
                _ = v27
                _ = v28
            }
        }
    } else {
        var v142 uint32 = 42
        var v143 uint32 = 2147483647
        _ = v142
        _ = v143
    }
    v144 := GenS11{F0: 772, F1: 100, F2: 128}
    v145 := []int{1, -867, -9223372036854775808}
    fmt.Println(v8)
    for fi146 := 0; fi146 < 16; fi146++ {
        if false {
            v147 := GenS20{F0: 672, F1: fi146}
            _ = v147
        } else {
            for fi148 := 0; fi148 < 8; fi148++ {
                if (-473 > 647) {
                    v149 := GenS4{F0: 1, F1: -128}
                    if ((!((true && (1.5 <= 1e10)) || (false || (1 < 0)))) && true) {
                        var v150 int8 = 100
                        _ = v150
                    }
                    var v151 uint8 = 255
                    _ = v149
                    _ = v151
                } else {
                    switch v9 {
                    case "hello":
                        v8 = 0.0
                    default:
                        switch "hello" {
                        case "":
                            if false {
                                v152 := GenS21{F0: 4294967295}
                                v153 := GenS11{F0: (fi146 | (fi148 & fi148)), F1: 100, F2: (fi148 / ((0) | 1))}
                                v9 = "hello"
                                _ = v152
                                _ = v153
                            }
                        case "\n\t":
                            if (float64(v8) > (1e10 / 1.0)) {
                                if (true || false) {
                                    switch (fi146 | (fi148 & (fi146 + (fi146 - -1)))) {
                                    case 1:
                                        var v154 uint8 = 2
                                        _ = v154
                                    default:
                                        var v155 uint64 = 2
                                        _ = v155
                                    }
                                }
                                if true {
                                    v156 := GenS6{F0: fi146, F1: int16(fi146)}
                                    _ = v156
                                }
                                var v157 int8 = -9
                                v158 := GenS22{F0: 368, F1: 0, F2: -2147483648}
                                v159 := GenS9{F0: 1}
                                _ = v157
                                _ = v158
                                _ = v159
                            }
                        default:
                            v160 := GenS23{F0: 1}
                            _ = v160
                        }
                    }
                    v9 = v9
                    if (-1 != 2147483647) {
                        v161 := GenS23{F0: 4294967295}
                        _ = v161
                    }
                }
                var v162 int64 = -365
                _ = fi148
                _ = v162
            }
        }
        switch v9 {
        case "":
            v9 = v9
        case "hello":
            v8 += (0.0 * v8)
        case "\n\t":
            v163 := GenS24{F0: 1, F1: fi146}
            _ = v163
        }
        var v164 int16 = 0
        _ = fi146
        _ = v164
    }
    _ = len(v145)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v145:", len(v145), v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v144
    _ = v145
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
