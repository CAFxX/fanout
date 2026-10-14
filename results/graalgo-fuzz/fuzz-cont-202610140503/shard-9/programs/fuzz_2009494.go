package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16) int8 {
    var v7 bool = false
    v8 := GenS1{F0: "a"}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int64 = 166
        v11 := GenS2{F0: -128, F1: 219, F2: (fi9 + fi9)}
        var v12 int64 = 1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int = 42
    var v14 bool = v7
    _ = p0
    _ = v13
    _ = v14
    _ = v7
    _ = v8
    return -61
}

func main() {
    var v1 int32 = 42
    v2 := []uint32{}
    var v3 uint8 = 1
    var v4 uint32 = 894
    v5 := []uint32{766, (v4 & (v4 % ((uint32(v3)) | 1)))}
    v6 := make(map[int]uint32)
    if (417 == v4) {
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS3{F0: 495}
            v17 := GenS2{F0: 100, F1: 65535, F2: (fi15 >> ((45) & 63))}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 float64 = 1e10
        _ = v18
    }
    var v19 int32 = (v1 * -2147483648)
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float64 = (0.0 + 0.0)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            delete(v6, (^(fi23)))
            fmt.Println(v4)
            if true {
                var v24 uint32 = (v4 | 2147483647)
                if (4294967295 > 0) {
                    v24 = (v4 - (^(v24)))
                    v3 += (v3 ^ v3)
                    var v25 int8 = 42
                    _ = v25
                }
                _ = v24
            }
            _ = fi23
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int16 = 1
            if (1 >= v4) {
                if (false || (776 == v27)) {
                    delete(v6, fi22)
                    var v28 int32 = (v1 % (((v19 ^ 2147483647)) | 1))
                    if true {
                        var v29 bool = false
                        switch "\n\t" {
                        case "foo bar":
                            var v30 uint8 = v3
                            _ = len(v6)
                            _ = v30
                        case "hello":
                            if v29 {
                                var v31 uint64 = 4294967295
                                if v29 {
                                    v31 -= (v31 - (^(v31)))
                                    switch fi26 {
                                    case 127:
                                        fmt.Println(v29, v1)
                                    }
                                } else {
                                    v4 += (v4 & (v4 % (((v4 * 2)) | 1)))
                                }
                                _ = v31
                            } else {
                                _ = len(v6)
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v2 = append(v2, v4)
                                }
                                v32 := GenS3{F0: v27}
                                _ = v32
                            }
                        default:
                            var v33 uint = uint(v3)
                            _ = v33
                        }
                        fmt.Println(fi22)
                        v28 -= v28
                        _ = v29
                    }
                    _ = v28
                }
                fmt.Println(v3)
            }
            _ = fi26
            _ = v27
        }
        _ = fi22
    }
    var v34 int8 = 127
    var v35 int16 = 14
    switch -1 {
    case -892:
        v35 = v35
    case 4294967295:
        v1 = v19
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS3{F0: (v35 / ((v35) | 1))}
        _ = fi36
        _ = v37
    }
    if (genFn0(v35) <= (v34 ^ genFn0(int16(v4)))) {
        fmt.Println(v19, v35, v4)
    }
    switch 127 {
    case -886:
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 uint8 = 2
            _ = fi38
            _ = v39
        }
    case 127:
        switch 32767 {
        case 32768:
            v34 -= v34
        }
    default:
        v5 = append(v5, 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
