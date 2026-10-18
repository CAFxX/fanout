package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 int8; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint8; F2 string }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint32 = 775
    v2 := make(map[uint32]int)
    var v3 string = ((("hello" + "a") + "a") + (("a" + "") + "hello"))
    var v4 uint = 4294967295
    v5 := []uint32{}
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 2; fi7++ {
        if ((v4 & 4294967295) <= v4) {
            v4 += 200
        }
        switch fi7 {
        case -9223372036854775808:
            v8 := GenS1{F0: -128, F1: v4, F2: 1}
            _ = v8
        default:
            switch ((v3 + ("foo bar" + "hello")) + (v3 + (v3 + (("\n\t" + v3) + ("foo bar" + v3))))) {
            case "":
                var v9 int = fi7
                _ = v9
            case "foo bar":
                v4 = v4
            default:
                var v10 bool = false
                _ = v10
            }
        }
        _ = len(v5)
        _ = fi7
    }
    v11 := GenS1{F0: 42, F1: v4, F2: -32768}
    v1 = (^(v1))
    v1 = v1
    v4 -= 42
    v6 = (v6 + 65535)
    for fi12 := 0; fi12 < 7; fi12++ {
        switch 127 {
        case 2147483647:
            var v13 string = v3
            _ = v13
        }
        var v14 uint16 = (v6 & v6)
        _ = fi12
        _ = v14
    }
    v5 = append(v5, (v1 + v1))
    v6 -= (^(v6))
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 = v4
        v6 = 2
        _ = fi15
    }
    var v16 int64 = int64(v6)
    var v17 float32 = 0.0
    v18 := v2[742]
    if ((-9223372036854775808 >= (-(v16))) || (339 > 23)) {
        switch 128 {
        case 42:
            for fi19 := 0; fi19 < 5; fi19++ {
                v17 -= ((v17 / ((0.0 - 0.0) + v17)) - float32(3.14159))
                v3 = (("foo bar" + v3) + v3)
                if ((v16 ^ int64(fi19)) == (v16 & v16)) {
                    fmt.Println(v16)
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if true {
                            switch v3 {
                            case "foo bar":
                                v21 := GenS2{F0: 1}
                                _ = v21
                            default:
                                if false {
                                    v16 = (-(v16))
                                }
                            }
                            _ = len(v2)
                        } else {
                            var v22 uint32 = uint32(v6)
                            var v23 uint = (v4 % (((v4 ^ (v4 ^ (v4 - v4)))) | 1))
                            fmt.Println(v4, v16)
                            _ = v22
                            _ = v23
                        }
                        var v24 bool = (false || true)
                        v25 := GenS3{F0: (v6 % ((v6) | 1)), F1: v16}
                        if (0 != (v16 % ((0) | 1))) {
                            fmt.Println(v1, v1, v4)
                            v4 += v4
                        }
                        _ = fi20
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if false {
                            if false {
                                var v27 uint16 = 1
                                _ = v27
                            } else {
                                if false {
                                    if ((v6 * v6) == 32767) {
                                        var v28 uint64 = 0
                                        if true {
                                            var v29 uint64 = 42
                                            var v30 uint16 = (v6 >> ((1) & 63))
                                            _ = v29
                                            _ = v30
                                        } else {
                                            v16 += 118
                                            v31 := GenS2{F0: -2147483648}
                                            _ = v31
                                        }
                                        _ = v28
                                    }
                                } else {
                                    var v32 int32 = -2147483648
                                    _ = v32
                                }
                                _ = len(v2)
                            }
                            fmt.Println(fi26)
                            var v33 int8 = 1
                            _ = v33
                        }
                        var v34 uint16 = uint16(v4)
                        v18 += 42
                        _ = fi26
                        _ = v34
                    }
                } else {
                    _ = len(v2)
                    if (42 != 1000) {
                        fmt.Println(v6, v18, v16)
                    } else {
                        v35 := v2[0]
                        _ = len(v5)
                        _ = v35
                    }
                }
                _ = fi19
            }
        case 44:
            fmt.Println(v4)
        case -392:
            if true {
                delete(v2, 2147483647)
                switch "" {
                case "\n\t":
                    v16 += v16
                case "hello":
                    if true {
                        fmt.Println(v17, v6)
                        v4 += (v4 * v4)
                        if false {
                            for fi36 := 0; fi36 < 1; fi36++ {
                                v37 := GenS4{F0: 2147483647, F1: 42, F2: (v6 << (((v4 >> (((v4 & v4)) & 63))) & 63))}
                                if ((float64(3.14159) * float64(v17)) <= 3.14159) {
                                    delete(v2, 393)
                                } else {
                                    if (255 <= (v6 * 914)) {
                                        var v38 int8 = -96
                                        var v39 uint = v4
                                        var v40 uint16 = v6
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    } else {
                                        var v41 float64 = 0.0
                                        if (((float64(v17) / 1.5) > v41) && (!true)) {
                                            switch (fi36 ^ fi36) {
                                            case 718:
                                                var v42 int8 = 82
                                                _ = v42
                                            case -130:
                                                v1 = v1
                                            case 4294967295:
                                                v43 := GenS5{F0: fi36}
                                                _ = v43
                                            }
                                            if true {
                                                v44 := GenS6{F0: 912, F1: 255, F2: v3}
                                                var v45 string = ("\n\t" + "\n\t")
                                                if false {
                                                    v4 = 387
                                                    var v46 int = v18
                                                    delete(v2, 2147483647)
                                                    _ = v46
                                                }
                                                _ = v44
                                                _ = v45
                                            }
                                            var v47 int16 = 658
                                            _ = v47
                                        } else {
                                            var v48 uint16 = v6
                                            var v49 bool = ((v4 * v4) <= (v4 | 579))
                                            _ = v48
                                            _ = v49
                                        }
                                        var v50 float64 = v41
                                        _ = v41
                                        _ = v50
                                    }
                                    var v51 uint16 = 326
                                    if true {
                                        var v52 int8 = 97
                                        v2[(v1 / (((v1 & v1)) | 1))] = 4294967295
                                        _ = v52
                                    }
                                    _ = v51
                                }
                                v53 := GenS1{F0: 0, F1: 433, F2: -32768}
                                _ = fi36
                                _ = v37
                                _ = v53
                            }
                            if false {
                                v3 = v3
                                if false {
                                    var v54 uint64 = 2
                                    _ = v54
                                } else {
                                    switch v18 {
                                    case -1:
                                        v16 += v16
                                    }
                                    if (0 >= -2147483648) {
                                        var v55 int8 = -41
                                        if (true && ((v16 ^ 507) == -9223372036854775808)) {
                                            v1 = v1
                                            switch int(v16) {
                                            case 9223372036854775807:
                                                if len(v5) > 0 {
                                                    v5[0] = v1
                                                }
                                            case 128:
                                                var v56 int16 = 12
                                                _ = v56
                                            }
                                            var v57 uint64 = 850
                                            _ = v57
                                        } else {
                                            v58 := GenS7{F0: 2, F1: 42}
                                            delete(v2, (v1 & v1))
                                            _ = v58
                                        }
                                        _ = v55
                                    }
                                }
                                fmt.Println(v1, v1, v1)
                            }
                        } else {
                            for fi59 := 0; fi59 < 15; fi59++ {
                                if len(v5) > 0 {
                                    v5[0] = (^(v1))
                                }
                                if true {
                                    switch v3 {
                                    case "\n\t":
                                        for fi60 := 0; fi60 < 8; fi60++ {
                                            switch "a" {
                                            case "":
                                                var v61 float64 = 3.14159
                                                _ = v61
                                            case "a":
                                                v6 -= (v6 >> (((v4 + v4)) & 63))
                                            case "foo bar":
                                                switch v3 {
                                                case "foo bar":
                                                    v62 := GenS8{F0: 4294967295}
                                                    _ = v62
                                                default:
                                                    _ = len(v5)
                                                }
                                                var v63 uint = (v4 + (v4 << (((v4 / ((uint(v1)) | 1))) & 63)))
                                                _ = v63
                                            }
                                            v64 := GenS9{F0: ("hello" != v3), F1: 255}
                                            _ = fi60
                                            _ = v64
                                        }
                                    case "":
                                        var v65 bool = (!((fi59 / ((0) | 1)) >= v18))
                                        _ = v65
                                    }
                                } else {
                                    v66 := GenS8{F0: 4294967295}
                                    _ = v66
                                }
                                v67 := GenS4{F0: (fi59 - -589), F1: int8(v4), F2: (^(v6))}
                                _ = fi59
                                _ = v67
                            }
                            var v68 uint8 = 1
                            _ = v68
                        }
                    } else {
                        v6 += 65535
                        var v69 uint32 = 813
                        _ = v69
                    }
                case "a":
                    v6 -= 776
                default:
                    v4 += v4
                }
            }
        default:
            for fi70 := 0; fi70 < 16; fi70++ {
                var v71 float64 = 1.5
                _ = fi70
                _ = v71
            }
        }
    } else {
        var v72 uint8 = 105
        var v73 float64 = 3.14159
        _ = v72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
