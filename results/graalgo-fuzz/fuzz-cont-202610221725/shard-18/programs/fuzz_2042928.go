package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0(p0 bool, p1 uint64, p2 int8) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p1 -= p1
        var v8 uint8 = 127
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    v10 := GenS1{F0: 1, F1: p0}
    switch 2147483647 {
    case 128:
        p2 -= (p2 + p2)
    case -1:
        var v11 int64 = -227
        _ = v11
    case 9223372036854775807:
        v12 := GenS2{F0: 680, F1: 0}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int16, p1 bool) int {
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float64 = (0.1 / 0.1)
        if (337 == 0) {
            var v15 uint16 = 232
            v14 -= (0.1 + (-2.25 + v14))
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    v16 := GenS1{F0: 42, F1: (!(692 == 65535))}
    var v17 int64 = 0
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -1
    v3 := []int{32767, -876, -699, -985}
    var v4 float64 = 3.14159
    var v5 uint32 = 0
    var v6 uint = 18446744073709551615
    if false {
        if true {
            v18 := GenS3{F0: 255}
            _ = v18
        }
        v1 = v1
    }
    if (v2 <= 0) {
        v3 = append(v3, 9223372036854775807)
        _ = len(v3)
    }
    v19 := GenS4{F0: v6}
    v2 += (v2 % (((v2 + v2)) | 1))
    fmt.Println(v5, v5, v2)
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int16 = 32767
        var v22 int16 = 32767
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v2)
        v5 -= 2147483648
        v24 := GenS5{F0: 42, F1: 386, F2: 556}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        _ = len(v3)
        _ = fi25
    }
    fmt.Println(v4, v2)
    v6 = 0
    var v26 int8 = (v2 | (v2 >> ((v6) & 63)))
    v27 := GenS1{F0: 53, F1: false}
    v1 = "\n\t"
    if len(v3) > 0 {
        v3[0] = -1
    }
    var v28 int8 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        for fi29 := 0; fi29 < 10; fi29++ {
            v1 = "\n\t"
            v4 = 1e10
            _ = fi29
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 uint64 = 851
            _ = fi30
            _ = v31
        }
        v26 = (v2 & (v26 / ((v2) | 1)))
    } else {
        var v32 int16 = 42
        v4 = v4
        v3 = append(v3, -1)
        _ = v32
    }
    var v33 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
