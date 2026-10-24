package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
func genFn0(p0 bool, p1 uint, p2 uint16) int32 {
    fmt.Println(p2, p0, p1)
    p0 = false
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int64, p1 uint8) uint32 {
    var v4 int16 = 32767
    p1 -= 127
    for fi5 := 0; fi5 < 9; fi5++ {
        switch fi5 {
        case 1:
            v4 += v4
        case 9223372036854775807:
            p0 = p0
        }
        _ = fi5
    }
    v6 := []int{-885}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 0
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = -607
        _ = fi8
        _ = v9
    }
    var v10 string = ""
    v1 = 310
    v11 := []int{42, 2147483647, 0}
    fmt.Println(v10)
    v12 := GenS1{F0: (v1 & (v1 ^ (v1 ^ 0))), F1: -83}
    var v13 int8 = -78
    for fi14 := 0; fi14 < 10; fi14++ {
        v7 = (v7 * ((1e10 - float64((1.5 * 0.1))) - 1e10))
        _ = fi14
    }
    if ((65535 < 65535) && true) {
        _ = len(v11)
        v7 = (v7 + v7)
        if ((42 < 1) && (18446744073709551615 < 1)) {
            for fi15 := 0; fi15 < 17; fi15++ {
                if false {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v13, fi15)
                    default:
                        fmt.Println(v7, v13, fi15)
                    }
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    v13 = v13
                    _ = fi16
                }
                v11 = append(v11, (fi15 | (-(fi15))))
                _ = fi15
            }
            v10 = (v10 + "\n\t")
            var v17 int8 = v13
            _ = v17
        } else {
            if false {
                v11 = append(v11, -656)
                if true {
                    v18 := GenS1{F0: v1, F1: -1}
                    if false {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            for fi20 := 0; fi20 < 10; fi20++ {
                                fmt.Println(v13, v10)
                                var v21 uint = 512
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(v13, v13)
                            _ = fi19
                        }
                    } else {
                        if true {
                            v22 := GenS2{F0: (v13 | -104), F1: uint64(v1)}
                            v1 -= v1
                            if (v13 < 127) {
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    var v24 int16 = 0
                                    if (!true) {
                                        v25 := GenS1{F0: 471, F1: v24}
                                        for fi26 := 0; fi26 < 9; fi26++ {
                                            v27 := GenS3{F0: v10, F1: (v10 + "hello"), F2: 282}
                                            _ = fi26
                                            _ = v27
                                        }
                                        _ = len(v3)
                                        _ = v25
                                    } else {
                                        var v28 int8 = v13
                                        v1 = 255
                                        var v29 int32 = -1
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                v30 := GenS2{F0: (-(v13)), F1: 0}
                                _ = v30
                            }
                            _ = v22
                        }
                        v10 = v10
                    }
                    var v31 int = 4294967295
                    _ = v18
                    _ = v31
                }
                if true {
                    fmt.Println(v1, v10)
                } else {
                    v10 = v10
                    if (v7 < v7) {
                        var v32 int = 1
                        _ = v32
                    }
                }
            } else {
                if true {
                    var v33 int = 2147483647
                    var v34 float64 = (v7 + 1e10)
                    _ = v33
                    _ = v34
                }
            }
            var v35 bool = true
            v1 += v1
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
