package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 int32; F1 int16; F2 uint }
type GenS11 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16; F1 uint16 }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    if (!false) {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: true, F1: 42, F2: 0}
            v8 := []int{}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 bool = true
        _ = v9
    }
    if false {
        fmt.Println(p0)
        var v10 uint = 4294967295
        _ = v10
    } else {
        v11 := []int{2147483647, (p0 << ((65535) & 63)), p0}
        v12 := GenS1{F0: (!true), F1: 335, F2: 200}
        p0 += p0
        _ = v11
        _ = v12
    }
    var v13 float64 = 0.1
    _ = p0
    _ = v13
    return 100
}
func genFn1(p0 uint16, p1 string, p2 uint16) uint16 {
    v14 := []int{127, -643, 4294967295}
    if true {
        var v15 uint = 768
        var v16 bool = false
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 1, F1: 4294967295, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return p2
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []string{("hello" + ("a" + ("hello" + ("" + "foo bar")))), ("" + (("hello" + ((((("" + "hello") + ("a" + ("foo bar" + "foo bar"))) + ("" + "foo bar")) + ("a" + "a")) + "\n\t")) + ("a" + ""))), "\n\t", ((("\n\t" + ("a" + "foo bar")) + ("\n\t" + "foo bar")) + "foo bar")}
    var v3 string = ""
    var v4 int = 533
    var v5 uint64 = 42
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            switch (v4 << ((952) & 63)) {
            case -9223372036854775808:
                fmt.Println(v4, v4)
            default:
                switch v4 {
                case 4294967295:
                    if (65535 < 522) {
                        if (100 != 18) {
                            v3 = v3
                            v20 := GenS3{F0: 0, F1: (v5 & v5), F2: 2147483647}
                            var v21 uint32 = 2
                            var v22 string = "\n\t"
                            v23 := GenS4{F0: (v1 + (fi19 * (-(fi18)))), F1: 65535, F2: true}
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            v24 := GenS4{F0: fi19, F1: 42, F2: true}
                            switch v3 {
                            case "hello":
                                v3 = (((((v3 + ("a" + v3)) + (((("\n\t" + "\n\t") + ("foo bar" + "\n\t")) + v3) + (v3 + "\n\t"))) + (v3 + (v3 + v3))) + "\n\t") + "a")
                            case "foo bar":
                                v5 = v5
                            }
                            v25 := GenS3{F0: -862, F1: 18446744073709551615, F2: 0}
                            _ = v24
                            _ = v25
                        }
                    }
                case 475:
                    if (197 > 128) {
                        var v26 bool = false
                        fmt.Println(v5)
                        _ = v26
                    } else {
                        v2 = append(v2, v3)
                        var v27 float64 = float64(1.5)
                        if true {
                            v28 := GenS5{F0: 42, F1: (!false)}
                            var v29 int8 = 63
                            _ = v28
                            _ = v29
                        } else {
                            v5 -= (v5 | (v5 ^ (v5 / (((v5 & (v5 * (v5 & (v5 & v5))))) | 1))))
                            v30 := GenS6{F0: 163, F1: 1}
                            _ = v30
                        }
                        _ = v27
                    }
                default:
                    v31 := GenS7{F0: 0}
                    _ = v31
                }
            }
            fmt.Println(v5, fi18, v3)
            v2 = append(v2, "foo bar")
            _ = fi19
        }
        if (!(!true)) {
            if true {
                switch v3 {
                case "a":
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v4, v3, v5)
                    }
                case "\n\t":
                    if false {
                        v32 := GenS8{F0: 869, F1: 1}
                        var v33 int16 = -854
                        var v34 string = "hello"
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                case "foo bar":
                    v1 = fi18
                default:
                    if true {
                        v1 = fi18
                    }
                }
                fmt.Println(v4, v3)
                if true {
                    var v35 string = v3
                    if (619 < 943) {
                        var v36 int16 = -32768
                        _ = v36
                    } else {
                        v37 := GenS3{F0: -940, F1: 42, F2: -1}
                        for fi38 := 0; fi38 < 19; fi38++ {
                            var v39 uint16 = 1
                            fmt.Println(v3)
                            v40 := GenS5{F0: -1, F1: true}
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v37
                    }
                    var v41 float32 = (float32((0.0 * float64(1e10))) * 0.0)
                    _ = v35
                    _ = v41
                } else {
                    switch 4294967295 {
                    case -293:
                        _ = len(v2)
                    case 1:
                        v4 = fi18
                    case -728:
                        var v42 uint = uint(v5)
                        _ = v42
                    }
                    v4 = v4
                    fmt.Println(fi18, v5, v5)
                }
            } else {
                var v43 uint64 = (v5 * (v5 >> ((18446744073709551615) & 63)))
                for fi44 := 0; fi44 < 12; fi44++ {
                    v4 = v4
                    var v45 uint = 1
                    _ = fi44
                    _ = v45
                }
                var v46 uint16 = 73
                _ = v43
                _ = v46
            }
            v47 := GenS9{F0: -35, F1: 0}
            _ = v47
        } else {
            var v48 string = ""
            v3 = ((("" + "a") + "a") + "hello")
            _ = v48
        }
        _ = fi18
    }
    v3 = v3
    if true {
        v49 := GenS5{F0: int32(v1), F1: true}
        if (!(false || (v4 <= (^(v1))))) {
            if true {
                v50 := GenS6{F0: 18446744073709551615, F1: 65535}
                switch "hello" {
                case "hello":
                    v3 = v3
                case "\n\t":
                    var v51 bool = (0 >= genFn1(0, v3, 359))
                    _ = v51
                default:
                    var v52 float32 = float32(-2.25)
                    _ = v52
                }
                _ = v50
            }
            switch 749 {
            case 42:
                if true {
                    var v53 float64 = -2.25
                    _ = v53
                }
            case -2147483648:
                for fi54 := 0; fi54 < 2; fi54++ {
                    fmt.Println(fi54, v4, fi54)
                    var v55 uint8 = 49
                    v55 += (v55 >> ((42) & 63))
                    _ = fi54
                    _ = v55
                }
            case -99:
                var v56 float64 = 0.0
                _ = v56
            default:
                v1 = (-(v4))
            }
        } else {
            switch "" {
            case "":
                v4 = 418
            case "hello":
                if true {
                    v1 += v1
                    v5 -= v5
                } else {
                    fmt.Println(v3)
                    switch v3 {
                    case "hello":
                        for fi57 := 0; fi57 < 19; fi57++ {
                            var v58 bool = false
                            v59 := GenS10{F0: 172, F1: 0, F2: 255}
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    }
                }
            default:
                v5 += 8
            }
        }
        fmt.Println(v5, v5, v3)
        _ = v49
    } else {
        switch v1 {
        case 127:
            switch v1 {
            case 0:
                for fi60 := 0; fi60 < 18; fi60++ {
                    if (42 <= 0) {
                        if (!false) {
                            v61 := GenS7{F0: 65535}
                            var v62 int8 = 127
                            var v63 int8 = 1
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        v2 = append(v2, (v3 + ((v3 + (v3 + ("hello" + (v3 + v3)))) + v3)))
                        fmt.Println(v1, v1)
                    } else {
                        for fi64 := 0; fi64 < 14; fi64++ {
                            var v65 float64 = 1.5
                            _ = fi64
                            _ = v65
                        }
                        if false {
                            for fi66 := 0; fi66 < 3; fi66++ {
                                var v67 int = (v4 * v1)
                                _ = len(v2)
                                if false {
                                    switch v3 {
                                    case "a":
                                        v5 = (v5 << ((39) & 63))
                                    default:
                                        switch (((v3 + v3) + (("\n\t" + (v3 + v3)) + "hello")) + (v3 + v3)) {
                                        case "foo bar":
                                            v2 = append(v2, v3)
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "a":
                                            if false {
                                                switch (v1 + 4294967295) {
                                                case 915:
                                                    switch ((("a" + v3) + v3) + "\n\t") {
                                                    case "hello":
                                                        v3 = v3
                                                    case "\n\t":
                                                        var v68 uint16 = 255
                                                        _ = v68
                                                    case "foo bar":
                                                        if true {
                                                            fmt.Println(v4, fi66)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            var v69 string = v3
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                                if true {
                                                    if len(v2) > 0 {
                                                        v2[0] = v3
                                                    }
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v70 := GenS4{F0: (v1 | 127), F1: 65535, F2: false}
                                                    _ = v70
                                                }
                                            }
                                        }
                                    }
                                    if true {
                                        v5 = v5
                                    } else {
                                        switch "" {
                                        case "":
                                            v4 = fi66
                                        case "hello":
                                            _ = len(v2)
                                        default:
                                            var v71 int32 = 901
                                            _ = v71
                                        }
                                        v72 := GenS11{F0: 32767, F1: uint8(v5), F2: 48}
                                        _ = v72
                                    }
                                    var v73 float64 = float64(0.0)
                                    var v74 float32 = float32(0.0)
                                    var v75 int8 = 1
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    v4 = 9223372036854775807
                                    v76 := GenS12{F0: 42}
                                    var v77 float64 = 1.5
                                    _ = v76
                                    _ = v77
                                }
                                _ = fi66
                                _ = v67
                            }
                        }
                        var v78 float64 = float64((1.5 - 1.5))
                        _ = v78
                    }
                    _ = fi60
                }
            default:
                v2 = append(v2, v3)
            }
        case 4294967295:
            v1 += (-(v1))
        case -2147483648:
            var v79 int16 = 291
            _ = v79
        default:
            for fi80 := 0; fi80 < 13; fi80++ {
                v4 -= v4
                _ = fi80
            }
        }
    }
    var v81 uint8 = uint8(v5)
    v3 = v3
    var v82 uint32 = 65535
    if true {
        v83 := GenS13{F0: -1, F1: 32767}
        for fi84 := 0; fi84 < 19; fi84++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi84
        }
        var v85 bool = true
        _ = v83
        _ = v85
    } else {
        v2 = append(v2, v3)
        v81 += v81
    }
    var v86 uint32 = 42
    for fi87 := 0; fi87 < 13; fi87++ {
        v88 := GenS5{F0: 1, F1: true}
        fmt.Println(v86, v1)
        if (true || (int16(v4) != 1000)) {
            for fi89 := 0; fi89 < 20; fi89++ {
                var v90 uint8 = v81
                var v91 int = (v4 / ((-391) | 1))
                _ = fi89
                _ = v90
                _ = v91
            }
        }
        _ = fi87
        _ = v88
    }
    var v92 string = ("foo bar" + "\n\t")
    var v93 int64 = -9223372036854775808
    v5 += v5
    if (!true) {
        for fi94 := 0; fi94 < 14; fi94++ {
            fmt.Println(v81, v86, v86)
            var v95 int64 = 0
            var v96 int16 = 42
            _ = fi94
            _ = v95
            _ = v96
        }
        _ = len(v2)
    }
    for fi97 := 0; fi97 < 1; fi97++ {
        var v98 uint16 = 42
        var v99 float64 = (float64(1e10) / (((3.14159 + 0.0) - 0.1) * (1.5 + 0.0)))
        var v100 int8 = 0
        _ = fi97
        _ = v100
        _ = v98
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v86:", v86)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v86
    _ = v92
    _ = v93
}
