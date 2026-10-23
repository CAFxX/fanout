package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    var v5 int64 = 0
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 82, F1: -385, F2: 4294967295}
        if false {
            for fi8 := 0; fi8 < 12; fi8++ {
                v5 = -9223372036854775808
                v5 -= v5
                var v9 int8 = 42
                _ = fi8
                _ = v9
            }
            v5 = v5
            v10 := []int{2147483647}
            _ = v10
        } else {
            v11 := GenS2{F0: 42}
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v12 := GenS2{F0: 4294967295}
    _ = v12
    _ = v5
    return 1
}
func genFn1(p0 uint32, p1 uint64) uint8 {
    v13 := []int{-2147483648, 32768}
    var v14 uint = 4294967295
    v13 = append(v13, 0)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 44
}

func main() {
    var v1 int8 = 0
    var v2 uint64 = 0
    var v3 int32 = 2147483647
    var v4 uint32 = 919
    v1 -= v1
    v15 := []int{4294967295, 32768}
    if ((v1 | v1) >= (-(v1))) {
        _ = len(v15)
        switch 423 {
        case 1:
            fmt.Println(v1, v1)
        }
    }
    v4 -= v4
    var v16 int8 = v1
    for fi17 := 0; fi17 < 14; fi17++ {
        if false {
            if true {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v2 = v2
                    var v19 float64 = 1.5
                    if len(v15) > 0 {
                        v15[0] = fi17
                    }
                    _ = fi18
                    _ = v19
                }
                v16 -= (v1 | (v1 >> ((2) & 63)))
            } else {
                switch fi17 {
                case -9223372036854775808:
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if len(v15) > 0 {
                            v15[0] = fi17
                        }
                        _ = fi20
                    }
                default:
                    switch "hello" {
                    case "\n\t":
                        v2 = 42
                    case "hello":
                        v21 := GenS3{F0: 0}
                        _ = v21
                    default:
                        v4 -= v4
                    }
                }
            }
            v2 = 1
            v4 -= 4294967295
        }
        _ = fi17
    }
    var v22 bool = true
    if ((v2 < (v2 + v2)) && v22) {
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(fi23)
            fmt.Println(v2, fi23, v2)
            _ = fi23
        }
        v2 = uint64(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
