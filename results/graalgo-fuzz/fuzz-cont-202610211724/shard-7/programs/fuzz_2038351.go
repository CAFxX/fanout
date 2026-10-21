package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
func genFn0(p0 bool, p1 uint32, p2 int32) uint8 {
    var v4 uint8 = 0
    fmt.Println(p2, v4, p0)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, 480, fi5, fi5}
        _ = fi5
        _ = v6
    }
    v7 := []int{1, -1, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return v4
}
func genFn1(p0 int16, p1 int8) int32 {
    p0 += (-(p0))
    for fi8 := 0; fi8 < 6; fi8++ {
        if true {
            var v9 bool = ((fi8 & fi8) != fi8)
            _ = v9
        } else {
            v10 := GenS1{F0: (-(p1)), F1: 4294967295, F2: 2147483648}
            v11 := GenS1{F0: (p1 % ((p1) | 1)), F1: 65535, F2: 334}
            switch (-(fi8)) {
            case 128:
                var v12 uint16 = 65535
                _ = v12
            default:
                if true {
                    p0 -= p0
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := []int{}
                        fmt.Println(p1)
                        var v15 bool = true
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        p1 = p1
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    v16 := GenS1{F0: (p1 / ((p1) | 1)), F1: 729, F2: 65535}
                    _ = v16
                }
            }
            _ = v10
            _ = v11
        }
        var v17 int8 = p1
        _ = fi8
        _ = v17
    }
    v18 := GenS2{F0: 4294967295}
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(p1, p0)
        _ = fi19
    }
    if (float64(float32(1e10)) >= 0.1) {
        v20 := GenS1{F0: (p1 | p1), F1: 1, F2: 2147483647}
        fmt.Println(p0, p0)
        _ = v20
    } else {
        var v21 uint = 65535
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(p1)
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := []int{-1, fi23, fi22}
                _ = fi23
                _ = v24
            }
            var v25 int32 = 2147483647
            _ = fi22
            _ = v25
        }
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v18
    return 0
}

func main() {
    v1 := []uint32{0}
    v2 := []uint32{1, 560, 196}
    var v3 int = 128
    if (!true) {
        if false {
            var v26 int32 = 637
            v27 := GenS1{F0: -128, F1: 4294967295, F2: 0}
            var v28 string = "\n\t"
            _ = v26
            _ = v27
            _ = v28
        }
        v3 -= 385
    } else {
        var v29 int16 = 1
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS3{F0: (fi30 - v3)}
        var v32 int = (v3 >> ((65535) & 63))
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v33 := GenS4{F0: v3}
    fmt.Println(v3)
    v3 = v3
    if false {
        var v34 float64 = 0.0
        _ = v34
    } else {
        v3 = v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = v3
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 uint64 = 0
    if (true || false) {
        var v36 int64 = 1
        _ = v36
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
}
