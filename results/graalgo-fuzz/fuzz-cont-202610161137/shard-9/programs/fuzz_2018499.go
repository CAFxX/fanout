package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint16; F1 int16; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int16; F1 int; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 255
        _ = fi6
        _ = v7
    }
    if true {
        var v8 int8 = 1
        if false {
            v9 := []int{-312, 128}
            var v10 int16 = 42
            v11 := GenS1{F0: 0, F1: 857}
            _ = v10
            _ = v11
            _ = v9
        } else {
            v8 = -1
            var v12 bool = false
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := []int{(fi13 * 674), (-(fi13)), fi13, 32768}
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        _ = v8
    }
    if (true || false) {
        v15 := []int{1, -1}
        v15 = append(v15, 216)
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint64 = 2
            for fi18 := 0; fi18 < 12; fi18++ {
                v17 += (^(v17))
                _ = fi18
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        if (!true) {
            var v19 float64 = 0.1
            _ = v19
        }
        if true {
            var v20 uint16 = 32767
            v21 := GenS2{F0: false, F1: (^(v20)), F2: ("a" + "a")}
            _ = v20
            _ = v21
        } else {
            v22 := GenS3{F0: 42, F1: 1, F2: 32767}
            _ = v22
        }
    }
    var v23 int = -151
    _ = v23
    return 269
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1.5
    var v3 uint8 = 255
    var v4 string = "foo bar"
    var v5 uint64 = 42
    fmt.Println(v3, v2)
    fmt.Println(v5)
    v5 += v5
    v5 = v5
    v24 := []int{-369, 127, -354}
    for fi25 := 0; fi25 < 13; fi25++ {
        fmt.Println(v2, v2, v3)
        v2 += v2
        _ = fi25
    }
    v4 = ((v4 + v4) + v4)
    var v26 uint = 1
    v2 = ((1.5 - (0.0 * (v2 - (1e10 / -2.25)))) - v2)
    _ = len(v1)
    var v27 uint16 = 2
    v24 = append(v24, 594)
    var v28 bool = false
    for fi29 := 0; fi29 < 14; fi29++ {
        switch (v4 + ((((((v4 + "") + v4) + ("" + "")) + v4) + "hello") + (v4 + v4))) {
        case "a":
            var v30 int16 = 42
            _ = v30
        case "\n\t":
            var v31 int = int(v5)
            _ = v31
        case "hello":
            if (128 != (v3 * v3)) {
                v3 = (v3 >> (((v26 & v26)) & 63))
            }
        default:
            var v32 int8 = 109
            _ = v32
        }
        switch 32768 {
        case 59:
            if v28 {
                if (("\n\t" < v4) && ("" <= ("\n\t" + v4))) {
                    v33 := GenS4{F0: 536, F1: v27}
                    var v34 int8 = -128
                    _ = v33
                    _ = v34
                }
            }
        default:
            if (0 == -833) {
                v35 := GenS1{F0: fi29, F1: -9223372036854775808}
                if ((v26 / ((v26) | 1)) >= (v26 << ((0) & 63))) {
                    v36 := v1["hello"]
                    v36 = 184
                    for fi37 := 0; fi37 < 20; fi37++ {
                        var v38 uint64 = (^(v5))
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                } else {
                    v39 := GenS4{F0: 813, F1: v27}
                    _ = v39
                }
                _ = v35
            }
        }
        _ = fi29
    }
    if v28 {
        if v28 {
            if ((42 >= -1) || v28) {
                if v28 {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        var v41 int64 = -1
                        _ = fi40
                        _ = v41
                    }
                    v42 := GenS5{F0: v27, F1: (^(v5))}
                    switch 647 {
                    case 1:
                        var v43 int64 = 1
                        _ = v43
                    case 628:
                        fmt.Println(v28, v3, v2)
                    case 32767:
                        v27 += uint16(v26)
                    default:
                        var v44 int64 = 1
                        _ = v44
                    }
                    _ = v42
                }
                var v45 int32 = -2147483648
                v46 := GenS6{F0: 42, F1: 128, F2: v28}
                _ = v45
                _ = v46
            } else {
                if len(v24) > 0 {
                    v24[0] = -2147483648
                }
                for fi47 := 0; fi47 < 14; fi47++ {
                    if v28 {
                        v5 = v5
                        if true {
                            if v28 {
                                v48 := GenS7{F0: genFn0()}
                                v4 = v4
                                _ = v48
                            }
                            v1["hello"] = 629
                        } else {
                            var v49 uint64 = uint64(v27)
                            _ = v49
                        }
                    } else {
                        if true {
                            if len(v24) > 0 {
                                v24[0] = fi47
                            }
                            v27 = 65535
                            if (!((v3 != 127) && ((0.1 * 3.14159) != (0.0 - 3.14159)))) {
                                v50 := GenS6{F0: -432, F1: (fi47 * 32768), F2: v28}
                                var v51 uint = 18446744073709551615
                                _ = v50
                                _ = v51
                            }
                        } else {
                            fmt.Println(v27)
                        }
                        for fi52 := 0; fi52 < 11; fi52++ {
                            fmt.Println(fi52)
                            _ = fi52
                        }
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    }
                    fmt.Println(v28, v27, v4)
                    if (genFn0() >= v27) {
                        v53 := GenS4{F0: (v5 & v5), F1: (^(v27))}
                        switch v4 {
                        case "foo bar":
                            if (uint16(v26) < v27) {
                                fmt.Println(v3, v5)
                                for fi54 := 0; fi54 < 13; fi54++ {
                                    switch (fi47 << (((^(v26))) & 63)) {
                                    case 32768:
                                        v27 -= (v27 >> (((^(v26))) & 63))
                                        v5 -= (v5 | (v5 << ((v26) & 63)))
                                    case -814:
                                        var v55 uint32 = 2
                                        switch ("" + "foo bar") {
                                        case "":
                                            var v56 float64 = (v2 + v2)
                                            _ = v56
                                        case "foo bar":
                                            switch fi54 {
                                            case 127:
                                                switch "a" {
                                                case "a":
                                                    fmt.Println(v3)
                                                case "foo bar":
                                                    if false {
                                                        fmt.Println(v3)
                                                        if v28 {
                                                            fmt.Println(v5)
                                                        } else {
                                                            if len(v24) > 0 {
                                                                v24[0] = fi54
                                                            }
                                                            var v57 uint16 = uint16(v26)
                                                            v27 = (v27 ^ uint16(v5))
                                                            _ = v57
                                                        }
                                                    }
                                                case "hello":
                                                    switch fi47 {
                                                    case 4294967295:
                                                        v24 = append(v24, (fi47 % (((fi47 - (fi47 - (fi54 / ((32768) | 1))))) | 1)))
                                                    case 1:
                                                        if len(v24) >= 2 {
                                                            v24 = v24[1:len(v24)]
                                                        }
                                                    default:
                                                        var v58 int = (fi47 + fi54)
                                                        _ = v58
                                                    }
                                                }
                                            case 1:
                                                var v59 int64 = -199
                                                _ = v59
                                            default:
                                                v26 = 64
                                            }
                                        case "hello":
                                            var v60 float64 = v2
                                            _ = v60
                                        }
                                        _ = v55
                                    case -869:
                                        v61 := GenS8{F0: (v26 - 0)}
                                        _ = v61
                                    default:
                                        var v62 int32 = -355
                                        _ = v62
                                    }
                                    _ = fi54
                                }
                                var v63 bool = false
                                _ = v63
                            } else {
                                if (951 > 751) {
                                    var v64 string = "a"
                                    _ = v64
                                } else {
                                    var v65 uint16 = v27
                                    v66 := v1[v4]
                                    _ = v65
                                    _ = v66
                                }
                            }
                        }
                        for fi67 := 0; fi67 < 1; fi67++ {
                            var v68 uint16 = (v27 % (((v27 - (v27 % (((v27 ^ genFn0())) | 1)))) | 1))
                            var v69 int8 = -27
                            fmt.Println(v28, fi67, v2)
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                        _ = v53
                    }
                    _ = fi47
                }
                fmt.Println(v2, v2)
            }
        } else {
            for fi70 := 0; fi70 < 12; fi70++ {
                for fi71 := 0; fi71 < 15; fi71++ {
                    var v72 uint64 = (v5 >> (((v26 % (((v26 / ((v26) | 1))) | 1))) & 63))
                    var v73 int64 = 1
                    v74 := GenS8{F0: 1}
                    _ = fi71
                    _ = v72
                    _ = v73
                    _ = v74
                }
                _ = fi70
            }
            v75 := GenS9{F0: v4, F1: 1, F2: 42}
            _ = v75
        }
        switch (v4 + v4) {
        case "":
            var v76 int8 = -1
            _ = v76
        default:
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
    }
    var v77 int = 4294967295
    fmt.Println(v27, v3, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
