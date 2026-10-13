package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }
type GenS5 struct { F0 int64; F1 bool; F2 string }
type GenS6 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint, p1 int8) uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 -= p0
        p1 = (p1 - (p1 + p1))
        _ = fi6
    }
    var v7 uint32 = 2147483648
    if true {
        fmt.Println(p1)
        if (-716 >= -2147483648) {
            v7 = (v7 % (((^(v7))) | 1))
        }
        v8 := []int{128, -540, 42}
        _ = v8
    } else {
        switch "a" {
        case "\n\t":
            p0 -= (p0 * (p0 + p0))
        case "a":
            for fi9 := 0; fi9 < 15; fi9++ {
                if false {
                    var v10 int32 = -819
                    _ = v10
                } else {
                    v11 := GenS1{F0: (p1 >> (((p0 & p0)) & 63)), F1: 255, F2: "\n\t"}
                    v12 := []int{fi9, fi9, 42}
                    var v13 float32 = (0.0 / 3.14159)
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        case "foo bar":
            v14 := GenS2{F0: -1}
            _ = v14
        default:
            v15 := []int{}
            _ = v15
        }
    }
    p1 = int8(v7)
    if ((true || (3.14159 != (((1.5 + (-2.25 - 1e10)) / (1e10 + -2.25)) * 1.5))) && (false || false)) {
        var v16 float32 = ((float32(1.5) - (0.1 / 1.5)) / (float32(((-2.25 + (1.5 / 3.14159)) + 0.1)) - 0.1))
        v16 = v16
        p0 += 0
        _ = v16
    } else {
        if false {
            v17 := GenS2{F0: 42}
            _ = v17
        }
        if (750 >= 9223372036854775807) {
            v18 := []int{1}
            var v19 int32 = -2147483648
            if false {
                fmt.Println(p0, p1)
                v20 := GenS3{F0: (true && true), F1: (v7 & v7)}
                _ = v20
            }
            _ = v18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1() uint {
    var v21 float32 = (0.1 - (float32(float64(-2.25)) + 1e10))
    fmt.Println(v21)
    _ = v21
    return 65535
}

func main() {
    v1 := []uint32{65535, 597, 2, 1}
    var v2 int = 48
    var v3 uint = 799
    var v4 float32 = 1e10
    var v5 uint32 = 963
    v4 = ((1.5 * float32((0.1 + (((float64(float32((1.5 + 3.14159))) + 1e10) * 0.1) / -2.25)))) - 3.14159)
    var v22 int8 = 42
    switch "foo bar" {
    case "foo bar":
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int16 = -205
            v2 = 32768
            var v25 int32 = -2147483648
            _ = fi23
            _ = v24
            _ = v25
        }
    case "hello":
        for fi26 := 0; fi26 < 13; fi26++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v5 = (v5 + (v5 * (v5 / (((v5 | 285)) | 1))))
            switch v2 {
            case -9223372036854775808:
                if true {
                    var v27 int32 = 1
                    for fi28 := 0; fi28 < 16; fi28++ {
                        v29 := GenS3{F0: true, F1: (v5 * v5)}
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                } else {
                    v30 := GenS4{F0: 423, F1: 544, F2: 218}
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if (("" + "") >= ("\n\t" + ("a" + ("\n\t" + "hello")))) {
                            v2 = 0
                            var v32 int = fi26
                            _ = v32
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch fi26 {
                            case 32767:
                                var v33 uint = (v3 & v3)
                                _ = v33
                            }
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                        _ = fi31
                    }
                    _ = v30
                }
            case 0:
                fmt.Println(v4, v2)
            default:
                v34 := GenS4{F0: 1, F1: 810, F2: 42}
                _ = v34
            }
            _ = fi26
        }
    case "a":
        if true {
            if true {
                v2 -= v2
                var v35 int = 818
                _ = v35
            }
            if false {
                if ((!true) && ((false && (!true)) || (!false))) {
                    var v36 int32 = 250
                    _ = v36
                }
                v22 += (v22 >> (((^(v3))) & 63))
                switch v2 {
                case -2147483648:
                    var v37 int32 = -665
                    _ = v37
                case -825:
                    v38 := GenS5{F0: 224, F1: ((v2 - 32767) == (v2 >> (((v3 - genFn1())) & 63))), F2: "foo bar"}
                    _ = v38
                }
            } else {
                var v39 uint64 = 632
                _ = v39
            }
            v5 -= v5
        }
    default:
        fmt.Println(v2)
    }
    if false {
        v4 = 0.0
        v4 += float32(0.0)
        var v40 uint = v3
        _ = v40
    }
    switch (v2 + 32768) {
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        if (((196 > -9223372036854775808) || (((v3 | (v3 + uint(v5))) >= v3) || true)) || (false && ((v5 | v5) >= 233))) {
            if true {
                var v42 int32 = -940
                _ = v42
            }
            if (true || true) {
                fmt.Println(v4)
            }
            fmt.Println(v22, fi41, v3)
        } else {
            v43 := GenS6{F0: genFn0(genFn1(), -128), F1: 1}
            _ = v43
        }
        if true {
            v44 := GenS5{F0: 9223372036854775807, F1: ((v5 >> (((^(v3))) & 63)) == v5), F2: "a"}
            var v45 int8 = 127
            if false {
                var v46 uint64 = 4294967295
                _ = v46
            } else {
                v47 := GenS2{F0: 32767}
                v45 += -123
                _ = v47
            }
            _ = v44
            _ = v45
        }
        _ = fi41
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        fmt.Println(fi48)
        _ = fi48
    }
    v1 = append(v1, 147)
    if ((!(true || false)) && (-2.25 <= 1.5)) {
        if false {
            var v49 float64 = 1.5
            var v50 uint64 = 4294967295
            _ = v49
            _ = v50
        } else {
            if len(v1) > 0 {
                v1[0] = 42
            }
            v22 = -11
        }
        v51 := GenS3{F0: (-2.25 >= ((((1.5 + ((-2.25 * 1e10) * (1e10 * ((-2.25 - 0.0) + float64((0.1 - (1.5 - 0.0))))))) - float64(v4)) * (1e10 * 3.14159)) + (1e10 * -2.25))), F1: v5}
        v2 = v2
        _ = v51
    }
    v52 := GenS2{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
