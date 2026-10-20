package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16; F1 int; F2 int }
type GenS5 struct { F0 int; F1 int64; F2 uint32 }
func genFn0() int {
    switch 0 {
    case -1:
        if true {
            v5 := GenS1{F0: 0, F1: 972}
            if (700 != 32767) {
                v6 := GenS1{F0: 37, F1: 4294967295}
                switch -9223372036854775808 {
                case 259:
                    var v7 uint8 = 63
                    _ = v7
                case -9223372036854775808:
                    switch ("foo bar" + ("a" + (("a" + "") + ("foo bar" + "")))) {
                    case "a":
                        for fi8 := 0; fi8 < 16; fi8++ {
                            for fi9 := 0; fi9 < 20; fi9++ {
                                if (!false) {
                                    var v10 int = fi8
                                    var v11 uint16 = 705
                                    v12 := GenS1{F0: 255, F1: 42}
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                                _ = fi9
                            }
                            _ = fi8
                        }
                    case "hello":
                        if true {
                            var v13 int32 = -174
                            var v14 uint = 255
                            _ = v13
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 17; fi15++ {
                                fmt.Println(fi15)
                                _ = fi15
                            }
                            var v16 int8 = 1
                            v16 = v16
                            _ = v16
                        }
                    case "\n\t":
                        var v17 float32 = (0.1 / 1.5)
                        _ = v17
                    default:
                        v18 := []int{0, 226, 370}
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := GenS1{F0: 255, F1: 1}
                        if true {
                            var v21 int64 = 42
                            fmt.Println(fi19, v21)
                            _ = v21
                        } else {
                            var v22 float64 = float64((0.1 + 3.14159))
                            var v23 int8 = -128
                            _ = v22
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                _ = v6
            }
            _ = v5
        }
    default:
        var v24 uint32 = 2
        _ = v24
    }
    var v25 uint = 42
    v25 += (v25 | 4294967295)
    _ = v25
    return 138
}

func main() {
    var v1 float32 = 1e10
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    v4 := []int{0}
    fmt.Println(v2)
    var v26 int64 = (v2 * v2)
    v26 = v26
    for fi27 := 0; fi27 < 8; fi27++ {
        if true {
            var v28 uint16 = 1
            v1 = v1
            _ = v28
        } else {
            v1 = float32(3.14159)
            var v29 float32 = (float32((0.1 + 0.0)) * 1e10)
            for fi30 := 0; fi30 < 1; fi30++ {
                if (!(false || true)) {
                    if (!((v2 + -9223372036854775808) <= (v26 % ((v2) | 1)))) {
                        var v31 uint8 = 1
                        var v32 uint32 = 1
                        var v33 uint16 = 353
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int64 = -1
                        if (!(!true)) {
                            v34 += v34
                        } else {
                            var v35 bool = (396 != 1)
                            _ = v35
                        }
                        v36 := GenS2{F0: 127}
                        _ = v34
                        _ = v36
                    }
                    switch 9223372036854775807 {
                    case 32767:
                        v37 := GenS3{F0: true, F1: fi27}
                        _ = len(v4)
                        _ = v37
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "":
                            v38 := v3[(-(fi30))]
                            _ = v38
                        default:
                            var v39 bool = (42 != 42)
                            _ = v39
                        }
                    default:
                        _ = len(v4)
                    }
                }
                if (42 < 191) {
                    v29 += 1.5
                }
                _ = fi30
            }
            _ = v29
        }
        _ = fi27
    }
    if (0 >= 128) {
        if len(v4) > 0 {
            v4[0] = 32767
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi40
        }
    }
    var v41 float64 = ((0.0 + ((-2.25 - 3.14159) * ((-2.25 + 1e10) * 1e10))) + -2.25)
    for fi42 := 0; fi42 < 3; fi42++ {
        v41 = ((float64(3.14159) + (v41 / (v41 / v41))) * v41)
        var v43 int32 = -2147483648
        _ = fi42
        _ = v43
    }
    fmt.Println(v1, v41)
    v44 := v3[-1]
    v44 = (^(v44))
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 int8 = -28
        fmt.Println(v44)
        if false {
            v44 = 42
        } else {
            v26 -= -1
            v41 = v41
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := GenS1{F0: 255, F1: v44}
                var v49 int32 = -2147483648
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        _ = fi45
        _ = v46
    }
    for fi50 := 0; fi50 < 5; fi50++ {
        for fi51 := 0; fi51 < 20; fi51++ {
            v44 = (^(v44))
            _ = fi51
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (v2 != v26) {
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 float32 = v1
                var v54 int32 = 1
                var v55 int8 = 100
                _ = fi52
                _ = v53
                _ = v54
                _ = v55
            }
            for fi56 := 0; fi56 < 9; fi56++ {
                if ("" == "\n\t") {
                    v57 := GenS4{F0: 1, F1: fi56, F2: fi50}
                    _ = v57
                }
                _ = fi56
            }
        } else {
            if (0 < 1000) {
                var v58 uint64 = 2
                _ = v58
            } else {
                switch (^(fi50)) {
                case 42:
                    v4 = append(v4, fi50)
                default:
                    var v59 uint64 = 751
                    _ = v59
                }
                fmt.Println(v41, fi50)
                var v60 float32 = (((-2.25 / ((((v1 / 1.5) + float32(v41)) * v1) - -2.25)) / -2.25) - v1)
                _ = v60
            }
            if false {
                if (3.14159 <= 1e10) {
                    switch fi50 {
                    case 2147483647:
                        _ = len(v4)
                    }
                } else {
                    v26 = v2
                }
                v2 = -1
            }
        }
        _ = fi50
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        v62 := GenS2{F0: 100}
        _ = fi61
        _ = v62
    }
    if (!((v44 * (^(v44))) >= (^(v44)))) {
        v63 := GenS5{F0: 0, F1: v2, F2: 65535}
        v41 -= 1e10
        _ = len(v4)
        _ = v63
    } else {
        var v64 float32 = ((v1 / v1) - ((v1 + -2.25) + -2.25))
        _ = v64
    }
    var v65 uint = 42
    v65 += 434
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v65
}
