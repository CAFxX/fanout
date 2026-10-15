package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0(p0 uint, p1 uint16) uint16 {
    if (!(136 <= p1)) {
        v7 := []int{-1, -1}
        p1 = (p1 & (p1 ^ (p1 & p1)))
        var v8 uint8 = 128
        _ = v7
        _ = v8
    }
    var v9 uint8 = 127
    var v10 int16 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p1
}
func genFn1(p0 uint8, p1 uint64, p2 int16) uint {
    var v11 string = (("" + ((("hello" + "") + "a") + "foo bar")) + "")
    if (!false) {
        fmt.Println(p0, v11, p2)
        if (((p0 + (p0 >> ((uint(p0)) & 63))) != (p0 << ((4294967295) & 63))) || true) {
            p1 = p1
            v12 := []int{-346}
            _ = v12
        } else {
            var v13 uint32 = 0
            var v14 uint64 = p1
            _ = v13
            _ = v14
        }
        var v15 uint32 = 1
        _ = v15
    } else {
        var v16 float64 = ((1.5 - 0.0) - float64(0.1))
        if false {
            v17 := GenS1{F0: p1, F1: 1}
            var v18 uint32 = 101
            _ = v17
            _ = v18
        } else {
            var v19 uint16 = uint16(p1)
            _ = v19
        }
        var v20 float32 = 1.5
        _ = v16
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 852
}

func main() {
    v1 := []string{}
    var v2 int = 4294967295
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    v5 := []int{int(v4), (-(v2)), (^(v2)), 425}
    var v6 bool = true
    delete(v3, 4294967295)
    v2 = v2
    fmt.Println(v4, v2)
    v21 := GenS2{F0: genFn0(2, 202), F1: 0}
    v22 := GenS1{F0: v4, F1: 2}
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint = 0
        var v25 uint = (^(v24))
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int32 = -47
        _ = len(v3)
        _ = fi26
        _ = v27
    }
    v2 = v2
    v2 -= 1
    var v28 int32 = -1
    v29 := GenS2{F0: 881, F1: 127}
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v28, v6, v6)
        _ = fi30
    }
    if v6 {
        v1 = append(v1, "foo bar")
        switch (v2 - (v2 + v2)) {
        case 0:
            var v31 bool = (0.1 < 3.14159)
            _ = v31
        case 127:
            _ = len(v3)
        case 128:
            v32 := GenS3{F0: 788, F1: 1, F2: 0}
            _ = v32
        }
        _ = len(v3)
    }
    if (1 == 42) {
        v2 += v2
        fmt.Println(v4, v2, v4)
    }
    if true {
        v4 += (^(v4))
        v33 := GenS3{F0: 2, F1: v2, F2: 2147483647}
        v34 := GenS2{F0: 0, F1: 1}
        _ = v33
        _ = v34
    }
    v5 = append(v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
