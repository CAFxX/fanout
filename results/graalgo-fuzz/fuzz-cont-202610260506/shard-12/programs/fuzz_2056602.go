package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string; F1 uint16; F2 bool }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 int32, p1 string, p2 int32) int8 {
    v4 := GenS1{F0: (p2 / ((p0) | 1)), F1: 1}
    p2 = (p2 >> ((4294967295) & 63))
    var v5 int32 = (p2 ^ -2147483648)
    fmt.Println(v5, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1() uint {
    var v6 float32 = (1.5 / (0.1 * 3.14159))
    switch ("\n\t" + "a") {
    case "\n\t":
        fmt.Println(v6)
    case "":
        v7 := GenS1{F0: 805, F1: 997}
        _ = v7
    case "hello":
        if (253 > 128) {
            var v8 string = ("hello" + (("hello" + ("foo bar" + ("foo bar" + ""))) + ("hello" + ("" + ("\n\t" + "")))))
            fmt.Println(v8, v6)
            _ = v8
        } else {
            v9 := []int{-815, -2147483648, -2147483648}
            _ = v9
        }
    }
    var v10 uint64 = 18446744073709551615
    _ = v10
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 uint16 = 182
    for fi11 := 0; fi11 < 3; fi11++ {
        v1 += v1
        _ = fi11
    }
    fmt.Println(v1, v3)
    for fi12 := 0; fi12 < 4; fi12++ {
        delete(v2, "a")
        v1 += (v1 % (((v1 - (v1 % (((v1 * v1)) | 1)))) | 1))
        _ = fi12
    }
    fmt.Println(v3)
    var v13 int8 = 42
    if true {
        if (4294967295 > 255) {
            fmt.Println(v3, v3)
            fmt.Println(v1, v13, v3)
        } else {
            v2["foo bar"] = 2147483648
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{fi14, 2147483647, fi14}
        switch ("foo bar" + ("foo bar" + "")) {
        case "a":
            if true {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint64 = 2
                    v18 := GenS1{F0: 923, F1: genFn1()}
                    var v19 int32 = -1
                    v17 = v17
                    var v20 int8 = 118
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                delete(v2, "\n\t")
                switch "a" {
                case "":
                    v21 := GenS2{F0: 2, F1: false, F2: (fi14 + (fi14 % ((fi14) | 1)))}
                    _ = v21
                default:
                    if (719 != 798) {
                        fmt.Println(v3)
                        var v22 uint = 4294967295
                        if true {
                            switch -2147483648 {
                            case -1:
                                switch 125 {
                                case -2147483648:
                                    v3 += (v3 ^ v3)
                                case 128:
                                    if true {
                                        var v23 float64 = -2.25
                                        if true {
                                            v23 = 1e10
                                            for fi24 := 0; fi24 < 9; fi24++ {
                                                switch "a" {
                                                case "a":
                                                    var v25 int64 = (v1 << ((v22) & 63))
                                                    _ = v25
                                                default:
                                                    _ = len(v2)
                                                }
                                                _ = fi24
                                            }
                                            v13 = (v13 << (((v22 ^ 768)) & 63))
                                        } else {
                                            for fi26 := 0; fi26 < 3; fi26++ {
                                                var v27 uint8 = 61
                                                switch (fi14 | (fi14 << ((4294967295) & 63))) {
                                                case 81:
                                                    v22 -= genFn1()
                                                case 42:
                                                    if false {
                                                        if ((3.14159 <= ((-2.25 - (0.1 * (1e10 - (float32(v23) / -2.25)))) / 1.5)) && (42 == 1)) {
                                                            var v28 int16 = 32767
                                                            _ = v28
                                                        }
                                                        v2["\n\t"] = 4294967295
                                                    }
                                                }
                                                fmt.Println(fi14)
                                                v29 := v2["hello"]
                                                if true {
                                                    var v30 int32 = int32(v13)
                                                    _ = v30
                                                }
                                                if (float64(1.5) <= 3.14159) {
                                                    switch fi26 {
                                                    case 42:
                                                        v31 := GenS1{F0: -665, F1: uint(v29)}
                                                        _ = v31
                                                    case -673:
                                                        _ = len(v15)
                                                    case -2147483648:
                                                        fmt.Println(v23, v13)
                                                    default:
                                                        var v32 int64 = v1
                                                        v33 := GenS3{F0: (v27 | v27), F1: ""}
                                                        v3 = 1
                                                        fmt.Println(v32)
                                                        var v34 uint = uint(v29)
                                                        v22 = genFn1()
                                                        _ = v32
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                }
                                                _ = fi26
                                                _ = v27
                                                _ = v29
                                            }
                                        }
                                        fmt.Println(v13, fi14)
                                        _ = v23
                                    } else {
                                        v3 += (^(v3))
                                        for fi35 := 0; fi35 < 5; fi35++ {
                                            v36 := GenS3{F0: 36, F1: ("foo bar" + "")}
                                            var v37 uint64 = 2
                                            var v38 uint64 = v37
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                case -1:
                                    if false {
                                        v22 -= genFn1()
                                        fmt.Println(v22, v1, v13)
                                        v13 = v13
                                    } else {
                                        v3 = v3
                                        for fi39 := 0; fi39 < 16; fi39++ {
                                            var v40 int32 = 1
                                            _ = fi39
                                            _ = v40
                                        }
                                        var v41 uint16 = 921
                                        _ = v41
                                    }
                                default:
                                    if false {
                                        if false {
                                            for fi42 := 0; fi42 < 8; fi42++ {
                                                v43 := GenS2{F0: 128, F1: (!false), F2: (fi42 - (fi42 % ((fi14) | 1)))}
                                                fmt.Println(v1)
                                                v13 += (v13 << ((2) & 63))
                                                _ = fi42
                                                _ = v43
                                            }
                                            switch (fi14 % (((fi14 ^ 9223372036854775807)) | 1)) {
                                            case -752:
                                                var v44 bool = true
                                                _ = v44
                                            default:
                                                var v45 int8 = genFn0(-600, "", -2147483648)
                                                _ = v45
                                            }
                                            v1 = v1
                                        } else {
                                            var v46 int64 = v1
                                            for fi47 := 0; fi47 < 10; fi47++ {
                                                switch 0 {
                                                case 4294967295:
                                                    fmt.Println(fi14, v3)
                                                case -2147483648:
                                                    v3 = v3
                                                default:
                                                    var v48 uint16 = v3
                                                    _ = v48
                                                }
                                                v15 = append(v15, (fi47 << ((v22) & 63)))
                                                var v49 uint32 = 2147483648
                                                var v50 uint32 = v49
                                                _ = fi47
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = v46
                                        }
                                    }
                                }
                            case 965:
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    var v52 int = (fi14 & fi14)
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        var v53 uint16 = v3
                                        _ = v53
                                    case "a":
                                        delete(v2, ((("\n\t" + "hello") + "foo bar") + "foo bar"))
                                    case "foo bar":
                                        v54 := GenS4{F0: "", F1: (v3 & (v3 & v3)), F2: false}
                                        _ = v54
                                    default:
                                        v52 += -752
                                    }
                                    switch fi51 {
                                    case 32767:
                                        if true {
                                            var v55 uint8 = 77
                                            switch ("hello" + "hello") {
                                            case "":
                                                v56 := GenS4{F0: "", F1: 294, F2: (1 == 4294967295)}
                                                _ = v56
                                            case "foo bar":
                                                fmt.Println(fi51, v22, v22)
                                                var v57 float32 = 0.1
                                                _ = v57
                                            }
                                            fmt.Println(fi14)
                                            _ = v55
                                        } else {
                                            _ = len(v15)
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                            var v58 uint = (v22 << ((v22) & 63))
                                            if false {
                                                switch (v52 << (((v58 << (((v58 - 18446744073709551615)) & 63))) & 63)) {
                                                case 128:
                                                    if (2 <= v3) {
                                                        v1 -= (v1 - v1)
                                                        switch "a" {
                                                        case "\n\t":
                                                            v13 -= (v13 % (((^(v13))) | 1))
                                                        }
                                                        var v59 int64 = (^(v1))
                                                        v60 := GenS5{F0: -1, F1: (!false)}
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                case 9223372036854775807:
                                                    var v61 int = 0
                                                    _ = v61
                                                default:
                                                    switch (^(fi14)) {
                                                    case 32768:
                                                        v62 := GenS1{F0: 0, F1: (v22 % (((v22 | (v58 ^ genFn1()))) | 1))}
                                                        _ = v62
                                                    case 55:
                                                        v58 += 42
                                                    case 32767:
                                                        var v63 int16 = 32767
                                                        _ = v63
                                                    default:
                                                        var v64 int8 = (v13 % (((v13 - (v13 & v13))) | 1))
                                                        _ = v64
                                                    }
                                                }
                                                if false {
                                                    if len(v15) > 0 {
                                                        v15[0] = 127
                                                    }
                                                }
                                                if false {
                                                    _ = len(v2)
                                                    v2[("" + "hello")] = 117
                                                }
                                            } else {
                                                _ = len(v15)
                                                var v65 float64 = 0.0
                                                _ = v65
                                            }
                                            _ = v58
                                        }
                                    case -828:
                                        if false {
                                            v66 := GenS4{F0: "", F1: 65535, F2: ((v3 >> (((v22 << ((230) & 63))) & 63)) >= v3)}
                                            v2["a"] = 2147483647
                                            _ = v66
                                        }
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v67 uint32 = 0
                                        _ = v67
                                    case "hello":
                                        v68 := GenS2{F0: 2, F1: false, F2: fi51}
                                        _ = v68
                                    case "\n\t":
                                        var v69 uint8 = 0
                                        _ = v69
                                    }
                                    _ = fi51
                                    _ = v52
                                }
                            case -541:
                                var v70 string = ("foo bar" + "hello")
                                _ = v70
                            }
                            var v71 float64 = float64((0.1 + 3.14159))
                            v72 := GenS3{F0: 0, F1: "foo bar"}
                            _ = v71
                            _ = v72
                        } else {
                            v1 += (v1 * -684)
                            if (true && false) {
                                switch "foo bar" {
                                case "\n\t":
                                    switch "" {
                                    case "a":
                                        _ = len(v15)
                                    }
                                default:
                                    for fi73 := 0; fi73 < 9; fi73++ {
                                        fmt.Println(fi73, v1, v22)
                                        var v74 int64 = v1
                                        v74 += (^(v74))
                                        fmt.Println(v1, fi73)
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        var v75 float32 = 0.1
                                        if len(v15) > 0 {
                                            v15[0] = fi73
                                        }
                                        var v76 float32 = v75
                                        var v77 float32 = ((0.1 / (v75 - (1e10 / -2.25))) * (((((v75 / 1e10) / (v76 + v75)) + v75) - ((v76 - v76) * (1e10 / 1.0))) - v76))
                                        v15 = append(v15, (-(fi14)))
                                        var v78 uint8 = 128
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            } else {
                                for fi79 := 0; fi79 < 14; fi79++ {
                                    fmt.Println(v13, fi79)
                                    _ = fi79
                                }
                            }
                        }
                        _ = v22
                    }
                }
            }
        case "\n\t":
            var v80 int64 = 1
            _ = v80
        case "hello":
            var v81 float32 = -2.25
            _ = v81
        }
        _ = fi14
        _ = v15
    }
    v1 = v1
    var v82 int16 = 1000
    var v83 uint64 = 0
    fmt.Println(v83, v83)
    var v84 uint16 = v3
    if (!true) {
        for fi85 := 0; fi85 < 1; fi85++ {
            fmt.Println(fi85, v83)
            delete(v2, ("a" + (("a" + ("foo bar" + "a")) + ((("a" + ("hello" + "hello")) + "a") + "foo bar"))))
            fmt.Println(v82, fi85)
            _ = fi85
        }
        v3 = (v3 * 42)
        var v86 uint32 = 336
        _ = v86
    } else {
        v3 -= v84
        var v87 uint8 = 42
        _ = v87
    }
    var v88 uint32 = 65535
    var v89 int16 = v82
    var v90 uint64 = (v83 ^ 42)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v82
    _ = v83
    _ = v84
    _ = v88
    _ = v89
    _ = v90
}
