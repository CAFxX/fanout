package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 uint; F1 uint; F2 int32 }
type GenS12 struct { F0 uint; F1 int16; F2 uint16 }
type GenS13 struct { F0 bool; F1 uint64 }
type GenS14 struct { F0 uint8; F1 uint16; F2 uint64 }

func main() {
    var v1 uint32 = 801
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    v1 = (^(v1))
    fmt.Println(v3, v1)
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 2, F1: 255, F2: 938}
        var v6 uint16 = 1
        v7 := GenS2{F0: 65535, F1: -270, F2: fi4}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v3 += -337
    v8 := GenS2{F0: 861, F1: 136, F2: 32768}
    v3 -= v3
    v1 += v1
    v1 += 390
    v2[0] = -250
    v1 = (^(v1))
    fmt.Println(v1, v1)
    v9 := v2[0]
    v3 += (v3 / (((v3 * v3)) | 1))
    v1 += (v1 % ((2) | 1))
    v10 := GenS1{F0: 0, F1: 127, F2: 42}
    switch "hello" {
    case "a":
        switch v9 {
        case 32768:
            var v11 uint8 = 2
            _ = v11
        case 128:
            fmt.Println(v3, v3, v9)
        case -1:
            v12 := []int{}
            _ = v12
        default:
            if (999 <= -161) {
                if (-32768 >= 702) {
                    switch 13 {
                    case 0:
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 float64 = 1.5
                            fmt.Println(v1, v1)
                            for fi15 := 0; fi15 < 1; fi15++ {
                                v14 = v14
                                _ = fi15
                            }
                            _ = fi13
                            _ = v14
                        }
                    case 1:
                        v16 := GenS2{F0: 2, F1: (-(v9)), F2: 42}
                        _ = v16
                    case 4294967295:
                        if (!(0 <= v3)) {
                            v17 := GenS3{F0: 403, F1: 0, F2: 856}
                            _ = v17
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v19 := []int{fi18}
                            for fi20 := 0; fi20 < 12; fi20++ {
                                switch ("foo bar" + ("foo bar" + "")) {
                                case "hello":
                                    fmt.Println(fi18, v1, fi20)
                                }
                                if ((^(fi18)) == (v9 ^ fi20)) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v21 uint8 = 2
                                        _ = v21
                                    case "a":
                                        v22 := GenS1{F0: 68, F1: 79, F2: 576}
                                        _ = v22
                                    case "foo bar":
                                        v23 := GenS4{F0: 32767}
                                        _ = v23
                                    default:
                                        v3 -= v3
                                    }
                                } else {
                                    var v24 int64 = -584
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v3, fi20)
                                    case "a":
                                        v24 = v3
                                    default:
                                        if (!true) {
                                            v25 := GenS2{F0: 636, F1: fi20, F2: (v9 | v9)}
                                            v2[fi18] = (fi20 - fi18)
                                            fmt.Println(v9, fi18, fi20)
                                            v26 := GenS2{F0: 659, F1: -1, F2: (^(fi20))}
                                            v27 := GenS5{F0: v3, F1: true, F2: 417}
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v9 = fi18
                                            v28 := GenS6{F0: 9223372036854775807}
                                            _ = v28
                                        }
                                    }
                                    _ = v24
                                }
                                _ = fi20
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        switch (v9 % (((v9 << ((2) & 63))) | 1)) {
                        case 80:
                            if true {
                                var v29 string = "a"
                                v29 = (v29 + ("foo bar" + (("a" + v29) + v29)))
                                _ = v29
                            } else {
                                _ = len(v2)
                                fmt.Println(v1)
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    fmt.Println(fi30, v9)
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        if false {
                                            _ = len(v2)
                                            v3 = (v3 * (v3 - v3))
                                        } else {
                                            var v32 int8 = 1
                                            if false {
                                                if true {
                                                    v33 := GenS4{F0: 42}
                                                    if ((0.1 == 1e10) && (true && (false || true))) {
                                                        v3 = (v3 * (v3 ^ (v3 + (-(v3)))))
                                                        switch v9 {
                                                        case 32768:
                                                            v1 = v1
                                                        case -720:
                                                            v9 = (-(fi31))
                                                        default:
                                                            v1 += 0
                                                        }
                                                        v34 := []int{(v9 & -9223372036854775808)}
                                                        _ = v34
                                                    }
                                                    var v35 uint64 = 130
                                                    _ = v33
                                                    _ = v35
                                                } else {
                                                    v1 = (v1 << ((2) & 63))
                                                    switch ("" + "\n\t") {
                                                    case "hello":
                                                        if true {
                                                            var v36 int32 = int32(v32)
                                                            _ = v36
                                                        }
                                                    }
                                                }
                                            } else {
                                                var v37 bool = false
                                                v2[(fi31 / (((-(v9))) | 1))] = fi30
                                                v32 -= -111
                                                _ = v37
                                            }
                                            switch (fi30 * v9) {
                                            case 32767:
                                                var v38 int8 = -35
                                                _ = v38
                                            case 2147483647:
                                                fmt.Println(v9, fi30)
                                            case -749:
                                                v3 = (v3 * v3)
                                            default:
                                                v39 := GenS7{F0: (v9 << ((96) & 63))}
                                                _ = v39
                                            }
                                            _ = v32
                                        }
                                        v1 = 659
                                        fmt.Println(v9)
                                        _ = fi31
                                    }
                                    var v40 int = (v9 % ((2147483647) | 1))
                                    _ = fi30
                                    _ = v40
                                }
                            }
                        case 32767:
                            v3 += 852
                        }
                    case "foo bar":
                        _ = len(v2)
                    default:
                        var v41 bool = (42 <= -42)
                        _ = v41
                    }
                }
                var v42 int64 = (v3 >> ((18446744073709551615) & 63))
                _ = v42
            } else {
                if (false || (false || false)) {
                    _ = len(v2)
                } else {
                    switch "" {
                    case "foo bar":
                        v3 = (v3 << ((18446744073709551615) & 63))
                    case "\n\t":
                        for fi43 := 0; fi43 < 20; fi43++ {
                            for fi44 := 0; fi44 < 13; fi44++ {
                                v1 = (v1 * 1)
                                if true {
                                    v9 -= 2147483647
                                    v3 = (v3 ^ 729)
                                    if true {
                                        if true {
                                            v1 = (v1 % (((v1 >> ((65535) & 63))) | 1))
                                        }
                                        v3 = int64(v1)
                                        v45 := []int{fi43, (v9 & (v9 + (fi43 >> ((0) & 63)))), 0}
                                        _ = v45
                                    } else {
                                        switch (-(fi44)) {
                                        case 674:
                                            v46 := GenS8{F0: -215}
                                            _ = v46
                                        default:
                                            var v47 uint16 = 255
                                            _ = v47
                                        }
                                        if (!((v3 << ((65535) & 63)) != v3)) {
                                            v3 += (v3 >> ((18446744073709551615) & 63))
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                v3 = (v3 << ((896) & 63))
                                            case "hello":
                                                delete(v2, int(v1))
                                            case "\n\t":
                                                v9 -= 32767
                                                switch "" {
                                                case "a":
                                                    v48 := GenS2{F0: 1, F1: (fi43 + (-(v9))), F2: (v9 ^ (v9 % (((v9 * fi44)) | 1)))}
                                                    _ = v48
                                                default:
                                                    v3 = 9223372036854775807
                                                }
                                            default:
                                                v49 := GenS2{F0: 18446744073709551615, F1: fi43, F2: (-(v9))}
                                                var v50 uint32 = (v1 - (v1 << ((65535) & 63)))
                                                if false {
                                                    if ((!false) || true) {
                                                        var v51 string = ("" + "a")
                                                        if (1000 == 42) {
                                                            v52 := []int{}
                                                            if (4294967295 >= 2) {
                                                                v51 = v51
                                                            }
                                                            fmt.Println(v9, v50, v50)
                                                            _ = v52
                                                        } else {
                                                            v2[-1] = fi44
                                                        }
                                                        if (!(!true)) {
                                                            if false {
                                                                v53 := GenS4{F0: 1}
                                                                var v54 bool = (fi44 == fi44)
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            v55 := GenS5{F0: v3, F1: true, F2: 2}
                                                            switch v51 {
                                                            case "":
                                                                var v56 uint64 = 42
                                                                _ = v56
                                                            case "a":
                                                                fmt.Println(fi43, v1, fi44)
                                                            }
                                                            var v57 uint32 = 42
                                                            _ = v55
                                                            _ = v57
                                                        } else {
                                                            v58 := GenS9{F0: 65535, F1: (v1 + v50)}
                                                            var v59 int = -1
                                                            if true {
                                                                if false {
                                                                    switch v51 {
                                                                    case "a":
                                                                        var v60 uint8 = 171
                                                                        _ = v60
                                                                    case "foo bar":
                                                                        v61 := v2[(-(fi43))]
                                                                        _ = v61
                                                                    default:
                                                                        fmt.Println(fi43)
                                                                    }
                                                                    v62 := GenS10{F0: 0, F1: 42}
                                                                    var v63 uint8 = 131
                                                                    switch (fi44 * (^(v9))) {
                                                                    case 1:
                                                                        var v64 uint8 = (v63 << ((2) & 63))
                                                                        _ = v64
                                                                    case -835:
                                                                        fmt.Println(v59)
                                                                    case 0:
                                                                        v50 = 1
                                                                    default:
                                                                        var v65 int32 = 728
                                                                        _ = v65
                                                                    }
                                                                    fmt.Println(fi43, v50)
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    v66 := GenS7{F0: v9}
                                                                    if false {
                                                                        if false {
                                                                            switch (^(v9)) {
                                                                            case -9223372036854775808:
                                                                                var v67 string = v51
                                                                                _ = v67
                                                                            }
                                                                            switch (v59 + fi44) {
                                                                            case -727:
                                                                                v50 = v1
                                                                            case -1:
                                                                                if true {
                                                                                    if ((v59 & fi44) != 2147483647) {
                                                                                        var v68 int64 = (v3 ^ v3)
                                                                                        _ = v68
                                                                                    }
                                                                                    fmt.Println(fi43)
                                                                                    switch fi44 {
                                                                                    case 32767:
                                                                                        v69 := []int{v9}
                                                                                        _ = v69
                                                                                    case -1:
                                                                                        v2[fi43] = fi44
                                                                                    }
                                                                                }
                                                                            case 127:
                                                                                fmt.Println(fi43, v50, fi43)
                                                                            }
                                                                            if true {
                                                                                var v70 int = fi43
                                                                                if true {
                                                                                    fmt.Println(fi43)
                                                                                    var v71 int8 = 42
                                                                                    v51 = "foo bar"
                                                                                    _ = v71
                                                                                }
                                                                                _ = v70
                                                                            } else {
                                                                                var v72 uint = 255
                                                                                _ = v72
                                                                            }
                                                                        } else {
                                                                            v73 := []int{fi43, fi44}
                                                                            var v74 int = (^(fi44))
                                                                            v59 = (v9 % ((fi43) | 1))
                                                                            v75 := GenS11{F0: 0, F1: 18446744073709551615, F2: 452}
                                                                            _ = v73
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        v76 := []int{}
                                                                        if true {
                                                                            fmt.Println(v3, v59)
                                                                        } else {
                                                                            v77 := GenS2{F0: 255, F1: v9, F2: fi44}
                                                                            var v78 uint8 = 2
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        var v79 float32 = -2.25
                                                                        var v80 uint = 18446744073709551615
                                                                        switch (fi44 + (-(fi44))) {
                                                                        case 828:
                                                                            var v81 float32 = float32(3.14159)
                                                                            _ = v81
                                                                        }
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    v50 += (v50 >> ((2) & 63))
                                                                    _ = v66
                                                                }
                                                                if true {
                                                                    switch fi44 {
                                                                    case -9223372036854775808:
                                                                        v82 := GenS11{F0: 255, F1: 1, F2: -621}
                                                                        _ = v82
                                                                    default:
                                                                        switch v51 {
                                                                        case "foo bar":
                                                                            v2[(v59 * (v59 >> ((4294967295) & 63)))] = (fi44 / ((0) | 1))
                                                                        case "":
                                                                            fmt.Println(fi43, v59)
                                                                        case "a":
                                                                            var v83 float32 = 0.1
                                                                            _ = v83
                                                                        default:
                                                                            switch fi43 {
                                                                            case -9223372036854775808:
                                                                                if ((!false) || false) {
                                                                                    var v84 float64 = float64(0.1)
                                                                                    if (v50 < (v1 / ((4294967295) | 1))) {
                                                                                        v51 = "a"
                                                                                    }
                                                                                    _ = v84
                                                                                } else {
                                                                                    v9 = fi43
                                                                                }
                                                                                v85 := GenS11{F0: 904, F1: 0, F2: 992}
                                                                                if (!(!(!true))) {
                                                                                    switch int(v3) {
                                                                                    case -253:
                                                                                        v9 += v9
                                                                                    case 32768:
                                                                                        var v86 int8 = 127
                                                                                        _ = v86
                                                                                    case 1:
                                                                                        v50 = 65535
                                                                                    default:
                                                                                        if ((false && true) && (-1 >= -32768)) {
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                var v87 uint32 = 4294967295
                                                                                                _ = v87
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    var v88 uint64 = 18446744073709551615
                                                                                                    var v89 int64 = 9223372036854775807
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                } else {
                                                                                                    var v90 int8 = 114
                                                                                                    _ = v90
                                                                                                }
                                                                                            case "":
                                                                                                if false {
                                                                                                    v51 = ("" + ((("hello" + "hello") + "a") + ("foo bar" + "hello")))
                                                                                                    v91 := GenS9{F0: 360, F1: v50}
                                                                                                    fmt.Println(fi44, v59, v50)
                                                                                                    v92 := []int{v59, (v9 + v59)}
                                                                                                    var v93 int8 = 0
                                                                                                    _ = v91
                                                                                                    _ = v92
                                                                                                    _ = v93
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v59, v50, fi43)
                                                                                            fmt.Println(v3, v9)
                                                                                            v50 += v1
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v94 uint = 255
                                                                                    v94 -= v94
                                                                                    if true {
                                                                                        v95 := v2[(v59 ^ int(v50))]
                                                                                        _ = v95
                                                                                    } else {
                                                                                        v96 := GenS7{F0: (fi44 - fi44)}
                                                                                        _ = v96
                                                                                    }
                                                                                    _ = v94
                                                                                }
                                                                                _ = v85
                                                                            case 128:
                                                                                fmt.Println(v1, v51)
                                                                            case 4294967295:
                                                                                v51 = ""
                                                                            }
                                                                        }
                                                                    }
                                                                    var v97 string = "foo bar"
                                                                    fmt.Println(v3, v9, v97)
                                                                    v50 = v50
                                                                    v98 := []int{}
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    if true {
                                                                        v59 = 127
                                                                        v2[(^(fi43))] = 32767
                                                                        v59 = int(v3)
                                                                    } else {
                                                                        fmt.Println(v3, fi43)
                                                                        v51 = (v51 + (v51 + ("" + v51)))
                                                                        delete(v2, fi44)
                                                                        fmt.Println(v51, v3, fi44)
                                                                    }
                                                                    v59 = fi43
                                                                    v99 := v2[(fi44 - fi44)]
                                                                    _ = v99
                                                                }
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v51
                                                    }
                                                    v3 += (v3 - (^(v3)))
                                                    var v100 int = (fi43 - (fi43 * (v9 | 9223372036854775807)))
                                                    var v101 uint64 = 18446744073709551615
                                                    _ = v100
                                                    _ = v101
                                                } else {
                                                    var v102 uint32 = (v1 - 1)
                                                    _ = v102
                                                }
                                                _ = v49
                                                _ = v50
                                            }
                                            fmt.Println(fi43, v1, v3)
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, v1, v9)
                                    v9 -= 1
                                }
                                v103 := GenS12{F0: 1, F1: 0, F2: 519}
                                _ = fi44
                                _ = v103
                            }
                            if false {
                                fmt.Println(v3, v9)
                                var v104 bool = (false && (!true))
                                _ = v104
                            }
                            v105 := GenS13{F0: (-272 > -2147483648), F1: 1}
                            _ = fi43
                            _ = v105
                        }
                    case "":
                        v3 += (v3 & (v3 / ((v3) | 1)))
                    }
                }
            }
        }
    case "hello":
        v3 += (v3 >> ((18446744073709551615) & 63))
    }
    if (812 <= 0) {
        if false {
            var v106 uint16 = 42
            _ = v106
        } else {
            for fi107 := 0; fi107 < 20; fi107++ {
                var v108 int16 = 0
                if (!true) {
                    v109 := GenS14{F0: 2, F1: 153, F2: 238}
                    for fi110 := 0; fi110 < 2; fi110++ {
                        v111 := v2[(fi107 & v9)]
                        _ = fi110
                        _ = v111
                    }
                    _ = v109
                }
                var v112 uint = 255
                _ = fi107
                _ = v108
                _ = v112
            }
            v113 := []int{1, (v9 ^ v9), v9}
            _ = v113
        }
        var v114 string = "a"
        switch v9 {
        case 0:
            var v115 bool = (1.5 < (float32((1e10 / (0.1 / -2.25))) / float32((((((((1e10 - float64(3.14159)) + 3.14159) - 0.0) - 0.0) * (1.5 + 0.1)) - 0.0) - 0.1))))
            _ = v115
        case 2147483647:
            for fi116 := 0; fi116 < 19; fi116++ {
                var v117 uint8 = 135
                _ = fi116
                _ = v117
            }
        case 829:
            fmt.Println(v1)
        }
        _ = v114
    } else {
        v118 := v2[v9]
        _ = v118
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
