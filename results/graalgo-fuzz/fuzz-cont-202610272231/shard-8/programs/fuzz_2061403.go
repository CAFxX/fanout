package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }

func main() {
    var v1 uint8 = 128
    var v2 float64 = ((0.0 / float64((((1.5 / 1e10) - 0.1) + float32(1.5)))) * float64(float32(-2.25)))
    var v3 int8 = 64
    var v4 int8 = -128
    v5 := []string{("" + "hello")}
    if (718 <= 2) {
        v6 := GenS1{F0: 0, F1: false, F2: -2147483648}
        _ = v6
    }
    var v7 uint = 255
    v7 = 65535
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 1, F1: false, F2: v8}
    for fi10 := 0; fi10 < 9; fi10++ {
        if (v7 == (^(v7))) {
            v7 += v7
        } else {
            var v11 bool = ((!false) || (-1 >= int64(v1)))
            v2 = 0.0
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v4, v1, v7)
    v12 := GenS2{F0: 584, F1: (v3 % ((0) | 1)), F2: v8}
    switch "foo bar" {
    case "foo bar":
        if true {
            if ((2 != 4294967295) && true) {
                switch -2147483648 {
                case -929:
                    v2 = v2
                case 32768:
                    for fi13 := 0; fi13 < 8; fi13++ {
                        var v14 int16 = 1
                        _ = fi13
                        _ = v14
                    }
                case 2147483647:
                    v3 -= v3
                }
            } else {
                switch 241 {
                case -1:
                    v5 = append(v5, ("foo bar" + "a"))
                case -2147483648:
                    switch -1 {
                    case 9223372036854775807:
                        v5 = append(v5, "foo bar")
                    case 826:
                        if (v2 < ((v2 - ((1e10 / (v2 / -2.25)) / (v2 / float64((0.1 / 1.0))))) + float64(0.1))) {
                            switch 2147483647 {
                            case -13:
                                v15 := GenS1{F0: 4294967295, F1: ((!true) && ((v2 < (v2 + (v2 - 1.5))) || false)), F2: (v8 + v8)}
                                _ = v15
                            case 32768:
                                for fi16 := 0; fi16 < 3; fi16++ {
                                    v4 = v3
                                    for fi17 := 0; fi17 < 5; fi17++ {
                                        fmt.Println(v8, v4, fi16)
                                        _ = fi17
                                    }
                                    _ = fi16
                                }
                            }
                            var v18 float64 = (v2 - v2)
                            _ = v18
                        } else {
                            var v19 float32 = (1.5 / 1.0)
                            _ = v19
                        }
                    }
                default:
                    var v20 int = 128
                    _ = v20
                }
                var v21 int32 = (^(v8))
                _ = v21
            }
            var v22 int8 = v3
            _ = v22
        } else {
            _ = len(v5)
            switch "foo bar" {
            case "hello":
                v23 := GenS1{F0: 1, F1: (!false), F2: (v8 >> ((v7) & 63))}
                _ = v23
            case "":
                v3 -= (v4 - v4)
            default:
                v4 = (v3 - v4)
            }
        }
    case "hello":
        v7 = v7
    case "":
        v8 = v8
    default:
        var v24 bool = (false || true)
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 int8 = (v3 % (((v4 + (v4 / (((v4 ^ v4)) | 1)))) | 1))
        _ = fi25
        _ = v26
    }
    var v27 int16 = -227
    var v28 int16 = v27
    v29 := GenS2{F0: 1, F1: v4, F2: (v8 / (((v8 * (v8 & v8))) | 1))}
    v7 = v7
    for fi30 := 0; fi30 < 11; fi30++ {
        v7 -= 1
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint32 = 441
            fmt.Println(fi31, v32, v2)
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    var v33 int16 = (v28 ^ 1000)
    var v34 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
