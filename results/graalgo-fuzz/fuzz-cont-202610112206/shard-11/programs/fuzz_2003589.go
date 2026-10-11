package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 bool) int16 {
    if (0 > 42) {
        v5 := GenS1{F0: 0, F1: "", F2: 9223372036854775807}
        _ = v5
    } else {
        p0 -= (^(p1))
        v6 := []int{-2147483648, 0, 479}
        _ = v6
    }
    fmt.Println(p1)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = true
        v9 := []int{(fi7 ^ 1), (fi7 >> ((294) & 63)), (fi7 * fi7)}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 2147483647, F1: 42, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1000
}

func main() {
    var v1 int32 = 159
    v2 := make(map[int]int)
    var v3 uint64 = 312
    v4 := []int{}
    v1 = v1
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int64 = 9223372036854775807
        v12 -= v12
        var v13 float64 = 1.5
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v3, v1)
    fmt.Println(v3)
    var v14 int = 435
    var v15 float32 = (1.5 / 1e10)
    v16 := GenS3{F0: "hello", F1: -155, F2: 420}
    v14 += (v14 + v14)
    fmt.Println(v14)
    fmt.Println(v15, v3, v1)
    if true {
        for fi17 := 0; fi17 < 9; fi17++ {
            v15 = float32(3.14159)
            _ = fi17
        }
        fmt.Println(v3, v14, v14)
        var v18 int8 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        if true {
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS4{F0: "a", F1: 4294967295, F2: 94}
                _ = fi20
                _ = v21
            }
        } else {
            v1 = v1
            if true {
                v15 -= -2.25
            } else {
                for fi22 := 0; fi22 < 5; fi22++ {
                    if true {
                        var v23 uint16 = 798
                        _ = v23
                    }
                    _ = fi22
                }
            }
            var v24 uint16 = 2
            _ = v24
        }
        delete(v2, (fi19 | (fi19 | v14)))
        _ = fi19
    }
    v25 := GenS2{F0: 2147483647, F1: (v1 >> ((65535) & 63)), F2: (^(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
