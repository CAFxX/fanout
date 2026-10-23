package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint32; F2 int }
type GenS7 struct { F0 uint16; F1 int; F2 uint8 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 bool; F2 int16 }
type GenS12 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 int32) int16 {
    var v6 uint16 = 445
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 9223372036854775807
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 141, F1: v6}
        _ = fi9
        _ = v10
    }
    v11 := []int{127, -2147483648, 814, 810}
    _ = p0
    _ = v11
    _ = v6
    return -196
}
func genFn1(p0 int8) int16 {
    p0 -= p0
    switch "" {
    case "":
        switch ("" + "a") {
        case "hello":
            if (0 <= 1) {
                if false {
                    fmt.Println(p0)
                    v12 := []int{}
                    _ = v12
                }
            } else {
                v13 := GenS2{F0: -2147483648, F1: 1}
                _ = v13
            }
        case "a":
            if (450 > 42) {
                var v14 string = ("foo bar" + "\n\t")
                if (!true) {
                    if false {
                        var v15 int8 = p0
                        fmt.Println(p0, p0, v14)
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 bool = false
                            v18 := GenS3{F0: -771, F1: v17, F2: 65535}
                            var v19 int8 = (p0 ^ p0)
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = v15
                    }
                    p0 = 42
                }
                v20 := []int{127, 2147483647, -96, 32767}
                _ = v14
                _ = v20
            }
        }
    case "foo bar":
        v21 := []int{42, 32767, 878}
        _ = v21
    }
    var v22 uint32 = 2147483647
    var v23 uint = 65535
    _ = p0
    _ = v22
    _ = v23
    return -393
}

func main() {
    var v1 int64 = -151
    v2 := make(map[uint32]int)
    var v3 string = ""
    var v4 uint16 = 0
    v5 := make(map[uint32]uint32)
    var v24 uint64 = 191
    if ((true || true) || true) {
        v4 = uint16(v24)
        if (v3 == (v3 + v3)) {
            v25 := GenS4{F0: v3, F1: 0, F2: (v1 - -1)}
            v1 = (v1 ^ v1)
            _ = v25
        } else {
            for fi26 := 0; fi26 < 9; fi26++ {
                switch fi26 {
                case 128:
                    var v27 uint8 = 20
                    _ = v27
                case 32768:
                    v24 = (^(v24))
                default:
                    v28 := GenS5{F0: (fi26 & 32767)}
                    _ = v28
                }
                _ = len(v5)
                for fi29 := 0; fi29 < 18; fi29++ {
                    if false {
                        if true {
                            v24 = 2
                        }
                        var v30 int8 = -46
                        var v31 int = fi29
                        _ = v30
                        _ = v31
                    } else {
                        if (1e10 != -2.25) {
                            var v32 string = v3
                            v33 := GenS6{F0: ("\n\t" + v32), F1: uint32(v4), F2: (fi26 << ((65535) & 63))}
                            _ = v32
                            _ = v33
                        }
                    }
                    v4 += 1
                    _ = fi29
                }
                _ = fi26
            }
            v34 := GenS6{F0: (v3 + v3), F1: 2147483648, F2: 128}
            v24 -= v24
            _ = v34
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            if true {
                v36 := GenS6{F0: v3, F1: 157, F2: (fi35 & (-(fi35)))}
                if (-2.25 >= 1.5) {
                    if (!(!(((((v3 + "foo bar") + (v3 + "")) + "foo bar") + ((((v3 + ((v3 + "") + "\n\t")) + (("a" + ("foo bar" + "\n\t")) + "hello")) + v3) + v3)) > v3))) {
                        for fi37 := 0; fi37 < 13; fi37++ {
                            var v38 uint = 18446744073709551615
                            v39 := GenS7{F0: (^(v4)), F1: fi35, F2: 255}
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    } else {
                        _ = len(v5)
                        var v40 int16 = int16(fi35)
                        _ = v40
                    }
                    v24 = (v24 ^ (v24 - 0))
                    var v41 string = v3
                    _ = v41
                } else {
                    v5[2] = 4294967295
                    for fi42 := 0; fi42 < 14; fi42++ {
                        v24 += v24
                        if false {
                            switch "hello" {
                            case "\n\t":
                                v4 -= v4
                                var v43 int = (fi42 & (^(fi35)))
                                v24 = (v24 << ((0) & 63))
                                var v44 int16 = 41
                                _ = v43
                                _ = v44
                            default:
                                switch (v3 + v3) {
                                case "\n\t":
                                    if false {
                                        switch (fi42 - fi42) {
                                        case 42:
                                            if (((128 < 255) || (false || false)) || false) {
                                                delete(v2, 65535)
                                            }
                                        default:
                                            v3 = "\n\t"
                                        }
                                        v1 = v1
                                        switch (fi35 << ((255) & 63)) {
                                        case 667:
                                            switch ((("a" + "foo bar") + ((v3 + "foo bar") + ((((v3 + v3) + "foo bar") + v3) + v3))) + "hello") {
                                            case "hello":
                                                switch (fi35 & (fi42 % ((99) | 1))) {
                                                case 946:
                                                    v4 -= (^(v4))
                                                case 128:
                                                    var v45 int32 = -1
                                                    _ = v45
                                                }
                                            case "\n\t":
                                                if false {
                                                    switch v3 {
                                                    case "a":
                                                        fmt.Println(v3, fi35, v3)
                                                    default:
                                                        v3 = (v3 + v3)
                                                        var v46 uint8 = 255
                                                        _ = v46
                                                    }
                                                    switch v3 {
                                                    case "hello":
                                                        if (!true) {
                                                            v1 = -1
                                                            if false {
                                                                if false {
                                                                    v4 = v4
                                                                }
                                                            }
                                                        } else {
                                                            v47 := GenS2{F0: fi42, F1: 986}
                                                            _ = v47
                                                        }
                                                    case "foo bar":
                                                        var v48 float64 = 0.0
                                                        _ = v48
                                                    case "a":
                                                        var v49 uint64 = v24
                                                        _ = v49
                                                    default:
                                                        v4 = (v4 / ((v4) | 1))
                                                    }
                                                    v50 := GenS4{F0: v3, F1: 4294967295, F2: v1}
                                                    _ = v50
                                                }
                                            case "foo bar":
                                                var v51 float32 = 1e10
                                                _ = v51
                                            }
                                        case 4294967295:
                                            fmt.Println(v4)
                                        }
                                    }
                                case "":
                                    var v52 int32 = 1
                                    _ = v52
                                }
                            }
                        } else {
                            v53 := GenS3{F0: v1, F1: (191 != 222), F2: v4}
                            v24 = (^(v24))
                            _ = v53
                        }
                        v3 = "\n\t"
                        fmt.Println(v3, v24, fi42)
                        _ = fi42
                    }
                }
                _ = v36
            }
            if (uint16(v24) >= (v4 ^ (v4 * 0))) {
                v3 = ((v3 + v3) + (("\n\t" + ((((v3 + "hello") + "a") + v3) + v3)) + ("" + ("a" + v3))))
                switch (-(fi35)) {
                case 32767:
                    for fi54 := 0; fi54 < 11; fi54++ {
                        switch fi35 {
                        case 32768:
                            var v55 uint32 = 779
                            _ = v55
                        case -1:
                            var v56 uint32 = 2
                            _ = v56
                        default:
                            v57 := GenS2{F0: 1, F1: (v4 - v4)}
                            _ = v57
                        }
                        switch v3 {
                        case "a":
                            v58 := GenS7{F0: (^(v4)), F1: (-(fi54)), F2: 205}
                            _ = v58
                        case "hello":
                            v2[0] = fi54
                        }
                        v24 = (v24 >> ((18446744073709551615) & 63))
                        _ = fi54
                    }
                default:
                    v59 := []int{(fi35 | fi35), int(v24)}
                    _ = v59
                }
                for fi60 := 0; fi60 < 2; fi60++ {
                    fmt.Println(v3, fi35)
                    v24 -= 2
                    switch ("\n\t" + v3) {
                    case "hello":
                        fmt.Println(fi60, v1, v24)
                    case "\n\t":
                        v61 := []int{1, -578}
                        _ = v61
                    case "a":
                        v62 := GenS8{F0: (v24 << ((1) & 63)), F1: genFn1(0)}
                        if false {
                            var v63 int64 = v1
                            var v64 int8 = 13
                            switch "" {
                            case "":
                                var v65 int = fi60
                                _ = v65
                            }
                            _ = v63
                            _ = v64
                        } else {
                            v4 = 708
                            fmt.Println(v1, v1)
                        }
                        _ = v62
                    default:
                        if true {
                            v66 := GenS9{F0: (v24 | uint64(v4)), F1: true}
                            _ = v66
                        }
                    }
                    _ = fi60
                }
            } else {
                v1 -= v1
                var v67 int16 = genFn0(1)
                _ = v67
            }
            _ = fi35
        }
    }
    var v68 int64 = (v1 >> ((2) & 63))
    if ((-70 <= 578) && true) {
        v3 = v3
        v69 := GenS10{F0: -1}
        _ = v69
    } else {
        delete(v5, 2147483648)
        v70 := GenS11{F0: -32768, F1: false, F2: 42}
        _ = v70
    }
    fmt.Println(v68)
    var v71 uint64 = v24
    if (65535 >= 18446744073709551615) {
        for fi72 := 0; fi72 < 4; fi72++ {
            v73 := GenS3{F0: (-(v1)), F1: false, F2: v4}
            _ = fi72
            _ = v73
        }
    }
    v3 = (v3 + v3)
    fmt.Println(v3)
    if (101 == -49) {
        var v74 string = v3
        _ = v74
    } else {
        v71 = (v71 >> ((4294967295) & 63))
        v75 := []int{32767, -723, -1, 671}
        var v76 bool = false
        _ = v75
        _ = v76
    }
    switch (v3 + "hello") {
    case "a":
        var v77 uint16 = 0
        _ = v77
    case "hello":
        for fi78 := 0; fi78 < 12; fi78++ {
            delete(v2, 707)
            _ = fi78
        }
    case "":
        v79 := GenS12{F0: 42, F1: -1, F2: (v71 & (v71 >> ((878) & 63)))}
        _ = v79
    default:
        var v80 float32 = 1e10
        _ = v80
    }
    v1 = (v1 / (((v68 - v68)) | 1))
    var v81 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v71
    _ = v81
}
