package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 string }

func main() {
    var v1 int32 = 1
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 uint32 = 4294967295
    var v5 int16 = int16(v1)
    for fi6 := 0; fi6 < 7; fi6++ {
        v5 = (v5 & (v5 + (v5 * (v5 | -569))))
        v3[(v4 % ((v4) | 1))] = fi6
        _ = fi6
    }
    v7 := v3[v4]
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v5)
                fmt.Println(v4, v7, v1)
                if true {
                    var v10 string = ("hello" + "a")
                    _ = len(v3)
                    var v11 float32 = 0.0
                    if ("\n\t" == v10) {
                        var v12 bool = (!true)
                        v13 := GenS1{F0: v5}
                        v14 := []int{fi8}
                        v1 -= (-(v1))
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v10, v2, v2)
                        case "":
                            if v12 {
                                var v15 int64 = v2
                                v5 -= v5
                                _ = v15
                            } else {
                                if ((v5 / (((v5 - v5)) | 1)) == -1) {
                                    if (4294967295 >= 4294967295) {
                                        if false {
                                            if v12 {
                                                var v16 uint16 = 2
                                                if v12 {
                                                    if (1 > v4) {
                                                        if v12 {
                                                            var v17 string = v10
                                                            _ = v17
                                                        }
                                                    } else {
                                                        v4 -= v4
                                                        v1 += (^(v1))
                                                        v4 = (v4 + 876)
                                                    }
                                                }
                                                if v12 {
                                                    fmt.Println(v4)
                                                } else {
                                                    v18 := GenS2{F0: -32768, F1: 65535, F2: ""}
                                                    var v19 int = int(v1)
                                                    v1 -= (v1 + v1)
                                                    _ = v18
                                                    _ = v19
                                                }
                                                v16 = (v16 % ((v16) | 1))
                                                if v12 {
                                                    fmt.Println(v16, v11)
                                                }
                                                _ = v16
                                            } else {
                                                var v20 int = fi9
                                                _ = v20
                                            }
                                            fmt.Println(fi8, v5, fi8)
                                            switch "" {
                                            case "":
                                                v1 = v1
                                            default:
                                                var v21 uint = 255
                                                _ = v21
                                            }
                                        } else {
                                            fmt.Println(v7, v7)
                                        }
                                        var v22 string = "\n\t"
                                        _ = v22
                                    }
                                }
                                switch v7 {
                                case 127:
                                    fmt.Println(fi8, v1, v12)
                                case 32767:
                                    fmt.Println(v12, fi8)
                                }
                                v1 = 42
                            }
                        default:
                            if (255 != 233) {
                                v5 += 1
                            } else {
                                fmt.Println(v4, v4, v11)
                                var v23 uint8 = 127
                                _ = v23
                            }
                        }
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    v11 -= v11
                    _ = v10
                    _ = v11
                }
                v5 += int16(v4)
                fmt.Println(fi8, v5)
                _ = fi9
            }
        }
        v24 := GenS2{F0: (^(v5)), F1: 515, F2: ""}
        var v25 int32 = -56
        _ = fi8
        _ = v24
        _ = v25
    }
    v7 = (v7 / (((v7 * (v7 ^ int(v4)))) | 1))
    v4 = 65535
    v26 := []int{42, v7}
    var v27 string = "hello"
    v28 := GenS1{F0: 1000}
    fmt.Println(v4, v4, v27)
    for fi29 := 0; fi29 < 13; fi29++ {
        switch v7 {
        case 4294967295:
            switch v27 {
            case "":
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(fi30)
                    v4 += v4
                    _ = fi30
                }
            }
        default:
            var v31 bool = (-1 != int8(v1))
            _ = v31
        }
        var v32 uint8 = 0
        _ = fi29
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
