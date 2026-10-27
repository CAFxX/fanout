package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0() int32 {
    v5 := GenS1{F0: 21, F1: 127, F2: 311}
    v6 := GenS1{F0: 100, F1: -621, F2: 30}
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi8)
            v9 := []int{}
            if false {
                var v10 uint16 = 0
                if false {
                    if (438 > 42) {
                        fmt.Println(v10, fi7, v10)
                        v10 -= (v10 % (((v10 - 32767)) | 1))
                        if true {
                            switch fi8 {
                            case 650:
                                v11 := GenS2{F0: (uint(v10) < 286), F1: -128}
                                _ = v11
                            case 2147483647:
                                v10 += (v10 * v10)
                            }
                        }
                    }
                    v10 += (v10 ^ v10)
                } else {
                    fmt.Println(v10, fi7)
                    _ = len(v9)
                }
                var v12 uint8 = uint8(v10)
                _ = v10
                _ = v12
            } else {
                v13 := GenS2{F0: true, F1: -34}
                fmt.Println(fi7, fi8)
                switch -2147483648 {
                case -2147483648:
                    var v14 uint16 = 0
                    _ = v14
                default:
                    var v15 int16 = -1
                    _ = v15
                }
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        var v16 uint16 = 255
        _ = fi7
        _ = v16
    }
    var v17 uint8 = 78
    _ = v17
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint32, p1 uint16) uint {
    var v18 string = "a"
    for fi19 := 0; fi19 < 11; fi19++ {
        p0 -= p0
        var v20 int64 = -9223372036854775808
        if false {
            for fi21 := 0; fi21 < 9; fi21++ {
                if (!false) {
                    v22 := GenS1{F0: 83, F1: fi21, F2: p1}
                    _ = v22
                } else {
                    switch 0 {
                    case 128:
                        v23 := []int{(fi21 / ((42) | 1))}
                        _ = v23
                    case -1:
                        var v24 int64 = (-(v20))
                        _ = v24
                    }
                    var v25 bool = true
                    _ = v25
                }
                _ = fi21
            }
        }
        _ = fi19
        _ = v20
    }
    v26 := []int{32767, 42, -1, 798}
    var v27 float32 = float32((1.5 * 0.1))
    var v28 uint16 = (p1 + (p1 % ((p1) | 1)))
    _ = p0
    _ = p1
    _ = v18
    _ = v26
    _ = v27
    _ = v28
    return 426
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v1 = 561
    var v29 string = "hello"
    v30 := GenS2{F0: false, F1: 14}
    var v31 float64 = ((1.5 + float64(1e10)) - 1.5)
    v32 := []int{-305, -2147483648, -9223372036854775808, 213}
    switch -326 {
    case -295:
        v31 = (float64((1.5 / 0.1)) / (v31 / v31))
    case 127:
        v1 = v1
    case 4294967295:
        switch 4294967295 {
        case 2147483647:
            fmt.Println(v1)
        default:
            v1 += v1
        }
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v31 = 3.14159
        v29 = "\n\t"
        _ = fi33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if true {
            if true {
                v35 := GenS3{F0: -1, F1: 1}
                fmt.Println(v29, v29)
                _ = v35
            }
        }
        v31 = v31
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
