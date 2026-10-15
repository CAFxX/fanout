package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 int8 = -128
    var v3 uint64 = 569
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    if false {
        v6 := GenS1{F0: 2}
        _ = v6
    } else {
        var v7 bool = (2147483648 != 2147483648)
        _ = v7
    }
    if (false && (42 != 1)) {
        var v8 float64 = (((1.5 - (float64(((0.1 / 1.0) - (float32((1e10 / 0.1)) - (0.0 + 3.14159)))) * 0.0)) + (0.0 - (0.1 * 0.0))) - (float64(((-2.25 - (1.5 / float32(((0.0 + ((3.14159 / 0.1) * -2.25)) * float64(float32(1.5)))))) * -2.25)) - float64((0.1 - (3.14159 + (1.5 + 0.0))))))
        _ = v8
    } else {
        _ = len(v4)
        var v9 uint = 42
        _ = v9
    }
    var v10 bool = false
    if (-1 == 42) {
        v11 := GenS2{F0: 0, F1: uint(v3), F2: 920}
        v3 = v3
        v12 := GenS2{F0: (v2 + (v2 % (((^(v2))) | 1))), F1: 65535, F2: 1}
        _ = v11
        _ = v12
    }
    v2 += (v2 * (-(v2)))
    switch 32768 {
    case -1:
        var v13 uint8 = v1
        _ = v13
    case 42:
        if (42 != 655) {
            var v14 uint16 = 42
            _ = v14
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    fmt.Println(v3, v10)
                    if v10 {
                        v2 += 33
                        v17 := GenS1{F0: 261}
                        _ = v17
                    }
                    _ = fi16
                }
                _ = fi15
            }
        }
    case -599:
        var v18 int16 = 42
        _ = v18
    default:
        v2 += (^(v2))
    }
    var v19 int8 = v2
    v20 := GenS1{F0: 255}
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch fi22 {
            case 2147483647:
                _ = len(v5)
            case 9223372036854775807:
                fmt.Println(v10, v1)
            case -9223372036854775808:
                v10 = (uint(v1) != 2)
            default:
                if (42 > 101) {
                    v23 := GenS3{F0: 1, F1: 402}
                    if (false && (9223372036854775807 == 9223372036854775807)) {
                        _ = len(v4)
                    } else {
                        var v24 int8 = v2
                        var v25 uint8 = 2
                        _ = v24
                        _ = v25
                    }
                    v10 = false
                    _ = v23
                } else {
                    v1 -= 142
                    if (false || (-1 != -1)) {
                        switch "foo bar" {
                        case "hello":
                            v10 = true
                        case "foo bar":
                            v26 := []int{fi22}
                            v27 := v5[fi21]
                            v1 += (v1 >> ((65535) & 63))
                            _ = v26
                            _ = v27
                        }
                        v3 -= (v3 >> ((42) & 63))
                        v28 := GenS4{F0: 65535}
                        _ = v28
                    } else {
                        fmt.Println(v2)
                    }
                    v29 := []int{fi21, (fi21 >> ((255) & 63)), fi22}
                    if (42 == 255) {
                        v4["hello"] = fi21
                        v30 := GenS5{F0: (v1 / ((v1) | 1))}
                        if (((1e10 + 3.14159) < (-2.25 - 0.0)) || true) {
                            _ = len(v5)
                            v2 += v2
                            v31 := GenS3{F0: 2147483647, F1: 2147483648}
                            _ = v31
                        } else {
                            if false {
                                if (fi21 == 127) {
                                    var v32 string = "hello"
                                    switch v32 {
                                    case "hello":
                                        v33 := GenS3{F0: 2, F1: 2}
                                        _ = v33
                                    case "a":
                                        v3 = v3
                                    case "":
                                        if true {
                                            var v34 uint = 2
                                            var v35 float32 = -2.25
                                            v19 = -128
                                            _ = v34
                                            _ = v35
                                        } else {
                                            if (v10 && (40 < 56)) {
                                                v1 -= (v1 % ((v1) | 1))
                                                v2 -= v19
                                                v3 = (v3 << ((uint(v1)) & 63))
                                                if (!(v10 && ((1.5 - 3.14159) < 0.1))) {
                                                    v36 := GenS6{F0: 0, F1: 419}
                                                    _ = v36
                                                } else {
                                                    v37 := GenS3{F0: 2147483648, F1: 0}
                                                    _ = v37
                                                }
                                            } else {
                                                var v38 float64 = 3.14159
                                                _ = v38
                                            }
                                            v1 -= v1
                                        }
                                    }
                                    var v39 int32 = -1
                                    _ = v32
                                    _ = v39
                                } else {
                                    v29 = append(v29, (fi21 / (((fi21 % ((fi22) | 1))) | 1)))
                                    var v40 int8 = v2
                                    _ = v40
                                }
                            }
                        }
                        _ = v30
                    }
                    _ = v29
                }
            }
            _ = fi22
        }
        v19 += (v2 ^ (v19 / (((v2 ^ v19)) | 1)))
        for fi41 := 0; fi41 < 5; fi41++ {
            fmt.Println(fi21, fi41)
            if false {
                if v10 {
                    v42 := []int{-2147483648}
                    if v10 {
                        var v43 int = 845
                        _ = v43
                    } else {
                        var v44 uint16 = 515
                        switch ("\n\t" + ("a" + ("hello" + ("foo bar" + "\n\t")))) {
                        case "":
                            v2 = v19
                        case "foo bar":
                            v45 := GenS6{F0: v44, F1: v44}
                            _ = v45
                        case "a":
                            v42 = append(v42, fi21)
                        default:
                            var v46 uint32 = 332
                            _ = v46
                        }
                        _ = v44
                    }
                    var v47 int = (fi41 % (((^(fi41))) | 1))
                    _ = v42
                    _ = v47
                }
                v3 += (v3 ^ 984)
            }
            if (!v10) {
                v2 = v2
                v48 := GenS7{F0: -1, F1: (-(fi41))}
                _ = v48
            } else {
                v19 = (v19 ^ (v19 & (v19 | 100)))
                if v10 {
                    v49 := GenS8{F0: 899, F1: (fi41 | (fi41 & fi41))}
                    v19 = v19
                    var v50 int32 = 2147483647
                    _ = v49
                    _ = v50
                } else {
                    v51 := GenS9{F0: 0, F1: 255}
                    var v52 bool = true
                    _ = v51
                    _ = v52
                }
            }
            v19 -= v2
            _ = fi41
        }
        _ = fi21
    }
    var v53 int16 = 775
    v54 := GenS10{F0: 65535}
    v55 := []int{32767, -829, 1}
    fmt.Println(v2, v10, v1)
    v56 := GenS8{F0: uint(v3), F1: 9223372036854775807}
    if false {
        v57 := v4[("\n\t" + "hello")]
        _ = v57
    } else {
        v58 := GenS1{F0: 18446744073709551615}
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
