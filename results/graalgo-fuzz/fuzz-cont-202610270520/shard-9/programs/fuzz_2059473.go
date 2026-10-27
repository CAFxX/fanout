package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 bool }

func main() {
    v1 := []uint32{}
    var v2 uint = 1
    v3 := []int{2147483647, 283, 32768}
    var v4 uint = v2
    var v5 string = "\n\t"
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println(v2, v4, v2)
    v1 = append(v1, 60)
    v6 := GenS1{F0: -868, F1: 0}
    v4 = (^(v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS2{F0: 4294967295, F1: -1, F2: 2147483648}
    v8 := GenS3{F0: 128, F1: 0, F2: 218}
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            if true {
                v11 := GenS4{F0: -1, F1: 96}
                v5 = ""
                _ = v11
            }
            v5 = (v5 + ("" + v5))
            _ = fi10
        }
        v12 := GenS5{F0: 2147483647, F1: 127}
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v2, fi9)
            _ = fi13
        }
        _ = fi9
        _ = v12
    }
    switch "\n\t" {
    case "foo bar":
        for fi14 := 0; fi14 < 11; fi14++ {
            v5 = v5
            v4 -= v4
            _ = fi14
        }
    case "hello":
        var v15 int8 = 116
        _ = v15
    case "":
        var v16 int32 = -1
        _ = v16
    default:
        fmt.Println(v4)
    }
    fmt.Println(v5)
    v5 = "\n\t"
    var v17 uint64 = 4294967295
    v18 := GenS6{F0: v2, F1: -32768, F2: false}
    var v19 int = 591
    var v20 int64 = 17
    switch "a" {
    case "a":
        if true {
            switch v19 {
            case -1:
                v19 += v19
            }
            if (!true) {
                fmt.Println(v17)
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v2 = v2
                        fmt.Println(fi22, v20)
                        var v23 uint64 = v17
                        _ = len(v1)
                        var v24 uint8 = 255
                        if true {
                            var v25 uint16 = 0
                            _ = v25
                        } else {
                            var v26 uint64 = 409
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
                v5 = v5
            }
            var v27 int = 48
            _ = v27
        } else {
            v28 := GenS6{F0: (v2 ^ (v2 >> (((v4 - v4)) & 63))), F1: -513, F2: (2 <= 42)}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
