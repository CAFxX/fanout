package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 uint }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64) string {
    var v4 int = -656
    var v5 uint8 = 127
    _ = p0
    _ = v4
    _ = v5
    return ""
}
func genFn1(p0 uint, p1 bool, p2 int32) int32 {
    v6 := GenS1{F0: 71, F1: (p0 - 784)}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p2 * (^(p2)))
}

func main() {
    v1 := []int{-185, 42, 32767, 612}
    var v2 uint16 = 32767
    var v3 uint16 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 float64 = float64((1.5 - float32(-2.25)))
    v9 := GenS2{F0: 18446744073709551615}
    var v10 uint8 = 1
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS3{F0: 1, F1: (fi11 + fi11), F2: 42}
            if (!true) {
                fmt.Println(v10, v2)
            } else {
                var v13 uint64 = 2
                v10 -= v10
                _ = v13
            }
            fmt.Println(v10, v2, fi11)
            _ = fi11
            _ = v12
        }
        v3 -= 32767
        fmt.Println(v2, v2, v2)
    }
    v10 += (^(v10))
    var v14 int8 = 42
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS4{F0: 703, F1: 9223372036854775807}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    var v18 uint16 = (v2 / ((v3) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    if false {
        v1 = append(v1, -9223372036854775808)
    } else {
        var v19 float64 = 0.1
        if true {
            fmt.Println(v14, v3)
        } else {
            var v20 uint16 = v2
            _ = v20
        }
        _ = v19
    }
    var v21 int16 = -1
    fmt.Println(v14, v10, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
