package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 uint64, p2 int8) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 string = ""
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                p2 += (p2 & p2)
                p1 = p1
                _ = fi6
            }
            if (2147483647 > 2147483647) {
                v7 := []int{}
                p2 = p2
                switch ("" + "\n\t") {
                case "a":
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        if len(v7) > 0 {
                            v7[0] = fi4
                        }
                    }
                default:
                    if false {
                        for fi8 := 0; fi8 < 11; fi8++ {
                            fmt.Println(fi8)
                            if true {
                                var v9 int32 = 2147483647
                                _ = v9
                            } else {
                                var v10 uint8 = 6
                                var v11 uint = 1
                                _ = v10
                                _ = v11
                            }
                            var v12 float64 = ((0.0 - 0.1) + 1e10)
                            _ = fi8
                            _ = v12
                        }
                        v13 := GenS1{F0: (fi4 + (fi4 & 0)), F1: -32768}
                        if true {
                            _ = len(v7)
                            fmt.Println(v5)
                        } else {
                            var v14 int = (fi4 / ((fi4) | 1))
                            _ = v14
                        }
                        _ = v13
                    }
                }
                _ = v7
            } else {
                p2 = int8(fi4)
            }
        }
        _ = fi4
        _ = v5
    }
    var v15 uint64 = 851
    v15 -= v15
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p2 | p2)
}

func main() {
    var v1 bool = (879 < 1)
    var v2 uint64 = 559
    var v3 int16 = 0
    switch "\n\t" {
    case "foo bar":
        switch "" {
        case "\n\t":
            fmt.Println(v2)
        case "hello":
            v16 := []int{-2147483648, -315, 1, 0}
            _ = v16
        }
    case "hello":
        v17 := GenS1{F0: -798, F1: v3}
        _ = v17
    case "\n\t":
        v18 := GenS2{F0: 38, F1: v2, F2: 9223372036854775807}
        _ = v18
    }
    v2 = (v2 ^ 616)
    v19 := []int{-746, -845, -1}
    if v1 {
        fmt.Println(v1)
        v20 := GenS3{F0: 399, F1: 1}
        v2 = 427
        _ = v20
    } else {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    }
    if (v1 || v1) {
        if len(v19) > 0 {
            v19[0] = 269
        }
        switch ("hello" + "a") {
        case "a":
            var v21 float64 = 0.1
            _ = v21
        default:
            fmt.Println(v1, v2, v3)
        }
    }
    var v22 uint64 = 0
    var v23 uint32 = 840
    v19 = append(v19, 127)
    v24 := GenS4{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
