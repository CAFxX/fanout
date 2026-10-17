package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint32) int16 {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 += (p0 << ((0) & 63))
        _ = fi4
    }
    v5 := []int{0}
    _ = p0
    _ = v5
    return -32768
}
func genFn1(p0 int32, p1 int) string {
    p1 = 0
    v6 := []int{42, (p1 ^ -263), 9223372036854775807}
    p1 = (p1 / (((p1 & (p1 + p1))) | 1))
    _ = p0
    _ = p1
    _ = v6
    return "hello"
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{2147483648, 65535, 785}
    v3 := make(map[int]uint32)
    var v7 float32 = 0.1
    if true {
        var v8 uint64 = 42
        _ = v8
    }
    var v9 int64 = 831
    for fi10 := 0; fi10 < 6; fi10++ {
        if (true && (true || (-471 == -798))) {
            switch -1 {
            case -279:
                fmt.Println(fi10)
            case -1:
                v7 = float32(-2.25)
            }
            var v11 int64 = v9
            fmt.Println(v11)
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            switch ("foo bar" + (("hello" + "a") + ("\n\t" + "\n\t"))) {
            case "hello":
                if ((-1 != v9) && false) {
                    v13 := GenS1{F0: 2, F1: fi12, F2: 255}
                    _ = v13
                } else {
                    var v14 string = ""
                    v9 += 1
                    _ = v14
                }
            case "\n\t":
                v7 -= ((0.1 - float32((((((0.0 + ((-2.25 - 1.5) - 0.0)) - 1.5) - 0.1) / (1e10 + 1.5)) - float64(v7)))) / ((float32(3.14159) + float32(((3.14159 * 0.1) / ((3.14159 - 1e10) + (0.1 - (1.5 * float64(v7))))))) * (v7 / v7)))
            default:
                var v15 int32 = -1
                _ = v15
            }
            _ = fi12
        }
        v9 -= (v9 << ((105) & 63))
        _ = fi10
    }
    fmt.Println(v1, v1)
    if ((false && false) && false) {
        for fi16 := 0; fi16 < 9; fi16++ {
            v1 = genFn0(2147483647)
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 65535, F1: fi16, F2: 255}
                var v19 uint64 = 0
                v9 += -1
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: false, F1: true, F2: v9}
            _ = fi16
            _ = v20
        }
        var v21 int8 = 42
        _ = v21
    } else {
        switch 814 {
        case 973:
            v9 = 1
        case -9223372036854775808:
            v22 := GenS3{F0: 1000, F1: 128, F2: -140}
            _ = v22
        }
        v1 -= v1
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float64 = 1.5
            var v26 int64 = (^(v9))
            if false {
                if ((fi24 > fi23) && (!(false || ((^(v26)) != (v26 ^ (v26 / ((v9) | 1))))))) {
                    var v27 int64 = (v26 - v9)
                    if ((v1 >> ((0) & 63)) < (v1 % ((1000) | 1))) {
                        v25 += (0.0 - (v25 - -2.25))
                        v28 := GenS2{F0: true, F1: true, F2: (v27 << ((18446744073709551615) & 63))}
                        v1 -= (v1 + v1)
                        _ = v28
                    }
                    if ((18446744073709551615 == 18446744073709551615) && ((false || true) && (v9 >= 42))) {
                        var v29 uint16 = 42
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v27, v26)
                            v27 = -1
                            switch "a" {
                            case "hello":
                                var v30 uint64 = 42
                                _ = v30
                            case "\n\t":
                                var v31 int64 = v9
                                _ = v31
                            case "foo bar":
                                if (genFn0(4294967295) == -32768) {
                                    if len(v2) > 0 {
                                        v2[0] = 2
                                    }
                                    v3[(fi24 % (((fi24 - fi23)) | 1))] = 0
                                }
                                _ = len(v2)
                                v32 := GenS2{F0: false, F1: (v25 >= (v25 + ((v25 + v25) / v25))), F2: (v26 - v27)}
                                _ = v32
                            }
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = 438
                            }
                        case "":
                            fmt.Println(v29)
                        default:
                            if false {
                                if (((3.14159 + ((v7 + v7) / (v7 * 1.5))) / v7) >= (-2.25 * (v7 + (float32(0.1) - v7)))) {
                                    v9 += (v27 ^ -473)
                                } else {
                                    v33 := GenS2{F0: false, F1: true, F2: (v9 / ((-1) | 1))}
                                    v34 := GenS2{F0: (v25 != ((v25 * (v25 + (-2.25 + v25))) / float64((((float32(((v25 * 1.5) + (v25 - v25))) / 1e10) - v7) - v7)))), F1: ((-1 <= -23) || true), F2: v27}
                                    v35 := GenS1{F0: 2147483648, F1: int(v26), F2: 42}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v1 += v1
                            } else {
                                v7 = v7
                                var v36 int64 = (-(v27))
                                fmt.Println(v25, v1, v36)
                                if true {
                                    v29 = 492
                                    var v37 int = fi23
                                    if ((fi23 | int(v1)) > v37) {
                                        var v38 float32 = (1e10 / v7)
                                        _ = v38
                                    }
                                    _ = v37
                                }
                                _ = v36
                            }
                        }
                        _ = v29
                    } else {
                        v39 := GenS1{F0: 707, F1: fi24, F2: 42}
                        v2 = append(v2, 0)
                        var v40 bool = true
                        _ = v39
                        _ = v40
                    }
                    _ = v27
                }
                switch 9223372036854775807 {
                case 4294967295:
                    var v41 int = 268
                    _ = v41
                case 810:
                    var v42 bool = true
                    _ = v42
                case 32768:
                    if true {
                        if (!true) {
                            switch ("\n\t" + "") {
                            case "\n\t":
                                var v43 int16 = v1
                                _ = v43
                            case "hello":
                                fmt.Println(v7)
                                v44 := GenS4{F0: 2147483648}
                                switch "\n\t" {
                                case "":
                                    if (!true) {
                                        var v45 int = fi23
                                        _ = v45
                                    } else {
                                        if true {
                                            var v46 int64 = (v26 - v26)
                                            _ = v46
                                        } else {
                                            v47 := GenS5{F0: 1000}
                                            if (234 == 2147483647) {
                                                v25 += (float64((v7 * float32((v25 + (v25 / -2.25))))) - ((v25 * v25) * (((((v25 * 1.5) / v25) / v25) * v25) - (v25 / (v25 - v25)))))
                                                v9 += (v26 / (((v9 | -800)) | 1))
                                            } else {
                                                if (((1e10 - -2.25) - 1.5) == 1.5) {
                                                    if true {
                                                        switch (-(fi24)) {
                                                        case 9223372036854775807:
                                                            v1 = 156
                                                        case 4294967295:
                                                            var v48 int64 = 9223372036854775807
                                                            _ = v48
                                                        case 2147483647:
                                                            var v49 uint16 = 42
                                                            _ = v49
                                                        }
                                                        var v50 uint64 = 18446744073709551615
                                                        _ = v50
                                                    }
                                                }
                                            }
                                            v26 = v9
                                            _ = v47
                                        }
                                    }
                                default:
                                    fmt.Println(v9)
                                }
                                _ = v44
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        v9 += (v9 + (v26 << ((18446744073709551615) & 63)))
                        v7 += v7
                        fmt.Println(fi23, v25, v9)
                        var v51 int8 = -52
                        _ = v51
                    } else {
                        var v52 int32 = -976
                        _ = v52
                    }
                }
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    v1 += 319
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v53 uint = 611
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v7
    _ = v9
}
