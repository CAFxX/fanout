package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 string; F1 int8; F2 int }
func genFn0(p0 uint32, p1 int8, p2 uint32) uint8 {
    var v7 int32 = 1
    fmt.Println(p1, p2)
    p0 += 815
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{0}
            v7 = -20
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: 0, F1: v7}
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            v7 += v7
            v13 := []int{fi12}
            if true {
                v7 = (v7 + (v7 | 451))
                var v14 int16 = 1
                v15 := GenS2{F0: "foo bar", F1: (p1 + p1), F2: fi12}
                _ = v14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        var v16 uint = 42
        v17 := GenS2{F0: "foo bar", F1: -100, F2: -9223372036854775808}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 uint16 = 369
    var v2 uint32 = 1
    var v3 int32 = 618
    v4 := []uint32{(v2 >> ((uint(v2)) & 63)), (^(v2))}
    v5 := make(map[string]int)
    v6 := make(map[int]int)
    if false {
        var v18 uint16 = v1
        var v19 int64 = 1
        _ = v18
        _ = v19
    }
    var v20 int8 = 127
    v20 = v20
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint64 = 351
        v23 := v5["a"]
        _ = fi21
        _ = v22
        _ = v23
    }
    v3 -= v3
    for fi24 := 0; fi24 < 3; fi24++ {
        switch -848 {
        case -1:
            v25 := GenS1{F0: 387, F1: -1}
            _ = v25
        }
        fmt.Println(fi24, v2, v2)
        _ = fi24
    }
    v26 := GenS2{F0: "\n\t", F1: 42, F2: -1}
    v1 -= (^(v1))
    if true {
        for fi27 := 0; fi27 < 8; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                switch "" {
                case "":
                    v1 -= v1
                case "hello":
                    v20 -= int8(v1)
                default:
                    var v29 uint64 = 1
                    _ = v29
                }
                var v30 uint64 = 498
                _ = fi28
                _ = v30
            }
            fmt.Println(fi27, v2)
            v31 := GenS1{F0: -391, F1: (^(v3))}
            _ = fi27
            _ = v31
        }
    }
    if true {
        v32 := GenS2{F0: "\n\t", F1: (v20 * (v20 << ((uint(v2)) & 63))), F2: 4294967295}
        _ = v32
    } else {
        v20 += v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
