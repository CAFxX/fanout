package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0(p0 uint64, p1 int32, p2 int16) int {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: -773, F1: -1}
        _ = fi6
        _ = v7
    }
    var v8 int32 = (p1 << ((187) & 63))
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            v10 := GenS1{F0: (p1 >> ((42) & 63)), F1: 9223372036854775807}
            if true {
                var v11 int16 = p2
                v8 -= p1
                if (1.5 <= float32(float64(0.0))) {
                    v12 := []int{}
                    _ = len(v12)
                    _ = v12
                }
                _ = v11
            }
            var v13 int8 = 42
            _ = v10
            _ = v13
        } else {
            switch fi9 {
            case 4294967295:
                v14 := []int{(fi9 >> ((2) & 63)), fi9, (fi9 * (-(fi9)))}
                _ = v14
            case -1:
                v15 := []int{(fi9 | 32768)}
                _ = v15
            case 0:
                if false {
                    fmt.Println(fi9)
                    var v16 uint = 237
                    _ = v16
                }
            }
            var v17 float64 = ((float64(((3.14159 + 0.1) * (3.14159 * 0.0))) / 0.1) / -2.25)
            _ = v17
        }
        _ = fi9
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        p0 = 2
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -920
}

func main() {
    var v1 uint = 1
    var v2 string = "a"
    var v3 int16 = -588
    v4 := []uint32{65535, 2}
    v5 := make(map[uint32]int)
    var v19 float64 = 0.1
    switch genFn0(2, -2147483648, (v3 ^ v3)) {
    case -9223372036854775808:
        v3 -= (v3 ^ -95)
    }
    v19 = float64(-2.25)
    fmt.Println(v2)
    var v20 uint64 = 42
    var v21 float64 = v19
    var v22 int = 32768
    v21 = -2.25
    if (!(genFn0(v20, int32(v3), (v3 + (v3 | v3))) <= (^(v22)))) {
        v20 += v20
    }
    switch v22 {
    case 297:
        v20 = v20
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 float32 = 1e10
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    fmt.Println(v21)
    v21 += v21
    v26 := GenS1{F0: 2147483647, F1: 1}
    var v27 string = "\n\t"
    var v28 int64 = int64(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
