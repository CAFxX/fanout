package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int }
func genFn0(p0 uint, p1 int32) uint64 {
    var v7 uint = p0
    var v8 float32 = 0.1
    if (!true) {
        p0 = v7
    } else {
        v9 := GenS1{F0: true}
        v10 := GenS2{F0: 42, F1: 1}
        v11 := GenS1{F0: ((0.0 - ((((0.1 - 1.5) / float64(1e10)) + float64(3.14159)) + 0.0)) <= 0.1)}
        _ = v10
        _ = v11
        _ = v9
    }
    v8 -= ((v8 / v8) - (v8 - (((1.5 + v8) / (v8 + (v8 / v8))) - (v8 + -2.25))))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 1
    var v3 int64 = 616
    v4 := []uint32{}
    var v5 int64 = (v3 + (v3 >> (((v1 - v1)) & 63)))
    var v6 float64 = ((float64(0.1) * -2.25) - -2.25)
    if (65535 != 32767) {
        v2 += 1
        v12 := GenS2{F0: -77, F1: v5}
        v2 = 42
        _ = v12
    } else {
        switch 127 {
        case 9223372036854775807:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 127:
            v5 = (v5 / (((v5 / (((-(v3))) | 1))) | 1))
        case -2147483648:
            var v13 uint8 = 253
            _ = v13
        default:
            fmt.Println(v5)
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            switch fi14 {
            case 9223372036854775807:
                if true {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        if true {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if false {
                                    v5 = (v5 / ((v3) | 1))
                                    switch "foo bar" {
                                    case "foo bar":
                                        v1 -= 0
                                    default:
                                        var v16 int8 = 127
                                        _ = v16
                                    }
                                } else {
                                    switch ("a" + ("hello" + "")) {
                                    case "":
                                        if false {
                                            if true {
                                                v4 = append(v4, 791)
                                            }
                                            _ = len(v4)
                                            v3 = (-(v3))
                                        } else {
                                            var v17 int16 = 1
                                            var v18 uint = 558
                                            var v19 bool = (!true)
                                            var v20 uint32 = 4294967295
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                }
                            }
                            if false {
                                if false {
                                    v2 -= (v2 << ((v1) & 63))
                                    var v21 uint16 = 1
                                    _ = v21
                                } else {
                                    v3 = 9223372036854775807
                                    v4 = append(v4, 2)
                                }
                            }
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = len(v4)
                        _ = fi15
                    }
                }
            case -2147483648:
                for fi22 := 0; fi22 < 4; fi22++ {
                    switch (fi22 + fi22) {
                    case 127:
                        fmt.Println(v5)
                    case 32768:
                        var v23 uint16 = 1
                        _ = v23
                    default:
                        v6 += v6
                    }
                    _ = fi22
                }
            case 127:
                fmt.Println(v2, v1)
            default:
                for fi24 := 0; fi24 < 5; fi24++ {
                    v1 += 18446744073709551615
                    var v25 int8 = -128
                    _ = fi24
                    _ = v25
                }
            }
            if ((v2 % ((genFn0(683, -569)) | 1)) == (v2 / ((1) | 1))) {
                if true {
                    fmt.Println(v2)
                }
            } else {
                v26 := GenS3{F0: 42}
                if true {
                    v27 := GenS1{F0: true}
                    v5 += v5
                    _ = v27
                }
                fmt.Println(fi14, v5, v5)
                _ = v26
            }
            _ = fi14
        }
        v28 := GenS2{F0: 42, F1: v3}
        _ = v28
    }
    var v29 int32 = 0
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS3{F0: 255}
        var v32 int8 = 42
        v1 = (v1 + v1)
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v29, v5)
    var v33 uint8 = 8
    v3 = v5
    v3 += v3
    v2 = v2
    if false {
        _ = len(v4)
        v34 := GenS1{F0: false}
        var v35 int16 = 1
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        v33 = v33
        v37 := GenS4{F0: 206, F1: (fi36 & (fi36 * fi36))}
        _ = fi36
        _ = v37
    }
    var v38 uint = (v1 / ((v1) | 1))
    v39 := GenS2{F0: -122, F1: (v5 - v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
