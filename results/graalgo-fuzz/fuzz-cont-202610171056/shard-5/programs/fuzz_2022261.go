package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 int32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []int{2147483647}
    var v2 uint64 = 147
    var v3 uint8 = 1
    var v4 int64 = -404
    _ = len(v1)
    var v5 int = 128
    var v6 int64 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v7 := GenS1{F0: v3}
        v1 = append(v1, -213)
        v2 = (^(v2))
        _ = v7
    } else {
        if ((^(v3)) != v3) {
            v8 := GenS1{F0: (v3 + (v3 ^ (v3 ^ v3)))}
            _ = v8
        } else {
            if true {
                v9 := GenS1{F0: (v3 + (v3 | (v3 * (v3 ^ v3))))}
                v3 = v3
                _ = v9
            } else {
                switch ("foo bar" + "hello") {
                case "":
                    v2 += (v2 * 2)
                case "foo bar":
                    fmt.Println(v5, v5)
                default:
                    var v10 bool = false
                    _ = v10
                }
                v11 := GenS1{F0: v3}
                _ = v11
            }
            var v12 uint64 = (v2 | v2)
            _ = v12
        }
    }
    if (18446744073709551615 > 0) {
        switch (^(v5)) {
        case 127:
            var v13 bool = (-647 <= -9223372036854775808)
            _ = v13
        case -751:
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint32 = 92
                if (255 < 23) {
                    var v16 float64 = (-2.25 / 3.14159)
                    v17 := GenS2{F0: true, F1: (v6 - 42), F2: (v6 & -1)}
                    var v18 int64 = v4
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v2 += (v2 << ((391) & 63))
                }
                _ = fi14
                _ = v15
            }
        case 32768:
            v5 -= v5
        }
    } else {
        v1 = append(v1, v5)
        fmt.Println(v2, v2)
    }
    switch ("hello" + "") {
    case "foo bar":
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v3, v4, v3)
            if true {
                for fi20 := 0; fi20 < 12; fi20++ {
                    fmt.Println(fi20, v2, v6)
                    switch fi19 {
                    case 32767:
                        v1 = append(v1, fi20)
                    case 127:
                        if false {
                            if ((v2 ^ 1) > 556) {
                                var v21 bool = (255 < 2)
                                if v21 {
                                    if false {
                                        v6 = 225
                                        v6 += int64(v3)
                                        var v22 bool = false
                                        _ = v22
                                    }
                                    fmt.Println(v5)
                                }
                                var v23 float32 = (1.5 + 3.14159)
                                _ = v21
                                _ = v23
                            }
                            var v24 uint16 = 813
                            _ = v24
                        } else {
                            v5 = fi19
                            v25 := GenS1{F0: (v3 | 108)}
                            v26 := GenS3{F0: v5, F1: -2147483648, F2: 255}
                            _ = v25
                            _ = v26
                        }
                    }
                    v5 -= fi19
                    _ = fi20
                }
                if true {
                    var v27 uint = 876
                    _ = v27
                } else {
                    var v28 int = 0
                    _ = v28
                }
            } else {
                v4 = (v6 >> ((42) & 63))
            }
            _ = fi19
        }
    default:
        for fi29 := 0; fi29 < 10; fi29++ {
            fmt.Println(v3, v3, fi29)
            for fi30 := 0; fi30 < 20; fi30++ {
                fmt.Println(v3, v5)
                _ = fi30
            }
            _ = fi29
        }
    }
    switch (v5 % (((v5 + 1)) | 1)) {
    case -814:
        v31 := GenS4{F0: 65535}
        _ = v31
    default:
        for fi32 := 0; fi32 < 18; fi32++ {
            if (1e10 > ((((3.14159 * (-2.25 - (0.0 + ((0.0 * -2.25) + (0.0 / (0.0 - ((0.1 + (1e10 + ((0.1 * (3.14159 + (0.1 / (-2.25 + 0.0)))) * (1.5 - ((1.5 / -2.25) - 0.1))))) - 1.5))))))) - 1e10) + 1e10) + 1e10)) {
                v5 -= 42
            }
            v4 -= v4
            _ = fi32
        }
    }
    v6 -= (v4 | v6)
    for fi33 := 0; fi33 < 12; fi33++ {
        v3 = (v3 ^ v3)
        v6 -= v4
        _ = fi33
    }
    var v34 int = v5
    v35 := GenS3{F0: -9223372036854775808, F1: 399, F2: 850}
    v36 := GenS5{F0: v4}
    if (!false) {
        var v37 uint32 = 1
        _ = v37
    } else {
        v5 -= 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
