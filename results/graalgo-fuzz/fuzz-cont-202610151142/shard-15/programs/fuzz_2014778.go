package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 int }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64) int {
    v4 := []int{}
    var v5 float32 = 1e10
    _ = p0
    _ = v4
    _ = v5
    return -109
}
func genFn1() uint {
    var v6 string = "foo bar"
    var v7 uint32 = 0
    var v8 uint32 = (^(v7))
    v9 := []int{-703, 2147483647, 9223372036854775807}
    switch -2147483648 {
    case -249:
        v10 := GenS1{F0: -1, F1: 467, F2: 633}
        _ = v10
    case 9223372036854775807:
        fmt.Println(v7, v8)
    case -2147483648:
        v8 += v7
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 int16 = 401
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch fi11 {
                case -109:
                    fmt.Println(fi11)
                case 32767:
                    fmt.Println(v2, v2, fi12)
                case 2147483647:
                    v13 := GenS1{F0: (-(v3)), F1: (^(fi11)), F2: 1}
                    _ = v13
                }
                v14 := []int{fi12, 56}
                _ = len(v1)
                _ = len(v14)
                _ = fi12
                _ = v14
            }
            v2 += 128
            if true {
                var v15 string = ""
                _ = v15
            } else {
                v16 := GenS2{F0: 0, F1: 9223372036854775807, F2: 1}
                v17 := GenS2{F0: 2, F1: 1, F2: (-(fi11))}
                _ = v16
                _ = v17
            }
        }
        v3 = 1
        var v18 int16 = (v3 >> ((2) & 63))
        _ = fi11
        _ = v18
    }
    v2 -= v2
    if false {
        fmt.Println(v2)
        var v19 int16 = (v3 + (v3 * v3))
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 uint16 = 139
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    _ = len(v1)
    v3 += v3
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS2{F0: 142, F1: 211, F2: (fi22 - fi23)}
            if false {
                var v25 float64 = 1.5
                _ = len(v1)
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        var v26 int64 = -9223372036854775808
        var v27 int16 = v3
        _ = fi22
        _ = v26
        _ = v27
    }
    v3 += (v3 - v3)
    v28 := GenS3{F0: -1}
    fmt.Println(v2, v3)
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v2, v3, fi29)
        for fi30 := 0; fi30 < 13; fi30++ {
            switch fi29 {
            case -966:
                if true {
                    v2 -= 42
                    fmt.Println(v2)
                }
            default:
                v31 := GenS3{F0: 29}
                _ = v31
            }
            _ = fi30
        }
        _ = fi29
    }
    if false {
        v3 -= -32768
    }
    v32 := []int{127}
    if len(v32) > 0 {
        v32[0] = -532
    }
    v3 = (v3 * -880)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
}
