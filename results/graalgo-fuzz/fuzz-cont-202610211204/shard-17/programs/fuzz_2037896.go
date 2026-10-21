package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    v4 := []int{-2147483648}
    v5 := []uint32{904, 65535, 651, 4294967295}
    var v6 int = 9223372036854775807
    var v7 int32 = 463
    fmt.Println(v6, v2, v7)
    if false {
        delete(v3, "a")
    }
    v2 += (^(v2))
    fmt.Println(v2, v7, v6)
    for fi8 := 0; fi8 < 10; fi8++ {
        switch fi8 {
        case 32767:
            for fi9 := 0; fi9 < 14; fi9++ {
                v2 -= (v2 % ((32767) | 1))
                v10 := v3[("a" + "")]
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v7 = 2147483647
                var v11 float64 = (((3.14159 * 1.5) - float64((((0.1 / 1e10) - (1.5 / 1.0)) * 0.1))) + (float64(0.0) / float64(-2.25)))
                v6 += (fi9 % ((fi9) | 1))
                v2 = (v2 << ((0) & 63))
                _ = fi9
                _ = v10
                _ = v11
            }
        case 1:
            v2 += (^(v2))
        case 0:
            var v12 uint32 = 1
            _ = v12
        }
        fmt.Println(v2)
        _ = fi8
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 float64 = 1.5
        var v15 int = 686
        _ = len(v5)
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        switch ((("hello" + "a") + "") + "") {
        case "foo bar":
            if (0.1 == -2.25) {
                switch "" {
                case "a":
                    v2 = 42
                case "\n\t":
                    fmt.Println(v6, v6, v2)
                case "":
                    var v17 float32 = ((1.5 - 0.0) * 0.0)
                    _ = v17
                default:
                    if (true || ((!((2 > 255) || false)) || false)) {
                        v6 = (v6 % (((v6 & -2147483648)) | 1))
                        v7 -= v7
                        var v18 int64 = 42
                        _ = v18
                    }
                }
                switch 94 {
                case 9223372036854775807:
                    var v19 uint8 = 2
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v21 := GenS1{F0: 255, F1: -873}
                        v2 -= v2
                        _ = fi20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS2{F0: ("\n\t" + "a"), F1: 160, F2: -123}
                    var v24 bool = true
                    if ("a" >= "foo bar") {
                        switch fi16 {
                        case 4294967295:
                            v6 = int(v2)
                        case -451:
                            var v25 uint = 18446744073709551615
                            _ = v25
                        case 127:
                            if (!v24) {
                                fmt.Println(v24)
                                v6 += (fi22 ^ (fi22 - fi16))
                                var v26 uint32 = 2147483648
                                _ = v26
                            } else {
                                fmt.Println(v6, fi16, v6)
                                switch "\n\t" {
                                case "":
                                    var v27 float32 = float32(1e10)
                                    _ = v27
                                default:
                                    v24 = (65535 == 0)
                                }
                                if ((v6 << ((255) & 63)) != fi22) {
                                    fmt.Println(fi22, v7, v2)
                                    v1[("a" + "foo bar")] = (122 != 103)
                                    v7 -= v7
                                }
                            }
                        default:
                            switch 127 {
                            case 1:
                                switch fi16 {
                                case -1:
                                    delete(v1, "hello")
                                case -147:
                                    v7 = (-(v7))
                                }
                            case 2147483647:
                                delete(v1, "hello")
                            case 594:
                                var v28 float32 = -2.25
                                _ = v28
                            default:
                                var v29 string = "hello"
                                _ = v29
                            }
                        }
                        v2 -= v2
                    } else {
                        v2 = v2
                        var v30 bool = ((v2 == 255) && v24)
                        v24 = false
                        _ = v30
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        default:
            for fi31 := 0; fi31 < 14; fi31++ {
                if (!(true || false)) {
                    _ = len(v1)
                }
                v32 := GenS3{F0: 2}
                _ = fi31
                _ = v32
            }
        }
        _ = fi16
    }
    if true {
        var v33 int32 = v7
        _ = v33
    }
    if false {
        fmt.Println(v6)
    } else {
        _ = len(v4)
        v6 -= v6
        v34 := GenS4{F0: -1, F1: v7, F2: -2147483648}
        _ = v34
    }
    v6 -= 2147483647
    v2 -= 1
    var v35 uint16 = v2
    v35 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
