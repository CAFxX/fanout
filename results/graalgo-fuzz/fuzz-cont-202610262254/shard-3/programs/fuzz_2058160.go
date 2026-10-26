package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 string; F1 int32; F2 string }
type GenS10 struct { F0 uint64; F1 int32; F2 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8; F1 uint }
func genFn0() uint16 {
    if false {
        if true {
            var v7 int64 = 0
            fmt.Println(v7)
            _ = v7
        } else {
            v8 := []int{}
            _ = v8
        }
        var v9 float64 = ((0.0 - 1e10) / (3.14159 / (float64(-2.25) * 3.14159)))
        v9 = 1e10
        _ = v9
    }
    var v10 uint64 = 688
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int16 = -12
        var v13 bool = (false && false)
        v14 := GenS1{F0: 2}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 string = "a"
    v10 -= 42
    _ = v10
    _ = v15
    return 42
}
func genFn1(p0 int16) string {
    p0 += p0
    var v16 float32 = float32(0.1)
    v16 = v16
    v16 = v16
    fmt.Println(p0)
    _ = p0
    _ = v16
    return ("hello" + "")
}

func main() {
    v1 := []int{32767, 128}
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    v6 := make(map[int]int)
    v1 = append(v1, -128)
    delete(v2, -923)
    v17 := GenS1{F0: 1}
    var v18 float32 = 1.5
    var v19 uint8 = 127
    if true {
        v18 -= 0.1
    }
    var v20 uint = 33
    var v21 uint32 = 638
    switch 2147483647 {
    case 413:
        v20 -= 796
    case -9223372036854775808:
        v18 = (v18 - ((((v18 / float32(float64(v18))) - v18) * (v18 / (1.5 / (v18 + ((v18 * v18) / (0.1 / v18)))))) - float32(3.14159)))
    case 128:
        var v22 int = 4294967295
        _ = v22
    default:
        var v23 uint = v20
        _ = v23
    }
    if false {
        v18 -= (v18 * (v18 + v18))
    } else {
        if (0 == 9223372036854775807) {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 uint8 = (v19 ^ 255)
                _ = fi24
                _ = v25
            }
            if (42 <= 65535) {
                switch 4 {
                case 885:
                    if ("a" <= ("" + "a")) {
                        var v26 int64 = -178
                        for fi27 := 0; fi27 < 18; fi27++ {
                            switch fi27 {
                            case 0:
                                v21 = (v21 % ((v21) | 1))
                            case 665:
                                v28 := GenS1{F0: 0}
                                _ = v28
                            }
                            v29 := GenS2{F0: 855, F1: fi27}
                            switch "hello" {
                            case "a":
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    switch 487 {
                                    case -627:
                                        if false {
                                            var v31 int = fi30
                                            _ = len(v1)
                                            switch (^(fi27)) {
                                            case 2147483647:
                                                v20 += 196
                                            case 9223372036854775807:
                                                if (0 < v19) {
                                                    v32 := GenS3{F0: false, F1: ((int64(v19) > (^(v26))) && (!false))}
                                                    _ = v32
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if false {
                                                        delete(v4, (^(v21)))
                                                        var v33 int8 = -46
                                                        var v34 int64 = (v26 << ((v20) & 63))
                                                        fmt.Println(v33, v26)
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        var v35 uint32 = 136
                                                        v35 = (^(v21))
                                                        var v36 bool = true
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                }
                                            default:
                                                v2[fi30] = "a"
                                            }
                                            _ = v31
                                        } else {
                                            v26 = 42
                                        }
                                    }
                                    if (true || true) {
                                        v37 := GenS2{F0: 193, F1: 32768}
                                        _ = v37
                                    } else {
                                        var v38 uint = 224
                                        if (-1 > 2147483647) {
                                            var v39 int64 = v26
                                            switch "" {
                                            case "foo bar":
                                                v20 = v38
                                            case "":
                                                v40 := GenS1{F0: (v19 | 93)}
                                                _ = v40
                                            default:
                                                v41 := GenS4{F0: 361, F1: fi27}
                                                _ = v41
                                            }
                                            _ = v39
                                        } else {
                                            var v42 int64 = v26
                                            var v43 uint32 = v21
                                            _ = v42
                                            _ = v43
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v38
                                    }
                                    var v44 int8 = 42
                                    v44 -= (v44 % (((v44 | -71)) | 1))
                                    switch fi30 {
                                    case 89:
                                        v26 += (v26 & v26)
                                    case 42:
                                        v6[42] = -1
                                    default:
                                        v19 = v19
                                    }
                                    _ = fi30
                                    _ = v44
                                }
                            case "":
                                var v45 uint32 = v21
                                _ = v45
                            case "foo bar":
                                var v46 int64 = v26
                                _ = v46
                            }
                            _ = fi27
                            _ = v29
                        }
                        for fi47 := 0; fi47 < 8; fi47++ {
                            for fi48 := 0; fi48 < 10; fi48++ {
                                var v49 uint = (v20 ^ (v20 + (v20 * (v20 | (v20 % (((v20 + 4294967295)) | 1))))))
                                v50 := GenS4{F0: 0, F1: fi48}
                                var v51 uint64 = uint64(v19)
                                _ = fi48
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            switch (fi47 | (fi47 | (fi47 - fi47))) {
                            case 1:
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    if true {
                                        v53 := GenS4{F0: 0, F1: (fi47 + (fi52 / (((fi47 * -9223372036854775808)) | 1)))}
                                        _ = v53
                                    }
                                    var v54 uint16 = genFn0()
                                    var v55 uint16 = (v54 & (v54 / ((v54) | 1)))
                                    var v56 int = 2147483647
                                    v21 += (v21 << ((0) & 63))
                                    _ = fi52
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            case 165:
                                v57 := GenS2{F0: 2, F1: fi47}
                                _ = v57
                            default:
                                var v58 float64 = float64((((v18 * v18) + 3.14159) / 1e10))
                                _ = v58
                            }
                            for fi59 := 0; fi59 < 19; fi59++ {
                                if (2147483647 == 42) {
                                    var v60 bool = true
                                    v61 := GenS5{F0: 1, F1: uint16(v21)}
                                    _ = v60
                                    _ = v61
                                }
                                var v62 uint32 = 65535
                                _ = fi59
                                _ = v62
                            }
                            _ = fi47
                        }
                        _ = v26
                    }
                case -2147483648:
                    switch "" {
                    case "":
                        var v63 uint32 = (^(v21))
                        _ = v63
                    case "foo bar":
                        if true {
                            if true {
                                v21 -= v21
                                if true {
                                    v21 += (v21 << ((v20) & 63))
                                    if true {
                                        if (!false) {
                                            if false {
                                                v20 = (v20 * v20)
                                                v21 += 331
                                            }
                                        }
                                        if false {
                                            var v64 int64 = -9223372036854775808
                                            if false {
                                                fmt.Println(v64)
                                                var v65 int64 = (v64 + (v64 / ((v64) | 1)))
                                                v18 += ((v18 * float32(0.0)) + 3.14159)
                                                _ = v65
                                            } else {
                                                v66 := v3["foo bar"]
                                                fmt.Println(v19, v20, v66)
                                                _ = v66
                                            }
                                            _ = v64
                                        }
                                    }
                                    for fi67 := 0; fi67 < 5; fi67++ {
                                        for fi68 := 0; fi68 < 9; fi68++ {
                                            v18 -= v18
                                            switch "" {
                                            case "":
                                                var v69 int8 = -104
                                                _ = v69
                                            }
                                            _ = fi68
                                        }
                                        _ = fi67
                                    }
                                } else {
                                    v20 = v20
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v21 = v21
                                }
                                v21 = (^(v21))
                            }
                            for fi70 := 0; fi70 < 14; fi70++ {
                                fmt.Println(v18)
                                switch fi70 {
                                case 128:
                                    if true {
                                        var v71 int8 = 100
                                        _ = v71
                                    } else {
                                        v72 := GenS6{F0: true, F1: false}
                                        for fi73 := 0; fi73 < 7; fi73++ {
                                            if ((int8(v21) != -128) && false) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 4294967295
                                                }
                                            } else {
                                                if (fi70 == 4294967295) {
                                                    v20 += v20
                                                    var v74 string = ""
                                                    if true {
                                                        if (false || (fi73 == -2147483648)) {
                                                            fmt.Println(fi70, v74, v18)
                                                        }
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v75 uint16 = 0
                                                    var v76 uint64 = 888
                                                    v77 := GenS7{F0: 32767, F1: 1}
                                                    var v78 float64 = 1.5
                                                    var v79 int8 = -128
                                                    v80 := GenS8{F0: v20, F1: 1}
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                switch ("a" + "hello") {
                                                case "\n\t":
                                                    var v81 int16 = 317
                                                    _ = v81
                                                default:
                                                    if (!((v20 << (((^(v20))) & 63)) > v20)) {
                                                        v82 := GenS9{F0: "a", F1: 532, F2: "hello"}
                                                        var v83 bool = false
                                                        v19 = v19
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v84 := GenS10{F0: 4294967295, F1: -1, F2: 42}
                                                        _ = v84
                                                    }
                                                }
                                                v85 := GenS11{F0: 1}
                                                var v86 int16 = 32767
                                                v19 -= 0
                                                _ = v85
                                                _ = v86
                                            }
                                            var v87 string = "hello"
                                            delete(v2, fi70)
                                            v21 = v21
                                            _ = fi73
                                            _ = v87
                                        }
                                        fmt.Println(v18, fi70)
                                        _ = v72
                                    }
                                case -503:
                                    var v88 string = ("a" + (("\n\t" + "a") + "a"))
                                    _ = v88
                                }
                                _ = fi70
                            }
                        }
                    case "hello":
                        switch 42 {
                        case 42:
                            switch 2147483647 {
                            case -2147483648:
                                v20 = v20
                            case 32767:
                                var v89 bool = true
                                _ = v89
                            case 0:
                                v20 = (v20 * (v20 - uint(v19)))
                            }
                        case 4294967295:
                            var v90 int = 9223372036854775807
                            _ = v90
                        }
                    }
                }
            } else {
                var v91 uint64 = 18446744073709551615
                switch -611 {
                case -158:
                    fmt.Println(v91, v19)
                default:
                    v1 = append(v1, -1)
                }
                fmt.Println(v20, v18, v91)
                _ = v91
            }
        } else {
            var v92 uint8 = 0
            _ = v92
        }
        var v93 uint64 = 1
        for fi94 := 0; fi94 < 2; fi94++ {
            v21 -= uint32(v19)
            _ = fi94
        }
        _ = v93
    }
    if (356 == -14) {
        var v95 uint16 = 382
        v96 := GenS4{F0: 4294967295, F1: -1}
        _ = v95
        _ = v96
    } else {
        v97 := GenS12{F0: (v19 ^ v19), F1: v20}
        if ("foo bar" != "a") {
            v98 := v4[(^(v21))]
            delete(v3, "foo bar")
            if ((-32768 == 1000) && true) {
                if v98 {
                    _ = len(v2)
                }
            }
            _ = v98
        } else {
            var v99 bool = (false && true)
            _ = v99
        }
        _ = v97
    }
    for fi100 := 0; fi100 < 15; fi100++ {
        fmt.Println(v18, v19)
        _ = fi100
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
