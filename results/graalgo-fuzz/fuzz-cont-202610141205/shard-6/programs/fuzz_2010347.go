package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int; F2 bool }

func main() {
    var v1 int8 = -1
    var v2 float64 = (3.14159 * (((1.5 + 0.1) - float64(3.14159)) + -2.25))
    v3 := make(map[string]uint32)
    var v4 int32 = 42
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v1)
        _ = fi6
    }
    v7 := GenS1{F0: 145, F1: 2}
    switch "hello" {
    case "hello":
        var v8 int = 0
        _ = v8
    case "a":
        var v9 int = 32767
        _ = v9
    }
    var v10 bool = (-1 >= (v1 ^ (v1 - -1)))
    v11 := GenS2{F0: v4}
    var v12 bool = (v4 > (v4 - v4))
    if false {
        if ((^(v5)) != v5) {
            for fi13 := 0; fi13 < 19; fi13++ {
                v1 = (v1 % (((v1 >> ((42) & 63))) | 1))
                if (32767 == 2) {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v10 = (true && v10)
                        v15 := GenS1{F0: 255, F1: 169}
                        switch fi14 {
                        case -159:
                            var v16 int16 = -486
                            _ = v16
                        }
                        var v17 float64 = v2
                        _ = fi14
                        _ = v15
                        _ = v17
                    }
                    var v18 int32 = v4
                    v19 := v3["a"]
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        switch "" {
                        case "foo bar":
                            v21 := GenS3{F0: v5}
                            _ = v21
                        }
                        _ = fi20
                    }
                }
                var v22 int32 = (v4 * v4)
                _ = fi13
                _ = v22
            }
            v23 := []int{}
            for fi24 := 0; fi24 < 4; fi24++ {
                _ = len(v3)
                v25 := GenS1{F0: 255, F1: 2}
                if v10 {
                    fmt.Println(v4, v2)
                }
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            delete(v3, ("" + "foo bar"))
            for fi26 := 0; fi26 < 8; fi26++ {
                v27 := GenS4{F0: (fi26 + 32767), F1: fi26, F2: v10}
                _ = fi26
                _ = v27
            }
            v4 = (v4 & v4)
        }
        v28 := []int{-2147483648, -335, 533, 9223372036854775807}
        _ = v28
    } else {
        v29 := v3["a"]
        if true {
            var v30 int16 = 32767
            _ = v30
        }
        switch "a" {
        case "\n\t":
            for fi31 := 0; fi31 < 17; fi31++ {
                v29 = (^(v29))
                _ = fi31
            }
        case "hello":
            v4 = (v4 / ((0) | 1))
        }
        _ = v29
    }
    switch "a" {
    case "":
        v1 += v1
    case "hello":
        v1 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
