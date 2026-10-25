package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 string; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint; F1 uint; F2 int16 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint64; F1 int8 }
func genFn0(p0 int64) int64 {
    v5 := GenS1{F0: 580, F1: 966}
    var v6 uint8 = 2
    _ = p0
    _ = v5
    _ = v6
    return (p0 / (((p0 & (-(p0)))) | 1))
}
func genFn1(p0 uint16, p1 bool, p2 int16) uint8 {
    p2 += p2
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 38
}

func main() {
    v1 := []int{0, 1, 9223372036854775807}
    v2 := make(map[int]bool)
    var v3 int = 128
    var v4 uint = 4294967295
    if (false && false) {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v4, v3)
            v8 := GenS2{F0: 42}
            for fi9 := 0; fi9 < 14; fi9++ {
                v4 -= (^(v4))
                var v10 int64 = 42
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    } else {
        var v11 string = ("a" + ("a" + "a"))
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 string = ""
                switch v11 {
                case "a":
                    if true {
                        fmt.Println(fi12, v3, v11)
                    }
                }
                _ = fi13
                _ = v14
            }
            _ = len(v2)
            _ = fi12
        }
        _ = v11
    }
    var v15 int = (v3 * (-(v3)))
    var v16 float64 = 1e10
    fmt.Println(v15)
    fmt.Println(v3, v3, v16)
    _ = len(v2)
    v16 = 3.14159
    var v17 int64 = 0
    var v18 string = (("\n\t" + (("\n\t" + "hello") + "")) + "")
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int8 = int8(v17)
        var v21 int8 = 100
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint8 = 42
    var v23 uint = (v4 * v4)
    v24 := GenS3{F0: -479}
    switch (v15 | (v3 << ((2) & 63))) {
    case 111:
        if (42 == 0) {
            var v25 uint32 = 1
            _ = v25
        } else {
            for fi26 := 0; fi26 < 12; fi26++ {
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 int16 = -1
                    if false {
                        switch v18 {
                        case "\n\t":
                            if true {
                                v28 = (v28 >> ((0) & 63))
                            } else {
                                v29 := GenS2{F0: 0}
                                if (!(v16 >= v16)) {
                                    v3 -= fi26
                                } else {
                                    if (v16 > v16) {
                                        if true {
                                            if ((v18 + (((v18 + "") + "") + ((("" + v18) + "\n\t") + v18))) <= v18) {
                                                var v30 int32 = 767
                                                v15 += fi27
                                                _ = v30
                                            } else {
                                                v23 += v23
                                                _ = len(v2)
                                                switch fi27 {
                                                case 42:
                                                    switch fi26 {
                                                    case 32767:
                                                        if false {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v1 = append(v1, 127)
                                                            case "a":
                                                                if (v22 == (v22 | (v22 | (v22 + uint8(v23))))) {
                                                                    if false {
                                                                        var v31 uint64 = 18446744073709551615
                                                                        _ = v31
                                                                    }
                                                                }
                                                                fmt.Println(v23, fi26, v28)
                                                                _ = len(v2)
                                                            default:
                                                                _ = len(v2)
                                                                v16 += v16
                                                            }
                                                            v32 := GenS3{F0: -32768}
                                                            _ = v32
                                                        } else {
                                                            var v33 int64 = (v17 & (-(v17)))
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v1 = append(v1, 9223372036854775807)
                                                            _ = v33
                                                        }
                                                        if (v16 != float64((3.14159 * 1e10))) {
                                                            if (((v16 * -2.25) < (-2.25 - ((v16 * v16) - 1e10))) || ((v28 == 32767) && (false || ((!(1 >= (v17 >> ((v4) & 63)))) || false)))) {
                                                                if (!true) {
                                                                    v15 -= v15
                                                                    switch v18 {
                                                                    case "\n\t":
                                                                        v16 += 0.0
                                                                    case "a":
                                                                        if ((v18 + v18) != v18) {
                                                                            var v34 float64 = float64((float32(float64(1e10)) + 1.5))
                                                                            _ = v34
                                                                        } else {
                                                                            v17 = (v17 | v17)
                                                                        }
                                                                    default:
                                                                        fmt.Println(v18, v15, v28)
                                                                    }
                                                                } else {
                                                                    var v35 int16 = (v28 << (((v23 & (v23 - 65535))) & 63))
                                                                    _ = v35
                                                                }
                                                                fmt.Println(v15, v28)
                                                            }
                                                            switch (-(fi26)) {
                                                            case 395:
                                                                var v36 float64 = 0.1
                                                                _ = v36
                                                            case 2147483647:
                                                                fmt.Println(v23, fi27, fi26)
                                                            case -688:
                                                                v17 += genFn0(v17)
                                                            default:
                                                                var v37 int8 = 42
                                                                _ = v37
                                                            }
                                                        } else {
                                                            if false {
                                                                var v38 uint = (v23 + 65535)
                                                                var v39 int8 = 127
                                                                delete(v2, fi26)
                                                                fmt.Println(v16, v28)
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            switch (fi26 % (((v15 << (((v23 << ((v23) & 63))) & 63))) | 1)) {
                                                            case -876:
                                                                v4 = v23
                                                            case 1:
                                                                if (!true) {
                                                                    var v40 float32 = 0.0
                                                                    _ = v40
                                                                } else {
                                                                    fmt.Println(v22)
                                                                    if ((v23 == (v4 << ((594) & 63))) && ((v15 >> ((4294967295) & 63)) >= v15)) {
                                                                        if (2 > 0) {
                                                                            var v41 float32 = (0.0 - (1.5 - (3.14159 - 3.14159)))
                                                                            var v42 uint16 = 197
                                                                            _ = v41
                                                                            _ = v42
                                                                        }
                                                                        switch v15 {
                                                                        case 127:
                                                                            var v43 string = (v18 + (v18 + v18))
                                                                            _ = v43
                                                                        case 0:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        case 32768:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (fi26 * -552)
                                                                            }
                                                                        }
                                                                        v15 = (v3 % ((int(v4)) | 1))
                                                                    } else {
                                                                        fmt.Println(v3, v3)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            switch (v18 + v18) {
                                                            case "hello":
                                                                var v44 float32 = 0.1
                                                                _ = v44
                                                            case "foo bar":
                                                                var v45 int8 = 0
                                                                _ = v45
                                                            case "\n\t":
                                                                v46 := GenS4{F0: 18446744073709551615}
                                                                _ = v46
                                                            default:
                                                                v4 += v23
                                                            }
                                                        case "":
                                                            v47 := GenS3{F0: 546}
                                                            _ = v47
                                                        default:
                                                            if true {
                                                                if false {
                                                                    v16 = (0.1 * v16)
                                                                    v48 := GenS5{F0: v28, F1: "\n\t", F2: genFn0(v17)}
                                                                    _ = v48
                                                                } else {
                                                                    var v49 int = fi27
                                                                    v18 = v18
                                                                    var v50 string = ("hello" + (v18 + v18))
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                if (((0.0 * 3.14159) / 3.14159) > 0.0) {
                                                                    fmt.Println(fi26, v4)
                                                                    fmt.Println(v22, v28, v16)
                                                                } else {
                                                                    v4 = 0
                                                                }
                                                                var v51 int32 = 653
                                                                v16 += (0.0 - float64(-2.25))
                                                                _ = v51
                                                            }
                                                        }
                                                    case 2147483647:
                                                        v23 += 18446744073709551615
                                                    case -9223372036854775808:
                                                        v1 = append(v1, fi26)
                                                    default:
                                                        if (true || true) {
                                                            switch (v18 + ((((v18 + v18) + ("" + v18)) + ("\n\t" + v18)) + "foo bar")) {
                                                            case "\n\t":
                                                                fmt.Println(v3, v4)
                                                            default:
                                                                v28 = (v28 >> ((v4) & 63))
                                                            }
                                                            v16 = (v16 / v16)
                                                        }
                                                    }
                                                }
                                            }
                                            if ((false && (1 != 0)) && true) {
                                                switch "\n\t" {
                                                case "":
                                                    v52 := GenS6{F0: 180}
                                                    _ = v52
                                                case "\n\t":
                                                    if (!true) {
                                                        if len(v1) > 0 {
                                                            v1[0] = v3
                                                        }
                                                        v2[(fi27 + -2147483648)] = false
                                                        v23 = (^(v23))
                                                        fmt.Println(v23)
                                                        v2[v3] = (0 < 18446744073709551615)
                                                    } else {
                                                        var v53 bool = (v4 != v23)
                                                        var v54 uint16 = 1
                                                        var v55 string = "foo bar"
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                case "foo bar":
                                                    v56 := GenS7{F0: genFn1(0, false, int16(v22)), F1: (v23 - v23)}
                                                    v57 := GenS8{F0: v4}
                                                    switch v18 {
                                                    case "foo bar":
                                                        var v58 float64 = float64((1e10 - 1.5))
                                                        _ = v58
                                                    case "\n\t":
                                                        switch (v3 - int(v22)) {
                                                        case 1:
                                                            v18 = (v18 + "foo bar")
                                                        }
                                                    }
                                                    if false {
                                                        if false {
                                                            v15 += -2147483648
                                                        } else {
                                                            fmt.Println(v28, v23, v23)
                                                        }
                                                        fmt.Println(v16, v16)
                                                    }
                                                    v59 := GenS9{F0: 0}
                                                    _ = v56
                                                    _ = v57
                                                    _ = v59
                                                }
                                            }
                                            var v60 uint32 = 42
                                            switch (fi27 + 1) {
                                            case 32768:
                                                v3 = -406
                                            case -9223372036854775808:
                                                if (1 != 4294967295) {
                                                    var v61 float32 = 3.14159
                                                    _ = v61
                                                } else {
                                                    var v62 bool = (!(v18 > v18))
                                                    _ = v62
                                                }
                                            case 2147483647:
                                                fmt.Println(v28)
                                            default:
                                                v15 = 128
                                            }
                                            _ = v60
                                        }
                                        v16 -= 3.14159
                                        if false {
                                            switch (v18 + "foo bar") {
                                            case "":
                                                if true {
                                                    switch (v15 & (fi27 / (((v15 | (fi26 ^ 141))) | 1))) {
                                                    case 32768:
                                                        var v63 string = v18
                                                        _ = v63
                                                    }
                                                    if true {
                                                        v64 := GenS1{F0: 836, F1: 2}
                                                        _ = v64
                                                    }
                                                    v18 = "hello"
                                                } else {
                                                    var v65 float32 = (-2.25 * 1e10)
                                                    var v66 int32 = 1
                                                    _ = v65
                                                    _ = v66
                                                }
                                            default:
                                                v3 = -652
                                            }
                                        }
                                        v67 := GenS3{F0: v28}
                                        var v68 float32 = (1e10 + -2.25)
                                        _ = len(v1)
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if (-450 == 302) {
                                            v28 = 42
                                        } else {
                                            switch ("\n\t" + v18) {
                                            case "a":
                                                v69 := GenS1{F0: v23, F1: 18446744073709551615}
                                                _ = v69
                                            default:
                                                var v70 uint32 = uint32(v22)
                                                _ = v70
                                            }
                                            var v71 bool = false
                                            _ = v71
                                        }
                                        v72 := GenS4{F0: 693}
                                        _ = v72
                                    }
                                    v73 := v2[(v15 ^ (^(v3)))]
                                    _ = v73
                                }
                                fmt.Println(v18, v3, v18)
                                _ = v29
                            }
                        case "foo bar":
                            _ = len(v1)
                        case "":
                            if (v3 == v3) {
                                var v74 uint16 = 32767
                                switch (v15 << (((v23 / ((uint(v74)) | 1))) & 63)) {
                                case 4294967295:
                                    var v75 int64 = v17
                                    _ = v75
                                case 2147483647:
                                    v76 := GenS10{F0: -9223372036854775808}
                                    _ = v76
                                }
                                if false {
                                    if false {
                                        var v77 string = "a"
                                        if true {
                                            switch "foo bar" {
                                            case "hello":
                                                v78 := GenS4{F0: 0}
                                                _ = v78
                                            case "foo bar":
                                                v79 := v2[(fi27 >> (((^(v23))) & 63))]
                                                _ = v79
                                            case "a":
                                                v16 -= v16
                                            }
                                        } else {
                                            v17 -= 42
                                            var v80 uint16 = (v74 | 1)
                                            v74 += 612
                                            _ = v80
                                        }
                                        _ = v77
                                    } else {
                                        v81 := GenS9{F0: 1}
                                        _ = v81
                                    }
                                    var v82 float64 = v16
                                    v83 := GenS11{F0: (^(v4)), F1: v4, F2: v28}
                                    fmt.Println(v23, fi26)
                                    var v84 uint32 = 65535
                                    if (false && true) {
                                        v82 += v82
                                        v3 = fi27
                                    }
                                    fmt.Println(v4)
                                    switch int(v22) {
                                    case 42:
                                        switch v3 {
                                        case -2147483648:
                                            var v85 uint = 0
                                            if false {
                                                var v86 string = v18
                                                _ = v86
                                            }
                                            var v87 int = -2147483648
                                            _ = v85
                                            _ = v87
                                        case 9223372036854775807:
                                            var v88 float32 = 1.5
                                            _ = v88
                                        }
                                        v89 := GenS9{F0: -849}
                                        fmt.Println(v82)
                                        _ = v89
                                    default:
                                        var v90 uint64 = 4294967295
                                        _ = v90
                                    }
                                    v28 -= v28
                                    v17 = v17
                                    var v91 int64 = (-(v17))
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                    _ = v91
                                }
                                _ = v74
                            } else {
                                v3 = fi27
                                _ = len(v2)
                            }
                        default:
                            var v92 string = "a"
                            _ = v92
                        }
                        v4 += 343
                        fmt.Println(v15)
                    } else {
                        v17 = v17
                        if true {
                            v93 := GenS12{F0: -1}
                            if (0.1 != 1e10) {
                                var v94 uint64 = 4294967295
                                fmt.Println(v22)
                                v15 += 4294967295
                                _ = v94
                            }
                            v22 = v22
                            _ = v93
                        }
                        v15 -= -2147483648
                    }
                    _ = fi27
                    _ = v28
                }
                var v95 uint64 = uint64(v22)
                v96 := GenS13{F0: 630, F1: -61}
                _ = fi26
                _ = v95
                _ = v96
            }
        }
    case 0:
        var v97 int = (^(v15))
        _ = v97
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
