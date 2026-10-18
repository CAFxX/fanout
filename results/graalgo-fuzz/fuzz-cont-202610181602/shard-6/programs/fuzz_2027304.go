package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 uint, p1 bool, p2 uint64) int32 {
    if p1 {
        if false {
            for fi4 := 0; fi4 < 13; fi4++ {
                v5 := GenS1{F0: 1000, F1: 0, F2: 42}
                for fi6 := 0; fi6 < 18; fi6++ {
                    v7 := []int{(fi4 ^ (-(fi4))), fi6, (fi4 & fi4)}
                    v8 := GenS1{F0: 0, F1: 871, F2: 112}
                    v9 := GenS2{F0: 1, F1: (^(p0)), F2: -522}
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
                p2 = 4294967295
                _ = fi4
                _ = v5
            }
        } else {
            var v10 uint8 = 1
            fmt.Println(v10, p1)
            _ = v10
        }
    }
    v11 := []int{-470, -2147483648, -757}
    var v12 string = "a"
    v11 = append(v11, 2147483647)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -1
}
func genFn1(p0 int16, p1 int16) int {
    v13 := GenS2{F0: 128, F1: 391, F2: 56}
    p0 = p0
    p1 = p1
    if false {
        fmt.Println(p0)
        fmt.Println(p0, p1)
        p0 = (p1 >> ((624) & 63))
    } else {
        switch 2147483647 {
        case 32767:
            var v14 float64 = 1e10
            _ = v14
        case -904:
            var v15 int8 = 100
            _ = v15
        default:
            var v16 uint64 = 18446744073709551615
            _ = v16
        }
        fmt.Println(p0, p0)
        switch -2147483648 {
        case 128:
            p1 = (p0 | p0)
        }
    }
    var v17 string = "\n\t"
    _ = p0
    _ = p1
    _ = v13
    _ = v17
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 int8 = -128
    var v3 int16 = 42
    v3 += v3
    v18 := GenS2{F0: 231, F1: 0, F2: 32767}
    v2 = v2
    v19 := GenS2{F0: 127, F1: 0, F2: int(v3)}
    var v20 int64 = 1
    for fi21 := 0; fi21 < 10; fi21++ {
        fmt.Println(v1, v2, v20)
        var v22 float64 = 3.14159
        _ = fi21
        _ = v22
    }
    v2 += v2
    fmt.Println(v20, v2, v1)
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 uint8 = 0
        var v25 bool = false
        var v26 int = fi23
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v20 -= 1
        var v28 int64 = v20
        for fi29 := 0; fi29 < 5; fi29++ {
            v20 -= (v28 + 0)
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    v30 := []int{9223372036854775807, 127, genFn1((v3 / ((1) | 1)), v3), -9223372036854775808}
    var v31 float32 = 0.0
    v20 = v20
    v32 := GenS2{F0: 212, F1: 255, F2: 127}
    v20 += (v20 << ((2) & 63))
    var v33 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
