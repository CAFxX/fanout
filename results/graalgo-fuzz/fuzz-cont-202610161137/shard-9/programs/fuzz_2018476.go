package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 string) int16 {
    v5 := []int{-419}
    v5 = append(v5, 1)
    p0 = ""
    var v6 uint16 = 255
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1() int64 {
    var v7 float64 = (0.1 / 1.5)
    v7 += v7
    _ = v7
    return 1
}

func main() {
    var v1 int32 = 1
    var v2 int8 = 0
    v3 := make(map[int]bool)
    var v4 float32 = 1e10
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 = 0
        _ = fi8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 71}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v1 = 1
            v13 := GenS1{F0: 17}
            if true {
                v1 += (v1 - v1)
            } else {
                v2 = (v2 % ((v2) | 1))
                if true {
                    var v14 int8 = v2
                    _ = v14
                } else {
                    var v15 int64 = 1
                    _ = v15
                }
                if false {
                    v16 := v3[fi9]
                    _ = v16
                } else {
                    var v17 int = 127
                    _ = v17
                }
                fmt.Println(v4, fi12)
            }
            _ = fi12
            _ = v13
        }
        _ = fi9
    }
    v2 += v2
    v2 += (v2 ^ 16)
    v18 := GenS1{F0: 128}
    fmt.Println(v4, v1, v4)
    v19 := []int{-1, 42, 348}
    for fi20 := 0; fi20 < 5; fi20++ {
        switch "" {
        case "foo bar":
            delete(v3, fi20)
        default:
            for fi21 := 0; fi21 < 17; fi21++ {
                if true {
                    v22 := GenS1{F0: 42}
                    v23 := GenS2{F0: "", F1: 873}
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v1)
                    v24 := GenS2{F0: "", F1: 255}
                    switch "\n\t" {
                    case "hello":
                        v1 += 1
                    }
                    _ = v24
                }
                _ = fi21
            }
        }
        var v25 float64 = (-2.25 + 1e10)
        v3[21] = (!false)
        _ = fi20
        _ = v25
    }
    v26 := GenS3{F0: -1}
    fmt.Println(v2, v4, v4)
    v27 := GenS1{F0: 255}
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
