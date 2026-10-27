package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 uint16; F1 uint8; F2 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 int64 }

func main() {
    v1 := []int{1, 32767}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 uint = 0
    v4 -= v4
    v4 = (^(v4))
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 32767, F1: (v4 % ((230) | 1))}
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    var v7 float64 = 1e10
    v8 := GenS2{F0: 127, F1: 4294967295, F2: 0}
    v9 := GenS1{F0: uint16(v4), F1: v4}
    var v10 int16 = -1
    if false {
        if false {
            for fi11 := 0; fi11 < 2; fi11++ {
                v10 = v10
                var v12 int16 = (v10 % ((1) | 1))
                _ = fi11
                _ = v12
            }
            if false {
                fmt.Println(v4)
                if true {
                    if ((-1 < -1) || (!("foo bar" != "foo bar"))) {
                        v1 = append(v1, 158)
                        v13 := GenS3{F0: 106, F1: 213, F2: (v4 & v4)}
                        if true {
                            if len(v1) > 0 {
                                v1[0] = -2147483648
                            }
                            for fi14 := 0; fi14 < 1; fi14++ {
                                var v15 uint64 = 0
                                v4 = (v4 / ((255) | 1))
                                _ = fi14
                                _ = v15
                            }
                            var v16 int64 = -478
                            _ = v16
                        }
                        _ = v13
                    } else {
                        fmt.Println(v7, v7, v10)
                        if ((((0.1 * ((float64(float32(v7)) / ((float64((-2.25 - (3.14159 * (float32(-2.25) + 0.0)))) - (v7 - 0.0)) + v7)) + float64(1.5))) - 1e10) <= (((v7 / (v7 + -2.25)) - (-2.25 - v7)) + v7)) || (((1 > 952) && false) && false)) {
                            if true {
                                v17 := v3[454]
                                if true {
                                    v18 := GenS4{F0: 140}
                                    for fi19 := 0; fi19 < 19; fi19++ {
                                        var v20 int64 = -1
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = v18
                                }
                                var v21 uint32 = 52
                                _ = v17
                                _ = v21
                            }
                            var v22 int8 = 127
                            v7 = 3.14159
                            _ = v22
                        }
                        var v23 int16 = (v10 % (((-(v10))) | 1))
                        _ = v23
                    }
                }
            } else {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := GenS4{F0: 128}
                    switch fi24 {
                    case -9223372036854775808:
                        if (!true) {
                            if ((fi24 >> (((v4 | v4)) & 63)) == (fi24 / (((fi24 * 0)) | 1))) {
                                v7 += 0.0
                            } else {
                                v10 = v10
                            }
                        }
                    case 42:
                        var v26 string = ""
                        _ = v26
                    case 128:
                        v4 += (v4 - (v4 << (((v4 ^ v4)) & 63)))
                    default:
                        v7 -= (v7 - 1e10)
                    }
                    _ = fi24
                    _ = v25
                }
            }
            if true {
                if true {
                    var v27 bool = (false && true)
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v1 = append(v1, (-(fi28)))
                        _ = fi28
                    }
                    v29 := GenS3{F0: 0, F1: 255, F2: v4}
                    fmt.Println(v7)
                    _ = v29
                }
                v7 += -2.25
            } else {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v30 int16 = 42
                v10 -= v30
                _ = v30
            }
        } else {
            v31 := GenS5{F0: -32768, F1: 1, F2: (v4 >> (((v4 / ((v4) | 1))) & 63))}
            if true {
                v10 += (v10 / (((v10 ^ v10)) | 1))
                v7 -= v7
            } else {
                var v32 uint64 = 202
                delete(v2, "a")
                _ = v32
            }
            v4 += (^(v4))
            _ = v31
        }
        v33 := GenS2{F0: 157, F1: 228, F2: 42}
        _ = v33
    } else {
        v10 = (v10 + v10)
    }
    if true {
        var v34 int16 = 912
        var v35 int8 = int8(v10)
        _ = v34
        _ = v35
    } else {
        v10 = (v10 % ((v10) | 1))
        var v36 uint32 = 42
        _ = v36
    }
    switch -507 {
    case -1:
        v4 -= v4
    case 485:
        for fi37 := 0; fi37 < 11; fi37++ {
            if (119 != -110) {
                switch "" {
                case "foo bar":
                    v38 := GenS6{F0: 85, F1: 2}
                    _ = v38
                default:
                    fmt.Println(v7, v10)
                }
                v39 := GenS7{F0: -1, F1: 2, F2: (^(v4))}
                _ = v39
            } else {
                v4 += 0
                if false {
                    delete(v2, "")
                    v40 := GenS8{F0: ("foo bar" + "foo bar"), F1: int16(v4)}
                    v41 := v2["\n\t"]
                    _ = v40
                    _ = v41
                } else {
                    switch "" {
                    case "\n\t":
                        v42 := GenS2{F0: 194, F1: 317, F2: 264}
                        _ = v42
                    default:
                        switch ("foo bar" + ("\n\t" + "\n\t")) {
                        case "hello":
                            v7 = v7
                        case "":
                            v4 -= v4
                        case "\n\t":
                            fmt.Println(v4, v7, v10)
                        }
                    }
                }
            }
            delete(v2, ("hello" + ("hello" + ("hello" + "a"))))
            _ = fi37
        }
    default:
        for fi43 := 0; fi43 < 11; fi43++ {
            if true {
                v44 := v2["a"]
                for fi45 := 0; fi45 < 9; fi45++ {
                    v46 := GenS7{F0: 42, F1: 1, F2: 4294967295}
                    v10 = v10
                    _ = fi45
                    _ = v46
                }
                _ = v44
            } else {
                v4 = 4294967295
                if true {
                    var v47 uint32 = 2147483648
                    if true {
                        if (v4 >= (^(v4))) {
                            v10 = v10
                        } else {
                            _ = len(v3)
                        }
                        for fi48 := 0; fi48 < 4; fi48++ {
                            var v49 uint = (v4 % ((42) | 1))
                            if (!false) {
                                v4 = 18446744073709551615
                                fmt.Println(fi43, fi48, fi43)
                            } else {
                                v50 := GenS9{F0: -1, F1: true}
                                var v51 uint64 = 700
                                v49 += 18446744073709551615
                                _ = v50
                                _ = v51
                            }
                            _ = fi48
                            _ = v49
                        }
                        if ((128 >= 153) && false) {
                            switch int(v10) {
                            case 127:
                                if true {
                                    if (!(uint(v47) > 42)) {
                                        switch (-(fi43)) {
                                        case 1:
                                            switch 1 {
                                            case 42:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        delete(v3, (v47 % (((v47 + (v47 ^ 42))) | 1)))
                                    }
                                }
                            case -959:
                                if true {
                                    var v52 int64 = 9223372036854775807
                                    _ = v52
                                } else {
                                    for fi53 := 0; fi53 < 18; fi53++ {
                                        if (18446744073709551615 < 18446744073709551615) {
                                            switch fi53 {
                                            case -311:
                                                v7 = float64(0.0)
                                            case 42:
                                                fmt.Println(v4)
                                            default:
                                                v7 += 0.0
                                            }
                                        }
                                        var v54 uint8 = 102
                                        _ = fi53
                                        _ = v54
                                    }
                                    v10 -= v10
                                    v4 = (v4 % (((v4 - 42)) | 1))
                                }
                            case 128:
                                var v55 int16 = (v10 >> (((v4 % ((v4) | 1))) & 63))
                                _ = v55
                            }
                            var v56 float32 = 0.1
                            _ = v56
                        } else {
                            for fi57 := 0; fi57 < 11; fi57++ {
                                v10 += v10
                                if len(v1) > 0 {
                                    v1[0] = fi57
                                }
                                _ = fi57
                            }
                            var v58 uint16 = 0
                            _ = v58
                        }
                    } else {
                        v59 := GenS10{F0: 255, F1: 1, F2: false}
                        _ = v59
                    }
                    _ = v47
                } else {
                    if ((((-(v10)) <= 1) && false) || false) {
                        if true {
                            v60 := GenS11{F0: fi43}
                            switch (fi43 << ((v4) & 63)) {
                            case -9223372036854775808:
                                for fi61 := 0; fi61 < 6; fi61++ {
                                    v62 := GenS2{F0: 201, F1: 2, F2: -1}
                                    switch (fi61 | (fi61 * (fi43 & int(v10)))) {
                                    case -1:
                                        v10 -= -32768
                                    }
                                    _ = fi61
                                    _ = v62
                                }
                            default:
                                if true {
                                    fmt.Println(fi43)
                                    if (-118 > -1) {
                                        v10 = (v10 >> (((v4 - (^(v4)))) & 63))
                                        var v63 int8 = 100
                                        _ = v63
                                    }
                                    v4 = (v4 + (v4 >> (((^(v4))) & 63)))
                                }
                            }
                            _ = len(v2)
                            _ = v60
                        } else {
                            v4 = (v4 & (v4 - v4))
                            fmt.Println(v10, fi43)
                        }
                        var v64 string = "a"
                        v10 = int16(fi43)
                        _ = v64
                    }
                }
            }
            v7 = (v7 * v7)
            _ = fi43
        }
    }
    for fi65 := 0; fi65 < 11; fi65++ {
        v66 := GenS10{F0: 249, F1: 127, F2: (false && false)}
        _ = fi65
        _ = v66
    }
    for fi67 := 0; fi67 < 2; fi67++ {
        var v68 uint16 = 455
        for fi69 := 0; fi69 < 20; fi69++ {
            v4 = 1
            if (false && true) {
                var v70 uint32 = uint32(v68)
                v71 := GenS8{F0: ("hello" + ("" + "a")), F1: v10}
                var v72 uint64 = 1
                var v73 int64 = -1
                var v74 bool = false
                var v75 float64 = v7
                _ = v70
                _ = v71
                _ = v72
                _ = v73
                _ = v74
                _ = v75
            } else {
                var v76 int32 = -465
                _ = v76
            }
            v68 += v68
            var v77 uint16 = (v68 / (((v68 >> (((^(v4))) & 63))) | 1))
            _ = fi69
            _ = v77
        }
        if true {
            v2[((("\n\t" + ("\n\t" + (("hello" + "") + "hello"))) + "") + (((("foo bar" + "hello") + "a") + "\n\t") + "foo bar"))] = (^(fi67))
            v1 = append(v1, (fi67 | fi67))
            v78 := v3[65535]
            _ = v78
        } else {
            if (0 > 1) {
                _ = len(v2)
            }
        }
        _ = fi67
        _ = v68
    }
    v79 := v2["foo bar"]
    v4 = 65535
    v79 += (v79 * -994)
    v80 := GenS12{F0: -811, F1: 42}
    var v81 int64 = 738
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v79
    _ = v8
    _ = v80
    _ = v81
    _ = v9
}
