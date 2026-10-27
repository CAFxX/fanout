package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int64, p1 uint16, p2 string) uint16 {
    var v7 bool = true
    if (!(!(42 != 2147483647))) {
        v8 := []int{-938, 9223372036854775807, -1, 42}
        v7 = (0 == 42)
        _ = v8
    } else {
        switch p2 {
        case "a":
            p0 = p0
        default:
            p2 = p2
        }
        p2 = "hello"
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1(p0 int8, p1 uint, p2 int8) uint64 {
    if (float64(3.14159) >= (((float64(1e10) * -2.25) + 3.14159) / 1.5)) {
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(p2)
            var v10 float64 = 3.14159
            _ = fi9
            _ = v10
        }
        var v11 float64 = -2.25
        var v12 float64 = (1.5 - ((v11 - 1e10) / float64((0.1 + float32(float64((1e10 / 1.0)))))))
        _ = v11
        _ = v12
    } else {
        p1 = (p1 / ((p1) | 1))
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int8 = 111
            fmt.Println(p1)
            _ = fi13
            _ = v14
        }
        v15 := []int{4294967295}
        _ = v15
    }
    var v16 uint = p1
    p1 = (v16 ^ p1)
    var v17 float32 = -2.25
    v18 := []int{42, 42, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 18446744073709551615
    var v3 int8 = -128
    var v4 uint = 869
    var v5 uint = 2
    var v6 int32 = -1
    var v19 uint8 = 0
    v20 := []int{2147483647}
    if false {
        switch 764 {
        case -2147483648:
            var v21 int8 = (v3 % ((-31) | 1))
            _ = v21
        case 4294967295:
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 float32 = float32(3.14159)
                switch (fi22 / (((fi22 & fi22)) | 1)) {
                case -55:
                    var v24 int = (fi22 + 127)
                    _ = v24
                case 2147483647:
                    for fi25 := 0; fi25 < 4; fi25++ {
                        _ = len(v20)
                        if true {
                            v26 := GenS1{F0: 148, F1: (v1 | v1), F2: 510}
                            if len(v20) > 0 {
                                v20[0] = 127
                            }
                            fmt.Println(v4)
                            _ = v26
                        } else {
                            if false {
                                var v27 int16 = 1
                                v4 -= (v4 << (((v4 - 94)) & 63))
                                _ = v27
                            } else {
                                v28 := GenS2{F0: 87}
                                v29 := GenS3{F0: (v4 % (((v4 / (((v5 + (v4 ^ (v4 >> (((v5 % (((v5 ^ uint(v2))) | 1))) & 63))))) | 1))) | 1)), F1: v1}
                                _ = v28
                                _ = v29
                            }
                        }
                        fmt.Println(fi25, v23, v19)
                        _ = fi25
                    }
                case 128:
                    v19 = v19
                }
                _ = fi22
                _ = v23
            }
        case 9223372036854775807:
            var v30 float64 = -2.25
            _ = v30
        }
        var v31 float64 = 0.1
        var v32 int8 = 39
        _ = v31
        _ = v32
    }
    var v33 uint32 = 0
    var v34 int64 = int64(v19)
    v35 := GenS2{F0: 42}
    if len(v20) > 0 {
        v20[0] = -133
    }
    var v36 uint8 = (v19 | 127)
    v6 -= (v6 & int32(v19))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
