package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint8; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int8; F1 bool; F2 int8 }
func genFn0() uint16 {
    switch "a" {
    case "foo bar":
        var v7 string = "\n\t"
        _ = v7
    case "a":
        v8 := GenS1{F0: 65535, F1: 700}
        _ = v8
    }
    switch 127 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint = 91
            v10 = (v10 ^ v10)
            v11 := []int{fi9, (fi9 << (((v10 << (((v10 + v10)) & 63))) & 63))}
            _ = fi9
            _ = v10
            _ = v11
        }
    case -328:
        if (false || false) {
            var v12 bool = (3.14159 <= (float32((-2.25 * (1e10 * (0.1 * 1.5)))) / 0.1))
            _ = v12
        } else {
            v13 := []int{-724, 9223372036854775807}
            _ = v13
        }
    }
    var v14 float32 = (-2.25 / -2.25)
    _ = v14
    return 2
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]int)
    var v3 int = -9223372036854775808
    var v4 int = (v3 / ((v3) | 1))
    var v5 float64 = (1e10 + 0.0)
    var v6 uint8 = 42
    if ((v1 & v1) <= -783) {
        var v15 int = (v3 - (-(v4)))
        _ = v15
    }
    switch "" {
    case "\n\t":
        v16 := GenS1{F0: 112, F1: 2147483647}
        _ = v16
    case "":
        v17 := GenS1{F0: 203, F1: 65535}
        _ = v17
    case "hello":
        var v18 float32 = 3.14159
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 42}
    v20 := GenS2{F0: 611}
    if (("" + "\n\t") >= ("" + "a")) {
        switch ("" + (("hello" + (("" + "hello") + "hello")) + ("\n\t" + "a"))) {
        case "foo bar":
            fmt.Println(v6)
        case "hello":
            if ((v3 >> ((833) & 63)) < 32768) {
                fmt.Println(v3)
                var v21 uint32 = 65535
                _ = v21
            }
        default:
            if (255 > 367) {
                v22 := GenS1{F0: 42, F1: 2147483648}
                v23 := []int{v3, v4, 32768, v3}
                _ = v22
                _ = v23
            }
        }
        v1 = -787
        v24 := []int{v3, (v4 % ((42) | 1)), v3}
        _ = v24
    }
    if false {
        var v25 uint32 = 2147483647
        _ = v25
    } else {
        v3 -= (-(v3))
        for fi26 := 0; fi26 < 12; fi26++ {
            v4 = fi26
            v27 := GenS3{F0: "\n\t", F1: 32767, F2: v4}
            _ = fi26
            _ = v27
        }
    }
    v3 += 841
    v28 := GenS3{F0: ("" + (((("\n\t" + "a") + ("foo bar" + "\n\t")) + "a") + "")), F1: -32768, F2: int(v1)}
    if ((((((2 <= 0) && false) || (true || (!(false || false)))) && true) && false) && false) {
        var v29 int8 = -1
        if (v5 <= (float64(float32((((1e10 / v5) + 0.0) - v5))) * -2.25)) {
            var v30 int32 = 42
            v31 := GenS4{F0: 2}
            if (v5 < 1e10) {
                for fi32 := 0; fi32 < 19; fi32++ {
                    fmt.Println(v6, v1)
                    v30 -= (v30 | (v30 + (v30 >> ((127) & 63))))
                    if true {
                        var v33 uint64 = 18446744073709551615
                        _ = v33
                    } else {
                        var v34 float64 = ((0.1 / float64(float32(((v5 * v5) + 0.0)))) / 1e10)
                        _ = v34
                    }
                    _ = fi32
                }
                delete(v2, (v4 | (v4 << ((4294967295) & 63))))
                v35 := GenS5{F0: 1, F1: v6, F2: -420}
                _ = v35
            }
            _ = v30
            _ = v31
        } else {
            if (!((((v6 <= 2) && ((v6 & 42) > (v6 - (v6 - 127)))) || (-1 != 1)) && true)) {
                var v36 bool = (v4 == (v3 - v3))
                for fi37 := 0; fi37 < 4; fi37++ {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        switch "hello" {
                        case "":
                            var v39 int = -777
                            _ = v39
                        case "hello":
                            v40 := GenS4{F0: 255}
                            _ = v40
                        default:
                            if v36 {
                                v6 -= v6
                            } else {
                                v5 -= (float64((0.1 * (((float32((v5 - (float64(0.0) + -2.25))) * 0.0) - (1.5 - -2.25)) + 0.1))) + (float64((((1.5 / 1.0) + -2.25) / ((0.0 / 1e10) / (((0.1 / 3.14159) / 1.5) / float32(v5))))) / (v5 - -2.25)))
                                v41 := GenS6{F0: 1}
                                var v42 string = "foo bar"
                                if v36 {
                                    if (v3 == (fi37 % (((v3 + (fi38 + (v4 & fi37)))) | 1))) {
                                        var v43 int16 = -32768
                                        var v44 int8 = v29
                                        _ = v43
                                        _ = v44
                                    } else {
                                        switch 32767 {
                                        case 50:
                                            v42 = v42
                                        default:
                                            var v45 string = (v42 + v42)
                                            _ = v45
                                        }
                                    }
                                    v46 := GenS7{F0: 1}
                                    _ = v46
                                } else {
                                    fmt.Println(v36)
                                    switch v42 {
                                    case "\n\t":
                                        v47 := []int{(v3 / ((4294967295) | 1)), 32768, int(v29), (v4 % ((127) | 1))}
                                        _ = v47
                                    case "foo bar":
                                        if v36 {
                                            v4 += (v3 + (^(fi37)))
                                            if ((((v1 << ((18446744073709551615) & 63)) != -901) && (4294967295 >= 0)) || v36) {
                                                v29 -= 62
                                                var v48 string = ((v42 + v42) + v42)
                                                _ = v48
                                            } else {
                                                v42 = (v42 + v42)
                                                v49 := GenS2{F0: 318}
                                                var v50 int = (v4 + (-(v4)))
                                                _ = v49
                                                _ = v50
                                            }
                                            fmt.Println(v5, v42, v1)
                                        } else {
                                            fmt.Println(v4, v29, fi38)
                                        }
                                    case "hello":
                                        if v36 {
                                            v51 := GenS2{F0: 2147483647}
                                            _ = v51
                                        } else {
                                            v52 := GenS3{F0: "hello", F1: 1000, F2: (^(v4))}
                                            _ = v52
                                        }
                                    default:
                                        var v53 uint16 = genFn0()
                                        _ = v53
                                    }
                                    v5 -= v5
                                }
                                v1 += v1
                                _ = v41
                                _ = v42
                            }
                        }
                        _ = fi38
                    }
                    _ = fi37
                }
                _ = v36
            } else {
                var v54 float32 = (1e10 + 1e10)
                fmt.Println(v3, v5, v6)
                _ = v54
            }
            v55 := GenS8{F0: "\n\t", F1: 1}
            _ = v55
        }
        _ = v29
    } else {
        v6 -= v6
        for fi56 := 0; fi56 < 19; fi56++ {
            if true {
                for fi57 := 0; fi57 < 15; fi57++ {
                    fmt.Println(v3, v5)
                    _ = fi57
                }
                v58 := GenS7{F0: 666}
                _ = v58
            }
            _ = fi56
        }
        if (2147483647 < uint32(v6)) {
            switch "\n\t" {
            case "hello":
                var v59 uint64 = 18446744073709551615
                _ = v59
            case "foo bar":
                v3 = v3
            case "\n\t":
                v60 := GenS1{F0: genFn0(), F1: 0}
                _ = v60
            }
            switch ("hello" + ("hello" + "hello")) {
            case "a":
                var v61 string = ((("" + ("a" + ("a" + "foo bar"))) + "foo bar") + "")
                _ = v61
            case "hello":
                var v62 int32 = 0
                _ = v62
            case "\n\t":
                if true {
                    switch ("" + (("a" + "") + "\n\t")) {
                    case "hello":
                        switch "a" {
                        case "foo bar":
                            var v63 int32 = 42
                            _ = v63
                        default:
                            var v64 uint = 18446744073709551615
                            _ = v64
                        }
                    }
                    var v65 bool = (0.1 >= 1e10)
                    v66 := GenS6{F0: -121}
                    _ = v65
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 10; fi67++ {
                        v4 = v4
                        _ = fi67
                    }
                }
            }
            for fi68 := 0; fi68 < 7; fi68++ {
                var v69 bool = false
                _ = fi68
                _ = v69
            }
        } else {
            switch "a" {
            case "foo bar":
                v1 -= (v1 % ((v1) | 1))
            default:
                switch ((("a" + "hello") + ("" + "hello")) + (("a" + "") + "\n\t")) {
                case "foo bar":
                    v1 += (v1 >> ((0) & 63))
                case "":
                    var v70 int = (v4 | v4)
                    _ = v70
                }
            }
        }
    }
    v71 := GenS9{F0: -39, F1: false, F2: -100}
    if (true && true) {
        var v72 int16 = 32767
        for fi73 := 0; fi73 < 11; fi73++ {
            v74 := GenS4{F0: 0}
            v3 -= v3
            switch "a" {
            case "hello":
                var v75 bool = (!(-128 >= 1))
                _ = v75
            default:
                fmt.Println(v6, v3)
            }
            _ = fi73
            _ = v74
        }
        v72 = (^(v72))
        _ = v72
    } else {
        if true {
            fmt.Println(v4, v1)
            v76 := GenS9{F0: -128, F1: true, F2: 1}
            for fi77 := 0; fi77 < 6; fi77++ {
                fmt.Println(v1)
                _ = fi77
            }
            _ = v76
        } else {
            switch v4 {
            case -886:
                if (!(-114 > 42)) {
                    fmt.Println(v1, v4, v5)
                    var v78 uint8 = 2
                    var v79 float32 = 1.5
                    _ = v78
                    _ = v79
                }
            case 2147483647:
                for fi80 := 0; fi80 < 18; fi80++ {
                    if true {
                        v5 += v5
                    }
                    fmt.Println(v6, v3)
                    for fi81 := 0; fi81 < 20; fi81++ {
                        v82 := []int{}
                        if len(v82) > 0 {
                            v82[0] = (fi80 + 1)
                        }
                        _ = fi81
                        _ = v82
                    }
                    _ = fi80
                }
            }
            var v83 uint32 = 42
            var v84 uint64 = 4294967295
            _ = v83
            _ = v84
        }
        v3 -= v4
    }
    v3 = v4
    for fi85 := 0; fi85 < 16; fi85++ {
        if ((v4 >> ((255) & 63)) >= (fi85 / (((fi85 - (v3 / (((^(v3))) | 1)))) | 1))) {
            v86 := []int{(-(v4)), (v3 | fi85), (fi85 | (v4 * (-(v3))))}
            _ = v86
        } else {
            var v87 uint16 = 0
            switch v3 {
            case 1:
                v5 = v5
            case 2147483647:
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1)
                default:
                    for fi88 := 0; fi88 < 6; fi88++ {
                        var v89 float32 = (-2.25 * (((((((0.1 + 3.14159) - ((float32(((float64((1e10 * (((float32(v5) / float32(v5)) / 3.14159) / 0.1))) * v5) + v5)) - 3.14159) - 1e10)) + 1e10) * float32(v5)) / float32(3.14159)) + 1e10) + 0.0))
                        v90 := []int{fi88, int(v87)}
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                }
            }
            _ = v87
        }
        var v91 uint16 = 65535
        v92 := GenS5{F0: 671, F1: v6, F2: 0}
        _ = fi85
        _ = v91
        _ = v92
    }
    for fi93 := 0; fi93 < 10; fi93++ {
        var v94 int32 = -2147483648
        fmt.Println(v1)
        _ = fi93
        _ = v94
    }
    var v95 float64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v95
}
