package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 uint64 }
func genFn0() string {
    v4 := GenS1{F0: 20, F1: -509}
    v5 := GenS2{F0: 65535}
    v6 := []int{411, -2147483648, 557, 451}
    _ = v4
    _ = v5
    _ = v6
    return ("hello" + "\n\t")
}
func genFn1(p0 string) int64 {
    if false {
        if false {
            v7 := []int{}
            if false {
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 uint32 = 28
                    _ = fi8
                    _ = v9
                }
                var v10 uint32 = 955
                if (!false) {
                    v10 = (v10 % (((v10 >> ((0) & 63))) | 1))
                } else {
                    var v11 int8 = -1
                    var v12 uint16 = 572
                    v10 = v10
                    _ = v11
                    _ = v12
                }
                _ = v10
            }
            if (false && true) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    if (false || false) {
                        p0 = ((p0 + "hello") + p0)
                        p0 = p0
                        for fi14 := 0; fi14 < 16; fi14++ {
                            if true {
                                if true {
                                    p0 = "a"
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v15 int = fi13
                                    _ = v15
                                }
                                var v16 uint32 = 4294967295
                                _ = v16
                            } else {
                                if (!true) {
                                    p0 = (p0 + p0)
                                    var v17 bool = (false && (!(-2.25 < 0.1)))
                                    switch (fi13 | fi14) {
                                    case -511:
                                        switch "a" {
                                        case "\n\t":
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        case "":
                                            v18 := GenS1{F0: -1, F1: 1}
                                            _ = v18
                                        case "a":
                                            if ("hello" > (p0 + "")) {
                                                v19 := GenS2{F0: 0}
                                                _ = v19
                                            }
                                        }
                                    case -9223372036854775808:
                                        if v17 {
                                            var v20 int8 = -128
                                            v21 := GenS3{F0: p0, F1: 577, F2: p0}
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = v20
                                            _ = v21
                                        }
                                    default:
                                        p0 = p0
                                    }
                                    _ = v17
                                } else {
                                    fmt.Println(fi14, p0, p0)
                                    var v22 uint16 = 737
                                    _ = v22
                                }
                                if ((((0.1 == 0.1) && (true || (255 < 127))) || (2147483647 >= 1)) && true) {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                                fmt.Println(fi14, fi13, p0)
                            }
                            v7 = append(v7, -1)
                            var v23 uint32 = 0
                            _ = fi14
                            _ = v23
                        }
                    }
                    p0 = p0
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 int16 = 908
                        if true {
                            p0 = p0
                            if len(v7) > 0 {
                                v7[0] = (-(fi24))
                            }
                        }
                        if (1 == 4294967295) {
                            fmt.Println(fi13)
                        } else {
                            if len(v7) > 0 {
                                v7[0] = fi24
                            }
                            var v26 bool = false
                            var v27 int8 = 127
                            _ = v26
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi13
                }
                v7 = append(v7, 732)
                p0 = "foo bar"
            }
            _ = v7
        } else {
            if true {
                v28 := GenS1{F0: 116, F1: 2147483647}
                v29 := []int{128, 32768, 2147483647, 4294967295}
                _ = v28
                _ = v29
            } else {
                var v30 float64 = 0.1
                _ = v30
            }
        }
    }
    v31 := GenS4{F0: (255 == 127)}
    fmt.Println(p0)
    var v32 float64 = (1.5 * -2.25)
    _ = p0
    _ = v31
    _ = v32
    return -1
}

func main() {
    v1 := []uint32{573, 733}
    var v2 int32 = -79
    var v3 float32 = 0.1
    v2 = -2147483648
    for fi33 := 0; fi33 < 11; fi33++ {
        _ = len(v1)
        _ = fi33
    }
    var v34 float64 = 3.14159
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(v34, v2, v34)
    case "hello":
        v34 = 3.14159
    case "foo bar":
        var v35 int8 = 16
        _ = v35
    default:
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v36 float32 = (-2.25 * v3)
            _ = v36
        } else {
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 uint = 42
                _ = fi37
                _ = v38
            }
        }
    }
    fmt.Println(v3, v2)
    if (42 <= 4294967295) {
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 string = "foo bar"
            for fi41 := 0; fi41 < 13; fi41++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi39)
                _ = fi41
            }
            if (!(fi39 < 127)) {
                switch fi39 {
                case 783:
                    if (fi39 == (^(fi39))) {
                        if false {
                            if (true || (v2 <= v2)) {
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    switch 32768 {
                                    case 374:
                                        fmt.Println(v34, v2, v40)
                                    default:
                                        if false {
                                            v40 = v40
                                        } else {
                                            var v43 uint32 = 207
                                            switch (fi42 << ((0) & 63)) {
                                            case 32767:
                                                v43 = (v43 >> ((18446744073709551615) & 63))
                                            case 2147483647:
                                                switch 485 {
                                                case -528:
                                                    if true {
                                                        if true {
                                                            v44 := GenS4{F0: true}
                                                            v34 += v34
                                                            _ = v44
                                                        } else {
                                                            if true {
                                                                var v45 uint = 2
                                                                _ = v45
                                                            }
                                                            v3 = (((1.5 + v3) + (v3 * (float32((float64((1.5 / (0.0 - v3))) - 0.1)) - -2.25))) * v3)
                                                        }
                                                        var v46 uint8 = 7
                                                        fmt.Println(v46, v40, v46)
                                                        _ = v46
                                                    } else {
                                                        if ((v34 + (1e10 / v34)) == 1.5) {
                                                            if false {
                                                                fmt.Println(v43, fi39, v40)
                                                                switch fi42 {
                                                                case -2147483648:
                                                                    var v47 int16 = 1
                                                                    _ = v47
                                                                case 0:
                                                                    var v48 int8 = 44
                                                                    _ = v48
                                                                }
                                                                v3 += v3
                                                            } else {
                                                                v3 -= (v3 / float32(v34))
                                                            }
                                                        } else {
                                                            fmt.Println(v40, fi42)
                                                        }
                                                        var v49 uint8 = 42
                                                        _ = v49
                                                    }
                                                    v50 := GenS5{F0: -555, F1: 158}
                                                    var v51 uint64 = 4294967295
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case -9223372036854775808:
                                                var v52 int8 = 42
                                                _ = v52
                                            }
                                            var v53 bool = true
                                            _ = v43
                                            _ = v53
                                        }
                                    }
                                    _ = fi42
                                }
                                for fi54 := 0; fi54 < 2; fi54++ {
                                    var v55 int64 = -100
                                    _ = fi54
                                    _ = v55
                                }
                                for fi56 := 0; fi56 < 5; fi56++ {
                                    v57 := GenS5{F0: -955, F1: 185}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                            }
                        } else {
                            if (206 == 1) {
                                v40 = v40
                                v40 = (("hello" + (v40 + (v40 + v40))) + v40)
                            } else {
                                v3 = (float32(1.5) - ((v3 - 0.0) + ((0.1 * 1.5) + (((((float32(v34) * 0.0) - float32(0.0)) / 1e10) * v3) * (v3 / float32(3.14159))))))
                                v58 := GenS5{F0: 42, F1: 255}
                                var v59 uint64 = 2
                                _ = v58
                                _ = v59
                            }
                            v40 = v40
                            var v60 float32 = 1.5
                            _ = v60
                        }
                    }
                case -329:
                    if false {
                        var v61 uint8 = 8
                        v62 := GenS6{F0: (v2 >> ((18446744073709551615) & 63)), F1: 0}
                        v34 = ((float64(0.1) * v34) - v34)
                        _ = v61
                        _ = v62
                    }
                case -465:
                    switch (fi39 % (((^(fi39))) | 1)) {
                    case 0:
                        v63 := GenS3{F0: (("a" + "foo bar") + "\n\t"), F1: 209, F2: v40}
                        _ = v63
                    }
                default:
                    v1 = append(v1, 2)
                }
                switch fi39 {
                case 32768:
                    v40 = v40
                default:
                    v34 += (v34 - v34)
                }
                if true {
                    for fi64 := 0; fi64 < 6; fi64++ {
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                        _ = fi64
                    }
                    v65 := GenS7{F0: 636, F1: 529, F2: 1}
                    _ = v65
                }
            } else {
                if false {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        switch v40 {
                        case "\n\t":
                            var v67 uint32 = 335
                            _ = v67
                        case "":
                            v34 += ((float64(v3) * v34) + v34)
                        default:
                            var v68 uint8 = 255
                            _ = v68
                        }
                        _ = fi66
                    }
                    if (!true) {
                        if len(v1) > 0 {
                            v1[0] = 595
                        }
                        v34 -= 1e10
                    }
                }
            }
            _ = fi39
            _ = v40
        }
        fmt.Println(v34, v3, v3)
    }
    v69 := GenS8{F0: true, F1: 9223372036854775807}
    switch -2147483648 {
    case 128:
        if (true || false) {
            for fi70 := 0; fi70 < 12; fi70++ {
                v3 += 1.5
                fmt.Println(fi70)
                _ = fi70
            }
        }
    case 269:
        var v71 uint32 = 2
        _ = v71
    case -792:
        var v72 int16 = 32767
        _ = v72
    }
    v73 := GenS8{F0: true, F1: 42}
    var v74 float32 = 0.1
    v74 = 1e10
    switch "hello" {
    case "hello":
        v2 = v2
    case "\n\t":
        v2 = v2
    case "foo bar":
        for fi75 := 0; fi75 < 1; fi75++ {
            switch fi75 {
            case -466:
                if true {
                    var v76 int16 = -291
                    var v77 uint64 = 42
                    _ = v76
                    _ = v77
                } else {
                    v34 = float64((v3 + v74))
                }
            }
            fmt.Println(v34)
            _ = fi75
        }
    }
    v2 = (^(v2))
    switch "hello" {
    case "":
        v78 := GenS8{F0: (!true), F1: -474}
        _ = v78
    case "hello":
        v79 := GenS9{F0: 778}
        _ = v79
    }
    var v80 float64 = (((1.5 / v34) + float64(3.14159)) - v34)
    var v81 uint32 = 889
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v74:", v74)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v69
    _ = v73
    _ = v74
    _ = v80
    _ = v81
}
