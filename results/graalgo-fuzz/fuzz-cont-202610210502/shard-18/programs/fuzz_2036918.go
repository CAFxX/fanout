package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 bool; F2 string }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 int64; F2 int64 }
type GenS11 struct { F0 int8; F1 int32; F2 bool }
type GenS12 struct { F0 int16; F1 int64; F2 int8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string; F1 int32; F2 int }
type GenS15 struct { F0 uint8; F1 bool }
type GenS16 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8, p1 uint16) int8 {
    var v5 uint32 = 42
    var v6 bool = (42 == -9223372036854775808)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 42
    v2 := []string{"foo bar"}
    v3 := []string{"hello", "\n\t", "\n\t"}
    var v4 int = 1
    var v7 string = "hello"
    if (false || false) {
        fmt.Println(v4)
        fmt.Println(v7, v4)
        v8 := GenS1{F0: 9223372036854775807}
        _ = v8
    } else {
        if len(v2) > 0 {
            v2[0] = (v7 + v7)
        }
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v7 = (v7 + v7)
        switch "hello" {
        case "\n\t":
            v1 = (v1 | v1)
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v1)
            _ = fi10
        }
        _ = fi9
    }
    if (false || false) {
        fmt.Println(v4)
        v1 += 128
    } else {
        if true {
            if ((float32(1e10) != 3.14159) && (761 < 1)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v4)
                v1 -= (^(v1))
            } else {
                if false {
                    var v11 uint = 594
                    if true {
                        v1 = v1
                        if (1.5 <= float32(0.0)) {
                            for fi12 := 0; fi12 < 1; fi12++ {
                                var v13 uint8 = v1
                                _ = fi12
                                _ = v13
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    } else {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            for fi15 := 0; fi15 < 7; fi15++ {
                                switch v7 {
                                case "hello":
                                    v16 := GenS1{F0: -270}
                                    _ = v16
                                case "foo bar":
                                    var v17 float64 = 3.14159
                                    _ = v17
                                default:
                                    v18 := GenS1{F0: -1}
                                    _ = v18
                                }
                                _ = fi15
                            }
                            _ = fi14
                        }
                        _ = len(v3)
                    }
                    _ = v11
                }
                var v19 int16 = 1000
                _ = v19
            }
            var v20 int64 = 9223372036854775807
            _ = v20
        } else {
            switch v7 {
            case "foo bar":
                var v21 int = 42
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 9; fi22++ {
                    if (42 != uint32(v1)) {
                        v23 := GenS1{F0: 807}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v24 string = "foo bar"
                        _ = v23
                        _ = v24
                    } else {
                        v4 = (v4 ^ 247)
                    }
                    _ = fi22
                }
            }
            fmt.Println(v1)
        }
    }
    _ = len(v2)
    v7 = v7
    fmt.Println(v1, v7, v1)
    for fi25 := 0; fi25 < 19; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v1, fi25)
            fmt.Println(fi26, v1, v4)
            if false {
                if len(v3) > 0 {
                    v3[0] = (v7 + (v7 + v7))
                }
                v1 -= 0
                var v27 uint8 = v1
                _ = v27
            } else {
                if false {
                    var v28 bool = false
                    _ = v28
                }
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            fmt.Println(fi29, v4, v4)
            _ = fi29
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            _ = len(v3)
            if false {
                v7 = ("" + v7)
                if (2 == 65535) {
                    v2 = append(v2, "\n\t")
                    if true {
                        if (!true) {
                            v7 = "a"
                            var v31 uint64 = 186
                            _ = v31
                        } else {
                            var v32 float32 = float32(1.5)
                            _ = v32
                        }
                    } else {
                        var v33 float32 = -2.25
                        if true {
                            v7 = ("foo bar" + v7)
                            v33 -= v33
                            var v34 float64 = ((-2.25 * 1e10) + (0.0 * float64(0.1)))
                            switch -81 {
                            case -831:
                                if ((false && true) && (true && (v33 >= float32((float64(((((3.14159 + (v33 - 0.0)) * 3.14159) * ((float32(1.5) + (v33 - 1.5)) + (3.14159 + 1e10))) - 0.0)) - v34))))) {
                                    switch v7 {
                                    case "a":
                                        switch v4 {
                                        case 32768:
                                            var v35 int8 = 68
                                            _ = v35
                                        case 108:
                                            v33 -= ((v33 + ((-2.25 / 0.1) * v33)) / (v33 / float32(v34)))
                                        case 42:
                                            v36 := GenS1{F0: 0}
                                            _ = v36
                                        default:
                                            fmt.Println(fi25, v34)
                                        }
                                        var v37 int64 = 9223372036854775807
                                        _ = len(v3)
                                        _ = v37
                                    case "hello":
                                        v7 = ((v7 + v7) + "hello")
                                    case "foo bar":
                                        var v38 uint8 = v1
                                        v39 := GenS1{F0: -9223372036854775808}
                                        fmt.Println(v7, v7)
                                        switch v7 {
                                        case "":
                                            if true {
                                                v34 = (((v34 / (3.14159 + (1.5 / 3.14159))) + 3.14159) + float64((v33 - (0.1 + ((-2.25 - 1.5) / v33)))))
                                                var v40 uint8 = 128
                                                _ = v40
                                            }
                                            v33 -= ((v33 + v33) / (((0.1 - 1e10) * float32(((v34 + v34) * float64((v33 - v33))))) * v33))
                                        default:
                                            var v41 uint64 = 4294967295
                                            _ = v41
                                        }
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v38, v34, v7)
                                        default:
                                            fmt.Println(v7, v34)
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                    v2 = append(v2, "foo bar")
                                } else {
                                    var v42 uint8 = v1
                                    v43 := GenS2{F0: 32767, F1: ((!(1 <= 4294967295)) && true)}
                                    if ((!false) || (0 != (fi25 - (fi30 - 9223372036854775807)))) {
                                        fmt.Println(v1)
                                        if (true || (((("\n\t" + "\n\t") + v7) + ((v7 + v7) + "\n\t")) <= ("a" + (v7 + "")))) {
                                            v44 := GenS3{F0: 60, F1: false, F2: v7}
                                            switch v4 {
                                            case -712:
                                                fmt.Println(fi30)
                                            case 2147483647:
                                                var v45 uint = 42
                                                _ = v45
                                            case 928:
                                                if false {
                                                    switch "hello" {
                                                    case "":
                                                        v42 = 224
                                                    case "a":
                                                        fmt.Println(v7, fi30)
                                                    default:
                                                        v34 = (0.1 * 0.1)
                                                    }
                                                }
                                            }
                                            _ = v44
                                        } else {
                                            v1 = (^(v1))
                                            var v46 uint64 = 4294967295
                                            _ = v46
                                        }
                                    } else {
                                        v47 := GenS4{F0: 65535, F1: ""}
                                        fmt.Println(v34, fi30, v33)
                                        _ = v47
                                    }
                                    _ = v42
                                    _ = v43
                                }
                            case -727:
                                fmt.Println(v7)
                            default:
                                fmt.Println(v1)
                            }
                            _ = v34
                        } else {
                            var v48 int64 = 1
                            _ = v48
                        }
                        _ = v33
                    }
                    var v49 int = (fi25 & (^(fi30)))
                    v4 -= (fi30 * 2147483647)
                    _ = v49
                } else {
                    v50 := GenS5{F0: 4294967295, F1: 2147483648}
                    _ = v50
                }
                if true {
                    var v51 uint16 = 0
                    fmt.Println(fi30, v1)
                    _ = len(v3)
                    _ = v51
                } else {
                    var v52 uint = 65535
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        if false {
                            switch v4 {
                            case 127:
                                _ = len(v3)
                            default:
                                var v53 int8 = genFn0(v1, 956)
                                _ = v53
                            }
                        } else {
                            switch ("hello" + v7) {
                            case "\n\t":
                                var v54 float64 = float64(0.1)
                                _ = v54
                            case "foo bar":
                                switch v4 {
                                case 42:
                                    fmt.Println(fi30)
                                case 32767:
                                    v55 := GenS6{F0: v52, F1: -193}
                                    _ = v55
                                }
                            case "":
                                v56 := GenS7{F0: (v52 % ((143) | 1)), F1: 32767, F2: 825}
                                _ = v56
                            }
                            if true {
                                if false {
                                    if (!(32767 == 2)) {
                                        var v57 int32 = 1
                                        if (!true) {
                                            if false {
                                                switch v4 {
                                                case 4294967295:
                                                    v58 := GenS7{F0: 42, F1: 957, F2: 594}
                                                    _ = v58
                                                case 377:
                                                    v52 -= v52
                                                case 127:
                                                    if false {
                                                        v7 = ""
                                                        if (!true) {
                                                            if len(v3) > 0 {
                                                                v3[0] = "\n\t"
                                                            }
                                                        }
                                                        v52 += 65535
                                                    }
                                                    fmt.Println(v4, v1, v4)
                                                    var v59 uint = (v52 | v52)
                                                    _ = v59
                                                }
                                            } else {
                                                if (-1 < -62) {
                                                    v60 := GenS5{F0: (v52 << ((v52) & 63)), F1: 815}
                                                    v3 = append(v3, v7)
                                                    fmt.Println(v4)
                                                    _ = v60
                                                } else {
                                                    var v61 uint64 = 2
                                                    var v62 bool = (false && ((v57 << ((v52) & 63)) > v57))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v61
                                                    _ = v62
                                                }
                                                v52 -= uint(v1)
                                                v1 = v1
                                                var v63 uint8 = v1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v63
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = v7
                                            }
                                            if true {
                                                var v64 int16 = 1
                                                v65 := GenS8{F0: (v52 % ((v52) | 1)), F1: v64}
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v4 = v4
                                            }
                                            v52 += (v52 + (v52 * (v52 | v52)))
                                            var v66 uint64 = 1
                                            _ = v66
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = ("hello" + v7)
                                            }
                                            if ((2 == 732) && (-1 >= 0)) {
                                                v67 := GenS9{F0: 127}
                                                _ = v67
                                            } else {
                                                var v68 uint = 42
                                                if true {
                                                    v52 -= (v52 / ((v68) | 1))
                                                    var v69 float32 = -2.25
                                                    if true {
                                                        var v70 uint64 = 2
                                                        var v71 float64 = 0.1
                                                        var v72 uint8 = (v1 << (((v52 % ((0) | 1))) & 63))
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    _ = v69
                                                }
                                                switch v4 {
                                                case -358:
                                                    var v73 float64 = (0.1 / 0.1)
                                                    _ = v73
                                                case 4294967295:
                                                    v68 = 775
                                                }
                                                var v74 int32 = (v57 - 42)
                                                _ = v68
                                                _ = v74
                                            }
                                            v57 = v57
                                        }
                                        fmt.Println(fi30)
                                        _ = v57
                                    } else {
                                        if ((!true) || true) {
                                            fmt.Println(v1, fi30, v4)
                                        } else {
                                            v75 := GenS6{F0: v52, F1: v4}
                                            v76 := GenS5{F0: 87, F1: 2147483648}
                                            switch v7 {
                                            case "a":
                                                _ = len(v2)
                                            case "foo bar":
                                                switch ((v7 + ("" + v7)) + v7) {
                                                case "foo bar":
                                                    var v77 uint16 = 42
                                                    _ = v77
                                                case "\n\t":
                                                    fmt.Println(fi25)
                                                default:
                                                    v78 := GenS10{F0: v4, F1: 9223372036854775807, F2: -1}
                                                    _ = v78
                                                }
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                        v79 := GenS11{F0: 105, F1: -504, F2: false}
                                        _ = v79
                                    }
                                    if (-77 == -133) {
                                        var v80 int8 = genFn0(v1, 2)
                                        if (4294967295 < 4294967295) {
                                            if false {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if true {
                                                        if true {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            fmt.Println(fi25)
                                                        } else {
                                                            v81 := GenS8{F0: v52, F1: 452}
                                                            _ = v81
                                                        }
                                                        v1 -= (^(v1))
                                                    } else {
                                                        v52 += (v52 + (v52 >> (((^(v52))) & 63)))
                                                    }
                                                case "\n\t":
                                                    v7 = v7
                                                case "hello":
                                                    fmt.Println(fi25, v80)
                                                    var v82 int64 = 0
                                                    _ = v82
                                                }
                                                var v83 int32 = -123
                                                if len(v2) > 0 {
                                                    v2[0] = v7
                                                }
                                                _ = v83
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            switch ((v7 + v7) + v7) {
                                            case "":
                                                v52 = v52
                                            default:
                                                v4 += (fi30 - (fi30 + (v4 % ((int(v80)) | 1))))
                                            }
                                        }
                                        if ((true || (!(42 < 42))) && ((0.1 / 1.0) >= -2.25)) {
                                            v84 := GenS12{F0: -543, F1: 42, F2: v80}
                                            _ = v84
                                        }
                                        _ = v80
                                    }
                                }
                                switch v7 {
                                case "hello":
                                    var v85 int32 = 1
                                    _ = v85
                                case "":
                                    v86 := GenS4{F0: (v52 / ((uint(v1)) | 1)), F1: v7}
                                    _ = v86
                                }
                                v87 := GenS13{F0: v52}
                                _ = v87
                            } else {
                                v7 = ((v7 + "") + v7)
                                var v88 int16 = -423
                                v7 = v7
                                _ = v88
                            }
                        }
                    }
                    if ((1.5 / 1e10) > -2.25) {
                        v89 := GenS14{F0: v7, F1: -1, F2: fi25}
                        _ = len(v3)
                        if (!false) {
                            v2 = append(v2, "a")
                            var v90 int32 = 1
                            v91 := GenS15{F0: v1, F1: true}
                            _ = v90
                            _ = v91
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        var v92 uint16 = 1
                        _ = v89
                        _ = v92
                    } else {
                        v1 = (v1 - v1)
                        fmt.Println(v7, v4, v7)
                        fmt.Println(v4, v4, fi30)
                        v93 := GenS16{F0: 45, F1: 42}
                        _ = v93
                    }
                    _ = v52
                }
            } else {
                switch (fi25 & fi30) {
                case -603:
                    var v94 uint8 = (v1 - (v1 - 10))
                    _ = v94
                default:
                    fmt.Println(v7, v1)
                }
                if len(v2) > 0 {
                    v2[0] = v7
                }
            }
            _ = fi30
        }
        _ = fi25
    }
    _ = len(v2)
    v3 = append(v3, v7)
    if ((!false) && ((3.14159 * 0.0) != ((((-2.25 - -2.25) / 3.14159) + 1e10) / 3.14159))) {
        for fi95 := 0; fi95 < 1; fi95++ {
            var v96 int32 = 42
            var v97 int32 = v96
            fmt.Println(v97, fi95, v97)
            _ = fi95
            _ = v96
            _ = v97
        }
        switch "" {
        case "a":
            v3 = append(v3, ((("foo bar" + v7) + (v7 + "\n\t")) + (v7 + "foo bar")))
        case "":
            v4 += v4
        case "\n\t":
            var v98 bool = (((((1.5 * 3.14159) + 1.5) <= (((0.0 + 3.14159) / float32(-2.25)) + 0.1)) || (("foo bar" + "\n\t") == v7)) && false)
            _ = v98
        }
    } else {
        for fi99 := 0; fi99 < 5; fi99++ {
            switch -755 {
            case 464:
                v4 -= v4
            case 1:
                var v100 uint64 = 373
                _ = v100
            }
            _ = fi99
        }
        var v101 uint32 = 2147483647
        fmt.Println(v7, v4, v7)
        _ = v101
    }
    var v102 int8 = 100
    v103 := GenS2{F0: 1, F1: true}
    var v104 int8 = v102
    var v105 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
