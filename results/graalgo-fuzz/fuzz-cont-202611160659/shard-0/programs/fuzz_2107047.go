package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 int32; F1 int32; F2 int32 }
type GenS10 struct { F0 uint64; F1 bool; F2 bool }
type GenS11 struct { F0 uint16; F1 uint16 }
type GenS12 struct { F0 int16; F1 int16; F2 int8 }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 int16; F1 uint64; F2 int }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    if false {
        var v6 int = p0
        _ = v6
    }
    _ = p0
    return 291
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    v3 := []uint32{32, 2147483647}
    var v4 int = -9223372036854775808
    v5 := make(map[string]uint32)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v2 = 1
            var v9 string = "hello"
            if (3.14159 >= 0.0) {
                v2 += (v2 * 18446744073709551615)
                if false {
                    v2 -= (v2 << (((^(v2))) & 63))
                } else {
                    v10 := GenS1{F0: 1, F1: 0, F2: -171}
                    var v11 bool = true
                    _ = v10
                    _ = v11
                }
                v2 = (^(v2))
            } else {
                v12 := GenS2{F0: 1, F1: (fi7 ^ (fi7 | (v4 | fi8)))}
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v4 -= (v4 | (v4 | -2147483648))
    fmt.Println(v4)
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v4, fi13)
        _ = fi13
    }
    var v14 int64 = -9223372036854775808
    var v15 bool = true
    if (((v4 <= 127) && (v14 <= (v14 >> (((v2 >> ((v2) & 63))) & 63)))) && (v15 && (!v15))) {
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(v14, fi16, v14)
            if v15 {
                v17 := GenS1{F0: 1, F1: 42, F2: 174}
                if (2147483647 <= -2147483648) {
                    v15 = (v15 || (!v15))
                    v18 := GenS2{F0: -32768, F1: fi16}
                    _ = v18
                }
                switch fi16 {
                case -637:
                    v2 -= (^(v2))
                default:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if (!v15) {
                            if (32767 <= 32767) {
                                v20 := GenS3{F0: 18446744073709551615, F1: 0}
                                v21 := GenS4{F0: 2147483648}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 bool = v15
                                var v23 int64 = (v14 % ((int64(fi16)) | 1))
                                switch (("a" + (("\n\t" + "") + "\n\t")) + "") {
                                case "":
                                    fmt.Println(fi19)
                                default:
                                    v14 = int64(fi19)
                                }
                                _ = v22
                                _ = v23
                            }
                            fmt.Println(v15)
                        } else {
                            if (100 != 1) {
                                var v24 int64 = v14
                                if (!true) {
                                    fmt.Println(v24)
                                } else {
                                    if (1 >= uint16(v2)) {
                                        if (-32768 != 1000) {
                                            v2 -= v2
                                        }
                                        switch "a" {
                                        case "":
                                            v24 -= (-(v24))
                                        case "hello":
                                            if (!v15) {
                                                v25 := GenS5{F0: v14, F1: 0, F2: 65535}
                                                _ = v25
                                            }
                                        }
                                        v24 -= (v14 & 518)
                                        fmt.Println(fi19, fi16, fi16)
                                        switch fi16 {
                                        case 137:
                                            v26 := GenS3{F0: 18446744073709551615, F1: 65535}
                                            _ = v26
                                        case 32767:
                                            var v27 float64 = 1e10
                                            _ = v27
                                        case -908:
                                            fmt.Println(v2)
                                        default:
                                            var v28 uint64 = 1
                                            _ = v28
                                        }
                                        v29 := GenS6{F0: v15, F1: v15}
                                        switch "" {
                                        case "foo bar":
                                            var v30 int64 = (v24 | v24)
                                            _ = v30
                                        case "":
                                            v24 += v14
                                        }
                                        _ = v29
                                    } else {
                                        var v31 int32 = int32(v24)
                                        if false {
                                            v2 = v2
                                            v15 = v15
                                            var v32 uint32 = 249
                                            var v33 int8 = 1
                                            switch -2147483648 {
                                            case 42:
                                                var v34 uint64 = 42
                                                _ = v34
                                            case 44:
                                                if (!(v2 <= 18446744073709551615)) {
                                                    v32 = v32
                                                    var v35 uint16 = 65535
                                                    v4 += int(v24)
                                                    _ = v35
                                                }
                                            default:
                                                v36 := GenS7{F0: 150, F1: v31}
                                                _ = v36
                                            }
                                            _ = v32
                                            _ = v33
                                        }
                                        var v37 int8 = 1
                                        v24 = v24
                                        if (!(846 <= (v2 / ((0) | 1)))) {
                                            if v15 {
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                                v38 := GenS8{F0: v2, F1: (v14 % ((v24) | 1))}
                                                v39 := GenS9{F0: (-(v31)), F1: (-(v31)), F2: (^(v31))}
                                                v37 -= (v37 * v37)
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v2 = 1
                                                v31 += (v31 / (((v31 / ((v31) | 1))) | 1))
                                                _ = len(v3)
                                            }
                                            v24 -= v24
                                        } else {
                                            v31 = v31
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = v31
                                        _ = v37
                                    }
                                    var v40 uint64 = 0
                                    _ = v40
                                }
                                v4 = (fi16 ^ v4)
                                _ = v24
                            }
                            v1[2147483647] = (((v15 || (42 != 18446744073709551615)) && v15) && ((255 == 56) && (2 < 956)))
                        }
                        var v41 int16 = 1000
                        _ = fi19
                        _ = v41
                    }
                }
                _ = v17
            }
            _ = len(v3)
            _ = fi16
        }
        if (0 != 65535) {
            v5["foo bar"] = 2
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            var v43 int8 = -1
            v43 = (v43 % ((v43) | 1))
            if ((779 <= 65535) || (18446744073709551615 >= 2)) {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    switch (fi42 ^ (-(fi42))) {
                    case -75:
                        v44 := GenS4{F0: 0}
                        _ = v44
                    case -1:
                        fmt.Println(v43)
                    default:
                        v45 := v1[2]
                        _ = v45
                    }
                case "":
                    v46 := v1[2]
                    _ = v46
                case "a":
                    switch -9223372036854775808 {
                    case 4294967295:
                        v15 = false
                    case 128:
                        v47 := GenS9{F0: -607, F1: -1, F2: 42}
                        _ = v47
                    case 32767:
                        for fi48 := 0; fi48 < 7; fi48++ {
                            v49 := GenS10{F0: 18446744073709551615, F1: v15, F2: false}
                            fmt.Println(v14, v15, v43)
                            switch "hello" {
                            case "hello":
                                v2 += 1
                            }
                            _ = fi48
                            _ = v49
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v50 int16 = 858
                _ = v50
            } else {
                _ = len(v3)
                for fi51 := 0; fi51 < 13; fi51++ {
                    v5["a"] = 2
                    var v52 uint16 = 0
                    v2 += 508
                    fmt.Println(v15, v14)
                    if true {
                        var v53 int32 = 2147483647
                        v1[2147483647] = (v15 && (3.14159 > 0.1))
                        if true {
                            if v15 {
                                switch "a" {
                                case "hello":
                                    if (-32768 >= 1) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v54 := GenS3{F0: 18446744073709551615, F1: (v52 * 42)}
                                            _ = v54
                                        case "":
                                            v55 := v5[("a" + "a")]
                                            _ = v55
                                        case "hello":
                                            v2 += v2
                                        }
                                        v43 += (v43 & 38)
                                        if (0 <= -2147483648) {
                                            v56 := GenS11{F0: (^(v52)), F1: 2}
                                            v5[""] = 2
                                            _ = v56
                                        } else {
                                            v14 += -1
                                        }
                                    }
                                case "foo bar":
                                    v57 := GenS12{F0: -32768, F1: 42, F2: (v43 * v43)}
                                    _ = v57
                                case "\n\t":
                                    v58 := GenS13{F0: 127, F1: 1}
                                    _ = v58
                                }
                                switch "" {
                                case "\n\t":
                                    v3 = append(v3, 65535)
                                    if v15 {
                                        v59 := GenS9{F0: (v53 >> ((v2) & 63)), F1: int32(v43), F2: (-(v53))}
                                        var v60 int16 = 0
                                        var v61 bool = v15
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v2 = v2
                                        v14 = (v14 & v14)
                                        v52 = 42
                                    }
                                case "hello":
                                    var v62 string = ("hello" + ("\n\t" + ("hello" + "a")))
                                    _ = v62
                                case "":
                                    if false {
                                        v14 = v14
                                    }
                                default:
                                    var v63 int16 = 1000
                                    _ = v63
                                }
                                switch "hello" {
                                case "foo bar":
                                    v64 := GenS7{F0: 255, F1: v53}
                                    _ = v64
                                case "hello":
                                    if ((v14 - -1) == -710) {
                                        v2 = 65535
                                        var v65 uint = 96
                                        _ = v65
                                    }
                                }
                                v52 = 65535
                                _ = len(v3)
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    v5[("hello" + "foo bar")] = 42
                                case "":
                                    var v66 int16 = 32767
                                    _ = v66
                                case "\n\t":
                                    _ = len(v5)
                                default:
                                    fmt.Println(v2, fi42)
                                }
                                var v67 int16 = -464
                                _ = v67
                            }
                        }
                        _ = v53
                    } else {
                        v68 := GenS14{F0: 42, F1: 0, F2: (fi51 << ((1) & 63))}
                        v14 += v14
                        if ((v4 * (fi42 % ((int(v14)) | 1))) >= (fi42 | v4)) {
                            v43 = (v43 | v43)
                            fmt.Println(fi42, v15, v2)
                        } else {
                            v14 -= (v14 * v14)
                            fmt.Println(v14, v52, v52)
                        }
                        if false {
                            var v69 int16 = -32768
                            var v70 int8 = v43
                            _ = v69
                            _ = v70
                        }
                        var v71 float32 = 0.1
                        _ = v68
                        _ = v71
                    }
                    _ = fi51
                    _ = v52
                }
            }
            _ = fi42
            _ = v43
        }
    }
    delete(v5, "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
