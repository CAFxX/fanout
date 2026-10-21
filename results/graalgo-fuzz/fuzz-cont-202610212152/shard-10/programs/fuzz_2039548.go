package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32; F1 int16; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int; F1 int; F2 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint64; F1 int32; F2 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int; F1 int16; F2 int }
type GenS14 struct { F0 int; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 int8) uint {
    v6 := GenS1{F0: 897, F1: p0}
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p1, fi7)
        _ = fi7
    }
    var v8 float32 = (3.14159 * 1e10)
    var v9 int32 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 0
    var v3 uint64 = 1
    var v4 int64 = v2
    var v5 int32 = 0
    var v10 int64 = (v2 - v2)
    var v11 uint = 42
    switch 1 {
    case 761:
        if (188 != 67) {
            v11 = genFn0((v10 & -9223372036854775808), -1)
        }
    }
    v5 += (v5 * v5)
    fmt.Println(v11, v5, v5)
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 float64 = 3.14159
        v14 := GenS1{F0: 2, F1: v4}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float32 = (0.0 / -2.25)
    v5 += v5
    v16 := GenS2{F0: 1, F1: 2147483647}
    var v17 int = -1
    for fi18 := 0; fi18 < 9; fi18++ {
        if false {
            v11 = (^(v11))
            switch "a" {
            case "a":
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 uint32 = 361
                    if (2 >= 42) {
                        if false {
                            var v21 string = ""
                            if true {
                                switch v21 {
                                case "hello":
                                    v22 := GenS1{F0: (v20 ^ v20), F1: v10}
                                    _ = v22
                                default:
                                    fmt.Println(v15)
                                }
                                if false {
                                    var v23 int32 = (v5 >> ((2) & 63))
                                    _ = v23
                                } else {
                                    var v24 int64 = (v2 >> ((0) & 63))
                                    v25 := GenS3{F0: -1, F1: 42, F2: 1}
                                    _ = v24
                                    _ = v25
                                }
                            }
                            switch v21 {
                            case "foo bar":
                                fmt.Println(v20)
                            case "\n\t":
                                if true {
                                    var v26 int64 = (v4 << ((v11) & 63))
                                    v27 := GenS4{F0: 32767, F1: 42}
                                    v11 = (v11 & v11)
                                    var v28 string = ((v21 + ((v21 + v21) + v21)) + "")
                                    v2 += 0
                                    switch 621 {
                                    case 9223372036854775807:
                                        v29 := []int{fi19, v17}
                                        _ = v29
                                    default:
                                        v30 := GenS4{F0: 32767, F1: 42}
                                        _ = v30
                                    }
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case "a":
                                v17 += fi18
                            default:
                                v31 := v1[v20]
                                _ = v31
                            }
                            _ = v21
                        }
                        if false {
                            var v32 uint16 = 374
                            switch (^(v17)) {
                            case -9223372036854775808:
                                var v33 int64 = (v2 ^ v10)
                                _ = v33
                            case 127:
                                v34 := []int{(fi18 / ((v17) | 1)), 127, 128}
                                if (fi19 <= (fi18 + v17)) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v20, v11)
                                        fmt.Println(fi18)
                                        if false {
                                            v11 = (v11 % (((v11 % (((v11 - 1)) | 1))) | 1))
                                        } else {
                                            v35 := GenS5{F0: (fi18 * (fi19 % ((9223372036854775807) | 1))), F1: (fi18 << (((v11 | (v11 * (v11 ^ (v11 | v11))))) & 63)), F2: (v17 % (((fi19 << ((v11) & 63))) | 1))}
                                            v36 := GenS2{F0: 255, F1: (v20 / ((v20) | 1))}
                                            var v37 float32 = v15
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := GenS5{F0: (v17 >> ((v11) & 63)), F1: 127, F2: (fi19 >> (((v11 | genFn0(v10, -1))) & 63))}
                                        v39 := GenS4{F0: 550, F1: 1}
                                        v32 = (^(v32))
                                        _ = v38
                                        _ = v39
                                    case "foo bar":
                                        switch (fi19 % ((fi18) | 1)) {
                                        case 9223372036854775807:
                                            var v40 int64 = v10
                                            _ = v40
                                        }
                                    }
                                } else {
                                    v17 = (fi19 << ((804) & 63))
                                    if ((true || (v10 == (v10 * 9223372036854775807))) && true) {
                                        if false {
                                            v3 += (v3 * (v3 - 1))
                                        } else {
                                            fmt.Println(v15, v2, v10)
                                        }
                                        fmt.Println(v15, fi18)
                                        if (!(!((true || false) || false))) {
                                            v41 := GenS6{F0: v10}
                                            _ = v41
                                        }
                                        v42 := GenS7{F0: 1}
                                        v34 = append(v34, (v17 >> ((42) & 63)))
                                        _ = v42
                                    }
                                }
                                if len(v34) > 0 {
                                    v34[0] = (fi19 + fi18)
                                }
                                _ = v34
                            }
                            v32 += 312
                            _ = v32
                        } else {
                            v43 := []int{v17, fi19, 32768, v17}
                            var v44 bool = false
                            _ = v43
                            _ = v44
                        }
                        if ((fi18 | int(v3)) == (^(fi19))) {
                            v45 := GenS8{F0: v3}
                            _ = v45
                        } else {
                            var v46 bool = false
                            _ = v46
                        }
                    }
                    v47 := GenS1{F0: (v20 | (^(v20))), F1: (v2 % (((v2 - v10)) | 1))}
                    _ = fi19
                    _ = v20
                    _ = v47
                }
            default:
                v48 := GenS9{F0: v5, F1: v3, F2: 2}
                _ = v48
            }
            if (!true) {
                v49 := GenS10{F0: v3, F1: (v5 >> ((genFn0((v2 >> (((^(v11))) & 63)), -1)) & 63)), F2: ("\n\t" + "hello")}
                var v50 uint = 220
                if true {
                    fmt.Println(v5, v17)
                    var v51 uint64 = v3
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 12; fi52++ {
                        if false {
                            v4 = v10
                            var v53 int = fi52
                            _ = v53
                        }
                        var v54 bool = true
                        _ = fi52
                        _ = v54
                    }
                    fmt.Println(v17, v10)
                    v55 := []int{(fi18 ^ fi18), v17, fi18, -58}
                    _ = v55
                }
                _ = v49
                _ = v50
            } else {
                var v56 int = (^(fi18))
                _ = v56
            }
        } else {
            v1[65535] = false
        }
        _ = fi18
    }
    if ("a" > ("hello" + "hello")) {
        v5 = (v5 >> ((650) & 63))
        v3 = (^(v3))
        v57 := GenS11{F0: int64(v3)}
        _ = v57
    } else {
        v58 := GenS2{F0: 196, F1: 4294967295}
        var v59 int64 = v10
        fmt.Println(v2, v11)
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 19; fi60++ {
        v61 := []int{(fi60 | (v17 % (((fi60 / (((v17 << ((v11) & 63))) | 1))) | 1))), 784, v17}
        var v62 int64 = (-(v10))
        _ = fi60
        _ = v61
        _ = v62
    }
    if (!true) {
        if true {
            v63 := GenS12{F0: -128}
            fmt.Println(v3, v10, v17)
            _ = v63
        } else {
            v64 := GenS13{F0: -203, F1: -32768, F2: (v17 >> ((636) & 63))}
            for fi65 := 0; fi65 < 19; fi65++ {
                for fi66 := 0; fi66 < 19; fi66++ {
                    v5 = (v5 % ((v5) | 1))
                    _ = fi66
                }
                for fi67 := 0; fi67 < 11; fi67++ {
                    v68 := []int{}
                    v69 := GenS6{F0: (v4 + v4)}
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = fi65
            }
            _ = v64
        }
        switch ("\n\t" + "foo bar") {
        case "":
            for fi70 := 0; fi70 < 4; fi70++ {
                v5 -= int32(v3)
                for fi71 := 0; fi71 < 5; fi71++ {
                    if ((!(false || (true || (2 >= 450)))) || (false || (!true))) {
                        if ((1.5 * 1e10) != 3.14159) {
                            v72 := []int{9223372036854775807, -926, fi71, (fi71 * -9223372036854775808)}
                            var v73 int8 = 50
                            v73 = int8(v5)
                            _ = v72
                            _ = v73
                        } else {
                            v10 = (v2 + v4)
                            v15 = ((v15 / (((((v15 + 0.0) * v15) / (v15 + (((1.5 / (v15 - (v15 * ((v15 / v15) + (v15 + (3.14159 * v15)))))) - v15) - (-2.25 - 1e10)))) * v15) + (float32((3.14159 * ((((float64(-2.25) - 0.1) * float64((v15 + float32(1e10)))) * float64(float32((0.0 - 0.1)))) + ((0.0 + (3.14159 - 1e10)) / ((float64(v15) + -2.25) - (0.0 + 0.1)))))) - v15))) + ((0.0 + v15) / v15))
                        }
                        v74 := GenS3{F0: 42, F1: -1, F2: 2}
                        _ = v74
                    } else {
                        v75 := []int{1}
                        if false {
                            v76 := v1[2]
                            _ = v76
                        }
                        _ = len(v75)
                        _ = v75
                    }
                    switch 42 {
                    case -1:
                        var v77 uint16 = 255
                        _ = v77
                    case -2147483648:
                        v78 := v1[343]
                        v10 = -137
                        v79 := []int{}
                        _ = v78
                        _ = v79
                    }
                    _ = fi71
                }
                _ = fi70
            }
        case "foo bar":
            var v80 float32 = 0.0
            _ = v80
        }
        var v81 int = v17
        _ = v81
    }
    for fi82 := 0; fi82 < 7; fi82++ {
        switch "" {
        case "a":
            var v83 uint = (v11 - (v11 | (v11 * (v11 / ((genFn0((v2 << ((18446744073709551615) & 63)), -10)) | 1)))))
            _ = v83
        case "hello":
            for fi84 := 0; fi84 < 1; fi84++ {
                v85 := GenS14{F0: 98, F1: -1, F2: 2}
                var v86 int64 = -156
                _ = fi84
                _ = v85
                _ = v86
            }
        case "foo bar":
            if true {
                v87 := GenS9{F0: v5, F1: 0, F2: 42}
                _ = v87
            } else {
                fmt.Println(v4, v17)
            }
        default:
            fmt.Println(v3, v2, v4)
        }
        _ = fi82
    }
    var v88 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
}
