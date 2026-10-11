package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 uint) int8 {
    p0 += (^(p0))
    if true {
        p0 += (p0 * p0)
        var v7 uint16 = p0
        switch "a" {
        case "a":
            fmt.Println(p1, v7)
        case "hello":
            fmt.Println(v7)
        default:
            var v8 bool = (-1 <= 42)
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int = -532
    var v2 int64 = 1
    var v3 uint32 = 0
    var v4 uint32 = v3
    v5 := make(map[int]bool)
    var v6 string = "a"
    var v9 uint32 = 2
    var v10 int = 2147483647
    var v11 uint8 = 242
    var v12 bool = ((!(((1 >= 42) || true) && (!false))) || (3.14159 < (1e10 - 0.1)))
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{(v10 | (v1 ^ (-(v1))))}
        _ = fi13
        _ = v14
    }
    var v15 uint = 0
    delete(v5, (^(v1)))
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int64 = int64(v15)
        var v18 int64 = (-(v17))
        if false {
            v19 := GenS1{F0: 780, F1: v6, F2: -1}
            _ = v19
        } else {
            if ((v9 & v4) != (^(v4))) {
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint16 = 42
                    _ = fi20
                    _ = v21
                }
                v5[v10] = (!v12)
                if (!v12) {
                    fmt.Println(v11, v2)
                    if v12 {
                        var v22 int32 = -117
                        _ = v22
                    }
                }
            }
            switch (v6 + v6) {
            case "hello":
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int64 = v17
                    if false {
                        _ = len(v5)
                        if v12 {
                            if (0 < 2147483647) {
                                v10 += 127
                            }
                        }
                        v25 := GenS1{F0: v4, F1: v6, F2: 42}
                        if (v12 || true) {
                            if v12 {
                                fmt.Println(v11, v24, v18)
                                v26 := GenS2{F0: 1, F1: 0, F2: ("hello" + (v6 + ("a" + (((v6 + (v6 + "hello")) + "foo bar") + "a"))))}
                                _ = v26
                            }
                            switch v1 {
                            case 1:
                                if v12 {
                                    var v27 uint16 = 2
                                    v6 = ((v6 + (v6 + (v6 + (v6 + (v6 + "a"))))) + (v6 + v6))
                                    _ = v27
                                }
                            case 4294967295:
                                var v28 int32 = 945
                                v29 := GenS1{F0: v9, F1: v6, F2: 1000}
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = len(v5)
                        if true {
                            switch ((v6 + "") + v6) {
                            case "foo bar":
                                v18 = 42
                            case "":
                                var v30 float64 = 3.14159
                                _ = v30
                            case "a":
                                var v31 string = v6
                                _ = v31
                            default:
                                var v32 int8 = int8(v15)
                                _ = v32
                            }
                            if false {
                                v3 += (v3 * v4)
                                var v33 uint64 = 616
                                var v34 string = v6
                                var v35 float64 = 3.14159
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        } else {
                            v4 -= (v3 / (((v3 & v4)) | 1))
                            var v36 uint32 = v9
                            fmt.Println(v9)
                            _ = v36
                        }
                        v12 = false
                        _ = v25
                    } else {
                        if (v6 < (v6 + (v6 + ("a" + "")))) {
                            var v37 float64 = 0.0
                            if (false || (131 >= v2)) {
                                v38 := GenS3{F0: (v3 + 859)}
                                _ = v38
                            }
                            v15 = v15
                            _ = v37
                        } else {
                            v12 = false
                            v4 += 1
                            var v39 uint64 = 2
                            _ = v39
                        }
                        fmt.Println(v1, v3, v6)
                        v40 := GenS2{F0: -1, F1: 32767, F2: ""}
                        _ = v40
                    }
                    switch (v10 | (v1 / (((v1 >> ((v15) & 63))) | 1))) {
                    case -736:
                        v41 := GenS4{F0: (^(v3)), F1: 2, F2: 1000}
                        var v42 uint64 = 733
                        fmt.Println(v11, v42)
                        var v43 uint = uint(v11)
                        if v12 {
                            switch "foo bar" {
                            case "a":
                                if false {
                                    v18 += v18
                                    var v44 int8 = genFn0(0, 2147483648, v15)
                                    v45 := GenS4{F0: v9, F1: 334, F2: 0}
                                    _ = v44
                                    _ = v45
                                } else {
                                    if v12 {
                                        var v46 float32 = 0.1
                                        if ((-2.25 - (1e10 - (3.14159 / 1.0))) != float64(float32(((0.0 * 0.0) - 1.5)))) {
                                            v47 := GenS2{F0: 80, F1: 0, F2: v6}
                                            fmt.Println(v9, fi16, v11)
                                            v48 := GenS4{F0: (v3 ^ (^(v3))), F1: 0, F2: 32767}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v49 uint8 = uint8(v9)
                                            var v50 uint8 = (^(v49))
                                            fmt.Println(v11, v12)
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v46
                                    }
                                    switch v6 {
                                    case "hello":
                                        if v12 {
                                            switch (v10 | (v10 + 2147483647)) {
                                            case -2147483648:
                                                v51 := []int{fi23, -2147483648, -9223372036854775808}
                                                _ = v51
                                            }
                                            v11 += v11
                                            switch 1 {
                                            case -2147483648:
                                                if (!(255 < 197)) {
                                                    _ = len(v5)
                                                    v3 -= (v9 + v4)
                                                } else {
                                                    v52 := []int{32767, v1, (v1 | (fi16 * (^(fi23)))), (v1 - (v1 ^ v1))}
                                                    v53 := GenS4{F0: uint32(v15), F1: 1, F2: 0}
                                                    _ = v52
                                                    _ = v53
                                                }
                                            case 32767:
                                                var v54 uint8 = (v11 - (v11 - (v11 << ((v43) & 63))))
                                                _ = v54
                                            }
                                        } else {
                                            v18 = v18
                                            var v55 int64 = (v17 - v18)
                                            var v56 int8 = 100
                                            fmt.Println(v1, v3, v9)
                                            switch v6 {
                                            case "\n\t":
                                                switch -859 {
                                                case -1:
                                                    switch 1 {
                                                    case 1:
                                                        fmt.Println(v24, v2)
                                                    case 2147483647:
                                                        v2 -= -1
                                                    }
                                                default:
                                                    switch (v6 + ((v6 + ("" + v6)) + (v6 + v6))) {
                                                    case "a":
                                                        var v57 uint = (v43 * (^(v43)))
                                                        _ = v57
                                                    case "foo bar":
                                                        var v58 int = (v1 ^ (fi16 & v10))
                                                        _ = v58
                                                    case "hello":
                                                        fmt.Println(fi16)
                                                    default:
                                                        switch (((((v6 + ("\n\t" + v6)) + v6) + v6) + (v6 + "\n\t")) + v6) {
                                                        case "hello":
                                                            v6 = v6
                                                        case "\n\t":
                                                            if (v42 != 2) {
                                                                fmt.Println(v11)
                                                                var v59 int16 = 1
                                                                if v12 {
                                                                    v60 := GenS3{F0: 499}
                                                                    _ = v60
                                                                }
                                                                _ = v59
                                                            }
                                                            v61 := []int{-9223372036854775808, fi16, fi16, fi23}
                                                            _ = v61
                                                        case "":
                                                            if ((!true) && v12) {
                                                                v62 := GenS5{F0: v4, F1: ((v24 * (v17 + v17)) <= (v2 % ((int64(fi23)) | 1)))}
                                                                v63 := v5[(v10 >> ((v43) & 63))]
                                                                _ = v62
                                                                _ = v63
                                                            } else {
                                                                if v12 {
                                                                    var v64 int8 = v56
                                                                    _ = v64
                                                                } else {
                                                                    var v65 int = (v1 >> (((v15 | (v15 * v43))) & 63))
                                                                    v66 := GenS6{F0: int16(v65), F1: int8(v9)}
                                                                    if v12 {
                                                                        var v67 float32 = (1.5 / 0.1)
                                                                        _ = v67
                                                                    }
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                fmt.Println(v55)
                                                                v68 := []int{(fi16 % ((-1) | 1)), 2147483647, v10}
                                                                _ = v68
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v4)
                                                }
                                            case "":
                                                v69 := GenS7{F0: 71, F1: v17}
                                                _ = v69
                                            default:
                                                fmt.Println(v2)
                                            }
                                            _ = v55
                                            _ = v56
                                        }
                                    case "\n\t":
                                        v2 += v18
                                    default:
                                        v24 = (v24 - v17)
                                    }
                                }
                            case "foo bar":
                                v6 = "foo bar"
                            case "":
                                var v70 uint = uint(v4)
                                var v71 int = fi16
                                _ = v70
                                _ = v71
                            default:
                                v24 += v24
                            }
                            fmt.Println(fi23, v17)
                        } else {
                            v72 := []int{}
                            switch 2147483647 {
                            case 0:
                                v73 := GenS8{F0: 1, F1: (v42 % ((4294967295) | 1))}
                                _ = v73
                            }
                            _ = v72
                        }
                        v10 -= (v10 % ((128) | 1))
                        _ = v41
                        _ = v42
                        _ = v43
                    case -472:
                        var v74 uint16 = 355
                        _ = v74
                    case -264:
                        switch (fi16 / (((fi23 ^ 2147483647)) | 1)) {
                        case -761:
                            v4 -= (^(v3))
                        default:
                            var v75 uint16 = 1
                            _ = v75
                        }
                        fmt.Println(v1, fi16, v15)
                        var v76 uint = v15
                        _ = v76
                    default:
                        v5[(v1 | fi16)] = v12
                    }
                    _ = fi23
                    _ = v24
                }
            case "\n\t":
                v4 = v4
            case "foo bar":
                v1 -= (^(v1))
            }
            for fi77 := 0; fi77 < 5; fi77++ {
                switch v6 {
                case "hello":
                    v18 = (^(v18))
                default:
                    switch (v6 + ("\n\t" + ("hello" + ((v6 + "\n\t") + v6)))) {
                    case "hello":
                        if (-128 == 1) {
                            v78 := v5[158]
                            _ = v78
                        }
                    }
                }
                switch -1 {
                case 32768:
                    v5[(fi77 - fi77)] = true
                case 260:
                    var v79 uint8 = 207
                    _ = v79
                }
                _ = fi77
            }
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v80 int32 = 2147483647
    v15 = v15
    var v81 int8 = 42
    for fi82 := 0; fi82 < 4; fi82++ {
        if (false || ((-2.25 - 0.1) >= (((1e10 - 0.0) * (-2.25 - -2.25)) + 1.5))) {
            v83 := GenS5{F0: (v9 >> ((557) & 63)), F1: (v12 && false)}
            v5[(v10 & (v1 % ((v10) | 1)))] = v12
            _ = v83
        }
        var v84 uint16 = 0
        _ = fi82
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v81
    _ = v9
}
