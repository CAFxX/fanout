package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    var v7 uint = 118
    fmt.Println(v7)
    var v8 int8 = -1
    for fi9 := 0; fi9 < 16; fi9++ {
        v8 += (v8 & v8)
        _ = fi9
    }
    v10 := GenS1{F0: (v8 ^ v8)}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32, p1 int8) uint {
    v11 := GenS1{F0: p1}
    var v12 float32 = 1e10
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            p0 = 90
            v14 := []int{fi13, -904, (-(fi13))}
            v15 := GenS2{F0: 660, F1: 407, F2: int16(p1)}
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := []int{128, 127, 4294967295, 0}
        for fi17 := 0; fi17 < 10; fi17++ {
            p0 = (p0 / (((p0 & 583)) | 1))
            _ = fi17
        }
        _ = v16
    } else {
        if (false || true) {
            v18 := []int{-128, -2147483648, -9223372036854775808, 2147483647}
            p0 += p0
            v18 = append(v18, int(p0))
            _ = v18
        }
    }
    v19 := []int{0, 128, 1, -921}
    v19 = append(v19, -9223372036854775808)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v19
    return 292
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 int = 32768
    var v4 uint64 = 0
    v5 := []uint32{}
    var v6 int16 = 1
    fmt.Println(v4)
    v20 := v1[65535]
    fmt.Println(v3, v20, v3)
    switch 342 {
    case 9223372036854775807:
        v4 -= (v4 | (v4 & v4))
    case 1:
        var v21 string = "a"
        _ = v21
    case -823:
        v22 := GenS1{F0: 100}
        _ = v22
    }
    v1[v20] = v20
    v1[v20] = (v20 & v20)
    if (0 >= v3) {
        v20 -= v20
    } else {
        if true {
            var v23 bool = (((v20 <= 135) && true) || false)
            if v23 {
                v5 = append(v5, (v20 | v20))
            } else {
                v20 = v20
                _ = len(v5)
                var v24 int32 = 16
                _ = v24
            }
            var v25 uint16 = 0
            _ = v23
            _ = v25
        }
        var v26 string = ("a" + "")
        if true {
            switch v3 {
            case -1:
                switch (^(v3)) {
                case 32768:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        switch 553 {
                        case 0:
                            v28 := v1[v20]
                            _ = v28
                        case 2147483647:
                            if (false || false) {
                                switch v26 {
                                case "":
                                    fmt.Println(v4, v20)
                                case "\n\t":
                                    switch v26 {
                                    case "hello":
                                        var v29 float64 = (float64(3.14159) * 0.1)
                                        _ = v29
                                    case "":
                                        var v30 int32 = 2147483647
                                        _ = v30
                                    case "a":
                                        v31 := GenS3{F0: 32767}
                                        _ = v31
                                    default:
                                        for fi32 := 0; fi32 < 20; fi32++ {
                                            switch "a" {
                                            case "hello":
                                                v3 += fi32
                                            case "":
                                                var v33 string = (v26 + ("foo bar" + "hello"))
                                                _ = v33
                                            case "\n\t":
                                                if ((1 >= 1) && false) {
                                                    _ = len(v5)
                                                    v20 = 4294967295
                                                }
                                            }
                                            v20 = 65535
                                            v34 := GenS3{F0: 255}
                                            _ = fi32
                                            _ = v34
                                        }
                                    }
                                }
                                fmt.Println(v20, v6)
                                v3 = (v3 % ((fi27) | 1))
                            } else {
                                var v35 int64 = -157
                                var v36 int8 = int8(fi27)
                                _ = v35
                                _ = v36
                            }
                        case 128:
                            _ = len(v5)
                        default:
                            for fi37 := 0; fi37 < 8; fi37++ {
                                if false {
                                    if len(v5) > 0 {
                                        v5[0] = (v20 & (^(v20)))
                                    }
                                    var v38 int8 = 100
                                    var v39 int = (v3 - v3)
                                    _ = v38
                                    _ = v39
                                }
                                v4 = (^(v4))
                                _ = len(v2)
                                fmt.Println(v20)
                                _ = fi37
                            }
                        }
                        var v40 int = 963
                        v3 = 9223372036854775807
                        _ = fi27
                        _ = v40
                    }
                default:
                    var v41 int = v3
                    _ = v41
                }
            case 2147483647:
                v4 += (v4 - (v4 / ((v4) | 1)))
            default:
                v20 = v20
            }
            var v42 uint = 751
            _ = v42
        }
        _ = v26
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        for fi44 := 0; fi44 < 8; fi44++ {
            v5 = append(v5, 65535)
            _ = fi44
        }
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
