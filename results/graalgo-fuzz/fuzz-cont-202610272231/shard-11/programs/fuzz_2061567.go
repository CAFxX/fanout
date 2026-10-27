package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int; F1 uint8; F2 int64 }
func genFn0() uint {
    var v5 bool = false
    var v6 string = "foo bar"
    _ = v5
    _ = v6
    return 0
}
func genFn1() int8 {
    var v7 float64 = -2.25
    if true {
        v8 := []int{769, 138, 843}
        fmt.Println(v7)
        v9 := GenS1{F0: 0}
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 0
        switch "\n\t" {
        case "foo bar":
            var v11 bool = false
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                if (895 >= 736) {
                    v13 := GenS1{F0: -11}
                    _ = v13
                } else {
                    var v14 uint16 = 2
                    fmt.Println(v7)
                    _ = v14
                }
                v10 = (-(v10))
                _ = fi12
            }
        }
        _ = v10
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: 127}
        for fi17 := 0; fi17 < 6; fi17++ {
            if (-128 == 7) {
                v7 -= v7
                var v18 uint16 = 1
                switch "foo bar" {
                case "a":
                    fmt.Println(v7)
                    if true {
                        v19 := GenS1{F0: 127}
                        fmt.Println(fi15, fi15)
                        v20 := []int{fi15, (fi17 << ((65535) & 63)), fi15, 128}
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            var v21 int32 = 385
                            var v22 int8 = 51
                            _ = v21
                            _ = v22
                        } else {
                            var v23 uint16 = v18
                            v18 -= v18
                            _ = v23
                        }
                        var v24 uint64 = 264
                        var v25 bool = (true || true)
                        _ = v24
                        _ = v25
                    }
                    switch (fi17 * fi15) {
                    case 1:
                        var v26 uint32 = 459
                        _ = v26
                    case -762:
                        v27 := []int{fi15, fi15, fi15}
                        _ = v27
                    case 4294967295:
                        if (true && (!false)) {
                            var v28 uint32 = 92
                            v29 := GenS1{F0: 42}
                            _ = v28
                            _ = v29
                        }
                        v7 -= 1.5
                        v30 := GenS2{F0: 352}
                        v31 := []int{fi17, fi17}
                        v32 := GenS3{F0: -32768}
                        if (false && false) {
                            v18 += v18
                            v33 := GenS3{F0: -32768}
                            _ = v33
                        }
                        v18 = (v18 ^ (v18 + (v18 & 65535)))
                        _ = v30
                        _ = v31
                        _ = v32
                    default:
                        if false {
                            v34 := []int{}
                            switch (^(fi17)) {
                            case 127:
                                var v35 uint = 650
                                _ = v35
                            default:
                                var v36 int8 = 100
                                _ = v36
                            }
                            if false {
                                switch fi17 {
                                case 42:
                                    if len(v34) > 0 {
                                        v34[0] = fi15
                                    }
                                    v18 = (v18 * v18)
                                case 127:
                                    v7 -= float64(0.0)
                                default:
                                    v37 := GenS4{F0: 255}
                                    _ = v37
                                }
                                if ((fi15 % (((fi15 - (fi15 & (-(fi17))))) | 1)) == (fi17 ^ fi15)) {
                                    var v38 int8 = -128
                                    if false {
                                        v39 := GenS3{F0: 42}
                                        var v40 int = (fi15 & (fi15 << ((18446744073709551615) & 63)))
                                        _ = v39
                                        _ = v40
                                    } else {
                                        if true {
                                            var v41 uint8 = 2
                                            if len(v34) > 0 {
                                                v34[0] = (fi17 << ((258) & 63))
                                            }
                                            _ = v41
                                        }
                                    }
                                    _ = v38
                                } else {
                                    if (-139 <= 1) {
                                        if false {
                                            v42 := GenS1{F0: 42}
                                            _ = v42
                                        } else {
                                            v18 = (v18 * v18)
                                        }
                                        var v43 float32 = float32((v7 - (1.5 + -2.25)))
                                        v43 -= v43
                                        _ = v43
                                    }
                                    v44 := GenS5{F0: 42, F1: 864, F2: 109}
                                    var v45 float32 = 1.5
                                    _ = v44
                                    _ = v45
                                }
                            }
                            _ = v34
                        } else {
                            v7 -= v7
                            fmt.Println(v18, v7)
                            var v46 int64 = -306
                            _ = v46
                        }
                    }
                case "\n\t":
                    v7 = 3.14159
                default:
                    v47 := GenS6{F0: "a"}
                    _ = v47
                }
                _ = v18
            } else {
                if true {
                    switch "hello" {
                    case "a":
                        var v48 float32 = 1.5
                        _ = v48
                    }
                }
                var v49 uint = 2
                v50 := []int{(fi15 + fi15), (fi17 + fi17), (fi17 ^ (fi15 ^ (fi15 + fi17)))}
                fmt.Println(v49, fi15, fi15)
                _ = v49
                _ = v50
            }
            _ = fi17
        }
        var v51 float64 = (float64(-2.25) / v7)
        _ = fi15
        _ = v16
        _ = v51
    }
    fmt.Println(v7)
    v7 = 3.14159
    _ = v7
    return 100
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 1e10
    v3 := make(map[string]int)
    var v4 int64 = 954
    if false {
        fmt.Println(v4)
    } else {
        v52 := GenS1{F0: -128}
        fmt.Println(v2)
        for fi53 := 0; fi53 < 12; fi53++ {
            if false {
                if false {
                    v54 := GenS3{F0: -637}
                    for fi55 := 0; fi55 < 12; fi55++ {
                        v4 += v4
                        _ = fi55
                    }
                    for fi56 := 0; fi56 < 5; fi56++ {
                        if true {
                            switch ("\n\t" + "") {
                            case "foo bar":
                                if false {
                                    v4 -= (^(v4))
                                    v4 = -917
                                    v57 := GenS1{F0: -1}
                                    _ = v57
                                }
                            case "\n\t":
                                var v58 int8 = -1
                                _ = v58
                            case "hello":
                                v2 += v2
                            default:
                                v2 -= v2
                            }
                            if ((0 >= (^(fi56))) && true) {
                                if false {
                                    v59 := GenS5{F0: 114, F1: 65535, F2: 1}
                                    var v60 bool = true
                                    _ = v59
                                    _ = v60
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        v2 = v2
                                    case "":
                                        if (((fi53 << ((18446744073709551615) & 63)) > -793) && ((138 <= 2) && (!(-32768 <= 32767)))) {
                                            switch 32768 {
                                            case -30:
                                                if true {
                                                    var v61 uint64 = 18446744073709551615
                                                    switch (-(fi56)) {
                                                    case 1:
                                                        _ = len(v1)
                                                    case 2147483647:
                                                        v62 := GenS2{F0: 1}
                                                        _ = v62
                                                    }
                                                    _ = v61
                                                }
                                            default:
                                                var v63 uint32 = 42
                                                _ = v63
                                            }
                                            var v64 float32 = float32(v2)
                                            if (!(false && false)) {
                                                if false {
                                                    switch -2147483648 {
                                                    case 9223372036854775807:
                                                        var v65 uint32 = 2
                                                        _ = v65
                                                    case -773:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v66 string = ("hello" + (("foo bar" + "\n\t") + "\n\t"))
                                                            _ = v66
                                                        default:
                                                            var v67 uint64 = 101
                                                            _ = v67
                                                        }
                                                    case -2147483648:
                                                        if false {
                                                            switch 0 {
                                                            case 1:
                                                                var v68 string = "hello"
                                                                _ = v68
                                                            case 32767:
                                                                v69 := GenS7{F0: 63}
                                                                _ = v69
                                                            case -899:
                                                                if true {
                                                                    var v70 int = (fi56 + -67)
                                                                    var v71 int16 = -288
                                                                    var v72 float64 = 1.5
                                                                    _ = v70
                                                                    _ = v71
                                                                    _ = v72
                                                                } else {
                                                                    fmt.Println(v64)
                                                                }
                                                            }
                                                            fmt.Println(v2, fi56)
                                                        } else {
                                                            if (false || (!(1 != 127))) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 42
                                                                }
                                                                var v73 bool = true
                                                                var v74 int32 = 1
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                delete(v3, "")
                                                                _ = len(v3)
                                                                switch "a" {
                                                                case "a":
                                                                    v4 -= v4
                                                                case "foo bar":
                                                                    fmt.Println(v2, v64, fi53)
                                                                case "":
                                                                    fmt.Println(v64)
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v75 float32 = (v64 / v64)
                                                        _ = v75
                                                    }
                                                    v76 := GenS8{F0: "", F1: -1}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v76
                                                }
                                            } else {
                                                var v77 uint64 = 42
                                                fmt.Println(v64, v4)
                                                var v78 int32 = -233
                                                _ = v77
                                                _ = v78
                                            }
                                            _ = v64
                                        }
                                    case "a":
                                        fmt.Println(v2, fi53)
                                        var v79 float64 = (v2 + (v2 * v2))
                                        v80 := GenS3{F0: -32768}
                                        _ = v79
                                        _ = v80
                                    default:
                                        v81 := GenS8{F0: "", F1: (^(v4))}
                                        _ = v81
                                    }
                                    v2 = v2
                                }
                            } else {
                                fmt.Println(v2, v2, v4)
                                v1 = append(v1, 643)
                            }
                            var v82 float32 = -2.25
                            _ = v82
                        }
                        _ = fi56
                    }
                    _ = v54
                }
                v83 := GenS3{F0: -778}
                _ = v83
            }
            _ = fi53
        }
        _ = v52
    }
    var v84 float32 = ((((1e10 / 1e10) * 0.0) / (3.14159 / (1e10 - (1.5 / -2.25)))) - 0.0)
    _ = len(v1)
    switch 2147483647 {
    case 9223372036854775807:
        fmt.Println(v4)
    default:
        if (2 == 42) {
            for fi85 := 0; fi85 < 5; fi85++ {
                v4 += (v4 + (^(v4)))
                _ = len(v3)
                var v86 uint16 = 429
                _ = fi85
                _ = v86
            }
        } else {
            v87 := GenS7{F0: genFn0()}
            _ = v87
        }
    }
    var v88 float64 = (float64(((0.1 / v84) + v84)) - v2)
    v84 = ((v84 * v84) / 1.0)
    v4 -= (v4 | (^(v4)))
    var v89 bool = (!false)
    switch "a" {
    case "hello":
        for fi90 := 0; fi90 < 12; fi90++ {
            v91 := GenS8{F0: ("hello" + "foo bar"), F1: 1}
            _ = fi90
            _ = v91
        }
    case "":
        delete(v3, "\n\t")
    case "foo bar":
        fmt.Println(v2, v89)
    }
    v84 += v84
    fmt.Println(v4)
    if ((((v84 + ((float32(3.14159) + v84) + -2.25)) / float32(v88)) != 3.14159) && (75 != (^(v4)))) {
        v92 := GenS9{F0: -1, F1: 42, F2: (v4 / ((v4) | 1))}
        var v93 bool = v89
        if v89 {
            v94 := v3[("" + ("a" + ("foo bar" + ("" + "a"))))]
            delete(v3, ("" + "hello"))
            _ = v94
        }
        _ = v92
        _ = v93
    } else {
        var v95 int8 = 100
        fmt.Println(v4)
        var v96 uint8 = 1
        _ = v95
        _ = v96
    }
    var v97 uint16 = 324
    _ = len(v1)
    var v98 uint8 = 128
    var v99 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v84
    _ = v88
    _ = v89
    _ = v97
    _ = v98
    _ = v99
}
