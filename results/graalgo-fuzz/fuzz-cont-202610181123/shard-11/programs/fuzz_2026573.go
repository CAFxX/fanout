package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int16; F1 bool; F2 int64 }
type GenS6 struct { F0 uint16; F1 bool; F2 bool }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 int64; F1 uint; F2 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS11 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 int16, p1 uint64, p2 uint8) uint {
    var v5 int8 = int8(p1)
    fmt.Println(p0, p2)
    if false {
        fmt.Println(v5, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 int) uint16 {
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 9; fi7++ {
        p1 += p1
        _ = fi7
    }
    switch "a" {
    case "foo bar":
        p1 = 604
    case "a":
        for fi8 := 0; fi8 < 4; fi8++ {
            switch "\n\t" {
            case "foo bar":
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := []int{}
                    _ = fi9
                    _ = v10
                }
            case "a":
                if (42 != 332) {
                    var v11 string = "foo bar"
                    _ = v11
                }
            }
            p0 += p0
            v12 := GenS2{F0: 0, F1: 2}
            _ = fi8
            _ = v12
        }
    default:
        var v13 bool = true
        _ = v13
    }
    if (42 <= 1000) {
        v14 := GenS3{F0: 1}
        _ = v14
    } else {
        p0 = (p0 >> ((255) & 63))
        if false {
            p0 += (p0 >> ((0) & 63))
            var v15 string = "hello"
            _ = v15
        } else {
            p1 = p1
            if false {
                for fi16 := 0; fi16 < 3; fi16++ {
                    if (!((0.1 + float32(-2.25)) >= 0.0)) {
                        var v17 int8 = 60
                        _ = v17
                    } else {
                        v18 := []int{p1, p1, (^(fi16))}
                        p0 = 1
                        _ = v18
                    }
                    var v19 int64 = 0
                    _ = fi16
                    _ = v19
                }
                fmt.Println(p0, p1)
                v20 := GenS3{F0: 723}
                _ = v20
            } else {
                var v21 uint = 772
                _ = v21
            }
            switch 128 {
            case 276:
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (745 < 4294967295) {
                        var v23 uint32 = 913
                        var v24 float32 = 0.0
                        v23 = (v23 & 2)
                        _ = v23
                        _ = v24
                    } else {
                        v25 := []int{p1}
                        v26 := GenS4{F0: -9223372036854775808, F1: 18446744073709551615}
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 7; fi27++ {
                        switch fi22 {
                        case 633:
                            var v28 int16 = -178
                            _ = v28
                        case 32767:
                            var v29 bool = (false || (true && (true || ((!("hello" > "a")) || true))))
                            _ = v29
                        case 860:
                            switch ("foo bar" + "hello") {
                            case "":
                                v30 := []int{}
                                _ = v30
                            case "foo bar":
                                p1 += (fi22 | (fi27 / ((-133) | 1)))
                                var v31 int16 = 1000
                                var v32 int16 = 740
                                _ = v31
                                _ = v32
                            default:
                                if ((fi27 << ((2) & 63)) != fi27) {
                                    if (!(false || (fi27 > (fi22 ^ 32767)))) {
                                        if (true || true) {
                                            v33 := []int{fi22}
                                            switch 32768 {
                                            case 0:
                                                v34 := GenS5{F0: 1, F1: false, F2: 9223372036854775807}
                                                _ = v34
                                            case -9223372036854775808:
                                                var v35 string = "foo bar"
                                                _ = v35
                                            case 2147483647:
                                                if (!(2 >= 4294967295)) {
                                                    var v36 bool = (1e10 >= 1.5)
                                                    if (127 != 62) {
                                                        var v37 string = "hello"
                                                        fmt.Println(p1, p0, v36)
                                                        v33 = append(v33, p1)
                                                        if v36 {
                                                            v38 := GenS1{F0: 1}
                                                            _ = v38
                                                        } else {
                                                            var v39 int16 = 42
                                                            v33 = append(v33, (p1 >> ((genFn0(v39, 906, 1)) & 63)))
                                                            _ = v39
                                                        }
                                                        if (!((p0 != p0) || v36)) {
                                                            var v40 int64 = 9223372036854775807
                                                            var v41 float32 = ((((float32(0.1) / float32(0.1)) - (0.1 - 0.1)) / -2.25) / 1.0)
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        _ = v37
                                                    } else {
                                                        _ = len(v33)
                                                        v42 := GenS5{F0: 539, F1: ((!(v36 && (!((0 != -110) || v36)))) && (255 >= 65535)), F2: -758}
                                                        var v43 uint16 = 0
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    _ = v36
                                                }
                                            default:
                                                fmt.Println(p0)
                                            }
                                            _ = v33
                                        } else {
                                            v44 := GenS6{F0: 32767, F1: (!true), F2: false}
                                            var v45 int8 = int8(p0)
                                            fmt.Println(v45, fi27, fi22)
                                            _ = v44
                                            _ = v45
                                        }
                                        fmt.Println(p0, fi22)
                                        p1 = (-(p1))
                                    } else {
                                        v46 := []int{}
                                        v47 := GenS6{F0: 32767, F1: (18446744073709551615 != 4294967295), F2: true}
                                        _ = v46
                                        _ = v47
                                    }
                                    p0 -= p0
                                    p1 = fi27
                                } else {
                                    v48 := []int{(^(fi22)), p1}
                                    v49 := GenS6{F0: 32767, F1: true, F2: false}
                                    _ = v48
                                    _ = v49
                                }
                            }
                        default:
                            v50 := GenS2{F0: 0, F1: 18446744073709551615}
                            v51 := []int{(p1 | 42), 914, (^(fi22)), p1}
                            _ = v50
                            _ = v51
                        }
                        v52 := GenS2{F0: 9223372036854775807, F1: 65535}
                        _ = fi27
                        _ = v52
                    }
                    _ = fi22
                }
            case -531:
                p1 = (p1 << ((317) & 63))
            case -1:
                v53 := []int{}
                _ = v53
            }
        }
        var v54 int16 = 32767
        _ = v54
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint = (^(v1))
    v3 := []int{127, 9223372036854775807, 32768}
    v4 := []string{"", "\n\t"}
    switch 708 {
    case 1:
        var v55 int32 = 527
        _ = v55
    }
    fmt.Println(v1)
    var v56 uint8 = 255
    for fi57 := 0; fi57 < 17; fi57++ {
        v58 := GenS7{F0: genFn0(32767, 4294967295, v56), F1: 1000}
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        var v59 uint = 1
        _ = fi57
        _ = v58
        _ = v59
    }
    v60 := GenS8{F0: -9223372036854775808, F1: v2, F2: (("foo bar" + ("" + "a")) + "foo bar")}
    for fi61 := 0; fi61 < 20; fi61++ {
        var v62 bool = (v56 > v56)
        v2 = 929
        v62 = ((255 >= (^(v1))) || v62)
        _ = fi61
        _ = v62
    }
    if (85 <= 42) {
        if false {
            v56 = (v56 | v56)
            var v63 float32 = 0.1
            v64 := GenS9{F0: true}
            _ = v63
            _ = v64
        }
    }
    v65 := GenS3{F0: genFn1(0, 128)}
    switch "\n\t" {
    case "foo bar":
        var v66 int16 = 32767
        _ = v66
    case "\n\t":
        var v67 int16 = 1000
        _ = v67
    default:
        switch (("foo bar" + "a") + "\n\t") {
        case "foo bar":
            var v68 uint8 = (v56 >> (((v1 * (^(v2)))) & 63))
            _ = v68
        }
    }
    if (!true) {
        var v69 int16 = -32768
        if true {
            v69 -= -721
        }
        _ = v69
    } else {
        var v70 uint16 = genFn1(0, -1)
        for fi71 := 0; fi71 < 8; fi71++ {
            if len(v3) > 0 {
                v3[0] = fi71
            }
            _ = fi71
        }
        v1 += v2
        _ = v70
    }
    if false {
        var v72 int8 = 100
        v1 = (v2 % ((811) | 1))
        _ = v72
    }
    v3 = append(v3, -1)
    var v73 int8 = -1
    switch 782 {
    case 128:
        for fi74 := 0; fi74 < 8; fi74++ {
            v75 := GenS6{F0: 32767, F1: ((((0 > int32(fi74)) || (523 <= 788)) || false) || (v73 > v73)), F2: (int64(v1) == 42)}
            _ = fi74
            _ = v75
        }
    default:
        switch "hello" {
        case "":
            if (true || false) {
                fmt.Println(v1)
            }
        default:
            v76 := GenS10{F0: 18446744073709551615, F1: 42, F2: 65535}
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 6; fi77++ {
        _ = len(v4)
        v2 += v2
        v2 += v2
        _ = fi77
    }
    if false {
        v78 := GenS6{F0: genFn1(2147483647, 9223372036854775807), F1: (0 != 2), F2: false}
        _ = v78
    } else {
        if false {
            v3 = append(v3, 0)
            var v79 int8 = v73
            v56 = 127
            _ = v79
        } else {
            v80 := GenS11{F0: 731, F1: 1, F2: 2}
            switch -725 {
            case 797:
                var v81 bool = ((-45 >= v73) && true)
                _ = v81
            }
            _ = v80
        }
    }
    if (149 >= 1) {
        v1 += v1
        v3 = append(v3, 28)
        fmt.Println(v1, v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v56:", v56)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v60
    _ = v65
    _ = v73
}
