package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint; F1 int16; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 int16 = 32767
    v4 := []int{32767}
    var v5 bool = true
    for fi6 := 0; fi6 < 4; fi6++ {
        v3 += -1
        v7 := GenS1{F0: 752, F1: (!(!v5))}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    if true {
        v5 = v5
    }
    var v8 float32 = (((0.0 - ((0.1 + 1e10) - (1.5 + 1e10))) * (1e10 / float32(0.1))) + 0.1)
    v4 = append(v4, -9223372036854775808)
    switch "\n\t" {
    case "a":
        for fi9 := 0; fi9 < 14; fi9++ {
            v3 = (^(v3))
            var v10 int64 = -1
            _ = fi9
            _ = v10
        }
    default:
        v11 := GenS2{F0: -852, F1: 32767, F2: 880}
        _ = v11
    }
    var v12 int = 127
    v1["\n\t"] = 0
    v8 += v8
    if len(v4) > 0 {
        v4[0] = v12
    }
    var v13 uint64 = 490
    fmt.Println(v12)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int8 = 127
        if true {
            v15 = v15
            v16 := GenS3{F0: 0, F1: 127}
            for fi17 := 0; fi17 < 20; fi17++ {
                v12 += fi14
                if (378 < 65535) {
                    v5 = v5
                    var v18 int64 = -9223372036854775808
                    v19 := GenS2{F0: -608, F1: 32767, F2: 255}
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v5, fi14, v8)
                    if (!false) {
                        if v5 {
                            var v20 uint = uint(v13)
                            if (178 >= -9223372036854775808) {
                                v4 = append(v4, 127)
                                switch "hello" {
                                case "":
                                    var v21 string = "\n\t"
                                    _ = v21
                                }
                            }
                            var v22 int = v12
                            _ = v20
                            _ = v22
                        }
                        var v23 uint64 = 586
                        v24 := GenS1{F0: 0, F1: false}
                        _ = len(v4)
                        _ = v23
                        _ = v24
                    }
                }
                var v25 string = "hello"
                _ = fi17
                _ = v25
            }
            _ = v16
        } else {
            switch fi14 {
            case 42:
                fmt.Println(v15)
            default:
                for fi26 := 0; fi26 < 2; fi26++ {
                    switch 1 {
                    case 734:
                        v27 := GenS4{F0: 65535, F1: (v3 % ((v3) | 1)), F2: 42}
                        _ = v27
                    case 522:
                        if (9223372036854775807 != (-(fi14))) {
                            v15 = (^(v15))
                            v1["\n\t"] = 2147483648
                            delete(v1, "a")
                        } else {
                            var v28 uint = 65535
                            _ = v28
                        }
                    default:
                        v8 += v8
                    }
                    v29 := GenS4{F0: 304, F1: v3, F2: -1}
                    v5 = (v5 && true)
                    _ = fi26
                    _ = v29
                }
            }
        }
        _ = fi14
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
