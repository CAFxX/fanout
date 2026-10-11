package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
func genFn0(p0 bool, p1 int16) string {
    switch 9223372036854775807 {
    case 2147483647:
        switch "hello" {
        case "":
            p1 -= p1
        case "\n\t":
            for fi4 := 0; fi4 < 5; fi4++ {
                v5 := []int{fi4, (fi4 << ((202) & 63)), 32768}
                _ = fi4
                _ = v5
            }
        }
    case 32767:
        switch -2147483648 {
        case 42:
            fmt.Println(p1, p1)
        }
    case -592:
        if false {
            switch "foo bar" {
            case "\n\t":
                p1 = p1
            case "a":
                switch "\n\t" {
                case "\n\t":
                    for fi6 := 0; fi6 < 3; fi6++ {
                        switch (fi6 % ((fi6) | 1)) {
                        case 42:
                            v7 := GenS1{F0: 42, F1: 100}
                            _ = v7
                        case 127:
                            for fi8 := 0; fi8 < 17; fi8++ {
                                p1 = p1
                                p0 = (255 > 32767)
                                v9 := []int{32767, 4294967295}
                                _ = fi8
                                _ = v9
                            }
                        case -9223372036854775808:
                            for fi10 := 0; fi10 < 11; fi10++ {
                                var v11 int32 = -90
                                _ = fi10
                                _ = v11
                            }
                        default:
                            p1 = p1
                        }
                        if false {
                            p0 = p0
                            v12 := []int{(fi6 >> ((0) & 63)), (fi6 ^ (fi6 | fi6)), (fi6 % (((fi6 + fi6)) | 1))}
                            fmt.Println(p0)
                            _ = v12
                        } else {
                            var v13 float32 = 3.14159
                            v13 = float32(1e10)
                            _ = v13
                        }
                        fmt.Println(p1, p0, p0)
                        _ = fi6
                    }
                case "":
                    v14 := GenS2{F0: 0, F1: 0}
                    _ = v14
                case "hello":
                    var v15 uint8 = 255
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v17 := GenS3{F0: 210, F1: -128}
                        p0 = (0.1 > 1.5)
                        v18 := GenS2{F0: 18446744073709551615, F1: 1}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 int8 = 93
                _ = fi19
                _ = v20
            }
            switch 4294967295 {
            case 2147483647:
                fmt.Println(p1, p1)
            default:
                p0 = ((!p0) || (!(4294967295 < 1)))
            }
        }
    }
    switch -981 {
    case 32768:
        v21 := GenS1{F0: 42, F1: 100}
        _ = v21
    case -974:
        v22 := []int{145, 2147483647}
        _ = v22
    default:
        p1 = 1
    }
    p0 = (492 <= 2)
    _ = p0
    _ = p1
    return "hello"
}
func genFn1(p0 int16, p1 int32, p2 uint32) int32 {
    var v23 uint = 0
    p2 -= (p2 + 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p1
}

func main() {
    v1 := []string{}
    var v2 uint8 = 128
    var v3 float32 = (3.14159 / 1.0)
    var v24 int = 42
    v25 := GenS1{F0: -8, F1: int8(v2)}
    v24 = (^(v24))
    if (!(32767 == 1000)) {
        var v26 int64 = -9223372036854775808
        _ = v26
    }
    v27 := GenS1{F0: -92, F1: -128}
    v28 := GenS4{F0: 0, F1: 802, F2: int64(v2)}
    v29 := GenS2{F0: 37, F1: 1}
    v30 := GenS5{F0: "hello", F1: -743, F2: 0}
    var v31 uint = 482
    v31 = v31
    var v32 uint16 = uint16(v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
