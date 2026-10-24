package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 uint8; F1 int64; F2 int8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 int8) uint {
    switch "\n\t" {
    case "hello":
        v4 := []int{32768, 971, 2147483647, 2147483647}
        _ = v4
    case "\n\t":
        switch "a" {
        case "a":
            v5 := []int{128}
            _ = v5
        case "hello":
            v6 := []int{}
            _ = v6
        }
    case "a":
        if true {
            v7 := []int{1, 42, 921}
            if true {
                p0 += (^(p0))
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 float32 = (((0.1 * (0.0 / 1.0)) / 1.0) * 0.0)
                    _ = fi8
                    _ = v9
                }
            } else {
                v10 := GenS1{F0: 217}
                var v11 int16 = -7
                _ = v10
                _ = v11
            }
            _ = v7
        } else {
            if (!(97 <= 255)) {
                var v12 string = "hello"
                v12 = v12
                var v13 int32 = -778
                _ = v12
                _ = v13
            } else {
                v14 := GenS2{F0: 42, F1: -2147483648, F2: -475}
                _ = v14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 float64 = float64(0.1)
                    if true {
                        if (0.0 != v17) {
                            v18 := GenS3{F0: (v17 > (((0.1 / 1.0) + (((0.1 / float64(1e10)) / v17) + v17)) * (((v17 + v17) / float64(((0.0 + 1.5) * ((0.0 * 0.0) + 3.14159)))) + 0.1)))}
                            _ = v18
                        }
                    }
                    v19 := []int{4294967295, (fi15 ^ (fi16 | 128)), int(p1)}
                    _ = fi16
                    _ = v17
                    _ = v19
                }
                if (p1 == p1) {
                    v20 := GenS4{F0: (fi15 + (fi15 >> ((42) & 63)))}
                    fmt.Println(fi15)
                    _ = v20
                } else {
                    if false {
                        p1 = (p1 + (p1 ^ int8(fi15)))
                        fmt.Println(p1, fi15, fi15)
                    } else {
                        v21 := []int{2147483647}
                        for fi22 := 0; fi22 < 2; fi22++ {
                            var v23 int64 = 42
                            var v24 uint16 = 255
                            var v25 int = (-(fi22))
                            fmt.Println(p0, fi22)
                            var v26 int64 = v23
                            var v27 float64 = 1e10
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = v21
                    }
                }
                v28 := []int{fi15}
                _ = fi15
                _ = v28
            }
            p0 += (p0 | 631)
        }
    default:
        v29 := GenS5{F0: (^(p1))}
        _ = v29
    }
    var v30 float32 = 0.1
    v31 := []int{1, 1}
    v32 := GenS3{F0: true}
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    _ = v32
    return 981
}
func genFn1(p0 string) uint32 {
    v33 := []int{-213, 0, 4294967295, 128}
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            p0 = ("hello" + "\n\t")
            _ = fi35
        }
        _ = fi34
    }
    switch p0 {
    case "foo bar":
        switch "foo bar" {
        case "":
            var v36 float64 = (1e10 - float64(1.5))
            _ = v36
        case "\n\t":
            var v37 int8 = 0
            _ = v37
        default:
            var v38 int64 = -933
            _ = v38
        }
    case "a":
        fmt.Println(p0)
    }
    switch "\n\t" {
    case "hello":
        p0 = p0
    case "":
        var v39 uint8 = 89
        _ = v39
    case "foo bar":
        var v40 uint64 = 18446744073709551615
        _ = v40
    }
    _ = p0
    _ = v33
    return 739
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 0.0
    v3 := []int{}
    var v41 uint = 410
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v42 := GenS6{F0: 0}
    v43 := GenS7{F0: 42, F1: 1}
    var v44 uint = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v45 float32 = 1e10
        v45 -= (v2 / -2.25)
        _ = v45
    } else {
        v46 := GenS8{F0: (v44 * genFn0(42, -1)), F1: (true || true)}
        var v47 uint64 = 1
        switch "" {
        case "\n\t":
            fmt.Println(v1, v44, v2)
        case "foo bar":
            v48 := GenS9{F0: 1}
            _ = v48
        default:
            v49 := GenS10{F0: 1}
            _ = v49
        }
        _ = v46
        _ = v47
    }
    v44 = v44
    v1 = v1
    switch "" {
    case "a":
        for fi50 := 0; fi50 < 9; fi50++ {
            for fi51 := 0; fi51 < 9; fi51++ {
                v3 = append(v3, fi50)
                fmt.Println(v41, v41, fi50)
                _ = fi51
            }
            switch fi50 {
            case 2147483647:
                var v52 bool = (-1 < 763)
                _ = v52
            case 4294967295:
                switch (("foo bar" + "") + "a") {
                case "foo bar":
                    if (("hello" + "\n\t") != (("" + "\n\t") + "a")) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (3.14159 == 1e10) {
                            if (true || (-128 >= -92)) {
                                v3 = append(v3, fi50)
                            } else {
                                var v53 float64 = (1e10 - 0.1)
                                var v54 uint16 = 255
                                if len(v3) > 0 {
                                    v3[0] = fi50
                                }
                                _ = v53
                                _ = v54
                            }
                        } else {
                            switch fi50 {
                            case -628:
                                v1 += (v1 % (((^(v1))) | 1))
                            case -1:
                                v3 = append(v3, fi50)
                            }
                        }
                    } else {
                        var v55 int16 = int16(fi50)
                        for fi56 := 0; fi56 < 9; fi56++ {
                            if (100 != 42) {
                                var v57 uint64 = 18446744073709551615
                                _ = len(v3)
                                var v58 bool = false
                                _ = v57
                                _ = v58
                            } else {
                                if (!(v2 < 0.0)) {
                                    var v59 uint8 = 196
                                    _ = v59
                                } else {
                                    fmt.Println(v41)
                                    var v60 int32 = 42
                                    var v61 uint16 = 131
                                    _ = v60
                                    _ = v61
                                }
                                var v62 uint32 = (v1 & v1)
                                _ = v62
                            }
                            v41 -= genFn0(4294967295, 1)
                            _ = fi56
                        }
                        switch "" {
                        case "hello":
                            v1 += 65535
                        case "":
                            switch (fi50 & 1) {
                            case 127:
                                if (uint64(v41) == 0) {
                                    v63 := GenS3{F0: (!((255 >= 176) || (841 != 658)))}
                                    v1 = v1
                                    _ = v63
                                } else {
                                    var v64 int8 = 0
                                    v44 = 255
                                    _ = v64
                                }
                            case 353:
                                v65 := GenS11{F0: 329, F1: -2147483648}
                                _ = v65
                            case -9223372036854775808:
                                if true {
                                    for fi66 := 0; fi66 < 1; fi66++ {
                                        v2 += ((v2 * (1.5 / (((float32(-2.25) / (((((v2 + (v2 / 1.5)) + v2) - 0.1) / (float32(0.0) - v2)) - ((float32((float64((v2 + (v2 * float32(0.0)))) * -2.25)) + ((((v2 + v2) + v2) - v2) * (v2 - float32(0.1)))) * v2))) * v2) - (v2 * -2.25)))) + float32((1.5 - float64(v2))))
                                        _ = fi66
                                    }
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        v55 = v55
                                        if (!false) {
                                            v44 = v44
                                        } else {
                                            _ = len(v3)
                                            if ((v1 >> (((^(v44))) & 63)) < genFn1("\n\t")) {
                                                if true {
                                                    switch fi50 {
                                                    case -2147483648:
                                                        if len(v3) > 0 {
                                                            v3[0] = fi50
                                                        }
                                                    case -9223372036854775808:
                                                        var v68 int8 = -128
                                                        _ = v68
                                                    case 4294967295:
                                                        if true {
                                                            var v69 int32 = 1
                                                            var v70 int32 = v69
                                                            var v71 uint32 = genFn1("")
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            var v72 float32 = v2
                                                            if (127 > 0) {
                                                                if (-2147483648 != 2147483647) {
                                                                    var v73 int = (fi67 & (-(fi50)))
                                                                    v44 = v44
                                                                    v2 -= v72
                                                                    var v74 int64 = 9223372036854775807
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                            } else {
                                                                if true {
                                                                    var v75 float32 = (0.0 * v2)
                                                                    _ = len(v3)
                                                                    switch fi50 {
                                                                    case 1:
                                                                        v76 := GenS12{F0: 128, F1: -530, F2: -1}
                                                                        _ = v76
                                                                    case 0:
                                                                        switch fi67 {
                                                                        case -1:
                                                                            if true {
                                                                                var v77 int = fi50
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (fi67 * (v77 ^ fi50))
                                                                                }
                                                                                v78 := GenS11{F0: 1, F1: -420}
                                                                                _ = v77
                                                                                _ = v78
                                                                            }
                                                                        case 128:
                                                                            if (false && false) {
                                                                                if false {
                                                                                    var v79 uint = v44
                                                                                    switch ("hello" + "hello") {
                                                                                    case "":
                                                                                        var v80 float64 = 1.5
                                                                                        _ = v80
                                                                                    case "a":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 259
                                                                                        }
                                                                                    default:
                                                                                        if (!false) {
                                                                                            switch fi67 {
                                                                                            case -124:
                                                                                                var v81 float64 = 3.14159
                                                                                                _ = v81
                                                                                            }
                                                                                            var v82 int32 = 0
                                                                                            if (!(42 >= 2)) {
                                                                                                var v83 uint8 = 42
                                                                                                _ = v83
                                                                                            } else {
                                                                                                switch 9223372036854775807 {
                                                                                                case 127:
                                                                                                    switch fi50 {
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v41, v1, v79)
                                                                                                    default:
                                                                                                        v75 -= (v2 * (0.1 - v72))
                                                                                                    }
                                                                                                default:
                                                                                                    var v84 uint = (v44 << (((v44 * (v79 / ((genFn0(0, 127)) | 1)))) & 63))
                                                                                                    _ = v84
                                                                                                }
                                                                                                var v85 float64 = ((3.14159 - 0.0) * 0.1)
                                                                                                var v86 int = (fi67 & 32767)
                                                                                                v1 = v1
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    v87 := GenS3{F0: (v82 != -2147483648)}
                                                                                                    _ = v87
                                                                                                default:
                                                                                                    v72 = v75
                                                                                                }
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                            }
                                                                                            _ = v82
                                                                                        } else {
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = fi50
                                                                                            }
                                                                                            v55 = 1000
                                                                                        }
                                                                                    }
                                                                                    var v88 uint16 = uint16(v1)
                                                                                    v3 = append(v3, fi50)
                                                                                    switch int(v1) {
                                                                                    case 32768:
                                                                                        if (!(false || true)) {
                                                                                            v44 = (v41 - 65535)
                                                                                        } else {
                                                                                            _ = len(v3)
                                                                                        }
                                                                                    }
                                                                                    _ = v79
                                                                                    _ = v88
                                                                                }
                                                                            } else {
                                                                                var v89 uint16 = 385
                                                                                switch 114 {
                                                                                case 128:
                                                                                    v41 = (v44 >> (((v44 * (^(v41)))) & 63))
                                                                                }
                                                                                v90 := GenS1{F0: v55}
                                                                                v1 = (v1 + 688)
                                                                                if (false && false) {
                                                                                    v75 -= float32(1e10)
                                                                                    v41 = genFn0(663, -1)
                                                                                } else {
                                                                                    fmt.Println(v2, v41, v44)
                                                                                }
                                                                                _ = v89
                                                                                _ = v90
                                                                            }
                                                                        default:
                                                                            v91 := GenS1{F0: (-(v55))}
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    _ = v75
                                                                }
                                                                fmt.Println(v1, v55)
                                                                var v92 uint16 = 65535
                                                                v2 = v2
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v92
                                                            }
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                                fmt.Println(v1, v2)
                                            } else {
                                                v93 := GenS13{F0: ""}
                                                _ = v93
                                            }
                                        }
                                        v55 = (v55 ^ 0)
                                        _ = fi67
                                    }
                                } else {
                                    var v94 float32 = ((1e10 + float32(3.14159)) + v2)
                                    if (((3.14159 - v2) < (1e10 - (v2 / -2.25))) && ((float32(3.14159) + 1e10) != v2)) {
                                        v94 -= 1e10
                                    } else {
                                        for fi95 := 0; fi95 < 10; fi95++ {
                                            var v96 int32 = 151
                                            _ = len(v3)
                                            _ = fi95
                                            _ = v96
                                        }
                                        for fi97 := 0; fi97 < 6; fi97++ {
                                            _ = len(v3)
                                            v3 = append(v3, int(v1))
                                            v2 = v94
                                            _ = fi97
                                        }
                                        var v98 uint64 = 18446744073709551615
                                        _ = v98
                                    }
                                    fmt.Println(v1, v1)
                                    _ = v94
                                }
                            }
                        }
                        _ = v55
                    }
                case "a":
                    v99 := GenS7{F0: 42, F1: -1}
                    _ = v99
                }
            case 9223372036854775807:
                if false {
                    v100 := GenS14{F0: (v41 ^ 376), F1: 65535, F2: (v1 % ((v1) | 1))}
                    _ = v100
                } else {
                    var v101 float64 = -2.25
                    for fi102 := 0; fi102 < 3; fi102++ {
                        v2 = ((v2 / v2) + (1e10 * v2))
                        var v103 int32 = 0
                        if len(v3) > 0 {
                            v3[0] = (fi50 / ((fi102) | 1))
                        }
                        _ = fi102
                        _ = v103
                    }
                    _ = v101
                }
            }
            var v104 float32 = -2.25
            _ = fi50
            _ = v104
        }
    case "":
        v3 = append(v3, 1)
    }
    v44 -= 855
    v105 := GenS10{F0: 117}
    v106 := GenS11{F0: 42, F1: -1}
    v1 = (v1 << (((v41 % ((v44) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
