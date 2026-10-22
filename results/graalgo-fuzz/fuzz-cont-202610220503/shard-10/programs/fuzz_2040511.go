package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
type GenS7 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS8 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 int, p1 int32, p2 int8) uint {
    var v6 uint16 = 1
    var v7 uint8 = 234
    var v8 uint8 = 72
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint16) uint16 {
    fmt.Println(p0)
    if true {
        p0 = p0
        p0 += 312
        if true {
            var v9 uint8 = 42
            _ = v9
        }
    }
    switch "" {
    case "":
        v10 := GenS1{F0: -59, F1: p0, F2: 205}
        _ = v10
    case "foo bar":
        var v11 int = 103
        _ = v11
    default:
        var v12 float32 = (float32(3.14159) / ((1e10 * float32(1.5)) / float32((((-2.25 + 3.14159) + ((3.14159 + ((0.1 - (1e10 - ((0.0 / ((-2.25 - 0.1) + -2.25)) / (0.0 - -2.25)))) - -2.25)) - 0.0)) / -2.25))))
        _ = v12
    }
    var v13 int16 = -855
    v14 := GenS2{F0: 1, F1: 569}
    _ = p0
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = v1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 float64 = ((0.1 + (1e10 - 0.0)) + ((0.1 + float64((float32(3.14159) + float32(1e10)))) - 0.0))
    var v15 float64 = 0.1
    var v16 float32 = 3.14159
    v5 += 3.14159
    if true {
        if (v5 <= ((v15 * 0.0) - ((0.1 * v15) - (1e10 + 1.5)))) {
            v1 -= (v1 + v2)
            v16 = 0.1
        } else {
            v17 := GenS3{F0: (^(v1)), F1: 1, F2: 409}
            _ = v17
        }
        var v18 float64 = v15
        _ = v18
    } else {
        v2 = 1
        v19 := GenS4{F0: -412, F1: 4294967295, F2: 127}
        delete(v3, 1)
        _ = v19
    }
    if ("a" > "hello") {
        switch ("" + "foo bar") {
        case "foo bar":
            for fi20 := 0; fi20 < 15; fi20++ {
                for fi21 := 0; fi21 < 13; fi21++ {
                    switch fi21 {
                    case 4294967295:
                        v15 += float64((float32((((v5 + float64(v16)) + v15) / ((v15 * ((0.1 / (v5 + 3.14159)) * 0.1)) * v5))) * float32(v15)))
                    }
                    _ = fi21
                }
                v16 = (v16 * v16)
                if (0 <= int32(fi20)) {
                    var v22 uint8 = v1
                    delete(v4, -493)
                    _ = v22
                } else {
                    v2 = 1
                }
                _ = fi20
            }
        case "a":
            var v23 bool = (true || false)
            _ = v23
        case "\n\t":
            var v24 float64 = (v5 / 1.5)
            _ = v24
        default:
            if false {
                v25 := GenS2{F0: -1, F1: 42}
                var v26 float64 = v15
                if false {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 float64 = v15
                        _ = fi27
                        _ = v28
                    }
                }
                _ = v25
                _ = v26
            } else {
                v29 := v4[128]
                v5 = 0.1
                _ = v29
            }
        }
    } else {
        var v30 int = -507
        fmt.Println(v1, v2)
        _ = v30
    }
    if (!(2147483647 > 535)) {
        v2 -= v2
    } else {
        var v31 uint8 = 128
        if ((false || ((42 >= 140) || false)) && false) {
            var v32 int = 108
            fmt.Println(v2, v32, v32)
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS4{F0: 0, F1: 2147483647, F2: v2}
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        _ = v31
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v15, v2)
        for fi36 := 0; fi36 < 16; fi36++ {
            v5 = (v15 * v5)
            v37 := v4[-640]
            if true {
                v37 = v37
                v16 -= ((1e10 - v16) + 0.0)
            }
            if true {
                switch (fi36 / (((fi35 >> ((871) & 63))) | 1)) {
                case -1:
                    v38 := GenS1{F0: 23, F1: 1, F2: v1}
                    _ = v38
                case 42:
                    if false {
                        var v39 bool = (1 > 1)
                        if false {
                            v5 = v15
                            v2 -= (v1 | 243)
                        } else {
                            v37 = (^(v37))
                            v40 := GenS4{F0: 32767, F1: 0, F2: (v1 & (v2 | (^(v1))))}
                            v39 = (0 == 475)
                            _ = v40
                        }
                        switch "" {
                        case "foo bar":
                            var v41 uint32 = (^(v37))
                            _ = v41
                        case "":
                            v1 = (v1 % ((v1) | 1))
                        case "\n\t":
                            v42 := GenS3{F0: (v1 / (((v2 >> ((2) & 63))) | 1)), F1: 2, F2: 431}
                            _ = v42
                        }
                        var v43 string = "a"
                        _ = v39
                        _ = v43
                    }
                case 128:
                    var v44 int8 = 100
                    _ = v44
                default:
                    fmt.Println(fi36, fi35, fi35)
                    v45 := GenS3{F0: (v1 * v2), F1: (v37 * v37), F2: 940}
                    v16 += v16
                    if (true && (uint16(v37) <= 748)) {
                        var v46 uint8 = (v2 & (v2 % (((^(v1))) | 1)))
                        fmt.Println(v1, fi36)
                        v47 := GenS4{F0: 1, F1: (v37 * 453), F2: v2}
                        _ = v46
                        _ = v47
                    } else {
                        var v48 uint8 = (v1 % (((^(v1))) | 1))
                        v1 -= 2
                        _ = v48
                    }
                    switch fi35 {
                    case 32767:
                        switch fi35 {
                        case 42:
                            if (true && (1 <= -2147483648)) {
                                v49 := GenS5{F0: (v37 - (^(v37))), F1: 33}
                                fmt.Println(fi36)
                                v15 -= (((((v15 * float64(v16)) + (float64(v16) - v5)) - v5) - v5) / v15)
                                _ = v49
                            } else {
                                if (!false) {
                                    var v50 uint16 = 1
                                    v1 = 94
                                    _ = v50
                                }
                                v16 -= v16
                                var v51 uint64 = 162
                                var v52 float64 = (1.5 * float64(float32(float64((v16 / 1.5)))))
                                var v53 uint8 = 128
                                v53 -= (^(v53))
                                if false {
                                    var v54 uint32 = v37
                                    v55 := []int{(fi36 * (fi35 << ((genFn0(fi36, 0, -103)) & 63))), 128, fi36, (fi35 >> ((195) & 63))}
                                    v56 := GenS4{F0: 32767, F1: 497, F2: 128}
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        }
                    case 183:
                        var v57 string = ("" + ("\n\t" + "hello"))
                        _ = v57
                    default:
                        if true {
                            _ = len(v4)
                            v58 := GenS6{F0: (fi35 / ((0) | 1)), F1: -1, F2: -1}
                            if (false || (!true)) {
                                v59 := GenS7{F0: 2, F1: 489, F2: 2}
                                v37 -= (v37 + 2147483647)
                                _ = v59
                            }
                            _ = v58
                        }
                    }
                    fmt.Println(v37, v15)
                    if (false || (42 > -1)) {
                        if true {
                            v5 += v15
                            v16 = 0.0
                        } else {
                            v16 -= v16
                            if true {
                                if ((32767 < 1000) && ((!true) || true)) {
                                    v60 := GenS2{F0: 127, F1: uint(v37)}
                                    v61 := []int{(-(fi35)), -9223372036854775808}
                                    _ = v60
                                    _ = v61
                                }
                                v62 := GenS2{F0: 1, F1: 2}
                                v16 -= (v16 - ((3.14159 / v16) + float32(v15)))
                                var v63 uint16 = 2
                                _ = v62
                                _ = v63
                            } else {
                                switch fi35 {
                                case 42:
                                    fmt.Println(v2)
                                default:
                                    v15 = 0.1
                                }
                            }
                        }
                        var v64 uint16 = 65535
                        var v65 uint8 = uint8(v37)
                        v15 = 3.14159
                        v1 -= 127
                        _ = v64
                        _ = v65
                    } else {
                        switch (fi36 << ((297) & 63)) {
                        case -274:
                            if (458 <= 2) {
                                v66 := GenS2{F0: int8(v2), F1: 1}
                                var v67 int16 = 461
                                switch ("foo bar" + ("a" + "\n\t")) {
                                case "":
                                    var v68 uint16 = 255
                                    _ = v68
                                }
                                _ = v66
                                _ = v67
                            }
                        case -79:
                            v37 = 2
                            fmt.Println(v15, v16)
                            if (!((v37 ^ 42) >= v37)) {
                                var v69 uint16 = 0
                                _ = v69
                            } else {
                                var v70 int = 128
                                v16 -= v16
                                _ = v70
                            }
                        default:
                            var v71 uint = 455
                            _ = v71
                        }
                        v2 = v2
                        var v72 int8 = 0
                        _ = v72
                    }
                    var v73 uint8 = v1
                    _ = v45
                    _ = v73
                }
                var v74 int64 = -9223372036854775808
                var v75 uint16 = 1
                _ = v74
                _ = v75
            }
            if true {
                v76 := GenS8{F0: 589, F1: -1, F2: v2}
                var v77 int8 = -104
                v16 += 0.1
                v78 := []int{fi36}
                v79 := GenS3{F0: (v2 | v2), F1: 297, F2: 0}
                _ = v76
                _ = v77
                _ = v78
                _ = v79
            } else {
                if false {
                    fmt.Println(v1)
                    v5 = v5
                    var v80 int16 = 1
                    _ = v80
                } else {
                    delete(v3, (^(v37)))
                    switch "\n\t" {
                    case "\n\t":
                        var v81 int8 = 42
                        _ = v81
                    case "foo bar":
                        v15 = v15
                    case "":
                        v82 := []int{}
                        v83 := GenS6{F0: fi36, F1: 0, F2: 0}
                        v5 = (v5 / (v5 * -2.25))
                        _ = v82
                        _ = v83
                    }
                    fmt.Println(v5)
                    v16 = 1e10
                    v2 = v1
                }
            }
            _ = fi36
            _ = v37
        }
        var v84 string = "\n\t"
        _ = fi35
        _ = v84
    }
    v5 -= v5
    v5 -= (3.14159 + float64(v16))
    fmt.Println(v16, v16, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
