package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    var v4 float32 = 0.1
    v5 := make(map[uint32]string)
    var v6 int32 = -1
    fmt.Println(v4, v6)
    v3 = (v3 % ((v3) | 1))
    switch -13 {
    case 32767:
        var v7 uint = 890
        _ = v7
    case 697:
        v3 += (v3 << ((65535) & 63))
    case -1:
        var v8 uint32 = 336
        _ = v8
    }
    var v9 uint16 = 32767
    var v10 string = "a"
    delete(v5, 65535)
    _ = len(v5)
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int16 = 42
        if true {
            switch v10 {
            case "foo bar":
                v3 = (v3 / ((196) | 1))
            case "":
                if false {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        fmt.Println(v9, v3, v6)
                        fmt.Println(v3, v9, fi13)
                        var v14 string = "foo bar"
                        _ = fi13
                        _ = v14
                    }
                    v15 := []int{int(v6), (-(fi11)), fi11}
                    _ = v15
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        if (false && true) {
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v17 := []int{(fi16 >> ((368) & 63))}
                                if false {
                                    v3 = (v3 * v3)
                                } else {
                                    v18 := v5[84]
                                    _ = v18
                                }
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            fmt.Println(fi11, v1)
                            if (v9 == (v9 & 65535)) {
                                switch "" {
                                case "a":
                                    for fi19 := 0; fi19 < 4; fi19++ {
                                        if (1e10 >= v4) {
                                            if (2 >= 1) {
                                                delete(v5, v1)
                                            }
                                            if (!false) {
                                                v12 += v12
                                                v9 = v9
                                                v4 -= 0.0
                                            } else {
                                                fmt.Println(v1, v3, fi19)
                                                var v20 int64 = 379
                                                if (float32(0.1) >= v4) {
                                                    if (true && true) {
                                                        v1 = (v1 & 4294967295)
                                                        v4 += float32(((((0.1 * (float64(float32(-2.25)) * 3.14159)) / (0.1 + 1e10)) + (1e10 * (1e10 / ((0.0 - 1.5) * float64((v4 / v4)))))) / 0.1))
                                                        fmt.Println(fi19)
                                                    }
                                                }
                                                _ = v20
                                            }
                                            v21 := GenS1{F0: 127, F1: (100 != -66), F2: -62}
                                            v10 = v10
                                            _ = v21
                                        } else {
                                            v10 = (v10 + (v10 + (v10 + (v10 + v10))))
                                        }
                                        _ = fi19
                                    }
                                default:
                                    v22 := GenS1{F0: -118, F1: true, F2: 42}
                                    _ = v22
                                }
                            }
                            fmt.Println(v12)
                        }
                    case "":
                        v3 -= v3
                    case "hello":
                        v9 = v9
                    }
                    v9 = 432
                }
            case "a":
                fmt.Println(v3)
            default:
                switch (v10 + ((v10 + (v10 + v10)) + (v10 + v10))) {
                case "foo bar":
                    switch "hello" {
                    case "a":
                        for fi23 := 0; fi23 < 7; fi23++ {
                            var v24 float64 = 1.5
                            var v25 uint16 = (v9 & 32767)
                            v26 := []int{(fi11 | fi11)}
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case "":
                        v10 = "a"
                    default:
                        var v27 int8 = -128
                        _ = v27
                    }
                }
            }
        } else {
            delete(v5, 2147483648)
            var v28 uint = 65535
            var v29 uint8 = v3
            _ = v28
            _ = v29
        }
        v30 := GenS1{F0: 100, F1: (42 != 100), F2: -128}
        _ = fi11
        _ = v12
        _ = v30
    }
    v31 := []int{}
    if (true || false) {
        if (true && false) {
            v6 -= (v6 + v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
