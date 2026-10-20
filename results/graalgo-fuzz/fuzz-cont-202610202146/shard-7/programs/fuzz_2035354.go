package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 bool }
type GenS9 struct { F0 uint64; F1 int8; F2 uint16 }
func genFn0(p0 uint, p1 int32) uint {
    p0 = (^(p0))
    p1 += p1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return (p0 * (p0 ^ p0))
}

func main() {
    var v1 bool = ((69 >= 32767) && false)
    v2 := []int{1, -1, -2147483648}
    var v3 int8 = 127
    v3 -= -1
    if (-9223372036854775808 <= 402) {
        switch "foo bar" {
        case "\n\t":
            v4 := GenS1{F0: 128}
            _ = v4
        case "hello":
            fmt.Println(v3, v1)
        }
        if false {
            for fi5 := 0; fi5 < 13; fi5++ {
                if v1 {
                    for fi6 := 0; fi6 < 3; fi6++ {
                        var v7 uint8 = 255
                        var v8 string = ""
                        v9 := GenS1{F0: (v7 ^ (v7 << ((0) & 63)))}
                        _ = fi6
                        _ = v7
                        _ = v8
                        _ = v9
                    }
                    v3 += (v3 % ((59) | 1))
                } else {
                    var v10 int = (fi5 >> ((1) & 63))
                    v10 += 776
                    _ = v10
                }
                if true {
                    switch ("\n\t" + "hello") {
                    case "":
                        v3 += (^(v3))
                    case "a":
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: 127}
                            v13 := GenS2{F0: -1, F1: 464}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    default:
                        v3 += v3
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v14 := GenS1{F0: 255}
                _ = fi5
                _ = v14
            }
        }
    }
    if true {
        if v1 {
            switch 0 {
            case -2147483648:
                var v15 uint64 = 176
                _ = v15
            }
            v1 = (-2147483648 > -2147483648)
            var v16 string = ("foo bar" + "\n\t")
            _ = v16
        } else {
            fmt.Println(v3, v1)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v3 = (v3 + (v3 / ((121) | 1)))
    }
    switch 0 {
    case 127:
        if v1 {
            var v17 uint32 = 2147483647
            v18 := GenS2{F0: -1, F1: 1}
            v19 := GenS3{F0: 65535, F1: 2}
            _ = v17
            _ = v18
            _ = v19
        } else {
            fmt.Println(v1, v3)
            for fi20 := 0; fi20 < 14; fi20++ {
                v3 += v3
                switch ("" + "\n\t") {
                case "a":
                    var v21 bool = ((false || v1) && (true && (2 < 854)))
                    _ = v21
                case "hello":
                    for fi22 := 0; fi22 < 5; fi22++ {
                        if (0.0 >= 0.0) {
                            switch "hello" {
                            case "\n\t":
                                var v23 uint16 = 65535
                                _ = v23
                            case "foo bar":
                                v24 := GenS1{F0: 12}
                                _ = v24
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    var v25 uint32 = 2147483647
                                    _ = v25
                                case "a":
                                    var v26 int = (fi20 * 390)
                                    _ = v26
                                }
                                if false {
                                    v1 = true
                                    var v27 bool = false
                                    v2 = append(v2, 0)
                                    v28 := GenS3{F0: 65535, F1: 2}
                                    if v27 {
                                        if (469 == 4294967295) {
                                            var v29 float64 = float64(0.0)
                                            _ = v29
                                        } else {
                                            if v1 {
                                                switch 32767 {
                                                case 128:
                                                    if v1 {
                                                        _ = len(v2)
                                                        switch ("" + ("\n\t" + ("a" + ((("foo bar" + "hello") + "a") + (("" + "\n\t") + ("hello" + "foo bar")))))) {
                                                        case "\n\t":
                                                            if (!false) {
                                                                fmt.Println(v3)
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi20
                                                                }
                                                                v3 += (v3 * 100)
                                                            }
                                                        case "a":
                                                            v27 = false
                                                        case "hello":
                                                            switch "" {
                                                            case "":
                                                                switch ("\n\t" + "a") {
                                                                case "hello":
                                                                    var v30 int16 = 0
                                                                    _ = v30
                                                                }
                                                            case "a":
                                                                var v31 uint16 = 42
                                                                _ = v31
                                                            }
                                                            v32 := GenS4{F0: (^(fi22)), F1: 0}
                                                            _ = v32
                                                        default:
                                                            v27 = ((!false) && true)
                                                        }
                                                    }
                                                case -125:
                                                    _ = len(v2)
                                                }
                                                switch "" {
                                                case "a":
                                                    var v33 int64 = -13
                                                    _ = v33
                                                case "":
                                                    v27 = (42 < 4294967295)
                                                case "foo bar":
                                                    var v34 int8 = v3
                                                    _ = v34
                                                default:
                                                    v27 = (!(-2.25 <= 0.1))
                                                }
                                            } else {
                                                var v35 uint8 = 127
                                                _ = len(v2)
                                                if true {
                                                    if ((-2.25 - 0.1) < 3.14159) {
                                                        switch ("" + "foo bar") {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi20 & (fi22 - (fi22 % (((fi22 >> ((803) & 63))) | 1))))
                                                            }
                                                        }
                                                        v3 -= (v3 + -33)
                                                        switch 32768 {
                                                        case -673:
                                                            v36 := GenS3{F0: 255, F1: 1}
                                                            if (42 != 0) {
                                                                v27 = v27
                                                                if v27 {
                                                                    v3 = (v3 << ((145) & 63))
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 0
                                                                    }
                                                                }
                                                            } else {
                                                                var v37 uint16 = 1
                                                                v37 -= (v37 | v37)
                                                                var v38 uint16 = (v37 >> ((1) & 63))
                                                                v27 = v27
                                                                var v39 int16 = -416
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                fmt.Println(fi20)
                                                            case "a":
                                                                var v40 float64 = -2.25
                                                                _ = v40
                                                            case "hello":
                                                                if false {
                                                                    v1 = ("\n\t" == "\n\t")
                                                                    v35 = v35
                                                                    v41 := GenS4{F0: fi22, F1: 2147483647}
                                                                    _ = v41
                                                                }
                                                            }
                                                            _ = v36
                                                        }
                                                        var v42 bool = ((!(((v1 || v27) && (("" < ("" + "foo bar")) || ((v35 ^ v35) <= (v35 | (^(v35)))))) && (1 >= 4294967295))) && (!true))
                                                        _ = v42
                                                    } else {
                                                        fmt.Println(fi20, v3, v1)
                                                        var v43 float64 = 1e10
                                                        v44 := GenS5{F0: 1}
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    v2 = append(v2, (fi20 / ((42) | 1)))
                                                    if len(v2) > 0 {
                                                        v2[0] = fi22
                                                    }
                                                }
                                                _ = v35
                                            }
                                            if (2 != 127) {
                                                v3 -= (-(v3))
                                            } else {
                                                switch (fi22 << ((genFn0(18446744073709551615, -2147483648)) & 63)) {
                                                case 4294967295:
                                                    v45 := GenS6{F0: 18446744073709551615}
                                                    _ = v45
                                                case -9223372036854775808:
                                                    v2 = append(v2, fi22)
                                                case 103:
                                                    v27 = false
                                                default:
                                                    var v46 uint64 = 2
                                                    _ = v46
                                                }
                                                v2 = append(v2, (fi20 >> ((0) & 63)))
                                                switch fi22 {
                                                case 539:
                                                    if v1 {
                                                        fmt.Println(fi20, fi22, v3)
                                                        v27 = v1
                                                    } else {
                                                        var v47 uint8 = 175
                                                        _ = v47
                                                    }
                                                case 23:
                                                    if len(v2) > 0 {
                                                        v2[0] = fi22
                                                    }
                                                }
                                            }
                                        }
                                        var v48 int16 = 1000
                                        fmt.Println(fi22, v48)
                                        _ = v48
                                    }
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v49 uint32 = 2
                                    var v50 bool = false
                                    v49 = (v49 / ((v49) | 1))
                                    var v51 uint64 = 2
                                    if (genFn0(uint(v49), 0) <= 4294967295) {
                                        v52 := GenS3{F0: 1, F1: genFn0(2, -1)}
                                        _ = len(v2)
                                        _ = v52
                                    }
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                _ = len(v2)
                            }
                            if v1 {
                                if v1 {
                                    if v1 {
                                        if (!v1) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = len(v2)
                                            fmt.Println(fi20, v3, fi20)
                                        }
                                        var v53 bool = v1
                                        switch "\n\t" {
                                        case "":
                                            v54 := GenS1{F0: 0}
                                            _ = v54
                                        case "a":
                                            var v55 uint8 = 0
                                            _ = v55
                                        default:
                                            if (false || (127 >= 0)) {
                                                v56 := GenS4{F0: fi22, F1: 197}
                                                _ = v56
                                            }
                                        }
                                        _ = v53
                                    }
                                    v57 := GenS2{F0: 42, F1: 255}
                                    _ = v57
                                }
                                var v58 int16 = -251
                                _ = v58
                            }
                            v3 -= (v3 / ((v3) | 1))
                            if v1 {
                                switch fi20 {
                                case -2147483648:
                                    v1 = (float32(-2.25) == -2.25)
                                case 4294967295:
                                    if (!v1) {
                                        v1 = v1
                                        if (!false) {
                                            v59 := GenS7{F0: 127, F1: 349, F2: 255}
                                            v1 = ((42 >= -721) && true)
                                            _ = v59
                                        } else {
                                            v3 = -10
                                        }
                                    }
                                }
                                v1 = (!(("a" == "hello") && (v1 && (361 < 717))))
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (fi22 / ((32767) | 1))
                                }
                                if v1 {
                                    v3 = v3
                                    if v1 {
                                        switch -2147483648 {
                                        case -2147483648:
                                            fmt.Println(v3)
                                        case -9223372036854775808:
                                            v3 += (^(v3))
                                        }
                                        v3 += (^(v3))
                                    } else {
                                        if (-9223372036854775808 > 9223372036854775807) {
                                            v1 = (!(18446744073709551615 >= 1))
                                            fmt.Println(fi22, fi20)
                                            var v60 uint8 = 1
                                            _ = v60
                                        }
                                    }
                                } else {
                                    fmt.Println(fi22, v3)
                                }
                                var v61 int32 = 0
                                _ = v61
                            }
                            v2 = append(v2, fi22)
                            var v62 string = "hello"
                            _ = v62
                        }
                        v3 = 0
                        var v63 bool = v1
                        v63 = v63
                        v64 := GenS8{F0: 0, F1: 558, F2: (128 > 0)}
                        _ = fi22
                        _ = v63
                        _ = v64
                    }
                default:
                    switch (fi20 + fi20) {
                    case 32768:
                        fmt.Println(v1)
                    case 32767:
                        if v1 {
                            if (32767 < 160) {
                                v65 := GenS8{F0: 255, F1: 4294967295, F2: v1}
                                var v66 int16 = 0
                                var v67 int16 = (v66 - (v66 / ((42) | 1)))
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            fmt.Println(fi20)
                        } else {
                            for fi68 := 0; fi68 < 12; fi68++ {
                                v3 -= (v3 >> ((genFn0(1, -629)) & 63))
                                _ = fi68
                            }
                            _ = len(v2)
                        }
                    }
                }
                _ = fi20
            }
            v3 = 0
        }
    case 128:
        var v69 bool = (9223372036854775807 < -296)
        _ = v69
    default:
        if (((0.0 + (0.0 * -2.25)) * 1e10) == (0.1 + 0.1)) {
            v3 = -1
        } else {
            for fi70 := 0; fi70 < 4; fi70++ {
                if (2147483647 >= 0) {
                    var v71 int64 = 0
                    var v72 int16 = -1
                    _ = v71
                    _ = v72
                } else {
                    var v73 uint64 = 0
                    _ = v73
                }
                _ = len(v2)
                v2 = append(v2, 32768)
                _ = fi70
            }
        }
    }
    v1 = ((true || ((333 != 106) && v1)) || ((1 < 4294967295) && (!v1)))
    v1 = ((false || v1) && v1)
    v3 += v3
    v2 = append(v2, 9223372036854775807)
    if (1 < 65535) {
        var v74 uint = 685
        _ = v74
    } else {
        var v75 int = 9223372036854775807
        v75 -= v75
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = v75
    }
    for fi76 := 0; fi76 < 11; fi76++ {
        fmt.Println(fi76, v1)
        switch "foo bar" {
        case "a":
            switch (fi76 - (^(fi76))) {
            case 1:
                fmt.Println(fi76)
            case -9223372036854775808:
                if v1 {
                    var v77 int8 = v3
                    v3 += v3
                    _ = v77
                }
            case 127:
                v78 := GenS9{F0: 990, F1: (v3 + (v3 & (v3 & v3))), F2: 32767}
                _ = v78
            default:
                var v79 int16 = 887
                _ = v79
            }
        }
        v1 = (!(((v1 || false) && (("hello" >= "\n\t") && true)) || (!false)))
        _ = fi76
    }
    v80 := GenS1{F0: 128}
    for fi81 := 0; fi81 < 16; fi81++ {
        if v1 {
            v3 = 42
            var v82 uint = 4294967295
            _ = v82
        } else {
            v83 := GenS3{F0: 1, F1: 1}
            _ = v83
        }
        _ = fi81
    }
    v1 = (!v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v80
}
