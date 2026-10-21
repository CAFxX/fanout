package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 string }
func genFn0() string {
    var v4 uint64 = 18446744073709551615
    if (false || false) {
        v5 := []int{182}
        _ = v5
    }
    fmt.Println(v4)
    var v6 uint16 = 863
    _ = v4
    _ = v6
    return (("hello" + "a") + "")
}
func genFn1(p0 int32) uint {
    fmt.Println(p0)
    var v7 int = -2147483648
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{(v7 & fi8), v7, (fi8 & (fi8 / ((v7) | 1))), -797}
        v7 = v7
        var v10 bool = (false && (true || false))
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = v7
    var v11 int64 = 1
    _ = p0
    _ = v11
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := []int{128, 1, 0}
    v3 = append(v3, -411)
    v12 := GenS1{F0: 0, F1: -763, F2: 73}
    var v13 int32 = 0
    _ = len(v3)
    v13 = v13
    switch 9223372036854775807 {
    case -2147483648:
        fmt.Println(v13)
    case 127:
        if true {
            if (0.0 >= 1e10) {
                v13 += (v13 & v13)
                for fi14 := 0; fi14 < 4; fi14++ {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = (fi14 ^ fi14)
                        }
                        for fi15 := 0; fi15 < 18; fi15++ {
                            v16 := GenS1{F0: 0, F1: 9223372036854775807, F2: 1}
                            v13 = v13
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        switch fi14 {
                        case -798:
                            switch (-(fi14)) {
                            case 497:
                                fmt.Println(v13, v13)
                            case 32767:
                                for fi17 := 0; fi17 < 10; fi17++ {
                                    if len(v3) > 0 {
                                        v3[0] = fi17
                                    }
                                    _ = fi17
                                }
                            default:
                                v18 := GenS1{F0: 2147483647, F1: 42, F2: 127}
                                _ = v18
                            }
                        default:
                            if true {
                                var v19 int32 = (v13 & (v13 << ((42) & 63)))
                                _ = v19
                            } else {
                                switch "a" {
                                case "":
                                    var v20 uint = 65535
                                    _ = v20
                                case "hello":
                                    var v21 uint16 = 1
                                    _ = v21
                                }
                            }
                        }
                        if (32767 == -1) {
                            var v22 int16 = 440
                            _ = v22
                        } else {
                            switch "" {
                            case "hello":
                                v13 += v13
                            case "foo bar":
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    fmt.Println(fi23, fi23, fi14)
                                    var v24 float64 = 3.14159
                                    var v25 bool = false
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                fmt.Println(fi14, fi14)
                            }
                            fmt.Println(v13, v13)
                        }
                        v13 = (v13 | (v13 ^ (v13 + v13)))
                    }
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS1{F0: 2147483647, F1: 0, F2: 3}
                        fmt.Println(fi26)
                        _ = fi26
                        _ = v27
                    }
                    _ = fi14
                }
            }
            v2[2147483647] = false
        }
    }
    v28 := v1[0]
    v13 = (v13 % ((v13) | 1))
    v29 := GenS2{F0: "\n\t"}
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
