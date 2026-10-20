package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int8; F2 bool }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 uint8; F2 uint }
type GenS9 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint32 }

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 1
    var v3 uint = 0
    var v4 int8 = 100
    v5 := GenS1{F0: 0, F1: (v2 | (v2 + (v2 >> ((v3) & 63)))), F2: v3}
    fmt.Println(v4)
    v6 := []int{511}
    v6 = append(v6, 0)
    if true {
        var v7 int = 1
        _ = v7
    } else {
        var v8 float64 = v1
        if false {
            switch 747 {
            case 1:
                v2 += (^(v2))
            case 42:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case -9223372036854775808:
                if (((false || false) && false) && false) {
                    if len(v6) > 0 {
                        v6[0] = 122
                    }
                    if false {
                        fmt.Println(v1, v1, v3)
                        var v9 uint16 = v2
                        v10 := GenS2{F0: 42}
                        _ = v10
                        _ = v9
                    }
                } else {
                    v8 = 3.14159
                    if false {
                        v8 = v1
                        var v11 uint8 = 44
                        v11 -= 127
                        _ = v11
                    }
                }
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                switch ((("foo bar" + ("hello" + "hello")) + "\n\t") + ("a" + "hello")) {
                case "foo bar":
                    if (1 > 128) {
                        var v13 bool = (false || false)
                        _ = v13
                    } else {
                        _ = len(v6)
                        v14 := GenS1{F0: 114, F1: (v2 << ((v3) & 63)), F2: (v3 ^ (v3 + (^(v3))))}
                        _ = v14
                    }
                default:
                    if ((3.14159 - (((float32(((v1 - (1.5 + ((((v8 - ((v8 + v8) * (1e10 + float64(0.1)))) + (float64(1.5) + v8)) / 0.1) - (-2.25 - (v8 + 3.14159))))) / v8)) - (float32(((v8 - v1) - v8)) / 0.1)) + -2.25) + 1.5)) >= -2.25) {
                        var v15 int32 = -1
                        _ = v15
                    }
                }
                var v16 string = ""
                v17 := GenS2{F0: 0}
                _ = fi12
                _ = v16
                _ = v17
            }
        }
        _ = v8
    }
    var v18 uint64 = 0
    var v19 int = -485
    v1 = float64(float32((((v1 - ((v1 - -2.25) + v1)) - v1) / float64(-2.25))))
    v1 = v1
    fmt.Println(v1, v2, v1)
    v20 := GenS3{F0: true, F1: (v4 << ((892) & 63)), F2: (false || false)}
    var v21 uint16 = v2
    v18 = (v18 - 265)
    switch 32768 {
    case 9223372036854775807:
        for fi22 := 0; fi22 < 1; fi22++ {
            v2 = 568
            if (!(!false)) {
                v23 := GenS1{F0: 255, F1: uint16(v18), F2: v3}
                v4 -= (v4 & (^(v4)))
                v6 = append(v6, v19)
                _ = v23
            }
            _ = fi22
        }
    case 1:
        if ((!false) || false) {
            v1 = (v1 + v1)
            fmt.Println(v19, v3)
            var v24 uint8 = 20
            _ = v24
        } else {
            v1 += v1
            switch int(v3) {
            case 2147483647:
                var v25 int32 = -2147483648
                _ = v25
            case 32767:
                for fi26 := 0; fi26 < 6; fi26++ {
                    _ = len(v6)
                    _ = fi26
                }
            case 0:
                v4 = (-(v4))
            }
            if true {
                v27 := GenS3{F0: false, F1: 42, F2: true}
                _ = v27
            } else {
                v28 := GenS4{F0: 949, F1: v3}
                _ = v28
            }
        }
    case 128:
        for fi29 := 0; fi29 < 4; fi29++ {
            v4 += (v4 << (((v3 >> ((0) & 63))) & 63))
            fmt.Println(v19, v2, v3)
            _ = fi29
        }
    default:
        if false {
            var v30 uint16 = v2
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 string = ("foo bar" + ("foo bar" + ""))
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 int32 = 263
                switch 576 {
                case 4294967295:
                    var v35 uint32 = 2147483647
                    _ = v35
                case 586:
                    v36 := GenS5{F0: ("foo bar" + "")}
                    _ = v36
                }
                if ((float64(0.0) / (v1 - v1)) > 3.14159) {
                    if false {
                        switch -801 {
                        case 4294967295:
                            v19 = 176
                        case 0:
                            fmt.Println(fi33, v3, v2)
                        case -2147483648:
                            switch "" {
                            case "hello":
                                v37 := GenS6{F0: (1 > 9223372036854775807), F1: (v34 << (((^(v3))) & 63))}
                                _ = v37
                            case "foo bar":
                                if false {
                                    for fi38 := 0; fi38 < 3; fi38++ {
                                        v18 -= (v18 ^ (v18 << (((v3 * uint(v30))) & 63)))
                                        var v39 uint32 = 42
                                        _ = fi38
                                        _ = v39
                                    }
                                    var v40 bool = true
                                    if v40 {
                                        fmt.Println(v40, v34)
                                        if v40 {
                                            for fi41 := 0; fi41 < 19; fi41++ {
                                                fmt.Println(v34)
                                                if v40 {
                                                    switch (v19 - int(v4)) {
                                                    case 127:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case 2147483647:
                                                        v6 = append(v6, (fi33 * int(v4)))
                                                    default:
                                                        var v42 int64 = 42
                                                        _ = v42
                                                    }
                                                }
                                                if (!(!((v34 << ((v3) & 63)) > v34))) {
                                                    fmt.Println(v30)
                                                    if v40 {
                                                        v34 -= int32(v4)
                                                    }
                                                } else {
                                                    if (2 == 2147483647) {
                                                        if len(v6) > 0 {
                                                            v6[0] = (fi33 ^ -566)
                                                        }
                                                        if (571 == v3) {
                                                            if true {
                                                                var v43 uint = (v3 & 42)
                                                                v44 := GenS7{F0: ((28 == 0) || ((4294967295 != 0) && v40))}
                                                                v21 = (v2 | (v21 * 0))
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            if v40 {
                                                                v6 = append(v6, (-(fi41)))
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v45 uint64 = (^(v18))
                                                                    _ = v45
                                                                case "\n\t":
                                                                    v46 := GenS8{F0: 42, F1: 0, F2: (^(v3))}
                                                                    _ = v46
                                                                case "":
                                                                    v30 -= 1
                                                                }
                                                                if v40 {
                                                                    v47 := GenS2{F0: -855}
                                                                    _ = v47
                                                                }
                                                            } else {
                                                                var v48 int8 = -128
                                                                _ = v48
                                                            }
                                                        } else {
                                                            if true {
                                                                v18 += (v18 | v18)
                                                            } else {
                                                                if (255 < (v3 * (^(v3)))) {
                                                                    fmt.Println(v30)
                                                                    if ((!v40) && false) {
                                                                        v40 = v40
                                                                    } else {
                                                                        var v49 uint16 = v21
                                                                        _ = v49
                                                                    }
                                                                } else {
                                                                    v19 = fi41
                                                                    var v50 int32 = 42
                                                                    if (554 >= 0) {
                                                                        v51 := GenS9{F0: 35, F1: (v2 * (v21 * 1)), F2: 2}
                                                                        v18 -= v18
                                                                        var v52 uint8 = 42
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                    _ = v50
                                                                }
                                                                v34 -= v34
                                                            }
                                                        }
                                                        v53 := GenS2{F0: 9223372036854775807}
                                                        _ = v53
                                                    } else {
                                                        fmt.Println(v40, v18)
                                                    }
                                                }
                                                v54 := GenS9{F0: 0, F1: 1, F2: v18}
                                                v55 := GenS4{F0: v34, F1: v3}
                                                if (v40 || v40) {
                                                    var v56 uint32 = 1
                                                    if (v18 < v18) {
                                                        if v40 {
                                                            var v57 int32 = v34
                                                            _ = v57
                                                        }
                                                        switch (v19 ^ fi33) {
                                                        case 127:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        case -9223372036854775808:
                                                            var v58 uint32 = (^(v56))
                                                            _ = v58
                                                        }
                                                    } else {
                                                        if ((v34 ^ (v34 / ((0) | 1))) <= 2147483647) {
                                                            var v59 int = fi41
                                                            switch v19 {
                                                            case -266:
                                                                v30 += (v21 * 0)
                                                            case 32:
                                                                if len(v6) > 0 {
                                                                    v6[0] = (v19 / (((fi41 / ((v19) | 1))) | 1))
                                                                }
                                                            case 32768:
                                                                switch ("\n\t" + ("a" + ("\n\t" + ("\n\t" + "\n\t")))) {
                                                                case "":
                                                                    v60 := GenS2{F0: 615}
                                                                    _ = v60
                                                                case "\n\t":
                                                                    v2 += (v2 % (((v30 ^ (^(v30)))) | 1))
                                                                default:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v40 = v40
                                                                    case "hello":
                                                                        if len(v6) > 0 {
                                                                            v6[0] = int(v18)
                                                                        }
                                                                        v40 = false
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                    case "\n\t":
                                                                        v61 := GenS1{F0: 1, F1: (^(v21)), F2: (^(v3))}
                                                                        _ = v61
                                                                    }
                                                                }
                                                            default:
                                                                var v62 uint = 1
                                                                _ = v62
                                                            }
                                                            _ = v59
                                                        }
                                                        if ((1 > 0) || (!true)) {
                                                            var v63 int16 = 1000
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v40 = v40
                                                            _ = v63
                                                        } else {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v64 := GenS6{F0: ((v40 || false) || ((v3 | 326) > 4294967295)), F1: v34}
                                                                _ = v64
                                                            case "a":
                                                                if false {
                                                                    var v65 bool = (!v40)
                                                                    _ = v65
                                                                } else {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    if (381 < uint16(v3)) {
                                                                        var v66 uint64 = (v18 | 1)
                                                                        _ = v66
                                                                    } else {
                                                                        fmt.Println(v3, fi33)
                                                                        var v67 uint8 = 255
                                                                        if v40 {
                                                                            var v68 uint16 = v30
                                                                            var v69 uint32 = v56
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                    var v70 string = "\n\t"
                                                                    _ = v70
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v4, v21, v4)
                                                            default:
                                                                switch ("foo bar" + "a") {
                                                                case "":
                                                                    if v40 {
                                                                        var v71 uint = v3
                                                                        v21 = (v30 ^ v2)
                                                                        if len(v6) > 0 {
                                                                            v6[0] = -964
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        if v40 {
                                                                            var v72 uint32 = (v56 << (((^(v3))) & 63))
                                                                            v1 = 0.1
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v34 = v34
                                                                default:
                                                                    v73 := GenS10{F0: ((v18 <= 894) && false), F1: (v56 + v56), F2: (^(v56))}
                                                                    _ = v73
                                                                }
                                                            }
                                                        }
                                                        v74 := GenS3{F0: (true || false), F1: -1, F2: (-1 != -1)}
                                                        _ = v74
                                                    }
                                                    var v75 int = (fi41 / (((fi41 << ((0) & 63))) | 1))
                                                    _ = v56
                                                    _ = v75
                                                } else {
                                                    v40 = (v40 && ((v3 / ((213) | 1)) >= (v3 >> ((v3) & 63))))
                                                }
                                                var v76 int64 = -962
                                                fmt.Println(v3, v3)
                                                var v77 float64 = v1
                                                _ = fi41
                                                _ = v54
                                                _ = v55
                                                _ = v76
                                                _ = v77
                                            }
                                        } else {
                                            var v78 int64 = 9223372036854775807
                                            _ = v78
                                        }
                                        var v79 float64 = -2.25
                                        _ = v79
                                    }
                                    _ = v40
                                } else {
                                    v4 += v4
                                    fmt.Println(v4, v4)
                                }
                            default:
                                v6 = append(v6, fi33)
                            }
                        default:
                            v2 += v30
                        }
                    }
                } else {
                    v6 = append(v6, v19)
                    v80 := GenS7{F0: true}
                    var v81 int16 = -941
                    _ = v80
                    _ = v81
                }
                _ = fi33
                _ = v34
            }
            _ = v30
        }
    }
    v6 = append(v6, v19)
    fmt.Println(v2, v3, v21)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
