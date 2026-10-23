package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int32; F1 string; F2 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = (((0.1 / (3.14159 - 0.1)) * ((-2.25 + float64(0.1)) * float64(0.0))) / -2.25)
    var v3 int8 = 72
    var v4 int32 = -2147483648
    v5 := make(map[uint32]uint32)
    _ = len(v5)
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := v5[2147483647]
            fmt.Println(v8)
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (42 > 126), F1: (-(v3)), F2: 42}
        for fi10 := 0; fi10 < 8; fi10++ {
            if (2147483648 <= 798) {
                v11 := GenS2{F0: true, F1: (v1 ^ 1)}
                v5[629] = 2147483648
                v3 -= v3
                _ = v11
            } else {
                v4 = (v4 + v4)
            }
            _ = fi10
        }
        _ = fi6
        _ = v9
    }
    if false {
        v1 += int64(v4)
        v4 = v4
    }
    v4 = (-(v4))
    v12 := []int{int(v3), 2147483647}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    switch 9223372036854775807 {
    case 32767:
        fmt.Println(v4)
    default:
        var v13 uint8 = 255
        _ = v13
    }
    if false {
        v14 := v5[1]
        var v15 string = ("a" + "hello")
        var v16 float64 = (float64(0.1) / v2)
        _ = v14
        _ = v15
        _ = v16
    }
    fmt.Println(v1, v2)
    for fi17 := 0; fi17 < 7; fi17++ {
        if false {
            if false {
                var v18 int32 = (v4 + v4)
                _ = v18
            } else {
                if (127 < v3) {
                    if true {
                        v19 := GenS3{F0: 194, F1: 0, F2: 617}
                        _ = v19
                    }
                    if (2 != 255) {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            switch ("\n\t" + ("foo bar" + "foo bar")) {
                            case "\n\t":
                                v4 -= 2147483647
                            case "a":
                                _ = len(v12)
                            default:
                                v3 += 35
                            }
                            v3 -= -25
                            _ = fi20
                        }
                    }
                    if ((413 != 255) || false) {
                        v21 := GenS3{F0: 56, F1: 358, F2: 2147483648}
                        _ = v21
                    }
                }
                if ((false || (v2 >= (float64(0.0) / v2))) && false) {
                    if (!false) {
                        v1 = -203
                        v22 := GenS4{F0: -42, F1: fi17, F2: 32767}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v24 := GenS1{F0: (255 < 221), F1: (v3 + -57), F2: 0}
                            switch (fi17 ^ (fi23 / ((0) | 1))) {
                            case -60:
                                delete(v5, 171)
                            case 2147483647:
                                v4 = v4
                            }
                            switch fi17 {
                            case -1:
                                _ = len(v5)
                            default:
                                v4 -= int32(fi23)
                            }
                            v4 += (v4 / ((-399) | 1))
                            fmt.Println(v4, fi23, v4)
                            if (0 != 255) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                fmt.Println(v1)
                            } else {
                                v4 -= 1
                            }
                            if false {
                                var v25 int32 = v4
                                _ = v25
                            } else {
                                var v26 string = ("foo bar" + "\n\t")
                                _ = v26
                            }
                            if ((65535 == 2147483647) && (!true)) {
                                v1 = v1
                                v2 += float64(3.14159)
                                v4 -= (v4 / (((-(v4))) | 1))
                            } else {
                                v2 += (v2 / (float64(-2.25) / (float64((-2.25 + 0.1)) - v2)))
                            }
                            _ = fi23
                            _ = v24
                        }
                    }
                } else {
                    v27 := GenS3{F0: 619, F1: v1, F2: 2}
                    _ = v27
                }
                switch "hello" {
                case "":
                    if true {
                        delete(v5, 65535)
                    } else {
                        v28 := GenS5{F0: 127}
                        var v29 float64 = v2
                        _ = v28
                        _ = v29
                    }
                }
            }
            switch "foo bar" {
            case "\n\t":
                if (false && true) {
                    var v30 bool = true
                    switch -359 {
                    case 2147483647:
                        var v31 bool = (!(false && (0 < 4294967295)))
                        _ = v31
                    }
                    var v32 int64 = (^(v1))
                    _ = v30
                    _ = v32
                }
            case "foo bar":
                v33 := GenS6{F0: (v3 / ((43) | 1)), F1: 42}
                _ = v33
            case "hello":
                var v34 bool = (true || false)
                _ = v34
            default:
                v1 -= v1
            }
            v3 -= (v3 % (((v3 & (^(v3)))) | 1))
        } else {
            _ = len(v5)
            switch fi17 {
            case 4294967295:
                fmt.Println(v1)
            case 2147483647:
                v4 -= -2147483648
            case 789:
                fmt.Println(v3)
            default:
                fmt.Println(v1, v1)
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := GenS7{F0: (-(v4)), F1: "hello", F2: 296}
                if (v2 != v2) {
                    v12 = append(v12, (fi17 >> ((769) & 63)))
                    v4 = v4
                }
                v2 = (v2 / v2)
                _ = fi35
                _ = v36
            }
        }
        _ = fi17
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint64 = 155
        fmt.Println(v2)
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
