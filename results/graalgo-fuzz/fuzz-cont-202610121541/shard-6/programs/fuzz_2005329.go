package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 int = -1
    _ = len(v1)
    v4 := GenS1{F0: 100}
    if v2 {
        var v5 float32 = ((0.1 / (3.14159 / 3.14159)) - 1e10)
        v6 := GenS1{F0: -20}
        v3 = v3
        _ = v5
        _ = v6
    }
    var v7 uint16 = 553
    switch (v3 - (v3 << ((0) & 63))) {
    case 491:
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                if (553 > 2) {
                    var v10 int8 = -65
                    _ = v10
                }
                _ = fi9
            }
            _ = fi8
        }
    case -1:
        fmt.Println(v3)
    case 2147483647:
        v3 -= v3
    default:
        var v11 int = -2147483648
        _ = v11
    }
    v3 += (v3 | 9223372036854775807)
    var v12 uint = 165
    v2 = (0 < 2)
    delete(v1, "\n\t")
    var v13 int8 = 99
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 float64 = float64(((3.14159 / 0.1) * 1e10))
            fmt.Println(fi15)
            v17 := GenS1{F0: 0}
            fmt.Println(fi15)
            v3 += 1
            _ = fi15
            _ = v16
            _ = v17
        }
        switch "foo bar" {
        case "hello":
            v18 := GenS2{F0: 642}
            _ = v18
        default:
            fmt.Println(v3)
        }
        _ = fi14
    }
    v3 = (v3 % ((int(v12)) | 1))
    v2 = v2
    fmt.Println(v2, v13)
    switch (v3 + v3) {
    case 1:
        v19 := []int{(v3 & v3), 128, v3, 32768}
        _ = v19
    }
    var v20 int = v3
    fmt.Println(v3)
    var v21 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
