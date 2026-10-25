package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int8; F2 int8 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int16; F1 int16; F2 int32 }
type GenS8 struct { F0 bool; F1 int16; F2 bool }
type GenS9 struct { F0 uint64 }
func genFn0() int {
    var v4 uint8 = 42
    if (0 == 1) {
        v5 := GenS1{F0: -1}
        for fi6 := 0; fi6 < 15; fi6++ {
            switch "\n\t" {
            case "hello":
                var v7 int = -2147483648
                _ = v7
            case "\n\t":
                if (false || true) {
                    v4 += v4
                    var v8 int = -9223372036854775808
                    _ = v8
                } else {
                    v9 := GenS1{F0: -9223372036854775808}
                    fmt.Println(v4, fi6)
                    _ = v9
                }
            case "foo bar":
                v10 := []int{(fi6 * fi6), fi6, 127, (fi6 ^ (fi6 / ((1) | 1)))}
                _ = v10
            }
            var v11 uint = 0
            _ = fi6
            _ = v11
        }
        var v12 int32 = -1
        _ = v12
        _ = v5
    }
    _ = v4
    return 32768
}
func genFn1() int64 {
    var v13 float64 = ((1e10 * 1.5) - 1e10)
    var v14 string = "hello"
    v15 := []int{128, -9223372036854775808}
    v13 += 1e10
    fmt.Println(v13, v14)
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []string{""}
    var v2 float32 = -2.25
    var v3 bool = (1 > 2)
    if true {
        v2 += float32(3.14159)
        var v16 int32 = 2147483647
        v17 := GenS1{F0: 1}
        _ = v16
        _ = v17
    }
    var v18 int8 = -1
    fmt.Println(v3)
    v2 -= 0.0
    for fi19 := 0; fi19 < 4; fi19++ {
        switch "\n\t" {
        case "\n\t":
            if ((!((v3 && v3) || (true || (2 <= 4294967295)))) && ("" != "a")) {
                var v20 bool = (3.14159 == 1e10)
                var v21 int16 = 391
                _ = v20
                _ = v21
            }
        case "hello":
            var v22 int16 = 0
            _ = v22
        case "":
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    v24 := GenS1{F0: genFn1()}
                    var v25 string = ""
                    var v26 float32 = ((-2.25 - float32(1e10)) * -2.25)
                    v26 = v2
                    v2 -= v2
                    switch v25 {
                    case "hello":
                        if v3 {
                            _ = len(v1)
                            if v3 {
                                fmt.Println(v3, fi23)
                                v2 -= v26
                            } else {
                                if v3 {
                                    var v27 int32 = 2147483647
                                    v1 = append(v1, v25)
                                    _ = v27
                                }
                                var v28 int8 = v18
                                _ = v28
                            }
                        } else {
                            var v29 float64 = ((3.14159 + 1e10) * 3.14159)
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            var v30 int64 = genFn1()
                            if len(v1) > 0 {
                                v1[0] = v25
                            }
                            _ = v29
                            _ = v30
                        }
                    case "a":
                        var v31 int16 = 351
                        _ = v31
                    case "\n\t":
                        var v32 uint = 255
                        _ = v32
                    default:
                        v3 = false
                    }
                    _ = v24
                    _ = v25
                    _ = v26
                }
                if false {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v1 = append(v1, "a")
                }
                _ = fi23
            }
        default:
            var v33 uint = 42
            _ = v33
        }
        _ = fi19
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS2{F0: (!false), F1: 29, F2: v3}
        _ = fi34
        _ = v35
    }
    if true {
        var v36 int16 = 1000
        switch "a" {
        case "foo bar":
            v36 = (v36 << ((4294967295) & 63))
        case "\n\t":
            if ((!(float64((v2 + 0.1)) == 3.14159)) && false) {
                fmt.Println(v2)
            }
        default:
            v37 := GenS3{F0: int32(v36)}
            _ = v37
        }
        _ = v36
    } else {
        v2 = 1e10
        fmt.Println(v2)
        if (!(((829 <= 1) && true) && ((float32(0.0) / (((0.0 / v2) + 1e10) + 0.0)) > v2))) {
            if true {
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 uint64 = 18446744073709551615
                    v3 = (!(v3 && (!v3)))
                    _ = fi38
                    _ = v39
                }
                v3 = false
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                v18 = (v18 & v18)
                for fi41 := 0; fi41 < 8; fi41++ {
                    v1 = append(v1, "a")
                    v1 = append(v1, ((("\n\t" + "a") + (("" + "a") + ("a" + "a"))) + ("hello" + "\n\t")))
                    _ = len(v1)
                    v3 = ((42 <= 9223372036854775807) && (52 > 128))
                    _ = fi41
                }
                switch "hello" {
                case "foo bar":
                    var v42 bool = (("hello" + "a") != "a")
                    _ = v42
                case "\n\t":
                    _ = len(v1)
                case "hello":
                    _ = len(v1)
                }
                _ = fi40
            }
            var v43 int32 = -197
            _ = v43
        } else {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
    }
    v44 := GenS3{F0: -1}
    _ = len(v1)
    if (42 > 65535) {
        if true {
            switch "foo bar" {
            case "":
                var v45 int64 = 42
                _ = v45
            case "\n\t":
                for fi46 := 0; fi46 < 15; fi46++ {
                    if ((255 >= 255) || (("foo bar" + "hello") >= "")) {
                        if true {
                            fmt.Println(v3)
                            v18 = v18
                            v47 := GenS4{F0: 0}
                            _ = v47
                        } else {
                            fmt.Println(fi46)
                        }
                    } else {
                        v2 = v2
                        if (fi46 < genFn0()) {
                            var v48 int = 32768
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            _ = v48
                        }
                        v2 = v2
                    }
                    v3 = v3
                    var v49 uint32 = 0
                    _ = fi46
                    _ = v49
                }
            }
            v2 += (v2 / (v2 / 1.0))
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v50 := GenS3{F0: -1}
            _ = v50
        }
        v51 := GenS1{F0: -176}
        for fi52 := 0; fi52 < 3; fi52++ {
            v53 := GenS2{F0: (!false), F1: 160, F2: (v18 != (v18 - -128))}
            var v54 uint32 = 0
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = v51
    } else {
        v18 = -1
        var v55 int64 = 1
        _ = v55
    }
    var v56 uint16 = 2
    if v3 {
        fmt.Println(v18, v2, v56)
        v2 += v2
    } else {
        fmt.Println(v18, v2)
        for fi57 := 0; fi57 < 7; fi57++ {
            v1 = append(v1, ("\n\t" + "a"))
            _ = fi57
        }
        v1 = append(v1, "")
    }
    var v58 uint = 65535
    if v3 {
        if true {
            switch ("foo bar" + ((("hello" + "") + "foo bar") + ("hello" + (("hello" + "hello") + "")))) {
            case "foo bar":
                v58 -= uint(v56)
            case "\n\t":
                v59 := GenS4{F0: 367}
                _ = v59
            }
            for fi60 := 0; fi60 < 9; fi60++ {
                if (v3 || false) {
                    v61 := GenS5{F0: 1, F1: -41, F2: int8(fi60)}
                    switch (fi60 | fi60) {
                    case 127:
                        for fi62 := 0; fi62 < 13; fi62++ {
                            if v3 {
                                fmt.Println(v18, v3, v18)
                                v1 = append(v1, "foo bar")
                                switch "\n\t" {
                                case "a":
                                    switch (("foo bar" + "hello") + "\n\t") {
                                    case "":
                                        v58 = v58
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                case "foo bar":
                                    _ = len(v1)
                                case "hello":
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    if (0.1 > 1.5) {
                                        fmt.Println(fi60)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "\n\t":
                                            if (3.14159 < (1e10 - ((float64((float32(1.5) * (v2 * v2))) + float64(v2)) + (float64((-2.25 / 0.1)) + 1.5)))) {
                                                var v63 int32 = -404
                                                v64 := GenS6{F0: int64(v56), F1: 18446744073709551615}
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v65 := GenS7{F0: -749, F1: 475, F2: 819}
                                                v2 += v2
                                                fmt.Println(v18, v2, v2)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v65
                                            }
                                        default:
                                            var v66 uint64 = 400
                                            _ = v66
                                        }
                                    }
                                    if ((v2 / ((3.14159 - v2) - v2)) <= v2) {
                                        fmt.Println(v58)
                                    }
                                default:
                                    var v67 uint64 = 18446744073709551615
                                    _ = v67
                                }
                            }
                            if v3 {
                                v68 := GenS7{F0: -32768, F1: -1, F2: 1}
                                switch int(v18) {
                                case -1:
                                    if (!(4294967295 != 1)) {
                                        var v69 uint8 = 170
                                        if v3 {
                                            v2 = v2
                                        }
                                        _ = v69
                                    } else {
                                        switch fi62 {
                                        case 127:
                                            if v3 {
                                                if (v3 && v3) {
                                                    var v70 int8 = v18
                                                    _ = v70
                                                }
                                                v18 += (^(v18))
                                                v18 = (v18 - (v18 / (((-(v18))) | 1)))
                                            } else {
                                                var v71 string = ("\n\t" + ("foo bar" + ""))
                                                if len(v1) > 0 {
                                                    v1[0] = v71
                                                }
                                                _ = v71
                                            }
                                        case 42:
                                            v72 := GenS3{F0: 1}
                                            _ = v72
                                        case 128:
                                            var v73 uint16 = v56
                                            _ = v73
                                        }
                                        v74 := GenS6{F0: -9223372036854775808, F1: 332}
                                        _ = v74
                                    }
                                case 2147483647:
                                    switch "" {
                                    case "foo bar":
                                        v75 := GenS4{F0: 42}
                                        _ = v75
                                    case "a":
                                        fmt.Println(v18)
                                    }
                                case 4294967295:
                                    v76 := GenS8{F0: ((255 == 255) && v3), F1: -32768, F2: ((0.1 - 0.1) > (v2 / (float32(1.5) - v2)))}
                                    _ = v76
                                default:
                                    _ = len(v1)
                                }
                                var v77 float64 = -2.25
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v3, v18, v18)
                                fmt.Println(fi60, fi60, fi60)
                                _ = v68
                                _ = v77
                            }
                            v78 := GenS2{F0: (v2 < v2), F1: 248, F2: (((false && (v3 || (!((fi62 & (fi62 - fi62)) >= -597)))) && true) || (!v3))}
                            _ = fi62
                            _ = v78
                        }
                    }
                    v1 = append(v1, (("hello" + "a") + "a"))
                    _ = v61
                } else {
                    if (1 != 611) {
                        for fi79 := 0; fi79 < 20; fi79++ {
                            switch (("hello" + "a") + "hello") {
                            case "foo bar":
                                var v80 bool = (true || false)
                                _ = v80
                            case "hello":
                                v2 -= v2
                                v58 += (v58 ^ (v58 - 0))
                                var v81 int = (-(fi79))
                                _ = v81
                            }
                            v82 := GenS9{F0: 2}
                            _ = fi79
                            _ = v82
                        }
                        var v83 int = 32768
                        if (genFn1() >= 0) {
                            v3 = v3
                            var v84 int32 = 2147483647
                            _ = v84
                        } else {
                            v58 = 42
                        }
                        _ = v83
                    }
                    v2 = ((1e10 + (v2 / 0.1)) * v2)
                }
                v56 -= v56
                fmt.Println(fi60, v18, v58)
                _ = fi60
            }
            switch int(v58) {
            case 9223372036854775807:
                var v85 int16 = 1
                _ = v85
            case 156:
                v86 := GenS8{F0: (false || (v58 < 0)), F1: 0, F2: v3}
                _ = v86
            case 32768:
                if v3 {
                    var v87 uint64 = 444
                    fmt.Println(v87, v18, v56)
                    _ = len(v1)
                    _ = v87
                }
            default:
                if v3 {
                    switch genFn0() {
                    case -9223372036854775808:
                        _ = len(v1)
                    case 4294967295:
                        if (v3 || (3.14159 <= v2)) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    default:
                        if v3 {
                            var v88 uint64 = 565
                            if true {
                                v3 = v3
                            }
                            _ = v88
                        }
                    }
                }
            }
        } else {
            if false {
                switch 311 {
                case -783:
                    fmt.Println(v2, v58, v56)
                default:
                    var v89 string = "a"
                    _ = v89
                }
            } else {
                for fi90 := 0; fi90 < 17; fi90++ {
                    if v3 {
                        v56 = v56
                        var v91 uint16 = v56
                        fmt.Println(v2, v58, v91)
                        _ = v91
                    } else {
                        var v92 uint = (v58 - 18446744073709551615)
                        if (42 == genFn1()) {
                            if len(v1) > 0 {
                                v1[0] = (("foo bar" + "") + (("foo bar" + "foo bar") + "\n\t"))
                            }
                            var v93 uint16 = v56
                            _ = v93
                        } else {
                            var v94 bool = (v18 < (v18 + 0))
                            v56 = v56
                            if (1 != 1) {
                                v18 = (-(v18))
                                for fi95 := 0; fi95 < 3; fi95++ {
                                    fmt.Println(v18)
                                    v92 -= uint(v56)
                                    _ = fi95
                                }
                            }
                            _ = v94
                        }
                        var v96 int = -353
                        _ = v92
                        _ = v96
                    }
                    _ = fi90
                }
                var v97 bool = v3
                if (!false) {
                    var v98 uint16 = (^(v56))
                    _ = v98
                }
                _ = v97
            }
            v58 = 4294967295
            if (0 == -266) {
                var v99 int = 128
                var v100 float64 = (3.14159 - float64(float32(-2.25)))
                _ = v100
                _ = v99
            } else {
                var v101 int64 = 42
                for fi102 := 0; fi102 < 15; fi102++ {
                    fmt.Println(v56)
                    _ = fi102
                }
                v56 -= v56
                _ = v101
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v44
    _ = v56
    _ = v58
}
