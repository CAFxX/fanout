package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int64) uint {
    p0 = (^(p0))
    v7 := GenS1{F0: "foo bar", F1: 42, F2: "\n\t"}
    var v8 uint32 = 2147483647
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []int{127, -447}
    var v2 uint = 255
    var v3 uint = 255
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    _ = len(v4)
    if true {
        var v9 uint = 18446744073709551615
        var v10 string = ""
        v3 = v3
        _ = v10
        _ = v9
    } else {
        v11 := v4[1]
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: v11, F1: 2, F2: "\n\t"}
            _ = fi12
            _ = v13
        }
        var v14 float64 = 0.0
        _ = v11
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint64 = 42
        switch 609 {
        case 42:
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(fi17)
                delete(v4, (fi17 / ((fi15) | 1)))
                if false {
                    switch -9223372036854775808 {
                    case -2147483648:
                        fmt.Println(fi15)
                    case 0:
                        fmt.Println(v2, fi15, v3)
                    case -1:
                        fmt.Println(v3, fi17)
                    }
                    v2 += 625
                    v3 += v3
                    switch "\n\t" {
                    case "a":
                        v1 = append(v1, fi15)
                    case "":
                        v18 := v5["hello"]
                        _ = v18
                    default:
                        var v19 int64 = -901
                        _ = v19
                    }
                } else {
                    if (true && (!(false && (69 < -17)))) {
                        var v20 uint = 1
                        var v21 uint8 = 0
                        var v22 uint8 = uint8(v16)
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        if false {
                            v23 := GenS1{F0: "\n\t", F1: 906, F2: "foo bar"}
                            if false {
                                v1 = append(v1, 118)
                            } else {
                                v3 = v2
                                fmt.Println(v3)
                                v2 = (^(v3))
                                var v24 int32 = -879
                                _ = v24
                            }
                            _ = v23
                        } else {
                            switch "a" {
                            case "\n\t":
                                var v25 uint = v3
                                _ = v25
                            }
                            fmt.Println(fi15, v3)
                        }
                    }
                    var v26 string = (("foo bar" + "hello") + "\n\t")
                    var v27 float32 = (float32(float64((1.5 + ((float32(-2.25) / 3.14159) * 0.1)))) / 1e10)
                    var v28 float64 = (0.0 * ((1e10 - (1.5 - 1e10)) * (1e10 + (1.5 - 1.5))))
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi17
            }
        case 1:
            v3 = 42
        }
        _ = fi15
        _ = v16
    }
    var v29 float64 = 1.5
    v6[1] = 1
    for fi30 := 0; fi30 < 11; fi30++ {
        if false {
            switch (fi30 ^ fi30) {
            case 657:
                for fi31 := 0; fi31 < 17; fi31++ {
                    fmt.Println(v29)
                    if true {
                        if false {
                            if (0 != 688) {
                                v29 = 1e10
                                if (true && false) {
                                    if false {
                                        var v32 bool = (!(false && false))
                                        var v33 int64 = 42
                                        _ = v32
                                        _ = v33
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = fi30
                                        }
                                        v34 := GenS2{F0: 18446744073709551615, F1: 32767, F2: 32767}
                                        _ = v34
                                    }
                                    fmt.Println(fi30, v2, fi30)
                                    v2 = 255
                                }
                                if true {
                                    if false {
                                        fmt.Println(v29)
                                    } else {
                                        v2 = (^(v3))
                                        v3 = v3
                                        var v35 uint64 = 0
                                        _ = v35
                                    }
                                    v3 = (v2 ^ 18446744073709551615)
                                    _ = len(v1)
                                    if true {
                                        fmt.Println(v29, v29)
                                        v36 := GenS1{F0: "foo bar", F1: 979, F2: ("hello" + "\n\t")}
                                        if false {
                                            if (("a" + "foo bar") <= ("\n\t" + ("foo bar" + ("" + "\n\t")))) {
                                                var v37 float32 = 3.14159
                                                _ = v37
                                            }
                                            if (!false) {
                                                v29 = float64(0.0)
                                                if ((0 == -12) || (!false)) {
                                                    v3 += v2
                                                    v29 = (v29 - 0.1)
                                                } else {
                                                    fmt.Println(v3)
                                                    var v38 uint8 = 128
                                                    v39 := GenS1{F0: ("\n\t" + ("a" + ("foo bar" + "foo bar"))), F1: uint16(v2), F2: "foo bar"}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                var v40 int8 = -119
                                                v1 = append(v1, fi31)
                                                v3 = (v3 ^ 1)
                                                _ = v40
                                            } else {
                                                v2 = v2
                                            }
                                        }
                                        var v41 int = (fi31 & (fi31 >> ((v2) & 63)))
                                        if true {
                                            var v42 int16 = 0
                                            v2 -= (v3 ^ 255)
                                            v43 := GenS2{F0: v3, F1: (v42 % ((v42) | 1)), F2: (v42 << (((^(v3))) & 63))}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if true {
                                                switch fi31 {
                                                case 359:
                                                    var v44 uint = (v3 >> ((genFn0(9223372036854775807)) & 63))
                                                    _ = v44
                                                case -360:
                                                    switch (v41 | fi31) {
                                                    case 2147483647:
                                                        v45 := GenS1{F0: "foo bar", F1: 49, F2: ""}
                                                        _ = v45
                                                    case 9223372036854775807:
                                                        v46 := GenS2{F0: v3, F1: -329, F2: 1000}
                                                        _ = v46
                                                    case -2147483648:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                case 1:
                                                    v47 := GenS3{F0: -642, F1: 42}
                                                    _ = v47
                                                default:
                                                    v48 := GenS4{F0: -1, F1: (v2 / (((v3 >> (((^(v3))) & 63))) | 1)), F2: 165}
                                                    _ = v48
                                                }
                                            } else {
                                                v49 := GenS4{F0: 475, F1: v3, F2: 42}
                                                var v50 int = -2147483648
                                                fmt.Println(v29)
                                                _ = v49
                                                _ = v50
                                            }
                                            var v51 uint64 = 917
                                            _ = v51
                                        }
                                        v29 += float64(-2.25)
                                        _ = v36
                                        _ = v41
                                    } else {
                                        v52 := GenS5{F0: false, F1: 32767}
                                        var v53 int8 = 42
                                        _ = v52
                                        _ = v53
                                    }
                                    v3 -= v3
                                }
                            } else {
                                var v54 int32 = 1
                                fmt.Println(v3, v54, v54)
                                v2 = (v3 & v2)
                                v2 += (v3 % ((genFn0(-9223372036854775808)) | 1))
                                v55 := GenS4{F0: 42, F1: 4294967295, F2: 255}
                                _ = v54
                                _ = v55
                            }
                            _ = len(v6)
                            v3 = 1
                        }
                        var v56 int = fi31
                        _ = v56
                    } else {
                        var v57 int8 = 1
                        var v58 string = "foo bar"
                        fmt.Println(v2, v58)
                        _ = v57
                        _ = v58
                    }
                    _ = fi31
                }
            case 128:
                fmt.Println(v2, v3, v29)
            case 32767:
                if true {
                    switch "foo bar" {
                    case "a":
                        for fi59 := 0; fi59 < 4; fi59++ {
                            v60 := GenS3{F0: (fi30 * fi30), F1: -875}
                            switch ("\n\t" + (((("\n\t" + "\n\t") + ((((("" + "") + "\n\t") + ((((("\n\t" + "") + "\n\t") + "") + "a") + "foo bar")) + "\n\t") + ("hello" + "\n\t"))) + "a") + "\n\t")) {
                            case "":
                                v29 -= v29
                            case "a":
                                var v61 int32 = 1
                                _ = v61
                            default:
                                var v62 int16 = 1000
                                _ = v62
                            }
                            _ = fi59
                            _ = v60
                        }
                    case "":
                        v63 := GenS3{F0: fi30, F1: -1}
                        _ = v63
                    case "hello":
                        v64 := GenS2{F0: (^(v2)), F1: 883, F2: 1000}
                        _ = v64
                    default:
                        fmt.Println(v3, v3)
                    }
                }
            default:
                switch fi30 {
                case 27:
                    switch "\n\t" {
                    case "hello":
                        if false {
                            for fi65 := 0; fi65 < 13; fi65++ {
                                fmt.Println(v29, v29, v2)
                                _ = fi65
                            }
                        }
                    case "foo bar":
                        var v66 int = 838
                        _ = v66
                    case "\n\t":
                        v67 := GenS6{F0: false, F1: (v2 >> ((v3) & 63))}
                        _ = v67
                    default:
                        for fi68 := 0; fi68 < 18; fi68++ {
                            var v69 bool = (980 > 403)
                            var v70 uint64 = 0
                            _ = fi68
                            _ = v69
                            _ = v70
                        }
                    }
                case 128:
                    var v71 int32 = -2147483648
                    _ = v71
                case -1:
                    switch "foo bar" {
                    case "foo bar":
                        if false {
                            v72 := GenS4{F0: -869, F1: genFn0(42), F2: 1}
                            _ = v72
                        }
                    default:
                        if (-715 == -2147483648) {
                            v73 := GenS3{F0: fi30, F1: 0}
                            _ = v73
                        }
                    }
                default:
                    if (32767 < 2) {
                        v2 = 34
                        fmt.Println(v29, v2, v29)
                    }
                }
            }
        } else {
            v3 += v3
            for fi74 := 0; fi74 < 4; fi74++ {
                v3 = genFn0(-9223372036854775808)
                switch int(v2) {
                case 32768:
                    v29 = 1.5
                case -792:
                    var v75 uint = 18446744073709551615
                    if (222 == 42) {
                        v75 = v2
                    }
                    var v76 uint = v75
                    fmt.Println(v76, fi30, v2)
                    var v77 int64 = 0
                    v78 := GenS6{F0: ((2 > 127) || true), F1: v3}
                    _ = v75
                    _ = v76
                    _ = v77
                    _ = v78
                case 32767:
                    switch fi30 {
                    case 42:
                        v79 := GenS7{F0: 32767}
                        _ = v79
                    }
                }
                _ = fi74
            }
        }
        var v80 uint32 = 42
        _ = fi30
        _ = v80
    }
    var v81 uint64 = 2
    v82 := GenS2{F0: genFn0(42), F1: 0, F2: 42}
    var v83 int64 = 42
    var v84 uint32 = 685
    var v85 bool = (977 < 32768)
    v86 := GenS6{F0: v85, F1: (v2 << ((v3) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
