package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16; F1 int8 }
type GenS12 struct { F0 bool; F1 int32 }
type GenS13 struct { F0 uint32; F1 int16; F2 string }

func main() {
    var v1 int8 = -36
    v2 := make(map[uint32]string)
    var v3 uint = 18446744073709551615
    v4 := []int{652, 32768, 128}
    v5 := make(map[string]string)
    var v6 uint64 = 693
    fmt.Println(v1)
    if true {
        var v7 uint = (v3 & (^(v3)))
        delete(v2, 4294967295)
        _ = v7
    } else {
        v6 = (^(v6))
        if (-1 > 2147483647) {
            v1 -= 127
            if false {
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 int8 = 72
                    var v10 string = ""
                    for fi11 := 0; fi11 < 1; fi11++ {
                        switch (v10 + v10) {
                        case "\n\t":
                            var v12 uint8 = 128
                            _ = v12
                        case "a":
                            v9 += v9
                            if (42 < 32767) {
                                var v13 float64 = -2.25
                                var v14 string = v10
                                _ = v13
                                _ = v14
                            }
                            var v15 bool = (!true)
                            v1 += 42
                            _ = v15
                        }
                        delete(v5, "foo bar")
                        switch fi8 {
                        case -1:
                            if true {
                                v3 = (v3 + v3)
                            }
                        }
                        _ = fi11
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                if (1 != 1) {
                    v16 := GenS1{F0: uint(v6), F1: 1}
                    _ = v16
                }
                v17 := GenS2{F0: v3, F1: -1, F2: 42}
                _ = v17
            } else {
                v18 := GenS3{F0: (!(3.14159 != 1e10)), F1: uint16(v3)}
                fmt.Println(v3, v1)
                _ = v18
            }
            v19 := GenS4{F0: 9223372036854775807, F1: v6}
            _ = v19
        }
        v1 += v1
    }
    v20 := GenS5{F0: 42, F1: 42, F2: 255}
    v21 := GenS6{F0: 1, F1: 2147483647, F2: v6}
    var v22 float64 = 0.0
    var v23 uint = (v3 - 4294967295)
    if false {
        for fi24 := 0; fi24 < 12; fi24++ {
            if false {
                v25 := v2[317]
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS5{F0: 65535, F1: 65535, F2: 250}
                    fmt.Println(v22)
                    v28 := GenS6{F0: 176, F1: 42, F2: 18446744073709551615}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                switch 32767 {
                case -2147483648:
                    var v29 int64 = -472
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        switch v25 {
                        case "foo bar":
                            if true {
                                v31 := GenS5{F0: 492, F1: 65535, F2: 255}
                                v25 = v25
                                _ = v31
                            }
                        case "\n\t":
                            var v32 string = "\n\t"
                            _ = v32
                        case "a":
                            v3 -= v3
                        }
                        _ = fi30
                    }
                }
                _ = v25
            }
            _ = fi24
        }
        v33 := GenS4{F0: -1, F1: 0}
        for fi34 := 0; fi34 < 11; fi34++ {
            if (!(!true)) {
                v3 = 536
            } else {
                var v35 uint = 4294967295
                v3 -= v3
                for fi36 := 0; fi36 < 9; fi36++ {
                    if true {
                        v23 = v35
                    } else {
                        var v37 int = 989
                        if true {
                            v38 := GenS7{F0: v1, F1: (42 > 4294967295)}
                            _ = v38
                        }
                        if ((v6 << ((4294967295) & 63)) >= (v6 & v6)) {
                            v22 = ((v22 * v22) - ((v22 * (((float64(-2.25) + (0.1 + v22)) * v22) + (v22 - v22))) / v22))
                            fmt.Println(v1, v3)
                            v39 := GenS2{F0: (^(v3)), F1: -1, F2: -1}
                            _ = v39
                        }
                        _ = v37
                    }
                    v3 = (v35 - (v23 - (v23 << (((v35 & v3)) & 63))))
                    _ = fi36
                }
                _ = v35
            }
            var v40 int8 = v1
            _ = fi34
            _ = v40
        }
        _ = v33
    } else {
        var v41 uint64 = 829
        switch ("foo bar" + ((("a" + ("" + "hello")) + (("hello" + (("\n\t" + "hello") + "foo bar")) + "")) + ("hello" + "foo bar"))) {
        case "hello":
            for fi42 := 0; fi42 < 1; fi42++ {
                v3 = 103
                for fi43 := 0; fi43 < 6; fi43++ {
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        switch fi43 {
                        case 127:
                            v44 := GenS2{F0: v3, F1: -9223372036854775808, F2: -32768}
                            var v45 uint64 = (v6 % ((v6) | 1))
                            v22 = 3.14159
                            _ = v44
                            _ = v45
                        case -9223372036854775808:
                            v1 = v1
                        default:
                            if ((-(v1)) >= v1) {
                                fmt.Println(v22)
                                if len(v4) > 0 {
                                    v4[0] = int(v6)
                                }
                                var v46 int32 = 1
                                _ = v46
                            } else {
                                var v47 float32 = 1.5
                                v41 = (v6 & 1)
                                if true {
                                    v48 := GenS8{F0: v6, F1: v6, F2: (v1 & v1)}
                                    v23 += v3
                                    v5[(("" + "hello") + ("hello" + "hello"))] = ""
                                    if false {
                                        _ = len(v5)
                                        fmt.Println(v3)
                                    } else {
                                        switch ("" + (("foo bar" + "hello") + "hello")) {
                                        case "\n\t":
                                            v49 := GenS9{F0: ("foo bar" + ((("hello" + "a") + "\n\t") + ""))}
                                            _ = v49
                                        case "hello":
                                            v47 = 0.1
                                            v50 := GenS9{F0: (("foo bar" + "\n\t") + ("hello" + ("" + ("" + "\n\t"))))}
                                            _ = v50
                                        case "":
                                            switch (("hello" + ("hello" + ((("\n\t" + "\n\t") + "a") + ""))) + "\n\t") {
                                            case "":
                                                v51 := GenS10{F0: -628}
                                                _ = v51
                                            case "\n\t":
                                                v5[""] = ("" + "")
                                            }
                                        }
                                        switch ("\n\t" + "a") {
                                        case "":
                                            _ = len(v2)
                                        default:
                                            v52 := GenS11{F0: 1000, F1: (v1 & (v1 & (v1 + int8(fi42))))}
                                            _ = v52
                                        }
                                        v41 -= (v6 / ((4294967295) | 1))
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 128
                                    }
                                    switch fi43 {
                                    case 4294967295:
                                        var v53 int8 = (v1 & v1)
                                        _ = v53
                                    case 1:
                                        var v54 int8 = (v1 + 15)
                                        _ = v54
                                    case -402:
                                        var v55 float32 = 1.5
                                        _ = v55
                                    default:
                                        fmt.Println(v23)
                                    }
                                    _ = v48
                                }
                                var v56 uint32 = 1
                                _ = v47
                                _ = v56
                            }
                        }
                    }
                    v41 += (v6 * v41)
                    var v57 int = (fi42 & 163)
                    if true {
                        var v58 uint8 = 92
                        var v59 int8 = -38
                        if true {
                            if len(v4) > 0 {
                                v4[0] = (fi42 ^ (^(fi42)))
                            }
                            switch fi42 {
                            case 9223372036854775807:
                                switch "\n\t" {
                                case "\n\t":
                                    v1 += v1
                                case "a":
                                    v6 = 221
                                default:
                                    v60 := GenS3{F0: (true && (true || true)), F1: 2}
                                    _ = v60
                                }
                            case 800:
                                v61 := GenS9{F0: ("hello" + ("hello" + "a"))}
                                _ = v61
                            case 42:
                                if true {
                                    if ((((v22 * 1e10) * v22) > ((0.0 - v22) / (v22 * -2.25))) && true) {
                                        var v62 string = "\n\t"
                                        v58 += (v58 - (v58 % (((v58 & 2)) | 1)))
                                        _ = v62
                                    } else {
                                        var v63 uint64 = v41
                                        if false {
                                            var v64 uint32 = 2147483648
                                            _ = v64
                                        } else {
                                            var v65 uint64 = v41
                                            v66 := GenS2{F0: (v3 + (v23 / ((v3) | 1))), F1: 727, F2: 1}
                                            if false {
                                                v4 = append(v4, fi43)
                                                var v67 float64 = ((v22 / 3.14159) / v22)
                                                _ = v67
                                            } else {
                                                if true {
                                                    v4 = append(v4, 32768)
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = -426
                                                    }
                                                    v68 := GenS9{F0: (("\n\t" + "foo bar") + ("\n\t" + "hello"))}
                                                    _ = v68
                                                }
                                                var v69 uint64 = 2
                                                var v70 uint32 = 437
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                        switch 32768 {
                                        case 2147483647:
                                            fmt.Println(v57, v1, v3)
                                        case 42:
                                            var v71 int32 = -374
                                            fmt.Println(v57, v23)
                                            _ = v71
                                        case 0:
                                            v41 += v63
                                        default:
                                            switch v57 {
                                            case -9223372036854775808:
                                                v59 = 42
                                            case 127:
                                                v72 := GenS8{F0: v63, F1: (v6 << (((v3 / ((65535) | 1))) & 63)), F2: v59}
                                                _ = v72
                                            }
                                        }
                                        _ = v63
                                    }
                                    var v73 int = (fi43 * -277)
                                    v1 = v1
                                    _ = v73
                                }
                            }
                        } else {
                            switch "foo bar" {
                            case "a":
                                if (v22 < 3.14159) {
                                    switch v57 {
                                    case 1:
                                        if true {
                                            fmt.Println(v23)
                                        }
                                    case 810:
                                        v74 := GenS12{F0: true, F1: 180}
                                        _ = v74
                                    }
                                    v6 = (v41 << ((v23) & 63))
                                    fmt.Println(fi42, v23, v6)
                                }
                            case "\n\t":
                                fmt.Println(v3, v58, v6)
                            case "hello":
                                if len(v4) > 0 {
                                    v4[0] = fi43
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi43
                                }
                                var v75 uint32 = 796
                                _ = v75
                            }
                            fmt.Println(v1)
                            switch (-(v57)) {
                            case -2147483648:
                                switch -990 {
                                case -168:
                                    switch (^(fi42)) {
                                    case 127:
                                        var v76 uint = v23
                                        _ = v76
                                    default:
                                        fmt.Println(v6, v3, v41)
                                    }
                                case 1:
                                    fmt.Println(v58, v3, v1)
                                case 942:
                                    var v77 int16 = 1
                                    _ = v77
                                default:
                                    v78 := GenS6{F0: uint16(v58), F1: -2147483648, F2: v6}
                                    _ = v78
                                }
                            }
                        }
                        _ = v58
                        _ = v59
                    }
                    _ = fi43
                    _ = v57
                }
                if false {
                    var v79 int32 = 526
                    _ = v79
                }
                _ = fi42
            }
        default:
            if false {
                fmt.Println(v23)
                for fi80 := 0; fi80 < 11; fi80++ {
                    for fi81 := 0; fi81 < 10; fi81++ {
                        fmt.Println(v6)
                        fmt.Println(v3, v23, v23)
                        v82 := GenS13{F0: 633, F1: 1000, F2: ("foo bar" + (("foo bar" + "hello") + "\n\t"))}
                        switch ("a" + "hello") {
                        case "hello":
                            switch (fi81 & fi81) {
                            case 748:
                                if false {
                                    var v83 uint64 = (v41 - v41)
                                    _ = v83
                                }
                            default:
                                v2[889] = "hello"
                            }
                        case "foo bar":
                            delete(v5, "a")
                        default:
                            v3 = v3
                        }
                        v41 = (v41 << ((255) & 63))
                        _ = fi81
                        _ = v82
                    }
                    _ = fi80
                }
            } else {
                v3 = v23
                var v84 float32 = float32(v22)
                _ = v84
            }
        }
        _ = v41
    }
    v6 -= (v6 | 78)
    v23 -= (v3 / ((42) | 1))
    var v85 uint = (v23 ^ v23)
    if (true || (false || true)) {
        for fi86 := 0; fi86 < 18; fi86++ {
            v85 -= (v23 - v3)
            fmt.Println(v85, v1, v6)
            v87 := GenS10{F0: 32767}
            _ = fi86
            _ = v87
        }
    } else {
        if (-1 != 9223372036854775807) {
            v6 += 4294967295
        }
        v6 += (v6 % ((166) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
}
