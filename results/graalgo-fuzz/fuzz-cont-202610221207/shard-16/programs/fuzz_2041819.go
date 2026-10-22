package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int; F1 string; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int32; F2 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 int16 }
func genFn0(p0 uint8, p1 int, p2 uint8) uint16 {
    var v4 int32 = 0
    p0 -= p0
    switch "foo bar" {
    case "hello":
        v4 = v4
    case "":
        switch 850 {
        case -514:
            v5 := []int{-1, p1}
            _ = v5
        case 540:
            fmt.Println(p0, p1, p2)
        default:
            fmt.Println(p0, v4)
        }
    case "foo bar":
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: fi6, F1: (42 >= int8(p2)), F2: 2}
            p1 += -641
            v8 := []int{fi6, (p1 >> ((644) & 63))}
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        if true {
            v9 := []int{p1, (p1 ^ (^(p1)))}
            var v10 string = "a"
            if true {
                var v11 uint16 = 42
                var v12 int16 = -32768
                _ = v11
                _ = v12
            } else {
                switch v10 {
                case "a":
                    v13 := GenS2{F0: -1, F1: 42, F2: 36}
                    _ = v13
                case "":
                    switch (p1 & -855) {
                    case -2147483648:
                        if false {
                            var v14 int32 = v4
                            _ = v14
                        }
                    case 1:
                        v15 := GenS3{F0: "foo bar"}
                        _ = v15
                    }
                case "hello":
                    if (((p0 << ((2) & 63)) == p0) || (false || (!true))) {
                        for fi16 := 0; fi16 < 13; fi16++ {
                            _ = len(v9)
                            switch v10 {
                            case "\n\t":
                                for fi17 := 0; fi17 < 15; fi17++ {
                                    if ((v4 ^ -1) <= v4) {
                                        var v18 uint8 = (p2 << ((0) & 63))
                                        var v19 string = ("foo bar" + ("hello" + ((v10 + "") + "foo bar")))
                                        _ = len(v9)
                                        fmt.Println(fi17, p2, fi17)
                                        p0 -= (p2 * (p0 >> ((18446744073709551615) & 63)))
                                        _ = v18
                                        _ = v19
                                    } else {
                                        if (!(!false)) {
                                            switch p1 {
                                            case 708:
                                                var v20 float64 = 3.14159
                                                _ = v20
                                            case -9223372036854775808:
                                                var v21 int = -271
                                                _ = v21
                                            default:
                                                var v22 uint8 = 2
                                                _ = v22
                                            }
                                            var v23 uint16 = 2
                                            var v24 uint64 = 42
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                    _ = fi17
                                }
                            case "":
                                for fi25 := 0; fi25 < 16; fi25++ {
                                    fmt.Println(v10)
                                    _ = fi25
                                }
                            default:
                                fmt.Println(p0, v10)
                            }
                            _ = fi16
                        }
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v27 := GenS4{F0: (fi26 & p1), F1: 42}
                            var v28 float32 = 0.0
                            var v29 int32 = v4
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                }
            }
            _ = v10
            _ = v9
        }
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 float32 = float32((float64(0.1) / 1.5))
        for fi32 := 0; fi32 < 5; fi32++ {
            v33 := GenS2{F0: -1, F1: 261, F2: 255}
            var v34 float64 = (3.14159 / 0.1)
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    v35 := GenS2{F0: 931, F1: 4294967295, F2: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v4
    return uint16(p0)
}
func genFn1() bool {
    v36 := GenS5{F0: -339, F1: ("foo bar" + "foo bar"), F2: 9223372036854775807}
    v37 := GenS6{F0: 255}
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := GenS6{F0: 42}
        if false {
            var v40 uint = 2
            var v41 int = 32768
            var v42 string = "a"
            _ = v40
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    if true {
        v43 := GenS3{F0: ("\n\t" + "a")}
        var v44 bool = (0.0 <= 0.1)
        _ = v43
        _ = v44
    }
    if (0 < 7) {
        v45 := []int{32767, 42, -97}
        _ = len(v45)
        var v46 string = ("\n\t" + (("a" + "\n\t") + "foo bar"))
        _ = v45
        _ = v46
    } else {
        var v47 int = 0
        v48 := GenS3{F0: ("foo bar" + ("\n\t" + "\n\t"))}
        _ = v47
        _ = v48
    }
    _ = v36
    _ = v37
    return true
}

func main() {
    var v1 uint64 = 0
    v2 := []string{}
    var v3 uint = 42
    if (9223372036854775807 == -9223372036854775808) {
        fmt.Println(v1)
    }
    v49 := GenS7{F0: false, F1: 2147483647, F2: -525}
    v3 += v3
    v50 := GenS5{F0: -9223372036854775808, F1: "\n\t", F2: 9223372036854775807}
    for fi51 := 0; fi51 < 18; fi51++ {
        v1 -= 18446744073709551615
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi51
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v3 += (v3 - (v3 << ((v3) & 63)))
    for fi52 := 0; fi52 < 8; fi52++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi52
    }
    v53 := GenS8{F0: (1.5 != (0.1 / (-2.25 + ((0.1 + (1.5 - -2.25)) + 0.1))))}
    switch "foo bar" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "hello")
        }
    case "":
        if len(v2) > 0 {
            v2[0] = ("" + "a")
        }
    }
    if (false && ("hello" > "\n\t")) {
        switch ("foo bar" + "hello") {
        case "\n\t":
            v3 = (v3 % (((v3 << ((v3) & 63))) | 1))
        case "a":
            if true {
                for fi54 := 0; fi54 < 17; fi54++ {
                    fmt.Println(v3, fi54)
                    for fi55 := 0; fi55 < 11; fi55++ {
                        switch (fi55 % ((fi55) | 1)) {
                        case 158:
                            v3 = v3
                        default:
                            var v56 int64 = -9223372036854775808
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v57 int8 = 100
                            var v58 string = ("hello" + "foo bar")
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v59 int64 = 9223372036854775807
                            _ = v56
                            _ = v57
                            _ = v58
                            _ = v59
                        }
                        var v60 uint8 = 227
                        _ = fi55
                        _ = v60
                    }
                    _ = fi54
                }
                var v61 bool = true
                for fi62 := 0; fi62 < 10; fi62++ {
                    if false {
                        var v63 int64 = 859
                        if ((0.1 / 1.0) >= (1e10 + 0.1)) {
                            v1 = (^(v1))
                            v3 += uint(v1)
                        }
                        _ = len(v2)
                        _ = v63
                    } else {
                        var v64 uint = (v3 / ((v3) | 1))
                        if v61 {
                            var v65 uint64 = (v1 | (v1 + v1))
                            _ = v65
                        } else {
                            if v61 {
                                for fi66 := 0; fi66 < 19; fi66++ {
                                    if len(v2) > 0 {
                                        v2[0] = ("" + ("a" + (("a" + "foo bar") + (("" + "hello") + ""))))
                                    }
                                    _ = fi66
                                }
                                if false {
                                    v67 := GenS6{F0: (^(v64))}
                                    v64 -= v64
                                    fmt.Println(v3, fi62)
                                    _ = v67
                                }
                            } else {
                                v3 += v64
                                v61 = v61
                                var v68 uint = v64
                                _ = v68
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (("a" + (("" + "a") + "\n\t")) + ("hello" + "\n\t"))
                        }
                        _ = v64
                    }
                    fmt.Println(fi62, v61, v61)
                    _ = fi62
                }
                _ = v61
            } else {
                if (-2.25 > ((float64(0.0) + 1.5) + (1.5 * 1e10))) {
                    for fi69 := 0; fi69 < 6; fi69++ {
                        for fi70 := 0; fi70 < 2; fi70++ {
                            if (65535 == 42) {
                                var v71 float32 = float32(-2.25)
                                v71 -= float32(0.0)
                                var v72 float32 = ((1.5 - (3.14159 / ((v71 - 1.5) / v71))) / v71)
                                if false {
                                    var v73 string = (("\n\t" + "hello") + "foo bar")
                                    v74 := GenS1{F0: (fi70 * 2147483647), F1: (1000 != -263), F2: 42}
                                    var v75 int64 = -831
                                    if (-2.25 <= 1.5) {
                                        v76 := GenS4{F0: 1, F1: 786}
                                        _ = v76
                                    }
                                    var v77 uint16 = genFn0(231, fi69, 128)
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                    _ = v77
                                }
                                _ = v71
                                _ = v72
                            } else {
                                fmt.Println(fi69)
                            }
                            _ = fi70
                        }
                        _ = fi69
                    }
                }
            }
        }
        fmt.Println(v1)
    }
    var v78 uint8 = 42
    v78 += v78
    fmt.Println(v3, v3, v1)
    var v79 uint64 = 593
    for fi80 := 0; fi80 < 7; fi80++ {
        v81 := GenS7{F0: false, F1: 1, F2: -2147483648}
        v78 += v78
        v82 := GenS9{F0: 564, F1: 1}
        _ = fi80
        _ = v81
        _ = v82
    }
    var v83 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v53
    _ = v78
    _ = v79
    _ = v83
}
