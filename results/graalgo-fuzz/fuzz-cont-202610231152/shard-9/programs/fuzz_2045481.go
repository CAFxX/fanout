package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 int16, p1 bool, p2 uint64) int8 {
    v4 := GenS1{F0: uint16(p2), F1: 32767}
    fmt.Println(p2, p0, p1)
    var v5 bool = (p1 && (2 > 812))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -74
}
func genFn1(p0 int8, p1 int8) uint {
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 607, F1: 65535}
            _ = fi7
            _ = v8
        }
        var v9 uint8 = 2
        if (4294967295 <= 2147483647) {
            var v10 int = fi6
            for fi11 := 0; fi11 < 1; fi11++ {
                switch "" {
                case "hello":
                    p1 += (^(p0))
                case "\n\t":
                    v12 := []int{fi6, (v10 >> ((122) & 63)), (fi6 | 9223372036854775807)}
                    _ = v12
                }
                var v13 uint8 = v9
                _ = fi11
                _ = v13
            }
            _ = v10
        } else {
            if (3.14159 >= 0.0) {
                p1 = (p0 ^ 1)
                var v14 uint8 = v9
                _ = v14
            }
        }
        _ = fi6
        _ = v9
    }
    v15 := GenS1{F0: 1, F1: 2}
    p1 = p1
    for fi16 := 0; fi16 < 15; fi16++ {
        if (!((fi16 >> ((18446744073709551615) & 63)) < -97)) {
            p0 = p0
            p0 -= (p0 * genFn0(1000, ((^(fi16)) >= fi16), 4294967295))
            p1 = p0
        } else {
            var v17 uint32 = 1
            _ = v17
        }
        if true {
            p1 -= -1
            for fi18 := 0; fi18 < 18; fi18++ {
                if true {
                    p1 = 1
                } else {
                    var v19 float32 = -2.25
                    _ = v19
                }
                var v20 int16 = 1000
                fmt.Println(fi18, fi18, p1)
                _ = fi18
                _ = v20
            }
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v15
    return 2
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int32 = -1
    if true {
        fmt.Println(v3)
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS1{F0: 1, F1: 32767}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint16 = 32767
            _ = fi23
            _ = v24
        }
        v3 = 949
    }
    v3 += v3
    fmt.Println(v3)
    var v25 string = "\n\t"
    v3 -= -2147483648
    v26 := GenS1{F0: 0, F1: 0}
    var v27 uint64 = 18446744073709551615
    if ((^(v27)) <= (v27 - v27)) {
        v28 := GenS1{F0: 0, F1: 175}
        fmt.Println(v27)
        fmt.Println(v25, v25, v27)
        _ = v28
    }
    v1[-1] = "\n\t"
    fmt.Println(v27, v3, v25)
    for fi29 := 0; fi29 < 3; fi29++ {
        v2[65535] = (-(fi29))
        v27 = (v27 | (v27 & (^(v27))))
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := []int{(fi29 >> ((genFn1(genFn0(32767, (!false), (v27 | 8)), 8)) & 63)), fi30}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v27 = v27
    var v32 float64 = 1.5
    v33 := GenS2{F0: -420, F1: true, F2: 1}
    v34 := GenS3{F0: v3, F1: 65535, F2: 46}
    fmt.Println(v32)
    switch v25 {
    case "a":
        switch -1 {
        case 959:
            v32 += v32
        case 328:
            v1[-1] = v25
        case 318:
            fmt.Println(v3)
        }
    default:
        v35 := GenS4{F0: -9223372036854775808, F1: (v27 / (((v27 | (^(v27)))) | 1))}
        _ = v35
    }
    var v36 uint16 = 330
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
}
