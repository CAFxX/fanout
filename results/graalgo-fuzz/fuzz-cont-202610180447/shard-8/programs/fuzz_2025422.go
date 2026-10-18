package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0() uint16 {
    v4 := GenS1{F0: false, F1: 1}
    var v5 int = 9223372036854775807
    v6 := []int{(v5 + v5), (^(v5))}
    fmt.Println(v5)
    v5 = v5
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16) uint64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint16 = p0
        var v9 int = (^(fi7))
        var v10 int64 = -1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch -706 {
    case 763:
        var v11 uint8 = 2
        _ = v11
    }
    var v12 int32 = 42
    _ = p0
    _ = v12
    return 96
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = (float64(1.5) - 0.0)
    var v3 uint = 18446744073709551615
    v3 = v3
    for fi13 := 0; fi13 < 20; fi13++ {
        if (0 < 1) {
            for fi14 := 0; fi14 < 16; fi14++ {
                v3 = (v3 & 255)
                v15 := []int{fi14, (fi13 - -9223372036854775808), 128}
                var v16 bool = ((!(!(v3 > (v3 & (v3 / (((^(v3))) | 1)))))) && false)
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS2{F0: 42, F1: (true && (0 >= -2147483648))}
            var v18 float64 = 1e10
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    var v19 uint8 = 42
    var v20 int32 = -1
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v20, v3)
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS2{F0: 1, F1: (!((fi22 % ((fi21) | 1)) <= (^(fi21))))}
            _ = len(v1)
            _ = fi22
            _ = v23
        }
        var v24 bool = (1 >= 4294967295)
        _ = fi21
        _ = v24
    }
    fmt.Println(v3)
    v25 := GenS2{F0: 979, F1: true}
    v3 -= 255
    v2 = -2.25
    var v26 int16 = 42
    v2 -= float64(float32((v2 / 0.1)))
    var v27 bool = true
    v28 := []int{2147483647, -662}
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS1{F0: (65535 == 2), F1: int16(v19)}
        _ = fi29
        _ = v30
    }
    fmt.Println(v26)
    fmt.Println(v26)
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        if (3.14159 >= 0.1) {
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 uint16 = 2
                v34 := GenS2{F0: 4294967295, F1: ((fi32 >= -9223372036854775808) || false)}
                _ = fi32
                _ = v33
                _ = v34
            }
            v28 = append(v28, (fi31 * (-(fi31))))
        }
        v35 := GenS3{F0: v19, F1: (-(v26)), F2: 225}
        _ = fi31
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
