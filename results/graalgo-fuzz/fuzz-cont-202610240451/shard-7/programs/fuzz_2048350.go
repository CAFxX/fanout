package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint) int8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{fi6, fi6, fi6}
        v8 := GenS1{F0: 2}
        v9 := GenS2{F0: (true && true), F1: -36, F2: p1}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 518
    v11 := GenS2{F0: true, F1: 0, F2: p1}
    var v12 string = "foo bar"
    v13 := GenS3{F0: -113}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 0
}
func genFn1(p0 int64) uint8 {
    p0 = p0
    fmt.Println(p0)
    p0 += 9223372036854775807
    _ = p0
    return 66
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483647, 1}
    v4 := make(map[int]int)
    var v5 float64 = 3.14159
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 string = ""
    if true {
        var v15 uint64 = 134
        fmt.Println(v15, v14)
        _ = v15
    } else {
        v16 := GenS1{F0: 347}
        v17 := v2[4294967295]
        _ = v16
        _ = v17
    }
    v14 = "a"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v14, v5)
    v18 := v2[42]
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS3{F0: (fi19 % (((-(fi19))) | 1))}
        var v21 int = fi19
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        var v22 int8 = -13
        if (607 < 18446744073709551615) {
            v14 = v14
            v5 -= ((v5 - ((float64(1.5) + 1e10) / (-2.25 + v5))) * (((v5 + (3.14159 * float64(1.5))) + (((v5 + (v5 / v5)) * 0.1) - ((v5 / (3.14159 + (v5 + (v5 - ((1e10 / (0.1 + float64(0.1))) - (v5 / (((v5 * v5) - (1e10 - (float64(1e10) - ((0.0 * ((v5 + 1.5) - (0.0 - ((((v5 * 1e10) / v5) + 0.1) / v5)))) / v5)))) / (1e10 + v5)))))))) * 0.0))) * v5))
            _ = len(v3)
        } else {
            v5 -= v5
        }
        v18 = (v18 + v14)
        _ = v22
    } else {
        var v23 float32 = ((1.5 + ((((3.14159 / 1.0) - 3.14159) - -2.25) - 1.5)) + float32((v5 / v5)))
        _ = v23
    }
    v24 := GenS4{F0: 42}
    v25 := GenS4{F0: 2}
    v18 = ((v18 + "") + v14)
    fmt.Println(v18, v18, v5)
    var v26 int16 = 892
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
