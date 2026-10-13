package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 int32, p2 int16) uint16 {
    fmt.Println(p2, p0, p1)
    v5 := []int{-524}
    v6 := GenS1{F0: -1, F1: 127, F2: (p2 / (((-(p2))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 702
    var v4 int8 = 0
    _ = len(v2)
    switch "" {
    case "a":
        var v7 int64 = -233
        _ = v7
    default:
        var v8 int16 = -331
        _ = v8
    }
    var v9 int = 42
    v10 := []int{(v9 ^ v9), (v9 / (((v9 & v9)) | 1)), v9}
    v3 = v3
    var v11 bool = ((42 > 245) || (!true))
    for fi12 := 0; fi12 < 14; fi12++ {
        v3 = (v3 >> ((255) & 63))
        for fi13 := 0; fi13 < 7; fi13++ {
            v9 += v9
            var v14 int64 = 127
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if (v4 <= (v4 & -128)) {
        v1 += (v1 * (v1 - 1.5))
        if (42 > 42) {
            v3 = v3
            var v15 uint8 = 0
            if ((^(v3)) != (^(v3))) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch (fi16 ^ 1) {
                    case -9223372036854775808:
                        v9 += 2147483647
                    }
                    switch fi16 {
                    case -116:
                        fmt.Println(v4, v15)
                    case -2147483648:
                        var v17 uint = 2
                        _ = v17
                    }
                    v18 := GenS1{F0: 4294967295, F1: (v9 & fi16), F2: -165}
                    _ = fi16
                    _ = v18
                }
            } else {
                fmt.Println(v9, v15)
            }
            _ = v15
        } else {
            for fi19 := 0; fi19 < 7; fi19++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                switch fi19 {
                case -9223372036854775808:
                    switch fi19 {
                    case -351:
                        switch "a" {
                        case "":
                            switch fi19 {
                            case 128:
                                v20 := GenS2{F0: -1}
                                _ = v20
                            case 0:
                                v3 = 0
                            case 127:
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    v22 := GenS1{F0: fi21, F1: v9, F2: -1}
                                    if (v4 == 0) {
                                        switch (fi19 & -1) {
                                        case -1:
                                            fmt.Println(v3, fi19)
                                        case -395:
                                            v4 -= v4
                                        case 4294967295:
                                            delete(v2, (v3 % ((v3) | 1)))
                                        }
                                        var v23 uint64 = 0
                                        v1 += (float32((-2.25 / 1.5)) * (float32((1.5 - -2.25)) + v1))
                                        _ = v23
                                    }
                                    if (v11 || v11) {
                                        v24 := GenS2{F0: 0}
                                        _ = v24
                                    }
                                    switch (-(fi19)) {
                                    case 0:
                                        var v25 float32 = (0.0 + 0.1)
                                        _ = v25
                                    case 2147483647:
                                        v9 += 82
                                    default:
                                        var v26 int = fi21
                                        _ = v26
                                    }
                                    var v27 int32 = -2147483648
                                    v28 := GenS1{F0: fi21, F1: fi19, F2: -368}
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    _ = fi21
                                    _ = v22
                                    _ = v27
                                    _ = v28
                                }
                            }
                        case "foo bar":
                            v29 := GenS1{F0: fi19, F1: (^(fi19)), F2: -1}
                            _ = v29
                        }
                    default:
                        v9 = fi19
                    }
                case -2147483648:
                    var v30 int8 = 120
                    _ = v30
                case 127:
                    var v31 uint = 65535
                    _ = v31
                default:
                    v9 -= fi19
                }
                _ = fi19
            }
            v4 = v4
            v1 = (-2.25 + v1)
        }
    }
    var v32 int32 = 0
    var v33 uint64 = 18446744073709551615
    v33 = 18446744073709551615
    v34 := GenS3{F0: int8(v9), F1: genFn0(v3, (v32 >> ((65535) & 63)), 42)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
