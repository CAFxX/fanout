package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 int) string {
    var v7 uint8 = 57
    var v8 int16 = 32767
    var v9 uint32 = 2
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint, p1 uint32, p2 bool) int16 {
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int8 = -128
        if true {
            for fi13 := 0; fi13 < 9; fi13++ {
                v12 = v12
                switch fi11 {
                case 1:
                    v14 := GenS1{F0: 1, F1: 67}
                    _ = v14
                case 24:
                    var v15 string = "foo bar"
                    _ = v15
                case 0:
                    p2 = (!(!(146 != -937)))
                }
                if p2 {
                    switch (-(fi11)) {
                    case 42:
                        v16 := []int{fi11, 128, fi11, fi11}
                        _ = v16
                    case -2147483648:
                        var v17 bool = (0.0 >= 1.5)
                        var v18 uint16 = 1
                        v19 := []int{fi11, fi13, (fi13 / (((fi11 % ((fi13) | 1))) | 1)), fi11}
                        _ = v17
                        _ = v18
                        _ = v19
                    case 1:
                        p1 = p1
                    default:
                        var v20 int16 = -794
                        _ = v20
                    }
                    p0 = (p0 * p0)
                    v12 = v12
                    switch fi13 {
                    case 4294967295:
                        var v21 int8 = (^(v12))
                        _ = v21
                    case 32767:
                        v22 := []int{}
                        _ = v22
                    default:
                        if (p0 <= (^(p0))) {
                            v23 := GenS2{F0: -810, F1: 1, F2: fi13}
                            v12 = v12
                            _ = v23
                        } else {
                            fmt.Println(v12)
                        }
                    }
                } else {
                    var v24 int16 = -7
                    _ = v24
                }
                _ = fi13
            }
        } else {
            var v25 int32 = -30
            _ = v25
        }
        _ = fi11
        _ = v12
    }
    var v26 int = -9223372036854775808
    p0 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 974
}

func main() {
    var v1 int16 = 1000
    v2 := []int{0, -497, 2147483647, 32767}
    var v3 uint16 = 2
    var v4 int8 = 0
    v5 := []uint32{1, 2, 2, 4294967295}
    v6 := make(map[uint32]bool)
    for fi27 := 0; fi27 < 20; fi27++ {
        if true {
            fmt.Println(fi27)
        } else {
            fmt.Println(v1, v3)
        }
        _ = fi27
    }
    fmt.Println(v1)
    delete(v6, 2)
    if (v3 > v3) {
        fmt.Println(v3, v3)
        var v28 string = "\n\t"
        _ = v28
    }
    delete(v6, 42)
    v1 = (v1 ^ (v1 & int16(v4)))
    var v29 int = 446
    var v30 uint = 0
    var v31 uint = (^(v30))
    var v32 float32 = 3.14159
    _ = len(v6)
    delete(v6, 2147483647)
    fmt.Println(v3, v29)
    v33 := GenS3{F0: 9223372036854775807}
    var v34 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
