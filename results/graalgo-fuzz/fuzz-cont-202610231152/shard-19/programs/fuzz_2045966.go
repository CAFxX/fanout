package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 uint }
type GenS7 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64, p1 uint) string {
    v7 := GenS1{F0: 136}
    var v8 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("\n\t" + (("a" + "") + ((("hello" + "\n\t") + "hello") + "foo bar")))
}
func genFn1(p0 int8) uint32 {
    fmt.Println(p0)
    var v9 uint = 2
    if (1.5 != 3.14159) {
        var v10 string = "\n\t"
        _ = v10
    }
    p0 -= (^(p0))
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    var v3 int = 1
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    var v6 float32 = (1.5 + -2.25)
    var v11 float32 = (3.14159 * float32((0.0 + ((0.0 - (float64(v6) * -2.25)) / -2.25))))
    v11 = (v11 + v6)
    v12 := GenS2{F0: (v3 - 2147483647), F1: 32767}
    v13 := GenS3{F0: 24, F1: 42}
    if (1e10 > v11) {
        var v14 uint64 = 42
        var v15 string = ""
        for fi16 := 0; fi16 < 6; fi16++ {
            if true {
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int32 = 0
                    var v19 int32 = (v18 << ((782) & 63))
                    v14 = (v14 * v14)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if (2 <= 0) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS4{F0: 42, F1: true}
                        v15 = ((v15 + "a") + (v15 + (("hello" + (v15 + v15)) + ((v15 + v15) + ""))))
                        _ = fi20
                        _ = v21
                    }
                    if (false || true) {
                        if true {
                            if true {
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    if (-138 < -1) {
                                        var v23 int16 = 32767
                                        var v24 string = (v15 + (("hello" + "hello") + v15))
                                        v25 := GenS5{F0: v1, F1: 128}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v14 += v14
                                        fmt.Println(v1)
                                        var v26 int16 = 99
                                        _ = v26
                                    }
                                    _ = fi22
                                }
                                v27 := []int{}
                                _ = v27
                            } else {
                                if false {
                                    var v28 uint32 = 2147483647
                                    var v29 uint16 = 156
                                    var v30 float32 = v6
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                v3 -= (fi16 >> ((255) & 63))
                            }
                        }
                    }
                } else {
                    if false {
                        v31 := GenS1{F0: genFn1((-(v1)))}
                        v2[fi16] = 121
                        _ = v31
                    }
                }
            }
            v32 := GenS5{F0: (v1 * -44), F1: 127}
            v33 := GenS6{F0: 2, F1: 0, F2: 18446744073709551615}
            _ = fi16
            _ = v32
            _ = v33
        }
        _ = v14
        _ = v15
    }
    v34 := GenS4{F0: 2, F1: ((v11 - (v6 + (0.1 / ((v6 + v11) + (((-2.25 - 0.0) / (1e10 - float32(0.1))) - v6))))) < v11)}
    v11 -= (0.0 + float32(3.14159))
    if (18446744073709551615 > 42) {
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS2{F0: v3, F1: 209}
            _ = fi35
            _ = v36
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            if (false && ((float64(3.14159) + 1e10) > -2.25)) {
                var v37 uint32 = 209
                _ = v37
            } else {
                v11 = (v6 - (v11 * v11))
            }
        case "a":
            v38 := GenS7{F0: 42, F1: 42}
            _ = v38
        default:
            v11 -= (v6 - v11)
        }
    }
    fmt.Println(v3, v6)
    v39 := GenS2{F0: 32768, F1: 551}
    v40 := GenS1{F0: 42}
    fmt.Println(v11)
    fmt.Println(v6, v6)
    _ = len(v4)
    v41 := []int{(^(v3)), v3, v3, v3}
    v11 += v11
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
