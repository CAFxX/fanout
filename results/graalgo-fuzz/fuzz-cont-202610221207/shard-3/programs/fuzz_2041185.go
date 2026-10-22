package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 int; F1 int; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = (((float32(float64(1e10)) * 1.5) + float32(((-2.25 / 0.1) * 3.14159))) - 0.0)
    v3 := []string{"a", "hello", "", ""}
    var v4 int = -9223372036854775808
    var v5 bool = (9 > 181)
    v6 := GenS1{F0: 1}
    if (42 != 42) {
        var v7 bool = (32767 != 42)
        v8 := GenS1{F0: 72}
        _ = v7
        _ = v8
    }
    if (true && (312 > 1)) {
        if v5 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi9 := 0; fi9 < 7; fi9++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi10 := 0; fi10 < 4; fi10++ {
                    if v5 {
                        switch "a" {
                        case "\n\t":
                            if v5 {
                                v3 = append(v3, ("" + "\n\t"))
                                fmt.Println(v5)
                            }
                        case "":
                            v2 = 1.5
                        case "foo bar":
                            var v11 int64 = -41
                            _ = v11
                        }
                        var v12 uint8 = 109
                        if (!(0.0 <= (1.5 * -2.25))) {
                            if (1000 < -564) {
                                v13 := GenS1{F0: uint16(v12)}
                                v5 = v5
                                fmt.Println(v4, v12)
                                _ = v13
                            } else {
                                v2 -= 0.0
                                v1 += (v1 / (((^(v1))) | 1))
                            }
                            var v14 uint8 = v12
                            v15 := GenS1{F0: 0}
                            _ = v14
                            _ = v15
                        } else {
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                fmt.Println(v1)
                            case "":
                                v16 := GenS2{F0: (v4 % ((-1) | 1)), F1: -77, F2: v12}
                                _ = v16
                            default:
                                if v5 {
                                    v12 = (v12 ^ 2)
                                    fmt.Println(fi10, v1)
                                }
                            }
                            fmt.Println(v4, v1, fi9)
                            v4 += (v4 ^ v4)
                        }
                        _ = v12
                    } else {
                        fmt.Println(v4, fi10, v2)
                        if true {
                            if v5 {
                                v1 = 2147483648
                                var v17 string = ""
                                if true {
                                    if (65535 == 674) {
                                        var v18 uint = 2
                                        var v19 int = (v4 ^ (fi10 - v4))
                                        v17 = (v17 + v17)
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v3 = append(v3, (v17 + (v17 + ((("\n\t" + v17) + "") + (v17 + ((((("" + ("" + "a")) + (v17 + "foo bar")) + "hello") + "hello") + v17))))))
                                        if v5 {
                                            var v20 uint = 42
                                            switch v17 {
                                            case "hello":
                                                v2 += (-2.25 - -2.25)
                                            case "foo bar":
                                                fmt.Println(fi10)
                                            default:
                                                if v5 {
                                                    fmt.Println(v1, v5, v1)
                                                    if len(v3) > 0 {
                                                        v3[0] = (("" + (v17 + v17)) + (v17 + ""))
                                                    }
                                                } else {
                                                    if true {
                                                        v17 = v17
                                                    }
                                                    var v21 int8 = 42
                                                    _ = v21
                                                }
                                                switch fi10 {
                                                case 42:
                                                    v22 := GenS3{F0: (v4 + (v4 & (fi9 % (((^(fi10))) | 1)))), F1: (fi9 & 4294967295), F2: v17}
                                                    v1 += (v1 | 0)
                                                    _ = v22
                                                case 1:
                                                    var v23 uint8 = 2
                                                    _ = v23
                                                case -2147483648:
                                                    var v24 float64 = ((0.0 * 1.5) - 3.14159)
                                                    _ = v24
                                                default:
                                                    v5 = false
                                                }
                                                v5 = ((float32(0.0) + (v2 / float32(1e10))) <= v2)
                                            }
                                            v25 := GenS2{F0: 4294967295, F1: -304, F2: 255}
                                            _ = v20
                                            _ = v25
                                        } else {
                                            var v26 int = 735
                                            _ = v26
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                    }
                                    _ = len(v3)
                                    var v27 uint = 2
                                    fmt.Println(v1, v4, fi10)
                                    _ = v27
                                } else {
                                    if v5 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                                _ = v17
                            } else {
                                var v28 uint16 = 42
                                if (-1 <= -1) {
                                    v5 = true
                                    v29 := GenS2{F0: -9223372036854775808, F1: 32767, F2: 0}
                                    v1 -= v1
                                    _ = v29
                                } else {
                                    var v30 int = v4
                                    v2 = float32(float64(v2))
                                    _ = v30
                                }
                                if true {
                                    var v31 uint8 = 127
                                    v5 = (v4 < (^(fi10)))
                                    _ = v31
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v32 int64 = 1
                                    var v33 int32 = 42
                                    v1 -= (v1 - (^(v1)))
                                    v34 := GenS2{F0: -321, F1: 42, F2: uint8(v28)}
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                _ = v28
                            }
                        } else {
                            switch (fi9 >> ((20) & 63)) {
                            case 127:
                                if v5 {
                                    var v35 int8 = 0
                                    _ = v35
                                }
                            default:
                                if (1 < 2) {
                                    if (v5 || ("a" >= (("\n\t" + "\n\t") + "foo bar"))) {
                                        if v5 {
                                            v5 = (0 < 255)
                                            _ = len(v3)
                                            fmt.Println(v4, v1)
                                            var v36 uint64 = 4294967295
                                            v2 = -2.25
                                            _ = v36
                                        } else {
                                            switch v4 {
                                            case 42:
                                                v5 = v5
                                            case -9223372036854775808:
                                                fmt.Println(v2, v1, v5)
                                            }
                                            v5 = (27 >= -86)
                                            var v37 float64 = 1e10
                                            _ = v37
                                        }
                                        switch 127 {
                                        case -1:
                                            fmt.Println(fi10, v4, v2)
                                        }
                                        var v38 string = "a"
                                        _ = v38
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v39 := GenS4{F0: v5}
                                    fmt.Println(v1, v1)
                                    _ = v39
                                }
                            }
                            var v40 int64 = 42
                            fmt.Println(fi10, fi9, fi9)
                            _ = v40
                        }
                        v5 = false
                    }
                    var v41 uint16 = 32767
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi10
                    _ = v41
                }
                _ = fi9
            }
            v4 = -1
        } else {
            fmt.Println(v2, v1)
            var v42 int8 = -1
            v43 := GenS5{F0: 42, F1: 0, F2: -490}
            _ = v42
            _ = v43
        }
    } else {
        if (65535 >= 65535) {
            v1 = (^(v1))
            var v44 int16 = 1000
            fmt.Println(v1, v4)
            _ = v44
        }
        v45 := GenS4{F0: (false && (127 == 100))}
        _ = v45
    }
    v46 := GenS1{F0: 0}
    v47 := GenS4{F0: false}
    for fi48 := 0; fi48 < 5; fi48++ {
        for fi49 := 0; fi49 < 3; fi49++ {
            var v50 int64 = 0
            v5 = (2 == 201)
            _ = fi49
            _ = v50
        }
        switch "foo bar" {
        case "foo bar":
            var v51 float64 = 3.14159
            _ = v51
        case "":
            if (0 > 1) {
                if ((-821 >= 1) || ((2 < 65535) && ((!v5) && (((-2.25 - 1.5) - 3.14159) >= float64(-2.25))))) {
                    v2 = 1.5
                    v52 := GenS3{F0: (v4 * -9223372036854775808), F1: 928, F2: "\n\t"}
                    var v53 int8 = -1
                    _ = v52
                    _ = v53
                }
                for fi54 := 0; fi54 < 7; fi54++ {
                    if v5 {
                        var v55 uint64 = 1
                        _ = v55
                    } else {
                        _ = len(v3)
                    }
                    if (4294967295 <= 482) {
                        fmt.Println(fi54)
                        v4 = fi48
                    } else {
                        v5 = (128 <= 127)
                        v1 = v1
                    }
                    var v56 uint32 = v1
                    var v57 string = ("hello" + ("hello" + ("\n\t" + "hello")))
                    var v58 uint16 = 42
                    if (1 != 318) {
                        if (!(((((v57 + v57) + v57) > v57) && false) && (0.1 < (-2.25 * 1.5)))) {
                            v59 := GenS1{F0: uint16(v1)}
                            _ = len(v3)
                            _ = v59
                        }
                        v60 := GenS5{F0: 128, F1: 0, F2: -220}
                        _ = v60
                    } else {
                        v57 = ((v57 + (v57 + v57)) + v57)
                    }
                    _ = fi54
                    _ = v56
                    _ = v57
                    _ = v58
                }
                v4 += v4
            } else {
                var v61 uint64 = 42
                _ = v61
            }
        default:
            for fi62 := 0; fi62 < 17; fi62++ {
                var v63 string = ("" + "hello")
                switch (^(fi48)) {
                case 562:
                    if (-9223372036854775808 == 9223372036854775807) {
                        switch "foo bar" {
                        case "":
                            if (!v5) {
                                var v64 int64 = 9223372036854775807
                                v1 = v1
                                if len(v3) > 0 {
                                    v3[0] = v63
                                }
                                _ = v64
                            } else {
                                v2 = (float32(3.14159) * 1e10)
                                v4 += v4
                            }
                        }
                        var v65 uint16 = 468
                        var v66 int16 = 32767
                        v67 := GenS2{F0: (-(fi48)), F1: v66, F2: 193}
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                case -2147483648:
                    v68 := GenS4{F0: (!(!v5))}
                    _ = v68
                case 9223372036854775807:
                    switch v63 {
                    case "foo bar":
                        switch v63 {
                        case "foo bar":
                            if (v5 || v5) {
                                v5 = (!((fi62 >> ((1) & 63)) <= 596))
                                var v69 int16 = -848
                                _ = v69
                            } else {
                                v70 := GenS2{F0: (fi62 & 2147483647), F1: 42, F2: 255}
                                if true {
                                    fmt.Println(fi62, v1)
                                }
                                _ = v70
                            }
                        case "\n\t":
                            var v71 uint32 = v1
                            fmt.Println(v71, v71, fi48)
                            _ = len(v3)
                            fmt.Println(fi62)
                            _ = v71
                        default:
                            v63 = "hello"
                        }
                    case "a":
                        v72 := GenS6{F0: 543}
                        _ = v72
                    }
                }
                if (16 <= 136) {
                    v73 := GenS1{F0: 2}
                    if true {
                        fmt.Println(v1)
                        if ((!(169 < 754)) || (!(0 != 0))) {
                            var v74 string = v63
                            v75 := GenS7{F0: 136}
                            v3 = append(v3, (((("\n\t" + v74) + v63) + (v74 + v74)) + v63))
                            _ = v74
                            _ = v75
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch fi48 {
                        case -53:
                            v3 = append(v3, v63)
                        case -9223372036854775808:
                            _ = len(v3)
                        case 127:
                            var v76 uint32 = 2
                            _ = v76
                        default:
                            v2 = (v2 / v2)
                        }
                    } else {
                        v2 = v2
                        v77 := GenS8{F0: -2147483648}
                        if ((v2 > v2) && (32767 >= 653)) {
                            var v78 float64 = 1.5
                            _ = len(v3)
                            _ = v78
                        } else {
                            var v79 uint16 = 42
                            v80 := GenS9{F0: 77, F1: (v79 - (v79 % (((v79 << ((18446744073709551615) & 63))) | 1)))}
                            switch 20 {
                            case 9223372036854775807:
                                switch ((v63 + (v63 + v63)) + (v63 + (("\n\t" + v63) + "hello"))) {
                                case "\n\t":
                                    switch ((((v63 + (v63 + v63)) + ((((v63 + "a") + "a") + v63) + "")) + "") + ("a" + v63)) {
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v81 := GenS10{F0: 58, F1: 0}
                                        _ = v81
                                    }
                                case "":
                                    var v82 uint16 = (v79 | 2)
                                    _ = v82
                                }
                            case 1:
                                v4 = (v4 - 32768)
                            default:
                                if true {
                                    fmt.Println(v79)
                                    v5 = v5
                                } else {
                                    var v83 uint16 = v79
                                    _ = v83
                                }
                            }
                            var v84 int = (-(fi48))
                            if (454 != (v4 | fi48)) {
                                v4 = 9223372036854775807
                            } else {
                                v5 = (!v5)
                                v1 = v1
                            }
                            var v85 bool = (!(!(42 <= -32768)))
                            _ = v79
                            _ = v80
                            _ = v84
                            _ = v85
                        }
                        _ = v77
                    }
                    _ = v73
                }
                _ = fi62
                _ = v63
            }
        }
        var v86 int16 = 1000
        _ = fi48
        _ = v86
    }
    var v87 string = "hello"
    var v88 bool = v5
    if true {
        var v89 int16 = 32767
        _ = v89
    } else {
        v2 = ((v2 + (v2 * v2)) * v2)
        for fi90 := 0; fi90 < 20; fi90++ {
            for fi91 := 0; fi91 < 12; fi91++ {
                v2 += v2
                v87 = v87
                v92 := GenS8{F0: -2147483648}
                _ = fi91
                _ = v92
            }
            if false {
                if ((0 != -1) && false) {
                    v93 := GenS4{F0: v5}
                    v87 = v87
                    _ = v93
                }
                if (1 == 2147483647) {
                    v4 -= -1
                }
                v87 = "hello"
            }
            v94 := GenS3{F0: -1, F1: fi90, F2: "hello"}
            _ = fi90
            _ = v94
        }
        switch 0 {
        case 42:
            fmt.Println(v88)
        default:
            switch v4 {
            case 7:
                v1 += (v1 - v1)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v87
    _ = v88
}
