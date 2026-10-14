package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 float64 = (3.14159 / 0.1)
    var v2 uint32 = 845
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    if (-128 >= 51) {
        v5 := []int{606, 797, 2147483647}
        var v6 float64 = float64(1.5)
        _ = v5
        _ = v6
    } else {
        if false {
            fmt.Println(v2)
            v1 += v1
        }
    }
    v7 := GenS1{F0: 774, F1: 42}
    v3[4294967295] = -392
    fmt.Println(v2)
    var v8 uint16 = 978
    switch ("" + "hello") {
    case "hello":
        v9 := GenS2{F0: 128, F1: false}
        _ = v9
    case "a":
        v10 := GenS2{F0: 1, F1: (42 != -2147483648)}
        _ = v10
    case "\n\t":
        v11 := GenS2{F0: 1, F1: (0 >= v8)}
        _ = v11
    }
    switch "hello" {
    case "foo bar":
        var v12 int32 = 1
        _ = v12
    case "":
        v13 := []int{-8, 685, -603}
        _ = v13
    default:
        var v14 int8 = 100
        _ = v14
    }
    v15 := []int{-855}
    switch 32768 {
    case 127:
        v2 = 0
    default:
        v1 += float64(-2.25)
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        delete(v4, ("" + ((("hello" + "\n\t") + ("" + "\n\t")) + ("" + "hello"))))
        _ = fi16
    }
    v8 += 1
    var v17 uint = 987
    fmt.Println(v1, v8, v17)
    if (!(!(1 > -64))) {
        switch "" {
        case "hello":
            var v18 uint8 = 1
            _ = v18
        }
        var v19 int64 = -1
        _ = v19
    }
    switch 32768 {
    case -962:
        if (!true) {
            var v20 int8 = -128
            var v21 uint64 = 1
            _ = v20
            _ = v21
        } else {
            v17 -= v17
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS1{F0: int16(v17), F1: 0}
                if (!(true || (2147483647 > 42))) {
                    var v24 uint8 = 127
                    var v25 float64 = ((-2.25 * (0.0 / v1)) * v1)
                    _ = v24
                    _ = v25
                }
                v15 = append(v15, int(v8))
                _ = fi22
                _ = v23
            }
            fmt.Println(v2)
        }
    }
    var v26 float32 = 0.1
    var v27 uint32 = (v2 | v2)
    fmt.Println(v27, v17)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
