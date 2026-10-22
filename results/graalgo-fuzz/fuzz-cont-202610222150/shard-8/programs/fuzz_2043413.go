package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    v3 := make(map[int]int)
    v2 = (v2 + v2)
    v4 := []int{128, 471, 32768, 950}
    v5 := GenS1{F0: -2147483648, F1: 42}
    if (42 == -1) {
        for fi6 := 0; fi6 < 16; fi6++ {
            if (!true) {
                v2 = v2
            } else {
                var v7 bool = false
                switch fi6 {
                case 807:
                    v7 = (!(9223372036854775807 >= 9223372036854775807))
                }
                _ = v7
            }
            v2 = ("" + v2)
            _ = fi6
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                switch fi8 {
                case -708:
                    var v10 int = fi8
                    _ = v10
                default:
                    v11 := GenS1{F0: 128, F1: 43}
                    _ = v11
                }
                fmt.Println(fi9)
                if ((0.1 / 3.14159) < 3.14159) {
                    var v12 uint32 = 2147483647
                    _ = v12
                } else {
                    var v13 uint16 = 359
                    var v14 int32 = 252
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v13
                    _ = v14
                }
                switch fi8 {
                case 128:
                    var v15 uint16 = 42
                    _ = len(v3)
                    _ = v15
                case 1:
                    switch v2 {
                    case "a":
                        if (0 >= 1000) {
                            v2 = "\n\t"
                            v16 := GenS1{F0: 0, F1: 42}
                            var v17 bool = false
                            if v17 {
                                var v18 float32 = -2.25
                                var v19 uint = 556
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        v20 := GenS1{F0: (fi9 ^ fi8), F1: 2147483647}
                        v21 := GenS2{F0: true}
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi9
            }
            delete(v1, -750)
            for fi22 := 0; fi22 < 20; fi22++ {
                v2 = ("hello" + v2)
                v1[2147483647] = 2
                var v23 int32 = -390
                _ = fi22
                _ = v23
            }
            _ = fi8
        }
        if (346 > 0) {
            switch v2 {
            case "":
                fmt.Println(v2)
            default:
                var v24 int32 = 591
                _ = v24
            }
            if true {
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
                var v25 int8 = -128
                _ = v25
            } else {
                _ = len(v4)
            }
        } else {
            var v26 int = 32767
            v3[v26] = (v26 / ((v26) | 1))
            _ = len(v4)
            _ = v26
        }
    } else {
        v4 = append(v4, 32767)
    }
    var v27 uint32 = 2
    delete(v1, -1)
    for fi28 := 0; fi28 < 9; fi28++ {
        delete(v3, (^(fi28)))
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 string = v2
            v31 := GenS1{F0: (^(fi29)), F1: 2}
            v2 = "hello"
            _ = fi29
            _ = v30
            _ = v31
        }
        v27 = v27
        _ = fi28
    }
    var v32 float32 = 1e10
    v2 = v2
    for fi33 := 0; fi33 < 12; fi33++ {
        if len(v4) > 0 {
            v4[0] = fi33
        }
        _ = fi33
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2)
        v32 += (((float32(float64(((v32 - (v32 / ((v32 - (v32 - v32)) * v32))) * 0.1))) - v32) - 0.1) / v32)
    } else {
        fmt.Println(v27, v2)
    }
    v2 = v2
    fmt.Println(v32)
    var v34 string = ("" + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
