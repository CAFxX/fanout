package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint64 }

func main() {
    var v1 string = ""
    var v2 uint = 4294967295
    v3 := []uint32{0, 942}
    fmt.Println(v2, v1)
    fmt.Println(v2)
    switch 32767 {
    case 0:
        var v4 int = 42
        _ = v4
    case 32768:
        if (((float64(float32(3.14159)) * (1.5 * 1.5)) + (float64(0.1) + 3.14159)) <= 0.1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v5 := GenS1{F0: (v2 | (v2 + (v2 | 2))), F1: 42, F2: v1}
    fmt.Println(v1, v2)
    v2 -= v2
    var v6 int32 = -1
    fmt.Println(v6, v6)
    if (329 > 65535) {
        v7 := GenS2{F0: 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int64 = 1
            v6 = (v6 | (v6 / ((0) | 1)))
            _ = fi9
            _ = v10
        }
        switch v1 {
        case "\n\t":
            v3 = append(v3, 0)
        case "a":
            fmt.Println(v6, v1)
        }
        v11 := GenS1{F0: v2, F1: 0, F2: v1}
        _ = fi8
        _ = v11
    }
    var v12 bool = (v1 < v1)
    if (v1 <= v1) {
        if true {
            v12 = v12
        }
    }
    v6 -= v6
    if true {
        if true {
            fmt.Println(v1, v12)
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                v6 -= 42
                _ = fi13
            }
            var v14 uint64 = 1
            v14 += v14
            _ = v14
        }
        v15 := GenS3{F0: "\n\t", F1: 731}
        var v16 int = 0
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
