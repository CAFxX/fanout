package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int; F2 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 string, p1 uint16, p2 uint8) int {
    v5 := []int{2147483647, 0, 1}
    if true {
        v6 := GenS1{F0: -1, F1: 42, F2: (p1 > 42)}
        _ = v6
    }
    var v7 int32 = -1
    switch -193 {
    case 0:
        var v8 uint64 = 911
        _ = v8
    case 937:
        if len(v5) > 0 {
            v5[0] = -774
        }
    case -445:
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v9 := GenS2{F0: 547, F1: 42}
            _ = v9
        } else {
            var v10 float32 = (3.14159 / float32(float64((float32(3.14159) / ((float32(-2.25) - 0.0) * 1e10)))))
            _ = v10
        }
    default:
        if (0 <= 976) {
            var v11 int16 = 378
            _ = v11
        }
    }
    var v12 string = (p0 + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    _ = v7
    return 184
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -1
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    switch "hello" {
    case "\n\t":
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{}
            for fi15 := 0; fi15 < 12; fi15++ {
                fmt.Println(fi13, fi13, fi15)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    case "a":
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 uint64 = 1
            fmt.Println(v2, fi16)
            _ = fi16
            _ = v17
        }
    }
    delete(v1, 766)
    var v18 string = "hello"
    if (!(true || false)) {
        v19 := GenS3{F0: 1000, F1: 18446744073709551615, F2: 650}
        var v20 uint = 270
        if true {
            v2 -= v2
        }
        _ = v19
        _ = v20
    }
    if false {
        v18 = v18
    } else {
        var v21 uint32 = 2147483648
        var v22 bool = true
        _ = v21
        _ = v22
    }
    if (v18 < v18) {
        var v23 int64 = 718
        for fi24 := 0; fi24 < 15; fi24++ {
            if (genFn0(v18, 32767, 87) >= fi24) {
                v25 := GenS4{F0: -706, F1: 576, F2: 2}
                _ = v25
            }
            if (42 != 1) {
                if false {
                    v26 := GenS4{F0: -2147483648, F1: 65535, F2: 2}
                    v2 = int8(v23)
                    v27 := []int{-1}
                    _ = v26
                    _ = v27
                }
                if true {
                    var v28 int16 = -127
                    delete(v4, fi24)
                    _ = v28
                }
                if false {
                    _ = len(v3)
                    v4[32768] = "a"
                } else {
                    var v29 float64 = 1.5
                    v30 := []int{}
                    _ = v29
                    _ = v30
                }
            } else {
                switch v18 {
                case "":
                    v23 = (v23 & (v23 | (^(v23))))
                case "\n\t":
                    if false {
                        var v31 uint = 255
                        if false {
                            if (42 != 538) {
                                var v32 uint32 = 210
                                _ = v32
                            } else {
                                _ = len(v4)
                                var v33 int16 = -581
                                switch (fi24 % ((genFn0(((v18 + v18) + ("a" + v18)), 42, 0)) | 1)) {
                                case 144:
                                    v34 := []int{}
                                    _ = v34
                                }
                                _ = v33
                            }
                        } else {
                            var v35 int = (fi24 - fi24)
                            _ = v35
                        }
                        v36 := GenS5{F0: "foo bar"}
                        _ = v31
                        _ = v36
                    }
                default:
                    v18 = v18
                }
                fmt.Println(v2, v23, v2)
            }
            v2 -= v2
            _ = fi24
        }
        _ = v23
    }
    if false {
        var v37 int32 = -838
        v37 -= (v37 & v37)
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        var v39 int32 = 390
        _ = fi38
        _ = v39
    }
    var v40 uint16 = 664
    if (true || ((!true) || (!(!true)))) {
        v2 = v2
        var v41 uint64 = 42
        var v42 bool = false
        _ = v41
        _ = v42
    }
    v4[343] = "\n\t"
    v43 := []int{}
    switch ("a" + ("foo bar" + "foo bar")) {
    case "":
        fmt.Println(v40)
    case "foo bar":
        if (false && false) {
            fmt.Println(v40, v40)
            v44 := GenS6{F0: true, F1: genFn0("", (v40 | v40), 1), F2: v2}
            v40 += v40
            _ = v44
        } else {
            fmt.Println(v2, v40)
            switch v18 {
            case "\n\t":
                v40 -= (v40 >> ((0) & 63))
            case "a":
                if false {
                    v2 = v2
                    var v45 uint32 = 50
                    _ = v45
                } else {
                    v46 := GenS2{F0: 0, F1: 42}
                    if ((v40 ^ v40) <= v40) {
                        fmt.Println(v2, v18)
                        for fi47 := 0; fi47 < 5; fi47++ {
                            for fi48 := 0; fi48 < 18; fi48++ {
                                var v49 int16 = 1000
                                v40 -= 65535
                                if true {
                                    v2 += v2
                                    if (!false) {
                                        fmt.Println(fi47, fi48)
                                    } else {
                                        v18 = ("foo bar" + "foo bar")
                                        if false {
                                            if false {
                                                fmt.Println(fi47, v18, v2)
                                            } else {
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                                var v50 int64 = -1
                                                _ = v50
                                            }
                                            if true {
                                                v51 := GenS6{F0: true, F1: (fi48 * (fi48 - -50)), F2: (v2 + v2)}
                                                if true {
                                                    v52 := GenS7{F0: 4294967295}
                                                    fmt.Println(fi48, fi48, fi47)
                                                    var v53 uint8 = 2
                                                    var v54 int = 4294967295
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v55 := GenS4{F0: -1, F1: (v40 << ((441) & 63)), F2: 198}
                                                _ = v51
                                                _ = v55
                                            } else {
                                                v56 := GenS1{F0: int32(fi47), F1: fi48, F2: (!true)}
                                                var v57 uint16 = 361
                                                v40 = (v40 & v40)
                                                fmt.Println(v57, v57, v18)
                                                v58 := GenS2{F0: 18446744073709551615, F1: -412}
                                                if (!(v2 >= (v2 ^ 109))) {
                                                    if (65535 <= 816) {
                                                        v59 := GenS7{F0: 2}
                                                        _ = v59
                                                    }
                                                    fmt.Println(fi47, v2)
                                                    v18 = ("" + v18)
                                                }
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        } else {
                                            v60 := GenS1{F0: -2147483648, F1: (fi48 << ((18446744073709551615) & 63)), F2: (true || true)}
                                            _ = v60
                                        }
                                    }
                                    if (3.14159 > 1e10) {
                                        var v61 float32 = (0.1 * (-2.25 - ((0.1 - (((0.1 + 1.5) / 1.0) - 0.1)) / (((0.1 - -2.25) - -2.25) + (0.1 / 3.14159)))))
                                        var v62 uint = 255
                                        if true {
                                            v49 = -847
                                            var v63 int8 = int8(fi48)
                                            if (((1 != 349) && false) || true) {
                                                v64 := v3[2]
                                                if (0 > 0) {
                                                    var v65 uint64 = 2
                                                    v43 = append(v43, fi48)
                                                    _ = v65
                                                } else {
                                                    v43 = append(v43, genFn0(v18, 32767, 128))
                                                }
                                                v64 = ("\n\t" + v18)
                                                _ = v64
                                            }
                                            _ = v63
                                        } else {
                                            v40 += 32767
                                        }
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    v66 := GenS6{F0: (-2.25 <= (1e10 - 1.5)), F1: fi47, F2: v2}
                                    _ = v66
                                }
                                _ = fi48
                                _ = v49
                            }
                            _ = fi47
                        }
                    }
                    _ = v46
                }
            case "foo bar":
                delete(v3, 227)
            }
        }
    case "a":
        fmt.Println(v2, v18, v40)
    default:
        for fi67 := 0; fi67 < 6; fi67++ {
            v40 -= v40
            _ = fi67
        }
    }
    var v68 string = v18
    switch genFn0("hello", 2, 1) {
    case 0:
        v43 = append(v43, -2147483648)
    }
    if false {
        v69 := GenS5{F0: "hello"}
        v70 := GenS4{F0: -1, F1: 1, F2: (v40 * v40)}
        _ = v69
        _ = v70
    } else {
        v43 = append(v43, genFn0(v68, 32767, 1))
        switch (v18 + "\n\t") {
        case "a":
            if (4294967295 != 18446744073709551615) {
                v1[1] = 90
                var v71 string = ""
                _ = v71
            } else {
                v72 := GenS8{F0: (^(v40)), F1: 171}
                _ = v72
            }
        case "foo bar":
            var v73 int8 = (v2 << ((1) & 63))
            _ = v73
        }
        if (32767 < 42) {
            for fi74 := 0; fi74 < 16; fi74++ {
                for fi75 := 0; fi75 < 20; fi75++ {
                    var v76 uint64 = 999
                    var v77 int64 = 337
                    if (!(2 >= 0)) {
                        v3[343] = "\n\t"
                        _ = len(v1)
                    } else {
                        if true {
                            v76 = v76
                            var v78 float32 = (-2.25 * 1e10)
                            var v79 int8 = v2
                            _ = v78
                            _ = v79
                        }
                        switch v68 {
                        case "foo bar":
                            fmt.Println(v68, fi74, v68)
                        case "hello":
                            v68 = ""
                        case "\n\t":
                            switch v18 {
                            case "":
                                delete(v3, 954)
                            case "hello":
                                if ("\n\t" <= v18) {
                                    v76 = v76
                                } else {
                                    switch genFn0("a", (v40 - v40), 187) {
                                    case 32768:
                                        v80 := GenS9{F0: -2147483648, F1: 42, F2: 68}
                                        _ = v80
                                    default:
                                        v1[65535] = 1
                                    }
                                    v68 = ""
                                    var v81 float64 = (float64((0.0 - 1e10)) - (1.5 * -2.25))
                                    _ = v81
                                }
                            case "\n\t":
                                if true {
                                    var v82 float32 = (0.1 / 1.0)
                                    var v83 uint64 = 42
                                    v43 = append(v43, fi74)
                                    _ = v82
                                    _ = v83
                                } else {
                                    var v84 int16 = 1000
                                    fmt.Println(v77)
                                    if len(v43) >= 2 {
                                        v43 = v43[1:len(v43)]
                                    }
                                    _ = v84
                                }
                            }
                        default:
                            switch v68 {
                            case "a":
                                switch (-(fi75)) {
                                case -2147483648:
                                    v77 -= (v77 ^ v77)
                                case 128:
                                    if (!(2147483647 <= 2147483647)) {
                                        v76 -= (v76 % (((^(v76))) | 1))
                                    }
                                default:
                                    if (1.5 > -2.25) {
                                        var v85 int = (fi75 & (fi74 & (fi74 % (((fi74 / (((^(fi74))) | 1))) | 1))))
                                        var v86 uint8 = 42
                                        v2 += (v2 * 42)
                                        _ = v85
                                        _ = v86
                                    } else {
                                        switch "" {
                                        case "":
                                            var v87 int32 = 2147483647
                                            _ = v87
                                        case "\n\t":
                                            if false {
                                                v88 := GenS1{F0: int32(v40), F1: (fi75 << ((0) & 63)), F2: false}
                                                if (v2 != v2) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v89 string = ((v18 + "hello") + ("a" + v68))
                                                        _ = v89
                                                    default:
                                                        var v90 float64 = 3.14159
                                                        _ = v90
                                                    }
                                                }
                                                _ = v88
                                            }
                                        case "foo bar":
                                            v43 = append(v43, genFn0("", (^(v40)), 128))
                                        default:
                                            v68 = v18
                                        }
                                    }
                                }
                            default:
                                fmt.Println(fi75, v40, v2)
                            }
                            var v91 int8 = v2
                            _ = v91
                        }
                    }
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                _ = fi74
            }
            var v92 uint16 = v40
            if true {
                v2 = (^(v2))
                v93 := GenS5{F0: ((v18 + ("foo bar" + v18)) + v18)}
                var v94 uint16 = 135
                _ = v93
                _ = v94
            } else {
                var v95 float32 = (-2.25 + 1.5)
                _ = v95
            }
            _ = v92
        } else {
            if (true && false) {
                for fi96 := 0; fi96 < 7; fi96++ {
                    v68 = "\n\t"
                    _ = fi96
                }
                var v97 uint64 = 42
                _ = v97
            }
            var v98 uint64 = 2
            _ = v98
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v68
}
