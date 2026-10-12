package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool) uint32 {
    if true {
        fmt.Println(p0)
        var v6 int32 = 0
        _ = v6
    }
    p0 = p0
    v7 := []int{9223372036854775807}
    var v8 float32 = (0.1 - (0.1 / (float32(1e10) + 0.0)))
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1() uint {
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 int = 42
                v11 -= (-(fi10))
                var v12 int = (v11 ^ v11)
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 42, F1: 0, F2: 253}
            _ = fi9
            _ = v13
        }
    } else {
        var v14 float64 = (1.5 + (3.14159 - 0.0))
        v15 := GenS2{F0: 0}
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: -32768, F1: 207, F2: 1}
    if (127 >= 12) {
        var v17 float32 = (1.5 / 1e10)
        _ = v17
    }
    _ = v16
    return 1
}

func main() {
    var v1 int16 = -925
    var v2 uint32 = 2147483647
    var v3 int64 = 1
    var v4 int8 = -124
    var v5 float32 = 1.5
    var v18 uint32 = v2
    fmt.Println(v4, v18)
    for fi19 := 0; fi19 < 13; fi19++ {
        v2 = 4294967295
        var v20 int8 = (v4 % ((100) | 1))
        _ = fi19
        _ = v20
    }
    v5 = 1e10
    v21 := []int{}
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS3{F0: 128, F1: v1, F2: genFn0(true)}
        if true {
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v5)
                v25 := GenS4{F0: (fi22 / ((fi24) | 1)), F1: (("" + "") + "foo bar")}
                if (((1e10 + v5) >= 1e10) && false) {
                    var v26 int64 = (v3 | 0)
                    _ = v26
                }
                switch "hello" {
                case "a":
                    v21 = append(v21, fi24)
                case "foo bar":
                    v18 -= (^(v2))
                case "hello":
                    var v27 int8 = (^(v4))
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        } else {
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 string = ""
                _ = fi28
                _ = v29
            }
            var v30 float32 = float32(3.14159)
            for fi31 := 0; fi31 < 13; fi31++ {
                if true {
                    v30 -= 0.0
                    v3 -= (-(v3))
                }
                fmt.Println(v3)
                _ = fi31
            }
            _ = v30
        }
        _ = fi22
        _ = v23
    }
    v18 += 2147483648
    v4 = -106
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := GenS5{F0: 1, F1: 794}
        _ = fi32
        _ = v33
    }
    var v34 uint64 = 34
    fmt.Println(v5, v18, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
