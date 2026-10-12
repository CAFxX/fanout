package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 int32 }
func genFn0() string {
    v4 := []int{9223372036854775807, 32767}
    for fi5 := 0; fi5 < 18; fi5++ {
        switch ("\n\t" + "a") {
        case "":
            fmt.Println(fi5)
        case "\n\t":
            switch "hello" {
            case "a":
                var v6 int8 = 1
                _ = v6
            default:
                var v7 int16 = -1
                _ = v7
            }
        case "foo bar":
            fmt.Println(fi5)
        default:
            fmt.Println(fi5)
        }
        if ((!false) && (((fi5 <= (fi5 + fi5)) && ((fi5 ^ (fi5 >> ((534) & 63))) != fi5)) || (-48 >= -1))) {
            if (!(false || (!(4294967295 < 61)))) {
                fmt.Println(fi5)
            }
        }
        _ = fi5
    }
    _ = v4
    return "foo bar"
}
func genFn1(p0 int32, p1 uint32, p2 uint) uint {
    var v8 float32 = ((3.14159 + 1.5) - 1e10)
    if false {
        fmt.Println(v8, p1)
        v9 := GenS1{F0: (p1 | 705), F1: "\n\t", F2: (110 < -1)}
        _ = v9
    }
    v8 = (3.14159 * (v8 + (float32(float64(v8)) + -2.25)))
    v10 := []int{2147483647, 9223372036854775807, 499}
    switch 32768 {
    case 2147483647:
        if (255 > 250) {
            v11 := GenS2{F0: 127, F1: 188, F2: 271}
            _ = v11
        }
    case 32768:
        var v12 uint32 = 2147483647
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    return p2
}

func main() {
    var v1 uint32 = 131
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    if (42 != 245) {
        v2 -= (1.5 + 1e10)
        fmt.Println(v2, v1)
        var v13 uint64 = 593
        _ = v13
    }
    var v14 float32 = 3.14159
    v1 = v1
    var v15 int32 = 42
    var v16 float64 = 0.0
    if ((v14 - float32(3.14159)) <= (v14 + v14)) {
        var v17 uint8 = 201
        v18 := []int{-1, -9223372036854775808}
        if true {
            switch "\n\t" {
            case "foo bar":
                var v19 uint16 = 42
                _ = v19
            default:
                v15 -= (v15 * -1)
            }
        } else {
            v20 := GenS2{F0: 128, F1: v17, F2: v1}
            switch "\n\t" {
            case "a":
                if false {
                    var v21 int16 = 836
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        _ = len(v18)
                        if false {
                            v16 = v16
                        } else {
                            for fi23 := 0; fi23 < 9; fi23++ {
                                v24 := GenS3{F0: 1}
                                var v25 int64 = 9223372036854775807
                                v2 = v2
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = fi22
                    }
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v1, v15)
                    default:
                        for fi26 := 0; fi26 < 5; fi26++ {
                            if false {
                                if (!(true || (46 > 0))) {
                                    fmt.Println(v2, v15, v1)
                                } else {
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        var v28 int = (fi26 / ((32767) | 1))
                                        var v29 uint8 = 42
                                        switch 127 {
                                        case 429:
                                            var v30 bool = (24 == v1)
                                            _ = v30
                                        case -281:
                                            fmt.Println(fi26)
                                        case 32768:
                                            switch (("" + "foo bar") + "") {
                                            case "\n\t":
                                                switch "" {
                                                case "foo bar":
                                                    var v31 int = fi27
                                                    _ = v31
                                                }
                                            default:
                                                v15 -= (v15 & v15)
                                                v2 += (v16 + v2)
                                                v15 -= (v15 / ((v15) | 1))
                                            }
                                        }
                                        switch 4294967295 {
                                        case 494:
                                            fmt.Println(v29, v1, fi27)
                                        default:
                                            var v32 bool = true
                                            _ = v32
                                        }
                                        v33 := GenS4{F0: 9223372036854775807, F1: 18446744073709551615}
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                        _ = v33
                                    }
                                    if false {
                                        for fi34 := 0; fi34 < 7; fi34++ {
                                            var v35 uint8 = v17
                                            if ((v1 & v1) <= v1) {
                                                var v36 uint32 = (v1 % (((v1 * (v1 >> ((2) & 63)))) | 1))
                                                var v37 float64 = (v2 / v16)
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v38 uint = genFn1((-(v15)), 19, 2)
                                                if (1 <= 255) {
                                                    v38 = genFn1(-732, 2147483647, genFn1((v15 >> ((531) & 63)), (v1 >> ((v38) & 63)), genFn1((v15 % ((int32(v1)) | 1)), v1, v38)))
                                                }
                                                v17 = v17
                                                _ = v38
                                            }
                                            v39 := GenS5{F0: 1, F1: 1000, F2: v15}
                                            _ = fi34
                                            _ = v35
                                            _ = v39
                                        }
                                        var v40 uint32 = (v1 ^ (v1 % (((^(v1))) | 1)))
                                        for fi41 := 0; fi41 < 18; fi41++ {
                                            var v42 uint32 = v40
                                            _ = fi41
                                            _ = v42
                                        }
                                        _ = v40
                                    }
                                }
                                for fi43 := 0; fi43 < 3; fi43++ {
                                    _ = len(v18)
                                    if (!true) {
                                        v17 = 127
                                        if len(v18) > 0 {
                                            v18[0] = (fi43 % (((fi26 + int(v1))) | 1))
                                        }
                                    } else {
                                        var v44 uint16 = 2
                                        fmt.Println(v17)
                                        var v45 string = ((("\n\t" + "hello") + "foo bar") + (("a" + "\n\t") + "a"))
                                        fmt.Println(fi43)
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = fi43
                                }
                            }
                            fmt.Println(v1)
                            _ = fi26
                        }
                    }
                }
            case "":
                v17 = (v17 >> ((823) & 63))
            case "hello":
                v2 = ((((1e10 * ((v2 / (v16 + (0.0 + -2.25))) + v16)) + (1e10 * v2)) + v16) * (-2.25 + -2.25))
            }
            _ = v20
        }
        _ = v17
        _ = v18
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        for fi47 := 0; fi47 < 19; fi47++ {
            delete(v3, (v1 / ((v1) | 1)))
            v16 = (v2 * (((v16 - (v16 / ((float64(float32(v2)) / 1.0) + 0.1))) * v16) - v2))
            if true {
                v15 -= (v15 * 42)
                if ((0.1 * 0.0) > v14) {
                    v15 = -1
                    var v48 uint64 = 0
                    _ = v48
                } else {
                    v16 = (3.14159 + ((v16 - (((float64(0.0) / ((1e10 - v2) / v2)) - float64(v14)) + 1e10)) / v2))
                }
            } else {
                if false {
                    fmt.Println(v2, v14, v14)
                    if (true || (1.5 != -2.25)) {
                        var v49 uint32 = v1
                        var v50 uint16 = 332
                        _ = v49
                        _ = v50
                    } else {
                        v51 := GenS2{F0: 2, F1: 2, F2: v1}
                        var v52 uint = 204
                        if true {
                            switch fi46 {
                            case 128:
                                v53 := []int{-484, fi47}
                                _ = v53
                            case 127:
                                v54 := []int{(fi46 % (((fi47 << ((65535) & 63))) | 1)), -1, fi47, 121}
                                var v55 int8 = -1
                                _ = v54
                                _ = v55
                            case -2147483648:
                                var v56 bool = true
                                _ = v56
                            }
                            fmt.Println(v2)
                            var v57 int64 = 134
                            _ = v57
                        } else {
                            if (v1 != (v1 * 1)) {
                                var v58 int8 = -1
                                _ = v58
                            } else {
                                v59 := []int{-199, int(v52), fi46}
                                _ = v59
                            }
                        }
                        _ = v51
                        _ = v52
                    }
                    v60 := []int{(^(fi46)), fi47}
                    _ = v60
                } else {
                    v2 = float64(1e10)
                }
            }
            _ = fi47
        }
        var v61 float32 = v14
        v62 := []int{fi46, (-(fi46)), 695, (-(fi46))}
        _ = fi46
        _ = v61
        _ = v62
    }
    var v63 uint8 = 0
    v64 := []int{296, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v63
    _ = v64
}
