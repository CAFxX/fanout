package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 2
    var v4 float64 = 1.5
    v5 := make(map[string]bool)
    var v6 float32 = (float32((0.0 - 0.0)) / 1e10)
    delete(v2, 42)
    var v7 uint = 1
    var v8 float32 = -2.25
    v9 := v5["a"]
    if true {
        v3 = 32767
        v10 := []int{127, -1, 2147483647, -964}
        _ = v10
    }
    if (-32768 > -446) {
        var v11 uint64 = 0
        _ = v11
    } else {
        v12 := []int{808, 42, -696}
        var v13 bool = v9
        switch "a" {
        case "a":
            for fi14 := 0; fi14 < 4; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := GenS1{F0: 171}
                    v17 := GenS2{F0: 1}
                    v18 := GenS3{F0: "\n\t"}
                    var v19 float32 = v6
                    v20 := GenS4{F0: -343, F1: 42, F2: 2}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                v6 -= float32(v4)
                _ = fi14
            }
        }
        _ = v12
        _ = v13
    }
    if v9 {
        var v21 int32 = 42
        switch -631 {
        case -9223372036854775808:
            if (!(-128 == 42)) {
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS3{F0: "hello"}
                    fmt.Println(v4, fi22)
                    if (uint(v3) < (v7 / (((v7 - (v7 % ((v7) | 1)))) | 1))) {
                        v4 = (v4 + ((((v4 + v4) + v4) + v4) / v4))
                        var v24 float32 = 3.14159
                        v8 += v6
                        _ = v24
                    } else {
                        v3 = (^(v3))
                        switch (fi22 | (fi22 << ((v7) & 63))) {
                        case 322:
                            delete(v2, 0)
                        case 32768:
                            for fi25 := 0; fi25 < 9; fi25++ {
                                if (2 <= 2) {
                                    v9 = v9
                                    v4 -= v4
                                    if (v8 == (v8 + v6)) {
                                        v8 += float32(v4)
                                    }
                                } else {
                                    fmt.Println(v8, v7, v8)
                                }
                                fmt.Println(fi25)
                                _ = fi25
                            }
                        }
                        switch fi22 {
                        case 32767:
                            v26 := GenS5{F0: (v3 >> (((v7 | uint(v3))) & 63))}
                            _ = v26
                        case -1:
                            v4 -= (float64(float32(v4)) / v4)
                        case 361:
                            var v27 uint32 = 1
                            _ = v27
                        default:
                            switch "\n\t" {
                            case "hello":
                                v6 -= 3.14159
                            }
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                var v28 int32 = v21
                v29 := []int{769, 127, 1, 127}
                _ = v28
                _ = v29
            } else {
                v30 := GenS6{F0: 1}
                fmt.Println(v4, v7)
                v8 = v8
                _ = v30
            }
        case -31:
            var v31 uint32 = 1
            _ = v31
        case -1:
            v8 -= -2.25
        }
        _ = v21
    } else {
        v4 += v4
    }
    var v32 uint = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
