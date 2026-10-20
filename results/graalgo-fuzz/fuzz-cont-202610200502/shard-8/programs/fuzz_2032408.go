package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 int16 = 1000
    v2 := []int{-591, 128, 32768}
    var v3 int16 = 42
    _ = len(v2)
    v4 := GenS1{F0: true}
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((true && true) && (1.5 >= (0.0 / (1e10 * float32(-2.25))))) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (127 >= -41) {
            v1 -= 1000
        } else {
            v3 = v3
            var v5 uint32 = 48
            _ = v5
        }
        if (565 >= 42) {
            v1 = (v1 | v3)
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := GenS1{F0: (948 >= 1)}
                _ = fi6
                _ = v7
            }
            if (!(!((783 > 95) || (!(!false))))) {
                fmt.Println(v1, v1)
                var v8 int16 = -181
                _ = v8
            } else {
                if (0.1 < (float64(float32(float64(float32((0.0 / (-2.25 - (3.14159 + (1.5 * ((1.5 * 0.0) - -2.25))))))))) * 0.0)) {
                    if true {
                        v1 -= 42
                        var v9 int64 = 9223372036854775807
                        _ = v9
                    } else {
                        v10 := GenS1{F0: true}
                        if len(v2) > 0 {
                            v2[0] = 682
                        }
                        v11 := GenS1{F0: true}
                        _ = v10
                        _ = v11
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 32768
                    }
                    switch ("\n\t" + "foo bar") {
                    case "":
                        if true {
                            var v12 uint8 = 128
                            if ((3.14159 * float64(0.0)) >= 1e10) {
                                var v13 uint8 = (^(v12))
                                _ = v13
                            }
                            _ = v12
                        }
                    case "\n\t":
                        var v14 int = 1
                        _ = v14
                    }
                }
                v15 := GenS1{F0: false}
                if false {
                    var v16 uint = 18446744073709551615
                    _ = v16
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (102 <= 127) {
                        v1 -= v3
                        v2 = append(v2, -1)
                        fmt.Println(v1, v3)
                    } else {
                        _ = len(v2)
                        fmt.Println(v3, v3)
                        v1 = -52
                    }
                    if false {
                        v1 -= v3
                        var v17 int16 = v1
                        _ = v17
                    }
                }
                _ = v15
            }
        }
    }
    switch "a" {
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            v3 = (v1 / (((v3 | v3)) | 1))
            v1 = (v1 + v1)
            if true {
                for fi19 := 0; fi19 < 19; fi19++ {
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            v3 = 42
                        case "a":
                            v20 := GenS2{F0: 42, F1: 127, F2: v1}
                            _ = v20
                        case "foo bar":
                            if (0.0 > 1e10) {
                                var v21 int16 = v3
                                v21 -= v1
                                fmt.Println(v3, v21, v3)
                                _ = v21
                            } else {
                                var v22 int32 = -1
                                if len(v2) > 0 {
                                    v2[0] = fi19
                                }
                                v3 = v1
                                _ = v22
                            }
                        default:
                            var v23 uint64 = 620
                            _ = v23
                        }
                    }
                    if false {
                        switch "\n\t" {
                        case "\n\t":
                            var v24 bool = false
                            _ = v24
                        case "hello":
                            v3 = v1
                        default:
                            switch ("\n\t" + "") {
                            case "foo bar":
                                switch fi19 {
                                case 32767:
                                    var v25 int = fi19
                                    _ = v25
                                case 9:
                                    if (!false) {
                                        v26 := GenS3{F0: (v3 | (v1 - (v1 ^ v1))), F1: false}
                                        v27 := GenS3{F0: v1, F1: true}
                                        _ = v26
                                        _ = v27
                                    }
                                case 42:
                                    v1 -= (v3 >> ((18446744073709551615) & 63))
                                default:
                                    v28 := GenS3{F0: (v3 - -589), F1: false}
                                    _ = v28
                                }
                            case "":
                                v1 = -32768
                            case "\n\t":
                                v3 -= (v3 | 42)
                            }
                        }
                        v1 = (v3 << ((65535) & 63))
                    }
                    _ = fi19
                }
            }
            _ = fi18
        }
    case "a":
        if (-9223372036854775808 != 0) {
            v3 -= (v1 / (((v3 * (v1 >> ((995) & 63)))) | 1))
        }
    default:
        v3 = -1
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    v3 += 0
    var v29 int = -9223372036854775808
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int8 = 42
        if (-1 == (v1 ^ v3)) {
            fmt.Println(fi30, v31, v29)
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
