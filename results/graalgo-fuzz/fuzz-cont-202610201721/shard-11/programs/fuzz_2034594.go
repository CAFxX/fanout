package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0() uint {
    var v4 uint64 = 0
    fmt.Println(v4)
    v4 = (^(v4))
    var v5 uint8 = 42
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 int32, p1 int) uint8 {
    p1 = p1
    if false {
        var v6 string = "foo bar"
        _ = v6
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float64 = (0.0 - 0.1)
    v2 := []uint32{}
    var v3 float32 = (-2.25 - (3.14159 - float32(((1e10 / 1e10) * 1.5))))
    var v7 int64 = -263
    if len(v2) > 0 {
        v2[0] = 955
    }
    if false {
        v3 += v3
        var v8 int16 = -1
        _ = v8
    } else {
        if true {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (false && true) {
                    switch "hello" {
                    case "hello":
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 326
                            }
                        default:
                            v1 = ((((v1 / v1) + ((float64(float32(v1)) / ((v1 - (v1 * float64((v3 - (v3 / v3))))) / (v1 * v1))) - (v1 * 0.0))) - (1.5 / v1)) / v1)
                        }
                    case "\n\t":
                        if true {
                            var v9 float64 = v1
                            v9 += v9
                            _ = v9
                        }
                    case "foo bar":
                        _ = len(v2)
                    default:
                        var v10 int8 = 42
                        _ = v10
                    }
                    fmt.Println(v7)
                    v11 := GenS1{F0: 42, F1: 32767, F2: -913}
                    _ = v11
                }
                var v12 uint = 967
                _ = v12
            } else {
                if len(v2) > 0 {
                    v2[0] = 807
                }
                if true {
                    v7 = v7
                    for fi13 := 0; fi13 < 2; fi13++ {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            var v15 uint64 = 378
                            if (0 <= 216) {
                                var v16 string = ("foo bar" + (("\n\t" + "hello") + "foo bar"))
                                v16 = (v16 + v16)
                                fmt.Println(v16)
                                _ = v16
                            }
                            var v17 string = "\n\t"
                            _ = fi14
                            _ = v15
                            _ = v17
                        }
                        _ = fi13
                    }
                }
                switch 1 {
                case -211:
                    _ = len(v2)
                case 1:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (((v3 + v3) >= 0.1) || (556 > 76)) {
                            for fi19 := 0; fi19 < 3; fi19++ {
                                v20 := GenS1{F0: 1, F1: 42, F2: (fi18 - fi18)}
                                var v21 int8 = -117
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            var v22 uint8 = 128
                            var v23 string = "hello"
                            _ = v22
                            _ = v23
                        }
                        fmt.Println(v3, v3, fi18)
                        _ = fi18
                    }
                default:
                    v1 += (v1 + v1)
                }
            }
            var v24 int16 = -32768
            var v25 uint32 = 389
            _ = v24
            _ = v25
        }
    }
    if (2147483648 < 972) {
        if (false || false) {
            switch -2147483648 {
            case -1:
                for fi26 := 0; fi26 < 9; fi26++ {
                    if (690 == -32768) {
                        var v27 float32 = v3
                        _ = v27
                    } else {
                        fmt.Println(v3)
                    }
                    switch (fi26 + fi26) {
                    case -2147483648:
                        v28 := GenS2{F0: -1}
                        _ = v28
                    default:
                        fmt.Println(v3, v7, v7)
                    }
                    switch fi26 {
                    case 42:
                        var v29 int64 = v7
                        _ = v29
                    default:
                        v30 := GenS3{F0: (fi26 & 42)}
                        _ = v30
                    }
                    _ = fi26
                }
            case 32768:
                switch "hello" {
                case "foo bar":
                    v7 += (v7 / (((v7 * (-(v7)))) | 1))
                }
            }
            var v31 uint32 = 884
            _ = v31
        } else {
            v2 = append(v2, 2)
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v1, v1)
                _ = fi32
            }
            v2 = append(v2, 327)
        }
        v3 += 3.14159
        _ = len(v2)
    } else {
        v3 = ((v3 / (v3 - (((-2.25 - float32(v1)) - float32(v1)) / v3))) - v3)
    }
    v3 -= (v3 * (-2.25 + v3))
    if (!(v7 < v7)) {
        v1 -= v1
        fmt.Println(v7, v3, v7)
        v1 = (v1 - ((float64(0.0) / float64(-2.25)) * ((-2.25 / (0.0 + v1)) * v1)))
    } else {
        if ((1 < genFn1(-2147483648, 9223372036854775807)) || (2147483647 < 644)) {
            v33 := GenS4{F0: 2147483648, F1: 32767, F2: -1}
            _ = v33
        }
        var v34 bool = false
        var v35 int64 = -9223372036854775808
        _ = v34
        _ = v35
    }
    var v36 uint16 = 384
    switch 2147483647 {
    case -812:
        fmt.Println(v1)
    }
    var v37 uint8 = 1
    switch "\n\t" {
    case "a":
        var v38 int16 = 42
        _ = v38
    case "hello":
        if true {
            var v39 string = ""
            _ = v39
        }
    case "":
        fmt.Println(v3, v37)
    default:
        var v40 float64 = v1
        _ = v40
    }
    v3 = (-2.25 / 0.1)
    var v41 uint8 = v37
    v2 = append(v2, 4294967295)
    if len(v2) > 0 {
        v2[0] = 73
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v41
    _ = v7
}
