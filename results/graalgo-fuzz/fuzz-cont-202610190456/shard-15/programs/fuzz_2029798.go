package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int16; F2 uint64 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int32; F1 int64; F2 int32 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 uint; F1 uint32; F2 int32 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 bool) string {
    v4 := GenS1{F0: 63, F1: 883, F2: 4294967295}
    var v5 string = ("hello" + (((("" + "hello") + "foo bar") + "hello") + (("foo bar" + ("a" + "")) + "foo bar")))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v5 + "\n\t")
}
func genFn1() int64 {
    v6 := GenS2{F0: 599, F1: -1}
    var v7 int8 = 1
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(fi8)
        v7 -= (v7 << ((4294967295) & 63))
        _ = fi8
    }
    switch "hello" {
    case "\n\t":
        if ("a" != (("a" + "hello") + "hello")) {
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "foo bar":
                for fi9 := 0; fi9 < 20; fi9++ {
                    fmt.Println(v7, v7)
                    v7 = 0
                    switch ((((("\n\t" + "a") + "\n\t") + "foo bar") + "foo bar") + ("" + "foo bar")) {
                    case "foo bar":
                        v10 := GenS3{F0: 32767, F1: 969, F2: 416}
                        _ = v10
                    case "":
                        v11 := []int{}
                        _ = v11
                    }
                    _ = fi9
                }
            case "":
                v12 := []int{}
                _ = v12
            default:
                if (32767 == 32767) {
                    var v13 bool = (true && false)
                    var v14 float64 = 0.0
                    v15 := GenS2{F0: 835, F1: 498}
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            v16 := GenS4{F0: -1}
            _ = v16
        } else {
            v17 := []int{-2147483648}
            _ = v17
        }
    case "a":
        fmt.Println(v7)
    case "hello":
        v18 := []int{}
        _ = v18
    default:
        v7 += v7
    }
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 uint8 = 99
    if ("a" > ((("hello" + ("" + "\n\t")) + ("" + "")) + "a")) {
        var v19 uint = 2
        if v2 {
            fmt.Println(v3, v19)
        } else {
            if v2 {
                switch int(v3) {
                case 42:
                    var v20 uint64 = 4294967295
                    _ = v20
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    v19 -= (v19 * v19)
                    if (-998 < 0) {
                        v3 = v3
                    } else {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := GenS5{F0: -9223372036854775808, F1: 1000, F2: 567}
                            v1[uint32(v19)] = 1
                            var v24 uint16 = 65535
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        var v25 uint16 = 32767
                        _ = v25
                    }
                    _ = fi21
                }
            } else {
                if ((3.14159 >= 1.5) && v2) {
                    v26 := GenS1{F0: 0, F1: 42, F2: 44}
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v2 = v2
                        _ = fi27
                    }
                    _ = v26
                }
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                fmt.Println(v3, fi28)
                var v29 uint64 = 215
                v19 -= (v19 ^ v19)
                _ = fi28
                _ = v29
            }
        }
        _ = v19
    }
    switch "a" {
    case "foo bar":
        for fi30 := 0; fi30 < 9; fi30++ {
            fmt.Println(v2)
            var v31 uint32 = 42
            _ = fi30
            _ = v31
        }
    case "hello":
        v32 := GenS4{F0: 32767}
        _ = v32
    case "\n\t":
        v2 = (((28 >= 4294967295) && v2) && true)
    default:
        for fi33 := 0; fi33 < 10; fi33++ {
            if false {
                if false {
                    v34 := GenS4{F0: 32767}
                    v35 := []int{}
                    for fi36 := 0; fi36 < 14; fi36++ {
                        v37 := GenS6{F0: -1, F1: 0}
                        v38 := GenS1{F0: 32767, F1: 2147483647, F2: 65535}
                        var v39 int64 = 9223372036854775807
                        v2 = (((1.5 - (1.5 - -2.25)) > (1.5 + 3.14159)) && ((^(v3)) >= (v3 | (v3 | 42))))
                        v3 = 0
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = v34
                    _ = v35
                }
            }
            _ = fi33
        }
    }
    fmt.Println(v3, v3)
    if (false && (1000 < 32767)) {
        var v40 string = "\n\t"
        switch "" {
        case "a":
            v41 := GenS7{F0: -1, F1: 836, F2: 0}
            _ = v41
        case "foo bar":
            var v42 int64 = -9223372036854775808
            _ = v42
        case "":
            v40 = v40
        default:
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 int32 = 0
                if (v40 >= "foo bar") {
                    v45 := []int{(fi43 >> ((0) & 63)), -9223372036854775808, 1}
                    v40 = (((v40 + v40) + (("hello" + v40) + "a")) + (v40 + "foo bar"))
                    _ = v45
                } else {
                    var v46 uint32 = 171
                    _ = v46
                }
                v47 := GenS8{F0: -9223372036854775808, F1: 0, F2: v3}
                _ = fi43
                _ = v44
                _ = v47
            }
        }
        _ = v40
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := GenS9{F0: 404, F1: 2147483648, F2: 378}
        if true {
            if false {
                if false {
                    var v50 float32 = 1.5
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        if v2 {
                            if v2 {
                                var v52 float64 = -2.25
                                var v53 uint16 = 670
                                _ = v52
                                _ = v53
                            } else {
                                v54 := GenS8{F0: genFn1(), F1: 908, F2: v3}
                                _ = v54
                            }
                            v3 -= (^(v3))
                            v55 := []int{(fi48 << ((932) & 63)), fi48}
                            if (int16(fi51) < -625) {
                                var v56 int32 = 2147483647
                                var v57 uint = 65535
                                _ = v56
                                _ = v57
                            }
                            _ = v55
                        }
                        _ = fi51
                    }
                }
                _ = len(v1)
                v58 := []int{(^(fi48)), fi48}
                _ = v58
            } else {
                fmt.Println(v3)
                switch "" {
                case "\n\t":
                    var v59 int64 = -9223372036854775808
                    _ = v59
                case "a":
                    v60 := GenS2{F0: 27, F1: 0}
                    _ = v60
                }
                if v2 {
                    var v61 int = (fi48 << ((2) & 63))
                    _ = v61
                } else {
                    v62 := GenS4{F0: 32767}
                    if (809 >= 4294967295) {
                        v63 := GenS4{F0: -56}
                        for fi64 := 0; fi64 < 19; fi64++ {
                            v2 = (v2 || v2)
                            fmt.Println(fi48)
                            var v65 bool = (316 >= 370)
                            _ = fi64
                            _ = v65
                        }
                        fmt.Println(v2)
                        _ = v63
                    }
                    _ = v62
                }
            }
        } else {
            if v2 {
                v66 := []int{(fi48 ^ fi48), fi48, 4294967295}
                v3 -= 2
                _ = v66
            } else {
                var v67 uint64 = 18446744073709551615
                var v68 uint8 = 159
                _ = v67
                _ = v68
            }
            v69 := []int{(fi48 >> ((18446744073709551615) & 63)), (^(fi48))}
            _ = v69
        }
        _ = fi48
        _ = v49
    }
    var v70 int32 = 407
    if (108 == 1) {
        for fi71 := 0; fi71 < 4; fi71++ {
            v3 += v3
            _ = fi71
        }
        if (1 < 964) {
            for fi72 := 0; fi72 < 17; fi72++ {
                v1[65535] = fi72
                v2 = false
                _ = fi72
            }
        } else {
            var v73 uint16 = 255
            _ = v73
        }
        v2 = true
    } else {
        if v2 {
            if true {
                v74 := GenS9{F0: 214, F1: 1, F2: 1}
                v75 := GenS10{F0: 1000, F1: 32768}
                _ = v74
                _ = v75
            } else {
                v2 = v2
                v76 := GenS11{F0: 54, F1: 2147483647, F2: (v70 / (((^(v70))) | 1))}
                switch 741 {
                case -9223372036854775808:
                    v3 -= 128
                case 32768:
                    for fi77 := 0; fi77 < 4; fi77++ {
                        switch ("hello" + "a") {
                        case "":
                            for fi78 := 0; fi78 < 15; fi78++ {
                                v3 -= 255
                                if (!false) {
                                    var v79 uint16 = 32767
                                    _ = v79
                                }
                                switch "" {
                                case "\n\t":
                                    switch 295 {
                                    case 9223372036854775807:
                                        switch "" {
                                        case "a":
                                            v70 -= v70
                                        default:
                                            fmt.Println(fi78)
                                        }
                                    default:
                                        v80 := GenS6{F0: 42, F1: v70}
                                        _ = v80
                                    }
                                case "":
                                    var v81 float64 = 0.1
                                    _ = v81
                                case "a":
                                    switch -187 {
                                    case 42:
                                        v3 += (v3 | (v3 + (v3 ^ v3)))
                                        var v82 uint8 = (v3 >> ((1) & 63))
                                        switch -904 {
                                        case -739:
                                            var v83 bool = (-2.25 == 1.5)
                                            var v84 uint16 = 285
                                            var v85 uint = 65535
                                            fmt.Println(v83, v3)
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                        case -1:
                                            var v86 bool = ((false || false) || (2147483648 == 4294967295))
                                            _ = v86
                                        }
                                        _ = v82
                                    case 447:
                                        v3 -= (v3 - 0)
                                    case -9223372036854775808:
                                        var v87 string = ("foo bar" + "a")
                                        _ = v87
                                    default:
                                        var v88 float32 = float32(1e10)
                                        _ = v88
                                    }
                                default:
                                    v89 := GenS12{F0: 2}
                                    _ = v89
                                }
                                v1[0] = fi77
                                _ = fi78
                            }
                        case "a":
                            switch "a" {
                            case "a":
                                v90 := GenS13{F0: (v70 - 1), F1: 9223372036854775807, F2: 0}
                                _ = v90
                            case "\n\t":
                                var v91 int32 = v70
                                _ = v91
                            }
                        }
                        _ = fi77
                    }
                default:
                    for fi92 := 0; fi92 < 5; fi92++ {
                        for fi93 := 0; fi93 < 8; fi93++ {
                            var v94 uint8 = (v3 - (v3 & (^(v3))))
                            var v95 uint8 = 198
                            _ = fi93
                            _ = v94
                            _ = v95
                        }
                        v70 = v70
                        _ = fi92
                    }
                }
                _ = v76
            }
            fmt.Println(v2, v70, v3)
        } else {
            fmt.Println(v70, v2, v2)
        }
        var v96 float64 = 0.0
        _ = v96
    }
    v97 := []int{1}
    for fi98 := 0; fi98 < 5; fi98++ {
        v2 = (((3.14159 * (float32(0.1) / 1e10)) >= (1.5 + float32(0.0))) && v2)
        _ = fi98
    }
    var v99 uint32 = 2
    var v100 float64 = 0.1
    v100 += v100
    if true {
        var v101 bool = (v70 <= 0)
        if (!v2) {
            fmt.Println(v3)
            v99 -= (v99 / ((v99) | 1))
        }
        _ = v101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v97:", len(v97), v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v70
    _ = v97
    _ = v99
}
