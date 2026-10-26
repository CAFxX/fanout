package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 int16; F1 uint; F2 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0(p0 int, p1 uint, p2 int8) int64 {
    v4 := []int{}
    fmt.Println(p2, p2)
    var v5 uint8 = 224
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 32767
    v3 := []uint32{42, 2147483647, 42, 147}
    var v6 uint8 = 195
    var v7 int16 = (v2 * (v2 - (v2 >> ((255) & 63))))
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            delete(v1, fi9)
            if (("\n\t" + "\n\t") <= ("\n\t" + "foo bar")) {
                v10 := GenS1{F0: (^(v6)), F1: 42, F2: 2147483648}
                if (!((100 > 1) && (65535 < 32767))) {
                    var v11 uint8 = (v6 ^ (v6 << ((77) & 63)))
                    if false {
                        if (!((!(39 > uint32(v6))) && false)) {
                            var v12 int64 = 42
                            fmt.Println(v2, fi8, v2)
                            v6 -= v6
                            _ = v12
                        } else {
                            delete(v1, 127)
                        }
                    }
                    v2 -= v2
                    _ = v11
                }
                if (-52 < -1) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    delete(v1, 32767)
                    v13 := GenS2{F0: false, F1: 42}
                    var v14 float32 = 0.0
                    _ = v13
                    _ = v14
                }
                _ = v10
            } else {
                var v15 string = "a"
                var v16 float32 = ((1.5 / -2.25) - float32(1.5))
                var v17 int8 = int8(v2)
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS3{F0: "a", F1: true, F2: 95}
            switch fi19 {
            case 128:
                var v21 int16 = v7
                _ = v21
            default:
                v22 := GenS3{F0: "foo bar", F1: true, F2: -180}
                _ = v22
            }
            if (false || ((((v6 == v6) || true) || (900 == 18446744073709551615)) || (("\n\t" + "hello") <= ("hello" + "hello")))) {
                if false {
                    v23 := GenS2{F0: (!false), F1: -1}
                    _ = v23
                }
                switch (fi18 - -1) {
                case 42:
                    fmt.Println(v6, v2, fi19)
                case 1:
                    var v24 int16 = -369
                    _ = v24
                case -1:
                    switch (fi19 % ((-9223372036854775808) | 1)) {
                    case 42:
                        v25 := GenS4{F0: v7, F1: ("a" + "")}
                        _ = v25
                    case 2147483647:
                        v26 := GenS5{F0: 42, F1: (v2 ^ 1000), F2: v6}
                        _ = v26
                    }
                default:
                    _ = len(v1)
                }
            } else {
                v1[1] = (fi18 ^ fi19)
            }
            _ = fi19
            _ = v20
        }
        var v27 int8 = -16
        v27 = v27
        _ = fi18
        _ = v27
    }
    v7 = v2
    var v28 float64 = -2.25
    if (-128 <= 100) {
        switch 32767 {
        case -9223372036854775808:
            if (1 < 995) {
                v1[-2147483648] = -9223372036854775808
            } else {
                if false {
                    if true {
                        v2 = (v2 | -32768)
                        var v29 float32 = 3.14159
                        var v30 float32 = v29
                        _ = v29
                        _ = v30
                    }
                }
                v6 = (v6 - (v6 % ((v6) | 1)))
                fmt.Println(v28, v6, v7)
            }
        case 127:
            switch 318 {
            case 9223372036854775807:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case 128:
            fmt.Println(v7, v7, v28)
        default:
            var v31 string = ((("\n\t" + "") + "\n\t") + "hello")
            _ = v31
        }
    }
    v32 := GenS6{F0: 42, F1: v6}
    v33 := GenS6{F0: 32768, F1: (v6 - 13)}
    v28 = ((float64(1.5) / v28) + v28)
    v34 := GenS5{F0: 42, F1: 32767, F2: (v6 - v6)}
    v6 = v6
    var v35 bool = (-540 == -823)
    for fi36 := 0; fi36 < 6; fi36++ {
        if v35 {
            var v37 uint = uint(v6)
            if (!(v35 || false)) {
                v3 = append(v3, 906)
                var v38 uint16 = 2
                var v39 float32 = 0.1
                _ = v38
                _ = v39
            }
            var v40 int32 = 42
            _ = v37
            _ = v40
        } else {
            switch fi36 {
            case 2147483647:
                switch "foo bar" {
                case "\n\t":
                    for fi41 := 0; fi41 < 11; fi41++ {
                        switch 697 {
                        case -879:
                            var v42 float32 = 3.14159
                            _ = v42
                        case 32768:
                            if v35 {
                                v35 = v35
                                var v43 uint8 = v6
                                var v44 uint = 486
                                var v45 uint16 = 42
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                v46 := v1[4294967295]
                                v46 -= (fi41 / (((fi41 ^ v46)) | 1))
                                _ = v46
                            }
                        }
                        if v35 {
                            var v47 string = ""
                            v35 = (v35 && v35)
                            v48 := GenS2{F0: (v6 >= (v6 + (^(v6)))), F1: -48}
                            v49 := GenS4{F0: (v2 >> ((42) & 63)), F1: "\n\t"}
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        var v50 float32 = ((0.0 * 1e10) / ((0.0 / 3.14159) * (float32(v28) + 0.0)))
                        _ = fi41
                        _ = v50
                    }
                case "a":
                    switch (("" + "\n\t") + (("foo bar" + "a") + (("foo bar" + "a") + (((("a" + "\n\t") + "\n\t") + "a") + "hello")))) {
                    case "":
                        if v35 {
                            if v35 {
                                var v51 bool = v35
                                var v52 int8 = 47
                                _ = v51
                                _ = v52
                            }
                            var v53 uint = 2
                            _ = v53
                        } else {
                            v54 := v1[fi36]
                            if (!true) {
                                if ((("\n\t" + "a") + ("" + ("foo bar" + "hello"))) <= "a") {
                                    var v55 uint32 = uint32(v6)
                                    _ = v55
                                }
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = 2
                                }
                            }
                            if v35 {
                                var v56 int32 = 2147483647
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    v54 = 4294967295
                                    var v58 int64 = 42
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v56
                            }
                            _ = v54
                        }
                    case "a":
                        if v35 {
                            if v35 {
                                fmt.Println(v28)
                            }
                            if v35 {
                                if v35 {
                                    fmt.Println(v6, v28, v28)
                                    v28 += v28
                                } else {
                                    v59 := GenS7{F0: -128, F1: v35}
                                    v60 := GenS1{F0: (v6 % (((v6 / ((v6) | 1))) | 1)), F1: (v7 % (((v7 | 727)) | 1)), F2: 42}
                                    _ = v59
                                    _ = v60
                                }
                                var v61 int = (-(fi36))
                                if len(v3) > 0 {
                                    v3[0] = 717
                                }
                                _ = v61
                            }
                        } else {
                            for fi62 := 0; fi62 < 3; fi62++ {
                                switch (fi62 % ((fi62) | 1)) {
                                case 745:
                                    var v63 uint16 = 704
                                    _ = v63
                                default:
                                    if v35 {
                                        if true {
                                            if true {
                                                switch fi62 {
                                                case 127:
                                                    var v64 uint16 = 32767
                                                    _ = v64
                                                case 9223372036854775807:
                                                    v65 := v1[(fi62 | (fi36 | (fi36 % (((fi36 | (fi36 - fi36))) | 1))))]
                                                    _ = v65
                                                default:
                                                    if false {
                                                        delete(v1, (fi36 / (((fi36 % (((fi62 & fi36)) | 1))) | 1)))
                                                        if v35 {
                                                            v7 = (v7 / ((v2) | 1))
                                                            v66 := GenS4{F0: (-(v7)), F1: ("\n\t" + ("a" + "foo bar"))}
                                                            v35 = (!false)
                                                            _ = v66
                                                        } else {
                                                            var v67 uint16 = 65535
                                                            _ = v67
                                                        }
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v2 = v7
                                                        var v68 uint32 = 4294967295
                                                        _ = v68
                                                    }
                                                }
                                                v69 := v1[int(v7)]
                                                _ = v69
                                            }
                                        } else {
                                            v28 = v28
                                        }
                                        v28 -= v28
                                        switch "hello" {
                                        case "\n\t":
                                            switch fi62 {
                                            case 0:
                                                v6 += 2
                                                fmt.Println(fi62, v35, fi62)
                                            case -9223372036854775808:
                                                switch fi62 {
                                                case 32767:
                                                    v28 -= v28
                                                case 32768:
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = 902
                                                        }
                                                        var v70 uint32 = 300
                                                        if true {
                                                            var v71 int = (fi36 ^ (fi62 & fi62))
                                                            v6 = 2
                                                            v72 := GenS4{F0: v2, F1: "\n\t"}
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        _ = v70
                                                    } else {
                                                        v6 = 255
                                                        var v73 float64 = v28
                                                        var v74 float32 = (3.14159 / (0.1 + 1.5))
                                                        switch "a" {
                                                        case "":
                                                            if v35 {
                                                                v3 = append(v3, 537)
                                                                var v75 string = ""
                                                                _ = v75
                                                            } else {
                                                                v7 = (v7 * int16(fi36))
                                                                v76 := GenS8{F0: (v2 % ((450) | 1)), F1: 2, F2: -9223372036854775808}
                                                                var v77 int8 = -128
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        default:
                                                            v6 -= v6
                                                            var v78 bool = (v6 >= (v6 & v6))
                                                            if (!true) {
                                                                v3 = append(v3, 324)
                                                                if ((-(v7)) > (v7 % (((^(v7))) | 1))) {
                                                                    v35 = (1 <= -9223372036854775808)
                                                                    var v79 int16 = v7
                                                                    _ = v79
                                                                }
                                                            } else {
                                                                fmt.Println(v78, v73, fi36)
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    if v35 {
                                                                        var v80 uint64 = 18446744073709551615
                                                                        v73 = v28
                                                                        _ = v80
                                                                    }
                                                                case "foo bar":
                                                                    v81 := GenS8{F0: v7, F1: 1, F2: (fi36 >> ((91) & 63))}
                                                                    _ = v81
                                                                default:
                                                                    if v78 {
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            v35 = v78
                                                                        case "":
                                                                            v82 := GenS4{F0: (v7 >> ((65535) & 63)), F1: "\n\t"}
                                                                            _ = v82
                                                                        case "\n\t":
                                                                            var v83 uint32 = 322
                                                                            _ = v83
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v78, fi62, v74)
                                                                    }
                                                                }
                                                            }
                                                            var v84 uint64 = 42
                                                            _ = v78
                                                            _ = v84
                                                        }
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                default:
                                                    v35 = (!true)
                                                }
                                            case 899:
                                                var v85 uint32 = 0
                                                _ = v85
                                            default:
                                                v28 = v28
                                            }
                                        case "a":
                                            v86 := GenS9{F0: -1}
                                            _ = v86
                                        default:
                                            v87 := GenS10{F0: 137, F1: 4294967295, F2: (v6 ^ v6)}
                                            _ = v87
                                        }
                                        v6 -= v6
                                        switch 42 {
                                        case 834:
                                            var v88 uint32 = 42
                                            _ = v88
                                        }
                                    } else {
                                        fmt.Println(fi36)
                                        _ = len(v1)
                                    }
                                }
                                _ = fi62
                            }
                            _ = len(v1)
                        }
                    case "foo bar":
                        v2 += (v2 & (v7 + 703))
                    }
                }
            case 254:
                var v89 uint16 = 2
                _ = v89
            default:
                v7 += v7
            }
            v6 = (v6 >> ((2) & 63))
            v28 = float64(1e10)
        }
        var v90 uint32 = 2147483647
        _ = fi36
        _ = v90
    }
    var v91 bool = (65535 < 918)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v6
    _ = v7
    _ = v91
}
