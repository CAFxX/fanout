package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS7 struct { F0 int; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 uint; F1 uint8 }
func genFn0(p0 int8, p1 int16) int32 {
    if false {
        var v7 uint64 = 943
        _ = v7
    } else {
        v8 := []int{128, 4294967295}
        var v9 string = "hello"
        _ = v8
        _ = v9
    }
    if true {
        var v10 uint64 = 2
        v11 := []int{9223372036854775807, 1}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 bool, p1 int16) int32 {
    switch (("" + ("" + "")) + "foo bar") {
    case "a":
        p1 -= p1
    case "\n\t":
        v12 := []int{0, 0, 290}
        _ = v12
    default:
        var v13 int8 = 42
        _ = v13
    }
    var v14 int64 = -9223372036854775808
    switch -2147483648 {
    case 9223372036854775807:
        var v15 uint = 4294967295
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v14
    return 2147483647
}

func main() {
    var v1 bool = (2147483648 != 2)
    var v2 uint16 = 376
    var v3 int8 = 1
    var v4 int64 = -543
    var v5 float32 = 1e10
    var v6 float64 = 0.1
    var v16 int32 = -644
    if ((-2147483648 >= 4294967295) && false) {
        v4 = (v4 - v4)
    } else {
        if (2147483647 >= 4294967295) {
            v17 := GenS1{F0: v2, F1: (v4 | v4), F2: (2 < 159)}
            var v18 int = -9223372036854775808
            _ = v17
            _ = v18
        } else {
            var v19 uint8 = 63
            switch "" {
            case "a":
                var v20 uint8 = v19
                _ = v20
            default:
                v21 := GenS2{F0: 2147483647, F1: -9223372036854775808, F2: (v4 << ((861) & 63))}
                _ = v21
            }
            _ = v19
        }
    }
    v22 := []int{123, -428, -9223372036854775808}
    if (false || v1) {
        switch -2147483648 {
        case 128:
            if v1 {
                v23 := GenS3{F0: 2}
                _ = v23
            } else {
                var v24 bool = true
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = v24
            }
        case -608:
            v2 = v2
        }
    } else {
        if v1 {
            if (true && (4294967295 >= 933)) {
                v1 = v1
                _ = len(v22)
            } else {
                v5 = v5
                v2 = (^(v2))
            }
            switch 128 {
            case 24:
                v25 := GenS4{F0: (^(v4)), F1: 1}
                _ = v25
            case -2147483648:
                var v26 int8 = 42
                _ = v26
            default:
                var v27 uint16 = (v2 + v2)
                _ = v27
            }
        } else {
            var v28 float64 = 1e10
            v29 := GenS5{F0: v2}
            var v30 int64 = (v4 / ((42) | 1))
            _ = v28
            _ = v29
            _ = v30
        }
        switch 127 {
        case -1:
            v4 = int64(v16)
        case 42:
            switch 32768 {
            case -9223372036854775808:
                fmt.Println(v6, v6)
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    if true {
                        switch fi31 {
                        case 793:
                            for fi32 := 0; fi32 < 10; fi32++ {
                                v2 += (v2 | 0)
                                _ = fi32
                            }
                        case -394:
                            v5 -= (v5 + ((3.14159 * (((((-2.25 * (v5 + v5)) / v5) - v5) - v5) + v5)) - (0.1 / (v5 + (v5 * v5)))))
                        case 556:
                            v33 := GenS1{F0: (v2 ^ 65535), F1: (v4 - int64(v16)), F2: (!(!("\n\t" != ((("hello" + ("\n\t" + ("" + "foo bar"))) + "foo bar") + "\n\t"))))}
                            _ = v33
                        default:
                            var v34 uint32 = 128
                            _ = v34
                        }
                    } else {
                        v22 = append(v22, (fi31 & -9223372036854775808))
                        for fi35 := 0; fi35 < 4; fi35++ {
                            _ = len(v22)
                            v1 = (1e10 > v5)
                            var v36 int64 = (v4 % ((v4) | 1))
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v16, v1, v2)
                    }
                    _ = fi31
                }
            }
        }
        v4 = (v4 | 0)
    }
    var v37 uint = 20
    if v1 {
        v38 := GenS5{F0: v2}
        switch "hello" {
        case "\n\t":
            for fi39 := 0; fi39 < 10; fi39++ {
                v3 += -128
                v22 = append(v22, fi39)
                _ = fi39
            }
        case "hello":
            var v40 uint = v37
            _ = v40
        case "a":
            switch "foo bar" {
            case "a":
                for fi41 := 0; fi41 < 1; fi41++ {
                    if (!((v6 != 0.1) || true)) {
                        v22 = append(v22, 128)
                        var v42 uint16 = 2
                        _ = v42
                    } else {
                        if len(v22) > 0 {
                            v22[0] = -253
                        }
                        v43 := GenS2{F0: 660, F1: fi41, F2: 1}
                        _ = v43
                    }
                    var v44 uint64 = uint64(v37)
                    _ = fi41
                    _ = v44
                }
            case "":
                v16 = genFn1((v2 > (^(v2))), -32768)
            case "hello":
                if (18446744073709551615 > 1) {
                    var v45 uint8 = 128
                    var v46 int = 42
                    if false {
                        if (42 < v2) {
                            if v1 {
                                switch (v46 * v46) {
                                case 2147483647:
                                    for fi47 := 0; fi47 < 15; fi47++ {
                                        if (!v1) {
                                            var v48 float64 = (0.1 * 0.0)
                                            _ = len(v22)
                                            _ = v48
                                        }
                                        _ = fi47
                                    }
                                default:
                                    if v1 {
                                        v3 = v3
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            var v49 int8 = (v3 - 11)
                                            _ = v49
                                        case "hello":
                                            v46 = v46
                                        case "":
                                            var v50 int32 = int32(v3)
                                            _ = v50
                                        default:
                                            var v51 int16 = 1000
                                            _ = v51
                                        }
                                        _ = len(v22)
                                    }
                                }
                                if v1 {
                                    var v52 uint32 = 4294967295
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v45, v37)
                                if ((v46 & int(v4)) == -2147483648) {
                                    v1 = (true && (!(1 != 1)))
                                } else {
                                    if v1 {
                                        if len(v22) > 0 {
                                            v22[0] = 292
                                        }
                                        fmt.Println(v5, v6)
                                        v3 += (v3 << (((^(v37))) & 63))
                                    } else {
                                        var v53 uint64 = 1
                                        if (v1 && ((float32((v6 - 3.14159)) == v5) || (-931 >= 0))) {
                                            switch ("hello" + "hello") {
                                            case "":
                                                for fi54 := 0; fi54 < 2; fi54++ {
                                                    v46 = -497
                                                    if ((true || (!((^(v2)) <= uint16(v45)))) && (v16 > v16)) {
                                                        switch "foo bar" {
                                                        case "":
                                                            v22 = append(v22, fi54)
                                                        case "hello":
                                                            for fi55 := 0; fi55 < 20; fi55++ {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    var v56 int8 = (v3 & -123)
                                                                    _ = v56
                                                                }
                                                                _ = fi55
                                                            }
                                                        default:
                                                            v57 := GenS3{F0: 2147483648}
                                                            _ = v57
                                                        }
                                                        if len(v22) > 0 {
                                                            v22[0] = (fi54 ^ 4294967295)
                                                        }
                                                        var v58 float64 = v6
                                                        _ = v58
                                                    } else {
                                                        v59 := GenS3{F0: 42}
                                                        _ = v59
                                                    }
                                                    var v60 uint = v37
                                                    _ = fi54
                                                    _ = v60
                                                }
                                            default:
                                                var v61 uint8 = v45
                                                _ = v61
                                            }
                                            v53 += v53
                                        }
                                        _ = v53
                                    }
                                    v62 := GenS1{F0: (v2 * v2), F1: 9223372036854775807, F2: v1}
                                    _ = v62
                                }
                            }
                            v63 := GenS5{F0: (^(v2))}
                            _ = v63
                        }
                    }
                    _ = v45
                    _ = v46
                }
            default:
                fmt.Println(v37)
            }
        }
        _ = v38
    } else {
        if ((!(v1 || true)) || (394 != v4)) {
            if (65535 < 2147483648) {
                var v64 int64 = v4
                v65 := GenS5{F0: (v2 / ((uint16(v37)) | 1))}
                v6 = v6
                _ = v64
                _ = v65
            } else {
                switch 32768 {
                case -601:
                    switch -348 {
                    case -1:
                        v5 = ((0.0 * v5) * (0.1 * 0.1))
                    case 42:
                        v16 -= genFn0((v3 % (((v3 << ((v37) & 63))) | 1)), -1)
                    }
                case 127:
                    if v1 {
                        if ((^(v16)) != int32(v2)) {
                            var v66 float32 = 3.14159
                            _ = v66
                        } else {
                            var v67 int = 193
                            v68 := GenS6{F0: 255, F1: v2, F2: (v16 >> (((v37 & v37)) & 63))}
                            _ = v67
                            _ = v68
                        }
                        for fi69 := 0; fi69 < 6; fi69++ {
                            if len(v22) > 0 {
                                v22[0] = (fi69 >> ((v37) & 63))
                            }
                            var v70 uint = (v37 % ((v37) | 1))
                            _ = fi69
                            _ = v70
                        }
                        v22 = append(v22, -756)
                    } else {
                        v22 = append(v22, 32767)
                    }
                case 4294967295:
                    for fi71 := 0; fi71 < 6; fi71++ {
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        v72 := GenS7{F0: fi71, F1: v16, F2: 3}
                        v22 = append(v22, fi71)
                        _ = fi71
                        _ = v72
                    }
                default:
                    if (v1 && v1) {
                        v16 = 0
                        fmt.Println(v5)
                        if v1 {
                            v5 = float32((v6 + v6))
                            var v73 int = 9223372036854775807
                            _ = v73
                        }
                    } else {
                        v16 = (v16 % (((^(v16))) | 1))
                        v2 = v2
                        v74 := GenS4{F0: v4, F1: 661}
                        _ = v74
                    }
                }
            }
            _ = len(v22)
        } else {
            v3 = 84
            v16 = 2147483647
            if v1 {
                v16 = (-(v16))
                var v75 float32 = v5
                _ = v75
            } else {
                v16 = v16
            }
        }
        for fi76 := 0; fi76 < 13; fi76++ {
            v2 -= (v2 ^ v2)
            for fi77 := 0; fi77 < 20; fi77++ {
                v78 := GenS8{F0: 235, F1: 32767}
                _ = fi77
                _ = v78
            }
            _ = fi76
        }
    }
    if false {
        v22 = append(v22, 42)
        v4 += v4
        v1 = (v5 > v5)
    } else {
        v79 := GenS2{F0: 32768, F1: 0, F2: v4}
        var v80 bool = (!v1)
        v81 := GenS9{F0: (^(v37)), F1: 42}
        _ = v79
        _ = v80
        _ = v81
    }
    v1 = (!(199 != 270))
    v37 += 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
