package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
func genFn0(p0 uint16, p1 int, p2 int64) int8 {
    var v6 bool = (p0 >= p0)
    switch ("foo bar" + "a") {
    case "\n\t":
        var v7 uint32 = 65535
        _ = v7
    default:
        switch "foo bar" {
        case "\n\t":
            var v8 uint64 = 18446744073709551615
            _ = v8
        default:
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 125
}
func genFn1(p0 uint, p1 uint, p2 uint) uint8 {
    v9 := []int{}
    for fi10 := 0; fi10 < 3; fi10++ {
        if false {
            v9 = append(v9, (fi10 | (fi10 << ((p0) & 63))))
            var v11 uint = p1
            switch (("" + ("foo bar" + ("foo bar" + "a"))) + "foo bar") {
            case "foo bar":
                var v12 int16 = -32768
                _ = v12
            case "a":
                var v13 int8 = genFn0(1, (fi10 << ((p1) & 63)), 42)
                _ = v13
            default:
                v14 := GenS1{F0: 42}
                _ = v14
            }
            _ = v11
        }
        _ = fi10
    }
    p2 += (p0 ^ (p2 - p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 uint8 = 195
    var v3 int16 = 397
    var v4 uint32 = 2147483647
    var v5 float64 = (0.0 / 1.0)
    var v15 uint8 = (^(v2))
    var v16 bool = false
    if v16 {
        switch -427 {
        case 4294967295:
            if (710 != 4294967295) {
                v17 := GenS1{F0: 42}
                _ = v17
            } else {
                var v18 int8 = v1
                _ = v18
            }
        case -1:
            switch 128 {
            case -982:
                var v19 uint = 255
                _ = v19
            case 32767:
                switch -300 {
                case -1:
                    v20 := GenS1{F0: -1}
                    _ = v20
                default:
                    if true {
                        v3 = v3
                        v15 += (^(v15))
                        v21 := GenS1{F0: 1}
                        _ = v21
                    }
                }
            }
        }
        fmt.Println(v5)
    } else {
        v15 += uint8(v4)
    }
    v4 += v4
    var v22 int64 = -1
    var v23 uint = 0
    for fi24 := 0; fi24 < 14; fi24++ {
        v22 = (v22 - (v22 | 80))
        if (fi24 < (fi24 * (fi24 ^ fi24))) {
            switch 32767 {
            case 4294967295:
                for fi25 := 0; fi25 < 3; fi25++ {
                    if true {
                        var v26 int64 = v22
                        switch ("" + "\n\t") {
                        case "\n\t":
                            v27 := GenS2{F0: uint16(v4), F1: (^(v26)), F2: (v15 % (((v15 % ((v15) | 1))) | 1))}
                            _ = v27
                        }
                        v28 := []int{}
                        _ = v26
                        _ = v28
                    }
                    _ = fi25
                }
            case 362:
                if false {
                    var v29 int = (fi24 | 128)
                    v2 -= v2
                    if true {
                        fmt.Println(v22, v3, fi24)
                        for fi30 := 0; fi30 < 8; fi30++ {
                            v31 := GenS2{F0: 0, F1: (^(v22)), F2: genFn1((v23 * (v23 ^ v23)), v23, 444)}
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        if (v16 || ((!(65535 != 1)) || v16)) {
                            v32 := GenS1{F0: 0}
                            if v16 {
                                var v33 int16 = (v3 >> ((v23) & 63))
                                fmt.Println(v15)
                                _ = v33
                            } else {
                                v4 += v4
                            }
                            var v34 int8 = (v1 / ((42) | 1))
                            _ = v32
                            _ = v34
                        } else {
                            if ((v4 / ((0) | 1)) >= 903) {
                                var v35 float64 = 1e10
                                _ = v35
                            }
                            if (float64(float32(v5)) >= v5) {
                                v1 += 100
                            }
                        }
                        v36 := []int{(fi24 % ((170) | 1)), (fi24 * v29), fi24, -9223372036854775808}
                        _ = v36
                    }
                    _ = v29
                } else {
                    if v16 {
                        var v37 uint = (^(v23))
                        var v38 string = "\n\t"
                        fmt.Println(v5)
                        _ = v37
                        _ = v38
                    } else {
                        fmt.Println(v5, v1, v3)
                        switch 4294967295 {
                        case 745:
                            v22 = (v22 - int64(v1))
                        case 9223372036854775807:
                            var v39 uint8 = 127
                            _ = v39
                        case 1:
                            v3 -= (v3 - (v3 * (^(v3))))
                        }
                        var v40 int = fi24
                        _ = v40
                    }
                    fmt.Println(v22, v22)
                }
            }
            if v16 {
                if v16 {
                    fmt.Println(fi24)
                    fmt.Println(v2)
                }
                for fi41 := 0; fi41 < 18; fi41++ {
                    v4 += v4
                    v23 = v23
                    v1 -= v1
                    _ = fi41
                }
                switch fi24 {
                case 2147483647:
                    v5 = 0.1
                }
            }
            v16 = (("a" == "") && v16)
        } else {
            for fi42 := 0; fi42 < 1; fi42++ {
                if ((2 <= 42) || v16) {
                    if (true || (true || (0 == uint8(v4)))) {
                        var v43 int64 = -48
                        _ = v43
                    }
                    var v44 int = 4294967295
                    _ = v44
                }
                _ = fi42
            }
            switch int(v2) {
            case -769:
                var v45 uint32 = v4
                _ = v45
            case -778:
                switch "\n\t" {
                case "\n\t":
                    if (v16 && v16) {
                        v46 := []int{fi24, 2147483647, fi24}
                        v2 += genFn1(v23, (v23 | (v23 & (v23 << ((v23) & 63)))), v23)
                        fmt.Println(v5, v2, v3)
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 16; fi47++ {
                            var v48 float64 = v5
                            _ = fi47
                            _ = v48
                        }
                        var v49 float32 = (3.14159 * ((float32((1.5 + v5)) - ((float32(v5) * 0.1) * 0.1)) - 0.1))
                        switch "a" {
                        case "hello":
                            if (!v16) {
                                v50 := []int{fi24}
                                if v16 {
                                    if false {
                                        switch fi24 {
                                        case 128:
                                            v49 = (v49 + v49)
                                        default:
                                            v15 = v2
                                        }
                                        var v51 bool = (!v16)
                                        switch fi24 {
                                        case 4294967295:
                                            var v52 uint8 = uint8(v4)
                                            _ = v52
                                        case 142:
                                            switch "a" {
                                            case "foo bar":
                                                switch 128 {
                                                case -620:
                                                    v5 += float64(0.0)
                                                case 128:
                                                    if v51 {
                                                        if false {
                                                            v53 := GenS2{F0: 1, F1: (v22 / ((v22) | 1)), F2: v2}
                                                            var v54 int = fi24
                                                            for fi55 := 0; fi55 < 13; fi55++ {
                                                                var v56 int64 = (v22 + (v22 | v22))
                                                                _ = fi55
                                                                _ = v56
                                                            }
                                                            _ = v53
                                                            _ = v54
                                                        } else {
                                                            v2 = 127
                                                            var v57 float64 = (float64(v49) - 1e10)
                                                            _ = v57
                                                        }
                                                    } else {
                                                        switch fi24 {
                                                        case 32768:
                                                            if (((v49 - (3.14159 / v49)) <= (1e10 * (v49 / v49))) || v16) {
                                                                if (!(!(v16 || ((true && false) || v16)))) {
                                                                    v22 += int64(v1)
                                                                    var v58 uint = (v23 & uint(v2))
                                                                    _ = v58
                                                                }
                                                            } else {
                                                                var v59 float64 = v5
                                                                for fi60 := 0; fi60 < 8; fi60++ {
                                                                    v50 = append(v50, (fi24 * 32767))
                                                                    switch fi60 {
                                                                    case 638:
                                                                        if (!v16) {
                                                                            v4 += v4
                                                                            var v61 uint32 = v4
                                                                            switch fi60 {
                                                                            case -849:
                                                                                v62 := GenS3{F0: 65535, F1: (-(v22)), F2: (v15 - (v2 / (((v2 ^ 84)) | 1)))}
                                                                                _ = v62
                                                                            case 42:
                                                                                var v63 int = fi60
                                                                                _ = v63
                                                                            }
                                                                            if len(v50) > 0 {
                                                                                v50[0] = (fi60 << ((42) & 63))
                                                                            }
                                                                            _ = v61
                                                                        }
                                                                    case 32768:
                                                                        switch 42 {
                                                                        case 0:
                                                                            var v64 bool = false
                                                                            _ = v64
                                                                        case 4294967295:
                                                                            v51 = v51
                                                                        default:
                                                                            fmt.Println(fi24, v51, v2)
                                                                        }
                                                                    }
                                                                    _ = fi60
                                                                }
                                                                _ = v59
                                                            }
                                                        case 42:
                                                            v4 -= uint32(v2)
                                                        case -857:
                                                            v51 = (!(!v16))
                                                        default:
                                                            v1 += (v1 * (-(v1)))
                                                        }
                                                        if len(v50) > 0 {
                                                            v50[0] = fi24
                                                        }
                                                    }
                                                default:
                                                    v16 = (!((v3 * -488) != v3))
                                                }
                                            case "":
                                                var v65 string = "a"
                                                _ = v65
                                            }
                                        case 32767:
                                            if len(v50) > 0 {
                                                v50[0] = 42
                                            }
                                        default:
                                            if len(v50) > 0 {
                                                v50[0] = int(v1)
                                            }
                                        }
                                        _ = v51
                                    } else {
                                        fmt.Println(v4, v2, v22)
                                    }
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        if (v16 && ((v5 < float64(v49)) && (!v16))) {
                                            var v67 uint = v23
                                            fmt.Println(v2, v67)
                                            v23 = (^(v67))
                                            _ = v67
                                        }
                                        _ = fi66
                                    }
                                    fmt.Println(v1, v5, v5)
                                }
                                var v68 uint64 = 52
                                _ = v50
                                _ = v68
                            }
                        case "":
                            v69 := GenS2{F0: 65535, F1: 42, F2: 75}
                            _ = v69
                        }
                        _ = v49
                    }
                case "":
                    v70 := GenS3{F0: v23, F1: int64(fi24), F2: (^(v15))}
                    _ = v70
                case "foo bar":
                    if (133 >= v15) {
                        var v71 int8 = genFn0(612, fi24, (v22 - int64(v23)))
                        _ = v71
                    } else {
                        var v72 int16 = v3
                        v73 := []int{}
                        _ = v72
                        _ = v73
                    }
                default:
                    var v74 int16 = v3
                    _ = v74
                }
            case 0:
                if v16 {
                    v75 := []int{(fi24 * fi24), (fi24 >> ((v23) & 63))}
                    switch "foo bar" {
                    case "foo bar":
                        if len(v75) >= 2 {
                            v75 = v75[1:len(v75)]
                        }
                    case "hello":
                        v76 := GenS4{F0: v3}
                        _ = v76
                    case "":
                        var v77 int16 = v3
                        _ = v77
                    }
                    _ = v75
                }
            }
            var v78 uint = v23
            _ = v78
        }
        switch "hello" {
        case "hello":
            for fi79 := 0; fi79 < 13; fi79++ {
                var v80 uint = (v23 + v23)
                _ = fi79
                _ = v80
            }
        case "foo bar":
            if (1e10 <= 1.5) {
                var v81 int16 = (v3 * v3)
                var v82 uint64 = 899
                _ = v81
                _ = v82
            }
        }
        _ = fi24
    }
    v22 = 0
    v83 := GenS3{F0: v23, F1: 9223372036854775807, F2: (v2 >> ((v23) & 63))}
    var v84 float64 = float64(0.0)
    for fi85 := 0; fi85 < 12; fi85++ {
        for fi86 := 0; fi86 < 18; fi86++ {
            v87 := []int{(fi86 << (((v23 | (v23 | (v23 + 199)))) & 63)), (fi85 >> ((558) & 63)), 9223372036854775807}
            _ = fi86
            _ = v87
        }
        if v16 {
            v16 = v16
            v88 := GenS5{F0: "foo bar", F1: (v22 + (^(v22))), F2: 0}
            _ = v88
        }
        _ = fi85
    }
    var v89 float64 = v84
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v89
}
