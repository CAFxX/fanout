package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int8; F2 int32 }
type GenS4 struct { F0 int32; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 string = ("a" + v2)
    if true {
        var v6 float32 = (1e10 - (0.0 / (1e10 / 1e10)))
        v6 = v6
        v2 = ""
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 string = ""
        fmt.Println(v2)
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 2
    v9 += v9
    v10 := GenS1{F0: 2}
    var v11 int8 = 100
    v12 := GenS1{F0: (^(v11))}
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            var v14 int = -314
            var v15 bool = false
            _ = v14
            _ = v15
        }
        if (((true || true) || (false || false)) && (921 > v9)) {
            v16 := GenS2{F0: -1}
            v17 := GenS3{F0: v2, F1: (-(v11)), F2: 0}
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v18 := []int{4294967295}
    switch -203 {
    case 9223372036854775807:
        fmt.Println(v2, v11, v9)
    default:
        switch -1 {
        case 9223372036854775807:
            var v19 bool = (true || false)
            _ = v19
        case -996:
            fmt.Println(v11, v5)
        }
    }
    var v20 int8 = v11
    for fi21 := 0; fi21 < 12; fi21++ {
        switch fi21 {
        case 2147483647:
            fmt.Println(v20, v11, fi21)
        case 42:
            v20 -= v11
        }
        var v22 uint8 = 2
        _ = fi21
        _ = v22
    }
    switch -2147483648 {
    case -9223372036854775808:
        switch -1 {
        case -2147483648:
            fmt.Println(v11, v9, v11)
        default:
            if len(v18) > 0 {
                v18[0] = 4294967295
            }
        }
    case -475:
        var v23 int8 = (v11 << ((1) & 63))
        _ = v23
    case 0:
        v24 := GenS4{F0: -2147483648, F1: true, F2: 430}
        _ = v24
    default:
        v9 = v9
    }
    fmt.Println(v5, v11)
    var v25 uint = 731
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
