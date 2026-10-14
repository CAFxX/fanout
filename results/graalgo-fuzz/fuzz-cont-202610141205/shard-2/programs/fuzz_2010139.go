package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int64; F2 int }
type GenS5 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 bool = (32767 < 255)
    var v4 int16 = -799
    v5 := []string{}
    v1 = v1
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        if len(v5) > 0 {
            v5[0] = ""
        }
        _ = fi7
    }
    if v3 {
        fmt.Println(v6, v1, v1)
        if (1 >= uint32(v6)) {
            v8 := GenS1{F0: 255, F1: 32768, F2: "a"}
            var v9 float64 = 1e10
            for fi10 := 0; fi10 < 20; fi10++ {
                v4 += 37
                v5 = append(v5, (("hello" + "foo bar") + "\n\t"))
                _ = fi10
            }
            _ = v8
            _ = v9
        }
        v11 := GenS1{F0: 255, F1: 42, F2: "\n\t"}
        _ = v11
    } else {
        v12 := GenS1{F0: 65535, F1: 32767, F2: "foo bar"}
        if (-1 != -272) {
            for fi13 := 0; fi13 < 3; fi13++ {
                switch (-(fi13)) {
                case -14:
                    var v14 int32 = v1
                    _ = v14
                }
                var v15 int64 = -867
                v4 += v4
                _ = fi13
                _ = v15
            }
            if v3 {
                var v16 float32 = 3.14159
                fmt.Println(v3)
                _ = v16
            } else {
                v17 := GenS1{F0: 32767, F1: 42, F2: ""}
                for fi18 := 0; fi18 < 9; fi18++ {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        switch 29 {
                        case 9223372036854775807:
                            v6 = (v6 & v6)
                        }
                        v6 = v6
                        var v20 float32 = 0.1
                        if (((65535 != 2147483648) || (1 >= 9223372036854775807)) && v3) {
                            v21 := GenS1{F0: 65535, F1: (fi18 - 127), F2: "hello"}
                            v22 := GenS1{F0: 2, F1: (fi18 >> ((65535) & 63)), F2: "hello"}
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int = fi19
                            var v24 uint32 = 1
                            fmt.Println(fi18, v20)
                            var v25 int = fi18
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v26 := GenS1{F0: 0, F1: fi19, F2: "hello"}
                        _ = fi19
                        _ = v20
                        _ = v26
                    }
                    switch 248 {
                    case 1:
                        var v27 bool = (32768 > fi18)
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 6; fi28++ {
                            if v3 {
                                switch "foo bar" {
                                case "a":
                                    if ((!v3) && v3) {
                                        var v29 string = ("a" + "")
                                        var v30 int64 = 187
                                        _ = v29
                                        _ = v30
                                    }
                                case "hello":
                                    v31 := v2[565]
                                    _ = v31
                                }
                                v3 = v3
                            } else {
                                v2[684] = v3
                                var v32 uint64 = 44
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        v1 = v1
                                        v33 := GenS1{F0: 42, F1: 32768, F2: (("a" + ("" + "\n\t")) + "\n\t")}
                                        _ = v33
                                    case "":
                                        if (42 >= -240) {
                                            v4 += v4
                                        }
                                    }
                                    if ((1.5 + ((((1e10 - 0.0) * 0.1) + ((1e10 * 1.5) + 0.0)) - 0.1)) != 1e10) {
                                        var v34 string = "\n\t"
                                        v35 := GenS1{F0: 629, F1: 9223372036854775807, F2: ("\n\t" + v34)}
                                        _ = v34
                                        _ = v35
                                    }
                                    v36 := GenS1{F0: 804, F1: (fi28 ^ fi28), F2: "hello"}
                                    if v3 {
                                        switch ("" + "hello") {
                                        case "":
                                            fmt.Println(v1, v32)
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = "foo bar"
                                            }
                                        }
                                        v6 = 42
                                        fmt.Println(v6)
                                    } else {
                                        var v37 bool = (v3 && v3)
                                        v3 = v37
                                        if (42 != 1) {
                                            if false {
                                                v3 = ((v1 >> ((18446744073709551615) & 63)) >= v1)
                                                fmt.Println(v1, fi18, v1)
                                            }
                                            v38 := GenS2{F0: "\n\t", F1: v4, F2: 823}
                                            v37 = ("\n\t" != "\n\t")
                                            _ = v38
                                        } else {
                                            v6 -= (^(v6))
                                            v1 = -2147483648
                                        }
                                        _ = v37
                                    }
                                    v39 := GenS2{F0: "a", F1: v4, F2: 233}
                                    _ = v36
                                    _ = v39
                                }
                                _ = v32
                            }
                            v3 = (v3 || (601 == 77))
                            _ = fi28
                        }
                    }
                    v40 := GenS2{F0: ((("hello" + "") + "a") + "hello"), F1: (^(v4)), F2: 570}
                    _ = fi18
                    _ = v40
                }
                _ = v17
            }
        } else {
            fmt.Println(v3, v4, v1)
            var v41 uint32 = 2147483648
            _ = v41
        }
        _ = v12
    }
    v42 := GenS1{F0: 65535, F1: 2147483647, F2: "\n\t"}
    for fi43 := 0; fi43 < 10; fi43++ {
        for fi44 := 0; fi44 < 2; fi44++ {
            if (v3 || false) {
                v3 = ("foo bar" < "hello")
            }
            _ = fi44
        }
        _ = fi43
    }
    if true {
        fmt.Println(v1)
        for fi45 := 0; fi45 < 13; fi45++ {
            if true {
                switch ("a" + "foo bar") {
                case "\n\t":
                    v46 := GenS3{F0: "hello"}
                    _ = v46
                default:
                    if true {
                        v47 := GenS1{F0: 2, F1: (fi45 * 32768), F2: ("hello" + "a")}
                        v6 += (v6 % (((v6 >> ((42) & 63))) | 1))
                        _ = v47
                    }
                }
                if (1 >= 1) {
                    switch ("\n\t" + "") {
                    case "":
                        switch "\n\t" {
                        case "a":
                            var v48 float32 = (((0.0 * 0.0) / (((0.0 - (1e10 * float32(((-2.25 / ((0.1 / 1e10) * ((0.1 * 3.14159) + ((1e10 / -2.25) + (0.0 / ((-2.25 + 3.14159) + 1.5)))))) + 1.5)))) / 0.1) - 3.14159)) + 0.0)
                            _ = v48
                        }
                    case "\n\t":
                        for fi49 := 0; fi49 < 17; fi49++ {
                            v3 = (("hello" + "\n\t") <= (("foo bar" + ((("hello" + "hello") + "hello") + "\n\t")) + "a"))
                            if v3 {
                                v4 = v4
                                if v3 {
                                    var v50 float64 = (1.5 / float64((1.5 / 3.14159)))
                                    _ = v50
                                } else {
                                    v51 := GenS3{F0: ("a" + ("" + "\n\t"))}
                                    _ = v51
                                }
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                switch ("a" + "a") {
                                case "\n\t":
                                    v4 = (v4 / ((32767) | 1))
                                case "":
                                    switch "foo bar" {
                                    case "hello":
                                        if v3 {
                                            if true {
                                                v3 = (0 == 0)
                                            } else {
                                                var v52 uint16 = 255
                                                var v53 bool = (!(4294967295 > 42))
                                                fmt.Println(v6, v3)
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    case "\n\t":
                                        if v3 {
                                            switch -1 {
                                            case 1:
                                                switch (fi49 >> ((863) & 63)) {
                                                case 445:
                                                    fmt.Println(v1, v3, v3)
                                                case 713:
                                                    switch (fi49 ^ (fi49 * (fi49 >> ((18446744073709551615) & 63)))) {
                                                    case -1:
                                                        v54 := v2[(-(fi45))]
                                                        v55 := GenS3{F0: "foo bar"}
                                                        var v56 float64 = -2.25
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    case 32767:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    default:
                                                        v57 := GenS2{F0: "a", F1: (v4 % ((42) | 1)), F2: 2}
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v58 float64 = (((float64(1.5) + 0.0) + 3.14159) + (0.0 * (0.1 / float64(1e10))))
                                                            _ = v58
                                                        case "":
                                                            v59 := GenS2{F0: "\n\t", F1: v4, F2: 661}
                                                            _ = v59
                                                        case "a":
                                                            v60 := v2[-2147483648]
                                                            var v61 int64 = 9223372036854775807
                                                            v62 := GenS4{F0: v1, F1: (^(v61)), F2: 127}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        _ = v57
                                                    }
                                                }
                                            case -760:
                                                v63 := GenS4{F0: v1, F1: 0, F2: 210}
                                                _ = v63
                                            case 0:
                                                fmt.Println(v6)
                                            }
                                        }
                                    }
                                }
                                v5 = append(v5, "foo bar")
                            }
                            _ = fi49
                        }
                    default:
                        v4 = 42
                    }
                }
            }
            v3 = v3
            _ = fi45
        }
        if (67 != 65535) {
            v1 += v1
            fmt.Println(v1, v6, v6)
        } else {
            v3 = v3
            if v3 {
                switch "a" {
                case "":
                    v2[-387] = v3
                }
                if true {
                    fmt.Println(v1, v3)
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                }
                v6 += v6
            } else {
                switch "a" {
                case "hello":
                    if ((1e10 - ((float32((1.5 / 1.0)) - float32(1.5)) - 1.5)) == 1.5) {
                        for fi64 := 0; fi64 < 7; fi64++ {
                            for fi65 := 0; fi65 < 6; fi65++ {
                                switch (fi65 + 128) {
                                case 42:
                                    var v66 uint8 = 127
                                    _ = v66
                                case 69:
                                    v3 = v3
                                }
                                fmt.Println(v6)
                                var v67 uint8 = (v6 << ((937) & 63))
                                _ = fi65
                                _ = v67
                            }
                            v4 += v4
                            var v68 int = fi64
                            _ = fi64
                            _ = v68
                        }
                    }
                }
                var v69 int32 = v1
                var v70 int64 = 42
                _ = v69
                _ = v70
            }
            fmt.Println(v6, v1)
        }
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        var v72 uint16 = 32767
        var v73 int16 = 32767
        switch fi71 {
        case 32767:
            var v74 float32 = -2.25
            _ = v74
        case 42:
            switch (fi71 | fi71) {
            case -860:
                v75 := GenS1{F0: (v72 & v72), F1: (-(fi71)), F2: "\n\t"}
                _ = v75
            }
        }
        _ = fi71
        _ = v72
        _ = v73
    }
    switch 127 {
    case 1:
        for fi76 := 0; fi76 < 8; fi76++ {
            if true {
                var v77 int32 = (^(v1))
                if (!false) {
                    v78 := GenS5{F0: 517, F1: -9223372036854775808, F2: uint32(v6)}
                    fmt.Println(v77)
                    var v79 int32 = v1
                    _ = v78
                    _ = v79
                }
                _ = v77
            }
            v80 := v2[(fi76 & 0)]
            if (1 <= 10) {
                var v81 int64 = 293
                v3 = (v80 && true)
                fmt.Println(v6, v6, v80)
                _ = v81
            } else {
                var v82 int64 = 42
                _ = v82
            }
            _ = fi76
            _ = v80
        }
    default:
        var v83 int16 = 1
        _ = v83
    }
    if v3 {
        v1 -= (v1 & (v1 - (v1 & v1)))
    }
    v4 = -251
    switch 72 {
    case -2147483648:
        for fi84 := 0; fi84 < 18; fi84++ {
            v4 = (v4 * v4)
            _ = fi84
        }
    case 42:
        v4 = (v4 + 32767)
    default:
        if len(v5) > 0 {
            v5[0] = ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + "")
        }
    }
    var v85 float32 = float32((0.0 + 3.14159))
    for fi86 := 0; fi86 < 19; fi86++ {
        for fi87 := 0; fi87 < 4; fi87++ {
            v88 := GenS4{F0: (v1 ^ v1), F1: 9223372036854775807, F2: fi87}
            if ((true && (false || (255 < 942))) && v3) {
                v89 := GenS4{F0: 505, F1: -9223372036854775808, F2: 4294967295}
                var v90 int = (fi87 + (fi86 - (-(fi86))))
                _ = v89
                _ = v90
            }
            if v3 {
                var v91 uint = 65535
                v6 = 2
                v1 = v1
                fmt.Println(v91)
                _ = v91
            } else {
                switch fi86 {
                case -2147483648:
                    _ = len(v5)
                case 42:
                    v3 = (fi86 < (^(fi86)))
                case -1:
                    fmt.Println(v3)
                default:
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                v1 -= v1
                _ = len(v5)
            }
            if (-2147483648 < fi87) {
                v6 -= 87
            } else {
                if false {
                    v4 += (v4 ^ (v4 / ((706) | 1)))
                    if (v4 < (v4 + v4)) {
                        v1 -= (v1 & 2147483647)
                    }
                } else {
                    fmt.Println(v4)
                }
            }
            _ = fi87
            _ = v88
        }
        _ = fi86
    }
    if v3 {
        switch 42 {
        case -445:
            for fi92 := 0; fi92 < 8; fi92++ {
                if v3 {
                    for fi93 := 0; fi93 < 18; fi93++ {
                        switch ("hello" + "") {
                        case "a":
                            var v94 uint32 = 2147483647
                            _ = v94
                        case "foo bar":
                            fmt.Println(v6)
                        default:
                            var v95 uint32 = 2147483647
                            _ = v95
                        }
                        v1 += v1
                        _ = fi93
                    }
                    var v96 float64 = -2.25
                    if ((!v3) || true) {
                        var v97 uint = 1
                        v98 := v2[fi92]
                        if false {
                            for fi99 := 0; fi99 < 20; fi99++ {
                                v1 = v1
                                _ = fi99
                            }
                        } else {
                            fmt.Println(v96)
                            fmt.Println(v3, v4)
                        }
                        _ = v97
                        _ = v98
                    }
                    _ = v96
                }
                _ = fi92
            }
        case 2147483647:
            var v100 int = -158
            _ = v100
        }
    }
    v1 -= -751
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v85
}
