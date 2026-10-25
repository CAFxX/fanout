package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 bool; F2 bool }

func main() {
    v1 := []string{"\n\t", ("foo bar" + ("" + "a")), ""}
    var v2 float64 = (((0.1 - (((-2.25 / (((0.1 * 1e10) * 1e10) / ((((((((0.1 + 1.5) + 0.0) * -2.25) - float64(1e10)) + ((float64(float32((1.5 * 1e10))) - 0.0) * 1e10)) * 1e10) + ((((0.0 - 3.14159) + 0.0) / -2.25) - (0.0 / 0.1))) - 1e10))) + (3.14159 * 0.0)) + (((0.1 * 1e10) * float64((1e10 + -2.25))) + 3.14159))) / 1.0) / 1.0)
    v3 := []string{"", "\n\t"}
    v4 := make(map[uint32]int)
    var v5 uint = 18446744073709551615
    var v6 int = 127
    delete(v4, 4)
    v5 = (v5 | v5)
    fmt.Println(v6, v6)
    switch "foo bar" {
    case "hello":
        v2 = v2
    case "foo bar":
        fmt.Println(v2, v6, v2)
    case "\n\t":
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint = (v5 % ((v5) | 1))
            var v9 uint = (v5 | (^(v5)))
            switch (-(v6)) {
            case 2147483647:
                switch 42 {
                case 4294967295:
                    var v10 uint64 = 42
                    _ = v10
                case 127:
                    switch "a" {
                    case "a":
                        var v11 uint64 = 0
                        _ = v11
                    case "":
                        v12 := GenS1{F0: 1, F1: 9223372036854775807, F2: -965}
                        _ = v12
                    case "hello":
                        switch 9223372036854775807 {
                        case 32767:
                            var v13 string = ("\n\t" + "a")
                            _ = v13
                        default:
                            v6 += (fi7 / ((v6) | 1))
                        }
                    default:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if true {
                                fmt.Println(v9, v9, v5)
                                v15 := v4[2147483648]
                                _ = v15
                            } else {
                                fmt.Println(v9)
                                delete(v4, 4294967295)
                            }
                            _ = fi14
                        }
                    }
                }
            case -981:
                fmt.Println(v9, v5, v2)
            default:
                var v16 float64 = 0.0
                _ = v16
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        if true {
            switch "" {
            case "foo bar":
                if ((!(false && (!((v2 - v2) >= v2)))) && true) {
                    v17 := GenS1{F0: -13, F1: -693, F2: 1}
                    _ = len(v4)
                    _ = v17
                } else {
                    v6 += v6
                    if false {
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v6 = v6
                            for fi19 := 0; fi19 < 10; fi19++ {
                                v2 += -2.25
                                v3 = append(v3, "")
                                v3 = append(v3, ("foo bar" + "a"))
                                _ = fi19
                            }
                            _ = fi18
                        }
                        var v20 uint = 0
                        _ = v20
                    } else {
                        v5 = (v5 ^ 115)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            case "":
                var v21 float32 = 0.0
                _ = v21
            }
        } else {
            v6 = -1
            switch (v6 & 310) {
            case -771:
                var v22 float32 = ((float32(float64(0.0)) * 1e10) + float32(v2))
                _ = v22
            case 1:
                v23 := GenS2{F0: 911, F1: 121}
                _ = v23
            case 4294967295:
                if (v2 != v2) {
                    _ = len(v1)
                    v24 := GenS3{F0: -1}
                    v25 := GenS4{F0: 0, F1: 65535, F2: -9223372036854775808}
                    _ = v24
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 7; fi26++ {
                    v6 = (v6 - 144)
                    fmt.Println(v2)
                    if false {
                        if (2 > 238) {
                            v27 := v4[211]
                            var v28 int8 = -40
                            _ = v27
                            _ = v28
                        } else {
                            _ = len(v3)
                            v29 := GenS3{F0: 0}
                            v2 = ((1e10 * float64((3.14159 / (-2.25 * 1e10)))) / ((((v2 - v2) * 3.14159) * v2) / ((v2 + v2) + (0.1 / v2))))
                            _ = v29
                        }
                        var v30 uint8 = 128
                        _ = v30
                    }
                    _ = fi26
                }
            }
        }
    }
    var v31 bool = true
    v32 := GenS5{F0: -315, F1: (!(9 < 113)), F2: true}
    var v33 uint64 = 0
    var v34 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
