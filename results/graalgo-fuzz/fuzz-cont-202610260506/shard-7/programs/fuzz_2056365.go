package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int8; F2 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS5 struct { F0 string; F1 uint64; F2 int }
type GenS6 struct { F0 bool; F1 string; F2 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 string }
func genFn0() uint8 {
    v4 := GenS1{F0: 348, F1: 25}
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(fi5, fi5)
            default:
                var v7 uint8 = 1
                _ = v7
            }
            v8 := GenS2{F0: 1}
            _ = fi6
            _ = v8
        }
        fmt.Println(fi5)
        if false {
            var v9 uint64 = 1
            _ = v9
        } else {
            var v10 float64 = 1.5
            _ = v10
        }
        _ = fi5
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int64 = -1
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float64 = 1.5
        for fi15 := 0; fi15 < 20; fi15++ {
            if (4294967295 != 42) {
                switch -1 {
                case 9223372036854775807:
                    var v16 int8 = 100
                    _ = v16
                default:
                    var v17 uint64 = 18446744073709551615
                    _ = v17
                }
                v18 := GenS1{F0: 18446744073709551615, F1: 475}
                v19 := []int{-9223372036854775808, fi15, fi13, 558}
                _ = v18
                _ = v19
            } else {
                if (-1 < 127) {
                    var v20 uint64 = 2
                    _ = v20
                }
                var v21 uint = 585
                v22 := GenS1{F0: v21, F1: 1}
                _ = v21
                _ = v22
            }
            fmt.Println(v14, v14)
            if true {
                var v23 int32 = -653
                _ = v23
            } else {
                v24 := []int{fi15, -779, (fi13 + (-(fi13))), fi15}
                _ = v24
            }
            if ((!((!(float64(1.5) == -2.25)) || false)) && (3.14159 > 1e10)) {
                var v25 uint16 = 255
                _ = v25
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    switch (("hello" + "\n\t") + ("hello" + "\n\t")) {
    case "\n\t":
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                if (!true) {
                    var v28 uint32 = 0
                    v28 = 2147483647
                    v29 := []int{fi27, fi27}
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            var v30 float64 = -2.25
            v31 := []int{4294967295, fi26, 32768}
            _ = fi26
            _ = v30
            _ = v31
        }
    case "a":
        var v32 uint16 = 2
        _ = v32
    default:
        v33 := []int{}
        _ = v33
    }
    _ = v4
    return 0
}
func genFn1(p0 int64) uint8 {
    var v34 uint8 = 129
    if (!false) {
        for fi35 := 0; fi35 < 6; fi35++ {
            v36 := []int{-9223372036854775808, (-(fi35))}
            _ = fi35
            _ = v36
        }
        v37 := GenS2{F0: 255}
        var v38 string = "foo bar"
        _ = v37
        _ = v38
    }
    switch "\n\t" {
    case "hello":
        v39 := []int{1}
        _ = v39
    case "a":
        var v40 int16 = 1
        _ = v40
    }
    p0 = (p0 * p0)
    for fi41 := 0; fi41 < 18; fi41++ {
        for fi42 := 0; fi42 < 10; fi42++ {
            var v43 string = "a"
            v43 = v43
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 19; fi44++ {
            p0 = p0
            _ = fi44
        }
        _ = fi41
    }
    _ = p0
    _ = v34
    return v34
}

func main() {
    v1 := []string{}
    v2 := []int{4294967295, 2147483647}
    var v3 float64 = (3.14159 * 1.5)
    fmt.Println(v3)
    var v45 bool = (1e10 >= (v3 / (((v3 + 0.1) / 1e10) - (1e10 * v3))))
    v46 := GenS2{F0: 65535}
    fmt.Println(v3, v45)
    v3 += 1.5
    _ = len(v1)
    v3 += (v3 * 0.1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        v48 := GenS3{F0: 739, F1: -33, F2: 0}
        _ = fi47
        _ = v48
    }
    var v49 uint16 = 42
    for fi50 := 0; fi50 < 10; fi50++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!(!true)) {
            v2 = append(v2, -1)
            for fi51 := 0; fi51 < 12; fi51++ {
                fmt.Println(fi51)
                _ = fi51
            }
        }
        var v52 uint16 = 1
        _ = fi50
        _ = v52
    }
    var v53 int32 = 0
    v49 -= 440
    for fi54 := 0; fi54 < 8; fi54++ {
        v55 := GenS4{F0: 4294967295, F1: -356, F2: 42}
        for fi56 := 0; fi56 < 9; fi56++ {
            if v45 {
                v57 := GenS5{F0: "hello", F1: 42, F2: -1}
                _ = len(v1)
                fmt.Println(fi54, v49, v53)
                _ = v57
            } else {
                _ = len(v1)
                switch "" {
                case "hello":
                    fmt.Println(v49, fi56, v45)
                    var v58 uint = 456
                    fmt.Println(v45)
                    _ = v58
                }
                if (4294967295 > 2) {
                    var v59 uint = 42
                    var v60 uint16 = (v49 | (v49 / ((409) | 1)))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v59
                    _ = v60
                }
            }
            switch (fi54 | (fi56 * (fi54 % ((fi54) | 1)))) {
            case 0:
                var v61 uint8 = 0
                _ = v61
            case 1:
                if (((0.0 + 1e10) >= 1e10) && true) {
                    if (!v45) {
                        var v62 uint = 0
                        switch 2147483647 {
                        case -9223372036854775808:
                            if (!(1e10 >= (-2.25 - ((1e10 + 0.1) * 1.5)))) {
                                v53 += v53
                            } else {
                                v63 := GenS2{F0: (v49 >> (((v62 ^ v62)) & 63))}
                                _ = v63
                            }
                        case 127:
                            v53 = v53
                        case -1:
                            fmt.Println(v45)
                        }
                        switch 127 {
                        case -674:
                            v45 = v45
                        case 2147483647:
                            v62 = (v62 % (((^(v62))) | 1))
                        case -1:
                            if v45 {
                                v64 := GenS5{F0: "\n\t", F1: uint64(v62), F2: fi54}
                                fmt.Println(v49)
                                switch "foo bar" {
                                case "a":
                                    var v65 bool = v45
                                    _ = v65
                                case "hello":
                                    fmt.Println(v3)
                                }
                                v49 = v49
                                if true {
                                    v53 = (v53 / ((-2147483648) | 1))
                                    var v66 int16 = -408
                                    _ = v66
                                }
                                _ = v64
                            }
                        }
                        _ = v62
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v45, v3, v53)
                        fmt.Println(fi54, v53)
                        if v45 {
                            switch -871 {
                            case 42:
                                fmt.Println(v49)
                            case 9223372036854775807:
                                var v67 float64 = 1e10
                                _ = v67
                            }
                        }
                    case "\n\t":
                        switch fi54 {
                        case 0:
                            if v45 {
                                v68 := GenS6{F0: false, F1: "foo bar", F2: 255}
                                v69 := GenS7{F0: "a"}
                                if ((!((("\n\t" + "\n\t") + (("\n\t" + (("\n\t" + "a") + "\n\t")) + "")) != ((("hello" + (("a" + "foo bar") + "")) + "\n\t") + ((("" + ("foo bar" + "\n\t")) + "") + ("" + (("foo bar" + ("\n\t" + "hello")) + ("hello" + "a"))))))) && v45) {
                                    var v70 int = (fi56 / (((fi54 >> ((1) & 63))) | 1))
                                    v45 = (2 > 18446744073709551615)
                                    _ = v70
                                }
                                _ = v68
                                _ = v69
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    v71 := GenS5{F0: ("\n\t" + ""), F1: 557, F2: fi56}
                                    _ = v71
                                default:
                                    fmt.Println(fi54, fi56, v45)
                                }
                                if (!(!(v3 == (0.1 + 0.1)))) {
                                    v72 := GenS7{F0: "hello"}
                                    v49 -= v49
                                    _ = v72
                                } else {
                                    v53 = v53
                                    v2 = append(v2, (fi54 / ((42) | 1)))
                                }
                                fmt.Println(v53)
                                var v73 float32 = 3.14159
                                if true {
                                    if true {
                                        switch fi56 {
                                        case -397:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case 590:
                                            fmt.Println(v49, fi54)
                                        }
                                        fmt.Println(v49, v53, v3)
                                        if v45 {
                                            var v74 int = (fi54 % (((fi56 & (fi54 & 716))) | 1))
                                            v1 = append(v1, "")
                                            _ = v74
                                        }
                                        fmt.Println(v3, fi56, fi56)
                                    }
                                }
                                _ = v73
                            }
                        case -1:
                            switch (("" + "\n\t") + "foo bar") {
                            case "\n\t":
                                v75 := GenS5{F0: ("hello" + "hello"), F1: 1, F2: -1}
                                _ = v75
                            }
                            v45 = false
                            if (255 != 27) {
                                v49 += v49
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 32767:
                            fmt.Println(fi56, v53)
                            v76 := GenS5{F0: "\n\t", F1: 1, F2: fi54}
                            var v77 uint8 = 0
                            v49 += 1
                            _ = v76
                            _ = v77
                        }
                    case "hello":
                        if (128 != 58) {
                            if v45 {
                                v53 = v53
                                fmt.Println(v49, fi54, v3)
                            }
                            fmt.Println(fi56, fi54)
                            fmt.Println(v49)
                        } else {
                            switch int(v53) {
                            case 9223372036854775807:
                                fmt.Println(fi56)
                            }
                            if len(v1) > 0 {
                                v1[0] = ("a" + ("\n\t" + ("" + "")))
                            }
                        }
                    }
                } else {
                    v45 = ((v49 << ((2) & 63)) <= 0)
                    if v45 {
                        var v78 string = "a"
                        if false {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                        } else {
                            switch (fi56 / ((fi56) | 1)) {
                            case 807:
                                var v79 int8 = -1
                                _ = v79
                            case 1:
                                var v80 string = v78
                                _ = v80
                            }
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        }
                        v3 += v3
                        v81 := GenS8{F0: 2, F1: 65535}
                        _ = v78
                        _ = v81
                    } else {
                        if (v3 >= (v3 - ((3.14159 + 0.0) * float64((0.1 - (1.5 + ((1e10 / 1.0) + -2.25))))))) {
                            _ = len(v2)
                            v82 := GenS9{F0: ("\n\t" + ("" + "foo bar"))}
                            v49 = (v49 / (((^(v49))) | 1))
                            _ = v82
                        } else {
                            fmt.Println(fi56, fi56, fi54)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v3, v45)
                        }
                        fmt.Println(v53)
                        var v83 int = fi54
                        _ = v83
                    }
                    v53 = -611
                }
                v49 += (v49 & 621)
                v84 := GenS9{F0: ""}
                var v85 float32 = 3.14159
                fmt.Println(v3)
                _ = v84
                _ = v85
            case 2147483647:
                v3 += (((v3 - (v3 - v3)) / v3) * ((1e10 * (v3 + float64(0.0))) * 0.0))
            default:
                if v45 {
                    var v86 uint8 = genFn1(1)
                    v1 = append(v1, "foo bar")
                    _ = v86
                }
            }
            v45 = (!(!(0.0 < (-2.25 - (1.5 * 1e10)))))
            _ = fi56
        }
        _ = fi54
        _ = v55
    }
    var v87 int8 = 0
    if true {
        var v88 uint = 518
        _ = v88
    } else {
        v87 = v87
        fmt.Println(v45, v53, v45)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v49
    _ = v53
    _ = v87
}
