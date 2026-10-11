package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 int32; F2 int }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 bool; F2 uint64 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 string; F2 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 uint32 }
type GenS15 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 int8) bool {
    v7 := []int{42, 1, 362}
    if (true || (42 > 1)) {
        p0 += (p0 ^ 0)
        p0 = 42
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: p0, F1: (("foo bar" < ("" + "hello")) || false), F2: 781}
        _ = fi8
        _ = v9
    }
    var v10 float32 = float32(((float64(1e10) / 1.0) / 3.14159))
    v11 := GenS2{F0: (p0 ^ (p0 - 0))}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return false
}
func genFn1() string {
    var v12 uint = 255
    v13 := GenS1{F0: 13, F1: (359 <= -266), F2: 42}
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 862
    v3 := []int{}
    var v4 uint = (v2 % ((v2) | 1))
    var v5 uint64 = 874
    var v6 int = -9223372036854775808
    switch "foo bar" {
    case "a":
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint8 = 127
            v5 = v5
            var v16 int64 = -1
            _ = fi14
            _ = v15
            _ = v16
        }
    case "hello":
        if (255 == 222) {
            v6 = 2147483647
        } else {
            v17 := GenS3{F0: ((((((1e10 / ((-2.25 - 3.14159) / (3.14159 - 0.1))) * 1e10) - ((1.5 - (float32(1.5) + ((3.14159 * 3.14159) + 0.0))) / (0.0 + (0.0 + 3.14159)))) - 3.14159) <= 1.5) && true), F1: "foo bar"}
            _ = v17
        }
    case "foo bar":
        v2 = v2
    }
    if len(v3) > 0 {
        v3[0] = (v6 ^ v6)
    }
    var v18 int16 = 1000
    if (!(!(!(!(v2 >= v2))))) {
        fmt.Println(v2)
    } else {
        var v19 int8 = -128
        _ = v19
    }
    v20 := GenS1{F0: 127, F1: (!true), F2: 18446744073709551615}
    var v21 string = ("foo bar" + "\n\t")
    fmt.Println(v5)
    v2 = (v4 + v2)
    if (false || (true || true)) {
        if (-77 >= 42) {
            switch (v6 ^ (v6 * 128)) {
            case 1:
                v5 = v5
            default:
                if false {
                    if (!true) {
                        var v22 int64 = 1
                        for fi23 := 0; fi23 < 7; fi23++ {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v25 := GenS2{F0: -1}
                                _ = fi24
                                _ = v25
                            }
                            _ = fi23
                        }
                        _ = v22
                    }
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                fmt.Println(v1)
                _ = fi26
            }
            fmt.Println(v18)
        } else {
            fmt.Println(v1, v4, v21)
            for fi27 := 0; fi27 < 4; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    if (130 == 255) {
                        v21 = (("hello" + ("foo bar" + v21)) + ((v21 + v21) + (v21 + ("a" + v21))))
                        if false {
                            v4 += (v2 << (((v4 - 334)) & 63))
                            if false {
                                var v29 int = (fi27 ^ fi28)
                                var v30 float64 = 1.5
                                var v31 uint32 = 1
                                v5 += (v5 - (v5 * (v5 + (v5 % (((v5 % (((^(v5))) | 1))) | 1)))))
                                v5 -= 42
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS3{F0: (42 != 127), F1: v21}
                            var v33 int32 = int32(v5)
                            v6 -= (fi27 - (^(v6)))
                            _ = v32
                            _ = v33
                        } else {
                            if (false && (68 != -362)) {
                                _ = len(v3)
                            } else {
                                switch v21 {
                                case "":
                                    if (fi27 < v6) {
                                        switch v6 {
                                        case 847:
                                            var v34 int8 = 127
                                            var v35 bool = (534 < uint64(v4))
                                            switch v21 {
                                            case "":
                                                var v36 uint64 = (v5 | (v5 + (v5 + v5)))
                                                _ = v36
                                            case "\n\t":
                                                if (v35 && (v18 == 32767)) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v2 = (v2 * (v2 ^ v2))
                                                }
                                            default:
                                                if (((3.14159 - 0.0) / (((float64(float32(((float64(3.14159) / 1.0) + 1.5))) - 0.0) - 1e10) + 1.5)) >= -2.25) {
                                                    fmt.Println(v6)
                                                    if len(v3) > 0 {
                                                        v3[0] = fi28
                                                    }
                                                    var v37 string = v21
                                                    _ = v37
                                                }
                                            }
                                            _ = v34
                                            _ = v35
                                        }
                                        v38 := GenS4{F0: (v18 & 423), F1: -2147483648, F2: (fi27 + (fi27 * 0))}
                                        var v39 string = (v21 + "")
                                        _ = v38
                                        _ = v39
                                    } else {
                                        if (-2.25 <= 0.1) {
                                            if (v18 > (v18 | (v18 | v18))) {
                                                if (v1 < v1) {
                                                    var v40 float32 = 1e10
                                                    _ = v40
                                                }
                                            } else {
                                                if false {
                                                    v41 := GenS1{F0: -128, F1: true, F2: v4}
                                                    v4 += v4
                                                    if false {
                                                        v1 += 65535
                                                        fmt.Println(v4, v6)
                                                        if (-32768 <= 42) {
                                                            var v42 int16 = (v18 >> ((v2) & 63))
                                                            _ = v42
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = v6
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v6, v5, v1)
                                                        v4 = v4
                                                        fmt.Println(v6, v5, fi27)
                                                    }
                                                    if (!true) {
                                                        v43 := GenS5{F0: 42, F1: 0}
                                                        v2 = (v2 * v2)
                                                        _ = v43
                                                    }
                                                    v18 = v18
                                                    _ = v41
                                                }
                                            }
                                        } else {
                                            var v44 string = (v21 + ((v21 + (v21 + ("foo bar" + v21))) + v21))
                                            _ = v44
                                        }
                                    }
                                case "a":
                                    if ((int16(v4) > v18) && ((3.14159 / 1.0) <= 0.0)) {
                                        v3 = append(v3, 410)
                                    } else {
                                        v45 := GenS6{F0: "\n\t"}
                                        _ = v45
                                    }
                                case "foo bar":
                                    v18 = int16(v6)
                                default:
                                    var v46 int8 = 5
                                    _ = v46
                                }
                                v47 := GenS7{F0: (fi28 + fi27), F1: false, F2: v5}
                                _ = v47
                            }
                            if true {
                                fmt.Println(v6, fi28, v21)
                                v48 := GenS8{F0: 642, F1: 42}
                                if ((1 != 42) || true) {
                                    v49 := GenS9{F0: (v1 | v1)}
                                    v4 = 693
                                    if ((true && true) && (true && false)) {
                                        v50 := GenS2{F0: 94}
                                        _ = v50
                                    } else {
                                        if (uint64(v4) > v5) {
                                            switch -2147483648 {
                                            case 128:
                                                if ((1e10 < ((3.14159 + 0.1) - 1.5)) && false) {
                                                    v2 = uint(v5)
                                                    _ = len(v3)
                                                    v4 += (^(v2))
                                                } else {
                                                    switch v21 {
                                                    case "hello":
                                                        fmt.Println(fi27, v2)
                                                    case "foo bar":
                                                        v51 := GenS2{F0: 127}
                                                        _ = v51
                                                    case "\n\t":
                                                        var v52 uint64 = (^(v5))
                                                        _ = v52
                                                    default:
                                                        var v53 string = "\n\t"
                                                        _ = v53
                                                    }
                                                    if false {
                                                        switch "" {
                                                        case "hello":
                                                            var v54 int64 = -87
                                                            _ = v54
                                                        case "":
                                                            fmt.Println(v21)
                                                        case "\n\t":
                                                            v55 := GenS10{F0: v21, F1: 238}
                                                            _ = v55
                                                        }
                                                        var v56 uint32 = v1
                                                        v18 = 42
                                                        _ = v56
                                                    } else {
                                                        if false {
                                                            switch ("hello" + "") {
                                                            case "foo bar":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case "hello":
                                                                var v57 uint16 = 255
                                                                _ = v57
                                                            default:
                                                                v5 += v5
                                                            }
                                                            v58 := GenS11{F0: (v4 >> (((v4 / ((uint(v5)) | 1))) & 63))}
                                                            v59 := GenS8{F0: int32(v18), F1: 0}
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v1 += (v1 << ((v2) & 63))
                                                            if ((v2 << (((v4 / ((v2) | 1))) & 63)) <= (v2 + v4)) {
                                                                v60 := GenS2{F0: 127}
                                                                _ = v60
                                                            }
                                                        }
                                                        if ("a" < "") {
                                                            v2 = 4294967295
                                                            var v61 int64 = -759
                                                            if (0.1 > ((0.1 * 1.5) - (0.1 * 1.5))) {
                                                                var v62 uint8 = uint8(v1)
                                                                _ = v62
                                                            }
                                                            _ = v61
                                                        } else {
                                                            v4 = 947
                                                            var v63 bool = (-1 > -764)
                                                            _ = v63
                                                        }
                                                        switch (-(fi28)) {
                                                        case 817:
                                                            _ = len(v3)
                                                        case 32768:
                                                            v2 += v2
                                                        case 2147483647:
                                                            if false {
                                                                _ = len(v3)
                                                                v64 := GenS12{F0: uint64(v1), F1: v21, F2: 239}
                                                                var v65 uint = 1
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                        default:
                                                            var v66 uint = uint(v1)
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                            default:
                                                fmt.Println(v4, v2)
                                                v67 := GenS4{F0: -1, F1: 1, F2: -9223372036854775808}
                                                _ = v67
                                            }
                                            fmt.Println(v21)
                                            if true {
                                                var v68 uint8 = 174
                                                v6 = v6
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v68
                                            } else {
                                                v2 += v4
                                            }
                                        } else {
                                            if false {
                                                v2 += v2
                                                v6 = 9223372036854775807
                                            } else {
                                                v5 -= uint64(v2)
                                            }
                                            v4 += 255
                                            v4 = (v2 % (((v2 - v4)) | 1))
                                        }
                                        _ = len(v3)
                                        fmt.Println(fi27, v18, fi28)
                                    }
                                    _ = v49
                                }
                                _ = v48
                            }
                        }
                        v4 += 2
                        var v69 string = v21
                        switch (v6 & (-(v6))) {
                        case -2147483648:
                            switch v69 {
                            case "hello":
                                var v70 uint64 = (^(v5))
                                _ = v70
                            case "a":
                                var v71 uint16 = 105
                                _ = v71
                            default:
                                v21 = v21
                            }
                        case 4294967295:
                            var v72 uint8 = 255
                            _ = v72
                        default:
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        if true {
                            switch (fi27 | v6) {
                            case -845:
                                var v73 bool = true
                                _ = v73
                            case 0:
                                fmt.Println(v2, v21)
                            case 524:
                                fmt.Println(v18, v1, v6)
                            }
                        } else {
                            v2 += (v2 % (((v4 | v4)) | 1))
                            switch ("a" + (v21 + v69)) {
                            case "a":
                                v74 := GenS13{F0: 2147483647}
                                _ = v74
                            case "foo bar":
                                switch v69 {
                                case "foo bar":
                                    v75 := GenS14{F0: 0, F1: v1}
                                    _ = v75
                                case "\n\t":
                                    switch fi27 {
                                    case -2147483648:
                                        v76 := GenS6{F0: "hello"}
                                        _ = v76
                                    default:
                                        var v77 float32 = -2.25
                                        _ = v77
                                    }
                                }
                            default:
                                var v78 float32 = ((1.5 + -2.25) * 0.0)
                                _ = v78
                            }
                            fmt.Println(fi27)
                        }
                        _ = v69
                    }
                    v2 -= v2
                    _ = fi28
                }
                var v79 int = fi27
                _ = fi27
                _ = v79
            }
        }
        if true {
            var v80 uint64 = 1
            for fi81 := 0; fi81 < 6; fi81++ {
                v2 = v2
                _ = fi81
            }
            _ = v80
        }
        switch v21 {
        case "":
            if true {
                v1 -= 42
            } else {
                if (!(v18 == v18)) {
                    var v82 uint8 = 249
                    v83 := GenS15{F0: (v6 % (((v6 + v6)) | 1)), F1: (v18 ^ v18), F2: (v1 >> ((18446744073709551615) & 63))}
                    _ = len(v3)
                    _ = v82
                    _ = v83
                } else {
                    fmt.Println(v6, v18, v1)
                }
                var v84 float32 = 3.14159
                _ = v84
            }
        }
    }
    if (157 > 2) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi85 := 0; fi85 < 17; fi85++ {
            for fi86 := 0; fi86 < 17; fi86++ {
                fmt.Println(v21, v4)
                _ = fi86
            }
            fmt.Println(v1, v18, v4)
            v4 = uint(v1)
            _ = fi85
        }
        switch v21 {
        case "":
            if (556 < v1) {
                v1 = v1
            }
        case "\n\t":
            switch v21 {
            case "":
                if (-544 > 727) {
                    var v87 uint = 646
                    v4 = v2
                    _ = v87
                } else {
                    v18 = v18
                }
            case "hello":
                var v88 int8 = 127
                _ = v88
            case "a":
                var v89 float64 = 1.5
                _ = v89
            default:
                if (!(true || false)) {
                    var v90 uint32 = 0
                    var v91 float64 = ((-2.25 + float64(1e10)) - (1.5 + 1e10))
                    _ = v90
                    _ = v91
                }
            }
        default:
            if len(v3) > 0 {
                v3[0] = v6
            }
        }
    }
    fmt.Println(v18, v6)
    var v92 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
