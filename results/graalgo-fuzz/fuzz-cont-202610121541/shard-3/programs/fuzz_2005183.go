package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 uint; F1 uint; F2 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 int64; F2 int }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint64) string {
    var v6 uint64 = p1
    for fi7 := 0; fi7 < 7; fi7++ {
        if (103 >= 255) {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := []int{(fi7 | (-(fi8)))}
                var v10 uint64 = v6
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(p0, p0)
            v11 := []int{0, fi7, fi7, fi7}
            _ = v11
        } else {
            if (!((p0 | p0) != (-(p0)))) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    switch "" {
                    case "a":
                        switch ("foo bar" + "hello") {
                        case "hello":
                            switch fi12 {
                            case 32767:
                                var v13 uint16 = 0
                                _ = v13
                            case 20:
                                p1 += 18446744073709551615
                            case -9223372036854775808:
                                v14 := GenS1{F0: -307, F1: 2147483648}
                                _ = v14
                            }
                        case "":
                            if false {
                                p0 = 970
                                if (true && (984 > (fi12 % (((fi7 + fi7)) | 1)))) {
                                    var v15 uint32 = 50
                                    fmt.Println(fi7, p0, v6)
                                    var v16 uint32 = (v15 | 65535)
                                    _ = v15
                                    _ = v16
                                }
                            }
                        }
                    case "hello":
                        if (false || (!(0 <= 1))) {
                            v17 := []int{fi12, -779, (fi12 / (((fi7 | fi7)) | 1)), fi12}
                            _ = v17
                        } else {
                            v18 := []int{127, 127}
                            v19 := GenS1{F0: 1, F1: 187}
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        p0 = p0
                    }
                    var v20 int32 = -15
                    v21 := []int{(fi12 / ((128) | 1))}
                    _ = fi12
                    _ = v20
                    _ = v21
                }
                fmt.Println(p0, p1)
                var v22 int16 = -491
                _ = v22
            }
            v23 := []int{(fi7 % (((fi7 << ((2) & 63))) | 1)), (fi7 / (((fi7 / (((-(fi7))) | 1))) | 1)), 0, fi7}
            _ = v23
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}
func genFn1(p0 uint, p1 uint8) uint32 {
    v24 := GenS1{F0: 0, F1: 2147483647}
    for fi25 := 0; fi25 < 14; fi25++ {
        switch fi25 {
        case -1:
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := []int{fi26}
                v28 := GenS2{F0: 1000}
                v29 := GenS1{F0: 1, F1: 992}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
        case -576:
            switch ("foo bar" + "hello") {
            case "":
                p0 -= p0
            case "foo bar":
                p0 += p0
            default:
                var v30 int16 = 42
                _ = v30
            }
        }
        _ = fi25
    }
    p1 = (p1 >> ((p0) & 63))
    _ = p0
    _ = p1
    _ = v24
    return 2147483648
}

func main() {
    v1 := []int{}
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 1
    var v4 int64 = 843
    var v5 float32 = (((1e10 + (1e10 + 0.0)) / ((float32(-2.25) + 3.14159) / 1e10)) * 1e10)
    var v31 bool = ((^(v2)) > (v2 * (v2 << ((18446744073709551615) & 63))))
    var v32 uint16 = 65535
    for fi33 := 0; fi33 < 14; fi33++ {
        if len(v1) > 0 {
            v1[0] = (fi33 << ((4294967295) & 63))
        }
        _ = fi33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        switch (("a" + "a") + "\n\t") {
        case "a":
            var v35 float32 = (0.0 / (float32(((1.5 - -2.25) * 3.14159)) * (((v5 + (float32(0.0) * v5)) + v5) + v5)))
            _ = v35
        default:
            if ((v2 - (^(v2))) > (v2 + (^(v2)))) {
                fmt.Println(v4)
                if false {
                    fmt.Println(v3, v31, v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v36 := GenS1{F0: (v4 / ((v4) | 1)), F1: 2}
                    _ = v36
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        switch "hello" {
        case "\n\t":
            var v37 uint = 0
            _ = v37
        case "foo bar":
            v4 += (v4 << ((453) & 63))
        default:
            switch "foo bar" {
            case "a":
                v1 = append(v1, (fi34 >> ((255) & 63)))
            case "\n\t":
                v2 -= v2
            case "hello":
                var v38 uint16 = (v32 | (v32 | v32))
                _ = v38
            }
        }
        _ = fi34
    }
    v32 += v32
    for fi39 := 0; fi39 < 1; fi39++ {
        fmt.Println(fi39, v32)
        fmt.Println(fi39, v2, v2)
        if (885 > (fi39 / (((fi39 ^ (fi39 | fi39))) | 1))) {
            var v40 bool = v31
            v40 = false
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v40
        } else {
            var v41 uint8 = 33
            v42 := GenS1{F0: (v4 - (v4 | int64(v3))), F1: 268}
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        switch "hello" {
        case "hello":
            switch -365 {
            case -1:
                var v44 string = ""
                _ = v44
            case 178:
                if v31 {
                    for fi45 := 0; fi45 < 4; fi45++ {
                        var v46 float32 = 1e10
                        var v47 uint16 = 0
                        v48 := GenS3{F0: -128, F1: 42, F2: v32}
                        _ = fi45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    var v49 uint = 255
                    for fi50 := 0; fi50 < 20; fi50++ {
                        var v51 int16 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "foo bar" {
                        case "hello":
                            v49 = v49
                        }
                        var v52 int = (^(fi43))
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    var v53 uint = 42
                    var v54 uint8 = 1
                    _ = v53
                    _ = v54
                }
            }
        case "a":
            v5 = v5
        case "foo bar":
            if ((v3 | (-(v3))) > (^(v3))) {
                if (!true) {
                    if (0.0 < -2.25) {
                        v55 := GenS1{F0: (v4 ^ (v4 - (v4 >> ((999) & 63)))), F1: 417}
                        for fi56 := 0; fi56 < 15; fi56++ {
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = (fi43 - (fi56 >> ((42) & 63)))
                                }
                                if true {
                                    v57 := GenS4{F0: (v2 ^ (v2 * (v2 & (v2 | 992)))), F1: 682, F2: 711}
                                    if len(v1) > 0 {
                                        v1[0] = (fi43 & 0)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = -247
                                    }
                                    if (2 > 127) {
                                        var v58 int64 = (v4 << ((860) & 63))
                                        v1 = append(v1, (fi43 & fi56))
                                        _ = v58
                                    } else {
                                        var v59 int16 = v3
                                        v5 -= v5
                                        v60 := GenS2{F0: v59}
                                        if false {
                                            if v31 {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v61 := GenS5{F0: (v2 << ((900) & 63)), F1: 42}
                                            v62 := GenS6{F0: (v32 / (((v32 << ((42) & 63))) | 1)), F1: 255, F2: uint16(v2)}
                                            _ = v61
                                            _ = v62
                                        }
                                        v63 := GenS5{F0: v2, F1: 201}
                                        _ = v59
                                        _ = v60
                                        _ = v63
                                    }
                                    v64 := GenS7{F0: 0, F1: -1}
                                    v5 = v5
                                    _ = v57
                                    _ = v64
                                } else {
                                    if true {
                                        if (!v31) {
                                            if len(v1) > 0 {
                                                v1[0] = 261
                                            }
                                        }
                                        v65 := GenS2{F0: (v3 / ((v3) | 1))}
                                        _ = v65
                                    }
                                    var v66 int8 = 111
                                    var v67 string = "hello"
                                    v2 = (v2 % (((v2 & v2)) | 1))
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                v68 := GenS8{F0: 95, F1: 2147483647}
                                switch 5 {
                                case 128:
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(fi56)
                                    case "":
                                        var v69 uint8 = 2
                                        _ = v69
                                    }
                                default:
                                    switch fi56 {
                                    case 42:
                                        switch fi43 {
                                        case -2147483648:
                                            var v70 float64 = (0.1 + 1e10)
                                            _ = v70
                                        case 4294967295:
                                            var v71 int8 = int8(fi43)
                                            _ = v71
                                        default:
                                            var v72 uint64 = 2
                                            _ = v72
                                        }
                                    case -344:
                                        fmt.Println(v32, v2, v3)
                                    default:
                                        if v31 {
                                            v73 := GenS9{F0: 18446744073709551615, F1: 0, F2: 0}
                                            var v74 int32 = 578
                                            fmt.Println(v5, v31)
                                            _ = v73
                                            _ = v74
                                        } else {
                                            var v75 int64 = -9223372036854775808
                                            _ = v75
                                        }
                                    }
                                }
                                _ = v68
                            }
                            fmt.Println(v5)
                            if v31 {
                                switch 128 {
                                case 32768:
                                    v76 := GenS6{F0: (v32 * (v32 / ((347) | 1))), F1: 139, F2: (^(v32))}
                                    v77 := GenS6{F0: 1, F1: 255, F2: v32}
                                    _ = v76
                                    _ = v77
                                }
                                switch fi56 {
                                case 42:
                                    if true {
                                        var v78 uint16 = 687
                                        v31 = false
                                        if (v31 || (-2.25 >= (0.0 * 1e10))) {
                                            v79 := GenS6{F0: 928, F1: 103, F2: (^(v32))}
                                            if (2 >= 0) {
                                                if v31 {
                                                    var v80 string = (("" + ("hello" + "foo bar")) + "foo bar")
                                                    _ = v80
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            } else {
                                                var v81 int8 = 127
                                                _ = v81
                                            }
                                            var v82 int8 = int8(v3)
                                            v82 = (v82 ^ (-(v82)))
                                            _ = v79
                                            _ = v82
                                        }
                                        _ = v78
                                    }
                                case -543:
                                    v3 -= (v3 & -563)
                                }
                                if v31 {
                                    var v83 int16 = (v3 + v3)
                                    v31 = v31
                                    switch (fi43 >> ((1) & 63)) {
                                    case 32768:
                                        _ = len(v1)
                                    case -1:
                                        v2 += 18446744073709551615
                                    case 128:
                                        fmt.Println(v5, fi56, v5)
                                        fmt.Println(v3, v5)
                                        v84 := GenS10{F0: 42}
                                        _ = v84
                                    default:
                                        fmt.Println(fi43)
                                    }
                                    _ = v83
                                }
                            } else {
                                if false {
                                    switch fi56 {
                                    case 0:
                                        v2 += (v2 >> ((1) & 63))
                                    case 32768:
                                        if v31 {
                                            v2 = uint64(v32)
                                            if v31 {
                                                v2 -= 896
                                                if false {
                                                    v85 := GenS11{F0: (v32 % ((v32) | 1)), F1: (v4 >> ((810) & 63)), F2: fi43}
                                                    _ = v85
                                                }
                                                if (false || (fi43 >= int(v4))) {
                                                    var v86 float64 = (3.14159 / 1.0)
                                                    if v31 {
                                                        v87 := GenS12{F0: 1}
                                                        _ = v87
                                                    }
                                                    _ = v86
                                                }
                                            }
                                            v4 -= (v4 & v4)
                                        } else {
                                            var v88 int16 = v3
                                            _ = v88
                                        }
                                    case -301:
                                        v5 += (0.1 / v5)
                                    }
                                    if false {
                                        fmt.Println(v31, v2, v31)
                                    } else {
                                        _ = len(v1)
                                        switch (fi56 & fi56) {
                                        case 0:
                                            var v89 int8 = 127
                                            _ = v89
                                        default:
                                            v5 -= ((0.1 + 0.1) * v5)
                                            v4 = 351
                                            v1 = append(v1, (fi56 * fi43))
                                        }
                                    }
                                } else {
                                    v32 = 0
                                    fmt.Println(fi43)
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v32 = 560
                            } else {
                                var v90 float64 = ((((float64((v5 / float32(1e10))) - ((1e10 - (0.1 / 1.0)) * 1e10)) * -2.25) * 1.5) / (3.14159 / -2.25))
                                if true {
                                    v90 += 1.5
                                    var v91 int32 = 936
                                    _ = v91
                                }
                                _ = v90
                            }
                            _ = fi56
                        }
                        for fi92 := 0; fi92 < 17; fi92++ {
                            var v93 int64 = v4
                            fmt.Println(v3, fi43)
                            _ = fi92
                            _ = v93
                        }
                        _ = v55
                    }
                } else {
                    fmt.Println(v5)
                    v1 = append(v1, 9223372036854775807)
                }
            }
        default:
            var v94 uint16 = 65535
            _ = v94
        }
        _ = fi43
    }
    v3 = v3
    fmt.Println(v31, v4)
    fmt.Println(v32, v4)
    if true {
        if (float32((-2.25 - ((0.1 - 3.14159) * 1.5))) != 1e10) {
            var v95 int = 32767
            v95 = (v95 - v95)
            _ = v95
        }
    } else {
        for fi96 := 0; fi96 < 11; fi96++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v97 uint16 = v32
            _ = fi96
            _ = v97
        }
        var v98 float64 = (3.14159 + 0.0)
        var v99 uint16 = v32
        _ = v98
        _ = v99
    }
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
