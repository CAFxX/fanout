package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint16; F1 uint }
type GenS13 struct { F0 bool; F1 int32; F2 int32 }
type GenS14 struct { F0 uint32 }
func genFn0() int16 {
    v5 := GenS1{F0: 255}
    v6 := []int{0}
    for fi7 := 0; fi7 < 10; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch 128 {
        case 128:
            var v8 int8 = -24
            _ = v8
        case -2147483648:
            var v9 uint8 = 128
            _ = v9
        }
        _ = fi7
    }
    var v10 float64 = -2.25
    _ = v10
    _ = v5
    _ = v6
    return 1000
}

func main() {
    v1 := []string{"", "", ""}
    var v2 string = ""
    var v3 float32 = float32(0.0)
    var v4 uint8 = 34
    var v11 uint8 = (v4 / ((v4) | 1))
    for fi12 := 0; fi12 < 10; fi12++ {
        if (((!((0.0 != -2.25) && true)) && false) || (42 == 575)) {
            v11 += (v11 ^ (v4 / (((^(v4))) | 1)))
            fmt.Println(v3, v11, v11)
            fmt.Println(v2, v3)
        } else {
            switch 42 {
            case 0:
                var v13 uint64 = 4294967295
                _ = v13
            case 9223372036854775807:
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(v11)
                    switch fi14 {
                    case 0:
                        if false {
                            if true {
                                if true {
                                    switch 127 {
                                    case 931:
                                        v3 -= v3
                                    case 1:
                                        switch fi14 {
                                        case 128:
                                            var v15 int32 = 1
                                            _ = v15
                                        case -828:
                                            var v16 float64 = 0.0
                                            _ = v16
                                        case -543:
                                            if true {
                                                if (42 == -34) {
                                                    v1 = append(v1, "a")
                                                    var v17 bool = (v2 < "hello")
                                                    _ = v17
                                                }
                                            }
                                        default:
                                            switch (fi14 | -294) {
                                            case -477:
                                                var v18 string = v2
                                                _ = v18
                                            }
                                        }
                                    }
                                    fmt.Println(v3, fi14, v11)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            } else {
                                var v19 int64 = -905
                                fmt.Println(v19, fi14, v19)
                                switch -9223372036854775808 {
                                case 1:
                                    var v20 float32 = ((float32(-2.25) * ((v3 / (-2.25 * 1e10)) - float32(0.1))) + (v3 * (v3 + v3)))
                                    _ = v20
                                case 335:
                                    switch fi12 {
                                    case 0:
                                        v3 -= ((v3 + v3) + (v3 - 0.1))
                                    case -929:
                                        v21 := GenS2{F0: v2}
                                        _ = v21
                                    case 436:
                                        v22 := GenS1{F0: 1}
                                        _ = v22
                                    }
                                default:
                                    v23 := GenS3{F0: 65535}
                                    _ = v23
                                }
                                _ = v19
                            }
                            var v24 int16 = 603
                            v4 = v4
                            v11 -= v11
                            _ = v24
                        }
                    case 504:
                        v4 += (v11 * v11)
                        if false {
                            switch "hello" {
                            case "":
                                var v25 int = -887
                                _ = v25
                            default:
                                switch ((((v2 + ((((("\n\t" + "a") + v2) + v2) + ("a" + "hello")) + ("\n\t" + ("" + (v2 + ""))))) + ("" + v2)) + v2) + v2) {
                                case "foo bar":
                                    var v26 int = (^(fi12))
                                    _ = v26
                                case "\n\t":
                                    var v27 uint8 = v4
                                    _ = v27
                                case "":
                                    if (0 <= -803) {
                                        if len(v1) > 0 {
                                            v1[0] = v2
                                        }
                                    } else {
                                        var v28 float32 = (v3 * ((float32((-2.25 + 1e10)) + 1.5) * 3.14159))
                                        if false {
                                            var v29 float32 = 0.0
                                            if (v2 < v2) {
                                                v30 := GenS2{F0: "a"}
                                                if len(v1) > 0 {
                                                    v1[0] = v2
                                                }
                                                _ = v30
                                            } else {
                                                v3 -= 0.1
                                                var v31 int16 = -32768
                                                if len(v1) > 0 {
                                                    v1[0] = v2
                                                }
                                                _ = v31
                                            }
                                            _ = v29
                                        }
                                        _ = v28
                                    }
                                    v4 = (v11 & (v11 << ((4294967295) & 63)))
                                }
                            }
                            v4 = (v4 & v11)
                        }
                        v3 -= (float32(0.0) - (v3 / 1.5))
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                        fmt.Println(v2)
                        if (!true) {
                            var v32 int = (fi14 << ((42) & 63))
                            _ = v32
                        }
                    case 32767:
                        v33 := GenS4{F0: 2}
                        _ = v33
                    default:
                        if (219 < 42) {
                            fmt.Println(v2, v4, v4)
                        } else {
                            if (!(!(0.1 >= ((float64((((0.1 / (v3 / 1.0)) - ((v3 - (float32(((0.1 / 3.14159) - 1e10)) * v3)) - (v3 + (v3 + (1e10 * ((float32((1.5 - (0.0 + (1e10 + ((float64(v3) / 0.1) / 1.5))))) + float32(3.14159)) - ((v3 + (v3 + v3)) * v3))))))) - v3)) * 1.5) / 1e10)))) {
                                var v34 bool = false
                                _ = v34
                            }
                        }
                    }
                    _ = fi14
                }
            }
            if ((!false) && false) {
                v2 = v2
                fmt.Println(fi12, v2)
                var v35 uint = 18446744073709551615
                _ = v35
            } else {
                if (true && (2147483647 != -1)) {
                    fmt.Println(v2, v3, fi12)
                } else {
                    if (v2 >= (v2 + v2)) {
                        if true {
                            fmt.Println(v4, v11)
                            var v36 float64 = (0.1 * 0.0)
                            var v37 int = 2147483647
                            _ = v36
                            _ = v37
                        }
                    }
                    var v38 string = "hello"
                    v2 = (("" + "\n\t") + v38)
                    _ = v38
                }
            }
            v39 := GenS2{F0: (("" + "\n\t") + v2)}
            _ = v39
        }
        _ = fi12
    }
    v2 = v2
    var v40 int16 = -997
    switch -152 {
    case 0:
        for fi41 := 0; fi41 < 20; fi41++ {
            fmt.Println(v11, v40)
            v40 += -617
            _ = fi41
        }
    case 4294967295:
        switch 1 {
        case 9223372036854775807:
            switch "hello" {
            case "hello":
                var v42 uint = 4294967295
                _ = v42
            case "a":
                v40 -= (v40 % (((^(v40))) | 1))
            case "foo bar":
                var v43 uint = 65535
                _ = v43
            default:
                v11 = 128
            }
        case 972:
            v44 := GenS1{F0: 255}
            _ = v44
        case -2147483648:
            var v45 int16 = v40
            _ = v45
        default:
            v46 := GenS1{F0: 833}
            _ = v46
        }
    }
    v47 := GenS2{F0: v2}
    if false {
        if len(v1) > 0 {
            v1[0] = ((("\n\t" + v2) + v2) + v2)
        }
        var v48 uint16 = uint16(v11)
        _ = v48
    } else {
        if (100 < -1) {
            var v49 uint = 0
            switch "" {
            case "hello":
                if (uint64(v11) >= 42) {
                    fmt.Println(v4, v49, v40)
                    for fi50 := 0; fi50 < 14; fi50++ {
                        fmt.Println(v2)
                        fmt.Println(v4, v11)
                        if (32767 >= 938) {
                            if (97 != 0) {
                                if (255 != v4) {
                                    v51 := GenS5{F0: 1, F1: 42}
                                    _ = v51
                                } else {
                                    v40 = (-(v40))
                                    fmt.Println(v3, v2, v3)
                                }
                            } else {
                                if (false && (!false)) {
                                    v4 += 0
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    var v52 uint32 = 1
                                    _ = v52
                                } else {
                                    for fi53 := 0; fi53 < 6; fi53++ {
                                        if false {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            switch ("foo bar" + v2) {
                                            case "\n\t":
                                                if (false || true) {
                                                    v54 := GenS6{F0: 0, F1: 505}
                                                    var v55 uint32 = 65535
                                                    if ((v11 ^ (^(v4))) != uint8(v49)) {
                                                        var v56 int32 = 503
                                                        var v57 int32 = (v56 | v56)
                                                        v58 := GenS3{F0: 255}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    if (false || false) {
                                                        v59 := GenS7{F0: (v11 / (((v11 | (v4 << (((v49 % (((v49 % ((65535) | 1))) | 1))) & 63)))) | 1))}
                                                        fmt.Println(v55)
                                                        _ = v59
                                                    } else {
                                                        v60 := GenS8{F0: 2147483647}
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(fi53, v3, v3)
                                                        _ = v60
                                                    }
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    if (v49 < 255) {
                                                        var v61 int8 = 0
                                                        _ = v61
                                                    } else {
                                                        v2 = (v2 + (v2 + ((v2 + ("foo bar" + v2)) + "\n\t")))
                                                        if true {
                                                            v62 := GenS4{F0: 4294967295}
                                                            v63 := GenS9{F0: false}
                                                            fmt.Println(fi53)
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            v64 := GenS6{F0: 2, F1: 255}
                                                            _ = v64
                                                        }
                                                    }
                                                    v65 := GenS10{F0: 56}
                                                    _ = v65
                                                }
                                            case "hello":
                                                v40 = v40
                                            case "a":
                                                if (float32(0.1) > ((v3 + ((0.0 / (v3 + v3)) + 3.14159)) * (1.5 / v3))) {
                                                    switch 32767 {
                                                    case -976:
                                                        v3 = v3
                                                        fmt.Println(v3)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    default:
                                                        fmt.Println(v49, fi53, v3)
                                                    }
                                                } else {
                                                    v3 -= float32(float64(((((v3 + v3) / v3) * -2.25) + (-2.25 / ((-2.25 + float32(0.1)) + (v3 * v3))))))
                                                    if true {
                                                        var v66 uint8 = v11
                                                        _ = v66
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = v2
                                                        }
                                                        v4 = v4
                                                    }
                                                    v67 := GenS1{F0: 2}
                                                    _ = len(v1)
                                                    _ = len(v1)
                                                    _ = v67
                                                }
                                            }
                                            if false {
                                                _ = len(v1)
                                                fmt.Println(v2)
                                                var v68 int16 = (v40 + (v40 - (v40 >> ((v49) & 63))))
                                                _ = v68
                                            } else {
                                                v2 = v2
                                                v69 := GenS11{F0: -412, F1: 42, F2: 42}
                                                var v70 uint16 = 65535
                                                fmt.Println(v70)
                                                var v71 uint8 = 128
                                                v72 := GenS5{F0: -24, F1: v71}
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                            fmt.Println(v3)
                                        }
                                        if false {
                                            if false {
                                                fmt.Println(fi53, fi50, fi50)
                                                fmt.Println(fi50, fi50)
                                            } else {
                                                var v73 int = (^(fi53))
                                                _ = v73
                                            }
                                        }
                                        _ = fi53
                                    }
                                }
                            }
                            if true {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    if true {
                                        v3 = v3
                                        if len(v1) > 0 {
                                            v1[0] = (("a" + "a") + v2)
                                        }
                                    }
                                    fmt.Println(v4, v2)
                                    if (!((838 == 0) || false)) {
                                        var v74 int8 = -128
                                        var v75 float32 = ((0.0 + v3) + v3)
                                        _ = v74
                                        _ = v75
                                    }
                                }
                                for fi76 := 0; fi76 < 20; fi76++ {
                                    var v77 float32 = 0.0
                                    var v78 bool = false
                                    v79 := GenS12{F0: 953, F1: 10}
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                }
                                if (v40 == (v40 % ((v40) | 1))) {
                                    var v80 uint32 = 42
                                    var v81 string = ("" + v2)
                                    _ = v80
                                    _ = v81
                                } else {
                                    v82 := GenS12{F0: 65535, F1: v49}
                                    _ = v82
                                }
                            }
                            var v83 int32 = -1
                            _ = v83
                        } else {
                            v4 = (v4 << (((v49 | (v49 << (((v49 & (v49 ^ 34))) & 63)))) & 63))
                        }
                        _ = fi50
                    }
                }
            case "foo bar":
                var v84 int = -9223372036854775808
                _ = v84
            case "\n\t":
                v85 := GenS10{F0: 106}
                _ = v85
            }
            _ = v49
        } else {
            v40 = v40
            v4 = (v11 | (v4 ^ v11))
            if (false || false) {
                v4 = v4
                fmt.Println(v40, v40, v3)
            }
        }
        var v86 uint16 = 42
        v11 = v11
        _ = v86
    }
    v2 = v2
    var v87 int8 = 127
    _ = len(v1)
    var v88 int8 = v87
    fmt.Println(v2)
    v88 = 42
    fmt.Println(v40, v87, v40)
    switch -759 {
    case -1:
        v89 := GenS13{F0: false, F1: 0, F2: 1}
        _ = v89
    }
    if true {
        v3 += v3
    } else {
        for fi90 := 0; fi90 < 3; fi90++ {
            for fi91 := 0; fi91 < 8; fi91++ {
                var v92 uint8 = v4
                if (!(752 <= 255)) {
                    switch ("\n\t" + ("" + ((v2 + (("foo bar" + (v2 + "a")) + (v2 + "hello"))) + (v2 + (((v2 + ("" + ((v2 + "") + "hello"))) + ("foo bar" + v2)) + ("" + (("\n\t" + v2) + v2))))))) {
                    case "a":
                        v88 -= -85
                    case "foo bar":
                        fmt.Println(v92, fi90)
                    default:
                        var v93 uint8 = 60
                        _ = v93
                    }
                    if false {
                        var v94 int8 = 56
                        _ = v94
                    }
                    v3 -= v3
                } else {
                    if ((v87 * v88) == v88) {
                        v95 := GenS14{F0: 2147483647}
                        var v96 int8 = v88
                        _ = v95
                        _ = v96
                    } else {
                        var v97 uint16 = 42
                        var v98 string = v2
                        _ = v97
                        _ = v98
                    }
                    v1 = append(v1, "hello")
                    v4 -= (v11 << ((2) & 63))
                }
                v99 := GenS10{F0: 0}
                _ = fi91
                _ = v92
                _ = v99
            }
            fmt.Println(v11, v3, v88)
            _ = fi90
        }
    }
    var v100 uint = 0
    for fi101 := 0; fi101 < 12; fi101++ {
        v1 = append(v1, v2)
        for fi102 := 0; fi102 < 1; fi102++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi102
        }
        _ = fi101
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v100
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v87
    _ = v88
}
