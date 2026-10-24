package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int; F1 string; F2 int }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 uint8; F1 int32; F2 string }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool; F1 uint; F2 int8 }

func main() {
    var v1 float64 = 1e10
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 233
    var v5 bool = true
    for fi6 := 0; fi6 < 13; fi6++ {
        switch (fi6 | (fi6 << ((2) & 63))) {
        case -536:
            for fi7 := 0; fi7 < 18; fi7++ {
                switch "" {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                case "":
                    if (!v5) {
                        v8 := GenS1{F0: 2}
                        _ = v8
                    }
                    var v9 uint8 = 255
                    var v10 bool = v5
                    _ = v10
                    _ = v9
                }
                if (2 == 4294967295) {
                    fmt.Println(fi6)
                    var v11 int = fi7
                    _ = v11
                }
                _ = fi7
            }
        case 223:
            switch "" {
            case "hello":
                v4 = (v4 ^ (v4 * v4))
            case "\n\t":
                switch (fi6 >> ((65535) & 63)) {
                case -403:
                    if (125 > (v4 * v4)) {
                        _ = len(v3)
                    } else {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v5 = (v4 > 1)
                            delete(v3, 2147483647)
                            if v5 {
                                switch (fi12 * fi6) {
                                case 32767:
                                    var v13 uint16 = 17
                                    _ = v13
                                case 296:
                                    if ((v1 / 1e10) >= ((v1 - float64(3.14159)) * v1)) {
                                        var v14 int32 = 49
                                        _ = v14
                                    }
                                default:
                                    switch fi12 {
                                    case -2147483648:
                                        v5 = (1 == 2)
                                    default:
                                        switch (fi12 * 2147483647) {
                                        case 127:
                                            if len(v2) > 0 {
                                                v2[0] = (fi6 << ((255) & 63))
                                            }
                                        default:
                                            var v15 string = "hello"
                                            _ = v15
                                        }
                                    }
                                }
                            } else {
                                if v5 {
                                    var v16 int32 = 922
                                    var v17 int16 = 145
                                    switch "hello" {
                                    case "a":
                                        v1 = (v1 / v1)
                                    case "\n\t":
                                        delete(v3, 0)
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                if v5 {
                                    v3[42] = 62
                                }
                                v1 = (3.14159 - ((float64(float32(v1)) * (1e10 * float64(0.0))) / 1.0))
                            }
                            _ = fi12
                        }
                    }
                default:
                    switch 73 {
                    case -593:
                        if v5 {
                            v4 = (v4 % (((v4 ^ (v4 & v4))) | 1))
                            for fi18 := 0; fi18 < 12; fi18++ {
                                switch "\n\t" {
                                case "hello":
                                    if (-698 == 42) {
                                        var v19 bool = (((fi6 & (fi18 | fi6)) == -313) && (("foo bar" <= "") && v5))
                                        var v20 int32 = 533
                                        v21 := GenS1{F0: 65535}
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 float64 = 3.14159
                                        v5 = false
                                        _ = v22
                                    }
                                case "a":
                                    if (4294967295 < 1) {
                                        v23 := GenS2{F0: 2}
                                        switch (("" + "") + "hello") {
                                        case "foo bar":
                                            switch fi18 {
                                            case 1:
                                                v24 := GenS3{F0: fi18, F1: 42, F2: 101}
                                                _ = v24
                                            }
                                        case "":
                                            v25 := GenS4{F0: 18446744073709551615}
                                            _ = v25
                                        case "a":
                                            v1 = (v1 + 1e10)
                                        default:
                                            _ = len(v2)
                                        }
                                        _ = v23
                                    }
                                case "":
                                    _ = len(v3)
                                default:
                                    v1 = 1e10
                                }
                                if (-1 != 78) {
                                    var v26 string = ((("" + ((("hello" + "foo bar") + "") + ("a" + "a"))) + "") + "")
                                    _ = v26
                                } else {
                                    v27 := GenS4{F0: 1}
                                    var v28 string = "a"
                                    if ((fi6 << ((142) & 63)) == (fi6 + (fi18 / (((fi6 << ((42) & 63))) | 1)))) {
                                        fmt.Println(v5, fi6)
                                    } else {
                                        fmt.Println(v28, v1)
                                        fmt.Println(fi18)
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                if false {
                                    v5 = (65535 <= 255)
                                } else {
                                    fmt.Println(v4, v1)
                                    if v5 {
                                        var v29 string = "foo bar"
                                        _ = v29
                                    }
                                    var v30 int = (-(fi6))
                                    if (!v5) {
                                        if (1 <= 628) {
                                            v31 := GenS4{F0: 1}
                                            _ = v31
                                        } else {
                                            _ = len(v2)
                                        }
                                        v32 := GenS5{F0: "foo bar", F1: 1}
                                        switch 127 {
                                        case 128:
                                            v33 := GenS6{F0: (fi6 * (^(fi18))), F1: "hello", F2: (fi18 + (-(v30)))}
                                            _ = v33
                                        case 42:
                                            v4 -= v4
                                        case -9223372036854775808:
                                            fmt.Println(v5)
                                        default:
                                            v1 += 0.1
                                        }
                                        if false {
                                            if v5 {
                                                if (!(!("a" == ""))) {
                                                    if true {
                                                        var v34 int32 = 664
                                                        var v35 bool = (32767 > 65535)
                                                        if v35 {
                                                            v4 -= v4
                                                            var v36 int16 = 0
                                                            v37 := GenS7{F0: -738, F1: (-(v36))}
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            switch ("\n\t" + ("" + (("a" + "\n\t") + "a"))) {
                                                            case "foo bar":
                                                                switch 221 {
                                                                case -644:
                                                                    if (127 >= -1) {
                                                                        if v5 {
                                                                            var v38 string = "\n\t"
                                                                            _ = len(v3)
                                                                            _ = v38
                                                                        } else {
                                                                            if v35 {
                                                                                var v39 int64 = 9223372036854775807
                                                                                v34 = (^(v34))
                                                                                _ = v39
                                                                            }
                                                                            fmt.Println(fi18)
                                                                            var v40 uint16 = 0
                                                                            _ = v40
                                                                        }
                                                                        var v41 int = (fi18 % ((-9223372036854775808) | 1))
                                                                        _ = v41
                                                                    } else {
                                                                        _ = len(v2)
                                                                        fmt.Println(fi18)
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (fi6 >> ((684) & 63))
                                                                        }
                                                                    }
                                                                    switch 82 {
                                                                    case -398:
                                                                        var v42 float32 = (1e10 / ((1e10 - -2.25) / 1.5))
                                                                        _ = v42
                                                                    case 1:
                                                                        if (!(32767 <= (fi18 & 42))) {
                                                                            v43 := GenS7{F0: 9223372036854775807, F1: 42}
                                                                            _ = v43
                                                                        }
                                                                    default:
                                                                        fmt.Println(v34)
                                                                    }
                                                                case 4294967295:
                                                                    v4 = v4
                                                                default:
                                                                    if v35 {
                                                                        switch ((("" + "foo bar") + ("\n\t" + "a")) + ("\n\t" + "a")) {
                                                                        case "a":
                                                                            var v44 uint = 42
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                }
                                                            case "":
                                                                var v45 float64 = 0.1
                                                                _ = v45
                                                            case "\n\t":
                                                                v34 += (v34 * v34)
                                                            default:
                                                                v5 = (((true && v5) || false) || true)
                                                            }
                                                            v2 = append(v2, fi6)
                                                            switch "\n\t" {
                                                            case "":
                                                                fmt.Println(v35, v34)
                                                            case "\n\t":
                                                                v46 := GenS7{F0: 657, F1: -998}
                                                                _ = v46
                                                            }
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    var v47 bool = false
                                                    _ = v47
                                                }
                                                v4 -= v4
                                                switch "foo bar" {
                                                case "":
                                                    if v5 {
                                                        var v48 int64 = 828
                                                        _ = v48
                                                    } else {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v5 = v5
                                                        case "":
                                                            v49 := GenS7{F0: 42, F1: 1}
                                                            _ = v49
                                                        case "\n\t":
                                                            var v50 float64 = (1.5 * (((1e10 / v1) / ((float64(1.5) + v1) + v1)) - 0.0))
                                                            v5 = v5
                                                            var v51 uint16 = 65535
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                case "foo bar":
                                                    switch fi6 {
                                                    case -2147483648:
                                                        if v5 {
                                                            var v52 uint = 302
                                                            _ = v52
                                                        }
                                                    }
                                                default:
                                                    var v53 int32 = 0
                                                    _ = v53
                                                }
                                                var v54 float32 = 0.0
                                                _ = v54
                                            }
                                            v55 := GenS1{F0: 125}
                                            fmt.Println(v1, fi6, v1)
                                            _ = v55
                                        }
                                        v4 = v4
                                        _ = v32
                                    } else {
                                        v1 = v1
                                        _ = len(v2)
                                        _ = len(v3)
                                        v5 = (!v5)
                                        var v56 string = ""
                                        _ = v56
                                    }
                                    if (!false) {
                                        if (!false) {
                                            if (v5 && v5) {
                                                if (v30 >= 9223372036854775807) {
                                                    v57 := GenS6{F0: fi18, F1: "foo bar", F2: (fi6 * fi18)}
                                                    v58 := GenS7{F0: 1, F1: 385}
                                                    var v59 int = (fi6 - -763)
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    var v60 bool = false
                                                    v61 := GenS8{F0: 39, F1: -1, F2: "\n\t"}
                                                    v30 = (fi18 / (((fi6 << ((655) & 63))) | 1))
                                                    v62 := GenS9{F0: 607, F1: 793}
                                                    var v63 string = "\n\t"
                                                    v30 = 0
                                                    v64 := GenS10{F0: -2147483648}
                                                    v30 -= 128
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                                if v5 {
                                                    var v65 int32 = 901
                                                    fmt.Println(fi18, fi18, fi18)
                                                    _ = v65
                                                }
                                            } else {
                                                var v66 uint32 = 2
                                                _ = v66
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v1 = 0.1
                                        }
                                        v67 := GenS4{F0: 42}
                                        _ = v67
                                    }
                                    _ = v30
                                }
                                _ = fi18
                            }
                        } else {
                            delete(v3, 4294967295)
                            if false {
                                fmt.Println(v1, v5, v5)
                                switch fi6 {
                                case -955:
                                    v68 := GenS1{F0: 1}
                                    _ = v68
                                }
                            }
                        }
                    }
                }
            default:
                for fi69 := 0; fi69 < 16; fi69++ {
                    var v70 int32 = -1
                    var v71 uint = 142
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
        default:
            fmt.Println(v5)
        }
        switch "" {
        case "foo bar":
            var v72 int16 = 0
            _ = v72
        default:
            if v5 {
                switch (fi6 / ((fi6) | 1)) {
                case -9223372036854775808:
                    v5 = v5
                case 42:
                    v73 := GenS11{F0: (((1.5 + 1e10) - ((0.0 + (1e10 / 1.0)) / (0.1 - (((3.14159 + (3.14159 - (3.14159 / float32(v1)))) / -2.25) + -2.25)))) < (-2.25 / float32(3.14159))), F1: 65535, F2: -80}
                    _ = v73
                }
            }
        }
        v74 := GenS11{F0: true, F1: 42, F2: -42}
        _ = fi6
        _ = v74
    }
    if (-2147483648 < 128) {
        var v75 float64 = (v1 + v1)
        var v76 float32 = -2.25
        switch "foo bar" {
        case "":
            if false {
                var v77 int16 = 122
                var v78 int32 = 699
                _ = v77
                _ = v78
            }
        }
        _ = v75
        _ = v76
    } else {
        var v79 uint16 = 2
        _ = v79
    }
    if (v4 < v4) {
        v4 += v4
    } else {
        v5 = (v5 || (65535 != 488))
        var v80 uint32 = 2147483648
        v1 -= v1
        _ = v80
    }
    if v5 {
        v81 := GenS11{F0: (42 < 100), F1: 2, F2: 127}
        if ((4294967295 >= 65535) || (!(226 <= 2))) {
            fmt.Println(v1, v4)
        }
        _ = v81
    }
    fmt.Println(v5, v4)
    if (v5 && (0 > 253)) {
        if false {
            var v82 uint32 = 888
            _ = v82
        } else {
            v5 = v5
            v4 += 213
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
