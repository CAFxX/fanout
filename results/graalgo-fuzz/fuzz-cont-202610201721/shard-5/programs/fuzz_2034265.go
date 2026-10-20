package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8; F1 int16; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 string }
func genFn0(p0 uint64, p1 int64, p2 int64) uint64 {
    fmt.Println(p2)
    v5 := []int{}
    p2 = (p1 & p2)
    if len(v5) > 0 {
        v5[0] = 32768
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1(p0 uint32, p1 int16, p2 uint8) uint32 {
    p0 = p0
    p0 += 2147483648
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]string)
    var v3 uint32 = 2147483648
    var v4 int32 = 2147483647
    v6 := GenS1{F0: 42, F1: (v1 & 42)}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int8 = -1
        switch "a" {
        case "":
            v9 := GenS1{F0: (^(v8)), F1: int16(v4)}
            _ = v9
        case "foo bar":
            v3 = (v3 & v3)
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: -128, F1: v1}
            v12 := GenS2{F0: 0, F1: v1, F2: -32768}
            fmt.Println(fi10)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v13 float32 = (float32(1e10) / 1e10)
    var v14 string = ("a" + "\n\t")
    fmt.Println(v1, v14)
    switch -9223372036854775808 {
    case 1:
        delete(v2, (v14 + v14))
    }
    var v15 int = -420
    if (1.5 < (float32(0.0) - v13)) {
        var v16 int = -9223372036854775808
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int16 = 441
            v4 = v4
            for fi19 := 0; fi19 < 20; fi19++ {
                if (genFn0(1, -215, 1) < 63) {
                    fmt.Println(v1)
                } else {
                    if ((("\n\t" + v14) + "foo bar") == (("" + v14) + "foo bar")) {
                        if (true && ((v18 & (v1 << ((18446744073709551615) & 63))) == 1000)) {
                            fmt.Println(v13, fi19, v4)
                            fmt.Println(fi17, v18, v14)
                        } else {
                            var v20 bool = false
                            v15 = (fi19 | int(v1))
                            _ = v20
                        }
                        v21 := GenS2{F0: 61, F1: v1, F2: v18}
                        if ((2 <= 83) && false) {
                            if true {
                                var v22 int = fi19
                                v23 := GenS3{F0: (v4 % (((v4 * (v4 | (v4 + (v4 << ((2) & 63)))))) | 1))}
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS3{F0: (v4 >> ((42) & 63))}
                                v25 := GenS4{F0: 65535, F1: 288, F2: 2}
                                _ = v24
                                _ = v25
                            }
                            v18 -= (v18 / ((-1) | 1))
                            v1 += v18
                        }
                        delete(v2, v14)
                        _ = v21
                    } else {
                        if false {
                            v1 -= -1
                        }
                        var v26 int32 = (v4 % (((v4 + v4)) | 1))
                        if (genFn1((v3 - v3), 32767, 1) > v3) {
                            if true {
                                delete(v2, v14)
                                v27 := GenS5{F0: 24, F1: (v3 / ((v3) | 1))}
                                v26 += (v4 >> ((1) & 63))
                                _ = v27
                            } else {
                                v28 := GenS3{F0: v4}
                                v2[v14] = v14
                                v29 := GenS6{F0: 255}
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(v3, v13, v14)
                            v15 = (fi19 ^ (v15 ^ (fi17 + 32768)))
                        } else {
                            var v30 int64 = -477
                            switch "" {
                            case "\n\t":
                                if false {
                                    v1 = 1
                                    switch fi17 {
                                    case 0:
                                        switch fi19 {
                                        case -2147483648:
                                            v1 -= v18
                                        case 0:
                                            v15 = v16
                                        case 9223372036854775807:
                                            var v31 uint32 = v3
                                            _ = v31
                                        }
                                    case 42:
                                        v30 = int64(v3)
                                    }
                                    if ((((((-(v30)) == 42) && true) && ((true && false) || true)) || true) || false) {
                                        v18 += 653
                                    } else {
                                        if false {
                                            if true {
                                                v32 := GenS6{F0: 2}
                                                v33 := []int{fi17}
                                                if true {
                                                    var v34 uint64 = 0
                                                    var v35 int = -9223372036854775808
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v26, v3)
                                                }
                                                _ = v32
                                                _ = v33
                                            } else {
                                                v36 := []int{v16}
                                                var v37 int32 = (v26 + (v4 / (((v4 & (v4 & (v4 ^ (^(v26)))))) | 1)))
                                                _ = v36
                                                _ = v37
                                            }
                                            fmt.Println(v30)
                                            v3 = genFn1(v3, (^(v18)), 118)
                                        }
                                        switch (v16 / (((fi17 % ((fi19) | 1))) | 1)) {
                                        case 0:
                                            var v38 uint32 = genFn1((v3 & 2147483648), v18, 78)
                                            _ = v38
                                        case 32768:
                                            switch fi19 {
                                            case -778:
                                                v14 = (v14 + "a")
                                            case 4294967295:
                                                v2[((("\n\t" + (v14 + v14)) + "") + (v14 + v14))] = v14
                                            default:
                                                v39 := []int{28}
                                                _ = v39
                                            }
                                            v30 -= v30
                                            var v40 float32 = float32(0.0)
                                            v2["a"] = ("a" + "\n\t")
                                            _ = v40
                                        default:
                                            if (!true) {
                                                v41 := []int{v16, (fi19 >> ((954) & 63))}
                                                v42 := GenS7{F0: ((int16(v30) == v1) && false), F1: 0}
                                                switch v14 {
                                                case "\n\t":
                                                    switch int(v30) {
                                                    case 4294967295:
                                                        var v43 uint8 = 255
                                                        _ = v43
                                                    case -347:
                                                        if true {
                                                            _ = len(v41)
                                                        } else {
                                                            var v44 uint8 = 226
                                                            _ = v44
                                                        }
                                                        v26 += v4
                                                    case 2147483647:
                                                        _ = len(v2)
                                                    }
                                                }
                                                _ = v41
                                                _ = v42
                                            } else {
                                                if (!true) {
                                                    fmt.Println(v13, v15, v13)
                                                    if false {
                                                        fmt.Println(fi19, v26, v18)
                                                        if (128 <= 127) {
                                                            switch (("" + ("" + ("foo bar" + "foo bar"))) + "\n\t") {
                                                            case "":
                                                                v4 -= (^(v26))
                                                            default:
                                                                switch v14 {
                                                                case "\n\t":
                                                                    if true {
                                                                        v13 += 0.1
                                                                    }
                                                                case "foo bar":
                                                                    switch v14 {
                                                                    case "a":
                                                                        fmt.Println(fi17)
                                                                    }
                                                                    if true {
                                                                        v16 += v16
                                                                        v30 += (v30 ^ (v30 >> ((18446744073709551615) & 63)))
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v45 bool = true
                                                                            _ = v45
                                                                        case "a":
                                                                            var v46 int8 = 127
                                                                            switch ((v14 + v14) + ("" + "hello")) {
                                                                            case "a":
                                                                                v2[""] = (v14 + v14)
                                                                            case "foo bar":
                                                                                v47 := []int{(fi17 & (v15 & fi17)), 9223372036854775807}
                                                                                _ = v47
                                                                            default:
                                                                                v48 := v2[v14]
                                                                                _ = v48
                                                                            }
                                                                            var v49 uint16 = 1
                                                                            _ = v46
                                                                            _ = v49
                                                                        case "foo bar":
                                                                            v3 = v3
                                                                        }
                                                                    }
                                                                case "":
                                                                    v50 := []int{127, 541}
                                                                    _ = v50
                                                                default:
                                                                    if false {
                                                                        if (-53 != 127) {
                                                                            if true {
                                                                                v16 = (v16 | (v15 ^ v16))
                                                                            }
                                                                            v18 -= v1
                                                                        }
                                                                        if (false && true) {
                                                                            v30 += (v30 / (((v30 - 42)) | 1))
                                                                            if (false || false) {
                                                                                var v51 uint8 = 0
                                                                                _ = v51
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v52 float64 = 1e10
                                                            v53 := GenS8{F0: v3, F1: (v3 + (v3 >> ((1) & 63)))}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        var v54 uint32 = genFn1(0, v1, 127)
                                                        var v55 uint16 = 777
                                                        fmt.Println(v54)
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        v56 := GenS5{F0: 1, F1: v3}
                                                        if (!((v4 / ((v4) | 1)) == 1)) {
                                                            var v57 int8 = -128
                                                            v3 = v3
                                                            _ = v57
                                                        } else {
                                                            fmt.Println(v1, fi17, v13)
                                                            var v58 int32 = (-(v26))
                                                            if true {
                                                                switch v14 {
                                                                case "\n\t":
                                                                    switch v14 {
                                                                    case "\n\t":
                                                                        if (true || true) {
                                                                            v30 += (v30 << ((1) & 63))
                                                                            switch ("foo bar" + v14) {
                                                                            case "foo bar":
                                                                                v59 := GenS2{F0: uint8(v3), F1: -500, F2: 0}
                                                                                _ = v59
                                                                            case "hello":
                                                                                v1 = (v1 * (v18 % ((v18) | 1)))
                                                                            case "a":
                                                                                delete(v2, "hello")
                                                                            }
                                                                            v60 := []int{(v16 / ((-589) | 1))}
                                                                            _ = v60
                                                                        }
                                                                    case "foo bar":
                                                                        switch v14 {
                                                                        case "\n\t":
                                                                            var v61 int64 = v30
                                                                            _ = v61
                                                                        case "":
                                                                            var v62 float64 = (1.5 + (3.14159 + 0.0))
                                                                            _ = v62
                                                                        }
                                                                    case "a":
                                                                        var v63 int8 = 100
                                                                        _ = v63
                                                                    }
                                                                case "foo bar":
                                                                    v64 := GenS6{F0: 255}
                                                                    _ = v64
                                                                case "":
                                                                    var v65 int8 = 42
                                                                    _ = v65
                                                                default:
                                                                    v66 := []int{32768, -1}
                                                                    _ = v66
                                                                }
                                                                v3 = 0
                                                            }
                                                            _ = v58
                                                        }
                                                        v67 := GenS9{F0: 255}
                                                        _ = v56
                                                        _ = v67
                                                    }
                                                }
                                                v68 := []int{fi17}
                                                fmt.Println(v3, fi19)
                                                _ = v68
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v4, fi17, fi19)
                                }
                            }
                            var v69 int32 = 2147483647
                            _ = v30
                            _ = v69
                        }
                        _ = v26
                    }
                    var v70 bool = false
                    var v71 int32 = (v4 | int32(v3))
                    if ((v3 & (v3 << ((42) & 63))) < v3) {
                        var v72 uint = 762
                        _ = v72
                    } else {
                        _ = len(v2)
                        v16 = -2147483648
                        var v73 int = fi19
                        v18 = int16(v4)
                        v73 += -1
                        var v74 bool = (v4 != -1)
                        v75 := GenS10{F0: ((("\n\t" + v14) + v14) + v14)}
                        if v70 {
                            v15 -= int(v71)
                            v76 := []int{-678, v73, (v73 & 32767), (^(v15))}
                            v77 := GenS3{F0: 0}
                            _ = v76
                            _ = v77
                        } else {
                            v16 = (v15 | -1)
                        }
                        if (1 > 667) {
                            fmt.Println(v70, v1, v74)
                            v78 := []int{9223372036854775807, v15, 32767}
                            v4 += v71
                            _ = v78
                        } else {
                            var v79 int8 = 127
                            _ = v79
                        }
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                    _ = v70
                    _ = v71
                }
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        if false {
            var v80 int64 = -465
            v81 := []int{}
            var v82 float64 = 1e10
            _ = v80
            _ = v81
            _ = v82
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v3)
            case "":
                for fi83 := 0; fi83 < 11; fi83++ {
                    var v84 uint64 = 42
                    _ = fi83
                    _ = v84
                }
            }
            v14 = ""
            if (!false) {
                var v85 int16 = 1000
                v13 -= float32(1.5)
                v3 = (v3 & 42)
                _ = v85
            } else {
                v13 += 1e10
                fmt.Println(v4, v3, v16)
            }
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
