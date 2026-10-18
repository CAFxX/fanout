package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
type GenS5 struct { F0 string }
func genFn0(p0 uint, p1 int, p2 string) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{128, (fi7 / ((p1) | 1))}
        for fi9 := 0; fi9 < 13; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 int8 = -51
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 127
}
func genFn1(p0 int32) uint {
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{1, (fi11 + fi11)}
        _ = len(v12)
        _ = len(v12)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 65535, F1: 65535, F2: -133}
    p0 -= p0
    _ = p0
    _ = v13
    return 808
}

func main() {
    var v1 uint16 = 411
    var v2 int16 = 32767
    var v3 int64 = 1
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 int = 32768
    v14 := GenS2{F0: (!(false || false))}
    var v15 int64 = (v3 * v3)
    v16 := []int{}
    var v17 float64 = 3.14159
    v17 = v17
    if len(v16) > 0 {
        v16[0] = -9223372036854775808
    }
    v1 = (v1 - 98)
    for fi18 := 0; fi18 < 11; fi18++ {
        v2 = (v2 | v2)
        fmt.Println(v1, v6)
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    var v20 float32 = ((3.14159 + 3.14159) / 1.0)
    if true {
        v21 := GenS3{F0: 42, F1: (^(v3))}
        v22 := GenS4{F0: 240, F1: "\n\t", F2: "hello"}
        _ = v21
        _ = v22
    }
    if true {
        v3 -= v15
    }
    v23 := GenS5{F0: "foo bar"}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
