package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 uint32) int32 {
    p0 -= 65535
    var v5 uint8 = 0
    var v6 int32 = 1
    var v7 float64 = 3.14159
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 640
}
func genFn1() uint {
    v8 := GenS1{F0: 127, F1: 42, F2: 2}
    for fi9 := 0; fi9 < 9; fi9++ {
        switch "\n\t" {
        case "\n\t":
            v10 := []int{fi9}
            _ = v10
        case "foo bar":
            v11 := GenS1{F0: 42, F1: 2, F2: 308}
            _ = v11
        case "":
            v12 := GenS2{F0: (1 == 477), F1: 414}
            _ = v12
        default:
            v13 := GenS2{F0: (-38 < -94), F1: 456}
            _ = v13
        }
        _ = fi9
    }
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 337
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    var v4 int8 = v2
    for fi14 := 0; fi14 < 3; fi14++ {
        switch fi14 {
        case 909:
            v15 := GenS2{F0: true, F1: 18446744073709551615}
            _ = v15
        default:
            var v16 uint32 = 2
            _ = v16
        }
        var v17 uint64 = 0
        _ = fi14
        _ = v17
    }
    if false {
        if false {
            v4 += (v4 & (v4 & v2))
            delete(v3, -637)
        }
    }
    if false {
        var v18 float32 = 3.14159
        _ = v18
    } else {
        v1 = 650
    }
    if (!((v1 <= v1) && false)) {
        v19 := GenS3{F0: genFn0(2147483647)}
        fmt.Println(v4)
        if (0 == 65535) {
            var v20 int = -2147483648
            v1 = 0
            _ = v20
        } else {
            v1 = v1
        }
        _ = v19
    } else {
        v1 = (v1 << ((genFn1()) & 63))
        var v21 bool = (1 < 1)
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        delete(v3, 0)
        var v23 bool = false
        delete(v3, -819)
        _ = fi22
        _ = v23
    }
    var v24 int = 9223372036854775807
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 int = fi25
        v1 += 2
        _ = fi25
        _ = v26
    }
    switch ("a" + "foo bar") {
    case "foo bar":
        v27 := GenS3{F0: 2147483647}
        _ = v27
    case "\n\t":
        switch "a" {
        case "hello":
            fmt.Println(v24)
        }
    case "":
        if true {
            if (false || true) {
                if true {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v1 = v1
                        _ = fi28
                    }
                    if (42 < 42) {
                        v2 = v4
                        for fi29 := 0; fi29 < 10; fi29++ {
                            fmt.Println(fi29)
                            if false {
                                switch v24 {
                                case 127:
                                    switch (^(v24)) {
                                    case 108:
                                        fmt.Println(v1, v1)
                                    case -2147483648:
                                        var v30 int64 = -1
                                        _ = v30
                                    default:
                                        var v31 uint = 18446744073709551615
                                        _ = v31
                                    }
                                default:
                                    if (1.5 > 0.1) {
                                        v4 += 91
                                        for fi32 := 0; fi32 < 14; fi32++ {
                                            v1 = 2
                                            _ = fi32
                                        }
                                        if true {
                                            fmt.Println(v1, fi29, v24)
                                            fmt.Println(v1, v4)
                                        }
                                    } else {
                                        v33 := GenS4{F0: 0, F1: 20}
                                        if (!(!true)) {
                                            var v34 int32 = 0
                                            if true {
                                                for fi35 := 0; fi35 < 14; fi35++ {
                                                    var v36 bool = false
                                                    _ = fi35
                                                    _ = v36
                                                }
                                                var v37 uint = 1
                                                switch (^(fi29)) {
                                                case 4294967295:
                                                    var v38 uint = (v37 + 255)
                                                    _ = v38
                                                case -199:
                                                    fmt.Println(fi29, fi29)
                                                }
                                                _ = v37
                                            } else {
                                                v39 := []int{(fi29 + -877), (-(v24))}
                                                v4 += (v4 >> ((genFn1()) & 63))
                                                _ = v39
                                            }
                                            switch ("\n\t" + "") {
                                            case "a":
                                                for fi40 := 0; fi40 < 18; fi40++ {
                                                    if true {
                                                        v2 += -1
                                                        v2 -= v4
                                                    }
                                                    if false {
                                                        fmt.Println(fi40, fi29, fi29)
                                                    } else {
                                                        v41 := []int{(fi40 ^ int(v1)), -2147483648, (fi29 * 4294967295), (fi40 - v24)}
                                                        v34 = (-(v34))
                                                        fmt.Println(v1)
                                                        _ = len(v41)
                                                        _ = v41
                                                    }
                                                    v2 += v2
                                                    _ = fi40
                                                }
                                            case "":
                                                if ((1 > 42) || (false || false)) {
                                                    v1 = v1
                                                    for fi42 := 0; fi42 < 3; fi42++ {
                                                        switch "\n\t" {
                                                        case "a":
                                                            if false {
                                                                v1 -= v1
                                                                if ((2147483647 < 0) || false) {
                                                                    _ = len(v3)
                                                                    var v43 uint = 1
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                fmt.Println(v4, fi29, v4)
                                                            }
                                                        case "foo bar":
                                                            switch -743 {
                                                            case -2147483648:
                                                                var v44 float32 = 0.1
                                                                _ = v44
                                                            case 4294967295:
                                                                fmt.Println(v24, v1, v1)
                                                            default:
                                                                var v45 bool = (!(779 != 1))
                                                                _ = v45
                                                            }
                                                        case "hello":
                                                            var v46 uint8 = 2
                                                            _ = v46
                                                        }
                                                        v2 += v4
                                                        var v47 int64 = -1
                                                        var v48 float32 = 0.0
                                                        var v49 uint8 = 127
                                                        _ = fi42
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                            }
                                            _ = v34
                                        } else {
                                            v3[32768] = 289
                                            v2 = (^(v4))
                                        }
                                        _ = v33
                                    }
                                }
                                v24 += -1
                                switch fi29 {
                                case 1:
                                    if true {
                                        v2 += v4
                                    }
                                case 127:
                                    v4 = -20
                                }
                            } else {
                                var v50 int32 = genFn0(1)
                                v4 = (v4 - 42)
                                v2 = (v4 >> ((44) & 63))
                                _ = v50
                            }
                            _ = fi29
                        }
                    } else {
                        var v51 uint16 = 300
                        for fi52 := 0; fi52 < 13; fi52++ {
                            if true {
                                var v53 int32 = 2147483647
                                _ = v53
                            }
                            var v54 bool = true
                            switch v24 {
                            case -963:
                                v55 := []int{v24}
                                _ = v55
                            case -675:
                                v56 := []int{fi52}
                                _ = v56
                            case 1:
                                var v57 uint64 = 0
                                _ = v57
                            }
                            _ = fi52
                            _ = v54
                        }
                        for fi58 := 0; fi58 < 6; fi58++ {
                            for fi59 := 0; fi59 < 9; fi59++ {
                                v51 = v51
                                var v60 uint32 = 4294967295
                                switch v24 {
                                case -1:
                                    var v61 string = ("a" + ("hello" + ("foo bar" + "\n\t")))
                                    _ = v61
                                case 1:
                                    v62 := []int{}
                                    _ = v62
                                default:
                                    var v63 string = "hello"
                                    _ = v63
                                }
                                _ = fi59
                                _ = v60
                            }
                            v1 = v51
                            for fi64 := 0; fi64 < 3; fi64++ {
                                v51 -= 731
                                fmt.Println(fi64)
                                _ = fi64
                            }
                            _ = fi58
                        }
                        _ = v51
                    }
                    switch "" {
                    case "\n\t":
                        v2 = (v2 - v4)
                    case "hello":
                        var v65 int32 = -748
                        _ = v65
                    }
                } else {
                    var v66 uint16 = v1
                    var v67 uint = 968
                    fmt.Println(v4, v4)
                    _ = v66
                    _ = v67
                }
                var v68 int64 = -1
                _ = v68
            }
            if (!(!((-9223372036854775808 >= v24) && ((-9 >= 38) || (42 == 904))))) {
                v24 = v24
            } else {
                var v69 float32 = (-2.25 + float32(0.0))
                v24 -= (v24 % (((v24 * v24)) | 1))
                v24 -= v24
                _ = v69
            }
        }
    default:
        if true {
            for fi70 := 0; fi70 < 9; fi70++ {
                for fi71 := 0; fi71 < 7; fi71++ {
                    v72 := GenS3{F0: -670}
                    v73 := GenS3{F0: -1}
                    v2 -= v2
                    _ = fi71
                    _ = v72
                    _ = v73
                }
                _ = fi70
            }
            if ((!(false || (1 != 42))) || false) {
                if ((0.1 / 1.0) >= 0.0) {
                    switch (("a" + "hello") + "hello") {
                    case "hello":
                        fmt.Println(v2)
                    case "a":
                        v74 := GenS5{F0: -1, F1: false, F2: 42}
                        _ = v74
                    case "":
                        v75 := []int{(v24 * (-(v24))), 719, (v24 << ((65535) & 63))}
                        _ = v75
                    default:
                        for fi76 := 0; fi76 < 6; fi76++ {
                            if (0.1 >= 0.0) {
                                var v77 float32 = (0.1 * 3.14159)
                                switch (("\n\t" + "a") + ("\n\t" + "foo bar")) {
                                case "hello":
                                    var v78 uint8 = 2
                                    _ = v78
                                default:
                                    v24 += 9223372036854775807
                                }
                                if ((!(!true)) && true) {
                                    v79 := GenS1{F0: 154, F1: v1, F2: 983}
                                    var v80 uint = genFn1()
                                    _ = v79
                                    _ = v80
                                }
                                _ = v77
                            }
                            _ = fi76
                        }
                    }
                    v1 = (v1 * (v1 + v1))
                    var v81 int = (v24 + (v24 - v24))
                    _ = v81
                } else {
                    v82 := GenS4{F0: 749, F1: 0}
                    var v83 int64 = int64(v24)
                    v4 += -69
                    _ = v82
                    _ = v83
                }
            } else {
                if ((0.1 - 1.5) <= 3.14159) {
                    v4 += (v2 << ((1) & 63))
                    v4 = -97
                    v1 = (v1 + (v1 | v1))
                } else {
                    if (true && true) {
                        for fi84 := 0; fi84 < 5; fi84++ {
                            _ = len(v3)
                            var v85 uint8 = 215
                            if (((v85 * 42) == uint8(v1)) || ((fi84 > fi84) || false)) {
                                var v86 uint32 = 543
                                var v87 int32 = -95
                                _ = v86
                                _ = v87
                            } else {
                                var v88 int64 = 0
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    var v89 float32 = 0.0
                                    _ = v89
                                }
                                _ = v88
                            }
                            _ = fi84
                            _ = v85
                        }
                        if (2 <= 127) {
                            v4 -= (-(v2))
                            v90 := v3[v24]
                            _ = v90
                        } else {
                            var v91 bool = ((v1 + 42) != v1)
                            var v92 int32 = 831
                            _ = v91
                            _ = v92
                        }
                        for fi93 := 0; fi93 < 2; fi93++ {
                            switch "\n\t" {
                            case "foo bar":
                                for fi94 := 0; fi94 < 18; fi94++ {
                                    var v95 float64 = (0.0 / 1.0)
                                    v96 := []int{}
                                    _ = fi94
                                    _ = v95
                                    _ = v96
                                }
                            case "":
                                var v97 int = (fi93 % ((9223372036854775807) | 1))
                                _ = v97
                            }
                            v4 += v2
                            _ = fi93
                        }
                    } else {
                        var v98 uint64 = 42
                        v99 := []int{v24}
                        _ = v98
                        _ = v99
                    }
                    v24 += 127
                    for fi100 := 0; fi100 < 14; fi100++ {
                        v2 = v2
                        _ = fi100
                    }
                }
                v3[v24] = 316
            }
        } else {
            for fi101 := 0; fi101 < 6; fi101++ {
                v24 = (^(v24))
                var v102 uint64 = 1
                v102 += uint64(v1)
                _ = fi101
                _ = v102
            }
            v4 += v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
