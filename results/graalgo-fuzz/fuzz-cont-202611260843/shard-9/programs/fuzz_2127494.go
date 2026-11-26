package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 int, p1 uint64) int16 {
    fmt.Println(p1, p0)
    if true {
        var v4 int8 = 19
        _ = v4
    } else {
        var v5 int16 = 851
        v5 = (v5 - v5)
        _ = v5
    }
    p1 = p1
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 int16 = 0
    v2 := []int{0, -1, -551, -386}
    var v3 uint = 42
    v6 := GenS1{F0: -335, F1: true}
    var v7 string = "hello"
    v2 = append(v2, 32767)
    fmt.Println(v1, v7)
    for fi8 := 0; fi8 < 18; fi8++ {
        if (!false) {
            for fi9 := 0; fi9 < 6; fi9++ {
                v1 = v1
                _ = fi9
            }
        } else {
            v10 := GenS2{F0: true, F1: 1}
            if true {
                var v11 float64 = 1e10
                v12 := GenS1{F0: 42, F1: (true || false)}
                v13 := GenS3{F0: 37, F1: (v1 % ((1000) | 1))}
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 uint16 = 988
            _ = v10
            _ = v14
        }
        _ = fi8
    }
    v15 := GenS3{F0: 0, F1: v1}
    v16 := GenS2{F0: true, F1: 136}
    var v17 uint = (v3 + (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
