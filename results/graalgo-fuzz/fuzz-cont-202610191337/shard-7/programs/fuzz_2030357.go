package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int8 }
func genFn0() uint64 {
    switch "hello" {
    case "":
        var v6 uint8 = 47
        _ = v6
    case "\n\t":
        var v7 uint = 1
        _ = v7
    case "a":
        var v8 int = 32767
        _ = v8
    }
    var v9 uint16 = 255
    if true {
        fmt.Println(v9)
        var v10 float64 = -2.25
        _ = v10
    }
    var v11 uint = 0
    _ = v11
    _ = v9
    return 711
}
func genFn1() uint16 {
    var v12 int16 = -833
    var v13 uint8 = 2
    v14 := GenS1{F0: (false && (0 == 0)), F1: v12}
    v15 := []int{-9223372036854775808, 128, 2147483647, 2147483647}
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 2
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 21
    var v3 float64 = 1e10
    var v4 int64 = 9223372036854775807
    var v5 int16 = -1
    if (42 <= 255) {
        switch -9223372036854775808 {
        case 96:
            v2 += v2
        case 0:
            v5 -= v5
        case 4294967295:
            v16 := GenS1{F0: true, F1: (v5 % (((v5 - -461)) | 1))}
            _ = v16
        default:
            v3 = (v3 - 0.1)
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: -75}
            v3 -= 1.5
            v19 := []int{(fi17 << ((65535) & 63)), fi17, fi17}
            _ = fi17
            _ = v18
            _ = v19
        }
        if ((v5 >> ((4294967295) & 63)) < int16(v1)) {
            v3 -= 0.0
        } else {
            fmt.Println(v3)
        }
    }
    var v20 string = ""
    if (477 <= 775) {
        if false {
            v21 := GenS2{F0: -59}
            v2 -= 42
            v22 := GenS1{F0: (!false), F1: v5}
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    if false {
                        var v25 float64 = v3
                        fmt.Println(fi23, v2, v20)
                        _ = v25
                    }
                    v26 := GenS1{F0: false, F1: 131}
                    v27 := GenS2{F0: -1}
                    _ = fi24
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 15; fi28++ {
                    if true {
                        fmt.Println(v20, v4, v1)
                        var v29 string = "\n\t"
                        v30 := GenS2{F0: -128}
                        _ = v29
                        _ = v30
                    } else {
                        if (!true) {
                            var v31 uint8 = 255
                            if (42 >= genFn1()) {
                                v32 := GenS3{F0: (^(v5))}
                                fmt.Println(v20, v2, v4)
                                _ = v32
                            }
                            if (2147483647 >= 2147483648) {
                                v20 = v20
                                if (true || false) {
                                    if (1.5 == -2.25) {
                                        v33 := GenS4{F0: fi23}
                                        _ = v33
                                    }
                                }
                                switch int(v4) {
                                case 560:
                                    var v34 int8 = -128
                                    _ = v34
                                case 127:
                                    v35 := []int{(fi23 + (-(fi28))), fi28, fi28, fi23}
                                    _ = v35
                                default:
                                    v36 := GenS3{F0: (^(v5))}
                                    _ = v36
                                }
                            }
                            _ = v31
                        } else {
                            var v37 bool = false
                            fmt.Println(v1)
                            if v37 {
                                v1 = v1
                                v2 = (^(v2))
                            } else {
                                if v37 {
                                    switch "a" {
                                    case "a":
                                        v38 := []int{32767, fi23}
                                        _ = v38
                                    case "\n\t":
                                        if true {
                                            v2 += 197
                                            var v39 uint32 = 38
                                            v20 = v20
                                            if true {
                                                v40 := []int{42, (fi23 + (fi23 << ((255) & 63)))}
                                                if (((16 <= 100) || true) && false) {
                                                    var v41 int32 = v1
                                                    var v42 bool = (0 != (^(v2)))
                                                    v1 -= v41
                                                    var v43 bool = (v1 < v1)
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = len(v40)
                                                _ = v40
                                            }
                                            v44 := []int{-9223372036854775808, fi23}
                                            _ = v39
                                            _ = v44
                                        } else {
                                            var v45 bool = ((!v37) && false)
                                            _ = v45
                                        }
                                    case "":
                                        var v46 string = v20
                                        _ = v46
                                    }
                                    var v47 int64 = (v4 - 0)
                                    _ = v47
                                }
                                switch "foo bar" {
                                case "a":
                                    var v48 string = v20
                                    _ = v48
                                default:
                                    fmt.Println(v1)
                                }
                                if true {
                                    var v49 uint8 = 127
                                    _ = v49
                                } else {
                                    var v50 int8 = 14
                                    v5 -= (v5 | (v5 * (-(v5))))
                                    fmt.Println(v5, v5, v4)
                                    v51 := []int{fi23, (fi23 % (((fi28 / (((fi28 | (fi28 << ((4294967295) & 63)))) | 1))) | 1)), fi23}
                                    var v52 int = 127
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                }
                            }
                            var v53 uint32 = 1
                            v54 := GenS2{F0: -128}
                            _ = v37
                            _ = v53
                            _ = v54
                        }
                        v55 := GenS1{F0: false, F1: -632}
                        switch v20 {
                        case "hello":
                            v56 := []int{(fi28 ^ 2147483647), 671, (^(fi28)), (fi23 << ((575) & 63))}
                            _ = v56
                        default:
                            var v57 uint64 = 465
                            _ = v57
                        }
                        var v58 int = fi23
                        v5 = int16(v4)
                        v4 = 42
                        v1 -= v1
                        _ = v55
                        _ = v58
                    }
                    if (127 != -33) {
                        var v59 int = (fi23 & fi28)
                        v3 = v3
                        if (false || (float32(1.5) == 1.5)) {
                            if (v5 < (v5 * v5)) {
                                v60 := GenS5{F0: (760 != 2), F1: genFn0(), F2: genFn1()}
                                var v61 uint32 = 856
                                v62 := GenS6{F0: -58}
                                v63 := []int{(-(fi28)), v59, 9223372036854775807, v59}
                                if ((v2 - v2) < (^(v2))) {
                                    var v64 uint8 = v2
                                    fmt.Println(v64)
                                    _ = v64
                                }
                                _ = v60
                                _ = v61
                                _ = v62
                                _ = v63
                            } else {
                                switch "" {
                                case "foo bar":
                                    if (!(genFn1() != 42)) {
                                        v65 := GenS1{F0: false, F1: -1}
                                        var v66 uint64 = 2
                                        if true {
                                            v67 := []int{fi23, fi28, (fi23 - fi23)}
                                            if len(v67) > 0 {
                                                v67[0] = v59
                                            }
                                            _ = v67
                                        } else {
                                            v68 := []int{int(v66)}
                                            var v69 float32 = 0.1
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = v65
                                        _ = v66
                                    }
                                case "a":
                                    var v70 float64 = 3.14159
                                    _ = v70
                                default:
                                    fmt.Println(v59, v4, v59)
                                    if true {
                                        var v71 uint32 = 433
                                        var v72 uint32 = 804
                                        var v73 uint64 = 0
                                        v73 = v73
                                        v2 = 100
                                        v74 := []int{(fi28 | (fi23 ^ (fi28 * v59)))}
                                        _ = len(v74)
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    var v75 int8 = -1
                                    _ = v75
                                }
                                var v76 bool = (!false)
                                _ = v76
                            }
                            v2 += v2
                            v77 := []int{128}
                            _ = v77
                        } else {
                            var v78 int32 = (v1 / (((v1 - (v1 - (-(v1))))) | 1))
                            if (!(!false)) {
                                if false {
                                    v1 = (v78 >> ((2) & 63))
                                    v78 = (v78 << ((255) & 63))
                                    var v79 int64 = v4
                                    _ = v79
                                }
                                v20 = v20
                                v59 -= fi23
                            } else {
                                v3 += 3.14159
                                v80 := GenS7{F0: genFn1(), F1: -128}
                                v59 -= (-(fi28))
                                _ = v80
                            }
                            _ = v78
                        }
                        _ = v59
                    }
                    _ = fi28
                }
                _ = fi23
            }
            var v81 float32 = (((float32((((((0.1 / (v3 * v3)) - v3) - v3) + (((v3 * (v3 - 0.1)) - v3) * v3)) * v3)) - ((0.0 / (1.5 / 1.0)) - (0.1 * (-2.25 - 0.0)))) * (1e10 + 0.1)) + 1.5)
            _ = v81
        }
        v5 = 0
    }
    fmt.Println(v2)
    v82 := []int{-2147483648}
    for fi83 := 0; fi83 < 3; fi83++ {
        v1 = (v1 & (-(v1)))
        _ = fi83
    }
    _ = len(v82)
    for fi84 := 0; fi84 < 1; fi84++ {
        var v85 uint = 502
        _ = fi84
        _ = v85
    }
    v82 = append(v82, 404)
    v86 := GenS5{F0: true, F1: 0, F2: 394}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v82
    _ = v86
}
