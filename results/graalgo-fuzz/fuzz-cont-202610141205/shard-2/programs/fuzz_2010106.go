package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint16; F1 int32; F2 int16 }
func genFn0() uint64 {
    var v5 uint = 28
    fmt.Println(v5)
    v6 := []int{0, 32768}
    v7 := GenS1{F0: 578, F1: 0}
    if len(v6) > 0 {
        v6[0] = 934
    }
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int) int32 {
    var v8 float64 = 0.0
    v9 := GenS2{F0: 32767}
    var v10 float32 = 1.5
    v11 := []int{}
    if (p0 >= p0) {
        for fi12 := 0; fi12 < 16; fi12++ {
            _ = len(v11)
            _ = fi12
        }
        fmt.Println(v8)
    } else {
        if ((-471 != p0) || false) {
            v13 := GenS1{F0: 2147483648, F1: 1}
            v8 = float64((v10 * v10))
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(v8, v10)
            if len(v11) > 0 {
                v11[0] = 489
            }
            _ = fi14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                if false {
                    if (false || (!false)) {
                        p0 = 127
                    } else {
                        v10 = v10
                        switch fi16 {
                        case 0:
                            if len(v11) > 0 {
                                v11[0] = (p0 >> ((197) & 63))
                            }
                        case 32768:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        default:
                            var v17 uint = 927
                            _ = v17
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v18 uint16 = 65535
                            _ = v18
                        default:
                            v19 := GenS3{F0: genFn0()}
                            _ = v19
                        }
                    }
                }
                v11 = append(v11, 42)
                _ = fi16
            }
            _ = fi15
        }
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{590, 2, 2147483647, 2147483647}
    var v2 int16 = -1
    var v3 uint64 = 0
    v4 := make(map[uint32]string)
    v20 := v4[65535]
    fmt.Println(v2)
    if (65535 < 65535) {
        var v21 int = 717
        v21 = (v21 - v21)
        for fi22 := 0; fi22 < 12; fi22++ {
            fmt.Println(fi22, fi22)
            fmt.Println(v3)
            var v23 int64 = -473
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        v24 := GenS1{F0: 42, F1: 1}
        _ = v24
    }
    var v25 uint32 = 4294967295
    if ((0.1 == 0.1) && (2 == 1)) {
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 float32 = (-2.25 + 1.5)
            v2 = v2
            if (false || true) {
                for fi28 := 0; fi28 < 13; fi28++ {
                    fmt.Println(v20, v27)
                    switch (fi28 ^ fi28) {
                    case -1:
                        var v29 uint16 = 32767
                        _ = v29
                    case 1:
                        if true {
                            fmt.Println(v2, v3, v27)
                            var v30 int = fi26
                            _ = v30
                        } else {
                            if (65535 > 0) {
                                v31 := GenS4{F0: -1, F1: true}
                                var v32 float32 = 0.1
                                _ = v31
                                _ = v32
                            } else {
                                v27 = v27
                                var v33 int16 = (-(v2))
                                _ = v33
                            }
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi28
                }
            } else {
                v27 = (v27 / 0.1)
                if true {
                    v27 += v27
                    _ = len(v4)
                    v34 := GenS2{F0: 2}
                    _ = v34
                } else {
                    var v35 string = (v20 + "\n\t")
                    v36 := v4[(v25 >> ((853) & 63))]
                    v37 := GenS4{F0: int(v3), F1: (1 > 255)}
                    _ = v35
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 17; fi38++ {
                    switch (fi38 << ((4294967295) & 63)) {
                    case 127:
                        if false {
                            var v39 uint16 = 1
                            v40 := GenS3{F0: (^(v3))}
                            v41 := GenS5{F0: v2, F1: 127}
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    case 0:
                        if true {
                            var v42 bool = (!(v2 <= (v2 - v2)))
                            v20 = ("\n\t" + "foo bar")
                            v2 = v2
                            v20 = v20
                            _ = v42
                        }
                        v20 = (("" + v20) + v20)
                    default:
                        if (1 < 2) {
                            v1 = append(v1, (v25 * 2147483647))
                            v43 := GenS6{F0: 0}
                            _ = v43
                        }
                    }
                    var v44 int8 = 42
                    _ = fi38
                    _ = v44
                }
            }
            _ = fi26
            _ = v27
        }
        fmt.Println(v20)
    } else {
        v45 := GenS3{F0: v3}
        for fi46 := 0; fi46 < 6; fi46++ {
            v47 := v4[v25]
            v20 = "\n\t"
            fmt.Println(fi46)
            _ = fi46
            _ = v47
        }
        switch 128 {
        case 331:
            if (1 == 1) {
                var v48 int32 = -1
                _ = v48
            } else {
                var v49 uint64 = 1
                for fi50 := 0; fi50 < 15; fi50++ {
                    for fi51 := 0; fi51 < 17; fi51++ {
                        v3 += uint64(v25)
                        switch -1 {
                        case -1:
                            if false {
                                v1 = append(v1, (v25 + (v25 | 2)))
                                var v52 int8 = -128
                                switch fi51 {
                                case 127:
                                    if true {
                                        _ = len(v4)
                                        var v53 uint8 = 127
                                        _ = v53
                                    } else {
                                        var v54 float32 = float32(1e10)
                                        var v55 uint64 = 375
                                        v2 -= v2
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v56 uint8 = 42
                                            var v57 int32 = -414
                                            var v58 float32 = v54
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        } else {
                                            if false {
                                                v59 := GenS7{F0: 2, F1: (v52 + (v52 | (v52 - v52)))}
                                                _ = v59
                                            }
                                            var v60 uint8 = 127
                                            _ = v60
                                        }
                                        switch (fi51 | -227) {
                                        case 199:
                                            v20 = (v20 + (v20 + "foo bar"))
                                        case 539:
                                            if (((1.5 * v54) * float32(3.14159)) != v54) {
                                                v3 = v49
                                                if len(v1) > 0 {
                                                    v1[0] = 910
                                                }
                                            } else {
                                                if true {
                                                    fmt.Println(v2, fi51, v55)
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(fi51, fi51, v25)
                                        default:
                                            switch fi51 {
                                            case 42:
                                                fmt.Println(v25, v3, v55)
                                            case 128:
                                                switch (-(fi50)) {
                                                case 379:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case 1:
                                                    if len(v1) > 0 {
                                                        v1[0] = v25
                                                    }
                                                }
                                            case 390:
                                                var v61 int32 = int32(fi51)
                                                _ = v61
                                            }
                                        }
                                        v20 = "a"
                                        fmt.Println(v3, fi51, v52)
                                        if len(v1) > 0 {
                                            v1[0] = v25
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                case 1:
                                    v4[v25] = v20
                                case 9223372036854775807:
                                    if (3.14159 < 1.5) {
                                        if ((248 == 0) || false) {
                                            var v62 string = "foo bar"
                                            v20 = "a"
                                            _ = v62
                                        } else {
                                            v63 := v4[v25]
                                            var v64 float32 = (3.14159 / (float32(1.5) + 0.1))
                                            _ = v63
                                            _ = v64
                                        }
                                        if false {
                                            v4[v25] = v20
                                            fmt.Println(v49)
                                        }
                                    } else {
                                        if true {
                                            v2 = 1000
                                            if true {
                                                switch fi51 {
                                                case 688:
                                                    if true {
                                                        delete(v4, 1)
                                                        var v65 int8 = (v52 * 86)
                                                        _ = v65
                                                    }
                                                case 127:
                                                    switch ((v20 + "") + "\n\t") {
                                                    case "\n\t":
                                                        v49 = 4294967295
                                                    case "foo bar":
                                                        var v66 uint32 = 2
                                                        _ = v66
                                                    default:
                                                        v67 := GenS1{F0: (v25 + 1), F1: 442}
                                                        _ = v67
                                                    }
                                                }
                                            } else {
                                                var v68 string = "\n\t"
                                                if false {
                                                    if false {
                                                        v69 := GenS2{F0: 255}
                                                        fmt.Println(v49, fi50, fi50)
                                                        if ((v49 & 385) <= (v3 % ((v49) | 1))) {
                                                            if true {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        }
                                                        _ = v69
                                                    } else {
                                                        switch fi50 {
                                                        case 9223372036854775807:
                                                            switch (fi51 % ((-41) | 1)) {
                                                            case -338:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case 480:
                                                            _ = len(v4)
                                                        }
                                                        if false {
                                                            var v70 int16 = 42
                                                            _ = v70
                                                        } else {
                                                            var v71 uint16 = 42
                                                            _ = v71
                                                        }
                                                        v1 = append(v1, 4294967295)
                                                    }
                                                    if (false || (((127 > 42) || true) && false)) {
                                                        v68 = v20
                                                        if (0.1 < ((1.5 - 0.0) / 1e10)) {
                                                            v3 -= (^(v3))
                                                        }
                                                        switch fi50 {
                                                        case 9223372036854775807:
                                                            var v72 int8 = (v52 * v52)
                                                            _ = v72
                                                        case 171:
                                                            _ = len(v1)
                                                        }
                                                    } else {
                                                        var v73 int8 = (v52 * (v52 << ((uint(v3)) & 63)))
                                                        _ = v73
                                                    }
                                                } else {
                                                    if true {
                                                        v4[42] = "foo bar"
                                                        var v74 uint16 = 65535
                                                        _ = v74
                                                    }
                                                    v25 -= 2147483647
                                                }
                                                _ = v68
                                            }
                                        } else {
                                            fmt.Println(v49)
                                            v75 := GenS8{F0: 128, F1: v2}
                                            _ = v75
                                        }
                                    }
                                }
                                _ = v52
                            }
                            _ = len(v1)
                            v20 = v20
                            delete(v4, (v25 + (v25 - (v25 >> ((378) & 63)))))
                            v76 := GenS8{F0: 42, F1: (v2 - (v2 >> ((649) & 63)))}
                            _ = v76
                        default:
                            if true {
                                switch (v20 + "") {
                                case "hello":
                                    v77 := GenS1{F0: v25, F1: 590}
                                    fmt.Println(v3, v20)
                                    var v78 string = (v20 + v20)
                                    _ = v77
                                    _ = v78
                                case "\n\t":
                                    var v79 uint32 = v25
                                    _ = v79
                                case "":
                                    if false {
                                        fmt.Println(v49, v20, fi50)
                                    } else {
                                        var v80 int = -2147483648
                                        v81 := v4[(v25 | v25)]
                                        if false {
                                            _ = len(v4)
                                            fmt.Println(v81, v25)
                                        }
                                        _ = v80
                                        _ = v81
                                    }
                                }
                                var v82 string = "hello"
                                _ = v82
                            } else {
                                fmt.Println(v25, v25)
                                v3 -= v49
                                if true {
                                    if false {
                                        v2 = v2
                                        var v83 bool = false
                                        v84 := GenS5{F0: (v2 / (((v2 * -850)) | 1)), F1: 1}
                                        _ = v83
                                        _ = v84
                                    }
                                    var v85 uint64 = 42
                                    fmt.Println(v3, v25)
                                    _ = v85
                                } else {
                                    v20 = v20
                                    v86 := GenS5{F0: (v2 - -32768), F1: 100}
                                    _ = v86
                                }
                            }
                        }
                        var v87 uint64 = v49
                        if (!(971 >= (fi50 | fi50))) {
                            fmt.Println(fi50, fi50, v87)
                            v88 := GenS1{F0: (v25 >> ((1) & 63)), F1: 1}
                            _ = v88
                        }
                        _ = fi51
                        _ = v87
                    }
                    _ = fi50
                }
                var v89 bool = (100 <= 79)
                _ = v49
                _ = v89
            }
        case -2147483648:
            v3 -= genFn0()
        case -157:
            fmt.Println(v20, v20)
        default:
            v90 := GenS3{F0: genFn0()}
            _ = v90
        }
        _ = v45
    }
    switch v20 {
    case "":
        for fi91 := 0; fi91 < 17; fi91++ {
            var v92 float32 = 1.5
            _ = fi91
            _ = v92
        }
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        v2 -= v2
    }
    if false {
        if (42 < -462) {
            var v93 int16 = v2
            v20 = v20
            _ = v93
        } else {
            v20 = ((("a" + ((v20 + (((v20 + "") + v20) + (((v20 + (v20 + "hello")) + "foo bar") + v20))) + (((((((v20 + v20) + "a") + (v20 + (v20 + v20))) + (v20 + (v20 + v20))) + ((v20 + v20) + v20)) + ("hello" + v20)) + ""))) + "") + ((v20 + v20) + "\n\t"))
        }
        for fi94 := 0; fi94 < 10; fi94++ {
            for fi95 := 0; fi95 < 1; fi95++ {
                var v96 string = ((((("foo bar" + ("foo bar" + "foo bar")) + v20) + v20) + v20) + "hello")
                v97 := GenS9{F0: v3, F1: v3, F2: -1}
                _ = fi95
                _ = v96
                _ = v97
            }
            v98 := GenS4{F0: -703, F1: true}
            v99 := GenS9{F0: (v3 / ((v3) | 1)), F1: genFn0(), F2: 0}
            _ = fi94
            _ = v98
            _ = v99
        }
    }
    var v100 float32 = float32(0.1)
    v2 = v2
    switch 0 {
    case -2147483648:
        v101 := GenS2{F0: 615}
        _ = v101
    case 4294967295:
        if false {
            v102 := GenS10{F0: 1, F1: -1, F2: -28}
            v3 += 926
            _ = v102
        }
    default:
        if false {
            delete(v4, v25)
            if len(v1) > 0 {
                v1[0] = (v25 ^ 0)
            }
            v3 += (v3 & (v3 ^ 349))
        } else {
            if true {
                v103 := GenS1{F0: (^(v25)), F1: -907}
                _ = v103
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v100
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
