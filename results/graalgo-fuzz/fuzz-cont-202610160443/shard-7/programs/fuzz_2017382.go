package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int8; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 int32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int64; F1 int8; F2 uint16 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0() uint32 {
    v4 := []int{838, 748, 32767}
    v5 := GenS1{F0: 0}
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "foo bar"
    v3 := make(map[int]bool)
    v1 += v1
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: -32768}
        _ = fi6
        _ = v7
    }
    var v8 bool = (float32((((0.0 + 3.14159) + 1e10) * 1e10)) == float32(3.14159))
    if (!(!false)) {
        v1 = (^(v1))
    } else {
        if v8 {
            v9 := v3[-9223372036854775808]
            var v10 uint64 = 442
            _ = v10
            _ = v9
        } else {
            v2 = ((v2 + v2) + v2)
            var v11 uint16 = 117
            _ = v11
        }
        var v12 float64 = (-2.25 - 0.1)
        _ = v12
    }
    if v8 {
        if (!(v1 == (-(v1)))) {
            fmt.Println(v1, v8)
            if false {
                var v13 float32 = float32(((0.1 + (-2.25 * (((1e10 - 0.1) * 1.5) / ((((-2.25 * (1e10 * 1e10)) * 0.0) + (-2.25 / float64(0.1))) + ((1.5 * 0.1) + (1.5 - float64(0.1))))))) - 1e10))
                v13 += (1e10 / (((float32((0.0 + -2.25)) / 1.5) * v13) / float32(3.14159)))
                for fi14 := 0; fi14 < 4; fi14++ {
                    v1 -= v1
                    _ = fi14
                }
                _ = v13
            }
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                fmt.Println(fi15, v8)
                if (true && v8) {
                    fmt.Println(v2, fi15)
                    for fi16 := 0; fi16 < 9; fi16++ {
                        var v17 uint64 = 2
                        fmt.Println(v17, v1)
                        fmt.Println(v1, v1, v17)
                        _ = fi16
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    v1 += (v1 % (((v1 + 546)) | 1))
                    _ = fi18
                }
                _ = fi15
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint64 = 1
                _ = fi19
                _ = v20
            }
            fmt.Println(v1, v2, v1)
        }
        if v8 {
            if false {
                if (!true) {
                    v1 += (v1 / (((v1 + v1)) | 1))
                }
                switch v2 {
                case "foo bar":
                    v21 := GenS2{F0: 770, F1: 42}
                    _ = v21
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v22 int = 64
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                v1 -= (^(v1))
                if (750 > 514) {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        switch "" {
                        case "a":
                            v8 = v8
                        case "hello":
                            v8 = (194 == 0)
                        case "\n\t":
                            switch 0 {
                            case 127:
                                switch fi23 {
                                case -9223372036854775808:
                                    if (!(((v1 | -1) == (v1 / ((-935) | 1))) || false)) {
                                        if v8 {
                                            v25 := v3[(fi24 * fi24)]
                                            var v26 int64 = (v1 / ((-9223372036854775808) | 1))
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    var v27 uint32 = 614
                                    fmt.Println(v27, fi24)
                                    _ = v27
                                case -2147483648:
                                    switch (fi24 | (fi24 - (fi24 ^ fi24))) {
                                    case 32768:
                                        v2 = (v2 + "a")
                                    case -414:
                                        v28 := GenS3{F0: 76, F1: 0, F2: -2147483648}
                                        _ = v28
                                    case -256:
                                        v2 = (v2 + "\n\t")
                                    default:
                                        var v29 uint32 = 4294967295
                                        _ = v29
                                    }
                                case 9223372036854775807:
                                    if ((1 != -758) && (1 != 0)) {
                                        if (genFn0() != 65535) {
                                            fmt.Println(v2, v1)
                                            var v30 float32 = 1e10
                                            _ = v30
                                        } else {
                                            v1 = v1
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            var v31 int32 = 1
                                            _ = v31
                                        case "\n\t":
                                            v2 = v2
                                        default:
                                            v2 = v2
                                        }
                                    } else {
                                        fmt.Println(fi24, fi24)
                                    }
                                }
                            default:
                                v32 := []int{fi23, (fi24 << ((309) & 63)), 0}
                                _ = v32
                            }
                        }
                        _ = fi24
                    }
                    fmt.Println(fi23)
                    var v33 int8 = 1
                    _ = v33
                } else {
                    delete(v3, (fi23 % ((fi23) | 1)))
                }
                if v8 {
                    v34 := GenS4{F0: 42, F1: -1, F2: int32(fi23)}
                    if v8 {
                        v1 = v1
                    }
                    v35 := []int{fi23, (fi23 / (((fi23 << ((65535) & 63))) | 1))}
                    _ = v34
                    _ = v35
                } else {
                    if true {
                        if v8 {
                            if (v8 || v8) {
                                var v36 int64 = v1
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    var v38 uint8 = 128
                                    var v39 int8 = 0
                                    var v40 uint16 = 888
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                switch (fi23 % (((fi23 + (fi23 + (-(fi23))))) | 1)) {
                                case -1:
                                    if (81 == 53) {
                                        var v41 uint8 = 255
                                        delete(v3, fi23)
                                        _ = v41
                                    } else {
                                        v42 := GenS5{F0: 213, F1: 42}
                                        _ = v42
                                    }
                                case 32768:
                                    v43 := []int{}
                                    _ = v43
                                }
                            }
                            _ = len(v3)
                        } else {
                            if true {
                                var v44 float64 = 0.0
                                var v45 bool = v8
                                _ = v44
                                _ = v45
                            }
                            v8 = (!false)
                            v8 = v8
                        }
                        if (!((-316 < 0) || false)) {
                            var v46 int64 = (v1 % ((v1) | 1))
                            _ = v46
                        } else {
                            for fi47 := 0; fi47 < 15; fi47++ {
                                fmt.Println(v8, v1)
                                v48 := GenS6{F0: v1, F1: 127, F2: 410}
                                v2 = (v2 + ("foo bar" + (v2 + v2)))
                                _ = fi47
                                _ = v48
                            }
                            var v49 uint = 42
                            var v50 uint64 = 935
                            _ = v49
                            _ = v50
                        }
                        switch fi23 {
                        case -9223372036854775808:
                            var v51 uint32 = 78
                            _ = v51
                        case -778:
                            switch (fi23 | fi23) {
                            case -604:
                                v52 := GenS1{F0: -121}
                                _ = v52
                            case 0:
                                switch v2 {
                                case "":
                                    var v53 int8 = -1
                                    _ = v53
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    var v54 int = fi23
                                    _ = v54
                                }
                            case 32767:
                                v55 := GenS7{F0: v8, F1: ""}
                                _ = v55
                            }
                        default:
                            for fi56 := 0; fi56 < 18; fi56++ {
                                v57 := GenS5{F0: -1, F1: 0}
                                v1 = v1
                                _ = fi56
                                _ = v57
                            }
                        }
                    } else {
                        var v58 bool = (255 > 2)
                        v1 = 288
                        fmt.Println(v2)
                        _ = v58
                    }
                    var v59 uint8 = 92
                    _ = v59
                }
                _ = fi23
            }
        }
    } else {
        for fi60 := 0; fi60 < 19; fi60++ {
            v61 := GenS6{F0: 9223372036854775807, F1: 0, F2: 255}
            var v62 uint8 = 51
            v1 += (^(v1))
            _ = fi60
            _ = v61
            _ = v62
        }
        v63 := GenS7{F0: (v8 || v8), F1: ((v2 + "") + "")}
        fmt.Println(v1)
        _ = v63
    }
    v64 := GenS3{F0: 100, F1: 688, F2: 1}
    v1 = v1
    v65 := GenS7{F0: v8, F1: v2}
    var v66 uint = 28
    v67 := GenS4{F0: 373, F1: -2147483648, F2: 2147483647}
    if (!v8) {
        v66 += 630
        v68 := v3[int(v1)]
        _ = v68
    } else {
        v69 := GenS8{F0: -970, F1: -1}
        _ = v69
    }
    v70 := GenS9{F0: v1, F1: 1, F2: 694}
    v71 := GenS4{F0: 1, F1: 390, F2: -75}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v66:", v66)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v70
    _ = v71
    _ = v8
}
