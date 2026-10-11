package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32, p1 uint8) uint16 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{(-(fi7))}
        v9 := GenS1{F0: 2}
        var v10 float64 = 1.5
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    p1 = (p1 ^ p1)
    switch "" {
    case "":
        p1 += (p1 ^ p1)
    case "a":
        v11 := GenS2{F0: 18446744073709551615, F1: 225}
        _ = v11
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint16 = 652
    v2 := make(map[int]string)
    var v3 float32 = 1.5
    var v4 int = -9223372036854775808
    var v5 uint = 65535
    var v6 bool = true
    v12 := GenS3{F0: -2147483648, F1: 1}
    delete(v2, 891)
    var v13 uint64 = 1
    v1 = (^(v1))
    for fi14 := 0; fi14 < 9; fi14++ {
        switch (fi14 | (fi14 >> ((v5) & 63))) {
        case -1:
            if v6 {
                var v15 bool = (1 >= 42)
                _ = v15
            }
        case -2147483648:
            v16 := []int{(fi14 << ((v5) & 63)), v4}
            _ = v16
        case -793:
            var v17 uint8 = 103
            _ = v17
        }
        var v18 int64 = -1
        _ = fi14
        _ = v18
    }
    v13 = (v13 - (v13 | (v13 * 647)))
    var v19 bool = (!(v6 || v6))
    var v20 uint = v5
    for fi21 := 0; fi21 < 1; fi21++ {
        if v19 {
            for fi22 := 0; fi22 < 8; fi22++ {
                _ = len(v2)
                var v23 int16 = int16(v1)
                _ = fi22
                _ = v23
            }
            var v24 int32 = 493
            v25 := []int{}
            _ = v24
            _ = v25
        } else {
            var v26 uint8 = 213
            switch v4 {
            case -835:
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := GenS4{F0: -1}
                    var v29 int16 = int16(v1)
                    var v30 float64 = (0.1 / (1e10 + -2.25))
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            case 32768:
                v31 := []int{fi21, (^(v4)), -1}
                _ = v31
            case 0:
                var v32 bool = true
                _ = v32
            default:
                var v33 int = v4
                _ = v33
            }
            _ = v26
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
