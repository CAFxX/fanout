package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 uint8; F1 int64; F2 int }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 uint }
func genFn0() int {
    var v4 int8 = -82
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            v4 += (v4 / ((-40) | 1))
            v4 = 127
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        v4 = (v4 << ((42) & 63))
        _ = fi5
    }
    _ = v4
    return -720
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 1
    var v3 int16 = -218
    fmt.Println(v2)
    delete(v1, 1)
    if false {
        v8 := []int{}
        fmt.Println(v2, v3)
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        _ = v8
    } else {
        if (!false) {
            v2 -= v2
        } else {
            v2 = v2
        }
        var v10 int64 = 0
        _ = v10
    }
    v11 := GenS1{F0: false, F1: v3, F2: -128}
    if true {
        v12 := GenS2{F0: v2, F1: 1, F2: 2147483647}
        v2 -= (v2 + (v2 >> ((255) & 63)))
        v13 := []int{genFn0(), 2147483647, genFn0(), -6}
        _ = v12
        _ = v13
    }
    v14 := []int{-638, genFn0(), 4294967295}
    v3 -= 1000
    if true {
        var v15 float64 = (((0.1 * 3.14159) / -2.25) + 0.1)
        v3 -= (v3 ^ (v3 | (v3 << ((1) & 63))))
        _ = v15
    } else {
        v16 := GenS3{F0: -243, F1: 518}
        for fi17 := 0; fi17 < 7; fi17++ {
            if false {
                v2 -= 255
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 bool = (1 > 1)
                    if v19 {
                        if v19 {
                            v14 = append(v14, fi17)
                        }
                        fmt.Println(fi18, fi17)
                        if (2147483647 >= 26) {
                            v3 -= v3
                            var v20 float32 = 1.5
                            if (!(953 <= 1)) {
                                fmt.Println(v2, v19, v20)
                                v2 += v2
                            } else {
                                if (!((v2 & v2) < v2)) {
                                    if (!v19) {
                                        v19 = true
                                        v19 = (!(!((fi17 + (fi18 | -9223372036854775808)) != int(v2))))
                                        v2 = v2
                                        switch "hello" {
                                        case "":
                                            v20 = v20
                                        case "a":
                                            v21 := GenS4{F0: genFn0()}
                                            _ = v21
                                        default:
                                            if v19 {
                                                delete(v1, 2)
                                                switch "hello" {
                                                case "":
                                                    if v19 {
                                                        v3 += v3
                                                    }
                                                case "\n\t":
                                                    v22 := GenS5{F0: v3, F1: fi17}
                                                    _ = v22
                                                }
                                                delete(v1, 2)
                                            }
                                        }
                                        var v23 string = "hello"
                                        _ = v23
                                    } else {
                                        var v24 int32 = 1
                                        _ = v24
                                    }
                                } else {
                                    fmt.Println(v20)
                                    fmt.Println(fi17, fi18, v20)
                                    v25 := GenS4{F0: 9223372036854775807}
                                    _ = v25
                                }
                                var v26 uint32 = 915
                                var v27 bool = (v19 || v19)
                                _ = v26
                                _ = v27
                            }
                            _ = v20
                        }
                    } else {
                        if (((-(fi18)) == -619) || (false || v19)) {
                            v28 := GenS1{F0: true, F1: (v3 ^ (v3 << ((255) & 63))), F2: 42}
                            _ = v28
                        } else {
                            var v29 string = "hello"
                            _ = v29
                        }
                        var v30 int8 = -128
                        var v31 string = ""
                        _ = v30
                        _ = v31
                    }
                    _ = fi18
                    _ = v19
                }
            }
            var v32 int16 = v3
            _ = fi17
            _ = v32
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v16
    }
    switch 128 {
    case 968:
        if true {
            if (("foo bar" + "a") < (("\n\t" + "foo bar") + "a")) {
                v2 = 1
                if (true || false) {
                    v2 = 255
                    v1[65535] = false
                } else {
                    v33 := GenS4{F0: genFn0()}
                    v34 := GenS3{F0: 127, F1: 0}
                    v35 := GenS6{F0: 242, F1: 127, F2: (v2 | v2)}
                    _ = v33
                    _ = v34
                    _ = v35
                }
                v3 += v3
            }
        } else {
            if (!false) {
                fmt.Println(v2, v2)
            }
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS6{F0: 4294967295, F1: 128, F2: (v2 * (v2 + 1))}
                v3 = v3
                _ = fi36
                _ = v37
            }
        }
    case 128:
        v38 := GenS7{F0: 127, F1: (v2 + (^(v2)))}
        _ = v38
    default:
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 int16 = (v3 >> ((4294967295) & 63))
            if false {
                var v41 uint8 = 42
                v3 = -427
                _ = v41
            } else {
                v42 := GenS8{F0: 889, F1: false}
                _ = v42
            }
            var v43 int8 = 0
            _ = fi39
            _ = v40
            _ = v43
        }
    }
    v14 = append(v14, 32767)
    _ = len(v14)
    for fi44 := 0; fi44 < 14; fi44++ {
        if false {
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 float64 = ((1.5 / 1.5) * (float64(((float32(1.5) * (float32((-2.25 / 3.14159)) * ((0.1 + 0.0) * float32(-2.25)))) - ((0.1 + 1.5) + 1.5))) + (((1e10 + (-2.25 * (float64(3.14159) * -2.25))) / 1.0) / float64(1e10))))
                fmt.Println(v46, v2, v46)
                var v47 float32 = ((1e10 - ((1e10 / (3.14159 - (3.14159 + 1e10))) / (3.14159 * 3.14159))) / (0.0 / ((3.14159 * (0.1 * float32(v46))) * 0.0)))
                _ = fi45
                _ = v46
                _ = v47
            }
            fmt.Println(fi44)
        }
        _ = fi44
    }
    for fi48 := 0; fi48 < 4; fi48++ {
        for fi49 := 0; fi49 < 18; fi49++ {
            v1[4294967295] = true
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            if true {
                var v50 uint32 = 254
                _ = v50
            }
            _ = fi49
        }
        _ = len(v14)
        for fi51 := 0; fi51 < 16; fi51++ {
            v3 = (v3 >> ((4294967295) & 63))
            v2 = (v2 | 0)
            var v52 float32 = ((((-2.25 + ((3.14159 / -2.25) / (1e10 / 0.1))) + ((-2.25 * -2.25) * 0.0)) + -2.25) + 1.5)
            var v53 float32 = (v52 * (v52 / float32(3.14159)))
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi48
    }
    for fi54 := 0; fi54 < 10; fi54++ {
        v3 += v3
        _ = fi54
    }
    for fi55 := 0; fi55 < 11; fi55++ {
        switch "" {
        case "hello":
            v56 := GenS3{F0: 32768, F1: 479}
            _ = v56
        case "a":
            v1[116] = (-1 > 2147483647)
        case "\n\t":
            if ((((-2.25 / ((-2.25 + -2.25) * 1.5)) - float32(1.5)) * (0.0 + float32(3.14159))) < 0.1) {
                _ = len(v14)
            }
        default:
            v2 += (v2 / ((v2) | 1))
        }
        v57 := GenS4{F0: fi55}
        _ = fi55
        _ = v57
    }
    if (184 > 127) {
        switch 2147483647 {
        case 563:
            if (!(false || (1e10 >= 3.14159))) {
                if (18446744073709551615 > 1) {
                    switch "foo bar" {
                    case "a":
                        v2 = (v2 >> ((4294967295) & 63))
                    case "foo bar":
                        if len(v14) > 0 {
                            v14[0] = genFn0()
                        }
                    case "":
                        v58 := v1[4294967295]
                        _ = v58
                    default:
                        if false {
                            v3 = v3
                            switch "" {
                            case "a":
                                if true {
                                    var v59 uint8 = 79
                                    for fi60 := 0; fi60 < 14; fi60++ {
                                        if (1e10 != 1e10) {
                                            var v61 float64 = ((0.1 - -2.25) - 1e10)
                                            _ = v61
                                        }
                                        _ = fi60
                                    }
                                    if (!(("hello" > "\n\t") && true)) {
                                        var v62 bool = (false && true)
                                        var v63 int64 = -9223372036854775808
                                        if len(v14) > 0 {
                                            v14[0] = 680
                                        }
                                        _ = v62
                                        _ = v63
                                    } else {
                                        v59 = (v59 ^ (v59 - v59))
                                        v14 = append(v14, -875)
                                        if false {
                                            v2 = (^(v2))
                                            switch "\n\t" {
                                            case "foo bar":
                                                if (42 == uint(v2)) {
                                                    v59 = v59
                                                    v64 := v1[65535]
                                                    _ = v64
                                                } else {
                                                    if true {
                                                        v65 := GenS9{F0: 127, F1: -55, F2: -1}
                                                        _ = v65
                                                    } else {
                                                        v66 := GenS3{F0: 852, F1: 7}
                                                        for fi67 := 0; fi67 < 5; fi67++ {
                                                            fmt.Println(v2, v2, v59)
                                                            var v68 uint64 = 4294967295
                                                            for fi69 := 0; fi69 < 15; fi69++ {
                                                                v68 = 4294967295
                                                                v59 = 179
                                                                _ = fi69
                                                            }
                                                            _ = fi67
                                                            _ = v68
                                                        }
                                                        _ = v66
                                                    }
                                                    switch 2147483647 {
                                                    case -9223372036854775808:
                                                        v70 := GenS6{F0: 959, F1: (^(v59)), F2: (v2 << ((355) & 63))}
                                                        _ = v70
                                                    case 9223372036854775807:
                                                        var v71 uint8 = v59
                                                        _ = v71
                                                    case 1:
                                                        for fi72 := 0; fi72 < 5; fi72++ {
                                                            for fi73 := 0; fi73 < 1; fi73++ {
                                                                fmt.Println(fi72, fi73, fi72)
                                                                _ = fi73
                                                            }
                                                            _ = fi72
                                                        }
                                                    }
                                                    for fi74 := 0; fi74 < 7; fi74++ {
                                                        v2 -= v2
                                                        _ = fi74
                                                    }
                                                }
                                            case "\n\t":
                                                for fi75 := 0; fi75 < 11; fi75++ {
                                                    var v76 int8 = 85
                                                    _ = fi75
                                                    _ = v76
                                                }
                                            default:
                                                v77 := GenS10{F0: 467, F1: 18446744073709551615}
                                                _ = v77
                                            }
                                            var v78 uint32 = 2147483647
                                            _ = v78
                                        } else {
                                            var v79 float32 = 1e10
                                            var v80 int8 = 110
                                            _ = len(v1)
                                            _ = v79
                                            _ = v80
                                        }
                                    }
                                    _ = v59
                                } else {
                                    v2 = (^(v2))
                                    v2 = (v2 >> ((255) & 63))
                                }
                            default:
                                if false {
                                    v2 += v2
                                    var v81 float32 = 0.0
                                    switch 32768 {
                                    case 32767:
                                        var v82 uint8 = 255
                                        _ = v82
                                    case 1:
                                        v3 = v3
                                    case 677:
                                        v81 = ((v81 / (v81 - v81)) - v81)
                                    default:
                                        for fi83 := 0; fi83 < 9; fi83++ {
                                            v84 := GenS8{F0: 255, F1: (!(("\n\t" + ("hello" + "")) <= ""))}
                                            v81 = v81
                                            _ = fi83
                                            _ = v84
                                        }
                                    }
                                    _ = v81
                                } else {
                                    var v85 float64 = ((1e10 * 3.14159) / (float64(-2.25) + (1.5 + 1e10)))
                                    _ = v85
                                }
                            }
                        } else {
                            fmt.Println(v3, v3)
                        }
                    }
                }
            } else {
                if true {
                    v86 := v1[451]
                    for fi87 := 0; fi87 < 6; fi87++ {
                        v88 := GenS11{F0: 73, F1: v3}
                        var v89 float64 = 0.1
                        v3 = 32767
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                    v90 := GenS4{F0: -238}
                    _ = v86
                    _ = v90
                } else {
                    v3 += v3
                    v91 := GenS5{F0: (^(v3)), F1: 42}
                    v1[1] = false
                    _ = v91
                }
                switch "hello" {
                case "":
                    var v92 int64 = -995
                    _ = v92
                case "\n\t":
                    v3 -= v3
                case "foo bar":
                    var v93 bool = true
                    _ = v93
                }
            }
        case -991:
            if false {
                v3 += v3
            }
        default:
            fmt.Println(v2, v3)
        }
        switch 748 {
        case 128:
            for fi94 := 0; fi94 < 1; fi94++ {
                for fi95 := 0; fi95 < 1; fi95++ {
                    var v96 float64 = (0.1 - (3.14159 + 0.0))
                    v97 := GenS12{F0: 285}
                    fmt.Println(fi94, fi94)
                    if true {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        v96 = (v96 + v96)
                        v98 := GenS10{F0: 2147483648, F1: 477}
                        _ = v98
                    }
                    v2 += 42
                    v99 := GenS4{F0: int(v2)}
                    var v100 uint32 = 777
                    _ = fi95
                    _ = v100
                    _ = v96
                    _ = v97
                    _ = v99
                }
                if true {
                    if len(v14) > 0 {
                        v14[0] = 0
                    }
                    v3 = v3
                    fmt.Println(v2, v2, v2)
                }
                _ = fi94
            }
        case 9223372036854775807:
            var v101 int32 = 307
            _ = v101
        default:
            delete(v1, 4294967295)
        }
        v2 += v2
    } else {
        for fi102 := 0; fi102 < 7; fi102++ {
            var v103 uint16 = v2
            var v104 uint = 42
            fmt.Println(v3, v2)
            _ = fi102
            _ = v103
            _ = v104
        }
    }
    for fi105 := 0; fi105 < 10; fi105++ {
        if true {
            v14 = append(v14, fi105)
            switch "hello" {
            case "hello":
                var v106 int16 = v3
                _ = v106
            case "a":
                fmt.Println(v2, fi105, fi105)
            }
        }
        var v107 int = fi105
        _ = fi105
        _ = v107
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
}
