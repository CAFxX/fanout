package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int8; F1 string; F2 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 int16; F2 string }
type GenS10 struct { F0 bool; F1 uint64; F2 bool }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 string; F1 uint64; F2 int32 }
type GenS16 struct { F0 bool; F1 int64 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int8; F1 bool; F2 int32 }
type GenS19 struct { F0 uint32; F1 int; F2 int32 }
type GenS20 struct { F0 uint64; F1 uint; F2 int }
type GenS21 struct { F0 uint32; F1 int; F2 uint16 }
type GenS22 struct { F0 uint32; F1 uint32; F2 bool }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 uint; F1 string }
type GenS26 struct { F0 uint8; F1 int8; F2 bool }
type GenS27 struct { F0 uint32 }
type GenS28 struct { F0 string; F1 uint8; F2 int16 }
type GenS29 struct { F0 int8; F1 uint8 }
type GenS30 struct { F0 int32; F1 int32; F2 string }
type GenS31 struct { F0 int; F1 bool; F2 uint }
type GenS32 struct { F0 uint16 }
type GenS33 struct { F0 string; F1 int }
type GenS34 struct { F0 uint32 }
type GenS35 struct { F0 uint16; F1 bool }
type GenS36 struct { F0 int16; F1 int16; F2 uint16 }
type GenS37 struct { F0 string }
type GenS38 struct { F0 uint16 }
func genFn0(p0 int8) int16 {
    v5 := []int{}
    p0 -= p0
    for fi6 := 0; fi6 < 20; fi6++ {
        if (1.5 < float64(0.1)) {
            var v7 bool = (!(!true))
            if v7 {
                var v8 uint64 = 0
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(fi6, p0)
                    var v10 float32 = 1e10
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                for fi11 := 0; fi11 < 17; fi11++ {
                    v7 = (false && v7)
                    var v12 bool = ((42 == -334) && (fi6 >= (fi6 - (fi6 ^ 0))))
                    v13 := GenS1{F0: fi11, F1: 24, F2: 2}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = v7
        } else {
            var v14 uint64 = 1
            fmt.Println(fi6, p0, p0)
            _ = v14
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint16 = 42
        v5 = append(v5, 9223372036854775807)
        var v17 int16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v5
    return -32768
}

func main() {
    v1 := []string{"hello", "", "hello", ("a" + ("\n\t" + ("hello" + "")))}
    v2 := make(map[string]int)
    var v3 bool = (true && (false || false))
    v4 := []int{}
    v18 := GenS2{F0: 410, F1: 53, F2: 42}
    if v3 {
        v3 = (0 == -401)
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            _ = len(v4)
            var v20 uint = 18446744073709551615
            _ = fi19
            _ = v20
        }
        delete(v2, ("hello" + "hello"))
    }
    v3 = false
    fmt.Println(v3)
    v4 = append(v4, -184)
    var v21 int16 = -1
    delete(v2, ("a" + "a"))
    switch 0 {
    case 286:
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS2{F0: 2147483647, F1: 91, F2: (v21 << ((255) & 63))}
            _ = fi22
            _ = v23
        }
    case -9223372036854775808:
        if v3 {
            fmt.Println(v21)
        } else {
            if true {
                v3 = (9223372036854775807 > -545)
                v21 += (v21 + -1)
                switch "a" {
                case "a":
                    v3 = v3
                case "\n\t":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        v25 := GenS2{F0: 2147483647, F1: 127, F2: v21}
                        fmt.Println(v21)
                        _ = fi24
                        _ = v25
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 3; fi26++ {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            var v28 int8 = 127
                            _ = fi27
                            _ = v28
                        }
                        v3 = v3
                        _ = fi26
                    }
                default:
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "\n\t":
                        v3 = false
                    case "foo bar":
                        if (v3 || (0 <= 622)) {
                            switch -2147483648 {
                            case -766:
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            case 32768:
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    fmt.Println(fi29, v21)
                                    v3 = v3
                                    if v3 {
                                        if v3 {
                                            fmt.Println(v3)
                                            fmt.Println(v3, v3, fi29)
                                            var v30 uint16 = 32767
                                            _ = v30
                                        } else {
                                            var v31 int8 = -6
                                            v32 := GenS2{F0: 4294967295, F1: 63, F2: -1}
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    _ = fi29
                                }
                            }
                            var v33 int32 = -2147483648
                            fmt.Println(v3, v3)
                            _ = v33
                        } else {
                            v21 -= v21
                            v34 := GenS1{F0: 4294967295, F1: 0, F2: 18446744073709551615}
                            _ = v34
                        }
                    case "":
                        fmt.Println(v3)
                    default:
                        v21 = v21
                    }
                }
            }
            var v35 uint64 = 4294967295
            _ = v35
        }
    }
    v3 = v3
    switch -1 {
    case 128:
        fmt.Println(v3, v3)
    case 2147483647:
        for fi36 := 0; fi36 < 14; fi36++ {
            switch (^(fi36)) {
            case 32768:
                if false {
                    v4 = append(v4, fi36)
                    var v37 int32 = -927
                    _ = v37
                }
            case -894:
                v21 = (v21 * (v21 + 1000))
            case 284:
                for fi38 := 0; fi38 < 15; fi38++ {
                    switch "" {
                    case "a":
                        v4 = append(v4, (fi38 + (fi36 - fi36)))
                        v3 = v3
                        if (267 >= 32767) {
                            var v39 float32 = float32(1e10)
                            v3 = false
                            _ = len(v4)
                            _ = v39
                        } else {
                            if (v3 || false) {
                                if true {
                                    if v3 {
                                        v3 = (("a" + "foo bar") < ("foo bar" + ("\n\t" + "\n\t")))
                                        if false {
                                            var v40 uint = 775
                                            fmt.Println(v3, v21)
                                            if v3 {
                                                switch fi38 {
                                                case -1:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v2[(("" + "foo bar") + "hello")] = (fi38 & -1)
                                                    case "":
                                                        _ = len(v2)
                                                    case "a":
                                                        v41 := GenS3{F0: -102}
                                                        _ = v41
                                                    default:
                                                        v21 = v21
                                                    }
                                                case 32767:
                                                    if (1 < -1) {
                                                        v3 = (float32(((1e10 + float64((1e10 * (0.1 + (1e10 / float32(0.1)))))) * (float64(0.0) * 0.0))) != (1.5 / 0.1))
                                                        fmt.Println(v40, v3)
                                                    } else {
                                                        var v42 uint = v40
                                                        var v43 int = fi36
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    v21 = (v21 % ((genFn0(100)) | 1))
                                                }
                                                v40 += v40
                                                v44 := GenS4{F0: "a"}
                                                _ = v44
                                            } else {
                                                _ = len(v2)
                                                var v45 uint8 = 42
                                                if (true && (!true)) {
                                                    v40 = v40
                                                    v21 = (v21 * 42)
                                                    if v3 {
                                                        if v3 {
                                                            v46 := GenS5{F0: 9223372036854775807, F1: 4294967295}
                                                            fmt.Println(v21, fi38)
                                                            _ = v46
                                                        } else {
                                                            fmt.Println(fi36, fi38)
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v47 float64 = ((((1e10 - (1.5 + -2.25)) / 1e10) * ((3.14159 + -2.25) * (((0.0 + 1e10) + 0.1) * ((3.14159 * (3.14159 - 1e10)) / 3.14159)))) * 3.14159)
                                                            _ = v47
                                                        }
                                                        v45 -= v45
                                                        switch 128 {
                                                        case 4294967295:
                                                            _ = len(v2)
                                                        case 556:
                                                            var v48 int64 = 643
                                                            _ = v48
                                                        case 32768:
                                                            v49 := GenS4{F0: "\n\t"}
                                                            fmt.Println(v45, v3)
                                                            if (int32(v40) != -711) {
                                                                var v50 uint8 = v45
                                                                v45 += (v45 >> (((v40 & v40)) & 63))
                                                                var v51 bool = v3
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            v21 = v21
                                                            if v3 {
                                                                v52 := GenS6{F0: 103, F1: "foo bar", F2: 505}
                                                                _ = v52
                                                            }
                                                            v45 = v45
                                                            switch 127 {
                                                            case -565:
                                                                v40 -= (^(v40))
                                                            default:
                                                                if v3 {
                                                                    if (4294967295 != 2) {
                                                                        if (v3 && (!(0 != 2))) {
                                                                            var v53 bool = v3
                                                                            v40 = v40
                                                                            fmt.Println(fi38)
                                                                            if true {
                                                                                switch fi38 {
                                                                                case 42:
                                                                                    v53 = v3
                                                                                    v54 := GenS6{F0: 100, F1: "foo bar", F2: -1}
                                                                                    _ = v54
                                                                                case 32768:
                                                                                    v2[(("" + "\n\t") + (("a" + "\n\t") + (((("a" + "foo bar") + ("" + "a")) + ("" + "\n\t")) + "hello")))] = 1
                                                                                case -1:
                                                                                    var v55 int = fi38
                                                                                    _ = v55
                                                                                }
                                                                                if ((v21 ^ v21) < (v21 + v21)) {
                                                                                    var v56 int64 = 388
                                                                                    var v57 int16 = int16(v45)
                                                                                    switch ("hello" + ("\n\t" + (("foo bar" + (("foo bar" + ("" + "\n\t")) + "")) + "a"))) {
                                                                                    case "a":
                                                                                        var v58 string = "foo bar"
                                                                                        _ = v58
                                                                                    }
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                } else {
                                                                                    v4 = append(v4, fi36)
                                                                                    if ((v40 - v40) > 1) {
                                                                                        _ = len(v2)
                                                                                        v59 := GenS7{F0: 810}
                                                                                        _ = v59
                                                                                    }
                                                                                    v3 = (!(v3 && (((float32(0.0) / 1.0) / (1.5 - 3.14159)) < 3.14159)))
                                                                                    if v53 {
                                                                                        v40 -= v40
                                                                                        if (v53 || (-1 > -32768)) {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            fmt.Println(fi38)
                                                                                            var v60 int64 = 0
                                                                                            _ = v60
                                                                                        }
                                                                                    } else {
                                                                                        v53 = (2147483647 != 2147483648)
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                v61 := GenS8{F0: -792}
                                                                                                _ = v61
                                                                                            case "hello":
                                                                                                if (true || v3) {
                                                                                                    var v62 int8 = 1
                                                                                                    var v63 uint16 = uint16(v45)
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                }
                                                                                            default:
                                                                                                v64 := GenS6{F0: -34, F1: "", F2: 9223372036854775807}
                                                                                                _ = v64
                                                                                            }
                                                                                        case "":
                                                                                            if false {
                                                                                                if (v3 && true) {
                                                                                                    switch (("hello" + "a") + ("\n\t" + "a")) {
                                                                                                    case "\n\t":
                                                                                                        var v65 uint64 = 4294967295
                                                                                                        _ = v65
                                                                                                    default:
                                                                                                        var v66 string = ("foo bar" + "\n\t")
                                                                                                        _ = v66
                                                                                                    }
                                                                                                    switch (fi36 | 9223372036854775807) {
                                                                                                    case 32767:
                                                                                                        delete(v2, "foo bar")
                                                                                                    case 4294967295:
                                                                                                        v67 := GenS1{F0: fi36, F1: -1, F2: uint64(v45)}
                                                                                                        _ = v67
                                                                                                    case -2147483648:
                                                                                                        var v68 uint8 = v45
                                                                                                        _ = v68
                                                                                                    default:
                                                                                                        v21 = genFn0(0)
                                                                                                    }
                                                                                                    var v69 uint64 = 42
                                                                                                    _ = v69
                                                                                                }
                                                                                                fmt.Println(fi36, v40)
                                                                                                var v70 uint64 = 4294967295
                                                                                                _ = v70
                                                                                            }
                                                                                        case "\n\t":
                                                                                            if v53 {
                                                                                                delete(v2, "hello")
                                                                                            } else {
                                                                                                v21 -= 253
                                                                                                fmt.Println(fi36, v45, v45)
                                                                                            }
                                                                                        default:
                                                                                            if ((523 >= 42) || (-837 > -523)) {
                                                                                                if v3 {
                                                                                                    v1 = append(v1, (("\n\t" + ((("\n\t" + (("a" + "hello") + (("foo bar" + (("foo bar" + "\n\t") + "a")) + "foo bar"))) + "foo bar") + (("" + ("foo bar" + "\n\t")) + ""))) + "a"))
                                                                                                    var v71 int16 = 1000
                                                                                                    fmt.Println(v3, fi38, v53)
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    var v72 int8 = -1
                                                                                                    var v73 int8 = 0
                                                                                                    _ = v72
                                                                                                    _ = v73
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v74 uint64 = 4294967295
                                                                                switch (fi36 % ((fi38) | 1)) {
                                                                                case 4294967295:
                                                                                    if (0.0 == -2.25) {
                                                                                        var v75 float32 = (3.14159 / float32(0.1))
                                                                                        var v76 string = ("hello" + "")
                                                                                        if (0 < 2147483648) {
                                                                                            var v77 uint8 = v45
                                                                                            v78 := GenS9{F0: (v45 | (v77 >> ((1) & 63))), F1: (-(v21)), F2: v76}
                                                                                            if (v45 != v77) {
                                                                                                v40 = (v40 % (((v40 / ((v40) | 1))) | 1))
                                                                                            }
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        }
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    } else {
                                                                                        if false {
                                                                                            var v79 uint32 = 765
                                                                                            var v80 uint32 = v79
                                                                                            v45 -= v45
                                                                                            v74 -= v74
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        }
                                                                                        var v81 int32 = 2147483647
                                                                                        switch 4294967295 {
                                                                                        case 531:
                                                                                            v21 += 32767
                                                                                        default:
                                                                                            v81 = v81
                                                                                            var v82 int32 = (v81 + 1)
                                                                                            _ = v82
                                                                                        }
                                                                                        _ = v81
                                                                                    }
                                                                                    var v83 bool = true
                                                                                    _ = v83
                                                                                case -2147483648:
                                                                                    var v84 uint8 = v45
                                                                                    _ = v84
                                                                                }
                                                                                _ = v74
                                                                            }
                                                                            _ = v53
                                                                        }
                                                                        v21 -= v21
                                                                        _ = len(v2)
                                                                        v21 += genFn0(-1)
                                                                        if v3 {
                                                                            var v85 string = "foo bar"
                                                                            v21 = -461
                                                                            v2[v85] = 1
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi36, v3, v45)
                                                                        _ = len(v2)
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v86 uint16 = 0
                                                                            _ = v86
                                                                        case "foo bar":
                                                                            v2["foo bar"] = (fi38 ^ 42)
                                                                        }
                                                                    }
                                                                    if v3 {
                                                                        v87 := v2[(("foo bar" + "") + "\n\t")]
                                                                        _ = v87
                                                                    }
                                                                    v88 := GenS10{F0: v3, F1: 48, F2: (1 == -9223372036854775808)}
                                                                    if (0.1 <= (3.14159 * 0.1)) {
                                                                        var v89 int32 = -448
                                                                        v40 = 65535
                                                                        fmt.Println(v45, v89)
                                                                        v21 -= (-(v21))
                                                                        if ((fi38 << ((v40) & 63)) <= (fi36 & -2147483648)) {
                                                                            var v90 int64 = -783
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                v90 -= (v90 % (((v90 << ((42) & 63))) | 1))
                                                                                if (!v3) {
                                                                                    if (608 > 2) {
                                                                                        v21 = v21
                                                                                        if v3 {
                                                                                            if (!(42 < 0)) {
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            } else {
                                                                                                if (v45 < 2) {
                                                                                                    v40 += (v40 % (((v40 & v40)) | 1))
                                                                                                    v91 := GenS8{F0: v21}
                                                                                                    fmt.Println(v40)
                                                                                                    _ = v91
                                                                                                }
                                                                                                if v3 {
                                                                                                    if v3 {
                                                                                                        var v92 bool = v3
                                                                                                        _ = v92
                                                                                                    } else {
                                                                                                        if v3 {
                                                                                                            _ = len(v4)
                                                                                                        }
                                                                                                        var v93 uint8 = v45
                                                                                                        var v94 bool = (-2.25 >= ((((-2.25 * (-2.25 - ((float64((1e10 * 1e10)) * (float64((float32(3.14159) / 1.0)) - 1.5)) - 1.5))) / 0.1) - float64((3.14159 * 0.1))) / (-2.25 + 3.14159)))
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    var v95 int = (fi36 | (fi36 % ((fi36) | 1)))
                                                                                                    _ = v95
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if v3 {
                                                                                                v21 = (v21 % ((v21) | 1))
                                                                                                v21 += v21
                                                                                                v3 = (v45 <= v45)
                                                                                            }
                                                                                            var v96 int16 = 1000
                                                                                            switch "foo bar" {
                                                                                            case "a":
                                                                                                v96 -= v21
                                                                                                v97 := GenS2{F0: 0, F1: v45, F2: genFn0(127)}
                                                                                                fmt.Println(v3, fi38)
                                                                                                _ = v97
                                                                                            default:
                                                                                                if ((v40 * v40) < 4294967295) {
                                                                                                    switch fi38 {
                                                                                                    case 120:
                                                                                                        v2[((("\n\t" + (((("foo bar" + "hello") + "") + "") + ("a" + "a"))) + ("hello" + "")) + "")] = fi38
                                                                                                    }
                                                                                                } else {
                                                                                                    var v98 uint32 = 2147483648
                                                                                                    if ((^(v40)) != v40) {
                                                                                                        if (-75 < 127) {
                                                                                                            if (v89 > v89) {
                                                                                                                if false {
                                                                                                                    v99 := GenS2{F0: 1, F1: (^(v45)), F2: (v96 | (v21 & v21))}
                                                                                                                    var v100 uint8 = v45
                                                                                                                    _ = v100
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                                if v3 {
                                                                                                                    switch fi38 {
                                                                                                                    case 32768:
                                                                                                                        v101 := GenS8{F0: (v96 + (v21 & genFn0(33)))}
                                                                                                                        _ = v101
                                                                                                                    case 752:
                                                                                                                        var v102 uint32 = (v98 / (((v98 - (^(v98)))) | 1))
                                                                                                                        _ = v102
                                                                                                                    case 128:
                                                                                                                        switch (^(fi36)) {
                                                                                                                        case 9223372036854775807:
                                                                                                                            var v103 uint32 = v98
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch (fi38 ^ fi36) {
                                                                                                                    case -2147483648:
                                                                                                                        fmt.Println(fi36, v89, v3)
                                                                                                                        fmt.Println(v90, v40)
                                                                                                                        var v104 int16 = 371
                                                                                                                        if ("\n\t" != "foo bar") {
                                                                                                                            v104 = 1000
                                                                                                                        }
                                                                                                                        v105 := GenS11{F0: 42}
                                                                                                                        var v106 int = 127
                                                                                                                        _ = v104
                                                                                                                        _ = v105
                                                                                                                        _ = v106
                                                                                                                    case -1:
                                                                                                                        if ((v89 | v89) != v89) {
                                                                                                                            switch ("" + "foo bar") {
                                                                                                                            case "hello":
                                                                                                                                _ = len(v2)
                                                                                                                            case "foo bar":
                                                                                                                                switch (("" + "a") + "a") {
                                                                                                                                case "hello":
                                                                                                                                    var v107 int16 = v21
                                                                                                                                    _ = v107
                                                                                                                                case "\n\t":
                                                                                                                                    var v108 int16 = genFn0(127)
                                                                                                                                    _ = v108
                                                                                                                                case "foo bar":
                                                                                                                                    v109 := GenS3{F0: 0}
                                                                                                                                    _ = v109
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                v110 := v2[("foo bar" + ("hello" + "\n\t"))]
                                                                                                                                _ = v110
                                                                                                                            }
                                                                                                                            v1 = append(v1, "a")
                                                                                                                        } else {
                                                                                                                            if v3 {
                                                                                                                                v96 -= (v96 % ((genFn0(100)) | 1))
                                                                                                                                switch "foo bar" {
                                                                                                                                case "foo bar":
                                                                                                                                    v89 += v89
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if v3 {
                                                                                                                                    var v111 uint16 = 42
                                                                                                                                    v96 = (v96 % (((v96 % ((int16(v111)) | 1))) | 1))
                                                                                                                                    _ = v111
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v112 float64 = 0.0
                                                                                                                            if len(v1) >= 2 {
                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                            }
                                                                                                                            v113 := v2[(("hello" + "\n\t") + ("a" + "a"))]
                                                                                                                            _ = v112
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                    case 0:
                                                                                                                        var v114 uint16 = 115
                                                                                                                        _ = v114
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch (("\n\t" + "a") + ("" + "a")) {
                                                                                                                    case "\n\t":
                                                                                                                        v40 -= v40
                                                                                                                    case "foo bar":
                                                                                                                        var v115 string = (("foo bar" + "") + "a")
                                                                                                                        _ = v115
                                                                                                                    case "":
                                                                                                                        switch "a" {
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(v98, v98, v45)
                                                                                                                        case "hello":
                                                                                                                            if (221 == (v45 << ((v40) & 63))) {
                                                                                                                                if (v3 && (-1 != (v89 | v89))) {
                                                                                                                                    _ = len(v1)
                                                                                                                                    var v116 float64 = (1.5 / (-2.25 / 1.5))
                                                                                                                                    _ = v116
                                                                                                                                } else {
                                                                                                                                    v3 = (!(42 >= v40))
                                                                                                                                }
                                                                                                                                var v117 bool = v3
                                                                                                                                v118 := GenS12{F0: int64(fi38), F1: (fi36 % (((fi38 ^ int(v45))) | 1))}
                                                                                                                                switch (fi38 ^ fi36) {
                                                                                                                                case -9223372036854775808:
                                                                                                                                    var v119 int64 = v90
                                                                                                                                    _ = v119
                                                                                                                                default:
                                                                                                                                    fmt.Println(v96)
                                                                                                                                }
                                                                                                                                v4 = append(v4, (fi36 % (((fi36 | -9223372036854775808)) | 1)))
                                                                                                                                var v120 int16 = genFn0(42)
                                                                                                                                v45 = (v45 + v45)
                                                                                                                                _ = v117
                                                                                                                                _ = v118
                                                                                                                                _ = v120
                                                                                                                            } else {
                                                                                                                                if len(v4) >= 2 {
                                                                                                                                    v4 = v4[1:len(v4)]
                                                                                                                                }
                                                                                                                                v121 := GenS13{F0: 1, F1: -128}
                                                                                                                                var v122 uint8 = (v45 - (v45 << ((553) & 63)))
                                                                                                                                _ = v121
                                                                                                                                _ = v122
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v89 = (-(v89))
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = int(v98)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v123 float32 = ((-2.25 * (((1.5 - 3.14159) * (float32(float64(0.0)) * 3.14159)) * 0.0)) * 1e10)
                                                                                                                    fmt.Println(v89)
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                                v98 = (v98 * 571)
                                                                                                                v90 += (v90 % (((-(v90))) | 1))
                                                                                                                if (true || true) {
                                                                                                                    var v124 bool = true
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                    _ = len(v2)
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                                fmt.Println(v89)
                                                                                                                var v125 int = (fi38 / ((-1) | 1))
                                                                                                                if true {
                                                                                                                    fmt.Println(v40)
                                                                                                                    var v126 int64 = (v90 >> (((v40 >> ((v40) & 63))) & 63))
                                                                                                                    fmt.Println(v125)
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                                v127 := GenS9{F0: (v45 & 133), F1: 0, F2: "foo bar"}
                                                                                                                _ = v125
                                                                                                                _ = v127
                                                                                                            }
                                                                                                            fmt.Println(v96)
                                                                                                        } else {
                                                                                                            var v128 int32 = (v89 | (^(v89)))
                                                                                                            switch "hello" {
                                                                                                            case "hello":
                                                                                                                var v129 uint16 = 0
                                                                                                                _ = v129
                                                                                                            case "\n\t":
                                                                                                                var v130 uint = v40
                                                                                                                _ = v130
                                                                                                            }
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v98
                                                                                                }
                                                                                            }
                                                                                            _ = v96
                                                                                        }
                                                                                    } else {
                                                                                        v40 = (v40 % ((144) | 1))
                                                                                        if true {
                                                                                            v131 := GenS5{F0: (v90 - (-(v90))), F1: v40}
                                                                                            v132 := GenS13{F0: -1, F1: int8(v89)}
                                                                                            var v133 int64 = -151
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = int(v45)
                                                                                            }
                                                                                            _ = v131
                                                                                            _ = v132
                                                                                            _ = v133
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                            v21 = (-(v21))
                                                                                            if v3 {
                                                                                                var v134 int64 = -1
                                                                                                if v3 {
                                                                                                    v90 += 0
                                                                                                }
                                                                                                var v135 uint = 18446744073709551615
                                                                                                v136 := GenS14{F0: (^(v89)), F1: 1}
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                                _ = v136
                                                                                            } else {
                                                                                                v137 := v2[("a" + "a")]
                                                                                                v40 += (v40 - (v40 >> (((^(v40))) & 63)))
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    v21 = genFn0(42)
                                                                                                case "":
                                                                                                    var v138 int32 = 2147483647
                                                                                                    _ = v138
                                                                                                case "hello":
                                                                                                    v21 -= (v21 & v21)
                                                                                                default:
                                                                                                    v90 = v90
                                                                                                }
                                                                                                _ = len(v2)
                                                                                                _ = v137
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    v139 := GenS6{F0: 42, F1: ("\n\t" + "foo bar"), F2: (^(v90))}
                                                                                    var v140 int64 = 0
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                } else {
                                                                                    if true {
                                                                                        v21 = genFn0(127)
                                                                                    }
                                                                                    v141 := GenS15{F0: "\n\t", F1: 185, F2: (v89 * (v89 / ((v89) | 1)))}
                                                                                    _ = v141
                                                                                }
                                                                                var v142 int64 = (v90 % ((v90) | 1))
                                                                                _ = v142
                                                                            }
                                                                            fmt.Println(v21, v3, v21)
                                                                            _ = v90
                                                                        }
                                                                        _ = v89
                                                                    } else {
                                                                        var v143 int32 = -102
                                                                        _ = v143
                                                                    }
                                                                    switch fi38 {
                                                                    case 9223372036854775807:
                                                                        if v3 {
                                                                            v3 = (1 > 1)
                                                                            v40 -= v40
                                                                        } else {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ""
                                                                            }
                                                                        }
                                                                    default:
                                                                        if v3 {
                                                                            switch ("\n\t" + ("\n\t" + ("" + "foo bar"))) {
                                                                            case "\n\t":
                                                                                v40 -= v40
                                                                            case "a":
                                                                                v2["hello"] = (fi36 - (fi36 & (fi38 & (fi38 - (fi38 % ((int(v40)) | 1))))))
                                                                            }
                                                                        } else {
                                                                            v40 -= 255
                                                                        }
                                                                    }
                                                                    _ = v88
                                                                } else {
                                                                    var v144 bool = (2 >= 593)
                                                                    var v145 float32 = 1.5
                                                                    _ = v144
                                                                    _ = v145
                                                                }
                                                            }
                                                            if (-2147483648 >= 437) {
                                                                var v146 uint8 = v45
                                                                v4 = append(v4, fi38)
                                                                v146 = 0
                                                                _ = v146
                                                            } else {
                                                                var v147 uint8 = v45
                                                                var v148 int8 = 42
                                                                v149 := GenS16{F0: (v3 || v3), F1: 696}
                                                                _ = v147
                                                                _ = v148
                                                                _ = v149
                                                            }
                                                            delete(v2, (("foo bar" + ("" + (((((("hello" + "a") + ("\n\t" + "hello")) + "hello") + "a") + "a") + "a"))) + ("" + "")))
                                                            var v150 float32 = float32((0.1 + 1e10))
                                                            var v151 int32 = -279
                                                            v3 = (v3 || ((((1e10 * (1.5 - 1e10)) * 0.0) - float64(((1e10 / float32(3.14159)) + 1e10))) >= float64((v150 + 3.14159))))
                                                            if ((!v3) || ((-(fi38)) > 834)) {
                                                                if (("" >= "foo bar") && (!(v40 >= 64))) {
                                                                    switch (fi38 << (((^(v40))) & 63)) {
                                                                    case -819:
                                                                        v150 += (v150 + ((-2.25 * v150) * 3.14159))
                                                                    case 1:
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    case -9223372036854775808:
                                                                        switch ("\n\t" + ("a" + "\n\t")) {
                                                                        case "\n\t":
                                                                            v150 += float32(1e10)
                                                                        case "a":
                                                                            switch int(v40) {
                                                                            case -778:
                                                                                fmt.Println(v40, v40)
                                                                            case 127:
                                                                                if true {
                                                                                    _ = len(v1)
                                                                                    v152 := GenS13{F0: 83, F1: -1}
                                                                                    _ = v152
                                                                                }
                                                                            case 2147483647:
                                                                                switch (fi38 / (((fi36 ^ fi36)) | 1)) {
                                                                                case 4294967295:
                                                                                    v1 = append(v1, ("" + "\n\t"))
                                                                                    if ((("" + "hello") + "a") <= "a") {
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    } else {
                                                                                        if (v40 > v40) {
                                                                                            _ = len(v2)
                                                                                            fmt.Println(v150)
                                                                                        } else {
                                                                                            var v153 uint32 = 826
                                                                                            fmt.Println(v21)
                                                                                            v154 := GenS15{F0: "\n\t", F1: 182, F2: v151}
                                                                                            _ = v153
                                                                                            _ = v154
                                                                                        }
                                                                                        if true {
                                                                                            v155 := GenS7{F0: 507}
                                                                                            _ = v155
                                                                                        }
                                                                                        if v3 {
                                                                                            v156 := GenS11{F0: 127}
                                                                                            v2["\n\t"] = -859
                                                                                            var v157 float32 = v150
                                                                                            _ = v156
                                                                                            _ = v157
                                                                                        } else {
                                                                                            v21 = (v21 | v21)
                                                                                        }
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (("" + "hello") + "hello")
                                                                                    }
                                                                                case -2147483648:
                                                                                    if (((((((v150 + v150) * (v150 + (1e10 * (v150 / ((v150 - float32(0.0)) - float32(0.1)))))) / (1.5 - v150)) + float32(0.0)) <= (-2.25 * v150)) || ((0 <= 42) || true)) || (127 > (v45 << (((^(v40))) & 63)))) {
                                                                                        var v158 int32 = (v151 | (v151 | v151))
                                                                                        fmt.Println(v45)
                                                                                        _ = v158
                                                                                    } else {
                                                                                        switch -808 {
                                                                                        case -155:
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                if (0 > uint64(v45)) {
                                                                                                    if (1 > v45) {
                                                                                                        var v159 uint32 = 296
                                                                                                        _ = v159
                                                                                                    }
                                                                                                    switch (fi36 - -9223372036854775808) {
                                                                                                    case -2147483648:
                                                                                                        var v160 uint16 = 255
                                                                                                        _ = v160
                                                                                                    case 4294967295:
                                                                                                        var v161 int8 = int8(v151)
                                                                                                        _ = v161
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v21)
                                                                                                }
                                                                                            case "hello":
                                                                                                switch (("\n\t" + "\n\t") + ("\n\t" + "foo bar")) {
                                                                                                case "hello":
                                                                                                    v1 = append(v1, ("a" + ""))
                                                                                                case "\n\t":
                                                                                                    v162 := GenS17{F0: "a"}
                                                                                                    _ = v162
                                                                                                case "a":
                                                                                                    switch fi36 {
                                                                                                    case 128:
                                                                                                        if (!v3) {
                                                                                                            if ((!true) && v3) {
                                                                                                                v3 = ((0 != 42) && v3)
                                                                                                                fmt.Println(v21, fi38, v151)
                                                                                                            } else {
                                                                                                                v150 = (0.0 * v150)
                                                                                                                switch fi38 {
                                                                                                                case 0:
                                                                                                                    v3 = v3
                                                                                                                case 42:
                                                                                                                    v45 += v45
                                                                                                                }
                                                                                                            }
                                                                                                            v21 = (v21 << (((^(v40))) & 63))
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = fi38
                                                                                                            }
                                                                                                            v150 = v150
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                    fmt.Println(v3)
                                                                                                }
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi36, fi38)
                                                                                            default:
                                                                                                v151 += v151
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v21 += -361
                                                                                }
                                                                            default:
                                                                                fmt.Println(v40, v21)
                                                                            }
                                                                        case "hello":
                                                                            if false {
                                                                                if false {
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v163 uint64 = 42
                                                                                        _ = v163
                                                                                    case "hello":
                                                                                        var v164 uint64 = 42
                                                                                        _ = v164
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v165 bool = true
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    fmt.Println(v40, v151, fi36)
                                                                                    var v166 float32 = v150
                                                                                    v167 := GenS18{F0: 42, F1: ((v3 && (42 > -1)) || ((v21 <= v21) && false)), F2: v151}
                                                                                    v45 = 42
                                                                                    v150 = 0.1
                                                                                    if (9223372036854775807 <= fi38) {
                                                                                        fmt.Println(v21, v21, v150)
                                                                                        v150 = -2.25
                                                                                    } else {
                                                                                        if v165 {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "\n\t"
                                                                                            }
                                                                                        }
                                                                                        v21 = (v21 >> ((uint(v45)) & 63))
                                                                                    }
                                                                                    if (0.0 < 3.14159) {
                                                                                        v168 := GenS19{F0: 4294967295, F1: 9223372036854775807, F2: -584}
                                                                                        fmt.Println(v165, v151, v166)
                                                                                        _ = v168
                                                                                    }
                                                                                    var v169 int16 = int16(v151)
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                    _ = v169
                                                                                default:
                                                                                    v170 := GenS20{F0: 363, F1: v40, F2: 9223372036854775807}
                                                                                    _ = v170
                                                                                }
                                                                                _ = v165
                                                                            }
                                                                        }
                                                                    default:
                                                                        v40 -= 255
                                                                    }
                                                                    delete(v2, "\n\t")
                                                                } else {
                                                                    var v171 uint8 = (v45 << ((65535) & 63))
                                                                    switch ("foo bar" + (("\n\t" + "") + (("\n\t" + "foo bar") + "hello"))) {
                                                                    case "foo bar":
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v172 uint8 = 127
                                                                            _ = v172
                                                                        case "a":
                                                                            v150 = v150
                                                                        default:
                                                                            fmt.Println(v3, v151)
                                                                        }
                                                                    default:
                                                                        if (632 != 822) {
                                                                            v173 := GenS5{F0: 42, F1: v40}
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "foo bar":
                                                                                if v3 {
                                                                                    v1 = append(v1, "hello")
                                                                                }
                                                                            case "hello":
                                                                                var v174 uint64 = 42
                                                                                _ = v174
                                                                            default:
                                                                                v21 -= (v21 / ((int16(v171)) | 1))
                                                                            }
                                                                            _ = v173
                                                                        } else {
                                                                            var v175 int8 = 1
                                                                            if v3 {
                                                                                var v176 float32 = float32((-2.25 - ((1.5 * 0.0) - -2.25)))
                                                                                _ = v176
                                                                            }
                                                                            v4 = append(v4, fi36)
                                                                            if (42 == 491) {
                                                                                var v177 uint32 = 781
                                                                                _ = v177
                                                                            } else {
                                                                                v175 = v175
                                                                                var v178 float32 = 3.14159
                                                                                if (!v3) {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 446
                                                                                    }
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v179 := GenS21{F0: 1, F1: (^(fi38)), F2: 32767}
                                                                                        _ = v179
                                                                                    }
                                                                                    v2["a"] = 712
                                                                                } else {
                                                                                    var v180 int32 = v151
                                                                                    var v181 uint8 = (v45 % ((128) | 1))
                                                                                    var v182 uint32 = 944
                                                                                    _ = v180
                                                                                    _ = v181
                                                                                    _ = v182
                                                                                }
                                                                                _ = v178
                                                                            }
                                                                            _ = v175
                                                                        }
                                                                    }
                                                                    v183 := GenS22{F0: 2147483647, F1: 42, F2: (true || false)}
                                                                    if v3 {
                                                                        if v3 {
                                                                            switch fi36 {
                                                                            case 0:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            case -9223372036854775808:
                                                                                v184 := GenS9{F0: uint8(v40), F1: 32767, F2: ("\n\t" + "")}
                                                                                _ = v184
                                                                            case 128:
                                                                                var v185 uint8 = 127
                                                                                _ = v185
                                                                            default:
                                                                                if (v3 && true) {
                                                                                    if (412 >= uint16(v171)) {
                                                                                        v40 = (v40 & v40)
                                                                                        v45 = (v171 & 127)
                                                                                    } else {
                                                                                        if true {
                                                                                            v186 := GenS23{F0: 42}
                                                                                            var v187 string = ""
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (fi36 * fi36)
                                                                                            }
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                        }
                                                                                    }
                                                                                    v21 += 42
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        _ = len(v4)
                                                                                    case "hello":
                                                                                        switch 4294967295 {
                                                                                        case -845:
                                                                                            if v3 {
                                                                                                var v188 float32 = v150
                                                                                                _ = v188
                                                                                            }
                                                                                        case -140:
                                                                                            v4 = append(v4, 667)
                                                                                        case 4294967295:
                                                                                            v150 = 3.14159
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v21, v171, v171)
                                                                                    v171 = v45
                                                                                }
                                                                            }
                                                                            _ = len(v4)
                                                                            v189 := GenS24{F0: -19}
                                                                            _ = v189
                                                                        }
                                                                        if false {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 128
                                                                            }
                                                                            var v190 uint8 = v45
                                                                            _ = v190
                                                                        } else {
                                                                            v2["hello"] = fi36
                                                                            v171 = v45
                                                                            var v191 int64 = -671
                                                                            v1 = append(v1, "")
                                                                            _ = v191
                                                                        }
                                                                    }
                                                                    _ = v171
                                                                    _ = v183
                                                                }
                                                                fmt.Println(fi38, v21)
                                                            } else {
                                                                if true {
                                                                    v3 = true
                                                                    v3 = v3
                                                                } else {
                                                                    if v3 {
                                                                        if (!false) {
                                                                            if v3 {
                                                                                fmt.Println(v150, v3, fi38)
                                                                                var v192 int8 = 42
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v40 = v40
                                                                                _ = v192
                                                                            }
                                                                            var v193 int8 = -128
                                                                            v194 := GenS25{F0: 235, F1: "a"}
                                                                            _ = v193
                                                                            _ = v194
                                                                        } else {
                                                                            if (1 < (v21 ^ (v21 % ((v21) | 1)))) {
                                                                                _ = len(v2)
                                                                                fmt.Println(v40, fi36, v40)
                                                                                if (v3 && v3) {
                                                                                    if v3 {
                                                                                        if v3 {
                                                                                            v195 := GenS26{F0: (v45 - v45), F1: 42, F2: v3}
                                                                                            if v3 {
                                                                                                var v196 uint8 = v45
                                                                                                switch fi38 {
                                                                                                case -570:
                                                                                                    if v3 {
                                                                                                        v150 = v150
                                                                                                        v3 = (v3 || v3)
                                                                                                    } else {
                                                                                                        fmt.Println(v3, fi38, v21)
                                                                                                    }
                                                                                                default:
                                                                                                    if true {
                                                                                                        var v197 uint64 = 42
                                                                                                        _ = v197
                                                                                                    }
                                                                                                }
                                                                                                switch (("a" + ("hello" + "hello")) + "foo bar") {
                                                                                                case "foo bar":
                                                                                                    switch "\n\t" {
                                                                                                    case "":
                                                                                                        v198 := GenS20{F0: 4294967295, F1: v40, F2: fi36}
                                                                                                        var v199 int = fi38
                                                                                                        var v200 float64 = 0.1
                                                                                                        if v3 {
                                                                                                            var v201 int32 = v151
                                                                                                            v1 = append(v1, "foo bar")
                                                                                                            _ = v201
                                                                                                        }
                                                                                                        _ = v198
                                                                                                        _ = v199
                                                                                                        _ = v200
                                                                                                    }
                                                                                                case "a":
                                                                                                    v45 += 128
                                                                                                case "":
                                                                                                    if ((^(v21)) == 0) {
                                                                                                        v45 -= v196
                                                                                                        fmt.Println(v3, v45)
                                                                                                    }
                                                                                                default:
                                                                                                    v40 -= v40
                                                                                                }
                                                                                                _ = v196
                                                                                            }
                                                                                            fmt.Println(v150, v21, v21)
                                                                                            _ = v195
                                                                                        } else {
                                                                                            v202 := GenS27{F0: 0}
                                                                                            var v203 uint16 = 255
                                                                                            _ = v202
                                                                                            _ = v203
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            v150 -= (1e10 / v150)
                                                                                        case "foo bar":
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                v204 := GenS23{F0: 147}
                                                                                                _ = v204
                                                                                            case "":
                                                                                                v150 -= v150
                                                                                                v205 := GenS26{F0: 1, F1: 1, F2: v3}
                                                                                                delete(v2, ("\n\t" + ""))
                                                                                                _ = v205
                                                                                            default:
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    switch ("" + ("" + "foo bar")) {
                                                                                                    case "\n\t":
                                                                                                        v40 = (v40 + 0)
                                                                                                    case "":
                                                                                                        if v3 {
                                                                                                            _ = len(v1)
                                                                                                            fmt.Println(v150)
                                                                                                        }
                                                                                                    }
                                                                                                case "a":
                                                                                                    var v206 uint = (v40 >> (((v40 >> ((42) & 63))) & 63))
                                                                                                    _ = v206
                                                                                                case "hello":
                                                                                                    var v207 float64 = 0.0
                                                                                                    _ = v207
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            v208 := GenS1{F0: 1, F1: 13, F2: 0}
                                                                                            _ = v208
                                                                                        }
                                                                                        switch int(v21) {
                                                                                        case 1:
                                                                                            var v209 int8 = 100
                                                                                            _ = v209
                                                                                        case 32768:
                                                                                            v210 := GenS17{F0: ("hello" + "a")}
                                                                                            _ = v210
                                                                                        case -855:
                                                                                            v211 := v2["\n\t"]
                                                                                            _ = v211
                                                                                        }
                                                                                        switch (fi36 | fi38) {
                                                                                        case 42:
                                                                                            v40 = (v40 ^ uint(v45))
                                                                                        case -2147483648:
                                                                                            switch "a" {
                                                                                            case "hello":
                                                                                                v3 = (uint64(v45) >= 543)
                                                                                            }
                                                                                        case 347:
                                                                                            var v212 int32 = (v151 - (v151 ^ (v151 + v151)))
                                                                                            _ = v212
                                                                                        default:
                                                                                            v213 := GenS28{F0: "", F1: (v45 * v45), F2: int16(fi38)}
                                                                                            _ = v213
                                                                                        }
                                                                                        v3 = v3
                                                                                    }
                                                                                    v150 = v150
                                                                                    v2["\n\t"] = (fi36 >> ((v40) & 63))
                                                                                }
                                                                            }
                                                                            _ = len(v2)
                                                                            var v214 int16 = v21
                                                                            _ = v214
                                                                        }
                                                                        v40 = v40
                                                                    }
                                                                }
                                                            }
                                                            _ = v150
                                                            _ = v151
                                                            _ = v49
                                                        }
                                                    }
                                                }
                                                _ = v45
                                            }
                                            _ = v40
                                        } else {
                                            if (!(!(2147483647 != 0))) {
                                                switch "" {
                                                case "\n\t":
                                                    v3 = ((32767 > 167) || (v3 || (!v3)))
                                                case "a":
                                                    v3 = true
                                                case "hello":
                                                    v215 := GenS15{F0: "", F1: 512, F2: -1}
                                                    _ = v215
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                }
                                                if (!true) {
                                                    fmt.Println(v3, v21)
                                                    fmt.Println(fi36, v3)
                                                    if len(v1) > 0 {
                                                        v1[0] = "a"
                                                    }
                                                }
                                            } else {
                                                var v216 uint32 = 0
                                                v217 := GenS28{F0: "foo bar", F1: 0, F2: v21}
                                                _ = v216
                                                _ = v217
                                            }
                                        }
                                        v218 := v2["foo bar"]
                                        v2["\n\t"] = v218
                                        _ = len(v4)
                                        _ = v218
                                    } else {
                                        v219 := GenS29{F0: -100, F1: 215}
                                        _ = v219
                                    }
                                    var v220 uint8 = 14
                                    var v221 uint16 = 2
                                    v3 = v3
                                    _ = v220
                                    _ = v221
                                } else {
                                    if false {
                                        fmt.Println(v3, fi38)
                                        var v222 string = ""
                                        _ = v222
                                    } else {
                                        delete(v2, "foo bar")
                                    }
                                    var v223 string = "a"
                                    _ = v223
                                }
                                var v224 int32 = 1
                                var v225 uint8 = 185
                                v226 := GenS30{F0: 639, F1: 1, F2: ("hello" + "\n\t")}
                                var v227 int64 = 850
                                _ = v224
                                _ = v225
                                _ = v226
                                _ = v227
                            }
                        }
                    default:
                        var v228 bool = ((("" + "hello") + "\n\t") != "foo bar")
                        v229 := GenS31{F0: (fi36 * fi36), F1: true, F2: 876}
                        var v230 int16 = v21
                        _ = v228
                        _ = v229
                        _ = v230
                    }
                    var v231 float32 = float32(1.5)
                    _ = fi38
                    _ = v231
                }
            }
            v21 += 32767
            _ = fi36
        }
    default:
        if (!(42 == 2)) {
            v21 = v21
            var v232 float32 = 0.0
            _ = v232
        } else {
            for fi233 := 0; fi233 < 19; fi233++ {
                v3 = v3
                if (40 < 0) {
                    v4 = append(v4, (-(fi233)))
                    switch ("\n\t" + "a") {
                    case "a":
                        v234 := GenS2{F0: 289, F1: 128, F2: v21}
                        _ = v234
                    case "":
                        var v235 int16 = v21
                        _ = v235
                    case "\n\t":
                        v3 = true
                    default:
                        v3 = false
                    }
                    var v236 uint = 0
                    _ = v236
                }
                var v237 int64 = -719
                _ = fi233
                _ = v237
            }
            fmt.Println(v3, v21)
        }
    }
    var v238 float32 = float32((0.0 + float64(3.14159)))
    for fi239 := 0; fi239 < 18; fi239++ {
        v21 = 42
        _ = fi239
    }
    fmt.Println(v3)
    switch ("foo bar" + "") {
    case "":
        v240 := v2[("a" + "hello")]
        _ = v240
    default:
        v241 := GenS25{F0: 18446744073709551615, F1: "\n\t"}
        _ = v241
    }
    var v242 int64 = -455
    if (2 < 1) {
        v243 := GenS31{F0: -131, F1: (-2147483648 < 127), F2: 42}
        v244 := GenS32{F0: 255}
        v245 := GenS33{F0: "foo bar", F1: -438}
        _ = v243
        _ = v244
        _ = v245
    } else {
        fmt.Println(v242)
    }
    var v246 uint32 = 42
    if (((3.14159 * ((0.0 - ((0.1 - 1e10) + 1.5)) - (-2.25 + 0.1))) - (3.14159 - float64(v238))) < -2.25) {
        for fi247 := 0; fi247 < 12; fi247++ {
            if (true && v3) {
                var v248 uint64 = 0
                var v249 float64 = ((((0.1 + 1e10) + float64(1e10)) + -2.25) * 3.14159)
                if (v3 && ((v242 >= v242) || v3)) {
                    fmt.Println(v238)
                    for fi250 := 0; fi250 < 14; fi250++ {
                        v251 := GenS34{F0: (v246 * (^(v246)))}
                        v21 -= -943
                        if v3 {
                            if (!v3) {
                                var v252 int16 = (v21 | genFn0(-128))
                                if (!false) {
                                    fmt.Println(v248, v3, v252)
                                    fmt.Println(fi247, v3, v249)
                                    v253 := GenS35{F0: 65535, F1: (v242 <= v242)}
                                    _ = v253
                                }
                                _ = v252
                            }
                            var v254 float64 = -2.25
                            if (!(v3 && v3)) {
                                v254 -= v249
                                v248 -= (v248 | 18446744073709551615)
                                v3 = v3
                            }
                            var v255 int64 = (v242 % ((v242) | 1))
                            if ((v248 * 0) > 787) {
                                var v256 int16 = (v21 + int16(v255))
                                v3 = ((v248 / (((v248 * v248)) | 1)) != (v248 >> ((244) & 63)))
                                _ = v256
                            }
                            v254 -= 1e10
                            switch (("" + "hello") + "") {
                            case "a":
                                _ = len(v2)
                            case "\n\t":
                                switch "\n\t" {
                                case "a":
                                    v257 := GenS36{F0: (v21 * v21), F1: (v21 & genFn0(110)), F2: 943}
                                    _ = v257
                                case "\n\t":
                                    var v258 bool = (int8(v246) > 42)
                                    _ = v258
                                }
                            }
                            _ = v254
                            _ = v255
                        } else {
                            var v259 uint64 = (v248 + (v248 * (v248 ^ v248)))
                            v260 := GenS27{F0: 42}
                            _ = v259
                            _ = v260
                        }
                        _ = fi250
                        _ = v251
                    }
                    v261 := GenS11{F0: 114}
                    _ = v261
                }
                _ = v248
                _ = v249
            } else {
                v262 := GenS28{F0: "", F1: 128, F2: v21}
                var v263 uint64 = 42
                _ = v262
                _ = v263
            }
            _ = fi247
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if ((v242 % ((v242) | 1)) >= (v242 << ((42) & 63))) {
            for fi264 := 0; fi264 < 12; fi264++ {
                var v265 uint32 = v246
                _ = fi264
                _ = v265
            }
        } else {
            v266 := v2[("\n\t" + "foo bar")]
            if (v266 > -475) {
                delete(v2, "hello")
                if (v3 && (1 != 1)) {
                    var v267 bool = (42 > -1)
                    v238 = 3.14159
                    for fi268 := 0; fi268 < 20; fi268++ {
                        v267 = (!v3)
                        _ = fi268
                    }
                    _ = v267
                } else {
                    v21 = (v21 >> ((4294967295) & 63))
                }
                v269 := GenS37{F0: "\n\t"}
                _ = v269
            } else {
                var v270 uint = 820
                _ = v270
            }
            _ = v266
        }
        switch 42 {
        case 42:
            if true {
                v246 = v246
                _ = len(v2)
                for fi271 := 0; fi271 < 5; fi271++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi271
                }
            } else {
                for fi272 := 0; fi272 < 7; fi272++ {
                    v4 = append(v4, 0)
                    _ = fi272
                }
                v246 = 0
                fmt.Println(v21, v238, v238)
            }
        case -777:
            if (65535 == v246) {
                var v273 uint64 = 4294967295
                switch "" {
                case "foo bar":
                    for fi274 := 0; fi274 < 6; fi274++ {
                        var v275 int32 = -2147483648
                        _ = fi274
                        _ = v275
                    }
                case "hello":
                    v21 = v21
                }
                if len(v4) > 0 {
                    v4[0] = -1
                }
                _ = v273
            } else {
                switch "\n\t" {
                case "a":
                    v2["a"] = 370
                case "":
                    v4 = append(v4, 2147483647)
                }
                if (!(!((!(0.1 < (1.5 * -2.25))) || (false && false)))) {
                    v276 := GenS38{F0: 1}
                    switch 0 {
                    case 128:
                        v277 := v2["foo bar"]
                        _ = v277
                    }
                    v238 = v238
                    _ = v276
                }
            }
        }
        _ = len(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v238:", v238)
    fmt.Println("v242:", v242)
    fmt.Println("v246:", v246)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v238
    _ = v242
    _ = v246
    _ = v3
    _ = v4
}
