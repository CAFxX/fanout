package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint64, p1 int64) uint64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = fi6
        v7 = fi6
        p1 += (p1 >> ((255) & 63))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                p1 = p1
                var v10 uint16 = 1
                if (177 == 2) {
                    v11 := []int{(^(fi9))}
                    _ = len(v11)
                    fmt.Println(fi9)
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            p1 = (p1 ^ p1)
            p1 = p1
            if true {
                if true {
                    p0 -= 3
                    if (!(747 <= 1000)) {
                        v13 := GenS1{F0: 255, F1: 2}
                        v14 := []int{fi8, fi8}
                        _ = v13
                        _ = v14
                    } else {
                        p1 = 42
                    }
                }
            }
            var v15 int = 9223372036854775807
            _ = fi12
            _ = v15
        }
        _ = fi8
    }
    p0 += 18446744073709551615
    v16 := GenS1{F0: 2, F1: p0}
    _ = p0
    _ = p1
    _ = v16
    return p0
}
func genFn1(p0 uint8) int32 {
    if true {
        fmt.Println(p0)
    } else {
        p0 += p0
    }
    p0 -= (p0 + 42)
    _ = p0
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{}
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 uint8 = 52
    v17 := GenS2{F0: 2147483648}
    switch 4294967295 {
    case 32767:
        if true {
            var v18 int64 = (v1 ^ (^(v1)))
            var v19 int16 = 1000
            for fi20 := 0; fi20 < 16; fi20++ {
                v21 := GenS1{F0: 127, F1: genFn0(223, (v18 / ((-9223372036854775808) | 1)))}
                for fi22 := 0; fi22 < 10; fi22++ {
                    _ = len(v2)
                    var v23 float64 = float64(1.5)
                    if false {
                        fmt.Println(v18, v5)
                        var v24 int32 = 0
                        fmt.Println(v19, v1)
                        _ = v24
                    }
                    var v25 float32 = -2.25
                    _ = fi22
                    _ = v23
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
            _ = v18
            _ = v19
        }
    case 9223372036854775807:
        _ = len(v3)
    case 128:
        v26 := GenS2{F0: 0}
        _ = v26
    default:
        var v27 uint32 = 65535
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 bool = true
        var v30 uint16 = 42
        _ = fi28
        _ = v29
        _ = v30
    }
    delete(v3, 42)
    var v31 int32 = 1
    var v32 float32 = (3.14159 / -2.25)
    v33 := GenS3{F0: v1, F1: v1}
    v31 += (v31 | v31)
    for fi34 := 0; fi34 < 12; fi34++ {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        _ = fi34
    }
    v35 := v3[121]
    var v36 float64 = 3.14159
    var v37 int32 = v31
    v1 = 9223372036854775807
    v38 := GenS1{F0: v5, F1: 688}
    fmt.Println(v5, v35)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
