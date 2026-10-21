package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int; F2 int32 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = (255 <= 1)
    v3 := make(map[int]string)
    if true {
        switch "a" {
        case "a":
            var v4 uint = 784
            _ = v4
        case "\n\t":
            v5 := []int{4294967295, -2147483648}
            _ = v5
        case "hello":
            var v6 uint8 = 128
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 255}
        v8 := []int{32768}
        v8 = append(v8, 858)
        _ = v7
        _ = v8
    }
    var v9 bool = v2
    if v9 {
        fmt.Println(v2, v9, v1)
    }
    var v10 int16 = -570
    v11 := v3[-884]
    fmt.Println(v2, v1)
    switch "a" {
    case "":
        v12 := GenS2{F0: 124}
        _ = v12
    case "foo bar":
        v13 := []int{-714, -1, -228}
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := []int{4294967295, int(v1), fi14}
            fmt.Println(v1)
            switch "a" {
            case "a":
                switch "\n\t" {
                case "":
                    v1 = (v1 | (v1 / (((v1 - v1)) | 1)))
                case "hello":
                    fmt.Println(v9)
                default:
                    if false {
                        v2 = (126 < 28)
                        var v16 float32 = (float32(0.0) / ((float32(3.14159) / 1.0) * 3.14159))
                        switch "\n\t" {
                        case "\n\t":
                            if (v1 == 42) {
                                var v17 uint16 = 0
                                v18 := GenS3{F0: 128}
                                v19 := GenS4{F0: (v1 - v1), F1: (^(fi14)), F2: -1}
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        case "foo bar":
                            v3[(fi14 >> ((42) & 63))] = v11
                        default:
                            v1 = 309
                        }
                        _ = v16
                    } else {
                        var v20 int = (fi14 >> ((2) & 63))
                        v2 = ((false || v9) && (v1 < 1))
                        delete(v3, (fi14 % ((v20) | 1)))
                        _ = v20
                    }
                }
            }
            _ = fi14
            _ = v15
        }
    default:
        if v2 {
            fmt.Println(v11)
            v1 = (v1 % (((v1 | v1)) | 1))
            v21 := GenS5{F0: 65535}
            _ = v21
        } else {
            v22 := []int{4294967295, 127, -500, 128}
            v11 = (v11 + ("" + v11))
            fmt.Println(v1, v9, v1)
            _ = v22
        }
    }
    v23 := GenS2{F0: 2}
    var v24 int8 = 127
    v25 := GenS2{F0: 0}
    v26 := []int{}
    v27 := GenS1{F0: 726}
    v2 = v9
    if len(v26) > 0 {
        v26[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
