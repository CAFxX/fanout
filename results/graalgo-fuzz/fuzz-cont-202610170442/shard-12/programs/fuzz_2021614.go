package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 string; F1 string; F2 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS9 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -835
    v3 := make(map[int]uint32)
    var v4 uint8 = 128
    switch "a" {
    case "hello":
        switch 32767 {
        case -239:
            v5 := v1[231]
            _ = v5
        case 42:
            var v6 int = -1
            _ = v6
        case -474:
            if (!(false && (!true))) {
                v7 := GenS1{F0: (133 >= 18446744073709551615), F1: 445, F2: 0}
                for fi8 := 0; fi8 < 8; fi8++ {
                    if false {
                        var v9 bool = (!(!false))
                        v10 := GenS1{F0: (!((-32768 != -530) || (!((v2 % (((v2 | 1)) | 1)) >= 1)))), F1: -9223372036854775808, F2: 100}
                        if ((fi8 / (((fi8 | fi8)) | 1)) >= 2147483647) {
                            var v11 uint32 = 847
                            v12 := []int{(fi8 - (fi8 ^ -2147483648)), 567}
                            _ = v11
                            _ = v12
                        } else {
                            var v13 string = "foo bar"
                            fmt.Println(v9, v2, fi8)
                            _ = v13
                        }
                        _ = v10
                        _ = v9
                    } else {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            switch (("a" + "foo bar") + ("a" + "")) {
                            case "":
                                var v15 int64 = 42
                                _ = v15
                            case "hello":
                                v16 := GenS1{F0: (2147483648 > 73), F1: -588, F2: 1}
                                _ = v16
                            case "a":
                                switch (("" + "") + "") {
                                case "":
                                    var v17 string = "\n\t"
                                    _ = v17
                                case "a":
                                    switch (fi8 / ((fi8) | 1)) {
                                    case 127:
                                        _ = len(v3)
                                    case 42:
                                        if true {
                                            fmt.Println(fi14)
                                            if true {
                                                if true {
                                                    var v18 uint = 62
                                                    v19 := []int{(-(fi8)), 9223372036854775807, (fi14 | (^(fi14)))}
                                                    if ((2 != 4294967295) && false) {
                                                        var v20 bool = (-1 > -32768)
                                                        if v20 {
                                                            var v21 int32 = v2
                                                            _ = v21
                                                        }
                                                        _ = v20
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if ((v2 + -2147483648) >= 42) {
                                                                var v22 uint8 = v4
                                                                _ = v22
                                                            }
                                                        }
                                                        if false {
                                                            var v23 int64 = 9223372036854775807
                                                            if false {
                                                                switch ("hello" + ("\n\t" + ("hello" + ((("foo bar" + ("" + ("hello" + ("\n\t" + "\n\t")))) + "") + "hello")))) {
                                                                case "\n\t":
                                                                    _ = len(v19)
                                                                case "a":
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        _ = len(v19)
                                                                    case "a":
                                                                        var v24 uint32 = 859
                                                                        _ = v24
                                                                    default:
                                                                        if true {
                                                                            var v25 int32 = 1
                                                                            _ = v25
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    v3[(fi8 | -803)] = 760
                                                                }
                                                                var v26 uint8 = v4
                                                                fmt.Println(v2)
                                                                _ = v26
                                                            } else {
                                                                _ = len(v19)
                                                                if (true && false) {
                                                                    if true {
                                                                        var v27 uint16 = 255
                                                                        v3[fi8] = 2
                                                                        _ = v27
                                                                    } else {
                                                                        v28 := GenS2{F0: 18446744073709551615, F1: 672, F2: 4294967295}
                                                                        var v29 uint32 = 42
                                                                        v2 = v2
                                                                        _ = v28
                                                                        _ = v29
                                                                    }
                                                                }
                                                                v30 := GenS3{F0: 42, F1: 0}
                                                                _ = v30
                                                            }
                                                            _ = v23
                                                        } else {
                                                            v31 := GenS4{F0: "foo bar", F1: "a", F2: 1}
                                                            fmt.Println(fi8, fi8, v18)
                                                            v32 := GenS2{F0: 42, F1: 32767, F2: 65535}
                                                            if (42 != 932) {
                                                                var v33 int16 = 980
                                                                v34 := GenS1{F0: ((2147483647 == 540) && ((3.14159 >= 3.14159) || true)), F1: 42, F2: 127}
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        v19 = append(v19, 419)
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v35 bool = true
                                                    _ = v35
                                                case "a":
                                                    v4 -= v4
                                                case "hello":
                                                    var v36 int64 = -1
                                                    _ = v36
                                                }
                                                v37 := GenS5{F0: -1}
                                                _ = v37
                                            }
                                            v38 := GenS4{F0: "", F1: "hello", F2: 255}
                                            _ = v38
                                        } else {
                                            if ("\n\t" > ("a" + "foo bar")) {
                                                var v39 float32 = 1e10
                                                fmt.Println(v4, v4)
                                                _ = v39
                                            } else {
                                                var v40 uint8 = 128
                                                if (false && true) {
                                                    fmt.Println(fi8, fi8)
                                                    fmt.Println(v4, fi8, v4)
                                                    v41 := GenS4{F0: "\n\t", F1: ("hello" + "\n\t"), F2: 0}
                                                    _ = v41
                                                }
                                                if false {
                                                    if false {
                                                        if true {
                                                            switch "foo bar" {
                                                            case "":
                                                                if (false || false) {
                                                                    v42 := []int{fi14, (fi8 % (((-(fi14))) | 1)), 32768, fi8}
                                                                    v43 := GenS1{F0: true, F1: -1, F2: 127}
                                                                    if true {
                                                                        v44 := GenS3{F0: 1, F1: 42}
                                                                        _ = v44
                                                                    } else {
                                                                        if (375 <= 0) {
                                                                            var v45 uint64 = 182
                                                                            _ = v45
                                                                        } else {
                                                                            var v46 int = fi14
                                                                            var v47 float32 = 1e10
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                        var v48 uint = 0
                                                                        _ = len(v42)
                                                                        v49 := GenS2{F0: 2, F1: 89, F2: 2147483648}
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                } else {
                                                                    var v50 uint = 65535
                                                                    var v51 uint64 = 1
                                                                    v40 = v4
                                                                    v40 = 128
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            }
                                                            var v52 float64 = float64((0.1 * (((-2.25 / ((0.1 + 1.5) * 1e10)) / (-2.25 * 1e10)) / (1.5 / -2.25))))
                                                            v3[4294967295] = 65535
                                                            _ = v52
                                                        }
                                                        v53 := []int{fi8}
                                                        fmt.Println(v2, fi8, v4)
                                                        _ = v53
                                                    }
                                                } else {
                                                    var v54 int = fi8
                                                    _ = v54
                                                }
                                                v55 := GenS6{F0: 42}
                                                _ = v40
                                                _ = v55
                                            }
                                            switch (^(fi14)) {
                                            case 32767:
                                                var v56 uint = 961
                                                _ = v56
                                            default:
                                                var v57 int8 = 127
                                                if true {
                                                    var v58 bool = true
                                                    var v59 uint16 = 455
                                                    v58 = true
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    var v60 int32 = v2
                                                    v4 = (v4 % ((v4) | 1))
                                                    _ = v60
                                                }
                                                _ = v57
                                            }
                                        }
                                    default:
                                        _ = len(v3)
                                    }
                                case "\n\t":
                                    v61 := []int{fi14, (fi14 << ((187) & 63)), (-(fi8))}
                                    _ = v61
                                default:
                                    v2 += 0
                                }
                            default:
                                v4 = v4
                                if true {
                                    var v62 int = 127
                                    fmt.Println(fi8, v2)
                                    if (-9223372036854775808 < 9223372036854775807) {
                                        v63 := GenS2{F0: 18446744073709551615, F1: 32767, F2: 0}
                                        _ = v63
                                    }
                                    v64 := GenS6{F0: 0}
                                    _ = v62
                                    _ = v64
                                } else {
                                    var v65 uint = 1
                                    _ = v65
                                }
                                var v66 uint16 = 255
                                _ = v66
                            }
                            _ = fi14
                        }
                    }
                    v2 = (v2 % ((v2) | 1))
                    v67 := GenS7{F0: uint16(v4)}
                    _ = fi8
                    _ = v67
                }
                if false {
                    v68 := []int{1, 2147483647, 4294967295}
                    _ = v68
                }
                _ = v7
            } else {
                v4 -= 255
            }
        }
    case "a":
        for fi69 := 0; fi69 < 3; fi69++ {
            if (0 != -1) {
                fmt.Println(fi69, v4, fi69)
            }
            _ = fi69
        }
    }
    var v70 uint8 = v4
    v2 += -391
    var v71 float32 = (1.5 / 3.14159)
    var v72 int = 4294967295
    if true {
        switch v72 {
        case 42:
            var v73 bool = true
            _ = v73
        case 106:
            var v74 uint64 = 190
            _ = v74
        case 4294967295:
            var v75 uint32 = 2
            _ = v75
        }
        for fi76 := 0; fi76 < 12; fi76++ {
            v77 := []int{v72, (-(v72)), fi76, v72}
            _ = fi76
            _ = v77
        }
        for fi78 := 0; fi78 < 2; fi78++ {
            v79 := []int{v72, (fi78 % (((fi78 >> ((65535) & 63))) | 1))}
            for fi80 := 0; fi80 < 20; fi80++ {
                delete(v3, 246)
                v81 := GenS1{F0: (true || (true || false)), F1: 42, F2: 32}
                switch v72 {
                case 9223372036854775807:
                    v82 := GenS8{F0: -804, F1: 42, F2: 206}
                    _ = v82
                case 0:
                    fmt.Println(v71, v72, v71)
                case 32768:
                    switch fi80 {
                    case 788:
                        if true {
                            fmt.Println(v2, v71, v70)
                            if false {
                                var v83 float64 = (0.1 - 0.1)
                                v84 := GenS6{F0: 860}
                                v71 += 1.5
                                var v85 string = ("a" + (("a" + ("hello" + "")) + "foo bar"))
                                v4 -= v70
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                            if len(v79) >= 2 {
                                v79 = v79[1:len(v79)]
                            }
                        }
                    default:
                        if (fi80 >= 9223372036854775807) {
                            v70 += (v4 - v70)
                            fmt.Println(fi78, v71)
                        }
                    }
                }
                _ = fi80
                _ = v81
            }
            v71 -= v71
            _ = fi78
            _ = v79
        }
    } else {
        v4 = v70
        var v86 int64 = 1
        var v87 int = (v72 & v72)
        _ = v86
        _ = v87
    }
    v88 := GenS3{F0: uint32(v4), F1: 158}
    _ = len(v1)
    switch ((("hello" + "\n\t") + ("\n\t" + ("foo bar" + ""))) + "foo bar") {
    case "\n\t":
        for fi89 := 0; fi89 < 19; fi89++ {
            var v90 uint8 = v70
            _ = fi89
            _ = v90
        }
    case "foo bar":
        var v91 int16 = -32768
        _ = v91
    }
    for fi92 := 0; fi92 < 2; fi92++ {
        v93 := GenS9{F0: 32767}
        var v94 uint8 = 136
        switch fi92 {
        case 624:
            delete(v1, v72)
        default:
            v95 := []int{(fi92 / ((fi92) | 1)), (-(v72)), 716}
            _ = v95
        }
        _ = fi92
        _ = v93
        _ = v94
    }
    var v96 bool = false
    v97 := v3[(v72 >> ((2) & 63))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v72
    _ = v88
    _ = v96
    _ = v97
}
