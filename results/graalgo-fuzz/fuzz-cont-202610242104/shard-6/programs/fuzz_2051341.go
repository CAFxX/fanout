package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint16) uint32 {
    v7 := []int{32767, 2147483647, -493}
    if (2 < 2) {
        fmt.Println(p0)
        var v8 uint16 = 42
        _ = v8
    } else {
        var v9 uint32 = 2
        if (!false) {
            fmt.Println(p1, p1)
            v10 := GenS1{F0: 1, F1: "hello", F2: 1}
            var v11 uint = 994
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                p1 = p1
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v7) > 0 {
                        v7[0] = fi13
                    }
                    p0 += (p0 ^ (-(p0)))
                    v14 := GenS2{F0: 32767, F1: (v9 + v9), F2: (fi13 + fi12)}
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
            var v15 uint8 = 127
            _ = v15
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 338
}

func main() {
    var v1 bool = (-765 <= -690)
    var v2 uint16 = 1
    var v3 int8 = -85
    var v4 float32 = (3.14159 / 3.14159)
    var v5 uint = 65535
    v6 := []string{}
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 15; fi17++ {
        if v1 {
            v3 = v3
            var v18 float32 = (0.1 - (v4 / v4))
            switch "hello" {
            case "":
                var v19 uint32 = 65535
                _ = v19
            default:
                v16 = (("" + ("\n\t" + "")) + "")
            }
            _ = v18
        }
        _ = fi17
    }
    if (!false) {
        switch 727 {
        case 128:
            v20 := GenS1{F0: 1, F1: "", F2: 1}
            _ = v20
        case -258:
            v2 = (v2 + 42)
        case 492:
            var v21 float64 = 0.0
            _ = v21
        }
        if ((!(42 > 1)) && v1) {
            var v22 string = v16
            _ = v22
        }
        if (v1 || (42 == -32768)) {
            v3 = v3
            fmt.Println(v5, v5)
        } else {
            if ((v4 >= v4) || ((!v1) || (v1 && ((float64(-2.25) * 0.0) != ((1.5 + (1.5 + (0.0 * 1e10))) * (-2.25 * (3.14159 / ((-2.25 + -2.25) * (((-2.25 * ((1.5 - 0.0) * ((-2.25 / (1e10 + float64((v4 / v4)))) * 3.14159))) / -2.25) + 3.14159))))))))) {
                if (false || (int64(v2) == -9223372036854775808)) {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        fmt.Println(v1, v16)
                        fmt.Println(v5, v5)
                        fmt.Println(v16, v3)
                        _ = fi23
                    }
                }
                fmt.Println(v1, v3, v16)
                v16 = ("a" + (v16 + v16))
            }
            fmt.Println(v16)
            v6 = append(v6, "foo bar")
        }
    } else {
        v24 := GenS1{F0: (v5 % ((42) | 1)), F1: v16, F2: -395}
        v25 := GenS1{F0: v5, F1: (v16 + ((v16 + ((v16 + v16) + "\n\t")) + v16)), F2: 42}
        _ = v24
        _ = v25
    }
    v1 = v1
    v4 = (v4 + float32(((0.1 * (float64((v4 - 0.0)) - (0.1 - 3.14159))) / float64(((((v4 / (-2.25 / 1.0)) - float32(1.5)) / v4) + 1e10)))))
    var v26 uint8 = 2
    if true {
        if v1 {
            v27 := GenS3{F0: (v26 % ((42) | 1)), F1: 747}
            v28 := GenS2{F0: -22, F1: 2, F2: -2147483648}
            _ = v27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int32 = 42
        for fi31 := 0; fi31 < 10; fi31++ {
            v5 += v5
            v6 = append(v6, (v16 + v16))
            v32 := GenS2{F0: -32768, F1: 2147483647, F2: fi31}
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if ((^(v5)) <= 901) {
        v5 = v5
        if (!true) {
            if v1 {
                switch v16 {
                case "hello":
                    v2 -= 255
                case "a":
                    if v1 {
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v1 = (!(!(v1 && (true && v1))))
                            _ = fi33
                        }
                        _ = len(v6)
                    }
                default:
                    if ((((0.1 + ((v4 - float32((-2.25 / 1.0))) - (float32((0.0 - (float64(float32((0.0 + 0.0))) * 1e10))) * (0.1 * 0.1)))) * (v4 * ((float32(-2.25) + (v4 - float32(0.0))) * (v4 + v4)))) == v4) && false) {
                        switch v16 {
                        case "a":
                            v34 := GenS4{F0: "", F1: 127}
                            _ = v34
                        case "foo bar":
                            var v35 float32 = ((1e10 + v4) - v4)
                            _ = v35
                        default:
                            if v1 {
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        fmt.Println(v3, v3)
                                    case "\n\t":
                                        if (float32(0.0) > v4) {
                                            var v37 int16 = -1
                                            for fi38 := 0; fi38 < 4; fi38++ {
                                                var v39 bool = (v5 == v5)
                                                switch v16 {
                                                case "\n\t":
                                                    if true {
                                                        if (!(((float32(0.0) / 0.1) <= (1.5 + ((v4 + float32((-2.25 / (1.5 * (float64(3.14159) - 0.1))))) + float32(3.14159)))) || (!v1))) {
                                                            v26 = v26
                                                            v6 = append(v6, (v16 + v16))
                                                            if true {
                                                                v37 = (v37 - v37)
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if ((v3 % ((v3) | 1)) < (-(v3))) {
                                                                        var v40 uint = v5
                                                                        v16 = (v16 + (v16 + v16))
                                                                        _ = v40
                                                                    } else {
                                                                        if false {
                                                                            switch v16 {
                                                                            case "":
                                                                                var v41 uint = (v5 << ((v5) & 63))
                                                                                _ = v41
                                                                            case "\n\t":
                                                                                v42 := GenS1{F0: (v5 | (v5 >> ((v5) & 63))), F1: v16, F2: (fi38 | fi38)}
                                                                                _ = v42
                                                                            case "hello":
                                                                                v1 = false
                                                                            }
                                                                            fmt.Println(v39, v5)
                                                                        } else {
                                                                            if v1 {
                                                                                switch v16 {
                                                                                case "\n\t":
                                                                                    if (793 <= -1) {
                                                                                        v1 = v39
                                                                                    } else {
                                                                                        switch (fi36 % ((fi36) | 1)) {
                                                                                        case 819:
                                                                                            if v39 {
                                                                                                var v43 int8 = v3
                                                                                                if v39 {
                                                                                                    var v44 int64 = -9223372036854775808
                                                                                                    switch 1 {
                                                                                                    case 32767:
                                                                                                        if ((v16 + "") <= ("a" + v16)) {
                                                                                                            if v1 {
                                                                                                                v5 = (v5 * v5)
                                                                                                            }
                                                                                                            var v45 float64 = 3.14159
                                                                                                            var v46 float64 = float64((v4 * v4))
                                                                                                            _ = v45
                                                                                                            _ = v46
                                                                                                        } else {
                                                                                                            v47 := GenS2{F0: (v37 + (v37 & v37)), F1: 2147483647, F2: (fi36 | fi38)}
                                                                                                            _ = v47
                                                                                                        }
                                                                                                    default:
                                                                                                        switch v16 {
                                                                                                        case "a":
                                                                                                            switch (fi38 - (fi36 << ((v5) & 63))) {
                                                                                                            case 4294967295:
                                                                                                                var v48 uint16 = (v2 - 255)
                                                                                                                _ = v48
                                                                                                            case 976:
                                                                                                                v49 := GenS5{F0: v26, F1: v26}
                                                                                                                fmt.Println(v2)
                                                                                                                v50 := GenS2{F0: (v37 + (v37 | 1)), F1: 88, F2: (^(fi38))}
                                                                                                                _ = v49
                                                                                                                _ = v50
                                                                                                            case 9223372036854775807:
                                                                                                                v51 := GenS4{F0: v16, F1: int8(fi36)}
                                                                                                                _ = v51
                                                                                                            default:
                                                                                                                if len(v6) > 0 {
                                                                                                                    v6[0] = ((v16 + v16) + v16)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v44
                                                                                                }
                                                                                                fmt.Println(v2, v39)
                                                                                                _ = v43
                                                                                            } else {
                                                                                                var v52 uint = v5
                                                                                                _ = v52
                                                                                            }
                                                                                        case 42:
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                        case -941:
                                                                                            v5 -= (v5 - uint(v26))
                                                                                        default:
                                                                                            if true {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = ((("foo bar" + ("hello" + "foo bar")) + (v16 + ((v16 + "") + (v16 + (((v16 + (("foo bar" + ("a" + "foo bar")) + v16)) + "hello") + (v16 + ((v16 + "") + ("a" + v16)))))))) + "hello")
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if len(v6) > 0 {
                                                                                            v6[0] = ("hello" + (v16 + (v16 + v16)))
                                                                                        }
                                                                                        v2 += (^(v2))
                                                                                    }
                                                                                case "a":
                                                                                    v53 := GenS5{F0: (v26 / ((255) | 1)), F1: 128}
                                                                                    _ = v53
                                                                                case "":
                                                                                    if ((v26 - 56) >= v26) {
                                                                                        if len(v6) > 0 {
                                                                                            v6[0] = v16
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            switch fi36 {
                                                                                            case -2147483648:
                                                                                                var v54 int8 = -128
                                                                                                v55 := GenS4{F0: v16, F1: (v3 | -1)}
                                                                                                var v56 int = (fi38 * -9223372036854775808)
                                                                                                _ = v54
                                                                                                _ = v55
                                                                                                _ = v56
                                                                                            default:
                                                                                                fmt.Println(v16, v1)
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v57 := GenS3{F0: 1, F1: 616}
                                                                                        v2 = (^(v2))
                                                                                        _ = v57
                                                                                    }
                                                                                default:
                                                                                    v2 += v2
                                                                                }
                                                                                var v58 int64 = 403
                                                                                var v59 bool = false
                                                                                v1 = v59
                                                                                v60 := GenS6{F0: "\n\t", F1: uint32(v5)}
                                                                                _ = v58
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                var v61 int32 = 0
                                                                                _ = v61
                                                                            }
                                                                            fmt.Println(v26, fi36)
                                                                            var v62 uint64 = 0
                                                                            var v63 int16 = 521
                                                                            v64 := GenS1{F0: 0, F1: v16, F2: fi38}
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                        v65 := GenS6{F0: v16, F1: 2}
                                                                        _ = v65
                                                                    }
                                                                case "\n\t":
                                                                    var v66 float64 = (-2.25 + float64(0.1))
                                                                    _ = v66
                                                                default:
                                                                    var v67 uint8 = v26
                                                                    _ = v67
                                                                }
                                                                var v68 int32 = 42
                                                                _ = v68
                                                            }
                                                        } else {
                                                            var v69 uint = (^(v5))
                                                            _ = v69
                                                        }
                                                        var v70 string = (v16 + "\n\t")
                                                        _ = v70
                                                    } else {
                                                        fmt.Println(v4)
                                                        v37 = v37
                                                    }
                                                case "hello":
                                                    v6 = append(v6, "hello")
                                                case "a":
                                                    v6 = append(v6, v16)
                                                default:
                                                    v1 = (v5 >= 701)
                                                }
                                                v4 = v4
                                                _ = fi38
                                                _ = v39
                                            }
                                            _ = v37
                                        }
                                    case "a":
                                        v71 := GenS3{F0: (v26 % (((v26 & v26)) | 1)), F1: 32767}
                                        _ = v71
                                    default:
                                        v6 = append(v6, v16)
                                    }
                                    v72 := GenS7{F0: 746, F1: 747}
                                    _ = fi36
                                    _ = v72
                                }
                            }
                        }
                        for fi73 := 0; fi73 < 17; fi73++ {
                            if ((!v1) && true) {
                                fmt.Println(v26, fi73)
                                v74 := GenS5{F0: (^(v26)), F1: (v26 / ((v26) | 1))}
                                _ = v74
                            } else {
                                fmt.Println(v1)
                            }
                            var v75 uint8 = v26
                            _ = fi73
                            _ = v75
                        }
                    } else {
                        var v76 uint = uint(v26)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v76
                    }
                }
            } else {
                v16 = ((("hello" + v16) + (v16 + (("foo bar" + "foo bar") + ((("a" + v16) + "\n\t") + "a")))) + (v16 + (v16 + ((v16 + (v16 + "\n\t")) + v16))))
            }
        }
        var v77 float64 = 1.5
        _ = v77
    } else {
        if v1 {
            _ = len(v6)
            v3 += int8(v5)
            _ = len(v6)
        } else {
            var v78 int = -9223372036854775808
            if ((v78 ^ (v78 << ((v5) & 63))) < (v78 - 127)) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (!((!true) && ((!v1) && true))) {
                    fmt.Println(v5, v5, v26)
                    var v79 int32 = 0
                    var v80 string = v16
                    _ = v79
                    _ = v80
                } else {
                    v81 := GenS6{F0: v16, F1: 476}
                    v82 := GenS6{F0: "foo bar", F1: 2147483648}
                    var v83 int8 = int8(v2)
                    _ = v81
                    _ = v82
                    _ = v83
                }
                for fi84 := 0; fi84 < 8; fi84++ {
                    fmt.Println(v78, v16)
                    v16 = ""
                    var v85 int8 = v3
                    _ = fi84
                    _ = v85
                }
            }
            v78 = (v78 ^ v78)
            _ = v78
        }
        v5 = v5
    }
    switch -9223372036854775808 {
    case 881:
        var v86 int16 = 1
        _ = v86
    default:
        v87 := GenS1{F0: v5, F1: v16, F2: 42}
        _ = v87
    }
    var v88 int64 = -1
    if false {
        v26 += v26
        switch "foo bar" {
        case "":
            var v89 float32 = 1.5
            _ = v89
        case "a":
            if ((!true) || (!false)) {
                var v90 uint32 = uint32(v2)
                _ = v90
            }
        }
        if (1.5 < 1.5) {
            if (v16 <= v16) {
                var v91 bool = (v16 != v16)
                if false {
                    v91 = (true || v91)
                    if len(v6) > 0 {
                        v6[0] = (("hello" + ("" + v16)) + v16)
                    }
                    var v92 int32 = -197
                    _ = v92
                } else {
                    v26 += v26
                }
                _ = v91
            }
            var v93 float32 = v4
            _ = v93
        } else {
            var v94 uint = 42
            _ = v94
        }
    } else {
        switch v16 {
        case "foo bar":
            if v1 {
                v95 := GenS8{F0: 1}
                _ = v95
            } else {
                var v96 int8 = v3
                _ = v96
            }
        case "hello":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "\n\t":
            for fi97 := 0; fi97 < 11; fi97++ {
                var v98 float32 = float32(-2.25)
                v99 := GenS1{F0: v5, F1: (v16 + (v16 + v16)), F2: 0}
                fmt.Println(v98, v26, v3)
                _ = fi97
                _ = v98
                _ = v99
            }
        }
    }
    v88 -= 9223372036854775807
    v6 = append(v6, v16)
    for fi100 := 0; fi100 < 4; fi100++ {
        v4 = (((3.14159 / (float32((float64(v4) * (1e10 * float64(-2.25)))) / float32(float64((v4 - 0.0))))) * -2.25) - ((v4 + ((((0.0 / v4) + v4) - ((((v4 * -2.25) / 0.1) / float32(1e10)) * v4)) / (((v4 + v4) / v4) * (3.14159 * v4)))) * (0.1 / v4)))
        v101 := GenS6{F0: v16, F1: 392}
        _ = len(v6)
        _ = fi100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
}
