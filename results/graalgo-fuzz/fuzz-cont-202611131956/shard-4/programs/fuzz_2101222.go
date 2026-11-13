package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 string; F1 bool; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int8; F1 bool; F2 uint8 }
type GenS9 struct { F0 uint; F1 int16; F2 int32 }

func main() {
    var v1 bool = true
    var v2 string = ("foo bar" + "foo bar")
    var v3 int32 = 598
    v4 := []int{-1, -2147483648, 2147483647}
    fmt.Println(v2)
    v2 = (v2 + v2)
    if true {
        switch v2 {
        case "\n\t":
            var v5 int32 = (^(v3))
            _ = v5
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            fmt.Println(v3, v1, v2)
        default:
            var v6 int64 = 1
            _ = v6
        }
        v7 := GenS1{F0: true}
        _ = v7
    } else {
        v2 = ("foo bar" + ("\n\t" + ("" + ("foo bar" + ("\n\t" + "hello")))))
        if false {
            switch (v2 + (v2 + v2)) {
            case "hello":
                fmt.Println(v1, v1, v2)
            case "":
                for fi8 := 0; fi8 < 16; fi8++ {
                    v1 = v1
                    switch 4294967295 {
                    case 42:
                        fmt.Println(v1)
                    case 127:
                        v9 := GenS2{F0: (v3 - v3)}
                        _ = v9
                    }
                    var v10 int32 = (^(v3))
                    _ = fi8
                    _ = v10
                }
            default:
                if (v2 != v2) {
                    if ((21 < 2) || true) {
                        if v1 {
                            v11 := GenS2{F0: v3}
                            if (1e10 >= 0.0) {
                                switch 0 {
                                case -83:
                                    for fi12 := 0; fi12 < 17; fi12++ {
                                        switch (fi12 * fi12) {
                                        case -9223372036854775808:
                                            if v1 {
                                                var v13 int16 = 1
                                                fmt.Println(fi12, fi12)
                                                _ = v13
                                            } else {
                                                v3 = (v3 >> ((18446744073709551615) & 63))
                                            }
                                        case -506:
                                            if (!(!false)) {
                                                v14 := GenS3{F0: 100, F1: "hello"}
                                                _ = v14
                                            }
                                        case 1:
                                            if len(v4) > 0 {
                                                v4[0] = fi12
                                            }
                                        default:
                                            fmt.Println(v1, v3, v2)
                                        }
                                        _ = fi12
                                    }
                                }
                                _ = len(v4)
                                var v15 bool = (!(42 > -9223372036854775808))
                                _ = v15
                            } else {
                                var v16 int = 32768
                                _ = v16
                            }
                            v17 := GenS1{F0: v1}
                            _ = v11
                            _ = v17
                        } else {
                            v18 := GenS2{F0: (v3 << ((1) & 63))}
                            _ = v18
                        }
                    }
                    v2 = ("hello" + "hello")
                } else {
                    if (1 >= -32768) {
                        if v1 {
                            v2 = (((v2 + (v2 + "hello")) + v2) + v2)
                        } else {
                            switch v2 {
                            case "":
                                switch v2 {
                                case "hello":
                                    if false {
                                        v4 = append(v4, 127)
                                        fmt.Println(v3)
                                        v19 := GenS1{F0: (!false)}
                                        _ = v19
                                    }
                                case "\n\t":
                                    var v20 float64 = ((1e10 / 1.0) - ((0.0 - 1.5) - 1e10))
                                    _ = v20
                                }
                            case "hello":
                                var v21 int64 = 663
                                _ = v21
                            case "foo bar":
                                switch "a" {
                                case "":
                                    if v1 {
                                        if (2 == 42) {
                                            v1 = v1
                                            v22 := GenS3{F0: -101, F1: v2}
                                            var v23 uint64 = 502
                                            _ = v22
                                            _ = v23
                                        }
                                    } else {
                                        var v24 int16 = -1
                                        if true {
                                            v2 = ("a" + "hello")
                                            switch (v2 + v2) {
                                            case "a":
                                                if v1 {
                                                    if (18446744073709551615 >= 42) {
                                                        v24 -= -1
                                                        fmt.Println(v3, v3, v3)
                                                    }
                                                    switch -870 {
                                                    case -1:
                                                        if (32767 == -748) {
                                                            if v1 {
                                                                v25 := GenS4{F0: v2, F1: (v2 >= v2), F2: (v24 % (((v24 | v24)) | 1))}
                                                                _ = v25
                                                            }
                                                            for fi26 := 0; fi26 < 16; fi26++ {
                                                                switch fi26 {
                                                                case 128:
                                                                    var v27 uint16 = 65535
                                                                    _ = v27
                                                                }
                                                                var v28 uint32 = 14
                                                                _ = fi26
                                                                _ = v28
                                                            }
                                                        } else {
                                                            var v29 string = "\n\t"
                                                            _ = v29
                                                        }
                                                    case 0:
                                                        if false {
                                                            var v30 float32 = 0.0
                                                            v4 = append(v4, 554)
                                                            if true {
                                                                v31 := GenS2{F0: 0}
                                                                v2 = v2
                                                                var v32 uint64 = 2
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            _ = v30
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = -1
                                                            }
                                                            v24 = 0
                                                            fmt.Println(v1, v24, v24)
                                                        }
                                                    default:
                                                        for fi33 := 0; fi33 < 13; fi33++ {
                                                            fmt.Println(fi33, fi33, v24)
                                                            var v34 uint8 = 127
                                                            _ = fi33
                                                            _ = v34
                                                        }
                                                    }
                                                } else {
                                                    v24 = (v24 ^ 1)
                                                }
                                            case "":
                                                if v1 {
                                                    v35 := GenS2{F0: v3}
                                                    _ = v35
                                                } else {
                                                    for fi36 := 0; fi36 < 19; fi36++ {
                                                        v24 = v24
                                                        if ((-35 == -1) || false) {
                                                            var v37 uint8 = 166
                                                            _ = v37
                                                        }
                                                        if v1 {
                                                            switch (fi36 + fi36) {
                                                            case 32767:
                                                                fmt.Println(v1)
                                                            case 42:
                                                                fmt.Println(v24)
                                                            default:
                                                                var v38 uint16 = 255
                                                                _ = v38
                                                            }
                                                        } else {
                                                            for fi39 := 0; fi39 < 10; fi39++ {
                                                                v24 += (v24 + 42)
                                                                var v40 int16 = (v24 | (^(v24)))
                                                                _ = fi39
                                                                _ = v40
                                                            }
                                                        }
                                                        _ = fi36
                                                    }
                                                }
                                            case "foo bar":
                                                v24 = v24
                                            default:
                                                v2 = (v2 + "a")
                                            }
                                        } else {
                                            fmt.Println(v24, v24)
                                        }
                                        switch -1 {
                                        case -541:
                                            v1 = v1
                                        case 32768:
                                            v3 = (v3 | (v3 >> ((4294967295) & 63)))
                                        }
                                        _ = v24
                                    }
                                case "hello":
                                    for fi41 := 0; fi41 < 12; fi41++ {
                                        var v42 uint64 = 37
                                        v43 := GenS1{F0: v1}
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                case "a":
                                    var v44 int32 = (v3 << ((1000) & 63))
                                    _ = v44
                                default:
                                    var v45 int16 = 0
                                    _ = v45
                                }
                            default:
                                for fi46 := 0; fi46 < 10; fi46++ {
                                    for fi47 := 0; fi47 < 15; fi47++ {
                                        var v48 float64 = -2.25
                                        var v49 float32 = -2.25
                                        fmt.Println(v48, fi47)
                                        v48 += -2.25
                                        var v50 uint = 4294967295
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                    var v51 float32 = 0.1
                                    _ = fi46
                                    _ = v51
                                }
                            }
                            for fi52 := 0; fi52 < 9; fi52++ {
                                v1 = (!false)
                                _ = fi52
                            }
                            v53 := GenS5{F0: 42}
                            _ = v53
                        }
                        var v54 int16 = 42
                        _ = v54
                    }
                }
            }
            switch 2147483647 {
            case 2147483647:
                switch ((v2 + ("\n\t" + v2)) + v2) {
                case "\n\t":
                    fmt.Println(v3, v1, v3)
                case "":
                    v3 = v3
                }
            case -39:
                v2 = "\n\t"
            case 42:
                if ((!(1 < 42)) || v1) {
                    var v55 uint64 = 0
                    _ = len(v4)
                    v3 = (v3 >> ((0) & 63))
                    _ = v55
                }
            }
            if v1 {
                v2 = v2
                v3 = v3
            } else {
                for fi56 := 0; fi56 < 19; fi56++ {
                    v1 = true
                    switch ("hello" + v2) {
                    case "":
                        fmt.Println(v2)
                    case "\n\t":
                        var v57 int32 = -1
                        _ = v57
                    }
                    _ = len(v4)
                    _ = fi56
                }
                var v58 int16 = -699
                _ = v58
            }
        }
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    switch 189 {
    case 42:
        v2 = (v2 + (v2 + v2))
    case -816:
        for fi59 := 0; fi59 < 9; fi59++ {
            if len(v4) > 0 {
                v4[0] = -2147483648
            }
            if false {
                v60 := GenS4{F0: "foo bar", F1: (!(32767 > 32767)), F2: 1}
                _ = v60
            } else {
                v2 = (("\n\t" + (v2 + ("foo bar" + v2))) + "a")
                for fi61 := 0; fi61 < 9; fi61++ {
                    switch (fi59 & fi61) {
                    case 42:
                        v62 := GenS5{F0: (^(v3))}
                        _ = v62
                    case 32767:
                        v63 := GenS6{F0: 255, F1: -520}
                        _ = v63
                    case -481:
                        var v64 int8 = 42
                        _ = v64
                    }
                    _ = fi61
                }
            }
            var v65 uint8 = 2
            _ = fi59
            _ = v65
        }
    default:
        if (42 != (v3 + -808)) {
            v1 = (0 < 100)
            if v1 {
                fmt.Println(v2, v1)
                for fi66 := 0; fi66 < 10; fi66++ {
                    switch v2 {
                    case "\n\t":
                        var v67 uint16 = 453
                        _ = v67
                    case "foo bar":
                        v2 = ((("foo bar" + "\n\t") + v2) + "")
                    default:
                        if v1 {
                            switch v2 {
                            case "foo bar":
                                v1 = v1
                            }
                            switch v2 {
                            case "":
                                for fi68 := 0; fi68 < 2; fi68++ {
                                    v1 = (!v1)
                                    var v69 int64 = -9223372036854775808
                                    v70 := GenS3{F0: -78, F1: v2}
                                    _ = fi68
                                    _ = v69
                                    _ = v70
                                }
                            case "\n\t":
                                for fi71 := 0; fi71 < 18; fi71++ {
                                    var v72 int32 = (v3 % ((v3) | 1))
                                    if false {
                                        v73 := GenS5{F0: v3}
                                        _ = len(v4)
                                        _ = v73
                                    } else {
                                        switch v2 {
                                        case "hello":
                                            v2 = v2
                                        case "":
                                            _ = len(v4)
                                        case "a":
                                            switch -467 {
                                            case -2147483648:
                                                var v74 uint16 = 42
                                                _ = v74
                                            case 32768:
                                                if ((^(v72)) != 1) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch ("\n\t" + ("\n\t" + (v2 + (v2 + "\n\t")))) {
                                                    case "":
                                                        v3 = (v3 + (v72 >> ((65535) & 63)))
                                                    }
                                                    v2 = (v2 + ((v2 + (v2 + "foo bar")) + v2))
                                                }
                                            case 127:
                                                var v75 int16 = 51
                                                _ = v75
                                            default:
                                                v76 := GenS7{F0: 1, F1: 1}
                                                _ = v76
                                            }
                                        }
                                        v2 = ("\n\t" + "\n\t")
                                        var v77 uint8 = 0
                                        _ = v77
                                    }
                                    var v78 uint32 = 967
                                    fmt.Println(fi71)
                                    _ = fi71
                                    _ = v72
                                    _ = v78
                                }
                            }
                        }
                    }
                    var v79 float64 = 0.1
                    _ = fi66
                    _ = v79
                }
            } else {
                v4 = append(v4, 156)
            }
        }
    }
    for fi80 := 0; fi80 < 15; fi80++ {
        v81 := GenS8{F0: -1, F1: (!(931 != 2)), F2: 2}
        _ = fi80
        _ = v81
    }
    v3 = v3
    _ = len(v4)
    if (9223372036854775807 > -9223372036854775808) {
        for fi82 := 0; fi82 < 11; fi82++ {
            var v83 int32 = (v3 | (v3 | v3))
            v1 = true
            if (!(127 > 130)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v84 int16 = -137
                for fi85 := 0; fi85 < 9; fi85++ {
                    var v86 string = v2
                    fmt.Println(v1)
                    _ = fi85
                    _ = v86
                }
                _ = v84
            } else {
                switch (fi82 - fi82) {
                case 9223372036854775807:
                    var v87 string = ""
                    _ = v87
                case -2147483648:
                    var v88 int32 = (v83 >> ((146) & 63))
                    _ = v88
                case -368:
                    v89 := GenS7{F0: 65535, F1: 1}
                    _ = v89
                }
                for fi90 := 0; fi90 < 9; fi90++ {
                    v83 -= (v3 + v83)
                    v83 = (v3 & (v83 * int32(fi82)))
                    switch (fi90 >> ((4294967295) & 63)) {
                    case 474:
                        var v91 string = ""
                        _ = v91
                    }
                    _ = fi90
                }
                if (42 == 1) {
                    for fi92 := 0; fi92 < 18; fi92++ {
                        v4 = append(v4, 127)
                        var v93 float64 = 1.5
                        _ = fi92
                        _ = v93
                    }
                    switch fi82 {
                    case 0:
                        v83 -= (-(v83))
                    case -9223372036854775808:
                        v3 -= (-(v3))
                    }
                } else {
                    v3 += (v3 * (^(v3)))
                }
            }
            _ = fi82
            _ = v83
        }
        v2 = "\n\t"
        if v1 {
            v1 = (v1 || v1)
            switch "\n\t" {
            case "":
                v2 = (v2 + v2)
            case "a":
                v94 := GenS8{F0: 0, F1: false, F2: 1}
                _ = v94
            case "hello":
                if ((98 != 42) || v1) {
                    for fi95 := 0; fi95 < 2; fi95++ {
                        var v96 int = (fi95 + 127)
                        _ = fi95
                        _ = v96
                    }
                    switch (((v2 + "\n\t") + "") + v2) {
                    case "":
                        v2 = v2
                    case "\n\t":
                        v1 = ("foo bar" == v2)
                    case "hello":
                        var v97 int64 = 42
                        _ = v97
                    }
                } else {
                    var v98 int16 = 212
                    v99 := GenS9{F0: 651, F1: v98, F2: v3}
                    var v100 float64 = (3.14159 / ((3.14159 / 1.0) + 3.14159))
                    _ = v100
                    _ = v98
                    _ = v99
                }
            }
        }
    } else {
        var v101 int32 = v3
        _ = v101
    }
    v4 = append(v4, 428)
    v3 -= v3
    var v102 bool = (2 < 2)
    v102 = (!true)
    if false {
        var v103 string = v2
        _ = v103
    }
    var v104 bool = v1
    var v105 uint = 196
    v105 -= v105
    v104 = true
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v102
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
}
