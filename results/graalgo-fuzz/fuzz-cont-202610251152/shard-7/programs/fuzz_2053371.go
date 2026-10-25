package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 string }

func main() {
    var v1 uint8 = 2
    var v2 int64 = 9223372036854775807
    var v3 uint = 0
    v4 := make(map[uint32]string)
    var v5 uint8 = 144
    if false {
        var v6 int32 = -1
        v7 := GenS1{F0: 18446744073709551615}
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                if ((fi9 % ((fi8) | 1)) >= 42) {
                    v10 := GenS1{F0: 65535}
                    _ = len(v4)
                    var v11 uint8 = (v1 | v1)
                    _ = v10
                    _ = v11
                }
                if false {
                    fmt.Println(fi9, v5, v3)
                } else {
                    fmt.Println(fi9, v5)
                    fmt.Println(v1, fi8)
                }
                v12 := GenS2{F0: -521, F1: (fi9 * (fi9 & fi9))}
                _ = fi9
                _ = v12
            }
            _ = fi8
        }
        v2 = (v2 & v2)
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := []int{(-(fi13))}
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    if true {
        v15 := GenS3{F0: 2}
        v16 := []int{127, 0, 128, 42}
        var v17 int = -9223372036854775808
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := []int{}
    var v19 float32 = 0.0
    fmt.Println(v19)
    delete(v4, 756)
    delete(v4, 42)
    v20 := GenS4{F0: -837, F1: "a"}
    v3 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
