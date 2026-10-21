package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 string, p2 bool) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(p1, fi7)
            p2 = p2
            v8 := []int{32767, (fi6 >> ((2) & 63)), (-(fi7)), fi6}
            _ = fi7
            _ = v8
        }
        v9 := []int{fi6, (-(fi6)), (fi6 - fi6)}
        _ = fi6
        _ = v9
    }
    var v10 int16 = 0
    v11 := GenS1{F0: 1, F1: -535}
    var v12 int8 = 127
    switch -9223372036854775808 {
    case 127:
        for fi13 := 0; fi13 < 19; fi13++ {
            v12 += v12
            v14 := GenS2{F0: -2147483648}
            var v15 int = (fi13 ^ (fi13 + fi13))
            _ = fi13
            _ = v14
            _ = v15
        }
    case 9223372036854775807:
        var v16 uint = 850
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 54
}
func genFn1(p0 int8) bool {
    var v17 uint64 = 0
    var v18 uint64 = (v17 ^ 2)
    var v19 int32 = 174
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(v17, v18)
        _ = fi20
    }
    v21 := GenS2{F0: -367}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    _ = v21
    return true
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 uint32 = 1
    v4 := []uint32{}
    var v5 string = "hello"
    var v22 uint = 255
    var v23 int32 = -1
    v4 = append(v4, v3)
    var v24 float64 = (3.14159 + 0.1)
    fmt.Println(v22)
    v24 += (float64(-2.25) / v24)
    for fi25 := 0; fi25 < 11; fi25++ {
        if true {
            switch v5 {
            case "\n\t":
                var v26 string = ("foo bar" + (v5 + v5))
                _ = v26
            default:
                var v27 float32 = 0.0
                _ = v27
            }
        }
        _ = fi25
    }
    if true {
        switch "a" {
        case "":
            var v28 float32 = float32((v24 - v24))
            _ = v28
        case "a":
            var v29 int8 = -91
            _ = v29
        default:
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS3{F0: -973}
                var v32 float64 = -2.25
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            var v34 int32 = -605
            _ = fi33
            _ = v34
        }
    } else {
        v35 := GenS3{F0: -48}
        _ = v35
    }
    fmt.Println(v22, v24)
    if false {
        var v36 bool = true
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
