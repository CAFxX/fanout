package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int64; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint8 }
func genFn0() bool {
    var v7 float32 = float32(((3.14159 + float64(0.1)) + 0.0))
    var v8 uint8 = 2
    var v9 uint = 661
    if false {
        var v10 bool = (841 == 2147483647)
        v8 -= (v8 - (^(v8)))
        switch -23 {
        case 1:
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 32767}
                _ = fi11
                _ = v12
            }
        case 196:
            var v13 uint8 = (v8 % (((v8 & 203)) | 1))
            _ = v13
        default:
            for fi14 := 0; fi14 < 14; fi14++ {
                if (4294967295 != 42) {
                    v15 := GenS1{F0: 81}
                    _ = v15
                }
                v16 := []int{fi14, fi14, 254, fi14}
                for fi17 := 0; fi17 < 2; fi17++ {
                    _ = len(v16)
                    if len(v16) > 0 {
                        v16[0] = fi17
                    }
                    if false {
                        v18 := GenS2{F0: 42, F1: 32767, F2: -128}
                        switch "" {
                        case "hello":
                            _ = len(v16)
                        }
                        v9 += 535
                        _ = v18
                    }
                    _ = fi17
                }
                _ = fi14
                _ = v16
            }
        }
        _ = v10
    } else {
        var v19 bool = (int(v9) > 559)
        v20 := GenS3{F0: -264, F1: 42}
        v21 := GenS3{F0: -196, F1: 1}
        _ = v19
        _ = v20
        _ = v21
    }
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []string{("foo bar" + (("" + ("hello" + ("\n\t" + "a"))) + "a")), "foo bar", "hello"}
    var v2 int64 = 42
    var v3 string = "foo bar"
    var v4 int = -582
    var v5 uint32 = 351
    var v6 int32 = 0
    if len(v1) > 0 {
        v1[0] = v3
    }
    v3 = v3
    fmt.Println(v4, v4, v5)
    var v22 int8 = 1
    var v23 string = "\n\t"
    var v24 uint16 = 32767
    v24 += (v24 << ((224) & 63))
    v6 = v6
    if false {
        var v25 int = -9223372036854775808
        switch "" {
        case "foo bar":
            fmt.Println(v22, v25)
        case "a":
            switch (v4 * v25) {
            case -704:
                var v26 float32 = 3.14159
                _ = v26
            }
        case "hello":
            var v27 int = (v25 & 658)
            _ = v27
        }
        v25 -= (v4 >> ((471) & 63))
        _ = v25
    }
    if (!((1e10 * ((-2.25 + ((1e10 * 0.0) + 1.5)) / (float32(float64(0.0)) - 1e10))) > -2.25)) {
        var v28 uint = 4294967295
        _ = v28
    } else {
        switch v23 {
        case "foo bar":
            v4 += (v4 << ((65535) & 63))
        case "hello":
            for fi29 := 0; fi29 < 5; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    _ = len(v1)
                    if (0 > (v6 & (v6 % ((v6) | 1)))) {
                        if (1.5 == (((float32(1e10) * 3.14159) + float32(((-2.25 - 3.14159) * (0.0 / ((1.5 * 1e10) + 0.0))))) + 0.1)) {
                            v4 -= 32768
                        } else {
                            fmt.Println(v6)
                            if true {
                                v24 += 0
                            } else {
                                switch int(v6) {
                                case 128:
                                    switch ("\n\t" + ((v3 + ((v3 + "\n\t") + ("foo bar" + ""))) + v3)) {
                                    case "\n\t":
                                        var v31 int8 = (v22 - (v22 & v22))
                                        _ = v31
                                    }
                                case 32767:
                                    if len(v1) > 0 {
                                        v1[0] = ("foo bar" + v3)
                                    }
                                }
                                _ = len(v1)
                            }
                            if false {
                                if (1e10 <= -2.25) {
                                    if ((3.14159 - 0.1) <= -2.25) {
                                        if (!true) {
                                            v3 = v23
                                            fmt.Println(v3)
                                            var v32 string = (v23 + (v3 + "foo bar"))
                                            _ = v32
                                        }
                                        var v33 int64 = (v2 & (v2 % ((v2) | 1)))
                                        _ = v33
                                    } else {
                                        if (!(false || true)) {
                                            v34 := GenS4{F0: v2, F1: ((v23 + v3) + ("" + v3)), F2: (^(v22))}
                                            v23 = "hello"
                                            fmt.Println(v24, fi30)
                                            _ = v34
                                        } else {
                                            v35 := GenS2{F0: v6, F1: v24, F2: (v22 % ((v22) | 1))}
                                            _ = v35
                                        }
                                        fmt.Println(v4, v6)
                                    }
                                    var v36 bool = (false || (((((0.0 / (1.5 + (0.1 / 1.5))) / -2.25) + 1.5) >= (((((3.14159 * float64(float32(0.0))) + 3.14159) + float64(0.0)) / 1e10) * 0.0)) || false))
                                    var v37 uint = 2
                                    v6 = v6
                                    _ = v36
                                    _ = v37
                                }
                                if (!false) {
                                    fmt.Println(v24)
                                    var v38 int32 = (v6 * 1)
                                    _ = v38
                                } else {
                                    var v39 float64 = 0.1
                                    var v40 uint64 = 4294967295
                                    v4 = 32767
                                    _ = v39
                                    _ = v40
                                }
                                fmt.Println(v2, v6, v3)
                            }
                            if true {
                                var v41 bool = (!false)
                                _ = v41
                            } else {
                                if true {
                                    fmt.Println(v6, v2)
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = v23
                                        }
                                    }
                                    if false {
                                        var v42 uint8 = 2
                                        _ = v42
                                    }
                                } else {
                                    _ = len(v1)
                                    _ = len(v1)
                                    v43 := GenS2{F0: v6, F1: v24, F2: (-(v22))}
                                    var v44 float32 = 0.0
                                    _ = v43
                                    _ = v44
                                }
                                var v45 uint64 = 18446744073709551615
                                _ = v45
                            }
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v22, v6, fi29)
                                case "a":
                                    var v46 uint16 = (^(v24))
                                    _ = v46
                                default:
                                    if (((fi29 >> ((4294967295) & 63)) < -4) && true) {
                                        if true {
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    if (false && (false || ((-439 == -1) && (true && (float64(1.5) != (3.14159 - ((3.14159 * 0.0) - -2.25))))))) {
                                                        if len(v1) > 0 {
                                                            v1[0] = v23
                                                        }
                                                        var v47 int32 = (v6 ^ 42)
                                                        if len(v1) > 0 {
                                                            v1[0] = "a"
                                                        }
                                                        _ = v47
                                                    } else {
                                                        if true {
                                                            switch 2147483647 {
                                                            case -131:
                                                                v5 = v5
                                                            default:
                                                                var v48 bool = false
                                                                _ = v48
                                                            }
                                                            switch fi29 {
                                                            case -9223372036854775808:
                                                                switch (v4 & (v4 + (fi29 ^ 32767))) {
                                                                case 42:
                                                                    var v49 uint8 = 42
                                                                    _ = v49
                                                                case -1:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                default:
                                                                    v1 = append(v1, v3)
                                                                }
                                                            case 128:
                                                                var v50 bool = (!true)
                                                                _ = v50
                                                            default:
                                                                v51 := GenS3{F0: v6, F1: 2}
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v51
                                                            }
                                                        }
                                                    }
                                                case "foo bar":
                                                    v1 = append(v1, v23)
                                                }
                                                fmt.Println(fi29)
                                                switch (fi30 << ((65535) & 63)) {
                                                case -504:
                                                    var v52 uint32 = (v5 | 2)
                                                    _ = v52
                                                default:
                                                    if false {
                                                        switch fi29 {
                                                        case 936:
                                                            v53 := GenS2{F0: (v6 + int32(v5)), F1: (v24 >> ((36) & 63)), F2: v22}
                                                            _ = v53
                                                        case 667:
                                                            v22 += v22
                                                        case -1:
                                                            v4 = fi29
                                                        default:
                                                            switch v3 {
                                                            case "a":
                                                                v4 += fi29
                                                            case "":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v54 string = "\n\t"
                                                                    _ = v54
                                                                }
                                                            case "hello":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            var v55 string = (("" + v23) + v23)
                                                            v55 = (v55 + v55)
                                                            _ = v55
                                                        }
                                                    }
                                                }
                                            }
                                            _ = len(v1)
                                        } else {
                                            if (v2 >= v2) {
                                                fmt.Println(v23, v4)
                                                fmt.Println(fi30, v2, fi29)
                                                var v56 uint32 = v5
                                                _ = v56
                                            }
                                            var v57 uint32 = uint32(v24)
                                            var v58 float64 = (0.1 / 0.1)
                                            _ = v57
                                            _ = v58
                                        }
                                        v1 = append(v1, (("hello" + v3) + (v23 + "")))
                                        if false {
                                            v5 = v5
                                        } else {
                                            v3 = v3
                                            switch v23 {
                                            case "hello":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                    }
                                }
                                v1 = append(v1, "\n\t")
                                v6 = -1
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        if false {
                            v24 += v24
                            v59 := GenS1{F0: -32768}
                            _ = v59
                        } else {
                            var v60 bool = (true && (!true))
                            if v60 {
                                _ = len(v1)
                                var v61 bool = v60
                                var v62 uint64 = 18446744073709551615
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v61
                                _ = v62
                            } else {
                                var v63 int = 9223372036854775807
                                if v60 {
                                    var v64 uint64 = 42
                                    var v65 float64 = 0.1
                                    _ = v64
                                    _ = v65
                                } else {
                                    var v66 uint = 2
                                    v22 -= v22
                                    var v67 int = -9223372036854775808
                                    var v68 uint8 = 128
                                    v1 = append(v1, (v3 + (v23 + "a")))
                                    switch fi29 {
                                    case 42:
                                        v1 = append(v1, v3)
                                    case -71:
                                        v67 += v63
                                    }
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                _ = v63
                            }
                            fmt.Println(v3, v2, v2)
                            _ = v60
                        }
                        if false {
                            var v69 int64 = v2
                            var v70 int8 = (v22 / (((v22 >> ((42) & 63))) | 1))
                            if ((v2 > (v2 << ((4294967295) & 63))) && true) {
                                v5 = 65535
                                var v71 uint64 = 42
                                _ = v71
                            }
                            v72 := GenS5{F0: 555, F1: 5}
                            _ = v69
                            _ = v70
                            _ = v72
                        } else {
                            if len(v1) > 0 {
                                v1[0] = (v23 + ((v3 + v23) + ("" + ("\n\t" + (v3 + "hello")))))
                            }
                            fmt.Println(v5, fi29, fi29)
                        }
                        v22 -= (v22 / ((100) | 1))
                    }
                    switch 1 {
                    case -9223372036854775808:
                        v2 += int64(v6)
                    case -37:
                        if (!(!true)) {
                            v73 := GenS5{F0: (v6 * v6), F1: uint8(v24)}
                            fmt.Println(v5)
                            _ = v73
                        }
                    case -2147483648:
                        v1 = append(v1, (v3 + v23))
                    default:
                        if false {
                            var v74 float32 = 1e10
                            _ = v74
                        } else {
                            if false {
                                fmt.Println(v5, v2)
                                v2 = -718
                                v75 := GenS3{F0: 2147483647, F1: 135}
                                v76 := GenS3{F0: 477, F1: 769}
                                if true {
                                    v23 = "hello"
                                } else {
                                    v77 := GenS3{F0: (-(v6)), F1: 611}
                                    if false {
                                        if true {
                                            switch (fi30 / (((fi30 * 127)) | 1)) {
                                            case 58:
                                                v78 := GenS2{F0: (v6 - 2147483647), F1: 2, F2: v22}
                                                _ = v78
                                            case 128:
                                                if len(v1) > 0 {
                                                    v1[0] = v3
                                                }
                                            case -105:
                                                v4 = fi30
                                            default:
                                                v79 := GenS6{F0: v22, F1: (false || false)}
                                                _ = v79
                                            }
                                            fmt.Println(v5)
                                            fmt.Println(v5, v2)
                                        }
                                        fmt.Println(v2, v24, fi30)
                                        if false {
                                            v23 = ""
                                            v4 += fi29
                                        } else {
                                            var v80 uint8 = 122
                                            var v81 int = (-(fi30))
                                            _ = v80
                                            _ = v81
                                        }
                                    } else {
                                        if (0.0 >= float64(3.14159)) {
                                            v6 += (v6 - v6)
                                            var v82 uint16 = 30
                                            v4 = v4
                                            _ = v82
                                        } else {
                                            v24 += v24
                                            if (v5 <= 1) {
                                                var v83 float64 = (float64((1e10 - (float32(((-2.25 * 3.14159) / -2.25)) + (0.0 * float32(0.1))))) - (3.14159 * 3.14159))
                                                var v84 int = fi29
                                                _ = v83
                                                _ = v84
                                            }
                                            if false {
                                                var v85 uint8 = 42
                                                v5 += 42
                                                _ = v85
                                            }
                                        }
                                        var v86 float32 = (0.0 / 1.0)
                                        _ = v86
                                    }
                                    _ = v77
                                }
                                _ = v75
                                _ = v76
                            } else {
                                v3 = ("\n\t" + (((("" + v23) + v3) + (v23 + (v3 + v3))) + v3))
                            }
                            v87 := GenS7{F0: 19}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v87
                        }
                    }
                    _ = fi30
                }
                for fi88 := 0; fi88 < 4; fi88++ {
                    v22 = (-(v22))
                    v23 = v23
                    _ = fi88
                }
                v4 += (-(v4))
                _ = fi29
            }
        }
        v24 -= v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
