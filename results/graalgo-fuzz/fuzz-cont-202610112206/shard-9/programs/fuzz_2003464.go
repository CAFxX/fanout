package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint; F2 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int }
func genFn0() bool {
    v7 := GenS1{F0: 2147483647, F1: 372, F2: 83}
    var v8 uint = 4294967295
    for fi9 := 0; fi9 < 16; fi9++ {
        v8 += (v8 + 1)
        v10 := GenS2{F0: 0}
        v11 := GenS3{F0: 2147483647}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int8 = -1
    v13 := GenS2{F0: 224}
    _ = v12
    _ = v13
    _ = v7
    _ = v8
    return ("\n\t" >= ("" + "foo bar"))
}
func genFn1() uint8 {
    var v14 float64 = (-2.25 / (0.0 + float64(1e10)))
    var v15 uint16 = 661
    var v16 int16 = 42
    _ = v14
    _ = v15
    _ = v16
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint = 871
    var v4 int8 = -128
    var v5 float32 = (((0.0 - 0.1) - 0.1) - (0.1 - -2.25))
    var v6 uint8 = 128
    if true {
        v6 = v6
        v17 := []int{}
        var v18 int8 = -74
        _ = v17
        _ = v18
    }
    var v19 float64 = float64(0.1)
    var v20 bool = false
    if (-229 < 701) {
        if ((-1 >= -2147483648) || v20) {
            if ((v19 * 3.14159) == (v19 / (v19 / 1e10))) {
                var v21 float64 = v19
                _ = v21
            } else {
                if false {
                    v22 := GenS4{F0: (^(v3))}
                    _ = v22
                }
                if v20 {
                    v6 += v6
                    v23 := GenS5{F0: 227, F1: uint(v6), F2: "foo bar"}
                    _ = v23
                } else {
                    var v24 uint8 = 42
                    for fi25 := 0; fi25 < 15; fi25++ {
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v27 := v1[296]
                            fmt.Println(v20, v4)
                            _ = fi26
                            _ = v27
                        }
                        if (2147483648 <= 2147483647) {
                            fmt.Println(v4)
                        } else {
                            var v28 uint8 = (v6 + 127)
                            if (v20 && true) {
                                fmt.Println(v20, v5)
                                v28 = (v6 - genFn1())
                                if (v20 || (688 <= 9223372036854775807)) {
                                    if (!(false || (!(42 < 4294967295)))) {
                                        for fi29 := 0; fi29 < 5; fi29++ {
                                            var v30 uint8 = (^(v28))
                                            var v31 uint64 = 0
                                            v6 = genFn1()
                                            var v32 uint32 = 65535
                                            var v33 int64 = 42
                                            _ = fi29
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        switch "hello" {
                                        case "":
                                            var v34 float64 = 1e10
                                            _ = v34
                                        case "\n\t":
                                            if (!v20) {
                                                v28 -= (v6 | (v24 + (^(v24))))
                                            } else {
                                                for fi35 := 0; fi35 < 9; fi35++ {
                                                    v36 := []int{fi35, -9223372036854775808}
                                                    if v20 {
                                                        switch ("a" + (("hello" + "hello") + ("" + "foo bar"))) {
                                                        case "":
                                                            v20 = false
                                                        case "foo bar":
                                                            var v37 bool = (197 < -976)
                                                            _ = v37
                                                        }
                                                        v38 := GenS1{F0: 0, F1: 65535, F2: v4}
                                                        _ = v38
                                                    } else {
                                                        switch -9223372036854775808 {
                                                        case 1:
                                                            switch ("hello" + "") {
                                                            case "foo bar":
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(fi35, v20, fi25)
                                                                case "foo bar":
                                                                    var v39 uint8 = v28
                                                                    _ = v39
                                                                default:
                                                                    var v40 int64 = 9223372036854775807
                                                                    _ = v40
                                                                }
                                                            }
                                                        case 2147483647:
                                                            fmt.Println(v6)
                                                        case 128:
                                                            v5 -= ((v5 - v5) - (v5 - v5))
                                                        default:
                                                            v6 -= v24
                                                        }
                                                        v1[2147483647] = ((v20 || (("a" + "foo bar") == "foo bar")) || (!((!false) && ((0 == 0) && (42 != 32767)))))
                                                    }
                                                    var v41 int16 = -32768
                                                    _ = fi35
                                                    _ = v36
                                                    _ = v41
                                                }
                                            }
                                        case "foo bar":
                                            switch (fi25 % ((fi25) | 1)) {
                                            case 42:
                                                v24 += (v6 % (((v24 & v24)) | 1))
                                            case -2147483648:
                                                v19 += 1e10
                                            }
                                        default:
                                            var v42 int16 = 1000
                                            _ = v42
                                        }
                                    }
                                } else {
                                    v43 := GenS4{F0: 42}
                                    for fi44 := 0; fi44 < 3; fi44++ {
                                        if (!(-9223372036854775808 > -9223372036854775808)) {
                                            if v20 {
                                                var v45 uint = (^(v3))
                                                var v46 int16 = 32767
                                                fmt.Println(v6, v28, v6)
                                                _ = v45
                                                _ = v46
                                            } else {
                                                v6 += 220
                                                v6 -= 0
                                                v47 := GenS6{F0: 2147483647}
                                                if false {
                                                    var v48 uint16 = 1
                                                    v4 -= v4
                                                    _ = v48
                                                }
                                                v49 := GenS2{F0: 1}
                                                var v50 uint = v3
                                                var v51 uint = 847
                                                var v52 int16 = 1
                                                _ = v47
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            }
                                            v4 = v4
                                        }
                                        _ = fi44
                                    }
                                    fmt.Println(v6, v5)
                                    _ = v43
                                }
                            } else {
                                switch "" {
                                case "a":
                                    var v53 uint16 = 885
                                    _ = v53
                                case "":
                                    if (fi25 > 122) {
                                        v24 = v24
                                    }
                                case "foo bar":
                                    switch fi25 {
                                    case -9223372036854775808:
                                        v24 = v24
                                    }
                                default:
                                    v5 = 0.1
                                }
                                var v54 uint32 = 2
                                _ = v54
                            }
                            _ = v28
                        }
                        v55 := GenS5{F0: fi25, F1: (v3 % ((v3) | 1)), F2: "foo bar"}
                        _ = fi25
                        _ = v55
                    }
                    if v20 {
                        var v56 int8 = (v4 % ((v4) | 1))
                        var v57 float32 = v5
                        _ = v56
                        _ = v57
                    } else {
                        v4 += (v4 + v4)
                        v3 = (^(v3))
                    }
                    _ = v24
                }
            }
        }
        v58 := GenS3{F0: -70}
        for fi59 := 0; fi59 < 7; fi59++ {
            v60 := []int{(-(fi59)), fi59, fi59, 1}
            fmt.Println(v19, v20)
            _ = fi59
            _ = v60
        }
        _ = v58
    } else {
        for fi61 := 0; fi61 < 2; fi61++ {
            for fi62 := 0; fi62 < 10; fi62++ {
                var v63 uint8 = 0
                var v64 int16 = 32767
                v4 -= v4
                if (1 != 623) {
                    if v20 {
                        if (((!((v64 | v64) <= (-(v64)))) && v20) || false) {
                            v65 := []int{}
                            fmt.Println(v19, fi61, v20)
                            fmt.Println(v5, v20)
                            _ = v65
                        }
                        v64 += v64
                        var v66 int32 = 2147483647
                        _ = v66
                    }
                    v4 = -128
                }
                if v20 {
                    fmt.Println(v3, v63, v64)
                }
                _ = fi62
                _ = v63
                _ = v64
            }
            _ = fi61
        }
        for fi67 := 0; fi67 < 20; fi67++ {
            var v68 string = "hello"
            v3 -= (^(v3))
            _ = fi67
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 1; fi69++ {
        v70 := GenS7{F0: -1}
        if false {
            for fi71 := 0; fi71 < 2; fi71++ {
                var v72 int64 = 1
                _ = fi71
                _ = v72
            }
            v4 = -62
            v73 := []int{-614, (fi69 / ((4294967295) | 1))}
            _ = v73
        }
        _ = fi69
        _ = v70
    }
    var v74 int64 = 0
    if (18446744073709551615 > 62) {
        v3 += 457
        var v75 uint32 = 2147483648
        for fi76 := 0; fi76 < 12; fi76++ {
            _ = len(v2)
            for fi77 := 0; fi77 < 16; fi77++ {
                fmt.Println(v4, v6, v5)
                v74 = (v74 ^ (v74 << ((v3) & 63)))
                var v78 int32 = -2147483648
                v79 := GenS2{F0: -1}
                var v80 bool = false
                _ = len(v1)
                _ = fi77
                _ = v78
                _ = v79
                _ = v80
            }
            _ = fi76
        }
        _ = v75
    }
    v5 -= 0.0
    v4 = (v4 ^ v4)
    delete(v2, 0)
    v81 := []int{}
    if v20 {
        var v82 uint = 2
        _ = v82
    }
    _ = len(v1)
    if ((v74 + 541) < (v74 | v74)) {
        var v83 float32 = (float32((0.0 * v19)) * (((v5 + v5) / -2.25) * 0.0))
        fmt.Println(v4, v19)
        _ = v83
    } else {
        if len(v81) > 0 {
            v81[0] = 128
        }
        for fi84 := 0; fi84 < 5; fi84++ {
            var v85 bool = v20
            v86 := GenS5{F0: fi84, F1: (v3 * 307), F2: (((("hello" + "") + "a") + (("" + "foo bar") + ("" + "foo bar"))) + "foo bar")}
            var v87 string = "a"
            _ = fi84
            _ = v85
            _ = v86
            _ = v87
        }
    }
    for fi88 := 0; fi88 < 20; fi88++ {
        v81 = append(v81, fi88)
        if (true || v20) {
            for fi89 := 0; fi89 < 13; fi89++ {
                if v20 {
                    _ = len(v81)
                    if true {
                        fmt.Println(fi89, v5)
                    } else {
                        var v90 float32 = (-2.25 - 0.0)
                        var v91 uint64 = 18446744073709551615
                        _ = v90
                        _ = v91
                    }
                    if len(v81) > 0 {
                        v81[0] = fi89
                    }
                    if ((fi89 != fi88) || v20) {
                        var v92 uint16 = 689
                        var v93 uint8 = (v6 ^ (v6 | v6))
                        _ = v92
                        _ = v93
                    }
                    v3 += v3
                }
                if v20 {
                    v81 = append(v81, (fi88 - (-(fi88))))
                } else {
                    switch "\n\t" {
                    case "a":
                        var v94 int64 = (v74 | v74)
                        _ = v94
                    case "":
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            switch "" {
                            case "a":
                                v1[0] = false
                            case "foo bar":
                                var v95 float64 = ((v19 / 3.14159) * v19)
                                _ = v95
                            }
                        case "hello":
                            switch -612 {
                            case -1:
                                fmt.Println(v3, v20)
                            case 2147483647:
                                v5 += v5
                            case 32768:
                                var v96 int32 = int32(fi89)
                                _ = v96
                            }
                        case "a":
                            var v97 int64 = v74
                            _ = v97
                        default:
                            if (fi88 > (fi88 & 42)) {
                                v74 -= (-(v74))
                                v6 = (v6 / ((v6) | 1))
                                var v98 float32 = 1e10
                                _ = v98
                            } else {
                                v99 := GenS4{F0: (v3 + v3)}
                                v100 := GenS8{F0: fi89, F1: fi89}
                                _ = v100
                                _ = v99
                            }
                        }
                    default:
                        if len(v81) > 0 {
                            v81[0] = 1
                        }
                    }
                }
                v74 += -1
                _ = fi89
            }
            if v20 {
                if len(v81) >= 2 {
                    v81 = v81[1:len(v81)]
                }
                fmt.Println(v6, v6)
            }
            v101 := GenS4{F0: (^(v3))}
            _ = v101
        } else {
            fmt.Println(v20)
        }
        var v102 string = "hello"
        _ = fi88
        _ = v102
    }
    if len(v81) >= 2 {
        v81 = v81[1:len(v81)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("slicelen v81:", len(v81), v81)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v81
}
