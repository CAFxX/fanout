package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint; F2 int64 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8, p1 int8, p2 uint32) uint64 {
    p1 -= -9
    if (int(p1) == 607) {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 float64 = 1.5
            fmt.Println(p0, p1, v6)
            var v7 int16 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v8 := GenS1{F0: (("hello" + "a") + (("" + ("\n\t" + ("" + "\n\t"))) + ("foo bar" + "a"))), F1: -1}
        v9 := GenS1{F0: ("\n\t" + "a"), F1: 1}
        _ = v8
        _ = v9
    }
    if false {
        v10 := []int{-758, -1}
        if false {
            p2 -= p2
        } else {
            _ = len(v10)
        }
        _ = v10
    } else {
        var v11 uint8 = 255
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 103
}
func genFn1(p0 string) uint32 {
    var v12 int = 127
    var v13 int8 = 42
    v14 := []int{(v12 | (v12 - -9223372036854775808)), v12, (v12 >> ((255) & 63))}
    v12 += v12
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 894
    var v3 float32 = (float32(0.1) + 1e10)
    v4 := []uint32{57}
    v3 = (v3 / v3)
    var v15 float64 = 3.14159
    v15 -= ((3.14159 / 1.5) - ((0.1 - 0.0) * float64(0.0)))
    if (int64(v2) > 964) {
        if len(v4) > 0 {
            v4[0] = 848
        }
        if false {
            for fi16 := 0; fi16 < 7; fi16++ {
                if ((285 > 42) && (uint16(v1) <= 255)) {
                    v15 = v15
                    fmt.Println(v15)
                    var v17 float64 = v15
                    _ = v17
                } else {
                    v18 := GenS2{F0: 323}
                    _ = v18
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch (fi16 - 902) {
                    case 105:
                        if len(v4) > 0 {
                            v4[0] = uint32(v2)
                        }
                    case -2147483648:
                        v20 := GenS2{F0: -1}
                        _ = v20
                    case 32767:
                        v21 := GenS1{F0: ("hello" + "\n\t"), F1: 0}
                        _ = v21
                    }
                    var v22 int = int(v1)
                    _ = fi19
                    _ = v22
                }
                for fi23 := 0; fi23 < 12; fi23++ {
                    v4 = append(v4, 42)
                    if false {
                        switch 0 {
                        case -892:
                            switch (fi23 + (fi23 + fi23)) {
                            case 2147483647:
                                switch fi23 {
                                case 1:
                                    if false {
                                        var v24 uint16 = 487
                                        switch (fi23 & (fi23 - 594)) {
                                        case 127:
                                            if len(v4) > 0 {
                                                v4[0] = 65535
                                            }
                                        case 32767:
                                            v25 := GenS1{F0: "\n\t", F1: -1}
                                            _ = v25
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v26 := GenS1{F0: "a", F1: -1}
                                                _ = v26
                                            }
                                        }
                                        _ = v24
                                    } else {
                                        v1 = v1
                                        v4 = append(v4, genFn1(("" + ((("\n\t" + "a") + "foo bar") + ("" + "")))))
                                    }
                                default:
                                    v27 := GenS2{F0: -1}
                                    _ = v27
                                }
                            case 861:
                                v28 := GenS3{F0: 918, F1: (-(fi23))}
                                _ = v28
                            case 802:
                                v29 := GenS3{F0: -1, F1: -411}
                                v30 := GenS2{F0: 42}
                                _ = v29
                                _ = v30
                            }
                        case -9223372036854775808:
                            if true {
                                var v31 bool = true
                                v32 := GenS4{F0: 32767, F1: -1, F2: 0}
                                if v31 {
                                    v33 := GenS4{F0: 65535, F1: 32767, F2: 105}
                                    v15 -= v15
                                    _ = v33
                                } else {
                                    if (v31 || false) {
                                        v2 += (v2 / ((v2) | 1))
                                        switch (fi23 * 388) {
                                        case 378:
                                            v31 = (42 > 100)
                                        case 0:
                                            var v34 int16 = -118
                                            _ = v34
                                        case 9223372036854775807:
                                            fmt.Println(v1)
                                        }
                                        if v31 {
                                            var v35 int = (fi16 / (((fi16 >> ((4294967295) & 63))) | 1))
                                            _ = v35
                                        } else {
                                            v36 := GenS5{F0: 1, F1: 18446744073709551615, F2: 992}
                                            v31 = (!true)
                                            _ = v36
                                        }
                                    } else {
                                        _ = len(v4)
                                        var v37 bool = (v15 != (v15 * float64((v3 / (v3 + v3)))))
                                        _ = v37
                                    }
                                    v31 = ("\n\t" != "foo bar")
                                }
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(fi23)
                                v38 := GenS4{F0: 0, F1: 1000, F2: 78}
                                _ = v38
                            }
                        case 128:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        switch "" {
                        case "a":
                            switch 185 {
                            case -160:
                                v4 = append(v4, 4294967295)
                            default:
                                if false {
                                    _ = len(v4)
                                    switch fi23 {
                                    case -2147483648:
                                        var v39 int = fi16
                                        _ = v39
                                    case 32768:
                                        switch ((("\n\t" + "hello") + "a") + "a") {
                                        case "a":
                                            v40 := GenS6{F0: 660, F1: 566, F2: 2147483647}
                                            _ = v40
                                        case "":
                                            v41 := GenS7{F0: (fi16 & (fi16 & (fi23 - fi23))), F1: -1}
                                            _ = v41
                                        case "hello":
                                            if (1000 < -709) {
                                                v15 += ((0.1 / (((v15 / v15) * (-2.25 + 1.5)) - 1.5)) - (((v15 + 0.0) - v15) * v15))
                                            } else {
                                                v1 += 255
                                            }
                                        default:
                                            var v42 float32 = (0.0 / 1.5)
                                            _ = v42
                                        }
                                    default:
                                        if true {
                                            if true {
                                                v3 = (float32((v15 + v15)) + 3.14159)
                                                var v43 int16 = 32767
                                                _ = v43
                                            }
                                            var v44 uint8 = (^(v1))
                                            v45 := GenS6{F0: 65535, F1: (v2 + genFn0(127, 1, 4294967295)), F2: 0}
                                            v46 := GenS8{F0: (v1 << ((277) & 63)), F1: 0}
                                            var v47 string = (("a" + "") + "foo bar")
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                } else {
                                    fmt.Println(v2, v15, fi16)
                                    fmt.Println(fi16, v3)
                                    var v48 float32 = ((3.14159 + v3) - -2.25)
                                    _ = v48
                                }
                            }
                        case "\n\t":
                            _ = len(v4)
                        case "":
                            if (v15 < ((0.1 / v15) + (1e10 + (float64(v3) - float64((v3 * v3)))))) {
                                fmt.Println(v3, v15, fi16)
                                var v49 bool = false
                                v2 += 969
                                _ = v49
                            } else {
                                if ((1 >= 9223372036854775807) && false) {
                                    var v50 int64 = -15
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v50
                                } else {
                                    v4 = append(v4, 1)
                                    v2 -= v2
                                }
                            }
                        default:
                            if (739 != 32767) {
                                var v51 uint64 = uint64(v1)
                                v1 = v1
                                var v52 uint8 = 255
                                _ = v51
                                _ = v52
                            } else {
                                if false {
                                    fmt.Println(fi23, fi16, v1)
                                    v2 = (v2 ^ 0)
                                } else {
                                    _ = len(v4)
                                    v15 += v15
                                    if false {
                                        v15 = 0.0
                                        switch fi16 {
                                        case 0:
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v53 uint8 = 60
                                                _ = v53
                                            case "foo bar":
                                                v3 = (v3 + (((v3 - v3) + v3) + ((v3 * float32(1.5)) * float32((0.0 * ((((1.5 / v15) + v15) + 1e10) * ((float64((v3 - v3)) - (((v15 - float64((-2.25 + v3))) - v15) * (v15 - 0.0))) - v15)))))))
                                            }
                                        case 32768:
                                            v54 := GenS5{F0: 42, F1: uint(v2), F2: 0}
                                            _ = v54
                                        }
                                        var v55 int8 = 100
                                        _ = v55
                                    }
                                }
                                if (!(("" < (("" + "hello") + "")) && true)) {
                                    var v56 bool = false
                                    var v57 uint16 = 42
                                    _ = v56
                                    _ = v57
                                }
                            }
                        }
                    } else {
                        switch fi16 {
                        case 100:
                            _ = len(v4)
                        case 0:
                            switch (-(fi23)) {
                            case -531:
                                v15 -= (v15 - (v15 - (v15 + v15)))
                            case 127:
                                var v58 int = (fi16 << ((4294967295) & 63))
                                _ = v58
                            }
                        case 4294967295:
                            fmt.Println(fi16, v3, v3)
                        }
                        var v59 uint16 = 255
                        fmt.Println(v1, v2)
                        v59 = v59
                        _ = v59
                    }
                    _ = fi23
                }
                _ = fi16
            }
            v2 -= (v2 + v2)
            v60 := GenS4{F0: 826, F1: 42, F2: -1}
            _ = v60
        }
    } else {
        v61 := GenS3{F0: -977, F1: 472}
        if (false || (v3 == v3)) {
            v15 += ((v15 + -2.25) * float64(1.5))
            for fi62 := 0; fi62 < 13; fi62++ {
                for fi63 := 0; fi63 < 7; fi63++ {
                    v64 := GenS8{F0: 2, F1: -35}
                    _ = fi63
                    _ = v64
                }
                var v65 int = (-(fi62))
                fmt.Println(v2)
                _ = fi62
                _ = v65
            }
        }
        v66 := GenS7{F0: -9223372036854775808, F1: 1000}
        _ = v61
        _ = v66
    }
    var v67 uint16 = 2
    v68 := GenS6{F0: 0, F1: 2, F2: 171}
    fmt.Println(v1, v1)
    _ = len(v4)
    v4 = append(v4, 42)
    var v69 uint8 = (^(v1))
    var v70 int32 = 0
    switch ("hello" + "") {
    case "a":
        var v71 uint16 = (v67 + (v67 << ((928) & 63)))
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
