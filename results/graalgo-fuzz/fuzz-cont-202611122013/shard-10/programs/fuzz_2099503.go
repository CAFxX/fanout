package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 uint64; F1 string; F2 string }

func main() {
    var v1 int32 = -1
    var v2 float64 = -2.25
    var v3 bool = false
    var v4 uint = 4294967295
    v4 = v4
    switch "\n\t" {
    case "a":
        switch "hello" {
        case "a":
            var v5 int = 42
            _ = v5
        case "":
            if v3 {
                v2 += (0.0 + 3.14159)
                var v6 uint64 = 235
                var v7 float32 = 1.5
                _ = v6
                _ = v7
            }
        case "\n\t":
            if v3 {
                v8 := GenS1{F0: (v1 ^ (v1 | (v1 ^ v1))), F1: (v4 << ((v4) & 63))}
                _ = v8
            } else {
                v9 := []int{9223372036854775807, -747, 32767}
                switch -229 {
                case -399:
                    switch ("" + ("\n\t" + "hello")) {
                    case "":
                        var v10 uint16 = 879
                        _ = v10
                    case "hello":
                        if true {
                            v11 := GenS1{F0: 0, F1: v4}
                            _ = v11
                        }
                    case "a":
                        if v3 {
                            _ = len(v9)
                            switch "foo bar" {
                            case "a":
                                if len(v9) > 0 {
                                    v9[0] = 42
                                }
                            case "":
                                var v12 float32 = 3.14159
                                _ = v12
                            case "foo bar":
                                v13 := GenS1{F0: v1, F1: v4}
                                _ = v13
                            }
                        } else {
                            v14 := GenS1{F0: v1, F1: v4}
                            v1 = v1
                            _ = v14
                        }
                    default:
                        if (!false) {
                            switch -2147483648 {
                            case 32768:
                                v15 := GenS1{F0: (-(v1)), F1: 4294967295}
                                _ = v15
                            case -858:
                                v16 := GenS1{F0: v1, F1: (v4 & (^(v4)))}
                                _ = v16
                            }
                            if v3 {
                                switch "a" {
                                case "foo bar":
                                    if len(v9) > 0 {
                                        v9[0] = -605
                                    }
                                case "":
                                    v17 := GenS1{F0: (v1 ^ (v1 % ((-1) | 1))), F1: (v4 / ((18446744073709551615) | 1))}
                                    _ = v17
                                }
                                if v3 {
                                    _ = len(v9)
                                    var v18 float64 = float64(float32(0.1))
                                    _ = v18
                                }
                                v1 = v1
                            }
                            var v19 uint16 = 2
                            _ = v19
                        } else {
                            if v3 {
                                switch "hello" {
                                case "\n\t":
                                    var v20 uint = 2
                                    _ = v20
                                case "foo bar":
                                    var v21 int8 = 76
                                    _ = v21
                                case "":
                                    var v22 int64 = 0
                                    _ = v22
                                default:
                                    v23 := GenS1{F0: (v1 + -2147483648), F1: (v4 | v4)}
                                    _ = v23
                                }
                            } else {
                                v24 := GenS1{F0: 0, F1: (^(v4))}
                                switch 2147483647 {
                                case -1:
                                    switch 1 {
                                    case 416:
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    case -2147483648:
                                        v4 = v4
                                    }
                                case 127:
                                    if (0 == 0) {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        v2 += float64(1.5)
                                        var v25 int64 = -194
                                        _ = v25
                                    }
                                default:
                                    fmt.Println(v2, v2, v4)
                                }
                                if v3 {
                                    if v3 {
                                        var v26 int32 = -53
                                        _ = v26
                                    } else {
                                        var v27 uint16 = 316
                                        v3 = (((v1 >> ((18446744073709551615) & 63)) > v1) && (!(-1 > int32(v27))))
                                        v1 = (-(v1))
                                        _ = v27
                                    }
                                }
                                _ = v24
                            }
                            fmt.Println(v2, v4)
                        }
                    }
                case 0:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        var v29 uint = (^(v4))
                        _ = v29
                    case "hello":
                        if true {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v30 int = 9223372036854775807
                                _ = v30
                            case "hello":
                                if v3 {
                                    for fi31 := 0; fi31 < 19; fi31++ {
                                        v32 := GenS2{F0: 742, F1: (1 > 631)}
                                        v33 := GenS2{F0: 2, F1: v3}
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                    v1 = v1
                                } else {
                                    if (true && ((^(v4)) != v4)) {
                                        if len(v9) > 0 {
                                            v9[0] = -189
                                        }
                                        v1 = (v1 - (v1 - v1))
                                        v34 := GenS1{F0: v1, F1: 2}
                                        _ = v34
                                    }
                                }
                            case "a":
                                var v35 int8 = -1
                                _ = v35
                            }
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            fmt.Println(v1)
                        }
                    case "a":
                        if false {
                            if v3 {
                                if (83 == 923) {
                                    if v3 {
                                        for fi36 := 0; fi36 < 19; fi36++ {
                                            switch ("foo bar" + "") {
                                            case "hello":
                                                if ((^(v4)) != (v4 << ((0) & 63))) {
                                                    _ = len(v9)
                                                    v3 = (v4 <= v4)
                                                    if v3 {
                                                        v4 = (v4 / ((v4) | 1))
                                                        var v37 int16 = 42
                                                        for fi38 := 0; fi38 < 20; fi38++ {
                                                            v3 = v3
                                                            v2 = float64(1.5)
                                                            fmt.Println(v1)
                                                            switch ((("\n\t" + "a") + "hello") + ("foo bar" + ("\n\t" + ("\n\t" + "\n\t")))) {
                                                            case "hello":
                                                                v4 = (v4 % ((v4) | 1))
                                                            case "foo bar":
                                                                if v3 {
                                                                    var v39 int32 = (v1 << ((1) & 63))
                                                                    switch (("a" + ((("" + "\n\t") + "a") + "hello")) + ("" + "a")) {
                                                                    case "\n\t":
                                                                        v37 = (-(v37))
                                                                        if len(v9) >= 2 {
                                                                            v9 = v9[1:len(v9)]
                                                                        }
                                                                    default:
                                                                        v40 := GenS3{F0: -19, F1: (v37 / (((v37 * v37)) | 1)), F2: v4}
                                                                        _ = v40
                                                                    }
                                                                    var v41 int = 111
                                                                    _ = v39
                                                                    _ = v41
                                                                }
                                                            }
                                                            v42 := GenS4{F0: 42, F1: v1}
                                                            var v43 uint16 = 346
                                                            v44 := GenS5{F0: fi38}
                                                            if ((!false) || ((-2.25 != v2) && false)) {
                                                                switch (fi38 % (((fi38 - fi36)) | 1)) {
                                                                case 0:
                                                                    if false {
                                                                        fmt.Println(fi38)
                                                                        v3 = (0.1 > v2)
                                                                        v45 := GenS1{F0: (v1 ^ (^(v1))), F1: v4}
                                                                        _ = v45
                                                                    }
                                                                    switch 127 {
                                                                    case 570:
                                                                        v46 := GenS5{F0: (fi36 << ((v4) & 63))}
                                                                        _ = v46
                                                                    case 154:
                                                                        var v47 int8 = 42
                                                                        _ = v47
                                                                    case 32767:
                                                                        v1 = v1
                                                                    default:
                                                                        var v48 float64 = v2
                                                                        _ = v48
                                                                    }
                                                                case -1:
                                                                    var v49 uint = (v4 | (^(v4)))
                                                                    _ = v49
                                                                }
                                                                if (0 != -1) {
                                                                    var v50 uint32 = 139
                                                                    if ((float32(((1e10 + ((-2.25 * ((v2 * float64(1.5)) + (0.0 * -2.25))) + v2)) - float64(0.0))) * (float32((-2.25 - v2)) - ((1e10 - 0.0) * (((1e10 - (1e10 / -2.25)) / ((3.14159 - 0.1) * (float32(v2) / 1e10))) / -2.25)))) != (1e10 - 1e10)) {
                                                                        var v51 uint = (v4 * (v4 >> ((v4) & 63)))
                                                                        v51 = v4
                                                                        _ = v51
                                                                    } else {
                                                                        switch ("hello" + ((("a" + "a") + "\n\t") + "foo bar")) {
                                                                        case "foo bar":
                                                                            if len(v9) > 0 {
                                                                                v9[0] = fi36
                                                                            }
                                                                        case "hello":
                                                                            if true {
                                                                                switch (fi36 * 32768) {
                                                                                case 1:
                                                                                    v52 := GenS6{F0: 0, F1: v3}
                                                                                    _ = v52
                                                                                case -1:
                                                                                    switch ("hello" + "\n\t") {
                                                                                    case "a":
                                                                                        fmt.Println(v4, v37)
                                                                                    }
                                                                                case -399:
                                                                                    fmt.Println(fi38, fi36)
                                                                                default:
                                                                                    var v53 int = (fi36 & fi38)
                                                                                    _ = v53
                                                                                }
                                                                                if v3 {
                                                                                    var v54 string = ("" + "foo bar")
                                                                                    _ = len(v9)
                                                                                    _ = v54
                                                                                }
                                                                                var v55 int = int(v1)
                                                                                var v56 uint16 = v43
                                                                                _ = v55
                                                                                _ = v56
                                                                            }
                                                                            v57 := GenS7{F0: (!(35 >= 199))}
                                                                            var v58 uint = (v4 >> ((v4) & 63))
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        var v59 int32 = (v1 / (((v1 << ((v4) & 63))) | 1))
                                                                        switch (fi36 % ((fi38) | 1)) {
                                                                        case 9223372036854775807:
                                                                            var v60 int64 = -863
                                                                            _ = v60
                                                                        default:
                                                                            v61 := GenS4{F0: 1, F1: v59}
                                                                            _ = v61
                                                                        }
                                                                        switch -1 {
                                                                        case 42:
                                                                            fmt.Println(v43, v3, v3)
                                                                        case 2147483647:
                                                                            v1 = (^(v59))
                                                                            var v62 int32 = -658
                                                                            v4 -= (v4 / ((v4) | 1))
                                                                            _ = v62
                                                                        default:
                                                                            v59 += 1
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                    v50 = (^(v50))
                                                                    _ = v50
                                                                } else {
                                                                    var v63 float64 = (v2 * -2.25)
                                                                    _ = v63
                                                                }
                                                            } else {
                                                                if false {
                                                                    v64 := GenS8{F0: ""}
                                                                    v1 = (v1 % (((v1 / (((v1 & (v1 + -1))) | 1))) | 1))
                                                                    fmt.Println(fi38, fi38)
                                                                    _ = v64
                                                                } else {
                                                                    if len(v9) > 0 {
                                                                        v9[0] = int(v1)
                                                                    }
                                                                    v65 := GenS2{F0: (v43 | 43), F1: (18446744073709551615 <= 0)}
                                                                    switch 4294967295 {
                                                                    case 128:
                                                                        _ = len(v9)
                                                                    case 9223372036854775807:
                                                                        var v66 float64 = v2
                                                                        _ = v66
                                                                    default:
                                                                        v43 += v43
                                                                    }
                                                                    _ = v65
                                                                }
                                                                var v67 int32 = v1
                                                                v68 := GenS8{F0: "hello"}
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            v1 += (-(v1))
                                                            _ = len(v9)
                                                            _ = fi38
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        _ = v37
                                                    }
                                                } else {
                                                    if (!v3) {
                                                        var v69 float32 = ((1.5 * -2.25) - (0.0 * 1e10))
                                                        _ = v69
                                                    } else {
                                                        v9 = append(v9, -124)
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        var v70 uint32 = 0
                                                        _ = v70
                                                    }
                                                    fmt.Println(fi36, v2)
                                                }
                                            }
                                            var v71 float64 = v2
                                            _ = fi36
                                            _ = v71
                                        }
                                        fmt.Println(v3, v4, v2)
                                        for fi72 := 0; fi72 < 3; fi72++ {
                                            var v73 int = 25
                                            _ = fi72
                                            _ = v73
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "":
                                        v74 := GenS2{F0: 0, F1: (783 > -844)}
                                        _ = v74
                                    default:
                                        if v3 {
                                            if false {
                                                v1 = v1
                                            } else {
                                                v3 = (121 > 42)
                                            }
                                            v1 = (v1 | v1)
                                            _ = len(v9)
                                        } else {
                                            var v75 int8 = 127
                                            _ = v75
                                        }
                                    }
                                }
                            } else {
                                switch ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + ("foo bar" + "")))) {
                                case "\n\t":
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                switch 279 {
                                case 127:
                                    var v76 int32 = (v1 + 45)
                                    _ = v76
                                case 2147483647:
                                    _ = len(v9)
                                }
                            }
                        } else {
                            v1 = 0
                        }
                    }
                }
                switch "foo bar" {
                case "hello":
                    for fi77 := 0; fi77 < 8; fi77++ {
                        v78 := GenS1{F0: (v1 ^ 1), F1: 65535}
                        for fi79 := 0; fi79 < 16; fi79++ {
                            v80 := GenS9{F0: "hello", F1: 42}
                            if (65535 >= 545) {
                                v9 = append(v9, 533)
                            }
                            var v81 int32 = (v1 | v1)
                            _ = fi79
                            _ = v80
                            _ = v81
                        }
                        v82 := GenS10{F0: 40}
                        _ = fi77
                        _ = v78
                        _ = v82
                    }
                default:
                    var v83 uint32 = 1
                    _ = v83
                }
                _ = v9
            }
        }
    case "foo bar":
        var v84 int32 = (v1 << ((v4) & 63))
        _ = v84
    case "":
        v85 := GenS11{F0: 128, F1: -945}
        _ = v85
    }
    fmt.Println(v2, v3)
    fmt.Println(v1, v4, v4)
    switch 4294967295 {
    case 127:
        if ((!true) && v3) {
            var v86 int16 = 1000
            _ = v86
        } else {
            if v3 {
                v87 := []int{42, 751, -882}
                _ = v87
            } else {
                v88 := GenS12{F0: -1, F1: (-(v1))}
                _ = v88
            }
            var v89 int64 = -9223372036854775808
            _ = v89
        }
    default:
        var v90 int8 = -128
        _ = v90
    }
    var v91 uint64 = 429
    fmt.Println(v3, v4)
    if v3 {
        v2 = (1.5 + (float64(float32(1e10)) - -2.25))
    } else {
        v92 := GenS13{F0: v91, F1: (("" + ("a" + ("" + "a"))) + "hello"), F2: ""}
        for fi93 := 0; fi93 < 2; fi93++ {
            v94 := GenS4{F0: -108, F1: v1}
            v1 -= v1
            var v95 uint = v4
            _ = fi93
            _ = v94
            _ = v95
        }
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v91
}
