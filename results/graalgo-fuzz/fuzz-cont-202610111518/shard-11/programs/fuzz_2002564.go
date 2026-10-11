package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 bool; F1 int; F2 uint }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS11 struct { F0 int64; F1 int8; F2 uint16 }
type GenS12 struct { F0 uint16; F1 int32; F2 bool }
type GenS13 struct { F0 bool }
func genFn0(p0 uint64, p1 int16) uint {
    v5 := []int{-2147483648, -1, 0}
    p0 = 18446744073709551615
    if true {
        if (((42 < 128) || false) && true) {
            var v6 uint16 = 295
            _ = v6
        }
    } else {
        var v7 uint8 = 96
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(p0, fi8, p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    return uint(p0)
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (true && (439 > 1000))
    v3 := make(map[string]int)
    var v4 int8 = 125
    if true {
        switch -1 {
        case 0:
            for fi9 := 0; fi9 < 12; fi9++ {
                v4 = -35
                fmt.Println(v4)
                var v10 string = ("" + "\n\t")
                _ = fi9
                _ = v10
            }
        case 312:
            switch "foo bar" {
            case "a":
                v11 := GenS1{F0: -41}
                _ = v11
            default:
                var v12 string = "a"
                _ = v12
            }
        default:
            v13 := GenS2{F0: 1}
            _ = v13
        }
        v2 = true
    }
    v4 += 42
    v14 := GenS1{F0: 42}
    v15 := GenS3{F0: 0, F1: -948, F2: (v4 & -109)}
    fmt.Println(v2, v2)
    v16 := GenS4{F0: (v4 % (((v4 + (v4 >> ((388) & 63)))) | 1)), F1: 127}
    var v17 float64 = 0.0
    if false {
        for fi18 := 0; fi18 < 7; fi18++ {
            if false {
                for fi19 := 0; fi19 < 18; fi19++ {
                    if v2 {
                        switch ("a" + (("foo bar" + "\n\t") + "\n\t")) {
                        case "\n\t":
                            switch 32768 {
                            case 0:
                                v20 := GenS2{F0: 393}
                                _ = v20
                            default:
                                switch ("hello" + "foo bar") {
                                case "a":
                                    switch (fi18 ^ (fi18 * 128)) {
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "hello":
                                            v21 := v3["foo bar"]
                                            _ = v21
                                        case "\n\t":
                                            v22 := GenS5{F0: (!((((!(v17 >= v17)) || true) && v2) && ((!(955 <= -32768)) || (1 < 1)))), F1: fi19, F2: 4294967295}
                                            _ = v22
                                        case "foo bar":
                                            switch (fi19 + fi19) {
                                            case 42:
                                                v4 = v4
                                            case 0:
                                                v17 = ((((v17 * (1e10 * ((-2.25 / 0.1) - ((((v17 / v17) + 3.14159) / 1.0) * float64(float32((((v17 - (v17 * v17)) + (v17 - (v17 + float64(3.14159)))) + 3.14159))))))) / 1.5) / float64(float32(1.5))) / 3.14159)
                                            default:
                                                v23 := []int{(-(fi18))}
                                                _ = v23
                                            }
                                        }
                                    case 32768:
                                        fmt.Println(fi18)
                                        v2 = (42 == 42)
                                        var v24 bool = (-2147483648 >= -1)
                                        _ = v24
                                    case 1:
                                        if v2 {
                                            if (2 != 100) {
                                                var v25 int16 = 32767
                                                var v26 string = "\n\t"
                                                v27 := GenS6{F0: int64(v4), F1: 32768}
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v4 -= (v4 % ((v4) | 1))
                                                v28 := []int{2147483647, 127, fi18, (-(fi18))}
                                                _ = v28
                                            }
                                            if (255 == 2) {
                                                v17 = ((0.0 + (v17 * ((1e10 + 3.14159) * v17))) * v17)
                                                switch (("\n\t" + "a") + "foo bar") {
                                                case "":
                                                    v29 := v3[""]
                                                    _ = v29
                                                case "foo bar":
                                                    if false {
                                                        if v2 {
                                                            var v30 int64 = 0
                                                            v17 = v17
                                                            _ = v30
                                                        }
                                                        var v31 uint8 = 2
                                                        fmt.Println(v4, v31, fi19)
                                                        _ = v31
                                                    } else {
                                                        fmt.Println(fi18, fi18, fi19)
                                                    }
                                                default:
                                                    if false {
                                                        if v2 {
                                                            v32 := GenS7{F0: 948, F1: 325}
                                                            var v33 int32 = 0
                                                            var v34 int32 = 2147483647
                                                            _ = v32
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        var v35 uint8 = 49
                                                        v2 = (0 != -968)
                                                        switch fi19 {
                                                        case 42:
                                                            switch (fi19 - fi19) {
                                                            case 32767:
                                                                v35 = 2
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                        _ = v35
                                                    } else {
                                                        if ("a" == ("foo bar" + "foo bar")) {
                                                            v36 := GenS8{F0: 127}
                                                            v4 -= 52
                                                            _ = v36
                                                        }
                                                        if v2 {
                                                            v1[("foo bar" + ("a" + ""))] = "a"
                                                            fmt.Println(v17, v17, fi19)
                                                        }
                                                    }
                                                    v4 = 42
                                                }
                                            } else {
                                                v37 := GenS9{F0: 42}
                                                _ = v37
                                            }
                                        }
                                    }
                                default:
                                    var v38 int32 = 2147483647
                                    _ = v38
                                }
                            }
                        default:
                            if (-1 >= 0) {
                                fmt.Println(fi18)
                            }
                        }
                        if (2147483647 <= 2) {
                            if (0 <= 1) {
                                var v39 string = ""
                                switch (fi19 ^ (fi19 & fi19)) {
                                case -1:
                                    var v40 uint16 = 201
                                    _ = v40
                                default:
                                    switch 726 {
                                    case 1:
                                        fmt.Println(fi18)
                                    case 127:
                                        v17 += (((0.1 - v17) * ((v17 - ((v17 - v17) - (((v17 + (v17 / -2.25)) + (float64(float32(v17)) * v17)) / v17))) - v17)) + v17)
                                    default:
                                        fmt.Println(v17, v17)
                                    }
                                }
                                v41 := GenS10{F0: 2, F1: 255, F2: 42}
                                _ = v39
                                _ = v41
                            } else {
                                if ((1.5 == 0.0) && v2) {
                                    fmt.Println(v2, v17)
                                }
                                if v2 {
                                    if false {
                                        v42 := GenS11{F0: -9223372036854775808, F1: v4, F2: 420}
                                        _ = v42
                                    } else {
                                        v3["foo bar"] = (fi19 ^ fi19)
                                    }
                                    v4 = (v4 >> ((564) & 63))
                                } else {
                                    _ = len(v3)
                                    v43 := []int{}
                                    _ = v43
                                }
                                v44 := GenS4{F0: -8, F1: 144}
                                var v45 int = (fi19 | fi19)
                                _ = v44
                                _ = v45
                            }
                        }
                        var v46 int32 = 2147483647
                        _ = v46
                    }
                    fmt.Println(v2, fi19)
                    _ = fi19
                }
                v47 := v3["foo bar"]
                _ = v47
            } else {
                if v2 {
                    if (((v4 * -1) <= 42) && (v2 || (v2 && true))) {
                        switch 42 {
                        case 2147483647:
                            for fi48 := 0; fi48 < 10; fi48++ {
                                switch "hello" {
                                case "":
                                    if ((884 <= 42) || false) {
                                        if true {
                                            if v2 {
                                                if (127 >= (v4 / ((127) | 1))) {
                                                    v49 := GenS12{F0: 32767, F1: -1, F2: (595 > 42)}
                                                    var v50 int16 = -24
                                                    v51 := []int{fi48, fi18, fi48, (fi48 ^ -9223372036854775808)}
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                                fmt.Println(v2)
                                                v3["a"] = 0
                                            } else {
                                                v52 := []int{(fi18 ^ (fi18 | fi18))}
                                                switch (fi18 >> ((2) & 63)) {
                                                case 614:
                                                    v17 = 0.1
                                                }
                                                _ = v52
                                            }
                                            if v2 {
                                                var v53 int = fi18
                                                _ = v53
                                            }
                                        } else {
                                            fmt.Println(fi48, v4)
                                        }
                                        switch (fi18 >> ((42) & 63)) {
                                        case -9223372036854775808:
                                            v54 := GenS8{F0: 2}
                                            _ = v54
                                        default:
                                            if v2 {
                                                v4 = -1
                                                v55 := []int{fi18}
                                                v56 := GenS13{F0: (!(128 == 2))}
                                                fmt.Println(v2, fi18, v2)
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        v4 -= v4
                                    }
                                case "foo bar":
                                    _ = len(v3)
                                    var v57 int32 = 539
                                    _ = v57
                                }
                                v4 -= (v4 | (v4 * (v4 & (-(v4)))))
                                v58 := GenS3{F0: 1, F1: -266, F2: v4}
                                _ = fi48
                                _ = v58
                            }
                        case 32768:
                            if true {
                                v59 := []int{fi18, -2147483648, fi18}
                                _ = v59
                            } else {
                                v60 := []int{0}
                                v17 += ((v17 * 1.5) * float64(0.0))
                                _ = v60
                            }
                        default:
                            if v2 {
                                if v2 {
                                    var v61 int8 = v4
                                    v62 := []int{fi18, (^(fi18)), fi18, fi18}
                                    fmt.Println(v2, fi18)
                                    _ = v61
                                    _ = v62
                                }
                            } else {
                                for fi63 := 0; fi63 < 5; fi63++ {
                                    v64 := []int{}
                                    _ = fi63
                                    _ = v64
                                }
                            }
                        }
                        for fi65 := 0; fi65 < 4; fi65++ {
                            v17 -= v17
                            var v66 float64 = (((((v17 - ((((1e10 - 3.14159) + 1.5) + (((v17 / float64(3.14159)) + 0.1) / v17)) - v17)) - ((float64((1e10 * (0.1 + 0.0))) + v17) / float64(float32((v17 - -2.25))))) - v17) - float64(0.0)) - ((v17 - v17) - v17))
                            _ = fi65
                            _ = v66
                        }
                    } else {
                        v2 = v2
                        for fi67 := 0; fi67 < 1; fi67++ {
                            switch (-(fi67)) {
                            case 0:
                                if (2147483648 != 4294967295) {
                                    fmt.Println(v4, fi67)
                                    var v68 float64 = ((v17 * (v17 - v17)) + v17)
                                    v69 := GenS9{F0: 127}
                                    _ = v68
                                    _ = v69
                                }
                            case -260:
                                var v70 int = 0
                                _ = v70
                            case 483:
                                v2 = (651 > -699)
                            default:
                                fmt.Println(v17, v4, v17)
                                var v71 float64 = v17
                                v2 = (518 < 562)
                                _ = v71
                            }
                            var v72 int = fi18
                            v73 := []int{-2147483648, fi67, 2147483647}
                            _ = fi67
                            _ = v72
                            _ = v73
                        }
                    }
                    switch fi18 {
                    case 701:
                        var v74 int32 = -986
                        _ = v74
                    default:
                        v75 := []int{fi18, fi18, 127}
                        _ = v75
                    }
                    if false {
                        v76 := GenS10{F0: 2147483648, F1: 42, F2: int64(v4)}
                        fmt.Println(v17, v2, fi18)
                        if (!v2) {
                            var v77 uint32 = 946
                            _ = v77
                        }
                        _ = v76
                    }
                } else {
                    if false {
                        var v78 int32 = 1
                        _ = v78
                    } else {
                        if true {
                            fmt.Println(fi18, v17, v4)
                            if false {
                                var v79 int16 = 0
                                _ = v79
                            }
                            _ = len(v1)
                        }
                        v80 := GenS5{F0: (v4 == v4), F1: 4294967295, F2: 18446744073709551615}
                        v81 := GenS2{F0: 2}
                        _ = v80
                        _ = v81
                    }
                }
                fmt.Println(fi18, v17)
            }
            var v82 string = "a"
            _ = fi18
            _ = v82
        }
    }
    var v83 float64 = (3.14159 / v17)
    var v84 int64 = 664
    if (42 > 42) {
        v2 = true
        switch 4294967295 {
        case 128:
            v84 -= (v84 & -1)
        case -424:
            v83 = 0.1
        case -2147483648:
            v84 -= v84
        default:
            fmt.Println(v83)
        }
    } else {
        v17 = (v17 * (3.14159 - (1e10 * 1.5)))
        if (!v2) {
            switch 2147483647 {
            case 0:
                _ = len(v1)
            default:
                if true {
                    v4 = -128
                } else {
                    v85 := []int{-9223372036854775808, -499, 9223372036854775807}
                    _ = v85
                }
            }
            fmt.Println(v84, v83, v2)
            for fi86 := 0; fi86 < 7; fi86++ {
                v1["hello"] = (("" + "") + ("a" + ("hello" + (("" + "\n\t") + ("" + "foo bar")))))
                fmt.Println(v2, v83)
                _ = fi86
            }
        } else {
            v87 := GenS1{F0: v84}
            _ = v87
        }
        var v88 int16 = 431
        _ = v88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
}
