package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint8) uint {
    v7 := []int{-9223372036854775808}
    v8 := GenS1{F0: "\n\t", F1: 42, F2: 767}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int) int8 {
    p0 = -1
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: "hello", F1: -603, F2: 255}
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: "\n\t", F1: 42, F2: genFn0(414, 205)}
            v13 := []int{-9223372036854775808, 0}
            v13 = append(v13, fi11)
            _ = fi11
            _ = v12
            _ = v13
        }
        p0 = (-(p0))
        if true {
            v14 := GenS2{F0: 1000, F1: 2}
            v15 := []int{p0}
            v16 := GenS1{F0: "a", F1: 106, F2: 18446744073709551615}
            _ = v14
            _ = v15
            _ = v16
        } else {
            p0 = 264
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int = (p0 * (^(p0)))
                v18 = -859
                var v19 uint16 = 2
                _ = fi17
                _ = v18
                _ = v19
            }
            fmt.Println(p0)
        }
    }
    p0 += (p0 >> ((genFn0(255, 164)) & 63))
    _ = p0
    return -128
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = 0.0
    var v3 uint8 = 89
    var v4 float32 = (float32((v2 / (float64((1.5 * ((3.14159 - 0.1) - 0.0))) / (float64((1e10 * 0.0)) / (0.0 - v2))))) + (1.5 * ((3.14159 * ((0.1 / 3.14159) - (0.1 + 3.14159))) / ((float32((v2 + ((float64(1e10) + v2) * 1e10))) - 0.0) - -2.25))))
    var v5 float64 = 0.0
    var v6 int16 = 32767
    var v20 int = -9223372036854775808
    if (!((v1 | (v1 ^ 55)) > 1)) {
        v21 := GenS1{F0: "foo bar", F1: -833, F2: 0}
        var v22 uint = 0
        if true {
            for fi23 := 0; fi23 < 2; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 int16 = 32767
                    if (false || (((false || true) && (42 != 0)) || false)) {
                        var v26 string = "hello"
                        fmt.Println(fi24, v1, v6)
                        _ = v26
                    } else {
                        var v27 int32 = 695
                        var v28 int64 = 903
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v2, fi23)
                v20 += (v20 / (((v20 + 0)) | 1))
                _ = fi23
            }
        } else {
            v5 += 3.14159
        }
        _ = v21
        _ = v22
    } else {
        v6 = 587
        v29 := []int{42, v20, 276, -413}
        _ = v29
    }
    if (true || (false && false)) {
        var v30 int = v20
        switch "foo bar" {
        case "a":
            if (true || false) {
                var v31 uint16 = (^(v1))
                v30 -= v20
                _ = v31
            } else {
                for fi32 := 0; fi32 < 11; fi32++ {
                    fmt.Println(fi32)
                    v20 = 0
                    for fi33 := 0; fi33 < 2; fi33++ {
                        fmt.Println(v3, v3, v2)
                        v34 := GenS1{F0: "\n\t", F1: v6, F2: genFn0(65535, (^(v3)))}
                        if ((((((v1 - (v1 * v1)) < 115) && true) || (v6 == v6)) || (v1 >= 42)) || false) {
                            var v35 uint16 = v1
                            _ = v35
                        }
                        _ = fi33
                        _ = v34
                    }
                    _ = fi32
                }
                v36 := []int{(v20 ^ (v30 | v20))}
                _ = v36
            }
        case "hello":
            fmt.Println(v1, v5)
        default:
            v37 := []int{}
            _ = v37
        }
        _ = v30
    } else {
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 float32 = v4
            v40 := GenS3{F0: 2147483647, F1: v20, F2: (^(v6))}
            v3 = (v3 / (((v3 << ((18446744073709551615) & 63))) | 1))
            _ = fi38
            _ = v39
            _ = v40
        }
        v2 -= (float64((((v4 * float32(((0.1 * v5) / float64(float32(float64(v4)))))) * v4) * (v4 / v4))) * (float64(0.0) - v2))
    }
    if (42 < 108) {
        v41 := []int{(v20 << ((18446744073709551615) & 63)), v20, 9223372036854775807}
        _ = v41
    } else {
        if false {
            v3 = v3
        } else {
            var v42 bool = true
            _ = v42
        }
        switch v20 {
        case 42:
            v43 := []int{(v20 << ((255) & 63))}
            _ = v43
        default:
            switch "" {
            case "a":
                v44 := GenS4{F0: false, F1: true}
                _ = v44
            case "foo bar":
                for fi45 := 0; fi45 < 10; fi45++ {
                    for fi46 := 0; fi46 < 10; fi46++ {
                        var v47 uint16 = 407
                        v48 := []int{fi45, 127}
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    if true {
                        var v49 uint16 = v1
                        v50 := []int{(fi45 >> ((uint(v3)) & 63)), fi45, fi45}
                        var v51 bool = false
                        _ = v49
                        _ = v50
                        _ = v51
                    } else {
                        switch v20 {
                        case -2147483648:
                            if false {
                                v6 = (v6 % (((v6 / ((v6) | 1))) | 1))
                                var v52 uint16 = v1
                                _ = v52
                            }
                        }
                    }
                    _ = fi45
                }
            default:
                v53 := GenS5{F0: 1}
                _ = v53
            }
        }
    }
    var v54 uint16 = 32767
    fmt.Println(v1)
    switch "foo bar" {
    case "foo bar":
        for fi55 := 0; fi55 < 7; fi55++ {
            for fi56 := 0; fi56 < 9; fi56++ {
                fmt.Println(v1, v54)
                v20 = (v20 % ((fi55) | 1))
                _ = fi56
            }
            _ = fi55
        }
    case "hello":
        var v57 bool = true
        _ = v57
    case "\n\t":
        var v58 int64 = 9223372036854775807
        _ = v58
    }
    var v59 bool = false
    switch v20 {
    case 6:
        switch "" {
        case "\n\t":
            v60 := []int{v20}
            _ = v60
        case "":
            fmt.Println(v3, v6, v5)
        case "a":
            v4 += 0.1
        }
    case -1:
        var v61 float32 = (1.5 - (((((1.5 - v4) + v4) - ((3.14159 + float32((v5 * (v2 + (float64(((((((v4 + float32(v5)) + float32(((((0.1 - v5) + v5) * v2) / v5))) * ((0.1 - 1e10) / ((1.5 / ((v4 + ((1.5 + 0.1) / 3.14159)) * ((v4 - v4) / v4))) / v4))) + ((v4 * v4) / ((v4 - v4) / v4))) / 0.1) + (1e10 - v4))) + v5))))) + float32((3.14159 * v2)))) * v4) / (v4 + v4)))
        _ = v61
    case -274:
        switch "a" {
        case "foo bar":
            v62 := []int{-944, (v20 & (v20 | v20)), (v20 * (v20 / (((v20 | v20)) | 1))), (v20 - (v20 >> ((730) & 63)))}
            _ = v62
        case "":
            fmt.Println(v54, v5, v4)
        default:
            v63 := GenS5{F0: 931}
            _ = v63
        }
    }
    if false {
        v1 = (v54 << ((18446744073709551615) & 63))
    }
    v20 = v20
    for fi64 := 0; fi64 < 17; fi64++ {
        for fi65 := 0; fi65 < 1; fi65++ {
            v1 -= (^(v1))
            _ = fi65
        }
        v1 += uint16(v3)
        v59 = v59
        _ = fi64
    }
    for fi66 := 0; fi66 < 14; fi66++ {
        v54 = (v54 & v54)
        _ = fi66
    }
    var v67 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v59
    _ = v6
    _ = v67
}
