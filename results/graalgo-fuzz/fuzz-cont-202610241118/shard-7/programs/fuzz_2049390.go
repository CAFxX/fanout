package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 int64 = 42
    var v2 string = "a"
    v3 := make(map[uint32]int)
    if (128 <= 171) {
        var v4 float32 = -2.25
        var v5 uint = 255
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: 187, F1: 2147483647}
        var v7 uint = 155
        _ = v6
        _ = v7
    }
    fmt.Println(v1, v2)
    v1 = (v1 & -579)
    switch v2 {
    case "hello":
        if false {
            var v8 int16 = int16(v1)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint8 = 0
                _ = fi9
                _ = v10
            }
        }
    }
    v11 := GenS1{F0: 127, F1: -158}
    v1 = (v1 ^ -9223372036854775808)
    switch 128 {
    case -431:
        switch int(v1) {
        case 32767:
            v1 -= v1
        case 0:
            if (false || (-98 >= 78)) {
                v12 := GenS1{F0: 127, F1: 32768}
                _ = v12
            }
        case -463:
            _ = len(v3)
        }
    case -746:
        var v13 int64 = -9223372036854775808
        _ = v13
    case -9223372036854775808:
        v2 = "\n\t"
    }
    fmt.Println(v2, v1)
    fmt.Println(v1)
    v14 := []int{2147483647}
    var v15 float64 = float64(-2.25)
    if true {
        v16 := GenS1{F0: 17, F1: 2147483647}
        fmt.Println(v1, v1, v2)
        _ = len(v14)
        _ = v16
    } else {
        v15 += 0.0
        fmt.Println(v15)
    }
    v17 := GenS1{F0: 1, F1: 97}
    v1 = v1
    var v18 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
