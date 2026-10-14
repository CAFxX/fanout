package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    var v1 uint32 = 660
    var v2 int16 = 32767
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: -515, F1: (v2 % ((-1) | 1))}
    v8 := GenS2{F0: 0}
    var v9 float32 = (0.0 - (float32(float64(0.0)) - (1.5 + ((0.1 - 0.0) / 1e10))))
    switch "" {
    case "a":
        var v10 int8 = 42
        _ = v10
    case "":
        v11 := GenS3{F0: -1, F1: -2147483648}
        _ = v11
    case "hello":
        var v12 float64 = (0.1 * float64(float32(1.5)))
        _ = v12
    }
    if ((128 == 255) || ((v2 | v2) <= 1)) {
        var v13 uint64 = 42
        switch 1 {
        case 1:
            if true {
                var v14 uint8 = 0
                _ = v14
            }
        case 32768:
            for fi15 := 0; fi15 < 9; fi15++ {
                fmt.Println(fi15, v1)
                _ = fi15
            }
        }
        _ = v13
    } else {
        v16 := v5[((("hello" + ("foo bar" + "\n\t")) + (("foo bar" + "foo bar") + "\n\t")) + "")]
        _ = len(v4)
        _ = v16
    }
    var v17 float64 = 0.0
    var v18 uint8 = 24
    var v19 int16 = (v2 >> ((uint(v1)) & 63))
    var v20 int32 = 737
    fmt.Println(v17, v19, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
