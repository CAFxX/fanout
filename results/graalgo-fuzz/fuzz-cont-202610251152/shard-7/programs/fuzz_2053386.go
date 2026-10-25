package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 int64 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int16; F2 uint8 }
func genFn0(p0 uint32, p1 bool, p2 uint64) int8 {
    var v5 string = (("foo bar" + "a") + "hello")
    v6 := GenS1{F0: int16(p2), F1: 43}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 111
}
func genFn1(p0 int16, p1 uint16) int16 {
    switch "hello" {
    case "":
        p1 += 0
    }
    v7 := []int{128}
    fmt.Println(p1)
    var v8 uint16 = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 2147483647
    v3 := make(map[int]uint32)
    v4 := []int{-588, 818, 1, 128}
    var v9 uint32 = (^(v2))
    fmt.Println(v9, v1)
    for fi10 := 0; fi10 < 9; fi10++ {
        if (true && false) {
            var v11 uint16 = 2
            switch "" {
            case "a":
                v11 -= (^(v11))
            case "":
                if true {
                    fmt.Println(v9, v2)
                    if true {
                        var v12 string = "a"
                        _ = v12
                    } else {
                        switch "a" {
                        case "\n\t":
                            var v13 uint8 = 128
                            _ = v13
                        case "hello":
                            v2 = v9
                        }
                        for fi14 := 0; fi14 < 2; fi14++ {
                            v15 := GenS1{F0: 0, F1: 100}
                            v16 := GenS2{F0: 4}
                            v9 = 4294967295
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        delete(v3, fi10)
                    }
                    fmt.Println(v2)
                } else {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS3{F0: -1, F1: (1 > (v9 % ((2147483648) | 1))), F2: 9223372036854775807}
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(v1)
                }
            }
            _ = v11
        }
        var v19 bool = (!(float32(v1) >= 1e10))
        if v19 {
            if (42 == uint(v2)) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v2, v1)
                    switch fi20 {
                    case 128:
                        _ = len(v3)
                    default:
                        var v21 int64 = 42
                        _ = v21
                    }
                    if len(v4) > 0 {
                        v4[0] = fi20
                    }
                    if (1e10 != 0.1) {
                        var v22 uint8 = 127
                        switch "" {
                        case "hello":
                            v23 := GenS2{F0: 0}
                            _ = v23
                        case "a":
                            var v24 float32 = (0.0 / 3.14159)
                            _ = v24
                        }
                        _ = v22
                    }
                    _ = fi20
                }
                v1 = v1
            }
            var v25 string = "\n\t"
            _ = v25
        }
        _ = fi10
        _ = v19
    }
    if (1.5 <= 1.5) {
        switch "\n\t" {
        case "hello":
            if ((v9 << ((0) & 63)) == 4294967295) {
                v1 = 3.14159
                if len(v4) > 0 {
                    v4[0] = -434
                }
            } else {
                for fi26 := 0; fi26 < 9; fi26++ {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if (-1 >= 1) {
                            if ((v1 - v1) != float64(1.5)) {
                                var v28 float32 = 1.5
                                var v29 uint32 = v2
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = fi27
                    }
                    v1 -= float64((-2.25 / 1.5))
                    _ = fi26
                }
            }
        case "a":
            v9 -= (v2 + (v2 * (v9 >> ((65535) & 63))))
        case "foo bar":
            v2 = v9
        default:
            fmt.Println(v1, v1, v9)
        }
        v30 := GenS4{F0: "", F1: 10}
        _ = v30
    } else {
        switch 32767 {
        case 449:
            var v31 bool = false
            _ = v31
        }
        fmt.Println(v9, v1)
        if (655 < 18446744073709551615) {
            if false {
                _ = len(v3)
            }
            v3[-246] = v9
            for fi32 := 0; fi32 < 8; fi32++ {
                if false {
                    fmt.Println(v2, v2, v2)
                    v1 = v1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v33 uint64 = 42
                    v9 = (^(v9))
                    _ = v33
                }
                var v34 string = ("a" + "hello")
                _ = fi32
                _ = v34
            }
        }
    }
    var v35 float64 = 1.5
    for fi36 := 0; fi36 < 3; fi36++ {
        v1 -= v1
        _ = fi36
    }
    var v37 uint16 = 38
    var v38 float32 = 1e10
    switch -1 {
    case -1:
        v4 = append(v4, 32768)
    default:
        if false {
            switch 410 {
            case 127:
                var v39 float32 = ((-2.25 / -2.25) - v38)
                _ = v39
            default:
                if false {
                    if true {
                        for fi40 := 0; fi40 < 7; fi40++ {
                            var v41 float64 = -2.25
                            switch "" {
                            case "foo bar":
                                v42 := GenS5{F0: -2147483648}
                                _ = v42
                            case "\n\t":
                                if (45 != genFn1(-1, (v37 % ((2) | 1)))) {
                                    var v43 int16 = -249
                                    for fi44 := 0; fi44 < 16; fi44++ {
                                        v38 = v38
                                        var v45 float64 = float64(0.1)
                                        _ = fi44
                                        _ = v45
                                    }
                                    fmt.Println(v37, v1, v35)
                                    _ = v43
                                }
                            default:
                                switch ("hello" + "foo bar") {
                                case "":
                                    v9 -= 1
                                case "\n\t":
                                    var v46 int32 = -2147483648
                                    _ = v46
                                case "foo bar":
                                    v47 := GenS6{F0: (fi40 << ((42) & 63)), F1: 42, F2: 2}
                                    _ = v47
                                default:
                                    var v48 string = ("foo bar" + "foo bar")
                                    _ = v48
                                }
                            }
                            _ = fi40
                            _ = v41
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        var v49 bool = false
                        _ = v49
                    case "foo bar":
                        var v50 uint8 = 42
                        _ = v50
                    default:
                        if false {
                            var v51 float32 = v38
                            if true {
                                var v52 bool = false
                                v35 -= (0.1 - 0.1)
                                var v53 int32 = -2147483648
                                _ = v52
                                _ = v53
                            } else {
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    if (0.1 != (v51 / (v51 - v51))) {
                                        _ = len(v4)
                                        fmt.Println(v1, v1)
                                    }
                                    _ = fi54
                                }
                                var v55 uint32 = v2
                                v55 -= (v55 >> ((4294967295) & 63))
                                _ = v55
                            }
                            switch -808 {
                            case 0:
                                var v56 float32 = v51
                                _ = v56
                            case -1:
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            }
                            _ = v51
                        } else {
                            var v57 float32 = v38
                            var v58 int8 = 42
                            _ = v57
                            _ = v58
                        }
                    }
                    for fi59 := 0; fi59 < 8; fi59++ {
                        var v60 int16 = 0
                        if false {
                            switch 42 {
                            case 1:
                                if true {
                                    v1 = v35
                                    var v61 int64 = 9223372036854775807
                                    if true {
                                        var v62 bool = true
                                        if false {
                                            _ = len(v4)
                                            if false {
                                                switch "foo bar" {
                                                case "":
                                                    var v63 int16 = v60
                                                    _ = v63
                                                case "a":
                                                    var v64 float32 = (0.1 + v38)
                                                    _ = v64
                                                case "\n\t":
                                                    if len(v4) > 0 {
                                                        v4[0] = fi59
                                                    }
                                                default:
                                                    for fi65 := 0; fi65 < 8; fi65++ {
                                                        v62 = v62
                                                        _ = fi65
                                                    }
                                                }
                                                var v66 int8 = genFn0(65535, (v62 || true), 18446744073709551615)
                                                v66 += (v66 - -52)
                                                _ = v66
                                            } else {
                                                if (false && v62) {
                                                    v9 = (^(v2))
                                                    if len(v4) > 0 {
                                                        v4[0] = (^(fi59))
                                                    }
                                                } else {
                                                    switch "a" {
                                                    case "a":
                                                        for fi67 := 0; fi67 < 1; fi67++ {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v68 float32 = v38
                                                                _ = v68
                                                            case "a":
                                                                var v69 uint64 = 1
                                                                _ = v69
                                                            case "hello":
                                                                var v70 int64 = 0
                                                                _ = v70
                                                            default:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v71 int32 = 621
                                                                    _ = v71
                                                                case "a":
                                                                    v72 := GenS1{F0: int16(v61), F1: -128}
                                                                    _ = v72
                                                                case "foo bar":
                                                                    switch (fi67 & fi59) {
                                                                    case 4294967295:
                                                                        var v73 int8 = genFn0(v9, false, 837)
                                                                        _ = v73
                                                                    }
                                                                    var v74 uint16 = v37
                                                                    _ = v74
                                                                }
                                                            }
                                                            v60 -= genFn1(genFn1(v60, 32767), (v37 & uint16(v9)))
                                                            switch "\n\t" {
                                                            case "":
                                                                v9 -= v9
                                                                var v75 uint8 = 2
                                                                _ = v75
                                                            case "a":
                                                                if v62 {
                                                                    if v62 {
                                                                        v9 = v9
                                                                    }
                                                                    var v76 int = fi67
                                                                    var v77 uint64 = 18446744073709551615
                                                                    _ = v76
                                                                    _ = v77
                                                                } else {
                                                                    var v78 bool = (376 == 306)
                                                                    v79 := GenS5{F0: 42}
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                            case "\n\t":
                                                                if (((((1.5 * 0.0) * ((((float64(v38) * 0.0) * -2.25) * ((v1 + float64(v38)) / (((0.0 + float64(float32((float64(float32(((1.5 + 1.5) + v35))) + (float64((v38 * 1.5)) - -2.25))))) - v1) / ((1.5 - (v35 / v1)) / v1)))) + -2.25)) / v35) == v35) && (-115 <= -54)) {
                                                                    var v80 uint32 = v2
                                                                    v60 += v60
                                                                    _ = v80
                                                                }
                                                            }
                                                            var v81 bool = false
                                                            _ = fi67
                                                            _ = v81
                                                        }
                                                    }
                                                    var v82 string = "hello"
                                                    _ = v82
                                                }
                                                switch "" {
                                                case "foo bar":
                                                    var v83 uint = 18446744073709551615
                                                    _ = v83
                                                case "":
                                                    for fi84 := 0; fi84 < 3; fi84++ {
                                                        var v85 uint16 = 789
                                                        var v86 float32 = v38
                                                        var v87 int64 = v61
                                                        _ = fi84
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                case "hello":
                                                    v61 = 1
                                                }
                                            }
                                            for fi88 := 0; fi88 < 9; fi88++ {
                                                _ = len(v4)
                                                var v89 int32 = 483
                                                _ = fi88
                                                _ = v89
                                            }
                                        }
                                        var v90 uint32 = (v2 ^ v9)
                                        _ = v62
                                        _ = v90
                                    }
                                    _ = v61
                                } else {
                                    var v91 uint64 = 0
                                    var v92 uint8 = 64
                                    var v93 bool = false
                                    _ = v91
                                    _ = v92
                                    _ = v93
                                }
                            case 42:
                                v37 = (v37 >> ((605) & 63))
                            default:
                                var v94 uint16 = v37
                                _ = v94
                            }
                        }
                        _ = fi59
                        _ = v60
                    }
                }
            }
            v95 := GenS2{F0: 0}
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v9
}
