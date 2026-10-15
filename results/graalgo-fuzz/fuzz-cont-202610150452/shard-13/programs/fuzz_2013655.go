package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 string }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint64, p1 bool) uint8 {
    p1 = p1
    var v6 int = 32768
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1(p0 int16) int32 {
    v7 := []int{0, 1, 42, -240}
    var v8 float64 = (-2.25 * (3.14159 - float64(0.0)))
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int64 = 886
        fmt.Println(v8, p0, fi9)
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "\n\t":
        p0 = p0
    case "":
        v7 = append(v7, -9223372036854775808)
    default:
        v11 := GenS1{F0: 32767, F1: 4294967295, F2: 712}
        _ = v11
    }
    var v12 int = 127
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 2
    var v3 int = -9223372036854775808
    var v4 uint64 = 0
    v5 := make(map[int]uint32)
    if (!(false && (351 > 42))) {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint32 = 851
            _ = fi13
            _ = v14
        }
        _ = len(v5)
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float64 = 3.14159
            if true {
                switch fi15 {
                case -9223372036854775808:
                    var v17 int8 = 110
                    _ = v17
                case 32768:
                    var v18 uint8 = 127
                    _ = v18
                case 2147483647:
                    if true {
                        if (128 != v2) {
                            if (true && (!(-1 == 514))) {
                                fmt.Println(v16, v1)
                                v3 -= (fi15 >> (((v1 + 42)) & 63))
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    var v20 float32 = 0.1
                                    if false {
                                        fmt.Println(v3, v4, v20)
                                    } else {
                                        v21 := GenS1{F0: fi15, F1: v1, F2: uint(v2)}
                                        v22 := GenS2{F0: 570, F1: 100}
                                        v16 += (v16 / v16)
                                        _ = v21
                                        _ = v22
                                    }
                                    v5[(fi15 % ((-115) | 1))] = 42
                                    v23 := GenS3{F0: true, F1: int16(v4), F2: "\n\t"}
                                    fmt.Println(fi19)
                                    _ = fi19
                                    _ = v20
                                    _ = v23
                                }
                            }
                        }
                        fmt.Println(v1)
                    }
                default:
                    v24 := GenS1{F0: -1, F1: v1, F2: (v1 + (v1 * v1))}
                    _ = v24
                }
            }
            _ = fi15
            _ = v16
        }
    } else {
        switch "foo bar" {
        case "hello":
            v4 += v4
        case "foo bar":
            _ = len(v5)
        }
        fmt.Println(v1, v3)
    }
    if ((v4 / (((^(v4))) | 1)) == v4) {
        v25 := GenS2{F0: v4, F1: 100}
        _ = v25
    } else {
        v1 = v1
        if true {
            if true {
                v4 += 0
                v26 := GenS4{F0: (v2 / (((v2 ^ 42)) | 1)), F1: 1000}
                _ = v26
            }
            var v27 int8 = 100
            _ = v27
        } else {
            v28 := GenS2{F0: v4, F1: 1}
            fmt.Println(v3, v4)
            _ = v28
        }
        v29 := []int{v3, -776, (v3 * (-(v3)))}
        _ = v29
    }
    v4 = (^(v4))
    v30 := GenS3{F0: ((^(v2)) != v2), F1: 42, F2: "a"}
    delete(v5, (^(v3)))
    var v31 int = v3
    delete(v5, 16)
    if (-1 < 0) {
        v2 = (^(v2))
        v32 := GenS2{F0: 437, F1: 1}
        _ = v32
    } else {
        fmt.Println(v3, v2)
        if false {
            var v33 uint32 = 42
            var v34 int32 = genFn1(0)
            for fi35 := 0; fi35 < 20; fi35++ {
                v34 = v34
                v34 -= v34
                _ = fi35
            }
            _ = v33
            _ = v34
        } else {
            if (v1 <= v1) {
                v31 = v31
                var v36 uint32 = 1
                v3 = (v31 - 4294967295)
                _ = v36
            }
            v31 -= v31
        }
    }
    v31 = v3
    v37 := []int{127, v31, 2147483647}
    fmt.Println(v31)
    var v38 int64 = 884
    var v39 int32 = 715
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
