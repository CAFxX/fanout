package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool; F1 int; F2 uint16 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 int8; F1 int64; F2 uint64 }
type GenS11 struct { F0 string }
func genFn0(p0 int32, p1 uint16, p2 uint32) string {
    switch "" {
    case "\n\t":
        fmt.Println(p0, p2)
    case "a":
        v7 := GenS1{F0: 42}
        _ = v7
    default:
        if true {
            var v8 bool = ((p2 * (p2 + (p2 % ((p2) | 1)))) > p2)
            _ = v8
        }
    }
    p0 = p0
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: 175, F1: int8(fi9)}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 int, p1 int8) int64 {
    fmt.Println(p1)
    switch p0 {
    case -2147483648:
        v11 := []int{-2147483648, 2147483647, -490, (p0 + p0)}
        _ = v11
    }
    fmt.Println(p1, p0)
    v12 := []int{p0, (p0 ^ p0), p0}
    v13 := GenS3{F0: 385, F1: 2}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 687
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 uint32 = 495
    var v4 uint32 = 55
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    switch ("hello" + ((("a" + "") + ("" + "foo bar")) + "hello")) {
    case "foo bar":
        var v14 int16 = -238
        _ = v14
    default:
        switch 733 {
        case -539:
            switch 9223372036854775807 {
            case -597:
                for fi15 := 0; fi15 < 18; fi15++ {
                    switch 4294967295 {
                    case -9223372036854775808:
                        if ((^(fi15)) > fi15) {
                            if true {
                                if true {
                                    switch fi15 {
                                    case 2147483647:
                                        if ("a" > "a") {
                                            switch (fi15 * 904) {
                                            case -2147483648:
                                                fmt.Println(v4, fi15, fi15)
                                            case 128:
                                                if false {
                                                    v16 := GenS4{F0: (v3 * v4)}
                                                    _ = v16
                                                }
                                            default:
                                                fmt.Println(v3, fi15, fi15)
                                            }
                                            v3 = v4
                                            var v17 uint8 = 255
                                            _ = v17
                                        } else {
                                            switch -1 {
                                            case 4294967295:
                                                _ = len(v2)
                                            default:
                                                v18 := GenS5{F0: int32(fi15), F1: v3}
                                                _ = v18
                                            }
                                        }
                                    }
                                }
                            } else {
                                v3 += v4
                                v4 += v4
                                switch ("\n\t" + ("\n\t" + "foo bar")) {
                                case "\n\t":
                                    for fi19 := 0; fi19 < 12; fi19++ {
                                        v3 -= v3
                                        delete(v5, (^(fi19)))
                                        _ = fi19
                                    }
                                default:
                                    v3 -= v3
                                }
                            }
                            switch ((("\n\t" + "") + "hello") + "\n\t") {
                            case "\n\t":
                                fmt.Println(v3, v3)
                            case "hello":
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    if false {
                                        v4 = (v4 ^ (v4 & 0))
                                    } else {
                                        if false {
                                            var v21 uint16 = 255
                                            switch (fi20 << ((65535) & 63)) {
                                            case -49:
                                                v22 := GenS1{F0: -1}
                                                _ = v22
                                            default:
                                                var v23 float64 = -2.25
                                                _ = v23
                                            }
                                            if false {
                                                var v24 int64 = 1
                                                v25 := v1[fi15]
                                                if true {
                                                    var v26 uint16 = v21
                                                    v27 := []int{(fi15 + (fi15 ^ 32767))}
                                                    switch int(v21) {
                                                    case 9223372036854775807:
                                                        var v28 uint8 = 255
                                                        _ = v28
                                                    case 1:
                                                        v3 += v3
                                                    case -546:
                                                        v21 = 1
                                                    }
                                                    var v29 uint16 = v26
                                                    _ = v26
                                                    _ = v27
                                                    _ = v29
                                                } else {
                                                    v2[fi20] = v3
                                                    fmt.Println(v21)
                                                    v3 -= (v3 / (((v3 / (((v3 | 65535)) | 1))) | 1))
                                                }
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = v21
                                        } else {
                                            if true {
                                                if (0.1 < -2.25) {
                                                    v30 := GenS2{F0: 1, F1: 0}
                                                    v31 := GenS5{F0: 42, F1: (v3 | 877)}
                                                    var v32 int8 = 1
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                }
                                                if false {
                                                    v3 -= v4
                                                } else {
                                                    _ = len(v6)
                                                }
                                                fmt.Println(v4, v4)
                                            }
                                            v33 := GenS2{F0: 396, F1: 100}
                                            _ = v33
                                        }
                                    }
                                    if true {
                                        v3 -= (v4 >> ((65535) & 63))
                                        if (4294967295 <= 910) {
                                            var v34 int16 = -1
                                            v35 := GenS6{F0: true, F1: (fi15 << ((4294967295) & 63)), F2: 2}
                                            fmt.Println(v3)
                                            if ("" > ("\n\t" + ("hello" + "\n\t"))) {
                                                var v36 bool = true
                                                fmt.Println(fi15)
                                                _ = v36
                                            } else {
                                                switch ("hello" + "") {
                                                case "":
                                                    var v37 int16 = (v34 | v34)
                                                    _ = v37
                                                case "hello":
                                                    v38 := GenS7{F0: (fi15 & 0), F1: 65535}
                                                    _ = v38
                                                case "\n\t":
                                                    if false {
                                                        fmt.Println(v4, fi20)
                                                        switch -9223372036854775808 {
                                                        case -2147483648:
                                                            if true {
                                                                fmt.Println(fi15, v4)
                                                                var v39 float32 = 0.0
                                                                _ = v39
                                                            }
                                                        case 128:
                                                            var v40 uint = 457
                                                            _ = v40
                                                        case -915:
                                                            switch fi20 {
                                                            case 32768:
                                                                if false {
                                                                    v41 := []int{(fi15 + 42), fi20, -263, fi20}
                                                                    _ = v41
                                                                }
                                                            default:
                                                                var v42 uint = 0
                                                                _ = v42
                                                            }
                                                        }
                                                    } else {
                                                        v43 := GenS1{F0: -692}
                                                        _ = v43
                                                    }
                                                }
                                                v4 -= (v4 / ((v3) | 1))
                                            }
                                            switch fi20 {
                                            case -2147483648:
                                                v3 = (v3 >> ((42) & 63))
                                            case -9223372036854775808:
                                                if true {
                                                    if (540 <= 65535) {
                                                        if false {
                                                            switch "a" {
                                                            case "":
                                                                v44 := GenS2{F0: 18446744073709551615, F1: 0}
                                                                _ = v44
                                                            default:
                                                                v4 -= 87
                                                            }
                                                            v4 += (^(v3))
                                                        }
                                                    }
                                                } else {
                                                    if (fi15 > (fi15 << ((444) & 63))) {
                                                        var v45 int = fi15
                                                        if true {
                                                            switch (((("foo bar" + ("hello" + (("\n\t" + "hello") + "foo bar"))) + "hello") + ("a" + "hello")) + "a") {
                                                            case "":
                                                                fmt.Println(v45, fi15)
                                                            }
                                                            v4 += 486
                                                            v34 = (v34 ^ v34)
                                                            v46 := GenS3{F0: v3, F1: 4294967295}
                                                            delete(v6, (^(v3)))
                                                            _ = v46
                                                        }
                                                        _ = v45
                                                    }
                                                    var v47 int = fi20
                                                    _ = v47
                                                }
                                            }
                                            v4 = (v3 % (((^(v3))) | 1))
                                            if false {
                                                v48 := GenS8{F0: 15, F1: 18446744073709551615, F2: 1}
                                                v4 = 816
                                                _ = v48
                                            } else {
                                                v4 += 0
                                                if (255 <= 42) {
                                                    var v49 uint32 = v4
                                                    v49 -= v3
                                                    v4 = v4
                                                    _ = v49
                                                }
                                                if true {
                                                    v50 := []int{0, 127, (fi15 * fi15)}
                                                    _ = v50
                                                }
                                            }
                                            var v51 int64 = 1
                                            _ = v34
                                            _ = v35
                                            _ = v51
                                        } else {
                                            switch (fi20 & 950) {
                                            case 127:
                                                var v52 uint = 255
                                                _ = v52
                                            case 835:
                                                if true {
                                                    var v53 uint = 65535
                                                    var v54 int8 = 42
                                                    var v55 float32 = (3.14159 + (((1.5 + 3.14159) - 0.1) * 1e10))
                                                    fmt.Println(fi20, fi20)
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    if false {
                                                        if false {
                                                            v56 := GenS9{F0: 765, F1: 42}
                                                            _ = v56
                                                        } else {
                                                            if false {
                                                                if ((v3 % ((65535) | 1)) != v3) {
                                                                    v57 := GenS2{F0: 414, F1: 123}
                                                                    _ = v57
                                                                }
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v58 := GenS9{F0: 42, F1: -102}
                                                                    _ = v58
                                                                case "\n\t":
                                                                    switch ((("a" + (("" + "hello") + ("foo bar" + "a"))) + "\n\t") + ("a" + "\n\t")) {
                                                                    case "":
                                                                        fmt.Println(fi20, fi20, fi15)
                                                                    }
                                                                case "hello":
                                                                    var v59 int32 = 2147483647
                                                                    _ = v59
                                                                }
                                                                v4 += v4
                                                            } else {
                                                                v60 := GenS6{F0: false, F1: (fi15 / ((-1) | 1)), F2: 2}
                                                                switch (fi20 >> ((252) & 63)) {
                                                                case 2147483647:
                                                                    var v61 int32 = 42
                                                                    _ = v61
                                                                case -872:
                                                                    var v62 int = fi20
                                                                    _ = v62
                                                                default:
                                                                    var v63 int32 = -102
                                                                    v6[v4] = (v3 / (((v4 & (v3 - 2147483648))) | 1))
                                                                    _ = v63
                                                                }
                                                                if true {
                                                                    if true {
                                                                        fmt.Println(fi20, fi15, v3)
                                                                        v64 := []int{(fi15 * (fi15 * -1))}
                                                                        var v65 uint32 = (v3 * (v3 % (((v3 | (^(v3)))) | 1)))
                                                                        _ = v64
                                                                        _ = v65
                                                                    } else {
                                                                        var v66 string = "\n\t"
                                                                        _ = v66
                                                                    }
                                                                    if (true && ((403 != -32768) && false)) {
                                                                        switch (fi15 - fi15) {
                                                                        case -324:
                                                                            v3 += v4
                                                                        case 1:
                                                                            v4 -= 65535
                                                                        }
                                                                        v3 = 936
                                                                        fmt.Println(v4, v4, fi15)
                                                                    }
                                                                }
                                                                _ = v60
                                                            }
                                                            switch -9223372036854775808 {
                                                            case -495:
                                                                fmt.Println(fi20, v4)
                                                            case 32767:
                                                                v67 := v6[4294967295]
                                                                _ = v67
                                                            case 2147483647:
                                                                var v68 float32 = 1.5
                                                                _ = v68
                                                            }
                                                        }
                                                        if ("" == "a") {
                                                            v69 := GenS2{F0: 0, F1: int8(fi20)}
                                                            fmt.Println(fi20)
                                                            _ = v69
                                                        } else {
                                                            v70 := GenS10{F0: -97, F1: -847, F2: 2}
                                                            switch "foo bar" {
                                                            case "":
                                                                v71 := []int{(^(fi15))}
                                                                _ = v71
                                                            }
                                                            v72 := GenS3{F0: 4294967295, F1: 483}
                                                            v73 := GenS7{F0: fi20, F1: 414}
                                                            fmt.Println(fi15, v4)
                                                            _ = v70
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        var v74 uint16 = 65535
                                                        _ = v74
                                                    } else {
                                                        fmt.Println(v4, v4)
                                                    }
                                                    v75 := []int{fi20}
                                                    v4 += v3
                                                    _ = v75
                                                }
                                            case 2147483647:
                                                fmt.Println(v4)
                                            default:
                                                switch 2147483647 {
                                                case -210:
                                                    if true {
                                                        if true {
                                                            v4 += (^(v4))
                                                            v3 += (v4 * 4294967295)
                                                            v76 := GenS6{F0: ((((1.5 * 1.5) - 1.5) != (((1e10 * (1.5 - 1e10)) / float64((-2.25 - 1e10))) - (-2.25 / 1e10))) || (-1 > -2147483648)), F1: (fi20 << ((42) & 63)), F2: 1}
                                                            _ = v76
                                                        } else {
                                                            switch fi15 {
                                                            case 127:
                                                                switch 32767 {
                                                                case -476:
                                                                    v4 -= v4
                                                                case 442:
                                                                    fmt.Println(v4)
                                                                case -9223372036854775808:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v77 int64 = 42
                                                                        _ = v77
                                                                    case "a":
                                                                        v3 -= (^(v4))
                                                                    default:
                                                                        var v78 int = -581
                                                                        _ = v78
                                                                    }
                                                                }
                                                            case 32768:
                                                                if false {
                                                                    v4 -= (v3 * v4)
                                                                }
                                                            case -761:
                                                                if true {
                                                                    fmt.Println(fi15)
                                                                    v3 += (^(v3))
                                                                } else {
                                                                    var v79 int = fi20
                                                                    _ = v79
                                                                }
                                                            }
                                                            fmt.Println(fi20, fi20)
                                                        }
                                                        v3 += 65535
                                                        v80 := GenS7{F0: (fi20 / ((fi15) | 1)), F1: 255}
                                                        _ = v80
                                                    } else {
                                                        v81 := v1[fi15]
                                                        v82 := []int{}
                                                        fmt.Println(v81)
                                                        v81 = v81
                                                        v83 := v5[(fi15 | 1)]
                                                        v84 := GenS11{F0: ("a" + v81)}
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    fmt.Println(v3, fi20, v4)
                                                case 945:
                                                    var v85 uint8 = 42
                                                    v85 -= (v85 % ((255) | 1))
                                                    _ = v85
                                                case -1:
                                                    var v86 uint = 1
                                                    _ = v86
                                                default:
                                                    var v87 uint = 18446744073709551615
                                                    _ = v87
                                                }
                                            }
                                        }
                                        var v88 int8 = -7
                                        v3 = (v3 * (v4 / ((1) | 1)))
                                        if (!(18446744073709551615 != 2)) {
                                            fmt.Println(v3)
                                            v4 += (v3 << ((65535) & 63))
                                        } else {
                                            v4 = (v3 << ((65535) & 63))
                                            var v89 float32 = ((0.1 * (0.1 * 0.1)) / (((float32(1e10) + 0.1) / 1.0) - 1e10))
                                            switch fi20 {
                                            case 42:
                                                fmt.Println(fi15)
                                            case -863:
                                                v89 = (v89 - v89)
                                            }
                                            _ = v89
                                        }
                                        fmt.Println(fi20, v4)
                                        _ = v88
                                    }
                                    _ = fi20
                                }
                            default:
                                v3 += (v4 & 42)
                            }
                            v3 = 2
                        }
                    case 9223372036854775807:
                        for fi90 := 0; fi90 < 7; fi90++ {
                            var v91 uint64 = 0
                            var v92 int = fi15
                            _ = fi90
                            _ = v91
                            _ = v92
                        }
                    default:
                        var v93 float32 = 1e10
                        _ = v93
                    }
                    _ = fi15
                }
            case 505:
                for fi94 := 0; fi94 < 15; fi94++ {
                    v4 = (v3 ^ v4)
                    var v95 int8 = -128
                    _ = fi94
                    _ = v95
                }
            }
        case 9223372036854775807:
            v96 := []int{512, 128}
            _ = v96
        case 1:
            var v97 float64 = float64(3.14159)
            _ = v97
        }
    }
    var v98 int32 = -1
    switch 4294967295 {
    case -2147483648:
        var v99 int = 1
        _ = v99
    }
    var v100 uint16 = 2
    var v101 uint32 = (v3 - 42)
    for fi102 := 0; fi102 < 7; fi102++ {
        v100 = (v100 / (((v100 & (v100 | 42))) | 1))
        _ = fi102
    }
    v4 = v101
    var v103 int16 = 934
    fmt.Println(v103, v3, v100)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v101
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
