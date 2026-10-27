package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0() string {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 57, F1: ""}
        var v8 uint64 = 958
        var v9 float64 = 1e10
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if ((673 > 2) && ((1 >= 0) && (!false))) {
        v10 := GenS1{F0: 0, F1: ""}
        _ = v10
    }
    var v11 uint8 = 255
    var v12 int16 = -1
    switch ("foo bar" + "") {
    case "foo bar":
        v13 := GenS1{F0: (v11 % ((v11) | 1)), F1: (("foo bar" + ("a" + "hello")) + "a")}
        _ = v13
    case "a":
        v14 := GenS2{F0: 713}
        _ = v14
    case "\n\t":
        switch "hello" {
        case "hello":
            v15 := []int{-738, -2147483648, -408}
            _ = v15
        case "":
            if false {
                v12 -= 1000
                if true {
                    v12 += (v12 ^ v12)
                }
                v16 := []int{}
                _ = v16
            }
        case "\n\t":
            fmt.Println(v11, v11)
        }
    default:
        v17 := []int{0, 2147483647, -2147483648}
        _ = v17
    }
    _ = v11
    _ = v12
    return (("a" + (("a" + "hello") + "foo bar")) + "")
}
func genFn1() uint64 {
    switch 5 {
    case 42:
        var v18 uint = 0
        _ = v18
    case 1:
        if (((-1 > -9) || true) || true) {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 float64 = 0.1
                _ = fi19
                _ = v20
            }
        }
    }
    v21 := []int{488}
    var v22 string = ""
    switch v22 {
    case "a":
        v22 = v22
    case "":
        if (0.1 <= (((1.5 / ((0.0 * 0.0) - (3.14159 - 1.5))) * (1e10 + 1e10)) + (float64(1.5) * 1.5))) {
            var v23 float64 = 0.1
            _ = len(v21)
            if (24 > 0) {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v23 = v23
                    switch fi24 {
                    case -480:
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    case -575:
                        for fi25 := 0; fi25 < 2; fi25++ {
                            v22 = v22
                            var v26 float32 = float32((((((1.5 / v23) * ((v23 - ((((v23 * v23) * v23) + v23) * v23)) * v23)) - v23) + 0.1) + v23))
                            var v27 string = ((((v22 + (("" + v22) + (v22 + v22))) + v22) + v22) + v22)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    case -2147483648:
                        v21 = append(v21, (fi24 >> ((0) & 63)))
                    }
                    _ = fi24
                }
                var v28 int8 = -117
                _ = v28
            } else {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
            _ = v23
        }
    case "\n\t":
        v29 := GenS3{F0: 652, F1: -1, F2: 42}
        _ = v29
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            if len(v21) > 0 {
                v21[0] = fi30
            }
            _ = fi30
        }
    }
    if true {
        v22 = v22
    } else {
        if false {
            if len(v21) > 0 {
                v21[0] = -253
            }
        } else {
            v31 := GenS2{F0: 690}
            v21 = append(v21, 127)
            _ = len(v21)
            _ = v31
        }
    }
    _ = v21
    _ = v22
    return 604
}

func main() {
    v1 := []string{"a", "foo bar", (("a" + "foo bar") + (("foo bar" + ("" + (("\n\t" + (((("\n\t" + ("\n\t" + ("" + ("a" + "")))) + ("foo bar" + "")) + ("foo bar" + ("\n\t" + "foo bar"))) + "\n\t")) + (((("" + "") + ("a" + ("a" + ("\n\t" + "\n\t")))) + ("\n\t" + "foo bar")) + "hello")))) + ("\n\t" + (("foo bar" + "hello") + "\n\t"))))}
    v2 := []uint32{222, 147, 2147483648}
    var v3 uint = 65535
    var v4 uint32 = 2
    var v5 int8 = -128
    if false {
        v32 := GenS2{F0: v4}
        for fi33 := 0; fi33 < 11; fi33++ {
            v1 = append(v1, "foo bar")
            fmt.Println(v4, v5, v4)
            _ = fi33
        }
        switch "\n\t" {
        case "a":
            var v34 bool = false
            _ = v34
        }
        _ = v32
    } else {
        var v35 uint32 = 2147483647
        v36 := GenS3{F0: 840, F1: -811, F2: 18446744073709551615}
        v5 += 42
        _ = v35
        _ = v36
    }
    var v37 bool = (float64((((3.14159 * 3.14159) / (1e10 / (0.1 - 0.0))) / 1.5)) > 1e10)
    for fi38 := 0; fi38 < 1; fi38++ {
        v4 = 921
        var v39 int8 = v5
        _ = fi38
        _ = v39
    }
    var v40 int64 = 0
    var v41 int32 = -809
    switch 4294967295 {
    case 357:
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS4{F0: -747}
            v2 = append(v2, (^(v4)))
            v2 = append(v2, v4)
            _ = fi42
            _ = v43
        }
    default:
        if v37 {
            switch 128 {
            case 127:
                for fi44 := 0; fi44 < 20; fi44++ {
                    fmt.Println(v37)
                    switch (fi44 + -1) {
                    case 128:
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v46 := GenS1{F0: 2, F1: "\n\t"}
                            _ = fi45
                            _ = v46
                        }
                    case 1:
                        for fi47 := 0; fi47 < 11; fi47++ {
                            if v37 {
                                if (v5 > (v5 >> ((v3) & 63))) {
                                    v48 := GenS5{F0: (-(v40)), F1: (-(v40))}
                                    var v49 float64 = -2.25
                                    v2 = append(v2, v4)
                                    _ = v48
                                    _ = v49
                                } else {
                                    if (v40 < (v40 * v40)) {
                                        if true {
                                            v3 = v3
                                            v40 += v40
                                        } else {
                                            v1 = append(v1, "a")
                                            var v50 int32 = (v41 / ((v41) | 1))
                                            _ = v50
                                        }
                                    }
                                }
                                var v51 string = ""
                                _ = v51
                            }
                            _ = fi47
                        }
                    case 32767:
                        if v37 {
                            var v52 int16 = 0
                            for fi53 := 0; fi53 < 19; fi53++ {
                                var v54 int64 = (v40 ^ 1)
                                fmt.Println(v40)
                                v52 += v52
                                var v55 uint8 = uint8(v3)
                                if (v37 && v37) {
                                    if len(v1) > 0 {
                                        v1[0] = (("hello" + ("hello" + "foo bar")) + "foo bar")
                                    }
                                    v40 = (-(v40))
                                    _ = len(v1)
                                } else {
                                    v41 = (v41 - 2147483647)
                                    v54 = (v40 / (((v54 ^ (v40 % (((-(v40))) | 1)))) | 1))
                                    v56 := GenS2{F0: v4}
                                    _ = v56
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v1 = append(v1, "hello")
                                switch "\n\t" {
                                case "foo bar":
                                    v57 := GenS5{F0: (-(v40)), F1: (v40 % ((-1) | 1))}
                                    _ = v57
                                case "hello":
                                    v58 := GenS1{F0: v55, F1: "hello"}
                                    _ = v58
                                case "\n\t":
                                    v59 := GenS1{F0: 128, F1: "\n\t"}
                                    _ = v59
                                }
                                switch "" {
                                case "\n\t":
                                    v52 -= (v52 / ((int16(v54)) | 1))
                                }
                                v5 += v5
                                var v60 int64 = int64(v55)
                                _ = len(v2)
                                switch (fi53 & (fi44 % (((^(fi44))) | 1))) {
                                case 4294967295:
                                    var v61 int16 = (^(v52))
                                    var v62 string = "\n\t"
                                    if true {
                                        if ((((fi53 >= 1) || v37) && v37) || false) {
                                            fmt.Println(v41, v37)
                                            _ = len(v2)
                                        } else {
                                            v52 = (v61 - 492)
                                            var v63 float32 = 0.0
                                            _ = v63
                                        }
                                    }
                                    _ = v61
                                    _ = v62
                                case 1:
                                    var v64 int = -2147483648
                                    _ = v64
                                default:
                                    if (-2.25 <= (0.0 + 0.1)) {
                                        switch "a" {
                                        case "\n\t":
                                            var v65 uint8 = (v55 & 202)
                                            _ = v65
                                        case "":
                                            v66 := GenS4{F0: 42}
                                            _ = v66
                                        case "foo bar":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        var v67 int8 = -128
                                        v3 = v3
                                        v4 += v4
                                        v68 := GenS1{F0: (v55 ^ 2), F1: ("hello" + ("foo bar" + ("" + "\n\t")))}
                                        _ = v67
                                        _ = v68
                                    }
                                }
                                var v69 string = "hello"
                                _ = fi53
                                _ = v54
                                _ = v55
                                _ = v60
                                _ = v69
                            }
                            v40 -= -800
                            _ = v52
                        } else {
                            fmt.Println(fi44, v41)
                            v1 = append(v1, "hello")
                            v4 = 65535
                        }
                    }
                    for fi70 := 0; fi70 < 5; fi70++ {
                        v71 := GenS4{F0: -756}
                        var v72 bool = (!(!(v40 == -226)))
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                    _ = fi44
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v73 bool = false
            _ = v73
        }
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        v75 := GenS5{F0: (^(v40)), F1: 9223372036854775807}
        fmt.Println(v3, v40, v3)
        _ = fi74
        _ = v75
    }
    v76 := GenS1{F0: 0, F1: ("\n\t" + "foo bar")}
    fmt.Println(v3)
    if (-2.25 >= (0.1 + ((-2.25 / 1.0) * 0.0))) {
        v3 = v3
        v37 = false
    } else {
        switch -2147483648 {
        case -507:
            var v77 string = "foo bar"
            _ = v77
        default:
            for fi78 := 0; fi78 < 15; fi78++ {
                v37 = (!v37)
                var v79 uint8 = 1
                _ = fi78
                _ = v79
            }
        }
        v41 = 2147483647
    }
    if (v37 && ((v41 | 2147483647) > -1)) {
        if true {
            var v80 int64 = v40
            v2 = append(v2, 0)
            fmt.Println(v41, v37)
            _ = v80
        }
        fmt.Println(v3, v5, v40)
    } else {
        v37 = v37
        v4 += 1
        v5 = 19
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println(v37, v37, v4)
    v4 = (v4 << (((v3 % (((v3 % ((255) | 1))) | 1))) & 63))
    v81 := GenS3{F0: 255, F1: v40, F2: 210}
    var v82 int8 = v5
    v3 -= 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v76
    _ = v81
    _ = v82
}
