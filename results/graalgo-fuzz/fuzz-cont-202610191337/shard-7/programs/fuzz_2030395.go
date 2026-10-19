package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint; F1 int32; F2 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int; F2 bool }
type GenS10 struct { F0 uint32; F1 int }

func main() {
    var v1 string = ("" + "")
    var v2 int8 = 0
    var v3 uint = 65535
    var v4 uint64 = 18446744073709551615
    var v5 int64 = -1
    switch v1 {
    case "a":
        switch "a" {
        case "hello":
            var v6 uint64 = (v4 << (((v3 + v3)) & 63))
            _ = v6
        default:
            v7 := GenS1{F0: uint(v4)}
            _ = v7
        }
    case "":
        switch -344 {
        case -500:
            v8 := []int{-9223372036854775808, -1, -305}
            _ = v8
        case 450:
            if false {
                v3 = v3
                var v9 int = 4294967295
                v10 := []int{v9, 4294967295}
                _ = v10
                _ = v9
            }
        case 0:
            v11 := GenS1{F0: 913}
            _ = v11
        }
    default:
        if false {
            v5 = v5
            if ((v4 ^ (v4 >> ((v3) & 63))) < 63) {
                var v12 int = -9223372036854775808
                var v13 float32 = ((1.5 * (3.14159 - ((0.1 + (1.5 + (0.1 / (-2.25 + 3.14159)))) / 0.1))) - float32((0.1 * 3.14159)))
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int16 = 0
                fmt.Println(v4, fi14, v5)
                var v16 float32 = 0.1
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 int32 = -2147483648
                    _ = fi18
                    _ = v19
                }
                var v20 float32 = 1.5
                var v21 int = (fi17 << ((v3) & 63))
                _ = fi17
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int64 = (v5 - 0)
                _ = fi22
                _ = v23
            }
            v24 := []int{1, -700}
            _ = v24
        }
    }
    v25 := GenS2{F0: v5}
    switch 32768 {
    case -45:
        v3 -= 255
    case 9223372036854775807:
        v26 := []int{}
        _ = v26
    case 32768:
        var v27 float32 = 3.14159
        _ = v27
    default:
        for fi28 := 0; fi28 < 2; fi28++ {
            switch fi28 {
            case -817:
                if false {
                    if (!(255 <= 42)) {
                        var v29 uint32 = 65535
                        _ = v29
                    } else {
                        var v30 int = 127
                        v31 := []int{int(v5), (^(fi28))}
                        _ = v30
                        _ = v31
                    }
                    v2 += v2
                    v4 += 765
                } else {
                    if false {
                        var v32 int32 = 42
                        var v33 uint8 = 42
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v35 := GenS2{F0: (v5 << (((^(v3))) & 63))}
                            v36 := GenS3{F0: 1, F1: -22, F2: v33}
                            if false {
                                if true {
                                    fmt.Println(v1, v4)
                                    var v37 int64 = v5
                                    _ = v37
                                }
                                var v38 int32 = v32
                                var v39 uint = 778
                                _ = v38
                                _ = v39
                            }
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        _ = v32
                        _ = v33
                    }
                    for fi40 := 0; fi40 < 18; fi40++ {
                        if ((-1 > int16(v5)) || false) {
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    var v41 int32 = 1
                                    _ = v41
                                }
                            default:
                                if true {
                                    v1 = v1
                                    if false {
                                        v42 := GenS3{F0: 2147483648, F1: -32768, F2: 242}
                                        _ = v42
                                    } else {
                                        switch fi40 {
                                        case 636:
                                            switch v1 {
                                            case "\n\t":
                                                v43 := GenS4{F0: (fi40 / ((-878) | 1)), F1: v5, F2: 0}
                                                _ = v43
                                            case "":
                                                var v44 uint = v3
                                                _ = v44
                                            case "a":
                                                var v45 float32 = (-2.25 - 0.0)
                                                _ = v45
                                            default:
                                                var v46 string = v1
                                                _ = v46
                                            }
                                        case 2147483647:
                                            fmt.Println(v2, v4, fi40)
                                        case 187:
                                            if false {
                                                v1 = v1
                                            } else {
                                                if false {
                                                    if (((v4 + v4) != (v4 % ((v4) | 1))) || ((!(false && (false || (0 == 2147483647)))) || true)) {
                                                        v47 := []int{(fi40 + fi40), 1, -108, fi28}
                                                        _ = v47
                                                    } else {
                                                        var v48 uint16 = 32767
                                                        _ = v48
                                                    }
                                                } else {
                                                    if true {
                                                        if false {
                                                            var v49 uint = v3
                                                            switch int(v3) {
                                                            case 42:
                                                                v4 += (v4 % ((v4) | 1))
                                                                fmt.Println(fi40, fi28, v1)
                                                                if true {
                                                                    switch v1 {
                                                                    case "foo bar":
                                                                        v1 = ""
                                                                    default:
                                                                        switch v1 {
                                                                        case "hello":
                                                                            v50 := GenS5{F0: (v2 ^ (v2 - v2))}
                                                                            v4 += 4294967295
                                                                            fmt.Println(v2)
                                                                            v51 := GenS4{F0: (fi40 / ((-2147483648) | 1)), F1: v5, F2: 267}
                                                                            _ = v50
                                                                            _ = v51
                                                                        case "a":
                                                                            v1 = (v1 + "foo bar")
                                                                        default:
                                                                            v52 := GenS6{F0: fi28, F1: 0, F2: v5}
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                } else {
                                                                    v53 := []int{-1, (fi28 * fi40), fi28, 210}
                                                                    var v54 int8 = v2
                                                                    switch (fi40 % ((128) | 1)) {
                                                                    case 32768:
                                                                        fmt.Println(v5, v1, v54)
                                                                    case 1:
                                                                        var v55 string = v1
                                                                        _ = v55
                                                                    case 380:
                                                                        v56 := GenS3{F0: 2, F1: 42, F2: 42}
                                                                        _ = v56
                                                                    }
                                                                    v57 := GenS4{F0: fi40, F1: (v5 * 1), F2: 65535}
                                                                    var v58 uint16 = 2
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                            case 2147483647:
                                                                fmt.Println(v3, v2, fi28)
                                                            }
                                                            var v59 int64 = 9223372036854775807
                                                            v60 := []int{}
                                                            _ = v49
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        switch v1 {
                                                        case "hello":
                                                            v61 := []int{fi40, int(v2)}
                                                            _ = v61
                                                        case "":
                                                            if true {
                                                                if true {
                                                                    fmt.Println(v4, fi40, fi28)
                                                                    var v62 uint8 = 2
                                                                    var v63 int16 = -734
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                var v64 int = (-(fi28))
                                                                _ = v64
                                                            }
                                                        case "\n\t":
                                                            v5 -= 1
                                                        }
                                                        if (!(518 != 951)) {
                                                            var v65 float32 = 3.14159
                                                            _ = v65
                                                        } else {
                                                            v1 = "foo bar"
                                                            v1 = v1
                                                        }
                                                        if ((!(!true)) && ((fi28 * (fi40 / (((fi40 % ((-284) | 1))) | 1))) == (fi28 + (fi28 - (-(fi40)))))) {
                                                            v66 := []int{(-(fi40))}
                                                            if false {
                                                                var v67 uint = v3
                                                                v66 = append(v66, fi28)
                                                                if (!false) {
                                                                    if len(v66) > 0 {
                                                                        v66[0] = (fi28 << (((v3 << ((v67) & 63))) & 63))
                                                                    }
                                                                    v4 = 1
                                                                    v68 := GenS5{F0: 100}
                                                                    switch fi28 {
                                                                    case 69:
                                                                        v67 += (v3 << ((1) & 63))
                                                                    case 4294967295:
                                                                        if (int8(v3) >= (v2 >> ((v3) & 63))) {
                                                                            var v69 float32 = 3.14159
                                                                            v70 := GenS1{F0: (v67 * (v67 / (((v67 ^ (v3 + 397))) | 1)))}
                                                                            _ = v69
                                                                            _ = v70
                                                                        }
                                                                    case 0:
                                                                        if (391 >= 42) {
                                                                            v5 -= v5
                                                                            v4 = 429
                                                                        }
                                                                    }
                                                                    v71 := GenS2{F0: (v5 ^ (v5 * (v5 ^ (v5 | (-(v5))))))}
                                                                    fmt.Println(v1)
                                                                    _ = v68
                                                                    _ = v71
                                                                } else {
                                                                    var v72 string = v1
                                                                    var v73 uint64 = (v4 / ((v4) | 1))
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                _ = v67
                                                            }
                                                            _ = v66
                                                        } else {
                                                            var v74 int16 = 32767
                                                            v75 := []int{(fi40 >> ((v3) & 63))}
                                                            if false {
                                                                var v76 uint8 = 151
                                                                _ = v76
                                                            }
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    } else {
                                                        switch (fi40 * fi40) {
                                                        case 32767:
                                                            var v77 float64 = ((-2.25 - 1.5) / 1.0)
                                                            _ = v77
                                                        }
                                                        v3 = (v3 ^ (^(v3)))
                                                    }
                                                }
                                                v78 := GenS6{F0: 9223372036854775807, F1: 637, F2: (v5 / ((-1) | 1))}
                                                switch int(v5) {
                                                case 1:
                                                    if ((205 == 42) || true) {
                                                        var v79 uint = v3
                                                        _ = v79
                                                    }
                                                case -9223372036854775808:
                                                    if true {
                                                        fmt.Println(v1)
                                                        switch v1 {
                                                        case "":
                                                            v3 += (v3 << ((v3) & 63))
                                                        case "hello":
                                                            var v80 uint32 = 65535
                                                            _ = v80
                                                        }
                                                        v81 := GenS1{F0: v3}
                                                        _ = v81
                                                    }
                                                default:
                                                    v82 := GenS7{F0: (v3 << ((866) & 63)), F1: -2147483648, F2: -32768}
                                                    _ = v82
                                                }
                                                _ = v78
                                            }
                                        default:
                                            if true {
                                                fmt.Println(fi40, v5, fi40)
                                            } else {
                                                var v83 int64 = v5
                                                v2 += v2
                                                v84 := GenS6{F0: (fi28 + fi40), F1: 65535, F2: (v83 - (v5 - v5))}
                                                var v85 string = v1
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                        var v86 float64 = 3.14159
                                        switch "hello" {
                                        case "\n\t":
                                            v5 = v5
                                        }
                                        switch fi40 {
                                        case 127:
                                            v87 := []int{(fi40 * fi40), (fi40 % (((fi28 >> ((uint(v4)) & 63))) | 1)), 9223372036854775807, 127}
                                            var v88 uint = v3
                                            _ = v87
                                            _ = v88
                                        default:
                                            v2 = v2
                                        }
                                        _ = v86
                                    }
                                    switch (fi40 % ((0) | 1)) {
                                    case 2147483647:
                                        var v89 uint64 = v4
                                        _ = v89
                                    }
                                }
                            }
                        } else {
                            if true {
                                if (uint8(v3) > 76) {
                                    fmt.Println(fi28)
                                    v90 := []int{fi28, (fi40 << ((v3) & 63)), (fi40 * (^(fi28)))}
                                    fmt.Println(fi28, v5)
                                    _ = v90
                                }
                                v4 = (v4 % ((4294967295) | 1))
                                v91 := GenS8{F0: (v2 >> (((v3 * 2)) & 63))}
                                _ = v91
                            } else {
                                v1 = ""
                                var v92 int8 = 88
                                _ = v92
                            }
                            v1 = ""
                            var v93 uint = v3
                            _ = v93
                        }
                        fmt.Println(v2, fi40)
                        _ = fi40
                    }
                }
            case 564:
                switch (fi28 << ((v3) & 63)) {
                case -9223372036854775808:
                    var v94 uint64 = v4
                    _ = v94
                case 2147483647:
                    var v95 uint8 = 2
                    _ = v95
                default:
                    var v96 uint = v3
                    _ = v96
                }
            }
            _ = fi28
        }
    }
    switch 128 {
    case -90:
        var v97 int64 = (v5 >> (((v3 & (v3 % (((v3 ^ 4294967295)) | 1)))) & 63))
        _ = v97
    case 9223372036854775807:
        switch 1 {
        case 9223372036854775807:
            v98 := GenS5{F0: -128}
            _ = v98
        case 954:
            v99 := GenS9{F0: 2147483648, F1: -269, F2: (false || (!((!(((0 == 848) || true) || (!false))) || (!true))))}
            _ = v99
        default:
            v100 := []int{}
            _ = v100
        }
    default:
        if false {
            var v101 int64 = (v5 % (((-(v5))) | 1))
            v3 -= (v3 * (v3 ^ (v3 / (((^(v3))) | 1))))
            var v102 int8 = (v2 / ((int8(v5)) | 1))
            _ = v101
            _ = v102
        } else {
            v103 := []int{42, 9223372036854775807}
            v103 = append(v103, 42)
            _ = v103
        }
    }
    v1 = (v1 + ((v1 + v1) + (v1 + "\n\t")))
    var v104 int = 0
    v104 -= (v104 >> (((v3 * (^(v3)))) & 63))
    var v105 string = (v1 + v1)
    var v106 uint = 0
    v107 := GenS10{F0: 42, F1: v104}
    switch ((v1 + v1) + "\n\t") {
    case "":
        v108 := GenS2{F0: v5}
        _ = v108
    case "hello":
        if (v2 != 1) {
            v5 = (v5 >> ((v106) & 63))
        }
    case "a":
        v1 = (v105 + v1)
    default:
        for fi109 := 0; fi109 < 20; fi109++ {
            var v110 uint64 = (v4 >> ((v106) & 63))
            _ = fi109
            _ = v110
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
