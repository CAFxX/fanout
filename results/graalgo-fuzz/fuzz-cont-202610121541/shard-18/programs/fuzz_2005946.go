package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int16 = -83
    var v2 uint32 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := []int{-251, -2147483648}
    v6 := GenS1{F0: "\n\t", F1: (v2 ^ v2)}
    v7 := GenS1{F0: ("\n\t" + "\n\t"), F1: (v2 ^ v2)}
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: "", F1: (v2 - (v2 - v2))}
        var v10 float64 = 1.5
        var v11 uint = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float32 = 0.1
    v12 = v12
    if false {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch fi13 {
            case -2147483648:
                if (!false) {
                    _ = len(v3)
                    v12 = (float32(float64(float32(float64(v12)))) * ((((v12 * (-2.25 - float32(-2.25))) + 3.14159) / (float32((((3.14159 + (-2.25 / 1.0)) - 1e10) - 0.0)) + 1e10)) + float32(3.14159)))
                    var v14 int64 = -9223372036854775808
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        switch fi15 {
                        case 966:
                            var v16 int8 = 73
                            _ = v16
                        }
                        if (("hello" == (("foo bar" + ("foo bar" + ("" + "hello"))) + "foo bar")) && false) {
                            switch ("" + ("hello" + "foo bar")) {
                            case "hello":
                                v17 := GenS2{F0: 42, F1: (fi13 >= -9223372036854775808)}
                                _ = v17
                            case "a":
                                var v18 uint64 = 563
                                _ = v18
                            default:
                                var v19 float64 = float64(v12)
                                _ = v19
                            }
                        }
                        if (!false) {
                            if (-2.25 != (float64((v12 + ((0.1 / v12) / 1.5))) + float64(v12))) {
                                if false {
                                    if true {
                                        v12 -= 0.0
                                        fmt.Println(v2, v2)
                                    } else {
                                        _ = len(v4)
                                    }
                                    if true {
                                        v20 := v3[(fi15 % ((fi13) | 1))]
                                        v2 = (v2 & (v2 | v20))
                                        if true {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v21 uint64 = 4294967295
                                            var v22 int64 = -433
                                            fmt.Println(v2, v12, fi15)
                                            _ = v21
                                            _ = v22
                                        } else {
                                            fmt.Println(v12, fi13, v12)
                                        }
                                        _ = v20
                                    }
                                } else {
                                    v5 = append(v5, (fi15 / (((-(fi13))) | 1)))
                                    v2 = (v2 >> ((4294967295) & 63))
                                    v12 = (((float32(1.5) + v12) * v12) / v12)
                                }
                            }
                            var v23 int64 = 9223372036854775807
                            v24 := GenS2{F0: 601, F1: (!true)}
                            _ = v23
                            _ = v24
                        } else {
                            if true {
                                fmt.Println(v12)
                                v5 = append(v5, (fi15 >> ((1) & 63)))
                                v2 -= v2
                            }
                        }
                        _ = fi15
                    }
                    v25 := GenS2{F0: 0, F1: true}
                    var v26 float32 = 0.0
                    _ = v25
                    _ = v26
                }
            case 42:
                for fi27 := 0; fi27 < 6; fi27++ {
                    v1 += (v1 - (v1 & v1))
                    switch ("\n\t" + "hello") {
                    case "hello":
                        if false {
                            if (1.5 == 1.5) {
                                v28 := GenS3{F0: (v2 | (v2 * v2))}
                                switch "\n\t" {
                                case "a":
                                    v2 = v2
                                }
                                fmt.Println(fi27)
                                _ = v28
                            }
                            var v29 int64 = 120
                            _ = v29
                        } else {
                            v12 = 0.0
                            switch ((("foo bar" + "hello") + "foo bar") + "") {
                            case "\n\t":
                                v12 = float32((3.14159 + 3.14159))
                            case "foo bar":
                                delete(v3, -1)
                                v12 -= (v12 + 0.0)
                                v12 = 0.1
                            case "a":
                                var v30 uint = 255
                                _ = v30
                            default:
                                var v31 string = "foo bar"
                                _ = v31
                            }
                        }
                    }
                    _ = fi27
                }
            case 4294967295:
                var v32 float32 = float32(-2.25)
                _ = v32
            default:
                v1 -= v1
            }
            v2 -= (^(v2))
            _ = fi13
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
