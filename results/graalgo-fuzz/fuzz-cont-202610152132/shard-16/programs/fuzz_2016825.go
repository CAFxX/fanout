package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0() int {
    if true {
        var v7 uint32 = 2147483648
        v7 += v7
        _ = v7
    }
    if (!(1 != 55)) {
        var v8 uint8 = 127
        if true {
            v8 -= (v8 - v8)
        } else {
            if (true || (!(!(32767 != 1000)))) {
                v8 += (^(v8))
                v9 := GenS1{F0: 4294967295}
                _ = v9
            } else {
                var v10 string = "a"
                v11 := GenS2{F0: (v8 | 128), F1: 0, F2: 845}
                _ = v10
                _ = v11
            }
            if (!(!((1.5 < 0.0) || true))) {
                var v12 uint8 = (^(v8))
                v12 = v12
                _ = v12
            }
        }
        fmt.Println(v8)
        _ = v8
    }
    v13 := []int{1, 6}
    _ = v13
    return 42
}
func genFn1(p0 string) uint8 {
    var v14 uint32 = 634
    switch genFn0() {
    case -636:
        var v15 uint64 = 582
        _ = v15
    }
    _ = p0
    _ = v14
    return 128
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 float64 = -2.25
    v5 := make(map[string]int)
    var v6 uint = 725
    v6 = 658
    for fi16 := 0; fi16 < 4; fi16++ {
        v6 -= 18446744073709551615
        _ = fi16
    }
    var v17 float32 = 0.1
    v6 += 630
    var v18 int16 = -32768
    var v19 uint32 = 4294967295
    var v20 int32 = 42
    v4 += (float64(-2.25) / v4)
    fmt.Println(v1, v1)
    var v21 float64 = v4
    for fi22 := 0; fi22 < 9; fi22++ {
        v19 = v19
        v19 = v19
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 float64 = float64(v17)
            v25 := v2[(v19 ^ 25)]
            v26 := []int{fi22}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
    }
    var v27 uint = v6
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int8 = -128
        _ = fi28
        _ = v29
    }
    fmt.Println(v20, v6)
    v30 := GenS2{F0: 119, F1: 1, F2: 69}
    var v31 uint8 = 2
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS3{F0: 1, F1: 800}
        var v34 int = (^(fi32))
        v4 = 0.1
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 bool = ((v17 > (1e10 / 1e10)) && (!(1.5 <= (0.1 / -2.25))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
