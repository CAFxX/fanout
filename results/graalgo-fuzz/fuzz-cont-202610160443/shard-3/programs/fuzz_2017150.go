package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 uint16, p2 string) int32 {
    v7 := GenS1{F0: -746, F1: 1}
    p2 = p2
    var v8 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 uint64) int {
    switch ("hello" + "") {
    case "foo bar":
        fmt.Println(p1)
    case "hello":
        v9 := []int{127}
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 0, F1: p1}
            if true {
                v13 := []int{fi10}
                _ = v13
            }
            p0 -= p0
            _ = fi11
            _ = v12
        }
        var v14 int32 = 42
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := []int{(fi15 ^ 32767), (fi15 * fi15), fi15, fi15}
        if len(v16) > 0 {
            v16[0] = 0
        }
        var v17 float64 = 0.0
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(p0, fi18, p0)
        v19 := []int{}
        _ = fi18
        _ = v19
    }
    if true {
        v20 := GenS1{F0: -1, F1: 592}
        switch 9223372036854775807 {
        case 4294967295:
            v21 := GenS1{F0: p0, F1: p1}
            _ = v21
        }
        _ = v20
    } else {
        v22 := GenS2{F0: 1}
        _ = v22
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 float32 = 1e10
    v4 := make(map[uint32]int)
    var v5 int8 = 85
    var v6 uint8 = 23
    for fi23 := 0; fi23 < 13; fi23++ {
        v5 = int8(v6)
        v4[2147483647] = (fi23 & genFn1(42, 0))
        switch (fi23 >> ((845) & 63)) {
        case -18:
            v24 := []int{}
            _ = v24
        case 4294967295:
            v6 -= v6
        default:
            v25 := []int{genFn1(-1, 1), fi23}
            _ = v25
        }
        _ = fi23
    }
    var v26 int = 0
    var v27 int8 = (^(v5))
    if (4294967295 != 1) {
        v28 := GenS3{F0: 428, F1: 759}
        if (((v6 / (((v6 * (v6 / (((v6 - (v6 + 0))) | 1)))) | 1)) <= 140) && true) {
            var v29 int64 = -1
            _ = v29
        }
        _ = len(v1)
        _ = v28
    } else {
        fmt.Println(v27, v3, v27)
        v30 := GenS3{F0: 32767, F1: 1}
        _ = v30
    }
    var v31 uint32 = 1
    v32 := []int{genFn1(42, 2), 32767, (v26 - genFn1(1000, 42))}
    _ = len(v32)
    var v33 float32 = float32(0.1)
    var v34 bool = true
    v35 := GenS4{F0: -2147483648}
    if v34 {
        var v36 string = "hello"
        v37 := GenS2{F0: 0}
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
