package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint }
type GenS6 struct { F0 int16; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int32; F2 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 uint32; F2 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 int64, p2 int) int {
    v4 := []int{(p2 + (p2 << ((0) & 63))), (p2 - (-(p2))), (p2 + -1)}
    switch p2 {
    case 1:
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                p2 -= fi6
                if true {
                    var v7 string = (p0 + p0)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v7
                } else {
                    var v8 int8 = 91
                    v9 := GenS1{F0: p1, F1: -9223372036854775808, F2: p1}
                    v10 := GenS1{F0: p1, F1: -474, F2: 503}
                    _ = v10
                    _ = v8
                    _ = v9
                }
                _ = fi6
            }
            _ = fi5
        }
    case 128:
        if true {
            if len(v4) > 0 {
                v4[0] = (p2 | p2)
            }
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                p1 += (p1 % ((-9223372036854775808) | 1))
                _ = fi11
            }
        }
    default:
        var v12 bool = (!true)
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        p0 = p0
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}
func genFn1(p0 int16, p1 int32, p2 uint8) int {
    fmt.Println(p0, p1, p2)
    if (!true) {
        p1 = -341
    } else {
        fmt.Println(p1, p2, p0)
        switch 4294967295 {
        case 4294967295:
            switch -9223372036854775808 {
            case -2147483648:
                p0 += p0
            case -9223372036854775808:
                fmt.Println(p1, p0)
            default:
                var v14 float32 = 3.14159
                _ = v14
            }
        default:
            if (true || false) {
                v15 := []int{127, 888}
                v16 := GenS1{F0: -77, F1: 42, F2: -450}
                _ = v15
                _ = v16
            }
        }
        var v17 uint64 = 922
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 65535
    var v3 uint = 649
    switch ((("a" + ("" + "a")) + (("foo bar" + "foo bar") + (("" + "foo bar") + ("hello" + "foo bar")))) + ("hello" + ("foo bar" + ("hello" + ("" + "foo bar"))))) {
    case "a":
        v18 := GenS2{F0: 32767}
        _ = v18
    }
    switch -9223372036854775808 {
    case 32767:
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 bool = (1 >= 100)
            v21 := GenS3{F0: v3, F1: 42}
            _ = fi19
            _ = v20
            _ = v21
        }
    case 1:
        var v22 string = "foo bar"
        _ = v22
    default:
        switch -1 {
        case -790:
            v2 += (^(v2))
        case 959:
            switch "foo bar" {
            case "foo bar":
                switch 0 {
                case 553:
                    v23 := GenS3{F0: (v3 << (((v3 << ((v3) & 63))) & 63)), F1: 4294967295}
                    _ = v23
                case -404:
                    v2 = 689
                }
            case "a":
                for fi24 := 0; fi24 < 18; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if ((0.1 < 0.0) || ((1.5 + 0.0) == 0.1)) {
                            var v26 float32 = (float32((0.1 + (1.5 + 0.0))) * (-2.25 * (1.5 + ((0.1 + 0.0) + float32((-2.25 * 0.0))))))
                            if (!(v2 < 65535)) {
                                switch (fi24 << ((uint(v2)) & 63)) {
                                case 42:
                                    if (1.5 > float32(3.14159)) {
                                        var v27 float32 = 0.0
                                        fmt.Println(v27)
                                        _ = v27
                                    }
                                }
                            }
                            _ = v26
                        }
                        _ = fi25
                    }
                    v2 += uint32(v1)
                    _ = fi24
                }
            case "hello":
                for fi28 := 0; fi28 < 8; fi28++ {
                    if (-9223372036854775808 > 9223372036854775807) {
                        v29 := []int{(fi28 ^ (fi28 % ((fi28) | 1))), int(v1), 963}
                        var v30 uint32 = (v2 << ((v3) & 63))
                        v31 := GenS2{F0: 32767}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        v3 = 1
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v33 := []int{int(v2), fi32, (fi28 / ((fi28) | 1))}
                            fmt.Println(v2, v2, v2)
                            _ = fi32
                            _ = v33
                        }
                        switch fi28 {
                        case 128:
                            var v34 int32 = 981
                            _ = v34
                        default:
                            var v35 uint64 = 230
                            _ = v35
                        }
                    }
                    _ = fi28
                }
            default:
                var v36 bool = true
                _ = v36
            }
        case 9223372036854775807:
            var v37 int8 = 121
            _ = v37
        }
    }
    if (false || true) {
        if false {
            v3 = v3
        } else {
            v1 = (v1 & (v1 % ((uint8(v2)) | 1)))
        }
    } else {
        if (false || (!false)) {
            v38 := []int{genFn1(450, -2147483648, 128), -2147483648, 9223372036854775807}
            _ = v38
        }
    }
    if false {
        for fi39 := 0; fi39 < 9; fi39++ {
            switch fi39 {
            case 0:
                v2 = 188
            case -820:
                var v40 float64 = (0.0 * 1e10)
                _ = v40
            }
            for fi41 := 0; fi41 < 16; fi41++ {
                v42 := []int{(fi39 + fi41), (^(fi41))}
                v42 = append(v42, 668)
                var v43 uint8 = 2
                switch fi41 {
                case 32767:
                    fmt.Println(fi39, fi39, v43)
                case -465:
                    switch ("" + "hello") {
                    case "hello":
                        _ = len(v42)
                    }
                }
                var v44 uint64 = 42
                _ = len(v42)
                if true {
                    switch "a" {
                    case "\n\t":
                        v45 := GenS1{F0: 0, F1: 44, F2: 713}
                        v46 := GenS4{F0: 0, F1: 18446744073709551615, F2: (v43 >> ((18446744073709551615) & 63))}
                        fmt.Println(v43)
                        _ = v45
                        _ = v46
                    }
                    var v47 int32 = -1
                    _ = v47
                } else {
                    if len(v42) > 0 {
                        v42[0] = fi41
                    }
                    if false {
                        v48 := GenS2{F0: 42}
                        _ = v48
                    }
                    v43 -= v43
                }
                v42 = append(v42, (^(fi39)))
                v43 -= v1
                var v49 uint64 = v44
                if (1 != 42) {
                    fmt.Println(fi39, v43)
                    var v50 bool = true
                    if v50 {
                        v3 = 4294967295
                        v42 = append(v42, fi41)
                        switch "\n\t" {
                        case "\n\t":
                            v51 := GenS5{F0: (v3 << ((uint(v49)) & 63)), F1: 2, F2: (v3 << (((v3 + v3)) & 63))}
                            _ = v51
                        case "a":
                            if len(v42) >= 2 {
                                v42 = v42[1:len(v42)]
                            }
                        case "":
                            _ = len(v42)
                        }
                        v3 -= v3
                        var v52 bool = (1e10 <= ((0.1 * 1.5) - 3.14159))
                        _ = v52
                    } else {
                        v53 := GenS6{F0: 211, F1: (v3 / ((65535) | 1)), F2: v2}
                        switch fi41 {
                        case 128:
                            var v54 uint16 = 88
                            _ = v54
                        case 9223372036854775807:
                            if (false || ((v3 ^ v3) > (v3 / (((v3 << ((380) & 63))) | 1)))) {
                                switch (fi41 >> (((^(v3))) & 63)) {
                                case -463:
                                    v55 := GenS4{F0: 1, F1: v44, F2: (v1 | uint8(v3))}
                                    _ = v55
                                case 2147483647:
                                    v2 = v2
                                }
                            } else {
                                v2 = v2
                            }
                        default:
                            v2 = (v2 | v2)
                        }
                        switch 32767 {
                        case 42:
                            v44 += (v44 >> ((225) & 63))
                        case 0:
                            v49 = 42
                        }
                        _ = v53
                    }
                    _ = v50
                } else {
                    v49 += 4294967295
                    var v56 uint64 = uint64(v3)
                    _ = v56
                }
                v57 := GenS7{F0: v43, F1: -1, F2: false}
                if true {
                    switch fi41 {
                    case 32768:
                        switch 32768 {
                        case 315:
                            var v58 uint8 = (^(v43))
                            _ = v58
                        case -9223372036854775808:
                            if (int64(v2) >= 1) {
                                if false {
                                    var v59 float32 = float32(1e10)
                                    fmt.Println(v3)
                                    _ = v59
                                }
                            } else {
                                v60 := GenS3{F0: (v3 ^ v3), F1: fi39}
                                _ = v60
                            }
                        case -1:
                            if false {
                                if (-2.25 != ((((((float64(1.5) * 0.0) - (float64((0.1 - (((float32(float64(1e10)) / (0.0 + (1.5 * float32(((0.1 / ((-2.25 * -2.25) / -2.25)) + 1e10))))) * -2.25) * ((3.14159 - 3.14159) * ((1.5 + (0.0 + (1e10 / 0.1))) - 3.14159))))) + 1.5)) / (0.1 / (3.14159 / 1.5))) * float64((0.1 * (((0.0 * -2.25) * float32(1e10)) / 3.14159)))) - 0.0) - ((1.5 + float64((1.5 / (3.14159 * -2.25)))) * 0.1))) {
                                    if ((false || (42 == 42)) && (true && (152 > v43))) {
                                        v61 := GenS8{F0: int32(v3)}
                                        if (42 >= -1) {
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                            v62 := GenS9{F0: (fi41 | fi41)}
                                            _ = v62
                                        }
                                        _ = v61
                                    } else {
                                        v63 := GenS10{F0: 0, F1: v2, F2: 1}
                                        _ = len(v42)
                                        var v64 float64 = 1e10
                                        v44 -= v44
                                        var v65 uint = 4294967295
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    }
                                    var v66 float32 = (1.5 * float32((0.0 - 1.5)))
                                    var v67 uint16 = 0
                                    _ = v66
                                    _ = v67
                                } else {
                                    switch 1 {
                                    case 42:
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                    case -561:
                                        var v68 bool = true
                                        _ = v68
                                    case 9223372036854775807:
                                        var v69 float64 = 1e10
                                        _ = v69
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            _ = len(v42)
                                        case "":
                                            v70 := GenS11{F0: v43}
                                            _ = v70
                                        case "a":
                                            if (!(!true)) {
                                                switch "\n\t" {
                                                case "hello":
                                                    var v71 int8 = -128
                                                    _ = v71
                                                default:
                                                    var v72 int64 = 902
                                                    _ = v72
                                                }
                                                var v73 bool = false
                                                v74 := GenS7{F0: v1, F1: 370, F2: (v3 != v3)}
                                                _ = v73
                                                _ = v74
                                            } else {
                                                var v75 uint8 = v43
                                                var v76 uint8 = 42
                                                v44 += v44
                                                var v77 uint16 = 255
                                                if true {
                                                    if true {
                                                        v77 += (v77 >> (((v3 % (((v3 >> ((v3) & 63))) | 1))) & 63))
                                                        var v78 float32 = (0.0 - 0.0)
                                                        v3 += (^(v3))
                                                        if (uint8(v3) != (v76 << ((v3) & 63))) {
                                                            var v79 float32 = float32(0.1)
                                                            v2 = (v2 - (v2 >> (((^(v3))) & 63)))
                                                            _ = v79
                                                        } else {
                                                            v80 := GenS1{F0: -1, F1: 750, F2: -9223372036854775808}
                                                            _ = v80
                                                        }
                                                        var v81 uint64 = (v44 * 42)
                                                        _ = v78
                                                        _ = v81
                                                    }
                                                }
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            }
                                            fmt.Println(v3)
                                        }
                                    }
                                    switch genFn0("foo bar", 576, fi41) {
                                    case 149:
                                        var v82 uint64 = uint64(v2)
                                        _ = v82
                                    default:
                                        var v83 float32 = 0.0
                                        _ = v83
                                    }
                                }
                                v3 = v3
                            } else {
                                if len(v42) > 0 {
                                    v42[0] = (^(fi39))
                                }
                                var v84 int = (fi39 * fi39)
                                _ = v84
                            }
                        default:
                            if true {
                                if false {
                                    v85 := GenS3{F0: 42, F1: fi39}
                                    _ = v85
                                } else {
                                    v42 = append(v42, (-(fi39)))
                                    v86 := GenS9{F0: (fi41 - (-(fi39)))}
                                    if true {
                                        var v87 int16 = 679
                                        if (!false) {
                                            v88 := GenS2{F0: 42}
                                            _ = v88
                                        } else {
                                            fmt.Println(v2)
                                            v89 := GenS12{F0: (!false), F1: (v87 & 32767), F2: 939}
                                            _ = v89
                                        }
                                        switch fi39 {
                                        case -9223372036854775808:
                                            v49 = v44
                                        case -1:
                                            v43 -= (v1 >> ((v3) & 63))
                                        case 0:
                                            if len(v42) >= 2 {
                                                v42 = v42[1:len(v42)]
                                            }
                                        default:
                                            v2 -= v2
                                        }
                                        _ = v87
                                    }
                                    _ = v86
                                }
                                var v90 uint = (v3 & 65535)
                                fmt.Println(fi41, v3)
                                _ = v90
                            }
                        }
                    case 2147483647:
                        v3 = 517
                    default:
                        var v91 string = "a"
                        _ = v91
                    }
                    var v92 float32 = 1.5
                    _ = v92
                } else {
                    var v93 float32 = ((1e10 * 0.1) + ((1.5 * (0.1 - 0.0)) * 0.0))
                    _ = v93
                }
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
                _ = v49
                _ = v57
            }
            if (-1 != -1) {
                fmt.Println(fi39)
            } else {
                v94 := []int{}
                v2 += v2
                v3 = (v3 % (((v3 - (^(v3)))) | 1))
                _ = v94
            }
            _ = fi39
        }
        var v95 float32 = 0.0
        var v96 uint32 = v2
        _ = v95
        _ = v96
    }
    var v97 uint16 = 255
    v98 := []int{42}
    v3 = 644
    v97 = 863
    switch "foo bar" {
    case "hello":
        v1 += (^(v1))
    }
    var v99 int = 681
    v100 := GenS5{F0: v3, F1: 0, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v97:", v97)
    fmt.Println("slicelen v98:", len(v98), v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v97
    _ = v98
    _ = v99
}
