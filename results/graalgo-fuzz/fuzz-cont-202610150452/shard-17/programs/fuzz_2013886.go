package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int) uint16 {
    fmt.Println(p1)
    v5 := GenS1{F0: 620, F1: -635}
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = (fi6 & 0)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8)
        var v9 uint16 = 1
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(fi8, fi8)
            p1 = (fi8 ^ fi8)
            p0 = (^(v9))
            p1 += (fi8 / (((p1 - 2147483647)) | 1))
            var v11 uint32 = 0
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 uint = 775
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return p0
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{65535, 297, 0}
    var v3 int64 = 10
    v4 := make(map[int]bool)
    for fi13 := 0; fi13 < 9; fi13++ {
        switch "foo bar" {
        case "a":
            v3 = (v3 % (((v3 - v3)) | 1))
        case "\n\t":
            switch "\n\t" {
            case "":
                v1 = v1
            case "hello":
                v14 := GenS2{F0: 336}
                _ = v14
            case "foo bar":
                v15 := GenS3{F0: 42}
                _ = v15
            }
        }
        _ = fi13
    }
    if true {
        var v16 int16 = -433
        var v17 int32 = -249
        _ = v16
        _ = v17
    }
    v3 = 42
    v3 -= (-(v3))
    fmt.Println(v1)
    v18 := GenS3{F0: 128}
    var v19 float32 = -2.25
    for fi20 := 0; fi20 < 10; fi20++ {
        switch 1 {
        case 456:
            var v21 bool = (190 >= 2)
            _ = v21
        case 4294967295:
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        switch 1 {
        case -2147483648:
            var v22 uint = 906
            _ = v22
        default:
            var v23 string = ("\n\t" + "foo bar")
            _ = v23
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi20
    }
    var v24 int8 = (v1 >> ((287) & 63))
    v19 = v19
    switch 32768 {
    case -397:
        v25 := GenS3{F0: 0}
        _ = v25
    default:
        var v26 uint8 = 177
        _ = v26
    }
    _ = len(v4)
    var v27 int32 = 554
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
