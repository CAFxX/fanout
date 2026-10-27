package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int32) uint64 {
    p0 += -174
    v6 := GenS1{F0: (p0 << ((18446744073709551615) & 63))}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: p0}
        if (32767 <= -32768) {
            for fi9 := 0; fi9 < 8; fi9++ {
                p0 = -1
                var v10 int64 = -9223372036854775808
                _ = fi9
                _ = v10
            }
        } else {
            p0 = (^(p0))
            v11 := GenS1{F0: (p0 - p0)}
            _ = v11
        }
        p0 = p0
        _ = fi7
        _ = v8
    }
    p0 -= p0
    _ = p0
    _ = v6
    return 738
}

func main() {
    v1 := []uint32{676, 4294967295}
    var v2 bool = (575 >= 42)
    var v3 int8 = -70
    var v4 int32 = -138
    var v5 string = "foo bar"
    if (!true) {
        var v12 int64 = 42
        for fi13 := 0; fi13 < 7; fi13++ {
            if v2 {
                v14 := GenS1{F0: (v4 ^ (v4 / ((0) | 1)))}
                v1 = append(v1, 4294967295)
                _ = v14
            } else {
                switch -2147483648 {
                case 9223372036854775807:
                    fmt.Println(v4, v2, v12)
                case -153:
                    if v2 {
                        var v15 int64 = -1
                        v2 = (557 < 842)
                        if ((3.14159 == (((1e10 - 1e10) / 1e10) * ((0.0 + 1e10) - -2.25))) || true) {
                            v15 -= (v12 >> ((464) & 63))
                            v2 = (-32768 < -462)
                            switch (v5 + ("foo bar" + ("a" + v5))) {
                            case "\n\t":
                                v1 = append(v1, 2)
                            case "":
                                var v16 float64 = (3.14159 + 1.5)
                                _ = v16
                            }
                        } else {
                            for fi17 := 0; fi17 < 5; fi17++ {
                                if (0 >= 1) {
                                    switch 32768 {
                                    case -2147483648:
                                        fmt.Println(v3, v3, v15)
                                    case 128:
                                        var v18 uint = 1
                                        _ = v18
                                    case -9223372036854775808:
                                        if v2 {
                                            _ = len(v1)
                                            switch v5 {
                                            case "\n\t":
                                                var v19 float32 = -2.25
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v20 int64 = v15
                                                var v21 int = (fi13 * fi17)
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                            v22 := GenS1{F0: (v4 % ((v4) | 1))}
                                            _ = v22
                                        } else {
                                            v15 = v15
                                            var v23 bool = (true && v2)
                                            _ = v23
                                        }
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                }
                                if (128 == fi17) {
                                    switch 32767 {
                                    case -9223372036854775808:
                                        var v24 uint16 = 971
                                        _ = v24
                                    case 4294967295:
                                        switch (fi13 / ((fi13) | 1)) {
                                        case 1:
                                            if ((v2 || (v4 == (-(v4)))) || true) {
                                                if len(v1) > 0 {
                                                    v1[0] = 1
                                                }
                                                var v25 string = (v5 + ((v5 + "") + "a"))
                                                _ = v25
                                            } else {
                                                var v26 uint16 = 0
                                                _ = v26
                                            }
                                        case 2147483647:
                                            var v27 uint = 0
                                            _ = v27
                                        case -208:
                                            switch -2147483648 {
                                            case -225:
                                                var v28 int = (fi13 - (fi17 + 127))
                                                _ = v28
                                            case -261:
                                                var v29 bool = (1.5 >= (0.1 - 0.0))
                                                _ = v29
                                            case 127:
                                                var v30 float32 = (((-2.25 / float32(0.1)) * 3.14159) + 3.14159)
                                                _ = v30
                                            default:
                                                var v31 bool = (((0.0 - 3.14159) < (1e10 / 1.0)) || (!(60 != 127)))
                                                _ = v31
                                            }
                                        }
                                    case 1:
                                        v2 = v2
                                    }
                                    v32 := GenS1{F0: (v4 ^ (^(v4)))}
                                    switch 0 {
                                    case 9223372036854775807:
                                        var v33 uint8 = 128
                                        _ = v33
                                    case 128:
                                        v4 += v4
                                    case 6:
                                        v5 = v5
                                    }
                                    _ = v32
                                }
                                if v2 {
                                    switch (-(fi13)) {
                                    case 0:
                                        if (4294967295 > 1) {
                                            switch v5 {
                                            case "a":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "hello":
                                                v34 := GenS2{F0: 42, F1: (v3 ^ (v3 | (v3 + (v3 % ((0) | 1))))), F2: -27}
                                                _ = v34
                                            case "\n\t":
                                                v35 := GenS1{F0: (v4 / ((v4) | 1))}
                                                _ = v35
                                            }
                                            var v36 int = fi17
                                            _ = v36
                                        }
                                    case 4294967295:
                                        if ((!v2) || true) {
                                            fmt.Println(fi17, v5, v12)
                                            fmt.Println(v3, v3, v5)
                                        } else {
                                            if (0 != 42) {
                                                v3 = (v3 * v3)
                                                _ = len(v1)
                                            }
                                        }
                                    case -992:
                                        if (0.0 >= (float32((1.5 / 1.5)) * (-2.25 * 1e10))) {
                                            fmt.Println(v4, v5, v5)
                                            if (!v2) {
                                                var v37 float32 = ((((float32((0.1 / 1.5)) * -2.25) + 1.5) * -2.25) * 0.1)
                                                _ = v37
                                            } else {
                                                if (!v2) {
                                                    v38 := GenS3{F0: 978, F1: 18446744073709551615}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v15 = 42
                                                    _ = v38
                                                }
                                                v4 = v4
                                            }
                                        }
                                    default:
                                        if v2 {
                                            if v2 {
                                                var v39 int = fi13
                                                _ = v39
                                            } else {
                                                fmt.Println(v5, v3)
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                if false {
                                                    v4 -= (v4 - v4)
                                                    var v40 string = v5
                                                    v2 = v2
                                                    if (101 == 0) {
                                                        var v41 uint32 = 0
                                                        if false {
                                                            if true {
                                                                v1 = append(v1, (^(v41)))
                                                                var v42 float32 = (float32(0.0) * (1.5 + (3.14159 * 1.5)))
                                                                _ = v42
                                                            }
                                                            if v2 {
                                                                v4 = (v4 ^ (v4 / ((-1) | 1)))
                                                                var v43 bool = (42 <= v41)
                                                                _ = v43
                                                            } else {
                                                                var v44 int8 = v3
                                                                var v45 uint16 = 645
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        } else {
                                                            v12 = (v15 + (-(v12)))
                                                            v46 := GenS1{F0: v4}
                                                            v47 := GenS3{F0: 18446744073709551615, F1: 255}
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        _ = v41
                                                    }
                                                    v2 = (v3 >= v3)
                                                    _ = v40
                                                }
                                            case "":
                                                if v2 {
                                                    v2 = v2
                                                } else {
                                                    if (-32768 > -949) {
                                                        fmt.Println(v2)
                                                        switch (((v5 + v5) + ((v5 + v5) + "hello")) + "") {
                                                        case "foo bar":
                                                            var v48 uint = 18446744073709551615
                                                            _ = v48
                                                        case "":
                                                            if true {
                                                                var v49 float64 = -2.25
                                                                if false {
                                                                    var v50 string = v5
                                                                    switch ((v5 + "") + "\n\t") {
                                                                    case "foo bar":
                                                                        v51 := GenS3{F0: 994, F1: 635}
                                                                        _ = v51
                                                                    case "a":
                                                                        if (((0.1 - ((float32(((v49 * v49) + v49)) * (float32(v49) / 1.5)) + (((0.0 + 3.14159) * 0.1) + float32((0.1 * (0.1 + ((((v49 - ((-2.25 + v49) / v49)) - v49) * v49) * (float64((3.14159 + ((((1e10 + 0.1) * 3.14159) - 0.0) + ((3.14159 + (-2.25 + 1.5)) + 0.0)))) - ((v49 / v49) / v49))))))))) / 3.14159) != 3.14159) {
                                                                            fmt.Println(v4)
                                                                            switch (v5 + "a") {
                                                                            case "":
                                                                                v12 += 0
                                                                            case "hello":
                                                                                var v52 bool = (!v2)
                                                                                _ = v52
                                                                            default:
                                                                                v3 -= (v3 | v3)
                                                                            }
                                                                            v50 = ("foo bar" + v50)
                                                                        } else {
                                                                            v12 -= (-(v15))
                                                                            switch (("foo bar" + "a") + "foo bar") {
                                                                            case "\n\t":
                                                                                fmt.Println(v4, v50, v49)
                                                                            }
                                                                            if v2 {
                                                                                v53 := GenS1{F0: 0}
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v54 float32 = (float32(v49) * (-2.25 / (1.5 + (-2.25 - (float32(v49) / 1.0)))))
                                                                        _ = v54
                                                                    default:
                                                                        v55 := GenS4{F0: 155}
                                                                        v56 := GenS2{F0: 128, F1: -128, F2: v12}
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    _ = v50
                                                                }
                                                                _ = v49
                                                            }
                                                        case "\n\t":
                                                            v57 := GenS4{F0: 890}
                                                            _ = v57
                                                        default:
                                                            if v2 {
                                                                v1 = append(v1, 65535)
                                                                var v58 int16 = 939
                                                                if false {
                                                                    var v59 uint16 = 1
                                                                    _ = v59
                                                                } else {
                                                                    v5 = v5
                                                                    var v60 int32 = v4
                                                                    _ = v60
                                                                }
                                                                _ = v58
                                                            }
                                                        }
                                                    } else {
                                                        var v61 uint = 18446744073709551615
                                                        _ = v61
                                                    }
                                                    var v62 uint = 4294967295
                                                    _ = v62
                                                }
                                            case "a":
                                                var v63 int16 = 32767
                                                _ = v63
                                            default:
                                                var v64 uint8 = 234
                                                v12 = 711
                                                var v65 uint8 = (v64 | v64)
                                                v66 := GenS3{F0: uint64(v64), F1: 0}
                                                fmt.Println(fi13, v5)
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                            v12 -= (v15 << ((18446744073709551615) & 63))
                                        } else {
                                            switch (^(fi17)) {
                                            case 2147483647:
                                                if len(v1) > 0 {
                                                    v1[0] = 112
                                                }
                                            default:
                                                switch v5 {
                                                case "foo bar":
                                                    fmt.Println(v15, fi13, v5)
                                                case "a":
                                                    v15 = (^(v15))
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(fi17, v15)
                                    var v67 float64 = 1.5
                                    _ = v67
                                } else {
                                    var v68 int = fi17
                                    _ = v68
                                }
                                _ = fi17
                            }
                            var v69 float32 = 1e10
                            if (((((((3.14159 / 1e10) + (((1e10 * (v69 * v69)) / v69) / ((((v69 - ((v69 - (3.14159 / 1.5)) - 0.0)) / v69) + v69) * v69))) - (v69 + (v69 * 1e10))) - -2.25) / (((v69 + v69) * 3.14159) * 0.0)) - (v69 * 0.0)) >= v69) {
                                var v70 uint = 255
                                _ = v70
                            }
                            _ = v69
                        }
                        _ = v15
                    } else {
                        fmt.Println(v12, v4)
                    }
                default:
                    for fi71 := 0; fi71 < 9; fi71++ {
                        v5 = "foo bar"
                        v2 = (((!v2) && ((v4 << ((443) & 63)) > -1)) && v2)
                        if true {
                            var v72 string = (v5 + v5)
                            v5 = v72
                            v4 = v4
                            _ = v72
                        } else {
                            if (1.5 >= 0.1) {
                                _ = len(v1)
                                if (255 >= 42) {
                                    var v73 uint16 = 0
                                    v12 += v12
                                    _ = v73
                                } else {
                                    switch fi71 {
                                    case 9223372036854775807:
                                        if v2 {
                                            if len(v1) > 0 {
                                                v1[0] = 630
                                            }
                                            v74 := GenS3{F0: 334, F1: 42}
                                            var v75 bool = (2 == 733)
                                            _ = v74
                                            _ = v75
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = 143
                                            }
                                            v5 = "\n\t"
                                            var v76 float32 = 1e10
                                            v3 = -9
                                            _ = v76
                                        }
                                    default:
                                        switch ("\n\t" + v5) {
                                        case "":
                                            if ((v12 & int64(v3)) > (v12 + (v12 << ((1) & 63)))) {
                                                v1 = append(v1, 854)
                                            }
                                        default:
                                            v12 += 9223372036854775807
                                        }
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                            }
                        }
                        _ = fi71
                    }
                }
                v77 := GenS4{F0: 255}
                _ = v77
            }
            if (389 != 1) {
                var v78 int8 = (v3 >> ((42) & 63))
                _ = v78
            }
            v79 := GenS5{F0: 2, F1: 1}
            _ = fi13
            _ = v79
        }
        _ = v12
    }
    if false {
        switch -835 {
        case -479:
            v2 = (568 == 0)
        case 127:
            if ((3.14159 < 0.1) || (-527 == int16(v4))) {
                switch v5 {
                case "a":
                    fmt.Println(v4, v4, v5)
                default:
                    var v80 int = 32767
                    _ = v80
                }
            }
        default:
            var v81 uint32 = 2147483648
            _ = v81
        }
        var v82 int16 = 0
        v83 := GenS3{F0: 2, F1: 1}
        _ = v82
        _ = v83
    }
    v2 = (true || true)
    if (485 != 1) {
        if v2 {
            if (false && (v2 || (!v2))) {
                fmt.Println(v4)
            }
            switch 390 {
            case 4294967295:
                var v84 float32 = ((1.5 * ((1e10 * 3.14159) + 1e10)) - 1.5)
                _ = v84
            case 128:
                v1 = append(v1, 595)
            }
            if true {
                fmt.Println(v3, v4)
                if v2 {
                    var v85 int = 958
                    v85 += 42
                    _ = v85
                }
            } else {
                fmt.Println(v4, v5)
                for fi86 := 0; fi86 < 4; fi86++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ((^(v4)) < (-(v4))) {
                        v5 = ("hello" + ((("\n\t" + "\n\t") + "") + ""))
                    }
                    var v87 bool = true
                    _ = fi86
                    _ = v87
                }
                if v2 {
                    v88 := GenS4{F0: 359}
                    _ = v88
                }
            }
        }
        var v89 uint32 = 65535
        var v90 int32 = (v4 ^ (v4 << ((65535) & 63)))
        _ = v89
        _ = v90
    } else {
        for fi91 := 0; fi91 < 1; fi91++ {
            if (v2 && ((!true) || (!v2))) {
                v92 := GenS3{F0: 1, F1: 255}
                v3 = (v3 - v3)
                _ = v92
            }
            for fi93 := 0; fi93 < 3; fi93++ {
                var v94 bool = (!v2)
                var v95 uint64 = 345
                _ = fi93
                _ = v94
                _ = v95
            }
            var v96 uint8 = 171
            _ = fi91
            _ = v96
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch ("a" + v5) {
        case "":
            v3 -= (v3 * (v3 << ((604) & 63)))
        default:
            for fi97 := 0; fi97 < 7; fi97++ {
                var v98 string = ((v5 + v5) + "")
                v3 = (v3 + v3)
                _ = fi97
                _ = v98
            }
        }
    }
    var v99 uint64 = 18446744073709551615
    var v100 float64 = float64(((1.5 * 1e10) + 1.5))
    fmt.Println(v2, v99, v99)
    if (v3 <= int8(v99)) {
        var v101 bool = v2
        fmt.Println(v2, v3, v3)
        _ = v101
    }
    var v102 uint64 = (v99 + (v99 | v99))
    var v103 int32 = (v4 + (v4 & -1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
