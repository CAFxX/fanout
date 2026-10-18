package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int16; F2 string }
type GenS6 struct { F0 uint8; F1 int64; F2 int32 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 2
    var v3 bool = false
    var v4 string = "\n\t"
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 4294967295, F1: 2, F2: 0}
    switch 825 {
    case -256:
        if v3 {
            fmt.Println(v3, v3)
        }
    case -9223372036854775808:
        var v7 uint8 = 42
        _ = v7
    case 0:
        fmt.Println(v3, v1, v4)
    default:
        v8 := GenS1{F0: 127, F1: uint(v1), F2: 18446744073709551615}
        _ = v8
    }
    v1 = (^(v1))
    var v9 bool = (!v3)
    v2 -= v2
    if false {
        var v10 int = 821
        _ = v10
    } else {
        fmt.Println(v4)
        fmt.Println(v1, v2)
    }
    switch 128 {
    case -84:
        if ((!false) || true) {
            v11 := GenS1{F0: 4294967295, F1: 1, F2: 42}
            _ = v11
        } else {
            v12 := []int{-2147483648, 4294967295, 137, 0}
            v13 := GenS2{F0: 127}
            if v3 {
                switch ((v4 + ((v4 + ((v4 + (v4 + v4)) + ("\n\t" + "\n\t"))) + v4)) + v4) {
                case "foo bar":
                    var v14 uint8 = (v2 - 255)
                    _ = v14
                case "a":
                    v12 = append(v12, -2147483648)
                case "":
                    switch 806 {
                    case 4294967295:
                        v3 = false
                    case -59:
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v5[v4] = (v1 % ((v1) | 1))
                            switch -2147483648 {
                            case 2147483647:
                                v2 += v2
                            case -676:
                                if true {
                                    for fi16 := 0; fi16 < 3; fi16++ {
                                        switch ((v4 + (v4 + ((v4 + "a") + (v4 + ("a" + v4))))) + "") {
                                        case "hello":
                                            if v3 {
                                                var v17 int = fi16
                                                _ = len(v12)
                                                if v9 {
                                                    v2 -= v2
                                                    v18 := GenS1{F0: (fi15 | 2147483647), F1: 255, F2: 342}
                                                    if (v3 || v3) {
                                                        v2 -= v2
                                                    } else {
                                                        var v19 uint32 = (v1 << ((255) & 63))
                                                        v20 := GenS1{F0: fi15, F1: 42, F2: 0}
                                                        _ = v19
                                                        _ = v20
                                                    }
                                                    _ = v18
                                                }
                                                _ = v17
                                            } else {
                                                var v21 uint64 = 169
                                                v9 = v9
                                                delete(v5, (((((v4 + v4) + "a") + v4) + "") + ((((v4 + (v4 + v4)) + v4) + ((v4 + v4) + "\n\t")) + v4)))
                                                _ = v21
                                            }
                                        }
                                        fmt.Println(v3, v1)
                                        v22 := GenS3{F0: 42, F1: (v1 - 4294967295)}
                                        _ = fi16
                                        _ = v22
                                    }
                                }
                            case 9223372036854775807:
                                v3 = (("a" <= v4) || (9223372036854775807 > 42))
                            default:
                                v2 += v2
                            }
                            if true {
                                v23 := GenS2{F0: -21}
                                if true {
                                    v24 := GenS1{F0: (fi15 + -495), F1: 0, F2: 815}
                                    v25 := GenS4{F0: 120}
                                    _ = v24
                                    _ = v25
                                }
                                v26 := GenS5{F0: (^(v2)), F1: 1000, F2: "foo bar"}
                                _ = v23
                                _ = v26
                            }
                            _ = fi15
                        }
                    case -243:
                        if (127 > 0) {
                            switch 42 {
                            case -1:
                                var v27 int16 = 42
                                _ = v27
                            case -745:
                                fmt.Println(v3, v9)
                            case 2147483647:
                                v28 := GenS6{F0: v2, F1: -9223372036854775808, F2: -2147483648}
                                _ = v28
                            }
                            if (!true) {
                                fmt.Println(v3, v9)
                                var v29 float32 = (((0.0 * (1e10 * float32(0.0))) + ((((1e10 - -2.25) * 1e10) / 1.5) + 1e10)) + 1e10)
                                _ = v29
                            } else {
                                if ((v1 >= (v1 % ((v1) | 1))) || false) {
                                    for fi30 := 0; fi30 < 12; fi30++ {
                                        _ = len(v12)
                                        if false {
                                            if (1 != 4294967295) {
                                                v31 := GenS7{F0: 0, F1: 1}
                                                _ = v31
                                            }
                                            fmt.Println(v2, v2)
                                        } else {
                                            for fi32 := 0; fi32 < 19; fi32++ {
                                                if (426 == 42) {
                                                    v4 = v4
                                                }
                                                v9 = (776 > 42)
                                                _ = fi32
                                            }
                                            v33 := GenS5{F0: 2, F1: -66, F2: ("" + v4)}
                                            _ = v33
                                        }
                                        var v34 int8 = 42
                                        _ = fi30
                                        _ = v34
                                    }
                                    switch -9223372036854775808 {
                                    case 42:
                                        var v35 uint8 = 127
                                        _ = v35
                                    case 4294967295:
                                        v36 := v5[(v4 + ((v4 + (v4 + v4)) + v4))]
                                        _ = v36
                                    case 1:
                                        switch v4 {
                                        case "a":
                                            var v37 int8 = -112
                                            _ = v37
                                        case "":
                                            if v9 {
                                                v1 = (v1 ^ v1)
                                                for fi38 := 0; fi38 < 9; fi38++ {
                                                    v39 := GenS8{F0: v9, F1: 32767}
                                                    v9 = true
                                                    switch fi38 {
                                                    case 990:
                                                        var v40 float32 = (-2.25 + 0.0)
                                                        _ = v40
                                                    case 205:
                                                        for fi41 := 0; fi41 < 20; fi41++ {
                                                            _ = len(v12)
                                                            v9 = false
                                                            switch (v4 + v4) {
                                                            case "\n\t":
                                                                _ = len(v12)
                                                                var v42 uint32 = (^(v1))
                                                                v4 = ("a" + "a")
                                                                _ = v42
                                                            }
                                                            _ = fi41
                                                        }
                                                    case 1:
                                                        switch fi38 {
                                                        case 760:
                                                            var v43 uint32 = (v1 - 479)
                                                            _ = v43
                                                        case -910:
                                                            v44 := GenS2{F0: -1}
                                                            _ = v44
                                                        case -926:
                                                            var v45 uint16 = 301
                                                            _ = v45
                                                        default:
                                                            v1 = v1
                                                        }
                                                    }
                                                    _ = fi38
                                                    _ = v39
                                                }
                                                switch v4 {
                                                case "\n\t":
                                                    var v46 int32 = 42
                                                    _ = v46
                                                default:
                                                    var v47 float32 = (-2.25 - 1.5)
                                                    _ = v47
                                                }
                                            } else {
                                                var v48 bool = false
                                                if (!v9) {
                                                    if len(v12) > 0 {
                                                        v12[0] = 4294967295
                                                    }
                                                    _ = len(v12)
                                                }
                                                _ = v48
                                            }
                                        case "hello":
                                            _ = len(v5)
                                        default:
                                            var v49 bool = v9
                                            _ = v49
                                        }
                                    default:
                                        if (!v3) {
                                            v12 = append(v12, 1)
                                            for fi50 := 0; fi50 < 2; fi50++ {
                                                fmt.Println(v4)
                                                for fi51 := 0; fi51 < 15; fi51++ {
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    var v52 uint32 = 65535
                                                    _ = fi51
                                                    _ = v52
                                                }
                                                v4 = "foo bar"
                                                _ = fi50
                                            }
                                        } else {
                                            switch "a" {
                                            case "":
                                                var v53 int = 673
                                                _ = v53
                                            case "foo bar":
                                                if v9 {
                                                    var v54 int16 = -1
                                                    v55 := GenS6{F0: (v2 - 253), F1: -758, F2: 1}
                                                    _ = v54
                                                    _ = v55
                                                }
                                            default:
                                                if (2 != (v2 % ((60) | 1))) {
                                                    var v56 int64 = 0
                                                    if (0.1 >= -2.25) {
                                                        var v57 int32 = -1
                                                        for fi58 := 0; fi58 < 5; fi58++ {
                                                            var v59 int16 = 1000
                                                            _ = fi58
                                                            _ = v59
                                                        }
                                                        _ = v57
                                                    } else {
                                                        var v60 float32 = ((1.5 + 0.0) - 3.14159)
                                                        _ = v60
                                                    }
                                                    for fi61 := 0; fi61 < 4; fi61++ {
                                                        v2 = 1
                                                        if (-1 < 100) {
                                                            var v62 uint64 = 2
                                                            for fi63 := 0; fi63 < 17; fi63++ {
                                                                v64 := GenS6{F0: (^(v2)), F1: (v56 << ((255) & 63)), F2: 42}
                                                                v2 = 255
                                                                var v65 uint32 = (v1 & (v1 & v1))
                                                                _ = fi63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            _ = v62
                                                        }
                                                        _ = fi61
                                                    }
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v9, v1)
                                                }
                                            }
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            for fi66 := 0; fi66 < 1; fi66++ {
                                                v4 = (("\n\t" + ("" + v4)) + v4)
                                                _ = fi66
                                            }
                                        }
                                    }
                                    var v67 int32 = 1
                                    _ = v67
                                } else {
                                    switch -457 {
                                    case -2147483648:
                                        if v9 {
                                            var v68 float32 = 1.5
                                            v1 = (v1 * v1)
                                            if v3 {
                                                v9 = v3
                                                switch 32768 {
                                                case 2147483647:
                                                    for fi69 := 0; fi69 < 6; fi69++ {
                                                        if v3 {
                                                            v3 = (((!true) && (v9 || ((v2 / (((v2 - v2)) | 1)) <= (v2 ^ v2)))) || false)
                                                        }
                                                        for fi70 := 0; fi70 < 6; fi70++ {
                                                            var v71 string = "\n\t"
                                                            v4 = ("\n\t" + (v71 + v71))
                                                            _ = fi70
                                                            _ = v71
                                                        }
                                                        _ = fi69
                                                    }
                                                default:
                                                    if v9 {
                                                        for fi72 := 0; fi72 < 15; fi72++ {
                                                            if v3 {
                                                                switch v4 {
                                                                case "\n\t":
                                                                    switch fi72 {
                                                                    case 127:
                                                                        var v73 int16 = 966
                                                                        _ = v73
                                                                    case -425:
                                                                        fmt.Println(v1, v3, fi72)
                                                                    case 4294967295:
                                                                        v2 += v2
                                                                    }
                                                                case "a":
                                                                    for fi74 := 0; fi74 < 19; fi74++ {
                                                                        fmt.Println(v9, v1)
                                                                        switch 127 {
                                                                        case 2147483647:
                                                                            v3 = (-9223372036854775808 != 0)
                                                                        default:
                                                                            v2 += v2
                                                                            if v9 {
                                                                                v4 = (v4 + (v4 + (v4 + v4)))
                                                                            }
                                                                        }
                                                                        switch ("foo bar" + (v4 + (("\n\t" + (("" + "\n\t") + ("\n\t" + v4))) + ("\n\t" + (v4 + (v4 + ((("" + v4) + v4) + (v4 + v4)))))))) {
                                                                        case "\n\t":
                                                                            v75 := GenS7{F0: -43, F1: 47}
                                                                            _ = v75
                                                                        case "":
                                                                            v76 := GenS9{F0: 311}
                                                                            _ = v76
                                                                        default:
                                                                            if ((v68 * ((v68 * (float32(0.0) + float32(0.0))) * float32(0.0))) < (-2.25 / (float32((3.14159 * 1.5)) * float32(((1e10 + 1e10) + -2.25))))) {
                                                                                fmt.Println(v4, v2, v3)
                                                                            } else {
                                                                                v77 := GenS6{F0: v2, F1: -9223372036854775808, F2: -662}
                                                                                _ = v77
                                                                            }
                                                                        }
                                                                        _ = fi74
                                                                    }
                                                                case "foo bar":
                                                                    if (v1 > (^(v1))) {
                                                                        _ = len(v5)
                                                                        var v78 uint64 = 0
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                            var v79 int8 = 0
                                                            if v3 {
                                                                switch (v4 + (("a" + v4) + "foo bar")) {
                                                                case "hello":
                                                                    var v80 int32 = -2147483648
                                                                    _ = v80
                                                                case "":
                                                                    var v81 int16 = 277
                                                                    _ = v81
                                                                default:
                                                                    v79 -= v79
                                                                }
                                                                v3 = (true || v9)
                                                            }
                                                            _ = fi72
                                                            _ = v79
                                                        }
                                                        for fi82 := 0; fi82 < 20; fi82++ {
                                                            v4 = ("\n\t" + (v4 + v4))
                                                            v68 = -2.25
                                                            _ = fi82
                                                        }
                                                    } else {
                                                        if len(v12) > 0 {
                                                            v12[0] = -1
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v68
                                        } else {
                                            fmt.Println(v9, v9, v4)
                                        }
                                    case 608:
                                        for fi83 := 0; fi83 < 1; fi83++ {
                                            for fi84 := 0; fi84 < 5; fi84++ {
                                                fmt.Println(v4, v2)
                                                _ = fi84
                                            }
                                            if false {
                                                for fi85 := 0; fi85 < 17; fi85++ {
                                                    var v86 float32 = 0.0
                                                    if v9 {
                                                        fmt.Println(v86, fi83)
                                                        var v87 string = v4
                                                        switch "a" {
                                                        case "foo bar":
                                                            fmt.Println(v9, v4, v4)
                                                        default:
                                                            var v88 int32 = 42
                                                            _ = v88
                                                        }
                                                        _ = v87
                                                    }
                                                    var v89 uint = 0
                                                    _ = fi85
                                                    _ = v86
                                                    _ = v89
                                                }
                                            } else {
                                                v90 := GenS3{F0: 1, F1: 862}
                                                if v3 {
                                                    var v91 int8 = 80
                                                    var v92 uint64 = 405
                                                    _ = v91
                                                    _ = v92
                                                }
                                                _ = v90
                                            }
                                            if false {
                                                _ = len(v5)
                                                for fi93 := 0; fi93 < 1; fi93++ {
                                                    if ((fi93 == fi83) || v3) {
                                                        var v94 float64 = -2.25
                                                        _ = v94
                                                    } else {
                                                        fmt.Println(v3, fi83)
                                                        switch v4 {
                                                        case "foo bar":
                                                            var v95 float32 = ((1.5 + (1e10 + 1e10)) * (1e10 * -2.25))
                                                            _ = v95
                                                        }
                                                        _ = len(v5)
                                                    }
                                                    _ = fi93
                                                }
                                            } else {
                                                for fi96 := 0; fi96 < 7; fi96++ {
                                                    var v97 int32 = -1
                                                    v4 = v4
                                                    _ = fi96
                                                    _ = v97
                                                }
                                            }
                                            _ = fi83
                                        }
                                    case 2147483647:
                                        var v98 uint64 = 0
                                        _ = v98
                                    }
                                }
                                if v3 {
                                    v9 = false
                                    v3 = false
                                } else {
                                    v99 := GenS7{F0: 1, F1: -1}
                                    _ = v99
                                }
                            }
                        }
                    default:
                        fmt.Println(v3, v3, v1)
                    }
                default:
                    var v100 int64 = -214
                    _ = v100
                }
                v1 += 467
                v101 := GenS10{F0: 518, F1: v1}
                _ = v101
            } else {
                v9 = (true || ("hello" == "\n\t"))
            }
            _ = v12
            _ = v13
        }
    case -755:
        for fi102 := 0; fi102 < 20; fi102++ {
            fmt.Println(v1, fi102, fi102)
            var v103 int8 = 0
            _ = fi102
            _ = v103
        }
    case 32767:
        v104 := GenS9{F0: v1}
        _ = v104
    }
    v9 = false
    v105 := v5[v4]
    var v106 int32 = 468
    if v3 {
        v3 = (v9 && (-64 > 0))
        if ((v106 << ((0) & 63)) >= 42) {
            v105 += v1
        }
        switch 4294967295 {
        case 4294967295:
            for fi107 := 0; fi107 < 6; fi107++ {
                fmt.Println(v2)
                for fi108 := 0; fi108 < 20; fi108++ {
                    v109 := GenS9{F0: 0}
                    var v110 int = fi107
                    v3 = (848 != 0)
                    _ = fi108
                    _ = v109
                    _ = v110
                }
                _ = fi107
            }
        case -789:
            if v9 {
                fmt.Println(v4, v105)
                _ = len(v5)
            }
        }
    }
    var v111 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v105
    _ = v106
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
