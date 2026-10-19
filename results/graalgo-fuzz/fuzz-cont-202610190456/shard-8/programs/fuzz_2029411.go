package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 bool }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 int8 }
type GenS8 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint16 = 255
    var v2 int64 = 1
    var v3 string = "hello"
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := GenS1{F0: v3}
            v7 := GenS1{F0: v3}
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(v1, v1)
        switch ("foo bar" + v3) {
        case "":
            switch v3 {
            case "\n\t":
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := []int{359, fi4, 127, (fi8 | 0)}
                    _ = fi8
                    _ = v9
                }
            case "":
                v10 := GenS2{F0: 2, F1: 2147483648}
                _ = v10
            case "foo bar":
                var v11 uint32 = 561
                _ = v11
            default:
                var v12 bool = (!false)
                _ = v12
            }
        case "hello":
            for fi13 := 0; fi13 < 16; fi13++ {
                if false {
                    v3 = "\n\t"
                    fmt.Println(fi13, v2, v3)
                } else {
                    var v14 float32 = 1e10
                    _ = v14
                }
                _ = fi13
            }
        }
        _ = fi4
    }
    v15 := GenS3{F0: 2, F1: v2, F2: true}
    v16 := []int{-314}
    fmt.Println(v1, v2)
    if (0 <= 481) {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        v2 = (v2 & v2)
        v2 -= 42
    } else {
        var v17 int16 = -571
        _ = v17
    }
    var v18 int32 = -333
    var v19 int8 = 42
    if (true && true) {
        v18 -= (v18 + int32(v1))
        for fi20 := 0; fi20 < 7; fi20++ {
            v3 = v3
            _ = fi20
        }
        var v21 uint64 = 1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                for fi23 := 0; fi23 < 19; fi23++ {
                    v24 := GenS2{F0: 2, F1: 982}
                    _ = fi23
                    _ = v24
                }
                if false {
                    fmt.Println(v18, v3)
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v16 = append(v16, fi25)
                        var v26 string = "\n\t"
                        var v27 uint16 = (v1 & (v1 % (((v1 - v1)) | 1)))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    if (127 != 1) {
                        v28 := GenS4{F0: 74, F1: 2147483647, F2: (100 <= v19)}
                        var v29 float32 = 1.5
                        _ = len(v16)
                        _ = v28
                        _ = v29
                    } else {
                        if (!true) {
                            v19 += v19
                        } else {
                            switch fi22 {
                            case 933:
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v1 += v1
                                    switch (fi30 / (((^(fi30))) | 1)) {
                                    case 147:
                                        if true {
                                            fmt.Println(fi30, v18, fi30)
                                            v18 -= (v18 | v18)
                                            v31 := GenS2{F0: 0, F1: 1}
                                            v19 = (v19 | (v19 << ((0) & 63)))
                                            _ = v31
                                        }
                                    default:
                                        v1 = (v1 / ((1) | 1))
                                    }
                                    var v32 uint8 = 128
                                    _ = fi30
                                    _ = v32
                                }
                            case 0:
                                var v33 int8 = (v19 * v19)
                                _ = v33
                            case 128:
                                var v34 uint16 = v1
                                _ = v34
                            }
                            var v35 int = fi22
                            for fi36 := 0; fi36 < 3; fi36++ {
                                fmt.Println(v2, fi36)
                                v3 = (v3 + (("a" + (("a" + "a") + v3)) + ("foo bar" + "hello")))
                                _ = fi36
                            }
                            _ = v35
                        }
                        v1 += v1
                        v37 := GenS2{F0: 4294967295, F1: 2147483648}
                        _ = v37
                    }
                } else {
                    v16 = append(v16, 32768)
                }
            } else {
                for fi38 := 0; fi38 < 9; fi38++ {
                    fmt.Println(v18)
                    _ = fi38
                }
                switch -9223372036854775808 {
                case 1:
                    v19 -= v19
                case 42:
                    v19 += (v19 << ((18446744073709551615) & 63))
                case 691:
                    var v39 uint = 134
                    _ = v39
                default:
                    switch v3 {
                    case "a":
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    case "hello":
                        var v40 float64 = (float64(0.0) - 1e10)
                        _ = v40
                    case "":
                        var v41 int32 = (-(v18))
                        _ = v41
                    }
                }
            }
            fmt.Println(v19, v1)
            _ = fi22
        }
        if (!true) {
            if false {
                if len(v16) > 0 {
                    v16[0] = -367
                }
                if true {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        v19 = -128
                        v43 := GenS4{F0: (fi42 - 32768), F1: v18, F2: (255 < 1)}
                        _ = fi42
                        _ = v43
                    }
                } else {
                    switch 482 {
                    case -9223372036854775808:
                        for fi44 := 0; fi44 < 14; fi44++ {
                            switch fi44 {
                            case 127:
                                if true {
                                    var v45 uint16 = 112
                                    var v46 uint8 = 210
                                    _ = v45
                                    _ = v46
                                }
                            }
                            _ = fi44
                        }
                    case 128:
                        v47 := GenS1{F0: (v3 + v3)}
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 8; fi48++ {
                        if len(v16) > 0 {
                            v16[0] = (fi48 ^ fi48)
                        }
                        for fi49 := 0; fi49 < 15; fi49++ {
                            fmt.Println(fi49, v3)
                            _ = len(v16)
                            _ = fi49
                        }
                        _ = fi48
                    }
                    switch (((v3 + "hello") + ("foo bar" + "\n\t")) + v3) {
                    case "a":
                        if true {
                            v50 := GenS1{F0: "\n\t"}
                            if (127 == 9223372036854775807) {
                                var v51 uint8 = 12
                                if len(v16) > 0 {
                                    v16[0] = -9223372036854775808
                                }
                                var v52 int64 = v2
                                _ = v51
                                _ = v52
                            }
                            _ = v50
                        } else {
                            v53 := GenS5{F0: 2147483647, F1: v1}
                            var v54 int = 128
                            v2 += (v2 + (v2 & (v2 ^ (v2 % ((v2) | 1)))))
                            _ = v53
                            _ = v54
                        }
                    default:
                        if false {
                            v19 -= v19
                            if len(v16) > 0 {
                                v16[0] = 127
                            }
                        }
                    }
                }
            } else {
                var v55 uint32 = 169
                _ = v55
            }
        } else {
            _ = len(v16)
        }
        var v56 string = "foo bar"
        _ = v56
    }
    var v57 int32 = (v18 * (v18 * -693))
    for fi58 := 0; fi58 < 6; fi58++ {
        if len(v16) > 0 {
            v16[0] = 1
        }
        _ = fi58
    }
    for fi59 := 0; fi59 < 3; fi59++ {
        v1 += 65535
        var v60 uint32 = 1
        v2 += v2
        _ = fi59
        _ = v60
    }
    v18 = (v57 << ((65535) & 63))
    if len(v16) > 0 {
        v16[0] = 32767
    }
    v1 -= v1
    v61 := GenS6{F0: 42, F1: 59}
    _ = len(v16)
    if (56 == -32768) {
        v62 := GenS7{F0: 127, F1: 127, F2: v19}
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        for fi63 := 0; fi63 < 3; fi63++ {
            v57 += (v57 >> ((1) & 63))
            for fi64 := 0; fi64 < 15; fi64++ {
                v3 = "hello"
                switch ("hello" + "hello") {
                case "\n\t":
                    if (4294967295 > 526) {
                        _ = len(v16)
                        v18 += (v18 ^ -342)
                        var v65 int8 = (^(v19))
                        _ = v65
                    } else {
                        v1 = 380
                        if (65535 >= 1) {
                            v19 += -1
                            var v66 int64 = (v2 % ((-399) | 1))
                            _ = v66
                        } else {
                            v1 += (^(v1))
                            var v67 int32 = 0
                            _ = v67
                        }
                        var v68 uint8 = 128
                        _ = v68
                    }
                case "":
                    v69 := GenS2{F0: 78, F1: 2}
                    _ = v69
                }
                var v70 string = v3
                if false {
                    v71 := GenS6{F0: 1, F1: 2}
                    _ = v71
                } else {
                    var v72 uint = 4294967295
                    _ = v72
                }
                fmt.Println(fi64, v1, v2)
                var v73 uint16 = v1
                var v74 int8 = (v19 | v19)
                var v75 string = ("" + "hello")
                _ = fi64
                _ = v70
                _ = v73
                _ = v74
                _ = v75
            }
            _ = fi63
        }
        _ = v62
    } else {
        v2 += v2
    }
    for fi76 := 0; fi76 < 10; fi76++ {
        v77 := GenS8{F0: 69, F1: 17}
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v57
    _ = v61
}
