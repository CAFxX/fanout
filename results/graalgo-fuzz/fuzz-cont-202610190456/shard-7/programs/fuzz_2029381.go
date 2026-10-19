package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []uint32{625, 2147483647, 953, 384}
    var v2 int32 = 2147483647
    var v3 int8 = 1
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[uint32]bool)
    var v7 float32 = 1.5
    if (true && true) {
        v7 -= v7
        if ((0 < 0) && false) {
            v8 := GenS1{F0: (-(v2)), F1: int32(v3), F2: 0}
            v2 = v2
            _ = v8
        } else {
            if (v7 <= v7) {
                v9 := v4[258]
                if (!false) {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        var v11 float64 = (((1.5 / 3.14159) / 1e10) * (((float64(float32((3.14159 - -2.25))) - (0.0 / (((1.5 - 1e10) / 3.14159) * float64(float32(1e10))))) - 3.14159) + -2.25))
                        if (true && false) {
                            v2 += v2
                        } else {
                            var v12 int8 = -111
                            v13 := GenS2{F0: 1, F1: 2147483647, F2: 730}
                            var v14 int = fi10
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                v3 = (v3 & (v3 ^ (v3 ^ v3)))
                _ = v9
            }
            var v15 float32 = v7
            v3 = -1
            _ = v15
        }
        if len(v1) > 0 {
            v1[0] = 2
        }
    } else {
        fmt.Println(v3, v2)
        v3 += (v3 & v3)
        var v16 int32 = (v2 * v2)
        _ = v16
    }
    v17 := GenS3{F0: 1000, F1: 0}
    v7 += (((v7 + 0.0) * v7) - 0.1)
    v5 = 2
    var v18 uint16 = 65535
    if (!true) {
        var v19 float32 = float32((((((1e10 + float64(0.0)) / 0.1) - float64(3.14159)) * 0.1) - 3.14159))
        if false {
            v4[2147483647] = 524
            v19 = v19
            fmt.Println(v3)
        } else {
            _ = len(v6)
            if (!(32767 == (^(v18)))) {
                v19 -= v7
                v2 = (-(v2))
                if (0 == v3) {
                    v20 := GenS3{F0: 76, F1: v18}
                    var v21 int = -1
                    switch -2147483648 {
                    case 1:
                        var v22 bool = false
                        _ = v22
                    case 127:
                        v23 := GenS2{F0: 1, F1: 1, F2: (v18 >> ((42) & 63))}
                        _ = v23
                    default:
                        v24 := GenS3{F0: 1, F1: 65535}
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
            } else {
                var v25 int32 = (v2 >> ((742) & 63))
                for fi26 := 0; fi26 < 10; fi26++ {
                    v1 = append(v1, 0)
                    if true {
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 string = "foo bar"
                            v19 += v19
                            _ = fi27
                            _ = v28
                        }
                        var v29 string = ("\n\t" + "hello")
                        _ = v29
                    }
                    var v30 int = (fi26 + 4294967295)
                    _ = fi26
                    _ = v30
                }
                if false {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        switch 458 {
                        case 127:
                            for fi32 := 0; fi32 < 11; fi32++ {
                                v18 += (^(v18))
                                var v33 float64 = 0.1
                                _ = fi32
                                _ = v33
                            }
                        }
                        var v34 uint64 = v5
                        var v35 int32 = (^(v2))
                        _ = fi31
                        _ = v34
                        _ = v35
                    }
                    delete(v4, 4294967295)
                }
                _ = v25
            }
        }
        var v36 int = 128
        _ = v19
        _ = v36
    } else {
        fmt.Println(v2, v2)
    }
    v2 -= -992
    if false {
        v1 = append(v1, 13)
    }
    fmt.Println(v2)
    var v37 int8 = v3
    fmt.Println(v2, v3, v37)
    if (32767 <= (v18 + (v18 >> ((0) & 63)))) {
        for fi38 := 0; fi38 < 8; fi38++ {
            fmt.Println(v7, v18, v18)
            _ = fi38
        }
        var v39 int = -799
        _ = v39
    }
    var v40 uint32 = 545
    var v41 uint32 = 65535
    if (false || true) {
        var v42 int16 = 42
        _ = v42
    } else {
        v4[(v40 & (v41 | uint32(v18)))] = v41
        delete(v4, (v40 << ((1) & 63)))
        if (1 == -1) {
            for fi43 := 0; fi43 < 12; fi43++ {
                if true {
                    v44 := GenS2{F0: 250, F1: 99, F2: uint16(v5)}
                    _ = v44
                }
                _ = fi43
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
