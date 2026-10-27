package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((0.0 * 0.1) / 1.0)
    var v3 string = (((("foo bar" + (("a" + ("foo bar" + "")) + ("foo bar" + ("\n\t" + (("a" + "a") + "foo bar"))))) + "a") + ("\n\t" + "a")) + "a")
    var v4 uint32 = 773
    if false {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int = -558
            if (v4 < v4) {
                var v7 int = int(v4)
                _ = v7
            }
            v1[(v4 * v4)] = false
            _ = fi5
            _ = v6
        }
        var v8 uint8 = 1
        v2 = (v2 - v2)
        _ = v8
    }
    switch 32768 {
    case 1:
        var v9 bool = (!true)
        _ = v9
    case 128:
        var v10 float32 = -2.25
        _ = v10
    case 32767:
        v4 -= 392
    default:
        v2 = v2
    }
    var v11 int8 = 42
    if (!false) {
        v3 = "a"
    }
    v12 := []int{4294967295, 2147483647}
    var v13 uint8 = 127
    v14 := GenS1{F0: 808, F1: 32768, F2: -1}
    if (((!((v3 < v3) && false)) || true) || true) {
        v11 += v11
        _ = len(v12)
        var v15 float32 = -2.25
        _ = v15
    }
    v4 -= v4
    var v16 int32 = -348
    v13 = 57
    v17 := GenS2{F0: v13, F1: v11}
    var v18 uint32 = v4
    var v19 uint64 = 2
    if true {
        _ = len(v12)
        if ((1 < 18446744073709551615) && true) {
            delete(v1, 65535)
            fmt.Println(v2, v3)
        }
    } else {
        if (128 < 0) {
            delete(v1, (^(v18)))
            if (!(0 <= -1)) {
                if (int16(v16) > 32767) {
                    var v20 int32 = v16
                    _ = v20
                } else {
                    var v21 uint64 = 595
                    for fi22 := 0; fi22 < 9; fi22++ {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            var v24 float64 = (v2 + 1.5)
                            v25 := GenS3{F0: 536, F1: v11}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 uint64 = (v21 / ((0) | 1))
                        _ = fi22
                        _ = v26
                    }
                    var v27 int8 = (v11 & (^(v11)))
                    _ = v21
                    _ = v27
                }
                var v28 uint8 = (v13 + 202)
                if (-480 >= v16) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        var v29 string = v3
                        _ = v29
                    case 127:
                        v12 = append(v12, 127)
                    default:
                        switch 9223372036854775807 {
                        case 260:
                            var v30 bool = true
                            _ = v30
                        default:
                            fmt.Println(v13)
                        }
                    }
                    for fi31 := 0; fi31 < 15; fi31++ {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            var v33 int16 = -720
                            switch 917 {
                            case 1:
                                v34 := GenS1{F0: 18446744073709551615, F1: fi31, F2: v33}
                                _ = v34
                            default:
                                if (255 >= 4294967295) {
                                    _ = len(v12)
                                    v35 := GenS2{F0: (v28 ^ 170), F1: v11}
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    v2 = 0.0
                                    _ = v35
                                } else {
                                    var v36 bool = (!false)
                                    if true {
                                        fmt.Println(v16)
                                    } else {
                                        v37 := GenS2{F0: 128, F1: 30}
                                        v38 := GenS1{F0: 2, F1: 42, F2: (v33 * v33)}
                                        var v39 int = (fi32 % ((fi32) | 1))
                                        var v40 uint64 = v19
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v36
                                }
                            }
                            _ = fi32
                            _ = v33
                        }
                        v16 = 42
                        _ = fi31
                    }
                }
                _ = v28
            }
        }
        v19 += (v19 >> ((0) & 63))
        v3 = "foo bar"
    }
    v12 = append(v12, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
