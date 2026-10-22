package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16; F2 int }
func genFn0(p0 string, p1 int32) int8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 536
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    var v8 float64 = float64((0.1 + -2.25))
    var v9 uint = 0
    var v10 int = -732
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int, p1 uint64) bool {
    switch (-(p0)) {
    case 1:
        v11 := []int{(-(p0)), (-(p0)), (p0 >> ((0) & 63))}
        _ = v11
    case 128:
        p0 = p0
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{(fi12 >> ((18446744073709551615) & 63))}
            _ = fi13
            _ = v14
        }
        p1 += (^(p1))
        _ = fi12
    }
    fmt.Println(p0, p1)
    v15 := GenS1{F0: 33}
    switch (p0 + (^(p0))) {
    case 42:
        v16 := []int{0}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v15
    return (0.0 >= -2.25)
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{1, 2, 42, 2147483648}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    v4[("\n\t" + "foo bar")] = 2147483648
    for fi17 := 0; fi17 < 16; fi17++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS2{F0: 31, F1: 9223372036854775807}
            v20 := GenS1{F0: 1}
            var v21 int8 = 100
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        fmt.Println(v1)
        _ = fi17
    }
    v3[("\n\t" + "")] = 1
    var v22 float64 = (-2.25 + float64(1.5))
    _ = len(v2)
    v1 = (v1 / ((v1) | 1))
    v23 := GenS3{F0: 881, F1: 412, F2: 128}
    v1 += (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
