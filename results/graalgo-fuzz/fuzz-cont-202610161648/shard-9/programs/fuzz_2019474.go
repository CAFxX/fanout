package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }

func main() {
    v1 := []int{126, 333, 750, 32767}
    var v2 uint = 255
    var v3 float32 = 3.14159
    var v4 int64 = 9223372036854775807
    var v5 string = ""
    var v6 int64 = (-(v4))
    var v7 int32 = 961
    var v8 int16 = 920
    var v9 int32 = (v7 >> ((v2) & 63))
    if ((65535 == 1) || true) {
        if true {
            v10 := GenS1{F0: (0 > 2), F1: v6}
            v8 = v8
            v11 := GenS1{F0: true, F1: (v4 / (((v4 + v6)) | 1))}
            _ = v10
            _ = v11
        }
        v8 -= v8
        fmt.Println(v2)
    } else {
        var v12 uint8 = 42
        v6 += (v6 ^ (v4 / ((v4) | 1)))
        _ = v12
    }
    fmt.Println(v9, v9, v7)
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 += ((v3 + -2.25) - 0.0)
        switch v5 {
        case "":
            v4 -= (v4 ^ 1)
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            v5 = v5
            _ = fi14
        }
        _ = fi13
    }
    _ = len(v1)
    for fi15 := 0; fi15 < 15; fi15++ {
        v1 = append(v1, (fi15 << ((v2) & 63)))
        _ = fi15
    }
    v6 = v6
    var v16 float32 = (v3 - float32((float64((((v3 - (1.5 + (v3 / v3))) + 0.0) + -2.25)) + 0.0)))
    switch "" {
    case "hello":
        v7 = v9
    case "\n\t":
        var v17 float64 = (-2.25 * 1e10)
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
