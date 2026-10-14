package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 string) uint64 {
    v7 := GenS1{F0: 65535}
    switch 32768 {
    case 32768:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 2147483648}
            v10 := []int{}
            fmt.Println(p1)
            _ = fi8
            _ = v10
            _ = v9
        }
    case 932:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint, p2 int8) int {
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 1}
        v13 := GenS1{F0: 387}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{128, 0, 42}
    var v15 uint8 = 2
    if (2147483647 > 4294967295) {
        p2 += p2
        v16 := GenS1{F0: 762}
        _ = v16
    } else {
        v17 := GenS2{F0: 0, F1: 1}
        v18 := GenS3{F0: 2147483648, F1: 2147483647, F2: (-2.25 == (-2.25 - -2.25))}
        v14 = append(v14, 42)
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 202
    var v2 uint = 65535
    v3 := make(map[int]bool)
    var v4 uint64 = 962
    var v5 bool = true
    v6 := make(map[int]bool)
    v4 -= uint64(v1)
    switch 42 {
    case 42:
        if (4294967295 > 42) {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint32 = 65535
                if v5 {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int64 = 791
                        var v23 int = (-(fi21))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 int8 = -46
                        v26 := GenS4{F0: 42}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                } else {
                    v27 := []int{(-(fi19))}
                    var v28 int8 = -1
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 4; fi29++ {
                    v30 := GenS2{F0: 196, F1: -1}
                    _ = len(v3)
                    v31 := GenS3{F0: 2147483647, F1: v20, F2: ((-2.25 - -2.25) <= 0.0)}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi19
                _ = v20
            }
        }
    case -9223372036854775808:
        v4 -= 469
    default:
        var v32 uint64 = 392
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        v34 := GenS3{F0: 65535, F1: 493, F2: v5}
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: 0, F1: 42}
    v4 -= uint64(v2)
    var v36 uint64 = genFn0((1e10 > (0.1 - ((((3.14159 * 0.1) * 0.0) * 1e10) + float32(0.1)))), "\n\t")
    if v5 {
        var v37 float32 = float32((0.1 / (-2.25 - (float64(1.5) - 3.14159))))
        _ = v37
    }
    v38 := GenS4{F0: 2147483647}
    fmt.Println(v36, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
