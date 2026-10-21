package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 int32; F1 int; F2 int }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int64; F2 int8 }

func main() {
    var v1 string = "hello"
    v2 := []string{"foo bar", ("foo bar" + ""), (v1 + "")}
    v3 := []uint32{2, 4294967295}
    v1 = ((("hello" + v1) + (v1 + v1)) + v1)
    var v4 uint64 = 42
    switch 924 {
    case -1:
        var v5 uint16 = 210
        _ = v5
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 697:
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 string = v1
            _ = fi6
            _ = v7
        }
    }
    v4 = 1
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            v9 := GenS1{F0: 68, F1: 128}
            _ = v9
        } else {
            var v10 bool = (18446744073709551615 >= v4)
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = v10
        }
        _ = fi8
    }
    if (!(v1 > "")) {
        v11 := GenS1{F0: 182, F1: 0}
        var v12 uint = 372
        _ = v11
        _ = v12
    }
    var v13 int8 = 1
    if (true || (!((!true) || (!true)))) {
        fmt.Println(v13)
        if len(v3) > 0 {
            v3[0] = 781
        }
        var v14 uint16 = 824
        _ = v14
    } else {
        var v15 int32 = 836
        v16 := GenS2{F0: 1, F1: 32767}
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(v1, v15)
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    if (float32(float64(0.1)) <= ((3.14159 / 1.0) - -2.25)) {
        v13 += v13
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v18 := GenS2{F0: (^(v13)), F1: 42}
    for fi19 := 0; fi19 < 1; fi19++ {
        v13 = v13
        _ = fi19
    }
    var v20 int16 = -284
    var v21 int8 = (^(v13))
    v4 += (v4 * (v4 / ((v4) | 1)))
    for fi22 := 0; fi22 < 13; fi22++ {
        if true {
            var v23 uint32 = 42
            _ = v23
        } else {
            if true {
                if (-382 != fi22) {
                    v20 += -450
                    for fi24 := 0; fi24 < 9; fi24++ {
                        if (v1 != ("a" + "")) {
                            if len(v3) > 0 {
                                v3[0] = 287
                            }
                            if (-17 < -955) {
                                v25 := GenS3{F0: 1, F1: -1, F2: 739}
                                _ = v25
                            } else {
                                v26 := GenS3{F0: 42, F1: 1, F2: 413}
                                _ = v26
                            }
                        }
                        v3 = append(v3, 42)
                        _ = fi24
                    }
                    if false {
                        var v27 int32 = 0
                        _ = v27
                    } else {
                        var v28 uint64 = (^(v4))
                        var v29 int16 = (^(v20))
                        for fi30 := 0; fi30 < 8; fi30++ {
                            var v31 int8 = 127
                            _ = fi30
                            _ = v31
                        }
                        _ = v28
                        _ = v29
                    }
                }
                if (int64(v13) == 9223372036854775807) {
                    if (!false) {
                        v20 -= v20
                        for fi32 := 0; fi32 < 14; fi32++ {
                            v33 := GenS3{F0: (^(v4)), F1: 623, F2: 1}
                            if false {
                                v4 -= 18446744073709551615
                                if len(v3) > 0 {
                                    v3[0] = uint32(v4)
                                }
                                _ = len(v2)
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                            } else {
                                switch int(v20) {
                                case 42:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v34 int16 = (v20 / ((v20) | 1))
                                    _ = v34
                                }
                                v20 = (v20 - v20)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            var v35 uint16 = 0
                            var v36 uint = 639
                            _ = fi32
                            _ = v33
                            _ = v35
                            _ = v36
                        }
                        v20 += 77
                    } else {
                        v13 += v13
                        if (65535 >= 2) {
                            fmt.Println(fi22, v20)
                            v37 := GenS4{F0: 1, F1: (-(fi22)), F2: (^(fi22))}
                            v38 := GenS5{F0: -9223372036854775808, F1: 42}
                            _ = v37
                            _ = v38
                        } else {
                            var v39 float32 = ((3.14159 - float32((((3.14159 - (float64(0.1) - 0.1)) / (0.1 / (((((0.0 * 0.1) - 0.0) * (0.0 / 1.0)) - 1.5) / 3.14159))) / -2.25))) - 0.1)
                            _ = v39
                        }
                    }
                    v4 = v4
                } else {
                    if true {
                        var v40 int = (fi22 << ((2) & 63))
                        _ = v40
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v4 = 1
                        switch "" {
                        case "hello":
                            switch v1 {
                            case "":
                                v13 = v13
                            case "a":
                                for fi41 := 0; fi41 < 6; fi41++ {
                                    v13 += v13
                                    v21 += (v13 / ((-26) | 1))
                                    if (true && false) {
                                        fmt.Println(fi22, v20)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v42 float32 = 1e10
                                    var v43 int8 = 100
                                    v44 := GenS3{F0: (v4 % ((v4) | 1)), F1: 42, F2: 255}
                                    v13 = (v21 + (v21 * (v13 * (v43 % ((v13) | 1)))))
                                    v45 := GenS2{F0: 100, F1: (v20 / (((^(v20))) | 1))}
                                    switch fi41 {
                                    case 32767:
                                        v1 = (v1 + (v1 + "\n\t"))
                                    }
                                    if false {
                                        var v46 bool = (v42 < ((v42 + ((-2.25 + v42) / float32(float64(v42)))) + (v42 - 3.14159)))
                                        var v47 int16 = v20
                                        _ = v46
                                        _ = v47
                                    } else {
                                        if ((fi41 >> ((871) & 63)) > 1) {
                                            fmt.Println(v1, v13)
                                        }
                                        v42 += v42
                                    }
                                    var v48 float32 = -2.25
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                    _ = v48
                                }
                            }
                        default:
                            var v49 string = v1
                            _ = v49
                        }
                    }
                    var v50 uint8 = 214
                    for fi51 := 0; fi51 < 12; fi51++ {
                        switch fi51 {
                        case 32767:
                            v52 := GenS6{F0: int16(v50)}
                            _ = v52
                        }
                        if (!(!(0 <= v50))) {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            fmt.Println(fi22)
                            var v53 int8 = (v13 ^ (^(v13)))
                            _ = v53
                        } else {
                            v20 += v20
                        }
                        v54 := GenS2{F0: v21, F1: -1}
                        _ = fi51
                        _ = v54
                    }
                    _ = v50
                }
                if (0.1 < float64((-2.25 - -2.25))) {
                    v55 := GenS7{F0: 128, F1: -642, F2: (v21 * -1)}
                    for fi56 := 0; fi56 < 15; fi56++ {
                        fmt.Println(v21, v21)
                        _ = fi56
                    }
                    switch (fi22 & (^(fi22))) {
                    case 128:
                        v13 += (v21 | v21)
                    case 1:
                        for fi57 := 0; fi57 < 13; fi57++ {
                            v2 = append(v2, "\n\t")
                            var v58 uint16 = 1
                            _ = fi57
                            _ = v58
                        }
                    default:
                        var v59 string = v1
                        _ = v59
                    }
                    _ = v55
                }
            } else {
                v1 = v1
                v20 -= 196
            }
            for fi60 := 0; fi60 < 2; fi60++ {
                v1 = v1
                var v61 int16 = v20
                fmt.Println(v4, fi22, v20)
                _ = fi60
                _ = v61
            }
            var v62 int64 = -9223372036854775808
            _ = v62
        }
        for fi63 := 0; fi63 < 5; fi63++ {
            var v64 int64 = 1
            v2 = append(v2, v1)
            v13 = 100
            _ = fi63
            _ = v64
        }
        v1 = (("foo bar" + v1) + v1)
        _ = fi22
    }
    switch -9223372036854775808 {
    case 2147483647:
        v4 += v4
    }
    var v65 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v65
}
