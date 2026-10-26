package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 float64 = 1e10
    var v2 int16 = 1
    var v3 uint = 42
    fmt.Println(v3, v1)
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(v1)
        _ = fi4
    }
    switch -9223372036854775808 {
    case 2147483647:
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := []int{}
            if false {
                if true {
                    var v7 float64 = (v1 - (v1 / v1))
                    for fi8 := 0; fi8 < 10; fi8++ {
                        switch fi8 {
                        case -423:
                            var v9 float32 = float32((v1 + ((0.0 / (v7 - v7)) - (float64(float32(-2.25)) / 1.0))))
                            _ = v9
                        case -9223372036854775808:
                            var v10 uint32 = 42
                            _ = v10
                        default:
                            switch "" {
                            case "":
                                if false {
                                    switch (fi5 | 32768) {
                                    case 0:
                                        v3 = 956
                                    case 32768:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    default:
                                        if true {
                                            v1 += 3.14159
                                            var v11 uint32 = 4294967295
                                            var v12 uint32 = v11
                                            var v13 uint8 = 128
                                            _ = len(v6)
                                            _ = v11
                                            _ = v12
                                            _ = v13
                                        }
                                    }
                                    _ = len(v6)
                                    v3 -= v3
                                    v3 = 18446744073709551615
                                } else {
                                    v3 -= v3
                                    var v14 int16 = v2
                                    _ = v14
                                }
                            case "hello":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            default:
                                v7 += (((float64((1.5 - -2.25)) / v1) + v7) + (v1 / -2.25))
                            }
                        }
                        if len(v6) > 0 {
                            v6[0] = fi8
                        }
                        if ((!(56 < -128)) || true) {
                            v7 -= (v1 / 1.0)
                            v15 := GenS1{F0: 255, F1: 32767, F2: 255}
                            _ = v15
                        }
                        _ = fi8
                    }
                    var v16 int32 = 0
                    _ = v16
                    _ = v7
                } else {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v2 = (v2 & (v2 >> ((v3) & 63)))
                        v2 = v2
                        v18 := GenS2{F0: 128, F1: 2147483647}
                        var v19 uint = v3
                        v19 = 1
                        var v20 uint16 = 255
                        fmt.Println(v3, fi17)
                        v1 += v1
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(v3, v3, fi5)
                    v21 := GenS3{F0: "hello"}
                    _ = v21
                }
                var v22 string = "\n\t"
                fmt.Println(v1, v22, v3)
                _ = v22
            }
            fmt.Println(v3, v3)
            _ = fi5
            _ = v6
        }
    case 42:
        v23 := GenS3{F0: "\n\t"}
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        if true {
            v25 := []int{fi24, fi24}
            _ = v25
        }
        _ = fi24
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v2 += (v2 - v2)
        _ = fi26
    }
    v3 -= v3
    if true {
        var v27 int8 = -43
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 uint16 = 65535
            v1 -= (v1 - v1)
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        for fi30 := 0; fi30 < 14; fi30++ {
            var v31 int64 = -332
            var v32 float32 = 0.1
            for fi33 := 0; fi33 < 19; fi33++ {
                v34 := []int{-1, fi30}
                _ = fi33
                _ = v34
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := GenS1{F0: uint8(v3), F1: 0, F2: 2}
            _ = fi35
            _ = v36
        }
    }
    v2 = (v2 & -1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
