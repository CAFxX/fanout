package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint32; F1 uint; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int; F2 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 int64; F2 int64 }
type GenS11 struct { F0 uint }
func genFn0(p0 int16, p1 bool, p2 int) bool {
    v6 := []int{32767, p2, (p2 * p2)}
    p0 = 63
    if (!true) {
        _ = len(v6)
    } else {
        p0 = (p0 | (p0 >> ((1) & 63)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 uint64 = 825
    var v2 uint = 1
    v3 := []uint32{2147483648, 1}
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    var v7 bool = true
    v2 = v2
    var v8 string = ("hello" + "")
    switch 4294967295 {
    case 32767:
        v7 = (v8 == v8)
    case -2147483648:
        _ = len(v4)
    case 42:
        delete(v4, 2147483647)
    }
    v1 = 1
    switch 127 {
    case 0:
        v2 -= (v2 * v2)
    case -9223372036854775808:
        switch (v8 + v8) {
        case "\n\t":
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 int8 = -128
                    switch -9223372036854775808 {
                    case 32768:
                        v1 = 0
                    }
                    var v12 int16 = 548
                    var v13 int = (fi10 * (fi10 / (((fi9 | (fi9 % (((fi9 ^ 9223372036854775807)) | 1)))) | 1)))
                    v11 = (v11 * (v11 / ((v11) | 1)))
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if v7 {
                    switch v8 {
                    case "hello":
                        switch ("hello" + v8) {
                        case "\n\t":
                            if false {
                                var v14 float64 = 0.1
                                _ = len(v3)
                                if (v7 && v7) {
                                    var v15 int = -799
                                    v8 = ""
                                    var v16 int = (v15 & fi9)
                                    _ = v15
                                    _ = v16
                                } else {
                                    for fi17 := 0; fi17 < 11; fi17++ {
                                        v18 := GenS1{F0: v8, F1: 42}
                                        _ = fi17
                                        _ = v18
                                    }
                                    for fi19 := 0; fi19 < 1; fi19++ {
                                        switch v8 {
                                        case "":
                                            switch (fi19 & -300) {
                                            case -891:
                                                if v7 {
                                                    v5 = (v5 + 1)
                                                    var v20 uint8 = 128
                                                    _ = v20
                                                } else {
                                                    if (32767 >= uint16(v1)) {
                                                        if (float32(v14) == (0.1 + 1.5)) {
                                                            var v21 int64 = 1
                                                            var v22 string = (v8 + v8)
                                                            switch fi19 {
                                                            case 42:
                                                                switch fi19 {
                                                                case 9223372036854775807:
                                                                    v23 := GenS2{F0: (v2 / ((42) | 1))}
                                                                    _ = v23
                                                                case -913:
                                                                    fmt.Println(v8, v22, v22)
                                                                case -9223372036854775808:
                                                                    fmt.Println(v22, v22, v7)
                                                                }
                                                            case 32767:
                                                                fmt.Println(v5, v22)
                                                            case 560:
                                                                var v24 uint8 = 42
                                                                _ = v24
                                                            default:
                                                                switch (fi19 - (fi9 & fi19)) {
                                                                case 563:
                                                                    switch (fi9 - (fi9 << ((uint(v1)) & 63))) {
                                                                    case 90:
                                                                        v7 = v7
                                                                    default:
                                                                        switch 128 {
                                                                        case -9223372036854775808:
                                                                            var v25 int = -2147483648
                                                                            _ = v25
                                                                        case -241:
                                                                            switch (fi19 ^ 658) {
                                                                            case 0:
                                                                                fmt.Println(v8, fi9)
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v14, fi9)
                                                                            default:
                                                                                v7 = (-630 > 2147483647)
                                                                            }
                                                                        }
                                                                    }
                                                                case -349:
                                                                    if (-1 < 1) {
                                                                        v2 += 255
                                                                    }
                                                                default:
                                                                    var v26 int8 = 44
                                                                    _ = v26
                                                                }
                                                            }
                                                            _ = v21
                                                            _ = v22
                                                        }
                                                    } else {
                                                        var v27 string = (v8 + v8)
                                                        v28 := GenS3{F0: 1, F1: uint64(v2), F2: 1}
                                                        fmt.Println(v27)
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    if false {
                                                        if (false || false) {
                                                            if (0 > 42) {
                                                                v4[65535] = v8
                                                                v29 := v4[1]
                                                                _ = v29
                                                            } else {
                                                                if v7 {
                                                                    v30 := GenS3{F0: -128, F1: v1, F2: 0}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2147483648
                                                                    }
                                                                    _ = v30
                                                                } else {
                                                                    var v31 bool = v7
                                                                    _ = v31
                                                                }
                                                            }
                                                            v32 := GenS4{F0: 32767, F1: 4294967295}
                                                            _ = v32
                                                        } else {
                                                            v33 := GenS5{F0: 358, F1: (v2 / ((v2) | 1)), F2: (fi9 % ((fi9) | 1))}
                                                            var v34 int64 = (v5 ^ v5)
                                                            var v35 int64 = (v34 | (-(v5)))
                                                            v34 += (v5 | -1)
                                                            v36 := GenS1{F0: v8, F1: -1}
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        fmt.Println(v8, v5, fi19)
                                                        v5 -= v5
                                                    } else {
                                                        v7 = (v2 != (v2 / (((v2 % (((^(v2))) | 1))) | 1)))
                                                    }
                                                    if (0 == 131) {
                                                        switch (fi9 % ((fi9) | 1)) {
                                                        case 151:
                                                            var v37 bool = v7
                                                            _ = v37
                                                        }
                                                        var v38 uint32 = 4294967295
                                                        _ = v38
                                                    }
                                                }
                                            }
                                            var v39 uint16 = 355
                                            var v40 string = v8
                                            _ = v39
                                            _ = v40
                                        case "hello":
                                            var v41 bool = (!v7)
                                            _ = v41
                                        case "\n\t":
                                            switch ("hello" + v8) {
                                            case "\n\t":
                                                v1 += 18446744073709551615
                                            default:
                                                switch v8 {
                                                case "\n\t":
                                                    fmt.Println(v2, v1)
                                                case "":
                                                    if (!(!false)) {
                                                        v42 := GenS6{F0: 171}
                                                        switch fi19 {
                                                        case 32768:
                                                            var v43 uint = v2
                                                            _ = v43
                                                        default:
                                                            var v44 uint16 = 42
                                                            _ = v44
                                                        }
                                                        if v7 {
                                                            v8 = (v8 + "a")
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = 578
                                                            }
                                                            var v45 float32 = 1e10
                                                            _ = v45
                                                        }
                                                        v46 := GenS7{F0: (v2 ^ 0), F1: (fi9 / ((fi19) | 1)), F2: (v2 >= v2)}
                                                        _ = v42
                                                        _ = v46
                                                    } else {
                                                        v47 := GenS6{F0: 592}
                                                        _ = v47
                                                    }
                                                case "hello":
                                                    v8 = v8
                                                }
                                            }
                                            switch (v8 + (v8 + v8)) {
                                            case "\n\t":
                                                v48 := GenS7{F0: v2, F1: fi9, F2: (!((v5 < 1) && false))}
                                                _ = v48
                                            case "hello":
                                                v49 := GenS2{F0: (v2 * v2)}
                                                _ = v49
                                            default:
                                                if v7 {
                                                    fmt.Println(v1)
                                                    var v50 uint64 = 82
                                                    var v51 int32 = 62
                                                    _ = v50
                                                    _ = v51
                                                } else {
                                                    v1 -= v1
                                                }
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                v52 := GenS6{F0: 1000}
                                                _ = v52
                                            default:
                                                if (-1 < (fi9 + fi9)) {
                                                    if true {
                                                        fmt.Println(v7, v1)
                                                        v53 := GenS1{F0: "\n\t", F1: int32(v1)}
                                                        _ = v53
                                                    }
                                                    fmt.Println(v5, v14, v7)
                                                }
                                            }
                                        default:
                                            var v54 int8 = 100
                                            _ = v54
                                        }
                                        v55 := GenS6{F0: -663}
                                        v56 := GenS7{F0: (v2 & (^(v2))), F1: 0, F2: v7}
                                        v7 = (-2.25 > (1e10 - 3.14159))
                                        _ = fi19
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                _ = v14
                            }
                        case "a":
                            v8 = ("foo bar" + v8)
                        case "hello":
                            switch (^(fi9)) {
                            case 32768:
                                var v57 float64 = 0.0
                                _ = v57
                            case 128:
                                var v58 uint64 = v1
                                _ = v58
                            case 42:
                                var v59 int8 = 1
                                _ = v59
                            }
                        default:
                            var v60 uint = v2
                            _ = v60
                        }
                    case "a":
                        var v61 int16 = -323
                        _ = v61
                    case "":
                        v1 += 1
                    }
                    v7 = v7
                    for fi62 := 0; fi62 < 1; fi62++ {
                        var v63 int = (^(fi62))
                        v7 = (-2.25 != 1.5)
                        v7 = true
                        _ = fi62
                        _ = v63
                    }
                } else {
                    var v64 uint16 = 0
                    if (65535 < 2) {
                        var v65 uint8 = 0
                        switch 4294967295 {
                        case 1:
                            for fi66 := 0; fi66 < 20; fi66++ {
                                v5 = (v5 - (v5 & -617))
                                switch v8 {
                                case "a":
                                    var v67 int32 = -82
                                    _ = v67
                                case "foo bar":
                                    fmt.Println(fi9, fi9)
                                case "":
                                    v68 := GenS8{F0: v2}
                                    _ = v68
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                }
                                var v69 uint8 = v65
                                _ = fi66
                                _ = v69
                            }
                        default:
                            _ = len(v3)
                        }
                        if (-2147483648 > 2147483647) {
                            v70 := GenS9{F0: "hello"}
                            v71 := GenS3{F0: 46, F1: (^(v1)), F2: 715}
                            _ = v70
                            _ = v71
                        } else {
                            v5 += v5
                        }
                        _ = v65
                    } else {
                        v1 = (^(v1))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v7 = v7
                    }
                    _ = v64
                }
                _ = fi9
            }
        case "hello":
            v7 = true
        }
    case 89:
        var v72 int = 128
        _ = v72
    }
    v73 := GenS5{F0: 2147483648, F1: v2, F2: 966}
    v2 += 42
    var v74 uint64 = (^(v1))
    var v75 int16 = 1
    var v76 uint64 = (v74 | 18446744073709551615)
    switch v8 {
    case "a":
        fmt.Println(v7, v74, v1)
    case "\n\t":
        if (false || (!(734 <= 255))) {
            v77 := GenS2{F0: v2}
            v78 := GenS10{F0: 42, F1: (v5 ^ (v5 ^ 326)), F2: -1}
            v79 := GenS11{F0: v2}
            _ = v77
            _ = v78
            _ = v79
        }
    case "hello":
        v1 += 0
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v7 = v7
    for fi80 := 0; fi80 < 14; fi80++ {
        switch int(v5) {
        case -743:
            v74 -= (v74 + v1)
        case -116:
            fmt.Println(v8)
        default:
            for fi81 := 0; fi81 < 8; fi81++ {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                v82 := GenS6{F0: (v75 >> ((v2) & 63))}
                _ = fi81
                _ = v82
            }
        }
        _ = fi80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v8
}
