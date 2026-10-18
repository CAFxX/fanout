package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int8 = 100
    v2 := []uint32{}
    var v3 bool = false
    var v4 bool = ((!((1.5 - ((((1.5 - (((1.5 / 1.0) - 0.0) * 0.0)) * 1.5) - ((0.0 + float32(0.0)) * 3.14159)) * 3.14159)) == float32(float64(0.0)))) || false)
    v5 := GenS1{F0: v1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((!true) && v3) {
        v6 := GenS1{F0: (-(v1))}
        v1 = v1
        if (true && v4) {
            if (65535 <= 1) {
                fmt.Println(v4)
            } else {
                v4 = (!false)
                fmt.Println(v1, v4, v1)
                v7 := GenS1{F0: (v1 & (v1 - (v1 / ((v1) | 1))))}
                _ = v7
            }
        } else {
            if true {
                _ = len(v2)
                fmt.Println(v1, v3, v1)
                v3 = v4
            }
        }
        _ = v6
    } else {
        var v8 int = -719
        if (int32(v8) != -691) {
            if (!v3) {
                v8 = 128
                var v9 uint = 18446744073709551615
                var v10 int32 = -1
                _ = v10
                _ = v9
            } else {
                _ = len(v2)
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 uint8 = 158
                    _ = fi11
                    _ = v12
                }
            }
        }
        _ = v8
    }
    switch "\n\t" {
    case "\n\t":
        v13 := GenS1{F0: v1}
        _ = v13
    case "hello":
        var v14 uint8 = 128
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS2{F0: 507, F1: 0}
            for fi17 := 0; fi17 < 15; fi17++ {
                if (871 < 0) {
                    switch ((("hello" + "\n\t") + "a") + "hello") {
                    case "a":
                        if (1000 < -161) {
                            var v18 bool = (!v3)
                            v1 = v1
                            if true {
                                v2 = append(v2, 1)
                                if v18 {
                                    v4 = v18
                                }
                            }
                            _ = len(v2)
                            _ = v18
                        } else {
                            if v3 {
                                var v19 uint64 = 42
                                if v4 {
                                    if false {
                                        v20 := GenS2{F0: 42, F1: 255}
                                        var v21 uint8 = 128
                                        _ = v20
                                        _ = v21
                                    } else {
                                        fmt.Println(fi17, v1)
                                    }
                                    v2 = append(v2, 23)
                                }
                                v22 := GenS3{F0: 100, F1: 191, F2: 0}
                                _ = v19
                                _ = v22
                            } else {
                                v2 = append(v2, 42)
                                if (v1 < -95) {
                                    v4 = (!((!(!(!false))) || ("" <= "")))
                                    var v23 uint16 = 102
                                    v23 = 745
                                    _ = v23
                                }
                            }
                        }
                    }
                } else {
                    v24 := GenS4{F0: (!v3)}
                    var v25 uint = 255
                    _ = v24
                    _ = v25
                }
                var v26 float64 = 1e10
                v27 := GenS3{F0: (v1 + v1), F1: 65535, F2: 764}
                _ = fi17
                _ = v26
                _ = v27
            }
            _ = fi15
            _ = v16
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v28 int8 = v1
    if ((!(v3 && (2 != 2147483647))) && v3) {
        v2 = append(v2, 502)
        switch 0 {
        case -1:
            if ((v28 << ((255) & 63)) >= v1) {
                v3 = v3
            }
        default:
            v1 = (v1 * v1)
        }
    }
    var v29 uint32 = 42
    for fi30 := 0; fi30 < 4; fi30++ {
        switch "" {
        case "\n\t":
            v29 = 667
        default:
            switch 0 {
            case -923:
                v29 -= (^(v29))
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        v31 := GenS4{F0: (!((!((0.1 * 3.14159) > ((float32((1e10 / (((float64(1e10) * (((0.1 * 1e10) + 0.1) * 0.1)) - float64(0.0)) * float64(1.5)))) - (0.0 * 0.0)) / (((0.0 - 1.5) + float32(1e10)) + ((3.14159 * -2.25) + 0.0))))) && v4))}
        _ = fi30
        _ = v31
    }
    v3 = (1 > -1)
    var v32 int64 = 1
    fmt.Println(v4, v4, v1)
    v32 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
