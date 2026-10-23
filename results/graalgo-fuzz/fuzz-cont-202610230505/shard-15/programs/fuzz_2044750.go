package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint64; F2 bool }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 bool; F1 int8; F2 int64 }
type GenS9 struct { F0 uint8; F1 int64 }
func genFn0() uint16 {
    var v4 int16 = -32768
    v5 := GenS1{F0: -128, F1: v4, F2: "foo bar"}
    _ = v4
    _ = v5
    return 86
}
func genFn1(p0 uint64) string {
    p0 = (^(p0))
    if (false || false) {
        v6 := []int{}
        if (!false) {
            p0 += (^(p0))
        }
        _ = v6
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int8 = 0
            if false {
                var v9 int64 = 9223372036854775807
                if true {
                    p0 -= (p0 | (p0 | 0))
                }
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    v10 := GenS1{F0: 0, F1: 1, F2: "hello"}
    switch 638 {
    case 533:
        switch -9223372036854775808 {
        case -2147483648:
            p0 = 1
        case 627:
            v11 := GenS2{F0: "foo bar"}
            _ = v11
        }
    case -9223372036854775808:
        if true {
            v12 := GenS2{F0: (("foo bar" + "foo bar") + ("foo bar" + "foo bar"))}
            _ = v12
        } else {
            p0 += (^(p0))
        }
    case 0:
        for fi13 := 0; fi13 < 5; fi13++ {
            switch fi13 {
            case 32767:
                v14 := GenS3{F0: -90}
                _ = v14
            case 42:
                v15 := []int{}
                _ = v15
            case 761:
                v16 := []int{fi13}
                _ = v16
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                if false {
                    var v18 float32 = 0.1
                    var v19 float32 = 0.1
                    v20 := GenS1{F0: -47, F1: -1, F2: ""}
                    fmt.Println(fi17)
                    v21 := GenS2{F0: (((("" + "a") + ("" + "foo bar")) + "\n\t") + "")}
                    v22 := []int{(fi17 + (fi17 << ((4294967295) & 63))), fi13, -139, fi17}
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    if (0 <= 107) {
                        if false {
                            if (!(0.0 > (((-2.25 * 0.0) * 3.14159) / ((1.5 + ((0.1 - (1.5 + ((((0.0 + 0.1) - 1e10) * -2.25) / 1.0))) - ((0.0 + 1e10) - (((1.5 - (1.5 - float64((float32(1.5) - 0.0)))) + float64(0.0)) - 0.1)))) / ((float64(1.5) * 0.1) * 1e10))))) {
                                fmt.Println(fi13)
                                v23 := []int{}
                                _ = v23
                            } else {
                                var v24 int8 = -50
                                v25 := []int{(fi13 | fi13)}
                                var v26 uint64 = (p0 ^ 395)
                                var v27 int32 = -32
                                p0 = (p0 >> ((42) & 63))
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            if (!(false && (2147483648 >= 42))) {
                                var v28 float64 = 1e10
                                var v29 float32 = 0.1
                                _ = v28
                                _ = v29
                            }
                            if (42 <= 0) {
                                p0 -= (p0 | (p0 & 4294967295))
                            }
                        }
                        v30 := []int{(fi13 & 42), (^(fi17)), fi13, (fi17 & fi13)}
                        v30 = append(v30, (fi17 / ((fi17) | 1)))
                        _ = v30
                    }
                    p0 = p0
                    if (32767 >= -470) {
                        v31 := GenS4{F0: 42, F1: 2147483647}
                        _ = v31
                    }
                    var v32 int8 = 100
                    _ = v32
                }
                _ = fi17
            }
            _ = fi13
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v10
    return "foo bar"
}

func main() {
    var v1 int = -1
    v2 := []int{(v1 + v1)}
    var v3 uint16 = 255
    v1 -= (-(v1))
    v33 := GenS2{F0: (((("hello" + "a") + "hello") + "a") + "a")}
    v2 = append(v2, 32768)
    fmt.Println(v3, v1)
    switch v1 {
    case 42:
        for fi34 := 0; fi34 < 17; fi34++ {
            switch -9223372036854775808 {
            case 9223372036854775807:
                v2 = append(v2, fi34)
            case 101:
                if (false || false) {
                    if ((true || false) && (0 >= 2147483647)) {
                        var v35 int16 = -408
                        switch "a" {
                        case "\n\t":
                            v3 = v3
                        case "a":
                            if true {
                                v36 := GenS5{F0: 1, F1: (v3 << ((1) & 63))}
                                if ((genFn0() > 255) || ("a" <= "hello")) {
                                    fmt.Println(v35)
                                    var v37 int32 = -1
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = (fi34 ^ (v1 / (((v1 << ((889) & 63))) | 1)))
                                        }
                                        v1 += (fi34 / (((-(v1))) | 1))
                                        v3 = (v3 - v3)
                                    }
                                    _ = v37
                                } else {
                                    _ = len(v2)
                                    switch fi34 {
                                    case 32768:
                                        switch "hello" {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = 42
                                            }
                                        case "\n\t":
                                            if true {
                                                var v38 int16 = int16(v3)
                                                _ = v38
                                            }
                                        case "a":
                                            var v39 int16 = (v35 >> ((2) & 63))
                                            _ = v39
                                        default:
                                            for fi40 := 0; fi40 < 14; fi40++ {
                                                if (42 == 0) {
                                                    v35 = 32767
                                                }
                                                fmt.Println(v3)
                                                var v41 int64 = 475
                                                _ = fi40
                                                _ = v41
                                            }
                                        }
                                    }
                                    v42 := GenS2{F0: ("foo bar" + "a")}
                                    _ = v42
                                }
                                for fi43 := 0; fi43 < 12; fi43++ {
                                    v3 += v3
                                    _ = fi43
                                }
                                _ = v36
                            } else {
                                var v44 uint = 0
                                _ = v44
                            }
                        case "hello":
                            if (0.0 != (-2.25 / ((0.1 * 1.5) * -2.25))) {
                                var v45 int32 = -1
                                _ = v45
                            } else {
                                fmt.Println(fi34, v35, v35)
                                fmt.Println(v35, v35)
                                _ = len(v2)
                            }
                        default:
                            if (2147483647 <= 1) {
                                fmt.Println(v35, v35, fi34)
                                var v46 uint64 = 18446744073709551615
                                fmt.Println(v3, v1, v46)
                                _ = v46
                            }
                        }
                        v47 := GenS4{F0: 0, F1: 65535}
                        _ = v35
                        _ = v47
                    } else {
                        var v48 uint64 = 18446744073709551615
                        for fi49 := 0; fi49 < 11; fi49++ {
                            _ = len(v2)
                            var v50 uint8 = 42
                            var v51 uint8 = (v50 * (v50 % ((1) | 1)))
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        v52 := GenS1{F0: 33, F1: 42, F2: ("\n\t" + (("" + ("foo bar" + (("hello" + "hello") + "hello"))) + "foo bar"))}
                        _ = v48
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 19; fi53++ {
                        var v54 int8 = -112
                        if false {
                            var v55 float32 = float32(-2.25)
                            switch "" {
                            case "\n\t":
                                var v56 uint = 4294967295
                                _ = v56
                            default:
                                var v57 string = "foo bar"
                                _ = v57
                            }
                            switch -2147483648 {
                            case -2147483648:
                                if true {
                                    switch "a" {
                                    case "":
                                        v3 -= v3
                                    case "foo bar":
                                        fmt.Println(v3)
                                    case "hello":
                                        v3 = genFn0()
                                    default:
                                        var v58 int8 = -11
                                        _ = v58
                                    }
                                }
                            case 9223372036854775807:
                                var v59 int = (v1 * (fi53 << ((402) & 63)))
                                _ = v59
                            case -264:
                                var v60 int8 = -111
                                _ = v60
                            default:
                                if false {
                                    if ((v3 > v3) && (382 != -1)) {
                                        var v61 uint = 270
                                        v62 := GenS2{F0: ("\n\t" + "hello")}
                                        v54 = 1
                                        _ = v61
                                        _ = v62
                                    }
                                }
                            }
                            _ = v55
                        } else {
                            v54 = v54
                            var v63 int8 = v54
                            v64 := GenS3{F0: v54}
                            _ = v63
                            _ = v64
                        }
                        var v65 int16 = 990
                        if (true && true) {
                            if (-311 < -526) {
                                v1 = (v1 >> ((255) & 63))
                            }
                            if len(v2) > 0 {
                                v2[0] = (v1 & (v1 - (fi34 & int(v54))))
                            }
                        }
                        _ = fi53
                        _ = v54
                        _ = v65
                    }
                } else {
                    _ = len(v2)
                }
            default:
                fmt.Println(v1, v3)
            }
            _ = fi34
        }
    case 32768:
        v1 -= (v1 + (v1 << ((42) & 63)))
    }
    if false {
        var v66 int32 = -2147483648
        _ = v66
    } else {
        switch "\n\t" {
        case "a":
            if true {
                v1 = (v1 << ((0) & 63))
                fmt.Println(v3, v3)
                v3 += v3
            }
        case "hello":
            for fi67 := 0; fi67 < 15; fi67++ {
                v1 += v1
                for fi68 := 0; fi68 < 13; fi68++ {
                    v69 := GenS6{F0: ((!false) || (false && (434 < 255))), F1: 4294967295, F2: true}
                    v3 -= 65535
                    v2 = append(v2, (fi68 % (((^(fi67))) | 1)))
                    if (!(!(!(0 != -1)))) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi68
                    _ = v69
                }
                _ = fi67
            }
        case "":
            var v70 uint8 = 123
            _ = v70
        default:
            var v71 bool = (false || (!false))
            _ = v71
        }
        v72 := GenS6{F0: false, F1: 265, F2: ((!(v1 > 244)) && false)}
        v73 := GenS5{F0: 2, F1: 32767}
        _ = v72
        _ = v73
    }
    v1 = v1
    v3 -= v3
    v1 = -930
    v74 := GenS7{F0: genFn0(), F1: (^(v1))}
    v75 := GenS8{F0: false, F1: -128, F2: 0}
    v2 = append(v2, (-(v1)))
    v76 := GenS9{F0: 42, F1: -736}
    var v77 float32 = float32(((1e10 * (-2.25 + float64(3.14159))) + (3.14159 + 1.5)))
    v1 -= v1
    var v78 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
