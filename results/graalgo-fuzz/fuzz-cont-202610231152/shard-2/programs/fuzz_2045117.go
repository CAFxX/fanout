package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint8, p2 uint) int64 {
    v6 := GenS1{F0: 507, F1: -580, F2: (^(p1))}
    fmt.Println(p1, p2)
    var v7 int64 = -102
    var v8 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (v7 & v7)
}
func genFn1(p0 string) uint8 {
    if true {
        switch 9223372036854775807 {
        case 32768:
            switch p0 {
            case "foo bar":
                p0 = p0
            default:
                fmt.Println(p0)
            }
        case -262:
            fmt.Println(p0)
        }
        var v9 uint16 = 313
        p0 = "foo bar"
        _ = v9
    }
    v10 := []int{-1}
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 uint = 42
    var v3 uint = 0
    var v4 int8 = -113
    v5 := make(map[int]string)
    fmt.Println(v3, v4)
    if true {
        if false {
            switch 32767 {
            case -20:
                v11 := GenS1{F0: -799, F1: 346, F2: 1}
                _ = v11
            case -440:
                v12 := GenS1{F0: -2147483648, F1: 2147483647, F2: genFn1("hello")}
                _ = v12
            default:
                delete(v5, 32768)
            }
            fmt.Println(v1, v4, v1)
        }
        if true {
            fmt.Println(v1, v2)
            var v13 uint8 = 54
            _ = v13
        } else {
            var v14 int8 = (v1 % (((v1 >> ((v3) & 63))) | 1))
            v15 := GenS1{F0: 932, F1: 1, F2: 42}
            _ = v14
            _ = v15
        }
    } else {
        v2 = v3
    }
    v3 -= (^(v3))
    v4 += (v4 & (v4 * (v4 / ((42) | 1))))
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v3, v1)
        var v17 uint64 = 2
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 982, F1: 32767}
    if true {
        fmt.Println(v2, v2, v3)
        if true {
            v19 := GenS2{F0: 1, F1: -32768}
            _ = v19
        } else {
            var v20 int32 = -1
            _ = v20
        }
    } else {
        v21 := v5[-214]
        fmt.Println(v21)
        _ = v21
    }
    fmt.Println(v3)
    v4 += (v1 * (v4 & (-(v4))))
    v2 -= v2
    var v22 uint16 = 32767
    if (3.14159 != float64((0.0 / -2.25))) {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 float64 = float64(float32(0.1))
            if true {
                v1 += 42
                v25 := GenS3{F0: 918, F1: 0, F2: 42}
                _ = v25
            }
            switch (fi23 - fi23) {
            case 36:
                var v26 uint = (v2 ^ v2)
                _ = v26
            case 426:
                var v27 int8 = v4
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        v28 := []int{42}
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS1{F0: -871, F1: 122, F2: 42}
            if true {
                if (255 > genFn1(("\n\t" + ""))) {
                    var v31 int = fi29
                    switch "" {
                    case "\n\t":
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                    case "hello":
                        var v32 float64 = (1e10 / 1.0)
                        _ = v32
                    case "":
                        for fi33 := 0; fi33 < 12; fi33++ {
                            fmt.Println(fi33)
                            if (2 <= 95) {
                                v3 = (v2 ^ 1)
                                v34 := GenS2{F0: 377, F1: -32768}
                                _ = v34
                            } else {
                                if (301 > genFn0((v22 % (((v22 | v22)) | 1)), genFn1("a"), v3)) {
                                    if false {
                                        var v35 int32 = 42
                                        v4 = v1
                                        v36 := v5[fi29]
                                        var v37 int = int(v1)
                                        switch "foo bar" {
                                        case "":
                                            if false {
                                                v31 = (fi29 | -797)
                                            } else {
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                            }
                                        }
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            }
                            if (fi33 >= fi29) {
                                v4 -= (v1 + (v1 | v4))
                                var v38 int = (v31 << ((v3) & 63))
                                var v39 string = "\n\t"
                                v31 += (-(v38))
                                if len(v28) > 0 {
                                    v28[0] = (v31 | v31)
                                }
                                v40 := GenS1{F0: -425, F1: -1, F2: 42}
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                var v41 string = ""
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                                if len(v28) > 0 {
                                    v28[0] = (fi33 / ((fi33) | 1))
                                }
                                if len(v28) > 0 {
                                    v28[0] = (fi33 << (((v3 + v3)) & 63))
                                }
                                _ = v41
                            }
                            v31 -= fi29
                            _ = fi33
                        }
                    default:
                        if (!(-2.25 > 0.1)) {
                            if (true && true) {
                                v4 = v4
                                _ = len(v28)
                            }
                            for fi42 := 0; fi42 < 19; fi42++ {
                                v28 = append(v28, (fi42 | (v31 - fi29)))
                                if (!(32767 <= -601)) {
                                    var v43 float32 = 1.5
                                    if false {
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        fmt.Println(fi29, v2, v43)
                                    } else {
                                        v44 := GenS4{F0: true}
                                        if true {
                                            var v45 uint16 = 2
                                            if (((-411 == 1) && false) && false) {
                                                var v46 bool = (v45 >= 847)
                                                v31 -= fi29
                                                if (v46 || v46) {
                                                    if len(v28) > 0 {
                                                        v28[0] = (fi29 / (((fi42 / ((9223372036854775807) | 1))) | 1))
                                                    }
                                                    if (18446744073709551615 == 806) {
                                                        v28 = append(v28, 42)
                                                    } else {
                                                        if v46 {
                                                            if true {
                                                                if v46 {
                                                                    var v47 float32 = v43
                                                                    fmt.Println(v2, v46, v46)
                                                                    _ = v47
                                                                }
                                                                fmt.Println(v3, v22, v31)
                                                                v5[(fi42 / ((fi29) | 1))] = "a"
                                                            }
                                                            var v48 uint8 = 42
                                                            var v49 int = fi42
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            var v50 int8 = (v4 % ((v4) | 1))
                                                            var v51 uint16 = (v22 - 2)
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        if false {
                                                            v52 := GenS2{F0: 17, F1: -1}
                                                            _ = v52
                                                        }
                                                        var v53 float32 = 1e10
                                                        _ = v53
                                                    }
                                                } else {
                                                    v54 := GenS5{F0: genFn0(v45, genFn1(""), v2), F1: ("hello" + "foo bar")}
                                                    v1 -= v4
                                                    v4 = (v4 >> ((v2) & 63))
                                                    _ = v54
                                                }
                                                if len(v28) > 0 {
                                                    v28[0] = (fi29 ^ v31)
                                                }
                                                switch 1 {
                                                case 293:
                                                    v5[742] = "a"
                                                }
                                                _ = v46
                                            }
                                            v31 -= int(v4)
                                            _ = v45
                                        }
                                        var v55 int8 = v1
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        if true {
                                            var v56 int32 = 42
                                            v57 := GenS3{F0: -1, F1: genFn0(v22, 128, (v3 % ((v2) | 1))), F2: 4294967295}
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = v44
                                        _ = v55
                                    }
                                    v58 := GenS1{F0: 192, F1: 804, F2: 1}
                                    var v59 uint = (v2 / (((v2 | 0)) | 1))
                                    _ = v43
                                    _ = v58
                                    _ = v59
                                } else {
                                    var v60 float64 = (1.5 * ((float64(1e10) * ((1.5 / 1.0) - ((1.5 + 0.1) - 0.1))) / 1.0))
                                    var v61 uint = v2
                                    switch fi42 {
                                    case 682:
                                        v62 := GenS6{F0: (^(v31)), F1: 42, F2: 327}
                                        _ = v62
                                    default:
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    }
                                    _ = v60
                                    _ = v61
                                }
                                _ = fi42
                            }
                        }
                    }
                    _ = v31
                } else {
                    if true {
                        for fi63 := 0; fi63 < 14; fi63++ {
                            v4 -= 1
                            fmt.Println(v22, v4, v1)
                            _ = fi63
                        }
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                    } else {
                        var v64 uint64 = 4294967295
                        for fi65 := 0; fi65 < 13; fi65++ {
                            if (true && false) {
                                v3 = 42
                                v1 = (v1 - v4)
                                if false {
                                    if (uint32(v22) <= 4294967295) {
                                        v3 = 42
                                        var v66 int64 = 829
                                        _ = v66
                                    }
                                    v67 := GenS7{F0: (^(v22))}
                                    _ = v67
                                } else {
                                    fmt.Println(v4, fi29, v1)
                                }
                                _ = len(v5)
                            } else {
                                v64 = (v64 + v64)
                            }
                            _ = fi65
                        }
                        _ = v64
                    }
                    v68 := GenS6{F0: fi29, F1: uint16(v2), F2: 4294967295}
                    _ = v68
                }
                var v69 uint8 = 255
                var v70 int32 = 42
                _ = v69
                _ = v70
            } else {
                v1 = 96
            }
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        var v71 uint8 = 231
        _ = v71
    }
    var v72 int64 = int64(v1)
    var v73 uint64 = 2
    var v74 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
}
