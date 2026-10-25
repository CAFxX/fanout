package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 int, p1 uint, p2 int8) uint16 {
    var v5 int64 = 0
    if (3.14159 < 0.0) {
        var v6 uint32 = 2
        v7 := GenS1{F0: 42, F1: (p2 + 1), F2: false}
        _ = v6
        _ = v7
    }
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 15; fi9++ {
        if (-2.25 >= 0.1) {
            if false {
                fmt.Println(p0, v8, fi9)
            }
            p2 = -20
            p2 -= (^(p2))
        } else {
            v5 = v5
            p2 += 42
        }
        fmt.Println(v5, v5)
        if (v5 != (-(v5))) {
            for fi10 := 0; fi10 < 2; fi10++ {
                if (((^(v5)) == (v5 / (((^(v5))) | 1))) && false) {
                    v11 := GenS1{F0: 0, F1: p2, F2: ((float64(0.0) * ((0.1 * (3.14159 * (float64(((((float32((0.0 / (-2.25 - 0.0))) * ((0.0 - 3.14159) * 0.0)) * 0.0) / ((-2.25 - (0.1 * float32((0.0 + (1.5 + 0.0))))) + -2.25)) - -2.25)) * 1.5))) * -2.25)) < -2.25)}
                    v12 := GenS1{F0: -2147483648, F1: p2, F2: true}
                    v13 := GenS1{F0: 42, F1: p2, F2: false}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                switch "hello" {
                case "a":
                    v14 := []int{127, fi10, (-(fi10))}
                    _ = v14
                default:
                    p0 += -1
                }
                _ = fi10
            }
            if (true || (p1 <= (p1 >> (((p1 << ((1) & 63))) & 63)))) {
                v8 += 255
                fmt.Println(v5, v5)
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := []int{fi9, p0, fi15}
                    switch (p0 | (fi9 * (fi15 - p0))) {
                    case -562:
                        v16 = append(v16, (fi15 << ((p1) & 63)))
                    default:
                        p1 += uint(v8)
                    }
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 int8 = -1
                v18 := []int{p0, p0, (fi9 ^ int(v8)), int(v17)}
                if true {
                    var v19 float32 = (-2.25 + 0.1)
                    _ = v19
                } else {
                    switch 127 {
                    case 986:
                        v20 := GenS1{F0: 0, F1: 107, F2: (int8(v5) <= 42)}
                        _ = v20
                    default:
                        var v21 uint8 = v8
                        _ = v21
                    }
                    switch ((("a" + "hello") + "foo bar") + "foo bar") {
                    case "\n\t":
                        fmt.Println(v17)
                    case "":
                        switch "\n\t" {
                        case "a":
                            if ((483 > 519) && false) {
                                v22 := GenS2{F0: v8, F1: -1, F2: "hello"}
                                p1 = 4294967295
                                v18 = append(v18, 32768)
                                _ = v22
                            }
                        default:
                            p0 += fi9
                        }
                    case "foo bar":
                        p1 += (p1 >> ((1) & 63))
                    }
                    v5 = v5
                }
                _ = v17
                _ = v18
            }
        } else {
            v23 := GenS1{F0: -1, F1: p2, F2: false}
            fmt.Println(v8, v8, v5)
            v24 := []int{fi9}
            _ = v23
            _ = v24
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 0
}
func genFn1(p0 int, p1 int64) uint16 {
    v25 := []int{p0, (p0 & p0)}
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int8 = 90
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v25
    return 255
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 18446744073709551615
    var v3 int32 = int32(v2)
    var v4 int64 = 1
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v2, v3)
        v1 += v1
        _ = fi28
    }
    if false {
        var v29 uint64 = 460
        _ = v29
    }
    v30 := GenS3{F0: false}
    v1 = (^(v1))
    v4 -= (v4 ^ v4)
    v2 = 4294967295
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v3, fi31)
        var v32 int = 127
        _ = fi31
        _ = v32
    }
    v3 += 0
    var v33 int64 = v4
    if (!true) {
        v2 -= 42
    } else {
        if (453 != 65535) {
            v34 := GenS4{F0: false}
            _ = v34
        } else {
            var v35 float64 = float64(1e10)
            v36 := GenS1{F0: (v3 - v3), F1: -128, F2: false}
            if false {
                for fi37 := 0; fi37 < 20; fi37++ {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        var v39 uint = 428
                        _ = fi38
                        _ = v39
                    }
                    switch "hello" {
                    case "foo bar":
                        for fi40 := 0; fi40 < 19; fi40++ {
                            switch fi40 {
                            case 540:
                                var v41 int32 = v3
                                _ = v41
                            default:
                                v1 += (v1 << ((uint(v2)) & 63))
                            }
                            v42 := []int{(fi37 | (fi40 * fi37))}
                            var v43 float32 = 0.1
                            _ = fi40
                            _ = v42
                            _ = v43
                        }
                    default:
                        if false {
                            var v44 uint = 466
                            var v45 bool = (165 > 255)
                            for fi46 := 0; fi46 < 3; fi46++ {
                                v44 -= v44
                                if (v1 > (v1 * (v1 % (((v1 << (((^(v44))) & 63))) | 1)))) {
                                    v33 -= (-(v33))
                                    switch fi46 {
                                    case 128:
                                        if false {
                                            v47 := []int{fi46}
                                            v47 = append(v47, fi37)
                                            var v48 float32 = (3.14159 + float32(v35))
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v49 := []int{(fi37 % (((fi46 / (((fi37 & -1)) | 1))) | 1)), 1}
                                            _ = v49
                                        }
                                    case 4294967295:
                                        if v45 {
                                            v35 -= (float64(float32(((float64(1e10) + ((-2.25 / v35) - 0.0)) * 1e10))) + v35)
                                            v50 := []int{fi37, (fi37 << ((42) & 63))}
                                            v51 := GenS5{F0: v4, F1: int64(v1)}
                                            switch 2147483647 {
                                            case 812:
                                                fmt.Println(v33, v35, v1)
                                            case -440:
                                                var v52 uint = (^(v44))
                                                _ = v52
                                            }
                                            switch (^(fi37)) {
                                            case 127:
                                                v44 += (v44 / ((v44) | 1))
                                            case 9223372036854775807:
                                                v35 = (float64(float32((v35 / v35))) + -2.25)
                                            case -489:
                                                v50 = append(v50, 757)
                                            }
                                            _ = v50
                                            _ = v51
                                        } else {
                                            fmt.Println(fi37, v3)
                                            var v53 int64 = v4
                                            _ = v53
                                        }
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(v35, v33)
                                    }
                                }
                                _ = fi46
                            }
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = fi37
                }
            }
            _ = v35
            _ = v36
        }
    }
    if true {
        fmt.Println(v4, v2, v2)
    }
    var v54 float64 = float64((((((3.14159 - float32((float64(-2.25) + 0.0))) * 0.0) / 1e10) + 3.14159) / (3.14159 - 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v54
}
