package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint32; F2 int }
type GenS9 struct { F0 uint32; F1 uint16; F2 int }
func genFn0() uint {
    if (false || false) {
        var v4 float32 = (0.1 - 0.0)
        v4 += v4
        var v5 bool = (false || (2 <= 42))
        _ = v4
        _ = v5
    }
    var v6 uint64 = 83
    var v7 int = 32768
    v8 := []int{}
    _ = len(v8)
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1() uint16 {
    var v9 int32 = 0
    var v10 int8 = -128
    v10 += (v10 ^ -128)
    _ = v10
    _ = v9
    return 879
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -1
    var v3 uint = 18446744073709551615
    var v11 float64 = 1.5
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "hello" {
        case "":
            if false {
                v3 += (v3 + (v3 * (v3 & (^(v3)))))
                if ((false && (0 <= v2)) && (true && false)) {
                    v3 = v3
                    if true {
                        if false {
                            v13 := GenS1{F0: (0 == -249), F1: fi12, F2: (v2 << ((42) & 63))}
                            _ = v13
                        } else {
                            fmt.Println(v1)
                            v14 := []int{1, (fi12 ^ (fi12 >> ((v3) & 63)))}
                            for fi15 := 0; fi15 < 15; fi15++ {
                                switch fi15 {
                                case 926:
                                    v1 = (v1 + (v1 / ((42) | 1)))
                                case 2147483647:
                                    _ = len(v14)
                                default:
                                    var v16 float64 = (-2.25 * v11)
                                    _ = v16
                                }
                                _ = len(v14)
                                v14 = append(v14, (fi12 & 692))
                                _ = fi15
                            }
                            _ = v14
                        }
                        for fi17 := 0; fi17 < 2; fi17++ {
                            var v18 int32 = -2147483648
                            v19 := GenS1{F0: (2 >= 32767), F1: fi12, F2: v2}
                            v3 = 65535
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    } else {
                        if ((v1 % ((42) | 1)) >= int64(fi12)) {
                            v1 -= 754
                            if ((fi12 - 127) <= 9223372036854775807) {
                                fmt.Println(v2)
                                v1 -= (v1 ^ 42)
                            } else {
                                for fi20 := 0; fi20 < 17; fi20++ {
                                    v11 = 0.1
                                    var v21 int8 = -128
                                    _ = fi20
                                    _ = v21
                                }
                                v22 := []int{}
                                switch (fi12 % ((fi12) | 1)) {
                                case 127:
                                    fmt.Println(v2, v2)
                                }
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 7; fi23++ {
                                v24 := []int{(fi12 >> ((273) & 63)), fi23}
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            v25 := GenS2{F0: (!(2147483648 > 863)), F1: fi12, F2: v1}
                            v26 := GenS2{F0: (true || (0.1 != 1.5)), F1: (fi12 << (((v3 + v3)) & 63)), F2: v1}
                            _ = v25
                            _ = v26
                        }
                        if (2 != (v3 * (v3 << ((199) & 63)))) {
                            v27 := []int{}
                            for fi28 := 0; fi28 < 19; fi28++ {
                                fmt.Println(v11)
                                if (3.14159 > 1e10) {
                                    var v29 int32 = 293
                                    switch "hello" {
                                    case "a":
                                        v1 = (v1 * 9223372036854775807)
                                    case "foo bar":
                                        var v30 uint8 = 0
                                        _ = v30
                                    case "\n\t":
                                        var v31 uint16 = genFn1()
                                        _ = v31
                                    }
                                    _ = v29
                                }
                                _ = fi28
                            }
                            _ = v27
                        }
                    }
                }
            } else {
                v2 += v2
                if false {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v3 += genFn0()
                        v33 := []int{(fi12 | (-(fi32))), (fi12 & fi12), (fi32 & 0)}
                        if true {
                            if (v1 > v1) {
                                var v34 uint16 = 1
                                switch "\n\t" {
                                case "a":
                                    var v35 bool = true
                                    _ = v35
                                case "hello":
                                    v36 := GenS1{F0: false, F1: (fi12 & fi12), F2: (v2 / ((v2) | 1))}
                                    _ = v36
                                case "\n\t":
                                    v37 := GenS3{F0: (v34 / ((955) | 1))}
                                    _ = v37
                                default:
                                    var v38 int64 = (^(v1))
                                    _ = v38
                                }
                                switch "a" {
                                case "":
                                    fmt.Println(v1)
                                    v34 = v34
                                case "a":
                                    v34 -= v34
                                case "\n\t":
                                    switch int(v2) {
                                    case -9223372036854775808:
                                        v39 := GenS4{F0: genFn1(), F1: 430}
                                        _ = v39
                                    }
                                }
                                var v40 bool = (1 != -288)
                                if (1e10 > ((0.1 * -2.25) + float32(v11))) {
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(fi12, fi32)
                                    case "hello":
                                        v34 = (v34 >> (((v3 | 65535)) & 63))
                                    case "a":
                                        var v41 int8 = -117
                                        _ = v41
                                    }
                                    var v42 int = fi32
                                    _ = v42
                                } else {
                                    switch (("a" + "\n\t") + "\n\t") {
                                    case "a":
                                        if len(v33) >= 2 {
                                            v33 = v33[1:len(v33)]
                                        }
                                    default:
                                        var v43 uint = v3
                                        _ = v43
                                    }
                                    var v44 float32 = 0.0
                                    _ = v44
                                }
                                fmt.Println(fi32)
                                _ = v34
                                _ = v40
                            }
                        } else {
                            v45 := GenS2{F0: true, F1: fi12, F2: v1}
                            v11 = -2.25
                            if true {
                                v11 -= v11
                                fmt.Println(v1, fi12)
                                v2 += (v2 / (((v2 + (v2 + v2))) | 1))
                            } else {
                                v1 = (^(v1))
                                var v46 string = (("" + "foo bar") + ("foo bar" + "foo bar"))
                                var v47 int16 = v2
                                _ = v46
                                _ = v47
                            }
                            _ = v45
                        }
                        _ = fi32
                        _ = v33
                    }
                    if (v3 < (^(v3))) {
                        var v48 uint16 = genFn1()
                        if false {
                            for fi49 := 0; fi49 < 1; fi49++ {
                                fmt.Println(v1, v3)
                                var v50 int = 42
                                var v51 uint8 = 42
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 13; fi52++ {
                                fmt.Println(fi12, fi52)
                                v53 := GenS1{F0: true, F1: 819, F2: (v2 / (((-(v2))) | 1))}
                                _ = fi52
                                _ = v53
                            }
                            v11 = v11
                        } else {
                            var v54 int8 = -128
                            _ = v54
                        }
                        switch "hello" {
                        case "\n\t":
                            v2 -= v2
                        default:
                            v55 := GenS5{F0: "\n\t", F1: 128, F2: -181}
                            _ = v55
                        }
                        _ = v48
                    }
                    switch ("" + (("hello" + ("hello" + "\n\t")) + "foo bar")) {
                    case "\n\t":
                        if (((-128 < 100) || (42 != (v1 ^ v1))) || (0 >= 2)) {
                            if true {
                                v2 += v2
                            }
                            v3 += (^(v3))
                            var v56 uint32 = 42
                            _ = v56
                        }
                    default:
                        v57 := []int{(fi12 | (fi12 * fi12)), fi12, fi12, fi12}
                        _ = v57
                    }
                } else {
                    var v58 string = ""
                    v59 := GenS6{F0: v3}
                    _ = v58
                    _ = v59
                }
            }
        case "\n\t":
            v60 := []int{-102, 0, (-(fi12))}
            _ = v60
        default:
            v61 := []int{fi12}
            _ = v61
        }
        _ = fi12
    }
    v62 := GenS4{F0: 967, F1: 2}
    if false {
        fmt.Println(v3, v2)
    }
    if ((v2 << (((^(v3))) & 63)) != (v2 / ((v2) | 1))) {
        var v63 uint8 = 255
        _ = v63
    } else {
        switch "foo bar" {
        case "":
            var v64 bool = true
            _ = v64
        case "hello":
            v1 = v1
        }
        v11 = v11
    }
    var v65 uint32 = 42
    fmt.Println(v11)
    switch "foo bar" {
    case "hello":
        fmt.Println(v11, v1, v2)
    default:
        v11 = float64(0.0)
    }
    switch "foo bar" {
    case "a":
        v11 += v11
    case "":
        v66 := []int{32768}
        _ = v66
    case "foo bar":
        var v67 string = "\n\t"
        _ = v67
    }
    for fi68 := 0; fi68 < 4; fi68++ {
        var v69 int32 = 1
        _ = fi68
        _ = v69
    }
    var v70 float32 = 1.5
    var v71 int64 = int64(v2)
    if true {
        if (((1 >= v65) && true) && (!true)) {
            for fi72 := 0; fi72 < 8; fi72++ {
                if (false && (1 != 100)) {
                    v3 += genFn0()
                    v73 := GenS7{F0: (v3 & 1)}
                    var v74 int16 = (v2 ^ v2)
                    _ = v73
                    _ = v74
                }
                v3 = v3
                var v75 uint = (v3 * 18446744073709551615)
                _ = fi72
                _ = v75
            }
        } else {
            var v76 int8 = -104
            v3 += genFn0()
            v77 := GenS3{F0: 2}
            _ = v76
            _ = v77
        }
    } else {
        if (false || true) {
            var v78 uint = v3
            if (("\n\t" > "") && (true && false)) {
                v79 := []int{798, 32768, 1, -766}
                v2 -= 1000
                v1 -= v1
                _ = v79
            } else {
                var v80 float32 = 0.1
                switch 464 {
                case 9223372036854775807:
                    var v81 uint32 = v65
                    _ = v81
                case 127:
                    v82 := GenS8{F0: v71, F1: v65, F2: -284}
                    _ = v82
                default:
                    switch "" {
                    case "":
                        v83 := []int{-9223372036854775808, 1, -511, 128}
                        _ = v83
                    case "\n\t":
                        if ("foo bar" >= "") {
                            v65 -= v65
                        } else {
                            for fi84 := 0; fi84 < 12; fi84++ {
                                for fi85 := 0; fi85 < 10; fi85++ {
                                    var v86 int64 = (v71 | (v71 * (v71 >> ((1) & 63))))
                                    fmt.Println(v86, v70, v71)
                                    v70 = ((float32(v11) / v80) * float32(1.5))
                                    _ = fi85
                                    _ = v86
                                }
                                _ = fi84
                            }
                        }
                    case "a":
                        v87 := GenS6{F0: 4294967295}
                        _ = v87
                    }
                }
                var v88 uint64 = 859
                _ = v80
                _ = v88
            }
            _ = v78
        }
        if (((v70 / 1.0) / (v70 - float32(((0.1 + v11) - 0.0)))) >= ((v70 / v70) - v70)) {
            if false {
                var v89 int16 = (v2 ^ v2)
                v90 := []int{291, -1, -100, 128}
                _ = v89
                _ = v90
            }
            for fi91 := 0; fi91 < 5; fi91++ {
                for fi92 := 0; fi92 < 12; fi92++ {
                    v71 = (v1 & 42)
                    var v93 int8 = -128
                    _ = fi92
                    _ = v93
                }
                var v94 string = "\n\t"
                v95 := GenS9{F0: v65, F1: 503, F2: -72}
                _ = fi91
                _ = v94
                _ = v95
            }
            for fi96 := 0; fi96 < 18; fi96++ {
                v70 += float32(v11)
                var v97 bool = true
                if false {
                    v11 = (v11 * v11)
                    for fi98 := 0; fi98 < 18; fi98++ {
                        if (848 >= 32767) {
                            fmt.Println(fi98, v70, v1)
                        } else {
                            v11 += (3.14159 + v11)
                            var v99 uint8 = 255
                            _ = v99
                        }
                        _ = fi98
                    }
                }
                _ = fi96
                _ = v97
            }
        }
    }
    v100 := []int{-585, 9223372036854775807, -1}
    fmt.Println(v71, v70, v11)
    if (-146 >= -1) {
        for fi101 := 0; fi101 < 9; fi101++ {
            if (!(false && (true || (113 != -1)))) {
                var v102 int64 = (v1 >> ((42) & 63))
                fmt.Println(v1, v2)
                var v103 int = fi101
                _ = v102
                _ = v103
            }
            fmt.Println(v2, v1, v70)
            _ = fi101
        }
        v100 = append(v100, 268)
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v100
    _ = v11
    _ = v2
    _ = v3
    _ = v62
    _ = v65
    _ = v70
    _ = v71
}
