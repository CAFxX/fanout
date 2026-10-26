package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 int16 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 uint32) uint8 {
    p0 += p0
    p0 = p0
    var v7 bool = (0 < 1)
    v8 := GenS1{F0: (2147483647 < -2147483648)}
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int64, p1 uint8) bool {
    var v9 float64 = (float64(3.14159) - ((3.14159 - -2.25) - -2.25))
    if (v9 <= 3.14159) {
        v9 = float64(3.14159)
    }
    v10 := []int{321, 128}
    switch 9223372036854775807 {
    case -372:
        fmt.Println(p0)
    case 127:
        fmt.Println(v9, v9, p1)
    case 9223372036854775807:
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (!true)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (2 < 128)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{}
    v11 := GenS1{F0: true}
    if len(v6) > 0 {
        v6[0] = 342
    }
    var v12 uint32 = 229
    for fi13 := 0; fi13 < 19; fi13++ {
        if (42 > 640) {
            switch (-(fi13)) {
            case -257:
                var v14 float64 = 1.5
                _ = v14
            }
            switch 9223372036854775807 {
            case -9223372036854775808:
                _ = len(v3)
            }
        }
        if ((!(fi13 != -867)) || (577 < -1)) {
            var v15 float32 = float32(0.1)
            _ = v15
        }
        fmt.Println(fi13, v2, fi13)
        _ = fi13
    }
    v12 += v12
    v16 := GenS1{F0: v2}
    var v17 int64 = -243
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS2{F0: genFn0(v12), F1: 0, F2: 1000}
        if true {
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := GenS3{F0: 65535, F1: v2}
                _ = fi20
                _ = v21
            }
        } else {
            switch fi18 {
            case -391:
                var v22 int64 = v17
                _ = v22
            case 563:
                v2 = ((3.14159 - 0.1) < 3.14159)
            case 32767:
                var v23 int32 = 42
                _ = v23
            default:
                v17 += 1
            }
            var v24 float32 = 1e10
            _ = v24
        }
        var v25 int = fi18
        _ = fi18
        _ = v19
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v17, v12, v12)
        v6 = append(v6, 164)
        _ = fi26
    }
    _ = len(v6)
    var v27 int64 = (^(v17))
    v12 -= v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
