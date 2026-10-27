package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0(p0 uint16, p1 uint32, p2 int32) uint {
    switch 2147483647 {
    case 550:
        p0 = p0
    case 32767:
        for fi7 := 0; fi7 < 19; fi7++ {
            p2 += int32(p1)
            _ = fi7
        }
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 2}
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = (p0 ^ p0)
    var v11 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]string)
    var v3 int8 = 42
    var v4 float32 = -2.25
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    v4 = (v4 / v4)
    _ = len(v2)
    v1 = (v1 >> ((genFn0((v1 | (v1 | 1)), 2, -819)) & 63))
    if ((((float64(((v4 - v4) + 0.1)) + 0.1) + (0.0 - 1.5)) + 1.5) != (float64(v4) / 1e10)) {
        v12 := GenS2{F0: 32767}
        fmt.Println(v3, v3)
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 127}
        if ((^(v1)) != v1) {
            v3 -= v3
            var v15 string = ""
            _ = v15
        } else {
            delete(v5, "")
            if (1 <= 65535) {
                var v16 int64 = 9223372036854775807
                _ = v16
            } else {
                if (!(false || false)) {
                    fmt.Println(fi13, v4)
                    var v17 int32 = 42
                    v18 := []int{(-(fi13)), fi13, fi13}
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint8 = 1
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 uint64 = 796
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 uint = 52
            var v24 uint32 = 2147483647
            _ = fi22
            _ = v23
            _ = v24
        }
        v4 += 0.0
        _ = fi20
        _ = v21
    }
    _ = len(v5)
    v25 := v2[2]
    var v26 int16 = 0
    var v27 int32 = 1
    v28 := GenS3{F0: 127, F1: v1, F2: (954 < 583)}
    v26 -= (v26 + -498)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
