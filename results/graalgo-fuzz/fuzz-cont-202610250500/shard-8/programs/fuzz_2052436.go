package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{-9223372036854775808, 470, -9223372036854775808, 797}
    var v2 uint8 = 2
    var v3 int8 = -95
    var v4 int = -346
    v3 = (v3 >> ((339) & 63))
    if false {
        switch (v4 - int(v3)) {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v4 = v4
        }
        var v5 int64 = 0
        _ = v5
    }
    v3 = -56
    if (!false) {
        var v6 int64 = 0
        v2 += v2
        v7 := GenS1{F0: ("" <= "\n\t"), F1: v4, F2: (v3 * (v3 - (-(v3))))}
        _ = v6
        _ = v7
    } else {
        var v8 bool = true
        var v9 bool = ("hello" <= "a")
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v3, fi10)
        v3 -= v3
        var v11 float64 = 1.5
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: true}
    v2 += 42
    if true {
        for fi13 := 0; fi13 < 7; fi13++ {
            v1 = append(v1, v4)
            var v14 int8 = 0
            switch 75 {
            case 32768:
                var v15 uint32 = 266
                _ = v15
            case 2147483647:
                if ((-2.25 + 3.14159) <= 1.5) {
                    var v16 uint32 = 217
                    v14 += (-(v3))
                    var v17 int8 = v3
                    _ = v16
                    _ = v17
                }
            default:
                var v18 float64 = (-2.25 - 0.0)
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        v4 = v4
    } else {
        v2 -= v2
        v1 = append(v1, v4)
    }
    var v19 int64 = 9223372036854775807
    var v20 int16 = 1
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v4, v3)
        _ = fi21
    }
    var v22 int64 = int64(v20)
    for fi23 := 0; fi23 < 10; fi23++ {
        _ = len(v1)
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
