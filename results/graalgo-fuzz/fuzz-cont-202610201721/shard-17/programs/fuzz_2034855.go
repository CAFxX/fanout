package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int16; F2 uint }
type GenS12 struct { F0 uint64; F1 int32; F2 int16 }
type GenS13 struct { F0 uint8; F1 int; F2 int8 }
type GenS14 struct { F0 int32; F1 string }
func genFn0(p0 uint16, p1 bool) int32 {
    if (0 != -875) {
        v4 := GenS1{F0: 220}
        p1 = (741 < 42)
        _ = v4
    } else {
        p0 = p0
        v5 := []int{9223372036854775807, 127, 787}
        _ = v5
    }
    v6 := GenS2{F0: (^(p0))}
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{(("" + "") + ("" + "")), "\n\t"}
    var v2 int32 = 42
    var v3 int32 = v2
    v2 = (v3 % ((v3) | 1))
    fmt.Println(v3)
    if true {
        var v7 int = 42
        v7 = -9223372036854775808
        _ = v7
    }
    if (-9223372036854775808 != -9223372036854775808) {
        if (!true) {
            v2 += v3
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        } else {
            if false {
                for fi8 := 0; fi8 < 2; fi8++ {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = fi8
                }
                for fi9 := 0; fi9 < 15; fi9++ {
                    v3 += genFn0(255, false)
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        fmt.Println(fi9, fi9, v2)
                        v11 := GenS3{F0: 2147483647}
                        var v12 uint16 = 2
                        v3 += genFn0(v12, (!(18446744073709551615 >= 255)))
                        v3 = v2
                        v13 := GenS4{F0: 1, F1: -1}
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi9
                }
                v1 = append(v1, "\n\t")
            } else {
                switch 9223372036854775807 {
                case 128:
                    switch "" {
                    case "a":
                        v3 = 175
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    case "\n\t":
                        switch 128 {
                        case 803:
                            _ = len(v1)
                        case -270:
                            for fi14 := 0; fi14 < 14; fi14++ {
                                var v15 uint8 = 127
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    if (!(v3 >= 672)) {
                                        v3 += (v3 - -210)
                                        v15 += v15
                                    } else {
                                        if (!false) {
                                            v15 -= v15
                                        } else {
                                            v3 = 0
                                            var v17 int64 = 9223372036854775807
                                            var v18 string = "a"
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    v19 := GenS5{F0: v3, F1: (v2 | (v3 / ((v2) | 1)))}
                                    _ = fi16
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                    var v21 uint32 = 2147483648
                                    v22 := GenS6{F0: (v2 / ((v3) | 1))}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = fi14
                                _ = v15
                            }
                        case 378:
                            v3 = genFn0(863, (!(false || true)))
                        default:
                            v23 := GenS2{F0: 32767}
                            _ = v23
                        }
                    }
                default:
                    switch "hello" {
                    case "foo bar":
                        v24 := GenS7{F0: 0}
                        _ = v24
                    }
                }
                switch 42 {
                case -2147483648:
                    v2 += (v3 + (v2 * v3))
                case -420:
                    switch 32768 {
                    case 1:
                        switch "foo bar" {
                        case "\n\t":
                            var v25 uint16 = 238
                            _ = v25
                        }
                    case 42:
                        fmt.Println(v3, v3)
                    case 127:
                        v26 := GenS8{F0: 4294967295}
                        _ = v26
                    }
                default:
                    v27 := GenS9{F0: 207}
                    _ = v27
                }
                var v28 int16 = 32767
                _ = v28
            }
            fmt.Println(v2)
        }
        var v29 string = ""
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int32 = -1
            var v32 float64 = ((1.5 * 3.14159) * -2.25)
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v33 := GenS10{F0: 1}
    v34 := GenS11{F0: -2147483648, F1: 1, F2: 65535}
    switch 127 {
    case 127:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    case 0:
        if (127 != 127) {
            v35 := GenS12{F0: 42, F1: (v3 | (^(v3))), F2: 0}
            _ = v35
        }
    case 1:
        var v36 int = -521
        _ = v36
    default:
        if (!false) {
            var v37 int = 0
            v37 -= v37
            for fi38 := 0; fi38 < 11; fi38++ {
                v39 := GenS1{F0: 42}
                for fi40 := 0; fi40 < 8; fi40++ {
                    v37 -= fi40
                    if false {
                        if (!true) {
                            if ((((0.0 - 0.0) - -2.25) + (1.5 + 0.1)) >= ((((float32((float64(((-2.25 - ((((0.1 + (3.14159 + ((0.1 + (1e10 + 1.5)) * 1.5))) * 1e10) + 0.1) / 1e10)) / 1.0)) * 0.1)) * 0.0) / 1.5) / float32(1.5)) / (3.14159 / (3.14159 / -2.25)))) {
                                if true {
                                    v37 -= -1
                                    v41 := GenS12{F0: 0, F1: v3, F2: 42}
                                    fmt.Println(v3, fi38, fi40)
                                    _ = v41
                                }
                            } else {
                                var v42 string = ("a" + (("a" + "\n\t") + "\n\t"))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v42
                            }
                            var v43 int32 = genFn0(2, (42 <= 42))
                            _ = v43
                        }
                        v2 = 42
                        fmt.Println(v2)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi40
                }
                for fi44 := 0; fi44 < 13; fi44++ {
                    v3 -= (v3 & 2147483647)
                    v2 += 0
                    var v45 bool = (true || (true && false))
                    _ = fi44
                    _ = v45
                }
                _ = fi38
                _ = v39
            }
            _ = v37
        } else {
            var v46 uint = 42
            _ = v46
        }
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "hello")
    }
    if (32767 > -276) {
        v47 := GenS13{F0: 2, F1: 128, F2: -1}
        _ = v47
    } else {
        var v48 float32 = float32(((1.5 + 1.5) - 3.14159))
        _ = v48
    }
    if false {
        v2 -= 0
    } else {
        var v49 uint64 = 133
        _ = v49
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        v1 = append(v1, ((("a" + ("\n\t" + "\n\t")) + "\n\t") + "hello"))
        v1 = append(v1, "\n\t")
        fmt.Println(fi50, v3)
        _ = fi50
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        fmt.Println(fi51, v2, v3)
        for fi52 := 0; fi52 < 8; fi52++ {
            if (-79 > -1) {
                if ("foo bar" < "\n\t") {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch (fi52 & 128) {
                        case -1:
                            if (4294967295 != 2) {
                                switch (^(fi52)) {
                                case 4294967295:
                                    var v53 int = fi52
                                    _ = v53
                                case 128:
                                    if (-67 > 769) {
                                        switch (^(fi52)) {
                                        case 1:
                                            var v54 int8 = -128
                                            _ = v54
                                        case 4294967295:
                                            if true {
                                                v2 += v3
                                                v3 = v2
                                            } else {
                                                var v55 uint16 = 42
                                                _ = v55
                                            }
                                        case -2147483648:
                                            v56 := GenS9{F0: -9223372036854775808}
                                            _ = v56
                                        default:
                                            _ = len(v1)
                                        }
                                        switch 427 {
                                        case 53:
                                            switch (fi51 * fi52) {
                                            case -1:
                                                var v57 uint = 0
                                                _ = v57
                                            }
                                        case -26:
                                            var v58 int16 = 323
                                            _ = v58
                                        case 4294967295:
                                            v2 = v2
                                        default:
                                            if (("\n\t" < (("foo bar" + "hello") + "foo bar")) && true) {
                                                if (-1 <= 1) {
                                                    var v59 uint64 = 18446744073709551615
                                                    fmt.Println(v59)
                                                    fmt.Println(fi51, fi51)
                                                    fmt.Println(v59, v2)
                                                    _ = v59
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                                v3 = (v3 << ((567) & 63))
                                                if len(v1) > 0 {
                                                    v1[0] = ("hello" + "a")
                                                }
                                                v60 := GenS13{F0: 42, F1: (fi51 ^ fi52), F2: -99}
                                                var v61 int = fi52
                                                if true {
                                                    if false {
                                                        switch "a" {
                                                        case "hello":
                                                            if (42 <= 0) {
                                                                var v62 int64 = 9223372036854775807
                                                                _ = v62
                                                            } else {
                                                                v3 += (v3 | v2)
                                                            }
                                                            var v63 float32 = 1.5
                                                            _ = v63
                                                        case "a":
                                                            v64 := GenS13{F0: 127, F1: -9223372036854775808, F2: -82}
                                                            _ = v64
                                                        }
                                                        var v65 uint32 = 1
                                                        if false {
                                                            fmt.Println(v3, v3, v3)
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v61 -= v61
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v66 := GenS14{F0: v3, F1: ("\n\t" + "a")}
                                                            if (!(1.5 != -2.25)) {
                                                                if (!false) {
                                                                    v67 := GenS8{F0: 382}
                                                                    var v68 int64 = -1
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    var v69 float64 = 3.14159
                                                                    if (!true) {
                                                                        if false {
                                                                            if true {
                                                                                if ((1 > 1) && true) {
                                                                                    v70 := GenS9{F0: 495}
                                                                                    _ = v70
                                                                                }
                                                                                fmt.Println(fi52, v3, v69)
                                                                            } else {
                                                                                _ = len(v1)
                                                                                v69 += 3.14159
                                                                            }
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v3 -= v2
                                                                            }
                                                                            var v71 int = -9223372036854775808
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    _ = v69
                                                                }
                                                                fmt.Println(fi52, v2)
                                                                v2 = (v3 * (^(v3)))
                                                                var v72 int64 = -348
                                                                _ = len(v1)
                                                                _ = v72
                                                            }
                                                            _ = v66
                                                        }
                                                        v3 += v3
                                                        v61 += 2147483647
                                                        switch fi52 {
                                                        case 32768:
                                                            _ = len(v1)
                                                            v73 := GenS7{F0: 42}
                                                            v3 = 424
                                                            _ = v73
                                                        case 220:
                                                            v74 := GenS11{F0: 431, F1: 0, F2: 42}
                                                            _ = v74
                                                        case -2147483648:
                                                            v65 = (v65 | (v65 / (((v65 + v65)) | 1)))
                                                        default:
                                                            var v75 uint8 = 127
                                                            _ = v75
                                                        }
                                                        var v76 uint16 = 820
                                                        _ = v65
                                                        _ = v76
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                var v77 uint16 = 1
                                                _ = len(v1)
                                                _ = v60
                                                _ = v61
                                                _ = v77
                                            }
                                        }
                                    } else {
                                        v1 = append(v1, "hello")
                                    }
                                case -595:
                                    fmt.Println(v3)
                                }
                                var v78 uint16 = 1
                                _ = v78
                            }
                            var v79 bool = true
                            _ = v79
                        case -860:
                            if false {
                                var v80 uint16 = 255
                                var v81 int32 = genFn0(v80, true)
                                if true {
                                    v82 := GenS9{F0: 0}
                                    v83 := GenS11{F0: (fi52 * fi52), F1: 1000, F2: 65535}
                                    _ = v82
                                    _ = v83
                                }
                                _ = v80
                                _ = v81
                            } else {
                                v3 -= (^(v2))
                                var v84 int16 = 0
                                _ = v84
                            }
                        default:
                            v2 += (v2 | v3)
                        }
                        _ = len(v1)
                    }
                    v2 = v3
                    v1 = append(v1, "")
                }
                v2 -= (-(v3))
                v3 = 2147483647
            } else {
                switch "foo bar" {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + ((("" + "hello") + ((((("" + ("\n\t" + "")) + "\n\t") + "foo bar") + "foo bar") + "hello")) + ("\n\t" + ("a" + ("foo bar" + "foo bar")))))
                    }
                }
            }
            _ = fi52
        }
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
