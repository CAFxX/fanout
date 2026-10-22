package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 string }

func main() {
    var v1 uint32 = 421
    var v2 int8 = -128
    v3 := []int{-794, 127}
    v4 := make(map[int]string)
    v5 := []string{"a", "\n\t", ("\n\t" + "")}
    delete(v4, 4294967295)
    if true {
        v1 -= (v1 * v1)
        if false {
            v6 := GenS1{F0: 128, F1: 127}
            v2 = v2
            _ = v6
        } else {
            if true {
                v1 = (v1 / ((v1) | 1))
                fmt.Println(v1)
                var v7 float64 = 1.5
                _ = v7
            }
            v3 = append(v3, -1)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v8 := GenS2{F0: 4294967295}
        _ = v8
    }
    if (128 < 127) {
        switch "hello" {
        case "":
            if (-2147483648 == 1) {
                if (0.0 <= 3.14159) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v1 = (v1 & 0)
                        v1 -= (v1 ^ (^(v1)))
                        _ = len(v3)
                        _ = fi9
                    }
                    fmt.Println(v2)
                    switch "a" {
                    case "a":
                        v1 = v1
                    case "\n\t":
                        if (1 <= 42) {
                            if true {
                                fmt.Println(v1)
                                v2 -= (v2 - -89)
                            }
                            v10 := GenS2{F0: 829}
                            _ = v10
                        } else {
                            var v11 uint8 = 128
                            switch 42 {
                            case 128:
                                for fi12 := 0; fi12 < 6; fi12++ {
                                    delete(v4, fi12)
                                    _ = fi12
                                }
                            case -1:
                                v13 := v4[-9223372036854775808]
                                _ = v13
                            case -832:
                                switch -768 {
                                case 682:
                                    var v14 int8 = v2
                                    _ = v14
                                case -785:
                                    v1 = 2147483647
                                default:
                                    var v15 uint8 = v11
                                    _ = v15
                                }
                            default:
                                for fi16 := 0; fi16 < 2; fi16++ {
                                    v11 -= (v11 ^ (v11 / (((v11 * 19)) | 1)))
                                    v2 += (v2 * -128)
                                    _ = fi16
                                }
                            }
                            v17 := GenS2{F0: 2}
                            _ = v11
                            _ = v17
                        }
                    case "hello":
                        v1 += (v1 * v1)
                    default:
                        switch 32767 {
                        case 2147483647:
                            var v18 int32 = -1
                            _ = v18
                        default:
                            if true {
                                var v19 uint8 = 1
                                fmt.Println(v19)
                                _ = v19
                            } else {
                                fmt.Println(v1)
                                fmt.Println(v2, v1)
                                if (467 <= 577) {
                                    switch 128 {
                                    case -360:
                                        v20 := GenS1{F0: 2, F1: 255}
                                        _ = v20
                                    }
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v1, v2)
                    if (701 != 841) {
                        v1 += (v1 << ((18446744073709551615) & 63))
                        v21 := GenS3{F0: v2}
                        switch "\n\t" {
                        case "hello":
                            switch -522 {
                            case 300:
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    if false {
                                        var v23 int64 = 42
                                        var v24 float64 = 0.1
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v25 := GenS4{F0: 135}
                                        _ = v25
                                    }
                                    v2 = (v2 >> ((0) & 63))
                                    _ = fi22
                                }
                            case 0:
                                _ = len(v4)
                            case 128:
                                v26 := GenS1{F0: 128, F1: 36}
                                _ = v26
                            default:
                                fmt.Println(v2, v1)
                            }
                        default:
                            fmt.Println(v2, v1)
                        }
                        _ = v21
                    }
                }
            }
        default:
            switch 4294967295 {
            case 127:
                v4[0] = ("a" + ("hello" + (("a" + ((("foo bar" + "a") + ("" + "foo bar")) + "foo bar")) + "a")))
            default:
                for fi27 := 0; fi27 < 19; fi27++ {
                    fmt.Println(v2)
                    _ = len(v3)
                    _ = fi27
                }
            }
        }
        if true {
            v4[4294967295] = "foo bar"
        } else {
            for fi28 := 0; fi28 < 20; fi28++ {
                if ((fi28 >> ((686) & 63)) == (^(fi28))) {
                    v29 := GenS2{F0: 0}
                    _ = v29
                }
                _ = fi28
            }
            for fi30 := 0; fi30 < 13; fi30++ {
                if len(v3) > 0 {
                    v3[0] = (fi30 + fi30)
                }
                _ = fi30
            }
        }
        var v31 uint8 = 128
        _ = v31
    }
    var v32 uint = 18446744073709551615
    v32 = (v32 << (((v32 >> (((v32 >> (((v32 ^ (v32 % (((v32 >> (((v32 | (^(v32)))) & 63))) | 1)))) & 63))) & 63))) & 63))
    if true {
        var v33 uint8 = 42
        _ = v33
    }
    v1 += v1
    v34 := GenS5{F0: 0, F1: "hello"}
    var v35 int16 = -779
    fmt.Println(v35, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
