package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
type GenS7 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 uint8, p1 int8, p2 int) uint64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(p2 / ((fi6) | 1)), fi6, (^(fi6)), (p2 | fi6)}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -96, F1: 803, F2: (p0 / ((128) | 1))}
    v9 := GenS2{F0: 18446744073709551615, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int16 = 208
    var v2 int8 = 52
    v3 := make(map[uint32]string)
    var v4 uint8 = 127
    v5 := []string{"hello", "foo bar"}
    for fi10 := 0; fi10 < 5; fi10++ {
        switch "a" {
        case "a":
            var v11 uint8 = v4
            _ = v11
        default:
            var v12 int = fi10
            _ = v12
        }
        v13 := GenS1{F0: 0, F1: 0, F2: v4}
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = fi10
        _ = v13
    }
    var v14 bool = (true && false)
    var v15 int16 = (v1 / ((-1) | 1))
    var v16 uint16 = 107
    v17 := GenS3{F0: uint64(v4), F1: (65535 != 61), F2: ("\n\t" + "a")}
    var v18 string = "foo bar"
    v15 = v1
    v16 = v16
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 uint16 = (^(v16))
            v4 = 1
            fmt.Println(v4, fi19)
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v22 uint32 = 2147483647
    v16 = (^(v16))
    v23 := GenS4{F0: 492}
    if true {
        fmt.Println(v4, v14)
        v2 += v2
        if ((v22 ^ (^(v22))) <= (v22 | v22)) {
            v24 := GenS3{F0: uint64(v4), F1: v14, F2: ""}
            _ = v24
        } else {
            _ = len(v3)
            _ = len(v5)
        }
    } else {
        var v25 uint32 = 42
        v25 += v22
        for fi26 := 0; fi26 < 4; fi26++ {
            v15 = (v15 + v1)
            if v14 {
                var v27 int16 = v1
                switch ("foo bar" + (v18 + v18)) {
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    v14 = (!(!v14))
                }
                if (!v14) {
                    var v28 uint = 811
                    for fi29 := 0; fi29 < 20; fi29++ {
                        var v30 float32 = float32(3.14159)
                        v31 := GenS3{F0: 620, F1: (v14 || (!true)), F2: "foo bar"}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    var v32 int32 = 2147483647
                    _ = v28
                    _ = v32
                }
                _ = v27
            } else {
                fmt.Println(v22)
                if (v1 <= 0) {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        switch fi26 {
                        case 42:
                            fmt.Println(v22)
                            if v14 {
                                switch v18 {
                                case "foo bar":
                                    switch fi26 {
                                    case 1:
                                        _ = len(v5)
                                    default:
                                        var v34 int64 = -9223372036854775808
                                        _ = v34
                                    }
                                case "hello":
                                    if (v16 > v16) {
                                        fmt.Println(v15)
                                    }
                                    if v14 {
                                        v14 = ((v18 + "") > "foo bar")
                                        var v35 int = -175
                                        fmt.Println(v35, v18, fi26)
                                        _ = v35
                                    } else {
                                        v3[(^(v25))] = (v18 + "foo bar")
                                        v36 := GenS5{F0: v18, F1: v1, F2: genFn0(v4, -1, 42)}
                                        var v37 string = v18
                                        _ = v36
                                        _ = v37
                                    }
                                    v1 = v1
                                case "a":
                                    v38 := v3[2]
                                    _ = v38
                                }
                                var v39 uint = 18446744073709551615
                                _ = len(v3)
                                v22 += 65535
                                v40 := GenS6{F0: -196, F1: v4, F2: -791}
                                if (v14 && (100 >= v2)) {
                                    if (v2 <= (v2 ^ int8(fi33))) {
                                        if (-2.25 <= float32((((0.0 * (((0.1 + 3.14159) * 1e10) + 0.0)) - (3.14159 + 3.14159)) + (float64((0.1 / (3.14159 - float32(float64(1e10))))) / 1.0)))) {
                                            fmt.Println(v25, v4, v1)
                                            v39 -= (v39 & (^(v39)))
                                        } else {
                                            var v41 int64 = 1
                                            v14 = true
                                            v16 = v16
                                            _ = v41
                                        }
                                        v42 := GenS7{F0: ((((v1 - -32768) < 75) && v14) || (!(!(uint32(v39) == (v22 << ((v39) & 63)))))), F1: uint8(v25), F2: int(v25)}
                                        _ = v42
                                    } else {
                                        v4 += 128
                                    }
                                    switch fi33 {
                                    case -561:
                                        v43 := GenS6{F0: -861, F1: (^(v4)), F2: -1}
                                        _ = v43
                                    default:
                                        v16 = 65535
                                    }
                                    if false {
                                        var v44 float64 = ((-2.25 * 0.0) + 0.0)
                                        if true {
                                            var v45 uint16 = (v16 | (v16 % ((0) | 1)))
                                            v46 := GenS1{F0: 42, F1: v39, F2: (v4 / (((v4 | v4)) | 1))}
                                            _ = v45
                                            _ = v46
                                        }
                                        v15 -= 42
                                        _ = v44
                                    }
                                }
                                if (v18 < "\n\t") {
                                    _ = len(v5)
                                }
                                _ = len(v5)
                                _ = v39
                                _ = v40
                            }
                        case 15:
                            v47 := GenS6{F0: 2147483647, F1: v4, F2: 1}
                            _ = v47
                        default:
                            switch fi26 {
                            case 1:
                                switch v18 {
                                case "a":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            case -1:
                                var v48 int64 = -9223372036854775808
                                _ = v48
                            case 2147483647:
                                fmt.Println(v16, v1)
                            }
                        }
                        v49 := v3[631]
                        _ = fi33
                        _ = v49
                    }
                    v25 += 2147483648
                } else {
                    if (!(!(v14 && (0.1 == 0.0)))) {
                        if v14 {
                            v2 = (v2 % (((v2 >> ((2) & 63))) | 1))
                        } else {
                            switch "a" {
                            case "\n\t":
                                v2 = -81
                            case "hello":
                                var v50 int64 = -1
                                _ = v50
                            }
                            if v14 {
                                v1 = (v15 ^ v1)
                                v3[(v22 / (((v25 / ((v25) | 1))) | 1))] = (((v18 + (v18 + (v18 + "foo bar"))) + ((v18 + v18) + (("a" + (v18 + (v18 + v18))) + (v18 + ((("" + v18) + "hello") + v18))))) + ((v18 + (v18 + v18)) + (v18 + "a")))
                                var v51 int = -9223372036854775808
                                _ = v51
                            }
                        }
                    }
                }
            }
            _ = len(v5)
            _ = fi26
        }
        _ = v25
    }
    var v52 int32 = -2147483648
    if (-78 >= (v52 / ((v52) | 1))) {
        for fi53 := 0; fi53 < 3; fi53++ {
            _ = len(v5)
            v16 = (^(v16))
            fmt.Println(v2, v52, v1)
            _ = fi53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
