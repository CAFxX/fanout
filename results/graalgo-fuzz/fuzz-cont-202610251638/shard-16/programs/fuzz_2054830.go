package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
func genFn0(p0 bool) int16 {
    if p0 {
        v7 := GenS1{F0: 42, F1: 912}
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint8 = 255
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        if p0 {
            for fi10 := 0; fi10 < 6; fi10++ {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if true {
                        p0 = p0
                        v12 := []int{1, fi10, fi11}
                        _ = v12
                    } else {
                        switch (fi11 & (fi10 ^ fi11)) {
                        case 128:
                            if p0 {
                                fmt.Println(p0, fi11, fi10)
                            }
                        default:
                            fmt.Println(fi10)
                        }
                        var v13 int64 = 0
                        _ = v13
                    }
                    _ = fi11
                }
                _ = fi10
            }
        } else {
            p0 = true
        }
        v14 := []int{976}
        for fi15 := 0; fi15 < 14; fi15++ {
            p0 = false
            _ = fi15
        }
        _ = v14
    }
    v16 := []int{}
    _ = p0
    _ = v16
    return 42
}
func genFn1() int16 {
    var v17 int16 = 32767
    var v18 uint64 = 42
    v19 := []int{-1, 20, 2147483647, -617}
    _ = v17
    _ = v18
    _ = v19
    return genFn0(false)
}

func main() {
    var v1 uint = 331
    v2 := make(map[int]int)
    v3 := []string{"\n\t"}
    var v4 uint = v1
    v5 := []string{}
    v6 := make(map[string]bool)
    for fi20 := 0; fi20 < 6; fi20++ {
        switch fi20 {
        case 32768:
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(fi20, v1, v4)
                _ = fi21
            }
        case 128:
            v1 = v4
        default:
            var v22 uint32 = 1
            _ = v22
        }
        v23 := v2[(fi20 - fi20)]
        _ = fi20
        _ = v23
    }
    switch "\n\t" {
    case "a":
        v24 := v6[""]
        _ = v24
    case "":
        v25 := GenS1{F0: 0, F1: 913}
        _ = v25
    case "\n\t":
        var v26 int64 = -9223372036854775808
        _ = v26
    }
    v1 = v1
    v4 -= (v1 * v1)
    var v27 int16 = -979
    v28 := GenS1{F0: 9223372036854775807, F1: 2147483647}
    _ = len(v5)
    var v29 int64 = -977
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
