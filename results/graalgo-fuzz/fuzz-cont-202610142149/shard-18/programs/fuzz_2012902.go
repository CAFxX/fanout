package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 uint16 }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    var v3 uint64 = 18446744073709551615
    var v4 int8 = 127
    v4 += (v4 + v4)
    v4 = 42
    v5 := []int{-967, 42, 510, -90}
    if len(v5) > 0 {
        v5[0] = 42
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        switch fi6 {
        case 633:
            var v7 string = "\n\t"
            _ = v7
        default:
            var v8 float32 = (1.5 - -2.25)
            _ = v8
        }
        if (182 > 127) {
            var v9 bool = false
            _ = v9
        } else {
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                var v10 uint16 = 0
                _ = v10
            case "a":
                var v11 uint = (v1 >> ((v1) & 63))
                _ = v11
            default:
                v1 = (v1 | 1)
            }
            if (1e10 == (3.14159 * 0.0)) {
                var v12 int64 = 1
                if true {
                    v1 = v1
                    if true {
                        if (!false) {
                            var v13 int16 = 1000
                            _ = v13
                        } else {
                            fmt.Println(fi6, v1)
                        }
                        switch -371 {
                        case -1:
                            v1 = (v1 & 65535)
                        case -9223372036854775808:
                            fmt.Println(v12)
                        default:
                            for fi14 := 0; fi14 < 4; fi14++ {
                                _ = len(v5)
                                v4 += -128
                                _ = fi14
                            }
                        }
                        switch fi6 {
                        case -761:
                            var v15 string = "hello"
                            _ = v15
                        case 127:
                            v3 -= (^(v3))
                        case 32767:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    } else {
                        var v16 bool = false
                        for fi17 := 0; fi17 < 18; fi17++ {
                            var v18 float64 = 1e10
                            _ = fi17
                            _ = v18
                        }
                        var v19 int16 = 1000
                        _ = v16
                        _ = v19
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    fmt.Println(v3, v1, v3)
                    switch "\n\t" {
                    case "":
                        v1 = (v1 << ((45) & 63))
                    case "\n\t":
                        var v20 string = ""
                        _ = v20
                    case "hello":
                        v12 = v12
                    }
                }
                _ = v12
            } else {
                v4 = 100
            }
        }
        _ = fi6
    }
    v21 := GenS1{F0: (v3 & (^(v3)))}
    var v22 float64 = float64(float32(3.14159))
    v23 := v2[-627]
    if len(v5) > 0 {
        v5[0] = -8
    }
    if (true && (!false)) {
        switch 2147483647 {
        case -9223372036854775808:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        v24 := GenS2{F0: 0, F1: int8(v1)}
        v25 := GenS2{F0: 42, F1: (v4 << ((18446744073709551615) & 63))}
        var v26 float32 = float32(1.5)
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println(v4, v4, v22)
    v27 := GenS2{F0: 2, F1: (v4 ^ (v4 | v4))}
    if true {
        fmt.Println(v4)
        fmt.Println(v4, v3, v1)
    }
    v1 -= v1
    v3 += 170
    switch (v23 + v23) {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 506
        }
    case "hello":
        var v28 uint16 = 255
        _ = v28
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = 42
        }
    default:
        var v29 uint32 = 2
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        if (42 < 127) {
            fmt.Println(v23, v4)
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 uint16 = 1
                switch v23 {
                case "hello":
                    switch fi30 {
                    case 852:
                        var v33 uint32 = 65535
                        _ = v33
                    default:
                        v34 := GenS2{F0: 128, F1: v4}
                        _ = v34
                    }
                case "\n\t":
                    var v35 uint8 = 58
                    _ = v35
                default:
                    v32 = 32767
                }
                switch v23 {
                case "a":
                    v1 += v1
                case "\n\t":
                    fmt.Println(v32)
                case "hello":
                    v36 := GenS3{F0: fi30, F1: uint32(v32), F2: 32767}
                    _ = v36
                }
                _ = fi31
                _ = v32
            }
            fmt.Println(v22, v22)
        }
        if (!((125 >= uint8(v1)) || (3.14159 < float32(((3.14159 / v22) - v22))))) {
            v1 += 42
            fmt.Println(v23)
        } else {
            var v37 int16 = int16(v1)
            v37 = v37
            _ = v37
        }
        _ = fi30
    }
    var v38 uint = 303
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
