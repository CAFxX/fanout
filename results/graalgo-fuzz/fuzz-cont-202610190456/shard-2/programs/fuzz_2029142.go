package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
type GenS4 struct { F0 bool; F1 string; F2 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int, p1 int32) uint64 {
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 13; fi5++ {
        if (4294967295 < 566) {
            v6 := []int{2147483647, fi5, 128, 4294967295}
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS1{F0: 255, F1: 42}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            p0 = -840
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v10 bool = false
                    _ = v10
                case "hello":
                    fmt.Println(fi8, p1, fi8)
                }
                v11 := []int{(fi9 | -519), 42, fi9, (fi9 | (fi8 >> ((2) & 63)))}
                p0 += p0
                _ = v11
            } else {
                v12 := []int{(fi9 * 4294967295), fi9, 127}
                if (-32768 < -699) {
                    if (100 <= -1) {
                        p1 = p1
                        fmt.Println(fi9)
                        fmt.Println(fi9, p1)
                    } else {
                        fmt.Println(v4, fi8)
                        if false {
                            var v13 float64 = 0.0
                            var v14 int64 = 525
                            _ = v13
                            _ = v14
                        } else {
                            var v15 int64 = 0
                            switch fi9 {
                            case 490:
                                v4 = v4
                            case 773:
                                v12 = append(v12, 127)
                            case 32768:
                                p1 -= p1
                            }
                            _ = v15
                        }
                    }
                }
                _ = v12
            }
            _ = fi9
        }
        _ = fi8
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: 42, F1: 72}
            _ = fi17
            _ = v18
        }
        v4 = v4
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 int16 = (-(v1))
    var v3 int = -9223372036854775808
    v19 := []int{v3, (v3 << ((0) & 63)), (v3 / (((v3 & v3)) | 1)), (-(v3))}
    v19 = append(v19, (^(v3)))
    var v20 int32 = 1
    switch "foo bar" {
    case "":
        switch v3 {
        case -262:
            if (!(true && true)) {
                v20 = (v20 * (v20 + v20))
            }
        case 0:
            for fi21 := 0; fi21 < 11; fi21++ {
                if false {
                    v1 += (v2 % (((v1 / (((-(v1))) | 1))) | 1))
                    if false {
                        if false {
                            v3 -= v3
                            if false {
                                if len(v19) > 0 {
                                    v19[0] = (v3 - v3)
                                }
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = fi22
                                }
                            }
                        }
                        switch "a" {
                        case "a":
                            if false {
                                var v23 string = "foo bar"
                                _ = v23
                            } else {
                                v1 -= -1
                                var v24 int8 = 100
                                _ = v24
                            }
                        case "hello":
                            if (genFn0(-1, -1) == 1) {
                                v20 = v20
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            } else {
                                v25 := GenS2{F0: 1}
                                _ = v25
                            }
                        case "\n\t":
                            v2 = v1
                        default:
                            for fi26 := 0; fi26 < 11; fi26++ {
                                v20 += (v20 ^ v20)
                                _ = fi26
                            }
                        }
                        v27 := GenS1{F0: 42, F1: 42}
                        _ = v27
                    }
                    if ((112 > 127) || (v20 < 0)) {
                        if false {
                            v20 += 2147483647
                            for fi28 := 0; fi28 < 13; fi28++ {
                                var v29 uint32 = 196
                                fmt.Println(v2)
                                _ = fi28
                                _ = v29
                            }
                            fmt.Println(v1, fi21, v1)
                        } else {
                            v30 := GenS3{F0: 2147483647, F1: -128, F2: 127}
                            _ = v30
                        }
                        var v31 int8 = -83
                        _ = v31
                    } else {
                        if (1 < v2) {
                            v1 = -413
                        } else {
                            switch (v3 << ((180) & 63)) {
                            case 744:
                                if len(v19) > 0 {
                                    v19[0] = 778
                                }
                            }
                        }
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 bool = (v20 <= v20)
                            v1 = (v2 * (v1 | (v1 / ((739) | 1))))
                            _ = fi32
                            _ = v33
                        }
                        fmt.Println(v20)
                    }
                }
                _ = fi21
            }
        default:
            v34 := GenS1{F0: 255, F1: 1}
            _ = v34
        }
    case "\n\t":
        v1 = (v1 | v2)
    case "a":
        if (true || false) {
            var v35 float64 = (1.5 / 0.1)
            v36 := GenS3{F0: (v3 / ((4294967295) | 1)), F1: 100, F2: -97}
            fmt.Println(v3, v1)
            _ = v35
            _ = v36
        }
    default:
        var v37 uint32 = 4294967295
        _ = v37
    }
    if (3.14159 == 1e10) {
        switch v3 {
        case 4294967295:
            switch 127 {
            case -2147483648:
                var v38 uint32 = 42
                _ = v38
            case 32768:
                var v39 uint16 = 331
                _ = v39
            case 42:
                v19 = append(v19, -735)
            default:
                v20 = -2147483648
            }
        case 916:
            fmt.Println(v1, v20, v3)
        default:
            v40 := GenS4{F0: (!false), F1: "foo bar", F2: 0}
            _ = v40
        }
    } else {
        var v41 int8 = 91
        v41 = 66
        _ = v41
    }
    var v42 int = v3
    fmt.Println(v1, v20)
    v2 = (v1 >> ((18446744073709551615) & 63))
    v43 := GenS1{F0: 2, F1: 2}
    fmt.Println(v3, v42, v3)
    v19 = append(v19, 1)
    v44 := GenS5{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
