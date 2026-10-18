package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int8; F1 int16; F2 int32 }
type GenS7 struct { F0 int8; F1 int; F2 int64 }
type GenS8 struct { F0 uint32; F1 bool; F2 uint }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 uint16; F1 int16; F2 int16 }
type GenS11 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0() bool {
    var v5 uint16 = 32767
    v5 -= (v5 + 32767)
    v6 := []int{}
    for fi7 := 0; fi7 < 7; fi7++ {
        v5 -= (^(v5))
        v5 += 160
        v5 = 65535
        _ = fi7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "\n\t" {
        case "":
            v5 = (v5 / (((v5 % (((^(v5))) | 1))) | 1))
        case "hello":
            fmt.Println(v5)
        case "foo bar":
            v5 -= v5
        }
        fmt.Println(v5, fi8)
        v5 += v5
        _ = fi8
    }
    _ = v5
    _ = v6
    return ((302 <= -966) && false)
}
func genFn1() uint32 {
    var v9 uint64 = 0
    v9 = (^(v9))
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int = fi10
        v12 := GenS1{F0: 2, F1: (true && false)}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v9 += (v9 / ((4294967295) | 1))
        v14 := []int{32768, fi13, (fi13 * fi13)}
        if len(v14) > 0 {
            v14[0] = (fi13 / (((fi13 ^ (fi13 & (fi13 * 1)))) | 1))
        }
        _ = fi13
        _ = v14
    }
    if false {
        if true {
            for fi15 := 0; fi15 < 1; fi15++ {
                switch (fi15 ^ (fi15 * (fi15 / ((fi15) | 1)))) {
                case 4294967295:
                    var v16 uint = 56
                    _ = v16
                case 900:
                    var v17 bool = false
                    _ = v17
                }
                var v18 uint8 = 128
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := []int{fi19, -1, fi15, (fi19 * (fi15 - fi15))}
                    _ = len(v20)
                    if ((!(true || false)) && (0.1 <= (0.1 * 1.5))) {
                        fmt.Println(v18, fi19)
                        if ((!true) && true) {
                            var v21 bool = true
                            _ = v21
                        } else {
                            fmt.Println(fi15)
                        }
                        v22 := GenS2{F0: false, F1: 0, F2: (!true)}
                        _ = v22
                    } else {
                        var v23 bool = true
                        fmt.Println(v18)
                        v24 := GenS1{F0: 128, F1: (42 != 277)}
                        _ = v23
                        _ = v24
                    }
                    _ = fi19
                    _ = v20
                }
                _ = fi15
                _ = v18
            }
        }
        switch 2147483647 {
        case 32767:
            var v25 int32 = 86
            _ = v25
        }
        if (!((-2.25 >= 0.1) || true)) {
            var v26 uint32 = 994
            switch ("foo bar" + "") {
            case "hello":
                v9 += 0
            default:
                v9 = (v9 + 862)
            }
            _ = v26
        }
    } else {
        v27 := GenS2{F0: true, F1: 9223372036854775807, F2: (166 >= 1)}
        v28 := GenS3{F0: 42, F1: (v9 | (^(v9)))}
        var v29 float32 = -2.25
        _ = v27
        _ = v28
        _ = v29
    }
    _ = v9
    return 764
}

func main() {
    var v1 uint8 = 109
    var v2 int64 = 589
    var v3 uint = 65535
    var v4 uint16 = 32767
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            v2 -= (v2 >> (((v3 >> ((18446744073709551615) & 63))) & 63))
            fmt.Println(v4, fi30)
            _ = fi31
        }
        var v32 string = ""
        _ = fi30
        _ = v32
    }
    var v33 int = 102
    v34 := GenS3{F0: -1, F1: 0}
    if true {
        var v35 int64 = v2
        var v36 int = v33
        _ = v35
        _ = v36
    } else {
        if true {
            v1 += (v1 % ((v1) | 1))
            v37 := []int{(v33 ^ v33), (v33 >> ((v3) & 63))}
            _ = v37
        }
        v1 += v1
    }
    if (!true) {
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 uint16 = (v4 & (v4 - (^(v4))))
            for fi40 := 0; fi40 < 4; fi40++ {
                v41 := GenS4{F0: 324}
                _ = fi40
                _ = v41
            }
            v42 := []int{}
            _ = fi38
            _ = v39
            _ = v42
        }
        fmt.Println(v4)
    } else {
        var v43 int16 = -340
        if (v43 <= (v43 % (((v43 / ((int16(v1)) | 1))) | 1))) {
            v3 -= 65535
        } else {
            v1 -= v1
            v1 += (v1 * (v1 - 128))
            v44 := GenS2{F0: ((0.0 - (1.5 + 0.1)) != 0.0), F1: (v2 & -9223372036854775808), F2: ((v43 % ((int16(v4)) | 1)) < (v43 << ((0) & 63)))}
            _ = v44
        }
        v4 = v4
        _ = v43
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 string = ""
        for fi47 := 0; fi47 < 16; fi47++ {
            v48 := []int{}
            _ = fi47
            _ = v48
        }
        var v49 uint16 = (v4 ^ v4)
        _ = fi45
        _ = v46
        _ = v49
    }
    if (1 <= 2147483647) {
        if false {
            var v50 bool = (true || (42 != -2147483648))
            _ = v50
        }
    }
    v51 := []int{}
    for fi52 := 0; fi52 < 20; fi52++ {
        for fi53 := 0; fi53 < 19; fi53++ {
            v54 := GenS5{F0: (v4 / ((v4) | 1)), F1: 4294967295}
            fmt.Println(v1, v33, fi52)
            switch (v33 << ((906) & 63)) {
            case -1:
                switch 9223372036854775807 {
                case 829:
                    var v55 uint16 = v4
                    _ = v55
                case 376:
                    fmt.Println(fi53)
                    v56 := GenS3{F0: 40, F1: 1}
                    var v57 int32 = 0
                    _ = v56
                    _ = v57
                case 128:
                    if ((((1.5 - -2.25) * (3.14159 + -2.25)) - 1.5) <= float64(1.5)) {
                        if (!(1 >= 127)) {
                            fmt.Println(fi53, v1)
                        }
                        var v58 int16 = 0
                        _ = v58
                    }
                default:
                    fmt.Println(v2)
                }
            default:
                if false {
                    var v59 bool = false
                    v59 = v59
                    if len(v51) > 0 {
                        v51[0] = (fi52 * (fi52 / ((fi53) | 1)))
                    }
                    v60 := GenS6{F0: -113, F1: 799, F2: 2147483647}
                    v1 = v1
                    fmt.Println(v59, v33)
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                    switch (^(v33)) {
                    case 32768:
                        v59 = v59
                    case 1:
                        if true {
                            switch "hello" {
                            case "a":
                                var v61 int16 = 0
                                _ = v61
                            case "hello":
                                var v62 bool = false
                                _ = v62
                            case "\n\t":
                                var v63 uint = v3
                                _ = v63
                            }
                        } else {
                            if false {
                                var v64 int16 = 1000
                                switch (-(fi53)) {
                                case 0:
                                    if ((v4 ^ uint16(v3)) <= v4) {
                                        var v65 int64 = int64(v64)
                                        _ = v65
                                    }
                                }
                                v51 = append(v51, (v33 << (((v3 | v3)) & 63)))
                                _ = v64
                            } else {
                                if (0.0 == float32((((-2.25 + float64((3.14159 * (float32(1e10) / (3.14159 / 1e10))))) - 3.14159) * 1.5))) {
                                    v66 := GenS7{F0: 42, F1: 4294967295, F2: -214}
                                    v51 = append(v51, v33)
                                    _ = v66
                                } else {
                                    if len(v51) >= 2 {
                                        v51 = v51[1:len(v51)]
                                    }
                                    fmt.Println(v59, v3, v3)
                                    if true {
                                        switch (v33 + fi53) {
                                        case 0:
                                            switch "hello" {
                                            case "a":
                                                switch (fi53 - int(v4)) {
                                                case 1:
                                                    var v67 uint64 = 910
                                                    _ = v67
                                                case -1:
                                                    if (!(v59 || v59)) {
                                                        var v68 float32 = (((((float32(0.0) * 0.0) / 3.14159) + 1e10) - 1e10) * (-2.25 + -2.25))
                                                        v51 = append(v51, fi53)
                                                        fmt.Println(v4, v59)
                                                        _ = v68
                                                    } else {
                                                        v69 := GenS5{F0: v4, F1: 0}
                                                        _ = v69
                                                    }
                                                case 128:
                                                    if len(v51) > 0 {
                                                        v51[0] = (v33 * 42)
                                                    }
                                                    v1 += 2
                                                default:
                                                    v70 := GenS8{F0: genFn1(), F1: v59, F2: 42}
                                                    _ = v70
                                                }
                                            }
                                        }
                                        var v71 uint8 = 2
                                        _ = v71
                                    }
                                }
                            }
                            v4 -= v4
                        }
                    case 0:
                        if v59 {
                            v33 += (^(fi53))
                            if v59 {
                                var v72 int64 = (v2 - (v2 | (v2 << ((v3) & 63))))
                                var v73 int8 = 127
                                switch ("foo bar" + "hello") {
                                case "hello":
                                    if v59 {
                                        v51 = append(v51, v33)
                                    } else {
                                        var v74 bool = (42 != 2)
                                        if len(v51) > 0 {
                                            v51[0] = (fi52 / ((fi52) | 1))
                                        }
                                        _ = v74
                                    }
                                case "":
                                    v2 = (v72 * 281)
                                case "\n\t":
                                    if v59 {
                                        v33 = int(v72)
                                        var v75 float32 = 0.1
                                        _ = v75
                                    }
                                }
                                _ = v72
                                _ = v73
                            }
                            if (((0.1 + 0.1) <= 1.5) && v59) {
                                fmt.Println(v1, v2, v1)
                                fmt.Println(v59, v2, v1)
                            } else {
                                if ((v4 | v4) != v4) {
                                    var v76 int = v33
                                    _ = v76
                                }
                            }
                        } else {
                            fmt.Println(v2)
                        }
                        if v59 {
                            var v77 int16 = -1
                            fmt.Println(v2, v77, v1)
                            _ = v77
                        }
                        switch "a" {
                        case "foo bar":
                            var v78 int32 = 0
                            _ = v78
                        case "hello":
                            v59 = true
                            v79 := GenS9{F0: 0, F1: 39}
                            _ = v79
                        default:
                            v4 = v4
                        }
                    default:
                        var v80 uint64 = 1
                        _ = v80
                    }
                    _ = v59
                    _ = v60
                }
            }
            _ = fi53
            _ = v54
        }
        for fi81 := 0; fi81 < 5; fi81++ {
            switch "foo bar" {
            case "a":
                var v82 string = "a"
                _ = v82
            }
            _ = fi81
        }
        if (v4 != (v4 % (((^(v4))) | 1))) {
            if len(v51) >= 2 {
                v51 = v51[1:len(v51)]
            }
            for fi83 := 0; fi83 < 5; fi83++ {
                var v84 int16 = 1000
                fmt.Println(v2, fi83, v1)
                _ = fi83
                _ = v84
            }
            var v85 uint = (v3 ^ (v3 * 65535))
            _ = v85
        } else {
            v86 := GenS2{F0: (v1 != 209), F1: (v2 - v2), F2: (true && false)}
            _ = v86
        }
        _ = fi52
    }
    if ((2 > 0) || (!(42 != 2))) {
        for fi87 := 0; fi87 < 20; fi87++ {
            for fi88 := 0; fi88 < 3; fi88++ {
                if true {
                    if (false && false) {
                        _ = len(v51)
                        v33 += int(v1)
                    }
                    _ = len(v51)
                    v89 := GenS2{F0: (false && true), F1: (v2 >> ((693) & 63)), F2: true}
                    _ = v89
                } else {
                    v51 = append(v51, v33)
                }
                v90 := GenS10{F0: (v4 | v4), F1: -389, F2: -1}
                _ = fi88
                _ = v90
            }
            if true {
                _ = len(v51)
                if (668 != 1) {
                    if (true || false) {
                        for fi91 := 0; fi91 < 12; fi91++ {
                            if (false || true) {
                                var v92 uint64 = 1
                                var v93 int16 = 1
                                _ = v92
                                _ = v93
                            }
                            switch -1 {
                            case 1:
                                switch (fi87 - 32768) {
                                case 127:
                                    v94 := GenS11{F0: 0, F1: -457, F2: 1}
                                    _ = v94
                                case 42:
                                    var v95 uint = (v3 | (v3 << (((v3 / ((780) | 1))) & 63)))
                                    _ = v95
                                case 2147483647:
                                    v4 = v4
                                default:
                                    if true {
                                        v2 -= (v2 << ((v3) & 63))
                                    }
                                }
                                v51 = append(v51, (fi87 / ((598) | 1)))
                                if true {
                                    v96 := GenS2{F0: false, F1: (v2 << (((v3 % (((v3 * 0)) | 1))) & 63)), F2: (fi91 > fi91)}
                                    var v97 uint16 = 1
                                    v33 -= 1
                                    _ = v96
                                    _ = v97
                                }
                            case 128:
                                v98 := GenS7{F0: -1, F1: fi91, F2: v2}
                                _ = v98
                            }
                            var v99 uint8 = 127
                            _ = fi91
                            _ = v99
                        }
                        var v100 uint32 = 42
                        v1 -= 2
                        _ = v100
                    }
                }
                if len(v51) > 0 {
                    v51[0] = 32767
                }
            }
            var v101 uint16 = 1
            _ = fi87
            _ = v101
        }
        var v102 uint8 = (v1 / ((128) | 1))
        _ = v102
    }
    fmt.Println(v3, v4, v1)
    fmt.Println(v2)
    v1 += 150
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v51
}
