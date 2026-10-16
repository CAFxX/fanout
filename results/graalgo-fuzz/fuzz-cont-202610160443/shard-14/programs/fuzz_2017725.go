package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0() uint {
    var v6 bool = false
    var v7 int16 = 0
    v8 := []int{2147483647, 127, -9223372036854775808, -89}
    var v9 uint32 = 2147483648
    v6 = v6
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 225
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 float32 = (0.1 - (0.0 * ((((((1e10 / (1e10 * 3.14159)) - 1.5) * (float32((1.5 - 1e10)) / (0.1 / 1.0))) - 3.14159) + 0.1) + 1e10)))
    var v5 int32 = 687
    v3[""] = "hello"
    v4 += (v4 / ((v4 * v4) + v4))
    var v10 uint64 = 18446744073709551615
    v11 := GenS1{F0: 0}
    v12 := []int{}
    switch 0 {
    case 127:
        fmt.Println(v10)
    default:
        v13 := GenS2{F0: 1, F1: 936}
        _ = v13
    }
    switch 926 {
    case 865:
        switch 32767 {
        case 2147483647:
            if len(v12) > 0 {
                v12[0] = int(v5)
            }
        case -856:
            switch "a" {
            case "a":
                if ((v4 + 1e10) < 3.14159) {
                    v4 += -2.25
                } else {
                    var v14 float64 = 0.0
                    v10 = (v10 % (((v10 / (((v10 % ((v10) | 1))) | 1))) | 1))
                    fmt.Println(v10, v5)
                    _ = v14
                }
            case "hello":
                switch "foo bar" {
                case "":
                    if (!false) {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            fmt.Println(fi15)
                            var v16 uint64 = (v10 / (((^(v10))) | 1))
                            v10 = (^(v10))
                            _ = fi15
                            _ = v16
                        }
                    }
                case "a":
                    var v17 int64 = 1
                    _ = v17
                case "\n\t":
                    var v18 int8 = -81
                    _ = v18
                default:
                    v19 := GenS2{F0: 1, F1: 65535}
                    _ = v19
                }
            case "":
                v10 = v10
            }
        }
    case 42:
        if false {
            var v20 int64 = 240
            var v21 float32 = v4
            v21 -= float32(1.5)
            _ = v20
            _ = v21
        }
    case -560:
        v5 -= (v5 | 257)
    default:
        v22 := GenS2{F0: -1, F1: 1}
        _ = v22
    }
    var v23 uint64 = 0
    switch 0 {
    case -2147483648:
        v4 += v4
    case 127:
        if (!(true || (-32768 > 1))) {
            v10 -= 1
            if (!false) {
                v23 = 0
                if true {
                    delete(v2, 2147483648)
                    if false {
                        if len(v12) > 0 {
                            v12[0] = 128
                        }
                        var v24 int64 = 42
                        v4 = v4
                        _ = v24
                    } else {
                        v4 = 0.0
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v5 -= -1
                            var v26 int32 = (v5 / (((^(v5))) | 1))
                            switch (fi25 ^ fi25) {
                            case 1:
                                delete(v2, 0)
                            case 32768:
                                if false {
                                    fmt.Println(v23)
                                    var v27 uint = 901
                                    var v28 int = 4294967295
                                    _ = v27
                                    _ = v28
                                } else {
                                    v5 = -2147483648
                                    v3[("hello" + "\n\t")] = "\n\t"
                                }
                            }
                            _ = fi25
                            _ = v26
                        }
                    }
                    v29 := GenS3{F0: (v23 & v10)}
                    _ = v29
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
