package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int16; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint16; F1 uint64 }
func genFn0() int64 {
    switch "a" {
    case "a":
        var v7 float32 = 1.5
        _ = v7
    case "foo bar":
        if (!(true && true)) {
            v8 := GenS1{F0: 255, F1: -16, F2: 393}
            var v9 int8 = -23
            _ = v8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int = (fi11 / ((fi11) | 1))
                v13 := GenS1{F0: 1, F1: 334, F2: 42}
                _ = fi11
                _ = v12
                _ = v13
            }
            switch "hello" {
            case "\n\t":
                var v14 string = ("a" + (("hello" + "hello") + ""))
                _ = v14
            case "":
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 uint8 = 161
                    _ = fi15
                    _ = v16
                }
            }
            var v17 uint = 65535
            _ = fi10
            _ = v17
        }
    }
    switch 117 {
    case 734:
        if (((-1 > 0) || false) || ((2147483647 == 0) || true)) {
            var v18 int = 32767
            _ = v18
        } else {
            switch 0 {
            case 9223372036854775807:
                v19 := []int{1, -9223372036854775808, 798}
                _ = v19
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 uint16 = 32767
                _ = fi20
                _ = v21
            }
            v22 := []int{}
            _ = v22
        }
    case 42:
        var v23 uint16 = 473
        _ = v23
    default:
        var v24 int32 = 42
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(fi25)
        switch "a" {
        case "foo bar":
            var v26 uint8 = 42
            _ = v26
        default:
            v27 := []int{fi25, fi25, (fi25 | fi25)}
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            if (fi28 != fi28) {
                var v29 string = "foo bar"
                v29 = v29
                _ = v29
            } else {
                var v30 int16 = 0
                switch "hello" {
                case "":
                    fmt.Println(v30, fi28, fi28)
                case "\n\t":
                    v31 := GenS1{F0: 6, F1: (v30 % ((-527) | 1)), F2: 0}
                    _ = v31
                case "foo bar":
                    v32 := []int{fi25, (-(fi28))}
                    var v33 uint32 = 2
                    _ = v32
                    _ = v33
                default:
                    var v34 float32 = -2.25
                    _ = v34
                }
                if (((v30 >> ((0) & 63)) >= (^(v30))) || false) {
                    if true {
                        v35 := GenS2{F0: 169}
                        _ = v35
                    }
                    v36 := GenS3{F0: 2147483647, F1: 100}
                    _ = v36
                } else {
                    if false {
                        if ((990 < 1) || (!((-(fi25)) > 42))) {
                            if true {
                                v37 := []int{980, fi28, fi25}
                                v38 := GenS1{F0: 2, F1: (v30 - v30), F2: 116}
                                var v39 uint8 = 142
                                var v40 float64 = ((((1e10 / 1.0) / float64(-2.25)) / 3.14159) * 1e10)
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        }
                        fmt.Println(fi25, fi28, fi25)
                        var v41 uint32 = 4294967295
                        _ = v41
                    } else {
                        v30 += (v30 - v30)
                    }
                    var v42 int64 = 0
                    _ = v42
                }
                _ = v30
            }
            _ = fi28
        }
        _ = fi25
    }
    v43 := GenS4{F0: 128}
    _ = v43
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 int16 = v2
    var v4 uint = 42
    var v5 int16 = 1
    var v6 int = -761
    _ = len(v1)
    for fi44 := 0; fi44 < 18; fi44++ {
        v5 -= (^(v2))
        v45 := GenS3{F0: -2147483648, F1: -1}
        if (1 == -1) {
            if false {
                var v46 float32 = -2.25
                v47 := []int{596, v6}
                _ = v46
                _ = v47
            } else {
                v48 := []int{115, (fi44 & fi44), (v6 ^ -9223372036854775808)}
                switch "" {
                case "a":
                    v3 -= v5
                case "":
                    for fi49 := 0; fi49 < 18; fi49++ {
                        v50 := GenS5{F0: fi49, F1: 1, F2: 0}
                        var v51 uint64 = 0
                        v2 = v3
                        v1[("" + "\n\t")] = "foo bar"
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                default:
                    v2 += v5
                }
                _ = v48
            }
            if true {
                fmt.Println(v2)
            } else {
                fmt.Println(v4, v2, v3)
                v4 -= v4
                v2 = -802
            }
            v2 = v5
        } else {
            v5 = v3
        }
        _ = fi44
        _ = v45
    }
    var v52 int64 = 0
    var v53 int8 = -46
    fmt.Println(v2)
    var v54 float64 = -2.25
    v54 = (3.14159 / -2.25)
    v55 := GenS4{F0: 249}
    if (-2147483648 == -1) {
        if false {
            v52 = (v52 << (((v4 << ((603) & 63))) & 63))
            fmt.Println(v4)
        } else {
            if (112 == 189) {
                switch v6 {
                case 32768:
                    v4 -= 2
                case 216:
                    for fi56 := 0; fi56 < 13; fi56++ {
                        for fi57 := 0; fi57 < 5; fi57++ {
                            fmt.Println(v5, v2, v5)
                            _ = fi57
                        }
                        if false {
                            switch (v6 << ((0) & 63)) {
                            case 0:
                                switch "hello" {
                                case "foo bar":
                                    switch v6 {
                                    case 128:
                                        v1[("a" + (((("foo bar" + "hello") + "") + "\n\t") + "foo bar"))] = ""
                                    case -40:
                                        var v58 float64 = (v54 / 3.14159)
                                        _ = v58
                                    }
                                case "hello":
                                    v54 -= (float64(-2.25) - v54)
                                case "":
                                    switch "hello" {
                                    case "hello":
                                        v52 += (v52 ^ v52)
                                    default:
                                        var v59 uint = v4
                                        _ = v59
                                    }
                                }
                            case 806:
                                var v60 int64 = 1
                                _ = v60
                            case -2147483648:
                                for fi61 := 0; fi61 < 13; fi61++ {
                                    v52 = v52
                                    _ = fi61
                                }
                            default:
                                if ((((4294967295 <= 2) || true) || (!("" >= ("foo bar" + "hello")))) && (!(!false))) {
                                    for fi62 := 0; fi62 < 6; fi62++ {
                                        var v63 float64 = (((float64(1.5) + 3.14159) * (float64(-2.25) / v54)) / 3.14159)
                                        var v64 int16 = -569
                                        _ = fi62
                                        _ = v63
                                        _ = v64
                                    }
                                    v52 += (-(v52))
                                    if false {
                                        v6 = fi56
                                    } else {
                                        var v65 int64 = v52
                                        var v66 int64 = (v52 | -1)
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    v53 -= v53
                                }
                            }
                            for fi67 := 0; fi67 < 2; fi67++ {
                                v68 := GenS6{F0: v53, F1: ("hello" + ("\n\t" + "hello"))}
                                v69 := GenS7{F0: 2, F1: 453, F2: 66}
                                v70 := []int{fi67}
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            if true {
                                var v71 int = (fi56 - (v6 >> ((304) & 63)))
                                v72 := []int{}
                                _ = len(v72)
                                _ = v71
                                _ = v72
                            } else {
                                v3 -= v5
                                if false {
                                    var v73 bool = true
                                    v74 := []int{127}
                                    _ = v73
                                    _ = v74
                                } else {
                                    var v75 int8 = v53
                                    if (382 < -2147483648) {
                                        for fi76 := 0; fi76 < 19; fi76++ {
                                            v77 := []int{}
                                            _ = fi76
                                            _ = v77
                                        }
                                    }
                                    fmt.Println(v52, v4)
                                    _ = v75
                                }
                            }
                        }
                        for fi78 := 0; fi78 < 6; fi78++ {
                            var v79 int = (fi56 + 0)
                            var v80 uint16 = 2
                            v81 := []int{v79}
                            _ = fi78
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        _ = fi56
                    }
                case 32767:
                    _ = len(v1)
                default:
                    v2 += (v2 * v3)
                }
                v52 = (v52 << (((^(v4))) & 63))
                v82 := GenS8{F0: v52, F1: 65535, F2: 1}
                _ = v82
            } else {
                delete(v1, "foo bar")
                v54 -= -2.25
            }
            var v83 uint64 = 2
            _ = v83
        }
        switch "foo bar" {
        case "a":
            switch ("hello" + "hello") {
            case "hello":
                if (true && true) {
                    v1["hello"] = "foo bar"
                    switch "hello" {
                    case "foo bar":
                        var v84 uint64 = 223
                        _ = v84
                    default:
                        v3 = v3
                    }
                } else {
                    v1["foo bar"] = ""
                    v85 := GenS9{F0: v3, F1: 0, F2: -58}
                    _ = v85
                }
            case "\n\t":
                for fi86 := 0; fi86 < 19; fi86++ {
                    var v87 bool = (42 < 2147483647)
                    switch (v6 >> (((^(v4))) & 63)) {
                    case 32767:
                        v52 = -284
                    case 1:
                        v88 := []int{v6}
                        _ = v88
                    }
                    _ = len(v1)
                    _ = fi86
                    _ = v87
                }
            case "a":
                v89 := GenS8{F0: int64(v53), F1: 2147483648, F2: 65535}
                _ = v89
            default:
                var v90 string = "foo bar"
                _ = v90
            }
        case "":
            switch v6 {
            case -534:
                if (42 > 614) {
                    v4 -= v4
                    v4 = 529
                } else {
                    var v91 uint64 = 978
                    _ = v91
                }
            }
        }
        var v92 string = ("hello" + "hello")
        _ = v92
    }
    for fi93 := 0; fi93 < 13; fi93++ {
        var v94 uint = v4
        for fi95 := 0; fi95 < 19; fi95++ {
            var v96 uint8 = 2
            _ = fi95
            _ = v96
        }
        var v97 int = -9223372036854775808
        _ = fi93
        _ = v94
        _ = v97
    }
    v98 := []int{(v6 ^ (v6 ^ v6))}
    if (17 >= (v53 - (-(v53)))) {
        for fi99 := 0; fi99 < 20; fi99++ {
            switch "foo bar" {
            case "a":
                v6 = fi99
            case "foo bar":
                switch "" {
                case "\n\t":
                    v52 += genFn0()
                case "foo bar":
                    var v100 float32 = 1e10
                    _ = v100
                }
            }
            v4 -= v4
            v52 = int64(v53)
            _ = fi99
        }
        var v101 int = v6
        var v102 uint8 = 2
        _ = v101
        _ = v102
    } else {
        var v103 int16 = (-(v3))
        if true {
            switch v6 {
            case 2147483647:
                v52 = (-(v52))
            }
        } else {
            for fi104 := 0; fi104 < 1; fi104++ {
                v6 -= (fi104 - fi104)
                _ = fi104
            }
            v103 = (v3 ^ v103)
        }
        _ = v103
    }
    v105 := GenS10{F0: 0, F1: 657}
    v53 = (^(v53))
    if ((v53 + -1) >= (v53 + (v53 - (v53 << ((42) & 63))))) {
        delete(v1, "a")
        fmt.Println(v54, v52)
    } else {
        switch "" {
        case "foo bar":
            v106 := GenS9{F0: (^(v3)), F1: 1, F2: 495}
            _ = v106
        case "hello":
            switch (^(v6)) {
            case 0:
                v107 := v1[("a" + "foo bar")]
                _ = v107
            case -1:
                for fi108 := 0; fi108 < 5; fi108++ {
                    v98 = append(v98, (fi108 | (v6 | v6)))
                    _ = fi108
                }
            case 112:
                for fi109 := 0; fi109 < 14; fi109++ {
                    v110 := GenS7{F0: 4294967295, F1: 1, F2: 30}
                    var v111 float32 = 1e10
                    _ = fi109
                    _ = v110
                    _ = v111
                }
            }
        }
        if true {
            delete(v1, "")
            _ = len(v98)
            var v112 uint64 = 422
            _ = v112
        } else {
            v2 += v2
        }
    }
    v2 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v98:", len(v98), v98)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
    _ = v98
}
