package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 bool; F1 uint8; F2 int }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS13 struct { F0 int32; F1 bool }
type GenS14 struct { F0 int; F1 uint; F2 uint8 }
type GenS15 struct { F0 bool; F1 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = true
    var v3 string = "foo bar"
    if v2 {
        var v4 int = 953
        _ = v4
    } else {
        v2 = true
        if (v2 || false) {
            fmt.Println(v1)
        } else {
            if (0 != 0) {
                if (-128 == -28) {
                    var v5 float64 = v1
                    _ = v5
                }
                if v2 {
                    if (32768 > -757) {
                        v3 = v3
                        if v2 {
                            v6 := []int{2147483647, 1, -1}
                            if true {
                                switch 42 {
                                case -617:
                                    fmt.Println(v2, v3, v1)
                                case 131:
                                    v2 = ((v2 && v2) && (42 != 1000))
                                case -2147483648:
                                    for fi7 := 0; fi7 < 15; fi7++ {
                                        if v2 {
                                            if true {
                                                var v8 int64 = -1
                                                var v9 int64 = v8
                                                var v10 int = (fi7 & (fi7 - (fi7 >> ((1) & 63))))
                                                _ = v10
                                                _ = v8
                                                _ = v9
                                            } else {
                                                if false {
                                                    fmt.Println(v1, fi7, v3)
                                                } else {
                                                    switch v3 {
                                                    case "foo bar":
                                                        if ((639 <= 42) && (42 != 9223372036854775807)) {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v3 = ""
                                                        } else {
                                                            switch fi7 {
                                                            case -860:
                                                                for fi11 := 0; fi11 < 20; fi11++ {
                                                                    if true {
                                                                        var v12 uint8 = 104
                                                                        switch -1 {
                                                                        case -2147483648:
                                                                            var v13 uint64 = 1
                                                                            _ = v13
                                                                        case 0:
                                                                            fmt.Println(fi11, v12)
                                                                        default:
                                                                            if (v2 || v2) {
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    if v2 {
                                                                                        if true {
                                                                                            var v14 bool = (((3.14159 * float32(v1)) >= 0.1) && v2)
                                                                                            _ = v14
                                                                                        } else {
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = fi11
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v6 = append(v6, (-(fi11)))
                                                                                }
                                                                            }
                                                                        }
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        _ = v12
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v2 = false
                                                                    case "a":
                                                                        v3 = ((("hello" + ((v3 + v3) + "\n\t")) + v3) + ((v3 + "hello") + v3))
                                                                    case "hello":
                                                                        switch fi7 {
                                                                        case 128:
                                                                            if (v2 && v2) {
                                                                                if v2 {
                                                                                    switch fi11 {
                                                                                    case -276:
                                                                                        v15 := GenS1{F0: -1, F1: 0}
                                                                                        _ = v15
                                                                                    }
                                                                                }
                                                                                if v2 {
                                                                                    v16 := GenS1{F0: 123, F1: 961}
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    _ = v16
                                                                                } else {
                                                                                    v17 := GenS1{F0: 0, F1: 1}
                                                                                    switch fi11 {
                                                                                    case 1:
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                        case "hello":
                                                                                            if ((!false) || true) {
                                                                                                switch 4294967295 {
                                                                                                case 32768:
                                                                                                    v1 = v1
                                                                                                case 4294967295:
                                                                                                    fmt.Println(fi7, v2)
                                                                                                }
                                                                                                if (3.14159 >= v1) {
                                                                                                    v3 = ((v3 + ("" + (v3 + v3))) + v3)
                                                                                                    fmt.Println(fi11, v1)
                                                                                                    switch (-(fi11)) {
                                                                                                    case 127:
                                                                                                        fmt.Println(v3, fi7)
                                                                                                        fmt.Println(fi7, v1, v2)
                                                                                                        v1 = (v1 + 3.14159)
                                                                                                    case 9223372036854775807:
                                                                                                        if false {
                                                                                                            var v18 uint32 = 42
                                                                                                            v19 := GenS1{F0: 127, F1: 255}
                                                                                                            _ = v18
                                                                                                            _ = v19
                                                                                                        }
                                                                                                    default:
                                                                                                        if len(v6) > 0 {
                                                                                                            v6[0] = (fi11 * 32768)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v20 int = fi11
                                                                                                if false {
                                                                                                    switch (v3 + v3) {
                                                                                                    case "foo bar":
                                                                                                        if true {
                                                                                                            if (v2 || ((("hello" + v3) >= ("a" + "foo bar")) && true)) {
                                                                                                                v20 += v20
                                                                                                            }
                                                                                                            var v21 uint8 = 112
                                                                                                            if (65535 != 2) {
                                                                                                                var v22 int64 = 0
                                                                                                                _ = v22
                                                                                                            }
                                                                                                            _ = v21
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v1 -= float64(((1.5 - ((0.1 / 1.5) * 0.0)) - (1.5 * (3.14159 + 0.0))))
                                                                                                                if v2 {
                                                                                                                    v2 = v2
                                                                                                                    v20 -= (-(v20))
                                                                                                                    if len(v6) > 0 {
                                                                                                                        v6[0] = (^(fi11))
                                                                                                                    }
                                                                                                                    switch (v3 + "\n\t") {
                                                                                                                    case "a":
                                                                                                                        fmt.Println(fi11, v3)
                                                                                                                    case "hello":
                                                                                                                        switch (v3 + v3) {
                                                                                                                        case "hello":
                                                                                                                            var v23 bool = v2
                                                                                                                            _ = v23
                                                                                                                        case "":
                                                                                                                            fmt.Println(v2, v20)
                                                                                                                        case "foo bar":
                                                                                                                            v6 = append(v6, v20)
                                                                                                                        default:
                                                                                                                            v24 := GenS1{F0: 127, F1: 65535}
                                                                                                                            _ = v24
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        var v25 int = (fi7 >> ((18446744073709551615) & 63))
                                                                                                                        _ = v25
                                                                                                                    default:
                                                                                                                        v26 := GenS2{F0: -1, F1: 514, F2: 32768}
                                                                                                                        var v27 float64 = 1e10
                                                                                                                        _ = v26
                                                                                                                        _ = v27
                                                                                                                    }
                                                                                                                    fmt.Println(v20, fi11, fi11)
                                                                                                                }
                                                                                                                var v28 int32 = 0
                                                                                                                _ = v28
                                                                                                            }
                                                                                                            fmt.Println(fi7)
                                                                                                            v29 := GenS3{F0: 1, F1: 891}
                                                                                                            if ((v1 - 0.0) >= v1) {
                                                                                                                v2 = (v2 || (-516 != fi7))
                                                                                                                if v2 {
                                                                                                                    switch v3 {
                                                                                                                    case "a":
                                                                                                                        v3 = ""
                                                                                                                    default:
                                                                                                                        v2 = v2
                                                                                                                    }
                                                                                                                }
                                                                                                                switch fi7 {
                                                                                                                case 32767:
                                                                                                                    if v2 {
                                                                                                                        fmt.Println(fi11, v1)
                                                                                                                        if true {
                                                                                                                            var v30 int8 = -21
                                                                                                                            if len(v6) >= 2 {
                                                                                                                                v6 = v6[1:len(v6)]
                                                                                                                            }
                                                                                                                            _ = v30
                                                                                                                        } else {
                                                                                                                            var v31 uint8 = 1
                                                                                                                            var v32 int = fi7
                                                                                                                            v3 = v3
                                                                                                                            _ = v31
                                                                                                                            _ = v32
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 32768:
                                                                                                                    v1 = v1
                                                                                                                case 128:
                                                                                                                    v1 += 1e10
                                                                                                                    v6 = append(v6, (v20 << ((65535) & 63)))
                                                                                                                    v3 = v3
                                                                                                                }
                                                                                                            }
                                                                                                            switch v3 {
                                                                                                            case "\n\t":
                                                                                                                v33 := GenS4{F0: 192, F1: 2}
                                                                                                                _ = v33
                                                                                                            case "a":
                                                                                                                if v2 {
                                                                                                                    v34 := GenS5{F0: (127 >= 127), F1: 31, F2: (^(fi11))}
                                                                                                                    var v35 uint = 1
                                                                                                                    _ = v34
                                                                                                                    _ = v35
                                                                                                                } else {
                                                                                                                    switch fi7 {
                                                                                                                    case -2147483648:
                                                                                                                        var v36 uint = 0
                                                                                                                        _ = v36
                                                                                                                    }
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                fmt.Println(v1)
                                                                                                            default:
                                                                                                                fmt.Println(fi7, v20)
                                                                                                            }
                                                                                                            _ = v29
                                                                                                        }
                                                                                                    default:
                                                                                                        var v37 uint64 = 423
                                                                                                        switch (v3 + "\n\t") {
                                                                                                        case "":
                                                                                                            fmt.Println(v1, fi11)
                                                                                                        }
                                                                                                        if (((-2.25 + v1) > float64(-2.25)) && true) {
                                                                                                            switch ("hello" + "") {
                                                                                                            case "\n\t":
                                                                                                                _ = len(v6)
                                                                                                            default:
                                                                                                                fmt.Println(v2, fi7, v37)
                                                                                                            }
                                                                                                            var v38 uint = 0
                                                                                                            _ = v38
                                                                                                        } else {
                                                                                                            v6 = append(v6, (v20 * (fi7 % (((fi11 % (((fi7 * fi7)) | 1))) | 1))))
                                                                                                            if false {
                                                                                                                if (32767 != 65535) {
                                                                                                                    if (255 > 2) {
                                                                                                                        var v39 uint8 = 132
                                                                                                                        v2 = (uint16(v39) <= 32767)
                                                                                                                        _ = v39
                                                                                                                    }
                                                                                                                    if (1 >= 4294967295) {
                                                                                                                        if false {
                                                                                                                            v3 = v3
                                                                                                                            v1 = v1
                                                                                                                            var v40 uint32 = 42
                                                                                                                            _ = v40
                                                                                                                        }
                                                                                                                        v6 = append(v6, (fi7 & 32768))
                                                                                                                    } else {
                                                                                                                        if v2 {
                                                                                                                            v20 = (^(fi11))
                                                                                                                            var v41 string = "foo bar"
                                                                                                                            _ = v41
                                                                                                                        }
                                                                                                                        fmt.Println(v3, v37)
                                                                                                                        if false {
                                                                                                                            v42 := GenS5{F0: v2, F1: 2, F2: -537}
                                                                                                                            _ = v42
                                                                                                                        } else {
                                                                                                                            var v43 uint8 = 2
                                                                                                                            _ = v43
                                                                                                                        }
                                                                                                                        var v44 uint = 65535
                                                                                                                        _ = v44
                                                                                                                    }
                                                                                                                    switch 32767 {
                                                                                                                    case 10:
                                                                                                                        v37 = v37
                                                                                                                    case -4:
                                                                                                                        _ = len(v6)
                                                                                                                    default:
                                                                                                                        if (v3 < (((((("foo bar" + v3) + v3) + "foo bar") + v3) + v3) + (v3 + v3))) {
                                                                                                                            v45 := GenS4{F0: 94, F1: 886}
                                                                                                                            _ = v45
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v46 := GenS1{F0: 42, F1: 2}
                                                                                                                    _ = v46
                                                                                                                }
                                                                                                                var v47 int32 = 0
                                                                                                                _ = v47
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v37
                                                                                                    }
                                                                                                    v2 = (1 != -9223372036854775808)
                                                                                                    v6 = append(v6, (v20 - v20))
                                                                                                    v3 = (v3 + v3)
                                                                                                    v2 = false
                                                                                                    var v48 int16 = 0
                                                                                                    switch v3 {
                                                                                                    case "hello":
                                                                                                        v49 := GenS5{F0: v2, F1: 111, F2: (fi11 | fi11)}
                                                                                                        if v2 {
                                                                                                            v1 = v1
                                                                                                            fmt.Println(v20, v2, v2)
                                                                                                        } else {
                                                                                                            if len(v6) >= 2 {
                                                                                                                v6 = v6[1:len(v6)]
                                                                                                            }
                                                                                                            if len(v6) > 0 {
                                                                                                                v6[0] = 32768
                                                                                                            }
                                                                                                            fmt.Println(v20, fi7)
                                                                                                        }
                                                                                                        _ = v49
                                                                                                    default:
                                                                                                        if ((v2 && v2) || (!v2)) {
                                                                                                            v1 -= v1
                                                                                                            if (1 < 127) {
                                                                                                                fmt.Println(fi11, v1)
                                                                                                                if len(v6) > 0 {
                                                                                                                    v6[0] = v20
                                                                                                                }
                                                                                                                v1 -= ((((v1 + v1) - float64(3.14159)) / float64(3.14159)) * float64((0.0 + 0.1)))
                                                                                                            } else {
                                                                                                                v20 -= (fi11 << ((255) & 63))
                                                                                                                v6 = append(v6, v20)
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v3)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    var v50 float32 = 1e10
                                                                                                    var v51 bool = ((!((0.0 - v1) >= v1)) || (-1 > 1))
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                }
                                                                                                _ = v20
                                                                                            }
                                                                                        }
                                                                                    case 122:
                                                                                        v52 := GenS6{F0: 1000, F1: 4294967295}
                                                                                        var v53 uint64 = 2
                                                                                        _ = v52
                                                                                        _ = v53
                                                                                    default:
                                                                                        fmt.Println(fi11, fi11, fi11)
                                                                                    }
                                                                                    v1 = (float64(float32((v1 + v1))) + (v1 * 3.14159))
                                                                                    _ = v17
                                                                                }
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = fi7
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v54 int16 = -454
                                                                    _ = fi11
                                                                    _ = v54
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v1 += v1
                                            }
                                            v2 = v2
                                        }
                                        var v55 float64 = -2.25
                                        _ = len(v6)
                                        _ = fi7
                                        _ = v55
                                    }
                                default:
                                    var v56 string = "a"
                                    _ = v56
                                }
                                for fi57 := 0; fi57 < 2; fi57++ {
                                    if len(v6) > 0 {
                                        v6[0] = fi57
                                    }
                                    var v58 int32 = 0
                                    fmt.Println(v2)
                                    _ = fi57
                                    _ = v58
                                }
                            }
                            _ = v6
                        } else {
                            switch -1 {
                            case 127:
                                if false {
                                    if (-2147483648 <= 0) {
                                        var v59 int = 0
                                        for fi60 := 0; fi60 < 9; fi60++ {
                                            v61 := []int{}
                                            fmt.Println(v3)
                                            v62 := GenS3{F0: 2, F1: 758}
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                        }
                                        v63 := GenS4{F0: 128, F1: 65535}
                                        _ = v59
                                        _ = v63
                                    }
                                    var v64 bool = (-1 != -13)
                                    fmt.Println(v2, v2)
                                    _ = v64
                                } else {
                                    for fi65 := 0; fi65 < 1; fi65++ {
                                        v3 = v3
                                        _ = fi65
                                    }
                                }
                            case 32767:
                                var v66 uint64 = 4294967295
                                _ = v66
                            default:
                                var v67 uint16 = 602
                                _ = v67
                            }
                        }
                    }
                }
            }
        }
    }
    if false {
        if (1e10 < float32(float64((float32(((0.1 / (v1 * v1)) * float64(((((-2.25 * 0.1) * 0.1) - 1.5) / ((0.0 * 0.0) * (float32(3.14159) / (float32((float64((3.14159 - 1.5)) / (v1 / v1))) / 3.14159))))))) / (0.0 - -2.25))))) {
            fmt.Println(v1, v1, v1)
        }
        v68 := []int{32767}
        for fi69 := 0; fi69 < 12; fi69++ {
            var v70 bool = (1 == 15)
            _ = fi69
            _ = v70
        }
        _ = v68
    }
    if (!true) {
        var v71 uint16 = 837
        var v72 uint8 = 255
        _ = v71
        _ = v72
    } else {
        var v73 int64 = 0
        _ = v73
    }
    for fi74 := 0; fi74 < 14; fi74++ {
        var v75 int = (fi74 / ((628) | 1))
        _ = fi74
        _ = v75
    }
    v76 := GenS7{F0: 2147483647}
    fmt.Println(v1, v3, v2)
    for fi77 := 0; fi77 < 19; fi77++ {
        if (42 <= -996) {
            v78 := GenS6{F0: -32768, F1: 1}
            v2 = v2
            for fi79 := 0; fi79 < 15; fi79++ {
                v3 = v3
                _ = fi79
            }
            _ = v78
        }
        var v80 uint32 = 2147483647
        var v81 uint8 = 255
        _ = fi77
        _ = v80
        _ = v81
    }
    if (-349 != 0) {
        switch -2147483648 {
        case 622:
            if ((18446744073709551615 <= 635) && (-9223372036854775808 == 0)) {
                if (((v3 + v3) != v3) || (true && false)) {
                    v82 := GenS6{F0: 1000, F1: 4294967295}
                    for fi83 := 0; fi83 < 4; fi83++ {
                        v84 := GenS6{F0: 349, F1: 1}
                        v85 := GenS3{F0: 302, F1: 1}
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    _ = v82
                }
                for fi86 := 0; fi86 < 18; fi86++ {
                    v3 = "foo bar"
                    v3 = ("a" + ("\n\t" + v3))
                    _ = fi86
                }
                for fi87 := 0; fi87 < 1; fi87++ {
                    v88 := GenS1{F0: 127, F1: 255}
                    switch v3 {
                    case "hello":
                        var v89 uint8 = 209
                        _ = v89
                    case "foo bar":
                        v2 = v2
                    default:
                        if v2 {
                            switch v3 {
                            case "a":
                                v90 := []int{fi87, (fi87 << ((1) & 63)), 181, (fi87 - fi87)}
                                _ = v90
                            case "foo bar":
                                v1 -= v1
                            default:
                                v91 := GenS3{F0: 0, F1: 84}
                                _ = v91
                            }
                            fmt.Println(v2, v3, fi87)
                        }
                    }
                    fmt.Println(v3, v2)
                    _ = fi87
                    _ = v88
                }
            } else {
                v92 := GenS1{F0: -84, F1: 0}
                _ = v92
            }
        case 0:
            v93 := GenS8{F0: 0, F1: 134}
            _ = v93
        case 4294967295:
            fmt.Println(v1, v2)
        }
        v94 := GenS9{F0: v3}
        _ = v94
    }
    v1 -= v1
    v95 := GenS10{F0: -756, F1: 65535}
    var v96 int32 = 1
    v97 := GenS9{F0: (v3 + ((v3 + v3) + (v3 + ("hello" + "hello"))))}
    v2 = v2
    fmt.Println(v1)
    for fi98 := 0; fi98 < 18; fi98++ {
        var v99 string = v3
        if (v2 || v2) {
            v2 = v2
            for fi100 := 0; fi100 < 15; fi100++ {
                if (((v96 ^ v96) >= (^(v96))) && v2) {
                    v101 := []int{}
                    v102 := GenS11{F0: v99, F1: 0}
                    _ = v101
                    _ = v102
                }
                v103 := GenS2{F0: 341, F1: (v96 - v96), F2: 421}
                if v2 {
                    v104 := GenS2{F0: 0, F1: v96, F2: 128}
                    switch (v99 + "hello") {
                    case "foo bar":
                        var v105 int8 = 0
                        _ = v105
                    case "\n\t":
                        v96 = (v96 << ((740) & 63))
                        fmt.Println(v2, v2, v2)
                    case "a":
                        v106 := []int{fi100}
                        _ = v106
                    }
                    _ = v104
                } else {
                    switch -715 {
                    case 534:
                        if true {
                            v2 = (!(true || true))
                        } else {
                            v1 += (v1 / float64(3.14159))
                        }
                    case 32767:
                        fmt.Println(v96, fi100)
                    }
                    v2 = (17 < -128)
                    v1 = 0.0
                    if (0 < -1) {
                        v107 := GenS12{F0: -32768, F1: 18446744073709551615, F2: 948}
                        fmt.Println(fi98, v99, v96)
                        _ = v107
                    }
                    v108 := GenS5{F0: false, F1: 1, F2: 9223372036854775807}
                    _ = v108
                }
                _ = fi100
                _ = v103
            }
        }
        _ = fi98
        _ = v99
    }
    v109 := []int{}
    _ = len(v109)
    if (9223372036854775807 != 32767) {
        fmt.Println(v3, v96)
    } else {
        v96 += v96
        switch 2147483647 {
        case 2147483647:
            if (true || (2147483648 >= 743)) {
                v110 := GenS3{F0: 4294967295, F1: 727}
                _ = v110
            }
        default:
            if v2 {
                for fi111 := 0; fi111 < 10; fi111++ {
                    if false {
                        if (float64((-2.25 + 3.14159)) > 1.5) {
                            if v2 {
                                var v112 int64 = 1
                                fmt.Println(v1)
                                _ = v112
                            } else {
                                switch v3 {
                                case "":
                                    _ = len(v109)
                                case "\n\t":
                                    if v2 {
                                        switch ((v3 + v3) + v3) {
                                        case "\n\t":
                                            v3 = (v3 + (((("" + v3) + ((v3 + "\n\t") + "foo bar")) + v3) + (("\n\t" + v3) + "foo bar")))
                                        case "foo bar":
                                            if (-1 == -128) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch ("foo bar" + v3) {
                                                    case "a":
                                                        var v113 int32 = (v96 >> ((255) & 63))
                                                        _ = v113
                                                    }
                                                case "":
                                                    v96 -= v96
                                                }
                                                v114 := GenS13{F0: (v96 & (v96 - v96)), F1: (v1 < v1)}
                                                _ = v114
                                            } else {
                                                for fi115 := 0; fi115 < 11; fi115++ {
                                                    v3 = (v3 + v3)
                                                    var v116 int64 = 9223372036854775807
                                                    _ = fi115
                                                    _ = v116
                                                }
                                                v117 := GenS13{F0: v96, F1: true}
                                                _ = v117
                                            }
                                        case "":
                                            v1 -= 0.1
                                        default:
                                            if false {
                                                var v118 string = v3
                                                v109 = append(v109, (fi111 + (fi111 - 906)))
                                                v119 := GenS8{F0: -9223372036854775808, F1: 255}
                                                _ = v118
                                                _ = v119
                                            }
                                        }
                                        switch (fi111 / ((-912) | 1)) {
                                        case -2147483648:
                                            for fi120 := 0; fi120 < 4; fi120++ {
                                                if (!v2) {
                                                    if len(v109) >= 2 {
                                                        v109 = v109[1:len(v109)]
                                                    }
                                                    var v121 uint32 = 92
                                                    _ = v121
                                                } else {
                                                    v122 := GenS1{F0: -128, F1: 32767}
                                                    v2 = (!true)
                                                    _ = v122
                                                }
                                                if len(v109) >= 2 {
                                                    v109 = v109[1:len(v109)]
                                                }
                                                _ = fi120
                                            }
                                        case 127:
                                            var v123 float64 = 0.0
                                            _ = v123
                                        case 42:
                                            var v124 uint64 = 913
                                            _ = v124
                                        }
                                        var v125 uint16 = 385
                                        _ = v125
                                    } else {
                                        var v126 uint8 = 245
                                        v127 := GenS1{F0: -63, F1: 65535}
                                        v3 = v3
                                        _ = v126
                                        _ = v127
                                    }
                                default:
                                    switch 556 {
                                    case -949:
                                        var v128 int8 = 100
                                        _ = v128
                                    case 42:
                                        var v129 int = (fi111 / ((fi111) | 1))
                                        _ = v129
                                    case -1:
                                        v1 -= float64(((0.0 + (0.0 / ((3.14159 * (-2.25 / (0.1 - (1.5 * float32((float64(-2.25) * (v1 * float64(0.0)))))))) / -2.25))) / 1.0))
                                    }
                                }
                                for fi130 := 0; fi130 < 3; fi130++ {
                                    v1 = (float64(float32(0.0)) / ((v1 - v1) + (v1 / ((-2.25 * v1) * v1))))
                                    _ = fi130
                                }
                                for fi131 := 0; fi131 < 4; fi131++ {
                                    var v132 uint16 = 0
                                    _ = fi131
                                    _ = v132
                                }
                            }
                            for fi133 := 0; fi133 < 4; fi133++ {
                                v3 = "hello"
                                v134 := GenS14{F0: fi133, F1: 2, F2: 2}
                                v96 -= 1
                                _ = fi133
                                _ = v134
                            }
                        } else {
                            for fi135 := 0; fi135 < 5; fi135++ {
                                switch "a" {
                                case "\n\t":
                                    v136 := GenS15{F0: v2, F1: 0}
                                    _ = v136
                                }
                                v137 := GenS15{F0: v2, F1: 798}
                                v138 := GenS8{F0: -1, F1: 841}
                                v96 += v96
                                _ = fi135
                                _ = v137
                                _ = v138
                            }
                            var v139 int64 = -1
                            _ = v139
                        }
                        var v140 float64 = (v1 / 0.1)
                        v141 := GenS6{F0: 1000, F1: 18446744073709551615}
                        _ = v140
                        _ = v141
                    }
                    _ = fi111
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v109:", len(v109), v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v76
    _ = v95
    _ = v96
    _ = v97
}
