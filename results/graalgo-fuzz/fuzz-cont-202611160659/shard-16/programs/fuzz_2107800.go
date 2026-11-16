package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    v3 := []string{}
    var v4 int16 = 32767
    if (!true) {
        v3 = append(v3, (("\n\t" + "foo bar") + "hello"))
    }
    v2 = ((0.1 / (((3.14159 - v2) + v2) + (((((v2 / 1.0) * 0.0) * (float64(float32((((3.14159 / 3.14159) + ((v2 * 1e10) - 0.1)) * float64(((((-2.25 * (float32(((v2 / v2) * (1e10 / v2))) - (0.1 - 1e10))) * 0.1) + (3.14159 * 1.5)) * ((((3.14159 - 0.0) - float32(v2)) + (0.0 - (0.0 / 1.0))) * 0.0)))))) / (0.1 * ((v2 + float64(float32(v2))) / (v2 / float64(3.14159)))))) / (float64(float32((-2.25 + float64((0.0 + float32((v2 * v2))))))) / v2)) / v2))) + -2.25)
    v2 = float64((float32(((0.1 * v2) / ((0.0 * float64((float32((v2 / v2)) / 1e10))) - v2))) + 1.5))
    if false {
        switch ("a" + "\n\t") {
        case "\n\t":
            v4 = 0
        }
        v2 += -2.25
        v2 += 1.5
    } else {
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := GenS1{F0: 1, F1: 0, F2: -959}
            if false {
                var v7 uint16 = 65535
                var v8 int = (fi5 & (fi5 % ((32768) | 1)))
                _ = v7
                _ = v8
            } else {
                switch "a" {
                case "\n\t":
                    if true {
                        fmt.Println(v2, v2, v2)
                        if false {
                            v4 = -372
                            v9 := GenS1{F0: 42, F1: 1, F2: 548}
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 19; fi10++ {
                            var v11 uint64 = 42
                            var v12 uint64 = 576
                            if (!(2 <= 1)) {
                                v2 -= (((v2 / (v2 * 1e10)) / ((float64((1.5 * 0.1)) / float64((0.0 / ((0.1 * ((0.1 + float32((float64(3.14159) * -2.25))) * (0.0 / (0.1 / -2.25)))) + (float32(v2) + 1.5))))) * 1.5)) - (v2 * 0.1))
                                switch "a" {
                                case "":
                                    if (int16(v11) <= -2) {
                                        v12 += 0
                                    }
                                }
                                v2 = v2
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    } else {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v14 := GenS1{F0: 0, F1: 0, F2: -63}
                                var v15 float64 = (((v2 - 0.1) / v2) * float64((-2.25 - 1.5)))
                                if (((23 == 128) || (true && (32767 <= (v4 << ((1) & 63))))) && (true && (!false))) {
                                    v15 = v2
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v16 float32 = 0.1
                                    _ = v16
                                }
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        }
                        var v17 uint8 = 93
                        var v18 float64 = v2
                        _ = v17
                        _ = v18
                    }
                case "foo bar":
                    v19 := GenS2{F0: (v4 / (((-(v4))) | 1)), F1: 42, F2: "hello"}
                    _ = v19
                }
                v20 := v1[2147483648]
                _ = v20
            }
            _ = fi5
            _ = v6
        }
    }
    if false {
        var v21 float64 = 3.14159
        _ = v21
    }
    v22 := GenS3{F0: 42, F1: v4}
    switch -1 {
    case 644:
        for fi23 := 0; fi23 < 9; fi23++ {
            switch "\n\t" {
            case "foo bar":
                for fi24 := 0; fi24 < 17; fi24++ {
                    switch -902 {
                    case -2147483648:
                        if ((!true) && (false && false)) {
                            v2 = ((0.1 + -2.25) / ((0.0 - (v2 + v2)) / v2))
                            var v25 int8 = 42
                            v4 = (-(v4))
                            _ = v25
                        } else {
                            switch "" {
                            case "foo bar":
                                var v26 int64 = 0
                                _ = v26
                            case "":
                                fmt.Println(fi24, v4, fi24)
                            default:
                                v27 := GenS3{F0: 0, F1: -1}
                                _ = v27
                            }
                        }
                    case 802:
                        if (!(fi23 == 20)) {
                            v1[852] = 21
                            var v28 string = "\n\t"
                            _ = v28
                        }
                        fmt.Println(v4, v2)
                    case 128:
                        v4 -= 1
                    }
                    if (42 >= 9) {
                        delete(v1, 1)
                        v1[42] = 2
                    } else {
                        var v29 uint32 = 226
                        v30 := GenS3{F0: 760, F1: v4}
                        _ = v29
                        _ = v30
                    }
                    var v31 int = 32768
                    _ = fi24
                    _ = v31
                }
            }
            fmt.Println(v4, fi23)
            _ = fi23
        }
    case 876:
        var v32 uint64 = 4294967295
        _ = v32
    }
    fmt.Println(v4, v2)
    v4 = v4
    var v33 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
