package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 string, p2 int) uint32 {
    if false {
        if false {
            var v7 uint8 = 127
            v8 := GenS1{F0: 42, F1: 42, F2: 255}
            _ = v7
            _ = v8
        } else {
            var v9 uint = 200
            p2 = (-(p2))
            v9 -= (v9 | v9)
            _ = v9
        }
        v10 := GenS2{F0: 107, F1: true}
        _ = v10
    } else {
        switch (p2 | p2) {
        case 128:
            var v11 float64 = 0.0
            _ = v11
        default:
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        v13 := []int{}
        _ = v13
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int = 487
    var v2 int16 = 831
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 int8 = 127
    var v6 float32 = (3.14159 / float32(-2.25))
    var v14 uint32 = 2
    v15 := []int{}
    for fi16 := 0; fi16 < 18; fi16++ {
        v5 -= v5
        _ = fi16
    }
    v17 := v4[(v14 | 4294967295)]
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int8 = v5
        for fi20 := 0; fi20 < 18; fi20++ {
            switch fi20 {
            case 127:
                v15 = append(v15, (-(fi18)))
            case -2147483648:
                var v21 float64 = (3.14159 - 0.1)
                _ = v21
            case 32768:
                switch fi20 {
                case 4294967295:
                    v22 := GenS3{F0: false}
                    _ = v22
                case -264:
                    var v23 int8 = 1
                    v24 := GenS4{F0: 1}
                    _ = v23
                    _ = v24
                case 9223372036854775807:
                    switch "" {
                    case "foo bar":
                        v1 -= fi18
                    case "":
                        v19 += 42
                    case "a":
                        var v25 uint16 = 2
                        _ = v25
                    }
                default:
                    fmt.Println(fi20, v6, v5)
                }
            default:
                v26 := GenS3{F0: v17}
                _ = v26
            }
            _ = fi20
        }
        v6 += (float32(3.14159) / v6)
        _ = fi18
        _ = v19
    }
    fmt.Println(v17, v2, v17)
    var v27 float32 = 0.1
    fmt.Println(v6, v27)
    v14 -= v14
    var v28 int = v1
    var v29 int64 = -537
    v30 := GenS1{F0: 107, F1: (v29 / (((-(v29))) | 1)), F2: 604}
    var v31 bool = v17
    var v32 int8 = (v5 / (((v5 << ((18446744073709551615) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
