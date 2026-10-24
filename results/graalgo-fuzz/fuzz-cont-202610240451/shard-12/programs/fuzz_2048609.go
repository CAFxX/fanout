package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int; F2 uint }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]string)
    var v3 string = ""
    var v4 float32 = 3.14159
    v1 -= float64((float32(-2.25) / -2.25))
    v5 := GenS1{F0: 9223372036854775807}
    if (42 >= 4294967295) {
        var v6 int16 = 1
        for fi7 := 0; fi7 < 17; fi7++ {
            delete(v2, 2)
            _ = fi7
        }
        v4 = -2.25
        _ = v6
    }
    fmt.Println(v1, v3, v1)
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 9223372036854775807}
        _ = fi8
        _ = v9
    }
    v10 := []int{-1, 127, -2147483648}
    var v11 uint64 = 0
    var v12 float32 = (v4 * (v4 + v4))
    v3 = ("foo bar" + v3)
    var v13 uint16 = 261
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS2{F0: false, F1: (fi14 - (fi14 + (fi14 | fi14))), F2: 4294967295}
        var v16 string = ""
        _ = fi14
        _ = v15
        _ = v16
    }
    switch (v3 + v3) {
    case "hello":
        var v17 int32 = 1
        _ = v17
    case "":
        var v18 float64 = ((v1 - (-2.25 / 3.14159)) * v1)
        _ = v18
    default:
        v19 := GenS1{F0: 0}
        _ = v19
    }
    v20 := v2[804]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
