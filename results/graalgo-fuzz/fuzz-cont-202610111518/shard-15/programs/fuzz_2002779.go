package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int16; F1 int64; F2 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int; F2 uint }
type GenS9 struct { F0 bool; F1 string; F2 uint }
type GenS10 struct { F0 uint8 }
func genFn0() int8 {
    var v6 bool = true
    if (!false) {
        var v7 int8 = 42
        for fi8 := 0; fi8 < 20; fi8++ {
            if v6 {
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int64 = -9223372036854775808
                    v11 := []int{(fi8 % (((fi8 | fi8)) | 1)), (fi9 % ((42) | 1)), 2147483647, fi9}
                    v12 := GenS1{F0: 65535, F1: 1}
                    v13 := GenS1{F0: 42, F1: 255}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 uint = 0
                    var v16 int8 = -1
                    fmt.Println(v7, fi14)
                    if v6 {
                        v17 := []int{fi14, (fi8 / ((-9223372036854775808) | 1)), 127}
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            v18 := GenS1{F0: 32767, F1: 255}
            _ = fi8
            _ = v18
        }
        fmt.Println(v7, v6)
        _ = v7
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            if true {
                v6 = (v6 && v6)
                if (fi19 <= fi19) {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 uint8 = 0
                        v6 = ((v6 && v6) && ((1e10 - 1e10) > float32(((0.0 / 1e10) - 0.0))))
                        _ = fi20
                        _ = v21
                    }
                    var v22 bool = ("hello" != "a")
                    v6 = v6
                    _ = v22
                } else {
                    var v23 uint16 = 0
                    for fi24 := 0; fi24 < 15; fi24++ {
                        fmt.Println(fi19, fi19)
                        if v6 {
                            v6 = false
                            switch "hello" {
                            case "hello":
                                fmt.Println(v23, fi19)
                            }
                            fmt.Println(v23, v6, fi19)
                        }
                        _ = fi24
                    }
                    fmt.Println(fi19, fi19, v6)
                    _ = v23
                }
            } else {
                var v25 uint8 = 127
                _ = v25
            }
            v26 := GenS2{F0: 2, F1: 2, F2: 255}
            var v27 int = (fi19 / (((^(fi19))) | 1))
            _ = fi19
            _ = v26
            _ = v27
        }
        if ((0 != 255) || v6) {
            if (2 >= 2) {
                var v28 string = ("foo bar" + "hello")
                _ = v28
            } else {
                v29 := []int{42, 42}
                v30 := GenS1{F0: 42, F1: 28}
                _ = v29
                _ = v30
            }
            var v31 uint32 = 2
            switch "" {
            case "foo bar":
                v32 := GenS3{F0: 42}
                _ = v32
            case "":
                var v33 uint = 42
                _ = v33
            case "a":
                var v34 bool = v6
                _ = v34
            }
            _ = v31
        }
    }
    _ = v6
    return -44
}
func genFn1() uint64 {
    if true {
        v35 := GenS1{F0: 997, F1: 2}
        _ = v35
    } else {
        switch "a" {
        case "a":
            v36 := GenS4{F0: true, F1: 2}
            _ = v36
        case "":
            var v37 int32 = 42
            _ = v37
        case "\n\t":
            var v38 string = "\n\t"
            _ = v38
        }
        var v39 int16 = -857
        if ((42 < 18446744073709551615) && (2147483647 != 2147483648)) {
            if ((1e10 > ((0.1 / -2.25) + 1.5)) || false) {
                if false {
                    v40 := GenS5{F0: v39, F1: 1, F2: 2}
                    var v41 uint32 = 57
                    _ = v40
                    _ = v41
                }
            } else {
                v39 = v39
                fmt.Println(v39)
                var v42 int32 = 642
                _ = v42
            }
            v43 := []int{961}
            var v44 float32 = (0.0 + (0.0 * 0.1))
            _ = v43
            _ = v44
        } else {
            fmt.Println(v39)
        }
        _ = v39
    }
    var v45 int16 = 42
    _ = v45
    return 417
}

func main() {
    var v1 uint64 = 738
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    v5 := []uint32{4294967295, 916}
    v46 := v3[("hello" + ("" + "hello"))]
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 uint8 = 127
        for fi49 := 0; fi49 < 16; fi49++ {
            v50 := GenS2{F0: v48, F1: v48, F2: (v48 + (v48 / ((2) | 1)))}
            _ = len(v5)
            if false {
                var v51 int = 70
                _ = v51
            } else {
                var v52 int64 = 42
                v46 -= fi47
                _ = v52
            }
            if (0 <= -143) {
                if ((3.14159 <= ((3.14159 - 1e10) * -2.25)) || false) {
                    v46 += 32767
                    if (0 > 1000) {
                        if (1e10 >= 3.14159) {
                            fmt.Println(v1)
                        }
                        if false {
                            if ((1e10 + 3.14159) > 1.5) {
                                v53 := GenS3{F0: 432}
                                v54 := GenS6{F0: true}
                                switch ("hello" + "foo bar") {
                                case "":
                                    v55 := GenS3{F0: 64}
                                    if (-2.25 < 1.5) {
                                        var v56 int64 = -282
                                        _ = len(v4)
                                        delete(v3, "foo bar")
                                        _ = v56
                                    } else {
                                        var v57 uint64 = 1
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v57
                                    }
                                    _ = v55
                                default:
                                    var v58 int64 = 1
                                    _ = v58
                                }
                                fmt.Println(v46)
                                _ = v53
                                _ = v54
                            }
                            var v59 int64 = 9223372036854775807
                            _ = v59
                        }
                        var v60 string = "a"
                        _ = v60
                    }
                } else {
                    v61 := GenS3{F0: 289}
                    if (((("a" + "foo bar") + "a") + "hello") >= (("hello" + "hello") + ("" + ("" + "\n\t")))) {
                        fmt.Println(v46, v2, v1)
                        var v62 int32 = 2147483647
                        if false {
                            v63 := GenS3{F0: 32767}
                            fmt.Println(v46)
                            v62 += v62
                            _ = v63
                        }
                        _ = v62
                    }
                    v64 := GenS5{F0: -32768, F1: -1, F2: genFn1()}
                    _ = v61
                    _ = v64
                }
                var v65 uint64 = v2
                v66 := GenS4{F0: ((((1e10 + (1.5 + 0.1)) - 0.0) - float64(((((((-2.25 * 3.14159) + -2.25) / 3.14159) - ((1.5 + 0.1) * (0.0 / 1e10))) - -2.25) * 1e10))) > 1.5), F1: (v48 << ((4294967295) & 63))}
                _ = v65
                _ = v66
            }
            v1 += 0
            if true {
                var v67 uint32 = 1
                _ = v67
            } else {
                v68 := v4[231]
                v69 := GenS7{F0: 506}
                v70 := GenS8{F0: 0, F1: (fi47 >> ((159) & 63)), F2: 711}
                v46 = 219
                _ = v68
                _ = v69
                _ = v70
            }
            _ = fi49
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        switch (fi71 ^ v46) {
        case -2147483648:
            var v72 int32 = 0
            _ = v72
        case 984:
            _ = len(v5)
        }
        fmt.Println(v2)
        _ = fi71
    }
    v73 := GenS9{F0: (!(294 == 0)), F1: ("" + ""), F2: 18446744073709551615}
    var v74 int8 = genFn0()
    v5 = append(v5, 0)
    v75 := GenS4{F0: (((!true) || (!true)) || false), F1: 165}
    v76 := GenS8{F0: genFn1(), F1: (v46 & (v46 % ((v46) | 1))), F2: 18446744073709551615}
    v77 := GenS10{F0: uint8(v1)}
    for fi78 := 0; fi78 < 18; fi78++ {
        var v79 uint64 = (^(v2))
        if (v79 > (v79 + (v2 * (v2 & (v2 | (v2 ^ v2)))))) {
            v2 -= v79
        } else {
            var v80 float32 = 1e10
            if (false && true) {
                for fi81 := 0; fi81 < 5; fi81++ {
                    _ = len(v5)
                    fmt.Println(v74)
                    v74 = (v74 - genFn0())
                    var v82 uint = 2
                    _ = fi81
                    _ = v82
                }
                v1 += v79
                var v83 int16 = 1
                _ = v83
            }
            _ = v80
        }
        _ = fi78
        _ = v79
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
