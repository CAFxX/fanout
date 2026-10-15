package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 string; F1 uint32; F2 uint64 }
type GenS12 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 int) int64 {
    var v5 uint64 = 0
    var v6 string = (((("" + "hello") + ("" + "foo bar")) + ("\n\t" + "foo bar")) + "foo bar")
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{p1, (p1 ^ fi8), (fi8 / ((32768) | 1))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (p0 % ((42) | 1))
}
func genFn1() int32 {
    var v10 int = 9223372036854775807
    switch ("\n\t" + "") {
    case "a":
        v11 := []int{v10}
        _ = v11
    case "hello":
        v12 := []int{}
        _ = v12
    default:
        v13 := []int{}
        _ = v13
    }
    if true {
        if (!true) {
            switch "" {
            case "":
                v14 := []int{159, (-(v10)), (v10 * (v10 / ((v10) | 1))), 32768}
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 uint16 = 255
                    v17 := GenS1{F0: 42, F1: "hello", F2: 42}
                    if false {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            var v19 int32 = 0
                            v16 -= 691
                            _ = fi18
                            _ = v19
                        }
                        if ((v16 ^ (^(v16))) <= 119) {
                            fmt.Println(v16, fi15)
                        }
                        if true {
                            if ("a" > "\n\t") {
                                if (true && ((-2.25 <= 0.1) && false)) {
                                    v10 -= fi15
                                    var v20 uint8 = 127
                                    v16 -= (v16 | 1)
                                    _ = v20
                                }
                                if (-347 < 2147483647) {
                                    if true {
                                        var v21 uint32 = 997
                                        v21 = (v21 << ((4294967295) & 63))
                                        v22 := []int{(fi15 % (((v10 >> ((42) & 63))) | 1)), 127, (^(fi15))}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v23 := []int{0, v10, 0, fi15}
                                        v23 = append(v23, fi15)
                                        for fi24 := 0; fi24 < 18; fi24++ {
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            v10 -= (^(fi24))
                                            fmt.Println(fi24)
                                            _ = fi24
                                        }
                                        _ = v23
                                    }
                                    v10 += 32768
                                    if (-1 != 115) {
                                        var v25 uint16 = 822
                                        v26 := []int{int(v25), (^(fi15)), (v10 + (fi15 >> ((4294967295) & 63)))}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v27 := []int{fi15, (^(v10)), 9223372036854775807}
                                        if (!false) {
                                            var v28 int32 = 149
                                            _ = v28
                                        }
                                        _ = v27
                                    }
                                }
                            } else {
                                fmt.Println(fi15, v16, fi15)
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    var v30 float32 = -2.25
                                    v31 := GenS1{F0: 2, F1: "\n\t", F2: 0}
                                    v30 += float32(1.5)
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                        } else {
                            var v32 float64 = (((((0.0 / 1.0) - 1.5) / ((float64(3.14159) + 1.5) / 1.0)) - 1.5) / float64(-2.25))
                            v33 := GenS2{F0: -1}
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                var v34 string = "foo bar"
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 12; fi35++ {
                switch 0 {
                case -2147483648:
                    var v36 uint32 = 4294967295
                    _ = v36
                case 127:
                    fmt.Println(fi35, v10)
                default:
                    v37 := []int{fi35}
                    _ = v37
                }
                switch v10 {
                case 32767:
                    switch "a" {
                    case "":
                        var v38 string = "a"
                        _ = v38
                    case "foo bar":
                        switch (fi35 / (((-(v10))) | 1)) {
                        case 9223372036854775807:
                            v10 = 127
                        default:
                            if true {
                                switch (fi35 / ((v10) | 1)) {
                                case 1:
                                    v10 = 700
                                case 0:
                                    v10 += fi35
                                case 32768:
                                    for fi39 := 0; fi39 < 15; fi39++ {
                                        v40 := GenS1{F0: 0, F1: "a", F2: -160}
                                        _ = fi39
                                        _ = v40
                                    }
                                default:
                                    for fi41 := 0; fi41 < 10; fi41++ {
                                        if true {
                                            fmt.Println(fi35)
                                            v10 += fi41
                                            fmt.Println(v10)
                                        } else {
                                            switch (^(fi41)) {
                                            case 127:
                                                var v42 bool = true
                                                _ = v42
                                            case 4294967295:
                                                switch "a" {
                                                case "a":
                                                    var v43 int = (v10 >> ((711) & 63))
                                                    _ = v43
                                                case "\n\t":
                                                    if false {
                                                        var v44 int64 = 239
                                                        _ = v44
                                                    }
                                                case "hello":
                                                    v10 -= v10
                                                }
                                                var v45 int16 = -975
                                                v45 = (v45 ^ v45)
                                                _ = v45
                                            default:
                                                switch (fi41 % ((4294967295) | 1)) {
                                                case 32768:
                                                    v46 := []int{fi35, (fi41 << ((0) & 63)), fi35}
                                                    _ = v46
                                                case -887:
                                                    var v47 int32 = 1
                                                    _ = v47
                                                default:
                                                    var v48 uint64 = 18446744073709551615
                                                    _ = v48
                                                }
                                            }
                                            v49 := GenS3{F0: 1, F1: 42, F2: -346}
                                            if false {
                                                v50 := []int{2147483647, -825, (v10 + (v10 % (((fi35 - fi35)) | 1)))}
                                                v51 := GenS3{F0: 0, F1: -341, F2: 0}
                                                fmt.Println(fi35)
                                                _ = v50
                                                _ = v51
                                            }
                                            v52 := []int{fi41, (fi35 * (fi41 * (v10 * (fi35 - (v10 | fi35))))), 991, (fi41 << ((2) & 63))}
                                            v10 = (fi41 / ((v10) | 1))
                                            var v53 uint = 995
                                            v53 = (v53 | (v53 - 1))
                                            v54 := GenS2{F0: 970}
                                            _ = v49
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = fi41
                                    }
                                }
                            }
                        }
                    case "a":
                        for fi55 := 0; fi55 < 8; fi55++ {
                            v56 := []int{(^(fi35)), (-(fi35)), 4294967295}
                            v10 -= 42
                            _ = fi55
                            _ = v56
                        }
                    }
                case -1:
                    var v57 float64 = -2.25
                    _ = v57
                case -4:
                    v10 = (-(fi35))
                default:
                    v58 := []int{v10}
                    _ = v58
                }
                fmt.Println(fi35, fi35)
                _ = fi35
            }
        }
        v10 = 1
        v59 := GenS4{F0: 1, F1: 0, F2: 1}
        _ = v59
    } else {
        switch 2147483647 {
        case 805:
            var v60 int8 = -1
            _ = v60
        }
        var v61 uint16 = 949
        _ = v61
    }
    v62 := []int{v10, -9223372036854775808, (-(v10)), (v10 * v10)}
    var v63 string = ""
    _ = v10
    _ = v62
    _ = v63
    return -249
}

func main() {
    var v1 uint16 = 198
    var v2 int64 = -1
    v3 := []int{9223372036854775807, -9223372036854775808, 2147483647}
    var v4 int = -2147483648
    var v64 uint32 = 1
    for fi65 := 0; fi65 < 18; fi65++ {
        var v66 uint16 = v1
        var v67 uint = 18446744073709551615
        v64 = 363
        _ = fi65
        _ = v66
        _ = v67
    }
    v3 = append(v3, (^(v4)))
    for fi68 := 0; fi68 < 17; fi68++ {
        fmt.Println(fi68, v64, v2)
        _ = fi68
    }
    v2 += -9223372036854775808
    v4 = 974
    v69 := GenS3{F0: v64, F1: 1000, F2: 443}
    v1 = (v1 - (v1 ^ v1))
    if true {
        v1 = v1
        for fi70 := 0; fi70 < 9; fi70++ {
            if ((0 != 42) || true) {
                fmt.Println(v1, v4)
            }
            _ = fi70
        }
        v71 := GenS5{F0: v1, F1: uint32(v1), F2: (v4 * v4)}
        _ = v71
    } else {
        var v72 uint8 = 72
        v73 := GenS6{F0: -2147483648}
        switch v4 {
        case -9223372036854775808:
            v2 += 0
        case -127:
            for fi74 := 0; fi74 < 8; fi74++ {
                var v75 uint64 = 42
                _ = fi74
                _ = v75
            }
        case -1:
            v76 := GenS2{F0: v2}
            _ = v76
        }
        _ = v72
        _ = v73
    }
    v77 := GenS7{F0: 605, F1: "\n\t"}
    _ = len(v3)
    if true {
        v78 := GenS8{F0: ("" + "a")}
        _ = v78
    } else {
        _ = len(v3)
        v79 := GenS9{F0: 673, F1: v1}
        _ = v79
    }
    v3 = append(v3, -881)
    if true {
        for fi80 := 0; fi80 < 3; fi80++ {
            v1 += (v1 >> ((0) & 63))
            _ = len(v3)
            v1 -= (v1 - v1)
            _ = fi80
        }
        switch ("" + "hello") {
        case "":
            switch "\n\t" {
            case "foo bar":
                if ((0.1 / 1.5) < (-2.25 - -2.25)) {
                    for fi81 := 0; fi81 < 5; fi81++ {
                        v82 := GenS2{F0: v2}
                        _ = fi81
                        _ = v82
                    }
                }
            case "a":
                _ = len(v3)
            case "hello":
                var v83 uint32 = (v64 % ((v64) | 1))
                _ = v83
            default:
                for fi84 := 0; fi84 < 1; fi84++ {
                    var v85 int = fi84
                    _ = fi84
                    _ = v85
                }
            }
        case "hello":
            v4 = 858
        default:
            v86 := GenS7{F0: 18446744073709551615, F1: (("foo bar" + "a") + "a")}
            _ = v86
        }
        var v87 float32 = 1e10
        _ = v87
    }
    for fi88 := 0; fi88 < 10; fi88++ {
        var v89 int16 = -32768
        for fi90 := 0; fi90 < 5; fi90++ {
            var v91 uint16 = (v1 / ((v1) | 1))
            v1 -= (v1 | v91)
            var v92 uint8 = 29
            _ = fi90
            _ = v91
            _ = v92
        }
        if (-30 >= -128) {
            v93 := GenS6{F0: 1}
            for fi94 := 0; fi94 < 5; fi94++ {
                v4 = fi88
                fmt.Println(v1, v2, v89)
                switch "\n\t" {
                case "foo bar":
                    var v95 int32 = -1
                    _ = v95
                default:
                    v96 := GenS10{F0: 2, F1: (fi88 | (fi88 | (v4 << ((2) & 63))))}
                    _ = v96
                }
                _ = fi94
            }
            var v97 uint32 = (v64 + v64)
            _ = v93
            _ = v97
        } else {
            v4 += (fi88 * (fi88 - (fi88 | -960)))
        }
        _ = fi88
        _ = v89
    }
    for fi98 := 0; fi98 < 18; fi98++ {
        v99 := GenS10{F0: 0, F1: (-(fi98))}
        v100 := GenS11{F0: "\n\t", F1: (v64 + v64), F2: 913}
        switch "" {
        case "foo bar":
            v101 := GenS12{F0: 2147483647, F1: (-128 == 0), F2: 2147483647}
            _ = v101
        case "hello":
            var v102 int32 = 42
            _ = v102
        case "a":
            var v103 int64 = (v2 / ((genFn0((v2 << ((492) & 63)), -9223372036854775808)) | 1))
            _ = v103
        }
        _ = fi98
        _ = v100
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v69
    _ = v77
}
