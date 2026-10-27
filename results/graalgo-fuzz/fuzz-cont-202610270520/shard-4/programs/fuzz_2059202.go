package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 int64, p1 int16) string {
    v4 := GenS1{F0: 117, F1: 2147483647, F2: 728}
    var v5 int = 4294967295
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "a"
}

func main() {
    v1 := []int{785}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v6 float32 = -2.25
    v6 -= 1.5
    v6 += v6
    var v7 bool = false
    if (32767 >= -32768) {
        var v8 int = 482
        _ = len(v3)
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                delete(v3, (fi9 + v8))
                if (42 == 32767) {
                    if false {
                        for fi10 := 0; fi10 < 10; fi10++ {
                            switch "\n\t" {
                            case "foo bar":
                                var v11 uint64 = 4294967295
                                _ = v11
                            case "a":
                                v7 = false
                            default:
                                fmt.Println(fi9, fi9, fi10)
                                var v12 uint32 = 1
                                v12 += v12
                                switch "foo bar" {
                                case "a":
                                    v7 = (!(!v7))
                                case "hello":
                                    var v13 int64 = 334
                                    _ = v13
                                default:
                                    v14 := GenS1{F0: -1, F1: 32767, F2: 45}
                                    _ = v14
                                }
                                _ = v12
                            }
                            var v15 uint8 = 128
                            v16 := GenS2{F0: 128, F1: 9223372036854775807}
                            v15 = v15
                            _ = fi10
                            _ = v15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 17; fi17++ {
                            fmt.Println(v8, v6, fi17)
                            if (32767 >= 1000) {
                                v8 -= (fi9 | -2147483648)
                                v6 = v6
                            }
                            _ = fi17
                        }
                        v7 = (!(!(42 < 725)))
                    }
                } else {
                    var v18 bool = v7
                    var v19 uint32 = 42
                    var v20 float64 = (((((0.0 + 0.0) / 0.1) - float64(v6)) - 0.0) - (((0.0 * 0.0) / ((3.14159 / (1.5 * 0.1)) / ((-2.25 / 0.1) + float64((v6 / (v6 / v6)))))) + (0.0 - float64((1.5 * (1e10 - v6))))))
                    _ = v18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v8)
                _ = fi9
            }
        } else {
            var v21 int16 = -78
            v8 = 775
            _ = v21
        }
        _ = v8
    } else {
        fmt.Println(v6, v6)
        v6 += (float32(0.1) * (((v6 + v6) / (v6 + v6)) * (((float32((3.14159 / 1.5)) + v6) - v6) * ((v6 + (v6 - float32(1.5))) * v6))))
        fmt.Println(v7, v6)
    }
    fmt.Println(v6, v6)
    var v22 uint = 694
    var v23 int32 = -2147483648
    v24 := GenS2{F0: 128, F1: 0}
    if v7 {
        var v25 int16 = 32767
        switch 1 {
        case -298:
            switch 9223372036854775807 {
            case -9223372036854775808:
                v26 := GenS3{F0: 65535, F1: 212}
                _ = v26
            }
        case 72:
            switch 32767 {
            case -64:
                v25 += -552
            case 128:
                v7 = v7
            case 127:
                v27 := v2[65535]
                _ = v27
            }
        }
        _ = v25
    } else {
        v23 += v23
    }
    v28 := GenS3{F0: 47, F1: uint8(v22)}
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 uint16 = 65535
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS4{F0: ("foo bar" + ("foo bar" + ("a" + "hello"))), F1: 2147483647}
            v23 += (v23 | -640)
            fmt.Println(v6)
            _ = fi31
            _ = v32
        }
        _ = len(v3)
        _ = fi29
        _ = v30
    }
    fmt.Println(v23, v23, v23)
    var v33 int = 32768
    var v34 int32 = (v23 / ((v23) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v6
    _ = v7
}
