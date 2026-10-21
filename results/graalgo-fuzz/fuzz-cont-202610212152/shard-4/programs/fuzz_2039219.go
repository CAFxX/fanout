package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint; F2 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64) int8 {
    v6 := GenS1{F0: -2147483648, F1: 2, F2: (2 < 2)}
    var v7 int32 = -249
    if true {
        v8 := []int{-1, 91, 128, 9223372036854775807}
        var v9 float64 = 0.0
        var v10 int = -2147483648
        _ = v10
        _ = v8
        _ = v9
    }
    p0 = (p0 * p0)
    v11 := GenS1{F0: (v7 << ((65535) & 63)), F1: 2147483648, F2: (!(-65 >= 5))}
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    var v12 uint16 = 192
    v13 := []int{-1, 1}
    v12 = 42
    if (!(!(!true))) {
        for fi14 := 0; fi14 < 13; fi14++ {
            if (0 == 114) {
                fmt.Println(fi14)
                switch -118 {
                case 419:
                    fmt.Println(v12)
                case 662:
                    fmt.Println(fi14)
                }
                switch (-(fi14)) {
                case -622:
                    if (0 >= -278) {
                        if ((("\n\t" >= "\n\t") || false) && false) {
                            v12 -= (v12 * v12)
                        }
                        v12 -= 604
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS1{F0: 691, F1: 2, F2: true}
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        if (!true) {
                            for fi17 := 0; fi17 < 19; fi17++ {
                                switch "" {
                                case "\n\t":
                                    if len(v13) > 0 {
                                        v13[0] = (fi14 ^ 2147483647)
                                    }
                                default:
                                    v12 = (v12 ^ v12)
                                }
                                fmt.Println(fi17)
                                v18 := GenS2{F0: 42}
                                switch fi14 {
                                case 473:
                                    v19 := GenS3{F0: 42, F1: -9223372036854775808}
                                    _ = v19
                                }
                                switch fi17 {
                                case 4294967295:
                                    fmt.Println(fi17, fi14, fi14)
                                default:
                                    if len(v13) > 0 {
                                        v13[0] = fi17
                                    }
                                }
                                var v20 float64 = 0.1
                                v21 := GenS1{F0: int32(v12), F1: 4294967295, F2: false}
                                v12 = v12
                                var v22 int16 = -32768
                                _ = fi17
                                _ = v18
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            if (-1 == 1000) {
                                v12 = v12
                            } else {
                                fmt.Println(fi14, fi14)
                                _ = len(v13)
                                var v23 int8 = -128
                                _ = v23
                            }
                        } else {
                            v12 += 0
                            switch 32767 {
                            case 98:
                                var v24 uint8 = 239
                                _ = v24
                            }
                            if len(v13) > 0 {
                                v13[0] = fi14
                            }
                        }
                        var v25 int = (fi14 % (((^(fi14))) | 1))
                        for fi26 := 0; fi26 < 6; fi26++ {
                            fmt.Println(fi14, v25, fi26)
                            v13 = append(v13, (fi26 * (-(fi26))))
                            var v27 uint8 = 42
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    }
                case 4294967295:
                    var v28 int64 = 42
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 3; fi29++ {
                    switch fi29 {
                    case -1:
                        var v30 uint16 = (^(v12))
                        _ = v30
                    case 128:
                        v12 -= 1
                    case 127:
                        var v31 uint16 = 42
                        v13 = append(v13, 209)
                        _ = v31
                    default:
                        switch fi14 {
                        case 42:
                            v13 = append(v13, -9223372036854775808)
                        default:
                            var v32 string = "\n\t"
                            if len(v13) > 0 {
                                v13[0] = (fi14 % (((fi29 / ((fi29) | 1))) | 1))
                            }
                            var v33 uint32 = 42
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v32
                            _ = v33
                        }
                    }
                    v12 += v12
                    if (false && (18446744073709551615 != 400)) {
                        v34 := GenS1{F0: 777, F1: 1, F2: false}
                        _ = v34
                    } else {
                        v35 := GenS4{F0: 42}
                        v12 += 0
                        _ = v35
                    }
                    _ = fi29
                }
            }
            v12 += (v12 | (v12 >> ((4294967295) & 63)))
            _ = fi14
        }
    } else {
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := GenS5{F0: 0}
            _ = fi36
            _ = v37
        }
    }
    _ = v12
    _ = v13
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := []uint32{629, 1, 402}
    v5 := make(map[uint32]uint32)
    var v38 int32 = -1
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 uint16 = 65535
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        v38 = (v38 << ((113) & 63))
        _ = fi41
    }
    v38 = (v38 % ((0) | 1))
    switch 979 {
    case 128:
        fmt.Println(v38)
    case 1:
        if false {
            switch "a" {
            case "a":
                var v42 int8 = -128
                _ = v42
            case "":
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 uint64 = 42
                    _ = fi43
                    _ = v44
                }
            }
            var v45 bool = true
            if v45 {
                fmt.Println(v45, v45)
                var v46 uint16 = 38
                var v47 int = 42
                _ = v46
                _ = v47
            }
            _ = v45
        } else {
            _ = len(v2)
            v38 = v38
        }
    default:
        v48 := GenS6{F0: 233, F1: 18446744073709551615, F2: 2}
        _ = v48
    }
    v49 := GenS1{F0: (v38 + 1), F1: 471, F2: false}
    var v50 int16 = 0
    for fi51 := 0; fi51 < 7; fi51++ {
        v52 := GenS7{F0: fi51}
        fmt.Println(v50, v38, v50)
        _ = fi51
        _ = v52
    }
    var v53 int = 1
    v54 := GenS2{F0: 255}
    v55 := GenS8{F0: 0}
    if false {
        v50 += v50
        if len(v4) > 0 {
            v4[0] = 2
        }
    } else {
        var v56 uint8 = 2
        v57 := GenS4{F0: (^(v38))}
        _ = v56
        _ = v57
    }
    v53 += v53
    for fi58 := 0; fi58 < 11; fi58++ {
        for fi59 := 0; fi59 < 1; fi59++ {
            var v60 uint32 = 955
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    if false {
        switch (v53 << ((4294967295) & 63)) {
        case -853:
            _ = len(v4)
        case 0:
            v53 -= v53
        }
    } else {
        v61 := GenS3{F0: v38, F1: -9223372036854775808}
        _ = len(v5)
        v50 += 402
        _ = v61
    }
    var v62 uint32 = 757
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi63 := 0; fi63 < 3; fi63++ {
        for fi64 := 0; fi64 < 9; fi64++ {
            var v65 uint = 926
            if (v62 != v62) {
                v66 := GenS4{F0: (v38 << ((v65) & 63))}
                v67 := GenS9{F0: 127, F1: v50, F2: 506}
                v53 -= (fi63 ^ v53)
                _ = v66
                _ = v67
            }
            switch 32767 {
            case -2147483648:
                var v68 int8 = -54
                _ = v68
            default:
                v62 -= (v62 & (v62 % (((v62 - v62)) | 1)))
            }
            _ = fi64
            _ = v65
        }
        for fi69 := 0; fi69 < 16; fi69++ {
            if (-2.25 <= 1.5) {
                var v70 int8 = 127
                v71 := v5[2147483648]
                v72 := GenS10{F0: 42}
                v73 := GenS11{F0: v38, F1: 265}
                _ = v70
                _ = v71
                _ = v72
                _ = v73
            } else {
                fmt.Println(v50)
            }
            _ = fi69
        }
        _ = fi63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v62
}
