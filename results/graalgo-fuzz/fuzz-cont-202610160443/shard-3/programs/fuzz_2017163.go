package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
func genFn0(p0 int8) uint32 {
    v7 := []int{-884, 1, 4294967295}
    for fi8 := 0; fi8 < 6; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 int8, p2 int) int32 {
    p1 += p1
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint8 = 127
        _ = fi9
        _ = v10
    }
    var v11 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := []uint32{}
    var v5 int8 = 100
    var v6 uint64 = 982
    var v12 uint8 = 128
    switch "a" {
    case "foo bar":
        var v13 uint16 = uint16(v12)
        _ = v13
    case "hello":
        var v14 uint16 = 42
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 15; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                v5 = 113
                _ = fi16
            }
            v12 += (v12 << ((18446744073709551615) & 63))
            _ = fi15
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if (false && false) {
            if false {
                v18 := GenS1{F0: 1}
                v19 := GenS1{F0: 0}
                var v20 int32 = -68
                _ = v18
                _ = v19
                _ = v20
            } else {
                v5 = (v5 + v5)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v21 int32 = 802
            _ = v21
        }
        v4 = append(v4, genFn0(v5))
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float32 = -2.25
            _ = fi22
            _ = v23
        }
        _ = fi17
    }
    v5 = v5
    switch "" {
    case "\n\t":
        fmt.Println(v12)
    case "":
        fmt.Println(v12)
    case "a":
        if false {
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 uint64 = v6
                var v26 uint8 = v12
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 int = 9223372036854775807
                for fi29 := 0; fi29 < 8; fi29++ {
                    v3[fi27] = "hello"
                    if true {
                        v30 := GenS2{F0: 2, F1: true, F2: v5}
                        var v31 string = "foo bar"
                        if ((-2.25 - 1.5) != 1e10) {
                            var v32 int = 1
                            v33 := GenS1{F0: 1}
                            if false {
                                var v34 int64 = 9223372036854775807
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        } else {
                            v35 := GenS3{F0: (118 > 4294967295), F1: 1, F2: 1}
                            if (!true) {
                                var v36 uint = 18446744073709551615
                                _ = v36
                            }
                            _ = v35
                        }
                        _ = v30
                        _ = v31
                    } else {
                        v28 -= v28
                        var v37 uint = 495
                        var v38 uint8 = (v12 ^ v12)
                        _ = v37
                        _ = v38
                    }
                    if false {
                        v5 -= (v5 / ((v5) | 1))
                        var v39 float32 = float32(float64(float32(-2.25)))
                        var v40 int = v28
                        _ = v39
                        _ = v40
                    }
                    _ = fi29
                }
                _ = fi27
                _ = v28
            }
        }
    }
    switch "hello" {
    case "\n\t":
        var v41 int = -1
        _ = v41
    case "foo bar":
        var v42 uint8 = v12
        _ = v42
    }
    v5 = (v5 * (-(v5)))
    var v43 float64 = ((1.5 - 3.14159) / 3.14159)
    v12 -= v12
    var v44 int16 = -690
    var v45 uint32 = 357
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
