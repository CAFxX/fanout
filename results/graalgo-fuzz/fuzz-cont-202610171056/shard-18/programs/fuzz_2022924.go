package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int32, p1 uint8, p2 int64) int64 {
    v7 := GenS1{F0: p2, F1: 108, F2: 537}
    v8 := GenS2{F0: 194}
    v9 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return int64(p0)
}

func main() {
    var v1 float64 = (3.14159 + -2.25)
    var v2 int16 = 43
    v3 := make(map[int]uint32)
    var v4 int = -381
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    for fi10 := 0; fi10 < 2; fi10++ {
        v5[v4] = true
        _ = fi10
    }
    v4 = -2147483648
    var v11 float32 = (-2.25 - (1e10 - float32((3.14159 + v1))))
    var v12 bool = (true || ((v11 + v11) > (v11 / v11)))
    var v13 uint16 = 407
    var v14 int32 = int32(v2)
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint32 = 2147483647
        _ = fi16
        _ = v17
    }
    v5[9223372036854775807] = (v15 >= -1)
    if (v14 <= v14) {
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := []int{-1, (-(v4)), v4}
            var v20 uint32 = 42
            _ = len(v19)
            _ = fi18
            _ = v19
            _ = v20
        }
        v2 -= v2
        v13 += (v13 * 22)
    } else {
        if v12 {
            v21 := v3[-1]
            _ = v21
        } else {
            var v22 int = 0
            v2 -= v15
            if (v12 && (0 != 2)) {
                v12 = (1 != 434)
                var v23 float32 = 3.14159
                _ = v23
            }
            _ = v22
        }
    }
    v24 := GenS3{F0: v13, F1: 0}
    v25 := []int{(v4 + v4)}
    var v26 bool = v12
    var v27 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
