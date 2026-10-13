package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 255
    var v2 float32 = ((-2.25 + float32(1.5)) * 3.14159)
    var v3 int16 = 420
    v4 := make(map[int]uint32)
    var v5 float64 = 1.5
    switch "" {
    case "\n\t":
        if true {
            v3 += -1
            v5 += float64((v2 / (((0.1 / v2) - -2.25) + 0.0)))
            var v6 uint8 = 128
            _ = v6
        } else {
            var v7 uint32 = 732
            var v8 uint32 = (v7 ^ v7)
            v8 -= (^(v7))
            _ = v7
            _ = v8
        }
    case "hello":
        v4[-775] = 2147483648
    case "foo bar":
        for fi9 := 0; fi9 < 12; fi9++ {
            switch fi9 {
            case 240:
                for fi10 := 0; fi10 < 20; fi10++ {
                    if (0 > 0) {
                        var v11 float32 = v2
                        v12 := GenS1{F0: "hello", F1: 2}
                        v13 := []int{}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v2 = -2.25
                        fmt.Println(fi9, v3, v3)
                    }
                    switch (("a" + "hello") + ("hello" + ("hello" + "hello"))) {
                    case "a":
                        var v14 int16 = (v3 - (v3 + 42))
                        _ = v14
                    case "\n\t":
                        _ = len(v4)
                    case "hello":
                        switch fi10 {
                        case -382:
                            var v15 uint64 = 42
                            _ = v15
                        case 747:
                            v4[(fi10 * (fi9 + (fi10 / ((fi10) | 1))))] = 455
                        default:
                            v16 := GenS1{F0: "", F1: 4294967295}
                            _ = v16
                        }
                    }
                    var v17 int8 = 1
                    _ = fi10
                    _ = v17
                }
            case -890:
                fmt.Println(v5, fi9, v5)
            case -9223372036854775808:
                var v18 uint = 42
                _ = v18
            }
            v3 = (v3 - (-(v3)))
            var v19 float32 = v2
            _ = fi9
            _ = v19
        }
    default:
        var v20 string = "a"
        _ = v20
    }
    v21 := GenS2{F0: (!false), F1: 108}
    v3 = v3
    var v22 string = (("foo bar" + ("hello" + "a")) + "\n\t")
    v4[-358] = 840
    v23 := []int{-951, 0, 869, -1}
    var v24 string = ""
    v25 := GenS1{F0: v22, F1: 1}
    fmt.Println(v3, v24)
    delete(v4, 367)
    v3 = (v3 * v3)
    var v26 uint32 = 1
    v24 = "a"
    v22 = "\n\t"
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := v4[(fi27 & (fi27 & fi27))]
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
