package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }

func main() {
    v1 := []uint32{0}
    var v2 int8 = 127
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 int = 284
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 -= fi6
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v7 := GenS1{F0: "foo bar", F1: 2, F2: 1}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint32 = 4294967295
        if true {
            var v10 float32 = 1e10
            if false {
                var v11 int = (^(v5))
                var v12 float64 = 1e10
                _ = v11
                _ = v12
            } else {
                fmt.Println(v10, v5)
                v2 += v2
                v2 = 0
            }
            switch v5 {
            case 2147483647:
                fmt.Println(v9)
            }
            _ = v10
        }
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if (42 == 4294967295) {
                v5 += -9223372036854775808
                switch "" {
                case "\n\t":
                    if (-1 == 32767) {
                        v3[(^(fi14))] = false
                        var v15 int = -53
                        _ = v15
                    }
                case "foo bar":
                    fmt.Println(v5, v2, fi14)
                default:
                    var v16 float32 = ((1e10 + (float32(((0.0 / 1.0) * 1e10)) / 3.14159)) * -2.25)
                    _ = v16
                }
            }
            _ = fi14
        }
        v17 := GenS1{F0: ("foo bar" + "hello"), F1: 1, F2: 1}
        _ = fi13
        _ = v17
    }
    fmt.Println(v2)
    var v18 int8 = -109
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint16 = 0
        if true {
            var v21 string = (("hello" + "a") + "a")
            _ = v21
        } else {
            v22 := v4[-1]
            var v23 int16 = 0
            for fi24 := 0; fi24 < 12; fi24++ {
                if (0.0 == -2.25) {
                    switch (v22 + (v22 + "\n\t")) {
                    case "hello":
                        v23 -= (v23 >> ((195) & 63))
                    default:
                        v20 += v20
                    }
                    switch fi19 {
                    case -9223372036854775808:
                        if (!false) {
                            var v25 int = (v5 & v5)
                            _ = len(v1)
                            _ = v25
                        } else {
                            v2 += (v2 & v18)
                            switch fi19 {
                            case 2147483647:
                                if true {
                                    var v26 uint8 = 127
                                    fmt.Println(v22, v22)
                                    var v27 uint64 = 18446744073709551615
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 int64 = 454
                                    _ = v28
                                }
                            case 1:
                                var v29 uint16 = 1
                                _ = v29
                            case -732:
                                var v30 int64 = -1
                                _ = v30
                            default:
                                v18 -= (^(v18))
                            }
                        }
                    case 32768:
                        if len(v1) > 0 {
                            v1[0] = 572
                        }
                    case 0:
                        var v31 int8 = (v2 ^ (v18 % ((v18) | 1)))
                        _ = v31
                    default:
                        _ = len(v3)
                    }
                } else {
                    v18 -= (v18 + -115)
                    var v32 uint8 = 42
                    _ = v32
                }
                _ = fi24
            }
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    var v33 bool = (!true)
    var v34 int64 = 1
    v2 -= (v18 << ((136) & 63))
    v35 := GenS1{F0: "\n\t", F1: 1, F2: 782}
    if len(v1) > 0 {
        v1[0] = 586
    }
    if (!((1.5 + 0.1) != 3.14159)) {
        v1 = append(v1, 2)
        var v36 int = v5
        if len(v1) > 0 {
            v1[0] = 451
        }
        _ = v36
    }
    _ = len(v4)
    fmt.Println(v18, v5, v33)
    if (1.5 >= 0.1) {
        v33 = (v34 == (-(v34)))
        var v37 int32 = 2147483647
        if v33 {
            v4[v5] = ("foo bar" + "a")
            var v38 uint16 = 255
            _ = v38
        }
        _ = v37
    } else {
        if ((-32768 > 613) && ((239 > 128) && v33)) {
            if v33 {
                if (!v33) {
                    delete(v3, v5)
                    var v39 int8 = (v18 - int8(v5))
                    _ = v39
                }
            }
        }
        var v40 string = "a"
        _ = v40
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
