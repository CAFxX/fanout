package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint32 = 42
    v3 := []uint32{}
    var v4 uint16 = 611
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    var v7 float32 = 0.1
    var v8 float32 = v7
    _ = len(v3)
    v9 := GenS1{F0: 0}
    fmt.Println(v8, v1, v8)
    v10 := GenS1{F0: 0}
    switch v1 {
    case 570:
        var v11 string = "a"
        _ = v11
    }
    if false {
        v2 = v2
    } else {
        var v12 int = 32768
        v12 = (v1 | v12)
        v1 = v1
        _ = v12
    }
    v13 := GenS2{F0: ("\n\t" + ("a" + ("a" + "foo bar")))}
    fmt.Println(v1, v8, v2)
    fmt.Println(v8, v4, v1)
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            v2 = (v2 << ((uint(v4)) & 63))
            _ = fi15
        }
        _ = fi14
    }
    var v16 float32 = 1e10
    v17 := v5["foo bar"]
    switch (-(v1)) {
    case 127:
        switch (("foo bar" + "a") + ("a" + "")) {
        case "foo bar":
            v16 -= 0.0
        case "\n\t":
            v2 -= 370
        case "hello":
            if (!(!(1 <= 42))) {
                var v18 uint32 = 1
                v19 := GenS2{F0: ("foo bar" + ("foo bar" + ("foo bar" + "hello")))}
                _ = v18
                _ = v19
            }
        default:
            v20 := GenS3{F0: 42}
            _ = v20
        }
    case 2147483647:
        var v21 float64 = float64(((((-2.25 - v16) / float32(1.5)) * v7) / (0.0 / ((((v8 + ((((v8 * (v7 / 3.14159)) / (v16 / v7)) / v7) - v16)) * v7) * (((v16 + (v16 / (float32(3.14159) * v7))) * 0.1) * (float32(float64((v16 + v8))) / v16))) * 3.14159))))
        _ = v21
    case -582:
        _ = len(v5)
    default:
        v22 := GenS4{F0: 255}
        _ = v22
    }
    var v23 float32 = (1.5 - v8)
    var v24 bool = true
    v25 := GenS5{F0: 65535, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
