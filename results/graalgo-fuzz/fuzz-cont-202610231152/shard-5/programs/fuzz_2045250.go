package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 uint32; F1 int16; F2 int16 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int, p2 uint8) uint32 {
    var v4 float64 = 3.14159
    p0 = 1
    if (uint64(p0) <= 244) {
        p1 = p1
        v5 := GenS1{F0: -2147483648}
        var v6 uint32 = 2147483647
        _ = v5
        _ = v6
    } else {
        switch p1 {
        case 32768:
            for fi7 := 0; fi7 < 7; fi7++ {
                v4 += (v4 / v4)
                _ = fi7
            }
        case -797:
            v8 := []int{4294967295, (-(p1)), (p1 | (p1 + (p1 - (p1 / ((128) | 1))))), p1}
            _ = v8
        case 719:
            switch ("\n\t" + "a") {
            case "\n\t":
                p0 = (^(p0))
            case "hello":
                v9 := []int{}
                _ = v9
            default:
                for fi10 := 0; fi10 < 3; fi10++ {
                    p0 += (p0 << ((2) & 63))
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
            }
        }
        switch "\n\t" {
        case "hello":
            for fi12 := 0; fi12 < 16; fi12++ {
                switch p1 {
                case 410:
                    var v13 uint8 = (^(p2))
                    _ = v13
                case 1:
                    var v14 int8 = 1
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int8 = 1
                        var v17 string = "foo bar"
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                var v18 int32 = 0
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch p1 {
                    case -2147483648:
                        var v20 uint32 = 65535
                        _ = v20
                    case -702:
                        var v21 uint32 = 65535
                        _ = v21
                    }
                    if ((327 != 1) || false) {
                        fmt.Println(p1, fi12, v4)
                    }
                    v4 -= 0.0
                    var v22 int16 = 949
                    _ = fi19
                    _ = v22
                }
                _ = fi12
                _ = v18
            }
        }
        var v23 uint8 = (p2 ^ (p2 - (p2 / ((p2) | 1))))
        _ = v23
    }
    if ("\n\t" == (("\n\t" + ("hello" + "\n\t")) + "\n\t")) {
        for fi24 := 0; fi24 < 10; fi24++ {
            switch (p1 ^ p1) {
            case 566:
                var v25 int = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 17; fi26++ {
                if (!(!("a" != ((((("a" + "foo bar") + "foo bar") + "\n\t") + ("hello" + "hello")) + (((("foo bar" + "\n\t") + "a") + ("foo bar" + ("a" + ""))) + "\n\t"))))) {
                    v27 := []int{-468, p1}
                    _ = len(v27)
                    _ = v27
                } else {
                    var v28 float32 = float32(v4)
                    var v29 uint16 = (p0 - 65535)
                    _ = v28
                    _ = v29
                }
                p2 += p2
                switch 32767 {
                case 32768:
                    var v30 float64 = 3.14159
                    v30 -= ((-2.25 - 1e10) - v30)
                    v4 = (v30 - 0.0)
                    _ = v30
                case 0:
                    v31 := GenS2{F0: "foo bar"}
                    _ = v31
                }
                _ = fi26
            }
            _ = fi24
        }
        v32 := GenS3{F0: int8(p2), F1: 1, F2: p0}
        _ = v32
    }
    switch "hello" {
    case "hello":
        v33 := GenS4{F0: 100, F1: (p2 * (p2 * (p2 * 128)))}
        _ = v33
    case "\n\t":
        if false {
            for fi34 := 0; fi34 < 12; fi34++ {
                for fi35 := 0; fi35 < 19; fi35++ {
                    v36 := []int{(fi35 & fi35), fi35, (fi34 >> ((1) & 63)), 1}
                    _ = fi35
                    _ = v36
                }
                for fi37 := 0; fi37 < 16; fi37++ {
                    switch (fi37 * (fi37 ^ -952)) {
                    case 127:
                        var v38 int = 127
                        _ = v38
                    default:
                        p1 += (^(fi37))
                    }
                    _ = fi37
                }
                _ = fi34
            }
            v39 := []int{(^(p1))}
            _ = v39
        }
    default:
        fmt.Println(p0, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 uint) int8 {
    if false {
        v40 := []int{426, 4294967295}
        v40 = append(v40, -2147483648)
        if ((0 < -1) && false) {
            switch 128 {
            case 64:
                for fi41 := 0; fi41 < 10; fi41++ {
                    fmt.Println(p0, fi41)
                    p0 = (p0 & p0)
                    _ = fi41
                }
            case -1:
                v40 = append(v40, 0)
            case 127:
                p0 = 42
            }
            p0 = p0
        } else {
            p0 -= p0
        }
        _ = v40
    }
    fmt.Println(p0)
    var v42 int = -1
    _ = p0
    _ = v42
    return 42
}

func main() {
    var v1 string = ("" + "foo bar")
    var v2 int8 = 1
    v3 := make(map[int]int)
    if true {
        var v43 float32 = -2.25
        var v44 uint16 = 32767
        _ = v43
        _ = v44
    }
    v2 = (v2 - (v2 >> ((924) & 63)))
    v45 := []int{573, 2147483647}
    v46 := GenS5{F0: 0}
    var v47 uint = 1
    if (!(-2.25 != 1.5)) {
        v48 := GenS5{F0: 713}
        fmt.Println(v47, v47)
        var v49 uint32 = 42
        _ = v48
        _ = v49
    } else {
        v3[-9223372036854775808] = 32767
        if false {
            v50 := GenS5{F0: 734}
            var v51 float32 = 1.5
            _ = v50
            _ = v51
        } else {
            for fi52 := 0; fi52 < 5; fi52++ {
                v53 := GenS3{F0: (v2 + genFn1(154)), F1: v2, F2: 2}
                var v54 int64 = 42
                _ = fi52
                _ = v53
                _ = v54
            }
        }
    }
    v55 := GenS2{F0: (v1 + v1)}
    var v56 int8 = (v2 >> (((v47 ^ 2)) & 63))
    var v57 uint16 = 260
    var v58 uint8 = 2
    if (!false) {
        if false {
            v2 = (v56 / ((genFn1(42)) | 1))
            switch v1 {
            case "":
                if false {
                    for fi59 := 0; fi59 < 1; fi59++ {
                        if false {
                            var v60 uint32 = 859
                            v56 = int8(v57)
                            _ = v60
                        }
                        _ = fi59
                    }
                }
            case "a":
                if true {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        var v62 int16 = -933
                        var v63 uint8 = v58
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    _ = len(v45)
                    if len(v45) > 0 {
                        v45[0] = 42
                    }
                } else {
                    fmt.Println(v1, v56)
                }
            default:
                for fi64 := 0; fi64 < 19; fi64++ {
                    var v65 uint = v47
                    _ = fi64
                    _ = v65
                }
            }
            fmt.Println(v56)
        } else {
            var v66 int = -905
            _ = len(v45)
            _ = v66
        }
        var v67 int32 = 42
        v2 = v2
        _ = v67
    } else {
        switch ("" + "") {
        case "\n\t":
            switch v1 {
            case "\n\t":
                v68 := GenS3{F0: v56, F1: genFn1(v47), F2: v57}
                _ = v68
            case "":
                if false {
                    var v69 float32 = float32((-2.25 / -2.25))
                    v70 := GenS2{F0: (v1 + v1)}
                    _ = v69
                    _ = v70
                }
            case "foo bar":
                v45 = append(v45, 0)
            default:
                for fi71 := 0; fi71 < 1; fi71++ {
                    for fi72 := 0; fi72 < 6; fi72++ {
                        var v73 uint64 = 32
                        var v74 uint8 = (v58 & 1)
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 16; fi75++ {
                        if false {
                            if false {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                            } else {
                                var v76 int16 = 32767
                                _ = v76
                            }
                        } else {
                            v77 := GenS5{F0: -160}
                            _ = v77
                        }
                        _ = fi75
                    }
                    switch -1 {
                    case -78:
                        if false {
                            fmt.Println(v56, v57, v47)
                            for fi78 := 0; fi78 < 16; fi78++ {
                                if (261 < -153) {
                                    v1 = (((((v1 + "hello") + v1) + "a") + (v1 + ((((("\n\t" + (v1 + "foo bar")) + "hello") + (v1 + (v1 + v1))) + "\n\t") + ((v1 + "foo bar") + v1)))) + ((v1 + v1) + "a"))
                                    var v79 int32 = -548
                                    _ = v79
                                } else {
                                    if (18446744073709551615 > 61) {
                                        switch (fi78 & fi71) {
                                        case 544:
                                            var v80 int64 = 9223372036854775807
                                            _ = v80
                                        case 127:
                                            fmt.Println(v47)
                                            fmt.Println(v2, fi78)
                                            delete(v3, fi71)
                                        default:
                                            v56 -= (v56 ^ v2)
                                        }
                                        switch 867 {
                                        case 2147483647:
                                            var v81 uint64 = 305
                                            _ = v81
                                        }
                                        v57 = (v57 - (v57 ^ v57))
                                    }
                                    v1 = ("hello" + v1)
                                }
                                _ = fi78
                            }
                        }
                    case -778:
                        v56 = genFn1((^(v47)))
                    default:
                        for fi82 := 0; fi82 < 16; fi82++ {
                            v3[fi82] = (fi71 % ((1) | 1))
                            v47 += (^(v47))
                            _ = fi82
                        }
                    }
                    _ = fi71
                }
            }
        case "foo bar":
            if true {
                if (129 != 1) {
                    fmt.Println(v57, v1)
                    for fi83 := 0; fi83 < 4; fi83++ {
                        v45 = append(v45, fi83)
                        _ = fi83
                    }
                } else {
                    _ = len(v3)
                    var v84 int8 = 1
                    _ = v84
                }
            } else {
                fmt.Println(v1, v56)
                if false {
                    for fi85 := 0; fi85 < 5; fi85++ {
                        switch v1 {
                        case "\n\t":
                            if len(v45) > 0 {
                                v45[0] = fi85
                            }
                        case "a":
                            v86 := GenS2{F0: (v1 + "hello")}
                            _ = v86
                        }
                        var v87 uint8 = 74
                        _ = fi85
                        _ = v87
                    }
                    v88 := GenS6{F0: 1, F1: (738 != uint32(v58))}
                    _ = v88
                } else {
                    v89 := GenS7{F0: 65535, F1: int16(v58), F2: 1000}
                    var v90 uint8 = v58
                    for fi91 := 0; fi91 < 20; fi91++ {
                        var v92 uint16 = (v57 - (v57 ^ (v57 / (((^(v57))) | 1))))
                        _ = fi91
                        _ = v92
                    }
                    _ = v89
                    _ = v90
                }
            }
        case "a":
            fmt.Println(v58)
        default:
            if (!true) {
                if true {
                    _ = len(v3)
                }
                if false {
                    v45 = append(v45, 42)
                } else {
                    for fi93 := 0; fi93 < 17; fi93++ {
                        for fi94 := 0; fi94 < 15; fi94++ {
                            delete(v3, -9223372036854775808)
                            _ = fi94
                        }
                        for fi95 := 0; fi95 < 7; fi95++ {
                            if (true && true) {
                                v96 := GenS3{F0: 127, F1: v2, F2: v57}
                                _ = v96
                            } else {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                            }
                            _ = fi95
                        }
                        var v97 bool = false
                        _ = fi93
                        _ = v97
                    }
                }
                for fi98 := 0; fi98 < 11; fi98++ {
                    v58 += v58
                    v56 += genFn1(v47)
                    _ = fi98
                }
            } else {
                v56 -= (v2 << (((v47 >> (((v47 ^ 0)) & 63))) & 63))
                v56 = (v56 ^ (v56 / (((v2 + v2)) | 1)))
                var v99 float64 = 1.5
                _ = v99
            }
        }
        v100 := GenS8{F0: v58}
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
