package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 42
    v3 := make(map[uint32]uint32)
    if false {
        var v4 float64 = float64(1e10)
        var v5 string = (("hello" + ("" + ("" + ""))) + "hello")
        _ = v4
        _ = v5
    }
    var v6 float64 = 0.0
    var v7 int8 = -1
    v8 := GenS1{F0: 32767}
    _ = len(v1)
    fmt.Println(v7)
    var v9 int8 = v7
    v2 += (v2 ^ v2)
    var v10 string = "\n\t"
    var v11 uint8 = 97
    switch -2147483648 {
    case -382:
        v9 = (v9 + (^(v7)))
    case -9223372036854775808:
        var v12 int32 = 689
        _ = v12
    case -1:
        var v13 uint8 = v11
        _ = v13
    default:
        if true {
            var v14 float64 = ((((((v6 * ((0.1 + 3.14159) * -2.25)) + v6) + ((v6 - v6) + ((float64(((0.1 / 0.1) / 1e10)) - 1.5) * v6))) / v6) / (v6 + 0.0)) * (1.5 + 1.5))
            fmt.Println(v10, v7)
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS2{F0: 65535, F1: -108, F2: 18446744073709551615}
                switch fi15 {
                case -2147483648:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        delete(v1, v10)
                        if (-2147483648 <= -1) {
                            v10 = ""
                            switch v10 {
                            case "":
                                if (0 > -601) {
                                    if (42 != -488) {
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(v6)
                                        }
                                        if ((^(v11)) < (v11 ^ 2)) {
                                            var v18 int64 = 1
                                            _ = v18
                                        } else {
                                            fmt.Println(v10)
                                            fmt.Println(v10, v9, v10)
                                        }
                                        v19 := GenS3{F0: 1, F1: -32768, F2: uint(v11)}
                                        _ = v19
                                    }
                                    var v20 int = (fi15 * fi17)
                                    v14 = float64(1.5)
                                    _ = v20
                                }
                            }
                            var v21 uint = 1
                            _ = v21
                        }
                        _ = fi17
                    }
                case 9223372036854775807:
                    for fi22 := 0; fi22 < 9; fi22++ {
                        var v23 int8 = (v7 / (((v9 << ((42) & 63))) | 1))
                        _ = fi22
                        _ = v23
                    }
                case 4294967295:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v25 := GenS1{F0: -618}
                        if (true && (!true)) {
                            v7 = -60
                            switch fi15 {
                            case -10:
                                v26 := []int{fi24}
                                _ = v26
                            default:
                                var v27 int32 = -2147483648
                                v28 := []int{}
                                _ = v27
                                _ = v28
                            }
                            if (!((v2 - v2) == v2)) {
                                v2 = (^(v2))
                                if true {
                                    v29 := GenS4{F0: v11, F1: fi24, F2: v10}
                                    _ = v29
                                } else {
                                    v30 := []int{fi15}
                                    v7 += 71
                                    _ = v30
                                }
                            } else {
                                switch (-(fi15)) {
                                case -713:
                                    v31 := []int{(-(fi15)), -2147483648}
                                    _ = v31
                                }
                                if (-1 < 9223372036854775807) {
                                    v32 := GenS1{F0: 42}
                                    v6 = (0.0 + ((1e10 - (((3.14159 - float64((float32(v14) + ((0.1 * float32(((1e10 * (float64(((0.1 / 1e10) - 1.5)) + v14)) - (v14 * (v14 + 3.14159))))) / 1.0)))) * v14) - ((float64(1.5) * v6) / ((0.0 / v14) * (v6 * (1e10 - 0.1)))))) * ((v14 + v6) + 0.1)))
                                    var v33 uint32 = 401
                                    v34 := []int{fi24, 42, fi15, fi24}
                                    v9 = (v7 - -125)
                                    v35 := GenS5{F0: 255, F1: int8(v2)}
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    v10 = v10
                                }
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            if (!(-2.25 > float32(1.5))) {
                switch v10 {
                case "hello":
                    for fi36 := 0; fi36 < 13; fi36++ {
                        if (0.0 < 0.1) {
                            v9 -= (v9 ^ (v9 & v9))
                            v10 = v10
                            var v37 float32 = 1.5
                            _ = v37
                        }
                        _ = fi36
                    }
                }
                v10 = v10
            } else {
                fmt.Println(v6, v7, v9)
                var v38 int32 = 2147483647
                _ = v38
            }
            var v39 int64 = 0
            fmt.Println(v10, v9)
            _ = v39
        }
    }
    v40 := GenS2{F0: 42, F1: -1, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v40
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
