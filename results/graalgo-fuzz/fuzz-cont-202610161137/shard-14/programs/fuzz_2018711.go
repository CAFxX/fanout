package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 int }
type GenS5 struct { F0 uint16; F1 int8; F2 int32 }
type GenS6 struct { F0 int32; F1 bool; F2 bool }
type GenS7 struct { F0 int; F1 uint8; F2 string }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS9 struct { F0 int32; F1 int8; F2 bool }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint, p2 string) string {
    v5 := GenS1{F0: 822}
    v6 := GenS2{F0: 4294967295, F1: false}
    var v7 uint8 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("foo bar" + "hello")
}

func main() {
    var v1 int = 32767
    var v2 int8 = -128
    v3 := []int{(v1 + 32767), v1}
    var v4 uint64 = 18446744073709551615
    if (false && (100 == v2)) {
        fmt.Println(v2, v4, v1)
        for fi8 := 0; fi8 < 15; fi8++ {
            if true {
                fmt.Println(v2, v1)
                switch "a" {
                case "a":
                    v9 := GenS2{F0: v1, F1: true}
                    _ = v9
                case "\n\t":
                    for fi10 := 0; fi10 < 19; fi10++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if true {
                            if true {
                                v3 = append(v3, fi10)
                                if false {
                                    v2 = v2
                                    v1 = -359
                                    fmt.Println(v2, v2, v4)
                                }
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = fi8
                                }
                            }
                            if (100 != v2) {
                                v3 = append(v3, fi8)
                                var v11 uint = 4294967295
                                v12 := GenS3{F0: 84, F1: 884, F2: -1}
                                _ = v11
                                _ = v12
                            } else {
                                if false {
                                    var v13 float64 = 0.0
                                    var v14 float32 = 0.0
                                    var v15 float32 = float32(v13)
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                } else {
                                    if (0.1 != -2.25) {
                                        if true {
                                            v16 := GenS3{F0: v2, F1: 1000, F2: 42}
                                            _ = v16
                                        } else {
                                            fmt.Println(v4, fi10)
                                            switch "a" {
                                            case "a":
                                                switch "foo bar" {
                                                case "":
                                                    v17 := GenS3{F0: v2, F1: int16(v4), F2: 42}
                                                    _ = v17
                                                default:
                                                    v18 := GenS2{F0: (fi8 * fi10), F1: false}
                                                    v2 += (v2 / ((42) | 1))
                                                    v4 = 4294967295
                                                    _ = v18
                                                }
                                            default:
                                                if false {
                                                    if false {
                                                        if (!true) {
                                                            var v19 int16 = -1
                                                            _ = v19
                                                        } else {
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(v2, v1, fi8)
                                                            case "hello":
                                                                var v20 int = int(v2)
                                                                _ = v20
                                                            default:
                                                                switch "" {
                                                                case "foo bar":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    fmt.Println(v4)
                                                                }
                                                            }
                                                            v2 = (v2 * (-(v2)))
                                                            fmt.Println(fi10, fi10, fi8)
                                                            fmt.Println(fi10)
                                                            var v21 int16 = 1
                                                            _ = v21
                                                        }
                                                        var v22 float32 = -2.25
                                                        _ = v22
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v2 = v2
                                                        case "foo bar":
                                                            var v23 uint16 = uint16(v4)
                                                            switch (fi10 ^ v1) {
                                                            case 128:
                                                                v4 = (v4 & (v4 - 1))
                                                            case -932:
                                                                _ = len(v3)
                                                            }
                                                            if (("" < "") || (-650 < fi10)) {
                                                                var v24 bool = (int32(v23) < 42)
                                                                _ = v24
                                                            } else {
                                                                if true {
                                                                    var v25 float32 = (1.5 + 0.0)
                                                                    var v26 float64 = 3.14159
                                                                    fmt.Println(fi8, v26, v4)
                                                                    _ = v25
                                                                    _ = v26
                                                                }
                                                                v27 := GenS2{F0: fi10, F1: true}
                                                                _ = v27
                                                            }
                                                            _ = v23
                                                        case "\n\t":
                                                            v28 := GenS3{F0: (v2 / ((v2) | 1)), F1: 1000, F2: 886}
                                                            _ = v28
                                                        default:
                                                            if (true || true) {
                                                                v4 -= (v4 % ((v4) | 1))
                                                            } else {
                                                                if true {
                                                                    var v29 string = ("\n\t" + "")
                                                                    v30 := GenS4{F0: uint16(v4), F1: (1 == 0), F2: v1}
                                                                    v31 := GenS3{F0: v2, F1: 916, F2: int64(v2)}
                                                                    if (42 != 730) {
                                                                        v32 := GenS5{F0: uint16(v4), F1: v2, F2: 586}
                                                                        fmt.Println(v29)
                                                                        if false {
                                                                            v29 = v29
                                                                            v4 -= (v4 * (^(v4)))
                                                                            var v33 uint = 4294967295
                                                                            _ = v33
                                                                        }
                                                                        _ = v32
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v1
                                                                        }
                                                                        v34 := GenS2{F0: fi8, F1: true}
                                                                        v29 = "foo bar"
                                                                        _ = v34
                                                                    }
                                                                    _ = v29
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                v2 = (v2 & v2)
                                                                if false {
                                                                    var v35 float32 = (1e10 * -2.25)
                                                                    if false {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            fmt.Println(v2, fi10)
                                                                            v36 := GenS6{F0: 742, F1: true, F2: (0 <= 17)}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v36
                                                                        case "\n\t":
                                                                            if false {
                                                                                var v37 int = (fi8 | (v1 << ((0) & 63)))
                                                                                var v38 uint16 = 0
                                                                                _ = v37
                                                                                _ = v38
                                                                            } else {
                                                                                if true {
                                                                                    var v39 uint32 = 0
                                                                                    v39 -= (v39 - 42)
                                                                                    if ((-761 < -2147483648) || true) {
                                                                                        var v40 int32 = -84
                                                                                        fmt.Println(fi10, fi10)
                                                                                        _ = v40
                                                                                    } else {
                                                                                        v39 = v39
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v1, v2)
                                                                                        v3 = append(v3, v1)
                                                                                        if (!((true && ((v4 - (v4 % (((v4 >> ((4294967295) & 63))) | 1))) < v4)) && true)) {
                                                                                            var v41 bool = false
                                                                                            _ = v41
                                                                                        }
                                                                                    }
                                                                                    _ = v39
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (fi10 & (-(fi10)))
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            v2 -= (v2 & v2)
                                                                        default:
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                if ((true || false) && false) {
                                                                                    var v42 string = ("\n\t" + ("\n\t" + "\n\t"))
                                                                                    v43 := GenS3{F0: (v2 % (((v2 << ((0) & 63))) | 1)), F1: -850, F2: 9223372036854775807}
                                                                                    v42 = ("a" + v42)
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                }
                                                                            case "hello":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case "a":
                                                                                if (197 < 42) {
                                                                                    fmt.Println(v1, v2)
                                                                                    var v44 int = fi10
                                                                                    _ = v44
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(fi10)
                                                                    }
                                                                    if false {
                                                                        v3 = append(v3, 9223372036854775807)
                                                                        v35 -= ((v35 - 0.0) - 1.5)
                                                                        if (!(false || false)) {
                                                                            fmt.Println(fi8, v1, v1)
                                                                            fmt.Println(v1, fi10, fi8)
                                                                        }
                                                                        var v45 bool = true
                                                                        _ = v45
                                                                    } else {
                                                                        v46 := GenS6{F0: -908, F1: false, F2: ((fi8 * -2147483648) >= (v1 >> ((65535) & 63)))}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (-(fi8))
                                                                        }
                                                                        _ = v46
                                                                    }
                                                                    _ = v35
                                                                } else {
                                                                    v4 = v4
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        var v47 uint64 = (v4 + (v4 % ((v4) | 1)))
                                                        _ = v47
                                                    case "foo bar":
                                                        _ = len(v3)
                                                    default:
                                                        if ("foo bar" <= "foo bar") {
                                                            v4 = v4
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v48 int8 = (^(v2))
                                                                _ = v48
                                                            }
                                                            v2 += -1
                                                            v2 += 42
                                                        } else {
                                                            v3 = append(v3, (v1 | (fi8 & (v1 & v1))))
                                                            if (false || false) {
                                                                v49 := GenS3{F0: v2, F1: 32767, F2: 460}
                                                                v1 = (v1 - fi8)
                                                                _ = v49
                                                            } else {
                                                                var v50 int8 = (v2 << ((1) & 63))
                                                                if ((-1 > 1) && false) {
                                                                    _ = len(v3)
                                                                    v51 := GenS7{F0: 32768, F1: 127, F2: "foo bar"}
                                                                    if true {
                                                                        v2 -= -1
                                                                        v4 -= 965
                                                                        fmt.Println(v2)
                                                                    }
                                                                    _ = v51
                                                                } else {
                                                                    var v52 int64 = 42
                                                                    var v53 float32 = 0.0
                                                                    var v54 uint = 255
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                switch (-(fi8)) {
                                                                case 32767:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            v2 = 42
                                                                        } else {
                                                                            fmt.Println(fi10, v2)
                                                                        }
                                                                    }
                                                                default:
                                                                    v50 -= -128
                                                                }
                                                                _ = v50
                                                            }
                                                            v55 := GenS7{F0: 32768, F1: 2, F2: ""}
                                                            v56 := GenS8{F0: 2147483648, F1: 255, F2: 978}
                                                            v3 = append(v3, int(v2))
                                                            v4 -= 2
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                    fmt.Println(v1)
                                                } else {
                                                    v2 -= (v2 * int8(fi10))
                                                    if true {
                                                        v57 := GenS9{F0: -2147483648, F1: v2, F2: true}
                                                        _ = v57
                                                    }
                                                }
                                            }
                                            var v58 int16 = int16(v2)
                                            v2 -= 105
                                            _ = v58
                                        }
                                    }
                                    switch (^(fi8)) {
                                    case 1:
                                        fmt.Println(v2, v1)
                                    case -2147483648:
                                        if (255 == 13) {
                                            var v59 int = v1
                                            _ = v59
                                        }
                                    default:
                                        if true {
                                            if (float64(0.1) <= (((-2.25 * 0.0) * (float64(3.14159) / ((0.1 + ((0.1 - (3.14159 / 1.5)) - ((0.1 - (1.5 * float64((float32(3.14159) / 1.5)))) - 1e10))) - 0.0))) + -2.25)) {
                                                if false {
                                                    v2 = (v2 * v2)
                                                    switch "foo bar" {
                                                    case "":
                                                        if (267 >= 65535) {
                                                            v3 = append(v3, (^(v1)))
                                                            var v60 int8 = v2
                                                            _ = v60
                                                        }
                                                    default:
                                                        var v61 int8 = (v2 / (((v2 + 24)) | 1))
                                                        _ = v61
                                                    }
                                                    v4 = v4
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(v1, v1, v4)
                                                    case "\n\t":
                                                        v1 = (v1 << ((963) & 63))
                                                    case "foo bar":
                                                        fmt.Println(v4)
                                                    }
                                                    if (!true) {
                                                        var v62 string = "a"
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v62
                                                    } else {
                                                        var v63 string = "a"
                                                        _ = v63
                                                    }
                                                    v2 = 100
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    fmt.Println(v1)
                                                }
                                            } else {
                                                var v64 int32 = 1
                                                _ = v64
                                            }
                                            fmt.Println(fi10)
                                            _ = len(v3)
                                            v4 -= v4
                                            if true {
                                                if false {
                                                    v65 := GenS9{F0: -688, F1: v2, F2: (v1 < (^(fi10)))}
                                                    v2 = v2
                                                    _ = v65
                                                }
                                                _ = len(v3)
                                            } else {
                                                v3 = append(v3, 1)
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v66 uint = 1
                                                _ = v66
                                            case "a":
                                                if (!(0 < 2)) {
                                                    v67 := GenS10{F0: 255}
                                                    if false {
                                                        var v68 string = "a"
                                                        _ = v68
                                                    } else {
                                                        fmt.Println(fi10)
                                                        v4 -= 42
                                                    }
                                                    _ = v67
                                                } else {
                                                    var v69 float32 = 3.14159
                                                    _ = v69
                                                }
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            default:
                                                fmt.Println(v4, fi8, v1)
                                                v1 -= v1
                                                fmt.Println(fi10, v2, v2)
                                            }
                                        } else {
                                            if (false || false) {
                                                v4 += (v4 * (v4 >> ((798) & 63)))
                                                if len(v3) > 0 {
                                                    v3[0] = (fi8 & (fi10 << ((18446744073709551615) & 63)))
                                                }
                                                var v70 int64 = -739
                                                v1 = fi10
                                                v71 := GenS8{F0: 42, F1: 255, F2: 0}
                                                _ = v70
                                                _ = v71
                                            } else {
                                                var v72 uint32 = 65535
                                                _ = v72
                                            }
                                            var v73 string = (("foo bar" + ("\n\t" + "\n\t")) + (("\n\t" + "") + "foo bar"))
                                            v1 -= fi8
                                            _ = v73
                                        }
                                    }
                                    var v74 int8 = v2
                                    _ = v74
                                }
                                v75 := GenS4{F0: 172, F1: false, F2: int(v4)}
                                _ = v75
                            }
                            v1 = -178
                        } else {
                            v1 = 2147483647
                            v3 = append(v3, (fi8 - (^(v1))))
                        }
                        _ = fi10
                    }
                }
            } else {
                var v76 string = ("\n\t" + (("foo bar" + (("a" + "hello") + "foo bar")) + "foo bar"))
                var v77 bool = (true || ((!false) && (1.5 >= 1e10)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v76
                _ = v77
            }
            for fi78 := 0; fi78 < 12; fi78++ {
                v3 = append(v3, fi78)
                _ = fi78
            }
            _ = fi8
        }
    } else {
        v1 -= 2147483647
        switch (v1 / (((v1 | v1)) | 1)) {
        case -126:
            var v79 int16 = 269
            _ = v79
        }
    }
    var v80 int64 = 42
    if (false || false) {
        var v81 uint = 1
        for fi82 := 0; fi82 < 6; fi82++ {
            var v83 float64 = ((0.0 / 1.5) / (-2.25 - 0.1))
            _ = fi82
            _ = v83
        }
        var v84 uint32 = 2147483647
        _ = v81
        _ = v84
    }
    var v85 string = "\n\t"
    var v86 int = (-(v1))
    if true {
        var v87 int32 = -1
        v2 -= -128
        _ = v87
    }
    v3 = append(v3, v1)
    v85 = (((v85 + "\n\t") + v85) + (v85 + (v85 + "\n\t")))
    var v88 uint8 = 12
    v3 = append(v3, -797)
    v85 = ("" + v85)
    v89 := GenS6{F0: -2147483648, F1: (-32768 > 32767), F2: false}
    for fi90 := 0; fi90 < 3; fi90++ {
        v86 = v1
        var v91 int8 = (v2 * (v2 ^ -128))
        _ = fi90
        _ = v91
    }
    switch (v85 + "a") {
    case "":
        var v92 uint16 = 0
        _ = v92
    default:
        v85 = (v85 + v85)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v85
    _ = v86
    _ = v88
    _ = v89
}
