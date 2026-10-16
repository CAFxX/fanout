package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 uint }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int; F1 uint; F2 uint8 }
type GenS6 struct { F0 string; F1 int32; F2 string }

func main() {
    var v1 string = "hello"
    var v2 uint8 = 139
    var v3 float32 = 1.5
    if (-20 < 0) {
        fmt.Println(v1, v1, v2)
    } else {
        for fi4 := 0; fi4 < 7; fi4++ {
            fmt.Println(v2, v1)
            for fi5 := 0; fi5 < 10; fi5++ {
                if (uint16(v2) >= 2) {
                    v6 := GenS1{F0: 0, F1: 0}
                    _ = v6
                }
                v7 := []int{fi5}
                _ = fi5
                _ = v7
            }
            _ = fi4
        }
        v3 += (-2.25 - 1.5)
        var v8 int = 879
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        switch v1 {
        case "\n\t":
            for fi10 := 0; fi10 < 2; fi10++ {
                if (-2.25 <= 0.0) {
                    v11 := GenS1{F0: 42, F1: 1}
                    v12 := []int{fi10, (^(fi9))}
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v3, v3, fi10)
                }
                var v13 uint16 = 0
                _ = fi10
                _ = v13
            }
        case "hello":
            v3 = 0.0
        default:
            var v14 int64 = -893
            _ = v14
        }
        _ = fi9
    }
    v15 := []int{}
    v3 = float32((float64((float32(1.5) + float32(float64(((v3 * (v3 / float32((-2.25 + ((0.1 * 0.1) / float64(float32(0.1))))))) + 0.1))))) * 3.14159))
    if (!false) {
        v16 := GenS1{F0: 63, F1: 322}
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 = ""
            if (v1 <= v1) {
                var v18 bool = true
                v1 = (v1 + v1)
                for fi19 := 0; fi19 < 16; fi19++ {
                    if false {
                        if (((v3 / v3) - v3) > v3) {
                            var v20 string = "hello"
                            var v21 uint = 19
                            _ = v20
                            _ = v21
                        }
                        v3 += -2.25
                        _ = len(v15)
                    } else {
                        if true {
                            fmt.Println(v3, v3)
                        } else {
                            if (!true) {
                                var v22 uint = 2
                                v23 := GenS2{F0: 0, F1: -2147483648}
                                var v24 uint8 = v2
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                fmt.Println(v2, fi19, fi19)
                                if len(v15) > 0 {
                                    v15[0] = (fi19 * (fi19 % (((fi17 * (-(fi17)))) | 1)))
                                }
                                var v25 int = 0
                                v26 := GenS3{F0: 128, F1: (fi19 ^ fi17), F2: 18446744073709551615}
                                fmt.Println(v3, v18)
                                v27 := GenS3{F0: (v2 & (v2 % (((v2 / ((v2) | 1))) | 1))), F1: (fi19 & (v25 >> ((42) & 63))), F2: 18446744073709551615}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            if (0 == 0) {
                                var v28 uint16 = 1
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                var v29 float64 = (3.14159 - 1.5)
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS2{F0: 42, F1: 2147483647}
                            _ = v30
                        }
                        v2 += v2
                    }
                    if len(v15) > 0 {
                        v15[0] = fi19
                    }
                    if (v18 && v18) {
                        _ = len(v15)
                    }
                    _ = fi19
                }
                _ = v18
            } else {
                _ = len(v15)
                v15 = append(v15, (fi17 + fi17))
            }
            _ = fi17
        }
        _ = v16
    } else {
        v31 := GenS4{F0: (0.1 == 0.0), F1: 4294967295}
        _ = v31
    }
    var v32 int64 = -9223372036854775808
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v33 float64 = 1e10
    for fi34 := 0; fi34 < 14; fi34++ {
        v3 -= v3
        if (true || false) {
            if len(v15) > 0 {
                v15[0] = (fi34 << ((103) & 63))
            }
            var v35 int16 = 0
            for fi36 := 0; fi36 < 20; fi36++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                if false {
                    var v37 float64 = 0.0
                    v38 := GenS5{F0: (fi34 * (fi36 * 972)), F1: 65535, F2: v2}
                    _ = v37
                    _ = v38
                }
                _ = fi36
            }
            _ = v35
        }
        _ = fi34
    }
    fmt.Println(v1)
    v39 := GenS6{F0: v1, F1: 2147483647, F2: (((v1 + "hello") + (v1 + (v1 + (v1 + "\n\t")))) + (((((v1 + "foo bar") + v1) + ("foo bar" + v1)) + "\n\t") + v1))}
    switch v1 {
    case "foo bar":
        var v40 uint64 = 42
        _ = v40
    default:
        v32 = 1
    }
    v32 += (v32 % (((v32 % ((v32) | 1))) | 1))
    switch "a" {
    case "hello":
        var v41 int16 = 42
        _ = v41
    }
    fmt.Println(v33, v1, v33)
    for fi42 := 0; fi42 < 7; fi42++ {
        var v43 int16 = 42
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
}
