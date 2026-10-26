package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0() string {
    var v5 uint8 = 42
    v6 := []int{}
    _ = v5
    _ = v6
    return ("foo bar" + "foo bar")
}
func genFn1(p0 uint32, p1 uint8, p2 int) int64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p2, p1)
        fmt.Println(p1)
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    p0 += (p0 % ((p0) | 1))
    var v9 uint8 = (p1 & (p1 / ((p1) | 1)))
    if true {
        var v10 uint8 = (v9 % ((42) | 1))
        _ = v10
    } else {
        var v11 int8 = -128
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    v4 := []int{0, 2147483647, 4294967295}
    var v12 uint = 0
    for fi13 := 0; fi13 < 2; fi13++ {
        switch "a" {
        case "hello":
            v4 = append(v4, 2147483647)
        case "a":
            switch "\n\t" {
            case "\n\t":
                var v14 float32 = 3.14159
                _ = v14
            case "hello":
                v15 := GenS1{F0: 4294967295}
                _ = v15
            default:
                v16 := GenS1{F0: 1}
                _ = v16
            }
        }
        v12 -= (v12 % (((v12 * v12)) | 1))
        _ = fi13
    }
    fmt.Println(v3, v3)
    _ = len(v4)
    _ = len(v4)
    if false {
        v17 := v1[65535]
        fmt.Println(v17, v17, v17)
        var v18 string = "foo bar"
        _ = v17
        _ = v18
    }
    switch -9223372036854775808 {
    case 127:
        v3 = v3
    case 809:
        if (2 < uint8(v12)) {
            v19 := GenS2{F0: 2, F1: 0}
            var v20 int64 = 0
            _ = v19
            _ = v20
        } else {
            v3 += (v3 >> (((v12 | (v12 % ((2) | 1)))) & 63))
            fmt.Println(v3, v3)
        }
    case 128:
        v12 += 4294967295
    default:
        switch 32768 {
        case -1:
            var v21 string = "foo bar"
            _ = v21
        case 32767:
            fmt.Println(v12, v12)
        case 4294967295:
            v1[42] = "\n\t"
        }
    }
    var v22 int64 = -555
    v3 -= v3
    v23 := GenS3{F0: 171, F1: (v22 | (v22 - v22))}
    switch "" {
    case "hello":
        v1[1] = ("hello" + "hello")
    case "foo bar":
        fmt.Println(v22)
    case "":
        var v24 uint = 65535
        _ = v24
    default:
        var v25 int32 = -830
        _ = v25
    }
    var v26 int = 4294967295
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int = 323
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
