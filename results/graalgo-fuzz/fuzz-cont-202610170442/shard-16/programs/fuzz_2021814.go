package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 bool }

func main() {
    v1 := []uint32{114, 2147483648, 225}
    var v2 float32 = 1e10
    var v3 int = 0
    var v4 uint = 1
    switch "" {
    case "":
        v5 := GenS1{F0: false, F1: 127}
        _ = v5
    default:
        switch (v3 & v3) {
        case 1:
            v2 = float32((float64((v2 * 0.0)) * (3.14159 / 1.5)))
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v6 bool = true
    var v7 int64 = -52
    if v6 {
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    if (v6 || v6) {
        for fi8 := 0; fi8 < 9; fi8++ {
            switch (fi8 & 539) {
            case 0:
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 float64 = 1.5
                    if ("foo bar" > "\n\t") {
                        if false {
                            var v11 uint32 = 0
                            var v12 uint16 = 255
                            switch fi8 {
                            case -776:
                                v13 := GenS2{F0: v4, F1: v4}
                                _ = v13
                            }
                            _ = v11
                            _ = v12
                        }
                    } else {
                        _ = len(v1)
                        v10 -= (v10 / (3.14159 * 1.5))
                    }
                    v3 += (fi9 << ((48) & 63))
                    _ = fi9
                    _ = v10
                }
            case 4294967295:
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 uint = v4
                    var v16 int = (fi8 % ((fi14) | 1))
                    fmt.Println(v2, fi14)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case -1:
                var v17 uint64 = 4294967295
                _ = v17
            default:
                switch v3 {
                case 32767:
                    v7 = v7
                case -1:
                    if (v4 >= (v4 - (v4 | v4))) {
                        fmt.Println(v3, v4, fi8)
                        v18 := GenS1{F0: v6, F1: 255}
                        _ = v18
                    }
                }
            }
            _ = fi8
        }
        v3 = v3
    } else {
        v19 := GenS2{F0: v4, F1: (v4 ^ v4)}
        if false {
            if (true && (85 >= 255)) {
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 int8 = -48
                    _ = fi20
                    _ = v21
                }
            }
        } else {
            if v6 {
                v22 := GenS3{F0: 10, F1: 2, F2: 835}
                if true {
                    v6 = (!v6)
                    v2 = float32(1.5)
                    v23 := GenS4{F0: (v3 | (v3 + (v3 % (((v3 << (((^(v4))) & 63))) | 1)))), F1: 18446744073709551615}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        fmt.Println(v6)
                        v25 := GenS5{F0: v6}
                        _ = fi24
                        _ = v25
                    }
                    if true {
                        v2 = v2
                        if false {
                            var v26 int16 = -27
                            var v27 string = ("" + ("" + "\n\t"))
                            _ = v26
                            _ = v27
                        } else {
                            v4 -= v4
                        }
                    } else {
                        switch (v3 << ((v4) & 63)) {
                        case 32767:
                            if (40 <= 817) {
                                var v28 int16 = 32767
                                v7 -= (^(v7))
                                _ = v28
                            } else {
                                var v29 int64 = (v7 * v7)
                                var v30 uint8 = 255
                                v31 := GenS3{F0: (v30 / (((^(v30))) | 1)), F1: 898, F2: 937}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        fmt.Println(v7)
                    }
                }
                _ = v22
            }
            if v6 {
                v1 = append(v1, 4294967295)
                v1 = append(v1, 4294967295)
            }
        }
        v32 := GenS4{F0: 4294967295, F1: 18446744073709551615}
        _ = v19
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 float32 = v2
        v34 = (0.0 * v2)
        for fi35 := 0; fi35 < 9; fi35++ {
            v3 -= (fi33 - v3)
            fmt.Println(v2, fi33, v34)
            _ = fi35
        }
        _ = fi33
        _ = v34
    }
    v1 = append(v1, 4)
    v3 = (-(v3))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 = ((v3 >= v3) && (v4 < (v4 >> ((255) & 63))))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v36 float32 = (float32((1e10 * (-2.25 / ((0.0 * (1.5 / (0.0 + (((float64((float32(float64((-2.25 + 3.14159))) * v2)) + (3.14159 - 1e10)) * 0.1) + (0.0 / 3.14159))))) - 0.1)))) + 0.1)
    var v37 int32 = 42
    if len(v1) > 0 {
        v1[0] = 759
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 -= (-(v7))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v6
    _ = v7
}
