package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint16; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint32) int8 {
    var v7 uint8 = 115
    p0 = p0
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    v4 := []uint32{65535, 2147483647, 558}
    var v5 uint8 = 1
    var v6 string = ("foo bar" + "hello")
    fmt.Println(v2, v2, v5)
    var v8 bool = true
    if len(v4) > 0 {
        v4[0] = 1
    }
    v9 := GenS1{F0: 32767}
    var v10 int16 = 372
    if (!(-239 == (^(v10)))) {
        fmt.Println(v5, v10)
        var v11 float32 = 1e10
        for fi12 := 0; fi12 < 6; fi12++ {
            switch fi12 {
            case 658:
                var v13 uint64 = (v2 << ((65535) & 63))
                _ = v13
            }
            if v8 {
                v10 += (-(v10))
            } else {
                if v8 {
                    var v14 float32 = v11
                    var v15 int64 = -1
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if (uint16(v2) == v1) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v15 = int64(v2)
                        } else {
                            if ((^(v2)) > v2) {
                                var v17 float64 = 1e10
                                if true {
                                    v1 += (^(v1))
                                    v8 = v8
                                }
                                _ = v17
                            } else {
                                delete(v3, 2)
                                switch fi12 {
                                case -1:
                                    v1 -= v1
                                case 0:
                                    var v18 uint = 2
                                    _ = v18
                                }
                            }
                            if false {
                                v3[42] = v8
                            } else {
                                var v19 int = fi16
                                _ = v19
                            }
                            fmt.Println(v8, v5, v5)
                        }
                        if v8 {
                            v20 := v3[299]
                            v21 := GenS2{F0: fi12}
                            _ = v20
                            _ = v21
                        } else {
                            v8 = (false && ((((v2 % ((18446744073709551615) | 1)) == v2) && ((((v6 + "\n\t") + ((v6 + v6) + v6)) + "\n\t") > v6)) && v8))
                            if true {
                                if v8 {
                                    v11 = 3.14159
                                }
                                fmt.Println(v5, v1, fi16)
                                if false {
                                    v8 = v8
                                    v10 = v10
                                } else {
                                    var v22 uint8 = 2
                                    _ = v22
                                }
                                v23 := GenS1{F0: fi16}
                                if (660 <= 2) {
                                    v2 += (^(v2))
                                    switch v6 {
                                    case "hello":
                                        if len(v4) > 0 {
                                            v4[0] = 2147483648
                                        }
                                    case "foo bar":
                                        var v24 int8 = -1
                                        _ = v24
                                    default:
                                        if v8 {
                                            v25 := GenS3{F0: 42}
                                            _ = v25
                                        } else {
                                            var v26 float64 = (float64((3.14159 - (v14 - v11))) + (3.14159 * (1.5 * ((1e10 + 3.14159) - ((-2.25 - -2.25) / 1e10)))))
                                            var v27 uint16 = v1
                                            if (2147483647 < -1) {
                                                v28 := GenS1{F0: (^(fi16))}
                                                v2 -= (v2 % ((0) | 1))
                                                if v8 {
                                                    var v29 uint32 = 42
                                                    fmt.Println(v1)
                                                    v29 -= v29
                                                    _ = v29
                                                }
                                                _ = v28
                                            }
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                    switch "hello" {
                                    case "a":
                                        switch (((v6 + "a") + (v6 + v6)) + v6) {
                                        case "":
                                            var v30 uint8 = (v5 ^ v5)
                                            _ = v30
                                        case "\n\t":
                                            v14 += v11
                                        case "foo bar":
                                            fmt.Println(fi12)
                                            fmt.Println(v11, fi12, v1)
                                        default:
                                            var v31 int16 = (v10 >> ((4294967295) & 63))
                                            _ = v31
                                        }
                                    case "\n\t":
                                        fmt.Println(v11, v2, v1)
                                    }
                                }
                                _ = v23
                            }
                        }
                        if true {
                            fmt.Println(v10, v11, v10)
                            if v8 {
                                var v32 bool = ("foo bar" <= "hello")
                                var v33 int32 = -2147483648
                                if v8 {
                                    v2 = v2
                                } else {
                                    var v34 uint16 = 310
                                    _ = v34
                                }
                                _ = v32
                                _ = v33
                            } else {
                                if v8 {
                                    v2 -= v2
                                    var v35 uint64 = v2
                                    var v36 float32 = v11
                                    _ = v35
                                    _ = v36
                                } else {
                                    v5 = (v5 | v5)
                                }
                            }
                            var v37 string = "a"
                            v8 = v8
                            switch v37 {
                            case "":
                                switch 778 {
                                case 127:
                                    _ = len(v4)
                                }
                            case "foo bar":
                                switch (fi16 | (fi12 << ((0) & 63))) {
                                case -9223372036854775808:
                                    v1 += v1
                                case -426:
                                    v38 := GenS1{F0: fi16}
                                    _ = v38
                                case 4294967295:
                                    var v39 string = v6
                                    _ = v39
                                default:
                                    fmt.Println(fi16, fi16, v10)
                                }
                            }
                            var v40 uint64 = v2
                            v1 += 255
                            _ = v37
                            _ = v40
                        }
                        _ = fi16
                    }
                    _ = v14
                    _ = v15
                }
                v41 := GenS3{F0: -1}
                var v42 float32 = (v11 / 1.5)
                _ = v41
                _ = v42
            }
            _ = fi12
        }
        _ = v11
    } else {
        v43 := GenS3{F0: int32(v5)}
        _ = v43
    }
    var v44 uint32 = 2
    switch 647 {
    case -1:
        for fi45 := 0; fi45 < 5; fi45++ {
            v8 = true
            _ = fi45
        }
    case 32768:
        if (int8(v2) < -128) {
            var v46 float64 = 1e10
            var v47 uint16 = v1
            fmt.Println(v8)
            _ = v46
            _ = v47
        }
    case -2147483648:
        for fi48 := 0; fi48 < 14; fi48++ {
            for fi49 := 0; fi49 < 13; fi49++ {
                v50 := GenS1{F0: fi48}
                if (genFn0((v44 >> ((4294967295) & 63))) != -25) {
                    var v51 int32 = 2147483647
                    var v52 int = fi49
                    _ = v51
                    _ = v52
                }
                _ = fi49
                _ = v50
            }
            for fi53 := 0; fi53 < 10; fi53++ {
                var v54 uint32 = v44
                if (!((1e10 / 1.5) >= 0.0)) {
                    _ = len(v4)
                    _ = len(v4)
                } else {
                    switch ("\n\t" + v6) {
                    case "hello":
                        v55 := GenS2{F0: fi53}
                        _ = v55
                    case "a":
                        var v56 bool = v8
                        _ = v56
                    case "":
                        fmt.Println(v5, fi53, v6)
                    default:
                        v10 -= v10
                    }
                }
                _ = fi53
                _ = v54
            }
            _ = fi48
        }
    }
    switch 128 {
    case 128:
        for fi57 := 0; fi57 < 11; fi57++ {
            switch "hello" {
            case "hello":
                var v58 uint = 255
                _ = v58
            case "foo bar":
                for fi59 := 0; fi59 < 2; fi59++ {
                    if ((float64((0.1 - 1e10)) >= float64(-2.25)) && v8) {
                        var v60 float64 = 0.1
                        var v61 bool = (uint16(v5) != (^(v1)))
                        _ = v60
                        _ = v61
                    }
                    var v62 float64 = 3.14159
                    if ((^(fi57)) <= (^(fi57))) {
                        var v63 int64 = int64(v10)
                        v3[(v44 - (v44 - v44))] = false
                        v63 += -356
                        fmt.Println(v10, v63)
                        if false {
                            fmt.Println(v2, v10, v2)
                            if (2147483647 >= 2147483647) {
                                v64 := GenS4{F0: 127, F1: 42, F2: (v1 < v1)}
                                var v65 uint16 = v1
                                if v8 {
                                    var v66 uint = 77
                                    var v67 int64 = -772
                                    _ = v66
                                    _ = v67
                                }
                                _ = v64
                                _ = v65
                            } else {
                                v62 = v62
                            }
                            switch 601 {
                            case 9223372036854775807:
                                v68 := v3[v44]
                                _ = v68
                            case 771:
                                if v8 {
                                    v63 += (v63 >> ((18446744073709551615) & 63))
                                    var v69 float64 = 1e10
                                    _ = v69
                                } else {
                                    delete(v3, (^(v44)))
                                    if v8 {
                                        switch ("\n\t" + ("hello" + "hello")) {
                                        case "":
                                            if true {
                                                v6 = v6
                                            } else {
                                                if v8 {
                                                    var v70 int16 = 562
                                                    v71 := GenS5{F0: v5}
                                                    if v8 {
                                                        if (128 != 127) {
                                                            var v72 string = v6
                                                            v5 = uint8(v1)
                                                            var v73 uint32 = v44
                                                            var v74 uint8 = (v5 + v5)
                                                            _ = v72
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        var v75 int16 = v70
                                                        v70 -= (v75 | v10)
                                                        _ = v75
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        }
                                        _ = len(v3)
                                        switch v6 {
                                        case "\n\t":
                                            var v76 uint32 = 42
                                            _ = v76
                                        case "hello":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case "foo bar":
                                            v77 := GenS5{F0: (v5 ^ 42)}
                                            var v78 uint16 = (v1 >> ((42) & 63))
                                            v2 = v2
                                            fmt.Println(v78)
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                    v79 := v3[v44]
                                    _ = v79
                                }
                            }
                        }
                        _ = v63
                    } else {
                        var v80 uint16 = 896
                        v80 = (v1 | uint16(v2))
                        _ = v80
                    }
                    _ = fi59
                    _ = v62
                }
            }
            v2 = 0
            var v81 uint8 = 128
            _ = fi57
            _ = v81
        }
    default:
        _ = len(v4)
    }
    v82 := GenS3{F0: 0}
    v6 = v6
    for fi83 := 0; fi83 < 9; fi83++ {
        fmt.Println(fi83)
        fmt.Println(v8, fi83, fi83)
        for fi84 := 0; fi84 < 14; fi84++ {
            var v85 int16 = v10
            fmt.Println(v5, v44)
            var v86 int16 = (^(v10))
            var v87 bool = true
            v88 := GenS6{F0: 837}
            _ = fi84
            _ = v85
            _ = v86
            _ = v87
            _ = v88
        }
        _ = fi83
    }
    if (!v8) {
        var v89 int = 32767
        _ = v89
    } else {
        v10 -= v10
        var v90 uint64 = (v2 * (^(v2)))
        if v8 {
            if (4294967295 > 245) {
                var v91 uint16 = 255
                var v92 uint32 = v44
                if v8 {
                    v44 -= (v92 ^ (v92 | v92))
                }
                _ = v91
                _ = v92
            }
            v90 = 372
        } else {
            var v93 uint16 = uint16(v5)
            fmt.Println(v1, v8, v2)
            v93 = (v93 >> ((1) & 63))
            _ = v93
        }
        _ = v90
    }
    var v94 float64 = float64(3.14159)
    if (false && v8) {
        for fi95 := 0; fi95 < 19; fi95++ {
            var v96 string = v6
            if (v8 && (456 < -999)) {
                v96 = ("foo bar" + ("" + (((v96 + "\n\t") + "") + (v6 + ((v6 + (("hello" + (v6 + ((v6 + ("hello" + ("" + ("foo bar" + ("\n\t" + "\n\t"))))) + (v96 + "foo bar")))) + (((((v96 + "hello") + "a") + v6) + (v96 + "")) + v6))) + "hello")))))
                v97 := GenS3{F0: -1}
                _ = v97
            } else {
                v3[v44] = (fi95 != fi95)
                fmt.Println(v44)
            }
            for fi98 := 0; fi98 < 20; fi98++ {
                if false {
                    v96 = "hello"
                    v8 = ((!(false && true)) && v8)
                    v10 += -380
                }
                v99 := GenS6{F0: uint64(v1)}
                var v100 float64 = ((v94 - v94) - (-2.25 - (1.5 + 1e10)))
                _ = fi98
                _ = v100
                _ = v99
            }
            _ = fi95
            _ = v96
        }
    }
    switch -701 {
    case 2147483647:
        var v101 bool = true
        _ = v101
    default:
        v4 = append(v4, v44)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v8
    _ = v82
    _ = v9
    _ = v94
}
