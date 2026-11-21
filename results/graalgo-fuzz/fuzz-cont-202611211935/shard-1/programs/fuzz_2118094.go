package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int; F2 int32 }
type GenS7 struct { F0 uint32; F1 int; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 int64; F2 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint16; F1 int }
func genFn0(p0 int32) uint {
    switch -2147483648 {
    case 32767:
        p0 -= 683
    case 9223372036854775807:
        v5 := []int{-165}
        _ = v5
    case -1:
        var v6 float64 = -2.25
        _ = v6
    }
    fmt.Println(p0)
    var v7 int16 = -32768
    var v8 uint32 = 79
    switch ("foo bar" + ("hello" + "")) {
    case "":
        var v9 uint = 18446744073709551615
        _ = v9
    case "hello":
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 bool = (false || true)
            v12 := GenS1{F0: fi10, F1: 130}
            var v13 uint32 = (v8 | (v8 + (v8 & (^(v8)))))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"", "a"}
    v3 := []uint32{219, 1, 283, 0}
    v4 := make(map[string]int)
    var v14 int = 42
    var v15 int64 = -217
    v16 := GenS2{F0: 35, F1: -286}
    var v17 uint16 = 466
    v18 := GenS3{F0: v15, F1: -150}
    v17 -= (^(v17))
    fmt.Println(v15)
    for fi19 := 0; fi19 < 6; fi19++ {
        if false {
            v1 += v1
            var v20 float32 = float32(1e10)
            _ = v20
        } else {
            v21 := v4["\n\t"]
            v17 = (v17 & (v17 - v17))
            _ = v21
        }
        v17 -= (^(v17))
        _ = fi19
    }
    if (false || false) {
        v22 := GenS4{F0: 42, F1: 1}
        v17 = (v17 + (v17 - v17))
        _ = v22
    } else {
        v23 := GenS5{F0: "foo bar"}
        v24 := GenS1{F0: v14, F1: 188}
        _ = v23
        _ = v24
    }
    if ((!(((v17 > 1) || (((1 == v17) || (-98 != -128)) && false)) || ((v14 ^ v14) == -130))) && true) {
        var v25 float32 = ((0.1 / 0.1) - ((0.1 + 0.0) / 1e10))
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 float32 = -2.25
                var v29 uint32 = 2147483648
                _ = fi27
                _ = v28
                _ = v29
            }
            if ((v1 ^ v1) < int16(v15)) {
                v14 = v14
            }
            _ = fi26
        }
    } else {
        if (1 == 259) {
            switch "\n\t" {
            case "\n\t":
                v30 := GenS6{F0: 42, F1: v14, F2: -1}
                _ = v30
            default:
                var v31 int8 = 100
                _ = v31
            }
            v32 := v4[(("foo bar" + "") + ("" + ""))]
            _ = v32
        } else {
            switch ("foo bar" + "foo bar") {
            case "a":
                var v33 float32 = (1.5 * 0.0)
                _ = v33
            case "":
                switch "\n\t" {
                case "":
                    if false {
                        v15 = v15
                    }
                }
            case "foo bar":
                for fi34 := 0; fi34 < 20; fi34++ {
                    if (!((0.0 / 3.14159) >= 0.1)) {
                        v15 += 1
                    } else {
                        v35 := v4[""]
                        var v36 bool = false
                        _ = v35
                        _ = v36
                    }
                    v17 = 255
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    _ = fi34
                }
            default:
                if (18446744073709551615 >= 67) {
                    switch (v14 >> ((genFn0(252)) & 63)) {
                    case -533:
                        fmt.Println(v15)
                    case -676:
                        fmt.Println(v15, v1)
                    case 9223372036854775807:
                        if (2 > 128) {
                            for fi37 := 0; fi37 < 2; fi37++ {
                                v3 = append(v3, 4294967295)
                                _ = fi37
                            }
                            var v38 bool = (1e10 != float64(1.5))
                            if (v1 != (v1 - int16(v17))) {
                                v39 := GenS7{F0: 65, F1: v14, F2: v14}
                                _ = v39
                            }
                            _ = v38
                        } else {
                            switch v14 {
                            case -1:
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    fmt.Println(fi40, v14)
                                    fmt.Println(v17, v17, v1)
                                    if true {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v15 += (v15 / (((^(v15))) | 1))
                                        case "hello":
                                            fmt.Println(v14, v17)
                                        case "":
                                            switch "" {
                                            case "\n\t":
                                                v14 = -243
                                            case "hello":
                                                v41 := GenS1{F0: (v14 | -526), F1: 1}
                                                _ = v41
                                            case "":
                                                switch int(v17) {
                                                case 32768:
                                                    v42 := GenS6{F0: 0, F1: 412, F2: -259}
                                                    _ = v42
                                                default:
                                                    switch 32767 {
                                                    case 2147483647:
                                                        for fi43 := 0; fi43 < 4; fi43++ {
                                                            switch ("foo bar" + (("\n\t" + ("" + "\n\t")) + "a")) {
                                                            case "\n\t":
                                                                if false {
                                                                    v44 := GenS8{F0: 15}
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v45 uint = genFn0(-2147483648)
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                            case "hello":
                                                                v46 := v4["a"]
                                                                _ = v46
                                                            case "a":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 42
                                                            }
                                                            if ((fi43 ^ fi43) <= 892) {
                                                                delete(v4, "\n\t")
                                                                _ = len(v4)
                                                            } else {
                                                                var v47 bool = ((^(v1)) == 42)
                                                                _ = v47
                                                            }
                                                            _ = fi43
                                                        }
                                                    default:
                                                        for fi48 := 0; fi48 < 3; fi48++ {
                                                            _ = len(v2)
                                                            v15 -= v15
                                                            fmt.Println(v15, v17, v14)
                                                            _ = fi48
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            if true {
                                                var v49 uint32 = 1
                                                v1 += (^(v1))
                                                _ = v49
                                            }
                                        }
                                    } else {
                                        v15 = -9223372036854775808
                                        delete(v4, "")
                                        v17 = 544
                                    }
                                    _ = fi40
                                }
                            case 772:
                                _ = len(v3)
                            case 128:
                                v14 = (v14 << ((342) & 63))
                            default:
                                fmt.Println(v15, v17)
                            }
                        }
                    default:
                        var v50 uint16 = (v17 | v17)
                        _ = v50
                    }
                    if (147 < 217) {
                        var v51 uint16 = v17
                        v51 = 628
                        if true {
                            v1 -= v1
                            for fi52 := 0; fi52 < 11; fi52++ {
                                if (!true) {
                                    for fi53 := 0; fi53 < 9; fi53++ {
                                        v54 := GenS5{F0: ("a" + "")}
                                        fmt.Println(v1, fi52, v15)
                                        var v55 float64 = float64(float32(-2.25))
                                        var v56 string = ("a" + "")
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    v51 -= v17
                                }
                                v15 -= (v15 << ((2) & 63))
                                _ = fi52
                            }
                            for fi57 := 0; fi57 < 19; fi57++ {
                                v58 := GenS9{F0: 1, F1: v15, F2: v15}
                                fmt.Println(v14, v14)
                                switch ("foo bar" + "") {
                                case "a":
                                    v59 := GenS2{F0: 69, F1: int32(v1)}
                                    _ = v59
                                case "foo bar":
                                    if true {
                                        delete(v4, "a")
                                        delete(v4, ("a" + "foo bar"))
                                    } else {
                                        v3 = append(v3, 4294967295)
                                        if ((v1 - (v1 | (v1 + v1))) != v1) {
                                            switch 127 {
                                            case 940:
                                                _ = len(v4)
                                            }
                                            switch (fi57 | -590) {
                                            case -204:
                                                for fi60 := 0; fi60 < 2; fi60++ {
                                                    var v61 float64 = -2.25
                                                    _ = len(v3)
                                                    _ = fi60
                                                    _ = v61
                                                }
                                            }
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                default:
                                    switch "a" {
                                    case "":
                                        var v62 int8 = -99
                                        _ = v62
                                    case "hello":
                                        var v63 uint = 751
                                        _ = v63
                                    default:
                                        var v64 uint64 = 596
                                        _ = v64
                                    }
                                }
                                _ = fi57
                                _ = v58
                            }
                        } else {
                            var v65 int64 = v15
                            v2 = append(v2, (("foo bar" + ("foo bar" + "a")) + "hello"))
                            _ = v65
                        }
                        _ = v51
                    }
                    switch v14 {
                    case 127:
                        for fi66 := 0; fi66 < 3; fi66++ {
                            if false {
                                switch fi66 {
                                case 32767:
                                    if true {
                                        if false {
                                            v67 := GenS2{F0: 2, F1: 1}
                                            _ = v67
                                        }
                                        if (v15 == int64(v17)) {
                                            v68 := GenS10{F0: 255}
                                            var v69 int64 = v15
                                            var v70 uint8 = 1
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                        }
                                        for fi71 := 0; fi71 < 14; fi71++ {
                                            fmt.Println(v17, fi66)
                                            _ = fi71
                                        }
                                    }
                                }
                            } else {
                                var v72 uint64 = 4294967295
                                var v73 int = int(v72)
                                _ = v72
                                _ = v73
                            }
                            delete(v4, ("" + "foo bar"))
                            var v74 int16 = -32768
                            _ = fi66
                            _ = v74
                        }
                    }
                }
            }
            if ((0.1 + 1.5) > (-2.25 + (3.14159 - ((-2.25 / 3.14159) / 0.1)))) {
                var v75 uint64 = 2
                _ = v75
            } else {
                v14 = (v14 - (v14 | v14))
                v76 := GenS11{F0: (v17 - v17), F1: v14}
                var v77 uint32 = 0
                _ = v76
                _ = v77
            }
        }
        var v78 float64 = 3.14159
        fmt.Println(v1, v78)
        _ = v78
    }
    v79 := GenS8{F0: 2}
    v17 -= 42
    switch "foo bar" {
    case "a":
        v17 = (v17 + (v17 << ((42) & 63)))
    case "\n\t":
        v17 -= 42
    }
    for fi80 := 0; fi80 < 7; fi80++ {
        v81 := GenS2{F0: 245, F1: int32(v17)}
        var v82 int8 = 1
        var v83 bool = true
        _ = fi80
        _ = v81
        _ = v82
        _ = v83
    }
    if true {
        var v84 int = (-(v14))
        _ = v84
    } else {
        if true {
            delete(v4, "foo bar")
            v1 += 42
        }
        if (!true) {
            var v85 int64 = (v15 % (((v15 | -948)) | 1))
            var v86 string = "a"
            _ = v85
            _ = v86
        } else {
            var v87 float64 = (1e10 / (3.14159 - 1.5))
            _ = v87
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v79
}
