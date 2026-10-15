package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
func genFn0(p0 uint8) uint16 {
    var v7 uint32 = 355
    if false {
        v8 := GenS1{F0: 0}
        _ = v8
    }
    _ = p0
    _ = v7
    return 32767
}

func main() {
    var v1 uint64 = 13
    var v2 uint = 771
    v3 := []uint32{763, 2}
    v4 := make(map[uint32]int)
    var v5 bool = (("hello" == "a") || (!true))
    v6 := []uint32{}
    var v9 uint16 = 1
    if false {
        fmt.Println(v2, v2)
        switch 0 {
        case 487:
            var v10 uint16 = (v9 | v9)
            _ = v10
        case 1:
            v11 := GenS1{F0: (v1 >> ((uint(v9)) & 63))}
            _ = v11
        }
        var v12 float32 = 3.14159
        _ = v12
    }
    switch -9223372036854775808 {
    case -268:
        _ = len(v4)
    case -2147483648:
        v13 := v4[2]
        _ = v13
    case 160:
        var v14 string = "foo bar"
        _ = v14
    }
    v15 := GenS2{F0: genFn0(255), F1: -399, F2: genFn0(127)}
    var v16 int16 = 1
    if ((1e10 * -2.25) == 3.14159) {
        v17 := GenS2{F0: genFn0(89), F1: 1, F2: v9}
        _ = v17
    } else {
        if v5 {
            if v5 {
                v1 -= (v1 | v1)
                for fi18 := 0; fi18 < 6; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        switch "foo bar" {
                        case "":
                            if (v2 > v2) {
                                var v20 uint8 = 1
                                _ = v20
                            }
                        case "foo bar":
                            v2 = v2
                        default:
                            v5 = v5
                        }
                        v5 = (!((1.5 / 1.0) == 3.14159))
                        switch -992 {
                        case 121:
                            v16 = (v16 ^ (^(v16)))
                        case -2147483648:
                            v5 = (!(886 < v1))
                        default:
                            v9 = v9
                            var v21 uint8 = uint8(v2)
                            _ = v21
                        }
                        _ = fi19
                    }
                    var v22 int32 = 1
                    var v23 uint32 = 2
                    _ = fi18
                    _ = v22
                    _ = v23
                }
                var v24 int16 = (-(v16))
                _ = v24
            }
            if (100 == 22) {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v1 += uint64(v9)
                    var v26 uint8 = uint8(v9)
                    _ = fi25
                    _ = v26
                }
                v27 := GenS3{F0: -1, F1: 258}
                v2 += (v2 % (((v2 % ((v2) | 1))) | 1))
                _ = v27
            }
        } else {
            if (127 != 128) {
                v2 = 65535
                switch "\n\t" {
                case "\n\t":
                    for fi28 := 0; fi28 < 10; fi28++ {
                        fmt.Println(v5, v5)
                        _ = fi28
                    }
                case "":
                    v29 := GenS1{F0: v1}
                    _ = v29
                default:
                    var v30 uint64 = (v1 + v1)
                    _ = v30
                }
            }
        }
    }
    v1 = (^(v1))
    var v31 uint = v2
    for fi32 := 0; fi32 < 6; fi32++ {
        v16 = v16
        var v33 uint8 = 197
        if v5 {
            v34 := GenS2{F0: genFn0((v33 | v33)), F1: 42, F2: v9}
            _ = v34
        } else {
            if true {
                for fi35 := 0; fi35 < 11; fi35++ {
                    v36 := GenS4{F0: -128}
                    v3 = append(v3, 4294967295)
                    var v37 float64 = (float64(float32(0.0)) - ((3.14159 / 1e10) * (0.1 - (float64(3.14159) / (1.5 / 1.5)))))
                    v5 = v5
                    if v5 {
                        v1 += (^(v1))
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            var v38 uint32 = 2147483647
            v9 = (v9 | v9)
            _ = v38
        }
        _ = fi32
        _ = v33
    }
    var v39 uint64 = 0
    var v40 bool = ("foo bar" == "")
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := GenS5{F0: 0, F1: v16}
        v4[2] = fi41
        if (1 < 218) {
            fmt.Println(fi41, v1)
            switch (fi41 ^ 1) {
            case 32768:
                var v43 float64 = (float64((0.1 - 0.1)) - 0.1)
                _ = v43
            case -1:
                if false {
                    var v44 uint8 = 128
                    for fi45 := 0; fi45 < 6; fi45++ {
                        if (!false) {
                            fmt.Println(v5, v9)
                            if (fi45 == (-(fi41))) {
                                v16 -= v16
                                v16 += 1
                                if v5 {
                                    fmt.Println(v2, v39)
                                }
                                fmt.Println(v5, v40)
                                if v40 {
                                    v16 = -155
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    v46 := GenS1{F0: 0}
                                    _ = v46
                                }
                                switch fi45 {
                                case -2147483648:
                                    v47 := GenS6{F0: -1}
                                    _ = v47
                                }
                                var v48 string = ""
                                _ = v48
                            } else {
                                if false {
                                    if ((v31 - 255) <= v31) {
                                        var v49 uint64 = v1
                                        fmt.Println(v16)
                                        fmt.Println(v16, fi45)
                                        _ = v49
                                    } else {
                                        var v50 int32 = 42
                                        _ = v50
                                    }
                                }
                                switch fi41 {
                                case 237:
                                    v16 += 209
                                case 32768:
                                    var v51 uint16 = 534
                                    _ = v51
                                default:
                                    fmt.Println(fi41)
                                }
                            }
                            if (!false) {
                                if v40 {
                                    v9 += 1
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = 6
                                    }
                                    v1 = v1
                                }
                            } else {
                                var v52 int64 = 1
                                var v53 int = fi45
                                _ = v52
                                _ = v53
                            }
                        }
                        if (!(v40 || ((43 > v44) || (!v5)))) {
                            _ = len(v4)
                        }
                        switch fi45 {
                        case 0:
                            if (1 == fi45) {
                                v44 -= (v44 << (((v2 | 624)) & 63))
                            }
                        case 127:
                            var v54 float64 = (0.1 / 1.0)
                            _ = v54
                        default:
                            if v5 {
                                v55 := GenS7{F0: (fi45 / ((fi45) | 1)), F1: 389}
                                _ = len(v4)
                                v56 := GenS2{F0: (v9 << ((v2) & 63)), F1: 0, F2: (v9 + 255)}
                                _ = v55
                                _ = v56
                            } else {
                                var v57 string = ""
                                switch fi45 {
                                case 9223372036854775807:
                                    switch fi45 {
                                    case 32768:
                                        v4[2147483648] = (fi41 / (((fi45 >> ((65535) & 63))) | 1))
                                    case 9223372036854775807:
                                        v58 := GenS8{F0: 4294967295}
                                        _ = v58
                                    }
                                case 128:
                                    v59 := v4[735]
                                    _ = v59
                                }
                                fmt.Println(fi45)
                                _ = v57
                            }
                        }
                        _ = fi45
                    }
                    for fi60 := 0; fi60 < 3; fi60++ {
                        if v40 {
                            if v5 {
                                v61 := GenS3{F0: 1, F1: genFn0(v44)}
                                _ = v61
                            } else {
                                v31 += v2
                            }
                            v62 := GenS7{F0: 0, F1: v1}
                            v63 := GenS9{F0: (v9 > genFn0(v44))}
                            _ = v62
                            _ = v63
                        }
                        v9 -= v9
                        _ = fi60
                    }
                    _ = v44
                }
            case 128:
                var v64 uint16 = 32767
                _ = v64
            }
            v2 += v2
        } else {
            for fi65 := 0; fi65 < 10; fi65++ {
                v39 += uint64(v9)
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi65
            }
            for fi66 := 0; fi66 < 20; fi66++ {
                fmt.Println(v2, v2, v2)
                var v67 int = (fi66 % ((fi66) | 1))
                switch fi41 {
                case 988:
                    v68 := GenS8{F0: fi66}
                    _ = v68
                default:
                    v9 -= 1
                }
                _ = fi66
                _ = v67
            }
            v9 += (v9 / (((v9 & 672)) | 1))
        }
        _ = fi41
        _ = v42
    }
    var v69 uint = 65535
    fmt.Println(v16, v31, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v69
    _ = v9
}
