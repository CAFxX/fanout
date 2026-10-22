package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 int32; F1 int16; F2 uint }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint32; F1 int32; F2 int }
func genFn0(p0 int, p1 string) string {
    var v7 uint16 = 876
    var v8 int16 = -32768
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint8 = 128
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v5 string = "hello"
    var v6 int = (-(v1))
    v9 := []int{}
    v5 = (v5 + v5)
    v9 = append(v9, v6)
    v5 = ((v5 + (v5 + (v5 + v5))) + v5)
    if false {
        v5 = (v5 + (v5 + v5))
        if (true && false) {
            v2 += 1
            switch "\n\t" {
            case "foo bar":
                for fi10 := 0; fi10 < 19; fi10++ {
                    _ = len(v3)
                    for fi11 := 0; fi11 < 10; fi11++ {
                        if (!true) {
                            v1 = 839
                            fmt.Println(v1)
                        }
                        _ = fi11
                    }
                    _ = fi10
                }
            default:
                _ = len(v9)
            }
        } else {
            if (!((v2 / ((v2) | 1)) < 1)) {
                var v12 string = v5
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch fi13 {
                    case 1:
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 int64 = 42
                            if len(v9) > 0 {
                                v9[0] = fi14
                            }
                            switch "" {
                            case "":
                                var v16 bool = (!true)
                                _ = v16
                            }
                            _ = fi14
                            _ = v15
                        }
                    }
                    var v17 uint8 = (^(v2))
                    _ = fi13
                    _ = v17
                }
                v18 := GenS1{F0: 42, F1: 1, F2: v2}
                _ = v12
                _ = v18
            } else {
                v19 := GenS2{F0: 2, F1: true, F2: v6}
                _ = v19
            }
            if true {
                fmt.Println(v6)
            }
            fmt.Println(v2)
        }
    } else {
        v6 = (-(v1))
    }
    switch v5 {
    case "hello":
        for fi20 := 0; fi20 < 17; fi20++ {
            if (32767 < 65535) {
                switch v5 {
                case "":
                    for fi21 := 0; fi21 < 13; fi21++ {
                        _ = len(v3)
                        if false {
                            v5 = v5
                            v22 := GenS1{F0: -56, F1: -188, F2: v2}
                            _ = v22
                        }
                        _ = fi21
                    }
                case "\n\t":
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 uint16 = 2
                        var v25 string = v5
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            } else {
                switch (v5 + "\n\t") {
                case "hello":
                    v26 := GenS2{F0: uint64(v2), F1: false, F2: (fi20 ^ v1)}
                    _ = v26
                default:
                    switch "" {
                    case "a":
                        fmt.Println(v6)
                    case "\n\t":
                        for fi27 := 0; fi27 < 11; fi27++ {
                            switch (((v5 + v5) + "\n\t") + (v5 + v5)) {
                            case "":
                                v28 := GenS3{F0: 59, F1: -1, F2: 611}
                                _ = v28
                            case "foo bar":
                                fmt.Println(v2, fi20, fi27)
                            case "hello":
                                v6 -= (fi20 >> ((2) & 63))
                            default:
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            if len(v9) > 0 {
                                v9[0] = (v6 * 32768)
                            }
                            _ = fi27
                        }
                    }
                }
                if false {
                    v29 := GenS1{F0: -1, F1: -2147483648, F2: v2}
                    _ = v29
                }
            }
            var v30 int64 = 42
            _ = fi20
            _ = v30
        }
    case "a":
        v2 = (v2 - v2)
    case "":
        var v31 float64 = (1.5 / 0.1)
        _ = v31
    default:
        switch 32768 {
        case 0:
            switch v5 {
            case "":
                v32 := GenS2{F0: 42, F1: false, F2: -9223372036854775808}
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 14; fi33++ {
                    if true {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            _ = len(v9)
                            v35 := GenS1{F0: -1, F1: -2147483648, F2: v2}
                            if false {
                                if true {
                                    var v36 int = -2147483648
                                    fmt.Println(v5, fi33)
                                    v5 = (v5 + (v5 + (("\n\t" + v5) + v5)))
                                    _ = v36
                                }
                                if false {
                                    v37 := GenS4{F0: 196, F1: v2}
                                    if false {
                                        switch v5 {
                                        case "":
                                            v1 -= (fi34 >> ((1) & 63))
                                        case "a":
                                            v38 := GenS3{F0: 682, F1: 0, F2: 0}
                                            _ = v38
                                        case "hello":
                                            switch 32767 {
                                            case -1:
                                                var v39 uint32 = 372
                                                _ = v39
                                            case 1:
                                                switch 42 {
                                                case 173:
                                                    switch (v5 + "hello") {
                                                    case "hello":
                                                        var v40 uint = 42
                                                        _ = v40
                                                    case "foo bar":
                                                        if true {
                                                            v6 -= 4294967295
                                                        }
                                                    case "":
                                                        v2 -= v2
                                                    }
                                                case 42:
                                                    if true {
                                                        switch (fi33 * fi33) {
                                                        case -1:
                                                            v6 += (fi33 - (v6 + -9223372036854775808))
                                                        }
                                                    }
                                                case 298:
                                                    var v41 int16 = 808
                                                    _ = v41
                                                default:
                                                    var v42 int = v1
                                                    _ = v42
                                                }
                                                var v43 bool = true
                                                if v43 {
                                                    var v44 float32 = (0.0 + 1e10)
                                                    if v43 {
                                                        if (1e10 == v44) {
                                                            if true {
                                                                v45 := GenS1{F0: -793, F1: 0, F2: v2}
                                                                fmt.Println(v44)
                                                                v46 := v4[42]
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        }
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        v1 -= fi33
                                                        var v47 int32 = 1
                                                        v6 += fi34
                                                        _ = v47
                                                    }
                                                    v2 = (v2 << ((369) & 63))
                                                    var v48 string = "\n\t"
                                                    if len(v9) > 0 {
                                                        v9[0] = (^(fi34))
                                                    }
                                                    v49 := GenS2{F0: 18446744073709551615, F1: true, F2: (fi33 & fi34)}
                                                    v6 = (fi34 - (-(v1)))
                                                    _ = v44
                                                    _ = v48
                                                    _ = v49
                                                }
                                                _ = v43
                                            }
                                        default:
                                            v50 := GenS1{F0: 973, F1: 0, F2: 240}
                                            _ = v50
                                        }
                                    }
                                    var v51 int64 = -479
                                    v52 := GenS5{F0: v51}
                                    if false {
                                        v9 = append(v9, fi33)
                                        v1 += fi34
                                    }
                                    _ = v37
                                    _ = v51
                                    _ = v52
                                } else {
                                    v1 = (v1 | (v1 | (v1 * fi33)))
                                    v9 = append(v9, -1)
                                    switch -14 {
                                    case -2147483648:
                                        v53 := GenS6{F0: v2, F1: 18446744073709551615}
                                        _ = v53
                                    default:
                                        v2 = v2
                                    }
                                }
                                var v54 uint8 = (^(v2))
                                _ = v54
                            }
                            _ = fi34
                            _ = v35
                        }
                    }
                    _ = fi33
                }
            case "foo bar":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            default:
                var v55 int8 = 12
                _ = v55
            }
        }
    }
    fmt.Println(v5, v1)
    var v56 int8 = 42
    for fi57 := 0; fi57 < 10; fi57++ {
        if true {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v1 -= fi57
        } else {
            v2 = v2
            fmt.Println(v6)
            if true {
                if ((0.0 * float32(-2.25)) != -2.25) {
                    switch (-(v6)) {
                    case 2147483647:
                        v56 += v56
                    case 1:
                        v9 = append(v9, (v1 << ((2) & 63)))
                    default:
                        var v58 float32 = 0.0
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 17; fi59++ {
                        delete(v3, "foo bar")
                        v5 = v5
                        switch 127 {
                        case 498:
                            v56 -= v56
                        case -2147483648:
                            var v60 uint8 = 44
                            _ = v60
                        case 0:
                            if false {
                                var v61 uint32 = 4294967295
                                v5 = v5
                                if false {
                                    var v62 uint16 = 65535
                                    fmt.Println(fi57)
                                    var v63 int8 = v56
                                    _ = len(v4)
                                    _ = v62
                                    _ = v63
                                } else {
                                    var v64 uint64 = 18446744073709551615
                                    v6 += (fi57 - v1)
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = v64
                                }
                                _ = v61
                            }
                        default:
                            v65 := GenS4{F0: 42, F1: (v2 | (^(v2)))}
                            _ = v65
                        }
                        _ = fi59
                    }
                } else {
                    if (!((!(!false)) && (!(!true)))) {
                        switch -9223372036854775808 {
                        case 32768:
                            for fi66 := 0; fi66 < 5; fi66++ {
                                v4[(fi66 ^ v6)] = ("foo bar" + v5)
                                v6 = (v6 - -2147483648)
                                if (18446744073709551615 <= 1) {
                                    v6 = 963
                                } else {
                                    fmt.Println(fi57, v2, v6)
                                    v67 := GenS3{F0: -274, F1: -1, F2: 2}
                                    v68 := GenS3{F0: -1, F1: -231, F2: 65535}
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi66
                            }
                        default:
                            if (true && false) {
                                fmt.Println(v1, v6, v5)
                                var v69 bool = (v5 >= (("foo bar" + ("\n\t" + "")) + v5))
                                if len(v9) > 0 {
                                    v9[0] = (fi57 * v1)
                                }
                                _ = v69
                            } else {
                                if false {
                                    v70 := GenS2{F0: 1, F1: ((!(!(!((float64(1e10) + float64(((1e10 * -2.25) * float32(0.0)))) <= 0.0)))) && (true || (0 == v2))), F2: (v1 + fi57)}
                                    if true {
                                        v6 = v1
                                        var v71 float64 = 0.0
                                        for fi72 := 0; fi72 < 14; fi72++ {
                                            fmt.Println(fi72)
                                            fmt.Println(v2, v1, v6)
                                            fmt.Println(v1)
                                            v9 = append(v9, (fi72 - v6))
                                            _ = fi72
                                        }
                                        _ = v71
                                    } else {
                                        switch v1 {
                                        case 611:
                                            if (!((v6 ^ 4294967295) == fi57)) {
                                                for fi73 := 0; fi73 < 2; fi73++ {
                                                    if false {
                                                        switch (-(v6)) {
                                                        case -1:
                                                            delete(v4, -9223372036854775808)
                                                        case 1:
                                                            delete(v4, int(v56))
                                                        case 9223372036854775807:
                                                            var v74 bool = true
                                                            _ = v74
                                                        }
                                                        fmt.Println(v2, v6, v6)
                                                        var v75 float64 = -2.25
                                                        var v76 int = (fi57 ^ -436)
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    _ = fi73
                                                }
                                            } else {
                                                switch -978 {
                                                case -2147483648:
                                                    _ = len(v9)
                                                }
                                            }
                                        default:
                                            v1 = (v6 / (((v6 + v1)) | 1))
                                        }
                                    }
                                    _ = v70
                                }
                                for fi77 := 0; fi77 < 3; fi77++ {
                                    v56 -= v56
                                    v56 += (^(v56))
                                    _ = fi77
                                }
                            }
                        }
                    } else {
                        v4[2147483647] = v5
                    }
                }
                var v78 uint8 = (v2 / ((v2) | 1))
                v2 = v78
                _ = v78
            } else {
                if false {
                    v79 := GenS7{F0: 2147483648, F1: -405, F2: (v1 << ((405) & 63))}
                    var v80 string = v5
                    _ = v79
                    _ = v80
                }
                v81 := GenS2{F0: 737, F1: false, F2: v1}
                _ = v81
            }
        }
        v1 += v1
        _ = fi57
    }
    v56 -= (v56 * 29)
    v4[v1] = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v9
}
