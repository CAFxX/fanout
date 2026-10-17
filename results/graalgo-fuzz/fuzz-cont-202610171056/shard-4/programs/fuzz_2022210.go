package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16; F1 string }
func genFn0(p0 bool, p1 uint32, p2 uint16) int16 {
    v5 := GenS1{F0: "\n\t", F1: 127}
    if p0 {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{(fi6 / (((fi6 / ((fi6) | 1))) | 1)), (fi6 ^ (fi6 * (-(fi6))))}
            _ = fi6
            _ = v7
        }
        switch 0 {
        case 32768:
            var v8 int64 = -651
            _ = v8
        case 42:
            var v9 uint16 = p2
            _ = v9
        case 322:
            var v10 int64 = -1
            _ = v10
        default:
            v11 := []int{32768, 32767, -298, 486}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{}
            v14 := GenS2{F0: 18446744073709551615, F1: 1, F2: (false && (-128 != 108))}
            _ = fi12
            _ = v13
            _ = v14
        }
        switch 0 {
        case 945:
            v15 := []int{1, 2147483647}
            _ = v15
        default:
            p0 = ((!(false || (-1 >= -2147483648))) || ("" <= ("\n\t" + ("" + "foo bar"))))
        }
        var v16 float64 = (1.5 / 1.0)
        _ = v16
    }
    var v17 uint32 = (p1 | p1)
    v17 -= (p1 - 4294967295)
    var v18 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v5
    return 32767
}
func genFn1() uint64 {
    v19 := GenS1{F0: ("\n\t" + ("hello" + "foo bar")), F1: 9223372036854775807}
    var v20 uint8 = 128
    v20 = v20
    v21 := GenS1{F0: ("" + "hello"), F1: 32767}
    var v22 int8 = -1
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return 18446744073709551615
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]bool)
    var v3 float64 = -2.25
    v4 := []int{}
    v1 = (v1 ^ genFn0(true, 921, 255))
    for fi23 := 0; fi23 < 18; fi23++ {
        if false {
            var v24 int = (fi23 + 931)
            fmt.Println(fi23, v24, v24)
            v24 -= (-(fi23))
            _ = v24
        } else {
            v25 := GenS2{F0: 528, F1: -474, F2: (((!false) && (-111 <= (fi23 << ((2) & 63)))) || (v3 < v3))}
            var v26 uint = 255
            if false {
                var v27 bool = true
                _ = v27
            }
            _ = v25
            _ = v26
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 float32 = 1.5
            _ = fi28
            _ = v29
        }
        _ = fi23
    }
    fmt.Println(v1, v3)
    v30 := GenS2{F0: 0, F1: 1, F2: (1 == 44)}
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 float32 = 1.5
        _ = fi31
        _ = v32
    }
    if (1 >= 2) {
        if (625 > 127) {
            v3 = 0.0
            for fi33 := 0; fi33 < 11; fi33++ {
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 bool = false
                    if v35 {
                        v36 := GenS2{F0: 42, F1: -891, F2: (!(fi34 < (fi33 * (fi33 - fi34))))}
                        _ = v36
                    }
                    switch fi34 {
                    case -957:
                        switch fi34 {
                        case 1:
                            v37 := GenS1{F0: "a", F1: (-(fi33))}
                            _ = v37
                        case 0:
                            var v38 int32 = 1
                            _ = v38
                        default:
                            if (758 >= 42) {
                                v1 = (v1 << ((255) & 63))
                            } else {
                                var v39 uint16 = 0
                                var v40 string = ""
                                var v41 string = (v40 + ("" + v40))
                                var v42 int8 = 42
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        }
                    default:
                        if (v35 && false) {
                            if (v1 <= (v1 ^ (^(v1)))) {
                                v1 += v1
                                if len(v4) > 0 {
                                    v4[0] = fi33
                                }
                                if v35 {
                                    v2[fi33] = v35
                                    v43 := GenS1{F0: ("hello" + "hello"), F1: fi33}
                                    v35 = true
                                    _ = v43
                                }
                                if false {
                                    switch 32767 {
                                    case 32767:
                                        v44 := GenS3{F0: 255}
                                        _ = v44
                                    case 9223372036854775807:
                                        v2[(fi34 ^ fi34)] = true
                                    case 183:
                                        var v45 uint = 0
                                        _ = v45
                                    default:
                                        var v46 int64 = -962
                                        _ = v46
                                    }
                                    var v47 string = ("foo bar" + "foo bar")
                                    _ = v47
                                }
                                fmt.Println(v35)
                                var v48 bool = v35
                                _ = v48
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    }
                    v49 := v2[fi33]
                    if false {
                        var v50 uint8 = 128
                        var v51 float32 = 3.14159
                        _ = v50
                        _ = v51
                    }
                    _ = fi34
                    _ = v35
                    _ = v49
                }
                v3 += v3
                _ = fi33
            }
            fmt.Println(v1)
        } else {
            var v52 uint64 = 4294967295
            var v53 float32 = (0.0 / 1e10)
            switch 1 {
            case 127:
                v1 = -32768
            case 32768:
                for fi54 := 0; fi54 < 15; fi54++ {
                    if true {
                        v52 = genFn1()
                        if true {
                            for fi55 := 0; fi55 < 15; fi55++ {
                                if ((2 <= 255) || false) {
                                    v56 := GenS4{F0: 51, F1: v52}
                                    switch (("a" + "\n\t") + "foo bar") {
                                    case "\n\t":
                                        var v57 uint64 = (v52 | v52)
                                        _ = v57
                                    case "":
                                        v53 = 1.5
                                    default:
                                        v4 = append(v4, -2147483648)
                                    }
                                    var v58 uint64 = genFn1()
                                    if false {
                                        if ((0 == 4294967295) || (true || true)) {
                                            switch (fi55 * fi55) {
                                            case 2147483647:
                                                if false {
                                                    var v59 uint16 = 218
                                                    v1 -= v1
                                                    _ = v59
                                                } else {
                                                    fmt.Println(v1, fi55, fi55)
                                                    var v60 int64 = 338
                                                    _ = v60
                                                }
                                            case 9:
                                                if ((1.5 + 0.1) == v3) {
                                                    var v61 uint8 = 0
                                                    var v62 int = (^(fi55))
                                                    fmt.Println(fi55, v1, fi54)
                                                    _ = v61
                                                    _ = v62
                                                }
                                            default:
                                                var v63 float32 = v53
                                                _ = v63
                                            }
                                        }
                                        fmt.Println(v53, v58, fi54)
                                        v3 = v3
                                    }
                                    fmt.Println(fi54)
                                    _ = v56
                                    _ = v58
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        if ("hello" <= (("hello" + ("hello" + "a")) + "a")) {
                                            var v64 string = ("hello" + "")
                                            _ = v64
                                        } else {
                                            var v65 int8 = 0
                                            fmt.Println(v52)
                                            if (231 >= 99) {
                                                var v66 int8 = (v65 + v65)
                                                _ = v66
                                            } else {
                                                v52 += 4294967295
                                                v67 := v2[-1]
                                                _ = v67
                                            }
                                            switch (fi54 & int(v1)) {
                                            case 655:
                                                fmt.Println(fi54, v53, v53)
                                            }
                                            _ = v65
                                        }
                                    case "foo bar":
                                        fmt.Println(v1, fi55)
                                    case "hello":
                                        v68 := GenS5{F0: "foo bar", F1: 1, F2: false}
                                        v53 -= v53
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                if false {
                                                    switch (fi54 * 32767) {
                                                    case 42:
                                                        var v69 uint16 = 1
                                                        _ = v69
                                                    default:
                                                        var v70 int64 = 9223372036854775807
                                                        _ = v70
                                                    }
                                                    v3 = (((0.1 / v3) + 0.0) * 1.5)
                                                } else {
                                                    var v71 int64 = 42
                                                    _ = v71
                                                }
                                            case "foo bar":
                                                var v72 uint = 255
                                                _ = v72
                                            case "\n\t":
                                                v1 = v1
                                            default:
                                                var v73 bool = (("a" + (("hello" + ("foo bar" + "hello")) + ("\n\t" + ("\n\t" + "")))) > "foo bar")
                                                _ = v73
                                            }
                                            if false {
                                                v53 = ((v53 - 0.1) - 3.14159)
                                                v52 += v52
                                            }
                                        } else {
                                            fmt.Println(v1, v1)
                                            delete(v2, fi54)
                                            fmt.Println(v1, fi54, fi55)
                                        }
                                        _ = v68
                                    }
                                    var v74 int8 = -128
                                    if true {
                                        v75 := GenS5{F0: ("foo bar" + ""), F1: 127, F2: true}
                                        if (genFn0((1 >= 1), 2147483647, 704) != v1) {
                                            var v76 int64 = 0
                                            v3 -= (v3 / v3)
                                            v74 = v74
                                            _ = v76
                                        } else {
                                            switch (fi55 / ((fi54) | 1)) {
                                            case 128:
                                                v77 := GenS4{F0: 0, F1: 693}
                                                _ = v77
                                            case 32768:
                                                var v78 uint16 = 2
                                                _ = v78
                                            }
                                            if ((v1 - 32767) < v1) {
                                                if (int32(fi55) >= 42) {
                                                    if (false && (!(!(true || (2147483647 != -1))))) {
                                                        var v79 string = (("a" + "") + ("" + ("\n\t" + "foo bar")))
                                                        _ = v79
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v80 uint8 = 42
                                                        _ = v80
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    v4 = append(v4, 32767)
                                                }
                                            } else {
                                                var v81 uint8 = 0
                                                var v82 uint64 = 2
                                                switch fi55 {
                                                case 2147483647:
                                                    switch (("a" + ("" + ("foo bar" + "foo bar"))) + "a") {
                                                    case "":
                                                        v53 = ((v53 / v53) / ((1e10 * ((v53 * 1e10) + 0.1)) / (v53 / 3.14159)))
                                                    case "\n\t":
                                                        fmt.Println(v52, v52, fi54)
                                                    default:
                                                        var v83 float64 = 1e10
                                                        _ = v83
                                                    }
                                                default:
                                                    v3 += float64(float32(v3))
                                                }
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                        var v84 int = (fi55 / ((fi54) | 1))
                                        v52 = v52
                                        _ = v75
                                        _ = v84
                                    }
                                    _ = v74
                                }
                                _ = fi55
                            }
                        } else {
                            switch "a" {
                            case "foo bar":
                                if true {
                                    v3 += v3
                                }
                            case "":
                                v85 := GenS5{F0: "\n\t", F1: 127, F2: (true || ((127 <= 41) && true))}
                                _ = v85
                            default:
                                switch "a" {
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "\n\t":
                                    if (0 > 0) {
                                        for fi86 := 0; fi86 < 15; fi86++ {
                                            v2[(fi54 + fi86)] = false
                                            v3 = (1.5 * 3.14159)
                                            v87 := GenS6{F0: 42, F1: "foo bar"}
                                            _ = fi86
                                            _ = v87
                                        }
                                        switch ("" + "foo bar") {
                                        case "hello":
                                            for fi88 := 0; fi88 < 4; fi88++ {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch fi54 {
                                                    case -9223372036854775808:
                                                        v53 -= float32(1e10)
                                                    case 9223372036854775807:
                                                        delete(v2, (fi54 << ((2) & 63)))
                                                    case 32768:
                                                        v1 -= 342
                                                    default:
                                                        v3 += v3
                                                    }
                                                case "":
                                                    v1 -= (v1 - (v1 % (((v1 / (((v1 << ((158) & 63))) | 1))) | 1)))
                                                case "a":
                                                    v52 = 4294967295
                                                default:
                                                    var v89 string = "foo bar"
                                                    _ = v89
                                                }
                                                _ = fi88
                                            }
                                        case "\n\t":
                                            for fi90 := 0; fi90 < 16; fi90++ {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = fi90
                                            }
                                        default:
                                            v91 := GenS3{F0: 851}
                                            _ = v91
                                        }
                                    } else {
                                        _ = len(v4)
                                        var v92 uint32 = 2
                                        _ = v92
                                    }
                                }
                            }
                            v1 += v1
                        }
                        if (v1 == (v1 - v1)) {
                            for fi93 := 0; fi93 < 2; fi93++ {
                                v1 += v1
                                v94 := GenS3{F0: 478}
                                var v95 uint = 2
                                v1 -= v1
                                v96 := GenS1{F0: "", F1: (^(fi93))}
                                _ = fi93
                                _ = v94
                                _ = v95
                                _ = v96
                            }
                        }
                    }
                    var v97 int64 = 9223372036854775807
                    _ = fi54
                    _ = v97
                }
            case 32767:
                fmt.Println(v3)
            default:
                var v98 uint = 255
                _ = v98
            }
            _ = v52
            _ = v53
        }
    } else {
        var v99 int16 = -32768
        var v100 int32 = 0
        _ = v100
        _ = v99
    }
    switch 32767 {
    case 127:
        var v101 int32 = 42
        _ = v101
    }
    var v102 uint8 = 42
    switch "hello" {
    case "foo bar":
        if (!(-822 <= 2147483647)) {
            v2[915] = (v1 == genFn0(true, 1, 1))
        } else {
            v102 = (v102 & v102)
            v4 = append(v4, 1)
        }
    case "\n\t":
        if (568 <= 543) {
            var v103 float32 = 0.0
            _ = v103
        } else {
            if len(v4) > 0 {
                v4[0] = 400
            }
            var v104 float64 = v3
            _ = v104
        }
    }
    v105 := v2[-74]
    fmt.Println(v1)
    v106 := GenS7{F0: 4294967295}
    v107 := GenS8{F0: v1, F1: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v102
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
